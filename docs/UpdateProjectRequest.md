# UpdateProjectRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique ID of the project | 
**Description** | Pointer to **NullableString** | Description of the project | [optional] 
**Name** | **string** | Name of the project | 
**IsFavorite** | Pointer to **NullableBool** | Indicates if the project is marked as favorite | [optional] 
**IsFlakyAuto** | Pointer to **NullableBool** | Indicates if the status \&quot;Flaky/Stable\&quot; sets automatically | [optional] 
**Type** | [**ProjectTypeModel**](ProjectTypeModel.md) |  | 

## Methods

### NewUpdateProjectRequest

`func NewUpdateProjectRequest(id string, name string, type_ ProjectTypeModel, ) *UpdateProjectRequest`

NewUpdateProjectRequest instantiates a new UpdateProjectRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateProjectRequestWithDefaults

`func NewUpdateProjectRequestWithDefaults() *UpdateProjectRequest`

NewUpdateProjectRequestWithDefaults instantiates a new UpdateProjectRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdateProjectRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateProjectRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateProjectRequest) SetId(v string)`

SetId sets Id field to given value.


### GetDescription

`func (o *UpdateProjectRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateProjectRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateProjectRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateProjectRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *UpdateProjectRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *UpdateProjectRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetName

`func (o *UpdateProjectRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateProjectRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateProjectRequest) SetName(v string)`

SetName sets Name field to given value.


### GetIsFavorite

`func (o *UpdateProjectRequest) GetIsFavorite() bool`

GetIsFavorite returns the IsFavorite field if non-nil, zero value otherwise.

### GetIsFavoriteOk

`func (o *UpdateProjectRequest) GetIsFavoriteOk() (*bool, bool)`

GetIsFavoriteOk returns a tuple with the IsFavorite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFavorite

`func (o *UpdateProjectRequest) SetIsFavorite(v bool)`

SetIsFavorite sets IsFavorite field to given value.

### HasIsFavorite

`func (o *UpdateProjectRequest) HasIsFavorite() bool`

HasIsFavorite returns a boolean if a field has been set.

### SetIsFavoriteNil

`func (o *UpdateProjectRequest) SetIsFavoriteNil(b bool)`

 SetIsFavoriteNil sets the value for IsFavorite to be an explicit nil

### UnsetIsFavorite
`func (o *UpdateProjectRequest) UnsetIsFavorite()`

UnsetIsFavorite ensures that no value is present for IsFavorite, not even an explicit nil
### GetIsFlakyAuto

`func (o *UpdateProjectRequest) GetIsFlakyAuto() bool`

GetIsFlakyAuto returns the IsFlakyAuto field if non-nil, zero value otherwise.

### GetIsFlakyAutoOk

`func (o *UpdateProjectRequest) GetIsFlakyAutoOk() (*bool, bool)`

GetIsFlakyAutoOk returns a tuple with the IsFlakyAuto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFlakyAuto

`func (o *UpdateProjectRequest) SetIsFlakyAuto(v bool)`

SetIsFlakyAuto sets IsFlakyAuto field to given value.

### HasIsFlakyAuto

`func (o *UpdateProjectRequest) HasIsFlakyAuto() bool`

HasIsFlakyAuto returns a boolean if a field has been set.

### SetIsFlakyAutoNil

`func (o *UpdateProjectRequest) SetIsFlakyAutoNil(b bool)`

 SetIsFlakyAutoNil sets the value for IsFlakyAuto to be an explicit nil

### UnsetIsFlakyAuto
`func (o *UpdateProjectRequest) UnsetIsFlakyAuto()`

UnsetIsFlakyAuto ensures that no value is present for IsFlakyAuto, not even an explicit nil
### GetType

`func (o *UpdateProjectRequest) GetType() ProjectTypeModel`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateProjectRequest) GetTypeOk() (*ProjectTypeModel, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateProjectRequest) SetType(v ProjectTypeModel)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


