# BackgroundJobFilterModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Types** | Pointer to [**[]BackgroundJobType**](BackgroundJobType.md) |  | [optional] 
**States** | Pointer to [**[]BackgroundJobState**](BackgroundJobState.md) |  | [optional] 
**IsDeleted** | Pointer to **NullableBool** |  | [optional] 
**StartDate** | Pointer to [**NullableDateTimeRangeSelectorModel**](DateTimeRangeSelectorModel.md) |  | [optional] 
**EndDate** | Pointer to [**NullableDateTimeRangeSelectorModel**](DateTimeRangeSelectorModel.md) |  | [optional] 

## Methods

### NewBackgroundJobFilterModel

`func NewBackgroundJobFilterModel() *BackgroundJobFilterModel`

NewBackgroundJobFilterModel instantiates a new BackgroundJobFilterModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackgroundJobFilterModelWithDefaults

`func NewBackgroundJobFilterModelWithDefaults() *BackgroundJobFilterModel`

NewBackgroundJobFilterModelWithDefaults instantiates a new BackgroundJobFilterModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTypes

`func (o *BackgroundJobFilterModel) GetTypes() []BackgroundJobType`

GetTypes returns the Types field if non-nil, zero value otherwise.

### GetTypesOk

`func (o *BackgroundJobFilterModel) GetTypesOk() (*[]BackgroundJobType, bool)`

GetTypesOk returns a tuple with the Types field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypes

`func (o *BackgroundJobFilterModel) SetTypes(v []BackgroundJobType)`

SetTypes sets Types field to given value.

### HasTypes

`func (o *BackgroundJobFilterModel) HasTypes() bool`

HasTypes returns a boolean if a field has been set.

### SetTypesNil

`func (o *BackgroundJobFilterModel) SetTypesNil(b bool)`

 SetTypesNil sets the value for Types to be an explicit nil

### UnsetTypes
`func (o *BackgroundJobFilterModel) UnsetTypes()`

UnsetTypes ensures that no value is present for Types, not even an explicit nil
### GetStates

`func (o *BackgroundJobFilterModel) GetStates() []BackgroundJobState`

GetStates returns the States field if non-nil, zero value otherwise.

### GetStatesOk

`func (o *BackgroundJobFilterModel) GetStatesOk() (*[]BackgroundJobState, bool)`

GetStatesOk returns a tuple with the States field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStates

`func (o *BackgroundJobFilterModel) SetStates(v []BackgroundJobState)`

SetStates sets States field to given value.

### HasStates

`func (o *BackgroundJobFilterModel) HasStates() bool`

HasStates returns a boolean if a field has been set.

### SetStatesNil

`func (o *BackgroundJobFilterModel) SetStatesNil(b bool)`

 SetStatesNil sets the value for States to be an explicit nil

### UnsetStates
`func (o *BackgroundJobFilterModel) UnsetStates()`

UnsetStates ensures that no value is present for States, not even an explicit nil
### GetIsDeleted

`func (o *BackgroundJobFilterModel) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *BackgroundJobFilterModel) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *BackgroundJobFilterModel) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *BackgroundJobFilterModel) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### SetIsDeletedNil

`func (o *BackgroundJobFilterModel) SetIsDeletedNil(b bool)`

 SetIsDeletedNil sets the value for IsDeleted to be an explicit nil

### UnsetIsDeleted
`func (o *BackgroundJobFilterModel) UnsetIsDeleted()`

UnsetIsDeleted ensures that no value is present for IsDeleted, not even an explicit nil
### GetStartDate

`func (o *BackgroundJobFilterModel) GetStartDate() DateTimeRangeSelectorModel`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *BackgroundJobFilterModel) GetStartDateOk() (*DateTimeRangeSelectorModel, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *BackgroundJobFilterModel) SetStartDate(v DateTimeRangeSelectorModel)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *BackgroundJobFilterModel) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### SetStartDateNil

`func (o *BackgroundJobFilterModel) SetStartDateNil(b bool)`

 SetStartDateNil sets the value for StartDate to be an explicit nil

### UnsetStartDate
`func (o *BackgroundJobFilterModel) UnsetStartDate()`

UnsetStartDate ensures that no value is present for StartDate, not even an explicit nil
### GetEndDate

`func (o *BackgroundJobFilterModel) GetEndDate() DateTimeRangeSelectorModel`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *BackgroundJobFilterModel) GetEndDateOk() (*DateTimeRangeSelectorModel, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *BackgroundJobFilterModel) SetEndDate(v DateTimeRangeSelectorModel)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *BackgroundJobFilterModel) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### SetEndDateNil

`func (o *BackgroundJobFilterModel) SetEndDateNil(b bool)`

 SetEndDateNil sets the value for EndDate to be an explicit nil

### UnsetEndDate
`func (o *BackgroundJobFilterModel) UnsetEndDate()`

UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


