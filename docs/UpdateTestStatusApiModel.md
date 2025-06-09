# UpdateTestStatusApiModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the status, must be unique | 
**Description** | Pointer to **NullableString** | Optional description of the status | [optional] 

## Methods

### NewUpdateTestStatusApiModel

`func NewUpdateTestStatusApiModel(name string, ) *UpdateTestStatusApiModel`

NewUpdateTestStatusApiModel instantiates a new UpdateTestStatusApiModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateTestStatusApiModelWithDefaults

`func NewUpdateTestStatusApiModelWithDefaults() *UpdateTestStatusApiModel`

NewUpdateTestStatusApiModelWithDefaults instantiates a new UpdateTestStatusApiModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateTestStatusApiModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateTestStatusApiModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateTestStatusApiModel) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *UpdateTestStatusApiModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateTestStatusApiModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateTestStatusApiModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateTestStatusApiModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *UpdateTestStatusApiModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *UpdateTestStatusApiModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


