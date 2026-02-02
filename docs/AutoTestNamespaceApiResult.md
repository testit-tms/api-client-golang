# AutoTestNamespaceApiResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | Autotest namespace | [optional] 
**Classes** | **[]string** | Autotest classnames | 

## Methods

### NewAutoTestNamespaceApiResult

`func NewAutoTestNamespaceApiResult(classes []string, ) *AutoTestNamespaceApiResult`

NewAutoTestNamespaceApiResult instantiates a new AutoTestNamespaceApiResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoTestNamespaceApiResultWithDefaults

`func NewAutoTestNamespaceApiResultWithDefaults() *AutoTestNamespaceApiResult`

NewAutoTestNamespaceApiResultWithDefaults instantiates a new AutoTestNamespaceApiResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AutoTestNamespaceApiResult) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AutoTestNamespaceApiResult) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AutoTestNamespaceApiResult) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AutoTestNamespaceApiResult) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *AutoTestNamespaceApiResult) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AutoTestNamespaceApiResult) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetClasses

`func (o *AutoTestNamespaceApiResult) GetClasses() []string`

GetClasses returns the Classes field if non-nil, zero value otherwise.

### GetClassesOk

`func (o *AutoTestNamespaceApiResult) GetClassesOk() (*[]string, bool)`

GetClassesOk returns a tuple with the Classes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClasses

`func (o *AutoTestNamespaceApiResult) SetClasses(v []string)`

SetClasses sets Classes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


