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

// checks if the NotificationModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NotificationModel{}

// NotificationModel struct for NotificationModel
type NotificationModel struct {
	Id *string `json:"id,omitempty"`
	CreatedDate NullableTime `json:"createdDate,omitempty"`
	IsRead *bool `json:"isRead,omitempty"`
	EntityId *string `json:"entityId,omitempty"`
	NotificationType NotificationTypeModel `json:"notificationType"`
	ProjectGlobalId NullableInt64 `json:"projectGlobalId,omitempty"`
	ProjectName NullableString `json:"projectName,omitempty"`
	TestPlanGlobalId *int64 `json:"testPlanGlobalId,omitempty"`
	TestPlanName *string `json:"testPlanName,omitempty"`
	WorkitemGlobalId NullableInt64 `json:"workitemGlobalId,omitempty"`
	Comment *string `json:"comment,omitempty"`
	WorkItemName *string `json:"workItemName,omitempty"`
	AttributeName NullableString `json:"attributeName,omitempty"`
	CreatedById *string `json:"createdById,omitempty"`
}

// NewNotificationModel instantiates a new NotificationModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationModel(notificationType NotificationTypeModel) *NotificationModel {
	this := NotificationModel{}
	this.NotificationType = notificationType
	return &this
}

// NewNotificationModelWithDefaults instantiates a new NotificationModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationModelWithDefaults() *NotificationModel {
	this := NotificationModel{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *NotificationModel) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationModel) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *NotificationModel) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *NotificationModel) SetId(v string) {
	o.Id = &v
}

// GetCreatedDate returns the CreatedDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NotificationModel) GetCreatedDate() time.Time {
	if o == nil || IsNil(o.CreatedDate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.CreatedDate.Get()
}

// GetCreatedDateOk returns a tuple with the CreatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NotificationModel) GetCreatedDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedDate.Get(), o.CreatedDate.IsSet()
}

// HasCreatedDate returns a boolean if a field has been set.
func (o *NotificationModel) HasCreatedDate() bool {
	if o != nil && o.CreatedDate.IsSet() {
		return true
	}

	return false
}

// SetCreatedDate gets a reference to the given NullableTime and assigns it to the CreatedDate field.
func (o *NotificationModel) SetCreatedDate(v time.Time) {
	o.CreatedDate.Set(&v)
}
// SetCreatedDateNil sets the value for CreatedDate to be an explicit nil
func (o *NotificationModel) SetCreatedDateNil() {
	o.CreatedDate.Set(nil)
}

// UnsetCreatedDate ensures that no value is present for CreatedDate, not even an explicit nil
func (o *NotificationModel) UnsetCreatedDate() {
	o.CreatedDate.Unset()
}

// GetIsRead returns the IsRead field value if set, zero value otherwise.
func (o *NotificationModel) GetIsRead() bool {
	if o == nil || IsNil(o.IsRead) {
		var ret bool
		return ret
	}
	return *o.IsRead
}

// GetIsReadOk returns a tuple with the IsRead field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationModel) GetIsReadOk() (*bool, bool) {
	if o == nil || IsNil(o.IsRead) {
		return nil, false
	}
	return o.IsRead, true
}

// HasIsRead returns a boolean if a field has been set.
func (o *NotificationModel) HasIsRead() bool {
	if o != nil && !IsNil(o.IsRead) {
		return true
	}

	return false
}

// SetIsRead gets a reference to the given bool and assigns it to the IsRead field.
func (o *NotificationModel) SetIsRead(v bool) {
	o.IsRead = &v
}

// GetEntityId returns the EntityId field value if set, zero value otherwise.
func (o *NotificationModel) GetEntityId() string {
	if o == nil || IsNil(o.EntityId) {
		var ret string
		return ret
	}
	return *o.EntityId
}

// GetEntityIdOk returns a tuple with the EntityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationModel) GetEntityIdOk() (*string, bool) {
	if o == nil || IsNil(o.EntityId) {
		return nil, false
	}
	return o.EntityId, true
}

// HasEntityId returns a boolean if a field has been set.
func (o *NotificationModel) HasEntityId() bool {
	if o != nil && !IsNil(o.EntityId) {
		return true
	}

	return false
}

// SetEntityId gets a reference to the given string and assigns it to the EntityId field.
func (o *NotificationModel) SetEntityId(v string) {
	o.EntityId = &v
}

// GetNotificationType returns the NotificationType field value
func (o *NotificationModel) GetNotificationType() NotificationTypeModel {
	if o == nil {
		var ret NotificationTypeModel
		return ret
	}

	return o.NotificationType
}

// GetNotificationTypeOk returns a tuple with the NotificationType field value
// and a boolean to check if the value has been set.
func (o *NotificationModel) GetNotificationTypeOk() (*NotificationTypeModel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NotificationType, true
}

// SetNotificationType sets field value
func (o *NotificationModel) SetNotificationType(v NotificationTypeModel) {
	o.NotificationType = v
}

// GetProjectGlobalId returns the ProjectGlobalId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NotificationModel) GetProjectGlobalId() int64 {
	if o == nil || IsNil(o.ProjectGlobalId.Get()) {
		var ret int64
		return ret
	}
	return *o.ProjectGlobalId.Get()
}

// GetProjectGlobalIdOk returns a tuple with the ProjectGlobalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NotificationModel) GetProjectGlobalIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProjectGlobalId.Get(), o.ProjectGlobalId.IsSet()
}

// HasProjectGlobalId returns a boolean if a field has been set.
func (o *NotificationModel) HasProjectGlobalId() bool {
	if o != nil && o.ProjectGlobalId.IsSet() {
		return true
	}

	return false
}

// SetProjectGlobalId gets a reference to the given NullableInt64 and assigns it to the ProjectGlobalId field.
func (o *NotificationModel) SetProjectGlobalId(v int64) {
	o.ProjectGlobalId.Set(&v)
}
// SetProjectGlobalIdNil sets the value for ProjectGlobalId to be an explicit nil
func (o *NotificationModel) SetProjectGlobalIdNil() {
	o.ProjectGlobalId.Set(nil)
}

// UnsetProjectGlobalId ensures that no value is present for ProjectGlobalId, not even an explicit nil
func (o *NotificationModel) UnsetProjectGlobalId() {
	o.ProjectGlobalId.Unset()
}

// GetProjectName returns the ProjectName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NotificationModel) GetProjectName() string {
	if o == nil || IsNil(o.ProjectName.Get()) {
		var ret string
		return ret
	}
	return *o.ProjectName.Get()
}

// GetProjectNameOk returns a tuple with the ProjectName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NotificationModel) GetProjectNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProjectName.Get(), o.ProjectName.IsSet()
}

// HasProjectName returns a boolean if a field has been set.
func (o *NotificationModel) HasProjectName() bool {
	if o != nil && o.ProjectName.IsSet() {
		return true
	}

	return false
}

// SetProjectName gets a reference to the given NullableString and assigns it to the ProjectName field.
func (o *NotificationModel) SetProjectName(v string) {
	o.ProjectName.Set(&v)
}
// SetProjectNameNil sets the value for ProjectName to be an explicit nil
func (o *NotificationModel) SetProjectNameNil() {
	o.ProjectName.Set(nil)
}

// UnsetProjectName ensures that no value is present for ProjectName, not even an explicit nil
func (o *NotificationModel) UnsetProjectName() {
	o.ProjectName.Unset()
}

// GetTestPlanGlobalId returns the TestPlanGlobalId field value if set, zero value otherwise.
func (o *NotificationModel) GetTestPlanGlobalId() int64 {
	if o == nil || IsNil(o.TestPlanGlobalId) {
		var ret int64
		return ret
	}
	return *o.TestPlanGlobalId
}

// GetTestPlanGlobalIdOk returns a tuple with the TestPlanGlobalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationModel) GetTestPlanGlobalIdOk() (*int64, bool) {
	if o == nil || IsNil(o.TestPlanGlobalId) {
		return nil, false
	}
	return o.TestPlanGlobalId, true
}

// HasTestPlanGlobalId returns a boolean if a field has been set.
func (o *NotificationModel) HasTestPlanGlobalId() bool {
	if o != nil && !IsNil(o.TestPlanGlobalId) {
		return true
	}

	return false
}

// SetTestPlanGlobalId gets a reference to the given int64 and assigns it to the TestPlanGlobalId field.
func (o *NotificationModel) SetTestPlanGlobalId(v int64) {
	o.TestPlanGlobalId = &v
}

// GetTestPlanName returns the TestPlanName field value if set, zero value otherwise.
func (o *NotificationModel) GetTestPlanName() string {
	if o == nil || IsNil(o.TestPlanName) {
		var ret string
		return ret
	}
	return *o.TestPlanName
}

// GetTestPlanNameOk returns a tuple with the TestPlanName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationModel) GetTestPlanNameOk() (*string, bool) {
	if o == nil || IsNil(o.TestPlanName) {
		return nil, false
	}
	return o.TestPlanName, true
}

// HasTestPlanName returns a boolean if a field has been set.
func (o *NotificationModel) HasTestPlanName() bool {
	if o != nil && !IsNil(o.TestPlanName) {
		return true
	}

	return false
}

// SetTestPlanName gets a reference to the given string and assigns it to the TestPlanName field.
func (o *NotificationModel) SetTestPlanName(v string) {
	o.TestPlanName = &v
}

// GetWorkitemGlobalId returns the WorkitemGlobalId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NotificationModel) GetWorkitemGlobalId() int64 {
	if o == nil || IsNil(o.WorkitemGlobalId.Get()) {
		var ret int64
		return ret
	}
	return *o.WorkitemGlobalId.Get()
}

// GetWorkitemGlobalIdOk returns a tuple with the WorkitemGlobalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NotificationModel) GetWorkitemGlobalIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.WorkitemGlobalId.Get(), o.WorkitemGlobalId.IsSet()
}

// HasWorkitemGlobalId returns a boolean if a field has been set.
func (o *NotificationModel) HasWorkitemGlobalId() bool {
	if o != nil && o.WorkitemGlobalId.IsSet() {
		return true
	}

	return false
}

// SetWorkitemGlobalId gets a reference to the given NullableInt64 and assigns it to the WorkitemGlobalId field.
func (o *NotificationModel) SetWorkitemGlobalId(v int64) {
	o.WorkitemGlobalId.Set(&v)
}
// SetWorkitemGlobalIdNil sets the value for WorkitemGlobalId to be an explicit nil
func (o *NotificationModel) SetWorkitemGlobalIdNil() {
	o.WorkitemGlobalId.Set(nil)
}

// UnsetWorkitemGlobalId ensures that no value is present for WorkitemGlobalId, not even an explicit nil
func (o *NotificationModel) UnsetWorkitemGlobalId() {
	o.WorkitemGlobalId.Unset()
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *NotificationModel) GetComment() string {
	if o == nil || IsNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationModel) GetCommentOk() (*string, bool) {
	if o == nil || IsNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *NotificationModel) HasComment() bool {
	if o != nil && !IsNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *NotificationModel) SetComment(v string) {
	o.Comment = &v
}

// GetWorkItemName returns the WorkItemName field value if set, zero value otherwise.
func (o *NotificationModel) GetWorkItemName() string {
	if o == nil || IsNil(o.WorkItemName) {
		var ret string
		return ret
	}
	return *o.WorkItemName
}

// GetWorkItemNameOk returns a tuple with the WorkItemName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationModel) GetWorkItemNameOk() (*string, bool) {
	if o == nil || IsNil(o.WorkItemName) {
		return nil, false
	}
	return o.WorkItemName, true
}

// HasWorkItemName returns a boolean if a field has been set.
func (o *NotificationModel) HasWorkItemName() bool {
	if o != nil && !IsNil(o.WorkItemName) {
		return true
	}

	return false
}

// SetWorkItemName gets a reference to the given string and assigns it to the WorkItemName field.
func (o *NotificationModel) SetWorkItemName(v string) {
	o.WorkItemName = &v
}

// GetAttributeName returns the AttributeName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NotificationModel) GetAttributeName() string {
	if o == nil || IsNil(o.AttributeName.Get()) {
		var ret string
		return ret
	}
	return *o.AttributeName.Get()
}

// GetAttributeNameOk returns a tuple with the AttributeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NotificationModel) GetAttributeNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AttributeName.Get(), o.AttributeName.IsSet()
}

// HasAttributeName returns a boolean if a field has been set.
func (o *NotificationModel) HasAttributeName() bool {
	if o != nil && o.AttributeName.IsSet() {
		return true
	}

	return false
}

// SetAttributeName gets a reference to the given NullableString and assigns it to the AttributeName field.
func (o *NotificationModel) SetAttributeName(v string) {
	o.AttributeName.Set(&v)
}
// SetAttributeNameNil sets the value for AttributeName to be an explicit nil
func (o *NotificationModel) SetAttributeNameNil() {
	o.AttributeName.Set(nil)
}

// UnsetAttributeName ensures that no value is present for AttributeName, not even an explicit nil
func (o *NotificationModel) UnsetAttributeName() {
	o.AttributeName.Unset()
}

// GetCreatedById returns the CreatedById field value if set, zero value otherwise.
func (o *NotificationModel) GetCreatedById() string {
	if o == nil || IsNil(o.CreatedById) {
		var ret string
		return ret
	}
	return *o.CreatedById
}

// GetCreatedByIdOk returns a tuple with the CreatedById field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationModel) GetCreatedByIdOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedById) {
		return nil, false
	}
	return o.CreatedById, true
}

// HasCreatedById returns a boolean if a field has been set.
func (o *NotificationModel) HasCreatedById() bool {
	if o != nil && !IsNil(o.CreatedById) {
		return true
	}

	return false
}

// SetCreatedById gets a reference to the given string and assigns it to the CreatedById field.
func (o *NotificationModel) SetCreatedById(v string) {
	o.CreatedById = &v
}

func (o NotificationModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NotificationModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.CreatedDate.IsSet() {
		toSerialize["createdDate"] = o.CreatedDate.Get()
	}
	if !IsNil(o.IsRead) {
		toSerialize["isRead"] = o.IsRead
	}
	if !IsNil(o.EntityId) {
		toSerialize["entityId"] = o.EntityId
	}
	toSerialize["notificationType"] = o.NotificationType
	if o.ProjectGlobalId.IsSet() {
		toSerialize["projectGlobalId"] = o.ProjectGlobalId.Get()
	}
	if o.ProjectName.IsSet() {
		toSerialize["projectName"] = o.ProjectName.Get()
	}
	if !IsNil(o.TestPlanGlobalId) {
		toSerialize["testPlanGlobalId"] = o.TestPlanGlobalId
	}
	if !IsNil(o.TestPlanName) {
		toSerialize["testPlanName"] = o.TestPlanName
	}
	if o.WorkitemGlobalId.IsSet() {
		toSerialize["workitemGlobalId"] = o.WorkitemGlobalId.Get()
	}
	if !IsNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	if !IsNil(o.WorkItemName) {
		toSerialize["workItemName"] = o.WorkItemName
	}
	if o.AttributeName.IsSet() {
		toSerialize["attributeName"] = o.AttributeName.Get()
	}
	if !IsNil(o.CreatedById) {
		toSerialize["createdById"] = o.CreatedById
	}
	return toSerialize, nil
}

type NullableNotificationModel struct {
	value *NotificationModel
	isSet bool
}

func (v NullableNotificationModel) Get() *NotificationModel {
	return v.value
}

func (v *NullableNotificationModel) Set(val *NotificationModel) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationModel) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationModel(val *NotificationModel) *NullableNotificationModel {
	return &NullableNotificationModel{value: val, isSet: true}
}

func (v NullableNotificationModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


