# ExternalFormModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fields** | [**[]ExternalFormFieldModel**](ExternalFormFieldModel.md) |  | 
**PossibleValues** | [**map[string][]ExternalFormAllowedValueModel**](array.md) |  | 

## Methods

### NewExternalFormModel

`func NewExternalFormModel(fields []ExternalFormFieldModel, possibleValues map[string][]ExternalFormAllowedValueModel, ) *ExternalFormModel`

NewExternalFormModel instantiates a new ExternalFormModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalFormModelWithDefaults

`func NewExternalFormModelWithDefaults() *ExternalFormModel`

NewExternalFormModelWithDefaults instantiates a new ExternalFormModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFields

`func (o *ExternalFormModel) GetFields() []ExternalFormFieldModel`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *ExternalFormModel) GetFieldsOk() (*[]ExternalFormFieldModel, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *ExternalFormModel) SetFields(v []ExternalFormFieldModel)`

SetFields sets Fields field to given value.


### GetPossibleValues

`func (o *ExternalFormModel) GetPossibleValues() map[string][]ExternalFormAllowedValueModel`

GetPossibleValues returns the PossibleValues field if non-nil, zero value otherwise.

### GetPossibleValuesOk

`func (o *ExternalFormModel) GetPossibleValuesOk() (*map[string][]ExternalFormAllowedValueModel, bool)`

GetPossibleValuesOk returns a tuple with the PossibleValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPossibleValues

`func (o *ExternalFormModel) SetPossibleValues(v map[string][]ExternalFormAllowedValueModel)`

SetPossibleValues sets PossibleValues field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


