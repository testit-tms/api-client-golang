# \BackgroundJobsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV2BackgroundJobsGet**](BackgroundJobsAPI.md#ApiV2BackgroundJobsGet) | **Get** /api/v2/backgroundJobs | 
[**ApiV2BackgroundJobsIdCancelPost**](BackgroundJobsAPI.md#ApiV2BackgroundJobsIdCancelPost) | **Post** /api/v2/backgroundJobs/{id}/cancel | Cancel current user background job
[**ApiV2BackgroundJobsIdGet**](BackgroundJobsAPI.md#ApiV2BackgroundJobsIdGet) | **Get** /api/v2/backgroundJobs/{id} | Get background job by ID
[**ApiV2BackgroundJobsIdStatusGet**](BackgroundJobsAPI.md#ApiV2BackgroundJobsIdStatusGet) | **Get** /api/v2/backgroundJobs/{id}/status | Get background job status by job ID
[**ApiV2BackgroundJobsSearchPost**](BackgroundJobsAPI.md#ApiV2BackgroundJobsSearchPost) | **Post** /api/v2/backgroundJobs/search | Search for user background jobs



## ApiV2BackgroundJobsGet

> []BackgroundJobGetModel ApiV2BackgroundJobsGet(ctx).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).Execute()



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
	resp, r, err := apiClient.BackgroundJobsAPI.ApiV2BackgroundJobsGet(context.Background()).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackgroundJobsAPI.ApiV2BackgroundJobsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2BackgroundJobsGet`: []BackgroundJobGetModel
	fmt.Fprintf(os.Stdout, "Response from `BackgroundJobsAPI.ApiV2BackgroundJobsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2BackgroundJobsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skip** | **int32** | Amount of items to be skipped (offset) | 
 **take** | **int32** | Amount of items to be taken (limit) | 
 **orderBy** | **string** | SQL-like  ORDER BY statement (column1 ASC|DESC , column2 ASC|DESC) | 
 **searchField** | **string** | Property name for searching | 
 **searchValue** | **string** | Value for searching | 

### Return type

[**[]BackgroundJobGetModel**](BackgroundJobGetModel.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2BackgroundJobsIdCancelPost

> ApiV2BackgroundJobsIdCancelPost(ctx, id).Execute()

Cancel current user background job

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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BackgroundJobsAPI.ApiV2BackgroundJobsIdCancelPost(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackgroundJobsAPI.ApiV2BackgroundJobsIdCancelPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2BackgroundJobsIdCancelPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2BackgroundJobsIdGet

> BackgroundJobGetModel ApiV2BackgroundJobsIdGet(ctx, id).Execute()

Get background job by ID

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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique ID of the background job

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackgroundJobsAPI.ApiV2BackgroundJobsIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackgroundJobsAPI.ApiV2BackgroundJobsIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2BackgroundJobsIdGet`: BackgroundJobGetModel
	fmt.Fprintf(os.Stdout, "Response from `BackgroundJobsAPI.ApiV2BackgroundJobsIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Unique ID of the background job | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2BackgroundJobsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BackgroundJobGetModel**](BackgroundJobGetModel.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2BackgroundJobsIdStatusGet

> BackgroundJobState ApiV2BackgroundJobsIdStatusGet(ctx, id).Execute()

Get background job status by job ID

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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique ID of the background job

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackgroundJobsAPI.ApiV2BackgroundJobsIdStatusGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackgroundJobsAPI.ApiV2BackgroundJobsIdStatusGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2BackgroundJobsIdStatusGet`: BackgroundJobState
	fmt.Fprintf(os.Stdout, "Response from `BackgroundJobsAPI.ApiV2BackgroundJobsIdStatusGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Unique ID of the background job | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2BackgroundJobsIdStatusGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BackgroundJobState**](BackgroundJobState.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2BackgroundJobsSearchPost

> []BackgroundJobGetModel ApiV2BackgroundJobsSearchPost(ctx).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).BackgroundJobFilterModel(backgroundJobFilterModel).Execute()

Search for user background jobs

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
	backgroundJobFilterModel := *openapiclient.NewBackgroundJobFilterModel() // BackgroundJobFilterModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackgroundJobsAPI.ApiV2BackgroundJobsSearchPost(context.Background()).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).BackgroundJobFilterModel(backgroundJobFilterModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackgroundJobsAPI.ApiV2BackgroundJobsSearchPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2BackgroundJobsSearchPost`: []BackgroundJobGetModel
	fmt.Fprintf(os.Stdout, "Response from `BackgroundJobsAPI.ApiV2BackgroundJobsSearchPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2BackgroundJobsSearchPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skip** | **int32** | Amount of items to be skipped (offset) | 
 **take** | **int32** | Amount of items to be taken (limit) | 
 **orderBy** | **string** | SQL-like  ORDER BY statement (column1 ASC|DESC , column2 ASC|DESC) | 
 **searchField** | **string** | Property name for searching | 
 **searchValue** | **string** | Value for searching | 
 **backgroundJobFilterModel** | [**BackgroundJobFilterModel**](BackgroundJobFilterModel.md) |  | 

### Return type

[**[]BackgroundJobGetModel**](BackgroundJobGetModel.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

