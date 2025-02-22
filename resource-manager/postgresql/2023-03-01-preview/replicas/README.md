
## `github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2023-03-01-preview/replicas` Documentation

The `replicas` SDK allows for interaction with the Azure Resource Manager Service `postgresql` (API Version `2023-03-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2023-03-01-preview/replicas"
```


### Client Initialization

```go
client := replicas.NewReplicasClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ReplicasClient.ListByServer`

```go
ctx := context.TODO()
id := replicas.NewFlexibleServerID("12345678-1234-9876-4563-123456789012", "example-resource-group", "flexibleServerValue")

read, err := client.ListByServer(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
