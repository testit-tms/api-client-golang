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

// checks if the PublicTestRunModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PublicTestRunModel{}

// PublicTestRunModel struct for PublicTestRunModel
type PublicTestRunModel struct {
	TestRunId *string `json:"testRunId,omitempty"`
	TestPlanId NullableString `json:"testPlanId,omitempty"`
	TestPlanGlobalId *int64 `json:"testPlanGlobalId,omitempty"`
	Name *string `json:"name,omitempty"`
	ProductName NullableString `json:"productName,omitempty"`
	Build NullableString `json:"build,omitempty"`
	Configurations []ConfigurationModel `json:"configurations,omitempty"`
	AutoTests []AutoTestModel `json:"autoTests,omitempty"`
	TestPoints []PublicTestPointModel `json:"testPoints,omitempty"`
	Status *string `json:"status,omitempty"`
}

// NewPublicTestRunModel instantiates a new PublicTestRunModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublicTestRunModel() *PublicTestRunModel {
	this := PublicTestRunModel{}
	return &this
}

// NewPublicTestRunModelWithDefaults instantiates a new PublicTestRunModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublicTestRunModelWithDefaults() *PublicTestRunModel {
	this := PublicTestRunModel{}
	return &this
}

// GetTestRunId returns the TestRunId field value if set, zero value otherwise.
func (o *PublicTestRunModel) GetTestRunId() string {
	if o == nil || IsNil(o.TestRunId) {
		var ret string
		return ret
	}
	return *o.TestRunId
}

// GetTestRunIdOk returns a tuple with the TestRunId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicTestRunModel) GetTestRunIdOk() (*string, bool) {
	if o == nil || IsNil(o.TestRunId) {
		return nil, false
	}
	return o.TestRunId, true
}

// HasTestRunId returns a boolean if a field has been set.
func (o *PublicTestRunModel) HasTestRunId() bool {
	if o != nil && !IsNil(o.TestRunId) {
		return true
	}

	return false
}

// SetTestRunId gets a reference to the given string and assigns it to the TestRunId field.
func (o *PublicTestRunModel) SetTestRunId(v string) {
	o.TestRunId = &v
}

// GetTestPlanId returns the TestPlanId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PublicTestRunModel) GetTestPlanId() string {
	if o == nil || IsNil(o.TestPlanId.Get()) {
		var ret string
		return ret
	}
	return *o.TestPlanId.Get()
}

// GetTestPlanIdOk returns a tuple with the TestPlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PublicTestRunModel) GetTestPlanIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TestPlanId.Get(), o.TestPlanId.IsSet()
}

// HasTestPlanId returns a boolean if a field has been set.
func (o *PublicTestRunModel) HasTestPlanId() bool {
	if o != nil && o.TestPlanId.IsSet() {
		return true
	}

	return false
}

// SetTestPlanId gets a reference to the given NullableString and assigns it to the TestPlanId field.
func (o *PublicTestRunModel) SetTestPlanId(v string) {
	o.TestPlanId.Set(&v)
}
// SetTestPlanIdNil sets the value for TestPlanId to be an explicit nil
func (o *PublicTestRunModel) SetTestPlanIdNil() {
	o.TestPlanId.Set(nil)
}

// UnsetTestPlanId ensures that no value is present for TestPlanId, not even an explicit nil
func (o *PublicTestRunModel) UnsetTestPlanId() {
	o.TestPlanId.Unset()
}

// GetTestPlanGlobalId returns the TestPlanGlobalId field value if set, zero value otherwise.
func (o *PublicTestRunModel) GetTestPlanGlobalId() int64 {
	if o == nil || IsNil(o.TestPlanGlobalId) {
		var ret int64
		return ret
	}
	return *o.TestPlanGlobalId
}

// GetTestPlanGlobalIdOk returns a tuple with the TestPlanGlobalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicTestRunModel) GetTestPlanGlobalIdOk() (*int64, bool) {
	if o == nil || IsNil(o.TestPlanGlobalId) {
		return nil, false
	}
	return o.TestPlanGlobalId, true
}

// HasTestPlanGlobalId returns a boolean if a field has been set.
func (o *PublicTestRunModel) HasTestPlanGlobalId() bool {
	if o != nil && !IsNil(o.TestPlanGlobalId) {
		return true
	}

	return false
}

// SetTestPlanGlobalId gets a reference to the given int64 and assigns it to the TestPlanGlobalId field.
func (o *PublicTestRunModel) SetTestPlanGlobalId(v int64) {
	o.TestPlanGlobalId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PublicTestRunModel) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicTestRunModel) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PublicTestRunModel) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PublicTestRunModel) SetName(v string) {
	o.Name = &v
}

// GetProductName returns the ProductName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PublicTestRunModel) GetProductName() string {
	if o == nil || IsNil(o.ProductName.Get()) {
		var ret string
		return ret
	}
	return *o.ProductName.Get()
}

// GetProductNameOk returns a tuple with the ProductName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PublicTestRunModel) GetProductNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProductName.Get(), o.ProductName.IsSet()
}

// HasProductName returns a boolean if a field has been set.
func (o *PublicTestRunModel) HasProductName() bool {
	if o != nil && o.ProductName.IsSet() {
		return true
	}

	return false
}

// SetProductName gets a reference to the given NullableString and assigns it to the ProductName field.
func (o *PublicTestRunModel) SetProductName(v string) {
	o.ProductName.Set(&v)
}
// SetProductNameNil sets the value for ProductName to be an explicit nil
func (o *PublicTestRunModel) SetProductNameNil() {
	o.ProductName.Set(nil)
}

// UnsetProductName ensures that no value is present for ProductName, not even an explicit nil
func (o *PublicTestRunModel) UnsetProductName() {
	o.ProductName.Unset()
}

// GetBuild returns the Build field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PublicTestRunModel) GetBuild() string {
	if o == nil || IsNil(o.Build.Get()) {
		var ret string
		return ret
	}
	return *o.Build.Get()
}

// GetBuildOk returns a tuple with the Build field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PublicTestRunModel) GetBuildOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Build.Get(), o.Build.IsSet()
}

// HasBuild returns a boolean if a field has been set.
func (o *PublicTestRunModel) HasBuild() bool {
	if o != nil && o.Build.IsSet() {
		return true
	}

	return false
}

// SetBuild gets a reference to the given NullableString and assigns it to the Build field.
func (o *PublicTestRunModel) SetBuild(v string) {
	o.Build.Set(&v)
}
// SetBuildNil sets the value for Build to be an explicit nil
func (o *PublicTestRunModel) SetBuildNil() {
	o.Build.Set(nil)
}

// UnsetBuild ensures that no value is present for Build, not even an explicit nil
func (o *PublicTestRunModel) UnsetBuild() {
	o.Build.Unset()
}

// GetConfigurations returns the Configurations field value if set, zero value otherwise.
func (o *PublicTestRunModel) GetConfigurations() []ConfigurationModel {
	if o == nil || IsNil(o.Configurations) {
		var ret []ConfigurationModel
		return ret
	}
	return o.Configurations
}

// GetConfigurationsOk returns a tuple with the Configurations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicTestRunModel) GetConfigurationsOk() ([]ConfigurationModel, bool) {
	if o == nil || IsNil(o.Configurations) {
		return nil, false
	}
	return o.Configurations, true
}

// HasConfigurations returns a boolean if a field has been set.
func (o *PublicTestRunModel) HasConfigurations() bool {
	if o != nil && !IsNil(o.Configurations) {
		return true
	}

	return false
}

// SetConfigurations gets a reference to the given []ConfigurationModel and assigns it to the Configurations field.
func (o *PublicTestRunModel) SetConfigurations(v []ConfigurationModel) {
	o.Configurations = v
}

// GetAutoTests returns the AutoTests field value if set, zero value otherwise.
func (o *PublicTestRunModel) GetAutoTests() []AutoTestModel {
	if o == nil || IsNil(o.AutoTests) {
		var ret []AutoTestModel
		return ret
	}
	return o.AutoTests
}

// GetAutoTestsOk returns a tuple with the AutoTests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicTestRunModel) GetAutoTestsOk() ([]AutoTestModel, bool) {
	if o == nil || IsNil(o.AutoTests) {
		return nil, false
	}
	return o.AutoTests, true
}

// HasAutoTests returns a boolean if a field has been set.
func (o *PublicTestRunModel) HasAutoTests() bool {
	if o != nil && !IsNil(o.AutoTests) {
		return true
	}

	return false
}

// SetAutoTests gets a reference to the given []AutoTestModel and assigns it to the AutoTests field.
func (o *PublicTestRunModel) SetAutoTests(v []AutoTestModel) {
	o.AutoTests = v
}

// GetTestPoints returns the TestPoints field value if set, zero value otherwise.
func (o *PublicTestRunModel) GetTestPoints() []PublicTestPointModel {
	if o == nil || IsNil(o.TestPoints) {
		var ret []PublicTestPointModel
		return ret
	}
	return o.TestPoints
}

// GetTestPointsOk returns a tuple with the TestPoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicTestRunModel) GetTestPointsOk() ([]PublicTestPointModel, bool) {
	if o == nil || IsNil(o.TestPoints) {
		return nil, false
	}
	return o.TestPoints, true
}

// HasTestPoints returns a boolean if a field has been set.
func (o *PublicTestRunModel) HasTestPoints() bool {
	if o != nil && !IsNil(o.TestPoints) {
		return true
	}

	return false
}

// SetTestPoints gets a reference to the given []PublicTestPointModel and assigns it to the TestPoints field.
func (o *PublicTestRunModel) SetTestPoints(v []PublicTestPointModel) {
	o.TestPoints = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *PublicTestRunModel) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicTestRunModel) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *PublicTestRunModel) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *PublicTestRunModel) SetStatus(v string) {
	o.Status = &v
}

func (o PublicTestRunModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PublicTestRunModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TestRunId) {
		toSerialize["testRunId"] = o.TestRunId
	}
	if o.TestPlanId.IsSet() {
		toSerialize["testPlanId"] = o.TestPlanId.Get()
	}
	if !IsNil(o.TestPlanGlobalId) {
		toSerialize["testPlanGlobalId"] = o.TestPlanGlobalId
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if o.ProductName.IsSet() {
		toSerialize["productName"] = o.ProductName.Get()
	}
	if o.Build.IsSet() {
		toSerialize["build"] = o.Build.Get()
	}
	if !IsNil(o.Configurations) {
		toSerialize["configurations"] = o.Configurations
	}
	if !IsNil(o.AutoTests) {
		toSerialize["autoTests"] = o.AutoTests
	}
	if !IsNil(o.TestPoints) {
		toSerialize["testPoints"] = o.TestPoints
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullablePublicTestRunModel struct {
	value *PublicTestRunModel
	isSet bool
}

func (v NullablePublicTestRunModel) Get() *PublicTestRunModel {
	return v.value
}

func (v *NullablePublicTestRunModel) Set(val *PublicTestRunModel) {
	v.value = val
	v.isSet = true
}

func (v NullablePublicTestRunModel) IsSet() bool {
	return v.isSet
}

func (v *NullablePublicTestRunModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePublicTestRunModel(val *PublicTestRunModel) *NullablePublicTestRunModel {
	return &NullablePublicTestRunModel{value: val, isSet: true}
}

func (v NullablePublicTestRunModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePublicTestRunModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


