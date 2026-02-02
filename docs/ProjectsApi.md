# \ProjectsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddGlobalAttributesToProject**](ProjectsAPI.md#AddGlobalAttributesToProject) | **Post** /api/v2/projects/{id}/globalAttributes | Add global attributes to project
[**ApiV2ProjectsIdDelete**](ProjectsAPI.md#ApiV2ProjectsIdDelete) | **Delete** /api/v2/projects/{id} | Archive project
[**ApiV2ProjectsIdFailureClassesGet**](ProjectsAPI.md#ApiV2ProjectsIdFailureClassesGet) | **Get** /api/v2/projects/{id}/failureClasses | Get failure classes
[**ApiV2ProjectsIdFavoritePut**](ProjectsAPI.md#ApiV2ProjectsIdFavoritePut) | **Put** /api/v2/projects/{id}/favorite | Mark Project as favorite
[**ApiV2ProjectsIdFiltersGet**](ProjectsAPI.md#ApiV2ProjectsIdFiltersGet) | **Get** /api/v2/projects/{id}/filters | Get Project filters
[**ApiV2ProjectsIdPatch**](ProjectsAPI.md#ApiV2ProjectsIdPatch) | **Patch** /api/v2/projects/{id} | Patch project
[**ApiV2ProjectsIdPurgePost**](ProjectsAPI.md#ApiV2ProjectsIdPurgePost) | **Post** /api/v2/projects/{id}/purge | Purge the project
[**ApiV2ProjectsIdRestorePost**](ProjectsAPI.md#ApiV2ProjectsIdRestorePost) | **Post** /api/v2/projects/{id}/restore | Restore archived project
[**ApiV2ProjectsIdTestPlansAttributeAttributeIdDelete**](ProjectsAPI.md#ApiV2ProjectsIdTestPlansAttributeAttributeIdDelete) | **Delete** /api/v2/projects/{id}/testPlans/attribute/{attributeId} | Delete attribute from project&#39;s test plans
[**ApiV2ProjectsIdTestPlansAttributePut**](ProjectsAPI.md#ApiV2ProjectsIdTestPlansAttributePut) | **Put** /api/v2/projects/{id}/testPlans/attribute | Update attribute of project&#39;s test plans
[**ApiV2ProjectsIdTestRunsFullGet**](ProjectsAPI.md#ApiV2ProjectsIdTestRunsFullGet) | **Get** /api/v2/projects/{id}/testRuns/full | Get Project TestRuns full models
[**ApiV2ProjectsNameNameExistsGet**](ProjectsAPI.md#ApiV2ProjectsNameNameExistsGet) | **Get** /api/v2/projects/name/{name}/exists | 
[**ApiV2ProjectsPurgeBulkPost**](ProjectsAPI.md#ApiV2ProjectsPurgeBulkPost) | **Post** /api/v2/projects/purge/bulk | Purge multiple projects
[**ApiV2ProjectsRestoreBulkPost**](ProjectsAPI.md#ApiV2ProjectsRestoreBulkPost) | **Post** /api/v2/projects/restore/bulk | Restore multiple projects
[**ApiV2ProjectsSearchPost**](ProjectsAPI.md#ApiV2ProjectsSearchPost) | **Post** /api/v2/projects/search | Search for projects
[**ApiV2ProjectsShortsPost**](ProjectsAPI.md#ApiV2ProjectsShortsPost) | **Post** /api/v2/projects/shorts | Get projects short models
[**CreateProject**](ProjectsAPI.md#CreateProject) | **Post** /api/v2/projects | Create project
[**DeleteProjectAutoTests**](ProjectsAPI.md#DeleteProjectAutoTests) | **Delete** /api/v2/projects/{id}/autoTests | Delete all autotests from project
[**GetAllProjects**](ProjectsAPI.md#GetAllProjects) | **Get** /api/v2/projects | Get all projects
[**GetAutoTestsNamespaces**](ProjectsAPI.md#GetAutoTestsNamespaces) | **Get** /api/v2/projects/{id}/autoTestsNamespaces | Get namespaces of autotests in project
[**GetProjectById**](ProjectsAPI.md#GetProjectById) | **Get** /api/v2/projects/{id} | Get project by ID
[**GetTestPlansByProjectId**](ProjectsAPI.md#GetTestPlansByProjectId) | **Get** /api/v2/projects/{id}/testPlans | Get project test plans
[**GetTestRunsByProjectId**](ProjectsAPI.md#GetTestRunsByProjectId) | **Get** /api/v2/projects/{id}/testRuns | Get project test runs
[**UpdateProject**](ProjectsAPI.md#UpdateProject) | **Put** /api/v2/projects | Update project



## AddGlobalAttributesToProject

> AddGlobalAttributesToProject(ctx, id).RequestBody(requestBody).Execute()

Add global attributes to project



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
	id := "id_example" // string | Project internal (UUID) or global (integer) identifier
	requestBody := []string{"Property_example"} // []string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectsAPI.AddGlobalAttributesToProject(context.Background(), id).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.AddGlobalAttributesToProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Project internal (UUID) or global (integer) identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddGlobalAttributesToProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestBody** | **[]string** |  | 

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


## ApiV2ProjectsIdDelete

> ApiV2ProjectsIdDelete(ctx, id).Execute()

Archive project

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
	id := "id_example" // string | Unique or global ID of the project

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectsAPI.ApiV2ProjectsIdDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.ApiV2ProjectsIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Unique or global ID of the project | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ProjectsIdDeleteRequest struct via the builder pattern


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


## ApiV2ProjectsIdFailureClassesGet

> []FailureCategoryApiResult ApiV2ProjectsIdFailureClassesGet(ctx, id).IsDeleted(isDeleted).Execute()

Get failure classes

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
	id := "id_example" // string | Unique or global ID of the project
	isDeleted := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.ApiV2ProjectsIdFailureClassesGet(context.Background(), id).IsDeleted(isDeleted).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.ApiV2ProjectsIdFailureClassesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2ProjectsIdFailureClassesGet`: []FailureCategoryApiResult
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.ApiV2ProjectsIdFailureClassesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Unique or global ID of the project | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ProjectsIdFailureClassesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **isDeleted** | **bool** |  | 

### Return type

[**[]FailureCategoryApiResult**](FailureCategoryApiResult.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2ProjectsIdFavoritePut

> ApiV2ProjectsIdFavoritePut(ctx, id).Execute()

Mark Project as favorite

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
	id := "id_example" // string | Project internal (UUID) or global (integer) identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectsAPI.ApiV2ProjectsIdFavoritePut(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.ApiV2ProjectsIdFavoritePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Project internal (UUID) or global (integer) identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ProjectsIdFavoritePutRequest struct via the builder pattern


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


## ApiV2ProjectsIdFiltersGet

> []FilterModel ApiV2ProjectsIdFiltersGet(ctx, id).Execute()

Get Project filters



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
	id := "id_example" // string | Project internal (UUID) or global (integer) identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.ApiV2ProjectsIdFiltersGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.ApiV2ProjectsIdFiltersGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2ProjectsIdFiltersGet`: []FilterModel
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.ApiV2ProjectsIdFiltersGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Project internal (UUID) or global (integer) identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ProjectsIdFiltersGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]FilterModel**](FilterModel.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2ProjectsIdPatch

> ApiV2ProjectsIdPatch(ctx, id).Operation(operation).Execute()

Patch project



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique or global Id of project
	operation := []openapiclient.Operation{*openapiclient.NewOperation()} // []Operation |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectsAPI.ApiV2ProjectsIdPatch(context.Background(), id).Operation(operation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.ApiV2ProjectsIdPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Unique or global Id of project | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ProjectsIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **operation** | [**[]Operation**](Operation.md) |  | 

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


## ApiV2ProjectsIdPurgePost

> ApiV2ProjectsIdPurgePost(ctx, id).Execute()

Purge the project

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
	id := "id_example" // string | Unique or global ID of the project

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectsAPI.ApiV2ProjectsIdPurgePost(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.ApiV2ProjectsIdPurgePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Unique or global ID of the project | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ProjectsIdPurgePostRequest struct via the builder pattern


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


## ApiV2ProjectsIdRestorePost

> ApiV2ProjectsIdRestorePost(ctx, id).Execute()

Restore archived project

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
	id := "id_example" // string | Unique or global ID of the project

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectsAPI.ApiV2ProjectsIdRestorePost(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.ApiV2ProjectsIdRestorePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Unique or global ID of the project | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ProjectsIdRestorePostRequest struct via the builder pattern


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


## ApiV2ProjectsIdTestPlansAttributeAttributeIdDelete

> ApiV2ProjectsIdTestPlansAttributeAttributeIdDelete(ctx, id, attributeId).Execute()

Delete attribute from project's test plans



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
	id := "id_example" // string | Project internal (UUID) or global (integer) identifier
	attributeId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectsAPI.ApiV2ProjectsIdTestPlansAttributeAttributeIdDelete(context.Background(), id, attributeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.ApiV2ProjectsIdTestPlansAttributeAttributeIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Project internal (UUID) or global (integer) identifier | 
**attributeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ProjectsIdTestPlansAttributeAttributeIdDeleteRequest struct via the builder pattern


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


## ApiV2ProjectsIdTestPlansAttributePut

> ApiV2ProjectsIdTestPlansAttributePut(ctx, id).CustomAttributeTestPlanProjectRelationPutModel(customAttributeTestPlanProjectRelationPutModel).Execute()

Update attribute of project's test plans



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
	id := "id_example" // string | Project internal (UUID) or global (integer) identifier
	customAttributeTestPlanProjectRelationPutModel := *openapiclient.NewCustomAttributeTestPlanProjectRelationPutModel("Id_example", false, false) // CustomAttributeTestPlanProjectRelationPutModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectsAPI.ApiV2ProjectsIdTestPlansAttributePut(context.Background(), id).CustomAttributeTestPlanProjectRelationPutModel(customAttributeTestPlanProjectRelationPutModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.ApiV2ProjectsIdTestPlansAttributePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Project internal (UUID) or global (integer) identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ProjectsIdTestPlansAttributePutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **customAttributeTestPlanProjectRelationPutModel** | [**CustomAttributeTestPlanProjectRelationPutModel**](CustomAttributeTestPlanProjectRelationPutModel.md) |  | 

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


## ApiV2ProjectsIdTestRunsFullGet

> []TestRunApiResult ApiV2ProjectsIdTestRunsFullGet(ctx, id).IncludeTestResults(includeTestResults).MustAggregateTestResults(mustAggregateTestResults).NotStarted(notStarted).InProgress(inProgress).Stopped(stopped).Completed(completed).CreatedDateFrom(createdDateFrom).CreatedDateTo(createdDateTo).TestPlanId(testPlanId).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).Execute()

Get Project TestRuns full models



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "id_example" // string | Project internal (UUID) or global (integer) identifier
	includeTestResults := true // bool |  (optional)
	mustAggregateTestResults := true // bool |  (optional)
	notStarted := true // bool |  (optional)
	inProgress := true // bool |  (optional)
	stopped := true // bool |  (optional)
	completed := true // bool |  (optional)
	createdDateFrom := time.Now() // time.Time |  (optional)
	createdDateTo := time.Now() // time.Time |  (optional)
	testPlanId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	skip := int32(56) // int32 | Amount of items to be skipped (offset) (optional)
	take := int32(56) // int32 | Amount of items to be taken (limit) (optional)
	orderBy := "orderBy_example" // string | SQL-like  ORDER BY statement (column1 ASC|DESC , column2 ASC|DESC) (optional)
	searchField := "searchField_example" // string | Property name for searching (optional)
	searchValue := "searchValue_example" // string | Value for searching (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.ApiV2ProjectsIdTestRunsFullGet(context.Background(), id).IncludeTestResults(includeTestResults).MustAggregateTestResults(mustAggregateTestResults).NotStarted(notStarted).InProgress(inProgress).Stopped(stopped).Completed(completed).CreatedDateFrom(createdDateFrom).CreatedDateTo(createdDateTo).TestPlanId(testPlanId).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.ApiV2ProjectsIdTestRunsFullGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2ProjectsIdTestRunsFullGet`: []TestRunApiResult
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.ApiV2ProjectsIdTestRunsFullGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Project internal (UUID) or global (integer) identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ProjectsIdTestRunsFullGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeTestResults** | **bool** |  | 
 **mustAggregateTestResults** | **bool** |  | 
 **notStarted** | **bool** |  | 
 **inProgress** | **bool** |  | 
 **stopped** | **bool** |  | 
 **completed** | **bool** |  | 
 **createdDateFrom** | **time.Time** |  | 
 **createdDateTo** | **time.Time** |  | 
 **testPlanId** | **string** |  | 
 **skip** | **int32** | Amount of items to be skipped (offset) | 
 **take** | **int32** | Amount of items to be taken (limit) | 
 **orderBy** | **string** | SQL-like  ORDER BY statement (column1 ASC|DESC , column2 ASC|DESC) | 
 **searchField** | **string** | Property name for searching | 
 **searchValue** | **string** | Value for searching | 

### Return type

[**[]TestRunApiResult**](TestRunApiResult.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2ProjectsNameNameExistsGet

> bool ApiV2ProjectsNameNameExistsGet(ctx, name).Execute()



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
	name := "name_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.ApiV2ProjectsNameNameExistsGet(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.ApiV2ProjectsNameNameExistsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2ProjectsNameNameExistsGet`: bool
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.ApiV2ProjectsNameNameExistsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ProjectsNameNameExistsGetRequest struct via the builder pattern


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


## ApiV2ProjectsPurgeBulkPost

> int64 ApiV2ProjectsPurgeBulkPost(ctx).ProjectSelectModel(projectSelectModel).Execute()

Purge multiple projects

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
	projectSelectModel := *openapiclient.NewProjectSelectModel() // ProjectSelectModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.ApiV2ProjectsPurgeBulkPost(context.Background()).ProjectSelectModel(projectSelectModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.ApiV2ProjectsPurgeBulkPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2ProjectsPurgeBulkPost`: int64
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.ApiV2ProjectsPurgeBulkPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ProjectsPurgeBulkPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectSelectModel** | [**ProjectSelectModel**](ProjectSelectModel.md) |  | 

### Return type

**int64**

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2ProjectsRestoreBulkPost

> int64 ApiV2ProjectsRestoreBulkPost(ctx).ProjectSelectModel(projectSelectModel).Execute()

Restore multiple projects

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
	projectSelectModel := *openapiclient.NewProjectSelectModel() // ProjectSelectModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.ApiV2ProjectsRestoreBulkPost(context.Background()).ProjectSelectModel(projectSelectModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.ApiV2ProjectsRestoreBulkPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2ProjectsRestoreBulkPost`: int64
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.ApiV2ProjectsRestoreBulkPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ProjectsRestoreBulkPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectSelectModel** | [**ProjectSelectModel**](ProjectSelectModel.md) |  | 

### Return type

**int64**

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2ProjectsSearchPost

> []ProjectShortModel ApiV2ProjectsSearchPost(ctx).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).ProjectsFilterModel(projectsFilterModel).Execute()

Search for projects

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
	projectsFilterModel := *openapiclient.NewProjectsFilterModel() // ProjectsFilterModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.ApiV2ProjectsSearchPost(context.Background()).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).ProjectsFilterModel(projectsFilterModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.ApiV2ProjectsSearchPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2ProjectsSearchPost`: []ProjectShortModel
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.ApiV2ProjectsSearchPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ProjectsSearchPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skip** | **int32** | Amount of items to be skipped (offset) | 
 **take** | **int32** | Amount of items to be taken (limit) | 
 **orderBy** | **string** | SQL-like  ORDER BY statement (column1 ASC|DESC , column2 ASC|DESC) | 
 **searchField** | **string** | Property name for searching | 
 **searchValue** | **string** | Value for searching | 
 **projectsFilterModel** | [**ProjectsFilterModel**](ProjectsFilterModel.md) |  | 

### Return type

[**[]ProjectShortModel**](ProjectShortModel.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2ProjectsShortsPost

> ProjectShortApiResultReply ApiV2ProjectsShortsPost(ctx).GetShortProjectsApiModel(getShortProjectsApiModel).Execute()

Get projects short models



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
	getShortProjectsApiModel := *openapiclient.NewGetShortProjectsApiModel() // GetShortProjectsApiModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.ApiV2ProjectsShortsPost(context.Background()).GetShortProjectsApiModel(getShortProjectsApiModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.ApiV2ProjectsShortsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2ProjectsShortsPost`: ProjectShortApiResultReply
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.ApiV2ProjectsShortsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ProjectsShortsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getShortProjectsApiModel** | [**GetShortProjectsApiModel**](GetShortProjectsApiModel.md) |  | 

### Return type

[**ProjectShortApiResultReply**](ProjectShortApiResultReply.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateProject

> ProjectApiResult CreateProject(ctx).CreateProjectApiModel(createProjectApiModel).Execute()

Create project



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
	createProjectApiModel := *openapiclient.NewCreateProjectApiModel("Name_example") // CreateProjectApiModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.CreateProject(context.Background()).CreateProjectApiModel(createProjectApiModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.CreateProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateProject`: ProjectApiResult
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.CreateProject`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createProjectApiModel** | [**CreateProjectApiModel**](CreateProjectApiModel.md) |  | 

### Return type

[**ProjectApiResult**](ProjectApiResult.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteProjectAutoTests

> DeleteProjectAutoTests(ctx, id).Execute()

Delete all autotests from project

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
	id := "id_example" // string | Unique or global ID of the project

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectsAPI.DeleteProjectAutoTests(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.DeleteProjectAutoTests``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Unique or global ID of the project | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProjectAutoTestsRequest struct via the builder pattern


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


## GetAllProjects

> []ProjectShortModel GetAllProjects(ctx).IsDeleted(isDeleted).ProjectName(projectName).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).Execute()

Get all projects



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
	isDeleted := true // bool | If result must consist of only actual/deleted parameters (optional)
	projectName := "projectName_example" // string |  (optional)
	skip := int32(56) // int32 | Amount of items to be skipped (offset) (optional)
	take := int32(56) // int32 | Amount of items to be taken (limit) (optional)
	orderBy := "orderBy_example" // string | SQL-like  ORDER BY statement (column1 ASC|DESC , column2 ASC|DESC) (optional)
	searchField := "searchField_example" // string | Property name for searching (optional)
	searchValue := "searchValue_example" // string | Value for searching (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.GetAllProjects(context.Background()).IsDeleted(isDeleted).ProjectName(projectName).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.GetAllProjects``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllProjects`: []ProjectShortModel
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.GetAllProjects`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllProjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **isDeleted** | **bool** | If result must consist of only actual/deleted parameters | 
 **projectName** | **string** |  | 
 **skip** | **int32** | Amount of items to be skipped (offset) | 
 **take** | **int32** | Amount of items to be taken (limit) | 
 **orderBy** | **string** | SQL-like  ORDER BY statement (column1 ASC|DESC , column2 ASC|DESC) | 
 **searchField** | **string** | Property name for searching | 
 **searchValue** | **string** | Value for searching | 

### Return type

[**[]ProjectShortModel**](ProjectShortModel.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAutoTestsNamespaces

> []AutoTestNamespaceApiResult GetAutoTestsNamespaces(ctx, id).Execute()

Get namespaces of autotests in project



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
	id := "id_example" // string | Project internal (UUID) or global (integer) identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.GetAutoTestsNamespaces(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.GetAutoTestsNamespaces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAutoTestsNamespaces`: []AutoTestNamespaceApiResult
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.GetAutoTestsNamespaces`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Project internal (UUID) or global (integer) identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAutoTestsNamespacesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]AutoTestNamespaceApiResult**](AutoTestNamespaceApiResult.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectById

> ProjectModel GetProjectById(ctx, id).Execute()

Get project by ID



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
	id := "id_example" // string | Project internal (UUID) or global (integer) identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.GetProjectById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.GetProjectById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectById`: ProjectModel
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.GetProjectById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Project internal (UUID) or global (integer) identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProjectModel**](ProjectModel.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTestPlansByProjectId

> []TestPlanModel GetTestPlansByProjectId(ctx, id).IsDeleted(isDeleted).Execute()

Get project test plans



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
	id := "id_example" // string | Project internal (UUID) or global (integer) identifier
	isDeleted := true // bool | If result must consist of only actual/archived test plans (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.GetTestPlansByProjectId(context.Background(), id).IsDeleted(isDeleted).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.GetTestPlansByProjectId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTestPlansByProjectId`: []TestPlanModel
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.GetTestPlansByProjectId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Project internal (UUID) or global (integer) identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTestPlansByProjectIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **isDeleted** | **bool** | If result must consist of only actual/archived test plans | 

### Return type

[**[]TestPlanModel**](TestPlanModel.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTestRunsByProjectId

> []TestRunV2ApiResult GetTestRunsByProjectId(ctx, id).NotStarted(notStarted).InProgress(inProgress).Stopped(stopped).Completed(completed).CreatedDateFrom(createdDateFrom).CreatedDateTo(createdDateTo).TestPlanId(testPlanId).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).Execute()

Get project test runs



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "id_example" // string | Project internal (UUID) or global (integer) identifier
	notStarted := true // bool | 
	inProgress := true // bool | 
	stopped := true // bool | 
	completed := true // bool | 
	createdDateFrom := time.Now() // time.Time |  (optional)
	createdDateTo := time.Now() // time.Time |  (optional)
	testPlanId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	skip := int32(56) // int32 | Amount of items to be skipped (offset) (optional)
	take := int32(56) // int32 | Amount of items to be taken (limit) (optional)
	orderBy := "orderBy_example" // string | SQL-like  ORDER BY statement (column1 ASC|DESC , column2 ASC|DESC) (optional)
	searchField := "searchField_example" // string | Property name for searching (optional)
	searchValue := "searchValue_example" // string | Value for searching (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.GetTestRunsByProjectId(context.Background(), id).NotStarted(notStarted).InProgress(inProgress).Stopped(stopped).Completed(completed).CreatedDateFrom(createdDateFrom).CreatedDateTo(createdDateTo).TestPlanId(testPlanId).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.GetTestRunsByProjectId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTestRunsByProjectId`: []TestRunV2ApiResult
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.GetTestRunsByProjectId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Project internal (UUID) or global (integer) identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTestRunsByProjectIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **notStarted** | **bool** |  | 
 **inProgress** | **bool** |  | 
 **stopped** | **bool** |  | 
 **completed** | **bool** |  | 
 **createdDateFrom** | **time.Time** |  | 
 **createdDateTo** | **time.Time** |  | 
 **testPlanId** | **string** |  | 
 **skip** | **int32** | Amount of items to be skipped (offset) | 
 **take** | **int32** | Amount of items to be taken (limit) | 
 **orderBy** | **string** | SQL-like  ORDER BY statement (column1 ASC|DESC , column2 ASC|DESC) | 
 **searchField** | **string** | Property name for searching | 
 **searchValue** | **string** | Value for searching | 

### Return type

[**[]TestRunV2ApiResult**](TestRunV2ApiResult.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateProject

> UpdateProject(ctx).UpdateProjectApiModel(updateProjectApiModel).Execute()

Update project



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
	updateProjectApiModel := *openapiclient.NewUpdateProjectApiModel("Id_example", "Name_example") // UpdateProjectApiModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectsAPI.UpdateProject(context.Background()).UpdateProjectApiModel(updateProjectApiModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.UpdateProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateProjectApiModel** | [**UpdateProjectApiModel**](UpdateProjectApiModel.md) |  | 

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

