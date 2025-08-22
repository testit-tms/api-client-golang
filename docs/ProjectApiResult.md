# ProjectApiResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique ID of the project | 
**Description** | Pointer to **NullableString** | Description of the project | [optional] 
**Name** | **string** | Name of the project | 
**IsFavorite** | **bool** | Indicates if the project is marked as favorite | 
**AttributesScheme** | Pointer to [**[]CustomAttributeModel**](CustomAttributeModel.md) | Collection of the project attributes | [optional] 
**TestPlansAttributesScheme** | Pointer to [**[]CustomAttributeModel**](CustomAttributeModel.md) | Collection of the project test plans attributes | [optional] 
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

### NewProjectApiResult

`func NewProjectApiResult(id string, name string, isFavorite bool, isDeleted bool, createdDate time.Time, createdById string, globalId int64, type_ ProjectTypeModel, isFlakyAuto bool, workflowId string, ) *ProjectApiResult`

NewProjectApiResult instantiates a new ProjectApiResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectApiResultWithDefaults

`func NewProjectApiResultWithDefaults() *ProjectApiResult`

NewProjectApiResultWithDefaults instantiates a new ProjectApiResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProjectApiResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProjectApiResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProjectApiResult) SetId(v string)`

SetId sets Id field to given value.


### GetDescription

`func (o *ProjectApiResult) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProjectApiResult) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProjectApiResult) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ProjectApiResult) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ProjectApiResult) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ProjectApiResult) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetName

`func (o *ProjectApiResult) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProjectApiResult) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProjectApiResult) SetName(v string)`

SetName sets Name field to given value.


### GetIsFavorite

`func (o *ProjectApiResult) GetIsFavorite() bool`

GetIsFavorite returns the IsFavorite field if non-nil, zero value otherwise.

### GetIsFavoriteOk

`func (o *ProjectApiResult) GetIsFavoriteOk() (*bool, bool)`

GetIsFavoriteOk returns a tuple with the IsFavorite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFavorite

`func (o *ProjectApiResult) SetIsFavorite(v bool)`

SetIsFavorite sets IsFavorite field to given value.


### GetAttributesScheme

`func (o *ProjectApiResult) GetAttributesScheme() []CustomAttributeModel`

GetAttributesScheme returns the AttributesScheme field if non-nil, zero value otherwise.

### GetAttributesSchemeOk

`func (o *ProjectApiResult) GetAttributesSchemeOk() (*[]CustomAttributeModel, bool)`

GetAttributesSchemeOk returns a tuple with the AttributesScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributesScheme

`func (o *ProjectApiResult) SetAttributesScheme(v []CustomAttributeModel)`

SetAttributesScheme sets AttributesScheme field to given value.

### HasAttributesScheme

`func (o *ProjectApiResult) HasAttributesScheme() bool`

HasAttributesScheme returns a boolean if a field has been set.

### SetAttributesSchemeNil

`func (o *ProjectApiResult) SetAttributesSchemeNil(b bool)`

 SetAttributesSchemeNil sets the value for AttributesScheme to be an explicit nil

### UnsetAttributesScheme
`func (o *ProjectApiResult) UnsetAttributesScheme()`

UnsetAttributesScheme ensures that no value is present for AttributesScheme, not even an explicit nil
### GetTestPlansAttributesScheme

`func (o *ProjectApiResult) GetTestPlansAttributesScheme() []CustomAttributeModel`

GetTestPlansAttributesScheme returns the TestPlansAttributesScheme field if non-nil, zero value otherwise.

### GetTestPlansAttributesSchemeOk

`func (o *ProjectApiResult) GetTestPlansAttributesSchemeOk() (*[]CustomAttributeModel, bool)`

GetTestPlansAttributesSchemeOk returns a tuple with the TestPlansAttributesScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestPlansAttributesScheme

`func (o *ProjectApiResult) SetTestPlansAttributesScheme(v []CustomAttributeModel)`

SetTestPlansAttributesScheme sets TestPlansAttributesScheme field to given value.

### HasTestPlansAttributesScheme

`func (o *ProjectApiResult) HasTestPlansAttributesScheme() bool`

HasTestPlansAttributesScheme returns a boolean if a field has been set.

### SetTestPlansAttributesSchemeNil

`func (o *ProjectApiResult) SetTestPlansAttributesSchemeNil(b bool)`

 SetTestPlansAttributesSchemeNil sets the value for TestPlansAttributesScheme to be an explicit nil

### UnsetTestPlansAttributesScheme
`func (o *ProjectApiResult) UnsetTestPlansAttributesScheme()`

UnsetTestPlansAttributesScheme ensures that no value is present for TestPlansAttributesScheme, not even an explicit nil
### GetTestCasesCount

`func (o *ProjectApiResult) GetTestCasesCount() int32`

GetTestCasesCount returns the TestCasesCount field if non-nil, zero value otherwise.

### GetTestCasesCountOk

`func (o *ProjectApiResult) GetTestCasesCountOk() (*int32, bool)`

GetTestCasesCountOk returns a tuple with the TestCasesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestCasesCount

`func (o *ProjectApiResult) SetTestCasesCount(v int32)`

SetTestCasesCount sets TestCasesCount field to given value.

### HasTestCasesCount

`func (o *ProjectApiResult) HasTestCasesCount() bool`

HasTestCasesCount returns a boolean if a field has been set.

### SetTestCasesCountNil

`func (o *ProjectApiResult) SetTestCasesCountNil(b bool)`

 SetTestCasesCountNil sets the value for TestCasesCount to be an explicit nil

### UnsetTestCasesCount
`func (o *ProjectApiResult) UnsetTestCasesCount()`

UnsetTestCasesCount ensures that no value is present for TestCasesCount, not even an explicit nil
### GetSharedStepsCount

`func (o *ProjectApiResult) GetSharedStepsCount() int32`

GetSharedStepsCount returns the SharedStepsCount field if non-nil, zero value otherwise.

### GetSharedStepsCountOk

`func (o *ProjectApiResult) GetSharedStepsCountOk() (*int32, bool)`

GetSharedStepsCountOk returns a tuple with the SharedStepsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedStepsCount

`func (o *ProjectApiResult) SetSharedStepsCount(v int32)`

SetSharedStepsCount sets SharedStepsCount field to given value.

### HasSharedStepsCount

`func (o *ProjectApiResult) HasSharedStepsCount() bool`

HasSharedStepsCount returns a boolean if a field has been set.

### SetSharedStepsCountNil

`func (o *ProjectApiResult) SetSharedStepsCountNil(b bool)`

 SetSharedStepsCountNil sets the value for SharedStepsCount to be an explicit nil

### UnsetSharedStepsCount
`func (o *ProjectApiResult) UnsetSharedStepsCount()`

UnsetSharedStepsCount ensures that no value is present for SharedStepsCount, not even an explicit nil
### GetCheckListsCount

`func (o *ProjectApiResult) GetCheckListsCount() int32`

GetCheckListsCount returns the CheckListsCount field if non-nil, zero value otherwise.

### GetCheckListsCountOk

`func (o *ProjectApiResult) GetCheckListsCountOk() (*int32, bool)`

GetCheckListsCountOk returns a tuple with the CheckListsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckListsCount

`func (o *ProjectApiResult) SetCheckListsCount(v int32)`

SetCheckListsCount sets CheckListsCount field to given value.

### HasCheckListsCount

`func (o *ProjectApiResult) HasCheckListsCount() bool`

HasCheckListsCount returns a boolean if a field has been set.

### SetCheckListsCountNil

`func (o *ProjectApiResult) SetCheckListsCountNil(b bool)`

 SetCheckListsCountNil sets the value for CheckListsCount to be an explicit nil

### UnsetCheckListsCount
`func (o *ProjectApiResult) UnsetCheckListsCount()`

UnsetCheckListsCount ensures that no value is present for CheckListsCount, not even an explicit nil
### GetAutoTestsCount

`func (o *ProjectApiResult) GetAutoTestsCount() int32`

GetAutoTestsCount returns the AutoTestsCount field if non-nil, zero value otherwise.

### GetAutoTestsCountOk

`func (o *ProjectApiResult) GetAutoTestsCountOk() (*int32, bool)`

GetAutoTestsCountOk returns a tuple with the AutoTestsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoTestsCount

`func (o *ProjectApiResult) SetAutoTestsCount(v int32)`

SetAutoTestsCount sets AutoTestsCount field to given value.

### HasAutoTestsCount

`func (o *ProjectApiResult) HasAutoTestsCount() bool`

HasAutoTestsCount returns a boolean if a field has been set.

### SetAutoTestsCountNil

`func (o *ProjectApiResult) SetAutoTestsCountNil(b bool)`

 SetAutoTestsCountNil sets the value for AutoTestsCount to be an explicit nil

### UnsetAutoTestsCount
`func (o *ProjectApiResult) UnsetAutoTestsCount()`

UnsetAutoTestsCount ensures that no value is present for AutoTestsCount, not even an explicit nil
### GetIsDeleted

`func (o *ProjectApiResult) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *ProjectApiResult) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *ProjectApiResult) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.


### GetCreatedDate

`func (o *ProjectApiResult) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *ProjectApiResult) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *ProjectApiResult) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.


### GetModifiedDate

`func (o *ProjectApiResult) GetModifiedDate() time.Time`

GetModifiedDate returns the ModifiedDate field if non-nil, zero value otherwise.

### GetModifiedDateOk

`func (o *ProjectApiResult) GetModifiedDateOk() (*time.Time, bool)`

GetModifiedDateOk returns a tuple with the ModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedDate

`func (o *ProjectApiResult) SetModifiedDate(v time.Time)`

SetModifiedDate sets ModifiedDate field to given value.

### HasModifiedDate

`func (o *ProjectApiResult) HasModifiedDate() bool`

HasModifiedDate returns a boolean if a field has been set.

### SetModifiedDateNil

`func (o *ProjectApiResult) SetModifiedDateNil(b bool)`

 SetModifiedDateNil sets the value for ModifiedDate to be an explicit nil

### UnsetModifiedDate
`func (o *ProjectApiResult) UnsetModifiedDate()`

UnsetModifiedDate ensures that no value is present for ModifiedDate, not even an explicit nil
### GetCreatedById

`func (o *ProjectApiResult) GetCreatedById() string`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *ProjectApiResult) GetCreatedByIdOk() (*string, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *ProjectApiResult) SetCreatedById(v string)`

SetCreatedById sets CreatedById field to given value.


### GetModifiedById

`func (o *ProjectApiResult) GetModifiedById() string`

GetModifiedById returns the ModifiedById field if non-nil, zero value otherwise.

### GetModifiedByIdOk

`func (o *ProjectApiResult) GetModifiedByIdOk() (*string, bool)`

GetModifiedByIdOk returns a tuple with the ModifiedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedById

`func (o *ProjectApiResult) SetModifiedById(v string)`

SetModifiedById sets ModifiedById field to given value.

### HasModifiedById

`func (o *ProjectApiResult) HasModifiedById() bool`

HasModifiedById returns a boolean if a field has been set.

### SetModifiedByIdNil

`func (o *ProjectApiResult) SetModifiedByIdNil(b bool)`

 SetModifiedByIdNil sets the value for ModifiedById to be an explicit nil

### UnsetModifiedById
`func (o *ProjectApiResult) UnsetModifiedById()`

UnsetModifiedById ensures that no value is present for ModifiedById, not even an explicit nil
### GetGlobalId

`func (o *ProjectApiResult) GetGlobalId() int64`

GetGlobalId returns the GlobalId field if non-nil, zero value otherwise.

### GetGlobalIdOk

`func (o *ProjectApiResult) GetGlobalIdOk() (*int64, bool)`

GetGlobalIdOk returns a tuple with the GlobalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalId

`func (o *ProjectApiResult) SetGlobalId(v int64)`

SetGlobalId sets GlobalId field to given value.


### GetType

`func (o *ProjectApiResult) GetType() ProjectTypeModel`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ProjectApiResult) GetTypeOk() (*ProjectTypeModel, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ProjectApiResult) SetType(v ProjectTypeModel)`

SetType sets Type field to given value.


### GetIsFlakyAuto

`func (o *ProjectApiResult) GetIsFlakyAuto() bool`

GetIsFlakyAuto returns the IsFlakyAuto field if non-nil, zero value otherwise.

### GetIsFlakyAutoOk

`func (o *ProjectApiResult) GetIsFlakyAutoOk() (*bool, bool)`

GetIsFlakyAutoOk returns a tuple with the IsFlakyAuto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFlakyAuto

`func (o *ProjectApiResult) SetIsFlakyAuto(v bool)`

SetIsFlakyAuto sets IsFlakyAuto field to given value.


### GetWorkflowId

`func (o *ProjectApiResult) GetWorkflowId() string`

GetWorkflowId returns the WorkflowId field if non-nil, zero value otherwise.

### GetWorkflowIdOk

`func (o *ProjectApiResult) GetWorkflowIdOk() (*string, bool)`

GetWorkflowIdOk returns a tuple with the WorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowId

`func (o *ProjectApiResult) SetWorkflowId(v string)`

SetWorkflowId sets WorkflowId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


