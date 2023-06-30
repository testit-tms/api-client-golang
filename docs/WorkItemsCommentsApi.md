# \WorkItemsCommentsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV2WorkItemsCommentsCommentIdDelete**](WorkItemsCommentsApi.md#ApiV2WorkItemsCommentsCommentIdDelete) | **Delete** /api/v2/workItems/comments/{commentId} | Delete WorkItem comment
[**ApiV2WorkItemsCommentsPost**](WorkItemsCommentsApi.md#ApiV2WorkItemsCommentsPost) | **Post** /api/v2/workItems/comments | Create WorkItem comment
[**ApiV2WorkItemsCommentsPut**](WorkItemsCommentsApi.md#ApiV2WorkItemsCommentsPut) | **Put** /api/v2/workItems/comments | Update work item comment
[**ApiV2WorkItemsIdCommentsGet**](WorkItemsCommentsApi.md#ApiV2WorkItemsIdCommentsGet) | **Get** /api/v2/workItems/{id}/comments | Get work item comments



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
    r, err := apiClient.WorkItemsCommentsApi.ApiV2WorkItemsCommentsCommentIdDelete(context.Background(), commentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkItemsCommentsApi.ApiV2WorkItemsCommentsCommentIdDelete``: %v\n", err)
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

> WorkItemCommentModel ApiV2WorkItemsCommentsPost(ctx).ApiV2WorkItemsCommentsPostRequest(apiV2WorkItemsCommentsPostRequest).Execute()

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
    apiV2WorkItemsCommentsPostRequest := *openapiclient.NewApiV2WorkItemsCommentsPostRequest("Text_example", "WorkItemId_example") // ApiV2WorkItemsCommentsPostRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkItemsCommentsApi.ApiV2WorkItemsCommentsPost(context.Background()).ApiV2WorkItemsCommentsPostRequest(apiV2WorkItemsCommentsPostRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkItemsCommentsApi.ApiV2WorkItemsCommentsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV2WorkItemsCommentsPost`: WorkItemCommentModel
    fmt.Fprintf(os.Stdout, "Response from `WorkItemsCommentsApi.ApiV2WorkItemsCommentsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2WorkItemsCommentsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiV2WorkItemsCommentsPostRequest** | [**ApiV2WorkItemsCommentsPostRequest**](ApiV2WorkItemsCommentsPostRequest.md) |  | 

### Return type

[**WorkItemCommentModel**](WorkItemCommentModel.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2WorkItemsCommentsPut

> ApiV2WorkItemsCommentsPut(ctx).ApiV2WorkItemsCommentsPutRequest(apiV2WorkItemsCommentsPutRequest).Execute()

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
    apiV2WorkItemsCommentsPutRequest := *openapiclient.NewApiV2WorkItemsCommentsPutRequest("Text_example", "Id_example") // ApiV2WorkItemsCommentsPutRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.WorkItemsCommentsApi.ApiV2WorkItemsCommentsPut(context.Background()).ApiV2WorkItemsCommentsPutRequest(apiV2WorkItemsCommentsPutRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkItemsCommentsApi.ApiV2WorkItemsCommentsPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2WorkItemsCommentsPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiV2WorkItemsCommentsPutRequest** | [**ApiV2WorkItemsCommentsPutRequest**](ApiV2WorkItemsCommentsPutRequest.md) |  | 

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


## ApiV2WorkItemsIdCommentsGet

> []WorkItemCommentModel ApiV2WorkItemsIdCommentsGet(ctx, id).Execute()

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
    resp, r, err := apiClient.WorkItemsCommentsApi.ApiV2WorkItemsIdCommentsGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkItemsCommentsApi.ApiV2WorkItemsIdCommentsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV2WorkItemsIdCommentsGet`: []WorkItemCommentModel
    fmt.Fprintf(os.Stdout, "Response from `WorkItemsCommentsApi.ApiV2WorkItemsIdCommentsGet`: %v\n", resp)
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

[**[]WorkItemCommentModel**](WorkItemCommentModel.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

