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

// checks if the WebHookPostModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebHookPostModel{}

// WebHookPostModel struct for WebHookPostModel
type WebHookPostModel struct {
	// Unique ID of the webhook project
	ProjectId string `json:"projectId"`
	EventType WebHookEventTypeModel `json:"eventType"`
	// Description of the webhook
	Description NullableString `json:"description,omitempty"`
	// Request URL of the webhook
	Url string `json:"url"`
	RequestType RequestTypeModel `json:"requestType"`
	// Indicates if the webhook sends body
	ShouldSendBody bool `json:"shouldSendBody"`
	// Collection of the webhook headers
	Headers map[string]string `json:"headers,omitempty"`
	// Collection of the webhook query parameters
	QueryParameters map[string]string `json:"queryParameters,omitempty"`
	// Indicates if the webhook is active
	IsEnabled bool `json:"isEnabled"`
	// Indicates if the webhook sends custom body
	ShouldSendCustomBody bool `json:"shouldSendCustomBody"`
	// Custom body of the webhook
	CustomBody NullableString `json:"customBody,omitempty"`
	// Indicates if the webhook injects parameters
	ShouldReplaceParameters bool `json:"shouldReplaceParameters"`
	// Indicates if the webhook escapes invalid characters in parameters
	ShouldEscapeParameters bool `json:"shouldEscapeParameters"`
	// Name of the webhook
	Name string `json:"name"`
}

// NewWebHookPostModel instantiates a new WebHookPostModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebHookPostModel(projectId string, eventType WebHookEventTypeModel, url string, requestType RequestTypeModel, shouldSendBody bool, isEnabled bool, shouldSendCustomBody bool, shouldReplaceParameters bool, shouldEscapeParameters bool, name string) *WebHookPostModel {
	this := WebHookPostModel{}
	this.ProjectId = projectId
	this.EventType = eventType
	this.Url = url
	this.RequestType = requestType
	this.ShouldSendBody = shouldSendBody
	this.IsEnabled = isEnabled
	this.ShouldSendCustomBody = shouldSendCustomBody
	this.ShouldReplaceParameters = shouldReplaceParameters
	this.ShouldEscapeParameters = shouldEscapeParameters
	this.Name = name
	return &this
}

// NewWebHookPostModelWithDefaults instantiates a new WebHookPostModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebHookPostModelWithDefaults() *WebHookPostModel {
	this := WebHookPostModel{}
	return &this
}

// GetProjectId returns the ProjectId field value
func (o *WebHookPostModel) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *WebHookPostModel) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *WebHookPostModel) SetProjectId(v string) {
	o.ProjectId = v
}

// GetEventType returns the EventType field value
func (o *WebHookPostModel) GetEventType() WebHookEventTypeModel {
	if o == nil {
		var ret WebHookEventTypeModel
		return ret
	}

	return o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value
// and a boolean to check if the value has been set.
func (o *WebHookPostModel) GetEventTypeOk() (*WebHookEventTypeModel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventType, true
}

// SetEventType sets field value
func (o *WebHookPostModel) SetEventType(v WebHookEventTypeModel) {
	o.EventType = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebHookPostModel) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebHookPostModel) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *WebHookPostModel) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *WebHookPostModel) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *WebHookPostModel) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *WebHookPostModel) UnsetDescription() {
	o.Description.Unset()
}

// GetUrl returns the Url field value
func (o *WebHookPostModel) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *WebHookPostModel) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *WebHookPostModel) SetUrl(v string) {
	o.Url = v
}

// GetRequestType returns the RequestType field value
func (o *WebHookPostModel) GetRequestType() RequestTypeModel {
	if o == nil {
		var ret RequestTypeModel
		return ret
	}

	return o.RequestType
}

// GetRequestTypeOk returns a tuple with the RequestType field value
// and a boolean to check if the value has been set.
func (o *WebHookPostModel) GetRequestTypeOk() (*RequestTypeModel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestType, true
}

// SetRequestType sets field value
func (o *WebHookPostModel) SetRequestType(v RequestTypeModel) {
	o.RequestType = v
}

// GetShouldSendBody returns the ShouldSendBody field value
func (o *WebHookPostModel) GetShouldSendBody() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ShouldSendBody
}

// GetShouldSendBodyOk returns a tuple with the ShouldSendBody field value
// and a boolean to check if the value has been set.
func (o *WebHookPostModel) GetShouldSendBodyOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShouldSendBody, true
}

// SetShouldSendBody sets field value
func (o *WebHookPostModel) SetShouldSendBody(v bool) {
	o.ShouldSendBody = v
}

// GetHeaders returns the Headers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebHookPostModel) GetHeaders() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}
	return o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebHookPostModel) GetHeadersOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Headers) {
		return nil, false
	}
	return &o.Headers, true
}

// HasHeaders returns a boolean if a field has been set.
func (o *WebHookPostModel) HasHeaders() bool {
	if o != nil && IsNil(o.Headers) {
		return true
	}

	return false
}

// SetHeaders gets a reference to the given map[string]string and assigns it to the Headers field.
func (o *WebHookPostModel) SetHeaders(v map[string]string) {
	o.Headers = v
}

// GetQueryParameters returns the QueryParameters field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebHookPostModel) GetQueryParameters() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}
	return o.QueryParameters
}

// GetQueryParametersOk returns a tuple with the QueryParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebHookPostModel) GetQueryParametersOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.QueryParameters) {
		return nil, false
	}
	return &o.QueryParameters, true
}

// HasQueryParameters returns a boolean if a field has been set.
func (o *WebHookPostModel) HasQueryParameters() bool {
	if o != nil && IsNil(o.QueryParameters) {
		return true
	}

	return false
}

// SetQueryParameters gets a reference to the given map[string]string and assigns it to the QueryParameters field.
func (o *WebHookPostModel) SetQueryParameters(v map[string]string) {
	o.QueryParameters = v
}

// GetIsEnabled returns the IsEnabled field value
func (o *WebHookPostModel) GetIsEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsEnabled
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value
// and a boolean to check if the value has been set.
func (o *WebHookPostModel) GetIsEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsEnabled, true
}

// SetIsEnabled sets field value
func (o *WebHookPostModel) SetIsEnabled(v bool) {
	o.IsEnabled = v
}

// GetShouldSendCustomBody returns the ShouldSendCustomBody field value
func (o *WebHookPostModel) GetShouldSendCustomBody() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ShouldSendCustomBody
}

// GetShouldSendCustomBodyOk returns a tuple with the ShouldSendCustomBody field value
// and a boolean to check if the value has been set.
func (o *WebHookPostModel) GetShouldSendCustomBodyOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShouldSendCustomBody, true
}

// SetShouldSendCustomBody sets field value
func (o *WebHookPostModel) SetShouldSendCustomBody(v bool) {
	o.ShouldSendCustomBody = v
}

// GetCustomBody returns the CustomBody field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebHookPostModel) GetCustomBody() string {
	if o == nil || IsNil(o.CustomBody.Get()) {
		var ret string
		return ret
	}
	return *o.CustomBody.Get()
}

// GetCustomBodyOk returns a tuple with the CustomBody field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebHookPostModel) GetCustomBodyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CustomBody.Get(), o.CustomBody.IsSet()
}

// HasCustomBody returns a boolean if a field has been set.
func (o *WebHookPostModel) HasCustomBody() bool {
	if o != nil && o.CustomBody.IsSet() {
		return true
	}

	return false
}

// SetCustomBody gets a reference to the given NullableString and assigns it to the CustomBody field.
func (o *WebHookPostModel) SetCustomBody(v string) {
	o.CustomBody.Set(&v)
}
// SetCustomBodyNil sets the value for CustomBody to be an explicit nil
func (o *WebHookPostModel) SetCustomBodyNil() {
	o.CustomBody.Set(nil)
}

// UnsetCustomBody ensures that no value is present for CustomBody, not even an explicit nil
func (o *WebHookPostModel) UnsetCustomBody() {
	o.CustomBody.Unset()
}

// GetShouldReplaceParameters returns the ShouldReplaceParameters field value
func (o *WebHookPostModel) GetShouldReplaceParameters() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ShouldReplaceParameters
}

// GetShouldReplaceParametersOk returns a tuple with the ShouldReplaceParameters field value
// and a boolean to check if the value has been set.
func (o *WebHookPostModel) GetShouldReplaceParametersOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShouldReplaceParameters, true
}

// SetShouldReplaceParameters sets field value
func (o *WebHookPostModel) SetShouldReplaceParameters(v bool) {
	o.ShouldReplaceParameters = v
}

// GetShouldEscapeParameters returns the ShouldEscapeParameters field value
func (o *WebHookPostModel) GetShouldEscapeParameters() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ShouldEscapeParameters
}

// GetShouldEscapeParametersOk returns a tuple with the ShouldEscapeParameters field value
// and a boolean to check if the value has been set.
func (o *WebHookPostModel) GetShouldEscapeParametersOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShouldEscapeParameters, true
}

// SetShouldEscapeParameters sets field value
func (o *WebHookPostModel) SetShouldEscapeParameters(v bool) {
	o.ShouldEscapeParameters = v
}

// GetName returns the Name field value
func (o *WebHookPostModel) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *WebHookPostModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *WebHookPostModel) SetName(v string) {
	o.Name = v
}

func (o WebHookPostModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebHookPostModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["projectId"] = o.ProjectId
	toSerialize["eventType"] = o.EventType
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	toSerialize["url"] = o.Url
	toSerialize["requestType"] = o.RequestType
	toSerialize["shouldSendBody"] = o.ShouldSendBody
	if o.Headers != nil {
		toSerialize["headers"] = o.Headers
	}
	if o.QueryParameters != nil {
		toSerialize["queryParameters"] = o.QueryParameters
	}
	toSerialize["isEnabled"] = o.IsEnabled
	toSerialize["shouldSendCustomBody"] = o.ShouldSendCustomBody
	if o.CustomBody.IsSet() {
		toSerialize["customBody"] = o.CustomBody.Get()
	}
	toSerialize["shouldReplaceParameters"] = o.ShouldReplaceParameters
	toSerialize["shouldEscapeParameters"] = o.ShouldEscapeParameters
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

type NullableWebHookPostModel struct {
	value *WebHookPostModel
	isSet bool
}

func (v NullableWebHookPostModel) Get() *WebHookPostModel {
	return v.value
}

func (v *NullableWebHookPostModel) Set(val *WebHookPostModel) {
	v.value = val
	v.isSet = true
}

func (v NullableWebHookPostModel) IsSet() bool {
	return v.isSet
}

func (v *NullableWebHookPostModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebHookPostModel(val *WebHookPostModel) *NullableWebHookPostModel {
	return &NullableWebHookPostModel{value: val, isSet: true}
}

func (v NullableWebHookPostModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebHookPostModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


