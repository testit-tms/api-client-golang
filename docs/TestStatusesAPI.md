# \TestStatusesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV2TestStatusesCodeCodeExistsGet**](TestStatusesAPI.md#ApiV2TestStatusesCodeCodeExistsGet) | **Get** /api/v2/testStatuses/code/{code}/exists | 
[**ApiV2TestStatusesIdDelete**](TestStatusesAPI.md#ApiV2TestStatusesIdDelete) | **Delete** /api/v2/testStatuses/{id} | 
[**ApiV2TestStatusesIdGet**](TestStatusesAPI.md#ApiV2TestStatusesIdGet) | **Get** /api/v2/testStatuses/{id} | 
[**ApiV2TestStatusesIdPut**](TestStatusesAPI.md#ApiV2TestStatusesIdPut) | **Put** /api/v2/testStatuses/{id} | 
[**ApiV2TestStatusesNameNameExistsGet**](TestStatusesAPI.md#ApiV2TestStatusesNameNameExistsGet) | **Get** /api/v2/testStatuses/name/{name}/exists | 
[**ApiV2TestStatusesPost**](TestStatusesAPI.md#ApiV2TestStatusesPost) | **Post** /api/v2/testStatuses | 
[**ApiV2TestStatusesSearchPost**](TestStatusesAPI.md#ApiV2TestStatusesSearchPost) | **Post** /api/v2/testStatuses/search | 



## ApiV2TestStatusesCodeCodeExistsGet

> bool ApiV2TestStatusesCodeCodeExistsGet(ctx, code).Execute()



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
	code := "code_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TestStatusesAPI.ApiV2TestStatusesCodeCodeExistsGet(context.Background(), code).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestStatusesAPI.ApiV2TestStatusesCodeCodeExistsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2TestStatusesCodeCodeExistsGet`: bool
	fmt.Fprintf(os.Stdout, "Response from `TestStatusesAPI.ApiV2TestStatusesCodeCodeExistsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2TestStatusesCodeCodeExistsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**bool**

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2TestStatusesIdDelete

> ApiV2TestStatusesIdDelete(ctx, id).Execute()



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TestStatusesAPI.ApiV2TestStatusesIdDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestStatusesAPI.ApiV2TestStatusesIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2TestStatusesIdDeleteRequest struct via the builder pattern


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


## ApiV2TestStatusesIdGet

> TestStatusApiResult ApiV2TestStatusesIdGet(ctx, id).Execute()



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TestStatusesAPI.ApiV2TestStatusesIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestStatusesAPI.ApiV2TestStatusesIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2TestStatusesIdGet`: TestStatusApiResult
	fmt.Fprintf(os.Stdout, "Response from `TestStatusesAPI.ApiV2TestStatusesIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2TestStatusesIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TestStatusApiResult**](TestStatusApiResult.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2TestStatusesIdPut

> ApiV2TestStatusesIdPut(ctx, id).UpdateTestStatusApiModel(updateTestStatusApiModel).Execute()



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
	updateTestStatusApiModel := *openapiclient.NewUpdateTestStatusApiModel("Name_example") // UpdateTestStatusApiModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TestStatusesAPI.ApiV2TestStatusesIdPut(context.Background(), id).UpdateTestStatusApiModel(updateTestStatusApiModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestStatusesAPI.ApiV2TestStatusesIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2TestStatusesIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateTestStatusApiModel** | [**UpdateTestStatusApiModel**](UpdateTestStatusApiModel.md) |  | 

### Return type

 (empty response body)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2TestStatusesNameNameExistsGet

> bool ApiV2TestStatusesNameNameExistsGet(ctx, name).Execute()



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
	name := "name_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TestStatusesAPI.ApiV2TestStatusesNameNameExistsGet(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestStatusesAPI.ApiV2TestStatusesNameNameExistsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2TestStatusesNameNameExistsGet`: bool
	fmt.Fprintf(os.Stdout, "Response from `TestStatusesAPI.ApiV2TestStatusesNameNameExistsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2TestStatusesNameNameExistsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**bool**

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2TestStatusesPost

> TestStatusApiResult ApiV2TestStatusesPost(ctx).CreateTestStatusApiModel(createTestStatusApiModel).Execute()



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
	createTestStatusApiModel := *openapiclient.NewCreateTestStatusApiModel("Name_example", openapiclient.TestStatusApiType("Pending"), "Code_example") // CreateTestStatusApiModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TestStatusesAPI.ApiV2TestStatusesPost(context.Background()).CreateTestStatusApiModel(createTestStatusApiModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestStatusesAPI.ApiV2TestStatusesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2TestStatusesPost`: TestStatusApiResult
	fmt.Fprintf(os.Stdout, "Response from `TestStatusesAPI.ApiV2TestStatusesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2TestStatusesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createTestStatusApiModel** | [**CreateTestStatusApiModel**](CreateTestStatusApiModel.md) |  | 

### Return type

[**TestStatusApiResult**](TestStatusApiResult.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2TestStatusesSearchPost

> TestStatusApiResultReply ApiV2TestStatusesSearchPost(ctx).SearchTestStatusesApiModel(searchTestStatusesApiModel).Execute()



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
	searchTestStatusesApiModel := *openapiclient.NewSearchTestStatusesApiModel() // SearchTestStatusesApiModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TestStatusesAPI.ApiV2TestStatusesSearchPost(context.Background()).SearchTestStatusesApiModel(searchTestStatusesApiModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestStatusesAPI.ApiV2TestStatusesSearchPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2TestStatusesSearchPost`: TestStatusApiResultReply
	fmt.Fprintf(os.Stdout, "Response from `TestStatusesAPI.ApiV2TestStatusesSearchPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2TestStatusesSearchPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchTestStatusesApiModel** | [**SearchTestStatusesApiModel**](SearchTestStatusesApiModel.md) |  | 

### Return type

[**TestStatusApiResultReply**](TestStatusApiResultReply.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

