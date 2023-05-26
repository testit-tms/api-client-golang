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

// checks if the AutoTestModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AutoTestModel{}

// AutoTestModel struct for AutoTestModel
type AutoTestModel struct {
	// Global ID of the autotest
	GlobalId *int64 `json:"globalId,omitempty"`
	// Indicates if the autotest is deleted
	IsDeleted *bool `json:"isDeleted,omitempty"`
	// Indicates if the autotest has unapproved changes from linked work items
	MustBeApproved *bool `json:"mustBeApproved,omitempty"`
	// Unique ID of the autotest
	Id *string `json:"id,omitempty"`
	// Creation date of the autotest
	CreatedDate *time.Time `json:"createdDate,omitempty"`
	// Last modification date of the project
	ModifiedDate NullableTime `json:"modifiedDate,omitempty"`
	// Unique ID of the project creator
	CreatedById *string `json:"createdById,omitempty"`
	// Unique ID of the project last editor
	ModifiedById NullableString `json:"modifiedById,omitempty"`
	// Unique ID of the autotest last test run
	LastTestRunId NullableString `json:"lastTestRunId,omitempty"`
	// Name of the autotest last test run
	LastTestRunName NullableString `json:"lastTestRunName,omitempty"`
	// Unique ID of the autotest last test result
	LastTestResultId NullableString `json:"lastTestResultId,omitempty"`
	// Outcome of the autotest last test result
	LastTestResultOutcome NullableString `json:"lastTestResultOutcome,omitempty"`
	// Stability percentage of the autotest
	StabilityPercentage NullableInt32 `json:"stabilityPercentage,omitempty"`
	// External ID of the autotest
	ExternalId string `json:"externalId"`
	// Collection of the autotest links
	Links []LinkPutModel `json:"links,omitempty"`
	// Unique ID of the autotest project
	ProjectId string `json:"projectId"`
	// Name of the autotest
	Name string `json:"name"`
	// Name of the autotest namespace
	Namespace NullableString `json:"namespace,omitempty"`
	// Name of the autotest class
	Classname NullableString `json:"classname,omitempty"`
	// Collection of the autotest steps
	Steps []AutoTestStepModel `json:"steps,omitempty"`
	// Collection of the autotest setup steps
	Setup []AutoTestStepModel `json:"setup,omitempty"`
	// Collection of the autotest teardown steps
	Teardown []AutoTestStepModel `json:"teardown,omitempty"`
	// Name of the autotest in autotest's card
	Title NullableString `json:"title,omitempty"`
	// Description of the autotest in autotest's card
	Description NullableString `json:"description,omitempty"`
	// Collection of the autotest labels
	Labels []LabelShortModel `json:"labels,omitempty"`
	// Indicates if the autotest is marked as flaky
	IsFlaky NullableBool `json:"isFlaky,omitempty"`
}

// NewAutoTestModel instantiates a new AutoTestModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAutoTestModel(externalId string, projectId string, name string) *AutoTestModel {
	this := AutoTestModel{}
	this.ExternalId = externalId
	this.ProjectId = projectId
	this.Name = name
	return &this
}

// NewAutoTestModelWithDefaults instantiates a new AutoTestModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAutoTestModelWithDefaults() *AutoTestModel {
	this := AutoTestModel{}
	return &this
}

// GetGlobalId returns the GlobalId field value if set, zero value otherwise.
func (o *AutoTestModel) GetGlobalId() int64 {
	if o == nil || IsNil(o.GlobalId) {
		var ret int64
		return ret
	}
	return *o.GlobalId
}

// GetGlobalIdOk returns a tuple with the GlobalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoTestModel) GetGlobalIdOk() (*int64, bool) {
	if o == nil || IsNil(o.GlobalId) {
		return nil, false
	}
	return o.GlobalId, true
}

// HasGlobalId returns a boolean if a field has been set.
func (o *AutoTestModel) HasGlobalId() bool {
	if o != nil && !IsNil(o.GlobalId) {
		return true
	}

	return false
}

// SetGlobalId gets a reference to the given int64 and assigns it to the GlobalId field.
func (o *AutoTestModel) SetGlobalId(v int64) {
	o.GlobalId = &v
}

// GetIsDeleted returns the IsDeleted field value if set, zero value otherwise.
func (o *AutoTestModel) GetIsDeleted() bool {
	if o == nil || IsNil(o.IsDeleted) {
		var ret bool
		return ret
	}
	return *o.IsDeleted
}

// GetIsDeletedOk returns a tuple with the IsDeleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoTestModel) GetIsDeletedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDeleted) {
		return nil, false
	}
	return o.IsDeleted, true
}

// HasIsDeleted returns a boolean if a field has been set.
func (o *AutoTestModel) HasIsDeleted() bool {
	if o != nil && !IsNil(o.IsDeleted) {
		return true
	}

	return false
}

// SetIsDeleted gets a reference to the given bool and assigns it to the IsDeleted field.
func (o *AutoTestModel) SetIsDeleted(v bool) {
	o.IsDeleted = &v
}

// GetMustBeApproved returns the MustBeApproved field value if set, zero value otherwise.
func (o *AutoTestModel) GetMustBeApproved() bool {
	if o == nil || IsNil(o.MustBeApproved) {
		var ret bool
		return ret
	}
	return *o.MustBeApproved
}

// GetMustBeApprovedOk returns a tuple with the MustBeApproved field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoTestModel) GetMustBeApprovedOk() (*bool, bool) {
	if o == nil || IsNil(o.MustBeApproved) {
		return nil, false
	}
	return o.MustBeApproved, true
}

// HasMustBeApproved returns a boolean if a field has been set.
func (o *AutoTestModel) HasMustBeApproved() bool {
	if o != nil && !IsNil(o.MustBeApproved) {
		return true
	}

	return false
}

// SetMustBeApproved gets a reference to the given bool and assigns it to the MustBeApproved field.
func (o *AutoTestModel) SetMustBeApproved(v bool) {
	o.MustBeApproved = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AutoTestModel) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoTestModel) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AutoTestModel) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AutoTestModel) SetId(v string) {
	o.Id = &v
}

// GetCreatedDate returns the CreatedDate field value if set, zero value otherwise.
func (o *AutoTestModel) GetCreatedDate() time.Time {
	if o == nil || IsNil(o.CreatedDate) {
		var ret time.Time
		return ret
	}
	return *o.CreatedDate
}

// GetCreatedDateOk returns a tuple with the CreatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoTestModel) GetCreatedDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedDate) {
		return nil, false
	}
	return o.CreatedDate, true
}

// HasCreatedDate returns a boolean if a field has been set.
func (o *AutoTestModel) HasCreatedDate() bool {
	if o != nil && !IsNil(o.CreatedDate) {
		return true
	}

	return false
}

// SetCreatedDate gets a reference to the given time.Time and assigns it to the CreatedDate field.
func (o *AutoTestModel) SetCreatedDate(v time.Time) {
	o.CreatedDate = &v
}

// GetModifiedDate returns the ModifiedDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AutoTestModel) GetModifiedDate() time.Time {
	if o == nil || IsNil(o.ModifiedDate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.ModifiedDate.Get()
}

// GetModifiedDateOk returns a tuple with the ModifiedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AutoTestModel) GetModifiedDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ModifiedDate.Get(), o.ModifiedDate.IsSet()
}

// HasModifiedDate returns a boolean if a field has been set.
func (o *AutoTestModel) HasModifiedDate() bool {
	if o != nil && o.ModifiedDate.IsSet() {
		return true
	}

	return false
}

// SetModifiedDate gets a reference to the given NullableTime and assigns it to the ModifiedDate field.
func (o *AutoTestModel) SetModifiedDate(v time.Time) {
	o.ModifiedDate.Set(&v)
}
// SetModifiedDateNil sets the value for ModifiedDate to be an explicit nil
func (o *AutoTestModel) SetModifiedDateNil() {
	o.ModifiedDate.Set(nil)
}

// UnsetModifiedDate ensures that no value is present for ModifiedDate, not even an explicit nil
func (o *AutoTestModel) UnsetModifiedDate() {
	o.ModifiedDate.Unset()
}

// GetCreatedById returns the CreatedById field value if set, zero value otherwise.
func (o *AutoTestModel) GetCreatedById() string {
	if o == nil || IsNil(o.CreatedById) {
		var ret string
		return ret
	}
	return *o.CreatedById
}

// GetCreatedByIdOk returns a tuple with the CreatedById field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoTestModel) GetCreatedByIdOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedById) {
		return nil, false
	}
	return o.CreatedById, true
}

// HasCreatedById returns a boolean if a field has been set.
func (o *AutoTestModel) HasCreatedById() bool {
	if o != nil && !IsNil(o.CreatedById) {
		return true
	}

	return false
}

// SetCreatedById gets a reference to the given string and assigns it to the CreatedById field.
func (o *AutoTestModel) SetCreatedById(v string) {
	o.CreatedById = &v
}

// GetModifiedById returns the ModifiedById field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AutoTestModel) GetModifiedById() string {
	if o == nil || IsNil(o.ModifiedById.Get()) {
		var ret string
		return ret
	}
	return *o.ModifiedById.Get()
}

// GetModifiedByIdOk returns a tuple with the ModifiedById field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AutoTestModel) GetModifiedByIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ModifiedById.Get(), o.ModifiedById.IsSet()
}

// HasModifiedById returns a boolean if a field has been set.
func (o *AutoTestModel) HasModifiedById() bool {
	if o != nil && o.ModifiedById.IsSet() {
		return true
	}

	return false
}

// SetModifiedById gets a reference to the given NullableString and assigns it to the ModifiedById field.
func (o *AutoTestModel) SetModifiedById(v string) {
	o.ModifiedById.Set(&v)
}
// SetModifiedByIdNil sets the value for ModifiedById to be an explicit nil
func (o *AutoTestModel) SetModifiedByIdNil() {
	o.ModifiedById.Set(nil)
}

// UnsetModifiedById ensures that no value is present for ModifiedById, not even an explicit nil
func (o *AutoTestModel) UnsetModifiedById() {
	o.ModifiedById.Unset()
}

// GetLastTestRunId returns the LastTestRunId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AutoTestModel) GetLastTestRunId() string {
	if o == nil || IsNil(o.LastTestRunId.Get()) {
		var ret string
		return ret
	}
	return *o.LastTestRunId.Get()
}

// GetLastTestRunIdOk returns a tuple with the LastTestRunId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AutoTestModel) GetLastTestRunIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastTestRunId.Get(), o.LastTestRunId.IsSet()
}

// HasLastTestRunId returns a boolean if a field has been set.
func (o *AutoTestModel) HasLastTestRunId() bool {
	if o != nil && o.LastTestRunId.IsSet() {
		return true
	}

	return false
}

// SetLastTestRunId gets a reference to the given NullableString and assigns it to the LastTestRunId field.
func (o *AutoTestModel) SetLastTestRunId(v string) {
	o.LastTestRunId.Set(&v)
}
// SetLastTestRunIdNil sets the value for LastTestRunId to be an explicit nil
func (o *AutoTestModel) SetLastTestRunIdNil() {
	o.LastTestRunId.Set(nil)
}

// UnsetLastTestRunId ensures that no value is present for LastTestRunId, not even an explicit nil
func (o *AutoTestModel) UnsetLastTestRunId() {
	o.LastTestRunId.Unset()
}

// GetLastTestRunName returns the LastTestRunName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AutoTestModel) GetLastTestRunName() string {
	if o == nil || IsNil(o.LastTestRunName.Get()) {
		var ret string
		return ret
	}
	return *o.LastTestRunName.Get()
}

// GetLastTestRunNameOk returns a tuple with the LastTestRunName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AutoTestModel) GetLastTestRunNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastTestRunName.Get(), o.LastTestRunName.IsSet()
}

// HasLastTestRunName returns a boolean if a field has been set.
func (o *AutoTestModel) HasLastTestRunName() bool {
	if o != nil && o.LastTestRunName.IsSet() {
		return true
	}

	return false
}

// SetLastTestRunName gets a reference to the given NullableString and assigns it to the LastTestRunName field.
func (o *AutoTestModel) SetLastTestRunName(v string) {
	o.LastTestRunName.Set(&v)
}
// SetLastTestRunNameNil sets the value for LastTestRunName to be an explicit nil
func (o *AutoTestModel) SetLastTestRunNameNil() {
	o.LastTestRunName.Set(nil)
}

// UnsetLastTestRunName ensures that no value is present for LastTestRunName, not even an explicit nil
func (o *AutoTestModel) UnsetLastTestRunName() {
	o.LastTestRunName.Unset()
}

// GetLastTestResultId returns the LastTestResultId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AutoTestModel) GetLastTestResultId() string {
	if o == nil || IsNil(o.LastTestResultId.Get()) {
		var ret string
		return ret
	}
	return *o.LastTestResultId.Get()
}

// GetLastTestResultIdOk returns a tuple with the LastTestResultId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AutoTestModel) GetLastTestResultIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastTestResultId.Get(), o.LastTestResultId.IsSet()
}

// HasLastTestResultId returns a boolean if a field has been set.
func (o *AutoTestModel) HasLastTestResultId() bool {
	if o != nil && o.LastTestResultId.IsSet() {
		return true
	}

	return false
}

// SetLastTestResultId gets a reference to the given NullableString and assigns it to the LastTestResultId field.
func (o *AutoTestModel) SetLastTestResultId(v string) {
	o.LastTestResultId.Set(&v)
}
// SetLastTestResultIdNil sets the value for LastTestResultId to be an explicit nil
func (o *AutoTestModel) SetLastTestResultIdNil() {
	o.LastTestResultId.Set(nil)
}

// UnsetLastTestResultId ensures that no value is present for LastTestResultId, not even an explicit nil
func (o *AutoTestModel) UnsetLastTestResultId() {
	o.LastTestResultId.Unset()
}

// GetLastTestResultOutcome returns the LastTestResultOutcome field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AutoTestModel) GetLastTestResultOutcome() string {
	if o == nil || IsNil(o.LastTestResultOutcome.Get()) {
		var ret string
		return ret
	}
	return *o.LastTestResultOutcome.Get()
}

// GetLastTestResultOutcomeOk returns a tuple with the LastTestResultOutcome field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AutoTestModel) GetLastTestResultOutcomeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastTestResultOutcome.Get(), o.LastTestResultOutcome.IsSet()
}

// HasLastTestResultOutcome returns a boolean if a field has been set.
func (o *AutoTestModel) HasLastTestResultOutcome() bool {
	if o != nil && o.LastTestResultOutcome.IsSet() {
		return true
	}

	return false
}

// SetLastTestResultOutcome gets a reference to the given NullableString and assigns it to the LastTestResultOutcome field.
func (o *AutoTestModel) SetLastTestResultOutcome(v string) {
	o.LastTestResultOutcome.Set(&v)
}
// SetLastTestResultOutcomeNil sets the value for LastTestResultOutcome to be an explicit nil
func (o *AutoTestModel) SetLastTestResultOutcomeNil() {
	o.LastTestResultOutcome.Set(nil)
}

// UnsetLastTestResultOutcome ensures that no value is present for LastTestResultOutcome, not even an explicit nil
func (o *AutoTestModel) UnsetLastTestResultOutcome() {
	o.LastTestResultOutcome.Unset()
}

// GetStabilityPercentage returns the StabilityPercentage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AutoTestModel) GetStabilityPercentage() int32 {
	if o == nil || IsNil(o.StabilityPercentage.Get()) {
		var ret int32
		return ret
	}
	return *o.StabilityPercentage.Get()
}

// GetStabilityPercentageOk returns a tuple with the StabilityPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AutoTestModel) GetStabilityPercentageOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.StabilityPercentage.Get(), o.StabilityPercentage.IsSet()
}

// HasStabilityPercentage returns a boolean if a field has been set.
func (o *AutoTestModel) HasStabilityPercentage() bool {
	if o != nil && o.StabilityPercentage.IsSet() {
		return true
	}

	return false
}

// SetStabilityPercentage gets a reference to the given NullableInt32 and assigns it to the StabilityPercentage field.
func (o *AutoTestModel) SetStabilityPercentage(v int32) {
	o.StabilityPercentage.Set(&v)
}
// SetStabilityPercentageNil sets the value for StabilityPercentage to be an explicit nil
func (o *AutoTestModel) SetStabilityPercentageNil() {
	o.StabilityPercentage.Set(nil)
}

// UnsetStabilityPercentage ensures that no value is present for StabilityPercentage, not even an explicit nil
func (o *AutoTestModel) UnsetStabilityPercentage() {
	o.StabilityPercentage.Unset()
}

// GetExternalId returns the ExternalId field value
func (o *AutoTestModel) GetExternalId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value
// and a boolean to check if the value has been set.
func (o *AutoTestModel) GetExternalIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExternalId, true
}

// SetExternalId sets field value
func (o *AutoTestModel) SetExternalId(v string) {
	o.ExternalId = v
}

// GetLinks returns the Links field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AutoTestModel) GetLinks() []LinkPutModel {
	if o == nil {
		var ret []LinkPutModel
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AutoTestModel) GetLinksOk() ([]LinkPutModel, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *AutoTestModel) HasLinks() bool {
	if o != nil && IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []LinkPutModel and assigns it to the Links field.
func (o *AutoTestModel) SetLinks(v []LinkPutModel) {
	o.Links = v
}

// GetProjectId returns the ProjectId field value
func (o *AutoTestModel) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *AutoTestModel) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *AutoTestModel) SetProjectId(v string) {
	o.ProjectId = v
}

// GetName returns the Name field value
func (o *AutoTestModel) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AutoTestModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AutoTestModel) SetName(v string) {
	o.Name = v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AutoTestModel) GetNamespace() string {
	if o == nil || IsNil(o.Namespace.Get()) {
		var ret string
		return ret
	}
	return *o.Namespace.Get()
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AutoTestModel) GetNamespaceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Namespace.Get(), o.Namespace.IsSet()
}

// HasNamespace returns a boolean if a field has been set.
func (o *AutoTestModel) HasNamespace() bool {
	if o != nil && o.Namespace.IsSet() {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given NullableString and assigns it to the Namespace field.
func (o *AutoTestModel) SetNamespace(v string) {
	o.Namespace.Set(&v)
}
// SetNamespaceNil sets the value for Namespace to be an explicit nil
func (o *AutoTestModel) SetNamespaceNil() {
	o.Namespace.Set(nil)
}

// UnsetNamespace ensures that no value is present for Namespace, not even an explicit nil
func (o *AutoTestModel) UnsetNamespace() {
	o.Namespace.Unset()
}

// GetClassname returns the Classname field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AutoTestModel) GetClassname() string {
	if o == nil || IsNil(o.Classname.Get()) {
		var ret string
		return ret
	}
	return *o.Classname.Get()
}

// GetClassnameOk returns a tuple with the Classname field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AutoTestModel) GetClassnameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Classname.Get(), o.Classname.IsSet()
}

// HasClassname returns a boolean if a field has been set.
func (o *AutoTestModel) HasClassname() bool {
	if o != nil && o.Classname.IsSet() {
		return true
	}

	return false
}

// SetClassname gets a reference to the given NullableString and assigns it to the Classname field.
func (o *AutoTestModel) SetClassname(v string) {
	o.Classname.Set(&v)
}
// SetClassnameNil sets the value for Classname to be an explicit nil
func (o *AutoTestModel) SetClassnameNil() {
	o.Classname.Set(nil)
}

// UnsetClassname ensures that no value is present for Classname, not even an explicit nil
func (o *AutoTestModel) UnsetClassname() {
	o.Classname.Unset()
}

// GetSteps returns the Steps field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AutoTestModel) GetSteps() []AutoTestStepModel {
	if o == nil {
		var ret []AutoTestStepModel
		return ret
	}
	return o.Steps
}

// GetStepsOk returns a tuple with the Steps field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AutoTestModel) GetStepsOk() ([]AutoTestStepModel, bool) {
	if o == nil || IsNil(o.Steps) {
		return nil, false
	}
	return o.Steps, true
}

// HasSteps returns a boolean if a field has been set.
func (o *AutoTestModel) HasSteps() bool {
	if o != nil && IsNil(o.Steps) {
		return true
	}

	return false
}

// SetSteps gets a reference to the given []AutoTestStepModel and assigns it to the Steps field.
func (o *AutoTestModel) SetSteps(v []AutoTestStepModel) {
	o.Steps = v
}

// GetSetup returns the Setup field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AutoTestModel) GetSetup() []AutoTestStepModel {
	if o == nil {
		var ret []AutoTestStepModel
		return ret
	}
	return o.Setup
}

// GetSetupOk returns a tuple with the Setup field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AutoTestModel) GetSetupOk() ([]AutoTestStepModel, bool) {
	if o == nil || IsNil(o.Setup) {
		return nil, false
	}
	return o.Setup, true
}

// HasSetup returns a boolean if a field has been set.
func (o *AutoTestModel) HasSetup() bool {
	if o != nil && IsNil(o.Setup) {
		return true
	}

	return false
}

// SetSetup gets a reference to the given []AutoTestStepModel and assigns it to the Setup field.
func (o *AutoTestModel) SetSetup(v []AutoTestStepModel) {
	o.Setup = v
}

// GetTeardown returns the Teardown field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AutoTestModel) GetTeardown() []AutoTestStepModel {
	if o == nil {
		var ret []AutoTestStepModel
		return ret
	}
	return o.Teardown
}

// GetTeardownOk returns a tuple with the Teardown field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AutoTestModel) GetTeardownOk() ([]AutoTestStepModel, bool) {
	if o == nil || IsNil(o.Teardown) {
		return nil, false
	}
	return o.Teardown, true
}

// HasTeardown returns a boolean if a field has been set.
func (o *AutoTestModel) HasTeardown() bool {
	if o != nil && IsNil(o.Teardown) {
		return true
	}

	return false
}

// SetTeardown gets a reference to the given []AutoTestStepModel and assigns it to the Teardown field.
func (o *AutoTestModel) SetTeardown(v []AutoTestStepModel) {
	o.Teardown = v
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AutoTestModel) GetTitle() string {
	if o == nil || IsNil(o.Title.Get()) {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AutoTestModel) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *AutoTestModel) HasTitle() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *AutoTestModel) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *AutoTestModel) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *AutoTestModel) UnsetTitle() {
	o.Title.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AutoTestModel) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AutoTestModel) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *AutoTestModel) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *AutoTestModel) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *AutoTestModel) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *AutoTestModel) UnsetDescription() {
	o.Description.Unset()
}

// GetLabels returns the Labels field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AutoTestModel) GetLabels() []LabelShortModel {
	if o == nil {
		var ret []LabelShortModel
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AutoTestModel) GetLabelsOk() ([]LabelShortModel, bool) {
	if o == nil || IsNil(o.Labels) {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *AutoTestModel) HasLabels() bool {
	if o != nil && IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []LabelShortModel and assigns it to the Labels field.
func (o *AutoTestModel) SetLabels(v []LabelShortModel) {
	o.Labels = v
}

// GetIsFlaky returns the IsFlaky field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AutoTestModel) GetIsFlaky() bool {
	if o == nil || IsNil(o.IsFlaky.Get()) {
		var ret bool
		return ret
	}
	return *o.IsFlaky.Get()
}

// GetIsFlakyOk returns a tuple with the IsFlaky field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AutoTestModel) GetIsFlakyOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsFlaky.Get(), o.IsFlaky.IsSet()
}

// HasIsFlaky returns a boolean if a field has been set.
func (o *AutoTestModel) HasIsFlaky() bool {
	if o != nil && o.IsFlaky.IsSet() {
		return true
	}

	return false
}

// SetIsFlaky gets a reference to the given NullableBool and assigns it to the IsFlaky field.
func (o *AutoTestModel) SetIsFlaky(v bool) {
	o.IsFlaky.Set(&v)
}
// SetIsFlakyNil sets the value for IsFlaky to be an explicit nil
func (o *AutoTestModel) SetIsFlakyNil() {
	o.IsFlaky.Set(nil)
}

// UnsetIsFlaky ensures that no value is present for IsFlaky, not even an explicit nil
func (o *AutoTestModel) UnsetIsFlaky() {
	o.IsFlaky.Unset()
}

func (o AutoTestModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AutoTestModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GlobalId) {
		toSerialize["globalId"] = o.GlobalId
	}
	if !IsNil(o.IsDeleted) {
		toSerialize["isDeleted"] = o.IsDeleted
	}
	if !IsNil(o.MustBeApproved) {
		toSerialize["mustBeApproved"] = o.MustBeApproved
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.CreatedDate) {
		toSerialize["createdDate"] = o.CreatedDate
	}
	if o.ModifiedDate.IsSet() {
		toSerialize["modifiedDate"] = o.ModifiedDate.Get()
	}
	if !IsNil(o.CreatedById) {
		toSerialize["createdById"] = o.CreatedById
	}
	if o.ModifiedById.IsSet() {
		toSerialize["modifiedById"] = o.ModifiedById.Get()
	}
	if o.LastTestRunId.IsSet() {
		toSerialize["lastTestRunId"] = o.LastTestRunId.Get()
	}
	if o.LastTestRunName.IsSet() {
		toSerialize["lastTestRunName"] = o.LastTestRunName.Get()
	}
	if o.LastTestResultId.IsSet() {
		toSerialize["lastTestResultId"] = o.LastTestResultId.Get()
	}
	if o.LastTestResultOutcome.IsSet() {
		toSerialize["lastTestResultOutcome"] = o.LastTestResultOutcome.Get()
	}
	if o.StabilityPercentage.IsSet() {
		toSerialize["stabilityPercentage"] = o.StabilityPercentage.Get()
	}
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
	if o.Title.IsSet() {
		toSerialize["title"] = o.Title.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.Labels != nil {
		toSerialize["labels"] = o.Labels
	}
	if o.IsFlaky.IsSet() {
		toSerialize["isFlaky"] = o.IsFlaky.Get()
	}
	return toSerialize, nil
}

type NullableAutoTestModel struct {
	value *AutoTestModel
	isSet bool
}

func (v NullableAutoTestModel) Get() *AutoTestModel {
	return v.value
}

func (v *NullableAutoTestModel) Set(val *AutoTestModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAutoTestModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAutoTestModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAutoTestModel(val *AutoTestModel) *NullableAutoTestModel {
	return &NullableAutoTestModel{value: val, isSet: true}
}

func (v NullableAutoTestModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAutoTestModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

