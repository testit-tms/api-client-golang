# TestPlanTestPointsWorkItemSearchApiResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**GlobalId** | **int64** |  | 
**VersionId** | **string** |  | 
**VersionNumber** | **int32** |  | 
**MedianDuration** | Pointer to **NullableInt64** |  | [optional] 
**IsDeleted** | Pointer to **NullableBool** |  | [optional] 
**Duration** | **int32** |  | 
**State** | [**WorkItemState**](WorkItemState.md) |  | 
**Tags** | **[]string** |  | 
**Attributes** | **map[string]interface{}** |  | 
**OrderRank** | Pointer to **NullableString** |  | [optional] 
**IsAutomated** | **bool** |  | 
**Name** | **string** |  | 
**Priority** | [**WorkItemPriority**](WorkItemPriority.md) |  | 
**Section** | [**TestPlanTestPointsSectionSearchApiResult**](TestPlanTestPointsSectionSearchApiResult.md) |  | 
**Created** | [**AuditApiResult**](AuditApiResult.md) |  | 
**Modified** | Pointer to [**NullableAuditApiResult**](AuditApiResult.md) |  | [optional] 

## Methods

### NewTestPlanTestPointsWorkItemSearchApiResult

`func NewTestPlanTestPointsWorkItemSearchApiResult(id string, globalId int64, versionId string, versionNumber int32, duration int32, state WorkItemState, tags []string, attributes map[string]interface{}, isAutomated bool, name string, priority WorkItemPriority, section TestPlanTestPointsSectionSearchApiResult, created AuditApiResult, ) *TestPlanTestPointsWorkItemSearchApiResult`

NewTestPlanTestPointsWorkItemSearchApiResult instantiates a new TestPlanTestPointsWorkItemSearchApiResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestPlanTestPointsWorkItemSearchApiResultWithDefaults

`func NewTestPlanTestPointsWorkItemSearchApiResultWithDefaults() *TestPlanTestPointsWorkItemSearchApiResult`

NewTestPlanTestPointsWorkItemSearchApiResultWithDefaults instantiates a new TestPlanTestPointsWorkItemSearchApiResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TestPlanTestPointsWorkItemSearchApiResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TestPlanTestPointsWorkItemSearchApiResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TestPlanTestPointsWorkItemSearchApiResult) SetId(v string)`

SetId sets Id field to given value.


### GetGlobalId

`func (o *TestPlanTestPointsWorkItemSearchApiResult) GetGlobalId() int64`

GetGlobalId returns the GlobalId field if non-nil, zero value otherwise.

### GetGlobalIdOk

`func (o *TestPlanTestPointsWorkItemSearchApiResult) GetGlobalIdOk() (*int64, bool)`

GetGlobalIdOk returns a tuple with the GlobalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalId

`func (o *TestPlanTestPointsWorkItemSearchApiResult) SetGlobalId(v int64)`

SetGlobalId sets GlobalId field to given value.


### GetVersionId

`func (o *TestPlanTestPointsWorkItemSearchApiResult) GetVersionId() string`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *TestPlanTestPointsWorkItemSearchApiResult) GetVersionIdOk() (*string, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *TestPlanTestPointsWorkItemSearchApiResult) SetVersionId(v string)`

SetVersionId sets VersionId field to given value.


### GetVersionNumber

`func (o *TestPlanTestPointsWorkItemSearchApiResult) GetVersionNumber() int32`

GetVersionNumber returns the VersionNumber field if non-nil, zero value otherwise.

### GetVersionNumberOk

`func (o *TestPlanTestPointsWorkItemSearchApiResult) GetVersionNumberOk() (*int32, bool)`

GetVersionNumberOk returns a tuple with the VersionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionNumber

`func (o *TestPlanTestPointsWorkItemSearchApiResult) SetVersionNumber(v int32)`

SetVersionNumber sets VersionNumber field to given value.


### GetMedianDuration

`func (o *TestPlanTestPointsWorkItemSearchApiResult) GetMedianDuration() int64`

GetMedianDuration returns the MedianDuration field if non-nil, zero value otherwise.

### GetMedianDurationOk

`func (o *TestPlanTestPointsWorkItemSearchApiResult) GetMedianDurationOk() (*int64, bool)`

GetMedianDurationOk returns a tuple with the MedianDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedianDuration

`func (o *TestPlanTestPointsWorkItemSearchApiResult) SetMedianDuration(v int64)`

SetMedianDuration sets MedianDuration field to given value.

### HasMedianDuration

`func (o *TestPlanTestPointsWorkItemSearchApiResult) HasMedianDuration() bool`

HasMedianDuration returns a boolean if a field has been set.

### SetMedianDurationNil

`func (o *TestPlanTestPointsWorkItemSearchApiResult) SetMedianDurationNil(b bool)`

 SetMedianDurationNil sets the value for MedianDuration to be an explicit nil

### UnsetMedianDuration
`func (o *TestPlanTestPointsWorkItemSearchApiResult) UnsetMedianDuration()`

UnsetMedianDuration ensures that no value is present for MedianDuration, not even an explicit nil
### GetIsDeleted

`func (o *TestPlanTestPointsWorkItemSearchApiResult) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *TestPlanTestPointsWorkItemSearchApiResult) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *TestPlanTestPointsWorkItemSearchApiResult) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *TestPlanTestPointsWorkItemSearchApiResult) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### SetIsDeletedNil

`func (o *TestPlanTestPointsWorkItemSearchApiResult) SetIsDeletedNil(b bool)`

 SetIsDeletedNil sets the value for IsDeleted to be an explicit nil

### UnsetIsDeleted
`func (o *TestPlanTestPointsWorkItemSearchApiResult) UnsetIsDeleted()`

UnsetIsDeleted ensures that no value is present for IsDeleted, not even an explicit nil
### GetDuration

`func (o *TestPlanTestPointsWorkItemSearchApiResult) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *TestPlanTestPointsWorkItemSearchApiResult) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *TestPlanTestPointsWorkItemSearchApiResult) SetDuration(v int32)`

SetDuration sets Duration field to given value.


### GetState

`func (o *TestPlanTestPointsWorkItemSearchApiResult) GetState() WorkItemState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *TestPlanTestPointsWorkItemSearchApiResult) GetStateOk() (*WorkItemState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *TestPlanTestPointsWorkItemSearchApiResult) SetState(v WorkItemState)`

SetState sets State field to given value.


### GetTags

`func (o *TestPlanTestPointsWorkItemSearchApiResult) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *TestPlanTestPointsWorkItemSearchApiResult) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *TestPlanTestPointsWorkItemSearchApiResult) SetTags(v []string)`

SetTags sets Tags field to given value.


### GetAttributes

`func (o *TestPlanTestPointsWorkItemSearchApiResult) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *TestPlanTestPointsWorkItemSearchApiResult) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *TestPlanTestPointsWorkItemSearchApiResult) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.


### GetOrderRank

`func (o *TestPlanTestPointsWorkItemSearchApiResult) GetOrderRank() string`

GetOrderRank returns the OrderRank field if non-nil, zero value otherwise.

### GetOrderRankOk

`func (o *TestPlanTestPointsWorkItemSearchApiResult) GetOrderRankOk() (*string, bool)`

GetOrderRankOk returns a tuple with the OrderRank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderRank

`func (o *TestPlanTestPointsWorkItemSearchApiResult) SetOrderRank(v string)`

SetOrderRank sets OrderRank field to given value.

### HasOrderRank

`func (o *TestPlanTestPointsWorkItemSearchApiResult) HasOrderRank() bool`

HasOrderRank returns a boolean if a field has been set.

### SetOrderRankNil

`func (o *TestPlanTestPointsWorkItemSearchApiResult) SetOrderRankNil(b bool)`

 SetOrderRankNil sets the value for OrderRank to be an explicit nil

### UnsetOrderRank
`func (o *TestPlanTestPointsWorkItemSearchApiResult) UnsetOrderRank()`

UnsetOrderRank ensures that no value is present for OrderRank, not even an explicit nil
### GetIsAutomated

`func (o *TestPlanTestPointsWorkItemSearchApiResult) GetIsAutomated() bool`

GetIsAutomated returns the IsAutomated field if non-nil, zero value otherwise.

### GetIsAutomatedOk

`func (o *TestPlanTestPointsWorkItemSearchApiResult) GetIsAutomatedOk() (*bool, bool)`

GetIsAutomatedOk returns a tuple with the IsAutomated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutomated

`func (o *TestPlanTestPointsWorkItemSearchApiResult) SetIsAutomated(v bool)`

SetIsAutomated sets IsAutomated field to given value.


### GetName

`func (o *TestPlanTestPointsWorkItemSearchApiResult) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TestPlanTestPointsWorkItemSearchApiResult) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TestPlanTestPointsWorkItemSearchApiResult) SetName(v string)`

SetName sets Name field to given value.


### GetPriority

`func (o *TestPlanTestPointsWorkItemSearchApiResult) GetPriority() WorkItemPriority`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *TestPlanTestPointsWorkItemSearchApiResult) GetPriorityOk() (*WorkItemPriority, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *TestPlanTestPointsWorkItemSearchApiResult) SetPriority(v WorkItemPriority)`

SetPriority sets Priority field to given value.


### GetSection

`func (o *TestPlanTestPointsWorkItemSearchApiResult) GetSection() TestPlanTestPointsSectionSearchApiResult`

GetSection returns the Section field if non-nil, zero value otherwise.

### GetSectionOk

`func (o *TestPlanTestPointsWorkItemSearchApiResult) GetSectionOk() (*TestPlanTestPointsSectionSearchApiResult, bool)`

GetSectionOk returns a tuple with the Section field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSection

`func (o *TestPlanTestPointsWorkItemSearchApiResult) SetSection(v TestPlanTestPointsSectionSearchApiResult)`

SetSection sets Section field to given value.


### GetCreated

`func (o *TestPlanTestPointsWorkItemSearchApiResult) GetCreated() AuditApiResult`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *TestPlanTestPointsWorkItemSearchApiResult) GetCreatedOk() (*AuditApiResult, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *TestPlanTestPointsWorkItemSearchApiResult) SetCreated(v AuditApiResult)`

SetCreated sets Created field to given value.


### GetModified

`func (o *TestPlanTestPointsWorkItemSearchApiResult) GetModified() AuditApiResult`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *TestPlanTestPointsWorkItemSearchApiResult) GetModifiedOk() (*AuditApiResult, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *TestPlanTestPointsWorkItemSearchApiResult) SetModified(v AuditApiResult)`

SetModified sets Modified field to given value.

### HasModified

`func (o *TestPlanTestPointsWorkItemSearchApiResult) HasModified() bool`

HasModified returns a boolean if a field has been set.

### SetModifiedNil

`func (o *TestPlanTestPointsWorkItemSearchApiResult) SetModifiedNil(b bool)`

 SetModifiedNil sets the value for Modified to be an explicit nil

### UnsetModified
`func (o *TestPlanTestPointsWorkItemSearchApiResult) UnsetModified()`

UnsetModified ensures that no value is present for Modified, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


