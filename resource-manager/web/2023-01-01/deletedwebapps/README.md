
## `github.com/hashicorp/go-azure-sdk/resource-manager/web/2023-01-01/deletedwebapps` Documentation

The `deletedwebapps` SDK allows for interaction with the Azure Resource Manager Service `web` (API Version `2023-01-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/web/2023-01-01/deletedwebapps"
```


### Client Initialization

```go
client := deletedwebapps.NewDeletedWebAppsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DeletedWebAppsClient.GetDeletedWebAppByLocation`

```go
ctx := context.TODO()
id := deletedwebapps.NewLocationDeletedSiteID("12345678-1234-9876-4563-123456789012", "locationValue", "deletedSiteIdValue")

read, err := client.GetDeletedWebAppByLocation(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeletedWebAppsClient.List`

```go
ctx := context.TODO()
id := deletedwebapps.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.List(ctx, id)` can be used to do batched pagination
items, err := client.ListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DeletedWebAppsClient.ListByLocation`

```go
ctx := context.TODO()
id := deletedwebapps.NewProviderLocationID("12345678-1234-9876-4563-123456789012", "locationValue")

// alternatively `client.ListByLocation(ctx, id)` can be used to do batched pagination
items, err := client.ListByLocationComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
