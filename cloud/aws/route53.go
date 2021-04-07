package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/route53domains"
)

func MigrateDomain(domain string, sourceAccountID string, sourceSecretKey string, targetAccountID string, targetSecretKey string) (string, string, error) {
	mySession := session.Must(session.NewSession())

	sourceClient := route53domains.New(mySession, aws.NewConfig().WithCredentials(credentials.NewStaticCredentials(sourceAccountID, sourceSecretKey, "")))
	transferParams := &route53domains.TransferDomainToAnotherAwsAccountInput{
		AccountId:  &targetAccountID,
		DomainName: &domain,
	}

	req, transferOutput := sourceClient.TransferDomainToAnotherAwsAccountRequest(transferParams)

	err := req.Send()
	if err != nil { // resp is now filled
		return "", "", err
	}

	targetClient := route53domains.New(mySession, aws.NewConfig().WithCredentials(credentials.NewStaticCredentials(targetAccountID, targetSecretKey, "")))
	acceptParams := &route53domains.AcceptDomainTransferFromAnotherAwsAccountInput{
		DomainName: &domain,
		Password:   transferOutput.Password,
	}

	acceptOutput, err := targetClient.AcceptDomainTransferFromAnotherAwsAccount(acceptParams)

	return *transferOutput.OperationId, *acceptOutput.OperationId, nil
}

func GetMigrateOperation(operationID string, accountID string, secretKey string) (string, error) {
	mySession := session.Must(session.NewSession())

	client := route53domains.New(mySession, aws.NewConfig().WithCredentials(credentials.NewStaticCredentials(accountID, secretKey, "")))
	operationDetailParams := &route53domains.GetOperationDetailInput{
		OperationId: &operationID,
	}

	operationDetailOutput, err := client.GetOperationDetail(operationDetailParams)
	if err != nil {
		return "", err
	}

	return *operationDetailOutput.Status, nil
}
