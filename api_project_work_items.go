/*
API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tmsclient

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
	"reflect"
)


// ProjectWorkItemsApiService ProjectWorkItemsApi service
type ProjectWorkItemsApiService service

type ApiApiV2ProjectsProjectIdWorkItemsSearchGroupedPostRequest struct {
	ctx context.Context
	ApiService *ProjectWorkItemsApiService
	projectId string
	skip *int32
	take *int32
	orderBy *string
	searchField *string
	searchValue *string
	apiV2ProjectsProjectIdWorkItemsSearchGroupedPostRequest *ApiV2ProjectsProjectIdWorkItemsSearchGroupedPostRequest
}

// Amount of items to be skipped (offset)
func (r ApiApiV2ProjectsProjectIdWorkItemsSearchGroupedPostRequest) Skip(skip int32) ApiApiV2ProjectsProjectIdWorkItemsSearchGroupedPostRequest {
	r.skip = &skip
	return r
}

// Amount of items to be taken (limit)
func (r ApiApiV2ProjectsProjectIdWorkItemsSearchGroupedPostRequest) Take(take int32) ApiApiV2ProjectsProjectIdWorkItemsSearchGroupedPostRequest {
	r.take = &take
	return r
}

// SQL-like  ORDER BY statement (column1 ASC|DESC , column2 ASC|DESC)
func (r ApiApiV2ProjectsProjectIdWorkItemsSearchGroupedPostRequest) OrderBy(orderBy string) ApiApiV2ProjectsProjectIdWorkItemsSearchGroupedPostRequest {
	r.orderBy = &orderBy
	return r
}

// Property name for searching
func (r ApiApiV2ProjectsProjectIdWorkItemsSearchGroupedPostRequest) SearchField(searchField string) ApiApiV2ProjectsProjectIdWorkItemsSearchGroupedPostRequest {
	r.searchField = &searchField
	return r
}

// Value for searching
func (r ApiApiV2ProjectsProjectIdWorkItemsSearchGroupedPostRequest) SearchValue(searchValue string) ApiApiV2ProjectsProjectIdWorkItemsSearchGroupedPostRequest {
	r.searchValue = &searchValue
	return r
}

func (r ApiApiV2ProjectsProjectIdWorkItemsSearchGroupedPostRequest) ApiV2ProjectsProjectIdWorkItemsSearchGroupedPostRequest(apiV2ProjectsProjectIdWorkItemsSearchGroupedPostRequest ApiV2ProjectsProjectIdWorkItemsSearchGroupedPostRequest) ApiApiV2ProjectsProjectIdWorkItemsSearchGroupedPostRequest {
	r.apiV2ProjectsProjectIdWorkItemsSearchGroupedPostRequest = &apiV2ProjectsProjectIdWorkItemsSearchGroupedPostRequest
	return r
}

func (r ApiApiV2ProjectsProjectIdWorkItemsSearchGroupedPostRequest) Execute() ([]WorkItemGroupModel, *http.Response, error) {
	return r.ApiService.ApiV2ProjectsProjectIdWorkItemsSearchGroupedPostExecute(r)
}

/*
ApiV2ProjectsProjectIdWorkItemsSearchGroupedPost Search for work items and group results by attribute

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projectId Unique or global ID of the project
 @return ApiApiV2ProjectsProjectIdWorkItemsSearchGroupedPostRequest
*/
func (a *ProjectWorkItemsApiService) ApiV2ProjectsProjectIdWorkItemsSearchGroupedPost(ctx context.Context, projectId string) ApiApiV2ProjectsProjectIdWorkItemsSearchGroupedPostRequest {
	return ApiApiV2ProjectsProjectIdWorkItemsSearchGroupedPostRequest{
		ApiService: a,
		ctx: ctx,
		projectId: projectId,
	}
}

// Execute executes the request
//  @return []WorkItemGroupModel
func (a *ProjectWorkItemsApiService) ApiV2ProjectsProjectIdWorkItemsSearchGroupedPostExecute(r ApiApiV2ProjectsProjectIdWorkItemsSearchGroupedPostRequest) ([]WorkItemGroupModel, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []WorkItemGroupModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProjectWorkItemsApiService.ApiV2ProjectsProjectIdWorkItemsSearchGroupedPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/projects/{projectId}/workItems/search/grouped"
	localVarPath = strings.Replace(localVarPath, "{"+"projectId"+"}", url.PathEscape(parameterValueToString(r.projectId, "projectId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.skip != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "Skip", r.skip, "")
	}
	if r.take != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "Take", r.take, "")
	}
	if r.orderBy != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "OrderBy", r.orderBy, "")
	}
	if r.searchField != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "SearchField", r.searchField, "")
	}
	if r.searchValue != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "SearchValue", r.searchValue, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.apiV2ProjectsProjectIdWorkItemsSearchGroupedPostRequest
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["Bearer or PrivateToken"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiApiV2ProjectsProjectIdWorkItemsSearchIdPostRequest struct {
	ctx context.Context
	ApiService *ProjectWorkItemsApiService
	projectId string
	skip *int32
	take *int32
	orderBy *string
	searchField *string
	searchValue *string
	apiV2ProjectsProjectIdWorkItemsSearchPostRequest *ApiV2ProjectsProjectIdWorkItemsSearchPostRequest
}

// Amount of items to be skipped (offset)
func (r ApiApiV2ProjectsProjectIdWorkItemsSearchIdPostRequest) Skip(skip int32) ApiApiV2ProjectsProjectIdWorkItemsSearchIdPostRequest {
	r.skip = &skip
	return r
}

// Amount of items to be taken (limit)
func (r ApiApiV2ProjectsProjectIdWorkItemsSearchIdPostRequest) Take(take int32) ApiApiV2ProjectsProjectIdWorkItemsSearchIdPostRequest {
	r.take = &take
	return r
}

// SQL-like  ORDER BY statement (column1 ASC|DESC , column2 ASC|DESC)
func (r ApiApiV2ProjectsProjectIdWorkItemsSearchIdPostRequest) OrderBy(orderBy string) ApiApiV2ProjectsProjectIdWorkItemsSearchIdPostRequest {
	r.orderBy = &orderBy
	return r
}

// Property name for searching
func (r ApiApiV2ProjectsProjectIdWorkItemsSearchIdPostRequest) SearchField(searchField string) ApiApiV2ProjectsProjectIdWorkItemsSearchIdPostRequest {
	r.searchField = &searchField
	return r
}

// Value for searching
func (r ApiApiV2ProjectsProjectIdWorkItemsSearchIdPostRequest) SearchValue(searchValue string) ApiApiV2ProjectsProjectIdWorkItemsSearchIdPostRequest {
	r.searchValue = &searchValue
	return r
}

func (r ApiApiV2ProjectsProjectIdWorkItemsSearchIdPostRequest) ApiV2ProjectsProjectIdWorkItemsSearchPostRequest(apiV2ProjectsProjectIdWorkItemsSearchPostRequest ApiV2ProjectsProjectIdWorkItemsSearchPostRequest) ApiApiV2ProjectsProjectIdWorkItemsSearchIdPostRequest {
	r.apiV2ProjectsProjectIdWorkItemsSearchPostRequest = &apiV2ProjectsProjectIdWorkItemsSearchPostRequest
	return r
}

func (r ApiApiV2ProjectsProjectIdWorkItemsSearchIdPostRequest) Execute() ([]string, *http.Response, error) {
	return r.ApiService.ApiV2ProjectsProjectIdWorkItemsSearchIdPostExecute(r)
}

/*
ApiV2ProjectsProjectIdWorkItemsSearchIdPost Search for work items and extract IDs only

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projectId Unique or global ID of the project
 @return ApiApiV2ProjectsProjectIdWorkItemsSearchIdPostRequest
*/
func (a *ProjectWorkItemsApiService) ApiV2ProjectsProjectIdWorkItemsSearchIdPost(ctx context.Context, projectId string) ApiApiV2ProjectsProjectIdWorkItemsSearchIdPostRequest {
	return ApiApiV2ProjectsProjectIdWorkItemsSearchIdPostRequest{
		ApiService: a,
		ctx: ctx,
		projectId: projectId,
	}
}

// Execute executes the request
//  @return []string
func (a *ProjectWorkItemsApiService) ApiV2ProjectsProjectIdWorkItemsSearchIdPostExecute(r ApiApiV2ProjectsProjectIdWorkItemsSearchIdPostRequest) ([]string, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProjectWorkItemsApiService.ApiV2ProjectsProjectIdWorkItemsSearchIdPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/projects/{projectId}/workItems/search/id"
	localVarPath = strings.Replace(localVarPath, "{"+"projectId"+"}", url.PathEscape(parameterValueToString(r.projectId, "projectId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.skip != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "Skip", r.skip, "")
	}
	if r.take != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "Take", r.take, "")
	}
	if r.orderBy != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "OrderBy", r.orderBy, "")
	}
	if r.searchField != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "SearchField", r.searchField, "")
	}
	if r.searchValue != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "SearchValue", r.searchValue, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.apiV2ProjectsProjectIdWorkItemsSearchPostRequest
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["Bearer or PrivateToken"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiApiV2ProjectsProjectIdWorkItemsSearchPostRequest struct {
	ctx context.Context
	ApiService *ProjectWorkItemsApiService
	projectId string
	skip *int32
	take *int32
	orderBy *string
	searchField *string
	searchValue *string
	apiV2ProjectsProjectIdWorkItemsSearchPostRequest *ApiV2ProjectsProjectIdWorkItemsSearchPostRequest
}

// Amount of items to be skipped (offset)
func (r ApiApiV2ProjectsProjectIdWorkItemsSearchPostRequest) Skip(skip int32) ApiApiV2ProjectsProjectIdWorkItemsSearchPostRequest {
	r.skip = &skip
	return r
}

// Amount of items to be taken (limit)
func (r ApiApiV2ProjectsProjectIdWorkItemsSearchPostRequest) Take(take int32) ApiApiV2ProjectsProjectIdWorkItemsSearchPostRequest {
	r.take = &take
	return r
}

// SQL-like  ORDER BY statement (column1 ASC|DESC , column2 ASC|DESC)
func (r ApiApiV2ProjectsProjectIdWorkItemsSearchPostRequest) OrderBy(orderBy string) ApiApiV2ProjectsProjectIdWorkItemsSearchPostRequest {
	r.orderBy = &orderBy
	return r
}

// Property name for searching
func (r ApiApiV2ProjectsProjectIdWorkItemsSearchPostRequest) SearchField(searchField string) ApiApiV2ProjectsProjectIdWorkItemsSearchPostRequest {
	r.searchField = &searchField
	return r
}

// Value for searching
func (r ApiApiV2ProjectsProjectIdWorkItemsSearchPostRequest) SearchValue(searchValue string) ApiApiV2ProjectsProjectIdWorkItemsSearchPostRequest {
	r.searchValue = &searchValue
	return r
}

func (r ApiApiV2ProjectsProjectIdWorkItemsSearchPostRequest) ApiV2ProjectsProjectIdWorkItemsSearchPostRequest(apiV2ProjectsProjectIdWorkItemsSearchPostRequest ApiV2ProjectsProjectIdWorkItemsSearchPostRequest) ApiApiV2ProjectsProjectIdWorkItemsSearchPostRequest {
	r.apiV2ProjectsProjectIdWorkItemsSearchPostRequest = &apiV2ProjectsProjectIdWorkItemsSearchPostRequest
	return r
}

func (r ApiApiV2ProjectsProjectIdWorkItemsSearchPostRequest) Execute() ([]WorkItemShortModel, *http.Response, error) {
	return r.ApiService.ApiV2ProjectsProjectIdWorkItemsSearchPostExecute(r)
}

/*
ApiV2ProjectsProjectIdWorkItemsSearchPost Search for work items

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projectId Unique or global ID of the project
 @return ApiApiV2ProjectsProjectIdWorkItemsSearchPostRequest
*/
func (a *ProjectWorkItemsApiService) ApiV2ProjectsProjectIdWorkItemsSearchPost(ctx context.Context, projectId string) ApiApiV2ProjectsProjectIdWorkItemsSearchPostRequest {
	return ApiApiV2ProjectsProjectIdWorkItemsSearchPostRequest{
		ApiService: a,
		ctx: ctx,
		projectId: projectId,
	}
}

// Execute executes the request
//  @return []WorkItemShortModel
func (a *ProjectWorkItemsApiService) ApiV2ProjectsProjectIdWorkItemsSearchPostExecute(r ApiApiV2ProjectsProjectIdWorkItemsSearchPostRequest) ([]WorkItemShortModel, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []WorkItemShortModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProjectWorkItemsApiService.ApiV2ProjectsProjectIdWorkItemsSearchPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/projects/{projectId}/workItems/search"
	localVarPath = strings.Replace(localVarPath, "{"+"projectId"+"}", url.PathEscape(parameterValueToString(r.projectId, "projectId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.skip != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "Skip", r.skip, "")
	}
	if r.take != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "Take", r.take, "")
	}
	if r.orderBy != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "OrderBy", r.orderBy, "")
	}
	if r.searchField != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "SearchField", r.searchField, "")
	}
	if r.searchValue != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "SearchValue", r.searchValue, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.apiV2ProjectsProjectIdWorkItemsSearchPostRequest
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["Bearer or PrivateToken"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiApiV2ProjectsProjectIdWorkItemsTagsGetRequest struct {
	ctx context.Context
	ApiService *ProjectWorkItemsApiService
	projectId string
	isDeleted *bool
}

func (r ApiApiV2ProjectsProjectIdWorkItemsTagsGetRequest) IsDeleted(isDeleted bool) ApiApiV2ProjectsProjectIdWorkItemsTagsGetRequest {
	r.isDeleted = &isDeleted
	return r
}

func (r ApiApiV2ProjectsProjectIdWorkItemsTagsGetRequest) Execute() ([]TagShortModel, *http.Response, error) {
	return r.ApiService.ApiV2ProjectsProjectIdWorkItemsTagsGetExecute(r)
}

/*
ApiV2ProjectsProjectIdWorkItemsTagsGet Get WorkItems Tags

<br>Use case
<br>User sets project internal identifier 
<br>User runs method execution
<br>System returns work items tags

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projectId Project internal (UUID) identifier
 @return ApiApiV2ProjectsProjectIdWorkItemsTagsGetRequest
*/
func (a *ProjectWorkItemsApiService) ApiV2ProjectsProjectIdWorkItemsTagsGet(ctx context.Context, projectId string) ApiApiV2ProjectsProjectIdWorkItemsTagsGetRequest {
	return ApiApiV2ProjectsProjectIdWorkItemsTagsGetRequest{
		ApiService: a,
		ctx: ctx,
		projectId: projectId,
	}
}

// Execute executes the request
//  @return []TagShortModel
func (a *ProjectWorkItemsApiService) ApiV2ProjectsProjectIdWorkItemsTagsGetExecute(r ApiApiV2ProjectsProjectIdWorkItemsTagsGetRequest) ([]TagShortModel, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []TagShortModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProjectWorkItemsApiService.ApiV2ProjectsProjectIdWorkItemsTagsGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/projects/{projectId}/workItems/tags"
	localVarPath = strings.Replace(localVarPath, "{"+"projectId"+"}", url.PathEscape(parameterValueToString(r.projectId, "projectId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.isDeleted != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "isDeleted", r.isDeleted, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["Bearer or PrivateToken"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetWorkItemsByProjectIdRequest struct {
	ctx context.Context
	ApiService *ProjectWorkItemsApiService
	projectId string
	isDeleted *bool
	tagNames *[]string
	includeIterations *bool
	skip *int32
	take *int32
	orderBy *string
	searchField *string
	searchValue *string
}

// If result must consist of only actual/deleted work items
func (r ApiGetWorkItemsByProjectIdRequest) IsDeleted(isDeleted bool) ApiGetWorkItemsByProjectIdRequest {
	r.isDeleted = &isDeleted
	return r
}

// List of tags to filter by
func (r ApiGetWorkItemsByProjectIdRequest) TagNames(tagNames []string) ApiGetWorkItemsByProjectIdRequest {
	r.tagNames = &tagNames
	return r
}

func (r ApiGetWorkItemsByProjectIdRequest) IncludeIterations(includeIterations bool) ApiGetWorkItemsByProjectIdRequest {
	r.includeIterations = &includeIterations
	return r
}

// Amount of items to be skipped (offset)
func (r ApiGetWorkItemsByProjectIdRequest) Skip(skip int32) ApiGetWorkItemsByProjectIdRequest {
	r.skip = &skip
	return r
}

// Amount of items to be taken (limit)
func (r ApiGetWorkItemsByProjectIdRequest) Take(take int32) ApiGetWorkItemsByProjectIdRequest {
	r.take = &take
	return r
}

// SQL-like  ORDER BY statement (column1 ASC|DESC , column2 ASC|DESC)
func (r ApiGetWorkItemsByProjectIdRequest) OrderBy(orderBy string) ApiGetWorkItemsByProjectIdRequest {
	r.orderBy = &orderBy
	return r
}

// Property name for searching
func (r ApiGetWorkItemsByProjectIdRequest) SearchField(searchField string) ApiGetWorkItemsByProjectIdRequest {
	r.searchField = &searchField
	return r
}

// Value for searching
func (r ApiGetWorkItemsByProjectIdRequest) SearchValue(searchValue string) ApiGetWorkItemsByProjectIdRequest {
	r.searchValue = &searchValue
	return r
}

func (r ApiGetWorkItemsByProjectIdRequest) Execute() ([]WorkItemShortModel, *http.Response, error) {
	return r.ApiService.GetWorkItemsByProjectIdExecute(r)
}

/*
GetWorkItemsByProjectId Get project work items

<br>Use case
<br>User sets project internal or global identifier
<br>[Optional] User sets isDeleted field value
<br>User runs method execution
<br>System search project
<br>[Optional] If User sets isDeleted field value as true, System search all deleted workitems related to project
<br>[Optional] If User sets isDeleted field value as false, System search all workitems related to project which are not deleted
<br>If User did not set isDeleted field value, System search all  workitems related to project
<br>System returns array of found workitems (listed in response model)

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projectId Project internal (UUID) or global (integer) identifier
 @return ApiGetWorkItemsByProjectIdRequest

Deprecated
*/
func (a *ProjectWorkItemsApiService) GetWorkItemsByProjectId(ctx context.Context, projectId string) ApiGetWorkItemsByProjectIdRequest {
	return ApiGetWorkItemsByProjectIdRequest{
		ApiService: a,
		ctx: ctx,
		projectId: projectId,
	}
}

// Execute executes the request
//  @return []WorkItemShortModel
// Deprecated
func (a *ProjectWorkItemsApiService) GetWorkItemsByProjectIdExecute(r ApiGetWorkItemsByProjectIdRequest) ([]WorkItemShortModel, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []WorkItemShortModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProjectWorkItemsApiService.GetWorkItemsByProjectId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/projects/{projectId}/workItems"
	localVarPath = strings.Replace(localVarPath, "{"+"projectId"+"}", url.PathEscape(parameterValueToString(r.projectId, "projectId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.isDeleted != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "isDeleted", r.isDeleted, "")
	}
	if r.tagNames != nil {
		t := *r.tagNames
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "tagNames", s.Index(i), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "tagNames", t, "multi")
		}
	}
	if r.includeIterations != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "includeIterations", r.includeIterations, "")
	}
	if r.skip != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "Skip", r.skip, "")
	}
	if r.take != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "Take", r.take, "")
	}
	if r.orderBy != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "OrderBy", r.orderBy, "")
	}
	if r.searchField != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "SearchField", r.searchField, "")
	}
	if r.searchValue != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "SearchValue", r.searchValue, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["Bearer or PrivateToken"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}