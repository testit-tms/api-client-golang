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

// checks if the UpdateTestPlanApiModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateTestPlanApiModel{}

// UpdateTestPlanApiModel struct for UpdateTestPlanApiModel
type UpdateTestPlanApiModel struct {
	// Test plan unique internal identifier
	Id string `json:"id"`
	// User who locked test plan modification internal identifier
	LockedById NullableString `json:"lockedById,omitempty"`
	// Test plan name
	Name string `json:"name"`
	// Date and time of test plan start
	StartDate NullableTime `json:"startDate,omitempty"`
	// Date and time of test plan end
	EndDate NullableTime `json:"endDate,omitempty"`
	// Test plan description
	Description NullableString `json:"description,omitempty"`
	// Build of the application on which test plan is executed
	Build NullableString `json:"build,omitempty"`
	// Project unique identifier
	ProjectId string `json:"projectId"`
	// Name of the testing product
	ProductName NullableString `json:"productName,omitempty"`
	// Boolean flag defines if test plan has automatic duration timer
	HasAutomaticDurationTimer NullableBool `json:"hasAutomaticDurationTimer,omitempty"`
	// Key value pair of test plan custom attributes
	Attributes map[string]interface{} `json:"attributes,omitempty"`
	// Test plan tag names collection
	Tags []TagApiModel `json:"tags,omitempty"`
}

type _UpdateTestPlanApiModel UpdateTestPlanApiModel

// NewUpdateTestPlanApiModel instantiates a new UpdateTestPlanApiModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateTestPlanApiModel(id string, name string, projectId string) *UpdateTestPlanApiModel {
	this := UpdateTestPlanApiModel{}
	this.Id = id
	this.Name = name
	this.ProjectId = projectId
	return &this
}

// NewUpdateTestPlanApiModelWithDefaults instantiates a new UpdateTestPlanApiModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateTestPlanApiModelWithDefaults() *UpdateTestPlanApiModel {
	this := UpdateTestPlanApiModel{}
	return &this
}

// GetId returns the Id field value
func (o *UpdateTestPlanApiModel) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *UpdateTestPlanApiModel) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *UpdateTestPlanApiModel) SetId(v string) {
	o.Id = v
}

// GetLockedById returns the LockedById field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateTestPlanApiModel) GetLockedById() string {
	if o == nil || IsNil(o.LockedById.Get()) {
		var ret string
		return ret
	}
	return *o.LockedById.Get()
}

// GetLockedByIdOk returns a tuple with the LockedById field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateTestPlanApiModel) GetLockedByIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LockedById.Get(), o.LockedById.IsSet()
}

// HasLockedById returns a boolean if a field has been set.
func (o *UpdateTestPlanApiModel) HasLockedById() bool {
	if o != nil && o.LockedById.IsSet() {
		return true
	}

	return false
}

// SetLockedById gets a reference to the given NullableString and assigns it to the LockedById field.
func (o *UpdateTestPlanApiModel) SetLockedById(v string) {
	o.LockedById.Set(&v)
}
// SetLockedByIdNil sets the value for LockedById to be an explicit nil
func (o *UpdateTestPlanApiModel) SetLockedByIdNil() {
	o.LockedById.Set(nil)
}

// UnsetLockedById ensures that no value is present for LockedById, not even an explicit nil
func (o *UpdateTestPlanApiModel) UnsetLockedById() {
	o.LockedById.Unset()
}

// GetName returns the Name field value
func (o *UpdateTestPlanApiModel) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UpdateTestPlanApiModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UpdateTestPlanApiModel) SetName(v string) {
	o.Name = v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateTestPlanApiModel) GetStartDate() time.Time {
	if o == nil || IsNil(o.StartDate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.StartDate.Get()
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateTestPlanApiModel) GetStartDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.StartDate.Get(), o.StartDate.IsSet()
}

// HasStartDate returns a boolean if a field has been set.
func (o *UpdateTestPlanApiModel) HasStartDate() bool {
	if o != nil && o.StartDate.IsSet() {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given NullableTime and assigns it to the StartDate field.
func (o *UpdateTestPlanApiModel) SetStartDate(v time.Time) {
	o.StartDate.Set(&v)
}
// SetStartDateNil sets the value for StartDate to be an explicit nil
func (o *UpdateTestPlanApiModel) SetStartDateNil() {
	o.StartDate.Set(nil)
}

// UnsetStartDate ensures that no value is present for StartDate, not even an explicit nil
func (o *UpdateTestPlanApiModel) UnsetStartDate() {
	o.StartDate.Unset()
}

// GetEndDate returns the EndDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateTestPlanApiModel) GetEndDate() time.Time {
	if o == nil || IsNil(o.EndDate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.EndDate.Get()
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateTestPlanApiModel) GetEndDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.EndDate.Get(), o.EndDate.IsSet()
}

// HasEndDate returns a boolean if a field has been set.
func (o *UpdateTestPlanApiModel) HasEndDate() bool {
	if o != nil && o.EndDate.IsSet() {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given NullableTime and assigns it to the EndDate field.
func (o *UpdateTestPlanApiModel) SetEndDate(v time.Time) {
	o.EndDate.Set(&v)
}
// SetEndDateNil sets the value for EndDate to be an explicit nil
func (o *UpdateTestPlanApiModel) SetEndDateNil() {
	o.EndDate.Set(nil)
}

// UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
func (o *UpdateTestPlanApiModel) UnsetEndDate() {
	o.EndDate.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateTestPlanApiModel) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateTestPlanApiModel) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *UpdateTestPlanApiModel) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *UpdateTestPlanApiModel) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *UpdateTestPlanApiModel) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *UpdateTestPlanApiModel) UnsetDescription() {
	o.Description.Unset()
}

// GetBuild returns the Build field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateTestPlanApiModel) GetBuild() string {
	if o == nil || IsNil(o.Build.Get()) {
		var ret string
		return ret
	}
	return *o.Build.Get()
}

// GetBuildOk returns a tuple with the Build field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateTestPlanApiModel) GetBuildOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Build.Get(), o.Build.IsSet()
}

// HasBuild returns a boolean if a field has been set.
func (o *UpdateTestPlanApiModel) HasBuild() bool {
	if o != nil && o.Build.IsSet() {
		return true
	}

	return false
}

// SetBuild gets a reference to the given NullableString and assigns it to the Build field.
func (o *UpdateTestPlanApiModel) SetBuild(v string) {
	o.Build.Set(&v)
}
// SetBuildNil sets the value for Build to be an explicit nil
func (o *UpdateTestPlanApiModel) SetBuildNil() {
	o.Build.Set(nil)
}

// UnsetBuild ensures that no value is present for Build, not even an explicit nil
func (o *UpdateTestPlanApiModel) UnsetBuild() {
	o.Build.Unset()
}

// GetProjectId returns the ProjectId field value
func (o *UpdateTestPlanApiModel) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *UpdateTestPlanApiModel) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *UpdateTestPlanApiModel) SetProjectId(v string) {
	o.ProjectId = v
}

// GetProductName returns the ProductName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateTestPlanApiModel) GetProductName() string {
	if o == nil || IsNil(o.ProductName.Get()) {
		var ret string
		return ret
	}
	return *o.ProductName.Get()
}

// GetProductNameOk returns a tuple with the ProductName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateTestPlanApiModel) GetProductNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProductName.Get(), o.ProductName.IsSet()
}

// HasProductName returns a boolean if a field has been set.
func (o *UpdateTestPlanApiModel) HasProductName() bool {
	if o != nil && o.ProductName.IsSet() {
		return true
	}

	return false
}

// SetProductName gets a reference to the given NullableString and assigns it to the ProductName field.
func (o *UpdateTestPlanApiModel) SetProductName(v string) {
	o.ProductName.Set(&v)
}
// SetProductNameNil sets the value for ProductName to be an explicit nil
func (o *UpdateTestPlanApiModel) SetProductNameNil() {
	o.ProductName.Set(nil)
}

// UnsetProductName ensures that no value is present for ProductName, not even an explicit nil
func (o *UpdateTestPlanApiModel) UnsetProductName() {
	o.ProductName.Unset()
}

// GetHasAutomaticDurationTimer returns the HasAutomaticDurationTimer field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateTestPlanApiModel) GetHasAutomaticDurationTimer() bool {
	if o == nil || IsNil(o.HasAutomaticDurationTimer.Get()) {
		var ret bool
		return ret
	}
	return *o.HasAutomaticDurationTimer.Get()
}

// GetHasAutomaticDurationTimerOk returns a tuple with the HasAutomaticDurationTimer field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateTestPlanApiModel) GetHasAutomaticDurationTimerOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.HasAutomaticDurationTimer.Get(), o.HasAutomaticDurationTimer.IsSet()
}

// HasHasAutomaticDurationTimer returns a boolean if a field has been set.
func (o *UpdateTestPlanApiModel) HasHasAutomaticDurationTimer() bool {
	if o != nil && o.HasAutomaticDurationTimer.IsSet() {
		return true
	}

	return false
}

// SetHasAutomaticDurationTimer gets a reference to the given NullableBool and assigns it to the HasAutomaticDurationTimer field.
func (o *UpdateTestPlanApiModel) SetHasAutomaticDurationTimer(v bool) {
	o.HasAutomaticDurationTimer.Set(&v)
}
// SetHasAutomaticDurationTimerNil sets the value for HasAutomaticDurationTimer to be an explicit nil
func (o *UpdateTestPlanApiModel) SetHasAutomaticDurationTimerNil() {
	o.HasAutomaticDurationTimer.Set(nil)
}

// UnsetHasAutomaticDurationTimer ensures that no value is present for HasAutomaticDurationTimer, not even an explicit nil
func (o *UpdateTestPlanApiModel) UnsetHasAutomaticDurationTimer() {
	o.HasAutomaticDurationTimer.Unset()
}

// GetAttributes returns the Attributes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateTestPlanApiModel) GetAttributes() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateTestPlanApiModel) GetAttributesOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Attributes) {
		return map[string]interface{}{}, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *UpdateTestPlanApiModel) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]interface{} and assigns it to the Attributes field.
func (o *UpdateTestPlanApiModel) SetAttributes(v map[string]interface{}) {
	o.Attributes = v
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateTestPlanApiModel) GetTags() []TagApiModel {
	if o == nil {
		var ret []TagApiModel
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateTestPlanApiModel) GetTagsOk() ([]TagApiModel, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *UpdateTestPlanApiModel) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []TagApiModel and assigns it to the Tags field.
func (o *UpdateTestPlanApiModel) SetTags(v []TagApiModel) {
	o.Tags = v
}

func (o UpdateTestPlanApiModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateTestPlanApiModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if o.LockedById.IsSet() {
		toSerialize["lockedById"] = o.LockedById.Get()
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
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	return toSerialize, nil
}

func (o *UpdateTestPlanApiModel) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"projectId",
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

	varUpdateTestPlanApiModel := _UpdateTestPlanApiModel{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUpdateTestPlanApiModel)

	if err != nil {
		return err
	}

	*o = UpdateTestPlanApiModel(varUpdateTestPlanApiModel)

	return err
}

type NullableUpdateTestPlanApiModel struct {
	value *UpdateTestPlanApiModel
	isSet bool
}

func (v NullableUpdateTestPlanApiModel) Get() *UpdateTestPlanApiModel {
	return v.value
}

func (v *NullableUpdateTestPlanApiModel) Set(val *UpdateTestPlanApiModel) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateTestPlanApiModel) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateTestPlanApiModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateTestPlanApiModel(val *UpdateTestPlanApiModel) *NullableUpdateTestPlanApiModel {
	return &NullableUpdateTestPlanApiModel{value: val, isSet: true}
}

func (v NullableUpdateTestPlanApiModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateTestPlanApiModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


