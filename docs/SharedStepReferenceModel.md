# SharedStepReferenceModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**GlobalId** | **int64** |  | 
**Name** | **string** |  | 
**EntityTypeName** | **string** |  | 
**HasThisSharedStepAsStep** | **bool** |  | 
**HasThisSharedStepAsPrecondition** | **bool** |  | 
**HasThisSharedStepAsPostcondition** | **bool** |  | 
**CreatedById** | **string** |  | 
**ModifiedById** | Pointer to **NullableString** |  | [optional] 
**CreatedDate** | Pointer to **NullableTime** |  | [optional] 
**ModifiedDate** | Pointer to **NullableTime** |  | [optional] 
**State** | **string** |  | 
**Priority** | [**WorkItemPriorityModel**](WorkItemPriorityModel.md) |  | 
**SourceType** | [**WorkItemSourceTypeModel**](WorkItemSourceTypeModel.md) |  | 
**IsDeleted** | **bool** |  | 
**VersionId** | **string** | used for versioning changes in workitem | 
**IsAutomated** | **bool** |  | 
**SectionId** | **string** |  | 
**Tags** | Pointer to [**[]TagModel**](TagModel.md) |  | [optional] 

## Methods

### NewSharedStepReferenceModel

`func NewSharedStepReferenceModel(id string, globalId int64, name string, entityTypeName string, hasThisSharedStepAsStep bool, hasThisSharedStepAsPrecondition bool, hasThisSharedStepAsPostcondition bool, createdById string, state string, priority WorkItemPriorityModel, sourceType WorkItemSourceTypeModel, isDeleted bool, versionId string, isAutomated bool, sectionId string, ) *SharedStepReferenceModel`

NewSharedStepReferenceModel instantiates a new SharedStepReferenceModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSharedStepReferenceModelWithDefaults

`func NewSharedStepReferenceModelWithDefaults() *SharedStepReferenceModel`

NewSharedStepReferenceModelWithDefaults instantiates a new SharedStepReferenceModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SharedStepReferenceModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SharedStepReferenceModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SharedStepReferenceModel) SetId(v string)`

SetId sets Id field to given value.


### GetGlobalId

`func (o *SharedStepReferenceModel) GetGlobalId() int64`

GetGlobalId returns the GlobalId field if non-nil, zero value otherwise.

### GetGlobalIdOk

`func (o *SharedStepReferenceModel) GetGlobalIdOk() (*int64, bool)`

GetGlobalIdOk returns a tuple with the GlobalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalId

`func (o *SharedStepReferenceModel) SetGlobalId(v int64)`

SetGlobalId sets GlobalId field to given value.


### GetName

`func (o *SharedStepReferenceModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SharedStepReferenceModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SharedStepReferenceModel) SetName(v string)`

SetName sets Name field to given value.


### GetEntityTypeName

`func (o *SharedStepReferenceModel) GetEntityTypeName() string`

GetEntityTypeName returns the EntityTypeName field if non-nil, zero value otherwise.

### GetEntityTypeNameOk

`func (o *SharedStepReferenceModel) GetEntityTypeNameOk() (*string, bool)`

GetEntityTypeNameOk returns a tuple with the EntityTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityTypeName

`func (o *SharedStepReferenceModel) SetEntityTypeName(v string)`

SetEntityTypeName sets EntityTypeName field to given value.


### GetHasThisSharedStepAsStep

`func (o *SharedStepReferenceModel) GetHasThisSharedStepAsStep() bool`

GetHasThisSharedStepAsStep returns the HasThisSharedStepAsStep field if non-nil, zero value otherwise.

### GetHasThisSharedStepAsStepOk

`func (o *SharedStepReferenceModel) GetHasThisSharedStepAsStepOk() (*bool, bool)`

GetHasThisSharedStepAsStepOk returns a tuple with the HasThisSharedStepAsStep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasThisSharedStepAsStep

`func (o *SharedStepReferenceModel) SetHasThisSharedStepAsStep(v bool)`

SetHasThisSharedStepAsStep sets HasThisSharedStepAsStep field to given value.


### GetHasThisSharedStepAsPrecondition

`func (o *SharedStepReferenceModel) GetHasThisSharedStepAsPrecondition() bool`

GetHasThisSharedStepAsPrecondition returns the HasThisSharedStepAsPrecondition field if non-nil, zero value otherwise.

### GetHasThisSharedStepAsPreconditionOk

`func (o *SharedStepReferenceModel) GetHasThisSharedStepAsPreconditionOk() (*bool, bool)`

GetHasThisSharedStepAsPreconditionOk returns a tuple with the HasThisSharedStepAsPrecondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasThisSharedStepAsPrecondition

`func (o *SharedStepReferenceModel) SetHasThisSharedStepAsPrecondition(v bool)`

SetHasThisSharedStepAsPrecondition sets HasThisSharedStepAsPrecondition field to given value.


### GetHasThisSharedStepAsPostcondition

`func (o *SharedStepReferenceModel) GetHasThisSharedStepAsPostcondition() bool`

GetHasThisSharedStepAsPostcondition returns the HasThisSharedStepAsPostcondition field if non-nil, zero value otherwise.

### GetHasThisSharedStepAsPostconditionOk

`func (o *SharedStepReferenceModel) GetHasThisSharedStepAsPostconditionOk() (*bool, bool)`

GetHasThisSharedStepAsPostconditionOk returns a tuple with the HasThisSharedStepAsPostcondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasThisSharedStepAsPostcondition

`func (o *SharedStepReferenceModel) SetHasThisSharedStepAsPostcondition(v bool)`

SetHasThisSharedStepAsPostcondition sets HasThisSharedStepAsPostcondition field to given value.


### GetCreatedById

`func (o *SharedStepReferenceModel) GetCreatedById() string`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *SharedStepReferenceModel) GetCreatedByIdOk() (*string, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *SharedStepReferenceModel) SetCreatedById(v string)`

SetCreatedById sets CreatedById field to given value.


### GetModifiedById

`func (o *SharedStepReferenceModel) GetModifiedById() string`

GetModifiedById returns the ModifiedById field if non-nil, zero value otherwise.

### GetModifiedByIdOk

`func (o *SharedStepReferenceModel) GetModifiedByIdOk() (*string, bool)`

GetModifiedByIdOk returns a tuple with the ModifiedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedById

`func (o *SharedStepReferenceModel) SetModifiedById(v string)`

SetModifiedById sets ModifiedById field to given value.

### HasModifiedById

`func (o *SharedStepReferenceModel) HasModifiedById() bool`

HasModifiedById returns a boolean if a field has been set.

### SetModifiedByIdNil

`func (o *SharedStepReferenceModel) SetModifiedByIdNil(b bool)`

 SetModifiedByIdNil sets the value for ModifiedById to be an explicit nil

### UnsetModifiedById
`func (o *SharedStepReferenceModel) UnsetModifiedById()`

UnsetModifiedById ensures that no value is present for ModifiedById, not even an explicit nil
### GetCreatedDate

`func (o *SharedStepReferenceModel) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *SharedStepReferenceModel) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *SharedStepReferenceModel) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *SharedStepReferenceModel) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### SetCreatedDateNil

`func (o *SharedStepReferenceModel) SetCreatedDateNil(b bool)`

 SetCreatedDateNil sets the value for CreatedDate to be an explicit nil

### UnsetCreatedDate
`func (o *SharedStepReferenceModel) UnsetCreatedDate()`

UnsetCreatedDate ensures that no value is present for CreatedDate, not even an explicit nil
### GetModifiedDate

`func (o *SharedStepReferenceModel) GetModifiedDate() time.Time`

GetModifiedDate returns the ModifiedDate field if non-nil, zero value otherwise.

### GetModifiedDateOk

`func (o *SharedStepReferenceModel) GetModifiedDateOk() (*time.Time, bool)`

GetModifiedDateOk returns a tuple with the ModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedDate

`func (o *SharedStepReferenceModel) SetModifiedDate(v time.Time)`

SetModifiedDate sets ModifiedDate field to given value.

### HasModifiedDate

`func (o *SharedStepReferenceModel) HasModifiedDate() bool`

HasModifiedDate returns a boolean if a field has been set.

### SetModifiedDateNil

`func (o *SharedStepReferenceModel) SetModifiedDateNil(b bool)`

 SetModifiedDateNil sets the value for ModifiedDate to be an explicit nil

### UnsetModifiedDate
`func (o *SharedStepReferenceModel) UnsetModifiedDate()`

UnsetModifiedDate ensures that no value is present for ModifiedDate, not even an explicit nil
### GetState

`func (o *SharedStepReferenceModel) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *SharedStepReferenceModel) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *SharedStepReferenceModel) SetState(v string)`

SetState sets State field to given value.


### GetPriority

`func (o *SharedStepReferenceModel) GetPriority() WorkItemPriorityModel`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *SharedStepReferenceModel) GetPriorityOk() (*WorkItemPriorityModel, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *SharedStepReferenceModel) SetPriority(v WorkItemPriorityModel)`

SetPriority sets Priority field to given value.


### GetSourceType

`func (o *SharedStepReferenceModel) GetSourceType() WorkItemSourceTypeModel`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *SharedStepReferenceModel) GetSourceTypeOk() (*WorkItemSourceTypeModel, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *SharedStepReferenceModel) SetSourceType(v WorkItemSourceTypeModel)`

SetSourceType sets SourceType field to given value.


### GetIsDeleted

`func (o *SharedStepReferenceModel) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *SharedStepReferenceModel) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *SharedStepReferenceModel) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.


### GetVersionId

`func (o *SharedStepReferenceModel) GetVersionId() string`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *SharedStepReferenceModel) GetVersionIdOk() (*string, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *SharedStepReferenceModel) SetVersionId(v string)`

SetVersionId sets VersionId field to given value.


### GetIsAutomated

`func (o *SharedStepReferenceModel) GetIsAutomated() bool`

GetIsAutomated returns the IsAutomated field if non-nil, zero value otherwise.

### GetIsAutomatedOk

`func (o *SharedStepReferenceModel) GetIsAutomatedOk() (*bool, bool)`

GetIsAutomatedOk returns a tuple with the IsAutomated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutomated

`func (o *SharedStepReferenceModel) SetIsAutomated(v bool)`

SetIsAutomated sets IsAutomated field to given value.


### GetSectionId

`func (o *SharedStepReferenceModel) GetSectionId() string`

GetSectionId returns the SectionId field if non-nil, zero value otherwise.

### GetSectionIdOk

`func (o *SharedStepReferenceModel) GetSectionIdOk() (*string, bool)`

GetSectionIdOk returns a tuple with the SectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectionId

`func (o *SharedStepReferenceModel) SetSectionId(v string)`

SetSectionId sets SectionId field to given value.


### GetTags

`func (o *SharedStepReferenceModel) GetTags() []TagModel`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *SharedStepReferenceModel) GetTagsOk() (*[]TagModel, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *SharedStepReferenceModel) SetTags(v []TagModel)`

SetTags sets Tags field to given value.

### HasTags

`func (o *SharedStepReferenceModel) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *SharedStepReferenceModel) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *SharedStepReferenceModel) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


