# \ProjectsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddGlobaAttributesToProject**](ProjectsApi.md#AddGlobaAttributesToProject) | **Post** /api/v2/projects/{id}/globalAttributes | Add global attributes to project
[**ApiV2ProjectsDemoPost**](ProjectsApi.md#ApiV2ProjectsDemoPost) | **Post** /api/v2/projects/demo | 
[**ApiV2ProjectsIdDelete**](ProjectsApi.md#ApiV2ProjectsIdDelete) | **Delete** /api/v2/projects/{id} | Archive project
[**ApiV2ProjectsIdFailureClassesGet**](ProjectsApi.md#ApiV2ProjectsIdFailureClassesGet) | **Get** /api/v2/projects/{id}/failureClasses | Get failure classes
[**ApiV2ProjectsIdFavoritePut**](ProjectsApi.md#ApiV2ProjectsIdFavoritePut) | **Put** /api/v2/projects/{id}/favorite | Mark Project as favorite
[**ApiV2ProjectsIdFiltersGet**](ProjectsApi.md#ApiV2ProjectsIdFiltersGet) | **Get** /api/v2/projects/{id}/filters | Get Project filters
[**ApiV2ProjectsIdPatch**](ProjectsApi.md#ApiV2ProjectsIdPatch) | **Patch** /api/v2/projects/{id} | Patch project
[**ApiV2ProjectsIdPurgePost**](ProjectsApi.md#ApiV2ProjectsIdPurgePost) | **Post** /api/v2/projects/{id}/purge | Purge the project
[**ApiV2ProjectsIdRestorePost**](ProjectsApi.md#ApiV2ProjectsIdRestorePost) | **Post** /api/v2/projects/{id}/restore | Restore archived project
[**ApiV2ProjectsIdTestPlansAttributeAttributeIdDelete**](ProjectsApi.md#ApiV2ProjectsIdTestPlansAttributeAttributeIdDelete) | **Delete** /api/v2/projects/{id}/testPlans/attribute/{attributeId} | Delete attribute from project&#39;s test plans
[**ApiV2ProjectsIdTestPlansAttributePut**](ProjectsApi.md#ApiV2ProjectsIdTestPlansAttributePut) | **Put** /api/v2/projects/{id}/testPlans/attribute | Update attribute of project&#39;s test plans
[**ApiV2ProjectsIdTestRunsActiveGet**](ProjectsApi.md#ApiV2ProjectsIdTestRunsActiveGet) | **Get** /api/v2/projects/{id}/testRuns/active | Get active Project TestRuns
[**ApiV2ProjectsIdTestRunsFullGet**](ProjectsApi.md#ApiV2ProjectsIdTestRunsFullGet) | **Get** /api/v2/projects/{id}/testRuns/full | Get Project TestRuns full models
[**ApiV2ProjectsNameNameExistsGet**](ProjectsApi.md#ApiV2ProjectsNameNameExistsGet) | **Get** /api/v2/projects/name/{name}/exists | 
[**ApiV2ProjectsPurgeBulkPost**](ProjectsApi.md#ApiV2ProjectsPurgeBulkPost) | **Post** /api/v2/projects/purge/bulk | Purge multiple projects
[**ApiV2ProjectsRestoreBulkPost**](ProjectsApi.md#ApiV2ProjectsRestoreBulkPost) | **Post** /api/v2/projects/restore/bulk | Restore multiple projects
[**ApiV2ProjectsSearchPost**](ProjectsApi.md#ApiV2ProjectsSearchPost) | **Post** /api/v2/projects/search | Search for projects
[**BackgroundImportProject**](ProjectsApi.md#BackgroundImportProject) | **Post** /api/v2/projects/import/json | Import project from JSON file in background job
[**BackgroundImportZipProject**](ProjectsApi.md#BackgroundImportZipProject) | **Post** /api/v2/projects/import/zip | Import project from Zip file in background job
[**CallImport**](ProjectsApi.md#CallImport) | **Post** /api/v2/projects/import | Import project from JSON file
[**CreateProject**](ProjectsApi.md#CreateProject) | **Post** /api/v2/projects | Create project
[**DeleteProjectAutoTests**](ProjectsApi.md#DeleteProjectAutoTests) | **Delete** /api/v2/projects/{id}/autoTests | Delete all autotests from project
[**GetAllProjects**](ProjectsApi.md#GetAllProjects) | **Get** /api/v2/projects | Get all projects
[**GetAutoTestsNamespaces**](ProjectsApi.md#GetAutoTestsNamespaces) | **Get** /api/v2/projects/{id}/autoTestsNamespaces | Get namespaces of autotests in project
[**GetProjectById**](ProjectsApi.md#GetProjectById) | **Get** /api/v2/projects/{id} | Get project by ID
[**GetTestPlansByProjectId**](ProjectsApi.md#GetTestPlansByProjectId) | **Get** /api/v2/projects/{id}/testPlans | Get project test plans
[**GetTestRunsByProjectId**](ProjectsApi.md#GetTestRunsByProjectId) | **Get** /api/v2/projects/{id}/testRuns | Get project test runs
[**UpdateProject**](ProjectsApi.md#UpdateProject) | **Put** /api/v2/projects | Update project



## AddGlobaAttributesToProject

> AddGlobaAttributesToProject(ctx, id).RequestBody(requestBody).Execute()

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
    r, err := apiClient.ProjectsApi.AddGlobaAttributesToProject(context.Background(), id).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.AddGlobaAttributesToProject``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAddGlobaAttributesToProjectRequest struct via the builder pattern


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


## ApiV2ProjectsDemoPost

> ProjectModel ApiV2ProjectsDemoPost(ctx).CreateProjectRequest(createProjectRequest).Execute()



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
    createProjectRequest := *openapiclient.NewCreateProjectRequest("Name_example") // CreateProjectRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.ApiV2ProjectsDemoPost(context.Background()).CreateProjectRequest(createProjectRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ApiV2ProjectsDemoPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV2ProjectsDemoPost`: ProjectModel
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.ApiV2ProjectsDemoPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ProjectsDemoPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createProjectRequest** | [**CreateProjectRequest**](CreateProjectRequest.md) |  | 

### Return type

[**ProjectModel**](ProjectModel.md)

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
    r, err := apiClient.ProjectsApi.ApiV2ProjectsIdDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ApiV2ProjectsIdDelete``: %v\n", err)
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

> []FailureClassModel ApiV2ProjectsIdFailureClassesGet(ctx, id).IsDeleted(isDeleted).Execute()

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
    resp, r, err := apiClient.ProjectsApi.ApiV2ProjectsIdFailureClassesGet(context.Background(), id).IsDeleted(isDeleted).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ApiV2ProjectsIdFailureClassesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV2ProjectsIdFailureClassesGet`: []FailureClassModel
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.ApiV2ProjectsIdFailureClassesGet`: %v\n", resp)
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

[**[]FailureClassModel**](FailureClassModel.md)

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
    r, err := apiClient.ProjectsApi.ApiV2ProjectsIdFavoritePut(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ApiV2ProjectsIdFavoritePut``: %v\n", err)
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
    resp, r, err := apiClient.ProjectsApi.ApiV2ProjectsIdFiltersGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ApiV2ProjectsIdFiltersGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV2ProjectsIdFiltersGet`: []FilterModel
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.ApiV2ProjectsIdFiltersGet`: %v\n", resp)
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
    r, err := apiClient.ProjectsApi.ApiV2ProjectsIdPatch(context.Background(), id).Operation(operation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ApiV2ProjectsIdPatch``: %v\n", err)
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
    r, err := apiClient.ProjectsApi.ApiV2ProjectsIdPurgePost(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ApiV2ProjectsIdPurgePost``: %v\n", err)
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
    r, err := apiClient.ProjectsApi.ApiV2ProjectsIdRestorePost(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ApiV2ProjectsIdRestorePost``: %v\n", err)
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
    r, err := apiClient.ProjectsApi.ApiV2ProjectsIdTestPlansAttributeAttributeIdDelete(context.Background(), id, attributeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ApiV2ProjectsIdTestPlansAttributeAttributeIdDelete``: %v\n", err)
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

> ApiV2ProjectsIdTestPlansAttributePut(ctx, id).UpdateCustomAttributeTestPlanProjectRelationsRequest(updateCustomAttributeTestPlanProjectRelationsRequest).Execute()

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
    updateCustomAttributeTestPlanProjectRelationsRequest := *openapiclient.NewUpdateCustomAttributeTestPlanProjectRelationsRequest("Id_example", false, false) // UpdateCustomAttributeTestPlanProjectRelationsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ProjectsApi.ApiV2ProjectsIdTestPlansAttributePut(context.Background(), id).UpdateCustomAttributeTestPlanProjectRelationsRequest(updateCustomAttributeTestPlanProjectRelationsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ApiV2ProjectsIdTestPlansAttributePut``: %v\n", err)
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

 **updateCustomAttributeTestPlanProjectRelationsRequest** | [**UpdateCustomAttributeTestPlanProjectRelationsRequest**](UpdateCustomAttributeTestPlanProjectRelationsRequest.md) |  | 

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


## ApiV2ProjectsIdTestRunsActiveGet

> []PublicTestRunModel ApiV2ProjectsIdTestRunsActiveGet(ctx, id).Execute()

Get active Project TestRuns



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
    resp, r, err := apiClient.ProjectsApi.ApiV2ProjectsIdTestRunsActiveGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ApiV2ProjectsIdTestRunsActiveGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV2ProjectsIdTestRunsActiveGet`: []PublicTestRunModel
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.ApiV2ProjectsIdTestRunsActiveGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Project internal (UUID) or global (integer) identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ProjectsIdTestRunsActiveGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]PublicTestRunModel**](PublicTestRunModel.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2ProjectsIdTestRunsFullGet

> []TestRunModel ApiV2ProjectsIdTestRunsFullGet(ctx, id).IncludeTestResults(includeTestResults).MustAggregateTestResults(mustAggregateTestResults).NotStarted(notStarted).InProgress(inProgress).Stopped(stopped).Completed(completed).CreatedDateFrom(createdDateFrom).CreatedDateTo(createdDateTo).TestPlanId(testPlanId).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).Execute()

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
    includeTestResults := true // bool |  (optional) (default to false)
    mustAggregateTestResults := true // bool |  (optional) (default to true)
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
    resp, r, err := apiClient.ProjectsApi.ApiV2ProjectsIdTestRunsFullGet(context.Background(), id).IncludeTestResults(includeTestResults).MustAggregateTestResults(mustAggregateTestResults).NotStarted(notStarted).InProgress(inProgress).Stopped(stopped).Completed(completed).CreatedDateFrom(createdDateFrom).CreatedDateTo(createdDateTo).TestPlanId(testPlanId).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ApiV2ProjectsIdTestRunsFullGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV2ProjectsIdTestRunsFullGet`: []TestRunModel
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.ApiV2ProjectsIdTestRunsFullGet`: %v\n", resp)
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

 **includeTestResults** | **bool** |  | [default to false]
 **mustAggregateTestResults** | **bool** |  | [default to true]
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

[**[]TestRunModel**](TestRunModel.md)

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
    resp, r, err := apiClient.ProjectsApi.ApiV2ProjectsNameNameExistsGet(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ApiV2ProjectsNameNameExistsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV2ProjectsNameNameExistsGet`: bool
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.ApiV2ProjectsNameNameExistsGet`: %v\n", resp)
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

> int64 ApiV2ProjectsPurgeBulkPost(ctx).ApiV2ProjectsRestoreBulkPostRequest(apiV2ProjectsRestoreBulkPostRequest).Execute()

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
    apiV2ProjectsRestoreBulkPostRequest := *openapiclient.NewApiV2ProjectsRestoreBulkPostRequest() // ApiV2ProjectsRestoreBulkPostRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.ApiV2ProjectsPurgeBulkPost(context.Background()).ApiV2ProjectsRestoreBulkPostRequest(apiV2ProjectsRestoreBulkPostRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ApiV2ProjectsPurgeBulkPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV2ProjectsPurgeBulkPost`: int64
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.ApiV2ProjectsPurgeBulkPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ProjectsPurgeBulkPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiV2ProjectsRestoreBulkPostRequest** | [**ApiV2ProjectsRestoreBulkPostRequest**](ApiV2ProjectsRestoreBulkPostRequest.md) |  | 

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

> int64 ApiV2ProjectsRestoreBulkPost(ctx).ApiV2ProjectsRestoreBulkPostRequest(apiV2ProjectsRestoreBulkPostRequest).Execute()

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
    apiV2ProjectsRestoreBulkPostRequest := *openapiclient.NewApiV2ProjectsRestoreBulkPostRequest() // ApiV2ProjectsRestoreBulkPostRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.ApiV2ProjectsRestoreBulkPost(context.Background()).ApiV2ProjectsRestoreBulkPostRequest(apiV2ProjectsRestoreBulkPostRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ApiV2ProjectsRestoreBulkPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV2ProjectsRestoreBulkPost`: int64
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.ApiV2ProjectsRestoreBulkPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ProjectsRestoreBulkPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiV2ProjectsRestoreBulkPostRequest** | [**ApiV2ProjectsRestoreBulkPostRequest**](ApiV2ProjectsRestoreBulkPostRequest.md) |  | 

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

> []ProjectModel ApiV2ProjectsSearchPost(ctx).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).ApiV2ProjectsSearchPostRequest(apiV2ProjectsSearchPostRequest).Execute()

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
    apiV2ProjectsSearchPostRequest := *openapiclient.NewApiV2ProjectsSearchPostRequest() // ApiV2ProjectsSearchPostRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.ApiV2ProjectsSearchPost(context.Background()).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).ApiV2ProjectsSearchPostRequest(apiV2ProjectsSearchPostRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ApiV2ProjectsSearchPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV2ProjectsSearchPost`: []ProjectModel
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.ApiV2ProjectsSearchPost`: %v\n", resp)
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
 **apiV2ProjectsSearchPostRequest** | [**ApiV2ProjectsSearchPostRequest**](ApiV2ProjectsSearchPostRequest.md) |  | 

### Return type

[**[]ProjectModel**](ProjectModel.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BackgroundImportProject

> string BackgroundImportProject(ctx).File(file).Execute()

Import project from JSON file in background job

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
    file := os.NewFile(1234, "some_file") // *os.File |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.BackgroundImportProject(context.Background()).File(file).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.BackgroundImportProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BackgroundImportProject`: string
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.BackgroundImportProject`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBackgroundImportProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | ***os.File** |  | 

### Return type

**string**

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: multipart/form-data, application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BackgroundImportZipProject

> string BackgroundImportZipProject(ctx).File(file).Execute()

Import project from Zip file in background job

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
    file := os.NewFile(1234, "some_file") // *os.File |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.BackgroundImportZipProject(context.Background()).File(file).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.BackgroundImportZipProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BackgroundImportZipProject`: string
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.BackgroundImportZipProject`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBackgroundImportZipProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | ***os.File** |  | 

### Return type

**string**

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: multipart/form-data, application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CallImport

> CallImport(ctx).IncludeAttachments(includeAttachments).File(file).Execute()

Import project from JSON file



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
    includeAttachments := true // bool | Enables attachment import. (optional) (default to false)
    file := os.NewFile(1234, "some_file") // *os.File | Select file (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ProjectsApi.CallImport(context.Background()).IncludeAttachments(includeAttachments).File(file).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.CallImport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCallImportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **includeAttachments** | **bool** | Enables attachment import. | [default to false]
 **file** | ***os.File** | Select file | 

### Return type

 (empty response body)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateProject

> ProjectModel CreateProject(ctx).CreateProjectRequest(createProjectRequest).Execute()

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
    createProjectRequest := *openapiclient.NewCreateProjectRequest("Name_example") // CreateProjectRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.CreateProject(context.Background()).CreateProjectRequest(createProjectRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.CreateProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateProject`: ProjectModel
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.CreateProject`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createProjectRequest** | [**CreateProjectRequest**](CreateProjectRequest.md) |  | 

### Return type

[**ProjectModel**](ProjectModel.md)

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
    r, err := apiClient.ProjectsApi.DeleteProjectAutoTests(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.DeleteProjectAutoTests``: %v\n", err)
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

> []ProjectModel GetAllProjects(ctx).IsDeleted(isDeleted).ProjectName(projectName).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).Execute()

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
    resp, r, err := apiClient.ProjectsApi.GetAllProjects(context.Background()).IsDeleted(isDeleted).ProjectName(projectName).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.GetAllProjects``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllProjects`: []ProjectModel
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.GetAllProjects`: %v\n", resp)
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

[**[]ProjectModel**](ProjectModel.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAutoTestsNamespaces

> []AutoTestNamespaceModel GetAutoTestsNamespaces(ctx, id).Execute()

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
    resp, r, err := apiClient.ProjectsApi.GetAutoTestsNamespaces(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.GetAutoTestsNamespaces``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAutoTestsNamespaces`: []AutoTestNamespaceModel
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.GetAutoTestsNamespaces`: %v\n", resp)
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

[**[]AutoTestNamespaceModel**](AutoTestNamespaceModel.md)

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
    resp, r, err := apiClient.ProjectsApi.GetProjectById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.GetProjectById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProjectById`: ProjectModel
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.GetProjectById`: %v\n", resp)
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
    resp, r, err := apiClient.ProjectsApi.GetTestPlansByProjectId(context.Background(), id).IsDeleted(isDeleted).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.GetTestPlansByProjectId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTestPlansByProjectId`: []TestPlanModel
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.GetTestPlansByProjectId`: %v\n", resp)
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

> []TestRunV2GetModel GetTestRunsByProjectId(ctx, id).NotStarted(notStarted).InProgress(inProgress).Stopped(stopped).Completed(completed).CreatedDateFrom(createdDateFrom).CreatedDateTo(createdDateTo).TestPlanId(testPlanId).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).Execute()

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
    resp, r, err := apiClient.ProjectsApi.GetTestRunsByProjectId(context.Background(), id).NotStarted(notStarted).InProgress(inProgress).Stopped(stopped).Completed(completed).CreatedDateFrom(createdDateFrom).CreatedDateTo(createdDateTo).TestPlanId(testPlanId).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.GetTestRunsByProjectId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTestRunsByProjectId`: []TestRunV2GetModel
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.GetTestRunsByProjectId`: %v\n", resp)
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

[**[]TestRunV2GetModel**](TestRunV2GetModel.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateProject

> UpdateProject(ctx).UpdateProjectRequest(updateProjectRequest).Execute()

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
    updateProjectRequest := *openapiclient.NewUpdateProjectRequest("Id_example", "Name_example", openapiclient.ProjectTypeModel("Regular")) // UpdateProjectRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ProjectsApi.UpdateProject(context.Background()).UpdateProjectRequest(updateProjectRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.UpdateProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateProjectRequest** | [**UpdateProjectRequest**](UpdateProjectRequest.md) |  | 

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

