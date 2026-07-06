# \FailureCategoriesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV2AutotestsFailureCategoriesGroupingSearchPost**](FailureCategoriesAPI.md#ApiV2AutotestsFailureCategoriesGroupingSearchPost) | **Post** /api/v2/autotests/failure-categories/grouping-search | Get failure categories with support for filtering, sorting and grouping
[**ApiV2AutotestsFailureCategoriesIdDelete**](FailureCategoriesAPI.md#ApiV2AutotestsFailureCategoriesIdDelete) | **Delete** /api/v2/autotests/failure-categories/{id} | Delete failure category
[**ApiV2AutotestsFailureCategoriesIdGet**](FailureCategoriesAPI.md#ApiV2AutotestsFailureCategoriesIdGet) | **Get** /api/v2/autotests/failure-categories/{id} | Get failure category by ID
[**ApiV2AutotestsFailureCategoriesNameNameExistsGet**](FailureCategoriesAPI.md#ApiV2AutotestsFailureCategoriesNameNameExistsGet) | **Get** /api/v2/autotests/failure-categories/name/{name}/exists | Check failure category with the specified name already exists
[**ApiV2AutotestsFailureCategoriesPost**](FailureCategoriesAPI.md#ApiV2AutotestsFailureCategoriesPost) | **Post** /api/v2/autotests/failure-categories | Create failure category
[**ApiV2AutotestsFailureCategoriesPut**](FailureCategoriesAPI.md#ApiV2AutotestsFailureCategoriesPut) | **Put** /api/v2/autotests/failure-categories | Update failure category
[**ApiV2AutotestsFailureCategoriesSearchPost**](FailureCategoriesAPI.md#ApiV2AutotestsFailureCategoriesSearchPost) | **Post** /api/v2/autotests/failure-categories/search | 
[**ApiV2AutotestsResultReasonsGroupingSearchPost**](FailureCategoriesAPI.md#ApiV2AutotestsResultReasonsGroupingSearchPost) | **Post** /api/v2/autotests/resultReasons/grouping-search | Get failure categories with support for filtering, sorting and grouping
[**ApiV2AutotestsResultReasonsIdDelete**](FailureCategoriesAPI.md#ApiV2AutotestsResultReasonsIdDelete) | **Delete** /api/v2/autotests/resultReasons/{id} | Delete failure category
[**ApiV2AutotestsResultReasonsIdGet**](FailureCategoriesAPI.md#ApiV2AutotestsResultReasonsIdGet) | **Get** /api/v2/autotests/resultReasons/{id} | Get failure category by ID
[**ApiV2AutotestsResultReasonsNameNameExistsGet**](FailureCategoriesAPI.md#ApiV2AutotestsResultReasonsNameNameExistsGet) | **Get** /api/v2/autotests/resultReasons/name/{name}/exists | Check failure category with the specified name already exists
[**ApiV2AutotestsResultReasonsPost**](FailureCategoriesAPI.md#ApiV2AutotestsResultReasonsPost) | **Post** /api/v2/autotests/resultReasons | Create failure category
[**ApiV2AutotestsResultReasonsPut**](FailureCategoriesAPI.md#ApiV2AutotestsResultReasonsPut) | **Put** /api/v2/autotests/resultReasons | Update failure category
[**ApiV2AutotestsResultReasonsSearchPost**](FailureCategoriesAPI.md#ApiV2AutotestsResultReasonsSearchPost) | **Post** /api/v2/autotests/resultReasons/search | 



## ApiV2AutotestsFailureCategoriesGroupingSearchPost

> FailureCategoryGroupItemApiResultReply ApiV2AutotestsFailureCategoriesGroupingSearchPost(ctx).FailureCategoryGroupSearchApiModel(failureCategoryGroupSearchApiModel).Execute()

Get failure categories with support for filtering, sorting and grouping

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
	failureCategoryGroupSearchApiModel := *openapiclient.NewFailureCategoryGroupSearchApiModel(*openapiclient.NewInquiry([]openapiclient.Order{*openapiclient.NewOrder("Field_example", openapiclient.ListSortDirection("Ascending"))})) // FailureCategoryGroupSearchApiModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FailureCategoriesAPI.ApiV2AutotestsFailureCategoriesGroupingSearchPost(context.Background()).FailureCategoryGroupSearchApiModel(failureCategoryGroupSearchApiModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FailureCategoriesAPI.ApiV2AutotestsFailureCategoriesGroupingSearchPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2AutotestsFailureCategoriesGroupingSearchPost`: FailureCategoryGroupItemApiResultReply
	fmt.Fprintf(os.Stdout, "Response from `FailureCategoriesAPI.ApiV2AutotestsFailureCategoriesGroupingSearchPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2AutotestsFailureCategoriesGroupingSearchPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **failureCategoryGroupSearchApiModel** | [**FailureCategoryGroupSearchApiModel**](FailureCategoryGroupSearchApiModel.md) |  | 

### Return type

[**FailureCategoryGroupItemApiResultReply**](FailureCategoryGroupItemApiResultReply.md)

### Authorization

[PrivateToken](../README.md#PrivateToken), [Cookies](../README.md#Cookies)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2AutotestsFailureCategoriesIdDelete

> ApiV2AutotestsFailureCategoriesIdDelete(ctx, id).Execute()

Delete failure category

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
	r, err := apiClient.FailureCategoriesAPI.ApiV2AutotestsFailureCategoriesIdDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FailureCategoriesAPI.ApiV2AutotestsFailureCategoriesIdDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiApiV2AutotestsFailureCategoriesIdDeleteRequest struct via the builder pattern


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


## ApiV2AutotestsFailureCategoriesIdGet

> FailureCategoryApiResult ApiV2AutotestsFailureCategoriesIdGet(ctx, id).IsDeleted(isDeleted).Execute()

Get failure category by ID

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
	isDeleted := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FailureCategoriesAPI.ApiV2AutotestsFailureCategoriesIdGet(context.Background(), id).IsDeleted(isDeleted).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FailureCategoriesAPI.ApiV2AutotestsFailureCategoriesIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2AutotestsFailureCategoriesIdGet`: FailureCategoryApiResult
	fmt.Fprintf(os.Stdout, "Response from `FailureCategoriesAPI.ApiV2AutotestsFailureCategoriesIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2AutotestsFailureCategoriesIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **isDeleted** | **bool** |  | 

### Return type

[**FailureCategoryApiResult**](FailureCategoryApiResult.md)

### Authorization

[PrivateToken](../README.md#PrivateToken), [Cookies](../README.md#Cookies)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2AutotestsFailureCategoriesNameNameExistsGet

> bool ApiV2AutotestsFailureCategoriesNameNameExistsGet(ctx, name).Execute()

Check failure category with the specified name already exists

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
	resp, r, err := apiClient.FailureCategoriesAPI.ApiV2AutotestsFailureCategoriesNameNameExistsGet(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FailureCategoriesAPI.ApiV2AutotestsFailureCategoriesNameNameExistsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2AutotestsFailureCategoriesNameNameExistsGet`: bool
	fmt.Fprintf(os.Stdout, "Response from `FailureCategoriesAPI.ApiV2AutotestsFailureCategoriesNameNameExistsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2AutotestsFailureCategoriesNameNameExistsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**bool**

### Authorization

[PrivateToken](../README.md#PrivateToken), [Cookies](../README.md#Cookies)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2AutotestsFailureCategoriesPost

> FailureCategoryApiResult ApiV2AutotestsFailureCategoriesPost(ctx).CreateFailureCategoryApiModel(createFailureCategoryApiModel).Execute()

Create failure category

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
	createFailureCategoryApiModel := *openapiclient.NewCreateFailureCategoryApiModel("Name_example", openapiclient.FailureCategory("InfrastructureDefect")) // CreateFailureCategoryApiModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FailureCategoriesAPI.ApiV2AutotestsFailureCategoriesPost(context.Background()).CreateFailureCategoryApiModel(createFailureCategoryApiModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FailureCategoriesAPI.ApiV2AutotestsFailureCategoriesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2AutotestsFailureCategoriesPost`: FailureCategoryApiResult
	fmt.Fprintf(os.Stdout, "Response from `FailureCategoriesAPI.ApiV2AutotestsFailureCategoriesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2AutotestsFailureCategoriesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createFailureCategoryApiModel** | [**CreateFailureCategoryApiModel**](CreateFailureCategoryApiModel.md) |  | 

### Return type

[**FailureCategoryApiResult**](FailureCategoryApiResult.md)

### Authorization

[PrivateToken](../README.md#PrivateToken), [Cookies](../README.md#Cookies)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2AutotestsFailureCategoriesPut

> ApiV2AutotestsFailureCategoriesPut(ctx).UpdateFailureCategoryApiModel(updateFailureCategoryApiModel).Execute()

Update failure category

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
	updateFailureCategoryApiModel := *openapiclient.NewUpdateFailureCategoryApiModel("Id_example", "Name_example", openapiclient.FailureCategory("InfrastructureDefect")) // UpdateFailureCategoryApiModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FailureCategoriesAPI.ApiV2AutotestsFailureCategoriesPut(context.Background()).UpdateFailureCategoryApiModel(updateFailureCategoryApiModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FailureCategoriesAPI.ApiV2AutotestsFailureCategoriesPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2AutotestsFailureCategoriesPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateFailureCategoryApiModel** | [**UpdateFailureCategoryApiModel**](UpdateFailureCategoryApiModel.md) |  | 

### Return type

 (empty response body)

### Authorization

[PrivateToken](../README.md#PrivateToken), [Cookies](../README.md#Cookies)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2AutotestsFailureCategoriesSearchPost

> []AutotestResultReasonShortGetModel ApiV2AutotestsFailureCategoriesSearchPost(ctx).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).AutotestResultReasonFilterModel(autotestResultReasonFilterModel).Execute()



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
	autotestResultReasonFilterModel := *openapiclient.NewAutotestResultReasonFilterModel() // AutotestResultReasonFilterModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FailureCategoriesAPI.ApiV2AutotestsFailureCategoriesSearchPost(context.Background()).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).AutotestResultReasonFilterModel(autotestResultReasonFilterModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FailureCategoriesAPI.ApiV2AutotestsFailureCategoriesSearchPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2AutotestsFailureCategoriesSearchPost`: []AutotestResultReasonShortGetModel
	fmt.Fprintf(os.Stdout, "Response from `FailureCategoriesAPI.ApiV2AutotestsFailureCategoriesSearchPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2AutotestsFailureCategoriesSearchPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skip** | **int32** | Amount of items to be skipped (offset) | 
 **take** | **int32** | Amount of items to be taken (limit) | 
 **orderBy** | **string** | SQL-like  ORDER BY statement (column1 ASC|DESC , column2 ASC|DESC) | 
 **searchField** | **string** | Property name for searching | 
 **searchValue** | **string** | Value for searching | 
 **autotestResultReasonFilterModel** | [**AutotestResultReasonFilterModel**](AutotestResultReasonFilterModel.md) |  | 

### Return type

[**[]AutotestResultReasonShortGetModel**](AutotestResultReasonShortGetModel.md)

### Authorization

[PrivateToken](../README.md#PrivateToken), [Cookies](../README.md#Cookies)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2AutotestsResultReasonsGroupingSearchPost

> FailureCategoryGroupItemApiResultReply ApiV2AutotestsResultReasonsGroupingSearchPost(ctx).FailureCategoryGroupSearchApiModel(failureCategoryGroupSearchApiModel).Execute()

Get failure categories with support for filtering, sorting and grouping

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
	failureCategoryGroupSearchApiModel := *openapiclient.NewFailureCategoryGroupSearchApiModel(*openapiclient.NewInquiry([]openapiclient.Order{*openapiclient.NewOrder("Field_example", openapiclient.ListSortDirection("Ascending"))})) // FailureCategoryGroupSearchApiModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FailureCategoriesAPI.ApiV2AutotestsResultReasonsGroupingSearchPost(context.Background()).FailureCategoryGroupSearchApiModel(failureCategoryGroupSearchApiModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FailureCategoriesAPI.ApiV2AutotestsResultReasonsGroupingSearchPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2AutotestsResultReasonsGroupingSearchPost`: FailureCategoryGroupItemApiResultReply
	fmt.Fprintf(os.Stdout, "Response from `FailureCategoriesAPI.ApiV2AutotestsResultReasonsGroupingSearchPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2AutotestsResultReasonsGroupingSearchPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **failureCategoryGroupSearchApiModel** | [**FailureCategoryGroupSearchApiModel**](FailureCategoryGroupSearchApiModel.md) |  | 

### Return type

[**FailureCategoryGroupItemApiResultReply**](FailureCategoryGroupItemApiResultReply.md)

### Authorization

[PrivateToken](../README.md#PrivateToken), [Cookies](../README.md#Cookies)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2AutotestsResultReasonsIdDelete

> ApiV2AutotestsResultReasonsIdDelete(ctx, id).Execute()

Delete failure category

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
	r, err := apiClient.FailureCategoriesAPI.ApiV2AutotestsResultReasonsIdDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FailureCategoriesAPI.ApiV2AutotestsResultReasonsIdDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiApiV2AutotestsResultReasonsIdDeleteRequest struct via the builder pattern


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


## ApiV2AutotestsResultReasonsIdGet

> FailureCategoryApiResult ApiV2AutotestsResultReasonsIdGet(ctx, id).IsDeleted(isDeleted).Execute()

Get failure category by ID

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
	isDeleted := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FailureCategoriesAPI.ApiV2AutotestsResultReasonsIdGet(context.Background(), id).IsDeleted(isDeleted).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FailureCategoriesAPI.ApiV2AutotestsResultReasonsIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2AutotestsResultReasonsIdGet`: FailureCategoryApiResult
	fmt.Fprintf(os.Stdout, "Response from `FailureCategoriesAPI.ApiV2AutotestsResultReasonsIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2AutotestsResultReasonsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **isDeleted** | **bool** |  | 

### Return type

[**FailureCategoryApiResult**](FailureCategoryApiResult.md)

### Authorization

[PrivateToken](../README.md#PrivateToken), [Cookies](../README.md#Cookies)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2AutotestsResultReasonsNameNameExistsGet

> bool ApiV2AutotestsResultReasonsNameNameExistsGet(ctx, name).Execute()

Check failure category with the specified name already exists

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
	resp, r, err := apiClient.FailureCategoriesAPI.ApiV2AutotestsResultReasonsNameNameExistsGet(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FailureCategoriesAPI.ApiV2AutotestsResultReasonsNameNameExistsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2AutotestsResultReasonsNameNameExistsGet`: bool
	fmt.Fprintf(os.Stdout, "Response from `FailureCategoriesAPI.ApiV2AutotestsResultReasonsNameNameExistsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2AutotestsResultReasonsNameNameExistsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**bool**

### Authorization

[PrivateToken](../README.md#PrivateToken), [Cookies](../README.md#Cookies)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2AutotestsResultReasonsPost

> FailureCategoryApiResult ApiV2AutotestsResultReasonsPost(ctx).CreateFailureCategoryApiModel(createFailureCategoryApiModel).Execute()

Create failure category

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
	createFailureCategoryApiModel := *openapiclient.NewCreateFailureCategoryApiModel("Name_example", openapiclient.FailureCategory("InfrastructureDefect")) // CreateFailureCategoryApiModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FailureCategoriesAPI.ApiV2AutotestsResultReasonsPost(context.Background()).CreateFailureCategoryApiModel(createFailureCategoryApiModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FailureCategoriesAPI.ApiV2AutotestsResultReasonsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2AutotestsResultReasonsPost`: FailureCategoryApiResult
	fmt.Fprintf(os.Stdout, "Response from `FailureCategoriesAPI.ApiV2AutotestsResultReasonsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2AutotestsResultReasonsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createFailureCategoryApiModel** | [**CreateFailureCategoryApiModel**](CreateFailureCategoryApiModel.md) |  | 

### Return type

[**FailureCategoryApiResult**](FailureCategoryApiResult.md)

### Authorization

[PrivateToken](../README.md#PrivateToken), [Cookies](../README.md#Cookies)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2AutotestsResultReasonsPut

> ApiV2AutotestsResultReasonsPut(ctx).UpdateFailureCategoryApiModel(updateFailureCategoryApiModel).Execute()

Update failure category

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
	updateFailureCategoryApiModel := *openapiclient.NewUpdateFailureCategoryApiModel("Id_example", "Name_example", openapiclient.FailureCategory("InfrastructureDefect")) // UpdateFailureCategoryApiModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FailureCategoriesAPI.ApiV2AutotestsResultReasonsPut(context.Background()).UpdateFailureCategoryApiModel(updateFailureCategoryApiModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FailureCategoriesAPI.ApiV2AutotestsResultReasonsPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2AutotestsResultReasonsPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateFailureCategoryApiModel** | [**UpdateFailureCategoryApiModel**](UpdateFailureCategoryApiModel.md) |  | 

### Return type

 (empty response body)

### Authorization

[PrivateToken](../README.md#PrivateToken), [Cookies](../README.md#Cookies)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2AutotestsResultReasonsSearchPost

> []AutotestResultReasonShortGetModel ApiV2AutotestsResultReasonsSearchPost(ctx).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).AutotestResultReasonFilterModel(autotestResultReasonFilterModel).Execute()



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
	autotestResultReasonFilterModel := *openapiclient.NewAutotestResultReasonFilterModel() // AutotestResultReasonFilterModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FailureCategoriesAPI.ApiV2AutotestsResultReasonsSearchPost(context.Background()).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).AutotestResultReasonFilterModel(autotestResultReasonFilterModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FailureCategoriesAPI.ApiV2AutotestsResultReasonsSearchPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2AutotestsResultReasonsSearchPost`: []AutotestResultReasonShortGetModel
	fmt.Fprintf(os.Stdout, "Response from `FailureCategoriesAPI.ApiV2AutotestsResultReasonsSearchPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2AutotestsResultReasonsSearchPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skip** | **int32** | Amount of items to be skipped (offset) | 
 **take** | **int32** | Amount of items to be taken (limit) | 
 **orderBy** | **string** | SQL-like  ORDER BY statement (column1 ASC|DESC , column2 ASC|DESC) | 
 **searchField** | **string** | Property name for searching | 
 **searchValue** | **string** | Value for searching | 
 **autotestResultReasonFilterModel** | [**AutotestResultReasonFilterModel**](AutotestResultReasonFilterModel.md) |  | 

### Return type

[**[]AutotestResultReasonShortGetModel**](AutotestResultReasonShortGetModel.md)

### Authorization

[PrivateToken](../README.md#PrivateToken), [Cookies](../README.md#Cookies)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

