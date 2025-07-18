# ExternalFormFieldModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FieldId** | Pointer to **NullableString** |  | [optional] 
**FieldName** | Pointer to **NullableString** |  | [optional] 
**HelpText** | Pointer to **NullableString** |  | [optional] 
**Type** | Pointer to **NullableString** |  | [optional] 
**ArrayValuesType** | Pointer to **NullableString** |  | [optional] 
**DefaultValue** | Pointer to **interface{}** |  | [optional] 
**IsCustomValueAllowed** | **bool** |  | 
**AutoCompleteUrl** | Pointer to **NullableString** |  | [optional] 
**ControlType** | Pointer to **NullableString** |  | [optional] 
**MinLength** | Pointer to **NullableFloat64** |  | [optional] 
**MaxLength** | Pointer to **NullableFloat64** |  | [optional] 
**IsRequired** | Pointer to **NullableBool** |  | [optional] 
**Min** | Pointer to **interface{}** |  | [optional] 
**Max** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewExternalFormFieldModel

`func NewExternalFormFieldModel(isCustomValueAllowed bool, ) *ExternalFormFieldModel`

NewExternalFormFieldModel instantiates a new ExternalFormFieldModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalFormFieldModelWithDefaults

`func NewExternalFormFieldModelWithDefaults() *ExternalFormFieldModel`

NewExternalFormFieldModelWithDefaults instantiates a new ExternalFormFieldModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFieldId

`func (o *ExternalFormFieldModel) GetFieldId() string`

GetFieldId returns the FieldId field if non-nil, zero value otherwise.

### GetFieldIdOk

`func (o *ExternalFormFieldModel) GetFieldIdOk() (*string, bool)`

GetFieldIdOk returns a tuple with the FieldId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldId

`func (o *ExternalFormFieldModel) SetFieldId(v string)`

SetFieldId sets FieldId field to given value.

### HasFieldId

`func (o *ExternalFormFieldModel) HasFieldId() bool`

HasFieldId returns a boolean if a field has been set.

### SetFieldIdNil

`func (o *ExternalFormFieldModel) SetFieldIdNil(b bool)`

 SetFieldIdNil sets the value for FieldId to be an explicit nil

### UnsetFieldId
`func (o *ExternalFormFieldModel) UnsetFieldId()`

UnsetFieldId ensures that no value is present for FieldId, not even an explicit nil
### GetFieldName

`func (o *ExternalFormFieldModel) GetFieldName() string`

GetFieldName returns the FieldName field if non-nil, zero value otherwise.

### GetFieldNameOk

`func (o *ExternalFormFieldModel) GetFieldNameOk() (*string, bool)`

GetFieldNameOk returns a tuple with the FieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldName

`func (o *ExternalFormFieldModel) SetFieldName(v string)`

SetFieldName sets FieldName field to given value.

### HasFieldName

`func (o *ExternalFormFieldModel) HasFieldName() bool`

HasFieldName returns a boolean if a field has been set.

### SetFieldNameNil

`func (o *ExternalFormFieldModel) SetFieldNameNil(b bool)`

 SetFieldNameNil sets the value for FieldName to be an explicit nil

### UnsetFieldName
`func (o *ExternalFormFieldModel) UnsetFieldName()`

UnsetFieldName ensures that no value is present for FieldName, not even an explicit nil
### GetHelpText

`func (o *ExternalFormFieldModel) GetHelpText() string`

GetHelpText returns the HelpText field if non-nil, zero value otherwise.

### GetHelpTextOk

`func (o *ExternalFormFieldModel) GetHelpTextOk() (*string, bool)`

GetHelpTextOk returns a tuple with the HelpText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelpText

`func (o *ExternalFormFieldModel) SetHelpText(v string)`

SetHelpText sets HelpText field to given value.

### HasHelpText

`func (o *ExternalFormFieldModel) HasHelpText() bool`

HasHelpText returns a boolean if a field has been set.

### SetHelpTextNil

`func (o *ExternalFormFieldModel) SetHelpTextNil(b bool)`

 SetHelpTextNil sets the value for HelpText to be an explicit nil

### UnsetHelpText
`func (o *ExternalFormFieldModel) UnsetHelpText()`

UnsetHelpText ensures that no value is present for HelpText, not even an explicit nil
### GetType

`func (o *ExternalFormFieldModel) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ExternalFormFieldModel) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ExternalFormFieldModel) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ExternalFormFieldModel) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *ExternalFormFieldModel) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *ExternalFormFieldModel) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetArrayValuesType

`func (o *ExternalFormFieldModel) GetArrayValuesType() string`

GetArrayValuesType returns the ArrayValuesType field if non-nil, zero value otherwise.

### GetArrayValuesTypeOk

`func (o *ExternalFormFieldModel) GetArrayValuesTypeOk() (*string, bool)`

GetArrayValuesTypeOk returns a tuple with the ArrayValuesType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrayValuesType

`func (o *ExternalFormFieldModel) SetArrayValuesType(v string)`

SetArrayValuesType sets ArrayValuesType field to given value.

### HasArrayValuesType

`func (o *ExternalFormFieldModel) HasArrayValuesType() bool`

HasArrayValuesType returns a boolean if a field has been set.

### SetArrayValuesTypeNil

`func (o *ExternalFormFieldModel) SetArrayValuesTypeNil(b bool)`

 SetArrayValuesTypeNil sets the value for ArrayValuesType to be an explicit nil

### UnsetArrayValuesType
`func (o *ExternalFormFieldModel) UnsetArrayValuesType()`

UnsetArrayValuesType ensures that no value is present for ArrayValuesType, not even an explicit nil
### GetDefaultValue

`func (o *ExternalFormFieldModel) GetDefaultValue() interface{}`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *ExternalFormFieldModel) GetDefaultValueOk() (*interface{}, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *ExternalFormFieldModel) SetDefaultValue(v interface{})`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *ExternalFormFieldModel) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### SetDefaultValueNil

`func (o *ExternalFormFieldModel) SetDefaultValueNil(b bool)`

 SetDefaultValueNil sets the value for DefaultValue to be an explicit nil

### UnsetDefaultValue
`func (o *ExternalFormFieldModel) UnsetDefaultValue()`

UnsetDefaultValue ensures that no value is present for DefaultValue, not even an explicit nil
### GetIsCustomValueAllowed

`func (o *ExternalFormFieldModel) GetIsCustomValueAllowed() bool`

GetIsCustomValueAllowed returns the IsCustomValueAllowed field if non-nil, zero value otherwise.

### GetIsCustomValueAllowedOk

`func (o *ExternalFormFieldModel) GetIsCustomValueAllowedOk() (*bool, bool)`

GetIsCustomValueAllowedOk returns a tuple with the IsCustomValueAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCustomValueAllowed

`func (o *ExternalFormFieldModel) SetIsCustomValueAllowed(v bool)`

SetIsCustomValueAllowed sets IsCustomValueAllowed field to given value.


### GetAutoCompleteUrl

`func (o *ExternalFormFieldModel) GetAutoCompleteUrl() string`

GetAutoCompleteUrl returns the AutoCompleteUrl field if non-nil, zero value otherwise.

### GetAutoCompleteUrlOk

`func (o *ExternalFormFieldModel) GetAutoCompleteUrlOk() (*string, bool)`

GetAutoCompleteUrlOk returns a tuple with the AutoCompleteUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoCompleteUrl

`func (o *ExternalFormFieldModel) SetAutoCompleteUrl(v string)`

SetAutoCompleteUrl sets AutoCompleteUrl field to given value.

### HasAutoCompleteUrl

`func (o *ExternalFormFieldModel) HasAutoCompleteUrl() bool`

HasAutoCompleteUrl returns a boolean if a field has been set.

### SetAutoCompleteUrlNil

`func (o *ExternalFormFieldModel) SetAutoCompleteUrlNil(b bool)`

 SetAutoCompleteUrlNil sets the value for AutoCompleteUrl to be an explicit nil

### UnsetAutoCompleteUrl
`func (o *ExternalFormFieldModel) UnsetAutoCompleteUrl()`

UnsetAutoCompleteUrl ensures that no value is present for AutoCompleteUrl, not even an explicit nil
### GetControlType

`func (o *ExternalFormFieldModel) GetControlType() string`

GetControlType returns the ControlType field if non-nil, zero value otherwise.

### GetControlTypeOk

`func (o *ExternalFormFieldModel) GetControlTypeOk() (*string, bool)`

GetControlTypeOk returns a tuple with the ControlType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlType

`func (o *ExternalFormFieldModel) SetControlType(v string)`

SetControlType sets ControlType field to given value.

### HasControlType

`func (o *ExternalFormFieldModel) HasControlType() bool`

HasControlType returns a boolean if a field has been set.

### SetControlTypeNil

`func (o *ExternalFormFieldModel) SetControlTypeNil(b bool)`

 SetControlTypeNil sets the value for ControlType to be an explicit nil

### UnsetControlType
`func (o *ExternalFormFieldModel) UnsetControlType()`

UnsetControlType ensures that no value is present for ControlType, not even an explicit nil
### GetMinLength

`func (o *ExternalFormFieldModel) GetMinLength() float64`

GetMinLength returns the MinLength field if non-nil, zero value otherwise.

### GetMinLengthOk

`func (o *ExternalFormFieldModel) GetMinLengthOk() (*float64, bool)`

GetMinLengthOk returns a tuple with the MinLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinLength

`func (o *ExternalFormFieldModel) SetMinLength(v float64)`

SetMinLength sets MinLength field to given value.

### HasMinLength

`func (o *ExternalFormFieldModel) HasMinLength() bool`

HasMinLength returns a boolean if a field has been set.

### SetMinLengthNil

`func (o *ExternalFormFieldModel) SetMinLengthNil(b bool)`

 SetMinLengthNil sets the value for MinLength to be an explicit nil

### UnsetMinLength
`func (o *ExternalFormFieldModel) UnsetMinLength()`

UnsetMinLength ensures that no value is present for MinLength, not even an explicit nil
### GetMaxLength

`func (o *ExternalFormFieldModel) GetMaxLength() float64`

GetMaxLength returns the MaxLength field if non-nil, zero value otherwise.

### GetMaxLengthOk

`func (o *ExternalFormFieldModel) GetMaxLengthOk() (*float64, bool)`

GetMaxLengthOk returns a tuple with the MaxLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLength

`func (o *ExternalFormFieldModel) SetMaxLength(v float64)`

SetMaxLength sets MaxLength field to given value.

### HasMaxLength

`func (o *ExternalFormFieldModel) HasMaxLength() bool`

HasMaxLength returns a boolean if a field has been set.

### SetMaxLengthNil

`func (o *ExternalFormFieldModel) SetMaxLengthNil(b bool)`

 SetMaxLengthNil sets the value for MaxLength to be an explicit nil

### UnsetMaxLength
`func (o *ExternalFormFieldModel) UnsetMaxLength()`

UnsetMaxLength ensures that no value is present for MaxLength, not even an explicit nil
### GetIsRequired

`func (o *ExternalFormFieldModel) GetIsRequired() bool`

GetIsRequired returns the IsRequired field if non-nil, zero value otherwise.

### GetIsRequiredOk

`func (o *ExternalFormFieldModel) GetIsRequiredOk() (*bool, bool)`

GetIsRequiredOk returns a tuple with the IsRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRequired

`func (o *ExternalFormFieldModel) SetIsRequired(v bool)`

SetIsRequired sets IsRequired field to given value.

### HasIsRequired

`func (o *ExternalFormFieldModel) HasIsRequired() bool`

HasIsRequired returns a boolean if a field has been set.

### SetIsRequiredNil

`func (o *ExternalFormFieldModel) SetIsRequiredNil(b bool)`

 SetIsRequiredNil sets the value for IsRequired to be an explicit nil

### UnsetIsRequired
`func (o *ExternalFormFieldModel) UnsetIsRequired()`

UnsetIsRequired ensures that no value is present for IsRequired, not even an explicit nil
### GetMin

`func (o *ExternalFormFieldModel) GetMin() interface{}`

GetMin returns the Min field if non-nil, zero value otherwise.

### GetMinOk

`func (o *ExternalFormFieldModel) GetMinOk() (*interface{}, bool)`

GetMinOk returns a tuple with the Min field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMin

`func (o *ExternalFormFieldModel) SetMin(v interface{})`

SetMin sets Min field to given value.

### HasMin

`func (o *ExternalFormFieldModel) HasMin() bool`

HasMin returns a boolean if a field has been set.

### SetMinNil

`func (o *ExternalFormFieldModel) SetMinNil(b bool)`

 SetMinNil sets the value for Min to be an explicit nil

### UnsetMin
`func (o *ExternalFormFieldModel) UnsetMin()`

UnsetMin ensures that no value is present for Min, not even an explicit nil
### GetMax

`func (o *ExternalFormFieldModel) GetMax() interface{}`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *ExternalFormFieldModel) GetMaxOk() (*interface{}, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *ExternalFormFieldModel) SetMax(v interface{})`

SetMax sets Max field to given value.

### HasMax

`func (o *ExternalFormFieldModel) HasMax() bool`

HasMax returns a boolean if a field has been set.

### SetMaxNil

`func (o *ExternalFormFieldModel) SetMaxNil(b bool)`

 SetMaxNil sets the value for Max to be an explicit nil

### UnsetMax
`func (o *ExternalFormFieldModel) UnsetMax()`

UnsetMax ensures that no value is present for Max, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


