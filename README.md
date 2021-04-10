# Some libs for go

## Redirect 
Basic lib to implicitely redirect an incoming connexion to a target endpoint

#### cli usage
- redirect --sourceHost=127.0.0.1 --sourcePort=3000 --targetHost=com.example.endpoint --targetPort=5000 --protocol=http

#### lib usage
- connectors.HttpConnector(&connectors.Connexion{ SourceHost: "127.0.0.1", SourcePort: 3000, TargetHost: "com.example.endpoint",TargetPort: 5000})

## Eventbus 
Some generic eventbus implementation

#### lib usage
- eb := &eventbus.Eventbus{}
- eb.Subscribe("EventName", subscriber)
- eb.Publish("EventName", publisher)

## Generic 
Some untyped generic methods for datastructures

#### arrays
- func Sort(input interface{}, less func(a interface{}, b interface{}) bool) (interface{}, error)
- func Find(element interface{}, input interface{}) (index, bool, error)

#### maps
- func Keys(input interface{}) (interface{}, error)
- func Values(input interface{}) (interface{}, error)
- func KeysValues(input interface{}) (interface{}, interface{}, error)
- func FindKey(key interface{}, input interface{}) (bool, error)

#### types
- func GeneralizeSlice(input interface{}) ([]interface{}, error)
- func GeneralizeMap(input interface{}) (map[interface{}]interface{}, error)

#### utils
- func Traverse(input interface{}, handler func(i interface{}), handleBeforeTraverse bool, handleAfterTraverse bool) error

## Cloud 
Some methods for cloud operations

#### aws
- func MigrateDomain(domain string, targetAccountID string, sourceAccessKeyID string, sourceSecretKey string, targetAccessKeyID string, targetSecretKey string) (string, error)
- func TransferDomain(domain string, targetAccountID string, accessKeyID string, secretKey string) (string, string, error)
- func AcceptDomainTransfer(domain string, password string, accessKeyID string, secretKey string) (string, error)
- func CancelDomainTransfer(domain string, accessKeyID string, secretKey string) (string, error)
- func GetDomainOperationStatus(operationID string, accessKeyID string, secretKey string) (string, error)
- func ListDomainOperations(accessKeyID string, secretKey string) ([]*route53domains.OperationSummary, error)
- func MigrateHostedZones(domain string, sourceAccessKeyID string, sourceSecretKey string, targetAccessKeyID string, targetSecretKey string) ([]string, error)
- func TransferHostedZone(domain string, hostedZoneID string, sourceAccessKeyID string, sourceSecretKey string, targetAccessKeyID string, targetSecretKey string) (string, error)
- func ListHostedZones(accessKeyID string, secretKey string) ([]*route53.HostedZone, error)
- func CreateHostedZone(name string, accessKeyID string, secretKey string) (string, error)
- func DestroyHostedZone(hostedZoneID string, accessKeyID string, secretKey string) (string, error)