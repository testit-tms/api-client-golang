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

// checks if the ProjectModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectModel{}

// ProjectModel struct for ProjectModel
type ProjectModel struct {
	// Unique ID of the project
	Id string `json:"id"`
	// Description of the project
	Description NullableString `json:"description,omitempty"`
	// Name of the project
	Name string `json:"name"`
	// Indicates if the project is marked as favorite
	IsFavorite bool `json:"isFavorite"`
	// Collection of the project attributes
	AttributesScheme []CustomAttributeModel `json:"attributesScheme,omitempty"`
	// Collection of the project test plans attributes
	TestPlansAttributesScheme []CustomAttributeModel `json:"testPlansAttributesScheme,omitempty"`
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
	Type ProjectTypeModel `json:"type"`
	// Indicates if the status \"Flaky/Stable\" sets automatically
	IsFlakyAuto bool `json:"isFlakyAuto"`
}

// NewProjectModel instantiates a new ProjectModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectModel(id string, name string, isFavorite bool, isDeleted bool, createdDate time.Time, createdById string, globalId int64, type_ ProjectTypeModel, isFlakyAuto bool) *ProjectModel {
	this := ProjectModel{}
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

// NewProjectModelWithDefaults instantiates a new ProjectModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectModelWithDefaults() *ProjectModel {
	this := ProjectModel{}
	return &this
}

// GetId returns the Id field value
func (o *ProjectModel) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ProjectModel) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ProjectModel) SetId(v string) {
	o.Id = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectModel) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectModel) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *ProjectModel) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *ProjectModel) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *ProjectModel) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *ProjectModel) UnsetDescription() {
	o.Description.Unset()
}

// GetName returns the Name field value
func (o *ProjectModel) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ProjectModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ProjectModel) SetName(v string) {
	o.Name = v
}

// GetIsFavorite returns the IsFavorite field value
func (o *ProjectModel) GetIsFavorite() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsFavorite
}

// GetIsFavoriteOk returns a tuple with the IsFavorite field value
// and a boolean to check if the value has been set.
func (o *ProjectModel) GetIsFavoriteOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsFavorite, true
}

// SetIsFavorite sets field value
func (o *ProjectModel) SetIsFavorite(v bool) {
	o.IsFavorite = v
}

// GetAttributesScheme returns the AttributesScheme field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectModel) GetAttributesScheme() []CustomAttributeModel {
	if o == nil {
		var ret []CustomAttributeModel
		return ret
	}
	return o.AttributesScheme
}

// GetAttributesSchemeOk returns a tuple with the AttributesScheme field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectModel) GetAttributesSchemeOk() ([]CustomAttributeModel, bool) {
	if o == nil || IsNil(o.AttributesScheme) {
		return nil, false
	}
	return o.AttributesScheme, true
}

// HasAttributesScheme returns a boolean if a field has been set.
func (o *ProjectModel) HasAttributesScheme() bool {
	if o != nil && IsNil(o.AttributesScheme) {
		return true
	}

	return false
}

// SetAttributesScheme gets a reference to the given []CustomAttributeModel and assigns it to the AttributesScheme field.
func (o *ProjectModel) SetAttributesScheme(v []CustomAttributeModel) {
	o.AttributesScheme = v
}

// GetTestPlansAttributesScheme returns the TestPlansAttributesScheme field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectModel) GetTestPlansAttributesScheme() []CustomAttributeModel {
	if o == nil {
		var ret []CustomAttributeModel
		return ret
	}
	return o.TestPlansAttributesScheme
}

// GetTestPlansAttributesSchemeOk returns a tuple with the TestPlansAttributesScheme field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectModel) GetTestPlansAttributesSchemeOk() ([]CustomAttributeModel, bool) {
	if o == nil || IsNil(o.TestPlansAttributesScheme) {
		return nil, false
	}
	return o.TestPlansAttributesScheme, true
}

// HasTestPlansAttributesScheme returns a boolean if a field has been set.
func (o *ProjectModel) HasTestPlansAttributesScheme() bool {
	if o != nil && IsNil(o.TestPlansAttributesScheme) {
		return true
	}

	return false
}

// SetTestPlansAttributesScheme gets a reference to the given []CustomAttributeModel and assigns it to the TestPlansAttributesScheme field.
func (o *ProjectModel) SetTestPlansAttributesScheme(v []CustomAttributeModel) {
	o.TestPlansAttributesScheme = v
}

// GetTestCasesCount returns the TestCasesCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectModel) GetTestCasesCount() int32 {
	if o == nil || IsNil(o.TestCasesCount.Get()) {
		var ret int32
		return ret
	}
	return *o.TestCasesCount.Get()
}

// GetTestCasesCountOk returns a tuple with the TestCasesCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectModel) GetTestCasesCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.TestCasesCount.Get(), o.TestCasesCount.IsSet()
}

// HasTestCasesCount returns a boolean if a field has been set.
func (o *ProjectModel) HasTestCasesCount() bool {
	if o != nil && o.TestCasesCount.IsSet() {
		return true
	}

	return false
}

// SetTestCasesCount gets a reference to the given NullableInt32 and assigns it to the TestCasesCount field.
func (o *ProjectModel) SetTestCasesCount(v int32) {
	o.TestCasesCount.Set(&v)
}
// SetTestCasesCountNil sets the value for TestCasesCount to be an explicit nil
func (o *ProjectModel) SetTestCasesCountNil() {
	o.TestCasesCount.Set(nil)
}

// UnsetTestCasesCount ensures that no value is present for TestCasesCount, not even an explicit nil
func (o *ProjectModel) UnsetTestCasesCount() {
	o.TestCasesCount.Unset()
}

// GetSharedStepsCount returns the SharedStepsCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectModel) GetSharedStepsCount() int32 {
	if o == nil || IsNil(o.SharedStepsCount.Get()) {
		var ret int32
		return ret
	}
	return *o.SharedStepsCount.Get()
}

// GetSharedStepsCountOk returns a tuple with the SharedStepsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectModel) GetSharedStepsCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.SharedStepsCount.Get(), o.SharedStepsCount.IsSet()
}

// HasSharedStepsCount returns a boolean if a field has been set.
func (o *ProjectModel) HasSharedStepsCount() bool {
	if o != nil && o.SharedStepsCount.IsSet() {
		return true
	}

	return false
}

// SetSharedStepsCount gets a reference to the given NullableInt32 and assigns it to the SharedStepsCount field.
func (o *ProjectModel) SetSharedStepsCount(v int32) {
	o.SharedStepsCount.Set(&v)
}
// SetSharedStepsCountNil sets the value for SharedStepsCount to be an explicit nil
func (o *ProjectModel) SetSharedStepsCountNil() {
	o.SharedStepsCount.Set(nil)
}

// UnsetSharedStepsCount ensures that no value is present for SharedStepsCount, not even an explicit nil
func (o *ProjectModel) UnsetSharedStepsCount() {
	o.SharedStepsCount.Unset()
}

// GetCheckListsCount returns the CheckListsCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectModel) GetCheckListsCount() int32 {
	if o == nil || IsNil(o.CheckListsCount.Get()) {
		var ret int32
		return ret
	}
	return *o.CheckListsCount.Get()
}

// GetCheckListsCountOk returns a tuple with the CheckListsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectModel) GetCheckListsCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.CheckListsCount.Get(), o.CheckListsCount.IsSet()
}

// HasCheckListsCount returns a boolean if a field has been set.
func (o *ProjectModel) HasCheckListsCount() bool {
	if o != nil && o.CheckListsCount.IsSet() {
		return true
	}

	return false
}

// SetCheckListsCount gets a reference to the given NullableInt32 and assigns it to the CheckListsCount field.
func (o *ProjectModel) SetCheckListsCount(v int32) {
	o.CheckListsCount.Set(&v)
}
// SetCheckListsCountNil sets the value for CheckListsCount to be an explicit nil
func (o *ProjectModel) SetCheckListsCountNil() {
	o.CheckListsCount.Set(nil)
}

// UnsetCheckListsCount ensures that no value is present for CheckListsCount, not even an explicit nil
func (o *ProjectModel) UnsetCheckListsCount() {
	o.CheckListsCount.Unset()
}

// GetAutoTestsCount returns the AutoTestsCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectModel) GetAutoTestsCount() int32 {
	if o == nil || IsNil(o.AutoTestsCount.Get()) {
		var ret int32
		return ret
	}
	return *o.AutoTestsCount.Get()
}

// GetAutoTestsCountOk returns a tuple with the AutoTestsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectModel) GetAutoTestsCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.AutoTestsCount.Get(), o.AutoTestsCount.IsSet()
}

// HasAutoTestsCount returns a boolean if a field has been set.
func (o *ProjectModel) HasAutoTestsCount() bool {
	if o != nil && o.AutoTestsCount.IsSet() {
		return true
	}

	return false
}

// SetAutoTestsCount gets a reference to the given NullableInt32 and assigns it to the AutoTestsCount field.
func (o *ProjectModel) SetAutoTestsCount(v int32) {
	o.AutoTestsCount.Set(&v)
}
// SetAutoTestsCountNil sets the value for AutoTestsCount to be an explicit nil
func (o *ProjectModel) SetAutoTestsCountNil() {
	o.AutoTestsCount.Set(nil)
}

// UnsetAutoTestsCount ensures that no value is present for AutoTestsCount, not even an explicit nil
func (o *ProjectModel) UnsetAutoTestsCount() {
	o.AutoTestsCount.Unset()
}

// GetIsDeleted returns the IsDeleted field value
func (o *ProjectModel) GetIsDeleted() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsDeleted
}

// GetIsDeletedOk returns a tuple with the IsDeleted field value
// and a boolean to check if the value has been set.
func (o *ProjectModel) GetIsDeletedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsDeleted, true
}

// SetIsDeleted sets field value
func (o *ProjectModel) SetIsDeleted(v bool) {
	o.IsDeleted = v
}

// GetCreatedDate returns the CreatedDate field value
func (o *ProjectModel) GetCreatedDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedDate
}

// GetCreatedDateOk returns a tuple with the CreatedDate field value
// and a boolean to check if the value has been set.
func (o *ProjectModel) GetCreatedDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedDate, true
}

// SetCreatedDate sets field value
func (o *ProjectModel) SetCreatedDate(v time.Time) {
	o.CreatedDate = v
}

// GetModifiedDate returns the ModifiedDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectModel) GetModifiedDate() time.Time {
	if o == nil || IsNil(o.ModifiedDate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.ModifiedDate.Get()
}

// GetModifiedDateOk returns a tuple with the ModifiedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectModel) GetModifiedDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ModifiedDate.Get(), o.ModifiedDate.IsSet()
}

// HasModifiedDate returns a boolean if a field has been set.
func (o *ProjectModel) HasModifiedDate() bool {
	if o != nil && o.ModifiedDate.IsSet() {
		return true
	}

	return false
}

// SetModifiedDate gets a reference to the given NullableTime and assigns it to the ModifiedDate field.
func (o *ProjectModel) SetModifiedDate(v time.Time) {
	o.ModifiedDate.Set(&v)
}
// SetModifiedDateNil sets the value for ModifiedDate to be an explicit nil
func (o *ProjectModel) SetModifiedDateNil() {
	o.ModifiedDate.Set(nil)
}

// UnsetModifiedDate ensures that no value is present for ModifiedDate, not even an explicit nil
func (o *ProjectModel) UnsetModifiedDate() {
	o.ModifiedDate.Unset()
}

// GetCreatedById returns the CreatedById field value
func (o *ProjectModel) GetCreatedById() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedById
}

// GetCreatedByIdOk returns a tuple with the CreatedById field value
// and a boolean to check if the value has been set.
func (o *ProjectModel) GetCreatedByIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedById, true
}

// SetCreatedById sets field value
func (o *ProjectModel) SetCreatedById(v string) {
	o.CreatedById = v
}

// GetModifiedById returns the ModifiedById field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectModel) GetModifiedById() string {
	if o == nil || IsNil(o.ModifiedById.Get()) {
		var ret string
		return ret
	}
	return *o.ModifiedById.Get()
}

// GetModifiedByIdOk returns a tuple with the ModifiedById field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectModel) GetModifiedByIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ModifiedById.Get(), o.ModifiedById.IsSet()
}

// HasModifiedById returns a boolean if a field has been set.
func (o *ProjectModel) HasModifiedById() bool {
	if o != nil && o.ModifiedById.IsSet() {
		return true
	}

	return false
}

// SetModifiedById gets a reference to the given NullableString and assigns it to the ModifiedById field.
func (o *ProjectModel) SetModifiedById(v string) {
	o.ModifiedById.Set(&v)
}
// SetModifiedByIdNil sets the value for ModifiedById to be an explicit nil
func (o *ProjectModel) SetModifiedByIdNil() {
	o.ModifiedById.Set(nil)
}

// UnsetModifiedById ensures that no value is present for ModifiedById, not even an explicit nil
func (o *ProjectModel) UnsetModifiedById() {
	o.ModifiedById.Unset()
}

// GetGlobalId returns the GlobalId field value
func (o *ProjectModel) GetGlobalId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.GlobalId
}

// GetGlobalIdOk returns a tuple with the GlobalId field value
// and a boolean to check if the value has been set.
func (o *ProjectModel) GetGlobalIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GlobalId, true
}

// SetGlobalId sets field value
func (o *ProjectModel) SetGlobalId(v int64) {
	o.GlobalId = v
}

// GetType returns the Type field value
func (o *ProjectModel) GetType() ProjectTypeModel {
	if o == nil {
		var ret ProjectTypeModel
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ProjectModel) GetTypeOk() (*ProjectTypeModel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ProjectModel) SetType(v ProjectTypeModel) {
	o.Type = v
}

// GetIsFlakyAuto returns the IsFlakyAuto field value
func (o *ProjectModel) GetIsFlakyAuto() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsFlakyAuto
}

// GetIsFlakyAutoOk returns a tuple with the IsFlakyAuto field value
// and a boolean to check if the value has been set.
func (o *ProjectModel) GetIsFlakyAutoOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsFlakyAuto, true
}

// SetIsFlakyAuto sets field value
func (o *ProjectModel) SetIsFlakyAuto(v bool) {
	o.IsFlakyAuto = v
}

func (o ProjectModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	toSerialize["name"] = o.Name
	toSerialize["isFavorite"] = o.IsFavorite
	if o.AttributesScheme != nil {
		toSerialize["attributesScheme"] = o.AttributesScheme
	}
	if o.TestPlansAttributesScheme != nil {
		toSerialize["testPlansAttributesScheme"] = o.TestPlansAttributesScheme
	}
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

type NullableProjectModel struct {
	value *ProjectModel
	isSet bool
}

func (v NullableProjectModel) Get() *ProjectModel {
	return v.value
}

func (v *NullableProjectModel) Set(val *ProjectModel) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectModel) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectModel(val *ProjectModel) *NullableProjectModel {
	return &NullableProjectModel{value: val, isSet: true}
}

func (v NullableProjectModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


