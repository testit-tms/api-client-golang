# \ExternalIssuesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV2ExternalIssuesSuggestionsPost**](ExternalIssuesAPI.md#ApiV2ExternalIssuesSuggestionsPost) | **Post** /api/v2/external-issues/suggestions | Returns list of suggestions from available external issues



## ApiV2ExternalIssuesSuggestionsPost

> ExternalIssueApiFieldSuggestionReply ApiV2ExternalIssuesSuggestionsPost(ctx).GetExternalIssueSuggestionsApiModel(getExternalIssueSuggestionsApiModel).Execute()

Returns list of suggestions from available external issues

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
	getExternalIssueSuggestionsApiModel := *openapiclient.NewGetExternalIssueSuggestionsApiModel(openapiclient.ExternalIssueApiField("Id")) // GetExternalIssueSuggestionsApiModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExternalIssuesAPI.ApiV2ExternalIssuesSuggestionsPost(context.Background()).GetExternalIssueSuggestionsApiModel(getExternalIssueSuggestionsApiModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExternalIssuesAPI.ApiV2ExternalIssuesSuggestionsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2ExternalIssuesSuggestionsPost`: ExternalIssueApiFieldSuggestionReply
	fmt.Fprintf(os.Stdout, "Response from `ExternalIssuesAPI.ApiV2ExternalIssuesSuggestionsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ExternalIssuesSuggestionsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getExternalIssueSuggestionsApiModel** | [**GetExternalIssueSuggestionsApiModel**](GetExternalIssueSuggestionsApiModel.md) |  | 

### Return type

[**ExternalIssueApiFieldSuggestionReply**](ExternalIssueApiFieldSuggestionReply.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

