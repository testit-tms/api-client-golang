# \AttachmentsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV2AttachmentsIdDelete**](AttachmentsAPI.md#ApiV2AttachmentsIdDelete) | **Delete** /api/v2/attachments/{id} | Delete attachment file
[**ApiV2AttachmentsIdGet**](AttachmentsAPI.md#ApiV2AttachmentsIdGet) | **Get** /api/v2/attachments/{id} | Download attachment file
[**ApiV2AttachmentsIdMetadataGet**](AttachmentsAPI.md#ApiV2AttachmentsIdMetadataGet) | **Get** /api/v2/attachments/{id}/metadata | Get attachment metadata
[**ApiV2AttachmentsOccupiedFileStorageSizeGet**](AttachmentsAPI.md#ApiV2AttachmentsOccupiedFileStorageSizeGet) | **Get** /api/v2/attachments/occupiedFileStorageSize | Get size of attachments storage in bytes
[**ApiV2AttachmentsPost**](AttachmentsAPI.md#ApiV2AttachmentsPost) | **Post** /api/v2/attachments | Upload new attachment file



## ApiV2AttachmentsIdDelete

> ApiV2AttachmentsIdDelete(ctx, id).Execute()

Delete attachment file

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
	r, err := apiClient.AttachmentsAPI.ApiV2AttachmentsIdDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AttachmentsAPI.ApiV2AttachmentsIdDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiApiV2AttachmentsIdDeleteRequest struct via the builder pattern


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


## ApiV2AttachmentsIdGet

> ApiV2AttachmentsIdGet(ctx, id).Width(width).Height(height).ResizeType(resizeType).BackgroundColor(backgroundColor).Preview(preview).Execute()

Download attachment file

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
	width := int32(56) // int32 | Width of the result image (optional)
	height := int32(56) // int32 | Height of the result image (optional)
	resizeType := openapiclient.ImageResizeType("Crop") // ImageResizeType | Type of resizing to apply to the result image (optional)
	backgroundColor := "backgroundColor_example" // string | Color of the background if the `resizeType` is `AddBackgroundStripes` (optional)
	preview := true // bool | If image must be converted to a preview (lower quality, no animation) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AttachmentsAPI.ApiV2AttachmentsIdGet(context.Background(), id).Width(width).Height(height).ResizeType(resizeType).BackgroundColor(backgroundColor).Preview(preview).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AttachmentsAPI.ApiV2AttachmentsIdGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiApiV2AttachmentsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **width** | **int32** | Width of the result image | 
 **height** | **int32** | Height of the result image | 
 **resizeType** | [**ImageResizeType**](ImageResizeType.md) | Type of resizing to apply to the result image | 
 **backgroundColor** | **string** | Color of the background if the &#x60;resizeType&#x60; is &#x60;AddBackgroundStripes&#x60; | 
 **preview** | **bool** | If image must be converted to a preview (lower quality, no animation) | 

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


## ApiV2AttachmentsIdMetadataGet

> AttachmentModel ApiV2AttachmentsIdMetadataGet(ctx, id).Execute()

Get attachment metadata

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
	resp, r, err := apiClient.AttachmentsAPI.ApiV2AttachmentsIdMetadataGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AttachmentsAPI.ApiV2AttachmentsIdMetadataGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2AttachmentsIdMetadataGet`: AttachmentModel
	fmt.Fprintf(os.Stdout, "Response from `AttachmentsAPI.ApiV2AttachmentsIdMetadataGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2AttachmentsIdMetadataGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AttachmentModel**](AttachmentModel.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2AttachmentsOccupiedFileStorageSizeGet

> int64 ApiV2AttachmentsOccupiedFileStorageSizeGet(ctx).Execute()

Get size of attachments storage in bytes

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
	resp, r, err := apiClient.AttachmentsAPI.ApiV2AttachmentsOccupiedFileStorageSizeGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AttachmentsAPI.ApiV2AttachmentsOccupiedFileStorageSizeGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2AttachmentsOccupiedFileStorageSizeGet`: int64
	fmt.Fprintf(os.Stdout, "Response from `AttachmentsAPI.ApiV2AttachmentsOccupiedFileStorageSizeGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2AttachmentsOccupiedFileStorageSizeGetRequest struct via the builder pattern


### Return type

**int64**

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2AttachmentsPost

> AttachmentModel ApiV2AttachmentsPost(ctx).File(file).Execute()

Upload new attachment file



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
	file := os.NewFile(1234, "some_file") // *os.File |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AttachmentsAPI.ApiV2AttachmentsPost(context.Background()).File(file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AttachmentsAPI.ApiV2AttachmentsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2AttachmentsPost`: AttachmentModel
	fmt.Fprintf(os.Stdout, "Response from `AttachmentsAPI.ApiV2AttachmentsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2AttachmentsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | ***os.File** |  | 

### Return type

[**AttachmentModel**](AttachmentModel.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

