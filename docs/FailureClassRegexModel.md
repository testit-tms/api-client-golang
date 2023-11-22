# FailureClassRegexModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RegexText** | **string** |  | 
**FailureClassId** | Pointer to **NullableString** |  | [optional] 
**Id** | **string** | Unique ID of the entity | 
**IsDeleted** | **bool** | Indicates if the entity is deleted | 

## Methods

### NewFailureClassRegexModel

`func NewFailureClassRegexModel(regexText string, id string, isDeleted bool, ) *FailureClassRegexModel`

NewFailureClassRegexModel instantiates a new FailureClassRegexModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFailureClassRegexModelWithDefaults

`func NewFailureClassRegexModelWithDefaults() *FailureClassRegexModel`

NewFailureClassRegexModelWithDefaults instantiates a new FailureClassRegexModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegexText

`func (o *FailureClassRegexModel) GetRegexText() string`

GetRegexText returns the RegexText field if non-nil, zero value otherwise.

### GetRegexTextOk

`func (o *FailureClassRegexModel) GetRegexTextOk() (*string, bool)`

GetRegexTextOk returns a tuple with the RegexText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegexText

`func (o *FailureClassRegexModel) SetRegexText(v string)`

SetRegexText sets RegexText field to given value.


### GetFailureClassId

`func (o *FailureClassRegexModel) GetFailureClassId() string`

GetFailureClassId returns the FailureClassId field if non-nil, zero value otherwise.

### GetFailureClassIdOk

`func (o *FailureClassRegexModel) GetFailureClassIdOk() (*string, bool)`

GetFailureClassIdOk returns a tuple with the FailureClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureClassId

`func (o *FailureClassRegexModel) SetFailureClassId(v string)`

SetFailureClassId sets FailureClassId field to given value.

### HasFailureClassId

`func (o *FailureClassRegexModel) HasFailureClassId() bool`

HasFailureClassId returns a boolean if a field has been set.

### SetFailureClassIdNil

`func (o *FailureClassRegexModel) SetFailureClassIdNil(b bool)`

 SetFailureClassIdNil sets the value for FailureClassId to be an explicit nil

### UnsetFailureClassId
`func (o *FailureClassRegexModel) UnsetFailureClassId()`

UnsetFailureClassId ensures that no value is present for FailureClassId, not even an explicit nil
### GetId

`func (o *FailureClassRegexModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FailureClassRegexModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FailureClassRegexModel) SetId(v string)`

SetId sets Id field to given value.


### GetIsDeleted

`func (o *FailureClassRegexModel) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *FailureClassRegexModel) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *FailureClassRegexModel) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


