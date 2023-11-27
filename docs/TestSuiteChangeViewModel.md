# TestSuiteChangeViewModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Configurations** | Pointer to [**[]ShortConfiguration**](ShortConfiguration.md) |  | [optional] 
**WorkItemCount** | **int64** |  | 

## Methods

### NewTestSuiteChangeViewModel

`func NewTestSuiteChangeViewModel(id string, name string, workItemCount int64, ) *TestSuiteChangeViewModel`

NewTestSuiteChangeViewModel instantiates a new TestSuiteChangeViewModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestSuiteChangeViewModelWithDefaults

`func NewTestSuiteChangeViewModelWithDefaults() *TestSuiteChangeViewModel`

NewTestSuiteChangeViewModelWithDefaults instantiates a new TestSuiteChangeViewModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TestSuiteChangeViewModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TestSuiteChangeViewModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TestSuiteChangeViewModel) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *TestSuiteChangeViewModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TestSuiteChangeViewModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TestSuiteChangeViewModel) SetName(v string)`

SetName sets Name field to given value.


### GetConfigurations

`func (o *TestSuiteChangeViewModel) GetConfigurations() []ShortConfiguration`

GetConfigurations returns the Configurations field if non-nil, zero value otherwise.

### GetConfigurationsOk

`func (o *TestSuiteChangeViewModel) GetConfigurationsOk() (*[]ShortConfiguration, bool)`

GetConfigurationsOk returns a tuple with the Configurations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurations

`func (o *TestSuiteChangeViewModel) SetConfigurations(v []ShortConfiguration)`

SetConfigurations sets Configurations field to given value.

### HasConfigurations

`func (o *TestSuiteChangeViewModel) HasConfigurations() bool`

HasConfigurations returns a boolean if a field has been set.

### SetConfigurationsNil

`func (o *TestSuiteChangeViewModel) SetConfigurationsNil(b bool)`

 SetConfigurationsNil sets the value for Configurations to be an explicit nil

### UnsetConfigurations
`func (o *TestSuiteChangeViewModel) UnsetConfigurations()`

UnsetConfigurations ensures that no value is present for Configurations, not even an explicit nil
### GetWorkItemCount

`func (o *TestSuiteChangeViewModel) GetWorkItemCount() int64`

GetWorkItemCount returns the WorkItemCount field if non-nil, zero value otherwise.

### GetWorkItemCountOk

`func (o *TestSuiteChangeViewModel) GetWorkItemCountOk() (*int64, bool)`

GetWorkItemCountOk returns a tuple with the WorkItemCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkItemCount

`func (o *TestSuiteChangeViewModel) SetWorkItemCount(v int64)`

SetWorkItemCount sets WorkItemCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


