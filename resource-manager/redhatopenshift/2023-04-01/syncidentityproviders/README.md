
## `github.com/hashicorp/go-azure-sdk/resource-manager/redhatopenshift/2023-04-01/syncidentityproviders` Documentation

The `syncidentityproviders` SDK allows for interaction with the Azure Resource Manager Service `redhatopenshift` (API Version `2023-04-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/redhatopenshift/2023-04-01/syncidentityproviders"
```


### Client Initialization

```go
client := syncidentityproviders.NewSyncIdentityProvidersClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SyncIdentityProvidersClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := syncidentityproviders.NewSyncIdentityProviderID("12345678-1234-9876-4563-123456789012", "example-resource-group", "openShiftClusterValue", "syncIdentityProviderValue")

payload := syncidentityproviders.SyncIdentityProvider{
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


### Example Usage: `SyncIdentityProvidersClient.Delete`

```go
ctx := context.TODO()
id := syncidentityproviders.NewSyncIdentityProviderID("12345678-1234-9876-4563-123456789012", "example-resource-group", "openShiftClusterValue", "syncIdentityProviderValue")

read, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SyncIdentityProvidersClient.Get`

```go
ctx := context.TODO()
id := syncidentityproviders.NewSyncIdentityProviderID("12345678-1234-9876-4563-123456789012", "example-resource-group", "openShiftClusterValue", "syncIdentityProviderValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SyncIdentityProvidersClient.List`

```go
ctx := context.TODO()
id := syncidentityproviders.NewOpenShiftClusterID("12345678-1234-9876-4563-123456789012", "example-resource-group", "openShiftClusterValue")

// alternatively `client.List(ctx, id)` can be used to do batched pagination
items, err := client.ListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `SyncIdentityProvidersClient.Update`

```go
ctx := context.TODO()
id := syncidentityproviders.NewSyncIdentityProviderID("12345678-1234-9876-4563-123456789012", "example-resource-group", "openShiftClusterValue", "syncIdentityProviderValue")

payload := syncidentityproviders.SyncIdentityProviderUpdate{
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
