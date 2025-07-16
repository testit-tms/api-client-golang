# CreateParameterApiModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Key of the parameter | 
**Value** | **string** | Value of the parameter | 
**ProjectIds** | Pointer to **[]string** | List of projects where parameter should be available | [optional] 

## Methods

### NewCreateParameterApiModel

`func NewCreateParameterApiModel(name string, value string, ) *CreateParameterApiModel`

NewCreateParameterApiModel instantiates a new CreateParameterApiModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateParameterApiModelWithDefaults

`func NewCreateParameterApiModelWithDefaults() *CreateParameterApiModel`

NewCreateParameterApiModelWithDefaults instantiates a new CreateParameterApiModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateParameterApiModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateParameterApiModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateParameterApiModel) SetName(v string)`

SetName sets Name field to given value.


### GetValue

`func (o *CreateParameterApiModel) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CreateParameterApiModel) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CreateParameterApiModel) SetValue(v string)`

SetValue sets Value field to given value.


### GetProjectIds

`func (o *CreateParameterApiModel) GetProjectIds() []string`

GetProjectIds returns the ProjectIds field if non-nil, zero value otherwise.

### GetProjectIdsOk

`func (o *CreateParameterApiModel) GetProjectIdsOk() (*[]string, bool)`

GetProjectIdsOk returns a tuple with the ProjectIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectIds

`func (o *CreateParameterApiModel) SetProjectIds(v []string)`

SetProjectIds sets ProjectIds field to given value.

### HasProjectIds

`func (o *CreateParameterApiModel) HasProjectIds() bool`

HasProjectIds returns a boolean if a field has been set.

### SetProjectIdsNil

`func (o *CreateParameterApiModel) SetProjectIdsNil(b bool)`

 SetProjectIdsNil sets the value for ProjectIds to be an explicit nil

### UnsetProjectIds
`func (o *CreateParameterApiModel) UnsetProjectIds()`

UnsetProjectIds ensures that no value is present for ProjectIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


