# TestStatusCreateModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Type** | [**TestStatusType**](TestStatusType.md) |  | 
**Code** | **string** |  | 
**Description** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewTestStatusCreateModel

`func NewTestStatusCreateModel(name string, type_ TestStatusType, code string, ) *TestStatusCreateModel`

NewTestStatusCreateModel instantiates a new TestStatusCreateModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestStatusCreateModelWithDefaults

`func NewTestStatusCreateModelWithDefaults() *TestStatusCreateModel`

NewTestStatusCreateModelWithDefaults instantiates a new TestStatusCreateModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TestStatusCreateModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TestStatusCreateModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TestStatusCreateModel) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *TestStatusCreateModel) GetType() TestStatusType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TestStatusCreateModel) GetTypeOk() (*TestStatusType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TestStatusCreateModel) SetType(v TestStatusType)`

SetType sets Type field to given value.


### GetCode

`func (o *TestStatusCreateModel) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *TestStatusCreateModel) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *TestStatusCreateModel) SetCode(v string)`

SetCode sets Code field to given value.


### GetDescription

`func (o *TestStatusCreateModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TestStatusCreateModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TestStatusCreateModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TestStatusCreateModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *TestStatusCreateModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *TestStatusCreateModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


