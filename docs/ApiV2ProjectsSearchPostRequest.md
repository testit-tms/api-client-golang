# ApiV2ProjectsSearchPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | Specifies a project name to search for | [optional] 
**IsFavorite** | Pointer to **NullableBool** | Specifies a project favorite status to search for | [optional] 
**IsDeleted** | Pointer to **NullableBool** | Specifies a project deleted status to search for | [optional] 
**TestCasesCount** | Pointer to [**NullableProjectsFilterModelTestCasesCount**](ProjectsFilterModelTestCasesCount.md) |  | [optional] 
**ChecklistsCount** | Pointer to [**NullableProjectsFilterModelChecklistsCount**](ProjectsFilterModelChecklistsCount.md) |  | [optional] 
**SharedStepsCount** | Pointer to [**NullableProjectsFilterModelSharedStepsCount**](ProjectsFilterModelSharedStepsCount.md) |  | [optional] 
**AutotestsCount** | Pointer to [**NullableProjectsFilterModelAutotestsCount**](ProjectsFilterModelAutotestsCount.md) |  | [optional] 
**GlobalIds** | Pointer to **[]int64** | Specifies a project global IDs to search for | [optional] 
**CreatedDate** | Pointer to [**NullableProjectsFilterModelCreatedDate**](ProjectsFilterModelCreatedDate.md) |  | [optional] 
**CreatedByIds** | Pointer to **[]string** | Specifies an autotest creator IDs to search for | [optional] 
**Types** | Pointer to [**[]ProjectTypeModel**](ProjectTypeModel.md) | Collection of project types to search for | [optional] 

## Methods

### NewApiV2ProjectsSearchPostRequest

`func NewApiV2ProjectsSearchPostRequest() *ApiV2ProjectsSearchPostRequest`

NewApiV2ProjectsSearchPostRequest instantiates a new ApiV2ProjectsSearchPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV2ProjectsSearchPostRequestWithDefaults

`func NewApiV2ProjectsSearchPostRequestWithDefaults() *ApiV2ProjectsSearchPostRequest`

NewApiV2ProjectsSearchPostRequestWithDefaults instantiates a new ApiV2ProjectsSearchPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ApiV2ProjectsSearchPostRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiV2ProjectsSearchPostRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiV2ProjectsSearchPostRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApiV2ProjectsSearchPostRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ApiV2ProjectsSearchPostRequest) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ApiV2ProjectsSearchPostRequest) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetIsFavorite

`func (o *ApiV2ProjectsSearchPostRequest) GetIsFavorite() bool`

GetIsFavorite returns the IsFavorite field if non-nil, zero value otherwise.

### GetIsFavoriteOk

`func (o *ApiV2ProjectsSearchPostRequest) GetIsFavoriteOk() (*bool, bool)`

GetIsFavoriteOk returns a tuple with the IsFavorite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFavorite

`func (o *ApiV2ProjectsSearchPostRequest) SetIsFavorite(v bool)`

SetIsFavorite sets IsFavorite field to given value.

### HasIsFavorite

`func (o *ApiV2ProjectsSearchPostRequest) HasIsFavorite() bool`

HasIsFavorite returns a boolean if a field has been set.

### SetIsFavoriteNil

`func (o *ApiV2ProjectsSearchPostRequest) SetIsFavoriteNil(b bool)`

 SetIsFavoriteNil sets the value for IsFavorite to be an explicit nil

### UnsetIsFavorite
`func (o *ApiV2ProjectsSearchPostRequest) UnsetIsFavorite()`

UnsetIsFavorite ensures that no value is present for IsFavorite, not even an explicit nil
### GetIsDeleted

`func (o *ApiV2ProjectsSearchPostRequest) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *ApiV2ProjectsSearchPostRequest) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *ApiV2ProjectsSearchPostRequest) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *ApiV2ProjectsSearchPostRequest) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### SetIsDeletedNil

`func (o *ApiV2ProjectsSearchPostRequest) SetIsDeletedNil(b bool)`

 SetIsDeletedNil sets the value for IsDeleted to be an explicit nil

### UnsetIsDeleted
`func (o *ApiV2ProjectsSearchPostRequest) UnsetIsDeleted()`

UnsetIsDeleted ensures that no value is present for IsDeleted, not even an explicit nil
### GetTestCasesCount

`func (o *ApiV2ProjectsSearchPostRequest) GetTestCasesCount() ProjectsFilterModelTestCasesCount`

GetTestCasesCount returns the TestCasesCount field if non-nil, zero value otherwise.

### GetTestCasesCountOk

`func (o *ApiV2ProjectsSearchPostRequest) GetTestCasesCountOk() (*ProjectsFilterModelTestCasesCount, bool)`

GetTestCasesCountOk returns a tuple with the TestCasesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestCasesCount

`func (o *ApiV2ProjectsSearchPostRequest) SetTestCasesCount(v ProjectsFilterModelTestCasesCount)`

SetTestCasesCount sets TestCasesCount field to given value.

### HasTestCasesCount

`func (o *ApiV2ProjectsSearchPostRequest) HasTestCasesCount() bool`

HasTestCasesCount returns a boolean if a field has been set.

### SetTestCasesCountNil

`func (o *ApiV2ProjectsSearchPostRequest) SetTestCasesCountNil(b bool)`

 SetTestCasesCountNil sets the value for TestCasesCount to be an explicit nil

### UnsetTestCasesCount
`func (o *ApiV2ProjectsSearchPostRequest) UnsetTestCasesCount()`

UnsetTestCasesCount ensures that no value is present for TestCasesCount, not even an explicit nil
### GetChecklistsCount

`func (o *ApiV2ProjectsSearchPostRequest) GetChecklistsCount() ProjectsFilterModelChecklistsCount`

GetChecklistsCount returns the ChecklistsCount field if non-nil, zero value otherwise.

### GetChecklistsCountOk

`func (o *ApiV2ProjectsSearchPostRequest) GetChecklistsCountOk() (*ProjectsFilterModelChecklistsCount, bool)`

GetChecklistsCountOk returns a tuple with the ChecklistsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecklistsCount

`func (o *ApiV2ProjectsSearchPostRequest) SetChecklistsCount(v ProjectsFilterModelChecklistsCount)`

SetChecklistsCount sets ChecklistsCount field to given value.

### HasChecklistsCount

`func (o *ApiV2ProjectsSearchPostRequest) HasChecklistsCount() bool`

HasChecklistsCount returns a boolean if a field has been set.

### SetChecklistsCountNil

`func (o *ApiV2ProjectsSearchPostRequest) SetChecklistsCountNil(b bool)`

 SetChecklistsCountNil sets the value for ChecklistsCount to be an explicit nil

### UnsetChecklistsCount
`func (o *ApiV2ProjectsSearchPostRequest) UnsetChecklistsCount()`

UnsetChecklistsCount ensures that no value is present for ChecklistsCount, not even an explicit nil
### GetSharedStepsCount

`func (o *ApiV2ProjectsSearchPostRequest) GetSharedStepsCount() ProjectsFilterModelSharedStepsCount`

GetSharedStepsCount returns the SharedStepsCount field if non-nil, zero value otherwise.

### GetSharedStepsCountOk

`func (o *ApiV2ProjectsSearchPostRequest) GetSharedStepsCountOk() (*ProjectsFilterModelSharedStepsCount, bool)`

GetSharedStepsCountOk returns a tuple with the SharedStepsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedStepsCount

`func (o *ApiV2ProjectsSearchPostRequest) SetSharedStepsCount(v ProjectsFilterModelSharedStepsCount)`

SetSharedStepsCount sets SharedStepsCount field to given value.

### HasSharedStepsCount

`func (o *ApiV2ProjectsSearchPostRequest) HasSharedStepsCount() bool`

HasSharedStepsCount returns a boolean if a field has been set.

### SetSharedStepsCountNil

`func (o *ApiV2ProjectsSearchPostRequest) SetSharedStepsCountNil(b bool)`

 SetSharedStepsCountNil sets the value for SharedStepsCount to be an explicit nil

### UnsetSharedStepsCount
`func (o *ApiV2ProjectsSearchPostRequest) UnsetSharedStepsCount()`

UnsetSharedStepsCount ensures that no value is present for SharedStepsCount, not even an explicit nil
### GetAutotestsCount

`func (o *ApiV2ProjectsSearchPostRequest) GetAutotestsCount() ProjectsFilterModelAutotestsCount`

GetAutotestsCount returns the AutotestsCount field if non-nil, zero value otherwise.

### GetAutotestsCountOk

`func (o *ApiV2ProjectsSearchPostRequest) GetAutotestsCountOk() (*ProjectsFilterModelAutotestsCount, bool)`

GetAutotestsCountOk returns a tuple with the AutotestsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutotestsCount

`func (o *ApiV2ProjectsSearchPostRequest) SetAutotestsCount(v ProjectsFilterModelAutotestsCount)`

SetAutotestsCount sets AutotestsCount field to given value.

### HasAutotestsCount

`func (o *ApiV2ProjectsSearchPostRequest) HasAutotestsCount() bool`

HasAutotestsCount returns a boolean if a field has been set.

### SetAutotestsCountNil

`func (o *ApiV2ProjectsSearchPostRequest) SetAutotestsCountNil(b bool)`

 SetAutotestsCountNil sets the value for AutotestsCount to be an explicit nil

### UnsetAutotestsCount
`func (o *ApiV2ProjectsSearchPostRequest) UnsetAutotestsCount()`

UnsetAutotestsCount ensures that no value is present for AutotestsCount, not even an explicit nil
### GetGlobalIds

`func (o *ApiV2ProjectsSearchPostRequest) GetGlobalIds() []int64`

GetGlobalIds returns the GlobalIds field if non-nil, zero value otherwise.

### GetGlobalIdsOk

`func (o *ApiV2ProjectsSearchPostRequest) GetGlobalIdsOk() (*[]int64, bool)`

GetGlobalIdsOk returns a tuple with the GlobalIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalIds

`func (o *ApiV2ProjectsSearchPostRequest) SetGlobalIds(v []int64)`

SetGlobalIds sets GlobalIds field to given value.

### HasGlobalIds

`func (o *ApiV2ProjectsSearchPostRequest) HasGlobalIds() bool`

HasGlobalIds returns a boolean if a field has been set.

### SetGlobalIdsNil

`func (o *ApiV2ProjectsSearchPostRequest) SetGlobalIdsNil(b bool)`

 SetGlobalIdsNil sets the value for GlobalIds to be an explicit nil

### UnsetGlobalIds
`func (o *ApiV2ProjectsSearchPostRequest) UnsetGlobalIds()`

UnsetGlobalIds ensures that no value is present for GlobalIds, not even an explicit nil
### GetCreatedDate

`func (o *ApiV2ProjectsSearchPostRequest) GetCreatedDate() ProjectsFilterModelCreatedDate`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *ApiV2ProjectsSearchPostRequest) GetCreatedDateOk() (*ProjectsFilterModelCreatedDate, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *ApiV2ProjectsSearchPostRequest) SetCreatedDate(v ProjectsFilterModelCreatedDate)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *ApiV2ProjectsSearchPostRequest) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### SetCreatedDateNil

`func (o *ApiV2ProjectsSearchPostRequest) SetCreatedDateNil(b bool)`

 SetCreatedDateNil sets the value for CreatedDate to be an explicit nil

### UnsetCreatedDate
`func (o *ApiV2ProjectsSearchPostRequest) UnsetCreatedDate()`

UnsetCreatedDate ensures that no value is present for CreatedDate, not even an explicit nil
### GetCreatedByIds

`func (o *ApiV2ProjectsSearchPostRequest) GetCreatedByIds() []string`

GetCreatedByIds returns the CreatedByIds field if non-nil, zero value otherwise.

### GetCreatedByIdsOk

`func (o *ApiV2ProjectsSearchPostRequest) GetCreatedByIdsOk() (*[]string, bool)`

GetCreatedByIdsOk returns a tuple with the CreatedByIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByIds

`func (o *ApiV2ProjectsSearchPostRequest) SetCreatedByIds(v []string)`

SetCreatedByIds sets CreatedByIds field to given value.

### HasCreatedByIds

`func (o *ApiV2ProjectsSearchPostRequest) HasCreatedByIds() bool`

HasCreatedByIds returns a boolean if a field has been set.

### SetCreatedByIdsNil

`func (o *ApiV2ProjectsSearchPostRequest) SetCreatedByIdsNil(b bool)`

 SetCreatedByIdsNil sets the value for CreatedByIds to be an explicit nil

### UnsetCreatedByIds
`func (o *ApiV2ProjectsSearchPostRequest) UnsetCreatedByIds()`

UnsetCreatedByIds ensures that no value is present for CreatedByIds, not even an explicit nil
### GetTypes

`func (o *ApiV2ProjectsSearchPostRequest) GetTypes() []ProjectTypeModel`

GetTypes returns the Types field if non-nil, zero value otherwise.

### GetTypesOk

`func (o *ApiV2ProjectsSearchPostRequest) GetTypesOk() (*[]ProjectTypeModel, bool)`

GetTypesOk returns a tuple with the Types field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypes

`func (o *ApiV2ProjectsSearchPostRequest) SetTypes(v []ProjectTypeModel)`

SetTypes sets Types field to given value.

### HasTypes

`func (o *ApiV2ProjectsSearchPostRequest) HasTypes() bool`

HasTypes returns a boolean if a field has been set.

### SetTypesNil

`func (o *ApiV2ProjectsSearchPostRequest) SetTypesNil(b bool)`

 SetTypesNil sets the value for Types to be an explicit nil

### UnsetTypes
`func (o *ApiV2ProjectsSearchPostRequest) UnsetTypes()`

UnsetTypes ensures that no value is present for Types, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


