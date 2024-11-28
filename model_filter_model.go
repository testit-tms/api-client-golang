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

// checks if the FilterModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FilterModel{}

// FilterModel struct for FilterModel
type FilterModel struct {
	CreatedDate time.Time `json:"createdDate"`
	ModifiedDate NullableTime `json:"modifiedDate,omitempty"`
	CreatedById string `json:"createdById"`
	ModifiedById NullableString `json:"modifiedById,omitempty"`
	Data WorkItemSearchQueryModel `json:"data"`
	ProjectId string `json:"projectId"`
	FieldsToShow interface{} `json:"fieldsToShow,omitempty"`
	Name string `json:"name"`
	// Unique ID of the entity
	Id string `json:"id"`
	// Indicates if the entity is deleted
	IsDeleted bool `json:"isDeleted"`
}

type _FilterModel FilterModel

// NewFilterModel instantiates a new FilterModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFilterModel(createdDate time.Time, createdById string, data WorkItemSearchQueryModel, projectId string, name string, id string, isDeleted bool) *FilterModel {
	this := FilterModel{}
	this.CreatedDate = createdDate
	this.CreatedById = createdById
	this.Data = data
	this.ProjectId = projectId
	this.Name = name
	this.Id = id
	this.IsDeleted = isDeleted
	return &this
}

// NewFilterModelWithDefaults instantiates a new FilterModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFilterModelWithDefaults() *FilterModel {
	this := FilterModel{}
	return &this
}

// GetCreatedDate returns the CreatedDate field value
func (o *FilterModel) GetCreatedDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedDate
}

// GetCreatedDateOk returns a tuple with the CreatedDate field value
// and a boolean to check if the value has been set.
func (o *FilterModel) GetCreatedDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedDate, true
}

// SetCreatedDate sets field value
func (o *FilterModel) SetCreatedDate(v time.Time) {
	o.CreatedDate = v
}

// GetModifiedDate returns the ModifiedDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FilterModel) GetModifiedDate() time.Time {
	if o == nil || IsNil(o.ModifiedDate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.ModifiedDate.Get()
}

// GetModifiedDateOk returns a tuple with the ModifiedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FilterModel) GetModifiedDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ModifiedDate.Get(), o.ModifiedDate.IsSet()
}

// HasModifiedDate returns a boolean if a field has been set.
func (o *FilterModel) HasModifiedDate() bool {
	if o != nil && o.ModifiedDate.IsSet() {
		return true
	}

	return false
}

// SetModifiedDate gets a reference to the given NullableTime and assigns it to the ModifiedDate field.
func (o *FilterModel) SetModifiedDate(v time.Time) {
	o.ModifiedDate.Set(&v)
}
// SetModifiedDateNil sets the value for ModifiedDate to be an explicit nil
func (o *FilterModel) SetModifiedDateNil() {
	o.ModifiedDate.Set(nil)
}

// UnsetModifiedDate ensures that no value is present for ModifiedDate, not even an explicit nil
func (o *FilterModel) UnsetModifiedDate() {
	o.ModifiedDate.Unset()
}

// GetCreatedById returns the CreatedById field value
func (o *FilterModel) GetCreatedById() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedById
}

// GetCreatedByIdOk returns a tuple with the CreatedById field value
// and a boolean to check if the value has been set.
func (o *FilterModel) GetCreatedByIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedById, true
}

// SetCreatedById sets field value
func (o *FilterModel) SetCreatedById(v string) {
	o.CreatedById = v
}

// GetModifiedById returns the ModifiedById field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FilterModel) GetModifiedById() string {
	if o == nil || IsNil(o.ModifiedById.Get()) {
		var ret string
		return ret
	}
	return *o.ModifiedById.Get()
}

// GetModifiedByIdOk returns a tuple with the ModifiedById field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FilterModel) GetModifiedByIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ModifiedById.Get(), o.ModifiedById.IsSet()
}

// HasModifiedById returns a boolean if a field has been set.
func (o *FilterModel) HasModifiedById() bool {
	if o != nil && o.ModifiedById.IsSet() {
		return true
	}

	return false
}

// SetModifiedById gets a reference to the given NullableString and assigns it to the ModifiedById field.
func (o *FilterModel) SetModifiedById(v string) {
	o.ModifiedById.Set(&v)
}
// SetModifiedByIdNil sets the value for ModifiedById to be an explicit nil
func (o *FilterModel) SetModifiedByIdNil() {
	o.ModifiedById.Set(nil)
}

// UnsetModifiedById ensures that no value is present for ModifiedById, not even an explicit nil
func (o *FilterModel) UnsetModifiedById() {
	o.ModifiedById.Unset()
}

// GetData returns the Data field value
func (o *FilterModel) GetData() WorkItemSearchQueryModel {
	if o == nil {
		var ret WorkItemSearchQueryModel
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *FilterModel) GetDataOk() (*WorkItemSearchQueryModel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *FilterModel) SetData(v WorkItemSearchQueryModel) {
	o.Data = v
}

// GetProjectId returns the ProjectId field value
func (o *FilterModel) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *FilterModel) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *FilterModel) SetProjectId(v string) {
	o.ProjectId = v
}

// GetFieldsToShow returns the FieldsToShow field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FilterModel) GetFieldsToShow() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.FieldsToShow
}

// GetFieldsToShowOk returns a tuple with the FieldsToShow field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FilterModel) GetFieldsToShowOk() (*interface{}, bool) {
	if o == nil || IsNil(o.FieldsToShow) {
		return nil, false
	}
	return &o.FieldsToShow, true
}

// HasFieldsToShow returns a boolean if a field has been set.
func (o *FilterModel) HasFieldsToShow() bool {
	if o != nil && !IsNil(o.FieldsToShow) {
		return true
	}

	return false
}

// SetFieldsToShow gets a reference to the given interface{} and assigns it to the FieldsToShow field.
func (o *FilterModel) SetFieldsToShow(v interface{}) {
	o.FieldsToShow = v
}

// GetName returns the Name field value
func (o *FilterModel) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *FilterModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *FilterModel) SetName(v string) {
	o.Name = v
}

// GetId returns the Id field value
func (o *FilterModel) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *FilterModel) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *FilterModel) SetId(v string) {
	o.Id = v
}

// GetIsDeleted returns the IsDeleted field value
func (o *FilterModel) GetIsDeleted() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsDeleted
}

// GetIsDeletedOk returns a tuple with the IsDeleted field value
// and a boolean to check if the value has been set.
func (o *FilterModel) GetIsDeletedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsDeleted, true
}

// SetIsDeleted sets field value
func (o *FilterModel) SetIsDeleted(v bool) {
	o.IsDeleted = v
}

func (o FilterModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FilterModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["createdDate"] = o.CreatedDate
	if o.ModifiedDate.IsSet() {
		toSerialize["modifiedDate"] = o.ModifiedDate.Get()
	}
	toSerialize["createdById"] = o.CreatedById
	if o.ModifiedById.IsSet() {
		toSerialize["modifiedById"] = o.ModifiedById.Get()
	}
	toSerialize["data"] = o.Data
	toSerialize["projectId"] = o.ProjectId
	if o.FieldsToShow != nil {
		toSerialize["fieldsToShow"] = o.FieldsToShow
	}
	toSerialize["name"] = o.Name
	toSerialize["id"] = o.Id
	toSerialize["isDeleted"] = o.IsDeleted
	return toSerialize, nil
}

func (o *FilterModel) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"createdDate",
		"createdById",
		"data",
		"projectId",
		"name",
		"id",
		"isDeleted",
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

	varFilterModel := _FilterModel{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFilterModel)

	if err != nil {
		return err
	}

	*o = FilterModel(varFilterModel)

	return err
}

type NullableFilterModel struct {
	value *FilterModel
	isSet bool
}

func (v NullableFilterModel) Get() *FilterModel {
	return v.value
}

func (v *NullableFilterModel) Set(val *FilterModel) {
	v.value = val
	v.isSet = true
}

func (v NullableFilterModel) IsSet() bool {
	return v.isSet
}

func (v *NullableFilterModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFilterModel(val *FilterModel) *NullableFilterModel {
	return &NullableFilterModel{value: val, isSet: true}
}

func (v NullableFilterModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFilterModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


