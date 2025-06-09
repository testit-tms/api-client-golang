# TestPlanTestPointsSearchApiResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Created** | [**AuditApiResult**](AuditApiResult.md) |  | 
**Modified** | Pointer to [**NullableAuditApiResult**](AuditApiResult.md) |  | [optional] 
**Status** | **string** |  | 
**StatusModel** | [**TestStatusShortApiResult**](TestStatusShortApiResult.md) |  | 
**InProgress** | **bool** |  | 
**Configuration** | [**ConfigurationShortApiResult**](ConfigurationShortApiResult.md) |  | 
**Tester** | Pointer to [**NullableUserNameApiResult**](UserNameApiResult.md) |  | [optional] 
**TestSuite** | [**TestPlanTestPointsTestSuiteSearchApiResult**](TestPlanTestPointsTestSuiteSearchApiResult.md) |  | 
**WorkItem** | [**TestPlanTestPointsWorkItemSearchApiResult**](TestPlanTestPointsWorkItemSearchApiResult.md) |  | 
**Parameters** | [**[]ParameterShortApiResult**](ParameterShortApiResult.md) |  | 
**LastTestResult** | Pointer to [**NullableLastTestResultApiResult**](LastTestResultApiResult.md) |  | [optional] 

## Methods

### NewTestPlanTestPointsSearchApiResult

`func NewTestPlanTestPointsSearchApiResult(id string, created AuditApiResult, status string, statusModel TestStatusShortApiResult, inProgress bool, configuration ConfigurationShortApiResult, testSuite TestPlanTestPointsTestSuiteSearchApiResult, workItem TestPlanTestPointsWorkItemSearchApiResult, parameters []ParameterShortApiResult, ) *TestPlanTestPointsSearchApiResult`

NewTestPlanTestPointsSearchApiResult instantiates a new TestPlanTestPointsSearchApiResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestPlanTestPointsSearchApiResultWithDefaults

`func NewTestPlanTestPointsSearchApiResultWithDefaults() *TestPlanTestPointsSearchApiResult`

NewTestPlanTestPointsSearchApiResultWithDefaults instantiates a new TestPlanTestPointsSearchApiResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TestPlanTestPointsSearchApiResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TestPlanTestPointsSearchApiResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TestPlanTestPointsSearchApiResult) SetId(v string)`

SetId sets Id field to given value.


### GetCreated

`func (o *TestPlanTestPointsSearchApiResult) GetCreated() AuditApiResult`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *TestPlanTestPointsSearchApiResult) GetCreatedOk() (*AuditApiResult, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *TestPlanTestPointsSearchApiResult) SetCreated(v AuditApiResult)`

SetCreated sets Created field to given value.


### GetModified

`func (o *TestPlanTestPointsSearchApiResult) GetModified() AuditApiResult`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *TestPlanTestPointsSearchApiResult) GetModifiedOk() (*AuditApiResult, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *TestPlanTestPointsSearchApiResult) SetModified(v AuditApiResult)`

SetModified sets Modified field to given value.

### HasModified

`func (o *TestPlanTestPointsSearchApiResult) HasModified() bool`

HasModified returns a boolean if a field has been set.

### SetModifiedNil

`func (o *TestPlanTestPointsSearchApiResult) SetModifiedNil(b bool)`

 SetModifiedNil sets the value for Modified to be an explicit nil

### UnsetModified
`func (o *TestPlanTestPointsSearchApiResult) UnsetModified()`

UnsetModified ensures that no value is present for Modified, not even an explicit nil
### GetStatus

`func (o *TestPlanTestPointsSearchApiResult) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TestPlanTestPointsSearchApiResult) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TestPlanTestPointsSearchApiResult) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetStatusModel

`func (o *TestPlanTestPointsSearchApiResult) GetStatusModel() TestStatusShortApiResult`

GetStatusModel returns the StatusModel field if non-nil, zero value otherwise.

### GetStatusModelOk

`func (o *TestPlanTestPointsSearchApiResult) GetStatusModelOk() (*TestStatusShortApiResult, bool)`

GetStatusModelOk returns a tuple with the StatusModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusModel

`func (o *TestPlanTestPointsSearchApiResult) SetStatusModel(v TestStatusShortApiResult)`

SetStatusModel sets StatusModel field to given value.


### GetInProgress

`func (o *TestPlanTestPointsSearchApiResult) GetInProgress() bool`

GetInProgress returns the InProgress field if non-nil, zero value otherwise.

### GetInProgressOk

`func (o *TestPlanTestPointsSearchApiResult) GetInProgressOk() (*bool, bool)`

GetInProgressOk returns a tuple with the InProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInProgress

`func (o *TestPlanTestPointsSearchApiResult) SetInProgress(v bool)`

SetInProgress sets InProgress field to given value.


### GetConfiguration

`func (o *TestPlanTestPointsSearchApiResult) GetConfiguration() ConfigurationShortApiResult`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *TestPlanTestPointsSearchApiResult) GetConfigurationOk() (*ConfigurationShortApiResult, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *TestPlanTestPointsSearchApiResult) SetConfiguration(v ConfigurationShortApiResult)`

SetConfiguration sets Configuration field to given value.


### GetTester

`func (o *TestPlanTestPointsSearchApiResult) GetTester() UserNameApiResult`

GetTester returns the Tester field if non-nil, zero value otherwise.

### GetTesterOk

`func (o *TestPlanTestPointsSearchApiResult) GetTesterOk() (*UserNameApiResult, bool)`

GetTesterOk returns a tuple with the Tester field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTester

`func (o *TestPlanTestPointsSearchApiResult) SetTester(v UserNameApiResult)`

SetTester sets Tester field to given value.

### HasTester

`func (o *TestPlanTestPointsSearchApiResult) HasTester() bool`

HasTester returns a boolean if a field has been set.

### SetTesterNil

`func (o *TestPlanTestPointsSearchApiResult) SetTesterNil(b bool)`

 SetTesterNil sets the value for Tester to be an explicit nil

### UnsetTester
`func (o *TestPlanTestPointsSearchApiResult) UnsetTester()`

UnsetTester ensures that no value is present for Tester, not even an explicit nil
### GetTestSuite

`func (o *TestPlanTestPointsSearchApiResult) GetTestSuite() TestPlanTestPointsTestSuiteSearchApiResult`

GetTestSuite returns the TestSuite field if non-nil, zero value otherwise.

### GetTestSuiteOk

`func (o *TestPlanTestPointsSearchApiResult) GetTestSuiteOk() (*TestPlanTestPointsTestSuiteSearchApiResult, bool)`

GetTestSuiteOk returns a tuple with the TestSuite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestSuite

`func (o *TestPlanTestPointsSearchApiResult) SetTestSuite(v TestPlanTestPointsTestSuiteSearchApiResult)`

SetTestSuite sets TestSuite field to given value.


### GetWorkItem

`func (o *TestPlanTestPointsSearchApiResult) GetWorkItem() TestPlanTestPointsWorkItemSearchApiResult`

GetWorkItem returns the WorkItem field if non-nil, zero value otherwise.

### GetWorkItemOk

`func (o *TestPlanTestPointsSearchApiResult) GetWorkItemOk() (*TestPlanTestPointsWorkItemSearchApiResult, bool)`

GetWorkItemOk returns a tuple with the WorkItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkItem

`func (o *TestPlanTestPointsSearchApiResult) SetWorkItem(v TestPlanTestPointsWorkItemSearchApiResult)`

SetWorkItem sets WorkItem field to given value.


### GetParameters

`func (o *TestPlanTestPointsSearchApiResult) GetParameters() []ParameterShortApiResult`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *TestPlanTestPointsSearchApiResult) GetParametersOk() (*[]ParameterShortApiResult, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *TestPlanTestPointsSearchApiResult) SetParameters(v []ParameterShortApiResult)`

SetParameters sets Parameters field to given value.


### GetLastTestResult

`func (o *TestPlanTestPointsSearchApiResult) GetLastTestResult() LastTestResultApiResult`

GetLastTestResult returns the LastTestResult field if non-nil, zero value otherwise.

### GetLastTestResultOk

`func (o *TestPlanTestPointsSearchApiResult) GetLastTestResultOk() (*LastTestResultApiResult, bool)`

GetLastTestResultOk returns a tuple with the LastTestResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTestResult

`func (o *TestPlanTestPointsSearchApiResult) SetLastTestResult(v LastTestResultApiResult)`

SetLastTestResult sets LastTestResult field to given value.

### HasLastTestResult

`func (o *TestPlanTestPointsSearchApiResult) HasLastTestResult() bool`

HasLastTestResult returns a boolean if a field has been set.

### SetLastTestResultNil

`func (o *TestPlanTestPointsSearchApiResult) SetLastTestResultNil(b bool)`

 SetLastTestResultNil sets the value for LastTestResult to be an explicit nil

### UnsetLastTestResult
`func (o *TestPlanTestPointsSearchApiResult) UnsetLastTestResult()`

UnsetLastTestResult ensures that no value is present for LastTestResult, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


