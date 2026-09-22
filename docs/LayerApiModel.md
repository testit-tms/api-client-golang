# LayerApiModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the test pyramid layer.              Available layers:              * Unit * Component * API * UI * E2E * Contract | 
**Source** | [**LayerSource**](LayerSource.md) | Source of the test pyramid layer. | 

## Methods

### NewLayerApiModel

`func NewLayerApiModel(name string, source LayerSource, ) *LayerApiModel`

NewLayerApiModel instantiates a new LayerApiModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLayerApiModelWithDefaults

`func NewLayerApiModelWithDefaults() *LayerApiModel`

NewLayerApiModelWithDefaults instantiates a new LayerApiModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *LayerApiModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LayerApiModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LayerApiModel) SetName(v string)`

SetName sets Name field to given value.


### GetSource

`func (o *LayerApiModel) GetSource() LayerSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *LayerApiModel) GetSourceOk() (*LayerSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *LayerApiModel) SetSource(v LayerSource)`

SetSource sets Source field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


