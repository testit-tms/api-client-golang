# UpdateParameterApiModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID&#39;s of the parameter | 
**Name** | **string** | Key of the parameter | 
**Value** | **string** | Value of the parameter | 
**ProjectIds** | Pointer to **[]string** | List of projects where parameter should be available | [optional] 

## Methods

### NewUpdateParameterApiModel

`func NewUpdateParameterApiModel(id string, name string, value string, ) *UpdateParameterApiModel`

NewUpdateParameterApiModel instantiates a new UpdateParameterApiModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateParameterApiModelWithDefaults

`func NewUpdateParameterApiModelWithDefaults() *UpdateParameterApiModel`

NewUpdateParameterApiModelWithDefaults instantiates a new UpdateParameterApiModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdateParameterApiModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateParameterApiModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateParameterApiModel) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *UpdateParameterApiModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateParameterApiModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateParameterApiModel) SetName(v string)`

SetName sets Name field to given value.


### GetValue

`func (o *UpdateParameterApiModel) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *UpdateParameterApiModel) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *UpdateParameterApiModel) SetValue(v string)`

SetValue sets Value field to given value.


### GetProjectIds

`func (o *UpdateParameterApiModel) GetProjectIds() []string`

GetProjectIds returns the ProjectIds field if non-nil, zero value otherwise.

### GetProjectIdsOk

`func (o *UpdateParameterApiModel) GetProjectIdsOk() (*[]string, bool)`

GetProjectIdsOk returns a tuple with the ProjectIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectIds

`func (o *UpdateParameterApiModel) SetProjectIds(v []string)`

SetProjectIds sets ProjectIds field to given value.

### HasProjectIds

`func (o *UpdateParameterApiModel) HasProjectIds() bool`

HasProjectIds returns a boolean if a field has been set.

### SetProjectIdsNil

`func (o *UpdateParameterApiModel) SetProjectIdsNil(b bool)`

 SetProjectIdsNil sets the value for ProjectIds to be an explicit nil

### UnsetProjectIds
`func (o *UpdateParameterApiModel) UnsetProjectIds()`

UnsetProjectIds ensures that no value is present for ProjectIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


