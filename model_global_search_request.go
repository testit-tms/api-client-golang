/*
API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tmsclient

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the GlobalSearchRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GlobalSearchRequest{}

// GlobalSearchRequest struct for GlobalSearchRequest
type GlobalSearchRequest struct {
	Query string `json:"query"`
	ResourceType NullableString `json:"resourceType,omitempty"`
	Take int32 `json:"take"`
	Skip int32 `json:"skip"`
}

type _GlobalSearchRequest GlobalSearchRequest

// NewGlobalSearchRequest instantiates a new GlobalSearchRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGlobalSearchRequest(query string, take int32, skip int32) *GlobalSearchRequest {
	this := GlobalSearchRequest{}
	this.Query = query
	this.Take = take
	this.Skip = skip
	return &this
}

// NewGlobalSearchRequestWithDefaults instantiates a new GlobalSearchRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGlobalSearchRequestWithDefaults() *GlobalSearchRequest {
	this := GlobalSearchRequest{}
	return &this
}

// GetQuery returns the Query field value
func (o *GlobalSearchRequest) GetQuery() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *GlobalSearchRequest) GetQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value
func (o *GlobalSearchRequest) SetQuery(v string) {
	o.Query = v
}

// GetResourceType returns the ResourceType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GlobalSearchRequest) GetResourceType() string {
	if o == nil || IsNil(o.ResourceType.Get()) {
		var ret string
		return ret
	}
	return *o.ResourceType.Get()
}

// GetResourceTypeOk returns a tuple with the ResourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GlobalSearchRequest) GetResourceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ResourceType.Get(), o.ResourceType.IsSet()
}

// HasResourceType returns a boolean if a field has been set.
func (o *GlobalSearchRequest) HasResourceType() bool {
	if o != nil && o.ResourceType.IsSet() {
		return true
	}

	return false
}

// SetResourceType gets a reference to the given NullableString and assigns it to the ResourceType field.
func (o *GlobalSearchRequest) SetResourceType(v string) {
	o.ResourceType.Set(&v)
}
// SetResourceTypeNil sets the value for ResourceType to be an explicit nil
func (o *GlobalSearchRequest) SetResourceTypeNil() {
	o.ResourceType.Set(nil)
}

// UnsetResourceType ensures that no value is present for ResourceType, not even an explicit nil
func (o *GlobalSearchRequest) UnsetResourceType() {
	o.ResourceType.Unset()
}

// GetTake returns the Take field value
func (o *GlobalSearchRequest) GetTake() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Take
}

// GetTakeOk returns a tuple with the Take field value
// and a boolean to check if the value has been set.
func (o *GlobalSearchRequest) GetTakeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Take, true
}

// SetTake sets field value
func (o *GlobalSearchRequest) SetTake(v int32) {
	o.Take = v
}

// GetSkip returns the Skip field value
func (o *GlobalSearchRequest) GetSkip() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Skip
}

// GetSkipOk returns a tuple with the Skip field value
// and a boolean to check if the value has been set.
func (o *GlobalSearchRequest) GetSkipOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Skip, true
}

// SetSkip sets field value
func (o *GlobalSearchRequest) SetSkip(v int32) {
	o.Skip = v
}

func (o GlobalSearchRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GlobalSearchRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["query"] = o.Query
	if o.ResourceType.IsSet() {
		toSerialize["resourceType"] = o.ResourceType.Get()
	}
	toSerialize["take"] = o.Take
	toSerialize["skip"] = o.Skip
	return toSerialize, nil
}

func (o *GlobalSearchRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"query",
		"take",
		"skip",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varGlobalSearchRequest := _GlobalSearchRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGlobalSearchRequest)

	if err != nil {
		return err
	}

	*o = GlobalSearchRequest(varGlobalSearchRequest)

	return err
}

type NullableGlobalSearchRequest struct {
	value *GlobalSearchRequest
	isSet bool
}

func (v NullableGlobalSearchRequest) Get() *GlobalSearchRequest {
	return v.value
}

func (v *NullableGlobalSearchRequest) Set(val *GlobalSearchRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGlobalSearchRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGlobalSearchRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGlobalSearchRequest(val *GlobalSearchRequest) *NullableGlobalSearchRequest {
	return &NullableGlobalSearchRequest{value: val, isSet: true}
}

func (v NullableGlobalSearchRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGlobalSearchRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


