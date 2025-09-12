# CreateProjectFailureCategoryApiModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Failure category name | 
**FailureCategory** | [**FailureCategory**](FailureCategory.md) | Category type | 
**FailureClassRegexes** | Pointer to [**[]CreateFailureClassRegexApiModel**](CreateFailureClassRegexApiModel.md) | Failure category regexes | [optional] 

## Methods

### NewCreateProjectFailureCategoryApiModel

`func NewCreateProjectFailureCategoryApiModel(name string, failureCategory FailureCategory, ) *CreateProjectFailureCategoryApiModel`

NewCreateProjectFailureCategoryApiModel instantiates a new CreateProjectFailureCategoryApiModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateProjectFailureCategoryApiModelWithDefaults

`func NewCreateProjectFailureCategoryApiModelWithDefaults() *CreateProjectFailureCategoryApiModel`

NewCreateProjectFailureCategoryApiModelWithDefaults instantiates a new CreateProjectFailureCategoryApiModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateProjectFailureCategoryApiModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateProjectFailureCategoryApiModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateProjectFailureCategoryApiModel) SetName(v string)`

SetName sets Name field to given value.


### GetFailureCategory

`func (o *CreateProjectFailureCategoryApiModel) GetFailureCategory() FailureCategory`

GetFailureCategory returns the FailureCategory field if non-nil, zero value otherwise.

### GetFailureCategoryOk

`func (o *CreateProjectFailureCategoryApiModel) GetFailureCategoryOk() (*FailureCategory, bool)`

GetFailureCategoryOk returns a tuple with the FailureCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureCategory

`func (o *CreateProjectFailureCategoryApiModel) SetFailureCategory(v FailureCategory)`

SetFailureCategory sets FailureCategory field to given value.


### GetFailureClassRegexes

`func (o *CreateProjectFailureCategoryApiModel) GetFailureClassRegexes() []CreateFailureClassRegexApiModel`

GetFailureClassRegexes returns the FailureClassRegexes field if non-nil, zero value otherwise.

### GetFailureClassRegexesOk

`func (o *CreateProjectFailureCategoryApiModel) GetFailureClassRegexesOk() (*[]CreateFailureClassRegexApiModel, bool)`

GetFailureClassRegexesOk returns a tuple with the FailureClassRegexes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureClassRegexes

`func (o *CreateProjectFailureCategoryApiModel) SetFailureClassRegexes(v []CreateFailureClassRegexApiModel)`

SetFailureClassRegexes sets FailureClassRegexes field to given value.

### HasFailureClassRegexes

`func (o *CreateProjectFailureCategoryApiModel) HasFailureClassRegexes() bool`

HasFailureClassRegexes returns a boolean if a field has been set.

### SetFailureClassRegexesNil

`func (o *CreateProjectFailureCategoryApiModel) SetFailureClassRegexesNil(b bool)`

 SetFailureClassRegexesNil sets the value for FailureClassRegexes to be an explicit nil

### UnsetFailureClassRegexes
`func (o *CreateProjectFailureCategoryApiModel) UnsetFailureClassRegexes()`

UnsetFailureClassRegexes ensures that no value is present for FailureClassRegexes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


