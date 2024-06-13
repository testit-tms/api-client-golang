# \ProjectAttributeTemplatesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV2ProjectsProjectIdAttributesTemplatesSearchPost**](ProjectAttributeTemplatesAPI.md#ApiV2ProjectsProjectIdAttributesTemplatesSearchPost) | **Post** /api/v2/projects/{projectId}/attributes/templates/search | Search for custom attributes templates
[**ApiV2ProjectsProjectIdAttributesTemplatesTemplateIdDelete**](ProjectAttributeTemplatesAPI.md#ApiV2ProjectsProjectIdAttributesTemplatesTemplateIdDelete) | **Delete** /api/v2/projects/{projectId}/attributes/templates/{templateId} | Delete CustomAttributeTemplate from Project
[**ApiV2ProjectsProjectIdAttributesTemplatesTemplateIdPost**](ProjectAttributeTemplatesAPI.md#ApiV2ProjectsProjectIdAttributesTemplatesTemplateIdPost) | **Post** /api/v2/projects/{projectId}/attributes/templates/{templateId} | Add CustomAttributeTemplate to Project



## ApiV2ProjectsProjectIdAttributesTemplatesSearchPost

> []ProjectCustomAttributeTemplateGetModel ApiV2ProjectsProjectIdAttributesTemplatesSearchPost(ctx, projectId).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).ProjectCustomAttributesTemplatesFilterModel(projectCustomAttributesTemplatesFilterModel).Execute()

Search for custom attributes templates

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
	projectId := "projectId_example" // string | 
	skip := int32(56) // int32 | Amount of items to be skipped (offset) (optional)
	take := int32(56) // int32 | Amount of items to be taken (limit) (optional)
	orderBy := "orderBy_example" // string | SQL-like  ORDER BY statement (column1 ASC|DESC , column2 ASC|DESC) (optional)
	searchField := "searchField_example" // string | Property name for searching (optional)
	searchValue := "searchValue_example" // string | Value for searching (optional)
	projectCustomAttributesTemplatesFilterModel := *openapiclient.NewProjectCustomAttributesTemplatesFilterModel() // ProjectCustomAttributesTemplatesFilterModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAttributeTemplatesAPI.ApiV2ProjectsProjectIdAttributesTemplatesSearchPost(context.Background(), projectId).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).ProjectCustomAttributesTemplatesFilterModel(projectCustomAttributesTemplatesFilterModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAttributeTemplatesAPI.ApiV2ProjectsProjectIdAttributesTemplatesSearchPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2ProjectsProjectIdAttributesTemplatesSearchPost`: []ProjectCustomAttributeTemplateGetModel
	fmt.Fprintf(os.Stdout, "Response from `ProjectAttributeTemplatesAPI.ApiV2ProjectsProjectIdAttributesTemplatesSearchPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ProjectsProjectIdAttributesTemplatesSearchPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **skip** | **int32** | Amount of items to be skipped (offset) | 
 **take** | **int32** | Amount of items to be taken (limit) | 
 **orderBy** | **string** | SQL-like  ORDER BY statement (column1 ASC|DESC , column2 ASC|DESC) | 
 **searchField** | **string** | Property name for searching | 
 **searchValue** | **string** | Value for searching | 
 **projectCustomAttributesTemplatesFilterModel** | [**ProjectCustomAttributesTemplatesFilterModel**](ProjectCustomAttributesTemplatesFilterModel.md) |  | 

### Return type

[**[]ProjectCustomAttributeTemplateGetModel**](ProjectCustomAttributeTemplateGetModel.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2ProjectsProjectIdAttributesTemplatesTemplateIdDelete

> ApiV2ProjectsProjectIdAttributesTemplatesTemplateIdDelete(ctx, projectId, templateId).Execute()

Delete CustomAttributeTemplate from Project



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
	templateId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | CustomAttributeTemplate internal (UUID) identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectAttributeTemplatesAPI.ApiV2ProjectsProjectIdAttributesTemplatesTemplateIdDelete(context.Background(), projectId, templateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAttributeTemplatesAPI.ApiV2ProjectsProjectIdAttributesTemplatesTemplateIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project internal (UUID) or global (integer) identifier | 
**templateId** | **string** | CustomAttributeTemplate internal (UUID) identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ProjectsProjectIdAttributesTemplatesTemplateIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2ProjectsProjectIdAttributesTemplatesTemplateIdPost

> ApiV2ProjectsProjectIdAttributesTemplatesTemplateIdPost(ctx, projectId, templateId).Execute()

Add CustomAttributeTemplate to Project



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
	templateId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | CustomAttributeTemplate internal (UUID) identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectAttributeTemplatesAPI.ApiV2ProjectsProjectIdAttributesTemplatesTemplateIdPost(context.Background(), projectId, templateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAttributeTemplatesAPI.ApiV2ProjectsProjectIdAttributesTemplatesTemplateIdPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project internal (UUID) or global (integer) identifier | 
**templateId** | **string** | CustomAttributeTemplate internal (UUID) identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ProjectsProjectIdAttributesTemplatesTemplateIdPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

