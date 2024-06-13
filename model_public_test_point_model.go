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

// checks if the PublicTestPointModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PublicTestPointModel{}

// PublicTestPointModel struct for PublicTestPointModel
type PublicTestPointModel struct {
	ConfigurationId string `json:"configurationId"`
	ConfigurationGlobalId int64 `json:"configurationGlobalId"`
	AutoTestIds []string `json:"autoTestIds,omitempty"`
	IterationId string `json:"iterationId"`
	ParameterModels []ParameterShortModel `json:"parameterModels,omitempty"`
	Id string `json:"id"`
}

type _PublicTestPointModel PublicTestPointModel

// NewPublicTestPointModel instantiates a new PublicTestPointModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublicTestPointModel(configurationId string, configurationGlobalId int64, iterationId string, id string) *PublicTestPointModel {
	this := PublicTestPointModel{}
	this.ConfigurationId = configurationId
	this.ConfigurationGlobalId = configurationGlobalId
	this.IterationId = iterationId
	this.Id = id
	return &this
}

// NewPublicTestPointModelWithDefaults instantiates a new PublicTestPointModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublicTestPointModelWithDefaults() *PublicTestPointModel {
	this := PublicTestPointModel{}
	return &this
}

// GetConfigurationId returns the ConfigurationId field value
func (o *PublicTestPointModel) GetConfigurationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConfigurationId
}

// GetConfigurationIdOk returns a tuple with the ConfigurationId field value
// and a boolean to check if the value has been set.
func (o *PublicTestPointModel) GetConfigurationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConfigurationId, true
}

// SetConfigurationId sets field value
func (o *PublicTestPointModel) SetConfigurationId(v string) {
	o.ConfigurationId = v
}

// GetConfigurationGlobalId returns the ConfigurationGlobalId field value
func (o *PublicTestPointModel) GetConfigurationGlobalId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ConfigurationGlobalId
}

// GetConfigurationGlobalIdOk returns a tuple with the ConfigurationGlobalId field value
// and a boolean to check if the value has been set.
func (o *PublicTestPointModel) GetConfigurationGlobalIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConfigurationGlobalId, true
}

// SetConfigurationGlobalId sets field value
func (o *PublicTestPointModel) SetConfigurationGlobalId(v int64) {
	o.ConfigurationGlobalId = v
}

// GetAutoTestIds returns the AutoTestIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PublicTestPointModel) GetAutoTestIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.AutoTestIds
}

// GetAutoTestIdsOk returns a tuple with the AutoTestIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PublicTestPointModel) GetAutoTestIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.AutoTestIds) {
		return nil, false
	}
	return o.AutoTestIds, true
}

// HasAutoTestIds returns a boolean if a field has been set.
func (o *PublicTestPointModel) HasAutoTestIds() bool {
	if o != nil && IsNil(o.AutoTestIds) {
		return true
	}

	return false
}

// SetAutoTestIds gets a reference to the given []string and assigns it to the AutoTestIds field.
func (o *PublicTestPointModel) SetAutoTestIds(v []string) {
	o.AutoTestIds = v
}

// GetIterationId returns the IterationId field value
func (o *PublicTestPointModel) GetIterationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IterationId
}

// GetIterationIdOk returns a tuple with the IterationId field value
// and a boolean to check if the value has been set.
func (o *PublicTestPointModel) GetIterationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IterationId, true
}

// SetIterationId sets field value
func (o *PublicTestPointModel) SetIterationId(v string) {
	o.IterationId = v
}

// GetParameterModels returns the ParameterModels field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PublicTestPointModel) GetParameterModels() []ParameterShortModel {
	if o == nil {
		var ret []ParameterShortModel
		return ret
	}
	return o.ParameterModels
}

// GetParameterModelsOk returns a tuple with the ParameterModels field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PublicTestPointModel) GetParameterModelsOk() ([]ParameterShortModel, bool) {
	if o == nil || IsNil(o.ParameterModels) {
		return nil, false
	}
	return o.ParameterModels, true
}

// HasParameterModels returns a boolean if a field has been set.
func (o *PublicTestPointModel) HasParameterModels() bool {
	if o != nil && IsNil(o.ParameterModels) {
		return true
	}

	return false
}

// SetParameterModels gets a reference to the given []ParameterShortModel and assigns it to the ParameterModels field.
func (o *PublicTestPointModel) SetParameterModels(v []ParameterShortModel) {
	o.ParameterModels = v
}

// GetId returns the Id field value
func (o *PublicTestPointModel) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PublicTestPointModel) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PublicTestPointModel) SetId(v string) {
	o.Id = v
}

func (o PublicTestPointModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PublicTestPointModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["configurationId"] = o.ConfigurationId
	toSerialize["configurationGlobalId"] = o.ConfigurationGlobalId
	if o.AutoTestIds != nil {
		toSerialize["autoTestIds"] = o.AutoTestIds
	}
	toSerialize["iterationId"] = o.IterationId
	if o.ParameterModels != nil {
		toSerialize["parameterModels"] = o.ParameterModels
	}
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

func (o *PublicTestPointModel) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"configurationId",
		"configurationGlobalId",
		"iterationId",
		"id",
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

	varPublicTestPointModel := _PublicTestPointModel{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPublicTestPointModel)

	if err != nil {
		return err
	}

	*o = PublicTestPointModel(varPublicTestPointModel)

	return err
}

type NullablePublicTestPointModel struct {
	value *PublicTestPointModel
	isSet bool
}

func (v NullablePublicTestPointModel) Get() *PublicTestPointModel {
	return v.value
}

func (v *NullablePublicTestPointModel) Set(val *PublicTestPointModel) {
	v.value = val
	v.isSet = true
}

func (v NullablePublicTestPointModel) IsSet() bool {
	return v.isSet
}

func (v *NullablePublicTestPointModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePublicTestPointModel(val *PublicTestPointModel) *NullablePublicTestPointModel {
	return &NullablePublicTestPointModel{value: val, isSet: true}
}

func (v NullablePublicTestPointModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePublicTestPointModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


