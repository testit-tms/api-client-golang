# ExternalFormCreateModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PossibleValues** | [**map[string][]ExternalFormAllowedValueModel**](array.md) |  | 
**Fields** | [**[]ExternalFormFieldModel**](ExternalFormFieldModel.md) |  | 
**Links** | [**[]ExternalFormLinkModel**](ExternalFormLinkModel.md) |  | 
**Values** | **map[string]interface{}** |  | 

## Methods

### NewExternalFormCreateModel

`func NewExternalFormCreateModel(possibleValues map[string][]ExternalFormAllowedValueModel, fields []ExternalFormFieldModel, links []ExternalFormLinkModel, values map[string]interface{}, ) *ExternalFormCreateModel`

NewExternalFormCreateModel instantiates a new ExternalFormCreateModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalFormCreateModelWithDefaults

`func NewExternalFormCreateModelWithDefaults() *ExternalFormCreateModel`

NewExternalFormCreateModelWithDefaults instantiates a new ExternalFormCreateModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPossibleValues

`func (o *ExternalFormCreateModel) GetPossibleValues() map[string][]ExternalFormAllowedValueModel`

GetPossibleValues returns the PossibleValues field if non-nil, zero value otherwise.

### GetPossibleValuesOk

`func (o *ExternalFormCreateModel) GetPossibleValuesOk() (*map[string][]ExternalFormAllowedValueModel, bool)`

GetPossibleValuesOk returns a tuple with the PossibleValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPossibleValues

`func (o *ExternalFormCreateModel) SetPossibleValues(v map[string][]ExternalFormAllowedValueModel)`

SetPossibleValues sets PossibleValues field to given value.


### GetFields

`func (o *ExternalFormCreateModel) GetFields() []ExternalFormFieldModel`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *ExternalFormCreateModel) GetFieldsOk() (*[]ExternalFormFieldModel, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *ExternalFormCreateModel) SetFields(v []ExternalFormFieldModel)`

SetFields sets Fields field to given value.


### GetLinks

`func (o *ExternalFormCreateModel) GetLinks() []ExternalFormLinkModel`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ExternalFormCreateModel) GetLinksOk() (*[]ExternalFormLinkModel, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ExternalFormCreateModel) SetLinks(v []ExternalFormLinkModel)`

SetLinks sets Links field to given value.


### GetValues

`func (o *ExternalFormCreateModel) GetValues() map[string]interface{}`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *ExternalFormCreateModel) GetValuesOk() (*map[string]interface{}, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *ExternalFormCreateModel) SetValues(v map[string]interface{})`

SetValues sets Values field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


