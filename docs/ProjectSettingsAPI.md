# \ProjectSettingsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV2ProjectsProjectIdSettingsAutotestsPost**](ProjectSettingsAPI.md#ApiV2ProjectsProjectIdSettingsAutotestsPost) | **Post** /api/v2/projects/{projectId}/settings/autotests | Set autotest project settings.
[**GetAutotestProjectSettings**](ProjectSettingsAPI.md#GetAutotestProjectSettings) | **Get** /api/v2/projects/{projectId}/settings/autotests | Get autotest project settings.



## ApiV2ProjectsProjectIdSettingsAutotestsPost

> ApiV2ProjectsProjectIdSettingsAutotestsPost(ctx, projectId).AutoTestProjectSettingsPostModel(autoTestProjectSettingsPostModel).Execute()

Set autotest project settings.

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
	autoTestProjectSettingsPostModel := *openapiclient.NewAutoTestProjectSettingsPostModel(false, int32(123)) // AutoTestProjectSettingsPostModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectSettingsAPI.ApiV2ProjectsProjectIdSettingsAutotestsPost(context.Background(), projectId).AutoTestProjectSettingsPostModel(autoTestProjectSettingsPostModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectSettingsAPI.ApiV2ProjectsProjectIdSettingsAutotestsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ProjectsProjectIdSettingsAutotestsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **autoTestProjectSettingsPostModel** | [**AutoTestProjectSettingsPostModel**](AutoTestProjectSettingsPostModel.md) |  | 

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


## GetAutotestProjectSettings

> AutoTestProjectSettingsGetModel GetAutotestProjectSettings(ctx, projectId).Execute()

Get autotest project settings.

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectSettingsAPI.GetAutotestProjectSettings(context.Background(), projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectSettingsAPI.GetAutotestProjectSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAutotestProjectSettings`: AutoTestProjectSettingsGetModel
	fmt.Fprintf(os.Stdout, "Response from `ProjectSettingsAPI.GetAutotestProjectSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAutotestProjectSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AutoTestProjectSettingsGetModel**](AutoTestProjectSettingsGetModel.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

