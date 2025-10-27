# UpdateFailureCategoryApiModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Failure category identifier | 
**Name** | **string** | Failure category name | 
**FailureCategory** | [**FailureCategory**](FailureCategory.md) | Category type | 
**FailureClassRegexes** | Pointer to [**[]UpdateFailureClassRegexApiModel**](UpdateFailureClassRegexApiModel.md) | Failure category regexes | [optional] 
**ProjectIds** | Pointer to **[]string** | Projects identifiers | [optional] 

## Methods

### NewUpdateFailureCategoryApiModel

`func NewUpdateFailureCategoryApiModel(id string, name string, failureCategory FailureCategory, ) *UpdateFailureCategoryApiModel`

NewUpdateFailureCategoryApiModel instantiates a new UpdateFailureCategoryApiModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateFailureCategoryApiModelWithDefaults

`func NewUpdateFailureCategoryApiModelWithDefaults() *UpdateFailureCategoryApiModel`

NewUpdateFailureCategoryApiModelWithDefaults instantiates a new UpdateFailureCategoryApiModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdateFailureCategoryApiModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateFailureCategoryApiModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateFailureCategoryApiModel) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *UpdateFailureCategoryApiModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateFailureCategoryApiModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateFailureCategoryApiModel) SetName(v string)`

SetName sets Name field to given value.


### GetFailureCategory

`func (o *UpdateFailureCategoryApiModel) GetFailureCategory() FailureCategory`

GetFailureCategory returns the FailureCategory field if non-nil, zero value otherwise.

### GetFailureCategoryOk

`func (o *UpdateFailureCategoryApiModel) GetFailureCategoryOk() (*FailureCategory, bool)`

GetFailureCategoryOk returns a tuple with the FailureCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureCategory

`func (o *UpdateFailureCategoryApiModel) SetFailureCategory(v FailureCategory)`

SetFailureCategory sets FailureCategory field to given value.


### GetFailureClassRegexes

`func (o *UpdateFailureCategoryApiModel) GetFailureClassRegexes() []UpdateFailureClassRegexApiModel`

GetFailureClassRegexes returns the FailureClassRegexes field if non-nil, zero value otherwise.

### GetFailureClassRegexesOk

`func (o *UpdateFailureCategoryApiModel) GetFailureClassRegexesOk() (*[]UpdateFailureClassRegexApiModel, bool)`

GetFailureClassRegexesOk returns a tuple with the FailureClassRegexes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureClassRegexes

`func (o *UpdateFailureCategoryApiModel) SetFailureClassRegexes(v []UpdateFailureClassRegexApiModel)`

SetFailureClassRegexes sets FailureClassRegexes field to given value.

### HasFailureClassRegexes

`func (o *UpdateFailureCategoryApiModel) HasFailureClassRegexes() bool`

HasFailureClassRegexes returns a boolean if a field has been set.

### SetFailureClassRegexesNil

`func (o *UpdateFailureCategoryApiModel) SetFailureClassRegexesNil(b bool)`

 SetFailureClassRegexesNil sets the value for FailureClassRegexes to be an explicit nil

### UnsetFailureClassRegexes
`func (o *UpdateFailureCategoryApiModel) UnsetFailureClassRegexes()`

UnsetFailureClassRegexes ensures that no value is present for FailureClassRegexes, not even an explicit nil
### GetProjectIds

`func (o *UpdateFailureCategoryApiModel) GetProjectIds() []string`

GetProjectIds returns the ProjectIds field if non-nil, zero value otherwise.

### GetProjectIdsOk

`func (o *UpdateFailureCategoryApiModel) GetProjectIdsOk() (*[]string, bool)`

GetProjectIdsOk returns a tuple with the ProjectIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectIds

`func (o *UpdateFailureCategoryApiModel) SetProjectIds(v []string)`

SetProjectIds sets ProjectIds field to given value.

### HasProjectIds

`func (o *UpdateFailureCategoryApiModel) HasProjectIds() bool`

HasProjectIds returns a boolean if a field has been set.

### SetProjectIdsNil

`func (o *UpdateFailureCategoryApiModel) SetProjectIdsNil(b bool)`

 SetProjectIdsNil sets the value for ProjectIds to be an explicit nil

### UnsetProjectIds
`func (o *UpdateFailureCategoryApiModel) UnsetProjectIds()`

UnsetProjectIds ensures that no value is present for ProjectIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


