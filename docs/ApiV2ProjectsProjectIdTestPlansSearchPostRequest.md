# ApiV2ProjectsProjectIdTestPlansSearchPostRequest

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

### NewApiV2ProjectsProjectIdTestPlansSearchPostRequest

`func NewApiV2ProjectsProjectIdTestPlansSearchPostRequest() *ApiV2ProjectsProjectIdTestPlansSearchPostRequest`

NewApiV2ProjectsProjectIdTestPlansSearchPostRequest instantiates a new ApiV2ProjectsProjectIdTestPlansSearchPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV2ProjectsProjectIdTestPlansSearchPostRequestWithDefaults

`func NewApiV2ProjectsProjectIdTestPlansSearchPostRequestWithDefaults() *ApiV2ProjectsProjectIdTestPlansSearchPostRequest`

NewApiV2ProjectsProjectIdTestPlansSearchPostRequestWithDefaults instantiates a new ApiV2ProjectsProjectIdTestPlansSearchPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetBuild

`func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) GetBuild() string`

GetBuild returns the Build field if non-nil, zero value otherwise.

### GetBuildOk

`func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) GetBuildOk() (*string, bool)`

GetBuildOk returns a tuple with the Build field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuild

`func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) SetBuild(v string)`

SetBuild sets Build field to given value.

### HasBuild

`func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) HasBuild() bool`

HasBuild returns a boolean if a field has been set.

### SetBuildNil

`func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) SetBuildNil(b bool)`

 SetBuildNil sets the value for Build to be an explicit nil

### UnsetBuild
`func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) UnsetBuild()`

UnsetBuild ensures that no value is present for Build, not even an explicit nil
### GetProductName

`func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) GetProductName() string`

GetProductName returns the ProductName field if non-nil, zero value otherwise.

### GetProductNameOk

`func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) GetProductNameOk() (*string, bool)`

GetProductNameOk returns a tuple with the ProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductName

`func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) SetProductName(v string)`

SetProductName sets ProductName field to given value.

### HasProductName

`func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) HasProductName() bool`

HasProductName returns a boolean if a field has been set.

### SetProductNameNil

`func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) SetProductNameNil(b bool)`

 SetProductNameNil sets the value for ProductName to be an explicit nil

### UnsetProductName
`func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) UnsetProductName()`

UnsetProductName ensures that no value is present for ProductName, not even an explicit nil
### GetStatus

`func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) GetStatus() []TestPlanStatusModel`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) GetStatusOk() (*[]TestPlanStatusModel, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) SetStatus(v []TestPlanStatusModel)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetGlobalIds

`func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) GetGlobalIds() []int64`

GetGlobalIds returns the GlobalIds field if non-nil, zero value otherwise.

### GetGlobalIdsOk

`func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) GetGlobalIdsOk() (*[]int64, bool)`

GetGlobalIdsOk returns a tuple with the GlobalIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalIds

`func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) SetGlobalIds(v []int64)`

SetGlobalIds sets GlobalIds field to given value.

### HasGlobalIds

`func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) HasGlobalIds() bool`

HasGlobalIds returns a boolean if a field has been set.

### SetGlobalIdsNil

`func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) SetGlobalIdsNil(b bool)`

 SetGlobalIdsNil sets the value for GlobalIds to be an explicit nil

### UnsetGlobalIds
`func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) UnsetGlobalIds()`

UnsetGlobalIds ensures that no value is present for GlobalIds, not even an explicit nil
### GetIsLocked

`func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) GetIsLocked() bool`

GetIsLocked returns the IsLocked field if non-nil, zero value otherwise.

### GetIsLockedOk

`func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) GetIsLockedOk() (*bool, bool)`

GetIsLockedOk returns a tuple with the IsLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocked

`func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) SetIsLocked(v bool)`

SetIsLocked sets IsLocked field to given value.

### HasIsLocked

`func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) HasIsLocked() bool`

HasIsLocked returns a boolean if a field has been set.

### SetIsLockedNil

`func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) SetIsLockedNil(b bool)`

 SetIsLockedNil sets the value for IsLocked to be an explicit nil

### UnsetIsLocked
`func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) UnsetIsLocked()`

UnsetIsLocked ensures that no value is present for IsLocked, not even an explicit nil
### GetLockedDate

`func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) GetLockedDate() DateTimeRangeSelectorModel`

GetLockedDate returns the LockedDate field if non-nil, zero value otherwise.

### GetLockedDateOk

`func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) GetLockedDateOk() (*DateTimeRangeSelectorModel, bool)`

GetLockedDateOk returns a tuple with the LockedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockedDate

`func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) SetLockedDate(v DateTimeRangeSelectorModel)`

SetLockedDate sets LockedDate field to given value.

### HasLockedDate

`func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) HasLockedDate() bool`

HasLockedDate returns a boolean if a field has been set.

### SetLockedDateNil

`func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) SetLockedDateNil(b bool)`

 SetLockedDateNil sets the value for LockedDate to be an explicit nil

### UnsetLockedDate
`func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) UnsetLockedDate()`

UnsetLockedDate ensures that no value is present for LockedDate, not even an explicit nil
### GetAutomaticDurationTimer

`func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) GetAutomaticDurationTimer() []bool`

GetAutomaticDurationTimer returns the AutomaticDurationTimer field if non-nil, zero value otherwise.

### GetAutomaticDurationTimerOk

`func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) GetAutomaticDurationTimerOk() (*[]bool, bool)`

GetAutomaticDurationTimerOk returns a tuple with the AutomaticDurationTimer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomaticDurationTimer

`func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) SetAutomaticDurationTimer(v []bool)`

SetAutomaticDurationTimer sets AutomaticDurationTimer field to given value.

### HasAutomaticDurationTimer

`func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) HasAutomaticDurationTimer() bool`

HasAutomaticDurationTimer returns a boolean if a field has been set.

### SetAutomaticDurationTimerNil

`func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) SetAutomaticDurationTimerNil(b bool)`

 SetAutomaticDurationTimerNil sets the value for AutomaticDurationTimer to be an explicit nil

### UnsetAutomaticDurationTimer
`func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) UnsetAutomaticDurationTimer()`

UnsetAutomaticDurationTimer ensures that no value is present for AutomaticDurationTimer, not even an explicit nil
### GetCreatedByIds

`func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) GetCreatedByIds() []string`

GetCreatedByIds returns the CreatedByIds field if non-nil, zero value otherwise.

### GetCreatedByIdsOk

`func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) GetCreatedByIdsOk() (*[]string, bool)`

GetCreatedByIdsOk returns a tuple with the CreatedByIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByIds

`func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) SetCreatedByIds(v []string)`

SetCreatedByIds sets CreatedByIds field to given value.

### HasCreatedByIds

`func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) HasCreatedByIds() bool`

HasCreatedByIds returns a boolean if a field has been set.

### SetCreatedByIdsNil

`func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) SetCreatedByIdsNil(b bool)`

 SetCreatedByIdsNil sets the value for CreatedByIds to be an explicit nil

### UnsetCreatedByIds
`func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) UnsetCreatedByIds()`

UnsetCreatedByIds ensures that no value is present for CreatedByIds, not even an explicit nil
### GetCreatedDate

`func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) GetCreatedDate() DateTimeRangeSelectorModel`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) GetCreatedDateOk() (*DateTimeRangeSelectorModel, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) SetCreatedDate(v DateTimeRangeSelectorModel)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### SetCreatedDateNil

`func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) SetCreatedDateNil(b bool)`

 SetCreatedDateNil sets the value for CreatedDate to be an explicit nil

### UnsetCreatedDate
`func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) UnsetCreatedDate()`

UnsetCreatedDate ensures that no value is present for CreatedDate, not even an explicit nil
### GetStartDate

`func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) GetStartDate() DateTimeRangeSelectorModel`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) GetStartDateOk() (*DateTimeRangeSelectorModel, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) SetStartDate(v DateTimeRangeSelectorModel)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### SetStartDateNil

`func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) SetStartDateNil(b bool)`

 SetStartDateNil sets the value for StartDate to be an explicit nil

### UnsetStartDate
`func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) UnsetStartDate()`

UnsetStartDate ensures that no value is present for StartDate, not even an explicit nil
### GetEndDate

`func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) GetEndDate() DateTimeRangeSelectorModel`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) GetEndDateOk() (*DateTimeRangeSelectorModel, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) SetEndDate(v DateTimeRangeSelectorModel)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### SetEndDateNil

`func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) SetEndDateNil(b bool)`

 SetEndDateNil sets the value for EndDate to be an explicit nil

### UnsetEndDate
`func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) UnsetEndDate()`

UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
### GetTagNames

`func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) GetTagNames() []string`

GetTagNames returns the TagNames field if non-nil, zero value otherwise.

### GetTagNamesOk

`func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) GetTagNamesOk() (*[]string, bool)`

GetTagNamesOk returns a tuple with the TagNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagNames

`func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) SetTagNames(v []string)`

SetTagNames sets TagNames field to given value.

### HasTagNames

`func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) HasTagNames() bool`

HasTagNames returns a boolean if a field has been set.

### SetTagNamesNil

`func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) SetTagNamesNil(b bool)`

 SetTagNamesNil sets the value for TagNames to be an explicit nil

### UnsetTagNames
`func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) UnsetTagNames()`

UnsetTagNames ensures that no value is present for TagNames, not even an explicit nil
### GetAttributes

`func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) GetAttributes() map[string][]string`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) GetAttributesOk() (*map[string][]string, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) SetAttributes(v map[string][]string)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### SetAttributesNil

`func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) SetAttributesNil(b bool)`

 SetAttributesNil sets the value for Attributes to be an explicit nil

### UnsetAttributes
`func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) UnsetAttributes()`

UnsetAttributes ensures that no value is present for Attributes, not even an explicit nil
### GetIsDeleted

`func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### SetIsDeletedNil

`func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) SetIsDeletedNil(b bool)`

 SetIsDeletedNil sets the value for IsDeleted to be an explicit nil

### UnsetIsDeleted
`func (o *ApiV2ProjectsProjectIdTestPlansSearchPostRequest) UnsetIsDeleted()`

UnsetIsDeleted ensures that no value is present for IsDeleted, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


