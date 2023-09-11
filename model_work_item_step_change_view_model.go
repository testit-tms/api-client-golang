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

// checks if the WorkItemStepChangeViewModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WorkItemStepChangeViewModel{}

// WorkItemStepChangeViewModel struct for WorkItemStepChangeViewModel
type WorkItemStepChangeViewModel struct {
	Action NullableString `json:"action,omitempty"`
	Expected NullableString `json:"expected,omitempty"`
	Comments NullableString `json:"comments,omitempty"`
	TestData NullableString `json:"testData,omitempty"`
	Index int32 `json:"index"`
	WorkItemId NullableString `json:"workItemId,omitempty"`
	WorkItem NullableSharedStepChangeViewModel `json:"workItem,omitempty"`
}

// NewWorkItemStepChangeViewModel instantiates a new WorkItemStepChangeViewModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkItemStepChangeViewModel(index int32) *WorkItemStepChangeViewModel {
	this := WorkItemStepChangeViewModel{}
	this.Index = index
	return &this
}

// NewWorkItemStepChangeViewModelWithDefaults instantiates a new WorkItemStepChangeViewModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkItemStepChangeViewModelWithDefaults() *WorkItemStepChangeViewModel {
	this := WorkItemStepChangeViewModel{}
	return &this
}

// GetAction returns the Action field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkItemStepChangeViewModel) GetAction() string {
	if o == nil || IsNil(o.Action.Get()) {
		var ret string
		return ret
	}
	return *o.Action.Get()
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkItemStepChangeViewModel) GetActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Action.Get(), o.Action.IsSet()
}

// HasAction returns a boolean if a field has been set.
func (o *WorkItemStepChangeViewModel) HasAction() bool {
	if o != nil && o.Action.IsSet() {
		return true
	}

	return false
}

// SetAction gets a reference to the given NullableString and assigns it to the Action field.
func (o *WorkItemStepChangeViewModel) SetAction(v string) {
	o.Action.Set(&v)
}
// SetActionNil sets the value for Action to be an explicit nil
func (o *WorkItemStepChangeViewModel) SetActionNil() {
	o.Action.Set(nil)
}

// UnsetAction ensures that no value is present for Action, not even an explicit nil
func (o *WorkItemStepChangeViewModel) UnsetAction() {
	o.Action.Unset()
}

// GetExpected returns the Expected field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkItemStepChangeViewModel) GetExpected() string {
	if o == nil || IsNil(o.Expected.Get()) {
		var ret string
		return ret
	}
	return *o.Expected.Get()
}

// GetExpectedOk returns a tuple with the Expected field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkItemStepChangeViewModel) GetExpectedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Expected.Get(), o.Expected.IsSet()
}

// HasExpected returns a boolean if a field has been set.
func (o *WorkItemStepChangeViewModel) HasExpected() bool {
	if o != nil && o.Expected.IsSet() {
		return true
	}

	return false
}

// SetExpected gets a reference to the given NullableString and assigns it to the Expected field.
func (o *WorkItemStepChangeViewModel) SetExpected(v string) {
	o.Expected.Set(&v)
}
// SetExpectedNil sets the value for Expected to be an explicit nil
func (o *WorkItemStepChangeViewModel) SetExpectedNil() {
	o.Expected.Set(nil)
}

// UnsetExpected ensures that no value is present for Expected, not even an explicit nil
func (o *WorkItemStepChangeViewModel) UnsetExpected() {
	o.Expected.Unset()
}

// GetComments returns the Comments field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkItemStepChangeViewModel) GetComments() string {
	if o == nil || IsNil(o.Comments.Get()) {
		var ret string
		return ret
	}
	return *o.Comments.Get()
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkItemStepChangeViewModel) GetCommentsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Comments.Get(), o.Comments.IsSet()
}

// HasComments returns a boolean if a field has been set.
func (o *WorkItemStepChangeViewModel) HasComments() bool {
	if o != nil && o.Comments.IsSet() {
		return true
	}

	return false
}

// SetComments gets a reference to the given NullableString and assigns it to the Comments field.
func (o *WorkItemStepChangeViewModel) SetComments(v string) {
	o.Comments.Set(&v)
}
// SetCommentsNil sets the value for Comments to be an explicit nil
func (o *WorkItemStepChangeViewModel) SetCommentsNil() {
	o.Comments.Set(nil)
}

// UnsetComments ensures that no value is present for Comments, not even an explicit nil
func (o *WorkItemStepChangeViewModel) UnsetComments() {
	o.Comments.Unset()
}

// GetTestData returns the TestData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkItemStepChangeViewModel) GetTestData() string {
	if o == nil || IsNil(o.TestData.Get()) {
		var ret string
		return ret
	}
	return *o.TestData.Get()
}

// GetTestDataOk returns a tuple with the TestData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkItemStepChangeViewModel) GetTestDataOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TestData.Get(), o.TestData.IsSet()
}

// HasTestData returns a boolean if a field has been set.
func (o *WorkItemStepChangeViewModel) HasTestData() bool {
	if o != nil && o.TestData.IsSet() {
		return true
	}

	return false
}

// SetTestData gets a reference to the given NullableString and assigns it to the TestData field.
func (o *WorkItemStepChangeViewModel) SetTestData(v string) {
	o.TestData.Set(&v)
}
// SetTestDataNil sets the value for TestData to be an explicit nil
func (o *WorkItemStepChangeViewModel) SetTestDataNil() {
	o.TestData.Set(nil)
}

// UnsetTestData ensures that no value is present for TestData, not even an explicit nil
func (o *WorkItemStepChangeViewModel) UnsetTestData() {
	o.TestData.Unset()
}

// GetIndex returns the Index field value
func (o *WorkItemStepChangeViewModel) GetIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *WorkItemStepChangeViewModel) GetIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *WorkItemStepChangeViewModel) SetIndex(v int32) {
	o.Index = v
}

// GetWorkItemId returns the WorkItemId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkItemStepChangeViewModel) GetWorkItemId() string {
	if o == nil || IsNil(o.WorkItemId.Get()) {
		var ret string
		return ret
	}
	return *o.WorkItemId.Get()
}

// GetWorkItemIdOk returns a tuple with the WorkItemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkItemStepChangeViewModel) GetWorkItemIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.WorkItemId.Get(), o.WorkItemId.IsSet()
}

// HasWorkItemId returns a boolean if a field has been set.
func (o *WorkItemStepChangeViewModel) HasWorkItemId() bool {
	if o != nil && o.WorkItemId.IsSet() {
		return true
	}

	return false
}

// SetWorkItemId gets a reference to the given NullableString and assigns it to the WorkItemId field.
func (o *WorkItemStepChangeViewModel) SetWorkItemId(v string) {
	o.WorkItemId.Set(&v)
}
// SetWorkItemIdNil sets the value for WorkItemId to be an explicit nil
func (o *WorkItemStepChangeViewModel) SetWorkItemIdNil() {
	o.WorkItemId.Set(nil)
}

// UnsetWorkItemId ensures that no value is present for WorkItemId, not even an explicit nil
func (o *WorkItemStepChangeViewModel) UnsetWorkItemId() {
	o.WorkItemId.Unset()
}

// GetWorkItem returns the WorkItem field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkItemStepChangeViewModel) GetWorkItem() SharedStepChangeViewModel {
	if o == nil || IsNil(o.WorkItem.Get()) {
		var ret SharedStepChangeViewModel
		return ret
	}
	return *o.WorkItem.Get()
}

// GetWorkItemOk returns a tuple with the WorkItem field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkItemStepChangeViewModel) GetWorkItemOk() (*SharedStepChangeViewModel, bool) {
	if o == nil {
		return nil, false
	}
	return o.WorkItem.Get(), o.WorkItem.IsSet()
}

// HasWorkItem returns a boolean if a field has been set.
func (o *WorkItemStepChangeViewModel) HasWorkItem() bool {
	if o != nil && o.WorkItem.IsSet() {
		return true
	}

	return false
}

// SetWorkItem gets a reference to the given NullableSharedStepChangeViewModel and assigns it to the WorkItem field.
func (o *WorkItemStepChangeViewModel) SetWorkItem(v SharedStepChangeViewModel) {
	o.WorkItem.Set(&v)
}
// SetWorkItemNil sets the value for WorkItem to be an explicit nil
func (o *WorkItemStepChangeViewModel) SetWorkItemNil() {
	o.WorkItem.Set(nil)
}

// UnsetWorkItem ensures that no value is present for WorkItem, not even an explicit nil
func (o *WorkItemStepChangeViewModel) UnsetWorkItem() {
	o.WorkItem.Unset()
}

func (o WorkItemStepChangeViewModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkItemStepChangeViewModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Action.IsSet() {
		toSerialize["action"] = o.Action.Get()
	}
	if o.Expected.IsSet() {
		toSerialize["expected"] = o.Expected.Get()
	}
	if o.Comments.IsSet() {
		toSerialize["comments"] = o.Comments.Get()
	}
	if o.TestData.IsSet() {
		toSerialize["testData"] = o.TestData.Get()
	}
	toSerialize["index"] = o.Index
	if o.WorkItemId.IsSet() {
		toSerialize["workItemId"] = o.WorkItemId.Get()
	}
	if o.WorkItem.IsSet() {
		toSerialize["workItem"] = o.WorkItem.Get()
	}
	return toSerialize, nil
}

type NullableWorkItemStepChangeViewModel struct {
	value *WorkItemStepChangeViewModel
	isSet bool
}

func (v NullableWorkItemStepChangeViewModel) Get() *WorkItemStepChangeViewModel {
	return v.value
}

func (v *NullableWorkItemStepChangeViewModel) Set(val *WorkItemStepChangeViewModel) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkItemStepChangeViewModel) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkItemStepChangeViewModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkItemStepChangeViewModel(val *WorkItemStepChangeViewModel) *NullableWorkItemStepChangeViewModel {
	return &NullableWorkItemStepChangeViewModel{value: val, isSet: true}
}

func (v NullableWorkItemStepChangeViewModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkItemStepChangeViewModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


