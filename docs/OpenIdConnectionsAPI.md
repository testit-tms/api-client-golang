# \OpenIdConnectionsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV2OpenidConnectionsGet**](OpenIdConnectionsAPI.md#ApiV2OpenidConnectionsGet) | **Get** /api/v2/openid-connections | 



## ApiV2OpenidConnectionsGet

> []OpenIdConnectionClientShortModel ApiV2OpenidConnectionsGet(ctx).Execute()



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OpenIdConnectionsAPI.ApiV2OpenidConnectionsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpenIdConnectionsAPI.ApiV2OpenidConnectionsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2OpenidConnectionsGet`: []OpenIdConnectionClientShortModel
	fmt.Fprintf(os.Stdout, "Response from `OpenIdConnectionsAPI.ApiV2OpenidConnectionsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2OpenidConnectionsGetRequest struct via the builder pattern


### Return type

[**[]OpenIdConnectionClientShortModel**](OpenIdConnectionClientShortModel.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

