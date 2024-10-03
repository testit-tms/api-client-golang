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
)


// ProjectAttributeTemplatesApiService ProjectAttributeTemplatesApi service
type ProjectAttributeTemplatesApiService service

type ApiApiV2ProjectsProjectIdAttributesTemplatesSearchPostRequest struct {
	ctx context.Context
	ApiService *ProjectAttributeTemplatesApiService
	projectId string
	skip *int32
	take *int32
	orderBy *string
	searchField *string
	searchValue *string
	apiV2ProjectsProjectIdAttributesTemplatesSearchPostRequest *ApiV2ProjectsProjectIdAttributesTemplatesSearchPostRequest
}

// Amount of items to be skipped (offset)
func (r ApiApiV2ProjectsProjectIdAttributesTemplatesSearchPostRequest) Skip(skip int32) ApiApiV2ProjectsProjectIdAttributesTemplatesSearchPostRequest {
	r.skip = &skip
	return r
}

// Amount of items to be taken (limit)
func (r ApiApiV2ProjectsProjectIdAttributesTemplatesSearchPostRequest) Take(take int32) ApiApiV2ProjectsProjectIdAttributesTemplatesSearchPostRequest {
	r.take = &take
	return r
}

// SQL-like  ORDER BY statement (column1 ASC|DESC , column2 ASC|DESC)
func (r ApiApiV2ProjectsProjectIdAttributesTemplatesSearchPostRequest) OrderBy(orderBy string) ApiApiV2ProjectsProjectIdAttributesTemplatesSearchPostRequest {
	r.orderBy = &orderBy
	return r
}

// Property name for searching
func (r ApiApiV2ProjectsProjectIdAttributesTemplatesSearchPostRequest) SearchField(searchField string) ApiApiV2ProjectsProjectIdAttributesTemplatesSearchPostRequest {
	r.searchField = &searchField
	return r
}

// Value for searching
func (r ApiApiV2ProjectsProjectIdAttributesTemplatesSearchPostRequest) SearchValue(searchValue string) ApiApiV2ProjectsProjectIdAttributesTemplatesSearchPostRequest {
	r.searchValue = &searchValue
	return r
}

func (r ApiApiV2ProjectsProjectIdAttributesTemplatesSearchPostRequest) ApiV2ProjectsProjectIdAttributesTemplatesSearchPostRequest(apiV2ProjectsProjectIdAttributesTemplatesSearchPostRequest ApiV2ProjectsProjectIdAttributesTemplatesSearchPostRequest) ApiApiV2ProjectsProjectIdAttributesTemplatesSearchPostRequest {
	r.apiV2ProjectsProjectIdAttributesTemplatesSearchPostRequest = &apiV2ProjectsProjectIdAttributesTemplatesSearchPostRequest
	return r
}

func (r ApiApiV2ProjectsProjectIdAttributesTemplatesSearchPostRequest) Execute() ([]ProjectCustomAttributeTemplateGetModel, *http.Response, error) {
	return r.ApiService.ApiV2ProjectsProjectIdAttributesTemplatesSearchPostExecute(r)
}

/*
ApiV2ProjectsProjectIdAttributesTemplatesSearchPost Search for custom attributes templates

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projectId
 @return ApiApiV2ProjectsProjectIdAttributesTemplatesSearchPostRequest
*/
func (a *ProjectAttributeTemplatesApiService) ApiV2ProjectsProjectIdAttributesTemplatesSearchPost(ctx context.Context, projectId string) ApiApiV2ProjectsProjectIdAttributesTemplatesSearchPostRequest {
	return ApiApiV2ProjectsProjectIdAttributesTemplatesSearchPostRequest{
		ApiService: a,
		ctx: ctx,
		projectId: projectId,
	}
}

// Execute executes the request
//  @return []ProjectCustomAttributeTemplateGetModel
func (a *ProjectAttributeTemplatesApiService) ApiV2ProjectsProjectIdAttributesTemplatesSearchPostExecute(r ApiApiV2ProjectsProjectIdAttributesTemplatesSearchPostRequest) ([]ProjectCustomAttributeTemplateGetModel, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []ProjectCustomAttributeTemplateGetModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProjectAttributeTemplatesApiService.ApiV2ProjectsProjectIdAttributesTemplatesSearchPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/projects/{projectId}/attributes/templates/search"
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
	localVarPostBody = r.apiV2ProjectsProjectIdAttributesTemplatesSearchPostRequest
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v ValidationProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
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
		if localVarHTTPResponse.StatusCode == 404 {
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
		if localVarHTTPResponse.StatusCode == 409 {
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
		if localVarHTTPResponse.StatusCode == 422 {
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

type ApiApiV2ProjectsProjectIdAttributesTemplatesTemplateIdDeleteRequest struct {
	ctx context.Context
	ApiService *ProjectAttributeTemplatesApiService
	projectId string
	templateId string
}

func (r ApiApiV2ProjectsProjectIdAttributesTemplatesTemplateIdDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.ApiV2ProjectsProjectIdAttributesTemplatesTemplateIdDeleteExecute(r)
}

/*
ApiV2ProjectsProjectIdAttributesTemplatesTemplateIdDelete Delete CustomAttributeTemplate from Project


Use case

User sets project internal or global identifier 

User sets attribute template internal identifier 

User runs method execution

System delete attribute template from project

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projectId Project internal (UUID) or global (integer) identifier
 @param templateId CustomAttributeTemplate internal (UUID) identifier
 @return ApiApiV2ProjectsProjectIdAttributesTemplatesTemplateIdDeleteRequest
*/
func (a *ProjectAttributeTemplatesApiService) ApiV2ProjectsProjectIdAttributesTemplatesTemplateIdDelete(ctx context.Context, projectId string, templateId string) ApiApiV2ProjectsProjectIdAttributesTemplatesTemplateIdDeleteRequest {
	return ApiApiV2ProjectsProjectIdAttributesTemplatesTemplateIdDeleteRequest{
		ApiService: a,
		ctx: ctx,
		projectId: projectId,
		templateId: templateId,
	}
}

// Execute executes the request
func (a *ProjectAttributeTemplatesApiService) ApiV2ProjectsProjectIdAttributesTemplatesTemplateIdDeleteExecute(r ApiApiV2ProjectsProjectIdAttributesTemplatesTemplateIdDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProjectAttributeTemplatesApiService.ApiV2ProjectsProjectIdAttributesTemplatesTemplateIdDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/projects/{projectId}/attributes/templates/{templateId}"
	localVarPath = strings.Replace(localVarPath, "{"+"projectId"+"}", url.PathEscape(parameterValueToString(r.projectId, "projectId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"templateId"+"}", url.PathEscape(parameterValueToString(r.templateId, "templateId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ValidationProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 409 {
			var v ProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v ProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiApiV2ProjectsProjectIdAttributesTemplatesTemplateIdPostRequest struct {
	ctx context.Context
	ApiService *ProjectAttributeTemplatesApiService
	projectId string
	templateId string
}

func (r ApiApiV2ProjectsProjectIdAttributesTemplatesTemplateIdPostRequest) Execute() (*http.Response, error) {
	return r.ApiService.ApiV2ProjectsProjectIdAttributesTemplatesTemplateIdPostExecute(r)
}

/*
ApiV2ProjectsProjectIdAttributesTemplatesTemplateIdPost Add CustomAttributeTemplate to Project


Use case

User sets project internal or global identifier 

User sets attribute template internal identifier 

User runs method execution

System add attribute template to project

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projectId Project internal (UUID) or global (integer) identifier
 @param templateId CustomAttributeTemplate internal (UUID) identifier
 @return ApiApiV2ProjectsProjectIdAttributesTemplatesTemplateIdPostRequest
*/
func (a *ProjectAttributeTemplatesApiService) ApiV2ProjectsProjectIdAttributesTemplatesTemplateIdPost(ctx context.Context, projectId string, templateId string) ApiApiV2ProjectsProjectIdAttributesTemplatesTemplateIdPostRequest {
	return ApiApiV2ProjectsProjectIdAttributesTemplatesTemplateIdPostRequest{
		ApiService: a,
		ctx: ctx,
		projectId: projectId,
		templateId: templateId,
	}
}

// Execute executes the request
func (a *ProjectAttributeTemplatesApiService) ApiV2ProjectsProjectIdAttributesTemplatesTemplateIdPostExecute(r ApiApiV2ProjectsProjectIdAttributesTemplatesTemplateIdPostRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProjectAttributeTemplatesApiService.ApiV2ProjectsProjectIdAttributesTemplatesTemplateIdPost")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/projects/{projectId}/attributes/templates/{templateId}"
	localVarPath = strings.Replace(localVarPath, "{"+"projectId"+"}", url.PathEscape(parameterValueToString(r.projectId, "projectId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"templateId"+"}", url.PathEscape(parameterValueToString(r.templateId, "templateId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ValidationProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 409 {
			var v ProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v ProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
