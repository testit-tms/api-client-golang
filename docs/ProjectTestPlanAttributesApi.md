# \ProjectTestPlanAttributesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCustomAttributeTestPlanProjectRelations**](ProjectTestPlanAttributesAPI.md#CreateCustomAttributeTestPlanProjectRelations) | **Post** /api/v2/projects/{projectId}/testPlans/attributes | Add attributes to project&#39;s test plans
[**DeleteCustomAttributeTestPlanProjectRelations**](ProjectTestPlanAttributesAPI.md#DeleteCustomAttributeTestPlanProjectRelations) | **Delete** /api/v2/projects/{projectId}/testPlans/attributes/{attributeId} | Delete attribute from project&#39;s test plans
[**GetCustomAttributeTestPlanProjectRelations**](ProjectTestPlanAttributesAPI.md#GetCustomAttributeTestPlanProjectRelations) | **Get** /api/v2/projects/{projectId}/testPlans/attributes | Get project&#39;s test plan attributes
[**SearchTestPlanAttributesInProject**](ProjectTestPlanAttributesAPI.md#SearchTestPlanAttributesInProject) | **Post** /api/v2/projects/{projectId}/testPlans/attributes/search | Search for attributes used in the project test plans
[**UpdateCustomAttributeTestPlanProjectRelations**](ProjectTestPlanAttributesAPI.md#UpdateCustomAttributeTestPlanProjectRelations) | **Put** /api/v2/projects/{projectId}/testPlans/attributes | Update attribute of project&#39;s test plans



## CreateCustomAttributeTestPlanProjectRelations

> CreateCustomAttributeTestPlanProjectRelations(ctx, projectId).RequestBody(requestBody).Execute()

Add attributes to project's test plans



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
	requestBody := []string{"Property_example"} // []string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectTestPlanAttributesAPI.CreateCustomAttributeTestPlanProjectRelations(context.Background(), projectId).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectTestPlanAttributesAPI.CreateCustomAttributeTestPlanProjectRelations``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCreateCustomAttributeTestPlanProjectRelationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestBody** | **[]string** |  | 

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


## DeleteCustomAttributeTestPlanProjectRelations

> DeleteCustomAttributeTestPlanProjectRelations(ctx, projectId, attributeId).Execute()

Delete attribute from project's test plans



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
	attributeId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectTestPlanAttributesAPI.DeleteCustomAttributeTestPlanProjectRelations(context.Background(), projectId, attributeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectTestPlanAttributesAPI.DeleteCustomAttributeTestPlanProjectRelations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project internal (UUID) or global (integer) identifier | 
**attributeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCustomAttributeTestPlanProjectRelationsRequest struct via the builder pattern


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


## GetCustomAttributeTestPlanProjectRelations

> []CustomAttributeModel GetCustomAttributeTestPlanProjectRelations(ctx, projectId).Execute()

Get project's test plan attributes



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectTestPlanAttributesAPI.GetCustomAttributeTestPlanProjectRelations(context.Background(), projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectTestPlanAttributesAPI.GetCustomAttributeTestPlanProjectRelations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCustomAttributeTestPlanProjectRelations`: []CustomAttributeModel
	fmt.Fprintf(os.Stdout, "Response from `ProjectTestPlanAttributesAPI.GetCustomAttributeTestPlanProjectRelations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project internal (UUID) or global (integer) identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomAttributeTestPlanProjectRelationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## SearchTestPlanAttributesInProject

> []CustomAttributeGetModel SearchTestPlanAttributesInProject(ctx, projectId).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).ProjectAttributesFilterModel(projectAttributesFilterModel).Execute()

Search for attributes used in the project test plans

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
	resp, r, err := apiClient.ProjectTestPlanAttributesAPI.SearchTestPlanAttributesInProject(context.Background(), projectId).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).ProjectAttributesFilterModel(projectAttributesFilterModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectTestPlanAttributesAPI.SearchTestPlanAttributesInProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchTestPlanAttributesInProject`: []CustomAttributeGetModel
	fmt.Fprintf(os.Stdout, "Response from `ProjectTestPlanAttributesAPI.SearchTestPlanAttributesInProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Unique or global project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchTestPlanAttributesInProjectRequest struct via the builder pattern


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


## UpdateCustomAttributeTestPlanProjectRelations

> UpdateCustomAttributeTestPlanProjectRelations(ctx, projectId).CustomAttributeTestPlanProjectRelationPutModel(customAttributeTestPlanProjectRelationPutModel).Execute()

Update attribute of project's test plans



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
	customAttributeTestPlanProjectRelationPutModel := *openapiclient.NewCustomAttributeTestPlanProjectRelationPutModel("Id_example", false, false) // CustomAttributeTestPlanProjectRelationPutModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectTestPlanAttributesAPI.UpdateCustomAttributeTestPlanProjectRelations(context.Background(), projectId).CustomAttributeTestPlanProjectRelationPutModel(customAttributeTestPlanProjectRelationPutModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectTestPlanAttributesAPI.UpdateCustomAttributeTestPlanProjectRelations``: %v\n", err)
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

Other parameters are passed through a pointer to a apiUpdateCustomAttributeTestPlanProjectRelationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **customAttributeTestPlanProjectRelationPutModel** | [**CustomAttributeTestPlanProjectRelationPutModel**](CustomAttributeTestPlanProjectRelationPutModel.md) |  | 

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

