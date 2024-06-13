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
	"bytes"
	"fmt"
)

// checks if the TestPlanPutModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TestPlanPutModel{}

// TestPlanPutModel struct for TestPlanPutModel
type TestPlanPutModel struct {
	Id string `json:"id"`
	LockedById NullableString `json:"lockedById,omitempty"`
	Tags []TagPostModel `json:"tags,omitempty"`
	Name string `json:"name"`
	// Used for analytics
	StartDate NullableTime `json:"startDate,omitempty"`
	// Used for analytics
	EndDate NullableTime `json:"endDate,omitempty"`
	Description NullableString `json:"description,omitempty"`
	Build NullableString `json:"build,omitempty"`
	ProjectId string `json:"projectId"`
	ProductName NullableString `json:"productName,omitempty"`
	HasAutomaticDurationTimer NullableBool `json:"hasAutomaticDurationTimer,omitempty"`
	Attributes map[string]interface{} `json:"attributes"`
}

type _TestPlanPutModel TestPlanPutModel

// NewTestPlanPutModel instantiates a new TestPlanPutModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTestPlanPutModel(id string, name string, projectId string, attributes map[string]interface{}) *TestPlanPutModel {
	this := TestPlanPutModel{}
	this.Id = id
	this.Name = name
	this.ProjectId = projectId
	this.Attributes = attributes
	return &this
}

// NewTestPlanPutModelWithDefaults instantiates a new TestPlanPutModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTestPlanPutModelWithDefaults() *TestPlanPutModel {
	this := TestPlanPutModel{}
	return &this
}

// GetId returns the Id field value
func (o *TestPlanPutModel) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TestPlanPutModel) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TestPlanPutModel) SetId(v string) {
	o.Id = v
}

// GetLockedById returns the LockedById field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestPlanPutModel) GetLockedById() string {
	if o == nil || IsNil(o.LockedById.Get()) {
		var ret string
		return ret
	}
	return *o.LockedById.Get()
}

// GetLockedByIdOk returns a tuple with the LockedById field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestPlanPutModel) GetLockedByIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LockedById.Get(), o.LockedById.IsSet()
}

// HasLockedById returns a boolean if a field has been set.
func (o *TestPlanPutModel) HasLockedById() bool {
	if o != nil && o.LockedById.IsSet() {
		return true
	}

	return false
}

// SetLockedById gets a reference to the given NullableString and assigns it to the LockedById field.
func (o *TestPlanPutModel) SetLockedById(v string) {
	o.LockedById.Set(&v)
}
// SetLockedByIdNil sets the value for LockedById to be an explicit nil
func (o *TestPlanPutModel) SetLockedByIdNil() {
	o.LockedById.Set(nil)
}

// UnsetLockedById ensures that no value is present for LockedById, not even an explicit nil
func (o *TestPlanPutModel) UnsetLockedById() {
	o.LockedById.Unset()
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestPlanPutModel) GetTags() []TagPostModel {
	if o == nil {
		var ret []TagPostModel
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestPlanPutModel) GetTagsOk() ([]TagPostModel, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *TestPlanPutModel) HasTags() bool {
	if o != nil && IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []TagPostModel and assigns it to the Tags field.
func (o *TestPlanPutModel) SetTags(v []TagPostModel) {
	o.Tags = v
}

// GetName returns the Name field value
func (o *TestPlanPutModel) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TestPlanPutModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TestPlanPutModel) SetName(v string) {
	o.Name = v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestPlanPutModel) GetStartDate() time.Time {
	if o == nil || IsNil(o.StartDate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.StartDate.Get()
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestPlanPutModel) GetStartDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.StartDate.Get(), o.StartDate.IsSet()
}

// HasStartDate returns a boolean if a field has been set.
func (o *TestPlanPutModel) HasStartDate() bool {
	if o != nil && o.StartDate.IsSet() {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given NullableTime and assigns it to the StartDate field.
func (o *TestPlanPutModel) SetStartDate(v time.Time) {
	o.StartDate.Set(&v)
}
// SetStartDateNil sets the value for StartDate to be an explicit nil
func (o *TestPlanPutModel) SetStartDateNil() {
	o.StartDate.Set(nil)
}

// UnsetStartDate ensures that no value is present for StartDate, not even an explicit nil
func (o *TestPlanPutModel) UnsetStartDate() {
	o.StartDate.Unset()
}

// GetEndDate returns the EndDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestPlanPutModel) GetEndDate() time.Time {
	if o == nil || IsNil(o.EndDate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.EndDate.Get()
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestPlanPutModel) GetEndDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.EndDate.Get(), o.EndDate.IsSet()
}

// HasEndDate returns a boolean if a field has been set.
func (o *TestPlanPutModel) HasEndDate() bool {
	if o != nil && o.EndDate.IsSet() {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given NullableTime and assigns it to the EndDate field.
func (o *TestPlanPutModel) SetEndDate(v time.Time) {
	o.EndDate.Set(&v)
}
// SetEndDateNil sets the value for EndDate to be an explicit nil
func (o *TestPlanPutModel) SetEndDateNil() {
	o.EndDate.Set(nil)
}

// UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
func (o *TestPlanPutModel) UnsetEndDate() {
	o.EndDate.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestPlanPutModel) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestPlanPutModel) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *TestPlanPutModel) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *TestPlanPutModel) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *TestPlanPutModel) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *TestPlanPutModel) UnsetDescription() {
	o.Description.Unset()
}

// GetBuild returns the Build field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestPlanPutModel) GetBuild() string {
	if o == nil || IsNil(o.Build.Get()) {
		var ret string
		return ret
	}
	return *o.Build.Get()
}

// GetBuildOk returns a tuple with the Build field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestPlanPutModel) GetBuildOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Build.Get(), o.Build.IsSet()
}

// HasBuild returns a boolean if a field has been set.
func (o *TestPlanPutModel) HasBuild() bool {
	if o != nil && o.Build.IsSet() {
		return true
	}

	return false
}

// SetBuild gets a reference to the given NullableString and assigns it to the Build field.
func (o *TestPlanPutModel) SetBuild(v string) {
	o.Build.Set(&v)
}
// SetBuildNil sets the value for Build to be an explicit nil
func (o *TestPlanPutModel) SetBuildNil() {
	o.Build.Set(nil)
}

// UnsetBuild ensures that no value is present for Build, not even an explicit nil
func (o *TestPlanPutModel) UnsetBuild() {
	o.Build.Unset()
}

// GetProjectId returns the ProjectId field value
func (o *TestPlanPutModel) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *TestPlanPutModel) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *TestPlanPutModel) SetProjectId(v string) {
	o.ProjectId = v
}

// GetProductName returns the ProductName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestPlanPutModel) GetProductName() string {
	if o == nil || IsNil(o.ProductName.Get()) {
		var ret string
		return ret
	}
	return *o.ProductName.Get()
}

// GetProductNameOk returns a tuple with the ProductName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestPlanPutModel) GetProductNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProductName.Get(), o.ProductName.IsSet()
}

// HasProductName returns a boolean if a field has been set.
func (o *TestPlanPutModel) HasProductName() bool {
	if o != nil && o.ProductName.IsSet() {
		return true
	}

	return false
}

// SetProductName gets a reference to the given NullableString and assigns it to the ProductName field.
func (o *TestPlanPutModel) SetProductName(v string) {
	o.ProductName.Set(&v)
}
// SetProductNameNil sets the value for ProductName to be an explicit nil
func (o *TestPlanPutModel) SetProductNameNil() {
	o.ProductName.Set(nil)
}

// UnsetProductName ensures that no value is present for ProductName, not even an explicit nil
func (o *TestPlanPutModel) UnsetProductName() {
	o.ProductName.Unset()
}

// GetHasAutomaticDurationTimer returns the HasAutomaticDurationTimer field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestPlanPutModel) GetHasAutomaticDurationTimer() bool {
	if o == nil || IsNil(o.HasAutomaticDurationTimer.Get()) {
		var ret bool
		return ret
	}
	return *o.HasAutomaticDurationTimer.Get()
}

// GetHasAutomaticDurationTimerOk returns a tuple with the HasAutomaticDurationTimer field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestPlanPutModel) GetHasAutomaticDurationTimerOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.HasAutomaticDurationTimer.Get(), o.HasAutomaticDurationTimer.IsSet()
}

// HasHasAutomaticDurationTimer returns a boolean if a field has been set.
func (o *TestPlanPutModel) HasHasAutomaticDurationTimer() bool {
	if o != nil && o.HasAutomaticDurationTimer.IsSet() {
		return true
	}

	return false
}

// SetHasAutomaticDurationTimer gets a reference to the given NullableBool and assigns it to the HasAutomaticDurationTimer field.
func (o *TestPlanPutModel) SetHasAutomaticDurationTimer(v bool) {
	o.HasAutomaticDurationTimer.Set(&v)
}
// SetHasAutomaticDurationTimerNil sets the value for HasAutomaticDurationTimer to be an explicit nil
func (o *TestPlanPutModel) SetHasAutomaticDurationTimerNil() {
	o.HasAutomaticDurationTimer.Set(nil)
}

// UnsetHasAutomaticDurationTimer ensures that no value is present for HasAutomaticDurationTimer, not even an explicit nil
func (o *TestPlanPutModel) UnsetHasAutomaticDurationTimer() {
	o.HasAutomaticDurationTimer.Unset()
}

// GetAttributes returns the Attributes field value
func (o *TestPlanPutModel) GetAttributes() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *TestPlanPutModel) GetAttributesOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Attributes, true
}

// SetAttributes sets field value
func (o *TestPlanPutModel) SetAttributes(v map[string]interface{}) {
	o.Attributes = v
}

func (o TestPlanPutModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TestPlanPutModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if o.LockedById.IsSet() {
		toSerialize["lockedById"] = o.LockedById.Get()
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	toSerialize["name"] = o.Name
	if o.StartDate.IsSet() {
		toSerialize["startDate"] = o.StartDate.Get()
	}
	if o.EndDate.IsSet() {
		toSerialize["endDate"] = o.EndDate.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.Build.IsSet() {
		toSerialize["build"] = o.Build.Get()
	}
	toSerialize["projectId"] = o.ProjectId
	if o.ProductName.IsSet() {
		toSerialize["productName"] = o.ProductName.Get()
	}
	if o.HasAutomaticDurationTimer.IsSet() {
		toSerialize["hasAutomaticDurationTimer"] = o.HasAutomaticDurationTimer.Get()
	}
	toSerialize["attributes"] = o.Attributes
	return toSerialize, nil
}

func (o *TestPlanPutModel) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"projectId",
		"attributes",
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

	varTestPlanPutModel := _TestPlanPutModel{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTestPlanPutModel)

	if err != nil {
		return err
	}

	*o = TestPlanPutModel(varTestPlanPutModel)

	return err
}

type NullableTestPlanPutModel struct {
	value *TestPlanPutModel
	isSet bool
}

func (v NullableTestPlanPutModel) Get() *TestPlanPutModel {
	return v.value
}

func (v *NullableTestPlanPutModel) Set(val *TestPlanPutModel) {
	v.value = val
	v.isSet = true
}

func (v NullableTestPlanPutModel) IsSet() bool {
	return v.isSet
}

func (v *NullableTestPlanPutModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestPlanPutModel(val *TestPlanPutModel) *NullableTestPlanPutModel {
	return &NullableTestPlanPutModel{value: val, isSet: true}
}

func (v NullableTestPlanPutModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTestPlanPutModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


