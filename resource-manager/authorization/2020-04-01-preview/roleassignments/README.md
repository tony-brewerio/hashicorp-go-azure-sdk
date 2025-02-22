
## `github.com/hashicorp/go-azure-sdk/resource-manager/authorization/2020-04-01-preview/roleassignments` Documentation

The `roleassignments` SDK allows for interaction with the Azure Resource Manager Service `authorization` (API Version `2020-04-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/authorization/2020-04-01-preview/roleassignments"
```


### Client Initialization

```go
client := roleassignments.NewRoleAssignmentsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `RoleAssignmentsClient.Create`

```go
ctx := context.TODO()
id := roleassignments.NewScopedRoleAssignmentID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group", "roleAssignmentValue")

payload := roleassignments.RoleAssignmentCreateParameters{
	// ...
}


read, err := client.Create(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `RoleAssignmentsClient.CreateById`

```go
ctx := context.TODO()
id := roleassignments.NewRoleIdID("roleIdValue")

payload := roleassignments.RoleAssignmentCreateParameters{
	// ...
}


read, err := client.CreateById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `RoleAssignmentsClient.Delete`

```go
ctx := context.TODO()
id := roleassignments.NewScopedRoleAssignmentID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group", "roleAssignmentValue")

read, err := client.Delete(ctx, id, roleassignments.DefaultDeleteOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `RoleAssignmentsClient.DeleteById`

```go
ctx := context.TODO()
id := roleassignments.NewRoleIdID("roleIdValue")

read, err := client.DeleteById(ctx, id, roleassignments.DefaultDeleteByIdOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `RoleAssignmentsClient.Get`

```go
ctx := context.TODO()
id := roleassignments.NewScopedRoleAssignmentID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group", "roleAssignmentValue")

read, err := client.Get(ctx, id, roleassignments.DefaultGetOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `RoleAssignmentsClient.GetById`

```go
ctx := context.TODO()
id := roleassignments.NewRoleIdID("roleIdValue")

read, err := client.GetById(ctx, id, roleassignments.DefaultGetByIdOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `RoleAssignmentsClient.List`

```go
ctx := context.TODO()
id := roleassignments.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.List(ctx, id, roleassignments.DefaultListOperationOptions())` can be used to do batched pagination
items, err := client.ListComplete(ctx, id, roleassignments.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `RoleAssignmentsClient.ListForResource`

```go
ctx := context.TODO()
id := roleassignments.NewScopeID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group")

// alternatively `client.ListForResource(ctx, id, roleassignments.DefaultListForResourceOperationOptions())` can be used to do batched pagination
items, err := client.ListForResourceComplete(ctx, id, roleassignments.DefaultListForResourceOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `RoleAssignmentsClient.ListForResourceGroup`

```go
ctx := context.TODO()
id := roleassignments.NewResourceGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group")

// alternatively `client.ListForResourceGroup(ctx, id, roleassignments.DefaultListForResourceGroupOperationOptions())` can be used to do batched pagination
items, err := client.ListForResourceGroupComplete(ctx, id, roleassignments.DefaultListForResourceGroupOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `RoleAssignmentsClient.ListForScope`

```go
ctx := context.TODO()
id := roleassignments.NewScopeID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group")

// alternatively `client.ListForScope(ctx, id, roleassignments.DefaultListForScopeOperationOptions())` can be used to do batched pagination
items, err := client.ListForScopeComplete(ctx, id, roleassignments.DefaultListForScopeOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
