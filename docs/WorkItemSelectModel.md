# WorkItemSelectModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filter** | [**WorkItemSelectModelFilter**](WorkItemSelectModelFilter.md) |  | 
**ExtractionModel** | Pointer to [**NullableWorkItemLocalSelectModelExtractionModel**](WorkItemLocalSelectModelExtractionModel.md) |  | [optional] 

## Methods

### NewWorkItemSelectModel

`func NewWorkItemSelectModel(filter WorkItemSelectModelFilter, ) *WorkItemSelectModel`

NewWorkItemSelectModel instantiates a new WorkItemSelectModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkItemSelectModelWithDefaults

`func NewWorkItemSelectModelWithDefaults() *WorkItemSelectModel`

NewWorkItemSelectModelWithDefaults instantiates a new WorkItemSelectModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilter

`func (o *WorkItemSelectModel) GetFilter() WorkItemSelectModelFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *WorkItemSelectModel) GetFilterOk() (*WorkItemSelectModelFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *WorkItemSelectModel) SetFilter(v WorkItemSelectModelFilter)`

SetFilter sets Filter field to given value.


### GetExtractionModel

`func (o *WorkItemSelectModel) GetExtractionModel() WorkItemLocalSelectModelExtractionModel`

GetExtractionModel returns the ExtractionModel field if non-nil, zero value otherwise.

### GetExtractionModelOk

`func (o *WorkItemSelectModel) GetExtractionModelOk() (*WorkItemLocalSelectModelExtractionModel, bool)`

GetExtractionModelOk returns a tuple with the ExtractionModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtractionModel

`func (o *WorkItemSelectModel) SetExtractionModel(v WorkItemLocalSelectModelExtractionModel)`

SetExtractionModel sets ExtractionModel field to given value.

### HasExtractionModel

`func (o *WorkItemSelectModel) HasExtractionModel() bool`

HasExtractionModel returns a boolean if a field has been set.

### SetExtractionModelNil

`func (o *WorkItemSelectModel) SetExtractionModelNil(b bool)`

 SetExtractionModelNil sets the value for ExtractionModel to be an explicit nil

### UnsetExtractionModel
`func (o *WorkItemSelectModel) UnsetExtractionModel()`

UnsetExtractionModel ensures that no value is present for ExtractionModel, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


