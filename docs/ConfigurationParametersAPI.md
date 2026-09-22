# \ConfigurationParametersAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV2ConfigurationParametersConfigurationParameterIdDelete**](ConfigurationParametersAPI.md#ApiV2ConfigurationParametersConfigurationParameterIdDelete) | **Delete** /api/v2/configuration-parameters/{configurationParameterId} | Deletes configuration parameter
[**ApiV2ConfigurationParametersConfigurationParameterIdGet**](ConfigurationParametersAPI.md#ApiV2ConfigurationParametersConfigurationParameterIdGet) | **Get** /api/v2/configuration-parameters/{configurationParameterId} | Gets configuration parameter by its identifier
[**ApiV2ConfigurationParametersConfigurationParameterIdPut**](ConfigurationParametersAPI.md#ApiV2ConfigurationParametersConfigurationParameterIdPut) | **Put** /api/v2/configuration-parameters/{configurationParameterId} | Updates configuration parameter
[**ApiV2ConfigurationParametersPost**](ConfigurationParametersAPI.md#ApiV2ConfigurationParametersPost) | **Post** /api/v2/configuration-parameters | Creates new configuration parameter
[**ApiV2ConfigurationParametersSearchPost**](ConfigurationParametersAPI.md#ApiV2ConfigurationParametersSearchPost) | **Post** /api/v2/configuration-parameters/search | Searches for configuration parameters



## ApiV2ConfigurationParametersConfigurationParameterIdDelete

> ApiV2ConfigurationParametersConfigurationParameterIdDelete(ctx, configurationParameterId).Execute()

Deletes configuration parameter

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
	configurationParameterId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ConfigurationParametersAPI.ApiV2ConfigurationParametersConfigurationParameterIdDelete(context.Background(), configurationParameterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationParametersAPI.ApiV2ConfigurationParametersConfigurationParameterIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configurationParameterId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ConfigurationParametersConfigurationParameterIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[PrivateToken](../README.md#PrivateToken), [Identity.Application](../README.md#Identity.Application)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2ConfigurationParametersConfigurationParameterIdGet

> ConfigurationParameterApiResult ApiV2ConfigurationParametersConfigurationParameterIdGet(ctx, configurationParameterId).Execute()

Gets configuration parameter by its identifier

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
	configurationParameterId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigurationParametersAPI.ApiV2ConfigurationParametersConfigurationParameterIdGet(context.Background(), configurationParameterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationParametersAPI.ApiV2ConfigurationParametersConfigurationParameterIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2ConfigurationParametersConfigurationParameterIdGet`: ConfigurationParameterApiResult
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationParametersAPI.ApiV2ConfigurationParametersConfigurationParameterIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configurationParameterId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ConfigurationParametersConfigurationParameterIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ConfigurationParameterApiResult**](ConfigurationParameterApiResult.md)

### Authorization

[PrivateToken](../README.md#PrivateToken), [Identity.Application](../README.md#Identity.Application)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2ConfigurationParametersConfigurationParameterIdPut

> ApiV2ConfigurationParametersConfigurationParameterIdPut(ctx, configurationParameterId).ConfigurationParameterApiModel(configurationParameterApiModel).Execute()

Updates configuration parameter

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
	configurationParameterId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	configurationParameterApiModel := *openapiclient.NewConfigurationParameterApiModel("Name_example", []openapiclient.ConfigurationParameterValueApiModel{*openapiclient.NewConfigurationParameterValueApiModel("Value_example")}) // ConfigurationParameterApiModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ConfigurationParametersAPI.ApiV2ConfigurationParametersConfigurationParameterIdPut(context.Background(), configurationParameterId).ConfigurationParameterApiModel(configurationParameterApiModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationParametersAPI.ApiV2ConfigurationParametersConfigurationParameterIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configurationParameterId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ConfigurationParametersConfigurationParameterIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **configurationParameterApiModel** | [**ConfigurationParameterApiModel**](ConfigurationParameterApiModel.md) |  | 

### Return type

 (empty response body)

### Authorization

[PrivateToken](../README.md#PrivateToken), [Identity.Application](../README.md#Identity.Application)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2ConfigurationParametersPost

> ConfigurationParameterApiResult ApiV2ConfigurationParametersPost(ctx).ConfigurationParameterApiModel(configurationParameterApiModel).Execute()

Creates new configuration parameter

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
	configurationParameterApiModel := *openapiclient.NewConfigurationParameterApiModel("Name_example", []openapiclient.ConfigurationParameterValueApiModel{*openapiclient.NewConfigurationParameterValueApiModel("Value_example")}) // ConfigurationParameterApiModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigurationParametersAPI.ApiV2ConfigurationParametersPost(context.Background()).ConfigurationParameterApiModel(configurationParameterApiModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationParametersAPI.ApiV2ConfigurationParametersPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2ConfigurationParametersPost`: ConfigurationParameterApiResult
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationParametersAPI.ApiV2ConfigurationParametersPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ConfigurationParametersPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **configurationParameterApiModel** | [**ConfigurationParameterApiModel**](ConfigurationParameterApiModel.md) |  | 

### Return type

[**ConfigurationParameterApiResult**](ConfigurationParameterApiResult.md)

### Authorization

[PrivateToken](../README.md#PrivateToken), [Identity.Application](../README.md#Identity.Application)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2ConfigurationParametersSearchPost

> ConfigurationParameterPreviewApiResultIReply ApiV2ConfigurationParametersSearchPost(ctx).SearchConfigurationParametersApiModel(searchConfigurationParametersApiModel).Execute()

Searches for configuration parameters

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
	searchConfigurationParametersApiModel := *openapiclient.NewSearchConfigurationParametersApiModel() // SearchConfigurationParametersApiModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigurationParametersAPI.ApiV2ConfigurationParametersSearchPost(context.Background()).SearchConfigurationParametersApiModel(searchConfigurationParametersApiModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationParametersAPI.ApiV2ConfigurationParametersSearchPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2ConfigurationParametersSearchPost`: ConfigurationParameterPreviewApiResultIReply
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationParametersAPI.ApiV2ConfigurationParametersSearchPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ConfigurationParametersSearchPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchConfigurationParametersApiModel** | [**SearchConfigurationParametersApiModel**](SearchConfigurationParametersApiModel.md) |  | 

### Return type

[**ConfigurationParameterPreviewApiResultIReply**](ConfigurationParameterPreviewApiResultIReply.md)

### Authorization

[PrivateToken](../README.md#PrivateToken), [Identity.Application](../README.md#Identity.Application)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

