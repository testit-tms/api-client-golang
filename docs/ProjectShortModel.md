# ProjectShortModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique ID of the project | 
**Description** | Pointer to **NullableString** | Description of the project | [optional] 
**Name** | **string** | Name of the project | 
**IsFavorite** | **bool** | Indicates if the project is marked as favorite | 
**TestCasesCount** | Pointer to **NullableInt32** | Number of test cases in the project | [optional] 
**SharedStepsCount** | Pointer to **NullableInt32** | Number of shared steps in the project | [optional] 
**CheckListsCount** | Pointer to **NullableInt32** | Number of checklists in the project | [optional] 
**AutoTestsCount** | Pointer to **NullableInt32** | Number of autotests in the project | [optional] 
**IsDeleted** | **bool** | Indicates if the project is deleted | 
**CreatedDate** | **time.Time** | Creation date of the project | 
**ModifiedDate** | Pointer to **NullableTime** | Last modification date of the project | [optional] 
**CreatedById** | **string** | Unique ID of the project creator | 
**ModifiedById** | Pointer to **NullableString** | Unique ID of the project last editor | [optional] 
**GlobalId** | **int64** | Global ID of the project | 
**Type** | [**ProjectTypeModel**](ProjectTypeModel.md) | Type of the project | 
**IsFlakyAuto** | **bool** | Indicates if the status \&quot;Flaky/Stable\&quot; sets automatically | 
**WorkflowId** | **string** |  | 

## Methods

### NewProjectShortModel

`func NewProjectShortModel(id string, name string, isFavorite bool, isDeleted bool, createdDate time.Time, createdById string, globalId int64, type_ ProjectTypeModel, isFlakyAuto bool, workflowId string, ) *ProjectShortModel`

NewProjectShortModel instantiates a new ProjectShortModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectShortModelWithDefaults

`func NewProjectShortModelWithDefaults() *ProjectShortModel`

NewProjectShortModelWithDefaults instantiates a new ProjectShortModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProjectShortModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProjectShortModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProjectShortModel) SetId(v string)`

SetId sets Id field to given value.


### GetDescription

`func (o *ProjectShortModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProjectShortModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProjectShortModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ProjectShortModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ProjectShortModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ProjectShortModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetName

`func (o *ProjectShortModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProjectShortModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProjectShortModel) SetName(v string)`

SetName sets Name field to given value.


### GetIsFavorite

`func (o *ProjectShortModel) GetIsFavorite() bool`

GetIsFavorite returns the IsFavorite field if non-nil, zero value otherwise.

### GetIsFavoriteOk

`func (o *ProjectShortModel) GetIsFavoriteOk() (*bool, bool)`

GetIsFavoriteOk returns a tuple with the IsFavorite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFavorite

`func (o *ProjectShortModel) SetIsFavorite(v bool)`

SetIsFavorite sets IsFavorite field to given value.


### GetTestCasesCount

`func (o *ProjectShortModel) GetTestCasesCount() int32`

GetTestCasesCount returns the TestCasesCount field if non-nil, zero value otherwise.

### GetTestCasesCountOk

`func (o *ProjectShortModel) GetTestCasesCountOk() (*int32, bool)`

GetTestCasesCountOk returns a tuple with the TestCasesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestCasesCount

`func (o *ProjectShortModel) SetTestCasesCount(v int32)`

SetTestCasesCount sets TestCasesCount field to given value.

### HasTestCasesCount

`func (o *ProjectShortModel) HasTestCasesCount() bool`

HasTestCasesCount returns a boolean if a field has been set.

### SetTestCasesCountNil

`func (o *ProjectShortModel) SetTestCasesCountNil(b bool)`

 SetTestCasesCountNil sets the value for TestCasesCount to be an explicit nil

### UnsetTestCasesCount
`func (o *ProjectShortModel) UnsetTestCasesCount()`

UnsetTestCasesCount ensures that no value is present for TestCasesCount, not even an explicit nil
### GetSharedStepsCount

`func (o *ProjectShortModel) GetSharedStepsCount() int32`

GetSharedStepsCount returns the SharedStepsCount field if non-nil, zero value otherwise.

### GetSharedStepsCountOk

`func (o *ProjectShortModel) GetSharedStepsCountOk() (*int32, bool)`

GetSharedStepsCountOk returns a tuple with the SharedStepsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedStepsCount

`func (o *ProjectShortModel) SetSharedStepsCount(v int32)`

SetSharedStepsCount sets SharedStepsCount field to given value.

### HasSharedStepsCount

`func (o *ProjectShortModel) HasSharedStepsCount() bool`

HasSharedStepsCount returns a boolean if a field has been set.

### SetSharedStepsCountNil

`func (o *ProjectShortModel) SetSharedStepsCountNil(b bool)`

 SetSharedStepsCountNil sets the value for SharedStepsCount to be an explicit nil

### UnsetSharedStepsCount
`func (o *ProjectShortModel) UnsetSharedStepsCount()`

UnsetSharedStepsCount ensures that no value is present for SharedStepsCount, not even an explicit nil
### GetCheckListsCount

`func (o *ProjectShortModel) GetCheckListsCount() int32`

GetCheckListsCount returns the CheckListsCount field if non-nil, zero value otherwise.

### GetCheckListsCountOk

`func (o *ProjectShortModel) GetCheckListsCountOk() (*int32, bool)`

GetCheckListsCountOk returns a tuple with the CheckListsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckListsCount

`func (o *ProjectShortModel) SetCheckListsCount(v int32)`

SetCheckListsCount sets CheckListsCount field to given value.

### HasCheckListsCount

`func (o *ProjectShortModel) HasCheckListsCount() bool`

HasCheckListsCount returns a boolean if a field has been set.

### SetCheckListsCountNil

`func (o *ProjectShortModel) SetCheckListsCountNil(b bool)`

 SetCheckListsCountNil sets the value for CheckListsCount to be an explicit nil

### UnsetCheckListsCount
`func (o *ProjectShortModel) UnsetCheckListsCount()`

UnsetCheckListsCount ensures that no value is present for CheckListsCount, not even an explicit nil
### GetAutoTestsCount

`func (o *ProjectShortModel) GetAutoTestsCount() int32`

GetAutoTestsCount returns the AutoTestsCount field if non-nil, zero value otherwise.

### GetAutoTestsCountOk

`func (o *ProjectShortModel) GetAutoTestsCountOk() (*int32, bool)`

GetAutoTestsCountOk returns a tuple with the AutoTestsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoTestsCount

`func (o *ProjectShortModel) SetAutoTestsCount(v int32)`

SetAutoTestsCount sets AutoTestsCount field to given value.

### HasAutoTestsCount

`func (o *ProjectShortModel) HasAutoTestsCount() bool`

HasAutoTestsCount returns a boolean if a field has been set.

### SetAutoTestsCountNil

`func (o *ProjectShortModel) SetAutoTestsCountNil(b bool)`

 SetAutoTestsCountNil sets the value for AutoTestsCount to be an explicit nil

### UnsetAutoTestsCount
`func (o *ProjectShortModel) UnsetAutoTestsCount()`

UnsetAutoTestsCount ensures that no value is present for AutoTestsCount, not even an explicit nil
### GetIsDeleted

`func (o *ProjectShortModel) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *ProjectShortModel) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *ProjectShortModel) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.


### GetCreatedDate

`func (o *ProjectShortModel) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *ProjectShortModel) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *ProjectShortModel) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.


### GetModifiedDate

`func (o *ProjectShortModel) GetModifiedDate() time.Time`

GetModifiedDate returns the ModifiedDate field if non-nil, zero value otherwise.

### GetModifiedDateOk

`func (o *ProjectShortModel) GetModifiedDateOk() (*time.Time, bool)`

GetModifiedDateOk returns a tuple with the ModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedDate

`func (o *ProjectShortModel) SetModifiedDate(v time.Time)`

SetModifiedDate sets ModifiedDate field to given value.

### HasModifiedDate

`func (o *ProjectShortModel) HasModifiedDate() bool`

HasModifiedDate returns a boolean if a field has been set.

### SetModifiedDateNil

`func (o *ProjectShortModel) SetModifiedDateNil(b bool)`

 SetModifiedDateNil sets the value for ModifiedDate to be an explicit nil

### UnsetModifiedDate
`func (o *ProjectShortModel) UnsetModifiedDate()`

UnsetModifiedDate ensures that no value is present for ModifiedDate, not even an explicit nil
### GetCreatedById

`func (o *ProjectShortModel) GetCreatedById() string`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *ProjectShortModel) GetCreatedByIdOk() (*string, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *ProjectShortModel) SetCreatedById(v string)`

SetCreatedById sets CreatedById field to given value.


### GetModifiedById

`func (o *ProjectShortModel) GetModifiedById() string`

GetModifiedById returns the ModifiedById field if non-nil, zero value otherwise.

### GetModifiedByIdOk

`func (o *ProjectShortModel) GetModifiedByIdOk() (*string, bool)`

GetModifiedByIdOk returns a tuple with the ModifiedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedById

`func (o *ProjectShortModel) SetModifiedById(v string)`

SetModifiedById sets ModifiedById field to given value.

### HasModifiedById

`func (o *ProjectShortModel) HasModifiedById() bool`

HasModifiedById returns a boolean if a field has been set.

### SetModifiedByIdNil

`func (o *ProjectShortModel) SetModifiedByIdNil(b bool)`

 SetModifiedByIdNil sets the value for ModifiedById to be an explicit nil

### UnsetModifiedById
`func (o *ProjectShortModel) UnsetModifiedById()`

UnsetModifiedById ensures that no value is present for ModifiedById, not even an explicit nil
### GetGlobalId

`func (o *ProjectShortModel) GetGlobalId() int64`

GetGlobalId returns the GlobalId field if non-nil, zero value otherwise.

### GetGlobalIdOk

`func (o *ProjectShortModel) GetGlobalIdOk() (*int64, bool)`

GetGlobalIdOk returns a tuple with the GlobalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalId

`func (o *ProjectShortModel) SetGlobalId(v int64)`

SetGlobalId sets GlobalId field to given value.


### GetType

`func (o *ProjectShortModel) GetType() ProjectTypeModel`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ProjectShortModel) GetTypeOk() (*ProjectTypeModel, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ProjectShortModel) SetType(v ProjectTypeModel)`

SetType sets Type field to given value.


### GetIsFlakyAuto

`func (o *ProjectShortModel) GetIsFlakyAuto() bool`

GetIsFlakyAuto returns the IsFlakyAuto field if non-nil, zero value otherwise.

### GetIsFlakyAutoOk

`func (o *ProjectShortModel) GetIsFlakyAutoOk() (*bool, bool)`

GetIsFlakyAutoOk returns a tuple with the IsFlakyAuto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFlakyAuto

`func (o *ProjectShortModel) SetIsFlakyAuto(v bool)`

SetIsFlakyAuto sets IsFlakyAuto field to given value.


### GetWorkflowId

`func (o *ProjectShortModel) GetWorkflowId() string`

GetWorkflowId returns the WorkflowId field if non-nil, zero value otherwise.

### GetWorkflowIdOk

`func (o *ProjectShortModel) GetWorkflowIdOk() (*string, bool)`

GetWorkflowIdOk returns a tuple with the WorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowId

`func (o *ProjectShortModel) SetWorkflowId(v string)`

SetWorkflowId sets WorkflowId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


