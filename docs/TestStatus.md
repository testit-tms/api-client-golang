# TestStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Type** | [**TestStatusType**](TestStatusType.md) |  | 
**IsBased** | **bool** |  | 
**IsDefault** | **bool** |  | 
**Code** | **string** |  | 
**Description** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewTestStatus

`func NewTestStatus(id string, name string, type_ TestStatusType, isBased bool, isDefault bool, code string, ) *TestStatus`

NewTestStatus instantiates a new TestStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestStatusWithDefaults

`func NewTestStatusWithDefaults() *TestStatus`

NewTestStatusWithDefaults instantiates a new TestStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TestStatus) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TestStatus) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TestStatus) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *TestStatus) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TestStatus) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TestStatus) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *TestStatus) GetType() TestStatusType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TestStatus) GetTypeOk() (*TestStatusType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TestStatus) SetType(v TestStatusType)`

SetType sets Type field to given value.


### GetIsBased

`func (o *TestStatus) GetIsBased() bool`

GetIsBased returns the IsBased field if non-nil, zero value otherwise.

### GetIsBasedOk

`func (o *TestStatus) GetIsBasedOk() (*bool, bool)`

GetIsBasedOk returns a tuple with the IsBased field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBased

`func (o *TestStatus) SetIsBased(v bool)`

SetIsBased sets IsBased field to given value.


### GetIsDefault

`func (o *TestStatus) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *TestStatus) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *TestStatus) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.


### GetCode

`func (o *TestStatus) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *TestStatus) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *TestStatus) SetCode(v string)`

SetCode sets Code field to given value.


### GetDescription

`func (o *TestStatus) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TestStatus) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TestStatus) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TestStatus) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *TestStatus) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *TestStatus) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


