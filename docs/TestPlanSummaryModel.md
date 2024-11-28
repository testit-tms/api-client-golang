# TestPlanSummaryModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalTestPointsCount** | **int32** |  | 
**ManualTestPointsCount** | **int32** |  | 
**AutomatedTestPointsCount** | **int32** |  | 
**CompletedTestPointsCount** | **int32** |  | 
**DefectsCount** | **int32** |  | 
**PlannedTestPointsDuration** | **int64** |  | 
**SpentTestPointsDuration** | Pointer to **NullableInt64** |  | [optional] 

## Methods

### NewTestPlanSummaryModel

`func NewTestPlanSummaryModel(totalTestPointsCount int32, manualTestPointsCount int32, automatedTestPointsCount int32, completedTestPointsCount int32, defectsCount int32, plannedTestPointsDuration int64, ) *TestPlanSummaryModel`

NewTestPlanSummaryModel instantiates a new TestPlanSummaryModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestPlanSummaryModelWithDefaults

`func NewTestPlanSummaryModelWithDefaults() *TestPlanSummaryModel`

NewTestPlanSummaryModelWithDefaults instantiates a new TestPlanSummaryModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalTestPointsCount

`func (o *TestPlanSummaryModel) GetTotalTestPointsCount() int32`

GetTotalTestPointsCount returns the TotalTestPointsCount field if non-nil, zero value otherwise.

### GetTotalTestPointsCountOk

`func (o *TestPlanSummaryModel) GetTotalTestPointsCountOk() (*int32, bool)`

GetTotalTestPointsCountOk returns a tuple with the TotalTestPointsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTestPointsCount

`func (o *TestPlanSummaryModel) SetTotalTestPointsCount(v int32)`

SetTotalTestPointsCount sets TotalTestPointsCount field to given value.


### GetManualTestPointsCount

`func (o *TestPlanSummaryModel) GetManualTestPointsCount() int32`

GetManualTestPointsCount returns the ManualTestPointsCount field if non-nil, zero value otherwise.

### GetManualTestPointsCountOk

`func (o *TestPlanSummaryModel) GetManualTestPointsCountOk() (*int32, bool)`

GetManualTestPointsCountOk returns a tuple with the ManualTestPointsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManualTestPointsCount

`func (o *TestPlanSummaryModel) SetManualTestPointsCount(v int32)`

SetManualTestPointsCount sets ManualTestPointsCount field to given value.


### GetAutomatedTestPointsCount

`func (o *TestPlanSummaryModel) GetAutomatedTestPointsCount() int32`

GetAutomatedTestPointsCount returns the AutomatedTestPointsCount field if non-nil, zero value otherwise.

### GetAutomatedTestPointsCountOk

`func (o *TestPlanSummaryModel) GetAutomatedTestPointsCountOk() (*int32, bool)`

GetAutomatedTestPointsCountOk returns a tuple with the AutomatedTestPointsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomatedTestPointsCount

`func (o *TestPlanSummaryModel) SetAutomatedTestPointsCount(v int32)`

SetAutomatedTestPointsCount sets AutomatedTestPointsCount field to given value.


### GetCompletedTestPointsCount

`func (o *TestPlanSummaryModel) GetCompletedTestPointsCount() int32`

GetCompletedTestPointsCount returns the CompletedTestPointsCount field if non-nil, zero value otherwise.

### GetCompletedTestPointsCountOk

`func (o *TestPlanSummaryModel) GetCompletedTestPointsCountOk() (*int32, bool)`

GetCompletedTestPointsCountOk returns a tuple with the CompletedTestPointsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedTestPointsCount

`func (o *TestPlanSummaryModel) SetCompletedTestPointsCount(v int32)`

SetCompletedTestPointsCount sets CompletedTestPointsCount field to given value.


### GetDefectsCount

`func (o *TestPlanSummaryModel) GetDefectsCount() int32`

GetDefectsCount returns the DefectsCount field if non-nil, zero value otherwise.

### GetDefectsCountOk

`func (o *TestPlanSummaryModel) GetDefectsCountOk() (*int32, bool)`

GetDefectsCountOk returns a tuple with the DefectsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefectsCount

`func (o *TestPlanSummaryModel) SetDefectsCount(v int32)`

SetDefectsCount sets DefectsCount field to given value.


### GetPlannedTestPointsDuration

`func (o *TestPlanSummaryModel) GetPlannedTestPointsDuration() int64`

GetPlannedTestPointsDuration returns the PlannedTestPointsDuration field if non-nil, zero value otherwise.

### GetPlannedTestPointsDurationOk

`func (o *TestPlanSummaryModel) GetPlannedTestPointsDurationOk() (*int64, bool)`

GetPlannedTestPointsDurationOk returns a tuple with the PlannedTestPointsDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedTestPointsDuration

`func (o *TestPlanSummaryModel) SetPlannedTestPointsDuration(v int64)`

SetPlannedTestPointsDuration sets PlannedTestPointsDuration field to given value.


### GetSpentTestPointsDuration

`func (o *TestPlanSummaryModel) GetSpentTestPointsDuration() int64`

GetSpentTestPointsDuration returns the SpentTestPointsDuration field if non-nil, zero value otherwise.

### GetSpentTestPointsDurationOk

`func (o *TestPlanSummaryModel) GetSpentTestPointsDurationOk() (*int64, bool)`

GetSpentTestPointsDurationOk returns a tuple with the SpentTestPointsDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpentTestPointsDuration

`func (o *TestPlanSummaryModel) SetSpentTestPointsDuration(v int64)`

SetSpentTestPointsDuration sets SpentTestPointsDuration field to given value.

### HasSpentTestPointsDuration

`func (o *TestPlanSummaryModel) HasSpentTestPointsDuration() bool`

HasSpentTestPointsDuration returns a boolean if a field has been set.

### SetSpentTestPointsDurationNil

`func (o *TestPlanSummaryModel) SetSpentTestPointsDurationNil(b bool)`

 SetSpentTestPointsDurationNil sets the value for SpentTestPointsDuration to be an explicit nil

### UnsetSpentTestPointsDuration
`func (o *TestPlanSummaryModel) UnsetSpentTestPointsDuration()`

UnsetSpentTestPointsDuration ensures that no value is present for SpentTestPointsDuration, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


