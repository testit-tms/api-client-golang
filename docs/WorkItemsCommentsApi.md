# \WorkItemsCommentsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV2WorkItemsCommentsCommentIdDelete**](WorkItemsCommentsAPI.md#ApiV2WorkItemsCommentsCommentIdDelete) | **Delete** /api/v2/workItems/comments/{commentId} | Delete WorkItem comment
[**ApiV2WorkItemsCommentsPost**](WorkItemsCommentsAPI.md#ApiV2WorkItemsCommentsPost) | **Post** /api/v2/workItems/comments | Create WorkItem comment
[**ApiV2WorkItemsCommentsPut**](WorkItemsCommentsAPI.md#ApiV2WorkItemsCommentsPut) | **Put** /api/v2/workItems/comments | Update work item comment
[**ApiV2WorkItemsIdCommentsCountGet**](WorkItemsCommentsAPI.md#ApiV2WorkItemsIdCommentsCountGet) | **Get** /api/v2/workItems/{id}/comments/count | Get work item comments count
[**ApiV2WorkItemsIdCommentsGet**](WorkItemsCommentsAPI.md#ApiV2WorkItemsIdCommentsGet) | **Get** /api/v2/workItems/{id}/comments | Get work item comments



## ApiV2WorkItemsCommentsCommentIdDelete

> ApiV2WorkItemsCommentsCommentIdDelete(ctx, commentId).Execute()

Delete WorkItem comment



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	commentId := "3fa85f64-5717-4562-b3fc-2c963f66afa6" // string | Comment internal (guid format) identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WorkItemsCommentsAPI.ApiV2WorkItemsCommentsCommentIdDelete(context.Background(), commentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkItemsCommentsAPI.ApiV2WorkItemsCommentsCommentIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**commentId** | **string** | Comment internal (guid format) identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2WorkItemsCommentsCommentIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2WorkItemsCommentsPost

> WorkItemCommentApiResult ApiV2WorkItemsCommentsPost(ctx).CreateWorkItemCommentApiModel(createWorkItemCommentApiModel).Execute()

Create WorkItem comment



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	createWorkItemCommentApiModel := *openapiclient.NewCreateWorkItemCommentApiModel("WorkItemId_example", "Text_example") // CreateWorkItemCommentApiModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkItemsCommentsAPI.ApiV2WorkItemsCommentsPost(context.Background()).CreateWorkItemCommentApiModel(createWorkItemCommentApiModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkItemsCommentsAPI.ApiV2WorkItemsCommentsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2WorkItemsCommentsPost`: WorkItemCommentApiResult
	fmt.Fprintf(os.Stdout, "Response from `WorkItemsCommentsAPI.ApiV2WorkItemsCommentsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2WorkItemsCommentsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createWorkItemCommentApiModel** | [**CreateWorkItemCommentApiModel**](CreateWorkItemCommentApiModel.md) |  | 

### Return type

[**WorkItemCommentApiResult**](WorkItemCommentApiResult.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2WorkItemsCommentsPut

> ApiV2WorkItemsCommentsPut(ctx).UpdateWorkItemCommentApiModel(updateWorkItemCommentApiModel).Execute()

Update work item comment

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	updateWorkItemCommentApiModel := *openapiclient.NewUpdateWorkItemCommentApiModel("Id_example", "Text_example") // UpdateWorkItemCommentApiModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WorkItemsCommentsAPI.ApiV2WorkItemsCommentsPut(context.Background()).UpdateWorkItemCommentApiModel(updateWorkItemCommentApiModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkItemsCommentsAPI.ApiV2WorkItemsCommentsPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2WorkItemsCommentsPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateWorkItemCommentApiModel** | [**UpdateWorkItemCommentApiModel**](UpdateWorkItemCommentApiModel.md) |  | 

### Return type

 (empty response body)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2WorkItemsIdCommentsCountGet

> int32 ApiV2WorkItemsIdCommentsCountGet(ctx, id).Execute()

Get work item comments count

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "id_example" // string | Unique or global ID of the work item

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkItemsCommentsAPI.ApiV2WorkItemsIdCommentsCountGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkItemsCommentsAPI.ApiV2WorkItemsIdCommentsCountGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2WorkItemsIdCommentsCountGet`: int32
	fmt.Fprintf(os.Stdout, "Response from `WorkItemsCommentsAPI.ApiV2WorkItemsIdCommentsCountGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Unique or global ID of the work item | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2WorkItemsIdCommentsCountGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**int32**

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2WorkItemsIdCommentsGet

> []WorkItemCommentApiResult ApiV2WorkItemsIdCommentsGet(ctx, id).Execute()

Get work item comments

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "id_example" // string | Unique or global ID of the work item

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkItemsCommentsAPI.ApiV2WorkItemsIdCommentsGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkItemsCommentsAPI.ApiV2WorkItemsIdCommentsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2WorkItemsIdCommentsGet`: []WorkItemCommentApiResult
	fmt.Fprintf(os.Stdout, "Response from `WorkItemsCommentsAPI.ApiV2WorkItemsIdCommentsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Unique or global ID of the work item | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2WorkItemsIdCommentsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]WorkItemCommentApiResult**](WorkItemCommentApiResult.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

