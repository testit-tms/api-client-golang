# \AutoTestsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV2AutoTestsDelete**](AutoTestsAPI.md#ApiV2AutoTestsDelete) | **Delete** /api/v2/autoTests | Delete autotests
[**ApiV2AutoTestsFlakyBulkPost**](AutoTestsAPI.md#ApiV2AutoTestsFlakyBulkPost) | **Post** /api/v2/autoTests/flaky/bulk | Set \&quot;Flaky\&quot; status for multiple autotests
[**ApiV2AutoTestsIdPatch**](AutoTestsAPI.md#ApiV2AutoTestsIdPatch) | **Patch** /api/v2/autoTests/{id} | Patch auto test
[**ApiV2AutoTestsIdTestResultsSearchPost**](AutoTestsAPI.md#ApiV2AutoTestsIdTestResultsSearchPost) | **Post** /api/v2/autoTests/{id}/testResults/search | Get test results history for autotest
[**ApiV2AutoTestsIdWorkItemsChangedIdGet**](AutoTestsAPI.md#ApiV2AutoTestsIdWorkItemsChangedIdGet) | **Get** /api/v2/autoTests/{id}/workItems/changed/id | Get identifiers of changed linked work items
[**ApiV2AutoTestsIdWorkItemsChangedWorkItemIdApprovePost**](AutoTestsAPI.md#ApiV2AutoTestsIdWorkItemsChangedWorkItemIdApprovePost) | **Post** /api/v2/autoTests/{id}/workItems/changed/{workItemId}/approve | Approve changes to work items linked to autotest
[**ApiV2AutoTestsSearchPost**](AutoTestsAPI.md#ApiV2AutoTestsSearchPost) | **Post** /api/v2/autoTests/search | Search for autotests
[**CreateAutoTest**](AutoTestsAPI.md#CreateAutoTest) | **Post** /api/v2/autoTests | Create autotest
[**CreateMultiple**](AutoTestsAPI.md#CreateMultiple) | **Post** /api/v2/autoTests/bulk | Create multiple autotests
[**DeleteAutoTest**](AutoTestsAPI.md#DeleteAutoTest) | **Delete** /api/v2/autoTests/{id} | Delete autotest
[**DeleteAutoTestLinkFromWorkItem**](AutoTestsAPI.md#DeleteAutoTestLinkFromWorkItem) | **Delete** /api/v2/autoTests/{id}/workItems | Unlink autotest from work item
[**GetAllAutoTests**](AutoTestsAPI.md#GetAllAutoTests) | **Get** /api/v2/autoTests | 
[**GetAutoTestAverageDuration**](AutoTestsAPI.md#GetAutoTestAverageDuration) | **Get** /api/v2/autoTests/{id}/averageDuration | Get average autotest duration
[**GetAutoTestById**](AutoTestsAPI.md#GetAutoTestById) | **Get** /api/v2/autoTests/{id} | Get autotest by internal or global ID
[**GetAutoTestChronology**](AutoTestsAPI.md#GetAutoTestChronology) | **Get** /api/v2/autoTests/{id}/chronology | Get autotest chronology
[**GetTestRuns**](AutoTestsAPI.md#GetTestRuns) | **Get** /api/v2/autoTests/{id}/testRuns | Get completed tests runs for autotests
[**GetWorkItemsLinkedToAutoTest**](AutoTestsAPI.md#GetWorkItemsLinkedToAutoTest) | **Get** /api/v2/autoTests/{id}/workItems | Get work items linked to autotest
[**LinkAutoTestToWorkItem**](AutoTestsAPI.md#LinkAutoTestToWorkItem) | **Post** /api/v2/autoTests/{id}/workItems | Link autotest with work items
[**UpdateAutoTest**](AutoTestsAPI.md#UpdateAutoTest) | **Put** /api/v2/autoTests | Update autotest
[**UpdateMultiple**](AutoTestsAPI.md#UpdateMultiple) | **Put** /api/v2/autoTests/bulk | Update multiple autotests



## ApiV2AutoTestsDelete

> AutoTestBulkDeleteApiResult ApiV2AutoTestsDelete(ctx).AutoTestBulkDeleteApiModel(autoTestBulkDeleteApiModel).Execute()

Delete autotests

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
	autoTestBulkDeleteApiModel := *openapiclient.NewAutoTestBulkDeleteApiModel(*openapiclient.NewAutoTestSelectModel(*openapiclient.NewAutoTestFilterModel())) // AutoTestBulkDeleteApiModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutoTestsAPI.ApiV2AutoTestsDelete(context.Background()).AutoTestBulkDeleteApiModel(autoTestBulkDeleteApiModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoTestsAPI.ApiV2AutoTestsDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2AutoTestsDelete`: AutoTestBulkDeleteApiResult
	fmt.Fprintf(os.Stdout, "Response from `AutoTestsAPI.ApiV2AutoTestsDelete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2AutoTestsDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **autoTestBulkDeleteApiModel** | [**AutoTestBulkDeleteApiModel**](AutoTestBulkDeleteApiModel.md) |  | 

### Return type

[**AutoTestBulkDeleteApiResult**](AutoTestBulkDeleteApiResult.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2AutoTestsFlakyBulkPost

> ApiV2AutoTestsFlakyBulkPost(ctx).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).AutoTestFlakyBulkApiModel(autoTestFlakyBulkApiModel).Execute()

Set \"Flaky\" status for multiple autotests



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
	autoTestFlakyBulkApiModel := *openapiclient.NewAutoTestFlakyBulkApiModel(*openapiclient.NewAutoTestSelectApiModel(), false) // AutoTestFlakyBulkApiModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AutoTestsAPI.ApiV2AutoTestsFlakyBulkPost(context.Background()).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).AutoTestFlakyBulkApiModel(autoTestFlakyBulkApiModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoTestsAPI.ApiV2AutoTestsFlakyBulkPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2AutoTestsFlakyBulkPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skip** | **int32** | Amount of items to be skipped (offset) | 
 **take** | **int32** | Amount of items to be taken (limit) | 
 **orderBy** | **string** | SQL-like  ORDER BY statement (column1 ASC|DESC , column2 ASC|DESC) | 
 **searchField** | **string** | Property name for searching | 
 **searchValue** | **string** | Value for searching | 
 **autoTestFlakyBulkApiModel** | [**AutoTestFlakyBulkApiModel**](AutoTestFlakyBulkApiModel.md) |  | 

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


## ApiV2AutoTestsIdPatch

> ApiV2AutoTestsIdPatch(ctx, id).Operation(operation).Execute()

Patch auto test



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Global Id of auto test
	operation := []openapiclient.Operation{*openapiclient.NewOperation()} // []Operation |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AutoTestsAPI.ApiV2AutoTestsIdPatch(context.Background(), id).Operation(operation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoTestsAPI.ApiV2AutoTestsIdPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Global Id of auto test | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2AutoTestsIdPatchRequest struct via the builder pattern


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


## ApiV2AutoTestsIdTestResultsSearchPost

> []AutoTestResultHistoryApiResult ApiV2AutoTestsIdTestResultsSearchPost(ctx, id).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).AutoTestResultHistorySelectApiModel(autoTestResultHistorySelectApiModel).Execute()

Get test results history for autotest



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
	id := "id_example" // string | Autotest identifier
	skip := int32(56) // int32 | Amount of items to be skipped (offset) (optional)
	take := int32(56) // int32 | Amount of items to be taken (limit) (optional)
	orderBy := "orderBy_example" // string | SQL-like  ORDER BY statement (column1 ASC|DESC , column2 ASC|DESC) (optional)
	searchField := "searchField_example" // string | Property name for searching (optional)
	searchValue := "searchValue_example" // string | Value for searching (optional)
	autoTestResultHistorySelectApiModel := *openapiclient.NewAutoTestResultHistorySelectApiModel() // AutoTestResultHistorySelectApiModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutoTestsAPI.ApiV2AutoTestsIdTestResultsSearchPost(context.Background(), id).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).AutoTestResultHistorySelectApiModel(autoTestResultHistorySelectApiModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoTestsAPI.ApiV2AutoTestsIdTestResultsSearchPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2AutoTestsIdTestResultsSearchPost`: []AutoTestResultHistoryApiResult
	fmt.Fprintf(os.Stdout, "Response from `AutoTestsAPI.ApiV2AutoTestsIdTestResultsSearchPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Autotest identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2AutoTestsIdTestResultsSearchPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **skip** | **int32** | Amount of items to be skipped (offset) | 
 **take** | **int32** | Amount of items to be taken (limit) | 
 **orderBy** | **string** | SQL-like  ORDER BY statement (column1 ASC|DESC , column2 ASC|DESC) | 
 **searchField** | **string** | Property name for searching | 
 **searchValue** | **string** | Value for searching | 
 **autoTestResultHistorySelectApiModel** | [**AutoTestResultHistorySelectApiModel**](AutoTestResultHistorySelectApiModel.md) |  | 

### Return type

[**[]AutoTestResultHistoryApiResult**](AutoTestResultHistoryApiResult.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2AutoTestsIdWorkItemsChangedIdGet

> []string ApiV2AutoTestsIdWorkItemsChangedIdGet(ctx, id).Execute()

Get identifiers of changed linked work items



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
	resp, r, err := apiClient.AutoTestsAPI.ApiV2AutoTestsIdWorkItemsChangedIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoTestsAPI.ApiV2AutoTestsIdWorkItemsChangedIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2AutoTestsIdWorkItemsChangedIdGet`: []string
	fmt.Fprintf(os.Stdout, "Response from `AutoTestsAPI.ApiV2AutoTestsIdWorkItemsChangedIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2AutoTestsIdWorkItemsChangedIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**[]string**

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2AutoTestsIdWorkItemsChangedWorkItemIdApprovePost

> ApiV2AutoTestsIdWorkItemsChangedWorkItemIdApprovePost(ctx, id, workItemId).Execute()

Approve changes to work items linked to autotest



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
	workItemId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AutoTestsAPI.ApiV2AutoTestsIdWorkItemsChangedWorkItemIdApprovePost(context.Background(), id, workItemId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoTestsAPI.ApiV2AutoTestsIdWorkItemsChangedWorkItemIdApprovePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**workItemId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2AutoTestsIdWorkItemsChangedWorkItemIdApprovePostRequest struct via the builder pattern


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


## ApiV2AutoTestsSearchPost

> []AutoTestApiResult ApiV2AutoTestsSearchPost(ctx).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).AutoTestSearchApiModel(autoTestSearchApiModel).Execute()

Search for autotests

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
	autoTestSearchApiModel := *openapiclient.NewAutoTestSearchApiModel() // AutoTestSearchApiModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutoTestsAPI.ApiV2AutoTestsSearchPost(context.Background()).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).AutoTestSearchApiModel(autoTestSearchApiModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoTestsAPI.ApiV2AutoTestsSearchPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2AutoTestsSearchPost`: []AutoTestApiResult
	fmt.Fprintf(os.Stdout, "Response from `AutoTestsAPI.ApiV2AutoTestsSearchPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2AutoTestsSearchPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skip** | **int32** | Amount of items to be skipped (offset) | 
 **take** | **int32** | Amount of items to be taken (limit) | 
 **orderBy** | **string** | SQL-like  ORDER BY statement (column1 ASC|DESC , column2 ASC|DESC) | 
 **searchField** | **string** | Property name for searching | 
 **searchValue** | **string** | Value for searching | 
 **autoTestSearchApiModel** | [**AutoTestSearchApiModel**](AutoTestSearchApiModel.md) |  | 

### Return type

[**[]AutoTestApiResult**](AutoTestApiResult.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAutoTest

> AutoTestApiResult CreateAutoTest(ctx).AutoTestCreateApiModel(autoTestCreateApiModel).Execute()

Create autotest



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
	autoTestCreateApiModel := *openapiclient.NewAutoTestCreateApiModel("ProjectId_example", "ExternalId_example", "Name_example") // AutoTestCreateApiModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutoTestsAPI.CreateAutoTest(context.Background()).AutoTestCreateApiModel(autoTestCreateApiModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoTestsAPI.CreateAutoTest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAutoTest`: AutoTestApiResult
	fmt.Fprintf(os.Stdout, "Response from `AutoTestsAPI.CreateAutoTest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAutoTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **autoTestCreateApiModel** | [**AutoTestCreateApiModel**](AutoTestCreateApiModel.md) |  | 

### Return type

[**AutoTestApiResult**](AutoTestApiResult.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMultiple

> []AutoTestApiResult CreateMultiple(ctx).AutoTestCreateApiModel(autoTestCreateApiModel).Execute()

Create multiple autotests



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
	autoTestCreateApiModel := []openapiclient.AutoTestCreateApiModel{*openapiclient.NewAutoTestCreateApiModel("ProjectId_example", "ExternalId_example", "Name_example")} // []AutoTestCreateApiModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutoTestsAPI.CreateMultiple(context.Background()).AutoTestCreateApiModel(autoTestCreateApiModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoTestsAPI.CreateMultiple``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateMultiple`: []AutoTestApiResult
	fmt.Fprintf(os.Stdout, "Response from `AutoTestsAPI.CreateMultiple`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateMultipleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **autoTestCreateApiModel** | [**[]AutoTestCreateApiModel**](AutoTestCreateApiModel.md) |  | 

### Return type

[**[]AutoTestApiResult**](AutoTestApiResult.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAutoTest

> DeleteAutoTest(ctx, id).Execute()

Delete autotest



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
	id := "id_example" // string | Autotest internal (UUID) or global (integer) identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AutoTestsAPI.DeleteAutoTest(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoTestsAPI.DeleteAutoTest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Autotest internal (UUID) or global (integer) identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAutoTestRequest struct via the builder pattern


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


## DeleteAutoTestLinkFromWorkItem

> DeleteAutoTestLinkFromWorkItem(ctx, id).WorkItemId(workItemId).Execute()

Unlink autotest from work item



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
	id := "id_example" // string | Autotest internal (UUID) or global (integer) identifier
	workItemId := "workItemId_example" // string | Work item internal (UUID) or global (integer) identifier (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AutoTestsAPI.DeleteAutoTestLinkFromWorkItem(context.Background(), id).WorkItemId(workItemId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoTestsAPI.DeleteAutoTestLinkFromWorkItem``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Autotest internal (UUID) or global (integer) identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAutoTestLinkFromWorkItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **workItemId** | **string** | Work item internal (UUID) or global (integer) identifier | 

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


## GetAllAutoTests

> []AutoTestModel GetAllAutoTests(ctx).ProjectId(projectId).ExternalId(externalId).GlobalId(globalId).Namespace(namespace).IsNamespaceNull(isNamespaceNull).IncludeEmptyNamespaces(includeEmptyNamespaces).ClassName(className).IsClassnameNull(isClassnameNull).IncludeEmptyClassNames(includeEmptyClassNames).IsDeleted(isDeleted).Deleted(deleted).Labels(labels).StabilityMinimal(stabilityMinimal).MinStability(minStability).StabilityMaximal(stabilityMaximal).MaxStability(maxStability).IsFlaky(isFlaky).Flaky(flaky).IncludeSteps(includeSteps).IncludeLabels(includeLabels).ExternalKey(externalKey).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).Execute()



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
	projectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Project internal ID (optional)
	externalId := "externalId_example" // string | Autotest external ID (optional)
	globalId := int64(789) // int64 | Autotest global ID (optional)
	namespace := "namespace_example" // string | Namespace in which autotest is located (optional)
	isNamespaceNull := true // bool | OBSOLETE: Use `includeEmptyNamespaces` instead (optional)
	includeEmptyNamespaces := true // bool | If result must contain autotests without namespace (optional)
	className := "className_example" // string | Name of class in which autotest is located (optional)
	isClassnameNull := true // bool | OBSOLETE: Use `includeEmptyClassNames` instead (optional)
	includeEmptyClassNames := true // bool | If result must contain autotests without class (optional)
	isDeleted := true // bool | OBSOLETE: Use `deleted` instead (optional)
	deleted := true // bool | Is autotest deleted (optional)
	labels := []string{"Inner_example"} // []string | Include only autotests with provided labels (optional)
	stabilityMinimal := int32(56) // int32 | OBSOLETE: Use `minStability` instead (optional)
	minStability := int32(56) // int32 | Minimum stability value of autotest (optional)
	stabilityMaximal := int32(56) // int32 | OBSOLETE: Use `maxStability` instead (optional)
	maxStability := int32(56) // int32 | Maximum stability value of autotest (optional)
	isFlaky := true // bool | OBSOLETE: Use `flaky` instead (optional)
	flaky := true // bool | Is autotest marked as \"Flaky\" (optional)
	includeSteps := true // bool | If result must also include autotest steps (optional)
	includeLabels := true // bool | If result must also include autotest labels (optional)
	externalKey := "externalKey_example" // string | External key of autotest (optional)
	skip := int32(56) // int32 | Amount of items to be skipped (offset) (optional)
	take := int32(56) // int32 | Amount of items to be taken (limit) (optional)
	orderBy := "orderBy_example" // string | SQL-like  ORDER BY statement (column1 ASC|DESC , column2 ASC|DESC) (optional)
	searchField := "searchField_example" // string | Property name for searching (optional)
	searchValue := "searchValue_example" // string | Value for searching (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutoTestsAPI.GetAllAutoTests(context.Background()).ProjectId(projectId).ExternalId(externalId).GlobalId(globalId).Namespace(namespace).IsNamespaceNull(isNamespaceNull).IncludeEmptyNamespaces(includeEmptyNamespaces).ClassName(className).IsClassnameNull(isClassnameNull).IncludeEmptyClassNames(includeEmptyClassNames).IsDeleted(isDeleted).Deleted(deleted).Labels(labels).StabilityMinimal(stabilityMinimal).MinStability(minStability).StabilityMaximal(stabilityMaximal).MaxStability(maxStability).IsFlaky(isFlaky).Flaky(flaky).IncludeSteps(includeSteps).IncludeLabels(includeLabels).ExternalKey(externalKey).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoTestsAPI.GetAllAutoTests``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllAutoTests`: []AutoTestModel
	fmt.Fprintf(os.Stdout, "Response from `AutoTestsAPI.GetAllAutoTests`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllAutoTestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectId** | **string** | Project internal ID | 
 **externalId** | **string** | Autotest external ID | 
 **globalId** | **int64** | Autotest global ID | 
 **namespace** | **string** | Namespace in which autotest is located | 
 **isNamespaceNull** | **bool** | OBSOLETE: Use &#x60;includeEmptyNamespaces&#x60; instead | 
 **includeEmptyNamespaces** | **bool** | If result must contain autotests without namespace | 
 **className** | **string** | Name of class in which autotest is located | 
 **isClassnameNull** | **bool** | OBSOLETE: Use &#x60;includeEmptyClassNames&#x60; instead | 
 **includeEmptyClassNames** | **bool** | If result must contain autotests without class | 
 **isDeleted** | **bool** | OBSOLETE: Use &#x60;deleted&#x60; instead | 
 **deleted** | **bool** | Is autotest deleted | 
 **labels** | **[]string** | Include only autotests with provided labels | 
 **stabilityMinimal** | **int32** | OBSOLETE: Use &#x60;minStability&#x60; instead | 
 **minStability** | **int32** | Minimum stability value of autotest | 
 **stabilityMaximal** | **int32** | OBSOLETE: Use &#x60;maxStability&#x60; instead | 
 **maxStability** | **int32** | Maximum stability value of autotest | 
 **isFlaky** | **bool** | OBSOLETE: Use &#x60;flaky&#x60; instead | 
 **flaky** | **bool** | Is autotest marked as \&quot;Flaky\&quot; | 
 **includeSteps** | **bool** | If result must also include autotest steps | 
 **includeLabels** | **bool** | If result must also include autotest labels | 
 **externalKey** | **string** | External key of autotest | 
 **skip** | **int32** | Amount of items to be skipped (offset) | 
 **take** | **int32** | Amount of items to be taken (limit) | 
 **orderBy** | **string** | SQL-like  ORDER BY statement (column1 ASC|DESC , column2 ASC|DESC) | 
 **searchField** | **string** | Property name for searching | 
 **searchValue** | **string** | Value for searching | 

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


## GetAutoTestAverageDuration

> AutoTestAverageDurationApiResult GetAutoTestAverageDuration(ctx, id).Execute()

Get average autotest duration



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
	id := "id_example" // string | Autotest internal (UUID) or global (integer) identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutoTestsAPI.GetAutoTestAverageDuration(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoTestsAPI.GetAutoTestAverageDuration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAutoTestAverageDuration`: AutoTestAverageDurationApiResult
	fmt.Fprintf(os.Stdout, "Response from `AutoTestsAPI.GetAutoTestAverageDuration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Autotest internal (UUID) or global (integer) identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAutoTestAverageDurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AutoTestAverageDurationApiResult**](AutoTestAverageDurationApiResult.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAutoTestById

> AutoTestApiResult GetAutoTestById(ctx, id).Execute()

Get autotest by internal or global ID



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
	id := "id_example" // string | Autotest internal (UUID) or global (integer) identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutoTestsAPI.GetAutoTestById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoTestsAPI.GetAutoTestById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAutoTestById`: AutoTestApiResult
	fmt.Fprintf(os.Stdout, "Response from `AutoTestsAPI.GetAutoTestById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Autotest internal (UUID) or global (integer) identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAutoTestByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AutoTestApiResult**](AutoTestApiResult.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAutoTestChronology

> []TestResultChronologyModel GetAutoTestChronology(ctx, id).Execute()

Get autotest chronology



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
	id := "id_example" // string | Autotest internal (UUID) or global (integer) identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutoTestsAPI.GetAutoTestChronology(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoTestsAPI.GetAutoTestChronology``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAutoTestChronology`: []TestResultChronologyModel
	fmt.Fprintf(os.Stdout, "Response from `AutoTestsAPI.GetAutoTestChronology`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Autotest internal (UUID) or global (integer) identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAutoTestChronologyRequest struct via the builder pattern


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


## GetTestRuns

> []TestRunByAutoTestApiResult GetTestRuns(ctx, id).Execute()

Get completed tests runs for autotests



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
	id := "id_example" // string | Autotest internal (UUID) or global (integer) identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutoTestsAPI.GetTestRuns(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoTestsAPI.GetTestRuns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTestRuns`: []TestRunByAutoTestApiResult
	fmt.Fprintf(os.Stdout, "Response from `AutoTestsAPI.GetTestRuns`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Autotest internal (UUID) or global (integer) identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTestRunsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]TestRunByAutoTestApiResult**](TestRunByAutoTestApiResult.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWorkItemsLinkedToAutoTest

> []AutoTestWorkItemIdentifierApiResult GetWorkItemsLinkedToAutoTest(ctx, id).IsDeleted(isDeleted).IsWorkItemDeleted(isWorkItemDeleted).Execute()

Get work items linked to autotest



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
	id := "id_example" // string | Specifies the autotest entity ID.  You can copy it from the address bar in your web browser or use autotest GUID.
	isDeleted := true // bool | Specifies that a test is deleted or still relevant. (optional)
	isWorkItemDeleted := true // bool | OBSOLETE: Use `isDeleted` instead (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutoTestsAPI.GetWorkItemsLinkedToAutoTest(context.Background(), id).IsDeleted(isDeleted).IsWorkItemDeleted(isWorkItemDeleted).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoTestsAPI.GetWorkItemsLinkedToAutoTest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWorkItemsLinkedToAutoTest`: []AutoTestWorkItemIdentifierApiResult
	fmt.Fprintf(os.Stdout, "Response from `AutoTestsAPI.GetWorkItemsLinkedToAutoTest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Specifies the autotest entity ID.  You can copy it from the address bar in your web browser or use autotest GUID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkItemsLinkedToAutoTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **isDeleted** | **bool** | Specifies that a test is deleted or still relevant. | 
 **isWorkItemDeleted** | **bool** | OBSOLETE: Use &#x60;isDeleted&#x60; instead | [default to false]

### Return type

[**[]AutoTestWorkItemIdentifierApiResult**](AutoTestWorkItemIdentifierApiResult.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LinkAutoTestToWorkItem

> LinkAutoTestToWorkItem(ctx, id).WorkItemIdApiModel(workItemIdApiModel).Execute()

Link autotest with work items



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
	id := "id_example" // string | Autotest internal (UUID) or global (integer) identifier
	workItemIdApiModel := *openapiclient.NewWorkItemIdApiModel("Id_example") // WorkItemIdApiModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AutoTestsAPI.LinkAutoTestToWorkItem(context.Background(), id).WorkItemIdApiModel(workItemIdApiModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoTestsAPI.LinkAutoTestToWorkItem``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Autotest internal (UUID) or global (integer) identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiLinkAutoTestToWorkItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **workItemIdApiModel** | [**WorkItemIdApiModel**](WorkItemIdApiModel.md) |  | 

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


## UpdateAutoTest

> UpdateAutoTest(ctx).AutoTestUpdateApiModel(autoTestUpdateApiModel).Execute()

Update autotest



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
	autoTestUpdateApiModel := *openapiclient.NewAutoTestUpdateApiModel("ProjectId_example", "ExternalId_example", "Name_example") // AutoTestUpdateApiModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AutoTestsAPI.UpdateAutoTest(context.Background()).AutoTestUpdateApiModel(autoTestUpdateApiModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoTestsAPI.UpdateAutoTest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAutoTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **autoTestUpdateApiModel** | [**AutoTestUpdateApiModel**](AutoTestUpdateApiModel.md) |  | 

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


## UpdateMultiple

> UpdateMultiple(ctx).AutoTestUpdateApiModel(autoTestUpdateApiModel).Execute()

Update multiple autotests



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
	autoTestUpdateApiModel := []openapiclient.AutoTestUpdateApiModel{*openapiclient.NewAutoTestUpdateApiModel("ProjectId_example", "ExternalId_example", "Name_example")} // []AutoTestUpdateApiModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AutoTestsAPI.UpdateMultiple(context.Background()).AutoTestUpdateApiModel(autoTestUpdateApiModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoTestsAPI.UpdateMultiple``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMultipleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **autoTestUpdateApiModel** | [**[]AutoTestUpdateApiModel**](AutoTestUpdateApiModel.md) |  | 

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

