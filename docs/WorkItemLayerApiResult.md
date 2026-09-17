# WorkItemLayerApiResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the test pyramid layer. | 
**Source** | [**LayerSource**](LayerSource.md) | Source of the test pyramid layer. | 

## Methods

### NewWorkItemLayerApiResult

`func NewWorkItemLayerApiResult(name string, source LayerSource, ) *WorkItemLayerApiResult`

NewWorkItemLayerApiResult instantiates a new WorkItemLayerApiResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkItemLayerApiResultWithDefaults

`func NewWorkItemLayerApiResultWithDefaults() *WorkItemLayerApiResult`

NewWorkItemLayerApiResultWithDefaults instantiates a new WorkItemLayerApiResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *WorkItemLayerApiResult) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkItemLayerApiResult) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkItemLayerApiResult) SetName(v string)`

SetName sets Name field to given value.


### GetSource

`func (o *WorkItemLayerApiResult) GetSource() LayerSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *WorkItemLayerApiResult) GetSourceOk() (*LayerSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *WorkItemLayerApiResult) SetSource(v LayerSource)`

SetSource sets Source field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


