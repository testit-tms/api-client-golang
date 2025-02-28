# TestRunByAutoTestApiResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique ID of the entity | 
**IsDeleted** | **bool** | Indicates if the entity is deleted | 
**StateName** | [**TestRunState**](TestRunState.md) | Test run state | 
**Status** | [**TestStatusApiResult**](TestStatusApiResult.md) | Test run status | 
**ProjectId** | **string** | Project internal identifier | 
**TestPlanId** | Pointer to **NullableString** | Test plan internal identifier | [optional] 
**Name** | Pointer to **NullableString** | Test run name | [optional] 
**Description** | Pointer to **NullableString** | Test run description | [optional] 

## Methods

### NewTestRunByAutoTestApiResult

`func NewTestRunByAutoTestApiResult(id string, isDeleted bool, stateName TestRunState, status TestStatusApiResult, projectId string, ) *TestRunByAutoTestApiResult`

NewTestRunByAutoTestApiResult instantiates a new TestRunByAutoTestApiResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestRunByAutoTestApiResultWithDefaults

`func NewTestRunByAutoTestApiResultWithDefaults() *TestRunByAutoTestApiResult`

NewTestRunByAutoTestApiResultWithDefaults instantiates a new TestRunByAutoTestApiResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TestRunByAutoTestApiResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TestRunByAutoTestApiResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TestRunByAutoTestApiResult) SetId(v string)`

SetId sets Id field to given value.


### GetIsDeleted

`func (o *TestRunByAutoTestApiResult) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *TestRunByAutoTestApiResult) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *TestRunByAutoTestApiResult) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.


### GetStateName

`func (o *TestRunByAutoTestApiResult) GetStateName() TestRunState`

GetStateName returns the StateName field if non-nil, zero value otherwise.

### GetStateNameOk

`func (o *TestRunByAutoTestApiResult) GetStateNameOk() (*TestRunState, bool)`

GetStateNameOk returns a tuple with the StateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateName

`func (o *TestRunByAutoTestApiResult) SetStateName(v TestRunState)`

SetStateName sets StateName field to given value.


### GetStatus

`func (o *TestRunByAutoTestApiResult) GetStatus() TestStatusApiResult`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TestRunByAutoTestApiResult) GetStatusOk() (*TestStatusApiResult, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TestRunByAutoTestApiResult) SetStatus(v TestStatusApiResult)`

SetStatus sets Status field to given value.


### GetProjectId

`func (o *TestRunByAutoTestApiResult) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *TestRunByAutoTestApiResult) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *TestRunByAutoTestApiResult) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetTestPlanId

`func (o *TestRunByAutoTestApiResult) GetTestPlanId() string`

GetTestPlanId returns the TestPlanId field if non-nil, zero value otherwise.

### GetTestPlanIdOk

`func (o *TestRunByAutoTestApiResult) GetTestPlanIdOk() (*string, bool)`

GetTestPlanIdOk returns a tuple with the TestPlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestPlanId

`func (o *TestRunByAutoTestApiResult) SetTestPlanId(v string)`

SetTestPlanId sets TestPlanId field to given value.

### HasTestPlanId

`func (o *TestRunByAutoTestApiResult) HasTestPlanId() bool`

HasTestPlanId returns a boolean if a field has been set.

### SetTestPlanIdNil

`func (o *TestRunByAutoTestApiResult) SetTestPlanIdNil(b bool)`

 SetTestPlanIdNil sets the value for TestPlanId to be an explicit nil

### UnsetTestPlanId
`func (o *TestRunByAutoTestApiResult) UnsetTestPlanId()`

UnsetTestPlanId ensures that no value is present for TestPlanId, not even an explicit nil
### GetName

`func (o *TestRunByAutoTestApiResult) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TestRunByAutoTestApiResult) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TestRunByAutoTestApiResult) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TestRunByAutoTestApiResult) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *TestRunByAutoTestApiResult) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *TestRunByAutoTestApiResult) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *TestRunByAutoTestApiResult) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TestRunByAutoTestApiResult) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TestRunByAutoTestApiResult) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TestRunByAutoTestApiResult) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *TestRunByAutoTestApiResult) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *TestRunByAutoTestApiResult) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


