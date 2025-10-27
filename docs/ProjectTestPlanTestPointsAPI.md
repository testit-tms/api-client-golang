# \ProjectTestPlanTestPointsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV2ProjectsProjectIdTestPlansTestPlanIdTestPointsAutotestsRerunPost**](ProjectTestPlanTestPointsAPI.md#ApiV2ProjectsProjectIdTestPlansTestPlanIdTestPointsAutotestsRerunPost) | **Post** /api/v2/projects/{projectId}/test-plans/{testPlanId}/test-points/autotests/rerun | Rerun autotests.
[**ApiV2ProjectsProjectIdTestPlansTestPlanIdTestPointsAutotestsRunPost**](ProjectTestPlanTestPointsAPI.md#ApiV2ProjectsProjectIdTestPlansTestPlanIdTestPointsAutotestsRunPost) | **Post** /api/v2/projects/{projectId}/test-plans/{testPlanId}/test-points/autotests/run | Run autotests.



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

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

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

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

