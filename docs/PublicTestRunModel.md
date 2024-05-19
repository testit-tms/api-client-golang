# PublicTestRunModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TestRunId** | **string** |  | 
**TestPlanId** | Pointer to **NullableString** |  | [optional] 
**TestPlanGlobalId** | **int64** |  | 
**Name** | **string** |  | 
**ProductName** | Pointer to **NullableString** |  | [optional] 
**Build** | Pointer to **NullableString** |  | [optional] 
**Configurations** | [**[]ConfigurationModel**](ConfigurationModel.md) |  | 
**AutoTests** | [**[]AutoTestModel**](AutoTestModel.md) |  | 
**TestPoints** | [**[]PublicTestPointModel**](PublicTestPointModel.md) |  | 
**Status** | **string** |  | 
**CustomParameters** | Pointer to **map[string]string** |  | [optional] 
**TestRunDescription** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewPublicTestRunModel

`func NewPublicTestRunModel(testRunId string, testPlanGlobalId int64, name string, configurations []ConfigurationModel, autoTests []AutoTestModel, testPoints []PublicTestPointModel, status string, ) *PublicTestRunModel`

NewPublicTestRunModel instantiates a new PublicTestRunModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicTestRunModelWithDefaults

`func NewPublicTestRunModelWithDefaults() *PublicTestRunModel`

NewPublicTestRunModelWithDefaults instantiates a new PublicTestRunModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTestRunId

`func (o *PublicTestRunModel) GetTestRunId() string`

GetTestRunId returns the TestRunId field if non-nil, zero value otherwise.

### GetTestRunIdOk

`func (o *PublicTestRunModel) GetTestRunIdOk() (*string, bool)`

GetTestRunIdOk returns a tuple with the TestRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestRunId

`func (o *PublicTestRunModel) SetTestRunId(v string)`

SetTestRunId sets TestRunId field to given value.


### GetTestPlanId

`func (o *PublicTestRunModel) GetTestPlanId() string`

GetTestPlanId returns the TestPlanId field if non-nil, zero value otherwise.

### GetTestPlanIdOk

`func (o *PublicTestRunModel) GetTestPlanIdOk() (*string, bool)`

GetTestPlanIdOk returns a tuple with the TestPlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestPlanId

`func (o *PublicTestRunModel) SetTestPlanId(v string)`

SetTestPlanId sets TestPlanId field to given value.

### HasTestPlanId

`func (o *PublicTestRunModel) HasTestPlanId() bool`

HasTestPlanId returns a boolean if a field has been set.

### SetTestPlanIdNil

`func (o *PublicTestRunModel) SetTestPlanIdNil(b bool)`

 SetTestPlanIdNil sets the value for TestPlanId to be an explicit nil

### UnsetTestPlanId
`func (o *PublicTestRunModel) UnsetTestPlanId()`

UnsetTestPlanId ensures that no value is present for TestPlanId, not even an explicit nil
### GetTestPlanGlobalId

`func (o *PublicTestRunModel) GetTestPlanGlobalId() int64`

GetTestPlanGlobalId returns the TestPlanGlobalId field if non-nil, zero value otherwise.

### GetTestPlanGlobalIdOk

`func (o *PublicTestRunModel) GetTestPlanGlobalIdOk() (*int64, bool)`

GetTestPlanGlobalIdOk returns a tuple with the TestPlanGlobalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestPlanGlobalId

`func (o *PublicTestRunModel) SetTestPlanGlobalId(v int64)`

SetTestPlanGlobalId sets TestPlanGlobalId field to given value.


### GetName

`func (o *PublicTestRunModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PublicTestRunModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PublicTestRunModel) SetName(v string)`

SetName sets Name field to given value.


### GetProductName

`func (o *PublicTestRunModel) GetProductName() string`

GetProductName returns the ProductName field if non-nil, zero value otherwise.

### GetProductNameOk

`func (o *PublicTestRunModel) GetProductNameOk() (*string, bool)`

GetProductNameOk returns a tuple with the ProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductName

`func (o *PublicTestRunModel) SetProductName(v string)`

SetProductName sets ProductName field to given value.

### HasProductName

`func (o *PublicTestRunModel) HasProductName() bool`

HasProductName returns a boolean if a field has been set.

### SetProductNameNil

`func (o *PublicTestRunModel) SetProductNameNil(b bool)`

 SetProductNameNil sets the value for ProductName to be an explicit nil

### UnsetProductName
`func (o *PublicTestRunModel) UnsetProductName()`

UnsetProductName ensures that no value is present for ProductName, not even an explicit nil
### GetBuild

`func (o *PublicTestRunModel) GetBuild() string`

GetBuild returns the Build field if non-nil, zero value otherwise.

### GetBuildOk

`func (o *PublicTestRunModel) GetBuildOk() (*string, bool)`

GetBuildOk returns a tuple with the Build field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuild

`func (o *PublicTestRunModel) SetBuild(v string)`

SetBuild sets Build field to given value.

### HasBuild

`func (o *PublicTestRunModel) HasBuild() bool`

HasBuild returns a boolean if a field has been set.

### SetBuildNil

`func (o *PublicTestRunModel) SetBuildNil(b bool)`

 SetBuildNil sets the value for Build to be an explicit nil

### UnsetBuild
`func (o *PublicTestRunModel) UnsetBuild()`

UnsetBuild ensures that no value is present for Build, not even an explicit nil
### GetConfigurations

`func (o *PublicTestRunModel) GetConfigurations() []ConfigurationModel`

GetConfigurations returns the Configurations field if non-nil, zero value otherwise.

### GetConfigurationsOk

`func (o *PublicTestRunModel) GetConfigurationsOk() (*[]ConfigurationModel, bool)`

GetConfigurationsOk returns a tuple with the Configurations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurations

`func (o *PublicTestRunModel) SetConfigurations(v []ConfigurationModel)`

SetConfigurations sets Configurations field to given value.


### GetAutoTests

`func (o *PublicTestRunModel) GetAutoTests() []AutoTestModel`

GetAutoTests returns the AutoTests field if non-nil, zero value otherwise.

### GetAutoTestsOk

`func (o *PublicTestRunModel) GetAutoTestsOk() (*[]AutoTestModel, bool)`

GetAutoTestsOk returns a tuple with the AutoTests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoTests

`func (o *PublicTestRunModel) SetAutoTests(v []AutoTestModel)`

SetAutoTests sets AutoTests field to given value.


### GetTestPoints

`func (o *PublicTestRunModel) GetTestPoints() []PublicTestPointModel`

GetTestPoints returns the TestPoints field if non-nil, zero value otherwise.

### GetTestPointsOk

`func (o *PublicTestRunModel) GetTestPointsOk() (*[]PublicTestPointModel, bool)`

GetTestPointsOk returns a tuple with the TestPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestPoints

`func (o *PublicTestRunModel) SetTestPoints(v []PublicTestPointModel)`

SetTestPoints sets TestPoints field to given value.


### GetStatus

`func (o *PublicTestRunModel) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PublicTestRunModel) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PublicTestRunModel) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetCustomParameters

`func (o *PublicTestRunModel) GetCustomParameters() map[string]string`

GetCustomParameters returns the CustomParameters field if non-nil, zero value otherwise.

### GetCustomParametersOk

`func (o *PublicTestRunModel) GetCustomParametersOk() (*map[string]string, bool)`

GetCustomParametersOk returns a tuple with the CustomParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomParameters

`func (o *PublicTestRunModel) SetCustomParameters(v map[string]string)`

SetCustomParameters sets CustomParameters field to given value.

### HasCustomParameters

`func (o *PublicTestRunModel) HasCustomParameters() bool`

HasCustomParameters returns a boolean if a field has been set.

### SetCustomParametersNil

`func (o *PublicTestRunModel) SetCustomParametersNil(b bool)`

 SetCustomParametersNil sets the value for CustomParameters to be an explicit nil

### UnsetCustomParameters
`func (o *PublicTestRunModel) UnsetCustomParameters()`

UnsetCustomParameters ensures that no value is present for CustomParameters, not even an explicit nil
### GetTestRunDescription

`func (o *PublicTestRunModel) GetTestRunDescription() string`

GetTestRunDescription returns the TestRunDescription field if non-nil, zero value otherwise.

### GetTestRunDescriptionOk

`func (o *PublicTestRunModel) GetTestRunDescriptionOk() (*string, bool)`

GetTestRunDescriptionOk returns a tuple with the TestRunDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestRunDescription

`func (o *PublicTestRunModel) SetTestRunDescription(v string)`

SetTestRunDescription sets TestRunDescription field to given value.

### HasTestRunDescription

`func (o *PublicTestRunModel) HasTestRunDescription() bool`

HasTestRunDescription returns a boolean if a field has been set.

### SetTestRunDescriptionNil

`func (o *PublicTestRunModel) SetTestRunDescriptionNil(b bool)`

 SetTestRunDescriptionNil sets the value for TestRunDescription to be an explicit nil

### UnsetTestRunDescription
`func (o *PublicTestRunModel) UnsetTestRunDescription()`

UnsetTestRunDescription ensures that no value is present for TestRunDescription, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


