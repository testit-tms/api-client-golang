# \ProjectExportApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Export**](ProjectExportApi.md#Export) | **Post** /api/v2/projects/{projectId}/export | Export project as JSON file
[**ExportProjectJson**](ProjectExportApi.md#ExportProjectJson) | **Post** /api/v2/projects/{projectId}/export/json | Export project as JSON file in background job
[**ExportProjectWithTestPlansJson**](ProjectExportApi.md#ExportProjectWithTestPlansJson) | **Post** /api/v2/projects/{projectId}/export/testPlans/json | Export project as JSON file with test plans in background job
[**ExportProjectWithTestPlansZip**](ProjectExportApi.md#ExportProjectWithTestPlansZip) | **Post** /api/v2/projects/{projectId}/export/testPlans/zip | Export project as Zip file with test plans in background job
[**ExportProjectZip**](ProjectExportApi.md#ExportProjectZip) | **Post** /api/v2/projects/{projectId}/export/zip | Export project as Zip file in background job



## Export

> *os.File Export(ctx, projectId).IncludeAttachments(includeAttachments).ExportProjectJsonRequest(exportProjectJsonRequest).Execute()

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
    exportProjectJsonRequest := *openapiclient.NewExportProjectJsonRequest() // ExportProjectJsonRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectExportApi.Export(context.Background(), projectId).IncludeAttachments(includeAttachments).ExportProjectJsonRequest(exportProjectJsonRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectExportApi.Export``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Export`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `ProjectExportApi.Export`: %v\n", resp)
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
 **exportProjectJsonRequest** | [**ExportProjectJsonRequest**](ExportProjectJsonRequest.md) |  | 

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

> string ExportProjectJson(ctx, projectId).TimeZoneOffsetInMinutes(timeZoneOffsetInMinutes).ExportProjectJsonRequest(exportProjectJsonRequest).Execute()

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
    exportProjectJsonRequest := *openapiclient.NewExportProjectJsonRequest() // ExportProjectJsonRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectExportApi.ExportProjectJson(context.Background(), projectId).TimeZoneOffsetInMinutes(timeZoneOffsetInMinutes).ExportProjectJsonRequest(exportProjectJsonRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectExportApi.ExportProjectJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportProjectJson`: string
    fmt.Fprintf(os.Stdout, "Response from `ProjectExportApi.ExportProjectJson`: %v\n", resp)
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
 **exportProjectJsonRequest** | [**ExportProjectJsonRequest**](ExportProjectJsonRequest.md) |  | 

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

> string ExportProjectWithTestPlansJson(ctx, projectId).TimeZoneOffsetInMinutes(timeZoneOffsetInMinutes).ExportProjectWithTestPlansJsonRequest(exportProjectWithTestPlansJsonRequest).Execute()

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
    exportProjectWithTestPlansJsonRequest := *openapiclient.NewExportProjectWithTestPlansJsonRequest() // ExportProjectWithTestPlansJsonRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectExportApi.ExportProjectWithTestPlansJson(context.Background(), projectId).TimeZoneOffsetInMinutes(timeZoneOffsetInMinutes).ExportProjectWithTestPlansJsonRequest(exportProjectWithTestPlansJsonRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectExportApi.ExportProjectWithTestPlansJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportProjectWithTestPlansJson`: string
    fmt.Fprintf(os.Stdout, "Response from `ProjectExportApi.ExportProjectWithTestPlansJson`: %v\n", resp)
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
 **exportProjectWithTestPlansJsonRequest** | [**ExportProjectWithTestPlansJsonRequest**](ExportProjectWithTestPlansJsonRequest.md) |  | 

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

> string ExportProjectWithTestPlansZip(ctx, projectId).TimeZoneOffsetInMinutes(timeZoneOffsetInMinutes).ExportProjectWithTestPlansJsonRequest(exportProjectWithTestPlansJsonRequest).Execute()

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
    exportProjectWithTestPlansJsonRequest := *openapiclient.NewExportProjectWithTestPlansJsonRequest() // ExportProjectWithTestPlansJsonRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectExportApi.ExportProjectWithTestPlansZip(context.Background(), projectId).TimeZoneOffsetInMinutes(timeZoneOffsetInMinutes).ExportProjectWithTestPlansJsonRequest(exportProjectWithTestPlansJsonRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectExportApi.ExportProjectWithTestPlansZip``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportProjectWithTestPlansZip`: string
    fmt.Fprintf(os.Stdout, "Response from `ProjectExportApi.ExportProjectWithTestPlansZip`: %v\n", resp)
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
 **exportProjectWithTestPlansJsonRequest** | [**ExportProjectWithTestPlansJsonRequest**](ExportProjectWithTestPlansJsonRequest.md) |  | 

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

> string ExportProjectZip(ctx, projectId).TimeZoneOffsetInMinutes(timeZoneOffsetInMinutes).ExportProjectJsonRequest(exportProjectJsonRequest).Execute()

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
    exportProjectJsonRequest := *openapiclient.NewExportProjectJsonRequest() // ExportProjectJsonRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectExportApi.ExportProjectZip(context.Background(), projectId).TimeZoneOffsetInMinutes(timeZoneOffsetInMinutes).ExportProjectJsonRequest(exportProjectJsonRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectExportApi.ExportProjectZip``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportProjectZip`: string
    fmt.Fprintf(os.Stdout, "Response from `ProjectExportApi.ExportProjectZip`: %v\n", resp)
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
 **exportProjectJsonRequest** | [**ExportProjectJsonRequest**](ExportProjectJsonRequest.md) |  | 

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

