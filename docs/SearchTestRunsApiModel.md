# SearchTestRunsApiModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**States** | Pointer to [**[]TestRunState**](TestRunState.md) |  | [optional] 
**StatusCodes** | Pointer to **[]string** |  | [optional] 
**StartedDate** | Pointer to [**NullableDateTimeRangeSelectorModel**](DateTimeRangeSelectorModel.md) |  | [optional] 
**CompletedDate** | Pointer to [**NullableDateTimeRangeSelectorModel**](DateTimeRangeSelectorModel.md) |  | [optional] 
**CreatedByIds** | Pointer to **[]string** |  | [optional] 
**ModifiedByIds** | Pointer to **[]string** |  | [optional] 

## Methods

### NewSearchTestRunsApiModel

`func NewSearchTestRunsApiModel() *SearchTestRunsApiModel`

NewSearchTestRunsApiModel instantiates a new SearchTestRunsApiModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchTestRunsApiModelWithDefaults

`func NewSearchTestRunsApiModelWithDefaults() *SearchTestRunsApiModel`

NewSearchTestRunsApiModelWithDefaults instantiates a new SearchTestRunsApiModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SearchTestRunsApiModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SearchTestRunsApiModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SearchTestRunsApiModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SearchTestRunsApiModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *SearchTestRunsApiModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *SearchTestRunsApiModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetStates

`func (o *SearchTestRunsApiModel) GetStates() []TestRunState`

GetStates returns the States field if non-nil, zero value otherwise.

### GetStatesOk

`func (o *SearchTestRunsApiModel) GetStatesOk() (*[]TestRunState, bool)`

GetStatesOk returns a tuple with the States field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStates

`func (o *SearchTestRunsApiModel) SetStates(v []TestRunState)`

SetStates sets States field to given value.

### HasStates

`func (o *SearchTestRunsApiModel) HasStates() bool`

HasStates returns a boolean if a field has been set.

### SetStatesNil

`func (o *SearchTestRunsApiModel) SetStatesNil(b bool)`

 SetStatesNil sets the value for States to be an explicit nil

### UnsetStates
`func (o *SearchTestRunsApiModel) UnsetStates()`

UnsetStates ensures that no value is present for States, not even an explicit nil
### GetStatusCodes

`func (o *SearchTestRunsApiModel) GetStatusCodes() []string`

GetStatusCodes returns the StatusCodes field if non-nil, zero value otherwise.

### GetStatusCodesOk

`func (o *SearchTestRunsApiModel) GetStatusCodesOk() (*[]string, bool)`

GetStatusCodesOk returns a tuple with the StatusCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCodes

`func (o *SearchTestRunsApiModel) SetStatusCodes(v []string)`

SetStatusCodes sets StatusCodes field to given value.

### HasStatusCodes

`func (o *SearchTestRunsApiModel) HasStatusCodes() bool`

HasStatusCodes returns a boolean if a field has been set.

### SetStatusCodesNil

`func (o *SearchTestRunsApiModel) SetStatusCodesNil(b bool)`

 SetStatusCodesNil sets the value for StatusCodes to be an explicit nil

### UnsetStatusCodes
`func (o *SearchTestRunsApiModel) UnsetStatusCodes()`

UnsetStatusCodes ensures that no value is present for StatusCodes, not even an explicit nil
### GetStartedDate

`func (o *SearchTestRunsApiModel) GetStartedDate() DateTimeRangeSelectorModel`

GetStartedDate returns the StartedDate field if non-nil, zero value otherwise.

### GetStartedDateOk

`func (o *SearchTestRunsApiModel) GetStartedDateOk() (*DateTimeRangeSelectorModel, bool)`

GetStartedDateOk returns a tuple with the StartedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedDate

`func (o *SearchTestRunsApiModel) SetStartedDate(v DateTimeRangeSelectorModel)`

SetStartedDate sets StartedDate field to given value.

### HasStartedDate

`func (o *SearchTestRunsApiModel) HasStartedDate() bool`

HasStartedDate returns a boolean if a field has been set.

### SetStartedDateNil

`func (o *SearchTestRunsApiModel) SetStartedDateNil(b bool)`

 SetStartedDateNil sets the value for StartedDate to be an explicit nil

### UnsetStartedDate
`func (o *SearchTestRunsApiModel) UnsetStartedDate()`

UnsetStartedDate ensures that no value is present for StartedDate, not even an explicit nil
### GetCompletedDate

`func (o *SearchTestRunsApiModel) GetCompletedDate() DateTimeRangeSelectorModel`

GetCompletedDate returns the CompletedDate field if non-nil, zero value otherwise.

### GetCompletedDateOk

`func (o *SearchTestRunsApiModel) GetCompletedDateOk() (*DateTimeRangeSelectorModel, bool)`

GetCompletedDateOk returns a tuple with the CompletedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedDate

`func (o *SearchTestRunsApiModel) SetCompletedDate(v DateTimeRangeSelectorModel)`

SetCompletedDate sets CompletedDate field to given value.

### HasCompletedDate

`func (o *SearchTestRunsApiModel) HasCompletedDate() bool`

HasCompletedDate returns a boolean if a field has been set.

### SetCompletedDateNil

`func (o *SearchTestRunsApiModel) SetCompletedDateNil(b bool)`

 SetCompletedDateNil sets the value for CompletedDate to be an explicit nil

### UnsetCompletedDate
`func (o *SearchTestRunsApiModel) UnsetCompletedDate()`

UnsetCompletedDate ensures that no value is present for CompletedDate, not even an explicit nil
### GetCreatedByIds

`func (o *SearchTestRunsApiModel) GetCreatedByIds() []string`

GetCreatedByIds returns the CreatedByIds field if non-nil, zero value otherwise.

### GetCreatedByIdsOk

`func (o *SearchTestRunsApiModel) GetCreatedByIdsOk() (*[]string, bool)`

GetCreatedByIdsOk returns a tuple with the CreatedByIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByIds

`func (o *SearchTestRunsApiModel) SetCreatedByIds(v []string)`

SetCreatedByIds sets CreatedByIds field to given value.

### HasCreatedByIds

`func (o *SearchTestRunsApiModel) HasCreatedByIds() bool`

HasCreatedByIds returns a boolean if a field has been set.

### SetCreatedByIdsNil

`func (o *SearchTestRunsApiModel) SetCreatedByIdsNil(b bool)`

 SetCreatedByIdsNil sets the value for CreatedByIds to be an explicit nil

### UnsetCreatedByIds
`func (o *SearchTestRunsApiModel) UnsetCreatedByIds()`

UnsetCreatedByIds ensures that no value is present for CreatedByIds, not even an explicit nil
### GetModifiedByIds

`func (o *SearchTestRunsApiModel) GetModifiedByIds() []string`

GetModifiedByIds returns the ModifiedByIds field if non-nil, zero value otherwise.

### GetModifiedByIdsOk

`func (o *SearchTestRunsApiModel) GetModifiedByIdsOk() (*[]string, bool)`

GetModifiedByIdsOk returns a tuple with the ModifiedByIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedByIds

`func (o *SearchTestRunsApiModel) SetModifiedByIds(v []string)`

SetModifiedByIds sets ModifiedByIds field to given value.

### HasModifiedByIds

`func (o *SearchTestRunsApiModel) HasModifiedByIds() bool`

HasModifiedByIds returns a boolean if a field has been set.

### SetModifiedByIdsNil

`func (o *SearchTestRunsApiModel) SetModifiedByIdsNil(b bool)`

 SetModifiedByIdsNil sets the value for ModifiedByIds to be an explicit nil

### UnsetModifiedByIds
`func (o *SearchTestRunsApiModel) UnsetModifiedByIds()`

UnsetModifiedByIds ensures that no value is present for ModifiedByIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


