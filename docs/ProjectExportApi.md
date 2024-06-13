# \ProjectExportAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Export**](ProjectExportAPI.md#Export) | **Post** /api/v2/projects/{projectId}/export | Export project as JSON file
[**ExportProjectJson**](ProjectExportAPI.md#ExportProjectJson) | **Post** /api/v2/projects/{projectId}/export/json | Export project as JSON file in background job
[**ExportProjectWithTestPlansJson**](ProjectExportAPI.md#ExportProjectWithTestPlansJson) | **Post** /api/v2/projects/{projectId}/export/testPlans/json | Export project as JSON file with test plans in background job
[**ExportProjectWithTestPlansZip**](ProjectExportAPI.md#ExportProjectWithTestPlansZip) | **Post** /api/v2/projects/{projectId}/export/testPlans/zip | Export project as Zip file with test plans in background job
[**ExportProjectZip**](ProjectExportAPI.md#ExportProjectZip) | **Post** /api/v2/projects/{projectId}/export/zip | Export project as Zip file in background job



## Export

> *os.File Export(ctx, projectId).IncludeAttachments(includeAttachments).ProjectExportQueryModel(projectExportQueryModel).Execute()

Export project as JSON file



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
	projectId := "projectId_example" // string | Specifies the ID of the project you want to export.
	includeAttachments := true // bool | Enables attachment export. (optional) (default to false)
	projectExportQueryModel := *openapiclient.NewProjectExportQueryModel() // ProjectExportQueryModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectExportAPI.Export(context.Background(), projectId).IncludeAttachments(includeAttachments).ProjectExportQueryModel(projectExportQueryModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectExportAPI.Export``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Export`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ProjectExportAPI.Export`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Specifies the ID of the project you want to export. | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeAttachments** | **bool** | Enables attachment export. | [default to false]
 **projectExportQueryModel** | [**ProjectExportQueryModel**](ProjectExportQueryModel.md) |  | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportProjectJson

> string ExportProjectJson(ctx, projectId).TimeZoneOffsetInMinutes(timeZoneOffsetInMinutes).ProjectExportQueryModel(projectExportQueryModel).Execute()

Export project as JSON file in background job

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
	timeZoneOffsetInMinutes := int64(789) // int64 |  (optional)
	projectExportQueryModel := *openapiclient.NewProjectExportQueryModel() // ProjectExportQueryModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectExportAPI.ExportProjectJson(context.Background(), projectId).TimeZoneOffsetInMinutes(timeZoneOffsetInMinutes).ProjectExportQueryModel(projectExportQueryModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectExportAPI.ExportProjectJson``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportProjectJson`: string
	fmt.Fprintf(os.Stdout, "Response from `ProjectExportAPI.ExportProjectJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project internal (UUID) or global (integer) identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportProjectJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **timeZoneOffsetInMinutes** | **int64** |  | 
 **projectExportQueryModel** | [**ProjectExportQueryModel**](ProjectExportQueryModel.md) |  | 

### Return type

**string**

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportProjectWithTestPlansJson

> string ExportProjectWithTestPlansJson(ctx, projectId).TimeZoneOffsetInMinutes(timeZoneOffsetInMinutes).ProjectExportWithTestPlansPostModel(projectExportWithTestPlansPostModel).Execute()

Export project as JSON file with test plans in background job

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
	timeZoneOffsetInMinutes := int64(789) // int64 |  (optional)
	projectExportWithTestPlansPostModel := *openapiclient.NewProjectExportWithTestPlansPostModel() // ProjectExportWithTestPlansPostModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectExportAPI.ExportProjectWithTestPlansJson(context.Background(), projectId).TimeZoneOffsetInMinutes(timeZoneOffsetInMinutes).ProjectExportWithTestPlansPostModel(projectExportWithTestPlansPostModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectExportAPI.ExportProjectWithTestPlansJson``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportProjectWithTestPlansJson`: string
	fmt.Fprintf(os.Stdout, "Response from `ProjectExportAPI.ExportProjectWithTestPlansJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project internal (UUID) or global (integer) identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportProjectWithTestPlansJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **timeZoneOffsetInMinutes** | **int64** |  | 
 **projectExportWithTestPlansPostModel** | [**ProjectExportWithTestPlansPostModel**](ProjectExportWithTestPlansPostModel.md) |  | 

### Return type

**string**

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportProjectWithTestPlansZip

> string ExportProjectWithTestPlansZip(ctx, projectId).TimeZoneOffsetInMinutes(timeZoneOffsetInMinutes).ProjectExportWithTestPlansPostModel(projectExportWithTestPlansPostModel).Execute()

Export project as Zip file with test plans in background job

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
	timeZoneOffsetInMinutes := int64(789) // int64 |  (optional)
	projectExportWithTestPlansPostModel := *openapiclient.NewProjectExportWithTestPlansPostModel() // ProjectExportWithTestPlansPostModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectExportAPI.ExportProjectWithTestPlansZip(context.Background(), projectId).TimeZoneOffsetInMinutes(timeZoneOffsetInMinutes).ProjectExportWithTestPlansPostModel(projectExportWithTestPlansPostModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectExportAPI.ExportProjectWithTestPlansZip``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportProjectWithTestPlansZip`: string
	fmt.Fprintf(os.Stdout, "Response from `ProjectExportAPI.ExportProjectWithTestPlansZip`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project internal (UUID) or global (integer) identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportProjectWithTestPlansZipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **timeZoneOffsetInMinutes** | **int64** |  | 
 **projectExportWithTestPlansPostModel** | [**ProjectExportWithTestPlansPostModel**](ProjectExportWithTestPlansPostModel.md) |  | 

### Return type

**string**

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportProjectZip

> string ExportProjectZip(ctx, projectId).TimeZoneOffsetInMinutes(timeZoneOffsetInMinutes).ProjectExportQueryModel(projectExportQueryModel).Execute()

Export project as Zip file in background job

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
	timeZoneOffsetInMinutes := int64(789) // int64 |  (optional)
	projectExportQueryModel := *openapiclient.NewProjectExportQueryModel() // ProjectExportQueryModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectExportAPI.ExportProjectZip(context.Background(), projectId).TimeZoneOffsetInMinutes(timeZoneOffsetInMinutes).ProjectExportQueryModel(projectExportQueryModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectExportAPI.ExportProjectZip``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportProjectZip`: string
	fmt.Fprintf(os.Stdout, "Response from `ProjectExportAPI.ExportProjectZip`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project internal (UUID) or global (integer) identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportProjectZipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **timeZoneOffsetInMinutes** | **int64** |  | 
 **projectExportQueryModel** | [**ProjectExportQueryModel**](ProjectExportQueryModel.md) |  | 

### Return type

**string**

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

