# \AIServicesAPIAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV2ExternalServicesIdAiModelsPost**](AIServicesAPIAPI.md#ApiV2ExternalServicesIdAiModelsPost) | **Post** /api/v2/external-services/{id}/ai/models | Ask for models with inquiry filter, cached



## ApiV2ExternalServicesIdAiModelsPost

> AIServiceModelApiResultIReply ApiV2ExternalServicesIdAiModelsPost(ctx, id).GetAIServiceModelsApiModel(getAIServiceModelsApiModel).Execute()

Ask for models with inquiry filter, cached

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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	getAIServiceModelsApiModel := *openapiclient.NewGetAIServiceModelsApiModel() // GetAIServiceModelsApiModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIServicesAPIAPI.ApiV2ExternalServicesIdAiModelsPost(context.Background(), id).GetAIServiceModelsApiModel(getAIServiceModelsApiModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIServicesAPIAPI.ApiV2ExternalServicesIdAiModelsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2ExternalServicesIdAiModelsPost`: AIServiceModelApiResultIReply
	fmt.Fprintf(os.Stdout, "Response from `AIServicesAPIAPI.ApiV2ExternalServicesIdAiModelsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ExternalServicesIdAiModelsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **getAIServiceModelsApiModel** | [**GetAIServiceModelsApiModel**](GetAIServiceModelsApiModel.md) |  | 

### Return type

[**AIServiceModelApiResultIReply**](AIServiceModelApiResultIReply.md)

### Authorization

[PrivateToken](../README.md#PrivateToken), [Cookies](../README.md#Cookies)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

