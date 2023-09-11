# ApiV2BackgroundJobsSearchPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Types** | Pointer to [**[]BackgroundJobType**](BackgroundJobType.md) |  | [optional] 
**States** | Pointer to [**[]BackgroundJobState**](BackgroundJobState.md) |  | [optional] 
**IsDeleted** | Pointer to **NullableBool** |  | [optional] 
**StartDate** | Pointer to [**NullableDateTimeRangeSelectorModel**](DateTimeRangeSelectorModel.md) |  | [optional] 
**EndDate** | Pointer to [**NullableDateTimeRangeSelectorModel**](DateTimeRangeSelectorModel.md) |  | [optional] 

## Methods

### NewApiV2BackgroundJobsSearchPostRequest

`func NewApiV2BackgroundJobsSearchPostRequest() *ApiV2BackgroundJobsSearchPostRequest`

NewApiV2BackgroundJobsSearchPostRequest instantiates a new ApiV2BackgroundJobsSearchPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV2BackgroundJobsSearchPostRequestWithDefaults

`func NewApiV2BackgroundJobsSearchPostRequestWithDefaults() *ApiV2BackgroundJobsSearchPostRequest`

NewApiV2BackgroundJobsSearchPostRequestWithDefaults instantiates a new ApiV2BackgroundJobsSearchPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTypes

`func (o *ApiV2BackgroundJobsSearchPostRequest) GetTypes() []BackgroundJobType`

GetTypes returns the Types field if non-nil, zero value otherwise.

### GetTypesOk

`func (o *ApiV2BackgroundJobsSearchPostRequest) GetTypesOk() (*[]BackgroundJobType, bool)`

GetTypesOk returns a tuple with the Types field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypes

`func (o *ApiV2BackgroundJobsSearchPostRequest) SetTypes(v []BackgroundJobType)`

SetTypes sets Types field to given value.

### HasTypes

`func (o *ApiV2BackgroundJobsSearchPostRequest) HasTypes() bool`

HasTypes returns a boolean if a field has been set.

### SetTypesNil

`func (o *ApiV2BackgroundJobsSearchPostRequest) SetTypesNil(b bool)`

 SetTypesNil sets the value for Types to be an explicit nil

### UnsetTypes
`func (o *ApiV2BackgroundJobsSearchPostRequest) UnsetTypes()`

UnsetTypes ensures that no value is present for Types, not even an explicit nil
### GetStates

`func (o *ApiV2BackgroundJobsSearchPostRequest) GetStates() []BackgroundJobState`

GetStates returns the States field if non-nil, zero value otherwise.

### GetStatesOk

`func (o *ApiV2BackgroundJobsSearchPostRequest) GetStatesOk() (*[]BackgroundJobState, bool)`

GetStatesOk returns a tuple with the States field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStates

`func (o *ApiV2BackgroundJobsSearchPostRequest) SetStates(v []BackgroundJobState)`

SetStates sets States field to given value.

### HasStates

`func (o *ApiV2BackgroundJobsSearchPostRequest) HasStates() bool`

HasStates returns a boolean if a field has been set.

### SetStatesNil

`func (o *ApiV2BackgroundJobsSearchPostRequest) SetStatesNil(b bool)`

 SetStatesNil sets the value for States to be an explicit nil

### UnsetStates
`func (o *ApiV2BackgroundJobsSearchPostRequest) UnsetStates()`

UnsetStates ensures that no value is present for States, not even an explicit nil
### GetIsDeleted

`func (o *ApiV2BackgroundJobsSearchPostRequest) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *ApiV2BackgroundJobsSearchPostRequest) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *ApiV2BackgroundJobsSearchPostRequest) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *ApiV2BackgroundJobsSearchPostRequest) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### SetIsDeletedNil

`func (o *ApiV2BackgroundJobsSearchPostRequest) SetIsDeletedNil(b bool)`

 SetIsDeletedNil sets the value for IsDeleted to be an explicit nil

### UnsetIsDeleted
`func (o *ApiV2BackgroundJobsSearchPostRequest) UnsetIsDeleted()`

UnsetIsDeleted ensures that no value is present for IsDeleted, not even an explicit nil
### GetStartDate

`func (o *ApiV2BackgroundJobsSearchPostRequest) GetStartDate() DateTimeRangeSelectorModel`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *ApiV2BackgroundJobsSearchPostRequest) GetStartDateOk() (*DateTimeRangeSelectorModel, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *ApiV2BackgroundJobsSearchPostRequest) SetStartDate(v DateTimeRangeSelectorModel)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *ApiV2BackgroundJobsSearchPostRequest) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### SetStartDateNil

`func (o *ApiV2BackgroundJobsSearchPostRequest) SetStartDateNil(b bool)`

 SetStartDateNil sets the value for StartDate to be an explicit nil

### UnsetStartDate
`func (o *ApiV2BackgroundJobsSearchPostRequest) UnsetStartDate()`

UnsetStartDate ensures that no value is present for StartDate, not even an explicit nil
### GetEndDate

`func (o *ApiV2BackgroundJobsSearchPostRequest) GetEndDate() DateTimeRangeSelectorModel`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *ApiV2BackgroundJobsSearchPostRequest) GetEndDateOk() (*DateTimeRangeSelectorModel, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *ApiV2BackgroundJobsSearchPostRequest) SetEndDate(v DateTimeRangeSelectorModel)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *ApiV2BackgroundJobsSearchPostRequest) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### SetEndDateNil

`func (o *ApiV2BackgroundJobsSearchPostRequest) SetEndDateNil(b bool)`

 SetEndDateNil sets the value for EndDate to be an explicit nil

### UnsetEndDate
`func (o *ApiV2BackgroundJobsSearchPostRequest) UnsetEndDate()`

UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


