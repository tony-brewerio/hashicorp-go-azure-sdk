
## `github.com/hashicorp/go-azure-sdk/resource-manager/sql/2021-11-01/sqlagent` Documentation

The `sqlagent` SDK allows for interaction with the Azure Resource Manager Service `sql` (API Version `2021-11-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/sql/2021-11-01/sqlagent"
```


### Client Initialization

```go
client := sqlagent.NewSqlAgentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SqlAgentClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := sqlagent.NewSqlManagedInstanceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedInstanceValue")

payload := sqlagent.SqlAgentConfiguration{
	// ...
}


read, err := client.CreateOrUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SqlAgentClient.Get`

```go
ctx := context.TODO()
id := sqlagent.NewSqlManagedInstanceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedInstanceValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
