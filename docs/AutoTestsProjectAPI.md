# \AutoTestsProjectAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV2ProjectsProjectIdAutoTestsImportPost**](AutoTestsProjectAPI.md#ApiV2ProjectsProjectIdAutoTestsImportPost) | **Post** /api/v2/projects/{projectId}/auto-tests/import | Schedules import of auto tests JUnit XML report.



## ApiV2ProjectsProjectIdAutoTestsImportPost

> ApiV2ProjectsProjectIdAutoTestsImportPost(ctx, projectId).ScheduleAutoTestsReportImportsApiModel(scheduleAutoTestsReportImportsApiModel).Execute()

Schedules import of auto tests JUnit XML report.

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
	projectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Project ID
	scheduleAutoTestsReportImportsApiModel := *openapiclient.NewScheduleAutoTestsReportImportsApiModel([]openapiclient.ScheduleAutoTestsReportImportApiModel{*openapiclient.NewScheduleAutoTestsReportImportApiModel("AttachmentId_example")}) // ScheduleAutoTestsReportImportsApiModel | The request model with collection of reports to import. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AutoTestsProjectAPI.ApiV2ProjectsProjectIdAutoTestsImportPost(context.Background(), projectId).ScheduleAutoTestsReportImportsApiModel(scheduleAutoTestsReportImportsApiModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoTestsProjectAPI.ApiV2ProjectsProjectIdAutoTestsImportPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ProjectsProjectIdAutoTestsImportPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **scheduleAutoTestsReportImportsApiModel** | [**ScheduleAutoTestsReportImportsApiModel**](ScheduleAutoTestsReportImportsApiModel.md) | The request model with collection of reports to import. | 

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

