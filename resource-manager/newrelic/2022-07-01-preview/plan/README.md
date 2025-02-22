
## `github.com/hashicorp/go-azure-sdk/resource-manager/newrelic/2022-07-01-preview/plan` Documentation

The `plan` SDK allows for interaction with the Azure Resource Manager Service `newrelic` (API Version `2022-07-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/newrelic/2022-07-01-preview/plan"
```


### Client Initialization

```go
client := plan.NewPlanClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `PlanClient.List`

```go
ctx := context.TODO()
id := plan.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.List(ctx, id, plan.DefaultListOperationOptions())` can be used to do batched pagination
items, err := client.ListComplete(ctx, id, plan.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
