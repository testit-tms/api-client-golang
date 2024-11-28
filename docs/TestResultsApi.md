# \TestResultsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV2TestResultsExternalProjectsExternalProjectIdDefectsExternalFormsPost**](TestResultsAPI.md#ApiV2TestResultsExternalProjectsExternalProjectIdDefectsExternalFormsPost) | **Post** /api/v2/testResults/external-projects/{externalProjectId}/defects/external-forms | 
[**ApiV2TestResultsExternalProjectsExternalProjectIdDefectsPost**](TestResultsAPI.md#ApiV2TestResultsExternalProjectsExternalProjectIdDefectsPost) | **Post** /api/v2/testResults/external-projects/{externalProjectId}/defects | 
[**ApiV2TestResultsIdAggregatedGet**](TestResultsAPI.md#ApiV2TestResultsIdAggregatedGet) | **Get** /api/v2/testResults/{id}/aggregated | Get test result by ID aggregated with previous results
[**ApiV2TestResultsIdAttachmentsAttachmentIdPut**](TestResultsAPI.md#ApiV2TestResultsIdAttachmentsAttachmentIdPut) | **Put** /api/v2/testResults/{id}/attachments/{attachmentId} | Attach file to the test result
[**ApiV2TestResultsIdAttachmentsInfoGet**](TestResultsAPI.md#ApiV2TestResultsIdAttachmentsInfoGet) | **Get** /api/v2/testResults/{id}/attachments/info | Get test result attachments meta-information
[**ApiV2TestResultsIdGet**](TestResultsAPI.md#ApiV2TestResultsIdGet) | **Get** /api/v2/testResults/{id} | Get test result by ID
[**ApiV2TestResultsIdPut**](TestResultsAPI.md#ApiV2TestResultsIdPut) | **Put** /api/v2/testResults/{id} | Edit test result by ID
[**ApiV2TestResultsIdRerunsGet**](TestResultsAPI.md#ApiV2TestResultsIdRerunsGet) | **Get** /api/v2/testResults/{id}/reruns | Get reruns
[**ApiV2TestResultsSearchPost**](TestResultsAPI.md#ApiV2TestResultsSearchPost) | **Post** /api/v2/testResults/search | Search for test results
[**ApiV2TestResultsStatisticsFilterPost**](TestResultsAPI.md#ApiV2TestResultsStatisticsFilterPost) | **Post** /api/v2/testResults/statistics/filter | Search for test results and extract statistics
[**CreateAttachment**](TestResultsAPI.md#CreateAttachment) | **Post** /api/v2/testResults/{id}/attachments | Upload and link attachment to TestResult
[**DeleteAttachment**](TestResultsAPI.md#DeleteAttachment) | **Delete** /api/v2/testResults/{id}/attachments/{attachmentId} | Remove attachment and unlink from TestResult
[**DownloadAttachment**](TestResultsAPI.md#DownloadAttachment) | **Get** /api/v2/testResults/{id}/attachments/{attachmentId} | Get attachment of TestResult
[**GetAttachment**](TestResultsAPI.md#GetAttachment) | **Get** /api/v2/testResults/{id}/attachments/{attachmentId}/info | Get Metadata of TestResult&#39;s attachment
[**GetAttachments**](TestResultsAPI.md#GetAttachments) | **Get** /api/v2/testResults/{id}/attachments | Get all attachments of TestResult



## ApiV2TestResultsExternalProjectsExternalProjectIdDefectsExternalFormsPost

> GetExternalFormApiResult ApiV2TestResultsExternalProjectsExternalProjectIdDefectsExternalFormsPost(ctx, externalProjectId).TestResultsSelectApiModel(testResultsSelectApiModel).Execute()



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
	externalProjectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	testResultsSelectApiModel := *openapiclient.NewTestResultsSelectApiModel(*openapiclient.NewTestResultsFilterRequest(), *openapiclient.NewTestResultsExtractionApiModel()) // TestResultsSelectApiModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TestResultsAPI.ApiV2TestResultsExternalProjectsExternalProjectIdDefectsExternalFormsPost(context.Background(), externalProjectId).TestResultsSelectApiModel(testResultsSelectApiModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestResultsAPI.ApiV2TestResultsExternalProjectsExternalProjectIdDefectsExternalFormsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2TestResultsExternalProjectsExternalProjectIdDefectsExternalFormsPost`: GetExternalFormApiResult
	fmt.Fprintf(os.Stdout, "Response from `TestResultsAPI.ApiV2TestResultsExternalProjectsExternalProjectIdDefectsExternalFormsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalProjectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2TestResultsExternalProjectsExternalProjectIdDefectsExternalFormsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **testResultsSelectApiModel** | [**TestResultsSelectApiModel**](TestResultsSelectApiModel.md) |  | 

### Return type

[**GetExternalFormApiResult**](GetExternalFormApiResult.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2TestResultsExternalProjectsExternalProjectIdDefectsPost

> DefectApiModel ApiV2TestResultsExternalProjectsExternalProjectIdDefectsPost(ctx, externalProjectId).CreateDefectApiModel(createDefectApiModel).Execute()



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
	externalProjectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	createDefectApiModel := *openapiclient.NewCreateDefectApiModel([]string{"TestResultIds_example"}, *openapiclient.NewExternalFormCreateModel(map[string][]openapiclient.ExternalFormAllowedValueModel{"key": []openapiclient.ExternalFormAllowedValueModel{*openapiclient.NewExternalFormAllowedValueModel(false)}}, []openapiclient.ExternalFormFieldModel{*openapiclient.NewExternalFormFieldModel(false)}, []openapiclient.ExternalFormLinkModel{*openapiclient.NewExternalFormLinkModel("Name_example", "Url_example")}, map[string]interface{}{"key": interface{}(123)})) // CreateDefectApiModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TestResultsAPI.ApiV2TestResultsExternalProjectsExternalProjectIdDefectsPost(context.Background(), externalProjectId).CreateDefectApiModel(createDefectApiModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestResultsAPI.ApiV2TestResultsExternalProjectsExternalProjectIdDefectsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2TestResultsExternalProjectsExternalProjectIdDefectsPost`: DefectApiModel
	fmt.Fprintf(os.Stdout, "Response from `TestResultsAPI.ApiV2TestResultsExternalProjectsExternalProjectIdDefectsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalProjectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2TestResultsExternalProjectsExternalProjectIdDefectsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createDefectApiModel** | [**CreateDefectApiModel**](CreateDefectApiModel.md) |  | 

### Return type

[**DefectApiModel**](DefectApiModel.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2TestResultsIdAggregatedGet

> TestResultResponse ApiV2TestResultsIdAggregatedGet(ctx, id).Execute()

Get test result by ID aggregated with previous results

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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Test result unique ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TestResultsAPI.ApiV2TestResultsIdAggregatedGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestResultsAPI.ApiV2TestResultsIdAggregatedGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2TestResultsIdAggregatedGet`: TestResultResponse
	fmt.Fprintf(os.Stdout, "Response from `TestResultsAPI.ApiV2TestResultsIdAggregatedGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Test result unique ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2TestResultsIdAggregatedGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TestResultResponse**](TestResultResponse.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2TestResultsIdAttachmentsAttachmentIdPut

> ApiV2TestResultsIdAttachmentsAttachmentIdPut(ctx, id, attachmentId).Execute()

Attach file to the test result

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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Test result unique ID
	attachmentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Attachment unique ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TestResultsAPI.ApiV2TestResultsIdAttachmentsAttachmentIdPut(context.Background(), id, attachmentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestResultsAPI.ApiV2TestResultsIdAttachmentsAttachmentIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Test result unique ID | 
**attachmentId** | **string** | Attachment unique ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2TestResultsIdAttachmentsAttachmentIdPutRequest struct via the builder pattern


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


## ApiV2TestResultsIdAttachmentsInfoGet

> []AttachmentModel ApiV2TestResultsIdAttachmentsInfoGet(ctx, id).Execute()

Get test result attachments meta-information

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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Test result unique ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TestResultsAPI.ApiV2TestResultsIdAttachmentsInfoGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestResultsAPI.ApiV2TestResultsIdAttachmentsInfoGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2TestResultsIdAttachmentsInfoGet`: []AttachmentModel
	fmt.Fprintf(os.Stdout, "Response from `TestResultsAPI.ApiV2TestResultsIdAttachmentsInfoGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Test result unique ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2TestResultsIdAttachmentsInfoGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]AttachmentModel**](AttachmentModel.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2TestResultsIdGet

> TestResultResponse ApiV2TestResultsIdGet(ctx, id).Execute()

Get test result by ID

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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Test result unique ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TestResultsAPI.ApiV2TestResultsIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestResultsAPI.ApiV2TestResultsIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2TestResultsIdGet`: TestResultResponse
	fmt.Fprintf(os.Stdout, "Response from `TestResultsAPI.ApiV2TestResultsIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Test result unique ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2TestResultsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TestResultResponse**](TestResultResponse.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2TestResultsIdPut

> ApiV2TestResultsIdPut(ctx, id).TestResultUpdateV2Request(testResultUpdateV2Request).Execute()

Edit test result by ID

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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Test result unique ID
	testResultUpdateV2Request := *openapiclient.NewTestResultUpdateV2Request() // TestResultUpdateV2Request |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TestResultsAPI.ApiV2TestResultsIdPut(context.Background(), id).TestResultUpdateV2Request(testResultUpdateV2Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestResultsAPI.ApiV2TestResultsIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Test result unique ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2TestResultsIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **testResultUpdateV2Request** | [**TestResultUpdateV2Request**](TestResultUpdateV2Request.md) |  | 

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


## ApiV2TestResultsIdRerunsGet

> RerunsModel ApiV2TestResultsIdRerunsGet(ctx, id).Execute()

Get reruns

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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Test result unique ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TestResultsAPI.ApiV2TestResultsIdRerunsGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestResultsAPI.ApiV2TestResultsIdRerunsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2TestResultsIdRerunsGet`: RerunsModel
	fmt.Fprintf(os.Stdout, "Response from `TestResultsAPI.ApiV2TestResultsIdRerunsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Test result unique ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2TestResultsIdRerunsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RerunsModel**](RerunsModel.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2TestResultsSearchPost

> []TestResultShortResponse ApiV2TestResultsSearchPost(ctx).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).TestResultsFilterRequest(testResultsFilterRequest).Execute()

Search for test results

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
	testResultsFilterRequest := *openapiclient.NewTestResultsFilterRequest() // TestResultsFilterRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TestResultsAPI.ApiV2TestResultsSearchPost(context.Background()).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).TestResultsFilterRequest(testResultsFilterRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestResultsAPI.ApiV2TestResultsSearchPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2TestResultsSearchPost`: []TestResultShortResponse
	fmt.Fprintf(os.Stdout, "Response from `TestResultsAPI.ApiV2TestResultsSearchPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2TestResultsSearchPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skip** | **int32** | Amount of items to be skipped (offset) | 
 **take** | **int32** | Amount of items to be taken (limit) | 
 **orderBy** | **string** | SQL-like  ORDER BY statement (column1 ASC|DESC , column2 ASC|DESC) | 
 **searchField** | **string** | Property name for searching | 
 **searchValue** | **string** | Value for searching | 
 **testResultsFilterRequest** | [**TestResultsFilterRequest**](TestResultsFilterRequest.md) |  | 

### Return type

[**[]TestResultShortResponse**](TestResultShortResponse.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2TestResultsStatisticsFilterPost

> TestResultsStatisticsResponse ApiV2TestResultsStatisticsFilterPost(ctx).TestResultsFilterRequest(testResultsFilterRequest).Execute()

Search for test results and extract statistics

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
	testResultsFilterRequest := *openapiclient.NewTestResultsFilterRequest() // TestResultsFilterRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TestResultsAPI.ApiV2TestResultsStatisticsFilterPost(context.Background()).TestResultsFilterRequest(testResultsFilterRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestResultsAPI.ApiV2TestResultsStatisticsFilterPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2TestResultsStatisticsFilterPost`: TestResultsStatisticsResponse
	fmt.Fprintf(os.Stdout, "Response from `TestResultsAPI.ApiV2TestResultsStatisticsFilterPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2TestResultsStatisticsFilterPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **testResultsFilterRequest** | [**TestResultsFilterRequest**](TestResultsFilterRequest.md) |  | 

### Return type

[**TestResultsStatisticsResponse**](TestResultsStatisticsResponse.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAttachment

> CreateAttachment(ctx, id).File(file).Execute()

Upload and link attachment to TestResult



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
	id := "3fa85f64-5717-4562-b3fc-2c963f66afa6" // string | Test result internal identifier (guid format)
	file := os.NewFile(1234, "some_file") // *os.File | Select file (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TestResultsAPI.CreateAttachment(context.Background(), id).File(file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestResultsAPI.CreateAttachment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Test result internal identifier (guid format) | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAttachmentRequest struct via the builder pattern


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


## DeleteAttachment

> DeleteAttachment(ctx, id, attachmentId).Execute()

Remove attachment and unlink from TestResult



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
	id := "3fa85f64-5717-4562-b3fc-2c963f66afa6" // string | Test result internal identifier (guid format)
	attachmentId := "3fa85f64-5717-4562-b3fc-2c963f66afa6" // string | Attachment internal identifier (guid format)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TestResultsAPI.DeleteAttachment(context.Background(), id, attachmentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestResultsAPI.DeleteAttachment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Test result internal identifier (guid format) | 
**attachmentId** | **string** | Attachment internal identifier (guid format) | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAttachmentRequest struct via the builder pattern


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


## DownloadAttachment

> DownloadAttachment(ctx, attachmentId, id).Width(width).Height(height).ResizeType(resizeType).BackgroundColor(backgroundColor).Preview(preview).Execute()

Get attachment of TestResult



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
	attachmentId := "3fa85f64-5717-4562-b3fc-2c963f66afa6" // string | Attachment internal identifier (guid format)
	id := "3fa85f64-5717-4562-b3fc-2c963f66afa6" // string | Test result internal identifier (guid format)
	width := int32(56) // int32 | Width of the result image (optional)
	height := int32(56) // int32 | Height of the result image (optional)
	resizeType := openapiclient.ImageResizeType("Crop") // ImageResizeType | Type of resizing to apply to the result image (optional)
	backgroundColor := "backgroundColor_example" // string | Color of the background if the `resizeType` is `AddBackgroundStripes` (optional)
	preview := true // bool | If image must be converted to a preview (lower quality, no animation) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TestResultsAPI.DownloadAttachment(context.Background(), attachmentId, id).Width(width).Height(height).ResizeType(resizeType).BackgroundColor(backgroundColor).Preview(preview).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestResultsAPI.DownloadAttachment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**attachmentId** | **string** | Attachment internal identifier (guid format) | 
**id** | **string** | Test result internal identifier (guid format) | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadAttachmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **width** | **int32** | Width of the result image | 
 **height** | **int32** | Height of the result image | 
 **resizeType** | [**ImageResizeType**](ImageResizeType.md) | Type of resizing to apply to the result image | 
 **backgroundColor** | **string** | Color of the background if the &#x60;resizeType&#x60; is &#x60;AddBackgroundStripes&#x60; | 
 **preview** | **bool** | If image must be converted to a preview (lower quality, no animation) | 

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


## GetAttachment

> AttachmentModel GetAttachment(ctx, id, attachmentId).Execute()

Get Metadata of TestResult's attachment



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
	id := "3fa85f64-5717-4562-b3fc-2c963f66afa6" // string | Test result internal identifier (guid format)
	attachmentId := "3fa85f64-5717-4562-b3fc-2c963f66afa6" // string | Attachment internal identifier (guid format)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TestResultsAPI.GetAttachment(context.Background(), id, attachmentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestResultsAPI.GetAttachment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAttachment`: AttachmentModel
	fmt.Fprintf(os.Stdout, "Response from `TestResultsAPI.GetAttachment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Test result internal identifier (guid format) | 
**attachmentId** | **string** | Attachment internal identifier (guid format) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAttachmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AttachmentModel**](AttachmentModel.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAttachments

> []AttachmentModel GetAttachments(ctx, id).Execute()

Get all attachments of TestResult



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
	id := "3fa85f64-5717-4562-b3fc-2c963f66afa6" // string | Test result internal identifier (guid format)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TestResultsAPI.GetAttachments(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestResultsAPI.GetAttachments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAttachments`: []AttachmentModel
	fmt.Fprintf(os.Stdout, "Response from `TestResultsAPI.GetAttachments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Test result internal identifier (guid format) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAttachmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]AttachmentModel**](AttachmentModel.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

