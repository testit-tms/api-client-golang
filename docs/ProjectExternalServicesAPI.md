# \ProjectExternalServicesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV2ProjectsIdExternalServicesExternalServiceIdDelete**](ProjectExternalServicesAPI.md#ApiV2ProjectsIdExternalServicesExternalServiceIdDelete) | **Delete** /api/v2/projects/{id}/external-services/{externalServiceId} | Disable an external service
[**ApiV2ProjectsIdExternalServicesExternalServiceIdGet**](ProjectExternalServicesAPI.md#ApiV2ProjectsIdExternalServicesExternalServiceIdGet) | **Get** /api/v2/projects/{id}/external-services/{externalServiceId} | Retrieves settings of an external service
[**ApiV2ProjectsIdExternalServicesExternalServiceIdPatch**](ProjectExternalServicesAPI.md#ApiV2ProjectsIdExternalServicesExternalServiceIdPatch) | **Patch** /api/v2/projects/{id}/external-services/{externalServiceId} | Replaces one active external service with another
[**ApiV2ProjectsIdExternalServicesExternalServiceIdPut**](ProjectExternalServicesAPI.md#ApiV2ProjectsIdExternalServicesExternalServiceIdPut) | **Put** /api/v2/projects/{id}/external-services/{externalServiceId} | Enable an external service
[**ApiV2ProjectsIdExternalServicesGet**](ProjectExternalServicesAPI.md#ApiV2ProjectsIdExternalServicesGet) | **Get** /api/v2/projects/{id}/external-services | Retrieves information about external services, including their integration status (enabled or not)
[**ApiV2ProjectsIdExternalServicesIssuesSearchPost**](ProjectExternalServicesAPI.md#ApiV2ProjectsIdExternalServicesIssuesSearchPost) | **Post** /api/v2/projects/{id}/external-services/issues/search | Searches for external issues using enabled external services in project



## ApiV2ProjectsIdExternalServicesExternalServiceIdDelete

> ApiV2ProjectsIdExternalServicesExternalServiceIdDelete(ctx, id, externalServiceId).Execute()

Disable an external service

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
	id := "id_example" // string | Project ID
	externalServiceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | External service ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectExternalServicesAPI.ApiV2ProjectsIdExternalServicesExternalServiceIdDelete(context.Background(), id, externalServiceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectExternalServicesAPI.ApiV2ProjectsIdExternalServicesExternalServiceIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Project ID | 
**externalServiceId** | **string** | External service ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ProjectsIdExternalServicesExternalServiceIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[PrivateToken](../README.md#PrivateToken), [Cookies](../README.md#Cookies)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2ProjectsIdExternalServicesExternalServiceIdGet

> ProjectExternalServiceSettingsApiResult ApiV2ProjectsIdExternalServicesExternalServiceIdGet(ctx, id, externalServiceId).Execute()

Retrieves settings of an external service

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
	id := "id_example" // string | Project ID
	externalServiceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | External service ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectExternalServicesAPI.ApiV2ProjectsIdExternalServicesExternalServiceIdGet(context.Background(), id, externalServiceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectExternalServicesAPI.ApiV2ProjectsIdExternalServicesExternalServiceIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2ProjectsIdExternalServicesExternalServiceIdGet`: ProjectExternalServiceSettingsApiResult
	fmt.Fprintf(os.Stdout, "Response from `ProjectExternalServicesAPI.ApiV2ProjectsIdExternalServicesExternalServiceIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Project ID | 
**externalServiceId** | **string** | External service ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ProjectsIdExternalServicesExternalServiceIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ProjectExternalServiceSettingsApiResult**](ProjectExternalServiceSettingsApiResult.md)

### Authorization

[PrivateToken](../README.md#PrivateToken), [Cookies](../README.md#Cookies)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2ProjectsIdExternalServicesExternalServiceIdPatch

> ApiV2ProjectsIdExternalServicesExternalServiceIdPatch(ctx, id, externalServiceId).ReplaceProjectExternalServiceApiModel(replaceProjectExternalServiceApiModel).Execute()

Replaces one active external service with another



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
	id := "id_example" // string | Project ID
	externalServiceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | External service ID
	replaceProjectExternalServiceApiModel := *openapiclient.NewReplaceProjectExternalServiceApiModel("NewExternalServiceId_example", interface{}(123)) // ReplaceProjectExternalServiceApiModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectExternalServicesAPI.ApiV2ProjectsIdExternalServicesExternalServiceIdPatch(context.Background(), id, externalServiceId).ReplaceProjectExternalServiceApiModel(replaceProjectExternalServiceApiModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectExternalServicesAPI.ApiV2ProjectsIdExternalServicesExternalServiceIdPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Project ID | 
**externalServiceId** | **string** | External service ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ProjectsIdExternalServicesExternalServiceIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **replaceProjectExternalServiceApiModel** | [**ReplaceProjectExternalServiceApiModel**](ReplaceProjectExternalServiceApiModel.md) |  | 

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


## ApiV2ProjectsIdExternalServicesExternalServiceIdPut

> ApiV2ProjectsIdExternalServicesExternalServiceIdPut(ctx, id, externalServiceId).EnableProjectExternalServiceApiModel(enableProjectExternalServiceApiModel).Execute()

Enable an external service

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
	id := "id_example" // string | Project ID
	externalServiceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | External service ID
	enableProjectExternalServiceApiModel := *openapiclient.NewEnableProjectExternalServiceApiModel(interface{}(123)) // EnableProjectExternalServiceApiModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectExternalServicesAPI.ApiV2ProjectsIdExternalServicesExternalServiceIdPut(context.Background(), id, externalServiceId).EnableProjectExternalServiceApiModel(enableProjectExternalServiceApiModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectExternalServicesAPI.ApiV2ProjectsIdExternalServicesExternalServiceIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Project ID | 
**externalServiceId** | **string** | External service ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ProjectsIdExternalServicesExternalServiceIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **enableProjectExternalServiceApiModel** | [**EnableProjectExternalServiceApiModel**](EnableProjectExternalServiceApiModel.md) |  | 

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


## ApiV2ProjectsIdExternalServicesGet

> ProjectExternalServicesApiResult ApiV2ProjectsIdExternalServicesGet(ctx, id).Category(category).Execute()

Retrieves information about external services, including their integration status (enabled or not)

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
	id := "id_example" // string | Project ID
	category := openapiclient.ApiExternalServiceCategory("AI") // ApiExternalServiceCategory |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectExternalServicesAPI.ApiV2ProjectsIdExternalServicesGet(context.Background(), id).Category(category).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectExternalServicesAPI.ApiV2ProjectsIdExternalServicesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2ProjectsIdExternalServicesGet`: ProjectExternalServicesApiResult
	fmt.Fprintf(os.Stdout, "Response from `ProjectExternalServicesAPI.ApiV2ProjectsIdExternalServicesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ProjectsIdExternalServicesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **category** | [**ApiExternalServiceCategory**](ApiExternalServiceCategory.md) |  | 

### Return type

[**ProjectExternalServicesApiResult**](ProjectExternalServicesApiResult.md)

### Authorization

[PrivateToken](../README.md#PrivateToken), [Cookies](../README.md#Cookies)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2ProjectsIdExternalServicesIssuesSearchPost

> []ExternalIssueApiResult ApiV2ProjectsIdExternalServicesIssuesSearchPost(ctx, id).SearchExternalIssuesApiModel(searchExternalIssuesApiModel).Execute()

Searches for external issues using enabled external services in project

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
	id := "id_example" // string | Internal (UUID) or global (integer) identifier
	searchExternalIssuesApiModel := *openapiclient.NewSearchExternalIssuesApiModel("Url_example") // SearchExternalIssuesApiModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectExternalServicesAPI.ApiV2ProjectsIdExternalServicesIssuesSearchPost(context.Background(), id).SearchExternalIssuesApiModel(searchExternalIssuesApiModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectExternalServicesAPI.ApiV2ProjectsIdExternalServicesIssuesSearchPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2ProjectsIdExternalServicesIssuesSearchPost`: []ExternalIssueApiResult
	fmt.Fprintf(os.Stdout, "Response from `ProjectExternalServicesAPI.ApiV2ProjectsIdExternalServicesIssuesSearchPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Internal (UUID) or global (integer) identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ProjectsIdExternalServicesIssuesSearchPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **searchExternalIssuesApiModel** | [**SearchExternalIssuesApiModel**](SearchExternalIssuesApiModel.md) |  | 

### Return type

[**[]ExternalIssueApiResult**](ExternalIssueApiResult.md)

### Authorization

[PrivateToken](../README.md#PrivateToken), [Cookies](../README.md#Cookies)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

