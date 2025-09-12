# UpdateFailureCategoryProjectApiModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Failure category identifier | 
**Name** | **string** | Failure category name | 
**FailureCategory** | [**FailureCategory**](FailureCategory.md) | Category type | 
**FailureClassRegexes** | Pointer to [**[]UpdateFailureClassRegexApiModel**](UpdateFailureClassRegexApiModel.md) | Failure category regexes | [optional] 

## Methods

### NewUpdateFailureCategoryProjectApiModel

`func NewUpdateFailureCategoryProjectApiModel(id string, name string, failureCategory FailureCategory, ) *UpdateFailureCategoryProjectApiModel`

NewUpdateFailureCategoryProjectApiModel instantiates a new UpdateFailureCategoryProjectApiModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateFailureCategoryProjectApiModelWithDefaults

`func NewUpdateFailureCategoryProjectApiModelWithDefaults() *UpdateFailureCategoryProjectApiModel`

NewUpdateFailureCategoryProjectApiModelWithDefaults instantiates a new UpdateFailureCategoryProjectApiModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdateFailureCategoryProjectApiModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateFailureCategoryProjectApiModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateFailureCategoryProjectApiModel) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *UpdateFailureCategoryProjectApiModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateFailureCategoryProjectApiModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateFailureCategoryProjectApiModel) SetName(v string)`

SetName sets Name field to given value.


### GetFailureCategory

`func (o *UpdateFailureCategoryProjectApiModel) GetFailureCategory() FailureCategory`

GetFailureCategory returns the FailureCategory field if non-nil, zero value otherwise.

### GetFailureCategoryOk

`func (o *UpdateFailureCategoryProjectApiModel) GetFailureCategoryOk() (*FailureCategory, bool)`

GetFailureCategoryOk returns a tuple with the FailureCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureCategory

`func (o *UpdateFailureCategoryProjectApiModel) SetFailureCategory(v FailureCategory)`

SetFailureCategory sets FailureCategory field to given value.


### GetFailureClassRegexes

`func (o *UpdateFailureCategoryProjectApiModel) GetFailureClassRegexes() []UpdateFailureClassRegexApiModel`

GetFailureClassRegexes returns the FailureClassRegexes field if non-nil, zero value otherwise.

### GetFailureClassRegexesOk

`func (o *UpdateFailureCategoryProjectApiModel) GetFailureClassRegexesOk() (*[]UpdateFailureClassRegexApiModel, bool)`

GetFailureClassRegexesOk returns a tuple with the FailureClassRegexes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureClassRegexes

`func (o *UpdateFailureCategoryProjectApiModel) SetFailureClassRegexes(v []UpdateFailureClassRegexApiModel)`

SetFailureClassRegexes sets FailureClassRegexes field to given value.

### HasFailureClassRegexes

`func (o *UpdateFailureCategoryProjectApiModel) HasFailureClassRegexes() bool`

HasFailureClassRegexes returns a boolean if a field has been set.

### SetFailureClassRegexesNil

`func (o *UpdateFailureCategoryProjectApiModel) SetFailureClassRegexesNil(b bool)`

 SetFailureClassRegexesNil sets the value for FailureClassRegexes to be an explicit nil

### UnsetFailureClassRegexes
`func (o *UpdateFailureCategoryProjectApiModel) UnsetFailureClassRegexes()`

UnsetFailureClassRegexes ensures that no value is present for FailureClassRegexes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


