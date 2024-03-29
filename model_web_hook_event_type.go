/*
API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tmsclient

import (
	"encoding/json"
	"fmt"
)

// WebHookEventType the model 'WebHookEventType'
type WebHookEventType string

// List of WebHookEventType
const (
	WEBHOOKEVENTTYPE_AUTOMATED_TEST_RUN_CREATED WebHookEventType = "AutomatedTestRunCreated"
	WEBHOOKEVENTTYPE_TEST_PLANS_STATUS_CHANGED WebHookEventType = "TestPlansStatusChanged"
	WEBHOOKEVENTTYPE_TEST_RUN_STOPPED WebHookEventType = "TestRunStopped"
	WEBHOOKEVENTTYPE_TEST_POINT_ASSIGNED WebHookEventType = "TestPointAssigned"
	WEBHOOKEVENTTYPE_TEST_RESULT_JIRA_ISSUE_CREATED WebHookEventType = "TestResultJiraIssueCreated"
	WEBHOOKEVENTTYPE_AUTO_TEST_FINISHED WebHookEventType = "AutoTestFinished"
	WEBHOOKEVENTTYPE_USER_MENTIONED_IN_COMMENT WebHookEventType = "UserMentionedInComment"
	WEBHOOKEVENTTYPE_USER_SELECTED_IN_WORK_ITEM_ATTRIBUTE WebHookEventType = "UserSelectedInWorkItemAttribute"
	WEBHOOKEVENTTYPE_ALL_TEST_POINTS_FINISHED WebHookEventType = "AllTestPointsFinished"
	WEBHOOKEVENTTYPE_ALL_AUTO_TESTS_FINISHED WebHookEventType = "AllAutoTestsFinished"
	WEBHOOKEVENTTYPE_AUTO_TEST_CHANGED WebHookEventType = "AutoTestChanged"
	WEBHOOKEVENTTYPE_WORK_ITEM_AUTO_TEST_RELATION_CHANGED WebHookEventType = "WorkItemAutoTestRelationChanged"
	WEBHOOKEVENTTYPE_WORK_ITEM_ATTRIBUTE_CHANGED WebHookEventType = "WorkItemAttributeChanged"
	WEBHOOKEVENTTYPE_WORK_ITEM_CHANGED WebHookEventType = "WorkItemChanged"
	WEBHOOKEVENTTYPE_CONFIGURATION_CHANGED WebHookEventType = "ConfigurationChanged"
	WEBHOOKEVENTTYPE_PROJECT_CHANGED WebHookEventType = "ProjectChanged"
	WEBHOOKEVENTTYPE_TEST_PLAN_CHANGED WebHookEventType = "TestPlanChanged"
)

// All allowed values of WebHookEventType enum
var AllowedWebHookEventTypeEnumValues = []WebHookEventType{
	"AutomatedTestRunCreated",
	"TestPlansStatusChanged",
	"TestRunStopped",
	"TestPointAssigned",
	"TestResultJiraIssueCreated",
	"AutoTestFinished",
	"UserMentionedInComment",
	"UserSelectedInWorkItemAttribute",
	"AllTestPointsFinished",
	"AllAutoTestsFinished",
	"AutoTestChanged",
	"WorkItemAutoTestRelationChanged",
	"WorkItemAttributeChanged",
	"WorkItemChanged",
	"ConfigurationChanged",
	"ProjectChanged",
	"TestPlanChanged",
}

func (v *WebHookEventType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := WebHookEventType(value)
	for _, existing := range AllowedWebHookEventTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid WebHookEventType", value)
}

// NewWebHookEventTypeFromValue returns a pointer to a valid WebHookEventType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewWebHookEventTypeFromValue(v string) (*WebHookEventType, error) {
	ev := WebHookEventType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for WebHookEventType: valid values are %v", v, AllowedWebHookEventTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v WebHookEventType) IsValid() bool {
	for _, existing := range AllowedWebHookEventTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to WebHookEventType value
func (v WebHookEventType) Ptr() *WebHookEventType {
	return &v
}

type NullableWebHookEventType struct {
	value *WebHookEventType
	isSet bool
}

func (v NullableWebHookEventType) Get() *WebHookEventType {
	return v.value
}

func (v *NullableWebHookEventType) Set(val *WebHookEventType) {
	v.value = val
	v.isSet = true
}

func (v NullableWebHookEventType) IsSet() bool {
	return v.isSet
}

func (v *NullableWebHookEventType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebHookEventType(val *WebHookEventType) *NullableWebHookEventType {
	return &NullableWebHookEventType{value: val, isSet: true}
}

func (v NullableWebHookEventType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebHookEventType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

