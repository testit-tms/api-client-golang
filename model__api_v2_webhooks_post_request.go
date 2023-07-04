/*
API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tmsclient

import (
	"encoding/json"
)

// checks if the ApiV2WebhooksPostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiV2WebhooksPostRequest{}

// ApiV2WebhooksPostRequest struct for ApiV2WebhooksPostRequest
type ApiV2WebhooksPostRequest struct {
	// Unique ID of the webhook project
	ProjectId string `json:"projectId"`
	EventType WebHookEventTypeModel `json:"eventType"`
	// Description of the webhook
	Description NullableString `json:"description,omitempty"`
	// Request URL of the webhook
	Url string `json:"url"`
	RequestType RequestTypeModel `json:"requestType"`
	// Indicates if the webhook sends body
	ShouldSendBody *bool `json:"shouldSendBody,omitempty"`
	// Collection of the webhook headers
	Headers *map[string]string `json:"headers,omitempty"`
	// Collection of the webhook query parameters
	QueryParameters *map[string]string `json:"queryParameters,omitempty"`
	// Indicates if the webhook is active
	IsEnabled *bool `json:"isEnabled,omitempty"`
	// Indicates if the webhook sends custom body
	ShouldSendCustomBody *bool `json:"shouldSendCustomBody,omitempty"`
	// Custom body of the webhook
	CustomBody NullableString `json:"customBody,omitempty"`
	// Indicates if the webhook injects parameters
	ShouldReplaceParameters *bool `json:"shouldReplaceParameters,omitempty"`
	// Indicates if the webhook escapes invalid characters in parameters
	ShouldEscapeParameters *bool `json:"shouldEscapeParameters,omitempty"`
	// Name of the webhook
	Name string `json:"name"`
}

// NewApiV2WebhooksPostRequest instantiates a new ApiV2WebhooksPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiV2WebhooksPostRequest(projectId string, eventType WebHookEventTypeModel, url string, requestType RequestTypeModel, name string) *ApiV2WebhooksPostRequest {
	this := ApiV2WebhooksPostRequest{}
	this.ProjectId = projectId
	this.EventType = eventType
	this.Url = url
	this.RequestType = requestType
	this.Name = name
	return &this
}

// NewApiV2WebhooksPostRequestWithDefaults instantiates a new ApiV2WebhooksPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiV2WebhooksPostRequestWithDefaults() *ApiV2WebhooksPostRequest {
	this := ApiV2WebhooksPostRequest{}
	return &this
}

// GetProjectId returns the ProjectId field value
func (o *ApiV2WebhooksPostRequest) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *ApiV2WebhooksPostRequest) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *ApiV2WebhooksPostRequest) SetProjectId(v string) {
	o.ProjectId = v
}

// GetEventType returns the EventType field value
func (o *ApiV2WebhooksPostRequest) GetEventType() WebHookEventTypeModel {
	if o == nil {
		var ret WebHookEventTypeModel
		return ret
	}

	return o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value
// and a boolean to check if the value has been set.
func (o *ApiV2WebhooksPostRequest) GetEventTypeOk() (*WebHookEventTypeModel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventType, true
}

// SetEventType sets field value
func (o *ApiV2WebhooksPostRequest) SetEventType(v WebHookEventTypeModel) {
	o.EventType = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2WebhooksPostRequest) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2WebhooksPostRequest) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *ApiV2WebhooksPostRequest) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *ApiV2WebhooksPostRequest) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *ApiV2WebhooksPostRequest) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *ApiV2WebhooksPostRequest) UnsetDescription() {
	o.Description.Unset()
}

// GetUrl returns the Url field value
func (o *ApiV2WebhooksPostRequest) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *ApiV2WebhooksPostRequest) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *ApiV2WebhooksPostRequest) SetUrl(v string) {
	o.Url = v
}

// GetRequestType returns the RequestType field value
func (o *ApiV2WebhooksPostRequest) GetRequestType() RequestTypeModel {
	if o == nil {
		var ret RequestTypeModel
		return ret
	}

	return o.RequestType
}

// GetRequestTypeOk returns a tuple with the RequestType field value
// and a boolean to check if the value has been set.
func (o *ApiV2WebhooksPostRequest) GetRequestTypeOk() (*RequestTypeModel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestType, true
}

// SetRequestType sets field value
func (o *ApiV2WebhooksPostRequest) SetRequestType(v RequestTypeModel) {
	o.RequestType = v
}

// GetShouldSendBody returns the ShouldSendBody field value if set, zero value otherwise.
func (o *ApiV2WebhooksPostRequest) GetShouldSendBody() bool {
	if o == nil || IsNil(o.ShouldSendBody) {
		var ret bool
		return ret
	}
	return *o.ShouldSendBody
}

// GetShouldSendBodyOk returns a tuple with the ShouldSendBody field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiV2WebhooksPostRequest) GetShouldSendBodyOk() (*bool, bool) {
	if o == nil || IsNil(o.ShouldSendBody) {
		return nil, false
	}
	return o.ShouldSendBody, true
}

// HasShouldSendBody returns a boolean if a field has been set.
func (o *ApiV2WebhooksPostRequest) HasShouldSendBody() bool {
	if o != nil && !IsNil(o.ShouldSendBody) {
		return true
	}

	return false
}

// SetShouldSendBody gets a reference to the given bool and assigns it to the ShouldSendBody field.
func (o *ApiV2WebhooksPostRequest) SetShouldSendBody(v bool) {
	o.ShouldSendBody = &v
}

// GetHeaders returns the Headers field value if set, zero value otherwise.
func (o *ApiV2WebhooksPostRequest) GetHeaders() map[string]string {
	if o == nil || IsNil(o.Headers) {
		var ret map[string]string
		return ret
	}
	return *o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiV2WebhooksPostRequest) GetHeadersOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Headers) {
		return nil, false
	}
	return o.Headers, true
}

// HasHeaders returns a boolean if a field has been set.
func (o *ApiV2WebhooksPostRequest) HasHeaders() bool {
	if o != nil && !IsNil(o.Headers) {
		return true
	}

	return false
}

// SetHeaders gets a reference to the given map[string]string and assigns it to the Headers field.
func (o *ApiV2WebhooksPostRequest) SetHeaders(v map[string]string) {
	o.Headers = &v
}

// GetQueryParameters returns the QueryParameters field value if set, zero value otherwise.
func (o *ApiV2WebhooksPostRequest) GetQueryParameters() map[string]string {
	if o == nil || IsNil(o.QueryParameters) {
		var ret map[string]string
		return ret
	}
	return *o.QueryParameters
}

// GetQueryParametersOk returns a tuple with the QueryParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiV2WebhooksPostRequest) GetQueryParametersOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.QueryParameters) {
		return nil, false
	}
	return o.QueryParameters, true
}

// HasQueryParameters returns a boolean if a field has been set.
func (o *ApiV2WebhooksPostRequest) HasQueryParameters() bool {
	if o != nil && !IsNil(o.QueryParameters) {
		return true
	}

	return false
}

// SetQueryParameters gets a reference to the given map[string]string and assigns it to the QueryParameters field.
func (o *ApiV2WebhooksPostRequest) SetQueryParameters(v map[string]string) {
	o.QueryParameters = &v
}

// GetIsEnabled returns the IsEnabled field value if set, zero value otherwise.
func (o *ApiV2WebhooksPostRequest) GetIsEnabled() bool {
	if o == nil || IsNil(o.IsEnabled) {
		var ret bool
		return ret
	}
	return *o.IsEnabled
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiV2WebhooksPostRequest) GetIsEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsEnabled) {
		return nil, false
	}
	return o.IsEnabled, true
}

// HasIsEnabled returns a boolean if a field has been set.
func (o *ApiV2WebhooksPostRequest) HasIsEnabled() bool {
	if o != nil && !IsNil(o.IsEnabled) {
		return true
	}

	return false
}

// SetIsEnabled gets a reference to the given bool and assigns it to the IsEnabled field.
func (o *ApiV2WebhooksPostRequest) SetIsEnabled(v bool) {
	o.IsEnabled = &v
}

// GetShouldSendCustomBody returns the ShouldSendCustomBody field value if set, zero value otherwise.
func (o *ApiV2WebhooksPostRequest) GetShouldSendCustomBody() bool {
	if o == nil || IsNil(o.ShouldSendCustomBody) {
		var ret bool
		return ret
	}
	return *o.ShouldSendCustomBody
}

// GetShouldSendCustomBodyOk returns a tuple with the ShouldSendCustomBody field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiV2WebhooksPostRequest) GetShouldSendCustomBodyOk() (*bool, bool) {
	if o == nil || IsNil(o.ShouldSendCustomBody) {
		return nil, false
	}
	return o.ShouldSendCustomBody, true
}

// HasShouldSendCustomBody returns a boolean if a field has been set.
func (o *ApiV2WebhooksPostRequest) HasShouldSendCustomBody() bool {
	if o != nil && !IsNil(o.ShouldSendCustomBody) {
		return true
	}

	return false
}

// SetShouldSendCustomBody gets a reference to the given bool and assigns it to the ShouldSendCustomBody field.
func (o *ApiV2WebhooksPostRequest) SetShouldSendCustomBody(v bool) {
	o.ShouldSendCustomBody = &v
}

// GetCustomBody returns the CustomBody field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2WebhooksPostRequest) GetCustomBody() string {
	if o == nil || IsNil(o.CustomBody.Get()) {
		var ret string
		return ret
	}
	return *o.CustomBody.Get()
}

// GetCustomBodyOk returns a tuple with the CustomBody field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2WebhooksPostRequest) GetCustomBodyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CustomBody.Get(), o.CustomBody.IsSet()
}

// HasCustomBody returns a boolean if a field has been set.
func (o *ApiV2WebhooksPostRequest) HasCustomBody() bool {
	if o != nil && o.CustomBody.IsSet() {
		return true
	}

	return false
}

// SetCustomBody gets a reference to the given NullableString and assigns it to the CustomBody field.
func (o *ApiV2WebhooksPostRequest) SetCustomBody(v string) {
	o.CustomBody.Set(&v)
}
// SetCustomBodyNil sets the value for CustomBody to be an explicit nil
func (o *ApiV2WebhooksPostRequest) SetCustomBodyNil() {
	o.CustomBody.Set(nil)
}

// UnsetCustomBody ensures that no value is present for CustomBody, not even an explicit nil
func (o *ApiV2WebhooksPostRequest) UnsetCustomBody() {
	o.CustomBody.Unset()
}

// GetShouldReplaceParameters returns the ShouldReplaceParameters field value if set, zero value otherwise.
func (o *ApiV2WebhooksPostRequest) GetShouldReplaceParameters() bool {
	if o == nil || IsNil(o.ShouldReplaceParameters) {
		var ret bool
		return ret
	}
	return *o.ShouldReplaceParameters
}

// GetShouldReplaceParametersOk returns a tuple with the ShouldReplaceParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiV2WebhooksPostRequest) GetShouldReplaceParametersOk() (*bool, bool) {
	if o == nil || IsNil(o.ShouldReplaceParameters) {
		return nil, false
	}
	return o.ShouldReplaceParameters, true
}

// HasShouldReplaceParameters returns a boolean if a field has been set.
func (o *ApiV2WebhooksPostRequest) HasShouldReplaceParameters() bool {
	if o != nil && !IsNil(o.ShouldReplaceParameters) {
		return true
	}

	return false
}

// SetShouldReplaceParameters gets a reference to the given bool and assigns it to the ShouldReplaceParameters field.
func (o *ApiV2WebhooksPostRequest) SetShouldReplaceParameters(v bool) {
	o.ShouldReplaceParameters = &v
}

// GetShouldEscapeParameters returns the ShouldEscapeParameters field value if set, zero value otherwise.
func (o *ApiV2WebhooksPostRequest) GetShouldEscapeParameters() bool {
	if o == nil || IsNil(o.ShouldEscapeParameters) {
		var ret bool
		return ret
	}
	return *o.ShouldEscapeParameters
}

// GetShouldEscapeParametersOk returns a tuple with the ShouldEscapeParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiV2WebhooksPostRequest) GetShouldEscapeParametersOk() (*bool, bool) {
	if o == nil || IsNil(o.ShouldEscapeParameters) {
		return nil, false
	}
	return o.ShouldEscapeParameters, true
}

// HasShouldEscapeParameters returns a boolean if a field has been set.
func (o *ApiV2WebhooksPostRequest) HasShouldEscapeParameters() bool {
	if o != nil && !IsNil(o.ShouldEscapeParameters) {
		return true
	}

	return false
}

// SetShouldEscapeParameters gets a reference to the given bool and assigns it to the ShouldEscapeParameters field.
func (o *ApiV2WebhooksPostRequest) SetShouldEscapeParameters(v bool) {
	o.ShouldEscapeParameters = &v
}

// GetName returns the Name field value
func (o *ApiV2WebhooksPostRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ApiV2WebhooksPostRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ApiV2WebhooksPostRequest) SetName(v string) {
	o.Name = v
}

func (o ApiV2WebhooksPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiV2WebhooksPostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["projectId"] = o.ProjectId
	toSerialize["eventType"] = o.EventType
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	toSerialize["url"] = o.Url
	toSerialize["requestType"] = o.RequestType
	if !IsNil(o.ShouldSendBody) {
		toSerialize["shouldSendBody"] = o.ShouldSendBody
	}
	if !IsNil(o.Headers) {
		toSerialize["headers"] = o.Headers
	}
	if !IsNil(o.QueryParameters) {
		toSerialize["queryParameters"] = o.QueryParameters
	}
	if !IsNil(o.IsEnabled) {
		toSerialize["isEnabled"] = o.IsEnabled
	}
	if !IsNil(o.ShouldSendCustomBody) {
		toSerialize["shouldSendCustomBody"] = o.ShouldSendCustomBody
	}
	if o.CustomBody.IsSet() {
		toSerialize["customBody"] = o.CustomBody.Get()
	}
	if !IsNil(o.ShouldReplaceParameters) {
		toSerialize["shouldReplaceParameters"] = o.ShouldReplaceParameters
	}
	if !IsNil(o.ShouldEscapeParameters) {
		toSerialize["shouldEscapeParameters"] = o.ShouldEscapeParameters
	}
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

type NullableApiV2WebhooksPostRequest struct {
	value *ApiV2WebhooksPostRequest
	isSet bool
}

func (v NullableApiV2WebhooksPostRequest) Get() *ApiV2WebhooksPostRequest {
	return v.value
}

func (v *NullableApiV2WebhooksPostRequest) Set(val *ApiV2WebhooksPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableApiV2WebhooksPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableApiV2WebhooksPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiV2WebhooksPostRequest(val *ApiV2WebhooksPostRequest) *NullableApiV2WebhooksPostRequest {
	return &NullableApiV2WebhooksPostRequest{value: val, isSet: true}
}

func (v NullableApiV2WebhooksPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiV2WebhooksPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


