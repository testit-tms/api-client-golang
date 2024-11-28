# AutoTestNamespaceCountApiModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**Count** | **int64** |  | 
**Classes** | [**[]AutoTestClassCountApiModel**](AutoTestClassCountApiModel.md) |  | 

## Methods

### NewAutoTestNamespaceCountApiModel

`func NewAutoTestNamespaceCountApiModel(count int64, classes []AutoTestClassCountApiModel, ) *AutoTestNamespaceCountApiModel`

NewAutoTestNamespaceCountApiModel instantiates a new AutoTestNamespaceCountApiModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoTestNamespaceCountApiModelWithDefaults

`func NewAutoTestNamespaceCountApiModelWithDefaults() *AutoTestNamespaceCountApiModel`

NewAutoTestNamespaceCountApiModelWithDefaults instantiates a new AutoTestNamespaceCountApiModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AutoTestNamespaceCountApiModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AutoTestNamespaceCountApiModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AutoTestNamespaceCountApiModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AutoTestNamespaceCountApiModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *AutoTestNamespaceCountApiModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AutoTestNamespaceCountApiModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetCount

`func (o *AutoTestNamespaceCountApiModel) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *AutoTestNamespaceCountApiModel) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *AutoTestNamespaceCountApiModel) SetCount(v int64)`

SetCount sets Count field to given value.


### GetClasses

`func (o *AutoTestNamespaceCountApiModel) GetClasses() []AutoTestClassCountApiModel`

GetClasses returns the Classes field if non-nil, zero value otherwise.

### GetClassesOk

`func (o *AutoTestNamespaceCountApiModel) GetClassesOk() (*[]AutoTestClassCountApiModel, bool)`

GetClassesOk returns a tuple with the Classes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClasses

`func (o *AutoTestNamespaceCountApiModel) SetClasses(v []AutoTestClassCountApiModel)`

SetClasses sets Classes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


