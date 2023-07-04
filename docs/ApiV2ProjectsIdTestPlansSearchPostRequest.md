# ApiV2ProjectsIdTestPlansSearchPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Build** | Pointer to **NullableString** |  | [optional] 
**ProductName** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to [**[]TestPlanStatusModel**](TestPlanStatusModel.md) |  | [optional] 
**GlobalIds** | Pointer to **[]int64** |  | [optional] 
**IsLocked** | Pointer to **NullableBool** |  | [optional] 
**LockedDate** | Pointer to [**NullableDateTimeRangeSelectorModel**](DateTimeRangeSelectorModel.md) |  | [optional] 
**AutomaticDurationTimer** | Pointer to **[]bool** |  | [optional] 
**CreatedByIds** | Pointer to **[]string** |  | [optional] 
**CreatedDate** | Pointer to [**NullableDateTimeRangeSelectorModel**](DateTimeRangeSelectorModel.md) |  | [optional] 
**StartDate** | Pointer to [**NullableDateTimeRangeSelectorModel**](DateTimeRangeSelectorModel.md) |  | [optional] 
**EndDate** | Pointer to [**NullableDateTimeRangeSelectorModel**](DateTimeRangeSelectorModel.md) |  | [optional] 
**TagNames** | Pointer to **[]string** |  | [optional] 
**Attributes** | Pointer to **map[string][]string** |  | [optional] 
**IsDeleted** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewApiV2ProjectsIdTestPlansSearchPostRequest

`func NewApiV2ProjectsIdTestPlansSearchPostRequest() *ApiV2ProjectsIdTestPlansSearchPostRequest`

NewApiV2ProjectsIdTestPlansSearchPostRequest instantiates a new ApiV2ProjectsIdTestPlansSearchPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV2ProjectsIdTestPlansSearchPostRequestWithDefaults

`func NewApiV2ProjectsIdTestPlansSearchPostRequestWithDefaults() *ApiV2ProjectsIdTestPlansSearchPostRequest`

NewApiV2ProjectsIdTestPlansSearchPostRequestWithDefaults instantiates a new ApiV2ProjectsIdTestPlansSearchPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetBuild

`func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) GetBuild() string`

GetBuild returns the Build field if non-nil, zero value otherwise.

### GetBuildOk

`func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) GetBuildOk() (*string, bool)`

GetBuildOk returns a tuple with the Build field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuild

`func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) SetBuild(v string)`

SetBuild sets Build field to given value.

### HasBuild

`func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) HasBuild() bool`

HasBuild returns a boolean if a field has been set.

### SetBuildNil

`func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) SetBuildNil(b bool)`

 SetBuildNil sets the value for Build to be an explicit nil

### UnsetBuild
`func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) UnsetBuild()`

UnsetBuild ensures that no value is present for Build, not even an explicit nil
### GetProductName

`func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) GetProductName() string`

GetProductName returns the ProductName field if non-nil, zero value otherwise.

### GetProductNameOk

`func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) GetProductNameOk() (*string, bool)`

GetProductNameOk returns a tuple with the ProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductName

`func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) SetProductName(v string)`

SetProductName sets ProductName field to given value.

### HasProductName

`func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) HasProductName() bool`

HasProductName returns a boolean if a field has been set.

### SetProductNameNil

`func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) SetProductNameNil(b bool)`

 SetProductNameNil sets the value for ProductName to be an explicit nil

### UnsetProductName
`func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) UnsetProductName()`

UnsetProductName ensures that no value is present for ProductName, not even an explicit nil
### GetStatus

`func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) GetStatus() []TestPlanStatusModel`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) GetStatusOk() (*[]TestPlanStatusModel, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) SetStatus(v []TestPlanStatusModel)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetGlobalIds

`func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) GetGlobalIds() []int64`

GetGlobalIds returns the GlobalIds field if non-nil, zero value otherwise.

### GetGlobalIdsOk

`func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) GetGlobalIdsOk() (*[]int64, bool)`

GetGlobalIdsOk returns a tuple with the GlobalIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalIds

`func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) SetGlobalIds(v []int64)`

SetGlobalIds sets GlobalIds field to given value.

### HasGlobalIds

`func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) HasGlobalIds() bool`

HasGlobalIds returns a boolean if a field has been set.

### SetGlobalIdsNil

`func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) SetGlobalIdsNil(b bool)`

 SetGlobalIdsNil sets the value for GlobalIds to be an explicit nil

### UnsetGlobalIds
`func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) UnsetGlobalIds()`

UnsetGlobalIds ensures that no value is present for GlobalIds, not even an explicit nil
### GetIsLocked

`func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) GetIsLocked() bool`

GetIsLocked returns the IsLocked field if non-nil, zero value otherwise.

### GetIsLockedOk

`func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) GetIsLockedOk() (*bool, bool)`

GetIsLockedOk returns a tuple with the IsLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocked

`func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) SetIsLocked(v bool)`

SetIsLocked sets IsLocked field to given value.

### HasIsLocked

`func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) HasIsLocked() bool`

HasIsLocked returns a boolean if a field has been set.

### SetIsLockedNil

`func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) SetIsLockedNil(b bool)`

 SetIsLockedNil sets the value for IsLocked to be an explicit nil

### UnsetIsLocked
`func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) UnsetIsLocked()`

UnsetIsLocked ensures that no value is present for IsLocked, not even an explicit nil
### GetLockedDate

`func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) GetLockedDate() DateTimeRangeSelectorModel`

GetLockedDate returns the LockedDate field if non-nil, zero value otherwise.

### GetLockedDateOk

`func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) GetLockedDateOk() (*DateTimeRangeSelectorModel, bool)`

GetLockedDateOk returns a tuple with the LockedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockedDate

`func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) SetLockedDate(v DateTimeRangeSelectorModel)`

SetLockedDate sets LockedDate field to given value.

### HasLockedDate

`func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) HasLockedDate() bool`

HasLockedDate returns a boolean if a field has been set.

### SetLockedDateNil

`func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) SetLockedDateNil(b bool)`

 SetLockedDateNil sets the value for LockedDate to be an explicit nil

### UnsetLockedDate
`func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) UnsetLockedDate()`

UnsetLockedDate ensures that no value is present for LockedDate, not even an explicit nil
### GetAutomaticDurationTimer

`func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) GetAutomaticDurationTimer() []bool`

GetAutomaticDurationTimer returns the AutomaticDurationTimer field if non-nil, zero value otherwise.

### GetAutomaticDurationTimerOk

`func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) GetAutomaticDurationTimerOk() (*[]bool, bool)`

GetAutomaticDurationTimerOk returns a tuple with the AutomaticDurationTimer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomaticDurationTimer

`func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) SetAutomaticDurationTimer(v []bool)`

SetAutomaticDurationTimer sets AutomaticDurationTimer field to given value.

### HasAutomaticDurationTimer

`func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) HasAutomaticDurationTimer() bool`

HasAutomaticDurationTimer returns a boolean if a field has been set.

### SetAutomaticDurationTimerNil

`func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) SetAutomaticDurationTimerNil(b bool)`

 SetAutomaticDurationTimerNil sets the value for AutomaticDurationTimer to be an explicit nil

### UnsetAutomaticDurationTimer
`func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) UnsetAutomaticDurationTimer()`

UnsetAutomaticDurationTimer ensures that no value is present for AutomaticDurationTimer, not even an explicit nil
### GetCreatedByIds

`func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) GetCreatedByIds() []string`

GetCreatedByIds returns the CreatedByIds field if non-nil, zero value otherwise.

### GetCreatedByIdsOk

`func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) GetCreatedByIdsOk() (*[]string, bool)`

GetCreatedByIdsOk returns a tuple with the CreatedByIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByIds

`func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) SetCreatedByIds(v []string)`

SetCreatedByIds sets CreatedByIds field to given value.

### HasCreatedByIds

`func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) HasCreatedByIds() bool`

HasCreatedByIds returns a boolean if a field has been set.

### SetCreatedByIdsNil

`func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) SetCreatedByIdsNil(b bool)`

 SetCreatedByIdsNil sets the value for CreatedByIds to be an explicit nil

### UnsetCreatedByIds
`func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) UnsetCreatedByIds()`

UnsetCreatedByIds ensures that no value is present for CreatedByIds, not even an explicit nil
### GetCreatedDate

`func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) GetCreatedDate() DateTimeRangeSelectorModel`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) GetCreatedDateOk() (*DateTimeRangeSelectorModel, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) SetCreatedDate(v DateTimeRangeSelectorModel)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### SetCreatedDateNil

`func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) SetCreatedDateNil(b bool)`

 SetCreatedDateNil sets the value for CreatedDate to be an explicit nil

### UnsetCreatedDate
`func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) UnsetCreatedDate()`

UnsetCreatedDate ensures that no value is present for CreatedDate, not even an explicit nil
### GetStartDate

`func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) GetStartDate() DateTimeRangeSelectorModel`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) GetStartDateOk() (*DateTimeRangeSelectorModel, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) SetStartDate(v DateTimeRangeSelectorModel)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### SetStartDateNil

`func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) SetStartDateNil(b bool)`

 SetStartDateNil sets the value for StartDate to be an explicit nil

### UnsetStartDate
`func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) UnsetStartDate()`

UnsetStartDate ensures that no value is present for StartDate, not even an explicit nil
### GetEndDate

`func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) GetEndDate() DateTimeRangeSelectorModel`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) GetEndDateOk() (*DateTimeRangeSelectorModel, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) SetEndDate(v DateTimeRangeSelectorModel)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### SetEndDateNil

`func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) SetEndDateNil(b bool)`

 SetEndDateNil sets the value for EndDate to be an explicit nil

### UnsetEndDate
`func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) UnsetEndDate()`

UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
### GetTagNames

`func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) GetTagNames() []string`

GetTagNames returns the TagNames field if non-nil, zero value otherwise.

### GetTagNamesOk

`func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) GetTagNamesOk() (*[]string, bool)`

GetTagNamesOk returns a tuple with the TagNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagNames

`func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) SetTagNames(v []string)`

SetTagNames sets TagNames field to given value.

### HasTagNames

`func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) HasTagNames() bool`

HasTagNames returns a boolean if a field has been set.

### SetTagNamesNil

`func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) SetTagNamesNil(b bool)`

 SetTagNamesNil sets the value for TagNames to be an explicit nil

### UnsetTagNames
`func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) UnsetTagNames()`

UnsetTagNames ensures that no value is present for TagNames, not even an explicit nil
### GetAttributes

`func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) GetAttributes() map[string][]string`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) GetAttributesOk() (*map[string][]string, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) SetAttributes(v map[string][]string)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### SetAttributesNil

`func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) SetAttributesNil(b bool)`

 SetAttributesNil sets the value for Attributes to be an explicit nil

### UnsetAttributes
`func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) UnsetAttributes()`

UnsetAttributes ensures that no value is present for Attributes, not even an explicit nil
### GetIsDeleted

`func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### SetIsDeletedNil

`func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) SetIsDeletedNil(b bool)`

 SetIsDeletedNil sets the value for IsDeleted to be an explicit nil

### UnsetIsDeleted
`func (o *ApiV2ProjectsIdTestPlansSearchPostRequest) UnsetIsDeleted()`

UnsetIsDeleted ensures that no value is present for IsDeleted, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


