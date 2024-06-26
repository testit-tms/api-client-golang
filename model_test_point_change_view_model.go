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

// checks if the TestPointChangeViewModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TestPointChangeViewModel{}

// TestPointChangeViewModel struct for TestPointChangeViewModel
type TestPointChangeViewModel struct {
	UserId string `json:"userId"`
	UserName NullableString `json:"userName,omitempty"`
	TestPointCount int64 `json:"testPointCount"`
}

type _TestPointChangeViewModel TestPointChangeViewModel

// NewTestPointChangeViewModel instantiates a new TestPointChangeViewModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTestPointChangeViewModel(userId string, testPointCount int64) *TestPointChangeViewModel {
	this := TestPointChangeViewModel{}
	this.UserId = userId
	this.TestPointCount = testPointCount
	return &this
}

// NewTestPointChangeViewModelWithDefaults instantiates a new TestPointChangeViewModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTestPointChangeViewModelWithDefaults() *TestPointChangeViewModel {
	this := TestPointChangeViewModel{}
	return &this
}

// GetUserId returns the UserId field value
func (o *TestPointChangeViewModel) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *TestPointChangeViewModel) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *TestPointChangeViewModel) SetUserId(v string) {
	o.UserId = v
}

// GetUserName returns the UserName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestPointChangeViewModel) GetUserName() string {
	if o == nil || IsNil(o.UserName.Get()) {
		var ret string
		return ret
	}
	return *o.UserName.Get()
}

// GetUserNameOk returns a tuple with the UserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestPointChangeViewModel) GetUserNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserName.Get(), o.UserName.IsSet()
}

// HasUserName returns a boolean if a field has been set.
func (o *TestPointChangeViewModel) HasUserName() bool {
	if o != nil && o.UserName.IsSet() {
		return true
	}

	return false
}

// SetUserName gets a reference to the given NullableString and assigns it to the UserName field.
func (o *TestPointChangeViewModel) SetUserName(v string) {
	o.UserName.Set(&v)
}
// SetUserNameNil sets the value for UserName to be an explicit nil
func (o *TestPointChangeViewModel) SetUserNameNil() {
	o.UserName.Set(nil)
}

// UnsetUserName ensures that no value is present for UserName, not even an explicit nil
func (o *TestPointChangeViewModel) UnsetUserName() {
	o.UserName.Unset()
}

// GetTestPointCount returns the TestPointCount field value
func (o *TestPointChangeViewModel) GetTestPointCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.TestPointCount
}

// GetTestPointCountOk returns a tuple with the TestPointCount field value
// and a boolean to check if the value has been set.
func (o *TestPointChangeViewModel) GetTestPointCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TestPointCount, true
}

// SetTestPointCount sets field value
func (o *TestPointChangeViewModel) SetTestPointCount(v int64) {
	o.TestPointCount = v
}

func (o TestPointChangeViewModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TestPointChangeViewModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["userId"] = o.UserId
	if o.UserName.IsSet() {
		toSerialize["userName"] = o.UserName.Get()
	}
	toSerialize["testPointCount"] = o.TestPointCount
	return toSerialize, nil
}

func (o *TestPointChangeViewModel) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"userId",
		"testPointCount",
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

	varTestPointChangeViewModel := _TestPointChangeViewModel{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTestPointChangeViewModel)

	if err != nil {
		return err
	}

	*o = TestPointChangeViewModel(varTestPointChangeViewModel)

	return err
}

type NullableTestPointChangeViewModel struct {
	value *TestPointChangeViewModel
	isSet bool
}

func (v NullableTestPointChangeViewModel) Get() *TestPointChangeViewModel {
	return v.value
}

func (v *NullableTestPointChangeViewModel) Set(val *TestPointChangeViewModel) {
	v.value = val
	v.isSet = true
}

func (v NullableTestPointChangeViewModel) IsSet() bool {
	return v.isSet
}

func (v *NullableTestPointChangeViewModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestPointChangeViewModel(val *TestPointChangeViewModel) *NullableTestPointChangeViewModel {
	return &NullableTestPointChangeViewModel{value: val, isSet: true}
}

func (v NullableTestPointChangeViewModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTestPointChangeViewModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


