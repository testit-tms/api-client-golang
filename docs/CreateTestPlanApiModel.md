# CreateTestPlanApiModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tags** | Pointer to [**[]TagPostModel**](TagPostModel.md) | Test plan tag names collection | [optional] 
**Name** | **string** | Test plan name | 
**StartDate** | Pointer to **NullableTime** | Date and time of test plan start | [optional] 
**EndDate** | Pointer to **NullableTime** | Date and time of test plan end | [optional] 
**Description** | Pointer to **NullableString** | Test plan description | [optional] 
**Build** | Pointer to **NullableString** | Build of the application on which test plan is executed | [optional] 
**ProjectId** | **string** | Project unique identifier | 
**ProductName** | Pointer to **NullableString** | Name of the testing product | [optional] 
**HasAutomaticDurationTimer** | Pointer to **NullableBool** | Boolean flag defines if test plan has automatic duration timer | [optional] 
**Attributes** | **map[string]interface{}** | Key value pair of test plan custom attributes | 
**TestSuite** | Pointer to [**NullableTestSuiteTestPlanApiModel**](TestSuiteTestPlanApiModel.md) |  | [optional] 

## Methods

### NewCreateTestPlanApiModel

`func NewCreateTestPlanApiModel(name string, projectId string, attributes map[string]interface{}, ) *CreateTestPlanApiModel`

NewCreateTestPlanApiModel instantiates a new CreateTestPlanApiModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTestPlanApiModelWithDefaults

`func NewCreateTestPlanApiModelWithDefaults() *CreateTestPlanApiModel`

NewCreateTestPlanApiModelWithDefaults instantiates a new CreateTestPlanApiModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTags

`func (o *CreateTestPlanApiModel) GetTags() []TagPostModel`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreateTestPlanApiModel) GetTagsOk() (*[]TagPostModel, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreateTestPlanApiModel) SetTags(v []TagPostModel)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CreateTestPlanApiModel) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *CreateTestPlanApiModel) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *CreateTestPlanApiModel) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetName

`func (o *CreateTestPlanApiModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateTestPlanApiModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateTestPlanApiModel) SetName(v string)`

SetName sets Name field to given value.


### GetStartDate

`func (o *CreateTestPlanApiModel) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *CreateTestPlanApiModel) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *CreateTestPlanApiModel) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *CreateTestPlanApiModel) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### SetStartDateNil

`func (o *CreateTestPlanApiModel) SetStartDateNil(b bool)`

 SetStartDateNil sets the value for StartDate to be an explicit nil

### UnsetStartDate
`func (o *CreateTestPlanApiModel) UnsetStartDate()`

UnsetStartDate ensures that no value is present for StartDate, not even an explicit nil
### GetEndDate

`func (o *CreateTestPlanApiModel) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *CreateTestPlanApiModel) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *CreateTestPlanApiModel) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *CreateTestPlanApiModel) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### SetEndDateNil

`func (o *CreateTestPlanApiModel) SetEndDateNil(b bool)`

 SetEndDateNil sets the value for EndDate to be an explicit nil

### UnsetEndDate
`func (o *CreateTestPlanApiModel) UnsetEndDate()`

UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
### GetDescription

`func (o *CreateTestPlanApiModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateTestPlanApiModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateTestPlanApiModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateTestPlanApiModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CreateTestPlanApiModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CreateTestPlanApiModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetBuild

`func (o *CreateTestPlanApiModel) GetBuild() string`

GetBuild returns the Build field if non-nil, zero value otherwise.

### GetBuildOk

`func (o *CreateTestPlanApiModel) GetBuildOk() (*string, bool)`

GetBuildOk returns a tuple with the Build field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuild

`func (o *CreateTestPlanApiModel) SetBuild(v string)`

SetBuild sets Build field to given value.

### HasBuild

`func (o *CreateTestPlanApiModel) HasBuild() bool`

HasBuild returns a boolean if a field has been set.

### SetBuildNil

`func (o *CreateTestPlanApiModel) SetBuildNil(b bool)`

 SetBuildNil sets the value for Build to be an explicit nil

### UnsetBuild
`func (o *CreateTestPlanApiModel) UnsetBuild()`

UnsetBuild ensures that no value is present for Build, not even an explicit nil
### GetProjectId

`func (o *CreateTestPlanApiModel) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *CreateTestPlanApiModel) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *CreateTestPlanApiModel) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetProductName

`func (o *CreateTestPlanApiModel) GetProductName() string`

GetProductName returns the ProductName field if non-nil, zero value otherwise.

### GetProductNameOk

`func (o *CreateTestPlanApiModel) GetProductNameOk() (*string, bool)`

GetProductNameOk returns a tuple with the ProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductName

`func (o *CreateTestPlanApiModel) SetProductName(v string)`

SetProductName sets ProductName field to given value.

### HasProductName

`func (o *CreateTestPlanApiModel) HasProductName() bool`

HasProductName returns a boolean if a field has been set.

### SetProductNameNil

`func (o *CreateTestPlanApiModel) SetProductNameNil(b bool)`

 SetProductNameNil sets the value for ProductName to be an explicit nil

### UnsetProductName
`func (o *CreateTestPlanApiModel) UnsetProductName()`

UnsetProductName ensures that no value is present for ProductName, not even an explicit nil
### GetHasAutomaticDurationTimer

`func (o *CreateTestPlanApiModel) GetHasAutomaticDurationTimer() bool`

GetHasAutomaticDurationTimer returns the HasAutomaticDurationTimer field if non-nil, zero value otherwise.

### GetHasAutomaticDurationTimerOk

`func (o *CreateTestPlanApiModel) GetHasAutomaticDurationTimerOk() (*bool, bool)`

GetHasAutomaticDurationTimerOk returns a tuple with the HasAutomaticDurationTimer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAutomaticDurationTimer

`func (o *CreateTestPlanApiModel) SetHasAutomaticDurationTimer(v bool)`

SetHasAutomaticDurationTimer sets HasAutomaticDurationTimer field to given value.

### HasHasAutomaticDurationTimer

`func (o *CreateTestPlanApiModel) HasHasAutomaticDurationTimer() bool`

HasHasAutomaticDurationTimer returns a boolean if a field has been set.

### SetHasAutomaticDurationTimerNil

`func (o *CreateTestPlanApiModel) SetHasAutomaticDurationTimerNil(b bool)`

 SetHasAutomaticDurationTimerNil sets the value for HasAutomaticDurationTimer to be an explicit nil

### UnsetHasAutomaticDurationTimer
`func (o *CreateTestPlanApiModel) UnsetHasAutomaticDurationTimer()`

UnsetHasAutomaticDurationTimer ensures that no value is present for HasAutomaticDurationTimer, not even an explicit nil
### GetAttributes

`func (o *CreateTestPlanApiModel) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *CreateTestPlanApiModel) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *CreateTestPlanApiModel) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.


### GetTestSuite

`func (o *CreateTestPlanApiModel) GetTestSuite() TestSuiteTestPlanApiModel`

GetTestSuite returns the TestSuite field if non-nil, zero value otherwise.

### GetTestSuiteOk

`func (o *CreateTestPlanApiModel) GetTestSuiteOk() (*TestSuiteTestPlanApiModel, bool)`

GetTestSuiteOk returns a tuple with the TestSuite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestSuite

`func (o *CreateTestPlanApiModel) SetTestSuite(v TestSuiteTestPlanApiModel)`

SetTestSuite sets TestSuite field to given value.

### HasTestSuite

`func (o *CreateTestPlanApiModel) HasTestSuite() bool`

HasTestSuite returns a boolean if a field has been set.

### SetTestSuiteNil

`func (o *CreateTestPlanApiModel) SetTestSuiteNil(b bool)`

 SetTestSuiteNil sets the value for TestSuite to be an explicit nil

### UnsetTestSuite
`func (o *CreateTestPlanApiModel) UnsetTestSuite()`

UnsetTestSuite ensures that no value is present for TestSuite, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


