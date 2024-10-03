# \TagsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV2TagsDelete**](TagsApi.md#ApiV2TagsDelete) | **Delete** /api/v2/tags | Delete tags
[**ApiV2TagsGet**](TagsApi.md#ApiV2TagsGet) | **Get** /api/v2/tags | Get all Tags
[**ApiV2TagsIdDelete**](TagsApi.md#ApiV2TagsIdDelete) | **Delete** /api/v2/tags/{id} | Delete tag
[**ApiV2TagsPost**](TagsApi.md#ApiV2TagsPost) | **Post** /api/v2/tags | Create tag
[**ApiV2TagsPut**](TagsApi.md#ApiV2TagsPut) | **Put** /api/v2/tags | Update tag
[**ApiV2TagsSearchGet**](TagsApi.md#ApiV2TagsSearchGet) | **Get** /api/v2/tags/search | Search tags
[**ApiV2TagsTestPlansTagsGet**](TagsApi.md#ApiV2TagsTestPlansTagsGet) | **Get** /api/v2/tags/testPlansTags | Get all Tags that are used in TestPlans



## ApiV2TagsDelete

> ApiV2TagsDelete(ctx).ApiV2TagsDeleteRequest(apiV2TagsDeleteRequest).Execute()

Delete tags



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
    apiV2TagsDeleteRequest := *openapiclient.NewApiV2TagsDeleteRequest() // ApiV2TagsDeleteRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TagsApi.ApiV2TagsDelete(context.Background()).ApiV2TagsDeleteRequest(apiV2TagsDeleteRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TagsApi.ApiV2TagsDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2TagsDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiV2TagsDeleteRequest** | [**ApiV2TagsDeleteRequest**](ApiV2TagsDeleteRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2TagsGet

> []TagModel ApiV2TagsGet(ctx).Execute()

Get all Tags



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TagsApi.ApiV2TagsGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TagsApi.ApiV2TagsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV2TagsGet`: []TagModel
    fmt.Fprintf(os.Stdout, "Response from `TagsApi.ApiV2TagsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2TagsGetRequest struct via the builder pattern


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


## ApiV2TagsIdDelete

> ApiV2TagsIdDelete(ctx, id).Execute()

Delete tag



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Tag internal (UUID) identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TagsApi.ApiV2TagsIdDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TagsApi.ApiV2TagsIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Tag internal (UUID) identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2TagsIdDeleteRequest struct via the builder pattern


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


## ApiV2TagsPost

> TagModel ApiV2TagsPost(ctx).ApiV2TagsPostRequest(apiV2TagsPostRequest).Execute()

Create tag



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
    apiV2TagsPostRequest := *openapiclient.NewApiV2TagsPostRequest("Name_example") // ApiV2TagsPostRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TagsApi.ApiV2TagsPost(context.Background()).ApiV2TagsPostRequest(apiV2TagsPostRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TagsApi.ApiV2TagsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV2TagsPost`: TagModel
    fmt.Fprintf(os.Stdout, "Response from `TagsApi.ApiV2TagsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2TagsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiV2TagsPostRequest** | [**ApiV2TagsPostRequest**](ApiV2TagsPostRequest.md) |  | 

### Return type

[**TagModel**](TagModel.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2TagsPut

> TagModel ApiV2TagsPut(ctx).Id(id).ApiV2TagsPutRequest(apiV2TagsPutRequest).Execute()

Update tag



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
    apiV2TagsPutRequest := *openapiclient.NewApiV2TagsPutRequest("Name_example") // ApiV2TagsPutRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TagsApi.ApiV2TagsPut(context.Background()).Id(id).ApiV2TagsPutRequest(apiV2TagsPutRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TagsApi.ApiV2TagsPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV2TagsPut`: TagModel
    fmt.Fprintf(os.Stdout, "Response from `TagsApi.ApiV2TagsPut`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2TagsPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** |  | 
 **apiV2TagsPutRequest** | [**ApiV2TagsPutRequest**](ApiV2TagsPutRequest.md) |  | 

### Return type

[**TagModel**](TagModel.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2TagsSearchGet

> []TagModel ApiV2TagsSearchGet(ctx).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).Execute()

Search tags



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
    skip := int32(56) // int32 | Amount of items to be skipped (offset) (optional)
    take := int32(56) // int32 | Amount of items to be taken (limit) (optional)
    orderBy := "orderBy_example" // string | SQL-like  ORDER BY statement (column1 ASC|DESC , column2 ASC|DESC) (optional)
    searchField := "searchField_example" // string | Property name for searching (optional)
    searchValue := "searchValue_example" // string | Value for searching (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TagsApi.ApiV2TagsSearchGet(context.Background()).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TagsApi.ApiV2TagsSearchGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV2TagsSearchGet`: []TagModel
    fmt.Fprintf(os.Stdout, "Response from `TagsApi.ApiV2TagsSearchGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2TagsSearchGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skip** | **int32** | Amount of items to be skipped (offset) | 
 **take** | **int32** | Amount of items to be taken (limit) | 
 **orderBy** | **string** | SQL-like  ORDER BY statement (column1 ASC|DESC , column2 ASC|DESC) | 
 **searchField** | **string** | Property name for searching | 
 **searchValue** | **string** | Value for searching | 

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


## ApiV2TagsTestPlansTagsGet

> []TagModel ApiV2TagsTestPlansTagsGet(ctx).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).Execute()

Get all Tags that are used in TestPlans



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
    skip := int32(56) // int32 | Amount of items to be skipped (offset) (optional)
    take := int32(56) // int32 | Amount of items to be taken (limit) (optional)
    orderBy := "orderBy_example" // string | SQL-like  ORDER BY statement (column1 ASC|DESC , column2 ASC|DESC) (optional)
    searchField := "searchField_example" // string | Property name for searching (optional)
    searchValue := "searchValue_example" // string | Value for searching (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TagsApi.ApiV2TagsTestPlansTagsGet(context.Background()).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TagsApi.ApiV2TagsTestPlansTagsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV2TagsTestPlansTagsGet`: []TagModel
    fmt.Fprintf(os.Stdout, "Response from `TagsApi.ApiV2TagsTestPlansTagsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2TagsTestPlansTagsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skip** | **int32** | Amount of items to be skipped (offset) | 
 **take** | **int32** | Amount of items to be taken (limit) | 
 **orderBy** | **string** | SQL-like  ORDER BY statement (column1 ASC|DESC , column2 ASC|DESC) | 
 **searchField** | **string** | Property name for searching | 
 **searchValue** | **string** | Value for searching | 

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

