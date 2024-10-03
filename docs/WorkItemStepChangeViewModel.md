# WorkItemStepChangeViewModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** |  | 
**Expected** | **string** |  | 
**Comments** | **string** |  | 
**TestData** | **string** |  | 
**Index** | **int32** |  | 
**WorkItemId** | Pointer to **NullableString** |  | [optional] 
**WorkItem** | [**WorkItemStepChangeViewModelWorkItem**](WorkItemStepChangeViewModelWorkItem.md) |  | 

## Methods

### NewWorkItemStepChangeViewModel

`func NewWorkItemStepChangeViewModel(action string, expected string, comments string, testData string, index int32, workItem WorkItemStepChangeViewModelWorkItem, ) *WorkItemStepChangeViewModel`

NewWorkItemStepChangeViewModel instantiates a new WorkItemStepChangeViewModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkItemStepChangeViewModelWithDefaults

`func NewWorkItemStepChangeViewModelWithDefaults() *WorkItemStepChangeViewModel`

NewWorkItemStepChangeViewModelWithDefaults instantiates a new WorkItemStepChangeViewModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *WorkItemStepChangeViewModel) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *WorkItemStepChangeViewModel) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *WorkItemStepChangeViewModel) SetAction(v string)`

SetAction sets Action field to given value.


### GetExpected

`func (o *WorkItemStepChangeViewModel) GetExpected() string`

GetExpected returns the Expected field if non-nil, zero value otherwise.

### GetExpectedOk

`func (o *WorkItemStepChangeViewModel) GetExpectedOk() (*string, bool)`

GetExpectedOk returns a tuple with the Expected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpected

`func (o *WorkItemStepChangeViewModel) SetExpected(v string)`

SetExpected sets Expected field to given value.


### GetComments

`func (o *WorkItemStepChangeViewModel) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *WorkItemStepChangeViewModel) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *WorkItemStepChangeViewModel) SetComments(v string)`

SetComments sets Comments field to given value.


### GetTestData

`func (o *WorkItemStepChangeViewModel) GetTestData() string`

GetTestData returns the TestData field if non-nil, zero value otherwise.

### GetTestDataOk

`func (o *WorkItemStepChangeViewModel) GetTestDataOk() (*string, bool)`

GetTestDataOk returns a tuple with the TestData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestData

`func (o *WorkItemStepChangeViewModel) SetTestData(v string)`

SetTestData sets TestData field to given value.


### GetIndex

`func (o *WorkItemStepChangeViewModel) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *WorkItemStepChangeViewModel) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *WorkItemStepChangeViewModel) SetIndex(v int32)`

SetIndex sets Index field to given value.


### GetWorkItemId

`func (o *WorkItemStepChangeViewModel) GetWorkItemId() string`

GetWorkItemId returns the WorkItemId field if non-nil, zero value otherwise.

### GetWorkItemIdOk

`func (o *WorkItemStepChangeViewModel) GetWorkItemIdOk() (*string, bool)`

GetWorkItemIdOk returns a tuple with the WorkItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkItemId

`func (o *WorkItemStepChangeViewModel) SetWorkItemId(v string)`

SetWorkItemId sets WorkItemId field to given value.

### HasWorkItemId

`func (o *WorkItemStepChangeViewModel) HasWorkItemId() bool`

HasWorkItemId returns a boolean if a field has been set.

### SetWorkItemIdNil

`func (o *WorkItemStepChangeViewModel) SetWorkItemIdNil(b bool)`

 SetWorkItemIdNil sets the value for WorkItemId to be an explicit nil

### UnsetWorkItemId
`func (o *WorkItemStepChangeViewModel) UnsetWorkItemId()`

UnsetWorkItemId ensures that no value is present for WorkItemId, not even an explicit nil
### GetWorkItem

`func (o *WorkItemStepChangeViewModel) GetWorkItem() WorkItemStepChangeViewModelWorkItem`

GetWorkItem returns the WorkItem field if non-nil, zero value otherwise.

### GetWorkItemOk

`func (o *WorkItemStepChangeViewModel) GetWorkItemOk() (*WorkItemStepChangeViewModelWorkItem, bool)`

GetWorkItemOk returns a tuple with the WorkItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkItem

`func (o *WorkItemStepChangeViewModel) SetWorkItem(v WorkItemStepChangeViewModelWorkItem)`

SetWorkItem sets WorkItem field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


