# TestPlanChangeModelTestPlanChangedFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to [**NullableStringChangedFieldWithDiffsViewModel**](StringChangedFieldWithDiffsViewModel.md) |  | [optional] 
**Description** | Pointer to [**NullableStringChangedFieldWithDiffsViewModel**](StringChangedFieldWithDiffsViewModel.md) |  | [optional] 
**ProductName** | Pointer to [**NullableStringChangedFieldViewModel**](StringChangedFieldViewModel.md) |  | [optional] 
**Build** | Pointer to [**NullableStringChangedFieldViewModel**](StringChangedFieldViewModel.md) |  | [optional] 
**Period** | Pointer to [**NullablePeriodViewModelChangedFieldViewModel**](PeriodViewModelChangedFieldViewModel.md) |  | [optional] 
**Status** | Pointer to [**NullableStringChangedFieldViewModel**](StringChangedFieldViewModel.md) |  | [optional] 
**Tags** | Pointer to [**NullableStringArrayChangedFieldViewModel**](StringArrayChangedFieldViewModel.md) |  | [optional] 
**TestSuite** | Pointer to [**NullableTestSuiteChangeViewModelChangedFieldViewModel**](TestSuiteChangeViewModelChangedFieldViewModel.md) |  | [optional] 
**TestPoints** | Pointer to [**NullableTestPointChangeViewModelChangedFieldViewModel**](TestPointChangeViewModelChangedFieldViewModel.md) |  | [optional] 
**TestResults** | Pointer to [**NullableTestResultChangeViewModelChangedFieldViewModel**](TestResultChangeViewModelChangedFieldViewModel.md) |  | [optional] 
**Locking** | Pointer to [**NullableBooleanChangedFieldViewModel**](BooleanChangedFieldViewModel.md) |  | [optional] 
**HasAutomaticDurationTimer** | Pointer to [**NullableBooleanNullableChangedFieldViewModel**](BooleanNullableChangedFieldViewModel.md) |  | [optional] 
**Attributes** | Pointer to [**map[string]CustomAttributeChangeModel**](CustomAttributeChangeModel.md) |  | [optional] 

## Methods

### NewTestPlanChangeModelTestPlanChangedFields

`func NewTestPlanChangeModelTestPlanChangedFields() *TestPlanChangeModelTestPlanChangedFields`

NewTestPlanChangeModelTestPlanChangedFields instantiates a new TestPlanChangeModelTestPlanChangedFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestPlanChangeModelTestPlanChangedFieldsWithDefaults

`func NewTestPlanChangeModelTestPlanChangedFieldsWithDefaults() *TestPlanChangeModelTestPlanChangedFields`

NewTestPlanChangeModelTestPlanChangedFieldsWithDefaults instantiates a new TestPlanChangeModelTestPlanChangedFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TestPlanChangeModelTestPlanChangedFields) GetName() StringChangedFieldWithDiffsViewModel`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TestPlanChangeModelTestPlanChangedFields) GetNameOk() (*StringChangedFieldWithDiffsViewModel, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TestPlanChangeModelTestPlanChangedFields) SetName(v StringChangedFieldWithDiffsViewModel)`

SetName sets Name field to given value.

### HasName

`func (o *TestPlanChangeModelTestPlanChangedFields) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *TestPlanChangeModelTestPlanChangedFields) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *TestPlanChangeModelTestPlanChangedFields) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *TestPlanChangeModelTestPlanChangedFields) GetDescription() StringChangedFieldWithDiffsViewModel`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TestPlanChangeModelTestPlanChangedFields) GetDescriptionOk() (*StringChangedFieldWithDiffsViewModel, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TestPlanChangeModelTestPlanChangedFields) SetDescription(v StringChangedFieldWithDiffsViewModel)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TestPlanChangeModelTestPlanChangedFields) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *TestPlanChangeModelTestPlanChangedFields) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *TestPlanChangeModelTestPlanChangedFields) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetProductName

`func (o *TestPlanChangeModelTestPlanChangedFields) GetProductName() StringChangedFieldViewModel`

GetProductName returns the ProductName field if non-nil, zero value otherwise.

### GetProductNameOk

`func (o *TestPlanChangeModelTestPlanChangedFields) GetProductNameOk() (*StringChangedFieldViewModel, bool)`

GetProductNameOk returns a tuple with the ProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductName

`func (o *TestPlanChangeModelTestPlanChangedFields) SetProductName(v StringChangedFieldViewModel)`

SetProductName sets ProductName field to given value.

### HasProductName

`func (o *TestPlanChangeModelTestPlanChangedFields) HasProductName() bool`

HasProductName returns a boolean if a field has been set.

### SetProductNameNil

`func (o *TestPlanChangeModelTestPlanChangedFields) SetProductNameNil(b bool)`

 SetProductNameNil sets the value for ProductName to be an explicit nil

### UnsetProductName
`func (o *TestPlanChangeModelTestPlanChangedFields) UnsetProductName()`

UnsetProductName ensures that no value is present for ProductName, not even an explicit nil
### GetBuild

`func (o *TestPlanChangeModelTestPlanChangedFields) GetBuild() StringChangedFieldViewModel`

GetBuild returns the Build field if non-nil, zero value otherwise.

### GetBuildOk

`func (o *TestPlanChangeModelTestPlanChangedFields) GetBuildOk() (*StringChangedFieldViewModel, bool)`

GetBuildOk returns a tuple with the Build field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuild

`func (o *TestPlanChangeModelTestPlanChangedFields) SetBuild(v StringChangedFieldViewModel)`

SetBuild sets Build field to given value.

### HasBuild

`func (o *TestPlanChangeModelTestPlanChangedFields) HasBuild() bool`

HasBuild returns a boolean if a field has been set.

### SetBuildNil

`func (o *TestPlanChangeModelTestPlanChangedFields) SetBuildNil(b bool)`

 SetBuildNil sets the value for Build to be an explicit nil

### UnsetBuild
`func (o *TestPlanChangeModelTestPlanChangedFields) UnsetBuild()`

UnsetBuild ensures that no value is present for Build, not even an explicit nil
### GetPeriod

`func (o *TestPlanChangeModelTestPlanChangedFields) GetPeriod() PeriodViewModelChangedFieldViewModel`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *TestPlanChangeModelTestPlanChangedFields) GetPeriodOk() (*PeriodViewModelChangedFieldViewModel, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *TestPlanChangeModelTestPlanChangedFields) SetPeriod(v PeriodViewModelChangedFieldViewModel)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *TestPlanChangeModelTestPlanChangedFields) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### SetPeriodNil

`func (o *TestPlanChangeModelTestPlanChangedFields) SetPeriodNil(b bool)`

 SetPeriodNil sets the value for Period to be an explicit nil

### UnsetPeriod
`func (o *TestPlanChangeModelTestPlanChangedFields) UnsetPeriod()`

UnsetPeriod ensures that no value is present for Period, not even an explicit nil
### GetStatus

`func (o *TestPlanChangeModelTestPlanChangedFields) GetStatus() StringChangedFieldViewModel`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TestPlanChangeModelTestPlanChangedFields) GetStatusOk() (*StringChangedFieldViewModel, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TestPlanChangeModelTestPlanChangedFields) SetStatus(v StringChangedFieldViewModel)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TestPlanChangeModelTestPlanChangedFields) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *TestPlanChangeModelTestPlanChangedFields) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *TestPlanChangeModelTestPlanChangedFields) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetTags

`func (o *TestPlanChangeModelTestPlanChangedFields) GetTags() StringArrayChangedFieldViewModel`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *TestPlanChangeModelTestPlanChangedFields) GetTagsOk() (*StringArrayChangedFieldViewModel, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *TestPlanChangeModelTestPlanChangedFields) SetTags(v StringArrayChangedFieldViewModel)`

SetTags sets Tags field to given value.

### HasTags

`func (o *TestPlanChangeModelTestPlanChangedFields) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *TestPlanChangeModelTestPlanChangedFields) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *TestPlanChangeModelTestPlanChangedFields) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetTestSuite

`func (o *TestPlanChangeModelTestPlanChangedFields) GetTestSuite() TestSuiteChangeViewModelChangedFieldViewModel`

GetTestSuite returns the TestSuite field if non-nil, zero value otherwise.

### GetTestSuiteOk

`func (o *TestPlanChangeModelTestPlanChangedFields) GetTestSuiteOk() (*TestSuiteChangeViewModelChangedFieldViewModel, bool)`

GetTestSuiteOk returns a tuple with the TestSuite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestSuite

`func (o *TestPlanChangeModelTestPlanChangedFields) SetTestSuite(v TestSuiteChangeViewModelChangedFieldViewModel)`

SetTestSuite sets TestSuite field to given value.

### HasTestSuite

`func (o *TestPlanChangeModelTestPlanChangedFields) HasTestSuite() bool`

HasTestSuite returns a boolean if a field has been set.

### SetTestSuiteNil

`func (o *TestPlanChangeModelTestPlanChangedFields) SetTestSuiteNil(b bool)`

 SetTestSuiteNil sets the value for TestSuite to be an explicit nil

### UnsetTestSuite
`func (o *TestPlanChangeModelTestPlanChangedFields) UnsetTestSuite()`

UnsetTestSuite ensures that no value is present for TestSuite, not even an explicit nil
### GetTestPoints

`func (o *TestPlanChangeModelTestPlanChangedFields) GetTestPoints() TestPointChangeViewModelChangedFieldViewModel`

GetTestPoints returns the TestPoints field if non-nil, zero value otherwise.

### GetTestPointsOk

`func (o *TestPlanChangeModelTestPlanChangedFields) GetTestPointsOk() (*TestPointChangeViewModelChangedFieldViewModel, bool)`

GetTestPointsOk returns a tuple with the TestPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestPoints

`func (o *TestPlanChangeModelTestPlanChangedFields) SetTestPoints(v TestPointChangeViewModelChangedFieldViewModel)`

SetTestPoints sets TestPoints field to given value.

### HasTestPoints

`func (o *TestPlanChangeModelTestPlanChangedFields) HasTestPoints() bool`

HasTestPoints returns a boolean if a field has been set.

### SetTestPointsNil

`func (o *TestPlanChangeModelTestPlanChangedFields) SetTestPointsNil(b bool)`

 SetTestPointsNil sets the value for TestPoints to be an explicit nil

### UnsetTestPoints
`func (o *TestPlanChangeModelTestPlanChangedFields) UnsetTestPoints()`

UnsetTestPoints ensures that no value is present for TestPoints, not even an explicit nil
### GetTestResults

`func (o *TestPlanChangeModelTestPlanChangedFields) GetTestResults() TestResultChangeViewModelChangedFieldViewModel`

GetTestResults returns the TestResults field if non-nil, zero value otherwise.

### GetTestResultsOk

`func (o *TestPlanChangeModelTestPlanChangedFields) GetTestResultsOk() (*TestResultChangeViewModelChangedFieldViewModel, bool)`

GetTestResultsOk returns a tuple with the TestResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestResults

`func (o *TestPlanChangeModelTestPlanChangedFields) SetTestResults(v TestResultChangeViewModelChangedFieldViewModel)`

SetTestResults sets TestResults field to given value.

### HasTestResults

`func (o *TestPlanChangeModelTestPlanChangedFields) HasTestResults() bool`

HasTestResults returns a boolean if a field has been set.

### SetTestResultsNil

`func (o *TestPlanChangeModelTestPlanChangedFields) SetTestResultsNil(b bool)`

 SetTestResultsNil sets the value for TestResults to be an explicit nil

### UnsetTestResults
`func (o *TestPlanChangeModelTestPlanChangedFields) UnsetTestResults()`

UnsetTestResults ensures that no value is present for TestResults, not even an explicit nil
### GetLocking

`func (o *TestPlanChangeModelTestPlanChangedFields) GetLocking() BooleanChangedFieldViewModel`

GetLocking returns the Locking field if non-nil, zero value otherwise.

### GetLockingOk

`func (o *TestPlanChangeModelTestPlanChangedFields) GetLockingOk() (*BooleanChangedFieldViewModel, bool)`

GetLockingOk returns a tuple with the Locking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocking

`func (o *TestPlanChangeModelTestPlanChangedFields) SetLocking(v BooleanChangedFieldViewModel)`

SetLocking sets Locking field to given value.

### HasLocking

`func (o *TestPlanChangeModelTestPlanChangedFields) HasLocking() bool`

HasLocking returns a boolean if a field has been set.

### SetLockingNil

`func (o *TestPlanChangeModelTestPlanChangedFields) SetLockingNil(b bool)`

 SetLockingNil sets the value for Locking to be an explicit nil

### UnsetLocking
`func (o *TestPlanChangeModelTestPlanChangedFields) UnsetLocking()`

UnsetLocking ensures that no value is present for Locking, not even an explicit nil
### GetHasAutomaticDurationTimer

`func (o *TestPlanChangeModelTestPlanChangedFields) GetHasAutomaticDurationTimer() BooleanNullableChangedFieldViewModel`

GetHasAutomaticDurationTimer returns the HasAutomaticDurationTimer field if non-nil, zero value otherwise.

### GetHasAutomaticDurationTimerOk

`func (o *TestPlanChangeModelTestPlanChangedFields) GetHasAutomaticDurationTimerOk() (*BooleanNullableChangedFieldViewModel, bool)`

GetHasAutomaticDurationTimerOk returns a tuple with the HasAutomaticDurationTimer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAutomaticDurationTimer

`func (o *TestPlanChangeModelTestPlanChangedFields) SetHasAutomaticDurationTimer(v BooleanNullableChangedFieldViewModel)`

SetHasAutomaticDurationTimer sets HasAutomaticDurationTimer field to given value.

### HasHasAutomaticDurationTimer

`func (o *TestPlanChangeModelTestPlanChangedFields) HasHasAutomaticDurationTimer() bool`

HasHasAutomaticDurationTimer returns a boolean if a field has been set.

### SetHasAutomaticDurationTimerNil

`func (o *TestPlanChangeModelTestPlanChangedFields) SetHasAutomaticDurationTimerNil(b bool)`

 SetHasAutomaticDurationTimerNil sets the value for HasAutomaticDurationTimer to be an explicit nil

### UnsetHasAutomaticDurationTimer
`func (o *TestPlanChangeModelTestPlanChangedFields) UnsetHasAutomaticDurationTimer()`

UnsetHasAutomaticDurationTimer ensures that no value is present for HasAutomaticDurationTimer, not even an explicit nil
### GetAttributes

`func (o *TestPlanChangeModelTestPlanChangedFields) GetAttributes() map[string]CustomAttributeChangeModel`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *TestPlanChangeModelTestPlanChangedFields) GetAttributesOk() (*map[string]CustomAttributeChangeModel, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *TestPlanChangeModelTestPlanChangedFields) SetAttributes(v map[string]CustomAttributeChangeModel)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *TestPlanChangeModelTestPlanChangedFields) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### SetAttributesNil

`func (o *TestPlanChangeModelTestPlanChangedFields) SetAttributesNil(b bool)`

 SetAttributesNil sets the value for Attributes to be an explicit nil

### UnsetAttributes
`func (o *TestPlanChangeModelTestPlanChangedFields) UnsetAttributes()`

UnsetAttributes ensures that no value is present for Attributes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


