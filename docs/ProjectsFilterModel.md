# ProjectsFilterModel

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

### NewProjectsFilterModel

`func NewProjectsFilterModel() *ProjectsFilterModel`

NewProjectsFilterModel instantiates a new ProjectsFilterModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectsFilterModelWithDefaults

`func NewProjectsFilterModelWithDefaults() *ProjectsFilterModel`

NewProjectsFilterModelWithDefaults instantiates a new ProjectsFilterModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ProjectsFilterModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProjectsFilterModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProjectsFilterModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProjectsFilterModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ProjectsFilterModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ProjectsFilterModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetIsFavorite

`func (o *ProjectsFilterModel) GetIsFavorite() bool`

GetIsFavorite returns the IsFavorite field if non-nil, zero value otherwise.

### GetIsFavoriteOk

`func (o *ProjectsFilterModel) GetIsFavoriteOk() (*bool, bool)`

GetIsFavoriteOk returns a tuple with the IsFavorite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFavorite

`func (o *ProjectsFilterModel) SetIsFavorite(v bool)`

SetIsFavorite sets IsFavorite field to given value.

### HasIsFavorite

`func (o *ProjectsFilterModel) HasIsFavorite() bool`

HasIsFavorite returns a boolean if a field has been set.

### SetIsFavoriteNil

`func (o *ProjectsFilterModel) SetIsFavoriteNil(b bool)`

 SetIsFavoriteNil sets the value for IsFavorite to be an explicit nil

### UnsetIsFavorite
`func (o *ProjectsFilterModel) UnsetIsFavorite()`

UnsetIsFavorite ensures that no value is present for IsFavorite, not even an explicit nil
### GetIsDeleted

`func (o *ProjectsFilterModel) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *ProjectsFilterModel) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *ProjectsFilterModel) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *ProjectsFilterModel) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### SetIsDeletedNil

`func (o *ProjectsFilterModel) SetIsDeletedNil(b bool)`

 SetIsDeletedNil sets the value for IsDeleted to be an explicit nil

### UnsetIsDeleted
`func (o *ProjectsFilterModel) UnsetIsDeleted()`

UnsetIsDeleted ensures that no value is present for IsDeleted, not even an explicit nil
### GetTestCasesCount

`func (o *ProjectsFilterModel) GetTestCasesCount() ProjectsFilterModelTestCasesCount`

GetTestCasesCount returns the TestCasesCount field if non-nil, zero value otherwise.

### GetTestCasesCountOk

`func (o *ProjectsFilterModel) GetTestCasesCountOk() (*ProjectsFilterModelTestCasesCount, bool)`

GetTestCasesCountOk returns a tuple with the TestCasesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestCasesCount

`func (o *ProjectsFilterModel) SetTestCasesCount(v ProjectsFilterModelTestCasesCount)`

SetTestCasesCount sets TestCasesCount field to given value.

### HasTestCasesCount

`func (o *ProjectsFilterModel) HasTestCasesCount() bool`

HasTestCasesCount returns a boolean if a field has been set.

### SetTestCasesCountNil

`func (o *ProjectsFilterModel) SetTestCasesCountNil(b bool)`

 SetTestCasesCountNil sets the value for TestCasesCount to be an explicit nil

### UnsetTestCasesCount
`func (o *ProjectsFilterModel) UnsetTestCasesCount()`

UnsetTestCasesCount ensures that no value is present for TestCasesCount, not even an explicit nil
### GetChecklistsCount

`func (o *ProjectsFilterModel) GetChecklistsCount() ProjectsFilterModelChecklistsCount`

GetChecklistsCount returns the ChecklistsCount field if non-nil, zero value otherwise.

### GetChecklistsCountOk

`func (o *ProjectsFilterModel) GetChecklistsCountOk() (*ProjectsFilterModelChecklistsCount, bool)`

GetChecklistsCountOk returns a tuple with the ChecklistsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecklistsCount

`func (o *ProjectsFilterModel) SetChecklistsCount(v ProjectsFilterModelChecklistsCount)`

SetChecklistsCount sets ChecklistsCount field to given value.

### HasChecklistsCount

`func (o *ProjectsFilterModel) HasChecklistsCount() bool`

HasChecklistsCount returns a boolean if a field has been set.

### SetChecklistsCountNil

`func (o *ProjectsFilterModel) SetChecklistsCountNil(b bool)`

 SetChecklistsCountNil sets the value for ChecklistsCount to be an explicit nil

### UnsetChecklistsCount
`func (o *ProjectsFilterModel) UnsetChecklistsCount()`

UnsetChecklistsCount ensures that no value is present for ChecklistsCount, not even an explicit nil
### GetSharedStepsCount

`func (o *ProjectsFilterModel) GetSharedStepsCount() ProjectsFilterModelSharedStepsCount`

GetSharedStepsCount returns the SharedStepsCount field if non-nil, zero value otherwise.

### GetSharedStepsCountOk

`func (o *ProjectsFilterModel) GetSharedStepsCountOk() (*ProjectsFilterModelSharedStepsCount, bool)`

GetSharedStepsCountOk returns a tuple with the SharedStepsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedStepsCount

`func (o *ProjectsFilterModel) SetSharedStepsCount(v ProjectsFilterModelSharedStepsCount)`

SetSharedStepsCount sets SharedStepsCount field to given value.

### HasSharedStepsCount

`func (o *ProjectsFilterModel) HasSharedStepsCount() bool`

HasSharedStepsCount returns a boolean if a field has been set.

### SetSharedStepsCountNil

`func (o *ProjectsFilterModel) SetSharedStepsCountNil(b bool)`

 SetSharedStepsCountNil sets the value for SharedStepsCount to be an explicit nil

### UnsetSharedStepsCount
`func (o *ProjectsFilterModel) UnsetSharedStepsCount()`

UnsetSharedStepsCount ensures that no value is present for SharedStepsCount, not even an explicit nil
### GetAutotestsCount

`func (o *ProjectsFilterModel) GetAutotestsCount() ProjectsFilterModelAutotestsCount`

GetAutotestsCount returns the AutotestsCount field if non-nil, zero value otherwise.

### GetAutotestsCountOk

`func (o *ProjectsFilterModel) GetAutotestsCountOk() (*ProjectsFilterModelAutotestsCount, bool)`

GetAutotestsCountOk returns a tuple with the AutotestsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutotestsCount

`func (o *ProjectsFilterModel) SetAutotestsCount(v ProjectsFilterModelAutotestsCount)`

SetAutotestsCount sets AutotestsCount field to given value.

### HasAutotestsCount

`func (o *ProjectsFilterModel) HasAutotestsCount() bool`

HasAutotestsCount returns a boolean if a field has been set.

### SetAutotestsCountNil

`func (o *ProjectsFilterModel) SetAutotestsCountNil(b bool)`

 SetAutotestsCountNil sets the value for AutotestsCount to be an explicit nil

### UnsetAutotestsCount
`func (o *ProjectsFilterModel) UnsetAutotestsCount()`

UnsetAutotestsCount ensures that no value is present for AutotestsCount, not even an explicit nil
### GetGlobalIds

`func (o *ProjectsFilterModel) GetGlobalIds() []int64`

GetGlobalIds returns the GlobalIds field if non-nil, zero value otherwise.

### GetGlobalIdsOk

`func (o *ProjectsFilterModel) GetGlobalIdsOk() (*[]int64, bool)`

GetGlobalIdsOk returns a tuple with the GlobalIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalIds

`func (o *ProjectsFilterModel) SetGlobalIds(v []int64)`

SetGlobalIds sets GlobalIds field to given value.

### HasGlobalIds

`func (o *ProjectsFilterModel) HasGlobalIds() bool`

HasGlobalIds returns a boolean if a field has been set.

### SetGlobalIdsNil

`func (o *ProjectsFilterModel) SetGlobalIdsNil(b bool)`

 SetGlobalIdsNil sets the value for GlobalIds to be an explicit nil

### UnsetGlobalIds
`func (o *ProjectsFilterModel) UnsetGlobalIds()`

UnsetGlobalIds ensures that no value is present for GlobalIds, not even an explicit nil
### GetCreatedDate

`func (o *ProjectsFilterModel) GetCreatedDate() ProjectsFilterModelCreatedDate`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *ProjectsFilterModel) GetCreatedDateOk() (*ProjectsFilterModelCreatedDate, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *ProjectsFilterModel) SetCreatedDate(v ProjectsFilterModelCreatedDate)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *ProjectsFilterModel) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### SetCreatedDateNil

`func (o *ProjectsFilterModel) SetCreatedDateNil(b bool)`

 SetCreatedDateNil sets the value for CreatedDate to be an explicit nil

### UnsetCreatedDate
`func (o *ProjectsFilterModel) UnsetCreatedDate()`

UnsetCreatedDate ensures that no value is present for CreatedDate, not even an explicit nil
### GetCreatedByIds

`func (o *ProjectsFilterModel) GetCreatedByIds() []string`

GetCreatedByIds returns the CreatedByIds field if non-nil, zero value otherwise.

### GetCreatedByIdsOk

`func (o *ProjectsFilterModel) GetCreatedByIdsOk() (*[]string, bool)`

GetCreatedByIdsOk returns a tuple with the CreatedByIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByIds

`func (o *ProjectsFilterModel) SetCreatedByIds(v []string)`

SetCreatedByIds sets CreatedByIds field to given value.

### HasCreatedByIds

`func (o *ProjectsFilterModel) HasCreatedByIds() bool`

HasCreatedByIds returns a boolean if a field has been set.

### SetCreatedByIdsNil

`func (o *ProjectsFilterModel) SetCreatedByIdsNil(b bool)`

 SetCreatedByIdsNil sets the value for CreatedByIds to be an explicit nil

### UnsetCreatedByIds
`func (o *ProjectsFilterModel) UnsetCreatedByIds()`

UnsetCreatedByIds ensures that no value is present for CreatedByIds, not even an explicit nil
### GetTypes

`func (o *ProjectsFilterModel) GetTypes() []ProjectTypeModel`

GetTypes returns the Types field if non-nil, zero value otherwise.

### GetTypesOk

`func (o *ProjectsFilterModel) GetTypesOk() (*[]ProjectTypeModel, bool)`

GetTypesOk returns a tuple with the Types field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypes

`func (o *ProjectsFilterModel) SetTypes(v []ProjectTypeModel)`

SetTypes sets Types field to given value.

### HasTypes

`func (o *ProjectsFilterModel) HasTypes() bool`

HasTypes returns a boolean if a field has been set.

### SetTypesNil

`func (o *ProjectsFilterModel) SetTypesNil(b bool)`

 SetTypesNil sets the value for Types to be an explicit nil

### UnsetTypes
`func (o *ProjectsFilterModel) UnsetTypes()`

UnsetTypes ensures that no value is present for Types, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


