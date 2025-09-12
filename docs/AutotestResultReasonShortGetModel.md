# AutotestResultReasonShortGetModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsDeleted** | **bool** |  | 
**FailureCategory** | [**AvailableFailureCategory**](AvailableFailureCategory.md) |  | 
**FailureCategoryId** | **int32** |  | 
**Name** | Pointer to **NullableString** |  | [optional] 
**RegexCount** | **int32** |  | 
**Id** | **string** |  | 
**CreatedDate** | **time.Time** |  | 
**CreatedById** | **string** |  | 
**ModifiedDate** | Pointer to **NullableTime** |  | [optional] 
**ModifiedById** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewAutotestResultReasonShortGetModel

`func NewAutotestResultReasonShortGetModel(isDeleted bool, failureCategory AvailableFailureCategory, failureCategoryId int32, regexCount int32, id string, createdDate time.Time, createdById string, ) *AutotestResultReasonShortGetModel`

NewAutotestResultReasonShortGetModel instantiates a new AutotestResultReasonShortGetModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutotestResultReasonShortGetModelWithDefaults

`func NewAutotestResultReasonShortGetModelWithDefaults() *AutotestResultReasonShortGetModel`

NewAutotestResultReasonShortGetModelWithDefaults instantiates a new AutotestResultReasonShortGetModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsDeleted

`func (o *AutotestResultReasonShortGetModel) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *AutotestResultReasonShortGetModel) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *AutotestResultReasonShortGetModel) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.


### GetFailureCategory

`func (o *AutotestResultReasonShortGetModel) GetFailureCategory() AvailableFailureCategory`

GetFailureCategory returns the FailureCategory field if non-nil, zero value otherwise.

### GetFailureCategoryOk

`func (o *AutotestResultReasonShortGetModel) GetFailureCategoryOk() (*AvailableFailureCategory, bool)`

GetFailureCategoryOk returns a tuple with the FailureCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureCategory

`func (o *AutotestResultReasonShortGetModel) SetFailureCategory(v AvailableFailureCategory)`

SetFailureCategory sets FailureCategory field to given value.


### GetFailureCategoryId

`func (o *AutotestResultReasonShortGetModel) GetFailureCategoryId() int32`

GetFailureCategoryId returns the FailureCategoryId field if non-nil, zero value otherwise.

### GetFailureCategoryIdOk

`func (o *AutotestResultReasonShortGetModel) GetFailureCategoryIdOk() (*int32, bool)`

GetFailureCategoryIdOk returns a tuple with the FailureCategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureCategoryId

`func (o *AutotestResultReasonShortGetModel) SetFailureCategoryId(v int32)`

SetFailureCategoryId sets FailureCategoryId field to given value.


### GetName

`func (o *AutotestResultReasonShortGetModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AutotestResultReasonShortGetModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AutotestResultReasonShortGetModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AutotestResultReasonShortGetModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *AutotestResultReasonShortGetModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AutotestResultReasonShortGetModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetRegexCount

`func (o *AutotestResultReasonShortGetModel) GetRegexCount() int32`

GetRegexCount returns the RegexCount field if non-nil, zero value otherwise.

### GetRegexCountOk

`func (o *AutotestResultReasonShortGetModel) GetRegexCountOk() (*int32, bool)`

GetRegexCountOk returns a tuple with the RegexCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegexCount

`func (o *AutotestResultReasonShortGetModel) SetRegexCount(v int32)`

SetRegexCount sets RegexCount field to given value.


### GetId

`func (o *AutotestResultReasonShortGetModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AutotestResultReasonShortGetModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AutotestResultReasonShortGetModel) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedDate

`func (o *AutotestResultReasonShortGetModel) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *AutotestResultReasonShortGetModel) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *AutotestResultReasonShortGetModel) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.


### GetCreatedById

`func (o *AutotestResultReasonShortGetModel) GetCreatedById() string`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *AutotestResultReasonShortGetModel) GetCreatedByIdOk() (*string, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *AutotestResultReasonShortGetModel) SetCreatedById(v string)`

SetCreatedById sets CreatedById field to given value.


### GetModifiedDate

`func (o *AutotestResultReasonShortGetModel) GetModifiedDate() time.Time`

GetModifiedDate returns the ModifiedDate field if non-nil, zero value otherwise.

### GetModifiedDateOk

`func (o *AutotestResultReasonShortGetModel) GetModifiedDateOk() (*time.Time, bool)`

GetModifiedDateOk returns a tuple with the ModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedDate

`func (o *AutotestResultReasonShortGetModel) SetModifiedDate(v time.Time)`

SetModifiedDate sets ModifiedDate field to given value.

### HasModifiedDate

`func (o *AutotestResultReasonShortGetModel) HasModifiedDate() bool`

HasModifiedDate returns a boolean if a field has been set.

### SetModifiedDateNil

`func (o *AutotestResultReasonShortGetModel) SetModifiedDateNil(b bool)`

 SetModifiedDateNil sets the value for ModifiedDate to be an explicit nil

### UnsetModifiedDate
`func (o *AutotestResultReasonShortGetModel) UnsetModifiedDate()`

UnsetModifiedDate ensures that no value is present for ModifiedDate, not even an explicit nil
### GetModifiedById

`func (o *AutotestResultReasonShortGetModel) GetModifiedById() string`

GetModifiedById returns the ModifiedById field if non-nil, zero value otherwise.

### GetModifiedByIdOk

`func (o *AutotestResultReasonShortGetModel) GetModifiedByIdOk() (*string, bool)`

GetModifiedByIdOk returns a tuple with the ModifiedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedById

`func (o *AutotestResultReasonShortGetModel) SetModifiedById(v string)`

SetModifiedById sets ModifiedById field to given value.

### HasModifiedById

`func (o *AutotestResultReasonShortGetModel) HasModifiedById() bool`

HasModifiedById returns a boolean if a field has been set.

### SetModifiedByIdNil

`func (o *AutotestResultReasonShortGetModel) SetModifiedByIdNil(b bool)`

 SetModifiedByIdNil sets the value for ModifiedById to be an explicit nil

### UnsetModifiedById
`func (o *AutotestResultReasonShortGetModel) UnsetModifiedById()`

UnsetModifiedById ensures that no value is present for ModifiedById, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


