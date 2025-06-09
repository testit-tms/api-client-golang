# CreateTestStatusApiModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the status, must be unique | 
**Type** | [**TestStatusApiType**](TestStatusApiType.md) | Type of the status | 
**Code** | **string** | Code of the status, must be unique | 
**Description** | Pointer to **NullableString** | Optional description of the status | [optional] 

## Methods

### NewCreateTestStatusApiModel

`func NewCreateTestStatusApiModel(name string, type_ TestStatusApiType, code string, ) *CreateTestStatusApiModel`

NewCreateTestStatusApiModel instantiates a new CreateTestStatusApiModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTestStatusApiModelWithDefaults

`func NewCreateTestStatusApiModelWithDefaults() *CreateTestStatusApiModel`

NewCreateTestStatusApiModelWithDefaults instantiates a new CreateTestStatusApiModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateTestStatusApiModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateTestStatusApiModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateTestStatusApiModel) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *CreateTestStatusApiModel) GetType() TestStatusApiType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateTestStatusApiModel) GetTypeOk() (*TestStatusApiType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateTestStatusApiModel) SetType(v TestStatusApiType)`

SetType sets Type field to given value.


### GetCode

`func (o *CreateTestStatusApiModel) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *CreateTestStatusApiModel) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *CreateTestStatusApiModel) SetCode(v string)`

SetCode sets Code field to given value.


### GetDescription

`func (o *CreateTestStatusApiModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateTestStatusApiModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateTestStatusApiModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateTestStatusApiModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CreateTestStatusApiModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CreateTestStatusApiModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


