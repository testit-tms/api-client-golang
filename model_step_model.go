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

// checks if the StepModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StepModel{}

// StepModel struct for StepModel
type StepModel struct {
	// Nested shared steps are allowed
	WorkItem NullableSharedStepModel `json:"workItem,omitempty"`
	Id string `json:"id"`
	Action NullableString `json:"action,omitempty"`
	Expected NullableString `json:"expected,omitempty"`
	TestData NullableString `json:"testData,omitempty"`
	Comments NullableString `json:"comments,omitempty"`
	WorkItemId NullableString `json:"workItemId,omitempty"`
}

type _StepModel StepModel

// NewStepModel instantiates a new StepModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStepModel(id string) *StepModel {
	this := StepModel{}
	this.Id = id
	return &this
}

// NewStepModelWithDefaults instantiates a new StepModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStepModelWithDefaults() *StepModel {
	this := StepModel{}
	return &this
}

// GetWorkItem returns the WorkItem field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StepModel) GetWorkItem() SharedStepModel {
	if o == nil || IsNil(o.WorkItem.Get()) {
		var ret SharedStepModel
		return ret
	}
	return *o.WorkItem.Get()
}

// GetWorkItemOk returns a tuple with the WorkItem field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StepModel) GetWorkItemOk() (*SharedStepModel, bool) {
	if o == nil {
		return nil, false
	}
	return o.WorkItem.Get(), o.WorkItem.IsSet()
}

// HasWorkItem returns a boolean if a field has been set.
func (o *StepModel) HasWorkItem() bool {
	if o != nil && o.WorkItem.IsSet() {
		return true
	}

	return false
}

// SetWorkItem gets a reference to the given NullableSharedStepModel and assigns it to the WorkItem field.
func (o *StepModel) SetWorkItem(v SharedStepModel) {
	o.WorkItem.Set(&v)
}
// SetWorkItemNil sets the value for WorkItem to be an explicit nil
func (o *StepModel) SetWorkItemNil() {
	o.WorkItem.Set(nil)
}

// UnsetWorkItem ensures that no value is present for WorkItem, not even an explicit nil
func (o *StepModel) UnsetWorkItem() {
	o.WorkItem.Unset()
}

// GetId returns the Id field value
func (o *StepModel) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *StepModel) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *StepModel) SetId(v string) {
	o.Id = v
}

// GetAction returns the Action field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StepModel) GetAction() string {
	if o == nil || IsNil(o.Action.Get()) {
		var ret string
		return ret
	}
	return *o.Action.Get()
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StepModel) GetActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Action.Get(), o.Action.IsSet()
}

// HasAction returns a boolean if a field has been set.
func (o *StepModel) HasAction() bool {
	if o != nil && o.Action.IsSet() {
		return true
	}

	return false
}

// SetAction gets a reference to the given NullableString and assigns it to the Action field.
func (o *StepModel) SetAction(v string) {
	o.Action.Set(&v)
}
// SetActionNil sets the value for Action to be an explicit nil
func (o *StepModel) SetActionNil() {
	o.Action.Set(nil)
}

// UnsetAction ensures that no value is present for Action, not even an explicit nil
func (o *StepModel) UnsetAction() {
	o.Action.Unset()
}

// GetExpected returns the Expected field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StepModel) GetExpected() string {
	if o == nil || IsNil(o.Expected.Get()) {
		var ret string
		return ret
	}
	return *o.Expected.Get()
}

// GetExpectedOk returns a tuple with the Expected field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StepModel) GetExpectedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Expected.Get(), o.Expected.IsSet()
}

// HasExpected returns a boolean if a field has been set.
func (o *StepModel) HasExpected() bool {
	if o != nil && o.Expected.IsSet() {
		return true
	}

	return false
}

// SetExpected gets a reference to the given NullableString and assigns it to the Expected field.
func (o *StepModel) SetExpected(v string) {
	o.Expected.Set(&v)
}
// SetExpectedNil sets the value for Expected to be an explicit nil
func (o *StepModel) SetExpectedNil() {
	o.Expected.Set(nil)
}

// UnsetExpected ensures that no value is present for Expected, not even an explicit nil
func (o *StepModel) UnsetExpected() {
	o.Expected.Unset()
}

// GetTestData returns the TestData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StepModel) GetTestData() string {
	if o == nil || IsNil(o.TestData.Get()) {
		var ret string
		return ret
	}
	return *o.TestData.Get()
}

// GetTestDataOk returns a tuple with the TestData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StepModel) GetTestDataOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TestData.Get(), o.TestData.IsSet()
}

// HasTestData returns a boolean if a field has been set.
func (o *StepModel) HasTestData() bool {
	if o != nil && o.TestData.IsSet() {
		return true
	}

	return false
}

// SetTestData gets a reference to the given NullableString and assigns it to the TestData field.
func (o *StepModel) SetTestData(v string) {
	o.TestData.Set(&v)
}
// SetTestDataNil sets the value for TestData to be an explicit nil
func (o *StepModel) SetTestDataNil() {
	o.TestData.Set(nil)
}

// UnsetTestData ensures that no value is present for TestData, not even an explicit nil
func (o *StepModel) UnsetTestData() {
	o.TestData.Unset()
}

// GetComments returns the Comments field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StepModel) GetComments() string {
	if o == nil || IsNil(o.Comments.Get()) {
		var ret string
		return ret
	}
	return *o.Comments.Get()
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StepModel) GetCommentsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Comments.Get(), o.Comments.IsSet()
}

// HasComments returns a boolean if a field has been set.
func (o *StepModel) HasComments() bool {
	if o != nil && o.Comments.IsSet() {
		return true
	}

	return false
}

// SetComments gets a reference to the given NullableString and assigns it to the Comments field.
func (o *StepModel) SetComments(v string) {
	o.Comments.Set(&v)
}
// SetCommentsNil sets the value for Comments to be an explicit nil
func (o *StepModel) SetCommentsNil() {
	o.Comments.Set(nil)
}

// UnsetComments ensures that no value is present for Comments, not even an explicit nil
func (o *StepModel) UnsetComments() {
	o.Comments.Unset()
}

// GetWorkItemId returns the WorkItemId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StepModel) GetWorkItemId() string {
	if o == nil || IsNil(o.WorkItemId.Get()) {
		var ret string
		return ret
	}
	return *o.WorkItemId.Get()
}

// GetWorkItemIdOk returns a tuple with the WorkItemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StepModel) GetWorkItemIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.WorkItemId.Get(), o.WorkItemId.IsSet()
}

// HasWorkItemId returns a boolean if a field has been set.
func (o *StepModel) HasWorkItemId() bool {
	if o != nil && o.WorkItemId.IsSet() {
		return true
	}

	return false
}

// SetWorkItemId gets a reference to the given NullableString and assigns it to the WorkItemId field.
func (o *StepModel) SetWorkItemId(v string) {
	o.WorkItemId.Set(&v)
}
// SetWorkItemIdNil sets the value for WorkItemId to be an explicit nil
func (o *StepModel) SetWorkItemIdNil() {
	o.WorkItemId.Set(nil)
}

// UnsetWorkItemId ensures that no value is present for WorkItemId, not even an explicit nil
func (o *StepModel) UnsetWorkItemId() {
	o.WorkItemId.Unset()
}

func (o StepModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StepModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.WorkItem.IsSet() {
		toSerialize["workItem"] = o.WorkItem.Get()
	}
	toSerialize["id"] = o.Id
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

func (o *StepModel) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varStepModel := _StepModel{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varStepModel)

	if err != nil {
		return err
	}

	*o = StepModel(varStepModel)

	return err
}

type NullableStepModel struct {
	value *StepModel
	isSet bool
}

func (v NullableStepModel) Get() *StepModel {
	return v.value
}

func (v *NullableStepModel) Set(val *StepModel) {
	v.value = val
	v.isSet = true
}

func (v NullableStepModel) IsSet() bool {
	return v.isSet
}

func (v *NullableStepModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStepModel(val *StepModel) *NullableStepModel {
	return &NullableStepModel{value: val, isSet: true}
}

func (v NullableStepModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStepModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


