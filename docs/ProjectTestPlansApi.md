# \ProjectTestPlansAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV2ProjectsProjectIdTestPlansAnalyticsGet**](ProjectTestPlansAPI.md#ApiV2ProjectsProjectIdTestPlansAnalyticsGet) | **Get** /api/v2/projects/{projectId}/testPlans/analytics | Get TestPlans analytics
[**ApiV2ProjectsProjectIdTestPlansDeleteBulkPost**](ProjectTestPlansAPI.md#ApiV2ProjectsProjectIdTestPlansDeleteBulkPost) | **Post** /api/v2/projects/{projectId}/testPlans/delete/bulk | Delete multiple test plans
[**ApiV2ProjectsProjectIdTestPlansNameExistsGet**](ProjectTestPlansAPI.md#ApiV2ProjectsProjectIdTestPlansNameExistsGet) | **Get** /api/v2/projects/{projectId}/testPlans/{name}/exists | Checks if TestPlan exists with the specified name exists for the project
[**ApiV2ProjectsProjectIdTestPlansPurgeBulkPost**](ProjectTestPlansAPI.md#ApiV2ProjectsProjectIdTestPlansPurgeBulkPost) | **Post** /api/v2/projects/{projectId}/testPlans/purge/bulk | Permanently delete multiple archived test plans
[**ApiV2ProjectsProjectIdTestPlansRestoreBulkPost**](ProjectTestPlansAPI.md#ApiV2ProjectsProjectIdTestPlansRestoreBulkPost) | **Post** /api/v2/projects/{projectId}/testPlans/restore/bulk | Restore multiple test plans
[**ApiV2ProjectsProjectIdTestPlansSearchPost**](ProjectTestPlansAPI.md#ApiV2ProjectsProjectIdTestPlansSearchPost) | **Post** /api/v2/projects/{projectId}/testPlans/search | Get Project TestPlans with analytics



## ApiV2ProjectsProjectIdTestPlansAnalyticsGet

> []TestPlanWithAnalyticModel ApiV2ProjectsProjectIdTestPlansAnalyticsGet(ctx, projectId).IsDeleted(isDeleted).MustUpdateCache(mustUpdateCache).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).Execute()

Get TestPlans analytics



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
	mustUpdateCache := true // bool |  (optional) (default to false)
	skip := int32(56) // int32 | Amount of items to be skipped (offset) (optional)
	take := int32(56) // int32 | Amount of items to be taken (limit) (optional)
	orderBy := "orderBy_example" // string | SQL-like  ORDER BY statement (column1 ASC|DESC , column2 ASC|DESC) (optional)
	searchField := "searchField_example" // string | Property name for searching (optional)
	searchValue := "searchValue_example" // string | Value for searching (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectTestPlansAPI.ApiV2ProjectsProjectIdTestPlansAnalyticsGet(context.Background(), projectId).IsDeleted(isDeleted).MustUpdateCache(mustUpdateCache).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectTestPlansAPI.ApiV2ProjectsProjectIdTestPlansAnalyticsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2ProjectsProjectIdTestPlansAnalyticsGet`: []TestPlanWithAnalyticModel
	fmt.Fprintf(os.Stdout, "Response from `ProjectTestPlansAPI.ApiV2ProjectsProjectIdTestPlansAnalyticsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project internal (UUID) identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ProjectsProjectIdTestPlansAnalyticsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **isDeleted** | **bool** |  | 
 **mustUpdateCache** | **bool** |  | [default to false]
 **skip** | **int32** | Amount of items to be skipped (offset) | 
 **take** | **int32** | Amount of items to be taken (limit) | 
 **orderBy** | **string** | SQL-like  ORDER BY statement (column1 ASC|DESC , column2 ASC|DESC) | 
 **searchField** | **string** | Property name for searching | 
 **searchValue** | **string** | Value for searching | 

### Return type

[**[]TestPlanWithAnalyticModel**](TestPlanWithAnalyticModel.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2ProjectsProjectIdTestPlansDeleteBulkPost

> []string ApiV2ProjectsProjectIdTestPlansDeleteBulkPost(ctx, projectId).TestPlanSelectModel(testPlanSelectModel).Execute()

Delete multiple test plans

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
	testPlanSelectModel := *openapiclient.NewTestPlanSelectModel(*openapiclient.NewProjectTestPlansFilterModel()) // TestPlanSelectModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectTestPlansAPI.ApiV2ProjectsProjectIdTestPlansDeleteBulkPost(context.Background(), projectId).TestPlanSelectModel(testPlanSelectModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectTestPlansAPI.ApiV2ProjectsProjectIdTestPlansDeleteBulkPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2ProjectsProjectIdTestPlansDeleteBulkPost`: []string
	fmt.Fprintf(os.Stdout, "Response from `ProjectTestPlansAPI.ApiV2ProjectsProjectIdTestPlansDeleteBulkPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Unique or global ID of the project | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ProjectsProjectIdTestPlansDeleteBulkPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **testPlanSelectModel** | [**TestPlanSelectModel**](TestPlanSelectModel.md) |  | 

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


## ApiV2ProjectsProjectIdTestPlansNameExistsGet

> bool ApiV2ProjectsProjectIdTestPlansNameExistsGet(ctx, projectId, name).Execute()

Checks if TestPlan exists with the specified name exists for the project



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
	projectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Project internal (UUID) or global (integer) identifier
	name := "name_example" // string | TestPlan name to check

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectTestPlansAPI.ApiV2ProjectsProjectIdTestPlansNameExistsGet(context.Background(), projectId, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectTestPlansAPI.ApiV2ProjectsProjectIdTestPlansNameExistsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2ProjectsProjectIdTestPlansNameExistsGet`: bool
	fmt.Fprintf(os.Stdout, "Response from `ProjectTestPlansAPI.ApiV2ProjectsProjectIdTestPlansNameExistsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project internal (UUID) or global (integer) identifier | 
**name** | **string** | TestPlan name to check | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ProjectsProjectIdTestPlansNameExistsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**bool**

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2ProjectsProjectIdTestPlansPurgeBulkPost

> ApiV2ProjectsProjectIdTestPlansPurgeBulkPost(ctx, projectId).TestPlanSelectModel(testPlanSelectModel).Execute()

Permanently delete multiple archived test plans

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
	testPlanSelectModel := *openapiclient.NewTestPlanSelectModel(*openapiclient.NewProjectTestPlansFilterModel()) // TestPlanSelectModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectTestPlansAPI.ApiV2ProjectsProjectIdTestPlansPurgeBulkPost(context.Background(), projectId).TestPlanSelectModel(testPlanSelectModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectTestPlansAPI.ApiV2ProjectsProjectIdTestPlansPurgeBulkPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Unique or global ID of the project | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ProjectsProjectIdTestPlansPurgeBulkPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **testPlanSelectModel** | [**TestPlanSelectModel**](TestPlanSelectModel.md) |  | 

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


## ApiV2ProjectsProjectIdTestPlansRestoreBulkPost

> []string ApiV2ProjectsProjectIdTestPlansRestoreBulkPost(ctx, projectId).TestPlanSelectModel(testPlanSelectModel).Execute()

Restore multiple test plans

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
	testPlanSelectModel := *openapiclient.NewTestPlanSelectModel(*openapiclient.NewProjectTestPlansFilterModel()) // TestPlanSelectModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectTestPlansAPI.ApiV2ProjectsProjectIdTestPlansRestoreBulkPost(context.Background(), projectId).TestPlanSelectModel(testPlanSelectModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectTestPlansAPI.ApiV2ProjectsProjectIdTestPlansRestoreBulkPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2ProjectsProjectIdTestPlansRestoreBulkPost`: []string
	fmt.Fprintf(os.Stdout, "Response from `ProjectTestPlansAPI.ApiV2ProjectsProjectIdTestPlansRestoreBulkPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Unique or global ID of the project | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ProjectsProjectIdTestPlansRestoreBulkPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **testPlanSelectModel** | [**TestPlanSelectModel**](TestPlanSelectModel.md) |  | 

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


## ApiV2ProjectsProjectIdTestPlansSearchPost

> []TestPlanWithAnalyticModel ApiV2ProjectsProjectIdTestPlansSearchPost(ctx, projectId).MustUpdateCache(mustUpdateCache).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).ProjectTestPlansFilterModel(projectTestPlansFilterModel).Execute()

Get Project TestPlans with analytics



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
	mustUpdateCache := true // bool |  (optional) (default to false)
	skip := int32(56) // int32 | Amount of items to be skipped (offset) (optional)
	take := int32(56) // int32 | Amount of items to be taken (limit) (optional)
	orderBy := "orderBy_example" // string | SQL-like  ORDER BY statement (column1 ASC|DESC , column2 ASC|DESC) (optional)
	searchField := "searchField_example" // string | Property name for searching (optional)
	searchValue := "searchValue_example" // string | Value for searching (optional)
	projectTestPlansFilterModel := *openapiclient.NewProjectTestPlansFilterModel() // ProjectTestPlansFilterModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectTestPlansAPI.ApiV2ProjectsProjectIdTestPlansSearchPost(context.Background(), projectId).MustUpdateCache(mustUpdateCache).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).ProjectTestPlansFilterModel(projectTestPlansFilterModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectTestPlansAPI.ApiV2ProjectsProjectIdTestPlansSearchPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2ProjectsProjectIdTestPlansSearchPost`: []TestPlanWithAnalyticModel
	fmt.Fprintf(os.Stdout, "Response from `ProjectTestPlansAPI.ApiV2ProjectsProjectIdTestPlansSearchPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project internal (UUID) or global (integer) identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ProjectsProjectIdTestPlansSearchPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mustUpdateCache** | **bool** |  | [default to false]
 **skip** | **int32** | Amount of items to be skipped (offset) | 
 **take** | **int32** | Amount of items to be taken (limit) | 
 **orderBy** | **string** | SQL-like  ORDER BY statement (column1 ASC|DESC , column2 ASC|DESC) | 
 **searchField** | **string** | Property name for searching | 
 **searchValue** | **string** | Value for searching | 
 **projectTestPlansFilterModel** | [**ProjectTestPlansFilterModel**](ProjectTestPlansFilterModel.md) |  | 

### Return type

[**[]TestPlanWithAnalyticModel**](TestPlanWithAnalyticModel.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

