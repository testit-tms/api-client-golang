# \ProjectImportApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BackgroundImportToExistingProject**](ProjectImportApi.md#BackgroundImportToExistingProject) | **Post** /api/v2/projects/{projectId}/import/json | Import project from JSON file into existing project in background job
[**BackgroundImportZipToExistingProject**](ProjectImportApi.md#BackgroundImportZipToExistingProject) | **Post** /api/v2/projects/{projectId}/import/zip | Import project from Zip file into existing project in background job
[**ImportToExistingProject**](ProjectImportApi.md#ImportToExistingProject) | **Post** /api/v2/projects/{projectId}/import | Import project from JSON file into existing project



## BackgroundImportToExistingProject

> string BackgroundImportToExistingProject(ctx, projectId).File(file).Execute()

Import project from JSON file into existing project in background job

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
    file := os.NewFile(1234, "some_file") // *os.File | Select file (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectImportApi.BackgroundImportToExistingProject(context.Background(), projectId).File(file).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectImportApi.BackgroundImportToExistingProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BackgroundImportToExistingProject`: string
    fmt.Fprintf(os.Stdout, "Response from `ProjectImportApi.BackgroundImportToExistingProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project internal (UUID) or global (integer) identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiBackgroundImportToExistingProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **file** | ***os.File** | Select file | 

### Return type

**string**

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BackgroundImportZipToExistingProject

> string BackgroundImportZipToExistingProject(ctx, projectId).File(file).Execute()

Import project from Zip file into existing project in background job

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
    file := os.NewFile(1234, "some_file") // *os.File | Select file (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectImportApi.BackgroundImportZipToExistingProject(context.Background(), projectId).File(file).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectImportApi.BackgroundImportZipToExistingProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BackgroundImportZipToExistingProject`: string
    fmt.Fprintf(os.Stdout, "Response from `ProjectImportApi.BackgroundImportZipToExistingProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project internal (UUID) or global (integer) identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiBackgroundImportZipToExistingProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **file** | ***os.File** | Select file | 

### Return type

**string**

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportToExistingProject

> ImportToExistingProject(ctx, projectId).IncludeAttachments(includeAttachments).File(file).Execute()

Import project from JSON file into existing project



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
    includeAttachments := true // bool |  (optional)
    file := os.NewFile(1234, "some_file") // *os.File | Select file (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ProjectImportApi.ImportToExistingProject(context.Background(), projectId).IncludeAttachments(includeAttachments).File(file).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectImportApi.ImportToExistingProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project internal (UUID) or global (integer) identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiImportToExistingProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeAttachments** | **bool** |  | 
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

