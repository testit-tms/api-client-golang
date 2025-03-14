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

// checks if the SectionWithStepsModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SectionWithStepsModel{}

// SectionWithStepsModel struct for SectionWithStepsModel
type SectionWithStepsModel struct {
	Attachments []AttachmentModel `json:"attachments,omitempty"`
	PreconditionSteps []StepModel `json:"preconditionSteps,omitempty"`
	PostconditionSteps []StepModel `json:"postconditionSteps,omitempty"`
	ProjectId NullableString `json:"projectId,omitempty"`
	ParentId NullableString `json:"parentId,omitempty"`
	IsDeleted bool `json:"isDeleted"`
	Id string `json:"id"`
	CreatedDate time.Time `json:"createdDate"`
	ModifiedDate NullableTime `json:"modifiedDate,omitempty"`
	CreatedById string `json:"createdById"`
	ModifiedById NullableString `json:"modifiedById,omitempty"`
	Name string `json:"name"`
}

type _SectionWithStepsModel SectionWithStepsModel

// NewSectionWithStepsModel instantiates a new SectionWithStepsModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSectionWithStepsModel(isDeleted bool, id string, createdDate time.Time, createdById string, name string) *SectionWithStepsModel {
	this := SectionWithStepsModel{}
	this.IsDeleted = isDeleted
	this.Id = id
	this.CreatedDate = createdDate
	this.CreatedById = createdById
	this.Name = name
	return &this
}

// NewSectionWithStepsModelWithDefaults instantiates a new SectionWithStepsModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSectionWithStepsModelWithDefaults() *SectionWithStepsModel {
	this := SectionWithStepsModel{}
	return &this
}

// GetAttachments returns the Attachments field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SectionWithStepsModel) GetAttachments() []AttachmentModel {
	if o == nil {
		var ret []AttachmentModel
		return ret
	}
	return o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SectionWithStepsModel) GetAttachmentsOk() ([]AttachmentModel, bool) {
	if o == nil || IsNil(o.Attachments) {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *SectionWithStepsModel) HasAttachments() bool {
	if o != nil && !IsNil(o.Attachments) {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given []AttachmentModel and assigns it to the Attachments field.
func (o *SectionWithStepsModel) SetAttachments(v []AttachmentModel) {
	o.Attachments = v
}

// GetPreconditionSteps returns the PreconditionSteps field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SectionWithStepsModel) GetPreconditionSteps() []StepModel {
	if o == nil {
		var ret []StepModel
		return ret
	}
	return o.PreconditionSteps
}

// GetPreconditionStepsOk returns a tuple with the PreconditionSteps field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SectionWithStepsModel) GetPreconditionStepsOk() ([]StepModel, bool) {
	if o == nil || IsNil(o.PreconditionSteps) {
		return nil, false
	}
	return o.PreconditionSteps, true
}

// HasPreconditionSteps returns a boolean if a field has been set.
func (o *SectionWithStepsModel) HasPreconditionSteps() bool {
	if o != nil && !IsNil(o.PreconditionSteps) {
		return true
	}

	return false
}

// SetPreconditionSteps gets a reference to the given []StepModel and assigns it to the PreconditionSteps field.
func (o *SectionWithStepsModel) SetPreconditionSteps(v []StepModel) {
	o.PreconditionSteps = v
}

// GetPostconditionSteps returns the PostconditionSteps field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SectionWithStepsModel) GetPostconditionSteps() []StepModel {
	if o == nil {
		var ret []StepModel
		return ret
	}
	return o.PostconditionSteps
}

// GetPostconditionStepsOk returns a tuple with the PostconditionSteps field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SectionWithStepsModel) GetPostconditionStepsOk() ([]StepModel, bool) {
	if o == nil || IsNil(o.PostconditionSteps) {
		return nil, false
	}
	return o.PostconditionSteps, true
}

// HasPostconditionSteps returns a boolean if a field has been set.
func (o *SectionWithStepsModel) HasPostconditionSteps() bool {
	if o != nil && !IsNil(o.PostconditionSteps) {
		return true
	}

	return false
}

// SetPostconditionSteps gets a reference to the given []StepModel and assigns it to the PostconditionSteps field.
func (o *SectionWithStepsModel) SetPostconditionSteps(v []StepModel) {
	o.PostconditionSteps = v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SectionWithStepsModel) GetProjectId() string {
	if o == nil || IsNil(o.ProjectId.Get()) {
		var ret string
		return ret
	}
	return *o.ProjectId.Get()
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SectionWithStepsModel) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProjectId.Get(), o.ProjectId.IsSet()
}

// HasProjectId returns a boolean if a field has been set.
func (o *SectionWithStepsModel) HasProjectId() bool {
	if o != nil && o.ProjectId.IsSet() {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given NullableString and assigns it to the ProjectId field.
func (o *SectionWithStepsModel) SetProjectId(v string) {
	o.ProjectId.Set(&v)
}
// SetProjectIdNil sets the value for ProjectId to be an explicit nil
func (o *SectionWithStepsModel) SetProjectIdNil() {
	o.ProjectId.Set(nil)
}

// UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
func (o *SectionWithStepsModel) UnsetProjectId() {
	o.ProjectId.Unset()
}

// GetParentId returns the ParentId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SectionWithStepsModel) GetParentId() string {
	if o == nil || IsNil(o.ParentId.Get()) {
		var ret string
		return ret
	}
	return *o.ParentId.Get()
}

// GetParentIdOk returns a tuple with the ParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SectionWithStepsModel) GetParentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ParentId.Get(), o.ParentId.IsSet()
}

// HasParentId returns a boolean if a field has been set.
func (o *SectionWithStepsModel) HasParentId() bool {
	if o != nil && o.ParentId.IsSet() {
		return true
	}

	return false
}

// SetParentId gets a reference to the given NullableString and assigns it to the ParentId field.
func (o *SectionWithStepsModel) SetParentId(v string) {
	o.ParentId.Set(&v)
}
// SetParentIdNil sets the value for ParentId to be an explicit nil
func (o *SectionWithStepsModel) SetParentIdNil() {
	o.ParentId.Set(nil)
}

// UnsetParentId ensures that no value is present for ParentId, not even an explicit nil
func (o *SectionWithStepsModel) UnsetParentId() {
	o.ParentId.Unset()
}

// GetIsDeleted returns the IsDeleted field value
func (o *SectionWithStepsModel) GetIsDeleted() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsDeleted
}

// GetIsDeletedOk returns a tuple with the IsDeleted field value
// and a boolean to check if the value has been set.
func (o *SectionWithStepsModel) GetIsDeletedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsDeleted, true
}

// SetIsDeleted sets field value
func (o *SectionWithStepsModel) SetIsDeleted(v bool) {
	o.IsDeleted = v
}

// GetId returns the Id field value
func (o *SectionWithStepsModel) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SectionWithStepsModel) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SectionWithStepsModel) SetId(v string) {
	o.Id = v
}

// GetCreatedDate returns the CreatedDate field value
func (o *SectionWithStepsModel) GetCreatedDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedDate
}

// GetCreatedDateOk returns a tuple with the CreatedDate field value
// and a boolean to check if the value has been set.
func (o *SectionWithStepsModel) GetCreatedDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedDate, true
}

// SetCreatedDate sets field value
func (o *SectionWithStepsModel) SetCreatedDate(v time.Time) {
	o.CreatedDate = v
}

// GetModifiedDate returns the ModifiedDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SectionWithStepsModel) GetModifiedDate() time.Time {
	if o == nil || IsNil(o.ModifiedDate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.ModifiedDate.Get()
}

// GetModifiedDateOk returns a tuple with the ModifiedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SectionWithStepsModel) GetModifiedDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ModifiedDate.Get(), o.ModifiedDate.IsSet()
}

// HasModifiedDate returns a boolean if a field has been set.
func (o *SectionWithStepsModel) HasModifiedDate() bool {
	if o != nil && o.ModifiedDate.IsSet() {
		return true
	}

	return false
}

// SetModifiedDate gets a reference to the given NullableTime and assigns it to the ModifiedDate field.
func (o *SectionWithStepsModel) SetModifiedDate(v time.Time) {
	o.ModifiedDate.Set(&v)
}
// SetModifiedDateNil sets the value for ModifiedDate to be an explicit nil
func (o *SectionWithStepsModel) SetModifiedDateNil() {
	o.ModifiedDate.Set(nil)
}

// UnsetModifiedDate ensures that no value is present for ModifiedDate, not even an explicit nil
func (o *SectionWithStepsModel) UnsetModifiedDate() {
	o.ModifiedDate.Unset()
}

// GetCreatedById returns the CreatedById field value
func (o *SectionWithStepsModel) GetCreatedById() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedById
}

// GetCreatedByIdOk returns a tuple with the CreatedById field value
// and a boolean to check if the value has been set.
func (o *SectionWithStepsModel) GetCreatedByIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedById, true
}

// SetCreatedById sets field value
func (o *SectionWithStepsModel) SetCreatedById(v string) {
	o.CreatedById = v
}

// GetModifiedById returns the ModifiedById field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SectionWithStepsModel) GetModifiedById() string {
	if o == nil || IsNil(o.ModifiedById.Get()) {
		var ret string
		return ret
	}
	return *o.ModifiedById.Get()
}

// GetModifiedByIdOk returns a tuple with the ModifiedById field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SectionWithStepsModel) GetModifiedByIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ModifiedById.Get(), o.ModifiedById.IsSet()
}

// HasModifiedById returns a boolean if a field has been set.
func (o *SectionWithStepsModel) HasModifiedById() bool {
	if o != nil && o.ModifiedById.IsSet() {
		return true
	}

	return false
}

// SetModifiedById gets a reference to the given NullableString and assigns it to the ModifiedById field.
func (o *SectionWithStepsModel) SetModifiedById(v string) {
	o.ModifiedById.Set(&v)
}
// SetModifiedByIdNil sets the value for ModifiedById to be an explicit nil
func (o *SectionWithStepsModel) SetModifiedByIdNil() {
	o.ModifiedById.Set(nil)
}

// UnsetModifiedById ensures that no value is present for ModifiedById, not even an explicit nil
func (o *SectionWithStepsModel) UnsetModifiedById() {
	o.ModifiedById.Unset()
}

// GetName returns the Name field value
func (o *SectionWithStepsModel) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SectionWithStepsModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SectionWithStepsModel) SetName(v string) {
	o.Name = v
}

func (o SectionWithStepsModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SectionWithStepsModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Attachments != nil {
		toSerialize["attachments"] = o.Attachments
	}
	if o.PreconditionSteps != nil {
		toSerialize["preconditionSteps"] = o.PreconditionSteps
	}
	if o.PostconditionSteps != nil {
		toSerialize["postconditionSteps"] = o.PostconditionSteps
	}
	if o.ProjectId.IsSet() {
		toSerialize["projectId"] = o.ProjectId.Get()
	}
	if o.ParentId.IsSet() {
		toSerialize["parentId"] = o.ParentId.Get()
	}
	toSerialize["isDeleted"] = o.IsDeleted
	toSerialize["id"] = o.Id
	toSerialize["createdDate"] = o.CreatedDate
	if o.ModifiedDate.IsSet() {
		toSerialize["modifiedDate"] = o.ModifiedDate.Get()
	}
	toSerialize["createdById"] = o.CreatedById
	if o.ModifiedById.IsSet() {
		toSerialize["modifiedById"] = o.ModifiedById.Get()
	}
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

func (o *SectionWithStepsModel) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"isDeleted",
		"id",
		"createdDate",
		"createdById",
		"name",
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

	varSectionWithStepsModel := _SectionWithStepsModel{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSectionWithStepsModel)

	if err != nil {
		return err
	}

	*o = SectionWithStepsModel(varSectionWithStepsModel)

	return err
}

type NullableSectionWithStepsModel struct {
	value *SectionWithStepsModel
	isSet bool
}

func (v NullableSectionWithStepsModel) Get() *SectionWithStepsModel {
	return v.value
}

func (v *NullableSectionWithStepsModel) Set(val *SectionWithStepsModel) {
	v.value = val
	v.isSet = true
}

func (v NullableSectionWithStepsModel) IsSet() bool {
	return v.isSet
}

func (v *NullableSectionWithStepsModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSectionWithStepsModel(val *SectionWithStepsModel) *NullableSectionWithStepsModel {
	return &NullableSectionWithStepsModel{value: val, isSet: true}
}

func (v NullableSectionWithStepsModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSectionWithStepsModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


