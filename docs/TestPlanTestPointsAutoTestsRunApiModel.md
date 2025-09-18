# TestPlanTestPointsAutoTestsRunApiModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filter** | Pointer to [**NullableTestPlanTestPointsSearchApiModel**](TestPlanTestPointsSearchApiModel.md) | Test points filters. | [optional] 
**ExtractionModel** | Pointer to [**NullableTestPlanTestPointsExtractionApiModel**](TestPlanTestPointsExtractionApiModel.md) | Test points extraction model. | [optional] 
**WebhookIds** | **[]string** | Webhook ids to run. | 
**Build** | Pointer to **NullableString** | Specifies the test run build. | [optional] 
**ResetNotActualAutomatedTestPoints** | **bool** | Reset test point status when actual work item does not automated. | 

## Methods

### NewTestPlanTestPointsAutoTestsRunApiModel

`func NewTestPlanTestPointsAutoTestsRunApiModel(webhookIds []string, resetNotActualAutomatedTestPoints bool, ) *TestPlanTestPointsAutoTestsRunApiModel`

NewTestPlanTestPointsAutoTestsRunApiModel instantiates a new TestPlanTestPointsAutoTestsRunApiModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestPlanTestPointsAutoTestsRunApiModelWithDefaults

`func NewTestPlanTestPointsAutoTestsRunApiModelWithDefaults() *TestPlanTestPointsAutoTestsRunApiModel`

NewTestPlanTestPointsAutoTestsRunApiModelWithDefaults instantiates a new TestPlanTestPointsAutoTestsRunApiModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilter

`func (o *TestPlanTestPointsAutoTestsRunApiModel) GetFilter() TestPlanTestPointsSearchApiModel`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *TestPlanTestPointsAutoTestsRunApiModel) GetFilterOk() (*TestPlanTestPointsSearchApiModel, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *TestPlanTestPointsAutoTestsRunApiModel) SetFilter(v TestPlanTestPointsSearchApiModel)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *TestPlanTestPointsAutoTestsRunApiModel) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### SetFilterNil

`func (o *TestPlanTestPointsAutoTestsRunApiModel) SetFilterNil(b bool)`

 SetFilterNil sets the value for Filter to be an explicit nil

### UnsetFilter
`func (o *TestPlanTestPointsAutoTestsRunApiModel) UnsetFilter()`

UnsetFilter ensures that no value is present for Filter, not even an explicit nil
### GetExtractionModel

`func (o *TestPlanTestPointsAutoTestsRunApiModel) GetExtractionModel() TestPlanTestPointsExtractionApiModel`

GetExtractionModel returns the ExtractionModel field if non-nil, zero value otherwise.

### GetExtractionModelOk

`func (o *TestPlanTestPointsAutoTestsRunApiModel) GetExtractionModelOk() (*TestPlanTestPointsExtractionApiModel, bool)`

GetExtractionModelOk returns a tuple with the ExtractionModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtractionModel

`func (o *TestPlanTestPointsAutoTestsRunApiModel) SetExtractionModel(v TestPlanTestPointsExtractionApiModel)`

SetExtractionModel sets ExtractionModel field to given value.

### HasExtractionModel

`func (o *TestPlanTestPointsAutoTestsRunApiModel) HasExtractionModel() bool`

HasExtractionModel returns a boolean if a field has been set.

### SetExtractionModelNil

`func (o *TestPlanTestPointsAutoTestsRunApiModel) SetExtractionModelNil(b bool)`

 SetExtractionModelNil sets the value for ExtractionModel to be an explicit nil

### UnsetExtractionModel
`func (o *TestPlanTestPointsAutoTestsRunApiModel) UnsetExtractionModel()`

UnsetExtractionModel ensures that no value is present for ExtractionModel, not even an explicit nil
### GetWebhookIds

`func (o *TestPlanTestPointsAutoTestsRunApiModel) GetWebhookIds() []string`

GetWebhookIds returns the WebhookIds field if non-nil, zero value otherwise.

### GetWebhookIdsOk

`func (o *TestPlanTestPointsAutoTestsRunApiModel) GetWebhookIdsOk() (*[]string, bool)`

GetWebhookIdsOk returns a tuple with the WebhookIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookIds

`func (o *TestPlanTestPointsAutoTestsRunApiModel) SetWebhookIds(v []string)`

SetWebhookIds sets WebhookIds field to given value.


### GetBuild

`func (o *TestPlanTestPointsAutoTestsRunApiModel) GetBuild() string`

GetBuild returns the Build field if non-nil, zero value otherwise.

### GetBuildOk

`func (o *TestPlanTestPointsAutoTestsRunApiModel) GetBuildOk() (*string, bool)`

GetBuildOk returns a tuple with the Build field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuild

`func (o *TestPlanTestPointsAutoTestsRunApiModel) SetBuild(v string)`

SetBuild sets Build field to given value.

### HasBuild

`func (o *TestPlanTestPointsAutoTestsRunApiModel) HasBuild() bool`

HasBuild returns a boolean if a field has been set.

### SetBuildNil

`func (o *TestPlanTestPointsAutoTestsRunApiModel) SetBuildNil(b bool)`

 SetBuildNil sets the value for Build to be an explicit nil

### UnsetBuild
`func (o *TestPlanTestPointsAutoTestsRunApiModel) UnsetBuild()`

UnsetBuild ensures that no value is present for Build, not even an explicit nil
### GetResetNotActualAutomatedTestPoints

`func (o *TestPlanTestPointsAutoTestsRunApiModel) GetResetNotActualAutomatedTestPoints() bool`

GetResetNotActualAutomatedTestPoints returns the ResetNotActualAutomatedTestPoints field if non-nil, zero value otherwise.

### GetResetNotActualAutomatedTestPointsOk

`func (o *TestPlanTestPointsAutoTestsRunApiModel) GetResetNotActualAutomatedTestPointsOk() (*bool, bool)`

GetResetNotActualAutomatedTestPointsOk returns a tuple with the ResetNotActualAutomatedTestPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetNotActualAutomatedTestPoints

`func (o *TestPlanTestPointsAutoTestsRunApiModel) SetResetNotActualAutomatedTestPoints(v bool)`

SetResetNotActualAutomatedTestPoints sets ResetNotActualAutomatedTestPoints field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


