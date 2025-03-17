# TestStatusApiResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Type** | [**TestStatusType**](TestStatusType.md) |  | 
**IsSystem** | **bool** |  | 
**Code** | **string** |  | 
**Description** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewTestStatusApiResult

`func NewTestStatusApiResult(id string, name string, type_ TestStatusType, isSystem bool, code string, ) *TestStatusApiResult`

NewTestStatusApiResult instantiates a new TestStatusApiResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestStatusApiResultWithDefaults

`func NewTestStatusApiResultWithDefaults() *TestStatusApiResult`

NewTestStatusApiResultWithDefaults instantiates a new TestStatusApiResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TestStatusApiResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TestStatusApiResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TestStatusApiResult) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *TestStatusApiResult) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TestStatusApiResult) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TestStatusApiResult) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *TestStatusApiResult) GetType() TestStatusType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TestStatusApiResult) GetTypeOk() (*TestStatusType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TestStatusApiResult) SetType(v TestStatusType)`

SetType sets Type field to given value.


### GetIsSystem

`func (o *TestStatusApiResult) GetIsSystem() bool`

GetIsSystem returns the IsSystem field if non-nil, zero value otherwise.

### GetIsSystemOk

`func (o *TestStatusApiResult) GetIsSystemOk() (*bool, bool)`

GetIsSystemOk returns a tuple with the IsSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSystem

`func (o *TestStatusApiResult) SetIsSystem(v bool)`

SetIsSystem sets IsSystem field to given value.


### GetCode

`func (o *TestStatusApiResult) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *TestStatusApiResult) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *TestStatusApiResult) SetCode(v string)`

SetCode sets Code field to given value.


### GetDescription

`func (o *TestStatusApiResult) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TestStatusApiResult) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TestStatusApiResult) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TestStatusApiResult) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *TestStatusApiResult) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *TestStatusApiResult) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


