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

// checks if the AutotestFilterModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AutotestFilterModel{}

// AutotestFilterModel struct for AutotestFilterModel
type AutotestFilterModel struct {
	// Specifies an autotest projects IDs to search for
	ProjectIds []string `json:"projectIds,omitempty"`
	// Specifies an autotest external IDs to search for
	ExternalIds []string `json:"externalIds,omitempty"`
	// Specifies an autotest global IDs to search for
	GlobalIds []int64 `json:"globalIds,omitempty"`
	// Specifies an autotest name to search for
	Name NullableString `json:"name,omitempty"`
	// Specifies an autotest flaky status to search for
	IsFlaky NullableBool `json:"isFlaky,omitempty"`
	// Specifies an autotest unapproved changes status to search for
	MustBeApproved NullableBool `json:"mustBeApproved,omitempty"`
	StabilityPercentage NullableAutotestFilterModelStabilityPercentage `json:"stabilityPercentage,omitempty"`
	CreatedDate NullableAutotestFilterModelCreatedDate `json:"createdDate,omitempty"`
	// Specifies an autotest creator IDs to search for
	CreatedByIds []string `json:"createdByIds,omitempty"`
	ModifiedDate NullableAutotestFilterModelModifiedDate `json:"modifiedDate,omitempty"`
	// Specifies an autotest last editor IDs to search for
	ModifiedByIds []string `json:"modifiedByIds,omitempty"`
	// Specifies an autotest deleted status to search for
	IsDeleted NullableBool `json:"isDeleted,omitempty"`
	// Specifies an autotest namespace to search for
	Namespace NullableString `json:"namespace,omitempty"`
	// Specifies an autotest namespace name presence status to search for
	IsEmptyNamespace NullableBool `json:"isEmptyNamespace,omitempty"`
	// Specifies an autotest class name to search for
	ClassName NullableString `json:"className,omitempty"`
	// Specifies an autotest class name presence status to search for
	IsEmptyClassName NullableBool `json:"isEmptyClassName,omitempty"`
	LastTestResultOutcome NullableAutotestResultOutcome `json:"lastTestResultOutcome,omitempty"`
}

// NewAutotestFilterModel instantiates a new AutotestFilterModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAutotestFilterModel() *AutotestFilterModel {
	this := AutotestFilterModel{}
	return &this
}

// NewAutotestFilterModelWithDefaults instantiates a new AutotestFilterModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAutotestFilterModelWithDefaults() *AutotestFilterModel {
	this := AutotestFilterModel{}
	return &this
}

// GetProjectIds returns the ProjectIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AutotestFilterModel) GetProjectIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ProjectIds
}

// GetProjectIdsOk returns a tuple with the ProjectIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AutotestFilterModel) GetProjectIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.ProjectIds) {
		return nil, false
	}
	return o.ProjectIds, true
}

// HasProjectIds returns a boolean if a field has been set.
func (o *AutotestFilterModel) HasProjectIds() bool {
	if o != nil && IsNil(o.ProjectIds) {
		return true
	}

	return false
}

// SetProjectIds gets a reference to the given []string and assigns it to the ProjectIds field.
func (o *AutotestFilterModel) SetProjectIds(v []string) {
	o.ProjectIds = v
}

// GetExternalIds returns the ExternalIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AutotestFilterModel) GetExternalIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ExternalIds
}

// GetExternalIdsOk returns a tuple with the ExternalIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AutotestFilterModel) GetExternalIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.ExternalIds) {
		return nil, false
	}
	return o.ExternalIds, true
}

// HasExternalIds returns a boolean if a field has been set.
func (o *AutotestFilterModel) HasExternalIds() bool {
	if o != nil && IsNil(o.ExternalIds) {
		return true
	}

	return false
}

// SetExternalIds gets a reference to the given []string and assigns it to the ExternalIds field.
func (o *AutotestFilterModel) SetExternalIds(v []string) {
	o.ExternalIds = v
}

// GetGlobalIds returns the GlobalIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AutotestFilterModel) GetGlobalIds() []int64 {
	if o == nil {
		var ret []int64
		return ret
	}
	return o.GlobalIds
}

// GetGlobalIdsOk returns a tuple with the GlobalIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AutotestFilterModel) GetGlobalIdsOk() ([]int64, bool) {
	if o == nil || IsNil(o.GlobalIds) {
		return nil, false
	}
	return o.GlobalIds, true
}

// HasGlobalIds returns a boolean if a field has been set.
func (o *AutotestFilterModel) HasGlobalIds() bool {
	if o != nil && IsNil(o.GlobalIds) {
		return true
	}

	return false
}

// SetGlobalIds gets a reference to the given []int64 and assigns it to the GlobalIds field.
func (o *AutotestFilterModel) SetGlobalIds(v []int64) {
	o.GlobalIds = v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AutotestFilterModel) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AutotestFilterModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *AutotestFilterModel) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *AutotestFilterModel) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *AutotestFilterModel) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *AutotestFilterModel) UnsetName() {
	o.Name.Unset()
}

// GetIsFlaky returns the IsFlaky field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AutotestFilterModel) GetIsFlaky() bool {
	if o == nil || IsNil(o.IsFlaky.Get()) {
		var ret bool
		return ret
	}
	return *o.IsFlaky.Get()
}

// GetIsFlakyOk returns a tuple with the IsFlaky field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AutotestFilterModel) GetIsFlakyOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsFlaky.Get(), o.IsFlaky.IsSet()
}

// HasIsFlaky returns a boolean if a field has been set.
func (o *AutotestFilterModel) HasIsFlaky() bool {
	if o != nil && o.IsFlaky.IsSet() {
		return true
	}

	return false
}

// SetIsFlaky gets a reference to the given NullableBool and assigns it to the IsFlaky field.
func (o *AutotestFilterModel) SetIsFlaky(v bool) {
	o.IsFlaky.Set(&v)
}
// SetIsFlakyNil sets the value for IsFlaky to be an explicit nil
func (o *AutotestFilterModel) SetIsFlakyNil() {
	o.IsFlaky.Set(nil)
}

// UnsetIsFlaky ensures that no value is present for IsFlaky, not even an explicit nil
func (o *AutotestFilterModel) UnsetIsFlaky() {
	o.IsFlaky.Unset()
}

// GetMustBeApproved returns the MustBeApproved field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AutotestFilterModel) GetMustBeApproved() bool {
	if o == nil || IsNil(o.MustBeApproved.Get()) {
		var ret bool
		return ret
	}
	return *o.MustBeApproved.Get()
}

// GetMustBeApprovedOk returns a tuple with the MustBeApproved field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AutotestFilterModel) GetMustBeApprovedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.MustBeApproved.Get(), o.MustBeApproved.IsSet()
}

// HasMustBeApproved returns a boolean if a field has been set.
func (o *AutotestFilterModel) HasMustBeApproved() bool {
	if o != nil && o.MustBeApproved.IsSet() {
		return true
	}

	return false
}

// SetMustBeApproved gets a reference to the given NullableBool and assigns it to the MustBeApproved field.
func (o *AutotestFilterModel) SetMustBeApproved(v bool) {
	o.MustBeApproved.Set(&v)
}
// SetMustBeApprovedNil sets the value for MustBeApproved to be an explicit nil
func (o *AutotestFilterModel) SetMustBeApprovedNil() {
	o.MustBeApproved.Set(nil)
}

// UnsetMustBeApproved ensures that no value is present for MustBeApproved, not even an explicit nil
func (o *AutotestFilterModel) UnsetMustBeApproved() {
	o.MustBeApproved.Unset()
}

// GetStabilityPercentage returns the StabilityPercentage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AutotestFilterModel) GetStabilityPercentage() AutotestFilterModelStabilityPercentage {
	if o == nil || IsNil(o.StabilityPercentage.Get()) {
		var ret AutotestFilterModelStabilityPercentage
		return ret
	}
	return *o.StabilityPercentage.Get()
}

// GetStabilityPercentageOk returns a tuple with the StabilityPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AutotestFilterModel) GetStabilityPercentageOk() (*AutotestFilterModelStabilityPercentage, bool) {
	if o == nil {
		return nil, false
	}
	return o.StabilityPercentage.Get(), o.StabilityPercentage.IsSet()
}

// HasStabilityPercentage returns a boolean if a field has been set.
func (o *AutotestFilterModel) HasStabilityPercentage() bool {
	if o != nil && o.StabilityPercentage.IsSet() {
		return true
	}

	return false
}

// SetStabilityPercentage gets a reference to the given NullableAutotestFilterModelStabilityPercentage and assigns it to the StabilityPercentage field.
func (o *AutotestFilterModel) SetStabilityPercentage(v AutotestFilterModelStabilityPercentage) {
	o.StabilityPercentage.Set(&v)
}
// SetStabilityPercentageNil sets the value for StabilityPercentage to be an explicit nil
func (o *AutotestFilterModel) SetStabilityPercentageNil() {
	o.StabilityPercentage.Set(nil)
}

// UnsetStabilityPercentage ensures that no value is present for StabilityPercentage, not even an explicit nil
func (o *AutotestFilterModel) UnsetStabilityPercentage() {
	o.StabilityPercentage.Unset()
}

// GetCreatedDate returns the CreatedDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AutotestFilterModel) GetCreatedDate() AutotestFilterModelCreatedDate {
	if o == nil || IsNil(o.CreatedDate.Get()) {
		var ret AutotestFilterModelCreatedDate
		return ret
	}
	return *o.CreatedDate.Get()
}

// GetCreatedDateOk returns a tuple with the CreatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AutotestFilterModel) GetCreatedDateOk() (*AutotestFilterModelCreatedDate, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedDate.Get(), o.CreatedDate.IsSet()
}

// HasCreatedDate returns a boolean if a field has been set.
func (o *AutotestFilterModel) HasCreatedDate() bool {
	if o != nil && o.CreatedDate.IsSet() {
		return true
	}

	return false
}

// SetCreatedDate gets a reference to the given NullableAutotestFilterModelCreatedDate and assigns it to the CreatedDate field.
func (o *AutotestFilterModel) SetCreatedDate(v AutotestFilterModelCreatedDate) {
	o.CreatedDate.Set(&v)
}
// SetCreatedDateNil sets the value for CreatedDate to be an explicit nil
func (o *AutotestFilterModel) SetCreatedDateNil() {
	o.CreatedDate.Set(nil)
}

// UnsetCreatedDate ensures that no value is present for CreatedDate, not even an explicit nil
func (o *AutotestFilterModel) UnsetCreatedDate() {
	o.CreatedDate.Unset()
}

// GetCreatedByIds returns the CreatedByIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AutotestFilterModel) GetCreatedByIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.CreatedByIds
}

// GetCreatedByIdsOk returns a tuple with the CreatedByIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AutotestFilterModel) GetCreatedByIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.CreatedByIds) {
		return nil, false
	}
	return o.CreatedByIds, true
}

// HasCreatedByIds returns a boolean if a field has been set.
func (o *AutotestFilterModel) HasCreatedByIds() bool {
	if o != nil && IsNil(o.CreatedByIds) {
		return true
	}

	return false
}

// SetCreatedByIds gets a reference to the given []string and assigns it to the CreatedByIds field.
func (o *AutotestFilterModel) SetCreatedByIds(v []string) {
	o.CreatedByIds = v
}

// GetModifiedDate returns the ModifiedDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AutotestFilterModel) GetModifiedDate() AutotestFilterModelModifiedDate {
	if o == nil || IsNil(o.ModifiedDate.Get()) {
		var ret AutotestFilterModelModifiedDate
		return ret
	}
	return *o.ModifiedDate.Get()
}

// GetModifiedDateOk returns a tuple with the ModifiedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AutotestFilterModel) GetModifiedDateOk() (*AutotestFilterModelModifiedDate, bool) {
	if o == nil {
		return nil, false
	}
	return o.ModifiedDate.Get(), o.ModifiedDate.IsSet()
}

// HasModifiedDate returns a boolean if a field has been set.
func (o *AutotestFilterModel) HasModifiedDate() bool {
	if o != nil && o.ModifiedDate.IsSet() {
		return true
	}

	return false
}

// SetModifiedDate gets a reference to the given NullableAutotestFilterModelModifiedDate and assigns it to the ModifiedDate field.
func (o *AutotestFilterModel) SetModifiedDate(v AutotestFilterModelModifiedDate) {
	o.ModifiedDate.Set(&v)
}
// SetModifiedDateNil sets the value for ModifiedDate to be an explicit nil
func (o *AutotestFilterModel) SetModifiedDateNil() {
	o.ModifiedDate.Set(nil)
}

// UnsetModifiedDate ensures that no value is present for ModifiedDate, not even an explicit nil
func (o *AutotestFilterModel) UnsetModifiedDate() {
	o.ModifiedDate.Unset()
}

// GetModifiedByIds returns the ModifiedByIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AutotestFilterModel) GetModifiedByIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ModifiedByIds
}

// GetModifiedByIdsOk returns a tuple with the ModifiedByIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AutotestFilterModel) GetModifiedByIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.ModifiedByIds) {
		return nil, false
	}
	return o.ModifiedByIds, true
}

// HasModifiedByIds returns a boolean if a field has been set.
func (o *AutotestFilterModel) HasModifiedByIds() bool {
	if o != nil && IsNil(o.ModifiedByIds) {
		return true
	}

	return false
}

// SetModifiedByIds gets a reference to the given []string and assigns it to the ModifiedByIds field.
func (o *AutotestFilterModel) SetModifiedByIds(v []string) {
	o.ModifiedByIds = v
}

// GetIsDeleted returns the IsDeleted field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AutotestFilterModel) GetIsDeleted() bool {
	if o == nil || IsNil(o.IsDeleted.Get()) {
		var ret bool
		return ret
	}
	return *o.IsDeleted.Get()
}

// GetIsDeletedOk returns a tuple with the IsDeleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AutotestFilterModel) GetIsDeletedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsDeleted.Get(), o.IsDeleted.IsSet()
}

// HasIsDeleted returns a boolean if a field has been set.
func (o *AutotestFilterModel) HasIsDeleted() bool {
	if o != nil && o.IsDeleted.IsSet() {
		return true
	}

	return false
}

// SetIsDeleted gets a reference to the given NullableBool and assigns it to the IsDeleted field.
func (o *AutotestFilterModel) SetIsDeleted(v bool) {
	o.IsDeleted.Set(&v)
}
// SetIsDeletedNil sets the value for IsDeleted to be an explicit nil
func (o *AutotestFilterModel) SetIsDeletedNil() {
	o.IsDeleted.Set(nil)
}

// UnsetIsDeleted ensures that no value is present for IsDeleted, not even an explicit nil
func (o *AutotestFilterModel) UnsetIsDeleted() {
	o.IsDeleted.Unset()
}

// GetNamespace returns the Namespace field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AutotestFilterModel) GetNamespace() string {
	if o == nil || IsNil(o.Namespace.Get()) {
		var ret string
		return ret
	}
	return *o.Namespace.Get()
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AutotestFilterModel) GetNamespaceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Namespace.Get(), o.Namespace.IsSet()
}

// HasNamespace returns a boolean if a field has been set.
func (o *AutotestFilterModel) HasNamespace() bool {
	if o != nil && o.Namespace.IsSet() {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given NullableString and assigns it to the Namespace field.
func (o *AutotestFilterModel) SetNamespace(v string) {
	o.Namespace.Set(&v)
}
// SetNamespaceNil sets the value for Namespace to be an explicit nil
func (o *AutotestFilterModel) SetNamespaceNil() {
	o.Namespace.Set(nil)
}

// UnsetNamespace ensures that no value is present for Namespace, not even an explicit nil
func (o *AutotestFilterModel) UnsetNamespace() {
	o.Namespace.Unset()
}

// GetIsEmptyNamespace returns the IsEmptyNamespace field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AutotestFilterModel) GetIsEmptyNamespace() bool {
	if o == nil || IsNil(o.IsEmptyNamespace.Get()) {
		var ret bool
		return ret
	}
	return *o.IsEmptyNamespace.Get()
}

// GetIsEmptyNamespaceOk returns a tuple with the IsEmptyNamespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AutotestFilterModel) GetIsEmptyNamespaceOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsEmptyNamespace.Get(), o.IsEmptyNamespace.IsSet()
}

// HasIsEmptyNamespace returns a boolean if a field has been set.
func (o *AutotestFilterModel) HasIsEmptyNamespace() bool {
	if o != nil && o.IsEmptyNamespace.IsSet() {
		return true
	}

	return false
}

// SetIsEmptyNamespace gets a reference to the given NullableBool and assigns it to the IsEmptyNamespace field.
func (o *AutotestFilterModel) SetIsEmptyNamespace(v bool) {
	o.IsEmptyNamespace.Set(&v)
}
// SetIsEmptyNamespaceNil sets the value for IsEmptyNamespace to be an explicit nil
func (o *AutotestFilterModel) SetIsEmptyNamespaceNil() {
	o.IsEmptyNamespace.Set(nil)
}

// UnsetIsEmptyNamespace ensures that no value is present for IsEmptyNamespace, not even an explicit nil
func (o *AutotestFilterModel) UnsetIsEmptyNamespace() {
	o.IsEmptyNamespace.Unset()
}

// GetClassName returns the ClassName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AutotestFilterModel) GetClassName() string {
	if o == nil || IsNil(o.ClassName.Get()) {
		var ret string
		return ret
	}
	return *o.ClassName.Get()
}

// GetClassNameOk returns a tuple with the ClassName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AutotestFilterModel) GetClassNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClassName.Get(), o.ClassName.IsSet()
}

// HasClassName returns a boolean if a field has been set.
func (o *AutotestFilterModel) HasClassName() bool {
	if o != nil && o.ClassName.IsSet() {
		return true
	}

	return false
}

// SetClassName gets a reference to the given NullableString and assigns it to the ClassName field.
func (o *AutotestFilterModel) SetClassName(v string) {
	o.ClassName.Set(&v)
}
// SetClassNameNil sets the value for ClassName to be an explicit nil
func (o *AutotestFilterModel) SetClassNameNil() {
	o.ClassName.Set(nil)
}

// UnsetClassName ensures that no value is present for ClassName, not even an explicit nil
func (o *AutotestFilterModel) UnsetClassName() {
	o.ClassName.Unset()
}

// GetIsEmptyClassName returns the IsEmptyClassName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AutotestFilterModel) GetIsEmptyClassName() bool {
	if o == nil || IsNil(o.IsEmptyClassName.Get()) {
		var ret bool
		return ret
	}
	return *o.IsEmptyClassName.Get()
}

// GetIsEmptyClassNameOk returns a tuple with the IsEmptyClassName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AutotestFilterModel) GetIsEmptyClassNameOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsEmptyClassName.Get(), o.IsEmptyClassName.IsSet()
}

// HasIsEmptyClassName returns a boolean if a field has been set.
func (o *AutotestFilterModel) HasIsEmptyClassName() bool {
	if o != nil && o.IsEmptyClassName.IsSet() {
		return true
	}

	return false
}

// SetIsEmptyClassName gets a reference to the given NullableBool and assigns it to the IsEmptyClassName field.
func (o *AutotestFilterModel) SetIsEmptyClassName(v bool) {
	o.IsEmptyClassName.Set(&v)
}
// SetIsEmptyClassNameNil sets the value for IsEmptyClassName to be an explicit nil
func (o *AutotestFilterModel) SetIsEmptyClassNameNil() {
	o.IsEmptyClassName.Set(nil)
}

// UnsetIsEmptyClassName ensures that no value is present for IsEmptyClassName, not even an explicit nil
func (o *AutotestFilterModel) UnsetIsEmptyClassName() {
	o.IsEmptyClassName.Unset()
}

// GetLastTestResultOutcome returns the LastTestResultOutcome field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AutotestFilterModel) GetLastTestResultOutcome() AutotestResultOutcome {
	if o == nil || IsNil(o.LastTestResultOutcome.Get()) {
		var ret AutotestResultOutcome
		return ret
	}
	return *o.LastTestResultOutcome.Get()
}

// GetLastTestResultOutcomeOk returns a tuple with the LastTestResultOutcome field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AutotestFilterModel) GetLastTestResultOutcomeOk() (*AutotestResultOutcome, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastTestResultOutcome.Get(), o.LastTestResultOutcome.IsSet()
}

// HasLastTestResultOutcome returns a boolean if a field has been set.
func (o *AutotestFilterModel) HasLastTestResultOutcome() bool {
	if o != nil && o.LastTestResultOutcome.IsSet() {
		return true
	}

	return false
}

// SetLastTestResultOutcome gets a reference to the given NullableAutotestResultOutcome and assigns it to the LastTestResultOutcome field.
func (o *AutotestFilterModel) SetLastTestResultOutcome(v AutotestResultOutcome) {
	o.LastTestResultOutcome.Set(&v)
}
// SetLastTestResultOutcomeNil sets the value for LastTestResultOutcome to be an explicit nil
func (o *AutotestFilterModel) SetLastTestResultOutcomeNil() {
	o.LastTestResultOutcome.Set(nil)
}

// UnsetLastTestResultOutcome ensures that no value is present for LastTestResultOutcome, not even an explicit nil
func (o *AutotestFilterModel) UnsetLastTestResultOutcome() {
	o.LastTestResultOutcome.Unset()
}

func (o AutotestFilterModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AutotestFilterModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ProjectIds != nil {
		toSerialize["projectIds"] = o.ProjectIds
	}
	if o.ExternalIds != nil {
		toSerialize["externalIds"] = o.ExternalIds
	}
	if o.GlobalIds != nil {
		toSerialize["globalIds"] = o.GlobalIds
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.IsFlaky.IsSet() {
		toSerialize["isFlaky"] = o.IsFlaky.Get()
	}
	if o.MustBeApproved.IsSet() {
		toSerialize["mustBeApproved"] = o.MustBeApproved.Get()
	}
	if o.StabilityPercentage.IsSet() {
		toSerialize["stabilityPercentage"] = o.StabilityPercentage.Get()
	}
	if o.CreatedDate.IsSet() {
		toSerialize["createdDate"] = o.CreatedDate.Get()
	}
	if o.CreatedByIds != nil {
		toSerialize["createdByIds"] = o.CreatedByIds
	}
	if o.ModifiedDate.IsSet() {
		toSerialize["modifiedDate"] = o.ModifiedDate.Get()
	}
	if o.ModifiedByIds != nil {
		toSerialize["modifiedByIds"] = o.ModifiedByIds
	}
	if o.IsDeleted.IsSet() {
		toSerialize["isDeleted"] = o.IsDeleted.Get()
	}
	if o.Namespace.IsSet() {
		toSerialize["namespace"] = o.Namespace.Get()
	}
	if o.IsEmptyNamespace.IsSet() {
		toSerialize["isEmptyNamespace"] = o.IsEmptyNamespace.Get()
	}
	if o.ClassName.IsSet() {
		toSerialize["className"] = o.ClassName.Get()
	}
	if o.IsEmptyClassName.IsSet() {
		toSerialize["isEmptyClassName"] = o.IsEmptyClassName.Get()
	}
	if o.LastTestResultOutcome.IsSet() {
		toSerialize["lastTestResultOutcome"] = o.LastTestResultOutcome.Get()
	}
	return toSerialize, nil
}

type NullableAutotestFilterModel struct {
	value *AutotestFilterModel
	isSet bool
}

func (v NullableAutotestFilterModel) Get() *AutotestFilterModel {
	return v.value
}

func (v *NullableAutotestFilterModel) Set(val *AutotestFilterModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAutotestFilterModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAutotestFilterModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAutotestFilterModel(val *AutotestFilterModel) *NullableAutotestFilterModel {
	return &NullableAutotestFilterModel{value: val, isSet: true}
}

func (v NullableAutotestFilterModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAutotestFilterModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


