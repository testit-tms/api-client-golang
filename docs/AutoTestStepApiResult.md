# AutoTestStepApiResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** | Step name. | 
**Description** | Pointer to **NullableString** | Detailed step description. It appears when the step is unfolded. | [optional] 
**Steps** | Pointer to [**[]AutoTestStepApiResult**](AutoTestStepApiResult.md) | Includes a nested step inside another step. The maximum nesting level is 15. | [optional] 

## Methods

### NewAutoTestStepApiResult

`func NewAutoTestStepApiResult(title string, ) *AutoTestStepApiResult`

NewAutoTestStepApiResult instantiates a new AutoTestStepApiResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoTestStepApiResultWithDefaults

`func NewAutoTestStepApiResultWithDefaults() *AutoTestStepApiResult`

NewAutoTestStepApiResultWithDefaults instantiates a new AutoTestStepApiResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *AutoTestStepApiResult) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AutoTestStepApiResult) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AutoTestStepApiResult) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *AutoTestStepApiResult) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AutoTestStepApiResult) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AutoTestStepApiResult) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AutoTestStepApiResult) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *AutoTestStepApiResult) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *AutoTestStepApiResult) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetSteps

`func (o *AutoTestStepApiResult) GetSteps() []AutoTestStepApiResult`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *AutoTestStepApiResult) GetStepsOk() (*[]AutoTestStepApiResult, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *AutoTestStepApiResult) SetSteps(v []AutoTestStepApiResult)`

SetSteps sets Steps field to given value.

### HasSteps

`func (o *AutoTestStepApiResult) HasSteps() bool`

HasSteps returns a boolean if a field has been set.

### SetStepsNil

`func (o *AutoTestStepApiResult) SetStepsNil(b bool)`

 SetStepsNil sets the value for Steps to be an explicit nil

### UnsetSteps
`func (o *AutoTestStepApiResult) UnsetSteps()`

UnsetSteps ensures that no value is present for Steps, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


