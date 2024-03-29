# \SectionsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV2SectionsIdPatch**](SectionsApi.md#ApiV2SectionsIdPatch) | **Patch** /api/v2/sections/{id} | Patch section
[**CreateSection**](SectionsApi.md#CreateSection) | **Post** /api/v2/sections | Create section
[**DeleteSection**](SectionsApi.md#DeleteSection) | **Delete** /api/v2/sections/{id} | Delete section
[**GetSectionById**](SectionsApi.md#GetSectionById) | **Get** /api/v2/sections/{id} | Get section
[**GetWorkItemsBySectionId**](SectionsApi.md#GetWorkItemsBySectionId) | **Get** /api/v2/sections/{id}/workItems | Get section work items
[**Move**](SectionsApi.md#Move) | **Post** /api/v2/sections/move | Move section with all work items into another section
[**Rename**](SectionsApi.md#Rename) | **Post** /api/v2/sections/rename | Rename section
[**UpdateSection**](SectionsApi.md#UpdateSection) | **Put** /api/v2/sections | Update section



## ApiV2SectionsIdPatch

> ApiV2SectionsIdPatch(ctx, id).Operation(operation).Execute()

Patch section



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Section internal (UUID) identifier
    operation := []openapiclient.Operation{*openapiclient.NewOperation()} // []Operation |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SectionsApi.ApiV2SectionsIdPatch(context.Background(), id).Operation(operation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SectionsApi.ApiV2SectionsIdPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Section internal (UUID) identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2SectionsIdPatchRequest struct via the builder pattern


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


## CreateSection

> SectionWithStepsModel CreateSection(ctx).CreateSectionRequest(createSectionRequest).Execute()

Create section



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
    createSectionRequest := *openapiclient.NewCreateSectionRequest("b21976c2-c9f3-497b-bd88-9e068d01bd4a", "b21976c2-c9f3-497b-bd88-9e068d01bd4a", []openapiclient.AttachmentPutModel{*openapiclient.NewAttachmentPutModel("Id_example")}) // CreateSectionRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SectionsApi.CreateSection(context.Background()).CreateSectionRequest(createSectionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SectionsApi.CreateSection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSection`: SectionWithStepsModel
    fmt.Fprintf(os.Stdout, "Response from `SectionsApi.CreateSection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createSectionRequest** | [**CreateSectionRequest**](CreateSectionRequest.md) |  | 

### Return type

[**SectionWithStepsModel**](SectionWithStepsModel.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSection

> DeleteSection(ctx, id).Execute()

Delete section



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Section internal (UUID) identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SectionsApi.DeleteSection(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SectionsApi.DeleteSection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Section internal (UUID) identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSectionRequest struct via the builder pattern


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


## GetSectionById

> SectionWithStepsModel GetSectionById(ctx, id).IsDeleted(isDeleted).Execute()

Get section



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Section internal (UUID) identifier
    isDeleted := openapiclient.DeletionState("Any") // DeletionState |  (optional) (default to "NotDeleted")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SectionsApi.GetSectionById(context.Background(), id).IsDeleted(isDeleted).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SectionsApi.GetSectionById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSectionById`: SectionWithStepsModel
    fmt.Fprintf(os.Stdout, "Response from `SectionsApi.GetSectionById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Section internal (UUID) identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSectionByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **isDeleted** | [**DeletionState**](DeletionState.md) |  | [default to &quot;NotDeleted&quot;]

### Return type

[**SectionWithStepsModel**](SectionWithStepsModel.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWorkItemsBySectionId

> []WorkItemShortModel GetWorkItemsBySectionId(ctx, id).IsDeleted(isDeleted).TagNames(tagNames).IncludeIterations(includeIterations).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).Execute()

Get section work items



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Section internal (UUID) identifier
    isDeleted := true // bool | Requested section is deleted (optional) (default to false)
    tagNames := []string{"Inner_example"} // []string | List of work item tags (optional)
    includeIterations := true // bool |  (optional) (default to true)
    skip := int32(56) // int32 | Amount of items to be skipped (offset) (optional)
    take := int32(56) // int32 | Amount of items to be taken (limit) (optional)
    orderBy := "orderBy_example" // string | SQL-like  ORDER BY statement (column1 ASC|DESC , column2 ASC|DESC) (optional)
    searchField := "searchField_example" // string | Property name for searching (optional)
    searchValue := "searchValue_example" // string | Value for searching (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SectionsApi.GetWorkItemsBySectionId(context.Background(), id).IsDeleted(isDeleted).TagNames(tagNames).IncludeIterations(includeIterations).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SectionsApi.GetWorkItemsBySectionId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorkItemsBySectionId`: []WorkItemShortModel
    fmt.Fprintf(os.Stdout, "Response from `SectionsApi.GetWorkItemsBySectionId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Section internal (UUID) identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkItemsBySectionIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **isDeleted** | **bool** | Requested section is deleted | [default to false]
 **tagNames** | **[]string** | List of work item tags | 
 **includeIterations** | **bool** |  | [default to true]
 **skip** | **int32** | Amount of items to be skipped (offset) | 
 **take** | **int32** | Amount of items to be taken (limit) | 
 **orderBy** | **string** | SQL-like  ORDER BY statement (column1 ASC|DESC , column2 ASC|DESC) | 
 **searchField** | **string** | Property name for searching | 
 **searchValue** | **string** | Value for searching | 

### Return type

[**[]WorkItemShortModel**](WorkItemShortModel.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Move

> Move(ctx).MoveRequest(moveRequest).Execute()

Move section with all work items into another section

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
    moveRequest := *openapiclient.NewMoveRequest("Id_example", "OldParentId_example", "ParentId_example") // MoveRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SectionsApi.Move(context.Background()).MoveRequest(moveRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SectionsApi.Move``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMoveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **moveRequest** | [**MoveRequest**](MoveRequest.md) |  | 

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


## Rename

> Rename(ctx).RenameRequest(renameRequest).Execute()

Rename section



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
    renameRequest := *openapiclient.NewRenameRequest("b21976c2-c9f3-497b-bd88-9e068d01bd4a", "New root section") // RenameRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SectionsApi.Rename(context.Background()).RenameRequest(renameRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SectionsApi.Rename``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRenameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **renameRequest** | [**RenameRequest**](RenameRequest.md) |  | 

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


## UpdateSection

> UpdateSection(ctx).UpdateSectionRequest(updateSectionRequest).Execute()

Update section



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
    updateSectionRequest := *openapiclient.NewUpdateSectionRequest("Id_example", "b21976c2-c9f3-497b-bd88-9e068d01bd4a", "b21976c2-c9f3-497b-bd88-9e068d01bd4a", []openapiclient.AttachmentPutModel{*openapiclient.NewAttachmentPutModel("Id_example")}) // UpdateSectionRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SectionsApi.UpdateSection(context.Background()).UpdateSectionRequest(updateSectionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SectionsApi.UpdateSection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateSectionRequest** | [**UpdateSectionRequest**](UpdateSectionRequest.md) |  | 

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

