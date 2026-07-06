# \ExternalServicesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV2ExternalServicesMetadataGet**](ExternalServicesAPI.md#ApiV2ExternalServicesMetadataGet) | **Get** /api/v2/external-services/metadata | Retrieves the metadata for all available external services



## ApiV2ExternalServicesMetadataGet

> ExternalServicesMetadataApiResult ApiV2ExternalServicesMetadataGet(ctx).Execute()

Retrieves the metadata for all available external services

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
	resp, r, err := apiClient.ExternalServicesAPI.ApiV2ExternalServicesMetadataGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExternalServicesAPI.ApiV2ExternalServicesMetadataGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2ExternalServicesMetadataGet`: ExternalServicesMetadataApiResult
	fmt.Fprintf(os.Stdout, "Response from `ExternalServicesAPI.ApiV2ExternalServicesMetadataGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ExternalServicesMetadataGetRequest struct via the builder pattern


### Return type

[**ExternalServicesMetadataApiResult**](ExternalServicesMetadataApiResult.md)

### Authorization

[PrivateToken](../README.md#PrivateToken), [Cookies](../README.md#Cookies)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

