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

// checks if the ApiV2ProjectsProjectIdTestPlansSearchPostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiV2ProjectsProjectIdTestPlansSearchPostRequest{}

// ApiV2ProjectsProjectIdTestPlansSearchPostRequest struct for ApiV2ProjectsProjectIdTestPlansSearchPostRequest
type ApiV2ProjectsProjectIdTestPlansSearchPostRequest struct {
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

// NewApiV2ProjectsProjectIdTestPlansSearchPostRequest instantiates a new ApiV2ProjectsProjectIdTestPlansSearchPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiV2ProjectsProjectIdTestPlansSearchPostRequest() *ApiV2ProjectsProjectIdTestPlansSearchPostRequest {
	this := ApiV2ProjectsProjectIdTestPlansSearchPostRequest{}
	return &this
}

// NewApiV2ProjectsProjectIdTestPlansSearchPostRequestWithDefaults instantiates a new ApiV2ProjectsProjectIdTestPlansSearchPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiV2ProjectsProjectIdTestPlansSearchPostRequestWithDefaults() *ApiV2ProjectsProjectIdTestPlansSearchPostRequest {
	this := ApiV2ProjectsProjectIdTestPlansSearchPostRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) UnsetName() {
	o.Name.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) UnsetDescription() {
	o.Description.Unset()
}

// GetBuild returns the Build field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) GetBuild() string {
	if o == nil || IsNil(o.Build.Get()) {
		var ret string
		return ret
	}
	return *o.Build.Get()
}

// GetBuildOk returns a tuple with the Build field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) GetBuildOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Build.Get(), o.Build.IsSet()
}

// HasBuild returns a boolean if a field has been set.
func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) HasBuild() bool {
	if o != nil && o.Build.IsSet() {
		return true
	}

	return false
}

// SetBuild gets a reference to the given NullableString and assigns it to the Build field.
func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) SetBuild(v string) {
	o.Build.Set(&v)
}
// SetBuildNil sets the value for Build to be an explicit nil
func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) SetBuildNil() {
	o.Build.Set(nil)
}

// UnsetBuild ensures that no value is present for Build, not even an explicit nil
func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) UnsetBuild() {
	o.Build.Unset()
}

// GetProductName returns the ProductName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) GetProductName() string {
	if o == nil || IsNil(o.ProductName.Get()) {
		var ret string
		return ret
	}
	return *o.ProductName.Get()
}

// GetProductNameOk returns a tuple with the ProductName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) GetProductNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProductName.Get(), o.ProductName.IsSet()
}

// HasProductName returns a boolean if a field has been set.
func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) HasProductName() bool {
	if o != nil && o.ProductName.IsSet() {
		return true
	}

	return false
}

// SetProductName gets a reference to the given NullableString and assigns it to the ProductName field.
func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) SetProductName(v string) {
	o.ProductName.Set(&v)
}
// SetProductNameNil sets the value for ProductName to be an explicit nil
func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) SetProductNameNil() {
	o.ProductName.Set(nil)
}

// UnsetProductName ensures that no value is present for ProductName, not even an explicit nil
func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) UnsetProductName() {
	o.ProductName.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) GetStatus() []TestPlanStatusModel {
	if o == nil {
		var ret []TestPlanStatusModel
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) GetStatusOk() ([]TestPlanStatusModel, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) HasStatus() bool {
	if o != nil && IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given []TestPlanStatusModel and assigns it to the Status field.
func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) SetStatus(v []TestPlanStatusModel) {
	o.Status = v
}

// GetGlobalIds returns the GlobalIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) GetGlobalIds() []int64 {
	if o == nil {
		var ret []int64
		return ret
	}
	return o.GlobalIds
}

// GetGlobalIdsOk returns a tuple with the GlobalIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) GetGlobalIdsOk() ([]int64, bool) {
	if o == nil || IsNil(o.GlobalIds) {
		return nil, false
	}
	return o.GlobalIds, true
}

// HasGlobalIds returns a boolean if a field has been set.
func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) HasGlobalIds() bool {
	if o != nil && IsNil(o.GlobalIds) {
		return true
	}

	return false
}

// SetGlobalIds gets a reference to the given []int64 and assigns it to the GlobalIds field.
func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) SetGlobalIds(v []int64) {
	o.GlobalIds = v
}

// GetIsLocked returns the IsLocked field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) GetIsLocked() bool {
	if o == nil || IsNil(o.IsLocked.Get()) {
		var ret bool
		return ret
	}
	return *o.IsLocked.Get()
}

// GetIsLockedOk returns a tuple with the IsLocked field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) GetIsLockedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsLocked.Get(), o.IsLocked.IsSet()
}

// HasIsLocked returns a boolean if a field has been set.
func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) HasIsLocked() bool {
	if o != nil && o.IsLocked.IsSet() {
		return true
	}

	return false
}

// SetIsLocked gets a reference to the given NullableBool and assigns it to the IsLocked field.
func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) SetIsLocked(v bool) {
	o.IsLocked.Set(&v)
}
// SetIsLockedNil sets the value for IsLocked to be an explicit nil
func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) SetIsLockedNil() {
	o.IsLocked.Set(nil)
}

// UnsetIsLocked ensures that no value is present for IsLocked, not even an explicit nil
func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) UnsetIsLocked() {
	o.IsLocked.Unset()
}

// GetLockedDate returns the LockedDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) GetLockedDate() DateTimeRangeSelectorModel {
	if o == nil || IsNil(o.LockedDate.Get()) {
		var ret DateTimeRangeSelectorModel
		return ret
	}
	return *o.LockedDate.Get()
}

// GetLockedDateOk returns a tuple with the LockedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) GetLockedDateOk() (*DateTimeRangeSelectorModel, bool) {
	if o == nil {
		return nil, false
	}
	return o.LockedDate.Get(), o.LockedDate.IsSet()
}

// HasLockedDate returns a boolean if a field has been set.
func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) HasLockedDate() bool {
	if o != nil && o.LockedDate.IsSet() {
		return true
	}

	return false
}

// SetLockedDate gets a reference to the given NullableDateTimeRangeSelectorModel and assigns it to the LockedDate field.
func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) SetLockedDate(v DateTimeRangeSelectorModel) {
	o.LockedDate.Set(&v)
}
// SetLockedDateNil sets the value for LockedDate to be an explicit nil
func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) SetLockedDateNil() {
	o.LockedDate.Set(nil)
}

// UnsetLockedDate ensures that no value is present for LockedDate, not even an explicit nil
func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) UnsetLockedDate() {
	o.LockedDate.Unset()
}

// GetAutomaticDurationTimer returns the AutomaticDurationTimer field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) GetAutomaticDurationTimer() []bool {
	if o == nil {
		var ret []bool
		return ret
	}
	return o.AutomaticDurationTimer
}

// GetAutomaticDurationTimerOk returns a tuple with the AutomaticDurationTimer field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) GetAutomaticDurationTimerOk() ([]bool, bool) {
	if o == nil || IsNil(o.AutomaticDurationTimer) {
		return nil, false
	}
	return o.AutomaticDurationTimer, true
}

// HasAutomaticDurationTimer returns a boolean if a field has been set.
func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) HasAutomaticDurationTimer() bool {
	if o != nil && IsNil(o.AutomaticDurationTimer) {
		return true
	}

	return false
}

// SetAutomaticDurationTimer gets a reference to the given []bool and assigns it to the AutomaticDurationTimer field.
func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) SetAutomaticDurationTimer(v []bool) {
	o.AutomaticDurationTimer = v
}

// GetCreatedByIds returns the CreatedByIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) GetCreatedByIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.CreatedByIds
}

// GetCreatedByIdsOk returns a tuple with the CreatedByIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) GetCreatedByIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.CreatedByIds) {
		return nil, false
	}
	return o.CreatedByIds, true
}

// HasCreatedByIds returns a boolean if a field has been set.
func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) HasCreatedByIds() bool {
	if o != nil && IsNil(o.CreatedByIds) {
		return true
	}

	return false
}

// SetCreatedByIds gets a reference to the given []string and assigns it to the CreatedByIds field.
func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) SetCreatedByIds(v []string) {
	o.CreatedByIds = v
}

// GetCreatedDate returns the CreatedDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) GetCreatedDate() DateTimeRangeSelectorModel {
	if o == nil || IsNil(o.CreatedDate.Get()) {
		var ret DateTimeRangeSelectorModel
		return ret
	}
	return *o.CreatedDate.Get()
}

// GetCreatedDateOk returns a tuple with the CreatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) GetCreatedDateOk() (*DateTimeRangeSelectorModel, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedDate.Get(), o.CreatedDate.IsSet()
}

// HasCreatedDate returns a boolean if a field has been set.
func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) HasCreatedDate() bool {
	if o != nil && o.CreatedDate.IsSet() {
		return true
	}

	return false
}

// SetCreatedDate gets a reference to the given NullableDateTimeRangeSelectorModel and assigns it to the CreatedDate field.
func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) SetCreatedDate(v DateTimeRangeSelectorModel) {
	o.CreatedDate.Set(&v)
}
// SetCreatedDateNil sets the value for CreatedDate to be an explicit nil
func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) SetCreatedDateNil() {
	o.CreatedDate.Set(nil)
}

// UnsetCreatedDate ensures that no value is present for CreatedDate, not even an explicit nil
func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) UnsetCreatedDate() {
	o.CreatedDate.Unset()
}

// GetStartDate returns the StartDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) GetStartDate() DateTimeRangeSelectorModel {
	if o == nil || IsNil(o.StartDate.Get()) {
		var ret DateTimeRangeSelectorModel
		return ret
	}
	return *o.StartDate.Get()
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) GetStartDateOk() (*DateTimeRangeSelectorModel, bool) {
	if o == nil {
		return nil, false
	}
	return o.StartDate.Get(), o.StartDate.IsSet()
}

// HasStartDate returns a boolean if a field has been set.
func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) HasStartDate() bool {
	if o != nil && o.StartDate.IsSet() {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given NullableDateTimeRangeSelectorModel and assigns it to the StartDate field.
func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) SetStartDate(v DateTimeRangeSelectorModel) {
	o.StartDate.Set(&v)
}
// SetStartDateNil sets the value for StartDate to be an explicit nil
func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) SetStartDateNil() {
	o.StartDate.Set(nil)
}

// UnsetStartDate ensures that no value is present for StartDate, not even an explicit nil
func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) UnsetStartDate() {
	o.StartDate.Unset()
}

// GetEndDate returns the EndDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) GetEndDate() DateTimeRangeSelectorModel {
	if o == nil || IsNil(o.EndDate.Get()) {
		var ret DateTimeRangeSelectorModel
		return ret
	}
	return *o.EndDate.Get()
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) GetEndDateOk() (*DateTimeRangeSelectorModel, bool) {
	if o == nil {
		return nil, false
	}
	return o.EndDate.Get(), o.EndDate.IsSet()
}

// HasEndDate returns a boolean if a field has been set.
func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) HasEndDate() bool {
	if o != nil && o.EndDate.IsSet() {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given NullableDateTimeRangeSelectorModel and assigns it to the EndDate field.
func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) SetEndDate(v DateTimeRangeSelectorModel) {
	o.EndDate.Set(&v)
}
// SetEndDateNil sets the value for EndDate to be an explicit nil
func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) SetEndDateNil() {
	o.EndDate.Set(nil)
}

// UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) UnsetEndDate() {
	o.EndDate.Unset()
}

// GetTagNames returns the TagNames field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) GetTagNames() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.TagNames
}

// GetTagNamesOk returns a tuple with the TagNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) GetTagNamesOk() ([]string, bool) {
	if o == nil || IsNil(o.TagNames) {
		return nil, false
	}
	return o.TagNames, true
}

// HasTagNames returns a boolean if a field has been set.
func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) HasTagNames() bool {
	if o != nil && IsNil(o.TagNames) {
		return true
	}

	return false
}

// SetTagNames gets a reference to the given []string and assigns it to the TagNames field.
func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) SetTagNames(v []string) {
	o.TagNames = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) GetAttributes() map[string][]string {
	if o == nil {
		var ret map[string][]string
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) GetAttributesOk() (*map[string][]string, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return &o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) HasAttributes() bool {
	if o != nil && IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string][]string and assigns it to the Attributes field.
func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) SetAttributes(v map[string][]string) {
	o.Attributes = v
}

// GetIsDeleted returns the IsDeleted field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) GetIsDeleted() bool {
	if o == nil || IsNil(o.IsDeleted.Get()) {
		var ret bool
		return ret
	}
	return *o.IsDeleted.Get()
}

// GetIsDeletedOk returns a tuple with the IsDeleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) GetIsDeletedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsDeleted.Get(), o.IsDeleted.IsSet()
}

// HasIsDeleted returns a boolean if a field has been set.
func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) HasIsDeleted() bool {
	if o != nil && o.IsDeleted.IsSet() {
		return true
	}

	return false
}

// SetIsDeleted gets a reference to the given NullableBool and assigns it to the IsDeleted field.
func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) SetIsDeleted(v bool) {
	o.IsDeleted.Set(&v)
}
// SetIsDeletedNil sets the value for IsDeleted to be an explicit nil
func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) SetIsDeletedNil() {
	o.IsDeleted.Set(nil)
}

// UnsetIsDeleted ensures that no value is present for IsDeleted, not even an explicit nil
func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) UnsetIsDeleted() {
	o.IsDeleted.Unset()
}

func (o ApiV2ProjectsProjectIdTestPlansSearchPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiV2ProjectsProjectIdTestPlansSearchPostRequest) ToMap() (map[string]interface{}, error) {
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

type NullableApiV2ProjectsProjectIdTestPlansSearchPostRequest struct {
	value *ApiV2ProjectsProjectIdTestPlansSearchPostRequest
	isSet bool
}

func (v NullableApiV2ProjectsProjectIdTestPlansSearchPostRequest) Get() *ApiV2ProjectsProjectIdTestPlansSearchPostRequest {
	return v.value
}

func (v *NullableApiV2ProjectsProjectIdTestPlansSearchPostRequest) Set(val *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableApiV2ProjectsProjectIdTestPlansSearchPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableApiV2ProjectsProjectIdTestPlansSearchPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiV2ProjectsProjectIdTestPlansSearchPostRequest(val *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) *NullableApiV2ProjectsProjectIdTestPlansSearchPostRequest {
	return &NullableApiV2ProjectsProjectIdTestPlansSearchPostRequest{value: val, isSet: true}
}

func (v NullableApiV2ProjectsProjectIdTestPlansSearchPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiV2ProjectsProjectIdTestPlansSearchPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


