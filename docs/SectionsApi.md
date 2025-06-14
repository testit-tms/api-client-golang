# \SectionsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV2SectionsIdPatch**](SectionsAPI.md#ApiV2SectionsIdPatch) | **Patch** /api/v2/sections/{id} | Patch section
[**CreateSection**](SectionsAPI.md#CreateSection) | **Post** /api/v2/sections | Create section
[**DeleteSection**](SectionsAPI.md#DeleteSection) | **Delete** /api/v2/sections/{id} | Delete section
[**GetSectionById**](SectionsAPI.md#GetSectionById) | **Get** /api/v2/sections/{id} | Get section
[**GetWorkItemsBySectionId**](SectionsAPI.md#GetWorkItemsBySectionId) | **Get** /api/v2/sections/{id}/workItems | Get section work items
[**Move**](SectionsAPI.md#Move) | **Post** /api/v2/sections/move | Move section with all work items into another section
[**Rename**](SectionsAPI.md#Rename) | **Post** /api/v2/sections/rename | Rename section
[**UpdateSection**](SectionsAPI.md#UpdateSection) | **Put** /api/v2/sections | Update section



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
	r, err := apiClient.SectionsAPI.ApiV2SectionsIdPatch(context.Background(), id).Operation(operation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SectionsAPI.ApiV2SectionsIdPatch``: %v\n", err)
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

> SectionWithStepsModel CreateSection(ctx).SectionPostModel(sectionPostModel).Execute()

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
	sectionPostModel := *openapiclient.NewSectionPostModel("0503a000-865d-4e08-8873-370a94517983", "0503a000-865d-4e08-8873-370a94517983", []openapiclient.AttachmentPutModel{*openapiclient.NewAttachmentPutModel("Id_example")}) // SectionPostModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SectionsAPI.CreateSection(context.Background()).SectionPostModel(sectionPostModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SectionsAPI.CreateSection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSection`: SectionWithStepsModel
	fmt.Fprintf(os.Stdout, "Response from `SectionsAPI.CreateSection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sectionPostModel** | [**SectionPostModel**](SectionPostModel.md) |  | 

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
	r, err := apiClient.SectionsAPI.DeleteSection(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SectionsAPI.DeleteSection``: %v\n", err)
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
	resp, r, err := apiClient.SectionsAPI.GetSectionById(context.Background(), id).IsDeleted(isDeleted).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SectionsAPI.GetSectionById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSectionById`: SectionWithStepsModel
	fmt.Fprintf(os.Stdout, "Response from `SectionsAPI.GetSectionById`: %v\n", resp)
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
	resp, r, err := apiClient.SectionsAPI.GetWorkItemsBySectionId(context.Background(), id).IsDeleted(isDeleted).TagNames(tagNames).IncludeIterations(includeIterations).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SectionsAPI.GetWorkItemsBySectionId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWorkItemsBySectionId`: []WorkItemShortModel
	fmt.Fprintf(os.Stdout, "Response from `SectionsAPI.GetWorkItemsBySectionId`: %v\n", resp)
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

> Move(ctx).SectionMoveModel(sectionMoveModel).Execute()

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
	sectionMoveModel := *openapiclient.NewSectionMoveModel("Id_example", "OldParentId_example", "ParentId_example") // SectionMoveModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SectionsAPI.Move(context.Background()).SectionMoveModel(sectionMoveModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SectionsAPI.Move``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMoveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sectionMoveModel** | [**SectionMoveModel**](SectionMoveModel.md) |  | 

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

> Rename(ctx).SectionRenameModel(sectionRenameModel).Execute()

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
	sectionRenameModel := *openapiclient.NewSectionRenameModel("0503a000-865d-4e08-8873-370a94517983", "New root section") // SectionRenameModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SectionsAPI.Rename(context.Background()).SectionRenameModel(sectionRenameModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SectionsAPI.Rename``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRenameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sectionRenameModel** | [**SectionRenameModel**](SectionRenameModel.md) |  | 

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

> UpdateSection(ctx).SectionPutModel(sectionPutModel).Execute()

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
	sectionPutModel := *openapiclient.NewSectionPutModel("Id_example", "0503a000-865d-4e08-8873-370a94517983", "0503a000-865d-4e08-8873-370a94517983", []openapiclient.AttachmentPutModel{*openapiclient.NewAttachmentPutModel("Id_example")}) // SectionPutModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SectionsAPI.UpdateSection(context.Background()).SectionPutModel(sectionPutModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SectionsAPI.UpdateSection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sectionPutModel** | [**SectionPutModel**](SectionPutModel.md) |  | 

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

