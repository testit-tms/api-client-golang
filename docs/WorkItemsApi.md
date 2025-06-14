# \WorkItemsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV2WorkItemsIdAttachmentsPost**](WorkItemsAPI.md#ApiV2WorkItemsIdAttachmentsPost) | **Post** /api/v2/workItems/{id}/attachments | Upload and link attachment to WorkItem
[**ApiV2WorkItemsIdCheckListTransformToTestCasePost**](WorkItemsAPI.md#ApiV2WorkItemsIdCheckListTransformToTestCasePost) | **Post** /api/v2/workItems/{id}/checkList/transformTo/testCase | Transform CheckList to TestCase
[**ApiV2WorkItemsIdHistoryGet**](WorkItemsAPI.md#ApiV2WorkItemsIdHistoryGet) | **Get** /api/v2/workItems/{id}/history | Get change history of WorkItem
[**ApiV2WorkItemsIdLikeDelete**](WorkItemsAPI.md#ApiV2WorkItemsIdLikeDelete) | **Delete** /api/v2/workItems/{id}/like | Delete like from WorkItem
[**ApiV2WorkItemsIdLikePost**](WorkItemsAPI.md#ApiV2WorkItemsIdLikePost) | **Post** /api/v2/workItems/{id}/like | Set like to WorkItem
[**ApiV2WorkItemsIdLikesCountGet**](WorkItemsAPI.md#ApiV2WorkItemsIdLikesCountGet) | **Get** /api/v2/workItems/{id}/likes/count | Get likes count of WorkItem
[**ApiV2WorkItemsIdLikesGet**](WorkItemsAPI.md#ApiV2WorkItemsIdLikesGet) | **Get** /api/v2/workItems/{id}/likes | Get likes of WorkItem
[**ApiV2WorkItemsIdTestResultsHistoryGet**](WorkItemsAPI.md#ApiV2WorkItemsIdTestResultsHistoryGet) | **Get** /api/v2/workItems/{id}/testResults/history | Get test results history of WorkItem
[**ApiV2WorkItemsIdVersionVersionIdActualPost**](WorkItemsAPI.md#ApiV2WorkItemsIdVersionVersionIdActualPost) | **Post** /api/v2/workItems/{id}/version/{versionId}/actual | Set WorkItem as actual
[**ApiV2WorkItemsLinksUrlsSearchPost**](WorkItemsAPI.md#ApiV2WorkItemsLinksUrlsSearchPost) | **Post** /api/v2/workItems/links/urls/search | 
[**ApiV2WorkItemsMovePost**](WorkItemsAPI.md#ApiV2WorkItemsMovePost) | **Post** /api/v2/workItems/move | Move WorkItem to another section
[**ApiV2WorkItemsSearchPost**](WorkItemsAPI.md#ApiV2WorkItemsSearchPost) | **Post** /api/v2/workItems/search | Search for work items
[**ApiV2WorkItemsSharedStepIdReferencesSectionsPost**](WorkItemsAPI.md#ApiV2WorkItemsSharedStepIdReferencesSectionsPost) | **Post** /api/v2/workItems/{sharedStepId}/references/sections | Get SharedStep references in sections
[**ApiV2WorkItemsSharedStepIdReferencesWorkItemsPost**](WorkItemsAPI.md#ApiV2WorkItemsSharedStepIdReferencesWorkItemsPost) | **Post** /api/v2/workItems/{sharedStepId}/references/workItems | Get SharedStep references in work items
[**ApiV2WorkItemsSharedStepsSharedStepIdReferencesGet**](WorkItemsAPI.md#ApiV2WorkItemsSharedStepsSharedStepIdReferencesGet) | **Get** /api/v2/workItems/sharedSteps/{sharedStepId}/references | Get SharedStep references
[**CreateWorkItem**](WorkItemsAPI.md#CreateWorkItem) | **Post** /api/v2/workItems | Create Test Case, Checklist or Shared Step
[**DeleteAllWorkItemsFromAutoTest**](WorkItemsAPI.md#DeleteAllWorkItemsFromAutoTest) | **Delete** /api/v2/workItems/{id}/autoTests | Delete all links AutoTests from WorkItem by Id or GlobalId
[**DeleteWorkItem**](WorkItemsAPI.md#DeleteWorkItem) | **Delete** /api/v2/workItems/{id} | Delete Test Case, Checklist or Shared Step by Id or GlobalId
[**GetAutoTestsForWorkItem**](WorkItemsAPI.md#GetAutoTestsForWorkItem) | **Get** /api/v2/workItems/{id}/autoTests | Get all AutoTests linked to WorkItem by Id or GlobalId
[**GetIterations**](WorkItemsAPI.md#GetIterations) | **Get** /api/v2/workItems/{id}/iterations | Get iterations by work item Id or GlobalId
[**GetWorkItemById**](WorkItemsAPI.md#GetWorkItemById) | **Get** /api/v2/workItems/{id} | Get Test Case, Checklist or Shared Step by Id or GlobalId
[**GetWorkItemChronology**](WorkItemsAPI.md#GetWorkItemChronology) | **Get** /api/v2/workItems/{id}/chronology | Get WorkItem chronology by Id or GlobalId
[**GetWorkItemVersions**](WorkItemsAPI.md#GetWorkItemVersions) | **Get** /api/v2/workItems/{id}/versions | Get WorkItem versions
[**PurgeWorkItem**](WorkItemsAPI.md#PurgeWorkItem) | **Post** /api/v2/workItems/{id}/purge | Permanently delete test case, checklist or shared steps from archive
[**RestoreWorkItem**](WorkItemsAPI.md#RestoreWorkItem) | **Post** /api/v2/workItems/{id}/restore | Restore test case, checklist or shared steps from archive
[**UpdateWorkItem**](WorkItemsAPI.md#UpdateWorkItem) | **Put** /api/v2/workItems | Update Test Case, Checklist or Shared Step



## ApiV2WorkItemsIdAttachmentsPost

> ApiV2WorkItemsIdAttachmentsPost(ctx, id).File(file).Execute()

Upload and link attachment to WorkItem



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
	id := "3fa85f64-5717-4562-b3fc-2c963f66afa6" // string | Work item internal identifier (guid format)
	file := os.NewFile(1234, "some_file") // *os.File | Select file (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WorkItemsAPI.ApiV2WorkItemsIdAttachmentsPost(context.Background(), id).File(file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkItemsAPI.ApiV2WorkItemsIdAttachmentsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Work item internal identifier (guid format) | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2WorkItemsIdAttachmentsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

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


## ApiV2WorkItemsIdCheckListTransformToTestCasePost

> WorkItemModel ApiV2WorkItemsIdCheckListTransformToTestCasePost(ctx, id).Execute()

Transform CheckList to TestCase



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
	resp, r, err := apiClient.WorkItemsAPI.ApiV2WorkItemsIdCheckListTransformToTestCasePost(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkItemsAPI.ApiV2WorkItemsIdCheckListTransformToTestCasePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2WorkItemsIdCheckListTransformToTestCasePost`: WorkItemModel
	fmt.Fprintf(os.Stdout, "Response from `WorkItemsAPI.ApiV2WorkItemsIdCheckListTransformToTestCasePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2WorkItemsIdCheckListTransformToTestCasePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WorkItemModel**](WorkItemModel.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2WorkItemsIdHistoryGet

> []WorkItemChangeModel ApiV2WorkItemsIdHistoryGet(ctx, id).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).Execute()

Get change history of WorkItem



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
	skip := int32(56) // int32 | Amount of items to be skipped (offset) (optional)
	take := int32(56) // int32 | Amount of items to be taken (limit) (optional)
	orderBy := "orderBy_example" // string | SQL-like  ORDER BY statement (column1 ASC|DESC , column2 ASC|DESC) (optional)
	searchField := "searchField_example" // string | Property name for searching (optional)
	searchValue := "searchValue_example" // string | Value for searching (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkItemsAPI.ApiV2WorkItemsIdHistoryGet(context.Background(), id).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkItemsAPI.ApiV2WorkItemsIdHistoryGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2WorkItemsIdHistoryGet`: []WorkItemChangeModel
	fmt.Fprintf(os.Stdout, "Response from `WorkItemsAPI.ApiV2WorkItemsIdHistoryGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2WorkItemsIdHistoryGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **skip** | **int32** | Amount of items to be skipped (offset) | 
 **take** | **int32** | Amount of items to be taken (limit) | 
 **orderBy** | **string** | SQL-like  ORDER BY statement (column1 ASC|DESC , column2 ASC|DESC) | 
 **searchField** | **string** | Property name for searching | 
 **searchValue** | **string** | Value for searching | 

### Return type

[**[]WorkItemChangeModel**](WorkItemChangeModel.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2WorkItemsIdLikeDelete

> ApiV2WorkItemsIdLikeDelete(ctx, id).Execute()

Delete like from WorkItem



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
	r, err := apiClient.WorkItemsAPI.ApiV2WorkItemsIdLikeDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkItemsAPI.ApiV2WorkItemsIdLikeDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiApiV2WorkItemsIdLikeDeleteRequest struct via the builder pattern


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


## ApiV2WorkItemsIdLikePost

> ApiV2WorkItemsIdLikePost(ctx, id).Execute()

Set like to WorkItem



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
	r, err := apiClient.WorkItemsAPI.ApiV2WorkItemsIdLikePost(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkItemsAPI.ApiV2WorkItemsIdLikePost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiApiV2WorkItemsIdLikePostRequest struct via the builder pattern


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


## ApiV2WorkItemsIdLikesCountGet

> int32 ApiV2WorkItemsIdLikesCountGet(ctx, id).Execute()

Get likes count of WorkItem



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
	resp, r, err := apiClient.WorkItemsAPI.ApiV2WorkItemsIdLikesCountGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkItemsAPI.ApiV2WorkItemsIdLikesCountGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2WorkItemsIdLikesCountGet`: int32
	fmt.Fprintf(os.Stdout, "Response from `WorkItemsAPI.ApiV2WorkItemsIdLikesCountGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2WorkItemsIdLikesCountGetRequest struct via the builder pattern


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


## ApiV2WorkItemsIdLikesGet

> []WorkItemLikeModel ApiV2WorkItemsIdLikesGet(ctx, id).Execute()

Get likes of WorkItem



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
	resp, r, err := apiClient.WorkItemsAPI.ApiV2WorkItemsIdLikesGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkItemsAPI.ApiV2WorkItemsIdLikesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2WorkItemsIdLikesGet`: []WorkItemLikeModel
	fmt.Fprintf(os.Stdout, "Response from `WorkItemsAPI.ApiV2WorkItemsIdLikesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2WorkItemsIdLikesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]WorkItemLikeModel**](WorkItemLikeModel.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2WorkItemsIdTestResultsHistoryGet

> []TestResultHistoryReportApiResult ApiV2WorkItemsIdTestResultsHistoryGet(ctx, id).From(from).To(to).ConfigurationIds(configurationIds).TestPlanIds(testPlanIds).UserIds(userIds).Outcomes(outcomes).StatusCodes(statusCodes).IsAutomated(isAutomated).Automated(automated).TestRunIds(testRunIds).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).Execute()

Get test results history of WorkItem



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	from := time.Now() // time.Time | Take results from this date (optional)
	to := time.Now() // time.Time | Take results until this date (optional)
	configurationIds := []string{"Inner_example"} // []string | Identifiers of test result configurations (optional)
	testPlanIds := []string{"Inner_example"} // []string | Identifiers of test plans which contain test results (optional)
	userIds := []string{"Inner_example"} // []string | Identifiers of users who set test results (optional)
	outcomes := []string{"Inner_example"} // []string | List of outcomes of test results (optional)
	statusCodes := []string{"Inner_example"} // []string | List of status codes of test results (optional)
	isAutomated := true // bool | OBSOLETE: Use `Automated` instead (optional)
	automated := true // bool | If result must consist of only manual/automated test results (optional)
	testRunIds := []string{"Inner_example"} // []string | Identifiers of test runs which contain test results (optional)
	skip := int32(56) // int32 | Amount of items to be skipped (offset) (optional)
	take := int32(56) // int32 | Amount of items to be taken (limit) (optional)
	orderBy := "orderBy_example" // string | SQL-like  ORDER BY statement (column1 ASC|DESC , column2 ASC|DESC) (optional)
	searchField := "searchField_example" // string | Property name for searching (optional)
	searchValue := "searchValue_example" // string | Value for searching (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkItemsAPI.ApiV2WorkItemsIdTestResultsHistoryGet(context.Background(), id).From(from).To(to).ConfigurationIds(configurationIds).TestPlanIds(testPlanIds).UserIds(userIds).Outcomes(outcomes).StatusCodes(statusCodes).IsAutomated(isAutomated).Automated(automated).TestRunIds(testRunIds).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkItemsAPI.ApiV2WorkItemsIdTestResultsHistoryGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2WorkItemsIdTestResultsHistoryGet`: []TestResultHistoryReportApiResult
	fmt.Fprintf(os.Stdout, "Response from `WorkItemsAPI.ApiV2WorkItemsIdTestResultsHistoryGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2WorkItemsIdTestResultsHistoryGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **from** | **time.Time** | Take results from this date | 
 **to** | **time.Time** | Take results until this date | 
 **configurationIds** | **[]string** | Identifiers of test result configurations | 
 **testPlanIds** | **[]string** | Identifiers of test plans which contain test results | 
 **userIds** | **[]string** | Identifiers of users who set test results | 
 **outcomes** | **[]string** | List of outcomes of test results | 
 **statusCodes** | **[]string** | List of status codes of test results | 
 **isAutomated** | **bool** | OBSOLETE: Use &#x60;Automated&#x60; instead | 
 **automated** | **bool** | If result must consist of only manual/automated test results | 
 **testRunIds** | **[]string** | Identifiers of test runs which contain test results | 
 **skip** | **int32** | Amount of items to be skipped (offset) | 
 **take** | **int32** | Amount of items to be taken (limit) | 
 **orderBy** | **string** | SQL-like  ORDER BY statement (column1 ASC|DESC , column2 ASC|DESC) | 
 **searchField** | **string** | Property name for searching | 
 **searchValue** | **string** | Value for searching | 

### Return type

[**[]TestResultHistoryReportApiResult**](TestResultHistoryReportApiResult.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2WorkItemsIdVersionVersionIdActualPost

> WorkItemModel ApiV2WorkItemsIdVersionVersionIdActualPost(ctx, id, versionId).Execute()

Set WorkItem as actual



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
	versionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkItemsAPI.ApiV2WorkItemsIdVersionVersionIdActualPost(context.Background(), id, versionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkItemsAPI.ApiV2WorkItemsIdVersionVersionIdActualPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2WorkItemsIdVersionVersionIdActualPost`: WorkItemModel
	fmt.Fprintf(os.Stdout, "Response from `WorkItemsAPI.ApiV2WorkItemsIdVersionVersionIdActualPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**versionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2WorkItemsIdVersionVersionIdActualPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**WorkItemModel**](WorkItemModel.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2WorkItemsLinksUrlsSearchPost

> SearchWorkItemLinkUrlsApiResult ApiV2WorkItemsLinksUrlsSearchPost(ctx).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).WorkItemLinkUrlApiModel(workItemLinkUrlApiModel).Execute()



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
	workItemLinkUrlApiModel := *openapiclient.NewWorkItemLinkUrlApiModel(*openapiclient.NewWorkItemLinkUrlFilterApiModel(), *openapiclient.NewWorkItemLinkExtractionApiModel()) // WorkItemLinkUrlApiModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkItemsAPI.ApiV2WorkItemsLinksUrlsSearchPost(context.Background()).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).WorkItemLinkUrlApiModel(workItemLinkUrlApiModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkItemsAPI.ApiV2WorkItemsLinksUrlsSearchPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2WorkItemsLinksUrlsSearchPost`: SearchWorkItemLinkUrlsApiResult
	fmt.Fprintf(os.Stdout, "Response from `WorkItemsAPI.ApiV2WorkItemsLinksUrlsSearchPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2WorkItemsLinksUrlsSearchPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skip** | **int32** | Amount of items to be skipped (offset) | 
 **take** | **int32** | Amount of items to be taken (limit) | 
 **orderBy** | **string** | SQL-like  ORDER BY statement (column1 ASC|DESC , column2 ASC|DESC) | 
 **searchField** | **string** | Property name for searching | 
 **searchValue** | **string** | Value for searching | 
 **workItemLinkUrlApiModel** | [**WorkItemLinkUrlApiModel**](WorkItemLinkUrlApiModel.md) |  | 

### Return type

[**SearchWorkItemLinkUrlsApiResult**](SearchWorkItemLinkUrlsApiResult.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2WorkItemsMovePost

> WorkItemShortModel ApiV2WorkItemsMovePost(ctx).WorkItemMovePostModel(workItemMovePostModel).Execute()

Move WorkItem to another section



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
	workItemMovePostModel := *openapiclient.NewWorkItemMovePostModel("Id_example", "NewSectionId_example") // WorkItemMovePostModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkItemsAPI.ApiV2WorkItemsMovePost(context.Background()).WorkItemMovePostModel(workItemMovePostModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkItemsAPI.ApiV2WorkItemsMovePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2WorkItemsMovePost`: WorkItemShortModel
	fmt.Fprintf(os.Stdout, "Response from `WorkItemsAPI.ApiV2WorkItemsMovePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2WorkItemsMovePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workItemMovePostModel** | [**WorkItemMovePostModel**](WorkItemMovePostModel.md) |  | 

### Return type

[**WorkItemShortModel**](WorkItemShortModel.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2WorkItemsSearchPost

> []WorkItemShortApiResult ApiV2WorkItemsSearchPost(ctx).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).WorkItemSelectApiModel(workItemSelectApiModel).Execute()

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
	skip := int32(56) // int32 | Amount of items to be skipped (offset) (optional)
	take := int32(56) // int32 | Amount of items to be taken (limit) (optional)
	orderBy := "orderBy_example" // string | SQL-like  ORDER BY statement (column1 ASC|DESC , column2 ASC|DESC) (optional)
	searchField := "searchField_example" // string | Property name for searching (optional)
	searchValue := "searchValue_example" // string | Value for searching (optional)
	workItemSelectApiModel := *openapiclient.NewWorkItemSelectApiModel(*openapiclient.NewWorkItemFilterApiModel()) // WorkItemSelectApiModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkItemsAPI.ApiV2WorkItemsSearchPost(context.Background()).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).WorkItemSelectApiModel(workItemSelectApiModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkItemsAPI.ApiV2WorkItemsSearchPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2WorkItemsSearchPost`: []WorkItemShortApiResult
	fmt.Fprintf(os.Stdout, "Response from `WorkItemsAPI.ApiV2WorkItemsSearchPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2WorkItemsSearchPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skip** | **int32** | Amount of items to be skipped (offset) | 
 **take** | **int32** | Amount of items to be taken (limit) | 
 **orderBy** | **string** | SQL-like  ORDER BY statement (column1 ASC|DESC , column2 ASC|DESC) | 
 **searchField** | **string** | Property name for searching | 
 **searchValue** | **string** | Value for searching | 
 **workItemSelectApiModel** | [**WorkItemSelectApiModel**](WorkItemSelectApiModel.md) |  | 

### Return type

[**[]WorkItemShortApiResult**](WorkItemShortApiResult.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2WorkItemsSharedStepIdReferencesSectionsPost

> []SharedStepReferenceSectionModel ApiV2WorkItemsSharedStepIdReferencesSectionsPost(ctx, sharedStepId).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).SharedStepReferenceSectionsQueryFilterModel(sharedStepReferenceSectionsQueryFilterModel).Execute()

Get SharedStep references in sections



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
	sharedStepId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	skip := int32(56) // int32 | Amount of items to be skipped (offset) (optional)
	take := int32(56) // int32 | Amount of items to be taken (limit) (optional)
	orderBy := "orderBy_example" // string | SQL-like  ORDER BY statement (column1 ASC|DESC , column2 ASC|DESC) (optional)
	searchField := "searchField_example" // string | Property name for searching (optional)
	searchValue := "searchValue_example" // string | Value for searching (optional)
	sharedStepReferenceSectionsQueryFilterModel := *openapiclient.NewSharedStepReferenceSectionsQueryFilterModel() // SharedStepReferenceSectionsQueryFilterModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkItemsAPI.ApiV2WorkItemsSharedStepIdReferencesSectionsPost(context.Background(), sharedStepId).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).SharedStepReferenceSectionsQueryFilterModel(sharedStepReferenceSectionsQueryFilterModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkItemsAPI.ApiV2WorkItemsSharedStepIdReferencesSectionsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2WorkItemsSharedStepIdReferencesSectionsPost`: []SharedStepReferenceSectionModel
	fmt.Fprintf(os.Stdout, "Response from `WorkItemsAPI.ApiV2WorkItemsSharedStepIdReferencesSectionsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sharedStepId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2WorkItemsSharedStepIdReferencesSectionsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **skip** | **int32** | Amount of items to be skipped (offset) | 
 **take** | **int32** | Amount of items to be taken (limit) | 
 **orderBy** | **string** | SQL-like  ORDER BY statement (column1 ASC|DESC , column2 ASC|DESC) | 
 **searchField** | **string** | Property name for searching | 
 **searchValue** | **string** | Value for searching | 
 **sharedStepReferenceSectionsQueryFilterModel** | [**SharedStepReferenceSectionsQueryFilterModel**](SharedStepReferenceSectionsQueryFilterModel.md) |  | 

### Return type

[**[]SharedStepReferenceSectionModel**](SharedStepReferenceSectionModel.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2WorkItemsSharedStepIdReferencesWorkItemsPost

> []SharedStepReferenceModel ApiV2WorkItemsSharedStepIdReferencesWorkItemsPost(ctx, sharedStepId).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).SharedStepReferencesQueryFilterModel(sharedStepReferencesQueryFilterModel).Execute()

Get SharedStep references in work items



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
	sharedStepId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	skip := int32(56) // int32 | Amount of items to be skipped (offset) (optional)
	take := int32(56) // int32 | Amount of items to be taken (limit) (optional)
	orderBy := "orderBy_example" // string | SQL-like  ORDER BY statement (column1 ASC|DESC , column2 ASC|DESC) (optional)
	searchField := "searchField_example" // string | Property name for searching (optional)
	searchValue := "searchValue_example" // string | Value for searching (optional)
	sharedStepReferencesQueryFilterModel := *openapiclient.NewSharedStepReferencesQueryFilterModel() // SharedStepReferencesQueryFilterModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkItemsAPI.ApiV2WorkItemsSharedStepIdReferencesWorkItemsPost(context.Background(), sharedStepId).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).SharedStepReferencesQueryFilterModel(sharedStepReferencesQueryFilterModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkItemsAPI.ApiV2WorkItemsSharedStepIdReferencesWorkItemsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2WorkItemsSharedStepIdReferencesWorkItemsPost`: []SharedStepReferenceModel
	fmt.Fprintf(os.Stdout, "Response from `WorkItemsAPI.ApiV2WorkItemsSharedStepIdReferencesWorkItemsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sharedStepId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2WorkItemsSharedStepIdReferencesWorkItemsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **skip** | **int32** | Amount of items to be skipped (offset) | 
 **take** | **int32** | Amount of items to be taken (limit) | 
 **orderBy** | **string** | SQL-like  ORDER BY statement (column1 ASC|DESC , column2 ASC|DESC) | 
 **searchField** | **string** | Property name for searching | 
 **searchValue** | **string** | Value for searching | 
 **sharedStepReferencesQueryFilterModel** | [**SharedStepReferencesQueryFilterModel**](SharedStepReferencesQueryFilterModel.md) |  | 

### Return type

[**[]SharedStepReferenceModel**](SharedStepReferenceModel.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2WorkItemsSharedStepsSharedStepIdReferencesGet

> []SharedStepReferenceModel ApiV2WorkItemsSharedStepsSharedStepIdReferencesGet(ctx, sharedStepId).Execute()

Get SharedStep references



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
	sharedStepId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkItemsAPI.ApiV2WorkItemsSharedStepsSharedStepIdReferencesGet(context.Background(), sharedStepId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkItemsAPI.ApiV2WorkItemsSharedStepsSharedStepIdReferencesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2WorkItemsSharedStepsSharedStepIdReferencesGet`: []SharedStepReferenceModel
	fmt.Fprintf(os.Stdout, "Response from `WorkItemsAPI.ApiV2WorkItemsSharedStepsSharedStepIdReferencesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sharedStepId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2WorkItemsSharedStepsSharedStepIdReferencesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]SharedStepReferenceModel**](SharedStepReferenceModel.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateWorkItem

> WorkItemModel CreateWorkItem(ctx).CreateWorkItemApiModel(createWorkItemApiModel).Execute()

Create Test Case, Checklist or Shared Step



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
	createWorkItemApiModel := *openapiclient.NewCreateWorkItemApiModel(openapiclient.WorkItemEntityTypes("TestCases"), openapiclient.WorkItemStates("NeedsWork"), openapiclient.WorkItemPriorityModel("Lowest"), []openapiclient.CreateStepApiModel{*openapiclient.NewCreateStepApiModel()}, []openapiclient.CreateStepApiModel{*openapiclient.NewCreateStepApiModel()}, []openapiclient.CreateStepApiModel{*openapiclient.NewCreateStepApiModel()}, int32(123), map[string]interface{}{"key": interface{}(123)}, []openapiclient.TagModel{*openapiclient.NewTagModel("Name_example")}, []openapiclient.CreateLinkApiModel{*openapiclient.NewCreateLinkApiModel("Url_example", false)}, "Name_example", "ProjectId_example", "SectionId_example") // CreateWorkItemApiModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkItemsAPI.CreateWorkItem(context.Background()).CreateWorkItemApiModel(createWorkItemApiModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkItemsAPI.CreateWorkItem``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateWorkItem`: WorkItemModel
	fmt.Fprintf(os.Stdout, "Response from `WorkItemsAPI.CreateWorkItem`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorkItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createWorkItemApiModel** | [**CreateWorkItemApiModel**](CreateWorkItemApiModel.md) |  | 

### Return type

[**WorkItemModel**](WorkItemModel.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAllWorkItemsFromAutoTest

> DeleteAllWorkItemsFromAutoTest(ctx, id).Execute()

Delete all links AutoTests from WorkItem by Id or GlobalId



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
	id := "3fa85f64-5717-4562-b3fc-2c963f66afa6" // string | WorkItem internal (guid format) or global(integer format) identifier\"

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WorkItemsAPI.DeleteAllWorkItemsFromAutoTest(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkItemsAPI.DeleteAllWorkItemsFromAutoTest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | WorkItem internal (guid format) or global(integer format) identifier\&quot; | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAllWorkItemsFromAutoTestRequest struct via the builder pattern


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


## DeleteWorkItem

> DeleteWorkItem(ctx, id).Execute()

Delete Test Case, Checklist or Shared Step by Id or GlobalId



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
	id := "3fa85f64-5717-4562-b3fc-2c963f66afa6" // string | WorkItem internal (guid format) or global(integer format) identifier\"

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WorkItemsAPI.DeleteWorkItem(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkItemsAPI.DeleteWorkItem``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | WorkItem internal (guid format) or global(integer format) identifier\&quot; | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWorkItemRequest struct via the builder pattern


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


## GetAutoTestsForWorkItem

> []AutoTestModel GetAutoTestsForWorkItem(ctx, id).Execute()

Get all AutoTests linked to WorkItem by Id or GlobalId



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
	id := "3fa85f64-5717-4562-b3fc-2c963f66afa6" // string | WorkItem internal (guid format) or global(integer format) identifier\"

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkItemsAPI.GetAutoTestsForWorkItem(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkItemsAPI.GetAutoTestsForWorkItem``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAutoTestsForWorkItem`: []AutoTestModel
	fmt.Fprintf(os.Stdout, "Response from `WorkItemsAPI.GetAutoTestsForWorkItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | WorkItem internal (guid format) or global(integer format) identifier\&quot; | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAutoTestsForWorkItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]AutoTestModel**](AutoTestModel.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIterations

> []IterationModel GetIterations(ctx, id).VersionId(versionId).VersionNumber(versionNumber).Execute()

Get iterations by work item Id or GlobalId

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
	id := "3fa85f64-5717-4562-b3fc-2c963f66afa6" // string | WorkItem internal (guid format) or global(integer format) identifier\"
	versionId := "00000000-0000-0000-0000-000000000000" // string | WorkItem version (guid format) identifier (optional)
	versionNumber := int32(0) // int32 | WorkItem version number (0 is the last version)\" (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkItemsAPI.GetIterations(context.Background(), id).VersionId(versionId).VersionNumber(versionNumber).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkItemsAPI.GetIterations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIterations`: []IterationModel
	fmt.Fprintf(os.Stdout, "Response from `WorkItemsAPI.GetIterations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | WorkItem internal (guid format) or global(integer format) identifier\&quot; | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIterationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **versionId** | **string** | WorkItem version (guid format) identifier | 
 **versionNumber** | **int32** | WorkItem version number (0 is the last version)\&quot; | 

### Return type

[**[]IterationModel**](IterationModel.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWorkItemById

> WorkItemModel GetWorkItemById(ctx, id).VersionId(versionId).VersionNumber(versionNumber).Execute()

Get Test Case, Checklist or Shared Step by Id or GlobalId



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
	id := "3fa85f64-5717-4562-b3fc-2c963f66afa6" // string | WorkItem internal (guid format) or global(integer format) identifier\"
	versionId := "00000000-0000-0000-0000-000000000000" // string | WorkItem version (guid format) identifier\" (optional)
	versionNumber := int32(0) // int32 | WorkItem version number (0 is the last version)\" (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkItemsAPI.GetWorkItemById(context.Background(), id).VersionId(versionId).VersionNumber(versionNumber).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkItemsAPI.GetWorkItemById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWorkItemById`: WorkItemModel
	fmt.Fprintf(os.Stdout, "Response from `WorkItemsAPI.GetWorkItemById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | WorkItem internal (guid format) or global(integer format) identifier\&quot; | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkItemByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **versionId** | **string** | WorkItem version (guid format) identifier\&quot; | 
 **versionNumber** | **int32** | WorkItem version number (0 is the last version)\&quot; | 

### Return type

[**WorkItemModel**](WorkItemModel.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWorkItemChronology

> []TestResultChronologyModel GetWorkItemChronology(ctx, id).Execute()

Get WorkItem chronology by Id or GlobalId



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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkItemsAPI.GetWorkItemChronology(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkItemsAPI.GetWorkItemChronology``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWorkItemChronology`: []TestResultChronologyModel
	fmt.Fprintf(os.Stdout, "Response from `WorkItemsAPI.GetWorkItemChronology`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkItemChronologyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]TestResultChronologyModel**](TestResultChronologyModel.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWorkItemVersions

> []WorkItemVersionModel GetWorkItemVersions(ctx, id).WorkItemVersionId(workItemVersionId).VersionNumber(versionNumber).Execute()

Get WorkItem versions



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
	id := "3fa85f64-5717-4562-b3fc-2c963f66afa6" // string | WorkItem internal (guid format) or global(integer format) identifier\"
	workItemVersionId := "3fa85f64-5717-4562-b3fc-2c963f66afa6" // string | WorkItem version (guid format) identifier\" (optional)
	versionNumber := int32(1) // int32 | WorkItem version (integer format) number\" (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkItemsAPI.GetWorkItemVersions(context.Background(), id).WorkItemVersionId(workItemVersionId).VersionNumber(versionNumber).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkItemsAPI.GetWorkItemVersions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWorkItemVersions`: []WorkItemVersionModel
	fmt.Fprintf(os.Stdout, "Response from `WorkItemsAPI.GetWorkItemVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | WorkItem internal (guid format) or global(integer format) identifier\&quot; | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkItemVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **workItemVersionId** | **string** | WorkItem version (guid format) identifier\&quot; | 
 **versionNumber** | **int32** | WorkItem version (integer format) number\&quot; | 

### Return type

[**[]WorkItemVersionModel**](WorkItemVersionModel.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PurgeWorkItem

> PurgeWorkItem(ctx, id).Execute()

Permanently delete test case, checklist or shared steps from archive

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
	r, err := apiClient.WorkItemsAPI.PurgeWorkItem(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkItemsAPI.PurgeWorkItem``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Unique or global ID of the work item | 

### Other Parameters

Other parameters are passed through a pointer to a apiPurgeWorkItemRequest struct via the builder pattern


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


## RestoreWorkItem

> RestoreWorkItem(ctx, id).Execute()

Restore test case, checklist or shared steps from archive

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
	r, err := apiClient.WorkItemsAPI.RestoreWorkItem(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkItemsAPI.RestoreWorkItem``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Unique or global ID of the work item | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestoreWorkItemRequest struct via the builder pattern


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


## UpdateWorkItem

> UpdateWorkItem(ctx).UpdateWorkItemApiModel(updateWorkItemApiModel).Execute()

Update Test Case, Checklist or Shared Step



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
	updateWorkItemApiModel := *openapiclient.NewUpdateWorkItemApiModel("Id_example", "SectionId_example", openapiclient.WorkItemStates("NeedsWork"), openapiclient.WorkItemPriorityModel("Lowest"), []openapiclient.UpdateStepApiModel{*openapiclient.NewUpdateStepApiModel("Id_example")}, []openapiclient.UpdateStepApiModel{*openapiclient.NewUpdateStepApiModel("Id_example")}, []openapiclient.UpdateStepApiModel{*openapiclient.NewUpdateStepApiModel("Id_example")}, int32(123), map[string]interface{}{"key": interface{}(123)}, []openapiclient.TagModel{*openapiclient.NewTagModel("Name_example")}, []openapiclient.UpdateLinkApiModel{*openapiclient.NewUpdateLinkApiModel("Url_example", false)}, "Name_example", []openapiclient.AssignAttachmentApiModel{*openapiclient.NewAssignAttachmentApiModel("Id_example")}) // UpdateWorkItemApiModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WorkItemsAPI.UpdateWorkItem(context.Background()).UpdateWorkItemApiModel(updateWorkItemApiModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkItemsAPI.UpdateWorkItem``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWorkItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateWorkItemApiModel** | [**UpdateWorkItemApiModel**](UpdateWorkItemApiModel.md) |  | 

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

