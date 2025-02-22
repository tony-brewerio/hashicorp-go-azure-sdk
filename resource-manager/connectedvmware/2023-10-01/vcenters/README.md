
## `github.com/hashicorp/go-azure-sdk/resource-manager/connectedvmware/2023-10-01/vcenters` Documentation

The `vcenters` SDK allows for interaction with the Azure Resource Manager Service `connectedvmware` (API Version `2023-10-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/connectedvmware/2023-10-01/vcenters"
```


### Client Initialization

```go
client := vcenters.NewVCentersClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `VCentersClient.Create`

```go
ctx := context.TODO()
id := vcenters.NewVCenterID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vCenterValue")

payload := vcenters.VCenter{
	// ...
}


if err := client.CreateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `VCentersClient.Delete`

```go
ctx := context.TODO()
id := vcenters.NewVCenterID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vCenterValue")

if err := client.DeleteThenPoll(ctx, id, vcenters.DefaultDeleteOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `VCentersClient.Get`

```go
ctx := context.TODO()
id := vcenters.NewVCenterID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vCenterValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VCentersClient.List`

```go
ctx := context.TODO()
id := vcenters.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.List(ctx, id)` can be used to do batched pagination
items, err := client.ListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `VCentersClient.ListByResourceGroup`

```go
ctx := context.TODO()
id := vcenters.NewResourceGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group")

// alternatively `client.ListByResourceGroup(ctx, id)` can be used to do batched pagination
items, err := client.ListByResourceGroupComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `VCentersClient.Update`

```go
ctx := context.TODO()
id := vcenters.NewVCenterID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vCenterValue")

payload := vcenters.ResourcePatch{
	// ...
}


read, err := client.Update(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
