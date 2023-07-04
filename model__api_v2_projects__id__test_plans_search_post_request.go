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

// checks if the ApiV2ProjectsIdTestPlansSearchPostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiV2ProjectsIdTestPlansSearchPostRequest{}

// ApiV2ProjectsIdTestPlansSearchPostRequest struct for ApiV2ProjectsIdTestPlansSearchPostRequest
type ApiV2ProjectsIdTestPlansSearchPostRequest struct {
	Name NullableString `json:"name,omitempty"`
	Description NullableString `json:"description,omitempty"`
	Build NullableString `json:"build,omitempty"`
	ProductName NullableString `json:"productName,omitempty"`
	Status []TestPlanStatusModel `json:"status,omitempty"`
	GlobalIds []int64 `json:"globalIds,omitempty"`
	IsLocked NullableBool `json:"isLocked,omitempty"`
	LockedDate NullableDateTimeRangeSelectorModel `json:"lockedDate,omitempty"`
	AutomaticDurationTimer []bool `json:"automaticDurationTimer,omitempty"`
	CreatedByIds []string `json:"createdByIds,omitempty"`
	CreatedDate NullableDateTimeRangeSelectorModel `json:"createdDate,omitempty"`
	StartDate NullableDateTimeRangeSelectorModel `json:"startDate,omitempty"`
	EndDate NullableDateTimeRangeSelectorModel `json:"endDate,omitempty"`
	TagNames []string `json:"tagNames,omitempty"`
	Attributes map[string][]string `json:"attributes,omitempty"`
	IsDeleted NullableBool `json:"isDeleted,omitempty"`
}

// NewApiV2ProjectsIdTestPlansSearchPostRequest instantiates a new ApiV2ProjectsIdTestPlansSearchPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiV2ProjectsIdTestPlansSearchPostRequest() *ApiV2ProjectsIdTestPlansSearchPostRequest {
	this := ApiV2ProjectsIdTestPlansSearchPostRequest{}
	return &this
}

// NewApiV2ProjectsIdTestPlansSearchPostRequestWithDefaults instantiates a new ApiV2ProjectsIdTestPlansSearchPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiV2ProjectsIdTestPlansSearchPostRequestWithDefaults() *ApiV2ProjectsIdTestPlansSearchPostRequest {
	this := ApiV2ProjectsIdTestPlansSearchPostRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) UnsetName() {
	o.Name.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) UnsetDescription() {
	o.Description.Unset()
}

// GetBuild returns the Build field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) GetBuild() string {
	if o == nil || IsNil(o.Build.Get()) {
		var ret string
		return ret
	}
	return *o.Build.Get()
}

// GetBuildOk returns a tuple with the Build field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) GetBuildOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Build.Get(), o.Build.IsSet()
}

// HasBuild returns a boolean if a field has been set.
func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) HasBuild() bool {
	if o != nil && o.Build.IsSet() {
		return true
	}

	return false
}

// SetBuild gets a reference to the given NullableString and assigns it to the Build field.
func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) SetBuild(v string) {
	o.Build.Set(&v)
}
// SetBuildNil sets the value for Build to be an explicit nil
func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) SetBuildNil() {
	o.Build.Set(nil)
}

// UnsetBuild ensures that no value is present for Build, not even an explicit nil
func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) UnsetBuild() {
	o.Build.Unset()
}

// GetProductName returns the ProductName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) GetProductName() string {
	if o == nil || IsNil(o.ProductName.Get()) {
		var ret string
		return ret
	}
	return *o.ProductName.Get()
}

// GetProductNameOk returns a tuple with the ProductName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) GetProductNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProductName.Get(), o.ProductName.IsSet()
}

// HasProductName returns a boolean if a field has been set.
func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) HasProductName() bool {
	if o != nil && o.ProductName.IsSet() {
		return true
	}

	return false
}

// SetProductName gets a reference to the given NullableString and assigns it to the ProductName field.
func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) SetProductName(v string) {
	o.ProductName.Set(&v)
}
// SetProductNameNil sets the value for ProductName to be an explicit nil
func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) SetProductNameNil() {
	o.ProductName.Set(nil)
}

// UnsetProductName ensures that no value is present for ProductName, not even an explicit nil
func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) UnsetProductName() {
	o.ProductName.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) GetStatus() []TestPlanStatusModel {
	if o == nil {
		var ret []TestPlanStatusModel
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) GetStatusOk() ([]TestPlanStatusModel, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) HasStatus() bool {
	if o != nil && IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given []TestPlanStatusModel and assigns it to the Status field.
func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) SetStatus(v []TestPlanStatusModel) {
	o.Status = v
}

// GetGlobalIds returns the GlobalIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) GetGlobalIds() []int64 {
	if o == nil {
		var ret []int64
		return ret
	}
	return o.GlobalIds
}

// GetGlobalIdsOk returns a tuple with the GlobalIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) GetGlobalIdsOk() ([]int64, bool) {
	if o == nil || IsNil(o.GlobalIds) {
		return nil, false
	}
	return o.GlobalIds, true
}

// HasGlobalIds returns a boolean if a field has been set.
func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) HasGlobalIds() bool {
	if o != nil && IsNil(o.GlobalIds) {
		return true
	}

	return false
}

// SetGlobalIds gets a reference to the given []int64 and assigns it to the GlobalIds field.
func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) SetGlobalIds(v []int64) {
	o.GlobalIds = v
}

// GetIsLocked returns the IsLocked field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) GetIsLocked() bool {
	if o == nil || IsNil(o.IsLocked.Get()) {
		var ret bool
		return ret
	}
	return *o.IsLocked.Get()
}

// GetIsLockedOk returns a tuple with the IsLocked field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) GetIsLockedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsLocked.Get(), o.IsLocked.IsSet()
}

// HasIsLocked returns a boolean if a field has been set.
func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) HasIsLocked() bool {
	if o != nil && o.IsLocked.IsSet() {
		return true
	}

	return false
}

// SetIsLocked gets a reference to the given NullableBool and assigns it to the IsLocked field.
func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) SetIsLocked(v bool) {
	o.IsLocked.Set(&v)
}
// SetIsLockedNil sets the value for IsLocked to be an explicit nil
func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) SetIsLockedNil() {
	o.IsLocked.Set(nil)
}

// UnsetIsLocked ensures that no value is present for IsLocked, not even an explicit nil
func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) UnsetIsLocked() {
	o.IsLocked.Unset()
}

// GetLockedDate returns the LockedDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) GetLockedDate() DateTimeRangeSelectorModel {
	if o == nil || IsNil(o.LockedDate.Get()) {
		var ret DateTimeRangeSelectorModel
		return ret
	}
	return *o.LockedDate.Get()
}

// GetLockedDateOk returns a tuple with the LockedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) GetLockedDateOk() (*DateTimeRangeSelectorModel, bool) {
	if o == nil {
		return nil, false
	}
	return o.LockedDate.Get(), o.LockedDate.IsSet()
}

// HasLockedDate returns a boolean if a field has been set.
func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) HasLockedDate() bool {
	if o != nil && o.LockedDate.IsSet() {
		return true
	}

	return false
}

// SetLockedDate gets a reference to the given NullableDateTimeRangeSelectorModel and assigns it to the LockedDate field.
func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) SetLockedDate(v DateTimeRangeSelectorModel) {
	o.LockedDate.Set(&v)
}
// SetLockedDateNil sets the value for LockedDate to be an explicit nil
func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) SetLockedDateNil() {
	o.LockedDate.Set(nil)
}

// UnsetLockedDate ensures that no value is present for LockedDate, not even an explicit nil
func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) UnsetLockedDate() {
	o.LockedDate.Unset()
}

// GetAutomaticDurationTimer returns the AutomaticDurationTimer field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) GetAutomaticDurationTimer() []bool {
	if o == nil {
		var ret []bool
		return ret
	}
	return o.AutomaticDurationTimer
}

// GetAutomaticDurationTimerOk returns a tuple with the AutomaticDurationTimer field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) GetAutomaticDurationTimerOk() ([]bool, bool) {
	if o == nil || IsNil(o.AutomaticDurationTimer) {
		return nil, false
	}
	return o.AutomaticDurationTimer, true
}

// HasAutomaticDurationTimer returns a boolean if a field has been set.
func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) HasAutomaticDurationTimer() bool {
	if o != nil && IsNil(o.AutomaticDurationTimer) {
		return true
	}

	return false
}

// SetAutomaticDurationTimer gets a reference to the given []bool and assigns it to the AutomaticDurationTimer field.
func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) SetAutomaticDurationTimer(v []bool) {
	o.AutomaticDurationTimer = v
}

// GetCreatedByIds returns the CreatedByIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) GetCreatedByIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.CreatedByIds
}

// GetCreatedByIdsOk returns a tuple with the CreatedByIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) GetCreatedByIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.CreatedByIds) {
		return nil, false
	}
	return o.CreatedByIds, true
}

// HasCreatedByIds returns a boolean if a field has been set.
func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) HasCreatedByIds() bool {
	if o != nil && IsNil(o.CreatedByIds) {
		return true
	}

	return false
}

// SetCreatedByIds gets a reference to the given []string and assigns it to the CreatedByIds field.
func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) SetCreatedByIds(v []string) {
	o.CreatedByIds = v
}

// GetCreatedDate returns the CreatedDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) GetCreatedDate() DateTimeRangeSelectorModel {
	if o == nil || IsNil(o.CreatedDate.Get()) {
		var ret DateTimeRangeSelectorModel
		return ret
	}
	return *o.CreatedDate.Get()
}

// GetCreatedDateOk returns a tuple with the CreatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) GetCreatedDateOk() (*DateTimeRangeSelectorModel, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedDate.Get(), o.CreatedDate.IsSet()
}

// HasCreatedDate returns a boolean if a field has been set.
func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) HasCreatedDate() bool {
	if o != nil && o.CreatedDate.IsSet() {
		return true
	}

	return false
}

// SetCreatedDate gets a reference to the given NullableDateTimeRangeSelectorModel and assigns it to the CreatedDate field.
func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) SetCreatedDate(v DateTimeRangeSelectorModel) {
	o.CreatedDate.Set(&v)
}
// SetCreatedDateNil sets the value for CreatedDate to be an explicit nil
func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) SetCreatedDateNil() {
	o.CreatedDate.Set(nil)
}

// UnsetCreatedDate ensures that no value is present for CreatedDate, not even an explicit nil
func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) UnsetCreatedDate() {
	o.CreatedDate.Unset()
}

// GetStartDate returns the StartDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) GetStartDate() DateTimeRangeSelectorModel {
	if o == nil || IsNil(o.StartDate.Get()) {
		var ret DateTimeRangeSelectorModel
		return ret
	}
	return *o.StartDate.Get()
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) GetStartDateOk() (*DateTimeRangeSelectorModel, bool) {
	if o == nil {
		return nil, false
	}
	return o.StartDate.Get(), o.StartDate.IsSet()
}

// HasStartDate returns a boolean if a field has been set.
func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) HasStartDate() bool {
	if o != nil && o.StartDate.IsSet() {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given NullableDateTimeRangeSelectorModel and assigns it to the StartDate field.
func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) SetStartDate(v DateTimeRangeSelectorModel) {
	o.StartDate.Set(&v)
}
// SetStartDateNil sets the value for StartDate to be an explicit nil
func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) SetStartDateNil() {
	o.StartDate.Set(nil)
}

// UnsetStartDate ensures that no value is present for StartDate, not even an explicit nil
func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) UnsetStartDate() {
	o.StartDate.Unset()
}

// GetEndDate returns the EndDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) GetEndDate() DateTimeRangeSelectorModel {
	if o == nil || IsNil(o.EndDate.Get()) {
		var ret DateTimeRangeSelectorModel
		return ret
	}
	return *o.EndDate.Get()
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) GetEndDateOk() (*DateTimeRangeSelectorModel, bool) {
	if o == nil {
		return nil, false
	}
	return o.EndDate.Get(), o.EndDate.IsSet()
}

// HasEndDate returns a boolean if a field has been set.
func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) HasEndDate() bool {
	if o != nil && o.EndDate.IsSet() {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given NullableDateTimeRangeSelectorModel and assigns it to the EndDate field.
func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) SetEndDate(v DateTimeRangeSelectorModel) {
	o.EndDate.Set(&v)
}
// SetEndDateNil sets the value for EndDate to be an explicit nil
func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) SetEndDateNil() {
	o.EndDate.Set(nil)
}

// UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) UnsetEndDate() {
	o.EndDate.Unset()
}

// GetTagNames returns the TagNames field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) GetTagNames() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.TagNames
}

// GetTagNamesOk returns a tuple with the TagNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) GetTagNamesOk() ([]string, bool) {
	if o == nil || IsNil(o.TagNames) {
		return nil, false
	}
	return o.TagNames, true
}

// HasTagNames returns a boolean if a field has been set.
func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) HasTagNames() bool {
	if o != nil && IsNil(o.TagNames) {
		return true
	}

	return false
}

// SetTagNames gets a reference to the given []string and assigns it to the TagNames field.
func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) SetTagNames(v []string) {
	o.TagNames = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) GetAttributes() map[string][]string {
	if o == nil {
		var ret map[string][]string
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) GetAttributesOk() (*map[string][]string, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return &o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) HasAttributes() bool {
	if o != nil && IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string][]string and assigns it to the Attributes field.
func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) SetAttributes(v map[string][]string) {
	o.Attributes = v
}

// GetIsDeleted returns the IsDeleted field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) GetIsDeleted() bool {
	if o == nil || IsNil(o.IsDeleted.Get()) {
		var ret bool
		return ret
	}
	return *o.IsDeleted.Get()
}

// GetIsDeletedOk returns a tuple with the IsDeleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) GetIsDeletedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsDeleted.Get(), o.IsDeleted.IsSet()
}

// HasIsDeleted returns a boolean if a field has been set.
func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) HasIsDeleted() bool {
	if o != nil && o.IsDeleted.IsSet() {
		return true
	}

	return false
}

// SetIsDeleted gets a reference to the given NullableBool and assigns it to the IsDeleted field.
func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) SetIsDeleted(v bool) {
	o.IsDeleted.Set(&v)
}
// SetIsDeletedNil sets the value for IsDeleted to be an explicit nil
func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) SetIsDeletedNil() {
	o.IsDeleted.Set(nil)
}

// UnsetIsDeleted ensures that no value is present for IsDeleted, not even an explicit nil
func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) UnsetIsDeleted() {
	o.IsDeleted.Unset()
}

func (o ApiV2ProjectsIdTestPlansSearchPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiV2ProjectsIdTestPlansSearchPostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.Build.IsSet() {
		toSerialize["build"] = o.Build.Get()
	}
	if o.ProductName.IsSet() {
		toSerialize["productName"] = o.ProductName.Get()
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.GlobalIds != nil {
		toSerialize["globalIds"] = o.GlobalIds
	}
	if o.IsLocked.IsSet() {
		toSerialize["isLocked"] = o.IsLocked.Get()
	}
	if o.LockedDate.IsSet() {
		toSerialize["lockedDate"] = o.LockedDate.Get()
	}
	if o.AutomaticDurationTimer != nil {
		toSerialize["automaticDurationTimer"] = o.AutomaticDurationTimer
	}
	if o.CreatedByIds != nil {
		toSerialize["createdByIds"] = o.CreatedByIds
	}
	if o.CreatedDate.IsSet() {
		toSerialize["createdDate"] = o.CreatedDate.Get()
	}
	if o.StartDate.IsSet() {
		toSerialize["startDate"] = o.StartDate.Get()
	}
	if o.EndDate.IsSet() {
		toSerialize["endDate"] = o.EndDate.Get()
	}
	if o.TagNames != nil {
		toSerialize["tagNames"] = o.TagNames
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	if o.IsDeleted.IsSet() {
		toSerialize["isDeleted"] = o.IsDeleted.Get()
	}
	return toSerialize, nil
}

type NullableApiV2ProjectsIdTestPlansSearchPostRequest struct {
	value *ApiV2ProjectsIdTestPlansSearchPostRequest
	isSet bool
}

func (v NullableApiV2ProjectsIdTestPlansSearchPostRequest) Get() *ApiV2ProjectsIdTestPlansSearchPostRequest {
	return v.value
}

func (v *NullableApiV2ProjectsIdTestPlansSearchPostRequest) Set(val *ApiV2ProjectsIdTestPlansSearchPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableApiV2ProjectsIdTestPlansSearchPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableApiV2ProjectsIdTestPlansSearchPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiV2ProjectsIdTestPlansSearchPostRequest(val *ApiV2ProjectsIdTestPlansSearchPostRequest) *NullableApiV2ProjectsIdTestPlansSearchPostRequest {
	return &NullableApiV2ProjectsIdTestPlansSearchPostRequest{value: val, isSet: true}
}

func (v NullableApiV2ProjectsIdTestPlansSearchPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiV2ProjectsIdTestPlansSearchPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


