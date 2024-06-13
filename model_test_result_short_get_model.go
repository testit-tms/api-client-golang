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

// checks if the TestResultShortGetModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TestResultShortGetModel{}

// TestResultShortGetModel struct for TestResultShortGetModel
type TestResultShortGetModel struct {
	// Unique ID of the test result
	Id string `json:"id"`
	// Name of autotest represented by the test result
	Name string `json:"name"`
	// Global ID of autotest represented by the test result
	AutotestGlobalId int64 `json:"autotestGlobalId"`
	// Unique ID of test run where the test result is located
	TestRunId string `json:"testRunId"`
	// Unique ID of configuration which the test result uses
	ConfigurationId string `json:"configurationId"`
	// Name of configuration which the test result uses
	ConfigurationName string `json:"configurationName"`
	// Outcome of the test result
	Outcome string `json:"outcome"`
	// Collection of result reasons which the test result have
	ResultReasons []AutotestResultReasonSubGetModel `json:"resultReasons"`
	// Comment to the test result
	Comment NullableString `json:"comment,omitempty"`
	// Date when the test result was completed or started or created
	// Deprecated
	Date time.Time `json:"date"`
	// Date when the test result has been created
	CreatedDate time.Time `json:"createdDate"`
	// Date when the test result has been modified
	ModifiedDate NullableTime `json:"modifiedDate,omitempty"`
	// Date when the test result has been started
	StartedOn NullableTime `json:"startedOn,omitempty"`
	// Date when the test result has been completed
	CompletedOn NullableTime `json:"completedOn,omitempty"`
	// Time which it took to run the test
	Duration NullableInt64 `json:"duration,omitempty"`
	// Collection of links attached to the test result
	Links []LinkSubGetModel `json:"links"`
	// Collection of files attached to the test result
	Attachments []AttachmentModel `json:"attachments"`
}

type _TestResultShortGetModel TestResultShortGetModel

// NewTestResultShortGetModel instantiates a new TestResultShortGetModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTestResultShortGetModel(id string, name string, autotestGlobalId int64, testRunId string, configurationId string, configurationName string, outcome string, resultReasons []AutotestResultReasonSubGetModel, date time.Time, createdDate time.Time, links []LinkSubGetModel, attachments []AttachmentModel) *TestResultShortGetModel {
	this := TestResultShortGetModel{}
	this.Id = id
	this.Name = name
	this.AutotestGlobalId = autotestGlobalId
	this.TestRunId = testRunId
	this.ConfigurationId = configurationId
	this.ConfigurationName = configurationName
	this.Outcome = outcome
	this.ResultReasons = resultReasons
	this.Date = date
	this.CreatedDate = createdDate
	this.Links = links
	this.Attachments = attachments
	return &this
}

// NewTestResultShortGetModelWithDefaults instantiates a new TestResultShortGetModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTestResultShortGetModelWithDefaults() *TestResultShortGetModel {
	this := TestResultShortGetModel{}
	return &this
}

// GetId returns the Id field value
func (o *TestResultShortGetModel) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TestResultShortGetModel) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TestResultShortGetModel) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *TestResultShortGetModel) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TestResultShortGetModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TestResultShortGetModel) SetName(v string) {
	o.Name = v
}

// GetAutotestGlobalId returns the AutotestGlobalId field value
func (o *TestResultShortGetModel) GetAutotestGlobalId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.AutotestGlobalId
}

// GetAutotestGlobalIdOk returns a tuple with the AutotestGlobalId field value
// and a boolean to check if the value has been set.
func (o *TestResultShortGetModel) GetAutotestGlobalIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AutotestGlobalId, true
}

// SetAutotestGlobalId sets field value
func (o *TestResultShortGetModel) SetAutotestGlobalId(v int64) {
	o.AutotestGlobalId = v
}

// GetTestRunId returns the TestRunId field value
func (o *TestResultShortGetModel) GetTestRunId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TestRunId
}

// GetTestRunIdOk returns a tuple with the TestRunId field value
// and a boolean to check if the value has been set.
func (o *TestResultShortGetModel) GetTestRunIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TestRunId, true
}

// SetTestRunId sets field value
func (o *TestResultShortGetModel) SetTestRunId(v string) {
	o.TestRunId = v
}

// GetConfigurationId returns the ConfigurationId field value
func (o *TestResultShortGetModel) GetConfigurationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConfigurationId
}

// GetConfigurationIdOk returns a tuple with the ConfigurationId field value
// and a boolean to check if the value has been set.
func (o *TestResultShortGetModel) GetConfigurationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConfigurationId, true
}

// SetConfigurationId sets field value
func (o *TestResultShortGetModel) SetConfigurationId(v string) {
	o.ConfigurationId = v
}

// GetConfigurationName returns the ConfigurationName field value
func (o *TestResultShortGetModel) GetConfigurationName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConfigurationName
}

// GetConfigurationNameOk returns a tuple with the ConfigurationName field value
// and a boolean to check if the value has been set.
func (o *TestResultShortGetModel) GetConfigurationNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConfigurationName, true
}

// SetConfigurationName sets field value
func (o *TestResultShortGetModel) SetConfigurationName(v string) {
	o.ConfigurationName = v
}

// GetOutcome returns the Outcome field value
func (o *TestResultShortGetModel) GetOutcome() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Outcome
}

// GetOutcomeOk returns a tuple with the Outcome field value
// and a boolean to check if the value has been set.
func (o *TestResultShortGetModel) GetOutcomeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Outcome, true
}

// SetOutcome sets field value
func (o *TestResultShortGetModel) SetOutcome(v string) {
	o.Outcome = v
}

// GetResultReasons returns the ResultReasons field value
func (o *TestResultShortGetModel) GetResultReasons() []AutotestResultReasonSubGetModel {
	if o == nil {
		var ret []AutotestResultReasonSubGetModel
		return ret
	}

	return o.ResultReasons
}

// GetResultReasonsOk returns a tuple with the ResultReasons field value
// and a boolean to check if the value has been set.
func (o *TestResultShortGetModel) GetResultReasonsOk() ([]AutotestResultReasonSubGetModel, bool) {
	if o == nil {
		return nil, false
	}
	return o.ResultReasons, true
}

// SetResultReasons sets field value
func (o *TestResultShortGetModel) SetResultReasons(v []AutotestResultReasonSubGetModel) {
	o.ResultReasons = v
}

// GetComment returns the Comment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestResultShortGetModel) GetComment() string {
	if o == nil || IsNil(o.Comment.Get()) {
		var ret string
		return ret
	}
	return *o.Comment.Get()
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestResultShortGetModel) GetCommentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Comment.Get(), o.Comment.IsSet()
}

// HasComment returns a boolean if a field has been set.
func (o *TestResultShortGetModel) HasComment() bool {
	if o != nil && o.Comment.IsSet() {
		return true
	}

	return false
}

// SetComment gets a reference to the given NullableString and assigns it to the Comment field.
func (o *TestResultShortGetModel) SetComment(v string) {
	o.Comment.Set(&v)
}
// SetCommentNil sets the value for Comment to be an explicit nil
func (o *TestResultShortGetModel) SetCommentNil() {
	o.Comment.Set(nil)
}

// UnsetComment ensures that no value is present for Comment, not even an explicit nil
func (o *TestResultShortGetModel) UnsetComment() {
	o.Comment.Unset()
}

// GetDate returns the Date field value
// Deprecated
func (o *TestResultShortGetModel) GetDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Date
}

// GetDateOk returns a tuple with the Date field value
// and a boolean to check if the value has been set.
// Deprecated
func (o *TestResultShortGetModel) GetDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Date, true
}

// SetDate sets field value
// Deprecated
func (o *TestResultShortGetModel) SetDate(v time.Time) {
	o.Date = v
}

// GetCreatedDate returns the CreatedDate field value
func (o *TestResultShortGetModel) GetCreatedDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedDate
}

// GetCreatedDateOk returns a tuple with the CreatedDate field value
// and a boolean to check if the value has been set.
func (o *TestResultShortGetModel) GetCreatedDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedDate, true
}

// SetCreatedDate sets field value
func (o *TestResultShortGetModel) SetCreatedDate(v time.Time) {
	o.CreatedDate = v
}

// GetModifiedDate returns the ModifiedDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestResultShortGetModel) GetModifiedDate() time.Time {
	if o == nil || IsNil(o.ModifiedDate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.ModifiedDate.Get()
}

// GetModifiedDateOk returns a tuple with the ModifiedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestResultShortGetModel) GetModifiedDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ModifiedDate.Get(), o.ModifiedDate.IsSet()
}

// HasModifiedDate returns a boolean if a field has been set.
func (o *TestResultShortGetModel) HasModifiedDate() bool {
	if o != nil && o.ModifiedDate.IsSet() {
		return true
	}

	return false
}

// SetModifiedDate gets a reference to the given NullableTime and assigns it to the ModifiedDate field.
func (o *TestResultShortGetModel) SetModifiedDate(v time.Time) {
	o.ModifiedDate.Set(&v)
}
// SetModifiedDateNil sets the value for ModifiedDate to be an explicit nil
func (o *TestResultShortGetModel) SetModifiedDateNil() {
	o.ModifiedDate.Set(nil)
}

// UnsetModifiedDate ensures that no value is present for ModifiedDate, not even an explicit nil
func (o *TestResultShortGetModel) UnsetModifiedDate() {
	o.ModifiedDate.Unset()
}

// GetStartedOn returns the StartedOn field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestResultShortGetModel) GetStartedOn() time.Time {
	if o == nil || IsNil(o.StartedOn.Get()) {
		var ret time.Time
		return ret
	}
	return *o.StartedOn.Get()
}

// GetStartedOnOk returns a tuple with the StartedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestResultShortGetModel) GetStartedOnOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.StartedOn.Get(), o.StartedOn.IsSet()
}

// HasStartedOn returns a boolean if a field has been set.
func (o *TestResultShortGetModel) HasStartedOn() bool {
	if o != nil && o.StartedOn.IsSet() {
		return true
	}

	return false
}

// SetStartedOn gets a reference to the given NullableTime and assigns it to the StartedOn field.
func (o *TestResultShortGetModel) SetStartedOn(v time.Time) {
	o.StartedOn.Set(&v)
}
// SetStartedOnNil sets the value for StartedOn to be an explicit nil
func (o *TestResultShortGetModel) SetStartedOnNil() {
	o.StartedOn.Set(nil)
}

// UnsetStartedOn ensures that no value is present for StartedOn, not even an explicit nil
func (o *TestResultShortGetModel) UnsetStartedOn() {
	o.StartedOn.Unset()
}

// GetCompletedOn returns the CompletedOn field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestResultShortGetModel) GetCompletedOn() time.Time {
	if o == nil || IsNil(o.CompletedOn.Get()) {
		var ret time.Time
		return ret
	}
	return *o.CompletedOn.Get()
}

// GetCompletedOnOk returns a tuple with the CompletedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestResultShortGetModel) GetCompletedOnOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.CompletedOn.Get(), o.CompletedOn.IsSet()
}

// HasCompletedOn returns a boolean if a field has been set.
func (o *TestResultShortGetModel) HasCompletedOn() bool {
	if o != nil && o.CompletedOn.IsSet() {
		return true
	}

	return false
}

// SetCompletedOn gets a reference to the given NullableTime and assigns it to the CompletedOn field.
func (o *TestResultShortGetModel) SetCompletedOn(v time.Time) {
	o.CompletedOn.Set(&v)
}
// SetCompletedOnNil sets the value for CompletedOn to be an explicit nil
func (o *TestResultShortGetModel) SetCompletedOnNil() {
	o.CompletedOn.Set(nil)
}

// UnsetCompletedOn ensures that no value is present for CompletedOn, not even an explicit nil
func (o *TestResultShortGetModel) UnsetCompletedOn() {
	o.CompletedOn.Unset()
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

// GetLinks returns the Links field value
func (o *TestResultShortGetModel) GetLinks() []LinkSubGetModel {
	if o == nil {
		var ret []LinkSubGetModel
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *TestResultShortGetModel) GetLinksOk() ([]LinkSubGetModel, bool) {
	if o == nil {
		return nil, false
	}
	return o.Links, true
}

// SetLinks sets field value
func (o *TestResultShortGetModel) SetLinks(v []LinkSubGetModel) {
	o.Links = v
}

// GetAttachments returns the Attachments field value
func (o *TestResultShortGetModel) GetAttachments() []AttachmentModel {
	if o == nil {
		var ret []AttachmentModel
		return ret
	}

	return o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value
// and a boolean to check if the value has been set.
func (o *TestResultShortGetModel) GetAttachmentsOk() ([]AttachmentModel, bool) {
	if o == nil {
		return nil, false
	}
	return o.Attachments, true
}

// SetAttachments sets field value
func (o *TestResultShortGetModel) SetAttachments(v []AttachmentModel) {
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
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["autotestGlobalId"] = o.AutotestGlobalId
	toSerialize["testRunId"] = o.TestRunId
	toSerialize["configurationId"] = o.ConfigurationId
	toSerialize["configurationName"] = o.ConfigurationName
	toSerialize["outcome"] = o.Outcome
	toSerialize["resultReasons"] = o.ResultReasons
	if o.Comment.IsSet() {
		toSerialize["comment"] = o.Comment.Get()
	}
	toSerialize["date"] = o.Date
	toSerialize["createdDate"] = o.CreatedDate
	if o.ModifiedDate.IsSet() {
		toSerialize["modifiedDate"] = o.ModifiedDate.Get()
	}
	if o.StartedOn.IsSet() {
		toSerialize["startedOn"] = o.StartedOn.Get()
	}
	if o.CompletedOn.IsSet() {
		toSerialize["completedOn"] = o.CompletedOn.Get()
	}
	if o.Duration.IsSet() {
		toSerialize["duration"] = o.Duration.Get()
	}
	toSerialize["links"] = o.Links
	toSerialize["attachments"] = o.Attachments
	return toSerialize, nil
}

func (o *TestResultShortGetModel) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"autotestGlobalId",
		"testRunId",
		"configurationId",
		"configurationName",
		"outcome",
		"resultReasons",
		"date",
		"createdDate",
		"links",
		"attachments",
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

	varTestResultShortGetModel := _TestResultShortGetModel{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTestResultShortGetModel)

	if err != nil {
		return err
	}

	*o = TestResultShortGetModel(varTestResultShortGetModel)

	return err
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


