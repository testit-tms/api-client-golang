# \ParametersAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV2ParametersBulkPost**](ParametersAPI.md#ApiV2ParametersBulkPost) | **Post** /api/v2/parameters/bulk | Create multiple parameters
[**ApiV2ParametersBulkPut**](ParametersAPI.md#ApiV2ParametersBulkPut) | **Put** /api/v2/parameters/bulk | Update multiple parameters
[**ApiV2ParametersGroupsGet**](ParametersAPI.md#ApiV2ParametersGroupsGet) | **Get** /api/v2/parameters/groups | Get parameters as group
[**ApiV2ParametersKeyNameNameExistsGet**](ParametersAPI.md#ApiV2ParametersKeyNameNameExistsGet) | **Get** /api/v2/parameters/key/name/{name}/exists | Check existence parameter key in system
[**ApiV2ParametersKeyValuesGet**](ParametersAPI.md#ApiV2ParametersKeyValuesGet) | **Get** /api/v2/parameters/{key}/values | Get all parameter key values
[**ApiV2ParametersKeysGet**](ParametersAPI.md#ApiV2ParametersKeysGet) | **Get** /api/v2/parameters/keys | Get all parameter keys
[**ApiV2ParametersSearchGroupsPost**](ParametersAPI.md#ApiV2ParametersSearchGroupsPost) | **Post** /api/v2/parameters/search/groups | Search for parameters as group
[**ApiV2ParametersSearchPost**](ParametersAPI.md#ApiV2ParametersSearchPost) | **Post** /api/v2/parameters/search | Search for parameters
[**CreateParameter**](ParametersAPI.md#CreateParameter) | **Post** /api/v2/parameters | Create parameter
[**DeleteByName**](ParametersAPI.md#DeleteByName) | **Delete** /api/v2/parameters/name/{name} | Delete parameter by name
[**DeleteByParameterKeyId**](ParametersAPI.md#DeleteByParameterKeyId) | **Delete** /api/v2/parameters/keyId/{keyId} | Delete parameters by parameter key identifier
[**DeleteParameter**](ParametersAPI.md#DeleteParameter) | **Delete** /api/v2/parameters/{id} | Delete parameter
[**GetAllParameters**](ParametersAPI.md#GetAllParameters) | **Get** /api/v2/parameters | Get all parameters
[**GetParameterById**](ParametersAPI.md#GetParameterById) | **Get** /api/v2/parameters/{id} | Get parameter by ID
[**UpdateParameter**](ParametersAPI.md#UpdateParameter) | **Put** /api/v2/parameters | Update parameter



## ApiV2ParametersBulkPost

> []ParameterApiResult ApiV2ParametersBulkPost(ctx).CreateParameterApiModel(createParameterApiModel).Execute()

Create multiple parameters



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
	createParameterApiModel := []openapiclient.CreateParameterApiModel{*openapiclient.NewCreateParameterApiModel("Name_example", "Value_example")} // []CreateParameterApiModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ParametersAPI.ApiV2ParametersBulkPost(context.Background()).CreateParameterApiModel(createParameterApiModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ParametersAPI.ApiV2ParametersBulkPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2ParametersBulkPost`: []ParameterApiResult
	fmt.Fprintf(os.Stdout, "Response from `ParametersAPI.ApiV2ParametersBulkPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ParametersBulkPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createParameterApiModel** | [**[]CreateParameterApiModel**](CreateParameterApiModel.md) |  | 

### Return type

[**[]ParameterApiResult**](ParameterApiResult.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2ParametersBulkPut

> ApiV2ParametersBulkPut(ctx).UpdateParameterApiModel(updateParameterApiModel).Execute()

Update multiple parameters



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
	updateParameterApiModel := []openapiclient.UpdateParameterApiModel{*openapiclient.NewUpdateParameterApiModel("Id_example", "Name_example", "Value_example")} // []UpdateParameterApiModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ParametersAPI.ApiV2ParametersBulkPut(context.Background()).UpdateParameterApiModel(updateParameterApiModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ParametersAPI.ApiV2ParametersBulkPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ParametersBulkPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateParameterApiModel** | [**[]UpdateParameterApiModel**](UpdateParameterApiModel.md) |  | 

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


## ApiV2ParametersGroupsGet

> []ParameterGroupApiResult ApiV2ParametersGroupsGet(ctx).ParameterKeyIds(parameterKeyIds).Name(name).IsDeleted(isDeleted).ProjectIds(projectIds).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).Execute()

Get parameters as group



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
	parameterKeyIds := []string{"Inner_example"} // []string |  (optional)
	name := "name_example" // string |  (optional)
	isDeleted := true // bool |  (optional)
	projectIds := []*string{"Inner_example"} // []*string |  (optional)
	skip := int32(56) // int32 | Amount of items to be skipped (offset) (optional)
	take := int32(56) // int32 | Amount of items to be taken (limit) (optional)
	orderBy := "orderBy_example" // string | SQL-like  ORDER BY statement (column1 ASC|DESC , column2 ASC|DESC) (optional)
	searchField := "searchField_example" // string | Property name for searching (optional)
	searchValue := "searchValue_example" // string | Value for searching (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ParametersAPI.ApiV2ParametersGroupsGet(context.Background()).ParameterKeyIds(parameterKeyIds).Name(name).IsDeleted(isDeleted).ProjectIds(projectIds).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ParametersAPI.ApiV2ParametersGroupsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2ParametersGroupsGet`: []ParameterGroupApiResult
	fmt.Fprintf(os.Stdout, "Response from `ParametersAPI.ApiV2ParametersGroupsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ParametersGroupsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **parameterKeyIds** | **[]string** |  | 
 **name** | **string** |  | 
 **isDeleted** | **bool** |  | 
 **projectIds** | **[]string** |  | 
 **skip** | **int32** | Amount of items to be skipped (offset) | 
 **take** | **int32** | Amount of items to be taken (limit) | 
 **orderBy** | **string** | SQL-like  ORDER BY statement (column1 ASC|DESC , column2 ASC|DESC) | 
 **searchField** | **string** | Property name for searching | 
 **searchValue** | **string** | Value for searching | 

### Return type

[**[]ParameterGroupApiResult**](ParameterGroupApiResult.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2ParametersKeyNameNameExistsGet

> bool ApiV2ParametersKeyNameNameExistsGet(ctx, name).Execute()

Check existence parameter key in system



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
	resp, r, err := apiClient.ParametersAPI.ApiV2ParametersKeyNameNameExistsGet(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ParametersAPI.ApiV2ParametersKeyNameNameExistsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2ParametersKeyNameNameExistsGet`: bool
	fmt.Fprintf(os.Stdout, "Response from `ParametersAPI.ApiV2ParametersKeyNameNameExistsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ParametersKeyNameNameExistsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**bool**

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2ParametersKeyValuesGet

> []string ApiV2ParametersKeyValuesGet(ctx, key).Execute()

Get all parameter key values



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
	key := "SomeKey" // string | Parameter key (string format)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ParametersAPI.ApiV2ParametersKeyValuesGet(context.Background(), key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ParametersAPI.ApiV2ParametersKeyValuesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2ParametersKeyValuesGet`: []string
	fmt.Fprintf(os.Stdout, "Response from `ParametersAPI.ApiV2ParametersKeyValuesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | Parameter key (string format) | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ParametersKeyValuesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**[]string**

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2ParametersKeysGet

> []string ApiV2ParametersKeysGet(ctx).ProjectIds(projectIds).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).Execute()

Get all parameter keys



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
	projectIds := []*string{"Inner_example"} // []*string |  (optional)
	skip := int32(56) // int32 | Amount of items to be skipped (offset) (optional)
	take := int32(56) // int32 | Amount of items to be taken (limit) (optional)
	orderBy := "orderBy_example" // string | SQL-like  ORDER BY statement (column1 ASC|DESC , column2 ASC|DESC) (optional)
	searchField := "searchField_example" // string | Property name for searching (optional)
	searchValue := "searchValue_example" // string | Value for searching (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ParametersAPI.ApiV2ParametersKeysGet(context.Background()).ProjectIds(projectIds).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ParametersAPI.ApiV2ParametersKeysGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2ParametersKeysGet`: []string
	fmt.Fprintf(os.Stdout, "Response from `ParametersAPI.ApiV2ParametersKeysGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ParametersKeysGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectIds** | **[]string** |  | 
 **skip** | **int32** | Amount of items to be skipped (offset) | 
 **take** | **int32** | Amount of items to be taken (limit) | 
 **orderBy** | **string** | SQL-like  ORDER BY statement (column1 ASC|DESC , column2 ASC|DESC) | 
 **searchField** | **string** | Property name for searching | 
 **searchValue** | **string** | Value for searching | 

### Return type

**[]string**

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2ParametersSearchGroupsPost

> []ParameterGroupApiResult ApiV2ParametersSearchGroupsPost(ctx).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).ParameterGroupsFilterApiModel(parameterGroupsFilterApiModel).Execute()

Search for parameters as group

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
	parameterGroupsFilterApiModel := *openapiclient.NewParameterGroupsFilterApiModel() // ParameterGroupsFilterApiModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ParametersAPI.ApiV2ParametersSearchGroupsPost(context.Background()).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).ParameterGroupsFilterApiModel(parameterGroupsFilterApiModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ParametersAPI.ApiV2ParametersSearchGroupsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2ParametersSearchGroupsPost`: []ParameterGroupApiResult
	fmt.Fprintf(os.Stdout, "Response from `ParametersAPI.ApiV2ParametersSearchGroupsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ParametersSearchGroupsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skip** | **int32** | Amount of items to be skipped (offset) | 
 **take** | **int32** | Amount of items to be taken (limit) | 
 **orderBy** | **string** | SQL-like  ORDER BY statement (column1 ASC|DESC , column2 ASC|DESC) | 
 **searchField** | **string** | Property name for searching | 
 **searchValue** | **string** | Value for searching | 
 **parameterGroupsFilterApiModel** | [**ParameterGroupsFilterApiModel**](ParameterGroupsFilterApiModel.md) |  | 

### Return type

[**[]ParameterGroupApiResult**](ParameterGroupApiResult.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2ParametersSearchPost

> []ParameterApiResult ApiV2ParametersSearchPost(ctx).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).ParametersFilterApiModel(parametersFilterApiModel).Execute()

Search for parameters

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
	parametersFilterApiModel := *openapiclient.NewParametersFilterApiModel() // ParametersFilterApiModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ParametersAPI.ApiV2ParametersSearchPost(context.Background()).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).ParametersFilterApiModel(parametersFilterApiModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ParametersAPI.ApiV2ParametersSearchPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2ParametersSearchPost`: []ParameterApiResult
	fmt.Fprintf(os.Stdout, "Response from `ParametersAPI.ApiV2ParametersSearchPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ParametersSearchPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skip** | **int32** | Amount of items to be skipped (offset) | 
 **take** | **int32** | Amount of items to be taken (limit) | 
 **orderBy** | **string** | SQL-like  ORDER BY statement (column1 ASC|DESC , column2 ASC|DESC) | 
 **searchField** | **string** | Property name for searching | 
 **searchValue** | **string** | Value for searching | 
 **parametersFilterApiModel** | [**ParametersFilterApiModel**](ParametersFilterApiModel.md) |  | 

### Return type

[**[]ParameterApiResult**](ParameterApiResult.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateParameter

> ParameterApiResult CreateParameter(ctx).CreateParameterApiModel(createParameterApiModel).Execute()

Create parameter



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
	createParameterApiModel := *openapiclient.NewCreateParameterApiModel("Name_example", "Value_example") // CreateParameterApiModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ParametersAPI.CreateParameter(context.Background()).CreateParameterApiModel(createParameterApiModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ParametersAPI.CreateParameter``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateParameter`: ParameterApiResult
	fmt.Fprintf(os.Stdout, "Response from `ParametersAPI.CreateParameter`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateParameterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createParameterApiModel** | [**CreateParameterApiModel**](CreateParameterApiModel.md) |  | 

### Return type

[**ParameterApiResult**](ParameterApiResult.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteByName

> DeleteByName(ctx, name).Execute()

Delete parameter by name



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
	name := "name_example" // string | Name of the parameter

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ParametersAPI.DeleteByName(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ParametersAPI.DeleteByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the parameter | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteByNameRequest struct via the builder pattern


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


## DeleteByParameterKeyId

> DeleteByParameterKeyId(ctx, keyId).Execute()

Delete parameters by parameter key identifier



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
	keyId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Identifier of the parameter key

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ParametersAPI.DeleteByParameterKeyId(context.Background(), keyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ParametersAPI.DeleteByParameterKeyId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyId** | **string** | Identifier of the parameter key | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteByParameterKeyIdRequest struct via the builder pattern


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


## DeleteParameter

> DeleteParameter(ctx, id).Execute()

Delete parameter



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Parameter internal (UUID) identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ParametersAPI.DeleteParameter(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ParametersAPI.DeleteParameter``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Parameter internal (UUID) identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteParameterRequest struct via the builder pattern


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


## GetAllParameters

> []ParameterApiResult GetAllParameters(ctx).IsDeleted(isDeleted).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).Execute()

Get all parameters



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
	isDeleted := true // bool | If result must consist of only actual/deleted parameters (optional)
	skip := int32(56) // int32 | Amount of items to be skipped (offset) (optional)
	take := int32(56) // int32 | Amount of items to be taken (limit) (optional)
	orderBy := "orderBy_example" // string | SQL-like  ORDER BY statement (column1 ASC|DESC , column2 ASC|DESC) (optional)
	searchField := "searchField_example" // string | Property name for searching (optional)
	searchValue := "searchValue_example" // string | Value for searching (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ParametersAPI.GetAllParameters(context.Background()).IsDeleted(isDeleted).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ParametersAPI.GetAllParameters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllParameters`: []ParameterApiResult
	fmt.Fprintf(os.Stdout, "Response from `ParametersAPI.GetAllParameters`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllParametersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **isDeleted** | **bool** | If result must consist of only actual/deleted parameters | 
 **skip** | **int32** | Amount of items to be skipped (offset) | 
 **take** | **int32** | Amount of items to be taken (limit) | 
 **orderBy** | **string** | SQL-like  ORDER BY statement (column1 ASC|DESC , column2 ASC|DESC) | 
 **searchField** | **string** | Property name for searching | 
 **searchValue** | **string** | Value for searching | 

### Return type

[**[]ParameterApiResult**](ParameterApiResult.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetParameterById

> ParameterApiResult GetParameterById(ctx, id).Execute()

Get parameter by ID



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Parameter internal (UUID) identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ParametersAPI.GetParameterById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ParametersAPI.GetParameterById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetParameterById`: ParameterApiResult
	fmt.Fprintf(os.Stdout, "Response from `ParametersAPI.GetParameterById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Parameter internal (UUID) identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetParameterByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ParameterApiResult**](ParameterApiResult.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateParameter

> UpdateParameter(ctx).UpdateParameterApiModel(updateParameterApiModel).Execute()

Update parameter



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
	updateParameterApiModel := *openapiclient.NewUpdateParameterApiModel("Id_example", "Name_example", "Value_example") // UpdateParameterApiModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ParametersAPI.UpdateParameter(context.Background()).UpdateParameterApiModel(updateParameterApiModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ParametersAPI.UpdateParameter``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateParameterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateParameterApiModel** | [**UpdateParameterApiModel**](UpdateParameterApiModel.md) |  | 

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

