# \ProjectWorkItemsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV2ProjectsProjectIdWorkItemsSearchGroupedPost**](ProjectWorkItemsApi.md#ApiV2ProjectsProjectIdWorkItemsSearchGroupedPost) | **Post** /api/v2/projects/{projectId}/workItems/search/grouped | Search for work items and group results by attribute
[**ApiV2ProjectsProjectIdWorkItemsSearchIdPost**](ProjectWorkItemsApi.md#ApiV2ProjectsProjectIdWorkItemsSearchIdPost) | **Post** /api/v2/projects/{projectId}/workItems/search/id | Search for work items and extract IDs only
[**ApiV2ProjectsProjectIdWorkItemsSearchPost**](ProjectWorkItemsApi.md#ApiV2ProjectsProjectIdWorkItemsSearchPost) | **Post** /api/v2/projects/{projectId}/workItems/search | Search for work items
[**ApiV2ProjectsProjectIdWorkItemsTagsGet**](ProjectWorkItemsApi.md#ApiV2ProjectsProjectIdWorkItemsTagsGet) | **Get** /api/v2/projects/{projectId}/workItems/tags | Get WorkItems Tags
[**GetWorkItemsByProjectId**](ProjectWorkItemsApi.md#GetWorkItemsByProjectId) | **Get** /api/v2/projects/{projectId}/workItems | Get project work items



## ApiV2ProjectsProjectIdWorkItemsSearchGroupedPost

> []WorkItemGroupModel ApiV2ProjectsProjectIdWorkItemsSearchGroupedPost(ctx, projectId).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).ApiV2ProjectsProjectIdWorkItemsSearchGroupedPostRequest(apiV2ProjectsProjectIdWorkItemsSearchGroupedPostRequest).Execute()

Search for work items and group results by attribute

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
    projectId := "projectId_example" // string | Unique or global ID of the project
    skip := int32(56) // int32 | Amount of items to be skipped (offset) (optional)
    take := int32(56) // int32 | Amount of items to be taken (limit) (optional)
    orderBy := "orderBy_example" // string | SQL-like  ORDER BY statement (column1 ASC|DESC , column2 ASC|DESC) (optional)
    searchField := "searchField_example" // string | Property name for searching (optional)
    searchValue := "searchValue_example" // string | Value for searching (optional)
    apiV2ProjectsProjectIdWorkItemsSearchGroupedPostRequest := *openapiclient.NewApiV2ProjectsProjectIdWorkItemsSearchGroupedPostRequest(openapiclient.WorkItemGroupType("Type")) // ApiV2ProjectsProjectIdWorkItemsSearchGroupedPostRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectWorkItemsApi.ApiV2ProjectsProjectIdWorkItemsSearchGroupedPost(context.Background(), projectId).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).ApiV2ProjectsProjectIdWorkItemsSearchGroupedPostRequest(apiV2ProjectsProjectIdWorkItemsSearchGroupedPostRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectWorkItemsApi.ApiV2ProjectsProjectIdWorkItemsSearchGroupedPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV2ProjectsProjectIdWorkItemsSearchGroupedPost`: []WorkItemGroupModel
    fmt.Fprintf(os.Stdout, "Response from `ProjectWorkItemsApi.ApiV2ProjectsProjectIdWorkItemsSearchGroupedPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Unique or global ID of the project | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ProjectsProjectIdWorkItemsSearchGroupedPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **skip** | **int32** | Amount of items to be skipped (offset) | 
 **take** | **int32** | Amount of items to be taken (limit) | 
 **orderBy** | **string** | SQL-like  ORDER BY statement (column1 ASC|DESC , column2 ASC|DESC) | 
 **searchField** | **string** | Property name for searching | 
 **searchValue** | **string** | Value for searching | 
 **apiV2ProjectsProjectIdWorkItemsSearchGroupedPostRequest** | [**ApiV2ProjectsProjectIdWorkItemsSearchGroupedPostRequest**](ApiV2ProjectsProjectIdWorkItemsSearchGroupedPostRequest.md) |  | 

### Return type

[**[]WorkItemGroupModel**](WorkItemGroupModel.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2ProjectsProjectIdWorkItemsSearchIdPost

> []string ApiV2ProjectsProjectIdWorkItemsSearchIdPost(ctx, projectId).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).ApiV2ProjectsProjectIdWorkItemsSearchPostRequest(apiV2ProjectsProjectIdWorkItemsSearchPostRequest).Execute()

Search for work items and extract IDs only

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
    projectId := "projectId_example" // string | Unique or global ID of the project
    skip := int32(56) // int32 | Amount of items to be skipped (offset) (optional)
    take := int32(56) // int32 | Amount of items to be taken (limit) (optional)
    orderBy := "orderBy_example" // string | SQL-like  ORDER BY statement (column1 ASC|DESC , column2 ASC|DESC) (optional)
    searchField := "searchField_example" // string | Property name for searching (optional)
    searchValue := "searchValue_example" // string | Value for searching (optional)
    apiV2ProjectsProjectIdWorkItemsSearchPostRequest := *openapiclient.NewApiV2ProjectsProjectIdWorkItemsSearchPostRequest(*openapiclient.NewWorkItemSelectModelFilter()) // ApiV2ProjectsProjectIdWorkItemsSearchPostRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectWorkItemsApi.ApiV2ProjectsProjectIdWorkItemsSearchIdPost(context.Background(), projectId).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).ApiV2ProjectsProjectIdWorkItemsSearchPostRequest(apiV2ProjectsProjectIdWorkItemsSearchPostRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectWorkItemsApi.ApiV2ProjectsProjectIdWorkItemsSearchIdPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV2ProjectsProjectIdWorkItemsSearchIdPost`: []string
    fmt.Fprintf(os.Stdout, "Response from `ProjectWorkItemsApi.ApiV2ProjectsProjectIdWorkItemsSearchIdPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Unique or global ID of the project | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ProjectsProjectIdWorkItemsSearchIdPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **skip** | **int32** | Amount of items to be skipped (offset) | 
 **take** | **int32** | Amount of items to be taken (limit) | 
 **orderBy** | **string** | SQL-like  ORDER BY statement (column1 ASC|DESC , column2 ASC|DESC) | 
 **searchField** | **string** | Property name for searching | 
 **searchValue** | **string** | Value for searching | 
 **apiV2ProjectsProjectIdWorkItemsSearchPostRequest** | [**ApiV2ProjectsProjectIdWorkItemsSearchPostRequest**](ApiV2ProjectsProjectIdWorkItemsSearchPostRequest.md) |  | 

### Return type

**[]string**

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2ProjectsProjectIdWorkItemsSearchPost

> []WorkItemShortModel ApiV2ProjectsProjectIdWorkItemsSearchPost(ctx, projectId).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).ApiV2ProjectsProjectIdWorkItemsSearchPostRequest(apiV2ProjectsProjectIdWorkItemsSearchPostRequest).Execute()

Search for work items

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
    projectId := "projectId_example" // string | Unique or global ID of the project
    skip := int32(56) // int32 | Amount of items to be skipped (offset) (optional)
    take := int32(56) // int32 | Amount of items to be taken (limit) (optional)
    orderBy := "orderBy_example" // string | SQL-like  ORDER BY statement (column1 ASC|DESC , column2 ASC|DESC) (optional)
    searchField := "searchField_example" // string | Property name for searching (optional)
    searchValue := "searchValue_example" // string | Value for searching (optional)
    apiV2ProjectsProjectIdWorkItemsSearchPostRequest := *openapiclient.NewApiV2ProjectsProjectIdWorkItemsSearchPostRequest(*openapiclient.NewWorkItemSelectModelFilter()) // ApiV2ProjectsProjectIdWorkItemsSearchPostRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectWorkItemsApi.ApiV2ProjectsProjectIdWorkItemsSearchPost(context.Background(), projectId).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).ApiV2ProjectsProjectIdWorkItemsSearchPostRequest(apiV2ProjectsProjectIdWorkItemsSearchPostRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectWorkItemsApi.ApiV2ProjectsProjectIdWorkItemsSearchPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV2ProjectsProjectIdWorkItemsSearchPost`: []WorkItemShortModel
    fmt.Fprintf(os.Stdout, "Response from `ProjectWorkItemsApi.ApiV2ProjectsProjectIdWorkItemsSearchPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Unique or global ID of the project | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ProjectsProjectIdWorkItemsSearchPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **skip** | **int32** | Amount of items to be skipped (offset) | 
 **take** | **int32** | Amount of items to be taken (limit) | 
 **orderBy** | **string** | SQL-like  ORDER BY statement (column1 ASC|DESC , column2 ASC|DESC) | 
 **searchField** | **string** | Property name for searching | 
 **searchValue** | **string** | Value for searching | 
 **apiV2ProjectsProjectIdWorkItemsSearchPostRequest** | [**ApiV2ProjectsProjectIdWorkItemsSearchPostRequest**](ApiV2ProjectsProjectIdWorkItemsSearchPostRequest.md) |  | 

### Return type

[**[]WorkItemShortModel**](WorkItemShortModel.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2ProjectsProjectIdWorkItemsTagsGet

> []TagModel ApiV2ProjectsProjectIdWorkItemsTagsGet(ctx, projectId).IsDeleted(isDeleted).Execute()

Get WorkItems Tags



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
    projectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Project internal (UUID) identifier
    isDeleted := true // bool |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectWorkItemsApi.ApiV2ProjectsProjectIdWorkItemsTagsGet(context.Background(), projectId).IsDeleted(isDeleted).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectWorkItemsApi.ApiV2ProjectsProjectIdWorkItemsTagsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV2ProjectsProjectIdWorkItemsTagsGet`: []TagModel
    fmt.Fprintf(os.Stdout, "Response from `ProjectWorkItemsApi.ApiV2ProjectsProjectIdWorkItemsTagsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project internal (UUID) identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ProjectsProjectIdWorkItemsTagsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **isDeleted** | **bool** |  | 

### Return type

[**[]TagModel**](TagModel.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWorkItemsByProjectId

> []WorkItemShortModel GetWorkItemsByProjectId(ctx, projectId).IsDeleted(isDeleted).TagNames(tagNames).IncludeIterations(includeIterations).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).Execute()

Get project work items



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
    projectId := "projectId_example" // string | Project internal (UUID) or global (integer) identifier
    isDeleted := true // bool | If result must consist of only actual/deleted work items (optional) (default to false)
    tagNames := []string{"Inner_example"} // []string | List of tags to filter by (optional)
    includeIterations := true // bool |  (optional) (default to true)
    skip := int32(56) // int32 | Amount of items to be skipped (offset) (optional)
    take := int32(56) // int32 | Amount of items to be taken (limit) (optional)
    orderBy := "orderBy_example" // string | SQL-like  ORDER BY statement (column1 ASC|DESC , column2 ASC|DESC) (optional)
    searchField := "searchField_example" // string | Property name for searching (optional)
    searchValue := "searchValue_example" // string | Value for searching (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectWorkItemsApi.GetWorkItemsByProjectId(context.Background(), projectId).IsDeleted(isDeleted).TagNames(tagNames).IncludeIterations(includeIterations).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectWorkItemsApi.GetWorkItemsByProjectId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorkItemsByProjectId`: []WorkItemShortModel
    fmt.Fprintf(os.Stdout, "Response from `ProjectWorkItemsApi.GetWorkItemsByProjectId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project internal (UUID) or global (integer) identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkItemsByProjectIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **isDeleted** | **bool** | If result must consist of only actual/deleted work items | [default to false]
 **tagNames** | **[]string** | List of tags to filter by | 
 **includeIterations** | **bool** |  | [default to true]
 **skip** | **int32** | Amount of items to be skipped (offset) | 
 **take** | **int32** | Amount of items to be taken (limit) | 
 **orderBy** | **string** | SQL-like  ORDER BY statement (column1 ASC|DESC , column2 ASC|DESC) | 
 **searchField** | **string** | Property name for searching | 
 **searchValue** | **string** | Value for searching | 

### Return type

[**[]WorkItemShortModel**](WorkItemShortModel.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

