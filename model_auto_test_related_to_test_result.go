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

// checks if the AutoTestRelatedToTestResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AutoTestRelatedToTestResult{}

// AutoTestRelatedToTestResult struct for AutoTestRelatedToTestResult
type AutoTestRelatedToTestResult struct {
	// This property is used to set autotest identifier from client system
	ExternalId string `json:"externalId"`
	Links []LinkModel `json:"links,omitempty"`
	// This property is used to link autotest with project
	ProjectId string `json:"projectId"`
	Name string `json:"name"`
	Namespace NullableString `json:"namespace,omitempty"`
	Classname NullableString `json:"classname,omitempty"`
	Steps []AutoTestStepModel `json:"steps,omitempty"`
	Setup []AutoTestStepModel `json:"setup,omitempty"`
	Teardown []AutoTestStepModel `json:"teardown,omitempty"`
	GlobalId int64 `json:"globalId"`
	CreatedDate NullableTime `json:"createdDate,omitempty"`
	ModifiedDate NullableTime `json:"modifiedDate,omitempty"`
	CreatedById string `json:"createdById"`
	ModifiedById NullableString `json:"modifiedById,omitempty"`
	Labels []LabelShortModel `json:"labels,omitempty"`
	ExternalKey NullableString `json:"externalKey,omitempty"`
	// Unique ID of the entity
	Id string `json:"id"`
	// Indicates if the entity is deleted
	IsDeleted bool `json:"isDeleted"`
}

type _AutoTestRelatedToTestResult AutoTestRelatedToTestResult

// NewAutoTestRelatedToTestResult instantiates a new AutoTestRelatedToTestResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAutoTestRelatedToTestResult(externalId string, projectId string, name string, globalId int64, createdById string, id string, isDeleted bool) *AutoTestRelatedToTestResult {
	this := AutoTestRelatedToTestResult{}
	this.ExternalId = externalId
	this.ProjectId = projectId
	this.Name = name
	this.GlobalId = globalId
	this.CreatedById = createdById
	this.Id = id
	this.IsDeleted = isDeleted
	return &this
}

// NewAutoTestRelatedToTestResultWithDefaults instantiates a new AutoTestRelatedToTestResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAutoTestRelatedToTestResultWithDefaults() *AutoTestRelatedToTestResult {
	this := AutoTestRelatedToTestResult{}
	return &this
}

// GetExternalId returns the ExternalId field value
func (o *AutoTestRelatedToTestResult) GetExternalId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value
// and a boolean to check if the value has been set.
func (o *AutoTestRelatedToTestResult) GetExternalIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExternalId, true
}

// SetExternalId sets field value
func (o *AutoTestRelatedToTestResult) SetExternalId(v string) {
	o.ExternalId = v
}

// GetLinks returns the Links field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AutoTestRelatedToTestResult) GetLinks() []LinkModel {
	if o == nil {
		var ret []LinkModel
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AutoTestRelatedToTestResult) GetLinksOk() ([]LinkModel, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *AutoTestRelatedToTestResult) HasLinks() bool {
	if o != nil && IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []LinkModel and assigns it to the Links field.
func (o *AutoTestRelatedToTestResult) SetLinks(v []LinkModel) {
	o.Links = v
}

// GetProjectId returns the ProjectId field value
func (o *AutoTestRelatedToTestResult) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *AutoTestRelatedToTestResult) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *AutoTestRelatedToTestResult) SetProjectId(v string) {
	o.ProjectId = v
}

// GetName returns the Name field value
func (o *AutoTestRelatedToTestResult) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AutoTestRelatedToTestResult) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AutoTestRelatedToTestResult) SetName(v string) {
	o.Name = v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AutoTestRelatedToTestResult) GetNamespace() string {
	if o == nil || IsNil(o.Namespace.Get()) {
		var ret string
		return ret
	}
	return *o.Namespace.Get()
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AutoTestRelatedToTestResult) GetNamespaceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Namespace.Get(), o.Namespace.IsSet()
}

// HasNamespace returns a boolean if a field has been set.
func (o *AutoTestRelatedToTestResult) HasNamespace() bool {
	if o != nil && o.Namespace.IsSet() {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given NullableString and assigns it to the Namespace field.
func (o *AutoTestRelatedToTestResult) SetNamespace(v string) {
	o.Namespace.Set(&v)
}
// SetNamespaceNil sets the value for Namespace to be an explicit nil
func (o *AutoTestRelatedToTestResult) SetNamespaceNil() {
	o.Namespace.Set(nil)
}

// UnsetNamespace ensures that no value is present for Namespace, not even an explicit nil
func (o *AutoTestRelatedToTestResult) UnsetNamespace() {
	o.Namespace.Unset()
}

// GetClassname returns the Classname field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AutoTestRelatedToTestResult) GetClassname() string {
	if o == nil || IsNil(o.Classname.Get()) {
		var ret string
		return ret
	}
	return *o.Classname.Get()
}

// GetClassnameOk returns a tuple with the Classname field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AutoTestRelatedToTestResult) GetClassnameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Classname.Get(), o.Classname.IsSet()
}

// HasClassname returns a boolean if a field has been set.
func (o *AutoTestRelatedToTestResult) HasClassname() bool {
	if o != nil && o.Classname.IsSet() {
		return true
	}

	return false
}

// SetClassname gets a reference to the given NullableString and assigns it to the Classname field.
func (o *AutoTestRelatedToTestResult) SetClassname(v string) {
	o.Classname.Set(&v)
}
// SetClassnameNil sets the value for Classname to be an explicit nil
func (o *AutoTestRelatedToTestResult) SetClassnameNil() {
	o.Classname.Set(nil)
}

// UnsetClassname ensures that no value is present for Classname, not even an explicit nil
func (o *AutoTestRelatedToTestResult) UnsetClassname() {
	o.Classname.Unset()
}

// GetSteps returns the Steps field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AutoTestRelatedToTestResult) GetSteps() []AutoTestStepModel {
	if o == nil {
		var ret []AutoTestStepModel
		return ret
	}
	return o.Steps
}

// GetStepsOk returns a tuple with the Steps field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AutoTestRelatedToTestResult) GetStepsOk() ([]AutoTestStepModel, bool) {
	if o == nil || IsNil(o.Steps) {
		return nil, false
	}
	return o.Steps, true
}

// HasSteps returns a boolean if a field has been set.
func (o *AutoTestRelatedToTestResult) HasSteps() bool {
	if o != nil && IsNil(o.Steps) {
		return true
	}

	return false
}

// SetSteps gets a reference to the given []AutoTestStepModel and assigns it to the Steps field.
func (o *AutoTestRelatedToTestResult) SetSteps(v []AutoTestStepModel) {
	o.Steps = v
}

// GetSetup returns the Setup field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AutoTestRelatedToTestResult) GetSetup() []AutoTestStepModel {
	if o == nil {
		var ret []AutoTestStepModel
		return ret
	}
	return o.Setup
}

// GetSetupOk returns a tuple with the Setup field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AutoTestRelatedToTestResult) GetSetupOk() ([]AutoTestStepModel, bool) {
	if o == nil || IsNil(o.Setup) {
		return nil, false
	}
	return o.Setup, true
}

// HasSetup returns a boolean if a field has been set.
func (o *AutoTestRelatedToTestResult) HasSetup() bool {
	if o != nil && IsNil(o.Setup) {
		return true
	}

	return false
}

// SetSetup gets a reference to the given []AutoTestStepModel and assigns it to the Setup field.
func (o *AutoTestRelatedToTestResult) SetSetup(v []AutoTestStepModel) {
	o.Setup = v
}

// GetTeardown returns the Teardown field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AutoTestRelatedToTestResult) GetTeardown() []AutoTestStepModel {
	if o == nil {
		var ret []AutoTestStepModel
		return ret
	}
	return o.Teardown
}

// GetTeardownOk returns a tuple with the Teardown field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AutoTestRelatedToTestResult) GetTeardownOk() ([]AutoTestStepModel, bool) {
	if o == nil || IsNil(o.Teardown) {
		return nil, false
	}
	return o.Teardown, true
}

// HasTeardown returns a boolean if a field has been set.
func (o *AutoTestRelatedToTestResult) HasTeardown() bool {
	if o != nil && IsNil(o.Teardown) {
		return true
	}

	return false
}

// SetTeardown gets a reference to the given []AutoTestStepModel and assigns it to the Teardown field.
func (o *AutoTestRelatedToTestResult) SetTeardown(v []AutoTestStepModel) {
	o.Teardown = v
}

// GetGlobalId returns the GlobalId field value
func (o *AutoTestRelatedToTestResult) GetGlobalId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.GlobalId
}

// GetGlobalIdOk returns a tuple with the GlobalId field value
// and a boolean to check if the value has been set.
func (o *AutoTestRelatedToTestResult) GetGlobalIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GlobalId, true
}

// SetGlobalId sets field value
func (o *AutoTestRelatedToTestResult) SetGlobalId(v int64) {
	o.GlobalId = v
}

// GetCreatedDate returns the CreatedDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AutoTestRelatedToTestResult) GetCreatedDate() time.Time {
	if o == nil || IsNil(o.CreatedDate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.CreatedDate.Get()
}

// GetCreatedDateOk returns a tuple with the CreatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AutoTestRelatedToTestResult) GetCreatedDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedDate.Get(), o.CreatedDate.IsSet()
}

// HasCreatedDate returns a boolean if a field has been set.
func (o *AutoTestRelatedToTestResult) HasCreatedDate() bool {
	if o != nil && o.CreatedDate.IsSet() {
		return true
	}

	return false
}

// SetCreatedDate gets a reference to the given NullableTime and assigns it to the CreatedDate field.
func (o *AutoTestRelatedToTestResult) SetCreatedDate(v time.Time) {
	o.CreatedDate.Set(&v)
}
// SetCreatedDateNil sets the value for CreatedDate to be an explicit nil
func (o *AutoTestRelatedToTestResult) SetCreatedDateNil() {
	o.CreatedDate.Set(nil)
}

// UnsetCreatedDate ensures that no value is present for CreatedDate, not even an explicit nil
func (o *AutoTestRelatedToTestResult) UnsetCreatedDate() {
	o.CreatedDate.Unset()
}

// GetModifiedDate returns the ModifiedDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AutoTestRelatedToTestResult) GetModifiedDate() time.Time {
	if o == nil || IsNil(o.ModifiedDate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.ModifiedDate.Get()
}

// GetModifiedDateOk returns a tuple with the ModifiedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AutoTestRelatedToTestResult) GetModifiedDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ModifiedDate.Get(), o.ModifiedDate.IsSet()
}

// HasModifiedDate returns a boolean if a field has been set.
func (o *AutoTestRelatedToTestResult) HasModifiedDate() bool {
	if o != nil && o.ModifiedDate.IsSet() {
		return true
	}

	return false
}

// SetModifiedDate gets a reference to the given NullableTime and assigns it to the ModifiedDate field.
func (o *AutoTestRelatedToTestResult) SetModifiedDate(v time.Time) {
	o.ModifiedDate.Set(&v)
}
// SetModifiedDateNil sets the value for ModifiedDate to be an explicit nil
func (o *AutoTestRelatedToTestResult) SetModifiedDateNil() {
	o.ModifiedDate.Set(nil)
}

// UnsetModifiedDate ensures that no value is present for ModifiedDate, not even an explicit nil
func (o *AutoTestRelatedToTestResult) UnsetModifiedDate() {
	o.ModifiedDate.Unset()
}

// GetCreatedById returns the CreatedById field value
func (o *AutoTestRelatedToTestResult) GetCreatedById() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedById
}

// GetCreatedByIdOk returns a tuple with the CreatedById field value
// and a boolean to check if the value has been set.
func (o *AutoTestRelatedToTestResult) GetCreatedByIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedById, true
}

// SetCreatedById sets field value
func (o *AutoTestRelatedToTestResult) SetCreatedById(v string) {
	o.CreatedById = v
}

// GetModifiedById returns the ModifiedById field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AutoTestRelatedToTestResult) GetModifiedById() string {
	if o == nil || IsNil(o.ModifiedById.Get()) {
		var ret string
		return ret
	}
	return *o.ModifiedById.Get()
}

// GetModifiedByIdOk returns a tuple with the ModifiedById field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AutoTestRelatedToTestResult) GetModifiedByIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ModifiedById.Get(), o.ModifiedById.IsSet()
}

// HasModifiedById returns a boolean if a field has been set.
func (o *AutoTestRelatedToTestResult) HasModifiedById() bool {
	if o != nil && o.ModifiedById.IsSet() {
		return true
	}

	return false
}

// SetModifiedById gets a reference to the given NullableString and assigns it to the ModifiedById field.
func (o *AutoTestRelatedToTestResult) SetModifiedById(v string) {
	o.ModifiedById.Set(&v)
}
// SetModifiedByIdNil sets the value for ModifiedById to be an explicit nil
func (o *AutoTestRelatedToTestResult) SetModifiedByIdNil() {
	o.ModifiedById.Set(nil)
}

// UnsetModifiedById ensures that no value is present for ModifiedById, not even an explicit nil
func (o *AutoTestRelatedToTestResult) UnsetModifiedById() {
	o.ModifiedById.Unset()
}

// GetLabels returns the Labels field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AutoTestRelatedToTestResult) GetLabels() []LabelShortModel {
	if o == nil {
		var ret []LabelShortModel
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AutoTestRelatedToTestResult) GetLabelsOk() ([]LabelShortModel, bool) {
	if o == nil || IsNil(o.Labels) {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *AutoTestRelatedToTestResult) HasLabels() bool {
	if o != nil && IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []LabelShortModel and assigns it to the Labels field.
func (o *AutoTestRelatedToTestResult) SetLabels(v []LabelShortModel) {
	o.Labels = v
}

// GetExternalKey returns the ExternalKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AutoTestRelatedToTestResult) GetExternalKey() string {
	if o == nil || IsNil(o.ExternalKey.Get()) {
		var ret string
		return ret
	}
	return *o.ExternalKey.Get()
}

// GetExternalKeyOk returns a tuple with the ExternalKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AutoTestRelatedToTestResult) GetExternalKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExternalKey.Get(), o.ExternalKey.IsSet()
}

// HasExternalKey returns a boolean if a field has been set.
func (o *AutoTestRelatedToTestResult) HasExternalKey() bool {
	if o != nil && o.ExternalKey.IsSet() {
		return true
	}

	return false
}

// SetExternalKey gets a reference to the given NullableString and assigns it to the ExternalKey field.
func (o *AutoTestRelatedToTestResult) SetExternalKey(v string) {
	o.ExternalKey.Set(&v)
}
// SetExternalKeyNil sets the value for ExternalKey to be an explicit nil
func (o *AutoTestRelatedToTestResult) SetExternalKeyNil() {
	o.ExternalKey.Set(nil)
}

// UnsetExternalKey ensures that no value is present for ExternalKey, not even an explicit nil
func (o *AutoTestRelatedToTestResult) UnsetExternalKey() {
	o.ExternalKey.Unset()
}

// GetId returns the Id field value
func (o *AutoTestRelatedToTestResult) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AutoTestRelatedToTestResult) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AutoTestRelatedToTestResult) SetId(v string) {
	o.Id = v
}

// GetIsDeleted returns the IsDeleted field value
func (o *AutoTestRelatedToTestResult) GetIsDeleted() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsDeleted
}

// GetIsDeletedOk returns a tuple with the IsDeleted field value
// and a boolean to check if the value has been set.
func (o *AutoTestRelatedToTestResult) GetIsDeletedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsDeleted, true
}

// SetIsDeleted sets field value
func (o *AutoTestRelatedToTestResult) SetIsDeleted(v bool) {
	o.IsDeleted = v
}

func (o AutoTestRelatedToTestResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AutoTestRelatedToTestResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["externalId"] = o.ExternalId
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	toSerialize["projectId"] = o.ProjectId
	toSerialize["name"] = o.Name
	if o.Namespace.IsSet() {
		toSerialize["namespace"] = o.Namespace.Get()
	}
	if o.Classname.IsSet() {
		toSerialize["classname"] = o.Classname.Get()
	}
	if o.Steps != nil {
		toSerialize["steps"] = o.Steps
	}
	if o.Setup != nil {
		toSerialize["setup"] = o.Setup
	}
	if o.Teardown != nil {
		toSerialize["teardown"] = o.Teardown
	}
	toSerialize["globalId"] = o.GlobalId
	if o.CreatedDate.IsSet() {
		toSerialize["createdDate"] = o.CreatedDate.Get()
	}
	if o.ModifiedDate.IsSet() {
		toSerialize["modifiedDate"] = o.ModifiedDate.Get()
	}
	toSerialize["createdById"] = o.CreatedById
	if o.ModifiedById.IsSet() {
		toSerialize["modifiedById"] = o.ModifiedById.Get()
	}
	if o.Labels != nil {
		toSerialize["labels"] = o.Labels
	}
	if o.ExternalKey.IsSet() {
		toSerialize["externalKey"] = o.ExternalKey.Get()
	}
	toSerialize["id"] = o.Id
	toSerialize["isDeleted"] = o.IsDeleted
	return toSerialize, nil
}

func (o *AutoTestRelatedToTestResult) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"externalId",
		"projectId",
		"name",
		"globalId",
		"createdById",
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

	varAutoTestRelatedToTestResult := _AutoTestRelatedToTestResult{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAutoTestRelatedToTestResult)

	if err != nil {
		return err
	}

	*o = AutoTestRelatedToTestResult(varAutoTestRelatedToTestResult)

	return err
}

type NullableAutoTestRelatedToTestResult struct {
	value *AutoTestRelatedToTestResult
	isSet bool
}

func (v NullableAutoTestRelatedToTestResult) Get() *AutoTestRelatedToTestResult {
	return v.value
}

func (v *NullableAutoTestRelatedToTestResult) Set(val *AutoTestRelatedToTestResult) {
	v.value = val
	v.isSet = true
}

func (v NullableAutoTestRelatedToTestResult) IsSet() bool {
	return v.isSet
}

func (v *NullableAutoTestRelatedToTestResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAutoTestRelatedToTestResult(val *AutoTestRelatedToTestResult) *NullableAutoTestRelatedToTestResult {
	return &NullableAutoTestRelatedToTestResult{value: val, isSet: true}
}

func (v NullableAutoTestRelatedToTestResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAutoTestRelatedToTestResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


