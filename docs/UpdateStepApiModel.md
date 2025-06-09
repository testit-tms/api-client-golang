# UpdateStepApiModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Step unique internal identifier | 
**Action** | Pointer to **NullableString** | Action applied by user | [optional] 
**Expected** | Pointer to **NullableString** | Expected system reaction | [optional] 
**TestData** | Pointer to **NullableString** | Test data for step | [optional] 
**Comments** | Pointer to **NullableString** | Comments for step | [optional] 
**WorkItemId** | Pointer to **NullableString** | Unique identifier of workitem which relates to the step | [optional] 

## Methods

### NewUpdateStepApiModel

`func NewUpdateStepApiModel(id string, ) *UpdateStepApiModel`

NewUpdateStepApiModel instantiates a new UpdateStepApiModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateStepApiModelWithDefaults

`func NewUpdateStepApiModelWithDefaults() *UpdateStepApiModel`

NewUpdateStepApiModelWithDefaults instantiates a new UpdateStepApiModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdateStepApiModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateStepApiModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateStepApiModel) SetId(v string)`

SetId sets Id field to given value.


### GetAction

`func (o *UpdateStepApiModel) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *UpdateStepApiModel) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *UpdateStepApiModel) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *UpdateStepApiModel) HasAction() bool`

HasAction returns a boolean if a field has been set.

### SetActionNil

`func (o *UpdateStepApiModel) SetActionNil(b bool)`

 SetActionNil sets the value for Action to be an explicit nil

### UnsetAction
`func (o *UpdateStepApiModel) UnsetAction()`

UnsetAction ensures that no value is present for Action, not even an explicit nil
### GetExpected

`func (o *UpdateStepApiModel) GetExpected() string`

GetExpected returns the Expected field if non-nil, zero value otherwise.

### GetExpectedOk

`func (o *UpdateStepApiModel) GetExpectedOk() (*string, bool)`

GetExpectedOk returns a tuple with the Expected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpected

`func (o *UpdateStepApiModel) SetExpected(v string)`

SetExpected sets Expected field to given value.

### HasExpected

`func (o *UpdateStepApiModel) HasExpected() bool`

HasExpected returns a boolean if a field has been set.

### SetExpectedNil

`func (o *UpdateStepApiModel) SetExpectedNil(b bool)`

 SetExpectedNil sets the value for Expected to be an explicit nil

### UnsetExpected
`func (o *UpdateStepApiModel) UnsetExpected()`

UnsetExpected ensures that no value is present for Expected, not even an explicit nil
### GetTestData

`func (o *UpdateStepApiModel) GetTestData() string`

GetTestData returns the TestData field if non-nil, zero value otherwise.

### GetTestDataOk

`func (o *UpdateStepApiModel) GetTestDataOk() (*string, bool)`

GetTestDataOk returns a tuple with the TestData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestData

`func (o *UpdateStepApiModel) SetTestData(v string)`

SetTestData sets TestData field to given value.

### HasTestData

`func (o *UpdateStepApiModel) HasTestData() bool`

HasTestData returns a boolean if a field has been set.

### SetTestDataNil

`func (o *UpdateStepApiModel) SetTestDataNil(b bool)`

 SetTestDataNil sets the value for TestData to be an explicit nil

### UnsetTestData
`func (o *UpdateStepApiModel) UnsetTestData()`

UnsetTestData ensures that no value is present for TestData, not even an explicit nil
### GetComments

`func (o *UpdateStepApiModel) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *UpdateStepApiModel) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *UpdateStepApiModel) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *UpdateStepApiModel) HasComments() bool`

HasComments returns a boolean if a field has been set.

### SetCommentsNil

`func (o *UpdateStepApiModel) SetCommentsNil(b bool)`

 SetCommentsNil sets the value for Comments to be an explicit nil

### UnsetComments
`func (o *UpdateStepApiModel) UnsetComments()`

UnsetComments ensures that no value is present for Comments, not even an explicit nil
### GetWorkItemId

`func (o *UpdateStepApiModel) GetWorkItemId() string`

GetWorkItemId returns the WorkItemId field if non-nil, zero value otherwise.

### GetWorkItemIdOk

`func (o *UpdateStepApiModel) GetWorkItemIdOk() (*string, bool)`

GetWorkItemIdOk returns a tuple with the WorkItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkItemId

`func (o *UpdateStepApiModel) SetWorkItemId(v string)`

SetWorkItemId sets WorkItemId field to given value.

### HasWorkItemId

`func (o *UpdateStepApiModel) HasWorkItemId() bool`

HasWorkItemId returns a boolean if a field has been set.

### SetWorkItemIdNil

`func (o *UpdateStepApiModel) SetWorkItemIdNil(b bool)`

 SetWorkItemIdNil sets the value for WorkItemId to be an explicit nil

### UnsetWorkItemId
`func (o *UpdateStepApiModel) UnsetWorkItemId()`

UnsetWorkItemId ensures that no value is present for WorkItemId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


