# UpdateAutoTestResultReasonProjectApiModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Failure category identifier | 
**Name** | **string** | Failure category name | 
**FailureCategory** | [**FailureCategory**](FailureCategory.md) | Category type | 
**FailureClassRegexes** | Pointer to [**[]UpdateFailureClassRegexApiModel**](UpdateFailureClassRegexApiModel.md) | Failure category regexes | [optional] 
**ProjectIds** | Pointer to **[]string** | Projects identifiers | [optional] 

## Methods

### NewUpdateAutoTestResultReasonProjectApiModel

`func NewUpdateAutoTestResultReasonProjectApiModel(id string, name string, failureCategory FailureCategory, ) *UpdateAutoTestResultReasonProjectApiModel`

NewUpdateAutoTestResultReasonProjectApiModel instantiates a new UpdateAutoTestResultReasonProjectApiModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAutoTestResultReasonProjectApiModelWithDefaults

`func NewUpdateAutoTestResultReasonProjectApiModelWithDefaults() *UpdateAutoTestResultReasonProjectApiModel`

NewUpdateAutoTestResultReasonProjectApiModelWithDefaults instantiates a new UpdateAutoTestResultReasonProjectApiModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdateAutoTestResultReasonProjectApiModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateAutoTestResultReasonProjectApiModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateAutoTestResultReasonProjectApiModel) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *UpdateAutoTestResultReasonProjectApiModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateAutoTestResultReasonProjectApiModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateAutoTestResultReasonProjectApiModel) SetName(v string)`

SetName sets Name field to given value.


### GetFailureCategory

`func (o *UpdateAutoTestResultReasonProjectApiModel) GetFailureCategory() FailureCategory`

GetFailureCategory returns the FailureCategory field if non-nil, zero value otherwise.

### GetFailureCategoryOk

`func (o *UpdateAutoTestResultReasonProjectApiModel) GetFailureCategoryOk() (*FailureCategory, bool)`

GetFailureCategoryOk returns a tuple with the FailureCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureCategory

`func (o *UpdateAutoTestResultReasonProjectApiModel) SetFailureCategory(v FailureCategory)`

SetFailureCategory sets FailureCategory field to given value.


### GetFailureClassRegexes

`func (o *UpdateAutoTestResultReasonProjectApiModel) GetFailureClassRegexes() []UpdateFailureClassRegexApiModel`

GetFailureClassRegexes returns the FailureClassRegexes field if non-nil, zero value otherwise.

### GetFailureClassRegexesOk

`func (o *UpdateAutoTestResultReasonProjectApiModel) GetFailureClassRegexesOk() (*[]UpdateFailureClassRegexApiModel, bool)`

GetFailureClassRegexesOk returns a tuple with the FailureClassRegexes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureClassRegexes

`func (o *UpdateAutoTestResultReasonProjectApiModel) SetFailureClassRegexes(v []UpdateFailureClassRegexApiModel)`

SetFailureClassRegexes sets FailureClassRegexes field to given value.

### HasFailureClassRegexes

`func (o *UpdateAutoTestResultReasonProjectApiModel) HasFailureClassRegexes() bool`

HasFailureClassRegexes returns a boolean if a field has been set.

### SetFailureClassRegexesNil

`func (o *UpdateAutoTestResultReasonProjectApiModel) SetFailureClassRegexesNil(b bool)`

 SetFailureClassRegexesNil sets the value for FailureClassRegexes to be an explicit nil

### UnsetFailureClassRegexes
`func (o *UpdateAutoTestResultReasonProjectApiModel) UnsetFailureClassRegexes()`

UnsetFailureClassRegexes ensures that no value is present for FailureClassRegexes, not even an explicit nil
### GetProjectIds

`func (o *UpdateAutoTestResultReasonProjectApiModel) GetProjectIds() []string`

GetProjectIds returns the ProjectIds field if non-nil, zero value otherwise.

### GetProjectIdsOk

`func (o *UpdateAutoTestResultReasonProjectApiModel) GetProjectIdsOk() (*[]string, bool)`

GetProjectIdsOk returns a tuple with the ProjectIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectIds

`func (o *UpdateAutoTestResultReasonProjectApiModel) SetProjectIds(v []string)`

SetProjectIds sets ProjectIds field to given value.

### HasProjectIds

`func (o *UpdateAutoTestResultReasonProjectApiModel) HasProjectIds() bool`

HasProjectIds returns a boolean if a field has been set.

### SetProjectIdsNil

`func (o *UpdateAutoTestResultReasonProjectApiModel) SetProjectIdsNil(b bool)`

 SetProjectIdsNil sets the value for ProjectIds to be an explicit nil

### UnsetProjectIds
`func (o *UpdateAutoTestResultReasonProjectApiModel) UnsetProjectIds()`

UnsetProjectIds ensures that no value is present for ProjectIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


