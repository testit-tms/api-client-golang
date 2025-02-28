# AutoTestResultHistoryApiResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**TestPlanId** | Pointer to **NullableString** |  | [optional] 
**TestPlanGlobalId** | Pointer to **NullableInt64** |  | [optional] 
**TestPlanName** | Pointer to **NullableString** |  | [optional] 
**Duration** | Pointer to **NullableInt64** |  | [optional] 
**TestRunId** | **string** |  | 
**TestRunName** | Pointer to **NullableString** |  | [optional] 
**ConfigurationId** | **string** |  | 
**ConfigurationName** | **string** |  | 
**Outcome** | [**AutotestResultOutcome**](AutotestResultOutcome.md) |  | 
**Status** | [**TestStatusApiResult**](TestStatusApiResult.md) |  | 
**LaunchSource** | Pointer to **NullableString** |  | [optional] 
**RerunCount** | **int32** |  | 
**RerunTestResults** | [**[]RerunTestResultApiResult**](RerunTestResultApiResult.md) |  | 
**CreatedDate** | **time.Time** |  | 
**CreatedById** | **string** |  | 
**CreatedByName** | Pointer to **NullableString** |  | [optional] 
**ModifiedDate** | Pointer to **NullableTime** |  | [optional] 
**ModifiedById** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewAutoTestResultHistoryApiResult

`func NewAutoTestResultHistoryApiResult(id string, testRunId string, configurationId string, configurationName string, outcome AutotestResultOutcome, status TestStatusApiResult, rerunCount int32, rerunTestResults []RerunTestResultApiResult, createdDate time.Time, createdById string, ) *AutoTestResultHistoryApiResult`

NewAutoTestResultHistoryApiResult instantiates a new AutoTestResultHistoryApiResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoTestResultHistoryApiResultWithDefaults

`func NewAutoTestResultHistoryApiResultWithDefaults() *AutoTestResultHistoryApiResult`

NewAutoTestResultHistoryApiResultWithDefaults instantiates a new AutoTestResultHistoryApiResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AutoTestResultHistoryApiResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AutoTestResultHistoryApiResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AutoTestResultHistoryApiResult) SetId(v string)`

SetId sets Id field to given value.


### GetTestPlanId

`func (o *AutoTestResultHistoryApiResult) GetTestPlanId() string`

GetTestPlanId returns the TestPlanId field if non-nil, zero value otherwise.

### GetTestPlanIdOk

`func (o *AutoTestResultHistoryApiResult) GetTestPlanIdOk() (*string, bool)`

GetTestPlanIdOk returns a tuple with the TestPlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestPlanId

`func (o *AutoTestResultHistoryApiResult) SetTestPlanId(v string)`

SetTestPlanId sets TestPlanId field to given value.

### HasTestPlanId

`func (o *AutoTestResultHistoryApiResult) HasTestPlanId() bool`

HasTestPlanId returns a boolean if a field has been set.

### SetTestPlanIdNil

`func (o *AutoTestResultHistoryApiResult) SetTestPlanIdNil(b bool)`

 SetTestPlanIdNil sets the value for TestPlanId to be an explicit nil

### UnsetTestPlanId
`func (o *AutoTestResultHistoryApiResult) UnsetTestPlanId()`

UnsetTestPlanId ensures that no value is present for TestPlanId, not even an explicit nil
### GetTestPlanGlobalId

`func (o *AutoTestResultHistoryApiResult) GetTestPlanGlobalId() int64`

GetTestPlanGlobalId returns the TestPlanGlobalId field if non-nil, zero value otherwise.

### GetTestPlanGlobalIdOk

`func (o *AutoTestResultHistoryApiResult) GetTestPlanGlobalIdOk() (*int64, bool)`

GetTestPlanGlobalIdOk returns a tuple with the TestPlanGlobalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestPlanGlobalId

`func (o *AutoTestResultHistoryApiResult) SetTestPlanGlobalId(v int64)`

SetTestPlanGlobalId sets TestPlanGlobalId field to given value.

### HasTestPlanGlobalId

`func (o *AutoTestResultHistoryApiResult) HasTestPlanGlobalId() bool`

HasTestPlanGlobalId returns a boolean if a field has been set.

### SetTestPlanGlobalIdNil

`func (o *AutoTestResultHistoryApiResult) SetTestPlanGlobalIdNil(b bool)`

 SetTestPlanGlobalIdNil sets the value for TestPlanGlobalId to be an explicit nil

### UnsetTestPlanGlobalId
`func (o *AutoTestResultHistoryApiResult) UnsetTestPlanGlobalId()`

UnsetTestPlanGlobalId ensures that no value is present for TestPlanGlobalId, not even an explicit nil
### GetTestPlanName

`func (o *AutoTestResultHistoryApiResult) GetTestPlanName() string`

GetTestPlanName returns the TestPlanName field if non-nil, zero value otherwise.

### GetTestPlanNameOk

`func (o *AutoTestResultHistoryApiResult) GetTestPlanNameOk() (*string, bool)`

GetTestPlanNameOk returns a tuple with the TestPlanName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestPlanName

`func (o *AutoTestResultHistoryApiResult) SetTestPlanName(v string)`

SetTestPlanName sets TestPlanName field to given value.

### HasTestPlanName

`func (o *AutoTestResultHistoryApiResult) HasTestPlanName() bool`

HasTestPlanName returns a boolean if a field has been set.

### SetTestPlanNameNil

`func (o *AutoTestResultHistoryApiResult) SetTestPlanNameNil(b bool)`

 SetTestPlanNameNil sets the value for TestPlanName to be an explicit nil

### UnsetTestPlanName
`func (o *AutoTestResultHistoryApiResult) UnsetTestPlanName()`

UnsetTestPlanName ensures that no value is present for TestPlanName, not even an explicit nil
### GetDuration

`func (o *AutoTestResultHistoryApiResult) GetDuration() int64`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *AutoTestResultHistoryApiResult) GetDurationOk() (*int64, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *AutoTestResultHistoryApiResult) SetDuration(v int64)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *AutoTestResultHistoryApiResult) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### SetDurationNil

`func (o *AutoTestResultHistoryApiResult) SetDurationNil(b bool)`

 SetDurationNil sets the value for Duration to be an explicit nil

### UnsetDuration
`func (o *AutoTestResultHistoryApiResult) UnsetDuration()`

UnsetDuration ensures that no value is present for Duration, not even an explicit nil
### GetTestRunId

`func (o *AutoTestResultHistoryApiResult) GetTestRunId() string`

GetTestRunId returns the TestRunId field if non-nil, zero value otherwise.

### GetTestRunIdOk

`func (o *AutoTestResultHistoryApiResult) GetTestRunIdOk() (*string, bool)`

GetTestRunIdOk returns a tuple with the TestRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestRunId

`func (o *AutoTestResultHistoryApiResult) SetTestRunId(v string)`

SetTestRunId sets TestRunId field to given value.


### GetTestRunName

`func (o *AutoTestResultHistoryApiResult) GetTestRunName() string`

GetTestRunName returns the TestRunName field if non-nil, zero value otherwise.

### GetTestRunNameOk

`func (o *AutoTestResultHistoryApiResult) GetTestRunNameOk() (*string, bool)`

GetTestRunNameOk returns a tuple with the TestRunName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestRunName

`func (o *AutoTestResultHistoryApiResult) SetTestRunName(v string)`

SetTestRunName sets TestRunName field to given value.

### HasTestRunName

`func (o *AutoTestResultHistoryApiResult) HasTestRunName() bool`

HasTestRunName returns a boolean if a field has been set.

### SetTestRunNameNil

`func (o *AutoTestResultHistoryApiResult) SetTestRunNameNil(b bool)`

 SetTestRunNameNil sets the value for TestRunName to be an explicit nil

### UnsetTestRunName
`func (o *AutoTestResultHistoryApiResult) UnsetTestRunName()`

UnsetTestRunName ensures that no value is present for TestRunName, not even an explicit nil
### GetConfigurationId

`func (o *AutoTestResultHistoryApiResult) GetConfigurationId() string`

GetConfigurationId returns the ConfigurationId field if non-nil, zero value otherwise.

### GetConfigurationIdOk

`func (o *AutoTestResultHistoryApiResult) GetConfigurationIdOk() (*string, bool)`

GetConfigurationIdOk returns a tuple with the ConfigurationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationId

`func (o *AutoTestResultHistoryApiResult) SetConfigurationId(v string)`

SetConfigurationId sets ConfigurationId field to given value.


### GetConfigurationName

`func (o *AutoTestResultHistoryApiResult) GetConfigurationName() string`

GetConfigurationName returns the ConfigurationName field if non-nil, zero value otherwise.

### GetConfigurationNameOk

`func (o *AutoTestResultHistoryApiResult) GetConfigurationNameOk() (*string, bool)`

GetConfigurationNameOk returns a tuple with the ConfigurationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationName

`func (o *AutoTestResultHistoryApiResult) SetConfigurationName(v string)`

SetConfigurationName sets ConfigurationName field to given value.


### GetOutcome

`func (o *AutoTestResultHistoryApiResult) GetOutcome() AutotestResultOutcome`

GetOutcome returns the Outcome field if non-nil, zero value otherwise.

### GetOutcomeOk

`func (o *AutoTestResultHistoryApiResult) GetOutcomeOk() (*AutotestResultOutcome, bool)`

GetOutcomeOk returns a tuple with the Outcome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcome

`func (o *AutoTestResultHistoryApiResult) SetOutcome(v AutotestResultOutcome)`

SetOutcome sets Outcome field to given value.


### GetStatus

`func (o *AutoTestResultHistoryApiResult) GetStatus() TestStatusApiResult`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AutoTestResultHistoryApiResult) GetStatusOk() (*TestStatusApiResult, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AutoTestResultHistoryApiResult) SetStatus(v TestStatusApiResult)`

SetStatus sets Status field to given value.


### GetLaunchSource

`func (o *AutoTestResultHistoryApiResult) GetLaunchSource() string`

GetLaunchSource returns the LaunchSource field if non-nil, zero value otherwise.

### GetLaunchSourceOk

`func (o *AutoTestResultHistoryApiResult) GetLaunchSourceOk() (*string, bool)`

GetLaunchSourceOk returns a tuple with the LaunchSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLaunchSource

`func (o *AutoTestResultHistoryApiResult) SetLaunchSource(v string)`

SetLaunchSource sets LaunchSource field to given value.

### HasLaunchSource

`func (o *AutoTestResultHistoryApiResult) HasLaunchSource() bool`

HasLaunchSource returns a boolean if a field has been set.

### SetLaunchSourceNil

`func (o *AutoTestResultHistoryApiResult) SetLaunchSourceNil(b bool)`

 SetLaunchSourceNil sets the value for LaunchSource to be an explicit nil

### UnsetLaunchSource
`func (o *AutoTestResultHistoryApiResult) UnsetLaunchSource()`

UnsetLaunchSource ensures that no value is present for LaunchSource, not even an explicit nil
### GetRerunCount

`func (o *AutoTestResultHistoryApiResult) GetRerunCount() int32`

GetRerunCount returns the RerunCount field if non-nil, zero value otherwise.

### GetRerunCountOk

`func (o *AutoTestResultHistoryApiResult) GetRerunCountOk() (*int32, bool)`

GetRerunCountOk returns a tuple with the RerunCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRerunCount

`func (o *AutoTestResultHistoryApiResult) SetRerunCount(v int32)`

SetRerunCount sets RerunCount field to given value.


### GetRerunTestResults

`func (o *AutoTestResultHistoryApiResult) GetRerunTestResults() []RerunTestResultApiResult`

GetRerunTestResults returns the RerunTestResults field if non-nil, zero value otherwise.

### GetRerunTestResultsOk

`func (o *AutoTestResultHistoryApiResult) GetRerunTestResultsOk() (*[]RerunTestResultApiResult, bool)`

GetRerunTestResultsOk returns a tuple with the RerunTestResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRerunTestResults

`func (o *AutoTestResultHistoryApiResult) SetRerunTestResults(v []RerunTestResultApiResult)`

SetRerunTestResults sets RerunTestResults field to given value.


### GetCreatedDate

`func (o *AutoTestResultHistoryApiResult) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *AutoTestResultHistoryApiResult) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *AutoTestResultHistoryApiResult) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.


### GetCreatedById

`func (o *AutoTestResultHistoryApiResult) GetCreatedById() string`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *AutoTestResultHistoryApiResult) GetCreatedByIdOk() (*string, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *AutoTestResultHistoryApiResult) SetCreatedById(v string)`

SetCreatedById sets CreatedById field to given value.


### GetCreatedByName

`func (o *AutoTestResultHistoryApiResult) GetCreatedByName() string`

GetCreatedByName returns the CreatedByName field if non-nil, zero value otherwise.

### GetCreatedByNameOk

`func (o *AutoTestResultHistoryApiResult) GetCreatedByNameOk() (*string, bool)`

GetCreatedByNameOk returns a tuple with the CreatedByName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByName

`func (o *AutoTestResultHistoryApiResult) SetCreatedByName(v string)`

SetCreatedByName sets CreatedByName field to given value.

### HasCreatedByName

`func (o *AutoTestResultHistoryApiResult) HasCreatedByName() bool`

HasCreatedByName returns a boolean if a field has been set.

### SetCreatedByNameNil

`func (o *AutoTestResultHistoryApiResult) SetCreatedByNameNil(b bool)`

 SetCreatedByNameNil sets the value for CreatedByName to be an explicit nil

### UnsetCreatedByName
`func (o *AutoTestResultHistoryApiResult) UnsetCreatedByName()`

UnsetCreatedByName ensures that no value is present for CreatedByName, not even an explicit nil
### GetModifiedDate

`func (o *AutoTestResultHistoryApiResult) GetModifiedDate() time.Time`

GetModifiedDate returns the ModifiedDate field if non-nil, zero value otherwise.

### GetModifiedDateOk

`func (o *AutoTestResultHistoryApiResult) GetModifiedDateOk() (*time.Time, bool)`

GetModifiedDateOk returns a tuple with the ModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedDate

`func (o *AutoTestResultHistoryApiResult) SetModifiedDate(v time.Time)`

SetModifiedDate sets ModifiedDate field to given value.

### HasModifiedDate

`func (o *AutoTestResultHistoryApiResult) HasModifiedDate() bool`

HasModifiedDate returns a boolean if a field has been set.

### SetModifiedDateNil

`func (o *AutoTestResultHistoryApiResult) SetModifiedDateNil(b bool)`

 SetModifiedDateNil sets the value for ModifiedDate to be an explicit nil

### UnsetModifiedDate
`func (o *AutoTestResultHistoryApiResult) UnsetModifiedDate()`

UnsetModifiedDate ensures that no value is present for ModifiedDate, not even an explicit nil
### GetModifiedById

`func (o *AutoTestResultHistoryApiResult) GetModifiedById() string`

GetModifiedById returns the ModifiedById field if non-nil, zero value otherwise.

### GetModifiedByIdOk

`func (o *AutoTestResultHistoryApiResult) GetModifiedByIdOk() (*string, bool)`

GetModifiedByIdOk returns a tuple with the ModifiedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedById

`func (o *AutoTestResultHistoryApiResult) SetModifiedById(v string)`

SetModifiedById sets ModifiedById field to given value.

### HasModifiedById

`func (o *AutoTestResultHistoryApiResult) HasModifiedById() bool`

HasModifiedById returns a boolean if a field has been set.

### SetModifiedByIdNil

`func (o *AutoTestResultHistoryApiResult) SetModifiedByIdNil(b bool)`

 SetModifiedByIdNil sets the value for ModifiedById to be an explicit nil

### UnsetModifiedById
`func (o *AutoTestResultHistoryApiResult) UnsetModifiedById()`

UnsetModifiedById ensures that no value is present for ModifiedById, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


