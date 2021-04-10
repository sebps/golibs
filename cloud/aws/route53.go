package aws

import (
	"errors"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/route53"
	"github.com/aws/aws-sdk-go/service/route53domains"
	"strconv"
	"time"
)

func MigrateDomain(domain string, targetAccountID string, sourceAccessKeyID string, sourceSecretKey string, targetAccessKeyID string, targetSecretKey string) (string, error) {
	_, password, err := TransferDomain(domain, targetAccountID, sourceAccessKeyID, sourceSecretKey)
	if err != nil {
		return "", err
	}

	operationId, err := AcceptDomainTransfer(domain, password, targetAccessKeyID, targetSecretKey)
	if err != nil {
		return "", err
	}

	return operationId, nil
}

func TransferDomain(domain string, targetAccountID string, accessKeyID string, secretKey string) (string, string, error) {
	mySession := session.Must(session.NewSession())
	sourceClient := route53domains.New(mySession, aws.NewConfig().WithRegion("us-east-1").WithCredentials(credentials.NewStaticCredentials(accessKeyID, secretKey, "")))

	transferParams := &route53domains.TransferDomainToAnotherAwsAccountInput{
		AccountId:  &targetAccountID,
		DomainName: &domain,
	}

	req, transferOutput := sourceClient.TransferDomainToAnotherAwsAccountRequest(transferParams)
	err := req.Send()
	if err != nil {
		return "", "", err
	}

	return *transferOutput.OperationId, *transferOutput.Password, nil
}

func AcceptDomainTransfer(domain string, password string, accessKeyID string, secretKey string) (string, error) {
	mySession := session.Must(session.NewSession())
	targetClient := route53domains.New(mySession, aws.NewConfig().WithRegion("us-east-1").WithCredentials(credentials.NewStaticCredentials(accessKeyID, secretKey, "")))

	acceptParams := &route53domains.AcceptDomainTransferFromAnotherAwsAccountInput{
		DomainName: &domain,
		Password:   &password,
	}
	acceptOutput, err := targetClient.AcceptDomainTransferFromAnotherAwsAccount(acceptParams)
	if err != nil {
		return "", err
	}

	return *acceptOutput.OperationId, nil
}

func CancelDomainTransfer(domain string, accessKeyID string, secretKey string) (string, error) {
	mySession := session.Must(session.NewSession())
	targetClient := route53domains.New(mySession, aws.NewConfig().WithRegion("us-east-1").WithCredentials(credentials.NewStaticCredentials(accessKeyID, secretKey, "")))

	cancelParams := &route53domains.CancelDomainTransferToAnotherAwsAccountInput{
		DomainName: &domain,
	}
	cancelOutput, err := targetClient.CancelDomainTransferToAnotherAwsAccount(cancelParams)
	if err != nil {
		return "", err
	}

	return *cancelOutput.OperationId, nil
}

func GetOperation(operationID string, accessKeyID string, secretKey string) (string, error) {
	mySession := session.Must(session.NewSession())
	client := route53domains.New(mySession, aws.NewConfig().WithRegion("us-east-1").WithCredentials(credentials.NewStaticCredentials(accessKeyID, secretKey, "")))

	operationDetailParams := &route53domains.GetOperationDetailInput{
		OperationId: &operationID,
	}
	operationDetailOutput, err := client.GetOperationDetail(operationDetailParams)
	if err != nil {
		return "", err
	}

	return *operationDetailOutput.Status, nil
}

func ListOperations(accessKeyID string, secretKey string) ([]*route53domains.OperationSummary, error) {
	mySession := session.Must(session.NewSession())
	client := route53domains.New(mySession, aws.NewConfig().WithRegion("us-east-1").WithCredentials(credentials.NewStaticCredentials(accessKeyID, secretKey, "")))

	listOperationsParams := &route53domains.ListOperationsInput{}
	operationList, err := client.ListOperations(listOperationsParams)
	if err != nil {
		return nil, err
	}

	return operationList.Operations, nil
}

func MigrateHostedZones(domain string, sourceAccessKeyID string, sourceSecretKey string, targetAccessKeyID string, targetSecretKey string) ([]string, error) {
	var changeIds []string
	// get all the records from the original hosted zone
	mySession := session.Must(session.NewSession())
	sourceClient := route53.New(mySession, aws.NewConfig().WithRegion("us-east-1").WithCredentials(credentials.NewStaticCredentials(sourceAccessKeyID, sourceSecretKey, "")))
	listHostedZonesOutput, err := sourceClient.ListHostedZonesByName(&route53.ListHostedZonesByNameInput{
		DNSName: &domain,
	})
	if err != nil {
		return nil, err
	}
	if len(listHostedZonesOutput.HostedZones) == 0 {
		return nil, errors.New("Domain not found")
	}

	// transfer each hosted zone
	for i := 0; i < len(listHostedZonesOutput.HostedZones); i++ {
		changeId, err := TransferHostedZone(*listHostedZonesOutput.HostedZones[i].Name, *listHostedZonesOutput.HostedZones[i].Id, sourceAccessKeyID, sourceSecretKey, targetAccessKeyID, targetSecretKey)
		if err != nil {
			return nil, err
		}

		changeIds = append(changeIds, changeId)
	}

	return changeIds, nil
}

func TransferHostedZone(domain string, hostedZoneID string, sourceAccessKeyID string, sourceSecretKey string, targetAccessKeyID string, targetSecretKey string) (string, error) {
	mySession := session.Must(session.NewSession())
	sourceClient := route53.New(mySession, aws.NewConfig().WithRegion("us-east-1").WithCredentials(credentials.NewStaticCredentials(sourceAccessKeyID, sourceSecretKey, "")))
	targetClient := route53.New(mySession, aws.NewConfig().WithRegion("us-east-1").WithCredentials(credentials.NewStaticCredentials(targetAccessKeyID, targetSecretKey, "")))
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)

	// check if hosted zone existing for domain
	listHostedZonesOutput, err := targetClient.ListHostedZonesByName(&route53.ListHostedZonesByNameInput{
		DNSName: &domain,
	})
	if err != nil {
		return "", err
	}

	var targetHostedZoneId string
	if len(listHostedZonesOutput.HostedZones) == 0 {
		// create a new hosted zone
		createHostedZoneOutput, err := targetClient.CreateHostedZone(&route53.CreateHostedZoneInput{
			CallerReference: &timestamp,
			Name:            &domain,
		})
		if err != nil {
			return "", err
		}

		targetHostedZoneId = *createHostedZoneOutput.HostedZone.Id
	} else {
		targetHostedZoneId = *listHostedZonesOutput.HostedZones[0].Id
	}

	// get all the resource records for the hosted zone
	listResourceRecordOutput, err := sourceClient.ListResourceRecordSets(&route53.ListResourceRecordSetsInput{
		HostedZoneId:    &hostedZoneID,
		StartRecordName: &domain,
	})
	if err != nil {
		return "", err
	}

	// create all the records in the new hosted zone
	var changes []*route53.Change
	var action = route53.ChangeActionCreate

	for i := 0; i < len(listResourceRecordOutput.ResourceRecordSets); i++ {
		changes = append(changes, &route53.Change{
			Action:            &action,
			ResourceRecordSet: listResourceRecordOutput.ResourceRecordSets[i],
		})
	}

	changeResourceRecordSetsOutput, err := targetClient.ChangeResourceRecordSets(&route53.ChangeResourceRecordSetsInput{
		ChangeBatch: &route53.ChangeBatch{
			Changes: changes,
		},
		HostedZoneId: &targetHostedZoneId,
	})
	if len(listHostedZonesOutput.HostedZones) > 0 {
		return "", errors.New("Resource not transferred")
	}

	return *changeResourceRecordSetsOutput.ChangeInfo.Id, nil
}

func ListHostedZones(accessKeyID string, secretKey string) ([]*route53.HostedZone, error) {
	mySession := session.Must(session.NewSession())
	client := route53.New(mySession, aws.NewConfig().WithRegion("us-east-1").WithCredentials(credentials.NewStaticCredentials(accessKeyID, secretKey, "")))

	listHostedZonesParams := &route53.ListHostedZonesInput{}
	zoneList, err := client.ListHostedZones(listHostedZonesParams)
	if err != nil {
		return nil, err
	}

	return zoneList.HostedZones, nil
}

func CreateHostedZone(name string, accessKeyID string, secretKey string) (string, error) {
	mySession := session.Must(session.NewSession())
	client := route53.New(mySession, aws.NewConfig().WithRegion("us-east-1").WithCredentials(credentials.NewStaticCredentials(accessKeyID, secretKey, "")))

	createHostedZoneOutput, err := client.CreateHostedZone(&route53.CreateHostedZoneInput{
		Name: &name,
	})
	if err != nil {
		return "", err
	}

	return *createHostedZoneOutput.HostedZone.Id, nil
}

func DestroyHostedZone(hostedZoneID string, accessKeyID string, secretKey string) (string, error) {
	mySession := session.Must(session.NewSession())
	client := route53.New(mySession, aws.NewConfig().WithRegion("us-east-1").WithCredentials(credentials.NewStaticCredentials(accessKeyID, secretKey, "")))

	deleteHostedZoneOutput, err := client.DeleteHostedZone(&route53.DeleteHostedZoneInput{
		Id: &hostedZoneID,
	})
	if err != nil {
		return "", err
	}

	return *deleteHostedZoneOutput.ChangeInfo.Id, nil
}
