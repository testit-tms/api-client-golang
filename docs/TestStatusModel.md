# TestStatusModel

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

### NewTestStatusModel

`func NewTestStatusModel(id string, name string, type_ TestStatusType, isSystem bool, code string, ) *TestStatusModel`

NewTestStatusModel instantiates a new TestStatusModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestStatusModelWithDefaults

`func NewTestStatusModelWithDefaults() *TestStatusModel`

NewTestStatusModelWithDefaults instantiates a new TestStatusModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TestStatusModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TestStatusModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TestStatusModel) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *TestStatusModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TestStatusModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TestStatusModel) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *TestStatusModel) GetType() TestStatusType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TestStatusModel) GetTypeOk() (*TestStatusType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TestStatusModel) SetType(v TestStatusType)`

SetType sets Type field to given value.


### GetIsSystem

`func (o *TestStatusModel) GetIsSystem() bool`

GetIsSystem returns the IsSystem field if non-nil, zero value otherwise.

### GetIsSystemOk

`func (o *TestStatusModel) GetIsSystemOk() (*bool, bool)`

GetIsSystemOk returns a tuple with the IsSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSystem

`func (o *TestStatusModel) SetIsSystem(v bool)`

SetIsSystem sets IsSystem field to given value.


### GetCode

`func (o *TestStatusModel) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *TestStatusModel) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *TestStatusModel) SetCode(v string)`

SetCode sets Code field to given value.


### GetDescription

`func (o *TestStatusModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TestStatusModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TestStatusModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TestStatusModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *TestStatusModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *TestStatusModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


