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

// checks if the StepPutModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StepPutModel{}

// StepPutModel struct for StepPutModel
type StepPutModel struct {
	Id *string `json:"id,omitempty"`
	Action NullableString `json:"action,omitempty"`
	Expected NullableString `json:"expected,omitempty"`
	TestData NullableString `json:"testData,omitempty"`
	Comments NullableString `json:"comments,omitempty"`
	WorkItemId NullableString `json:"workItemId,omitempty"`
}

// NewStepPutModel instantiates a new StepPutModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStepPutModel() *StepPutModel {
	this := StepPutModel{}
	return &this
}

// NewStepPutModelWithDefaults instantiates a new StepPutModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStepPutModelWithDefaults() *StepPutModel {
	this := StepPutModel{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *StepPutModel) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StepPutModel) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *StepPutModel) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *StepPutModel) SetId(v string) {
	o.Id = &v
}

// GetAction returns the Action field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StepPutModel) GetAction() string {
	if o == nil || IsNil(o.Action.Get()) {
		var ret string
		return ret
	}
	return *o.Action.Get()
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StepPutModel) GetActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Action.Get(), o.Action.IsSet()
}

// HasAction returns a boolean if a field has been set.
func (o *StepPutModel) HasAction() bool {
	if o != nil && o.Action.IsSet() {
		return true
	}

	return false
}

// SetAction gets a reference to the given NullableString and assigns it to the Action field.
func (o *StepPutModel) SetAction(v string) {
	o.Action.Set(&v)
}
// SetActionNil sets the value for Action to be an explicit nil
func (o *StepPutModel) SetActionNil() {
	o.Action.Set(nil)
}

// UnsetAction ensures that no value is present for Action, not even an explicit nil
func (o *StepPutModel) UnsetAction() {
	o.Action.Unset()
}

// GetExpected returns the Expected field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StepPutModel) GetExpected() string {
	if o == nil || IsNil(o.Expected.Get()) {
		var ret string
		return ret
	}
	return *o.Expected.Get()
}

// GetExpectedOk returns a tuple with the Expected field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StepPutModel) GetExpectedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Expected.Get(), o.Expected.IsSet()
}

// HasExpected returns a boolean if a field has been set.
func (o *StepPutModel) HasExpected() bool {
	if o != nil && o.Expected.IsSet() {
		return true
	}

	return false
}

// SetExpected gets a reference to the given NullableString and assigns it to the Expected field.
func (o *StepPutModel) SetExpected(v string) {
	o.Expected.Set(&v)
}
// SetExpectedNil sets the value for Expected to be an explicit nil
func (o *StepPutModel) SetExpectedNil() {
	o.Expected.Set(nil)
}

// UnsetExpected ensures that no value is present for Expected, not even an explicit nil
func (o *StepPutModel) UnsetExpected() {
	o.Expected.Unset()
}

// GetTestData returns the TestData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StepPutModel) GetTestData() string {
	if o == nil || IsNil(o.TestData.Get()) {
		var ret string
		return ret
	}
	return *o.TestData.Get()
}

// GetTestDataOk returns a tuple with the TestData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StepPutModel) GetTestDataOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TestData.Get(), o.TestData.IsSet()
}

// HasTestData returns a boolean if a field has been set.
func (o *StepPutModel) HasTestData() bool {
	if o != nil && o.TestData.IsSet() {
		return true
	}

	return false
}

// SetTestData gets a reference to the given NullableString and assigns it to the TestData field.
func (o *StepPutModel) SetTestData(v string) {
	o.TestData.Set(&v)
}
// SetTestDataNil sets the value for TestData to be an explicit nil
func (o *StepPutModel) SetTestDataNil() {
	o.TestData.Set(nil)
}

// UnsetTestData ensures that no value is present for TestData, not even an explicit nil
func (o *StepPutModel) UnsetTestData() {
	o.TestData.Unset()
}

// GetComments returns the Comments field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StepPutModel) GetComments() string {
	if o == nil || IsNil(o.Comments.Get()) {
		var ret string
		return ret
	}
	return *o.Comments.Get()
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StepPutModel) GetCommentsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Comments.Get(), o.Comments.IsSet()
}

// HasComments returns a boolean if a field has been set.
func (o *StepPutModel) HasComments() bool {
	if o != nil && o.Comments.IsSet() {
		return true
	}

	return false
}

// SetComments gets a reference to the given NullableString and assigns it to the Comments field.
func (o *StepPutModel) SetComments(v string) {
	o.Comments.Set(&v)
}
// SetCommentsNil sets the value for Comments to be an explicit nil
func (o *StepPutModel) SetCommentsNil() {
	o.Comments.Set(nil)
}

// UnsetComments ensures that no value is present for Comments, not even an explicit nil
func (o *StepPutModel) UnsetComments() {
	o.Comments.Unset()
}

// GetWorkItemId returns the WorkItemId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StepPutModel) GetWorkItemId() string {
	if o == nil || IsNil(o.WorkItemId.Get()) {
		var ret string
		return ret
	}
	return *o.WorkItemId.Get()
}

// GetWorkItemIdOk returns a tuple with the WorkItemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StepPutModel) GetWorkItemIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.WorkItemId.Get(), o.WorkItemId.IsSet()
}

// HasWorkItemId returns a boolean if a field has been set.
func (o *StepPutModel) HasWorkItemId() bool {
	if o != nil && o.WorkItemId.IsSet() {
		return true
	}

	return false
}

// SetWorkItemId gets a reference to the given NullableString and assigns it to the WorkItemId field.
func (o *StepPutModel) SetWorkItemId(v string) {
	o.WorkItemId.Set(&v)
}
// SetWorkItemIdNil sets the value for WorkItemId to be an explicit nil
func (o *StepPutModel) SetWorkItemIdNil() {
	o.WorkItemId.Set(nil)
}

// UnsetWorkItemId ensures that no value is present for WorkItemId, not even an explicit nil
func (o *StepPutModel) UnsetWorkItemId() {
	o.WorkItemId.Unset()
}

func (o StepPutModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StepPutModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Action.IsSet() {
		toSerialize["action"] = o.Action.Get()
	}
	if o.Expected.IsSet() {
		toSerialize["expected"] = o.Expected.Get()
	}
	if o.TestData.IsSet() {
		toSerialize["testData"] = o.TestData.Get()
	}
	if o.Comments.IsSet() {
		toSerialize["comments"] = o.Comments.Get()
	}
	if o.WorkItemId.IsSet() {
		toSerialize["workItemId"] = o.WorkItemId.Get()
	}
	return toSerialize, nil
}

type NullableStepPutModel struct {
	value *StepPutModel
	isSet bool
}

func (v NullableStepPutModel) Get() *StepPutModel {
	return v.value
}

func (v *NullableStepPutModel) Set(val *StepPutModel) {
	v.value = val
	v.isSet = true
}

func (v NullableStepPutModel) IsSet() bool {
	return v.isSet
}

func (v *NullableStepPutModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStepPutModel(val *StepPutModel) *NullableStepPutModel {
	return &NullableStepPutModel{value: val, isSet: true}
}

func (v NullableStepPutModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStepPutModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


