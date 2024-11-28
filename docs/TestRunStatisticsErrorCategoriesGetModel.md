# TestRunStatisticsErrorCategoriesGetModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InfrastructureDefect** | **int32** | Number of test results which outcomes were caused by some infrastructure defect | 
**ProductDefect** | **int32** | Number of test results which outcomes were caused by some tested product defect | 
**TestDefect** | **int32** | Number of test results which outcomes were caused by test itself | 

## Methods

### NewTestRunStatisticsErrorCategoriesGetModel

`func NewTestRunStatisticsErrorCategoriesGetModel(infrastructureDefect int32, productDefect int32, testDefect int32, ) *TestRunStatisticsErrorCategoriesGetModel`

NewTestRunStatisticsErrorCategoriesGetModel instantiates a new TestRunStatisticsErrorCategoriesGetModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestRunStatisticsErrorCategoriesGetModelWithDefaults

`func NewTestRunStatisticsErrorCategoriesGetModelWithDefaults() *TestRunStatisticsErrorCategoriesGetModel`

NewTestRunStatisticsErrorCategoriesGetModelWithDefaults instantiates a new TestRunStatisticsErrorCategoriesGetModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInfrastructureDefect

`func (o *TestRunStatisticsErrorCategoriesGetModel) GetInfrastructureDefect() int32`

GetInfrastructureDefect returns the InfrastructureDefect field if non-nil, zero value otherwise.

### GetInfrastructureDefectOk

`func (o *TestRunStatisticsErrorCategoriesGetModel) GetInfrastructureDefectOk() (*int32, bool)`

GetInfrastructureDefectOk returns a tuple with the InfrastructureDefect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrastructureDefect

`func (o *TestRunStatisticsErrorCategoriesGetModel) SetInfrastructureDefect(v int32)`

SetInfrastructureDefect sets InfrastructureDefect field to given value.


### GetProductDefect

`func (o *TestRunStatisticsErrorCategoriesGetModel) GetProductDefect() int32`

GetProductDefect returns the ProductDefect field if non-nil, zero value otherwise.

### GetProductDefectOk

`func (o *TestRunStatisticsErrorCategoriesGetModel) GetProductDefectOk() (*int32, bool)`

GetProductDefectOk returns a tuple with the ProductDefect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductDefect

`func (o *TestRunStatisticsErrorCategoriesGetModel) SetProductDefect(v int32)`

SetProductDefect sets ProductDefect field to given value.


### GetTestDefect

`func (o *TestRunStatisticsErrorCategoriesGetModel) GetTestDefect() int32`

GetTestDefect returns the TestDefect field if non-nil, zero value otherwise.

### GetTestDefectOk

`func (o *TestRunStatisticsErrorCategoriesGetModel) GetTestDefectOk() (*int32, bool)`

GetTestDefectOk returns a tuple with the TestDefect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestDefect

`func (o *TestRunStatisticsErrorCategoriesGetModel) SetTestDefect(v int32)`

SetTestDefect sets TestDefect field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


