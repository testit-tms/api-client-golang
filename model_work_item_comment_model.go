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

// checks if the WorkItemCommentModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WorkItemCommentModel{}

// WorkItemCommentModel struct for WorkItemCommentModel
type WorkItemCommentModel struct {
	Id string `json:"id"`
	Text string `json:"text"`
	User UserWithRankModel `json:"user"`
	CreatedById string `json:"createdById"`
	ModifiedById NullableString `json:"modifiedById,omitempty"`
	CreatedDate time.Time `json:"createdDate"`
	ModifiedDate NullableTime `json:"modifiedDate,omitempty"`
}

type _WorkItemCommentModel WorkItemCommentModel

// NewWorkItemCommentModel instantiates a new WorkItemCommentModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkItemCommentModel(id string, text string, user UserWithRankModel, createdById string, createdDate time.Time) *WorkItemCommentModel {
	this := WorkItemCommentModel{}
	this.Id = id
	this.Text = text
	this.User = user
	this.CreatedById = createdById
	this.CreatedDate = createdDate
	return &this
}

// NewWorkItemCommentModelWithDefaults instantiates a new WorkItemCommentModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkItemCommentModelWithDefaults() *WorkItemCommentModel {
	this := WorkItemCommentModel{}
	return &this
}

// GetId returns the Id field value
func (o *WorkItemCommentModel) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *WorkItemCommentModel) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *WorkItemCommentModel) SetId(v string) {
	o.Id = v
}

// GetText returns the Text field value
func (o *WorkItemCommentModel) GetText() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Text
}

// GetTextOk returns a tuple with the Text field value
// and a boolean to check if the value has been set.
func (o *WorkItemCommentModel) GetTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Text, true
}

// SetText sets field value
func (o *WorkItemCommentModel) SetText(v string) {
	o.Text = v
}

// GetUser returns the User field value
func (o *WorkItemCommentModel) GetUser() UserWithRankModel {
	if o == nil {
		var ret UserWithRankModel
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *WorkItemCommentModel) GetUserOk() (*UserWithRankModel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *WorkItemCommentModel) SetUser(v UserWithRankModel) {
	o.User = v
}

// GetCreatedById returns the CreatedById field value
func (o *WorkItemCommentModel) GetCreatedById() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedById
}

// GetCreatedByIdOk returns a tuple with the CreatedById field value
// and a boolean to check if the value has been set.
func (o *WorkItemCommentModel) GetCreatedByIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedById, true
}

// SetCreatedById sets field value
func (o *WorkItemCommentModel) SetCreatedById(v string) {
	o.CreatedById = v
}

// GetModifiedById returns the ModifiedById field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkItemCommentModel) GetModifiedById() string {
	if o == nil || IsNil(o.ModifiedById.Get()) {
		var ret string
		return ret
	}
	return *o.ModifiedById.Get()
}

// GetModifiedByIdOk returns a tuple with the ModifiedById field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkItemCommentModel) GetModifiedByIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ModifiedById.Get(), o.ModifiedById.IsSet()
}

// HasModifiedById returns a boolean if a field has been set.
func (o *WorkItemCommentModel) HasModifiedById() bool {
	if o != nil && o.ModifiedById.IsSet() {
		return true
	}

	return false
}

// SetModifiedById gets a reference to the given NullableString and assigns it to the ModifiedById field.
func (o *WorkItemCommentModel) SetModifiedById(v string) {
	o.ModifiedById.Set(&v)
}
// SetModifiedByIdNil sets the value for ModifiedById to be an explicit nil
func (o *WorkItemCommentModel) SetModifiedByIdNil() {
	o.ModifiedById.Set(nil)
}

// UnsetModifiedById ensures that no value is present for ModifiedById, not even an explicit nil
func (o *WorkItemCommentModel) UnsetModifiedById() {
	o.ModifiedById.Unset()
}

// GetCreatedDate returns the CreatedDate field value
func (o *WorkItemCommentModel) GetCreatedDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedDate
}

// GetCreatedDateOk returns a tuple with the CreatedDate field value
// and a boolean to check if the value has been set.
func (o *WorkItemCommentModel) GetCreatedDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedDate, true
}

// SetCreatedDate sets field value
func (o *WorkItemCommentModel) SetCreatedDate(v time.Time) {
	o.CreatedDate = v
}

// GetModifiedDate returns the ModifiedDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkItemCommentModel) GetModifiedDate() time.Time {
	if o == nil || IsNil(o.ModifiedDate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.ModifiedDate.Get()
}

// GetModifiedDateOk returns a tuple with the ModifiedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkItemCommentModel) GetModifiedDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ModifiedDate.Get(), o.ModifiedDate.IsSet()
}

// HasModifiedDate returns a boolean if a field has been set.
func (o *WorkItemCommentModel) HasModifiedDate() bool {
	if o != nil && o.ModifiedDate.IsSet() {
		return true
	}

	return false
}

// SetModifiedDate gets a reference to the given NullableTime and assigns it to the ModifiedDate field.
func (o *WorkItemCommentModel) SetModifiedDate(v time.Time) {
	o.ModifiedDate.Set(&v)
}
// SetModifiedDateNil sets the value for ModifiedDate to be an explicit nil
func (o *WorkItemCommentModel) SetModifiedDateNil() {
	o.ModifiedDate.Set(nil)
}

// UnsetModifiedDate ensures that no value is present for ModifiedDate, not even an explicit nil
func (o *WorkItemCommentModel) UnsetModifiedDate() {
	o.ModifiedDate.Unset()
}

func (o WorkItemCommentModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkItemCommentModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["text"] = o.Text
	toSerialize["user"] = o.User
	toSerialize["createdById"] = o.CreatedById
	if o.ModifiedById.IsSet() {
		toSerialize["modifiedById"] = o.ModifiedById.Get()
	}
	toSerialize["createdDate"] = o.CreatedDate
	if o.ModifiedDate.IsSet() {
		toSerialize["modifiedDate"] = o.ModifiedDate.Get()
	}
	return toSerialize, nil
}

func (o *WorkItemCommentModel) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"text",
		"user",
		"createdById",
		"createdDate",
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

	varWorkItemCommentModel := _WorkItemCommentModel{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varWorkItemCommentModel)

	if err != nil {
		return err
	}

	*o = WorkItemCommentModel(varWorkItemCommentModel)

	return err
}

type NullableWorkItemCommentModel struct {
	value *WorkItemCommentModel
	isSet bool
}

func (v NullableWorkItemCommentModel) Get() *WorkItemCommentModel {
	return v.value
}

func (v *NullableWorkItemCommentModel) Set(val *WorkItemCommentModel) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkItemCommentModel) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkItemCommentModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkItemCommentModel(val *WorkItemCommentModel) *NullableWorkItemCommentModel {
	return &NullableWorkItemCommentModel{value: val, isSet: true}
}

func (v NullableWorkItemCommentModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkItemCommentModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


