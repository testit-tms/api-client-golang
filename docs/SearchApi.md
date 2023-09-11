# \SearchApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV2SearchGlobalSearchPost**](SearchApi.md#ApiV2SearchGlobalSearchPost) | **Post** /api/v2/search/globalSearch | 



## ApiV2SearchGlobalSearchPost

> GlobalSearchResponse ApiV2SearchGlobalSearchPost(ctx).ApiV2SearchGlobalSearchPostRequest(apiV2SearchGlobalSearchPostRequest).Execute()



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
    apiV2SearchGlobalSearchPostRequest := *openapiclient.NewApiV2SearchGlobalSearchPostRequest("Query_example", int32(123), int32(123)) // ApiV2SearchGlobalSearchPostRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SearchApi.ApiV2SearchGlobalSearchPost(context.Background()).ApiV2SearchGlobalSearchPostRequest(apiV2SearchGlobalSearchPostRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchApi.ApiV2SearchGlobalSearchPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV2SearchGlobalSearchPost`: GlobalSearchResponse
    fmt.Fprintf(os.Stdout, "Response from `SearchApi.ApiV2SearchGlobalSearchPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2SearchGlobalSearchPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiV2SearchGlobalSearchPostRequest** | [**ApiV2SearchGlobalSearchPostRequest**](ApiV2SearchGlobalSearchPostRequest.md) |  | 

### Return type

[**GlobalSearchResponse**](GlobalSearchResponse.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

