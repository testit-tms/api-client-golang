# TestResultShortApiResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Outcome** | **string** |  | 
**Status** | [**TestStatusApiResult**](TestStatusApiResult.md) |  | 
**Traces** | Pointer to **NullableString** |  | [optional] 
**FailureType** | Pointer to **NullableString** |  | [optional] 
**Message** | Pointer to **NullableString** |  | [optional] 
**TestPoint** | Pointer to [**NullableTestPointShortApiResult**](TestPointShortApiResult.md) |  | [optional] 
**CreatedDate** | **time.Time** |  | 
**AutoTest** | Pointer to [**NullableAutoTestShortApiResult**](AutoTestShortApiResult.md) |  | [optional] 
**Attachments** | [**[]AttachmentApiResult**](AttachmentApiResult.md) |  | 

## Methods

### NewTestResultShortApiResult

`func NewTestResultShortApiResult(id string, outcome string, status TestStatusApiResult, createdDate time.Time, attachments []AttachmentApiResult, ) *TestResultShortApiResult`

NewTestResultShortApiResult instantiates a new TestResultShortApiResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestResultShortApiResultWithDefaults

`func NewTestResultShortApiResultWithDefaults() *TestResultShortApiResult`

NewTestResultShortApiResultWithDefaults instantiates a new TestResultShortApiResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TestResultShortApiResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TestResultShortApiResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TestResultShortApiResult) SetId(v string)`

SetId sets Id field to given value.


### GetOutcome

`func (o *TestResultShortApiResult) GetOutcome() string`

GetOutcome returns the Outcome field if non-nil, zero value otherwise.

### GetOutcomeOk

`func (o *TestResultShortApiResult) GetOutcomeOk() (*string, bool)`

GetOutcomeOk returns a tuple with the Outcome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcome

`func (o *TestResultShortApiResult) SetOutcome(v string)`

SetOutcome sets Outcome field to given value.


### GetStatus

`func (o *TestResultShortApiResult) GetStatus() TestStatusApiResult`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TestResultShortApiResult) GetStatusOk() (*TestStatusApiResult, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TestResultShortApiResult) SetStatus(v TestStatusApiResult)`

SetStatus sets Status field to given value.


### GetTraces

`func (o *TestResultShortApiResult) GetTraces() string`

GetTraces returns the Traces field if non-nil, zero value otherwise.

### GetTracesOk

`func (o *TestResultShortApiResult) GetTracesOk() (*string, bool)`

GetTracesOk returns a tuple with the Traces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraces

`func (o *TestResultShortApiResult) SetTraces(v string)`

SetTraces sets Traces field to given value.

### HasTraces

`func (o *TestResultShortApiResult) HasTraces() bool`

HasTraces returns a boolean if a field has been set.

### SetTracesNil

`func (o *TestResultShortApiResult) SetTracesNil(b bool)`

 SetTracesNil sets the value for Traces to be an explicit nil

### UnsetTraces
`func (o *TestResultShortApiResult) UnsetTraces()`

UnsetTraces ensures that no value is present for Traces, not even an explicit nil
### GetFailureType

`func (o *TestResultShortApiResult) GetFailureType() string`

GetFailureType returns the FailureType field if non-nil, zero value otherwise.

### GetFailureTypeOk

`func (o *TestResultShortApiResult) GetFailureTypeOk() (*string, bool)`

GetFailureTypeOk returns a tuple with the FailureType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureType

`func (o *TestResultShortApiResult) SetFailureType(v string)`

SetFailureType sets FailureType field to given value.

### HasFailureType

`func (o *TestResultShortApiResult) HasFailureType() bool`

HasFailureType returns a boolean if a field has been set.

### SetFailureTypeNil

`func (o *TestResultShortApiResult) SetFailureTypeNil(b bool)`

 SetFailureTypeNil sets the value for FailureType to be an explicit nil

### UnsetFailureType
`func (o *TestResultShortApiResult) UnsetFailureType()`

UnsetFailureType ensures that no value is present for FailureType, not even an explicit nil
### GetMessage

`func (o *TestResultShortApiResult) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *TestResultShortApiResult) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *TestResultShortApiResult) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *TestResultShortApiResult) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *TestResultShortApiResult) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *TestResultShortApiResult) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetTestPoint

`func (o *TestResultShortApiResult) GetTestPoint() TestPointShortApiResult`

GetTestPoint returns the TestPoint field if non-nil, zero value otherwise.

### GetTestPointOk

`func (o *TestResultShortApiResult) GetTestPointOk() (*TestPointShortApiResult, bool)`

GetTestPointOk returns a tuple with the TestPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestPoint

`func (o *TestResultShortApiResult) SetTestPoint(v TestPointShortApiResult)`

SetTestPoint sets TestPoint field to given value.

### HasTestPoint

`func (o *TestResultShortApiResult) HasTestPoint() bool`

HasTestPoint returns a boolean if a field has been set.

### SetTestPointNil

`func (o *TestResultShortApiResult) SetTestPointNil(b bool)`

 SetTestPointNil sets the value for TestPoint to be an explicit nil

### UnsetTestPoint
`func (o *TestResultShortApiResult) UnsetTestPoint()`

UnsetTestPoint ensures that no value is present for TestPoint, not even an explicit nil
### GetCreatedDate

`func (o *TestResultShortApiResult) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *TestResultShortApiResult) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *TestResultShortApiResult) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.


### GetAutoTest

`func (o *TestResultShortApiResult) GetAutoTest() AutoTestShortApiResult`

GetAutoTest returns the AutoTest field if non-nil, zero value otherwise.

### GetAutoTestOk

`func (o *TestResultShortApiResult) GetAutoTestOk() (*AutoTestShortApiResult, bool)`

GetAutoTestOk returns a tuple with the AutoTest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoTest

`func (o *TestResultShortApiResult) SetAutoTest(v AutoTestShortApiResult)`

SetAutoTest sets AutoTest field to given value.

### HasAutoTest

`func (o *TestResultShortApiResult) HasAutoTest() bool`

HasAutoTest returns a boolean if a field has been set.

### SetAutoTestNil

`func (o *TestResultShortApiResult) SetAutoTestNil(b bool)`

 SetAutoTestNil sets the value for AutoTest to be an explicit nil

### UnsetAutoTest
`func (o *TestResultShortApiResult) UnsetAutoTest()`

UnsetAutoTest ensures that no value is present for AutoTest, not even an explicit nil
### GetAttachments

`func (o *TestResultShortApiResult) GetAttachments() []AttachmentApiResult`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *TestResultShortApiResult) GetAttachmentsOk() (*[]AttachmentApiResult, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *TestResultShortApiResult) SetAttachments(v []AttachmentApiResult)`

SetAttachments sets Attachments field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


