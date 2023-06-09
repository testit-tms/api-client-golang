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

// checks if the TestResultsFilterModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TestResultsFilterModel{}

// TestResultsFilterModel struct for TestResultsFilterModel
type TestResultsFilterModel struct {
	// Specifies a test result test run IDs to search for
	TestRunIds []string `json:"testRunIds,omitempty"`
	// Specifies a test result configuration IDs to search for
	ConfigurationIds []string `json:"configurationIds,omitempty"`
	// Specifies a test result outcomes to search for
	Outcomes []TestResultOutcome `json:"outcomes,omitempty"`
	// Specifies a test result failure categories to search for
	FailureCategories []FailureCategoryModel `json:"failureCategories,omitempty"`
	// Specifies a test result namespace to search for
	Namespace NullableString `json:"namespace,omitempty"`
	// Specifies a test result class name to search for
	ClassName NullableString `json:"className,omitempty"`
}

// NewTestResultsFilterModel instantiates a new TestResultsFilterModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTestResultsFilterModel() *TestResultsFilterModel {
	this := TestResultsFilterModel{}
	return &this
}

// NewTestResultsFilterModelWithDefaults instantiates a new TestResultsFilterModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTestResultsFilterModelWithDefaults() *TestResultsFilterModel {
	this := TestResultsFilterModel{}
	return &this
}

// GetTestRunIds returns the TestRunIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestResultsFilterModel) GetTestRunIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.TestRunIds
}

// GetTestRunIdsOk returns a tuple with the TestRunIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestResultsFilterModel) GetTestRunIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.TestRunIds) {
		return nil, false
	}
	return o.TestRunIds, true
}

// HasTestRunIds returns a boolean if a field has been set.
func (o *TestResultsFilterModel) HasTestRunIds() bool {
	if o != nil && IsNil(o.TestRunIds) {
		return true
	}

	return false
}

// SetTestRunIds gets a reference to the given []string and assigns it to the TestRunIds field.
func (o *TestResultsFilterModel) SetTestRunIds(v []string) {
	o.TestRunIds = v
}

// GetConfigurationIds returns the ConfigurationIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestResultsFilterModel) GetConfigurationIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ConfigurationIds
}

// GetConfigurationIdsOk returns a tuple with the ConfigurationIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestResultsFilterModel) GetConfigurationIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.ConfigurationIds) {
		return nil, false
	}
	return o.ConfigurationIds, true
}

// HasConfigurationIds returns a boolean if a field has been set.
func (o *TestResultsFilterModel) HasConfigurationIds() bool {
	if o != nil && IsNil(o.ConfigurationIds) {
		return true
	}

	return false
}

// SetConfigurationIds gets a reference to the given []string and assigns it to the ConfigurationIds field.
func (o *TestResultsFilterModel) SetConfigurationIds(v []string) {
	o.ConfigurationIds = v
}

// GetOutcomes returns the Outcomes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestResultsFilterModel) GetOutcomes() []TestResultOutcome {
	if o == nil {
		var ret []TestResultOutcome
		return ret
	}
	return o.Outcomes
}

// GetOutcomesOk returns a tuple with the Outcomes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestResultsFilterModel) GetOutcomesOk() ([]TestResultOutcome, bool) {
	if o == nil || IsNil(o.Outcomes) {
		return nil, false
	}
	return o.Outcomes, true
}

// HasOutcomes returns a boolean if a field has been set.
func (o *TestResultsFilterModel) HasOutcomes() bool {
	if o != nil && IsNil(o.Outcomes) {
		return true
	}

	return false
}

// SetOutcomes gets a reference to the given []TestResultOutcome and assigns it to the Outcomes field.
func (o *TestResultsFilterModel) SetOutcomes(v []TestResultOutcome) {
	o.Outcomes = v
}

// GetFailureCategories returns the FailureCategories field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestResultsFilterModel) GetFailureCategories() []FailureCategoryModel {
	if o == nil {
		var ret []FailureCategoryModel
		return ret
	}
	return o.FailureCategories
}

// GetFailureCategoriesOk returns a tuple with the FailureCategories field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestResultsFilterModel) GetFailureCategoriesOk() ([]FailureCategoryModel, bool) {
	if o == nil || IsNil(o.FailureCategories) {
		return nil, false
	}
	return o.FailureCategories, true
}

// HasFailureCategories returns a boolean if a field has been set.
func (o *TestResultsFilterModel) HasFailureCategories() bool {
	if o != nil && IsNil(o.FailureCategories) {
		return true
	}

	return false
}

// SetFailureCategories gets a reference to the given []FailureCategoryModel and assigns it to the FailureCategories field.
func (o *TestResultsFilterModel) SetFailureCategories(v []FailureCategoryModel) {
	o.FailureCategories = v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestResultsFilterModel) GetNamespace() string {
	if o == nil || IsNil(o.Namespace.Get()) {
		var ret string
		return ret
	}
	return *o.Namespace.Get()
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestResultsFilterModel) GetNamespaceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Namespace.Get(), o.Namespace.IsSet()
}

// HasNamespace returns a boolean if a field has been set.
func (o *TestResultsFilterModel) HasNamespace() bool {
	if o != nil && o.Namespace.IsSet() {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given NullableString and assigns it to the Namespace field.
func (o *TestResultsFilterModel) SetNamespace(v string) {
	o.Namespace.Set(&v)
}
// SetNamespaceNil sets the value for Namespace to be an explicit nil
func (o *TestResultsFilterModel) SetNamespaceNil() {
	o.Namespace.Set(nil)
}

// UnsetNamespace ensures that no value is present for Namespace, not even an explicit nil
func (o *TestResultsFilterModel) UnsetNamespace() {
	o.Namespace.Unset()
}

// GetClassName returns the ClassName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestResultsFilterModel) GetClassName() string {
	if o == nil || IsNil(o.ClassName.Get()) {
		var ret string
		return ret
	}
	return *o.ClassName.Get()
}

// GetClassNameOk returns a tuple with the ClassName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestResultsFilterModel) GetClassNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClassName.Get(), o.ClassName.IsSet()
}

// HasClassName returns a boolean if a field has been set.
func (o *TestResultsFilterModel) HasClassName() bool {
	if o != nil && o.ClassName.IsSet() {
		return true
	}

	return false
}

// SetClassName gets a reference to the given NullableString and assigns it to the ClassName field.
func (o *TestResultsFilterModel) SetClassName(v string) {
	o.ClassName.Set(&v)
}
// SetClassNameNil sets the value for ClassName to be an explicit nil
func (o *TestResultsFilterModel) SetClassNameNil() {
	o.ClassName.Set(nil)
}

// UnsetClassName ensures that no value is present for ClassName, not even an explicit nil
func (o *TestResultsFilterModel) UnsetClassName() {
	o.ClassName.Unset()
}

func (o TestResultsFilterModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TestResultsFilterModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.TestRunIds != nil {
		toSerialize["testRunIds"] = o.TestRunIds
	}
	if o.ConfigurationIds != nil {
		toSerialize["configurationIds"] = o.ConfigurationIds
	}
	if o.Outcomes != nil {
		toSerialize["outcomes"] = o.Outcomes
	}
	if o.FailureCategories != nil {
		toSerialize["failureCategories"] = o.FailureCategories
	}
	if o.Namespace.IsSet() {
		toSerialize["namespace"] = o.Namespace.Get()
	}
	if o.ClassName.IsSet() {
		toSerialize["className"] = o.ClassName.Get()
	}
	return toSerialize, nil
}

type NullableTestResultsFilterModel struct {
	value *TestResultsFilterModel
	isSet bool
}

func (v NullableTestResultsFilterModel) Get() *TestResultsFilterModel {
	return v.value
}

func (v *NullableTestResultsFilterModel) Set(val *TestResultsFilterModel) {
	v.value = val
	v.isSet = true
}

func (v NullableTestResultsFilterModel) IsSet() bool {
	return v.isSet
}

func (v *NullableTestResultsFilterModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestResultsFilterModel(val *TestResultsFilterModel) *NullableTestResultsFilterModel {
	return &NullableTestResultsFilterModel{value: val, isSet: true}
}

func (v NullableTestResultsFilterModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTestResultsFilterModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


