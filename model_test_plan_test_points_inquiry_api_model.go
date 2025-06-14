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

// checks if the TestPlanTestPointsInquiryApiModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TestPlanTestPointsInquiryApiModel{}

// TestPlanTestPointsInquiryApiModel struct for TestPlanTestPointsInquiryApiModel
type TestPlanTestPointsInquiryApiModel struct {
	Order []Order `json:"order"`
	Page NullablePage `json:"page,omitempty"`
}

type _TestPlanTestPointsInquiryApiModel TestPlanTestPointsInquiryApiModel

// NewTestPlanTestPointsInquiryApiModel instantiates a new TestPlanTestPointsInquiryApiModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTestPlanTestPointsInquiryApiModel(order []Order) *TestPlanTestPointsInquiryApiModel {
	this := TestPlanTestPointsInquiryApiModel{}
	this.Order = order
	return &this
}

// NewTestPlanTestPointsInquiryApiModelWithDefaults instantiates a new TestPlanTestPointsInquiryApiModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTestPlanTestPointsInquiryApiModelWithDefaults() *TestPlanTestPointsInquiryApiModel {
	this := TestPlanTestPointsInquiryApiModel{}
	return &this
}

// GetOrder returns the Order field value
func (o *TestPlanTestPointsInquiryApiModel) GetOrder() []Order {
	if o == nil {
		var ret []Order
		return ret
	}

	return o.Order
}

// GetOrderOk returns a tuple with the Order field value
// and a boolean to check if the value has been set.
func (o *TestPlanTestPointsInquiryApiModel) GetOrderOk() ([]Order, bool) {
	if o == nil {
		return nil, false
	}
	return o.Order, true
}

// SetOrder sets field value
func (o *TestPlanTestPointsInquiryApiModel) SetOrder(v []Order) {
	o.Order = v
}

// GetPage returns the Page field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestPlanTestPointsInquiryApiModel) GetPage() Page {
	if o == nil || IsNil(o.Page.Get()) {
		var ret Page
		return ret
	}
	return *o.Page.Get()
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestPlanTestPointsInquiryApiModel) GetPageOk() (*Page, bool) {
	if o == nil {
		return nil, false
	}
	return o.Page.Get(), o.Page.IsSet()
}

// HasPage returns a boolean if a field has been set.
func (o *TestPlanTestPointsInquiryApiModel) HasPage() bool {
	if o != nil && o.Page.IsSet() {
		return true
	}

	return false
}

// SetPage gets a reference to the given NullablePage and assigns it to the Page field.
func (o *TestPlanTestPointsInquiryApiModel) SetPage(v Page) {
	o.Page.Set(&v)
}
// SetPageNil sets the value for Page to be an explicit nil
func (o *TestPlanTestPointsInquiryApiModel) SetPageNil() {
	o.Page.Set(nil)
}

// UnsetPage ensures that no value is present for Page, not even an explicit nil
func (o *TestPlanTestPointsInquiryApiModel) UnsetPage() {
	o.Page.Unset()
}

func (o TestPlanTestPointsInquiryApiModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TestPlanTestPointsInquiryApiModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["order"] = o.Order
	if o.Page.IsSet() {
		toSerialize["page"] = o.Page.Get()
	}
	return toSerialize, nil
}

func (o *TestPlanTestPointsInquiryApiModel) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"order",
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

	varTestPlanTestPointsInquiryApiModel := _TestPlanTestPointsInquiryApiModel{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTestPlanTestPointsInquiryApiModel)

	if err != nil {
		return err
	}

	*o = TestPlanTestPointsInquiryApiModel(varTestPlanTestPointsInquiryApiModel)

	return err
}

type NullableTestPlanTestPointsInquiryApiModel struct {
	value *TestPlanTestPointsInquiryApiModel
	isSet bool
}

func (v NullableTestPlanTestPointsInquiryApiModel) Get() *TestPlanTestPointsInquiryApiModel {
	return v.value
}

func (v *NullableTestPlanTestPointsInquiryApiModel) Set(val *TestPlanTestPointsInquiryApiModel) {
	v.value = val
	v.isSet = true
}

func (v NullableTestPlanTestPointsInquiryApiModel) IsSet() bool {
	return v.isSet
}

func (v *NullableTestPlanTestPointsInquiryApiModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestPlanTestPointsInquiryApiModel(val *TestPlanTestPointsInquiryApiModel) *NullableTestPlanTestPointsInquiryApiModel {
	return &NullableTestPlanTestPointsInquiryApiModel{value: val, isSet: true}
}

func (v NullableTestPlanTestPointsInquiryApiModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTestPlanTestPointsInquiryApiModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


