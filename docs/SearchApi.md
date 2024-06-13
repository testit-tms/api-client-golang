# \SearchAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV2SearchGlobalSearchPost**](SearchAPI.md#ApiV2SearchGlobalSearchPost) | **Post** /api/v2/search/globalSearch | 



## ApiV2SearchGlobalSearchPost

> GlobalSearchResponse ApiV2SearchGlobalSearchPost(ctx).GlobalSearchRequest(globalSearchRequest).Execute()



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
	globalSearchRequest := *openapiclient.NewGlobalSearchRequest("Query_example", int32(123), int32(123)) // GlobalSearchRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchAPI.ApiV2SearchGlobalSearchPost(context.Background()).GlobalSearchRequest(globalSearchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.ApiV2SearchGlobalSearchPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2SearchGlobalSearchPost`: GlobalSearchResponse
	fmt.Fprintf(os.Stdout, "Response from `SearchAPI.ApiV2SearchGlobalSearchPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2SearchGlobalSearchPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **globalSearchRequest** | [**GlobalSearchRequest**](GlobalSearchRequest.md) |  | 

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

