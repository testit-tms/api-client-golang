/*
API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tmsclient

import (
	"encoding/json"
	"time"
)

// checks if the TestPointFilterModelWorkItemModifiedDate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TestPointFilterModelWorkItemModifiedDate{}

// TestPointFilterModelWorkItemModifiedDate Specifies a work item range of last modification date to search for
type TestPointFilterModelWorkItemModifiedDate struct {
	From NullableTime `json:"from,omitempty"`
	To NullableTime `json:"to,omitempty"`
}

// NewTestPointFilterModelWorkItemModifiedDate instantiates a new TestPointFilterModelWorkItemModifiedDate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTestPointFilterModelWorkItemModifiedDate() *TestPointFilterModelWorkItemModifiedDate {
	this := TestPointFilterModelWorkItemModifiedDate{}
	return &this
}

// NewTestPointFilterModelWorkItemModifiedDateWithDefaults instantiates a new TestPointFilterModelWorkItemModifiedDate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTestPointFilterModelWorkItemModifiedDateWithDefaults() *TestPointFilterModelWorkItemModifiedDate {
	this := TestPointFilterModelWorkItemModifiedDate{}
	return &this
}

// GetFrom returns the From field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestPointFilterModelWorkItemModifiedDate) GetFrom() time.Time {
	if o == nil || IsNil(o.From.Get()) {
		var ret time.Time
		return ret
	}
	return *o.From.Get()
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestPointFilterModelWorkItemModifiedDate) GetFromOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.From.Get(), o.From.IsSet()
}

// HasFrom returns a boolean if a field has been set.
func (o *TestPointFilterModelWorkItemModifiedDate) HasFrom() bool {
	if o != nil && o.From.IsSet() {
		return true
	}

	return false
}

// SetFrom gets a reference to the given NullableTime and assigns it to the From field.
func (o *TestPointFilterModelWorkItemModifiedDate) SetFrom(v time.Time) {
	o.From.Set(&v)
}
// SetFromNil sets the value for From to be an explicit nil
func (o *TestPointFilterModelWorkItemModifiedDate) SetFromNil() {
	o.From.Set(nil)
}

// UnsetFrom ensures that no value is present for From, not even an explicit nil
func (o *TestPointFilterModelWorkItemModifiedDate) UnsetFrom() {
	o.From.Unset()
}

// GetTo returns the To field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestPointFilterModelWorkItemModifiedDate) GetTo() time.Time {
	if o == nil || IsNil(o.To.Get()) {
		var ret time.Time
		return ret
	}
	return *o.To.Get()
}

// GetToOk returns a tuple with the To field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestPointFilterModelWorkItemModifiedDate) GetToOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.To.Get(), o.To.IsSet()
}

// HasTo returns a boolean if a field has been set.
func (o *TestPointFilterModelWorkItemModifiedDate) HasTo() bool {
	if o != nil && o.To.IsSet() {
		return true
	}

	return false
}

// SetTo gets a reference to the given NullableTime and assigns it to the To field.
func (o *TestPointFilterModelWorkItemModifiedDate) SetTo(v time.Time) {
	o.To.Set(&v)
}
// SetToNil sets the value for To to be an explicit nil
func (o *TestPointFilterModelWorkItemModifiedDate) SetToNil() {
	o.To.Set(nil)
}

// UnsetTo ensures that no value is present for To, not even an explicit nil
func (o *TestPointFilterModelWorkItemModifiedDate) UnsetTo() {
	o.To.Unset()
}

func (o TestPointFilterModelWorkItemModifiedDate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TestPointFilterModelWorkItemModifiedDate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.From.IsSet() {
		toSerialize["from"] = o.From.Get()
	}
	if o.To.IsSet() {
		toSerialize["to"] = o.To.Get()
	}
	return toSerialize, nil
}

type NullableTestPointFilterModelWorkItemModifiedDate struct {
	value *TestPointFilterModelWorkItemModifiedDate
	isSet bool
}

func (v NullableTestPointFilterModelWorkItemModifiedDate) Get() *TestPointFilterModelWorkItemModifiedDate {
	return v.value
}

func (v *NullableTestPointFilterModelWorkItemModifiedDate) Set(val *TestPointFilterModelWorkItemModifiedDate) {
	v.value = val
	v.isSet = true
}

func (v NullableTestPointFilterModelWorkItemModifiedDate) IsSet() bool {
	return v.isSet
}

func (v *NullableTestPointFilterModelWorkItemModifiedDate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestPointFilterModelWorkItemModifiedDate(val *TestPointFilterModelWorkItemModifiedDate) *NullableTestPointFilterModelWorkItemModifiedDate {
	return &NullableTestPointFilterModelWorkItemModifiedDate{value: val, isSet: true}
}

func (v NullableTestPointFilterModelWorkItemModifiedDate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTestPointFilterModelWorkItemModifiedDate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


