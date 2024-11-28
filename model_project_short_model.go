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

// checks if the ProjectShortModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectShortModel{}

// ProjectShortModel struct for ProjectShortModel
type ProjectShortModel struct {
	// Unique ID of the project
	Id string `json:"id"`
	// Description of the project
	Description NullableString `json:"description,omitempty"`
	// Name of the project
	Name string `json:"name"`
	// Indicates if the project is marked as favorite
	IsFavorite bool `json:"isFavorite"`
	// Number of test cases in the project
	TestCasesCount NullableInt32 `json:"testCasesCount,omitempty"`
	// Number of shared steps in the project
	SharedStepsCount NullableInt32 `json:"sharedStepsCount,omitempty"`
	// Number of checklists in the project
	CheckListsCount NullableInt32 `json:"checkListsCount,omitempty"`
	// Number of autotests in the project
	AutoTestsCount NullableInt32 `json:"autoTestsCount,omitempty"`
	// Indicates if the project is deleted
	IsDeleted bool `json:"isDeleted"`
	// Creation date of the project
	CreatedDate time.Time `json:"createdDate"`
	// Last modification date of the project
	ModifiedDate NullableTime `json:"modifiedDate,omitempty"`
	// Unique ID of the project creator
	CreatedById string `json:"createdById"`
	// Unique ID of the project last editor
	ModifiedById NullableString `json:"modifiedById,omitempty"`
	// Global ID of the project
	GlobalId int64 `json:"globalId"`
	// Type of the project
	Type ProjectTypeModel `json:"type"`
	// Indicates if the status \"Flaky/Stable\" sets automatically
	// Deprecated
	IsFlakyAuto bool `json:"isFlakyAuto"`
}

type _ProjectShortModel ProjectShortModel

// NewProjectShortModel instantiates a new ProjectShortModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectShortModel(id string, name string, isFavorite bool, isDeleted bool, createdDate time.Time, createdById string, globalId int64, type_ ProjectTypeModel, isFlakyAuto bool) *ProjectShortModel {
	this := ProjectShortModel{}
	this.Id = id
	this.Name = name
	this.IsFavorite = isFavorite
	this.IsDeleted = isDeleted
	this.CreatedDate = createdDate
	this.CreatedById = createdById
	this.GlobalId = globalId
	this.Type = type_
	this.IsFlakyAuto = isFlakyAuto
	return &this
}

// NewProjectShortModelWithDefaults instantiates a new ProjectShortModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectShortModelWithDefaults() *ProjectShortModel {
	this := ProjectShortModel{}
	return &this
}

// GetId returns the Id field value
func (o *ProjectShortModel) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ProjectShortModel) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ProjectShortModel) SetId(v string) {
	o.Id = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectShortModel) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectShortModel) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *ProjectShortModel) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *ProjectShortModel) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *ProjectShortModel) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *ProjectShortModel) UnsetDescription() {
	o.Description.Unset()
}

// GetName returns the Name field value
func (o *ProjectShortModel) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ProjectShortModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ProjectShortModel) SetName(v string) {
	o.Name = v
}

// GetIsFavorite returns the IsFavorite field value
func (o *ProjectShortModel) GetIsFavorite() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsFavorite
}

// GetIsFavoriteOk returns a tuple with the IsFavorite field value
// and a boolean to check if the value has been set.
func (o *ProjectShortModel) GetIsFavoriteOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsFavorite, true
}

// SetIsFavorite sets field value
func (o *ProjectShortModel) SetIsFavorite(v bool) {
	o.IsFavorite = v
}

// GetTestCasesCount returns the TestCasesCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectShortModel) GetTestCasesCount() int32 {
	if o == nil || IsNil(o.TestCasesCount.Get()) {
		var ret int32
		return ret
	}
	return *o.TestCasesCount.Get()
}

// GetTestCasesCountOk returns a tuple with the TestCasesCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectShortModel) GetTestCasesCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.TestCasesCount.Get(), o.TestCasesCount.IsSet()
}

// HasTestCasesCount returns a boolean if a field has been set.
func (o *ProjectShortModel) HasTestCasesCount() bool {
	if o != nil && o.TestCasesCount.IsSet() {
		return true
	}

	return false
}

// SetTestCasesCount gets a reference to the given NullableInt32 and assigns it to the TestCasesCount field.
func (o *ProjectShortModel) SetTestCasesCount(v int32) {
	o.TestCasesCount.Set(&v)
}
// SetTestCasesCountNil sets the value for TestCasesCount to be an explicit nil
func (o *ProjectShortModel) SetTestCasesCountNil() {
	o.TestCasesCount.Set(nil)
}

// UnsetTestCasesCount ensures that no value is present for TestCasesCount, not even an explicit nil
func (o *ProjectShortModel) UnsetTestCasesCount() {
	o.TestCasesCount.Unset()
}

// GetSharedStepsCount returns the SharedStepsCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectShortModel) GetSharedStepsCount() int32 {
	if o == nil || IsNil(o.SharedStepsCount.Get()) {
		var ret int32
		return ret
	}
	return *o.SharedStepsCount.Get()
}

// GetSharedStepsCountOk returns a tuple with the SharedStepsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectShortModel) GetSharedStepsCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.SharedStepsCount.Get(), o.SharedStepsCount.IsSet()
}

// HasSharedStepsCount returns a boolean if a field has been set.
func (o *ProjectShortModel) HasSharedStepsCount() bool {
	if o != nil && o.SharedStepsCount.IsSet() {
		return true
	}

	return false
}

// SetSharedStepsCount gets a reference to the given NullableInt32 and assigns it to the SharedStepsCount field.
func (o *ProjectShortModel) SetSharedStepsCount(v int32) {
	o.SharedStepsCount.Set(&v)
}
// SetSharedStepsCountNil sets the value for SharedStepsCount to be an explicit nil
func (o *ProjectShortModel) SetSharedStepsCountNil() {
	o.SharedStepsCount.Set(nil)
}

// UnsetSharedStepsCount ensures that no value is present for SharedStepsCount, not even an explicit nil
func (o *ProjectShortModel) UnsetSharedStepsCount() {
	o.SharedStepsCount.Unset()
}

// GetCheckListsCount returns the CheckListsCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectShortModel) GetCheckListsCount() int32 {
	if o == nil || IsNil(o.CheckListsCount.Get()) {
		var ret int32
		return ret
	}
	return *o.CheckListsCount.Get()
}

// GetCheckListsCountOk returns a tuple with the CheckListsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectShortModel) GetCheckListsCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.CheckListsCount.Get(), o.CheckListsCount.IsSet()
}

// HasCheckListsCount returns a boolean if a field has been set.
func (o *ProjectShortModel) HasCheckListsCount() bool {
	if o != nil && o.CheckListsCount.IsSet() {
		return true
	}

	return false
}

// SetCheckListsCount gets a reference to the given NullableInt32 and assigns it to the CheckListsCount field.
func (o *ProjectShortModel) SetCheckListsCount(v int32) {
	o.CheckListsCount.Set(&v)
}
// SetCheckListsCountNil sets the value for CheckListsCount to be an explicit nil
func (o *ProjectShortModel) SetCheckListsCountNil() {
	o.CheckListsCount.Set(nil)
}

// UnsetCheckListsCount ensures that no value is present for CheckListsCount, not even an explicit nil
func (o *ProjectShortModel) UnsetCheckListsCount() {
	o.CheckListsCount.Unset()
}

// GetAutoTestsCount returns the AutoTestsCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectShortModel) GetAutoTestsCount() int32 {
	if o == nil || IsNil(o.AutoTestsCount.Get()) {
		var ret int32
		return ret
	}
	return *o.AutoTestsCount.Get()
}

// GetAutoTestsCountOk returns a tuple with the AutoTestsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectShortModel) GetAutoTestsCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.AutoTestsCount.Get(), o.AutoTestsCount.IsSet()
}

// HasAutoTestsCount returns a boolean if a field has been set.
func (o *ProjectShortModel) HasAutoTestsCount() bool {
	if o != nil && o.AutoTestsCount.IsSet() {
		return true
	}

	return false
}

// SetAutoTestsCount gets a reference to the given NullableInt32 and assigns it to the AutoTestsCount field.
func (o *ProjectShortModel) SetAutoTestsCount(v int32) {
	o.AutoTestsCount.Set(&v)
}
// SetAutoTestsCountNil sets the value for AutoTestsCount to be an explicit nil
func (o *ProjectShortModel) SetAutoTestsCountNil() {
	o.AutoTestsCount.Set(nil)
}

// UnsetAutoTestsCount ensures that no value is present for AutoTestsCount, not even an explicit nil
func (o *ProjectShortModel) UnsetAutoTestsCount() {
	o.AutoTestsCount.Unset()
}

// GetIsDeleted returns the IsDeleted field value
func (o *ProjectShortModel) GetIsDeleted() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsDeleted
}

// GetIsDeletedOk returns a tuple with the IsDeleted field value
// and a boolean to check if the value has been set.
func (o *ProjectShortModel) GetIsDeletedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsDeleted, true
}

// SetIsDeleted sets field value
func (o *ProjectShortModel) SetIsDeleted(v bool) {
	o.IsDeleted = v
}

// GetCreatedDate returns the CreatedDate field value
func (o *ProjectShortModel) GetCreatedDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedDate
}

// GetCreatedDateOk returns a tuple with the CreatedDate field value
// and a boolean to check if the value has been set.
func (o *ProjectShortModel) GetCreatedDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedDate, true
}

// SetCreatedDate sets field value
func (o *ProjectShortModel) SetCreatedDate(v time.Time) {
	o.CreatedDate = v
}

// GetModifiedDate returns the ModifiedDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectShortModel) GetModifiedDate() time.Time {
	if o == nil || IsNil(o.ModifiedDate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.ModifiedDate.Get()
}

// GetModifiedDateOk returns a tuple with the ModifiedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectShortModel) GetModifiedDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ModifiedDate.Get(), o.ModifiedDate.IsSet()
}

// HasModifiedDate returns a boolean if a field has been set.
func (o *ProjectShortModel) HasModifiedDate() bool {
	if o != nil && o.ModifiedDate.IsSet() {
		return true
	}

	return false
}

// SetModifiedDate gets a reference to the given NullableTime and assigns it to the ModifiedDate field.
func (o *ProjectShortModel) SetModifiedDate(v time.Time) {
	o.ModifiedDate.Set(&v)
}
// SetModifiedDateNil sets the value for ModifiedDate to be an explicit nil
func (o *ProjectShortModel) SetModifiedDateNil() {
	o.ModifiedDate.Set(nil)
}

// UnsetModifiedDate ensures that no value is present for ModifiedDate, not even an explicit nil
func (o *ProjectShortModel) UnsetModifiedDate() {
	o.ModifiedDate.Unset()
}

// GetCreatedById returns the CreatedById field value
func (o *ProjectShortModel) GetCreatedById() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedById
}

// GetCreatedByIdOk returns a tuple with the CreatedById field value
// and a boolean to check if the value has been set.
func (o *ProjectShortModel) GetCreatedByIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedById, true
}

// SetCreatedById sets field value
func (o *ProjectShortModel) SetCreatedById(v string) {
	o.CreatedById = v
}

// GetModifiedById returns the ModifiedById field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectShortModel) GetModifiedById() string {
	if o == nil || IsNil(o.ModifiedById.Get()) {
		var ret string
		return ret
	}
	return *o.ModifiedById.Get()
}

// GetModifiedByIdOk returns a tuple with the ModifiedById field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectShortModel) GetModifiedByIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ModifiedById.Get(), o.ModifiedById.IsSet()
}

// HasModifiedById returns a boolean if a field has been set.
func (o *ProjectShortModel) HasModifiedById() bool {
	if o != nil && o.ModifiedById.IsSet() {
		return true
	}

	return false
}

// SetModifiedById gets a reference to the given NullableString and assigns it to the ModifiedById field.
func (o *ProjectShortModel) SetModifiedById(v string) {
	o.ModifiedById.Set(&v)
}
// SetModifiedByIdNil sets the value for ModifiedById to be an explicit nil
func (o *ProjectShortModel) SetModifiedByIdNil() {
	o.ModifiedById.Set(nil)
}

// UnsetModifiedById ensures that no value is present for ModifiedById, not even an explicit nil
func (o *ProjectShortModel) UnsetModifiedById() {
	o.ModifiedById.Unset()
}

// GetGlobalId returns the GlobalId field value
func (o *ProjectShortModel) GetGlobalId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.GlobalId
}

// GetGlobalIdOk returns a tuple with the GlobalId field value
// and a boolean to check if the value has been set.
func (o *ProjectShortModel) GetGlobalIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GlobalId, true
}

// SetGlobalId sets field value
func (o *ProjectShortModel) SetGlobalId(v int64) {
	o.GlobalId = v
}

// GetType returns the Type field value
func (o *ProjectShortModel) GetType() ProjectTypeModel {
	if o == nil {
		var ret ProjectTypeModel
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ProjectShortModel) GetTypeOk() (*ProjectTypeModel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ProjectShortModel) SetType(v ProjectTypeModel) {
	o.Type = v
}

// GetIsFlakyAuto returns the IsFlakyAuto field value
// Deprecated
func (o *ProjectShortModel) GetIsFlakyAuto() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsFlakyAuto
}

// GetIsFlakyAutoOk returns a tuple with the IsFlakyAuto field value
// and a boolean to check if the value has been set.
// Deprecated
func (o *ProjectShortModel) GetIsFlakyAutoOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsFlakyAuto, true
}

// SetIsFlakyAuto sets field value
// Deprecated
func (o *ProjectShortModel) SetIsFlakyAuto(v bool) {
	o.IsFlakyAuto = v
}

func (o ProjectShortModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectShortModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	toSerialize["name"] = o.Name
	toSerialize["isFavorite"] = o.IsFavorite
	if o.TestCasesCount.IsSet() {
		toSerialize["testCasesCount"] = o.TestCasesCount.Get()
	}
	if o.SharedStepsCount.IsSet() {
		toSerialize["sharedStepsCount"] = o.SharedStepsCount.Get()
	}
	if o.CheckListsCount.IsSet() {
		toSerialize["checkListsCount"] = o.CheckListsCount.Get()
	}
	if o.AutoTestsCount.IsSet() {
		toSerialize["autoTestsCount"] = o.AutoTestsCount.Get()
	}
	toSerialize["isDeleted"] = o.IsDeleted
	toSerialize["createdDate"] = o.CreatedDate
	if o.ModifiedDate.IsSet() {
		toSerialize["modifiedDate"] = o.ModifiedDate.Get()
	}
	toSerialize["createdById"] = o.CreatedById
	if o.ModifiedById.IsSet() {
		toSerialize["modifiedById"] = o.ModifiedById.Get()
	}
	toSerialize["globalId"] = o.GlobalId
	toSerialize["type"] = o.Type
	toSerialize["isFlakyAuto"] = o.IsFlakyAuto
	return toSerialize, nil
}

func (o *ProjectShortModel) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"isFavorite",
		"isDeleted",
		"createdDate",
		"createdById",
		"globalId",
		"type",
		"isFlakyAuto",
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

	varProjectShortModel := _ProjectShortModel{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varProjectShortModel)

	if err != nil {
		return err
	}

	*o = ProjectShortModel(varProjectShortModel)

	return err
}

type NullableProjectShortModel struct {
	value *ProjectShortModel
	isSet bool
}

func (v NullableProjectShortModel) Get() *ProjectShortModel {
	return v.value
}

func (v *NullableProjectShortModel) Set(val *ProjectShortModel) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectShortModel) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectShortModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectShortModel(val *ProjectShortModel) *NullableProjectShortModel {
	return &NullableProjectShortModel{value: val, isSet: true}
}

func (v NullableProjectShortModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectShortModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


