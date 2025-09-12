# CreateAutoTestResultReasonProjectApiModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Failure category name | 
**FailureCategory** | [**FailureCategory**](FailureCategory.md) | Category type | 
**FailureClassRegexes** | Pointer to [**[]CreateFailureClassRegexApiModel**](CreateFailureClassRegexApiModel.md) | Failure category regexes | [optional] 
**ProjectIds** | Pointer to **[]string** | Projects identifiers | [optional] 

## Methods

### NewCreateAutoTestResultReasonProjectApiModel

`func NewCreateAutoTestResultReasonProjectApiModel(name string, failureCategory FailureCategory, ) *CreateAutoTestResultReasonProjectApiModel`

NewCreateAutoTestResultReasonProjectApiModel instantiates a new CreateAutoTestResultReasonProjectApiModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAutoTestResultReasonProjectApiModelWithDefaults

`func NewCreateAutoTestResultReasonProjectApiModelWithDefaults() *CreateAutoTestResultReasonProjectApiModel`

NewCreateAutoTestResultReasonProjectApiModelWithDefaults instantiates a new CreateAutoTestResultReasonProjectApiModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateAutoTestResultReasonProjectApiModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateAutoTestResultReasonProjectApiModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateAutoTestResultReasonProjectApiModel) SetName(v string)`

SetName sets Name field to given value.


### GetFailureCategory

`func (o *CreateAutoTestResultReasonProjectApiModel) GetFailureCategory() FailureCategory`

GetFailureCategory returns the FailureCategory field if non-nil, zero value otherwise.

### GetFailureCategoryOk

`func (o *CreateAutoTestResultReasonProjectApiModel) GetFailureCategoryOk() (*FailureCategory, bool)`

GetFailureCategoryOk returns a tuple with the FailureCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureCategory

`func (o *CreateAutoTestResultReasonProjectApiModel) SetFailureCategory(v FailureCategory)`

SetFailureCategory sets FailureCategory field to given value.


### GetFailureClassRegexes

`func (o *CreateAutoTestResultReasonProjectApiModel) GetFailureClassRegexes() []CreateFailureClassRegexApiModel`

GetFailureClassRegexes returns the FailureClassRegexes field if non-nil, zero value otherwise.

### GetFailureClassRegexesOk

`func (o *CreateAutoTestResultReasonProjectApiModel) GetFailureClassRegexesOk() (*[]CreateFailureClassRegexApiModel, bool)`

GetFailureClassRegexesOk returns a tuple with the FailureClassRegexes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureClassRegexes

`func (o *CreateAutoTestResultReasonProjectApiModel) SetFailureClassRegexes(v []CreateFailureClassRegexApiModel)`

SetFailureClassRegexes sets FailureClassRegexes field to given value.

### HasFailureClassRegexes

`func (o *CreateAutoTestResultReasonProjectApiModel) HasFailureClassRegexes() bool`

HasFailureClassRegexes returns a boolean if a field has been set.

### SetFailureClassRegexesNil

`func (o *CreateAutoTestResultReasonProjectApiModel) SetFailureClassRegexesNil(b bool)`

 SetFailureClassRegexesNil sets the value for FailureClassRegexes to be an explicit nil

### UnsetFailureClassRegexes
`func (o *CreateAutoTestResultReasonProjectApiModel) UnsetFailureClassRegexes()`

UnsetFailureClassRegexes ensures that no value is present for FailureClassRegexes, not even an explicit nil
### GetProjectIds

`func (o *CreateAutoTestResultReasonProjectApiModel) GetProjectIds() []string`

GetProjectIds returns the ProjectIds field if non-nil, zero value otherwise.

### GetProjectIdsOk

`func (o *CreateAutoTestResultReasonProjectApiModel) GetProjectIdsOk() (*[]string, bool)`

GetProjectIdsOk returns a tuple with the ProjectIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectIds

`func (o *CreateAutoTestResultReasonProjectApiModel) SetProjectIds(v []string)`

SetProjectIds sets ProjectIds field to given value.

### HasProjectIds

`func (o *CreateAutoTestResultReasonProjectApiModel) HasProjectIds() bool`

HasProjectIds returns a boolean if a field has been set.

### SetProjectIdsNil

`func (o *CreateAutoTestResultReasonProjectApiModel) SetProjectIdsNil(b bool)`

 SetProjectIdsNil sets the value for ProjectIds to be an explicit nil

### UnsetProjectIds
`func (o *CreateAutoTestResultReasonProjectApiModel) UnsetProjectIds()`

UnsetProjectIds ensures that no value is present for ProjectIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


