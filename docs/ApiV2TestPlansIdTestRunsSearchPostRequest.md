# ApiV2TestPlansIdTestRunsSearchPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**States** | Pointer to [**[]TestRunState**](TestRunState.md) |  | [optional] 
**StartedDate** | Pointer to [**NullableDateTimeRangeSelectorModel**](DateTimeRangeSelectorModel.md) |  | [optional] 
**CompletedDate** | Pointer to [**NullableDateTimeRangeSelectorModel**](DateTimeRangeSelectorModel.md) |  | [optional] 
**CreatedByIds** | Pointer to **[]string** |  | [optional] 
**ModifiedByIds** | Pointer to **[]string** |  | [optional] 

## Methods

### NewApiV2TestPlansIdTestRunsSearchPostRequest

`func NewApiV2TestPlansIdTestRunsSearchPostRequest() *ApiV2TestPlansIdTestRunsSearchPostRequest`

NewApiV2TestPlansIdTestRunsSearchPostRequest instantiates a new ApiV2TestPlansIdTestRunsSearchPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV2TestPlansIdTestRunsSearchPostRequestWithDefaults

`func NewApiV2TestPlansIdTestRunsSearchPostRequestWithDefaults() *ApiV2TestPlansIdTestRunsSearchPostRequest`

NewApiV2TestPlansIdTestRunsSearchPostRequestWithDefaults instantiates a new ApiV2TestPlansIdTestRunsSearchPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ApiV2TestPlansIdTestRunsSearchPostRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiV2TestPlansIdTestRunsSearchPostRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiV2TestPlansIdTestRunsSearchPostRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApiV2TestPlansIdTestRunsSearchPostRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ApiV2TestPlansIdTestRunsSearchPostRequest) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ApiV2TestPlansIdTestRunsSearchPostRequest) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetStates

`func (o *ApiV2TestPlansIdTestRunsSearchPostRequest) GetStates() []TestRunState`

GetStates returns the States field if non-nil, zero value otherwise.

### GetStatesOk

`func (o *ApiV2TestPlansIdTestRunsSearchPostRequest) GetStatesOk() (*[]TestRunState, bool)`

GetStatesOk returns a tuple with the States field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStates

`func (o *ApiV2TestPlansIdTestRunsSearchPostRequest) SetStates(v []TestRunState)`

SetStates sets States field to given value.

### HasStates

`func (o *ApiV2TestPlansIdTestRunsSearchPostRequest) HasStates() bool`

HasStates returns a boolean if a field has been set.

### SetStatesNil

`func (o *ApiV2TestPlansIdTestRunsSearchPostRequest) SetStatesNil(b bool)`

 SetStatesNil sets the value for States to be an explicit nil

### UnsetStates
`func (o *ApiV2TestPlansIdTestRunsSearchPostRequest) UnsetStates()`

UnsetStates ensures that no value is present for States, not even an explicit nil
### GetStartedDate

`func (o *ApiV2TestPlansIdTestRunsSearchPostRequest) GetStartedDate() DateTimeRangeSelectorModel`

GetStartedDate returns the StartedDate field if non-nil, zero value otherwise.

### GetStartedDateOk

`func (o *ApiV2TestPlansIdTestRunsSearchPostRequest) GetStartedDateOk() (*DateTimeRangeSelectorModel, bool)`

GetStartedDateOk returns a tuple with the StartedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedDate

`func (o *ApiV2TestPlansIdTestRunsSearchPostRequest) SetStartedDate(v DateTimeRangeSelectorModel)`

SetStartedDate sets StartedDate field to given value.

### HasStartedDate

`func (o *ApiV2TestPlansIdTestRunsSearchPostRequest) HasStartedDate() bool`

HasStartedDate returns a boolean if a field has been set.

### SetStartedDateNil

`func (o *ApiV2TestPlansIdTestRunsSearchPostRequest) SetStartedDateNil(b bool)`

 SetStartedDateNil sets the value for StartedDate to be an explicit nil

### UnsetStartedDate
`func (o *ApiV2TestPlansIdTestRunsSearchPostRequest) UnsetStartedDate()`

UnsetStartedDate ensures that no value is present for StartedDate, not even an explicit nil
### GetCompletedDate

`func (o *ApiV2TestPlansIdTestRunsSearchPostRequest) GetCompletedDate() DateTimeRangeSelectorModel`

GetCompletedDate returns the CompletedDate field if non-nil, zero value otherwise.

### GetCompletedDateOk

`func (o *ApiV2TestPlansIdTestRunsSearchPostRequest) GetCompletedDateOk() (*DateTimeRangeSelectorModel, bool)`

GetCompletedDateOk returns a tuple with the CompletedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedDate

`func (o *ApiV2TestPlansIdTestRunsSearchPostRequest) SetCompletedDate(v DateTimeRangeSelectorModel)`

SetCompletedDate sets CompletedDate field to given value.

### HasCompletedDate

`func (o *ApiV2TestPlansIdTestRunsSearchPostRequest) HasCompletedDate() bool`

HasCompletedDate returns a boolean if a field has been set.

### SetCompletedDateNil

`func (o *ApiV2TestPlansIdTestRunsSearchPostRequest) SetCompletedDateNil(b bool)`

 SetCompletedDateNil sets the value for CompletedDate to be an explicit nil

### UnsetCompletedDate
`func (o *ApiV2TestPlansIdTestRunsSearchPostRequest) UnsetCompletedDate()`

UnsetCompletedDate ensures that no value is present for CompletedDate, not even an explicit nil
### GetCreatedByIds

`func (o *ApiV2TestPlansIdTestRunsSearchPostRequest) GetCreatedByIds() []string`

GetCreatedByIds returns the CreatedByIds field if non-nil, zero value otherwise.

### GetCreatedByIdsOk

`func (o *ApiV2TestPlansIdTestRunsSearchPostRequest) GetCreatedByIdsOk() (*[]string, bool)`

GetCreatedByIdsOk returns a tuple with the CreatedByIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByIds

`func (o *ApiV2TestPlansIdTestRunsSearchPostRequest) SetCreatedByIds(v []string)`

SetCreatedByIds sets CreatedByIds field to given value.

### HasCreatedByIds

`func (o *ApiV2TestPlansIdTestRunsSearchPostRequest) HasCreatedByIds() bool`

HasCreatedByIds returns a boolean if a field has been set.

### SetCreatedByIdsNil

`func (o *ApiV2TestPlansIdTestRunsSearchPostRequest) SetCreatedByIdsNil(b bool)`

 SetCreatedByIdsNil sets the value for CreatedByIds to be an explicit nil

### UnsetCreatedByIds
`func (o *ApiV2TestPlansIdTestRunsSearchPostRequest) UnsetCreatedByIds()`

UnsetCreatedByIds ensures that no value is present for CreatedByIds, not even an explicit nil
### GetModifiedByIds

`func (o *ApiV2TestPlansIdTestRunsSearchPostRequest) GetModifiedByIds() []string`

GetModifiedByIds returns the ModifiedByIds field if non-nil, zero value otherwise.

### GetModifiedByIdsOk

`func (o *ApiV2TestPlansIdTestRunsSearchPostRequest) GetModifiedByIdsOk() (*[]string, bool)`

GetModifiedByIdsOk returns a tuple with the ModifiedByIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedByIds

`func (o *ApiV2TestPlansIdTestRunsSearchPostRequest) SetModifiedByIds(v []string)`

SetModifiedByIds sets ModifiedByIds field to given value.

### HasModifiedByIds

`func (o *ApiV2TestPlansIdTestRunsSearchPostRequest) HasModifiedByIds() bool`

HasModifiedByIds returns a boolean if a field has been set.

### SetModifiedByIdsNil

`func (o *ApiV2TestPlansIdTestRunsSearchPostRequest) SetModifiedByIdsNil(b bool)`

 SetModifiedByIdsNil sets the value for ModifiedByIds to be an explicit nil

### UnsetModifiedByIds
`func (o *ApiV2TestPlansIdTestRunsSearchPostRequest) UnsetModifiedByIds()`

UnsetModifiedByIds ensures that no value is present for ModifiedByIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


