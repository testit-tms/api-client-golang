# TestPlanChangedFieldsViewModel

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

### NewTestPlanChangedFieldsViewModel

`func NewTestPlanChangedFieldsViewModel() *TestPlanChangedFieldsViewModel`

NewTestPlanChangedFieldsViewModel instantiates a new TestPlanChangedFieldsViewModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestPlanChangedFieldsViewModelWithDefaults

`func NewTestPlanChangedFieldsViewModelWithDefaults() *TestPlanChangedFieldsViewModel`

NewTestPlanChangedFieldsViewModelWithDefaults instantiates a new TestPlanChangedFieldsViewModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TestPlanChangedFieldsViewModel) GetName() StringChangedFieldWithDiffsViewModel`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TestPlanChangedFieldsViewModel) GetNameOk() (*StringChangedFieldWithDiffsViewModel, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TestPlanChangedFieldsViewModel) SetName(v StringChangedFieldWithDiffsViewModel)`

SetName sets Name field to given value.

### HasName

`func (o *TestPlanChangedFieldsViewModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *TestPlanChangedFieldsViewModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *TestPlanChangedFieldsViewModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *TestPlanChangedFieldsViewModel) GetDescription() StringChangedFieldWithDiffsViewModel`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TestPlanChangedFieldsViewModel) GetDescriptionOk() (*StringChangedFieldWithDiffsViewModel, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TestPlanChangedFieldsViewModel) SetDescription(v StringChangedFieldWithDiffsViewModel)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TestPlanChangedFieldsViewModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *TestPlanChangedFieldsViewModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *TestPlanChangedFieldsViewModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetProductName

`func (o *TestPlanChangedFieldsViewModel) GetProductName() StringChangedFieldViewModel`

GetProductName returns the ProductName field if non-nil, zero value otherwise.

### GetProductNameOk

`func (o *TestPlanChangedFieldsViewModel) GetProductNameOk() (*StringChangedFieldViewModel, bool)`

GetProductNameOk returns a tuple with the ProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductName

`func (o *TestPlanChangedFieldsViewModel) SetProductName(v StringChangedFieldViewModel)`

SetProductName sets ProductName field to given value.

### HasProductName

`func (o *TestPlanChangedFieldsViewModel) HasProductName() bool`

HasProductName returns a boolean if a field has been set.

### SetProductNameNil

`func (o *TestPlanChangedFieldsViewModel) SetProductNameNil(b bool)`

 SetProductNameNil sets the value for ProductName to be an explicit nil

### UnsetProductName
`func (o *TestPlanChangedFieldsViewModel) UnsetProductName()`

UnsetProductName ensures that no value is present for ProductName, not even an explicit nil
### GetBuild

`func (o *TestPlanChangedFieldsViewModel) GetBuild() StringChangedFieldViewModel`

GetBuild returns the Build field if non-nil, zero value otherwise.

### GetBuildOk

`func (o *TestPlanChangedFieldsViewModel) GetBuildOk() (*StringChangedFieldViewModel, bool)`

GetBuildOk returns a tuple with the Build field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuild

`func (o *TestPlanChangedFieldsViewModel) SetBuild(v StringChangedFieldViewModel)`

SetBuild sets Build field to given value.

### HasBuild

`func (o *TestPlanChangedFieldsViewModel) HasBuild() bool`

HasBuild returns a boolean if a field has been set.

### SetBuildNil

`func (o *TestPlanChangedFieldsViewModel) SetBuildNil(b bool)`

 SetBuildNil sets the value for Build to be an explicit nil

### UnsetBuild
`func (o *TestPlanChangedFieldsViewModel) UnsetBuild()`

UnsetBuild ensures that no value is present for Build, not even an explicit nil
### GetPeriod

`func (o *TestPlanChangedFieldsViewModel) GetPeriod() PeriodViewModelChangedFieldViewModel`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *TestPlanChangedFieldsViewModel) GetPeriodOk() (*PeriodViewModelChangedFieldViewModel, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *TestPlanChangedFieldsViewModel) SetPeriod(v PeriodViewModelChangedFieldViewModel)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *TestPlanChangedFieldsViewModel) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### SetPeriodNil

`func (o *TestPlanChangedFieldsViewModel) SetPeriodNil(b bool)`

 SetPeriodNil sets the value for Period to be an explicit nil

### UnsetPeriod
`func (o *TestPlanChangedFieldsViewModel) UnsetPeriod()`

UnsetPeriod ensures that no value is present for Period, not even an explicit nil
### GetStatus

`func (o *TestPlanChangedFieldsViewModel) GetStatus() StringChangedFieldViewModel`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TestPlanChangedFieldsViewModel) GetStatusOk() (*StringChangedFieldViewModel, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TestPlanChangedFieldsViewModel) SetStatus(v StringChangedFieldViewModel)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TestPlanChangedFieldsViewModel) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *TestPlanChangedFieldsViewModel) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *TestPlanChangedFieldsViewModel) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetTags

`func (o *TestPlanChangedFieldsViewModel) GetTags() StringArrayChangedFieldViewModel`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *TestPlanChangedFieldsViewModel) GetTagsOk() (*StringArrayChangedFieldViewModel, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *TestPlanChangedFieldsViewModel) SetTags(v StringArrayChangedFieldViewModel)`

SetTags sets Tags field to given value.

### HasTags

`func (o *TestPlanChangedFieldsViewModel) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *TestPlanChangedFieldsViewModel) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *TestPlanChangedFieldsViewModel) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetTestSuite

`func (o *TestPlanChangedFieldsViewModel) GetTestSuite() TestSuiteChangeViewModelChangedFieldViewModel`

GetTestSuite returns the TestSuite field if non-nil, zero value otherwise.

### GetTestSuiteOk

`func (o *TestPlanChangedFieldsViewModel) GetTestSuiteOk() (*TestSuiteChangeViewModelChangedFieldViewModel, bool)`

GetTestSuiteOk returns a tuple with the TestSuite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestSuite

`func (o *TestPlanChangedFieldsViewModel) SetTestSuite(v TestSuiteChangeViewModelChangedFieldViewModel)`

SetTestSuite sets TestSuite field to given value.

### HasTestSuite

`func (o *TestPlanChangedFieldsViewModel) HasTestSuite() bool`

HasTestSuite returns a boolean if a field has been set.

### SetTestSuiteNil

`func (o *TestPlanChangedFieldsViewModel) SetTestSuiteNil(b bool)`

 SetTestSuiteNil sets the value for TestSuite to be an explicit nil

### UnsetTestSuite
`func (o *TestPlanChangedFieldsViewModel) UnsetTestSuite()`

UnsetTestSuite ensures that no value is present for TestSuite, not even an explicit nil
### GetTestPoints

`func (o *TestPlanChangedFieldsViewModel) GetTestPoints() TestPointChangeViewModelChangedFieldViewModel`

GetTestPoints returns the TestPoints field if non-nil, zero value otherwise.

### GetTestPointsOk

`func (o *TestPlanChangedFieldsViewModel) GetTestPointsOk() (*TestPointChangeViewModelChangedFieldViewModel, bool)`

GetTestPointsOk returns a tuple with the TestPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestPoints

`func (o *TestPlanChangedFieldsViewModel) SetTestPoints(v TestPointChangeViewModelChangedFieldViewModel)`

SetTestPoints sets TestPoints field to given value.

### HasTestPoints

`func (o *TestPlanChangedFieldsViewModel) HasTestPoints() bool`

HasTestPoints returns a boolean if a field has been set.

### SetTestPointsNil

`func (o *TestPlanChangedFieldsViewModel) SetTestPointsNil(b bool)`

 SetTestPointsNil sets the value for TestPoints to be an explicit nil

### UnsetTestPoints
`func (o *TestPlanChangedFieldsViewModel) UnsetTestPoints()`

UnsetTestPoints ensures that no value is present for TestPoints, not even an explicit nil
### GetTestResults

`func (o *TestPlanChangedFieldsViewModel) GetTestResults() TestResultChangeViewModelChangedFieldViewModel`

GetTestResults returns the TestResults field if non-nil, zero value otherwise.

### GetTestResultsOk

`func (o *TestPlanChangedFieldsViewModel) GetTestResultsOk() (*TestResultChangeViewModelChangedFieldViewModel, bool)`

GetTestResultsOk returns a tuple with the TestResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestResults

`func (o *TestPlanChangedFieldsViewModel) SetTestResults(v TestResultChangeViewModelChangedFieldViewModel)`

SetTestResults sets TestResults field to given value.

### HasTestResults

`func (o *TestPlanChangedFieldsViewModel) HasTestResults() bool`

HasTestResults returns a boolean if a field has been set.

### SetTestResultsNil

`func (o *TestPlanChangedFieldsViewModel) SetTestResultsNil(b bool)`

 SetTestResultsNil sets the value for TestResults to be an explicit nil

### UnsetTestResults
`func (o *TestPlanChangedFieldsViewModel) UnsetTestResults()`

UnsetTestResults ensures that no value is present for TestResults, not even an explicit nil
### GetLocking

`func (o *TestPlanChangedFieldsViewModel) GetLocking() BooleanChangedFieldViewModel`

GetLocking returns the Locking field if non-nil, zero value otherwise.

### GetLockingOk

`func (o *TestPlanChangedFieldsViewModel) GetLockingOk() (*BooleanChangedFieldViewModel, bool)`

GetLockingOk returns a tuple with the Locking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocking

`func (o *TestPlanChangedFieldsViewModel) SetLocking(v BooleanChangedFieldViewModel)`

SetLocking sets Locking field to given value.

### HasLocking

`func (o *TestPlanChangedFieldsViewModel) HasLocking() bool`

HasLocking returns a boolean if a field has been set.

### SetLockingNil

`func (o *TestPlanChangedFieldsViewModel) SetLockingNil(b bool)`

 SetLockingNil sets the value for Locking to be an explicit nil

### UnsetLocking
`func (o *TestPlanChangedFieldsViewModel) UnsetLocking()`

UnsetLocking ensures that no value is present for Locking, not even an explicit nil
### GetHasAutomaticDurationTimer

`func (o *TestPlanChangedFieldsViewModel) GetHasAutomaticDurationTimer() BooleanNullableChangedFieldViewModel`

GetHasAutomaticDurationTimer returns the HasAutomaticDurationTimer field if non-nil, zero value otherwise.

### GetHasAutomaticDurationTimerOk

`func (o *TestPlanChangedFieldsViewModel) GetHasAutomaticDurationTimerOk() (*BooleanNullableChangedFieldViewModel, bool)`

GetHasAutomaticDurationTimerOk returns a tuple with the HasAutomaticDurationTimer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAutomaticDurationTimer

`func (o *TestPlanChangedFieldsViewModel) SetHasAutomaticDurationTimer(v BooleanNullableChangedFieldViewModel)`

SetHasAutomaticDurationTimer sets HasAutomaticDurationTimer field to given value.

### HasHasAutomaticDurationTimer

`func (o *TestPlanChangedFieldsViewModel) HasHasAutomaticDurationTimer() bool`

HasHasAutomaticDurationTimer returns a boolean if a field has been set.

### SetHasAutomaticDurationTimerNil

`func (o *TestPlanChangedFieldsViewModel) SetHasAutomaticDurationTimerNil(b bool)`

 SetHasAutomaticDurationTimerNil sets the value for HasAutomaticDurationTimer to be an explicit nil

### UnsetHasAutomaticDurationTimer
`func (o *TestPlanChangedFieldsViewModel) UnsetHasAutomaticDurationTimer()`

UnsetHasAutomaticDurationTimer ensures that no value is present for HasAutomaticDurationTimer, not even an explicit nil
### GetAttributes

`func (o *TestPlanChangedFieldsViewModel) GetAttributes() map[string]CustomAttributeChangeModel`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *TestPlanChangedFieldsViewModel) GetAttributesOk() (*map[string]CustomAttributeChangeModel, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *TestPlanChangedFieldsViewModel) SetAttributes(v map[string]CustomAttributeChangeModel)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *TestPlanChangedFieldsViewModel) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### SetAttributesNil

`func (o *TestPlanChangedFieldsViewModel) SetAttributesNil(b bool)`

 SetAttributesNil sets the value for Attributes to be an explicit nil

### UnsetAttributes
`func (o *TestPlanChangedFieldsViewModel) UnsetAttributes()`

UnsetAttributes ensures that no value is present for Attributes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


