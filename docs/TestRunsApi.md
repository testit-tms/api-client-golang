# \TestRunsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV2TestRunsDelete**](TestRunsAPI.md#ApiV2TestRunsDelete) | **Delete** /api/v2/testRuns | Delete multiple test runs
[**ApiV2TestRunsIdAutoTestsNamespacesGet**](TestRunsAPI.md#ApiV2TestRunsIdAutoTestsNamespacesGet) | **Get** /api/v2/testRuns/{id}/autoTestsNamespaces | Get autotest classes and namespaces in test run
[**ApiV2TestRunsIdDelete**](TestRunsAPI.md#ApiV2TestRunsIdDelete) | **Delete** /api/v2/testRuns/{id} | Delete test run
[**ApiV2TestRunsIdPurgePost**](TestRunsAPI.md#ApiV2TestRunsIdPurgePost) | **Post** /api/v2/testRuns/{id}/purge | Permanently delete test run from archive
[**ApiV2TestRunsIdRerunsPost**](TestRunsAPI.md#ApiV2TestRunsIdRerunsPost) | **Post** /api/v2/testRuns/{id}/reruns | Manual autotests rerun in test run
[**ApiV2TestRunsIdRestorePost**](TestRunsAPI.md#ApiV2TestRunsIdRestorePost) | **Post** /api/v2/testRuns/{id}/restore | Restore test run from the archive
[**ApiV2TestRunsIdStatisticsFilterPost**](TestRunsAPI.md#ApiV2TestRunsIdStatisticsFilterPost) | **Post** /api/v2/testRuns/{id}/statistics/filter | Search for the test run test results and build statistics
[**ApiV2TestRunsIdTestPointsResultsGet**](TestRunsAPI.md#ApiV2TestRunsIdTestPointsResultsGet) | **Get** /api/v2/testRuns/{id}/testPoints/results | Get test results from the test run grouped by test points
[**ApiV2TestRunsIdTestResultsBulkPut**](TestRunsAPI.md#ApiV2TestRunsIdTestResultsBulkPut) | **Put** /api/v2/testRuns/{id}/testResults/bulk | Partial edit of multiple test results in the test run
[**ApiV2TestRunsIdTestResultsLastModifiedModificationDateGet**](TestRunsAPI.md#ApiV2TestRunsIdTestResultsLastModifiedModificationDateGet) | **Get** /api/v2/testRuns/{id}/testResults/lastModified/modificationDate | Get modification date of last test result of the test run
[**ApiV2TestRunsPurgeBulkPost**](TestRunsAPI.md#ApiV2TestRunsPurgeBulkPost) | **Post** /api/v2/testRuns/purge/bulk | Permanently delete multiple test runs from archive
[**ApiV2TestRunsRestoreBulkPost**](TestRunsAPI.md#ApiV2TestRunsRestoreBulkPost) | **Post** /api/v2/testRuns/restore/bulk | Restore multiple test runs from the archive
[**ApiV2TestRunsSearchPost**](TestRunsAPI.md#ApiV2TestRunsSearchPost) | **Post** /api/v2/testRuns/search | Search for test runs
[**ApiV2TestRunsUpdateMultiplePost**](TestRunsAPI.md#ApiV2TestRunsUpdateMultiplePost) | **Post** /api/v2/testRuns/updateMultiple | Update multiple test runs
[**CompleteTestRun**](TestRunsAPI.md#CompleteTestRun) | **Post** /api/v2/testRuns/{id}/complete | Complete TestRun
[**CreateAndFillByAutoTests**](TestRunsAPI.md#CreateAndFillByAutoTests) | **Post** /api/v2/testRuns/byAutoTests | Create test runs based on autotests and configurations
[**CreateAndFillByConfigurations**](TestRunsAPI.md#CreateAndFillByConfigurations) | **Post** /api/v2/testRuns/byConfigurations | Create test runs picking the needed test points
[**CreateAndFillByWorkItems**](TestRunsAPI.md#CreateAndFillByWorkItems) | **Post** /api/v2/testRuns/byWorkItems | Create test run based on configurations and work items
[**CreateEmpty**](TestRunsAPI.md#CreateEmpty) | **Post** /api/v2/testRuns | Create empty TestRun
[**GetTestRunById**](TestRunsAPI.md#GetTestRunById) | **Get** /api/v2/testRuns/{id} | Get TestRun by Id
[**SetAutoTestResultsForTestRun**](TestRunsAPI.md#SetAutoTestResultsForTestRun) | **Post** /api/v2/testRuns/{id}/testResults | Send test results to the test runs in the system
[**StartTestRun**](TestRunsAPI.md#StartTestRun) | **Post** /api/v2/testRuns/{id}/start | Start TestRun
[**StopTestRun**](TestRunsAPI.md#StopTestRun) | **Post** /api/v2/testRuns/{id}/stop | Stop TestRun
[**UpdateEmpty**](TestRunsAPI.md#UpdateEmpty) | **Put** /api/v2/testRuns | Update empty TestRun



## ApiV2TestRunsDelete

> int32 ApiV2TestRunsDelete(ctx).TestRunSelectApiModel(testRunSelectApiModel).Execute()

Delete multiple test runs



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
	testRunSelectApiModel := *openapiclient.NewTestRunSelectApiModel(*openapiclient.NewTestRunFilterApiModel(), *openapiclient.NewTestRunExtractionApiModel()) // TestRunSelectApiModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TestRunsAPI.ApiV2TestRunsDelete(context.Background()).TestRunSelectApiModel(testRunSelectApiModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestRunsAPI.ApiV2TestRunsDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2TestRunsDelete`: int32
	fmt.Fprintf(os.Stdout, "Response from `TestRunsAPI.ApiV2TestRunsDelete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2TestRunsDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **testRunSelectApiModel** | [**TestRunSelectApiModel**](TestRunSelectApiModel.md) |  | 

### Return type

**int32**

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2TestRunsIdAutoTestsNamespacesGet

> AutoTestNamespacesCountResponse ApiV2TestRunsIdAutoTestsNamespacesGet(ctx, id).Execute()

Get autotest classes and namespaces in test run

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
	resp, r, err := apiClient.TestRunsAPI.ApiV2TestRunsIdAutoTestsNamespacesGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestRunsAPI.ApiV2TestRunsIdAutoTestsNamespacesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2TestRunsIdAutoTestsNamespacesGet`: AutoTestNamespacesCountResponse
	fmt.Fprintf(os.Stdout, "Response from `TestRunsAPI.ApiV2TestRunsIdAutoTestsNamespacesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2TestRunsIdAutoTestsNamespacesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AutoTestNamespacesCountResponse**](AutoTestNamespacesCountResponse.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2TestRunsIdDelete

> ApiV2TestRunsIdDelete(ctx, id).Execute()

Delete test run



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Test run internal (UUID) identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TestRunsAPI.ApiV2TestRunsIdDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestRunsAPI.ApiV2TestRunsIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Test run internal (UUID) identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2TestRunsIdDeleteRequest struct via the builder pattern


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


## ApiV2TestRunsIdPurgePost

> ApiV2TestRunsIdPurgePost(ctx, id).Execute()

Permanently delete test run from archive



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Test run internal (UUID) identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TestRunsAPI.ApiV2TestRunsIdPurgePost(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestRunsAPI.ApiV2TestRunsIdPurgePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Test run internal (UUID) identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2TestRunsIdPurgePostRequest struct via the builder pattern


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


## ApiV2TestRunsIdRerunsPost

> ManualRerunApiResult ApiV2TestRunsIdRerunsPost(ctx, id).ManualRerunSelectTestResultsApiModel(manualRerunSelectTestResultsApiModel).Execute()

Manual autotests rerun in test run

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
	manualRerunSelectTestResultsApiModel := *openapiclient.NewManualRerunSelectTestResultsApiModel() // ManualRerunSelectTestResultsApiModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TestRunsAPI.ApiV2TestRunsIdRerunsPost(context.Background(), id).ManualRerunSelectTestResultsApiModel(manualRerunSelectTestResultsApiModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestRunsAPI.ApiV2TestRunsIdRerunsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2TestRunsIdRerunsPost`: ManualRerunApiResult
	fmt.Fprintf(os.Stdout, "Response from `TestRunsAPI.ApiV2TestRunsIdRerunsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2TestRunsIdRerunsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **manualRerunSelectTestResultsApiModel** | [**ManualRerunSelectTestResultsApiModel**](ManualRerunSelectTestResultsApiModel.md) |  | 

### Return type

[**ManualRerunApiResult**](ManualRerunApiResult.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2TestRunsIdRestorePost

> ApiV2TestRunsIdRestorePost(ctx, id).Execute()

Restore test run from the archive



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique ID of the test run

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TestRunsAPI.ApiV2TestRunsIdRestorePost(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestRunsAPI.ApiV2TestRunsIdRestorePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Unique ID of the test run | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2TestRunsIdRestorePostRequest struct via the builder pattern


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


## ApiV2TestRunsIdStatisticsFilterPost

> TestResultsStatisticsApiResult ApiV2TestRunsIdStatisticsFilterPost(ctx, id).TestRunStatisticsFilterApiModel(testRunStatisticsFilterApiModel).Execute()

Search for the test run test results and build statistics

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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Test run unique ID
	testRunStatisticsFilterApiModel := *openapiclient.NewTestRunStatisticsFilterApiModel() // TestRunStatisticsFilterApiModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TestRunsAPI.ApiV2TestRunsIdStatisticsFilterPost(context.Background(), id).TestRunStatisticsFilterApiModel(testRunStatisticsFilterApiModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestRunsAPI.ApiV2TestRunsIdStatisticsFilterPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2TestRunsIdStatisticsFilterPost`: TestResultsStatisticsApiResult
	fmt.Fprintf(os.Stdout, "Response from `TestRunsAPI.ApiV2TestRunsIdStatisticsFilterPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Test run unique ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2TestRunsIdStatisticsFilterPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **testRunStatisticsFilterApiModel** | [**TestRunStatisticsFilterApiModel**](TestRunStatisticsFilterApiModel.md) |  | 

### Return type

[**TestResultsStatisticsApiResult**](TestResultsStatisticsApiResult.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2TestRunsIdTestPointsResultsGet

> []TestPointResultApiResult ApiV2TestRunsIdTestPointsResultsGet(ctx, id).Execute()

Get test results from the test run grouped by test points

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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Test run unique ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TestRunsAPI.ApiV2TestRunsIdTestPointsResultsGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestRunsAPI.ApiV2TestRunsIdTestPointsResultsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2TestRunsIdTestPointsResultsGet`: []TestPointResultApiResult
	fmt.Fprintf(os.Stdout, "Response from `TestRunsAPI.ApiV2TestRunsIdTestPointsResultsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Test run unique ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2TestRunsIdTestPointsResultsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]TestPointResultApiResult**](TestPointResultApiResult.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2TestRunsIdTestResultsBulkPut

> ApiV2TestRunsIdTestResultsBulkPut(ctx, id).TestRunTestResultsPartialBulkSetModel(testRunTestResultsPartialBulkSetModel).Execute()

Partial edit of multiple test results in the test run

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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Test run unique ID
	testRunTestResultsPartialBulkSetModel := *openapiclient.NewTestRunTestResultsPartialBulkSetModel() // TestRunTestResultsPartialBulkSetModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TestRunsAPI.ApiV2TestRunsIdTestResultsBulkPut(context.Background(), id).TestRunTestResultsPartialBulkSetModel(testRunTestResultsPartialBulkSetModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestRunsAPI.ApiV2TestRunsIdTestResultsBulkPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Test run unique ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2TestRunsIdTestResultsBulkPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **testRunTestResultsPartialBulkSetModel** | [**TestRunTestResultsPartialBulkSetModel**](TestRunTestResultsPartialBulkSetModel.md) |  | 

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


## ApiV2TestRunsIdTestResultsLastModifiedModificationDateGet

> time.Time ApiV2TestRunsIdTestResultsLastModifiedModificationDateGet(ctx, id).Execute()

Get modification date of last test result of the test run

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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Test run unique ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TestRunsAPI.ApiV2TestRunsIdTestResultsLastModifiedModificationDateGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestRunsAPI.ApiV2TestRunsIdTestResultsLastModifiedModificationDateGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2TestRunsIdTestResultsLastModifiedModificationDateGet`: time.Time
	fmt.Fprintf(os.Stdout, "Response from `TestRunsAPI.ApiV2TestRunsIdTestResultsLastModifiedModificationDateGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Test run unique ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2TestRunsIdTestResultsLastModifiedModificationDateGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**time.Time**](time.Time.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2TestRunsPurgeBulkPost

> int32 ApiV2TestRunsPurgeBulkPost(ctx).TestRunSelectApiModel(testRunSelectApiModel).Execute()

Permanently delete multiple test runs from archive



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
	testRunSelectApiModel := *openapiclient.NewTestRunSelectApiModel(*openapiclient.NewTestRunFilterApiModel(), *openapiclient.NewTestRunExtractionApiModel()) // TestRunSelectApiModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TestRunsAPI.ApiV2TestRunsPurgeBulkPost(context.Background()).TestRunSelectApiModel(testRunSelectApiModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestRunsAPI.ApiV2TestRunsPurgeBulkPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2TestRunsPurgeBulkPost`: int32
	fmt.Fprintf(os.Stdout, "Response from `TestRunsAPI.ApiV2TestRunsPurgeBulkPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2TestRunsPurgeBulkPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **testRunSelectApiModel** | [**TestRunSelectApiModel**](TestRunSelectApiModel.md) |  | 

### Return type

**int32**

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2TestRunsRestoreBulkPost

> int32 ApiV2TestRunsRestoreBulkPost(ctx).TestRunSelectApiModel(testRunSelectApiModel).Execute()

Restore multiple test runs from the archive



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
	testRunSelectApiModel := *openapiclient.NewTestRunSelectApiModel(*openapiclient.NewTestRunFilterApiModel(), *openapiclient.NewTestRunExtractionApiModel()) // TestRunSelectApiModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TestRunsAPI.ApiV2TestRunsRestoreBulkPost(context.Background()).TestRunSelectApiModel(testRunSelectApiModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestRunsAPI.ApiV2TestRunsRestoreBulkPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2TestRunsRestoreBulkPost`: int32
	fmt.Fprintf(os.Stdout, "Response from `TestRunsAPI.ApiV2TestRunsRestoreBulkPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2TestRunsRestoreBulkPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **testRunSelectApiModel** | [**TestRunSelectApiModel**](TestRunSelectApiModel.md) |  | 

### Return type

**int32**

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2TestRunsSearchPost

> []TestRunShortApiResult ApiV2TestRunsSearchPost(ctx).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).TestRunFilterApiModel(testRunFilterApiModel).Execute()

Search for test runs

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
	testRunFilterApiModel := *openapiclient.NewTestRunFilterApiModel() // TestRunFilterApiModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TestRunsAPI.ApiV2TestRunsSearchPost(context.Background()).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).TestRunFilterApiModel(testRunFilterApiModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestRunsAPI.ApiV2TestRunsSearchPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2TestRunsSearchPost`: []TestRunShortApiResult
	fmt.Fprintf(os.Stdout, "Response from `TestRunsAPI.ApiV2TestRunsSearchPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2TestRunsSearchPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skip** | **int32** | Amount of items to be skipped (offset) | 
 **take** | **int32** | Amount of items to be taken (limit) | 
 **orderBy** | **string** | SQL-like  ORDER BY statement (column1 ASC|DESC , column2 ASC|DESC) | 
 **searchField** | **string** | Property name for searching | 
 **searchValue** | **string** | Value for searching | 
 **testRunFilterApiModel** | [**TestRunFilterApiModel**](TestRunFilterApiModel.md) |  | 

### Return type

[**[]TestRunShortApiResult**](TestRunShortApiResult.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2TestRunsUpdateMultiplePost

> ApiV2TestRunsUpdateMultiplePost(ctx).UpdateMultipleTestRunsApiModel(updateMultipleTestRunsApiModel).Execute()

Update multiple test runs

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
	updateMultipleTestRunsApiModel := *openapiclient.NewUpdateMultipleTestRunsApiModel(*openapiclient.NewTestRunSelectApiModel(*openapiclient.NewTestRunFilterApiModel(), *openapiclient.NewTestRunExtractionApiModel())) // UpdateMultipleTestRunsApiModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TestRunsAPI.ApiV2TestRunsUpdateMultiplePost(context.Background()).UpdateMultipleTestRunsApiModel(updateMultipleTestRunsApiModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestRunsAPI.ApiV2TestRunsUpdateMultiplePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2TestRunsUpdateMultiplePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateMultipleTestRunsApiModel** | [**UpdateMultipleTestRunsApiModel**](UpdateMultipleTestRunsApiModel.md) |  | 

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


## CompleteTestRun

> CompleteTestRun(ctx, id).Execute()

Complete TestRun



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
	id := "3fa85f64-5717-4562-b3fc-2c963f66afa6" // string | Test Run internal identifier (GUID format)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TestRunsAPI.CompleteTestRun(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestRunsAPI.CompleteTestRun``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Test Run internal identifier (GUID format) | 

### Other Parameters

Other parameters are passed through a pointer to a apiCompleteTestRunRequest struct via the builder pattern


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


## CreateAndFillByAutoTests

> TestRunV2ApiResult CreateAndFillByAutoTests(ctx).CreateTestRunAndFillByAutoTestsApiModel(createTestRunAndFillByAutoTestsApiModel).Execute()

Create test runs based on autotests and configurations



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
	createTestRunAndFillByAutoTestsApiModel := *openapiclient.NewCreateTestRunAndFillByAutoTestsApiModel("ProjectId_example", []string{"ConfigurationIds_example"}, []string{"AutoTestExternalIds_example"}) // CreateTestRunAndFillByAutoTestsApiModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TestRunsAPI.CreateAndFillByAutoTests(context.Background()).CreateTestRunAndFillByAutoTestsApiModel(createTestRunAndFillByAutoTestsApiModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestRunsAPI.CreateAndFillByAutoTests``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAndFillByAutoTests`: TestRunV2ApiResult
	fmt.Fprintf(os.Stdout, "Response from `TestRunsAPI.CreateAndFillByAutoTests`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAndFillByAutoTestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createTestRunAndFillByAutoTestsApiModel** | [**CreateTestRunAndFillByAutoTestsApiModel**](CreateTestRunAndFillByAutoTestsApiModel.md) |  | 

### Return type

[**TestRunV2ApiResult**](TestRunV2ApiResult.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAndFillByConfigurations

> TestRunV2ApiResult CreateAndFillByConfigurations(ctx).CreateTestRunAndFillByConfigurationsApiModel(createTestRunAndFillByConfigurationsApiModel).Execute()

Create test runs picking the needed test points



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
	createTestRunAndFillByConfigurationsApiModel := *openapiclient.NewCreateTestRunAndFillByConfigurationsApiModel("ProjectId_example", "TestPlanId_example", []openapiclient.TestPointSelector{*openapiclient.NewTestPointSelector("ConfigurationId_example", []string{"WorkItemIds_example"})}) // CreateTestRunAndFillByConfigurationsApiModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TestRunsAPI.CreateAndFillByConfigurations(context.Background()).CreateTestRunAndFillByConfigurationsApiModel(createTestRunAndFillByConfigurationsApiModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestRunsAPI.CreateAndFillByConfigurations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAndFillByConfigurations`: TestRunV2ApiResult
	fmt.Fprintf(os.Stdout, "Response from `TestRunsAPI.CreateAndFillByConfigurations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAndFillByConfigurationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createTestRunAndFillByConfigurationsApiModel** | [**CreateTestRunAndFillByConfigurationsApiModel**](CreateTestRunAndFillByConfigurationsApiModel.md) |  | 

### Return type

[**TestRunV2ApiResult**](TestRunV2ApiResult.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAndFillByWorkItems

> TestRunV2ApiResult CreateAndFillByWorkItems(ctx).CreateTestRunAndFillByWorkItemsApiModel(createTestRunAndFillByWorkItemsApiModel).Execute()

Create test run based on configurations and work items



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
	createTestRunAndFillByWorkItemsApiModel := *openapiclient.NewCreateTestRunAndFillByWorkItemsApiModel("ProjectId_example", "TestPlanId_example", []string{"ConfigurationIds_example"}, []string{"WorkItemIds_example"}) // CreateTestRunAndFillByWorkItemsApiModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TestRunsAPI.CreateAndFillByWorkItems(context.Background()).CreateTestRunAndFillByWorkItemsApiModel(createTestRunAndFillByWorkItemsApiModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestRunsAPI.CreateAndFillByWorkItems``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAndFillByWorkItems`: TestRunV2ApiResult
	fmt.Fprintf(os.Stdout, "Response from `TestRunsAPI.CreateAndFillByWorkItems`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAndFillByWorkItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createTestRunAndFillByWorkItemsApiModel** | [**CreateTestRunAndFillByWorkItemsApiModel**](CreateTestRunAndFillByWorkItemsApiModel.md) |  | 

### Return type

[**TestRunV2ApiResult**](TestRunV2ApiResult.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateEmpty

> TestRunV2ApiResult CreateEmpty(ctx).CreateEmptyTestRunApiModel(createEmptyTestRunApiModel).Execute()

Create empty TestRun



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
	createEmptyTestRunApiModel := *openapiclient.NewCreateEmptyTestRunApiModel("ProjectId_example") // CreateEmptyTestRunApiModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TestRunsAPI.CreateEmpty(context.Background()).CreateEmptyTestRunApiModel(createEmptyTestRunApiModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestRunsAPI.CreateEmpty``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateEmpty`: TestRunV2ApiResult
	fmt.Fprintf(os.Stdout, "Response from `TestRunsAPI.CreateEmpty`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateEmptyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createEmptyTestRunApiModel** | [**CreateEmptyTestRunApiModel**](CreateEmptyTestRunApiModel.md) |  | 

### Return type

[**TestRunV2ApiResult**](TestRunV2ApiResult.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTestRunById

> TestRunV2ApiResult GetTestRunById(ctx, id).Execute()

Get TestRun by Id



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
	id := "3fa85f64-5717-4562-b3fc-2c963f66afa6" // string | Test Run internal identifier (GUID format)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TestRunsAPI.GetTestRunById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestRunsAPI.GetTestRunById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTestRunById`: TestRunV2ApiResult
	fmt.Fprintf(os.Stdout, "Response from `TestRunsAPI.GetTestRunById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Test Run internal identifier (GUID format) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTestRunByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TestRunV2ApiResult**](TestRunV2ApiResult.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetAutoTestResultsForTestRun

> []string SetAutoTestResultsForTestRun(ctx, id).AutoTestResultsForTestRunModel(autoTestResultsForTestRunModel).Execute()

Send test results to the test runs in the system



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Test Run internal identifier (GUID format)
	autoTestResultsForTestRunModel := []openapiclient.AutoTestResultsForTestRunModel{*openapiclient.NewAutoTestResultsForTestRunModel("ConfigurationId_example", "AutoTestExternalId_example")} // []AutoTestResultsForTestRunModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TestRunsAPI.SetAutoTestResultsForTestRun(context.Background(), id).AutoTestResultsForTestRunModel(autoTestResultsForTestRunModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestRunsAPI.SetAutoTestResultsForTestRun``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetAutoTestResultsForTestRun`: []string
	fmt.Fprintf(os.Stdout, "Response from `TestRunsAPI.SetAutoTestResultsForTestRun`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Test Run internal identifier (GUID format) | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetAutoTestResultsForTestRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **autoTestResultsForTestRunModel** | [**[]AutoTestResultsForTestRunModel**](AutoTestResultsForTestRunModel.md) |  | 

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


## StartTestRun

> StartTestRun(ctx, id).Execute()

Start TestRun



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
	id := "3fa85f64-5717-4562-b3fc-2c963f66afa6" // string | Test Run internal identifier (GUID format)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TestRunsAPI.StartTestRun(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestRunsAPI.StartTestRun``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Test Run internal identifier (GUID format) | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartTestRunRequest struct via the builder pattern


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


## StopTestRun

> StopTestRun(ctx, id).Execute()

Stop TestRun



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
	id := "3fa85f64-5717-4562-b3fc-2c963f66afa6" // string | Test Run internal identifier (GUID format)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TestRunsAPI.StopTestRun(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestRunsAPI.StopTestRun``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Test Run internal identifier (GUID format) | 

### Other Parameters

Other parameters are passed through a pointer to a apiStopTestRunRequest struct via the builder pattern


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


## UpdateEmpty

> UpdateEmpty(ctx).UpdateEmptyTestRunApiModel(updateEmptyTestRunApiModel).Execute()

Update empty TestRun



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
	updateEmptyTestRunApiModel := *openapiclient.NewUpdateEmptyTestRunApiModel("Id_example", "Name_example") // UpdateEmptyTestRunApiModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TestRunsAPI.UpdateEmpty(context.Background()).UpdateEmptyTestRunApiModel(updateEmptyTestRunApiModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestRunsAPI.UpdateEmpty``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEmptyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateEmptyTestRunApiModel** | [**UpdateEmptyTestRunApiModel**](UpdateEmptyTestRunApiModel.md) |  | 

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

