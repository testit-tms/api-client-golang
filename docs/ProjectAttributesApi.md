# \ProjectAttributesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateProjectsAttribute**](ProjectAttributesAPI.md#CreateProjectsAttribute) | **Post** /api/v2/projects/{projectId}/attributes | Create project attribute
[**DeleteProjectsAttribute**](ProjectAttributesAPI.md#DeleteProjectsAttribute) | **Delete** /api/v2/projects/{projectId}/attributes/{attributeId} | Delete project attribute
[**GetAttributeByProjectId**](ProjectAttributesAPI.md#GetAttributeByProjectId) | **Get** /api/v2/projects/{projectId}/attributes/{attributeId} | Get project attribute
[**GetAttributesByProjectId**](ProjectAttributesAPI.md#GetAttributesByProjectId) | **Get** /api/v2/projects/{projectId}/attributes | Get project attributes
[**SearchAttributesInProject**](ProjectAttributesAPI.md#SearchAttributesInProject) | **Post** /api/v2/projects/{projectId}/attributes/search | Search for attributes used in the project
[**UpdateProjectsAttribute**](ProjectAttributesAPI.md#UpdateProjectsAttribute) | **Put** /api/v2/projects/{projectId}/attributes | Edit attribute of the project



## CreateProjectsAttribute

> CustomAttributeModel CreateProjectsAttribute(ctx, projectId).CustomAttributePostModel(customAttributePostModel).Execute()

Create project attribute



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
	customAttributePostModel := *openapiclient.NewCustomAttributePostModel(openapiclient.CustomAttributeTypesEnum("string"), "Name_example", false, false, false) // CustomAttributePostModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAttributesAPI.CreateProjectsAttribute(context.Background(), projectId).CustomAttributePostModel(customAttributePostModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAttributesAPI.CreateProjectsAttribute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateProjectsAttribute`: CustomAttributeModel
	fmt.Fprintf(os.Stdout, "Response from `ProjectAttributesAPI.CreateProjectsAttribute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project internal (UUID) or global (integer) identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateProjectsAttributeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **customAttributePostModel** | [**CustomAttributePostModel**](CustomAttributePostModel.md) |  | 

### Return type

[**CustomAttributeModel**](CustomAttributeModel.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteProjectsAttribute

> DeleteProjectsAttribute(ctx, projectId, attributeId).Execute()

Delete project attribute



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
	attributeId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Project attribute internal (UUID)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectAttributesAPI.DeleteProjectsAttribute(context.Background(), projectId, attributeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAttributesAPI.DeleteProjectsAttribute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project internal (UUID) or global (integer) identifier | 
**attributeId** | **string** | Project attribute internal (UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProjectsAttributeRequest struct via the builder pattern


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


## GetAttributeByProjectId

> CustomAttributeModel GetAttributeByProjectId(ctx, projectId, attributeId).Execute()

Get project attribute



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
	attributeId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Project attribute internal (UUID) or global (integer) identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAttributesAPI.GetAttributeByProjectId(context.Background(), projectId, attributeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAttributesAPI.GetAttributeByProjectId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAttributeByProjectId`: CustomAttributeModel
	fmt.Fprintf(os.Stdout, "Response from `ProjectAttributesAPI.GetAttributeByProjectId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project internal (UUID) or global (integer) identifier | 
**attributeId** | **string** | Project attribute internal (UUID) or global (integer) identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAttributeByProjectIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CustomAttributeModel**](CustomAttributeModel.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAttributesByProjectId

> []CustomAttributeModel GetAttributesByProjectId(ctx, projectId).IsDeleted(isDeleted).Execute()

Get project attributes



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
	isDeleted := openapiclient.DeletionState("Any") // DeletionState |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAttributesAPI.GetAttributesByProjectId(context.Background(), projectId).IsDeleted(isDeleted).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAttributesAPI.GetAttributesByProjectId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAttributesByProjectId`: []CustomAttributeModel
	fmt.Fprintf(os.Stdout, "Response from `ProjectAttributesAPI.GetAttributesByProjectId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project internal (UUID) or global (integer) identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAttributesByProjectIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **isDeleted** | [**DeletionState**](DeletionState.md) |  | 

### Return type

[**[]CustomAttributeModel**](CustomAttributeModel.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchAttributesInProject

> []CustomAttributeGetModel SearchAttributesInProject(ctx, projectId).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).ProjectAttributesFilterModel(projectAttributesFilterModel).Execute()

Search for attributes used in the project

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
	projectId := "projectId_example" // string | Unique or global project ID
	skip := int32(56) // int32 | Amount of items to be skipped (offset) (optional)
	take := int32(56) // int32 | Amount of items to be taken (limit) (optional)
	orderBy := "orderBy_example" // string | SQL-like  ORDER BY statement (column1 ASC|DESC , column2 ASC|DESC) (optional)
	searchField := "searchField_example" // string | Property name for searching (optional)
	searchValue := "searchValue_example" // string | Value for searching (optional)
	projectAttributesFilterModel := *openapiclient.NewProjectAttributesFilterModel("Name_example", []openapiclient.CustomAttributeTypesEnum{openapiclient.CustomAttributeTypesEnum("string")}) // ProjectAttributesFilterModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAttributesAPI.SearchAttributesInProject(context.Background(), projectId).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).ProjectAttributesFilterModel(projectAttributesFilterModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAttributesAPI.SearchAttributesInProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchAttributesInProject`: []CustomAttributeGetModel
	fmt.Fprintf(os.Stdout, "Response from `ProjectAttributesAPI.SearchAttributesInProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Unique or global project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchAttributesInProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **skip** | **int32** | Amount of items to be skipped (offset) | 
 **take** | **int32** | Amount of items to be taken (limit) | 
 **orderBy** | **string** | SQL-like  ORDER BY statement (column1 ASC|DESC , column2 ASC|DESC) | 
 **searchField** | **string** | Property name for searching | 
 **searchValue** | **string** | Value for searching | 
 **projectAttributesFilterModel** | [**ProjectAttributesFilterModel**](ProjectAttributesFilterModel.md) |  | 

### Return type

[**[]CustomAttributeGetModel**](CustomAttributeGetModel.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateProjectsAttribute

> UpdateProjectsAttribute(ctx, projectId).CustomAttributePutModel(customAttributePutModel).Execute()

Edit attribute of the project

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
	projectId := "projectId_example" // string | Unique or global project ID
	customAttributePutModel := *openapiclient.NewCustomAttributePutModel("Id_example", openapiclient.CustomAttributeTypesEnum("string"), false, "Name_example", false, false, false) // CustomAttributePutModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectAttributesAPI.UpdateProjectsAttribute(context.Background(), projectId).CustomAttributePutModel(customAttributePutModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAttributesAPI.UpdateProjectsAttribute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Unique or global project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateProjectsAttributeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **customAttributePutModel** | [**CustomAttributePutModel**](CustomAttributePutModel.md) |  | 

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

