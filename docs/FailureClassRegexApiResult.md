# FailureClassRegexApiResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Regex unique identifier | 
**IsDeleted** | **bool** | Indicates if the entity is deleted | 
**RegexText** | **string** | Regex value | 
**FailureClassId** | **string** | Failure category identifier | 

## Methods

### NewFailureClassRegexApiResult

`func NewFailureClassRegexApiResult(id string, isDeleted bool, regexText string, failureClassId string, ) *FailureClassRegexApiResult`

NewFailureClassRegexApiResult instantiates a new FailureClassRegexApiResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFailureClassRegexApiResultWithDefaults

`func NewFailureClassRegexApiResultWithDefaults() *FailureClassRegexApiResult`

NewFailureClassRegexApiResultWithDefaults instantiates a new FailureClassRegexApiResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FailureClassRegexApiResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FailureClassRegexApiResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FailureClassRegexApiResult) SetId(v string)`

SetId sets Id field to given value.


### GetIsDeleted

`func (o *FailureClassRegexApiResult) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *FailureClassRegexApiResult) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *FailureClassRegexApiResult) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.


### GetRegexText

`func (o *FailureClassRegexApiResult) GetRegexText() string`

GetRegexText returns the RegexText field if non-nil, zero value otherwise.

### GetRegexTextOk

`func (o *FailureClassRegexApiResult) GetRegexTextOk() (*string, bool)`

GetRegexTextOk returns a tuple with the RegexText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegexText

`func (o *FailureClassRegexApiResult) SetRegexText(v string)`

SetRegexText sets RegexText field to given value.


### GetFailureClassId

`func (o *FailureClassRegexApiResult) GetFailureClassId() string`

GetFailureClassId returns the FailureClassId field if non-nil, zero value otherwise.

### GetFailureClassIdOk

`func (o *FailureClassRegexApiResult) GetFailureClassIdOk() (*string, bool)`

GetFailureClassIdOk returns a tuple with the FailureClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureClassId

`func (o *FailureClassRegexApiResult) SetFailureClassId(v string)`

SetFailureClassId sets FailureClassId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


