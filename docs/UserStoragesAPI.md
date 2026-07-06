# \UserStoragesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV2UserStoragesResourceGet**](UserStoragesAPI.md#ApiV2UserStoragesResourceGet) | **Get** /api/v2/user-storages/{resource} | 
[**ApiV2UserStoragesResourcePost**](UserStoragesAPI.md#ApiV2UserStoragesResourcePost) | **Post** /api/v2/user-storages/{resource} | 



## ApiV2UserStoragesResourceGet

> ApiV2UserStoragesResourceGet(ctx, resource).Execute()



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
	resource := "resource_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UserStoragesAPI.ApiV2UserStoragesResourceGet(context.Background(), resource).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserStoragesAPI.ApiV2UserStoragesResourceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resource** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2UserStoragesResourceGetRequest struct via the builder pattern


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


## ApiV2UserStoragesResourcePost

> ApiV2UserStoragesResourcePost(ctx, resource).Execute()



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
	resource := "resource_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UserStoragesAPI.ApiV2UserStoragesResourcePost(context.Background(), resource).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserStoragesAPI.ApiV2UserStoragesResourcePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resource** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2UserStoragesResourcePostRequest struct via the builder pattern


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

