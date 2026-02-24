# WorkItemChangeModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**WorkItemId** | **string** |  | 
**OldVersionId** | **string** |  | 
**NewVersionId** | **string** |  | 
**WorkItemChangedFields** | [**WorkItemChangedFieldsViewModel**](WorkItemChangedFieldsViewModel.md) |  | 
**CreatedById** | **string** |  | 
**CreatedDate** | **time.Time** |  | 
**ModifiedById** | Pointer to **NullableString** |  | [optional] 
**ModifiedDate** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewWorkItemChangeModel

`func NewWorkItemChangeModel(id string, workItemId string, oldVersionId string, newVersionId string, workItemChangedFields WorkItemChangedFieldsViewModel, createdById string, createdDate time.Time, ) *WorkItemChangeModel`

NewWorkItemChangeModel instantiates a new WorkItemChangeModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkItemChangeModelWithDefaults

`func NewWorkItemChangeModelWithDefaults() *WorkItemChangeModel`

NewWorkItemChangeModelWithDefaults instantiates a new WorkItemChangeModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WorkItemChangeModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkItemChangeModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkItemChangeModel) SetId(v string)`

SetId sets Id field to given value.


### GetWorkItemId

`func (o *WorkItemChangeModel) GetWorkItemId() string`

GetWorkItemId returns the WorkItemId field if non-nil, zero value otherwise.

### GetWorkItemIdOk

`func (o *WorkItemChangeModel) GetWorkItemIdOk() (*string, bool)`

GetWorkItemIdOk returns a tuple with the WorkItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkItemId

`func (o *WorkItemChangeModel) SetWorkItemId(v string)`

SetWorkItemId sets WorkItemId field to given value.


### GetOldVersionId

`func (o *WorkItemChangeModel) GetOldVersionId() string`

GetOldVersionId returns the OldVersionId field if non-nil, zero value otherwise.

### GetOldVersionIdOk

`func (o *WorkItemChangeModel) GetOldVersionIdOk() (*string, bool)`

GetOldVersionIdOk returns a tuple with the OldVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldVersionId

`func (o *WorkItemChangeModel) SetOldVersionId(v string)`

SetOldVersionId sets OldVersionId field to given value.


### GetNewVersionId

`func (o *WorkItemChangeModel) GetNewVersionId() string`

GetNewVersionId returns the NewVersionId field if non-nil, zero value otherwise.

### GetNewVersionIdOk

`func (o *WorkItemChangeModel) GetNewVersionIdOk() (*string, bool)`

GetNewVersionIdOk returns a tuple with the NewVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewVersionId

`func (o *WorkItemChangeModel) SetNewVersionId(v string)`

SetNewVersionId sets NewVersionId field to given value.


### GetWorkItemChangedFields

`func (o *WorkItemChangeModel) GetWorkItemChangedFields() WorkItemChangedFieldsViewModel`

GetWorkItemChangedFields returns the WorkItemChangedFields field if non-nil, zero value otherwise.

### GetWorkItemChangedFieldsOk

`func (o *WorkItemChangeModel) GetWorkItemChangedFieldsOk() (*WorkItemChangedFieldsViewModel, bool)`

GetWorkItemChangedFieldsOk returns a tuple with the WorkItemChangedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkItemChangedFields

`func (o *WorkItemChangeModel) SetWorkItemChangedFields(v WorkItemChangedFieldsViewModel)`

SetWorkItemChangedFields sets WorkItemChangedFields field to given value.


### GetCreatedById

`func (o *WorkItemChangeModel) GetCreatedById() string`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *WorkItemChangeModel) GetCreatedByIdOk() (*string, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *WorkItemChangeModel) SetCreatedById(v string)`

SetCreatedById sets CreatedById field to given value.


### GetCreatedDate

`func (o *WorkItemChangeModel) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *WorkItemChangeModel) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *WorkItemChangeModel) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.


### GetModifiedById

`func (o *WorkItemChangeModel) GetModifiedById() string`

GetModifiedById returns the ModifiedById field if non-nil, zero value otherwise.

### GetModifiedByIdOk

`func (o *WorkItemChangeModel) GetModifiedByIdOk() (*string, bool)`

GetModifiedByIdOk returns a tuple with the ModifiedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedById

`func (o *WorkItemChangeModel) SetModifiedById(v string)`

SetModifiedById sets ModifiedById field to given value.

### HasModifiedById

`func (o *WorkItemChangeModel) HasModifiedById() bool`

HasModifiedById returns a boolean if a field has been set.

### SetModifiedByIdNil

`func (o *WorkItemChangeModel) SetModifiedByIdNil(b bool)`

 SetModifiedByIdNil sets the value for ModifiedById to be an explicit nil

### UnsetModifiedById
`func (o *WorkItemChangeModel) UnsetModifiedById()`

UnsetModifiedById ensures that no value is present for ModifiedById, not even an explicit nil
### GetModifiedDate

`func (o *WorkItemChangeModel) GetModifiedDate() time.Time`

GetModifiedDate returns the ModifiedDate field if non-nil, zero value otherwise.

### GetModifiedDateOk

`func (o *WorkItemChangeModel) GetModifiedDateOk() (*time.Time, bool)`

GetModifiedDateOk returns a tuple with the ModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedDate

`func (o *WorkItemChangeModel) SetModifiedDate(v time.Time)`

SetModifiedDate sets ModifiedDate field to given value.

### HasModifiedDate

`func (o *WorkItemChangeModel) HasModifiedDate() bool`

HasModifiedDate returns a boolean if a field has been set.

### SetModifiedDateNil

`func (o *WorkItemChangeModel) SetModifiedDateNil(b bool)`

 SetModifiedDateNil sets the value for ModifiedDate to be an explicit nil

### UnsetModifiedDate
`func (o *WorkItemChangeModel) UnsetModifiedDate()`

UnsetModifiedDate ensures that no value is present for ModifiedDate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


