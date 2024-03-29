# \TestSuitesApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddTestPointsToTestSuite**](TestSuitesApi.md#AddTestPointsToTestSuite) | **Post** /api/v2/testSuites/{id}/test-points | Add test-points to test suite
[**ApiV2TestSuitesIdPatch**](TestSuitesApi.md#ApiV2TestSuitesIdPatch) | **Patch** /api/v2/testSuites/{id} | Patch test suite
[**ApiV2TestSuitesIdRefreshPost**](TestSuitesApi.md#ApiV2TestSuitesIdRefreshPost) | **Post** /api/v2/testSuites/{id}/refresh | Refresh test suite. Only dynamic test suites are supported by this method
[**ApiV2TestSuitesIdWorkItemsPost**](TestSuitesApi.md#ApiV2TestSuitesIdWorkItemsPost) | **Post** /api/v2/testSuites/{id}/workItems | Set work items for test suite
[**ApiV2TestSuitesPost**](TestSuitesApi.md#ApiV2TestSuitesPost) | **Post** /api/v2/testSuites | Create test suite
[**ApiV2TestSuitesPut**](TestSuitesApi.md#ApiV2TestSuitesPut) | **Put** /api/v2/testSuites | Edit test suite
[**DeleteTestSuite**](TestSuitesApi.md#DeleteTestSuite) | **Delete** /api/v2/testSuites/{id} | Delete TestSuite
[**GetConfigurationsByTestSuiteId**](TestSuitesApi.md#GetConfigurationsByTestSuiteId) | **Get** /api/v2/testSuites/{id}/configurations | Get Configurations By Id
[**GetTestPointsById**](TestSuitesApi.md#GetTestPointsById) | **Get** /api/v2/testSuites/{id}/testPoints | Get TestPoints By Id
[**GetTestResultsById**](TestSuitesApi.md#GetTestResultsById) | **Get** /api/v2/testSuites/{id}/testResults | Get TestResults By Id
[**GetTestSuiteById**](TestSuitesApi.md#GetTestSuiteById) | **Get** /api/v2/testSuites/{id} | Get TestSuite by Id
[**SearchWorkItems**](TestSuitesApi.md#SearchWorkItems) | **Post** /api/v2/testSuites/{id}/workItems/search | Search WorkItems
[**SetConfigurationsByTestSuiteId**](TestSuitesApi.md#SetConfigurationsByTestSuiteId) | **Post** /api/v2/testSuites/{id}/configurations | Set Configurations By TestSuite Id



## AddTestPointsToTestSuite

> AddTestPointsToTestSuite(ctx, id).ApiV2ProjectsProjectIdWorkItemsSearchPostRequest(apiV2ProjectsProjectIdWorkItemsSearchPostRequest).Execute()

Add test-points to test suite

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
    id := "1ed608bf-8ac9-4ffd-b91e-ebdbbdce6132" // string | Test suite internal identifier
    apiV2ProjectsProjectIdWorkItemsSearchPostRequest := *openapiclient.NewApiV2ProjectsProjectIdWorkItemsSearchPostRequest(*openapiclient.NewWorkItemSelectModelFilter()) // ApiV2ProjectsProjectIdWorkItemsSearchPostRequest | Filter object to retrieve work items for test-suite's project (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TestSuitesApi.AddTestPointsToTestSuite(context.Background(), id).ApiV2ProjectsProjectIdWorkItemsSearchPostRequest(apiV2ProjectsProjectIdWorkItemsSearchPostRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TestSuitesApi.AddTestPointsToTestSuite``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Test suite internal identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddTestPointsToTestSuiteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiV2ProjectsProjectIdWorkItemsSearchPostRequest** | [**ApiV2ProjectsProjectIdWorkItemsSearchPostRequest**](ApiV2ProjectsProjectIdWorkItemsSearchPostRequest.md) | Filter object to retrieve work items for test-suite&#39;s project | 

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


## ApiV2TestSuitesIdPatch

> ApiV2TestSuitesIdPatch(ctx, id).Operation(operation).Execute()

Patch test suite



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Test Suite internal (UUID) identifier
    operation := []openapiclient.Operation{*openapiclient.NewOperation()} // []Operation |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TestSuitesApi.ApiV2TestSuitesIdPatch(context.Background(), id).Operation(operation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TestSuitesApi.ApiV2TestSuitesIdPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Test Suite internal (UUID) identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2TestSuitesIdPatchRequest struct via the builder pattern


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


## ApiV2TestSuitesIdRefreshPost

> ApiV2TestSuitesIdRefreshPost(ctx, id).Execute()

Refresh test suite. Only dynamic test suites are supported by this method

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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Test Suite internal (UUID) identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TestSuitesApi.ApiV2TestSuitesIdRefreshPost(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TestSuitesApi.ApiV2TestSuitesIdRefreshPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Test Suite internal (UUID) identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2TestSuitesIdRefreshPostRequest struct via the builder pattern


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


## ApiV2TestSuitesIdWorkItemsPost

> ApiV2TestSuitesIdWorkItemsPost(ctx, id).RequestBody(requestBody).Execute()

Set work items for test suite

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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique ID of the test suite
    requestBody := []string{"Property_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TestSuitesApi.ApiV2TestSuitesIdWorkItemsPost(context.Background(), id).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TestSuitesApi.ApiV2TestSuitesIdWorkItemsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Unique ID of the test suite | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2TestSuitesIdWorkItemsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestBody** | **[]string** |  | 

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


## ApiV2TestSuitesPost

> TestSuiteV2GetModel ApiV2TestSuitesPost(ctx).ApiV2TestSuitesPostRequest(apiV2TestSuitesPostRequest).Execute()

Create test suite

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
    apiV2TestSuitesPostRequest := *openapiclient.NewApiV2TestSuitesPostRequest("TestPlanId_example", "Name_example") // ApiV2TestSuitesPostRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TestSuitesApi.ApiV2TestSuitesPost(context.Background()).ApiV2TestSuitesPostRequest(apiV2TestSuitesPostRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TestSuitesApi.ApiV2TestSuitesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV2TestSuitesPost`: TestSuiteV2GetModel
    fmt.Fprintf(os.Stdout, "Response from `TestSuitesApi.ApiV2TestSuitesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2TestSuitesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiV2TestSuitesPostRequest** | [**ApiV2TestSuitesPostRequest**](ApiV2TestSuitesPostRequest.md) |  | 

### Return type

[**TestSuiteV2GetModel**](TestSuiteV2GetModel.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2TestSuitesPut

> ApiV2TestSuitesPut(ctx).ApiV2TestSuitesPutRequest(apiV2TestSuitesPutRequest).Execute()

Edit test suite

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
    apiV2TestSuitesPutRequest := *openapiclient.NewApiV2TestSuitesPutRequest("Id_example", "Name_example", false) // ApiV2TestSuitesPutRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TestSuitesApi.ApiV2TestSuitesPut(context.Background()).ApiV2TestSuitesPutRequest(apiV2TestSuitesPutRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TestSuitesApi.ApiV2TestSuitesPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2TestSuitesPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiV2TestSuitesPutRequest** | [**ApiV2TestSuitesPutRequest**](ApiV2TestSuitesPutRequest.md) |  | 

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


## DeleteTestSuite

> DeleteTestSuite(ctx, id).Execute()

Delete TestSuite



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
    id := "3fa85f64-5717-4562-b3fc-2c963f66afa6" // string | Test suite internal (guid format) identifier\"

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TestSuitesApi.DeleteTestSuite(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TestSuitesApi.DeleteTestSuite``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Test suite internal (guid format) identifier\&quot; | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTestSuiteRequest struct via the builder pattern


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


## GetConfigurationsByTestSuiteId

> []ConfigurationModel GetConfigurationsByTestSuiteId(ctx, id).Execute()

Get Configurations By Id



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
    id := "3fa85f64-5717-4562-b3fc-2c963f66afa6" // string | Test suite internal (guid format) identifier\"

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TestSuitesApi.GetConfigurationsByTestSuiteId(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TestSuitesApi.GetConfigurationsByTestSuiteId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConfigurationsByTestSuiteId`: []ConfigurationModel
    fmt.Fprintf(os.Stdout, "Response from `TestSuitesApi.GetConfigurationsByTestSuiteId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Test suite internal (guid format) identifier\&quot; | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConfigurationsByTestSuiteIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ConfigurationModel**](ConfigurationModel.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTestPointsById

> []TestPointByTestSuiteModel GetTestPointsById(ctx, id).Execute()

Get TestPoints By Id



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
    id := "3fa85f64-5717-4562-b3fc-2c963f66afa6" // string | Test suite internal (guid format) identifier\"

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TestSuitesApi.GetTestPointsById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TestSuitesApi.GetTestPointsById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTestPointsById`: []TestPointByTestSuiteModel
    fmt.Fprintf(os.Stdout, "Response from `TestSuitesApi.GetTestPointsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Test suite internal (guid format) identifier\&quot; | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTestPointsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]TestPointByTestSuiteModel**](TestPointByTestSuiteModel.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTestResultsById

> []TestResultV2ShortModel GetTestResultsById(ctx, id).Execute()

Get TestResults By Id



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
    id := "3fa85f64-5717-4562-b3fc-2c963f66afa6" // string | Test suite internal (guid format) identifier\"

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TestSuitesApi.GetTestResultsById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TestSuitesApi.GetTestResultsById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTestResultsById`: []TestResultV2ShortModel
    fmt.Fprintf(os.Stdout, "Response from `TestSuitesApi.GetTestResultsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Test suite internal (guid format) identifier\&quot; | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTestResultsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]TestResultV2ShortModel**](TestResultV2ShortModel.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTestSuiteById

> TestSuiteV2GetModel GetTestSuiteById(ctx, id).Execute()

Get TestSuite by Id



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
    id := "3fa85f64-5717-4562-b3fc-2c963f66afa6" // string | Test suite internal (guid format) identifier\"

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TestSuitesApi.GetTestSuiteById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TestSuitesApi.GetTestSuiteById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTestSuiteById`: TestSuiteV2GetModel
    fmt.Fprintf(os.Stdout, "Response from `TestSuitesApi.GetTestSuiteById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Test suite internal (guid format) identifier\&quot; | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTestSuiteByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TestSuiteV2GetModel**](TestSuiteV2GetModel.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchWorkItems

> []WorkItemShortModel SearchWorkItems(ctx, id).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).SearchWorkItemsRequest(searchWorkItemsRequest).Execute()

Search WorkItems



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
    id := "3fa85f64-5717-4562-b3fc-2c963f66afa6" // string | Test suite internal (guid format) identifier\"
    skip := int32(56) // int32 | Amount of items to be skipped (offset) (optional)
    take := int32(56) // int32 | Amount of items to be taken (limit) (optional)
    orderBy := "orderBy_example" // string | SQL-like  ORDER BY statement (column1 ASC|DESC , column2 ASC|DESC) (optional)
    searchField := "searchField_example" // string | Property name for searching (optional)
    searchValue := "searchValue_example" // string | Value for searching (optional)
    searchWorkItemsRequest := *openapiclient.NewSearchWorkItemsRequest() // SearchWorkItemsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TestSuitesApi.SearchWorkItems(context.Background(), id).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).SearchWorkItemsRequest(searchWorkItemsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TestSuitesApi.SearchWorkItems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchWorkItems`: []WorkItemShortModel
    fmt.Fprintf(os.Stdout, "Response from `TestSuitesApi.SearchWorkItems`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Test suite internal (guid format) identifier\&quot; | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchWorkItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **skip** | **int32** | Amount of items to be skipped (offset) | 
 **take** | **int32** | Amount of items to be taken (limit) | 
 **orderBy** | **string** | SQL-like  ORDER BY statement (column1 ASC|DESC , column2 ASC|DESC) | 
 **searchField** | **string** | Property name for searching | 
 **searchValue** | **string** | Value for searching | 
 **searchWorkItemsRequest** | [**SearchWorkItemsRequest**](SearchWorkItemsRequest.md) |  | 

### Return type

[**[]WorkItemShortModel**](WorkItemShortModel.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetConfigurationsByTestSuiteId

> SetConfigurationsByTestSuiteId(ctx, id).RequestBody(requestBody).Execute()

Set Configurations By TestSuite Id



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
    id := "3fa85f64-5717-4562-b3fc-2c963f66afa6" // string | Test suite internal (guid format) identifier\"
    requestBody := []string{"Property_example"} // []string | Collection of configuration identifiers\" (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TestSuitesApi.SetConfigurationsByTestSuiteId(context.Background(), id).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TestSuitesApi.SetConfigurationsByTestSuiteId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Test suite internal (guid format) identifier\&quot; | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetConfigurationsByTestSuiteIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestBody** | **[]string** | Collection of configuration identifiers\&quot; | 

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

