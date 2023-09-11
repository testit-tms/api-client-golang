# ProjectSelectModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filter** | Pointer to [**NullableProjectsFilterModel**](ProjectsFilterModel.md) |  | [optional] 
**ExtractionModel** | Pointer to [**NullableProjectExtractionModel**](ProjectExtractionModel.md) |  | [optional] 

## Methods

### NewProjectSelectModel

`func NewProjectSelectModel() *ProjectSelectModel`

NewProjectSelectModel instantiates a new ProjectSelectModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectSelectModelWithDefaults

`func NewProjectSelectModelWithDefaults() *ProjectSelectModel`

NewProjectSelectModelWithDefaults instantiates a new ProjectSelectModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilter

`func (o *ProjectSelectModel) GetFilter() ProjectsFilterModel`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *ProjectSelectModel) GetFilterOk() (*ProjectsFilterModel, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *ProjectSelectModel) SetFilter(v ProjectsFilterModel)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *ProjectSelectModel) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### SetFilterNil

`func (o *ProjectSelectModel) SetFilterNil(b bool)`

 SetFilterNil sets the value for Filter to be an explicit nil

### UnsetFilter
`func (o *ProjectSelectModel) UnsetFilter()`

UnsetFilter ensures that no value is present for Filter, not even an explicit nil
### GetExtractionModel

`func (o *ProjectSelectModel) GetExtractionModel() ProjectExtractionModel`

GetExtractionModel returns the ExtractionModel field if non-nil, zero value otherwise.

### GetExtractionModelOk

`func (o *ProjectSelectModel) GetExtractionModelOk() (*ProjectExtractionModel, bool)`

GetExtractionModelOk returns a tuple with the ExtractionModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtractionModel

`func (o *ProjectSelectModel) SetExtractionModel(v ProjectExtractionModel)`

SetExtractionModel sets ExtractionModel field to given value.

### HasExtractionModel

`func (o *ProjectSelectModel) HasExtractionModel() bool`

HasExtractionModel returns a boolean if a field has been set.

### SetExtractionModelNil

`func (o *ProjectSelectModel) SetExtractionModelNil(b bool)`

 SetExtractionModelNil sets the value for ExtractionModel to be an explicit nil

### UnsetExtractionModel
`func (o *ProjectSelectModel) UnsetExtractionModel()`

UnsetExtractionModel ensures that no value is present for ExtractionModel, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


