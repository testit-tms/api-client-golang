/*
API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tmsclient

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the TestRunFillByWorkItemsPostModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TestRunFillByWorkItemsPostModel{}

// TestRunFillByWorkItemsPostModel struct for TestRunFillByWorkItemsPostModel
type TestRunFillByWorkItemsPostModel struct {
	// Specifies the configuration GUIDs, from which test points are created. You can specify several GUIDs.
	ConfigurationIds []string `json:"configurationIds"`
	// Specifies the work item GUIDs, from which test points are created. You can specify several GUIDs.
	WorkItemIds []string `json:"workItemIds"`
	// Specifies the GUID of the project, in which a test run will be created.
	ProjectId string `json:"projectId"`
	// Specifies the GUID of the test plan, within which the test run will be created.
	TestPlanId string `json:"testPlanId"`
	// Specifies the name of the test run.
	Name NullableString `json:"name,omitempty"`
	// Specifies the test run description.
	Description NullableString `json:"description,omitempty"`
	// Specifies the test run launch source.
	LaunchSource NullableString `json:"launchSource,omitempty"`
	// Collection of attachment ids to relate to the test run
	Attachments []AttachmentPutModel `json:"attachments,omitempty"`
	// Collection of links to relate to the test run
	Links []LinkPostModel `json:"links,omitempty"`
}

type _TestRunFillByWorkItemsPostModel TestRunFillByWorkItemsPostModel

// NewTestRunFillByWorkItemsPostModel instantiates a new TestRunFillByWorkItemsPostModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTestRunFillByWorkItemsPostModel(configurationIds []string, workItemIds []string, projectId string, testPlanId string) *TestRunFillByWorkItemsPostModel {
	this := TestRunFillByWorkItemsPostModel{}
	this.ConfigurationIds = configurationIds
	this.WorkItemIds = workItemIds
	this.ProjectId = projectId
	this.TestPlanId = testPlanId
	return &this
}

// NewTestRunFillByWorkItemsPostModelWithDefaults instantiates a new TestRunFillByWorkItemsPostModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTestRunFillByWorkItemsPostModelWithDefaults() *TestRunFillByWorkItemsPostModel {
	this := TestRunFillByWorkItemsPostModel{}
	return &this
}

// GetConfigurationIds returns the ConfigurationIds field value
func (o *TestRunFillByWorkItemsPostModel) GetConfigurationIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ConfigurationIds
}

// GetConfigurationIdsOk returns a tuple with the ConfigurationIds field value
// and a boolean to check if the value has been set.
func (o *TestRunFillByWorkItemsPostModel) GetConfigurationIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConfigurationIds, true
}

// SetConfigurationIds sets field value
func (o *TestRunFillByWorkItemsPostModel) SetConfigurationIds(v []string) {
	o.ConfigurationIds = v
}

// GetWorkItemIds returns the WorkItemIds field value
func (o *TestRunFillByWorkItemsPostModel) GetWorkItemIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.WorkItemIds
}

// GetWorkItemIdsOk returns a tuple with the WorkItemIds field value
// and a boolean to check if the value has been set.
func (o *TestRunFillByWorkItemsPostModel) GetWorkItemIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.WorkItemIds, true
}

// SetWorkItemIds sets field value
func (o *TestRunFillByWorkItemsPostModel) SetWorkItemIds(v []string) {
	o.WorkItemIds = v
}

// GetProjectId returns the ProjectId field value
func (o *TestRunFillByWorkItemsPostModel) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *TestRunFillByWorkItemsPostModel) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *TestRunFillByWorkItemsPostModel) SetProjectId(v string) {
	o.ProjectId = v
}

// GetTestPlanId returns the TestPlanId field value
func (o *TestRunFillByWorkItemsPostModel) GetTestPlanId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TestPlanId
}

// GetTestPlanIdOk returns a tuple with the TestPlanId field value
// and a boolean to check if the value has been set.
func (o *TestRunFillByWorkItemsPostModel) GetTestPlanIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TestPlanId, true
}

// SetTestPlanId sets field value
func (o *TestRunFillByWorkItemsPostModel) SetTestPlanId(v string) {
	o.TestPlanId = v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestRunFillByWorkItemsPostModel) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestRunFillByWorkItemsPostModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *TestRunFillByWorkItemsPostModel) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *TestRunFillByWorkItemsPostModel) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *TestRunFillByWorkItemsPostModel) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *TestRunFillByWorkItemsPostModel) UnsetName() {
	o.Name.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestRunFillByWorkItemsPostModel) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestRunFillByWorkItemsPostModel) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *TestRunFillByWorkItemsPostModel) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *TestRunFillByWorkItemsPostModel) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *TestRunFillByWorkItemsPostModel) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *TestRunFillByWorkItemsPostModel) UnsetDescription() {
	o.Description.Unset()
}

// GetLaunchSource returns the LaunchSource field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestRunFillByWorkItemsPostModel) GetLaunchSource() string {
	if o == nil || IsNil(o.LaunchSource.Get()) {
		var ret string
		return ret
	}
	return *o.LaunchSource.Get()
}

// GetLaunchSourceOk returns a tuple with the LaunchSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestRunFillByWorkItemsPostModel) GetLaunchSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LaunchSource.Get(), o.LaunchSource.IsSet()
}

// HasLaunchSource returns a boolean if a field has been set.
func (o *TestRunFillByWorkItemsPostModel) HasLaunchSource() bool {
	if o != nil && o.LaunchSource.IsSet() {
		return true
	}

	return false
}

// SetLaunchSource gets a reference to the given NullableString and assigns it to the LaunchSource field.
func (o *TestRunFillByWorkItemsPostModel) SetLaunchSource(v string) {
	o.LaunchSource.Set(&v)
}
// SetLaunchSourceNil sets the value for LaunchSource to be an explicit nil
func (o *TestRunFillByWorkItemsPostModel) SetLaunchSourceNil() {
	o.LaunchSource.Set(nil)
}

// UnsetLaunchSource ensures that no value is present for LaunchSource, not even an explicit nil
func (o *TestRunFillByWorkItemsPostModel) UnsetLaunchSource() {
	o.LaunchSource.Unset()
}

// GetAttachments returns the Attachments field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestRunFillByWorkItemsPostModel) GetAttachments() []AttachmentPutModel {
	if o == nil {
		var ret []AttachmentPutModel
		return ret
	}
	return o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestRunFillByWorkItemsPostModel) GetAttachmentsOk() ([]AttachmentPutModel, bool) {
	if o == nil || IsNil(o.Attachments) {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *TestRunFillByWorkItemsPostModel) HasAttachments() bool {
	if o != nil && IsNil(o.Attachments) {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given []AttachmentPutModel and assigns it to the Attachments field.
func (o *TestRunFillByWorkItemsPostModel) SetAttachments(v []AttachmentPutModel) {
	o.Attachments = v
}

// GetLinks returns the Links field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestRunFillByWorkItemsPostModel) GetLinks() []LinkPostModel {
	if o == nil {
		var ret []LinkPostModel
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestRunFillByWorkItemsPostModel) GetLinksOk() ([]LinkPostModel, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *TestRunFillByWorkItemsPostModel) HasLinks() bool {
	if o != nil && IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []LinkPostModel and assigns it to the Links field.
func (o *TestRunFillByWorkItemsPostModel) SetLinks(v []LinkPostModel) {
	o.Links = v
}

func (o TestRunFillByWorkItemsPostModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TestRunFillByWorkItemsPostModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["configurationIds"] = o.ConfigurationIds
	toSerialize["workItemIds"] = o.WorkItemIds
	toSerialize["projectId"] = o.ProjectId
	toSerialize["testPlanId"] = o.TestPlanId
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.LaunchSource.IsSet() {
		toSerialize["launchSource"] = o.LaunchSource.Get()
	}
	if o.Attachments != nil {
		toSerialize["attachments"] = o.Attachments
	}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	return toSerialize, nil
}

func (o *TestRunFillByWorkItemsPostModel) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"configurationIds",
		"workItemIds",
		"projectId",
		"testPlanId",
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

	varTestRunFillByWorkItemsPostModel := _TestRunFillByWorkItemsPostModel{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTestRunFillByWorkItemsPostModel)

	if err != nil {
		return err
	}

	*o = TestRunFillByWorkItemsPostModel(varTestRunFillByWorkItemsPostModel)

	return err
}

type NullableTestRunFillByWorkItemsPostModel struct {
	value *TestRunFillByWorkItemsPostModel
	isSet bool
}

func (v NullableTestRunFillByWorkItemsPostModel) Get() *TestRunFillByWorkItemsPostModel {
	return v.value
}

func (v *NullableTestRunFillByWorkItemsPostModel) Set(val *TestRunFillByWorkItemsPostModel) {
	v.value = val
	v.isSet = true
}

func (v NullableTestRunFillByWorkItemsPostModel) IsSet() bool {
	return v.isSet
}

func (v *NullableTestRunFillByWorkItemsPostModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestRunFillByWorkItemsPostModel(val *TestRunFillByWorkItemsPostModel) *NullableTestRunFillByWorkItemsPostModel {
	return &NullableTestRunFillByWorkItemsPostModel{value: val, isSet: true}
}

func (v NullableTestRunFillByWorkItemsPostModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTestRunFillByWorkItemsPostModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


