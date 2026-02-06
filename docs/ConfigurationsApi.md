# \ConfigurationsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV2ConfigurationsCreateByParametersPost**](ConfigurationsAPI.md#ApiV2ConfigurationsCreateByParametersPost) | **Post** /api/v2/configurations/createByParameters | Create configurations by parameters
[**ApiV2ConfigurationsDeleteBulkPost**](ConfigurationsAPI.md#ApiV2ConfigurationsDeleteBulkPost) | **Post** /api/v2/configurations/delete/bulk | Delete multiple configurations
[**ApiV2ConfigurationsIdDelete**](ConfigurationsAPI.md#ApiV2ConfigurationsIdDelete) | **Delete** /api/v2/configurations/{id} | Delete configuration
[**ApiV2ConfigurationsIdPatch**](ConfigurationsAPI.md#ApiV2ConfigurationsIdPatch) | **Patch** /api/v2/configurations/{id} | Patch configuration
[**ApiV2ConfigurationsIdPurgePost**](ConfigurationsAPI.md#ApiV2ConfigurationsIdPurgePost) | **Post** /api/v2/configurations/{id}/purge | Permanently delete configuration from archive
[**ApiV2ConfigurationsIdRestorePost**](ConfigurationsAPI.md#ApiV2ConfigurationsIdRestorePost) | **Post** /api/v2/configurations/{id}/restore | Restore configuration from the archive
[**ApiV2ConfigurationsPurgeBulkPost**](ConfigurationsAPI.md#ApiV2ConfigurationsPurgeBulkPost) | **Post** /api/v2/configurations/purge/bulk | Permanently delete multiple archived configurations
[**ApiV2ConfigurationsPut**](ConfigurationsAPI.md#ApiV2ConfigurationsPut) | **Put** /api/v2/configurations | Edit configuration
[**ApiV2ConfigurationsRestoreBulkPost**](ConfigurationsAPI.md#ApiV2ConfigurationsRestoreBulkPost) | **Post** /api/v2/configurations/restore/bulk | Restore multiple configurations from the archive
[**ApiV2ConfigurationsSearchPost**](ConfigurationsAPI.md#ApiV2ConfigurationsSearchPost) | **Post** /api/v2/configurations/search | Search for configurations
[**CreateConfiguration**](ConfigurationsAPI.md#CreateConfiguration) | **Post** /api/v2/configurations | Create Configuration
[**GetConfigurationById**](ConfigurationsAPI.md#GetConfigurationById) | **Get** /api/v2/configurations/{id} | Get configuration by internal or global ID



## ApiV2ConfigurationsCreateByParametersPost

> []string ApiV2ConfigurationsCreateByParametersPost(ctx).ConfigurationByParametersModel(configurationByParametersModel).Execute()

Create configurations by parameters

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
	configurationByParametersModel := *openapiclient.NewConfigurationByParametersModel("a58827bc-4fbb-4b8d-8ddc-bfaa97dbd0d5", []string{"ParameterIds_example"}) // ConfigurationByParametersModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigurationsAPI.ApiV2ConfigurationsCreateByParametersPost(context.Background()).ConfigurationByParametersModel(configurationByParametersModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationsAPI.ApiV2ConfigurationsCreateByParametersPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2ConfigurationsCreateByParametersPost`: []string
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationsAPI.ApiV2ConfigurationsCreateByParametersPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ConfigurationsCreateByParametersPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **configurationByParametersModel** | [**ConfigurationByParametersModel**](ConfigurationByParametersModel.md) |  | 

### Return type

**[]string**

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2ConfigurationsDeleteBulkPost

> int32 ApiV2ConfigurationsDeleteBulkPost(ctx).ConfigurationSelectApiModel(configurationSelectApiModel).Execute()

Delete multiple configurations

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
	configurationSelectApiModel := *openapiclient.NewConfigurationSelectApiModel() // ConfigurationSelectApiModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigurationsAPI.ApiV2ConfigurationsDeleteBulkPost(context.Background()).ConfigurationSelectApiModel(configurationSelectApiModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationsAPI.ApiV2ConfigurationsDeleteBulkPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2ConfigurationsDeleteBulkPost`: int32
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationsAPI.ApiV2ConfigurationsDeleteBulkPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ConfigurationsDeleteBulkPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **configurationSelectApiModel** | [**ConfigurationSelectApiModel**](ConfigurationSelectApiModel.md) |  | 

### Return type

**int32**

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2ConfigurationsIdDelete

> ApiV2ConfigurationsIdDelete(ctx, id).Execute()

Delete configuration

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
	id := "id_example" // string | Unique or global ID of the configuration

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ConfigurationsAPI.ApiV2ConfigurationsIdDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationsAPI.ApiV2ConfigurationsIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Unique or global ID of the configuration | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ConfigurationsIdDeleteRequest struct via the builder pattern


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


## ApiV2ConfigurationsIdPatch

> ApiV2ConfigurationsIdPatch(ctx, id).Operation(operation).Execute()

Patch configuration



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique ID of the configuration
	operation := []openapiclient.Operation{*openapiclient.NewOperation()} // []Operation |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ConfigurationsAPI.ApiV2ConfigurationsIdPatch(context.Background(), id).Operation(operation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationsAPI.ApiV2ConfigurationsIdPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Unique ID of the configuration | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ConfigurationsIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **operation** | [**[]Operation**](Operation.md) |  | 

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


## ApiV2ConfigurationsIdPurgePost

> ApiV2ConfigurationsIdPurgePost(ctx, id).Execute()

Permanently delete configuration from archive

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
	id := "id_example" // string | Unique or global ID of the configuration

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ConfigurationsAPI.ApiV2ConfigurationsIdPurgePost(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationsAPI.ApiV2ConfigurationsIdPurgePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Unique or global ID of the configuration | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ConfigurationsIdPurgePostRequest struct via the builder pattern


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


## ApiV2ConfigurationsIdRestorePost

> ApiV2ConfigurationsIdRestorePost(ctx, id).Execute()

Restore configuration from the archive

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
	id := "id_example" // string | Unique or global ID of the configuration

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ConfigurationsAPI.ApiV2ConfigurationsIdRestorePost(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationsAPI.ApiV2ConfigurationsIdRestorePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Unique or global ID of the configuration | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ConfigurationsIdRestorePostRequest struct via the builder pattern


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


## ApiV2ConfigurationsPurgeBulkPost

> int32 ApiV2ConfigurationsPurgeBulkPost(ctx).ConfigurationSelectModel(configurationSelectModel).Execute()

Permanently delete multiple archived configurations

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
	configurationSelectModel := *openapiclient.NewConfigurationSelectModel() // ConfigurationSelectModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigurationsAPI.ApiV2ConfigurationsPurgeBulkPost(context.Background()).ConfigurationSelectModel(configurationSelectModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationsAPI.ApiV2ConfigurationsPurgeBulkPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2ConfigurationsPurgeBulkPost`: int32
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationsAPI.ApiV2ConfigurationsPurgeBulkPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ConfigurationsPurgeBulkPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **configurationSelectModel** | [**ConfigurationSelectModel**](ConfigurationSelectModel.md) |  | 

### Return type

**int32**

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2ConfigurationsPut

> ApiV2ConfigurationsPut(ctx).ConfigurationPutModel(configurationPutModel).Execute()

Edit configuration

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
	configurationPutModel := *openapiclient.NewConfigurationPutModel("a58827bc-4fbb-4b8d-8ddc-bfaa97dbd0d5", map[string]string{"key": "Inner_example"}, "ProjectId_example", true, "Default") // ConfigurationPutModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ConfigurationsAPI.ApiV2ConfigurationsPut(context.Background()).ConfigurationPutModel(configurationPutModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationsAPI.ApiV2ConfigurationsPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ConfigurationsPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **configurationPutModel** | [**ConfigurationPutModel**](ConfigurationPutModel.md) |  | 

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


## ApiV2ConfigurationsRestoreBulkPost

> int32 ApiV2ConfigurationsRestoreBulkPost(ctx).ConfigurationSelectModel(configurationSelectModel).Execute()

Restore multiple configurations from the archive

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
	configurationSelectModel := *openapiclient.NewConfigurationSelectModel() // ConfigurationSelectModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigurationsAPI.ApiV2ConfigurationsRestoreBulkPost(context.Background()).ConfigurationSelectModel(configurationSelectModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationsAPI.ApiV2ConfigurationsRestoreBulkPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2ConfigurationsRestoreBulkPost`: int32
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationsAPI.ApiV2ConfigurationsRestoreBulkPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ConfigurationsRestoreBulkPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **configurationSelectModel** | [**ConfigurationSelectModel**](ConfigurationSelectModel.md) |  | 

### Return type

**int32**

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2ConfigurationsSearchPost

> []ConfigurationModel ApiV2ConfigurationsSearchPost(ctx).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).ConfigurationFilterModel(configurationFilterModel).Execute()

Search for configurations

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
	skip := int32(56) // int32 | Amount of items to be skipped (offset) (optional)
	take := int32(56) // int32 | Amount of items to be taken (limit) (optional)
	orderBy := "orderBy_example" // string | SQL-like  ORDER BY statement (column1 ASC|DESC , column2 ASC|DESC) (optional)
	searchField := "searchField_example" // string | Property name for searching (optional)
	searchValue := "searchValue_example" // string | Value for searching (optional)
	configurationFilterModel := *openapiclient.NewConfigurationFilterModel() // ConfigurationFilterModel | Model containing all the filters (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigurationsAPI.ApiV2ConfigurationsSearchPost(context.Background()).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).ConfigurationFilterModel(configurationFilterModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationsAPI.ApiV2ConfigurationsSearchPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2ConfigurationsSearchPost`: []ConfigurationModel
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationsAPI.ApiV2ConfigurationsSearchPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ConfigurationsSearchPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skip** | **int32** | Amount of items to be skipped (offset) | 
 **take** | **int32** | Amount of items to be taken (limit) | 
 **orderBy** | **string** | SQL-like  ORDER BY statement (column1 ASC|DESC , column2 ASC|DESC) | 
 **searchField** | **string** | Property name for searching | 
 **searchValue** | **string** | Value for searching | 
 **configurationFilterModel** | [**ConfigurationFilterModel**](ConfigurationFilterModel.md) | Model containing all the filters | 

### Return type

[**[]ConfigurationModel**](ConfigurationModel.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateConfiguration

> ConfigurationModel CreateConfiguration(ctx).ConfigurationPostModel(configurationPostModel).Execute()

Create Configuration



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
	configurationPostModel := *openapiclient.NewConfigurationPostModel(map[string]string{"key": "Inner_example"}, "ProjectId_example", true, "Default") // ConfigurationPostModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigurationsAPI.CreateConfiguration(context.Background()).ConfigurationPostModel(configurationPostModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationsAPI.CreateConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateConfiguration`: ConfigurationModel
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationsAPI.CreateConfiguration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **configurationPostModel** | [**ConfigurationPostModel**](ConfigurationPostModel.md) |  | 

### Return type

[**ConfigurationModel**](ConfigurationModel.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConfigurationById

> ConfigurationModel GetConfigurationById(ctx, id).Execute()

Get configuration by internal or global ID



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
	id := "3fa85f64-5717-4562-b3fc-2c963f66afa6" // string | Configuration internal (guid format) or global (integer format) identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigurationsAPI.GetConfigurationById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationsAPI.GetConfigurationById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConfigurationById`: ConfigurationModel
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationsAPI.GetConfigurationById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Configuration internal (guid format) or global (integer format) identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConfigurationByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ConfigurationModel**](ConfigurationModel.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

