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

// checks if the TestResultShortGetModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TestResultShortGetModel{}

// TestResultShortGetModel struct for TestResultShortGetModel
type TestResultShortGetModel struct {
	// Unique ID of test result
	Id *string `json:"id,omitempty"`
	// Name of autotest represented by the test result
	Name *string `json:"name,omitempty"`
	// Global ID of autotest represented by test result
	AutotestGlobalId *int64 `json:"autotestGlobalId,omitempty"`
	// Unique ID of test run where test result is located
	TestRunId *string `json:"testRunId,omitempty"`
	// Unique ID of configuration which test result uses
	ConfigurationId *string `json:"configurationId,omitempty"`
	// Name of configuration which test result uses
	ConfigurationName *string `json:"configurationName,omitempty"`
	Outcome TestResultOutcome `json:"outcome"`
	// Collection of result reasons which test result have
	ResultReasons []AutotestResultReasonSubGetModel `json:"resultReasons,omitempty"`
	// Comment to test result
	Comment *string `json:"comment,omitempty"`
	// Date when test result has been set
	Date *time.Time `json:"date,omitempty"`
	// Time which it took to run the test
	Duration NullableInt64 `json:"duration,omitempty"`
	// Collection of links attached to test result
	Links []LinkSubGetModel `json:"links,omitempty"`
	// Collection of files attached to test result
	Attachments []AttachmentSubGetModel `json:"attachments,omitempty"`
}

// NewTestResultShortGetModel instantiates a new TestResultShortGetModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTestResultShortGetModel(outcome TestResultOutcome) *TestResultShortGetModel {
	this := TestResultShortGetModel{}
	this.Outcome = outcome
	return &this
}

// NewTestResultShortGetModelWithDefaults instantiates a new TestResultShortGetModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTestResultShortGetModelWithDefaults() *TestResultShortGetModel {
	this := TestResultShortGetModel{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TestResultShortGetModel) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestResultShortGetModel) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TestResultShortGetModel) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *TestResultShortGetModel) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TestResultShortGetModel) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestResultShortGetModel) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TestResultShortGetModel) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TestResultShortGetModel) SetName(v string) {
	o.Name = &v
}

// GetAutotestGlobalId returns the AutotestGlobalId field value if set, zero value otherwise.
func (o *TestResultShortGetModel) GetAutotestGlobalId() int64 {
	if o == nil || IsNil(o.AutotestGlobalId) {
		var ret int64
		return ret
	}
	return *o.AutotestGlobalId
}

// GetAutotestGlobalIdOk returns a tuple with the AutotestGlobalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestResultShortGetModel) GetAutotestGlobalIdOk() (*int64, bool) {
	if o == nil || IsNil(o.AutotestGlobalId) {
		return nil, false
	}
	return o.AutotestGlobalId, true
}

// HasAutotestGlobalId returns a boolean if a field has been set.
func (o *TestResultShortGetModel) HasAutotestGlobalId() bool {
	if o != nil && !IsNil(o.AutotestGlobalId) {
		return true
	}

	return false
}

// SetAutotestGlobalId gets a reference to the given int64 and assigns it to the AutotestGlobalId field.
func (o *TestResultShortGetModel) SetAutotestGlobalId(v int64) {
	o.AutotestGlobalId = &v
}

// GetTestRunId returns the TestRunId field value if set, zero value otherwise.
func (o *TestResultShortGetModel) GetTestRunId() string {
	if o == nil || IsNil(o.TestRunId) {
		var ret string
		return ret
	}
	return *o.TestRunId
}

// GetTestRunIdOk returns a tuple with the TestRunId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestResultShortGetModel) GetTestRunIdOk() (*string, bool) {
	if o == nil || IsNil(o.TestRunId) {
		return nil, false
	}
	return o.TestRunId, true
}

// HasTestRunId returns a boolean if a field has been set.
func (o *TestResultShortGetModel) HasTestRunId() bool {
	if o != nil && !IsNil(o.TestRunId) {
		return true
	}

	return false
}

// SetTestRunId gets a reference to the given string and assigns it to the TestRunId field.
func (o *TestResultShortGetModel) SetTestRunId(v string) {
	o.TestRunId = &v
}

// GetConfigurationId returns the ConfigurationId field value if set, zero value otherwise.
func (o *TestResultShortGetModel) GetConfigurationId() string {
	if o == nil || IsNil(o.ConfigurationId) {
		var ret string
		return ret
	}
	return *o.ConfigurationId
}

// GetConfigurationIdOk returns a tuple with the ConfigurationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestResultShortGetModel) GetConfigurationIdOk() (*string, bool) {
	if o == nil || IsNil(o.ConfigurationId) {
		return nil, false
	}
	return o.ConfigurationId, true
}

// HasConfigurationId returns a boolean if a field has been set.
func (o *TestResultShortGetModel) HasConfigurationId() bool {
	if o != nil && !IsNil(o.ConfigurationId) {
		return true
	}

	return false
}

// SetConfigurationId gets a reference to the given string and assigns it to the ConfigurationId field.
func (o *TestResultShortGetModel) SetConfigurationId(v string) {
	o.ConfigurationId = &v
}

// GetConfigurationName returns the ConfigurationName field value if set, zero value otherwise.
func (o *TestResultShortGetModel) GetConfigurationName() string {
	if o == nil || IsNil(o.ConfigurationName) {
		var ret string
		return ret
	}
	return *o.ConfigurationName
}

// GetConfigurationNameOk returns a tuple with the ConfigurationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestResultShortGetModel) GetConfigurationNameOk() (*string, bool) {
	if o == nil || IsNil(o.ConfigurationName) {
		return nil, false
	}
	return o.ConfigurationName, true
}

// HasConfigurationName returns a boolean if a field has been set.
func (o *TestResultShortGetModel) HasConfigurationName() bool {
	if o != nil && !IsNil(o.ConfigurationName) {
		return true
	}

	return false
}

// SetConfigurationName gets a reference to the given string and assigns it to the ConfigurationName field.
func (o *TestResultShortGetModel) SetConfigurationName(v string) {
	o.ConfigurationName = &v
}

// GetOutcome returns the Outcome field value
func (o *TestResultShortGetModel) GetOutcome() TestResultOutcome {
	if o == nil {
		var ret TestResultOutcome
		return ret
	}

	return o.Outcome
}

// GetOutcomeOk returns a tuple with the Outcome field value
// and a boolean to check if the value has been set.
func (o *TestResultShortGetModel) GetOutcomeOk() (*TestResultOutcome, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Outcome, true
}

// SetOutcome sets field value
func (o *TestResultShortGetModel) SetOutcome(v TestResultOutcome) {
	o.Outcome = v
}

// GetResultReasons returns the ResultReasons field value if set, zero value otherwise.
func (o *TestResultShortGetModel) GetResultReasons() []AutotestResultReasonSubGetModel {
	if o == nil || IsNil(o.ResultReasons) {
		var ret []AutotestResultReasonSubGetModel
		return ret
	}
	return o.ResultReasons
}

// GetResultReasonsOk returns a tuple with the ResultReasons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestResultShortGetModel) GetResultReasonsOk() ([]AutotestResultReasonSubGetModel, bool) {
	if o == nil || IsNil(o.ResultReasons) {
		return nil, false
	}
	return o.ResultReasons, true
}

// HasResultReasons returns a boolean if a field has been set.
func (o *TestResultShortGetModel) HasResultReasons() bool {
	if o != nil && !IsNil(o.ResultReasons) {
		return true
	}

	return false
}

// SetResultReasons gets a reference to the given []AutotestResultReasonSubGetModel and assigns it to the ResultReasons field.
func (o *TestResultShortGetModel) SetResultReasons(v []AutotestResultReasonSubGetModel) {
	o.ResultReasons = v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *TestResultShortGetModel) GetComment() string {
	if o == nil || IsNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestResultShortGetModel) GetCommentOk() (*string, bool) {
	if o == nil || IsNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *TestResultShortGetModel) HasComment() bool {
	if o != nil && !IsNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *TestResultShortGetModel) SetComment(v string) {
	o.Comment = &v
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *TestResultShortGetModel) GetDate() time.Time {
	if o == nil || IsNil(o.Date) {
		var ret time.Time
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestResultShortGetModel) GetDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Date) {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *TestResultShortGetModel) HasDate() bool {
	if o != nil && !IsNil(o.Date) {
		return true
	}

	return false
}

// SetDate gets a reference to the given time.Time and assigns it to the Date field.
func (o *TestResultShortGetModel) SetDate(v time.Time) {
	o.Date = &v
}

// GetDuration returns the Duration field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestResultShortGetModel) GetDuration() int64 {
	if o == nil || IsNil(o.Duration.Get()) {
		var ret int64
		return ret
	}
	return *o.Duration.Get()
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestResultShortGetModel) GetDurationOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Duration.Get(), o.Duration.IsSet()
}

// HasDuration returns a boolean if a field has been set.
func (o *TestResultShortGetModel) HasDuration() bool {
	if o != nil && o.Duration.IsSet() {
		return true
	}

	return false
}

// SetDuration gets a reference to the given NullableInt64 and assigns it to the Duration field.
func (o *TestResultShortGetModel) SetDuration(v int64) {
	o.Duration.Set(&v)
}
// SetDurationNil sets the value for Duration to be an explicit nil
func (o *TestResultShortGetModel) SetDurationNil() {
	o.Duration.Set(nil)
}

// UnsetDuration ensures that no value is present for Duration, not even an explicit nil
func (o *TestResultShortGetModel) UnsetDuration() {
	o.Duration.Unset()
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *TestResultShortGetModel) GetLinks() []LinkSubGetModel {
	if o == nil || IsNil(o.Links) {
		var ret []LinkSubGetModel
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestResultShortGetModel) GetLinksOk() ([]LinkSubGetModel, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *TestResultShortGetModel) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []LinkSubGetModel and assigns it to the Links field.
func (o *TestResultShortGetModel) SetLinks(v []LinkSubGetModel) {
	o.Links = v
}

// GetAttachments returns the Attachments field value if set, zero value otherwise.
func (o *TestResultShortGetModel) GetAttachments() []AttachmentSubGetModel {
	if o == nil || IsNil(o.Attachments) {
		var ret []AttachmentSubGetModel
		return ret
	}
	return o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestResultShortGetModel) GetAttachmentsOk() ([]AttachmentSubGetModel, bool) {
	if o == nil || IsNil(o.Attachments) {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *TestResultShortGetModel) HasAttachments() bool {
	if o != nil && !IsNil(o.Attachments) {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given []AttachmentSubGetModel and assigns it to the Attachments field.
func (o *TestResultShortGetModel) SetAttachments(v []AttachmentSubGetModel) {
	o.Attachments = v
}

func (o TestResultShortGetModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TestResultShortGetModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.AutotestGlobalId) {
		toSerialize["autotestGlobalId"] = o.AutotestGlobalId
	}
	if !IsNil(o.TestRunId) {
		toSerialize["testRunId"] = o.TestRunId
	}
	if !IsNil(o.ConfigurationId) {
		toSerialize["configurationId"] = o.ConfigurationId
	}
	if !IsNil(o.ConfigurationName) {
		toSerialize["configurationName"] = o.ConfigurationName
	}
	toSerialize["outcome"] = o.Outcome
	if !IsNil(o.ResultReasons) {
		toSerialize["resultReasons"] = o.ResultReasons
	}
	if !IsNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	if !IsNil(o.Date) {
		toSerialize["date"] = o.Date
	}
	if o.Duration.IsSet() {
		toSerialize["duration"] = o.Duration.Get()
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Attachments) {
		toSerialize["attachments"] = o.Attachments
	}
	return toSerialize, nil
}

type NullableTestResultShortGetModel struct {
	value *TestResultShortGetModel
	isSet bool
}

func (v NullableTestResultShortGetModel) Get() *TestResultShortGetModel {
	return v.value
}

func (v *NullableTestResultShortGetModel) Set(val *TestResultShortGetModel) {
	v.value = val
	v.isSet = true
}

func (v NullableTestResultShortGetModel) IsSet() bool {
	return v.isSet
}

func (v *NullableTestResultShortGetModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestResultShortGetModel(val *TestResultShortGetModel) *NullableTestResultShortGetModel {
	return &NullableTestResultShortGetModel{value: val, isSet: true}
}

func (v NullableTestResultShortGetModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTestResultShortGetModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


