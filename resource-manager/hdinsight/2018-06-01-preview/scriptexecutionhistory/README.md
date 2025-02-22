
## `github.com/hashicorp/go-azure-sdk/resource-manager/hdinsight/2018-06-01-preview/scriptexecutionhistory` Documentation

The `scriptexecutionhistory` SDK allows for interaction with the Azure Resource Manager Service `hdinsight` (API Version `2018-06-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/hdinsight/2018-06-01-preview/scriptexecutionhistory"
```


### Client Initialization

```go
client := scriptexecutionhistory.NewScriptExecutionHistoryClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ScriptExecutionHistoryClient.ListByCluster`

```go
ctx := context.TODO()
id := scriptexecutionhistory.NewHDInsightClusterID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterValue")

// alternatively `client.ListByCluster(ctx, id)` can be used to do batched pagination
items, err := client.ListByClusterComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ScriptExecutionHistoryClient.ScriptActionsGetExecutionDetail`

```go
ctx := context.TODO()
id := scriptexecutionhistory.NewScriptExecutionHistoryID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterValue", "scriptExecutionIdValue")

read, err := client.ScriptActionsGetExecutionDetail(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
