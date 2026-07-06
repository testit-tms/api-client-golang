# \ProjectFailureCategoriesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV2ProjectsProjectIdAutotestsFailureCategoriesGroupingSearchPost**](ProjectFailureCategoriesAPI.md#ApiV2ProjectsProjectIdAutotestsFailureCategoriesGroupingSearchPost) | **Post** /api/v2/projects/{projectId}/autotests/failure-categories/grouping-search | Get failure categories with support for filtering, sorting and grouping
[**ApiV2ProjectsProjectIdAutotestsFailureCategoriesIdDelete**](ProjectFailureCategoriesAPI.md#ApiV2ProjectsProjectIdAutotestsFailureCategoriesIdDelete) | **Delete** /api/v2/projects/{projectId}/autotests/failure-categories/{id} | Delete failure category
[**ApiV2ProjectsProjectIdAutotestsFailureCategoriesIdGet**](ProjectFailureCategoriesAPI.md#ApiV2ProjectsProjectIdAutotestsFailureCategoriesIdGet) | **Get** /api/v2/projects/{projectId}/autotests/failure-categories/{id} | Get failure category by ID
[**ApiV2ProjectsProjectIdAutotestsFailureCategoriesPost**](ProjectFailureCategoriesAPI.md#ApiV2ProjectsProjectIdAutotestsFailureCategoriesPost) | **Post** /api/v2/projects/{projectId}/autotests/failure-categories | Create failure category
[**ApiV2ProjectsProjectIdAutotestsFailureCategoriesPut**](ProjectFailureCategoriesAPI.md#ApiV2ProjectsProjectIdAutotestsFailureCategoriesPut) | **Put** /api/v2/projects/{projectId}/autotests/failure-categories | Update failure category



## ApiV2ProjectsProjectIdAutotestsFailureCategoriesGroupingSearchPost

> ProjectFailureCategoryGroupItemApiResultReply ApiV2ProjectsProjectIdAutotestsFailureCategoriesGroupingSearchPost(ctx, projectId).FailureCategoryGroupSearchApiModel(failureCategoryGroupSearchApiModel).Execute()

Get failure categories with support for filtering, sorting and grouping

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
	projectId := "projectId_example" // string | Internal (UUID) or global (integer) identifier
	failureCategoryGroupSearchApiModel := *openapiclient.NewFailureCategoryGroupSearchApiModel(*openapiclient.NewInquiry([]openapiclient.Order{*openapiclient.NewOrder("Field_example", openapiclient.ListSortDirection("Ascending"))})) // FailureCategoryGroupSearchApiModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectFailureCategoriesAPI.ApiV2ProjectsProjectIdAutotestsFailureCategoriesGroupingSearchPost(context.Background(), projectId).FailureCategoryGroupSearchApiModel(failureCategoryGroupSearchApiModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectFailureCategoriesAPI.ApiV2ProjectsProjectIdAutotestsFailureCategoriesGroupingSearchPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2ProjectsProjectIdAutotestsFailureCategoriesGroupingSearchPost`: ProjectFailureCategoryGroupItemApiResultReply
	fmt.Fprintf(os.Stdout, "Response from `ProjectFailureCategoriesAPI.ApiV2ProjectsProjectIdAutotestsFailureCategoriesGroupingSearchPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Internal (UUID) or global (integer) identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ProjectsProjectIdAutotestsFailureCategoriesGroupingSearchPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **failureCategoryGroupSearchApiModel** | [**FailureCategoryGroupSearchApiModel**](FailureCategoryGroupSearchApiModel.md) |  | 

### Return type

[**ProjectFailureCategoryGroupItemApiResultReply**](ProjectFailureCategoryGroupItemApiResultReply.md)

### Authorization

[PrivateToken](../README.md#PrivateToken), [Cookies](../README.md#Cookies)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2ProjectsProjectIdAutotestsFailureCategoriesIdDelete

> ApiV2ProjectsProjectIdAutotestsFailureCategoriesIdDelete(ctx, projectId, id).Execute()

Delete failure category

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
	projectId := "projectId_example" // string | Internal (UUID) or global (integer) identifier
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectFailureCategoriesAPI.ApiV2ProjectsProjectIdAutotestsFailureCategoriesIdDelete(context.Background(), projectId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectFailureCategoriesAPI.ApiV2ProjectsProjectIdAutotestsFailureCategoriesIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Internal (UUID) or global (integer) identifier | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ProjectsProjectIdAutotestsFailureCategoriesIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[PrivateToken](../README.md#PrivateToken), [Cookies](../README.md#Cookies)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2ProjectsProjectIdAutotestsFailureCategoriesIdGet

> ProjectDetailedFailureCategoryApiResult ApiV2ProjectsProjectIdAutotestsFailureCategoriesIdGet(ctx, projectId, id).Execute()

Get failure category by ID

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
	projectId := "projectId_example" // string | Internal (UUID) or global (integer) identifier
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectFailureCategoriesAPI.ApiV2ProjectsProjectIdAutotestsFailureCategoriesIdGet(context.Background(), projectId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectFailureCategoriesAPI.ApiV2ProjectsProjectIdAutotestsFailureCategoriesIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2ProjectsProjectIdAutotestsFailureCategoriesIdGet`: ProjectDetailedFailureCategoryApiResult
	fmt.Fprintf(os.Stdout, "Response from `ProjectFailureCategoriesAPI.ApiV2ProjectsProjectIdAutotestsFailureCategoriesIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Internal (UUID) or global (integer) identifier | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ProjectsProjectIdAutotestsFailureCategoriesIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ProjectDetailedFailureCategoryApiResult**](ProjectDetailedFailureCategoryApiResult.md)

### Authorization

[PrivateToken](../README.md#PrivateToken), [Cookies](../README.md#Cookies)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2ProjectsProjectIdAutotestsFailureCategoriesPost

> ProjectDetailedFailureCategoryApiResult ApiV2ProjectsProjectIdAutotestsFailureCategoriesPost(ctx, projectId).CreateProjectFailureCategoryApiModel(createProjectFailureCategoryApiModel).Execute()

Create failure category

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
	projectId := "projectId_example" // string | Internal (UUID) or global (integer) identifier
	createProjectFailureCategoryApiModel := *openapiclient.NewCreateProjectFailureCategoryApiModel("Name_example", openapiclient.FailureCategory("InfrastructureDefect")) // CreateProjectFailureCategoryApiModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectFailureCategoriesAPI.ApiV2ProjectsProjectIdAutotestsFailureCategoriesPost(context.Background(), projectId).CreateProjectFailureCategoryApiModel(createProjectFailureCategoryApiModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectFailureCategoriesAPI.ApiV2ProjectsProjectIdAutotestsFailureCategoriesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2ProjectsProjectIdAutotestsFailureCategoriesPost`: ProjectDetailedFailureCategoryApiResult
	fmt.Fprintf(os.Stdout, "Response from `ProjectFailureCategoriesAPI.ApiV2ProjectsProjectIdAutotestsFailureCategoriesPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Internal (UUID) or global (integer) identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ProjectsProjectIdAutotestsFailureCategoriesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createProjectFailureCategoryApiModel** | [**CreateProjectFailureCategoryApiModel**](CreateProjectFailureCategoryApiModel.md) |  | 

### Return type

[**ProjectDetailedFailureCategoryApiResult**](ProjectDetailedFailureCategoryApiResult.md)

### Authorization

[PrivateToken](../README.md#PrivateToken), [Cookies](../README.md#Cookies)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2ProjectsProjectIdAutotestsFailureCategoriesPut

> ApiV2ProjectsProjectIdAutotestsFailureCategoriesPut(ctx, projectId).UpdateFailureCategoryProjectApiModel(updateFailureCategoryProjectApiModel).Execute()

Update failure category

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
	projectId := "projectId_example" // string | Internal (UUID) or global (integer) identifier
	updateFailureCategoryProjectApiModel := *openapiclient.NewUpdateFailureCategoryProjectApiModel("Id_example", "Name_example", openapiclient.FailureCategory("InfrastructureDefect")) // UpdateFailureCategoryProjectApiModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectFailureCategoriesAPI.ApiV2ProjectsProjectIdAutotestsFailureCategoriesPut(context.Background(), projectId).UpdateFailureCategoryProjectApiModel(updateFailureCategoryProjectApiModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectFailureCategoriesAPI.ApiV2ProjectsProjectIdAutotestsFailureCategoriesPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Internal (UUID) or global (integer) identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ProjectsProjectIdAutotestsFailureCategoriesPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateFailureCategoryProjectApiModel** | [**UpdateFailureCategoryProjectApiModel**](UpdateFailureCategoryProjectApiModel.md) |  | 

### Return type

 (empty response body)

### Authorization

[PrivateToken](../README.md#PrivateToken), [Cookies](../README.md#Cookies)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

