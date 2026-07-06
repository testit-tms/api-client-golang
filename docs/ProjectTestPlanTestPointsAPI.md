# \ProjectTestPlanTestPointsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV2ProjectsProjectIdTestPlansTestPlanIdTestPointsAnalyticsPost**](ProjectTestPlanTestPointsAPI.md#ApiV2ProjectsProjectIdTestPlansTestPlanIdTestPointsAnalyticsPost) | **Post** /api/v2/projects/{projectId}/test-plans/{testPlanId}/test-points/analytics | Get test points analytics.
[**ApiV2ProjectsProjectIdTestPlansTestPlanIdTestPointsAutotestsRerunPost**](ProjectTestPlanTestPointsAPI.md#ApiV2ProjectsProjectIdTestPlansTestPlanIdTestPointsAutotestsRerunPost) | **Post** /api/v2/projects/{projectId}/test-plans/{testPlanId}/test-points/autotests/rerun | Rerun autotests.
[**ApiV2ProjectsProjectIdTestPlansTestPlanIdTestPointsAutotestsRunPost**](ProjectTestPlanTestPointsAPI.md#ApiV2ProjectsProjectIdTestPlansTestPlanIdTestPointsAutotestsRunPost) | **Post** /api/v2/projects/{projectId}/test-plans/{testPlanId}/test-points/autotests/run | Run autotests.
[**ApiV2ProjectsProjectIdTestPlansTestPlanIdTestPointsGroupingSearchPost**](ProjectTestPlanTestPointsAPI.md#ApiV2ProjectsProjectIdTestPlansTestPlanIdTestPointsGroupingSearchPost) | **Post** /api/v2/projects/{projectId}/test-plans/{testPlanId}/test-points/grouping-search | Search test points in test plan.
[**ApiV2ProjectsProjectIdTestPlansTestPlanIdTestPointsTestersPost**](ProjectTestPlanTestPointsAPI.md#ApiV2ProjectsProjectIdTestPlansTestPlanIdTestPointsTestersPost) | **Post** /api/v2/projects/{projectId}/test-plans/{testPlanId}/test-points/testers | Distribute test points between the users.



## ApiV2ProjectsProjectIdTestPlansTestPlanIdTestPointsAnalyticsPost

> TestPlanTestPointsAnalyticsApiResult ApiV2ProjectsProjectIdTestPlansTestPlanIdTestPointsAnalyticsPost(ctx, projectId, testPlanId).TestPlanTestPointsAnalyticsApiModel(testPlanTestPointsAnalyticsApiModel).Execute()

Get test points analytics.

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
	testPlanId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	testPlanTestPointsAnalyticsApiModel := *openapiclient.NewTestPlanTestPointsAnalyticsApiModel() // TestPlanTestPointsAnalyticsApiModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectTestPlanTestPointsAPI.ApiV2ProjectsProjectIdTestPlansTestPlanIdTestPointsAnalyticsPost(context.Background(), projectId, testPlanId).TestPlanTestPointsAnalyticsApiModel(testPlanTestPointsAnalyticsApiModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectTestPlanTestPointsAPI.ApiV2ProjectsProjectIdTestPlansTestPlanIdTestPointsAnalyticsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2ProjectsProjectIdTestPlansTestPlanIdTestPointsAnalyticsPost`: TestPlanTestPointsAnalyticsApiResult
	fmt.Fprintf(os.Stdout, "Response from `ProjectTestPlanTestPointsAPI.ApiV2ProjectsProjectIdTestPlansTestPlanIdTestPointsAnalyticsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Internal (UUID) or global (integer) identifier | 
**testPlanId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ProjectsProjectIdTestPlansTestPlanIdTestPointsAnalyticsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **testPlanTestPointsAnalyticsApiModel** | [**TestPlanTestPointsAnalyticsApiModel**](TestPlanTestPointsAnalyticsApiModel.md) |  | 

### Return type

[**TestPlanTestPointsAnalyticsApiResult**](TestPlanTestPointsAnalyticsApiResult.md)

### Authorization

[PrivateToken](../README.md#PrivateToken), [Cookies](../README.md#Cookies)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2ProjectsProjectIdTestPlansTestPlanIdTestPointsAutotestsRerunPost

> ApiV2ProjectsProjectIdTestPlansTestPlanIdTestPointsAutotestsRerunPost(ctx, projectId, testPlanId).TestPlanTestPointsAutoTestsRerunApiModel(testPlanTestPointsAutoTestsRerunApiModel).Execute()

Rerun autotests.

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
	testPlanId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	testPlanTestPointsAutoTestsRerunApiModel := *openapiclient.NewTestPlanTestPointsAutoTestsRerunApiModel() // TestPlanTestPointsAutoTestsRerunApiModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectTestPlanTestPointsAPI.ApiV2ProjectsProjectIdTestPlansTestPlanIdTestPointsAutotestsRerunPost(context.Background(), projectId, testPlanId).TestPlanTestPointsAutoTestsRerunApiModel(testPlanTestPointsAutoTestsRerunApiModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectTestPlanTestPointsAPI.ApiV2ProjectsProjectIdTestPlansTestPlanIdTestPointsAutotestsRerunPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Internal (UUID) or global (integer) identifier | 
**testPlanId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ProjectsProjectIdTestPlansTestPlanIdTestPointsAutotestsRerunPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **testPlanTestPointsAutoTestsRerunApiModel** | [**TestPlanTestPointsAutoTestsRerunApiModel**](TestPlanTestPointsAutoTestsRerunApiModel.md) |  | 

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


## ApiV2ProjectsProjectIdTestPlansTestPlanIdTestPointsAutotestsRunPost

> TestRunNameApiResult ApiV2ProjectsProjectIdTestPlansTestPlanIdTestPointsAutotestsRunPost(ctx, projectId, testPlanId).TestPlanTestPointsAutoTestsRunApiModel(testPlanTestPointsAutoTestsRunApiModel).Execute()

Run autotests.

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
	testPlanId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	testPlanTestPointsAutoTestsRunApiModel := *openapiclient.NewTestPlanTestPointsAutoTestsRunApiModel([]string{"WebhookIds_example"}, false) // TestPlanTestPointsAutoTestsRunApiModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectTestPlanTestPointsAPI.ApiV2ProjectsProjectIdTestPlansTestPlanIdTestPointsAutotestsRunPost(context.Background(), projectId, testPlanId).TestPlanTestPointsAutoTestsRunApiModel(testPlanTestPointsAutoTestsRunApiModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectTestPlanTestPointsAPI.ApiV2ProjectsProjectIdTestPlansTestPlanIdTestPointsAutotestsRunPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2ProjectsProjectIdTestPlansTestPlanIdTestPointsAutotestsRunPost`: TestRunNameApiResult
	fmt.Fprintf(os.Stdout, "Response from `ProjectTestPlanTestPointsAPI.ApiV2ProjectsProjectIdTestPlansTestPlanIdTestPointsAutotestsRunPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Internal (UUID) or global (integer) identifier | 
**testPlanId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ProjectsProjectIdTestPlansTestPlanIdTestPointsAutotestsRunPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **testPlanTestPointsAutoTestsRunApiModel** | [**TestPlanTestPointsAutoTestsRunApiModel**](TestPlanTestPointsAutoTestsRunApiModel.md) |  | 

### Return type

[**TestRunNameApiResult**](TestRunNameApiResult.md)

### Authorization

[PrivateToken](../README.md#PrivateToken), [Cookies](../README.md#Cookies)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2ProjectsProjectIdTestPlansTestPlanIdTestPointsGroupingSearchPost

> TestPlanTestPointsGroupSearchApiResult ApiV2ProjectsProjectIdTestPlansTestPlanIdTestPointsGroupingSearchPost(ctx, projectId, testPlanId).TestPlanTestPointsApiModel(testPlanTestPointsApiModel).Execute()

Search test points in test plan.

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
	testPlanId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	testPlanTestPointsApiModel := *openapiclient.NewTestPlanTestPointsApiModel() // TestPlanTestPointsApiModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectTestPlanTestPointsAPI.ApiV2ProjectsProjectIdTestPlansTestPlanIdTestPointsGroupingSearchPost(context.Background(), projectId, testPlanId).TestPlanTestPointsApiModel(testPlanTestPointsApiModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectTestPlanTestPointsAPI.ApiV2ProjectsProjectIdTestPlansTestPlanIdTestPointsGroupingSearchPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2ProjectsProjectIdTestPlansTestPlanIdTestPointsGroupingSearchPost`: TestPlanTestPointsGroupSearchApiResult
	fmt.Fprintf(os.Stdout, "Response from `ProjectTestPlanTestPointsAPI.ApiV2ProjectsProjectIdTestPlansTestPlanIdTestPointsGroupingSearchPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Internal (UUID) or global (integer) identifier | 
**testPlanId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ProjectsProjectIdTestPlansTestPlanIdTestPointsGroupingSearchPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **testPlanTestPointsApiModel** | [**TestPlanTestPointsApiModel**](TestPlanTestPointsApiModel.md) |  | 

### Return type

[**TestPlanTestPointsGroupSearchApiResult**](TestPlanTestPointsGroupSearchApiResult.md)

### Authorization

[PrivateToken](../README.md#PrivateToken), [Cookies](../README.md#Cookies)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2ProjectsProjectIdTestPlansTestPlanIdTestPointsTestersPost

> ApiV2ProjectsProjectIdTestPlansTestPlanIdTestPointsTestersPost(ctx, projectId, testPlanId).TestPlanTestPointsSetTestersApiModel(testPlanTestPointsSetTestersApiModel).Execute()

Distribute test points between the users.

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
	testPlanId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	testPlanTestPointsSetTestersApiModel := *openapiclient.NewTestPlanTestPointsSetTestersApiModel([]string{"TesterIds_example"}) // TestPlanTestPointsSetTestersApiModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectTestPlanTestPointsAPI.ApiV2ProjectsProjectIdTestPlansTestPlanIdTestPointsTestersPost(context.Background(), projectId, testPlanId).TestPlanTestPointsSetTestersApiModel(testPlanTestPointsSetTestersApiModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectTestPlanTestPointsAPI.ApiV2ProjectsProjectIdTestPlansTestPlanIdTestPointsTestersPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Internal (UUID) or global (integer) identifier | 
**testPlanId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ProjectsProjectIdTestPlansTestPlanIdTestPointsTestersPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **testPlanTestPointsSetTestersApiModel** | [**TestPlanTestPointsSetTestersApiModel**](TestPlanTestPointsSetTestersApiModel.md) |  | 

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

