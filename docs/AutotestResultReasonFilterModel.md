# AutotestResultReasonFilterModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FailureCategories** | Pointer to [**[]AvailableFailureCategory**](AvailableFailureCategory.md) |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**RegexCount** | Pointer to [**NullableInt32RangeSelectorModel**](Int32RangeSelectorModel.md) |  | [optional] 
**IsDeleted** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewAutotestResultReasonFilterModel

`func NewAutotestResultReasonFilterModel() *AutotestResultReasonFilterModel`

NewAutotestResultReasonFilterModel instantiates a new AutotestResultReasonFilterModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutotestResultReasonFilterModelWithDefaults

`func NewAutotestResultReasonFilterModelWithDefaults() *AutotestResultReasonFilterModel`

NewAutotestResultReasonFilterModelWithDefaults instantiates a new AutotestResultReasonFilterModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFailureCategories

`func (o *AutotestResultReasonFilterModel) GetFailureCategories() []AvailableFailureCategory`

GetFailureCategories returns the FailureCategories field if non-nil, zero value otherwise.

### GetFailureCategoriesOk

`func (o *AutotestResultReasonFilterModel) GetFailureCategoriesOk() (*[]AvailableFailureCategory, bool)`

GetFailureCategoriesOk returns a tuple with the FailureCategories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureCategories

`func (o *AutotestResultReasonFilterModel) SetFailureCategories(v []AvailableFailureCategory)`

SetFailureCategories sets FailureCategories field to given value.

### HasFailureCategories

`func (o *AutotestResultReasonFilterModel) HasFailureCategories() bool`

HasFailureCategories returns a boolean if a field has been set.

### SetFailureCategoriesNil

`func (o *AutotestResultReasonFilterModel) SetFailureCategoriesNil(b bool)`

 SetFailureCategoriesNil sets the value for FailureCategories to be an explicit nil

### UnsetFailureCategories
`func (o *AutotestResultReasonFilterModel) UnsetFailureCategories()`

UnsetFailureCategories ensures that no value is present for FailureCategories, not even an explicit nil
### GetName

`func (o *AutotestResultReasonFilterModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AutotestResultReasonFilterModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AutotestResultReasonFilterModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AutotestResultReasonFilterModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *AutotestResultReasonFilterModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AutotestResultReasonFilterModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetRegexCount

`func (o *AutotestResultReasonFilterModel) GetRegexCount() Int32RangeSelectorModel`

GetRegexCount returns the RegexCount field if non-nil, zero value otherwise.

### GetRegexCountOk

`func (o *AutotestResultReasonFilterModel) GetRegexCountOk() (*Int32RangeSelectorModel, bool)`

GetRegexCountOk returns a tuple with the RegexCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegexCount

`func (o *AutotestResultReasonFilterModel) SetRegexCount(v Int32RangeSelectorModel)`

SetRegexCount sets RegexCount field to given value.

### HasRegexCount

`func (o *AutotestResultReasonFilterModel) HasRegexCount() bool`

HasRegexCount returns a boolean if a field has been set.

### SetRegexCountNil

`func (o *AutotestResultReasonFilterModel) SetRegexCountNil(b bool)`

 SetRegexCountNil sets the value for RegexCount to be an explicit nil

### UnsetRegexCount
`func (o *AutotestResultReasonFilterModel) UnsetRegexCount()`

UnsetRegexCount ensures that no value is present for RegexCount, not even an explicit nil
### GetIsDeleted

`func (o *AutotestResultReasonFilterModel) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *AutotestResultReasonFilterModel) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *AutotestResultReasonFilterModel) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *AutotestResultReasonFilterModel) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### SetIsDeletedNil

`func (o *AutotestResultReasonFilterModel) SetIsDeletedNil(b bool)`

 SetIsDeletedNil sets the value for IsDeleted to be an explicit nil

### UnsetIsDeleted
`func (o *AutotestResultReasonFilterModel) UnsetIsDeleted()`

UnsetIsDeleted ensures that no value is present for IsDeleted, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


