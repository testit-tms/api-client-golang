# \UserRoleAssignmentsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV2UsersUserIdRolesRoleIdDelete**](UserRoleAssignmentsAPI.md#ApiV2UsersUserIdRolesRoleIdDelete) | **Delete** /api/v2/users/{userId}/roles/{roleId} | 
[**ApiV2UsersUserIdRolesRoleIdPost**](UserRoleAssignmentsAPI.md#ApiV2UsersUserIdRolesRoleIdPost) | **Post** /api/v2/users/{userId}/roles/{roleId} | 



## ApiV2UsersUserIdRolesRoleIdDelete

> ApiV2UsersUserIdRolesRoleIdDelete(ctx, userId, roleId).Execute()



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
	userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	roleId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UserRoleAssignmentsAPI.ApiV2UsersUserIdRolesRoleIdDelete(context.Background(), userId, roleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserRoleAssignmentsAPI.ApiV2UsersUserIdRolesRoleIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 
**roleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2UsersUserIdRolesRoleIdDeleteRequest struct via the builder pattern


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


## ApiV2UsersUserIdRolesRoleIdPost

> ApiV2UsersUserIdRolesRoleIdPost(ctx, userId, roleId).Execute()



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
	userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	roleId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UserRoleAssignmentsAPI.ApiV2UsersUserIdRolesRoleIdPost(context.Background(), userId, roleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserRoleAssignmentsAPI.ApiV2UsersUserIdRolesRoleIdPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 
**roleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2UsersUserIdRolesRoleIdPostRequest struct via the builder pattern


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

