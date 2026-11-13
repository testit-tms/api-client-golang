# CreateTestRunAndFillByAutoTestCasesApiModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | **string** | Specifies the GUID of the project, in which a test run will be created. | 
**Filter** | Pointer to [**NullableCompositeFilter**](CompositeFilter.md) | Specifies the filter for selecting autotests, from which test points are created. | [optional] 
**Name** | Pointer to **NullableString** | Specifies the name of the test run. | [optional] 
**ConfigurationIds** | **[]string** | Specifies the configuration GUIDs, from which test points are created. You can specify several GUIDs. | 
**Description** | Pointer to **NullableString** | Specifies the test run description. | [optional] 
**LaunchSource** | Pointer to **NullableString** | Specifies the test run launch source. | [optional] 
**Option** | [**TestRunLaunchOptionApiModel**](TestRunLaunchOptionApiModel.md) | Specifies the test run launch options. | 
**Tags** | Pointer to **[]string** | Collection of tags to assign to the test run | [optional] 

## Methods

### NewCreateTestRunAndFillByAutoTestCasesApiModel

`func NewCreateTestRunAndFillByAutoTestCasesApiModel(projectId string, configurationIds []string, option TestRunLaunchOptionApiModel, ) *CreateTestRunAndFillByAutoTestCasesApiModel`

NewCreateTestRunAndFillByAutoTestCasesApiModel instantiates a new CreateTestRunAndFillByAutoTestCasesApiModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTestRunAndFillByAutoTestCasesApiModelWithDefaults

`func NewCreateTestRunAndFillByAutoTestCasesApiModelWithDefaults() *CreateTestRunAndFillByAutoTestCasesApiModel`

NewCreateTestRunAndFillByAutoTestCasesApiModelWithDefaults instantiates a new CreateTestRunAndFillByAutoTestCasesApiModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *CreateTestRunAndFillByAutoTestCasesApiModel) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *CreateTestRunAndFillByAutoTestCasesApiModel) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *CreateTestRunAndFillByAutoTestCasesApiModel) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetFilter

`func (o *CreateTestRunAndFillByAutoTestCasesApiModel) GetFilter() CompositeFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *CreateTestRunAndFillByAutoTestCasesApiModel) GetFilterOk() (*CompositeFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *CreateTestRunAndFillByAutoTestCasesApiModel) SetFilter(v CompositeFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *CreateTestRunAndFillByAutoTestCasesApiModel) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### SetFilterNil

`func (o *CreateTestRunAndFillByAutoTestCasesApiModel) SetFilterNil(b bool)`

 SetFilterNil sets the value for Filter to be an explicit nil

### UnsetFilter
`func (o *CreateTestRunAndFillByAutoTestCasesApiModel) UnsetFilter()`

UnsetFilter ensures that no value is present for Filter, not even an explicit nil
### GetName

`func (o *CreateTestRunAndFillByAutoTestCasesApiModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateTestRunAndFillByAutoTestCasesApiModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateTestRunAndFillByAutoTestCasesApiModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateTestRunAndFillByAutoTestCasesApiModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CreateTestRunAndFillByAutoTestCasesApiModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CreateTestRunAndFillByAutoTestCasesApiModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetConfigurationIds

`func (o *CreateTestRunAndFillByAutoTestCasesApiModel) GetConfigurationIds() []string`

GetConfigurationIds returns the ConfigurationIds field if non-nil, zero value otherwise.

### GetConfigurationIdsOk

`func (o *CreateTestRunAndFillByAutoTestCasesApiModel) GetConfigurationIdsOk() (*[]string, bool)`

GetConfigurationIdsOk returns a tuple with the ConfigurationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationIds

`func (o *CreateTestRunAndFillByAutoTestCasesApiModel) SetConfigurationIds(v []string)`

SetConfigurationIds sets ConfigurationIds field to given value.


### GetDescription

`func (o *CreateTestRunAndFillByAutoTestCasesApiModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateTestRunAndFillByAutoTestCasesApiModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateTestRunAndFillByAutoTestCasesApiModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateTestRunAndFillByAutoTestCasesApiModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CreateTestRunAndFillByAutoTestCasesApiModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CreateTestRunAndFillByAutoTestCasesApiModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetLaunchSource

`func (o *CreateTestRunAndFillByAutoTestCasesApiModel) GetLaunchSource() string`

GetLaunchSource returns the LaunchSource field if non-nil, zero value otherwise.

### GetLaunchSourceOk

`func (o *CreateTestRunAndFillByAutoTestCasesApiModel) GetLaunchSourceOk() (*string, bool)`

GetLaunchSourceOk returns a tuple with the LaunchSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLaunchSource

`func (o *CreateTestRunAndFillByAutoTestCasesApiModel) SetLaunchSource(v string)`

SetLaunchSource sets LaunchSource field to given value.

### HasLaunchSource

`func (o *CreateTestRunAndFillByAutoTestCasesApiModel) HasLaunchSource() bool`

HasLaunchSource returns a boolean if a field has been set.

### SetLaunchSourceNil

`func (o *CreateTestRunAndFillByAutoTestCasesApiModel) SetLaunchSourceNil(b bool)`

 SetLaunchSourceNil sets the value for LaunchSource to be an explicit nil

### UnsetLaunchSource
`func (o *CreateTestRunAndFillByAutoTestCasesApiModel) UnsetLaunchSource()`

UnsetLaunchSource ensures that no value is present for LaunchSource, not even an explicit nil
### GetOption

`func (o *CreateTestRunAndFillByAutoTestCasesApiModel) GetOption() TestRunLaunchOptionApiModel`

GetOption returns the Option field if non-nil, zero value otherwise.

### GetOptionOk

`func (o *CreateTestRunAndFillByAutoTestCasesApiModel) GetOptionOk() (*TestRunLaunchOptionApiModel, bool)`

GetOptionOk returns a tuple with the Option field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOption

`func (o *CreateTestRunAndFillByAutoTestCasesApiModel) SetOption(v TestRunLaunchOptionApiModel)`

SetOption sets Option field to given value.


### GetTags

`func (o *CreateTestRunAndFillByAutoTestCasesApiModel) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreateTestRunAndFillByAutoTestCasesApiModel) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreateTestRunAndFillByAutoTestCasesApiModel) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CreateTestRunAndFillByAutoTestCasesApiModel) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *CreateTestRunAndFillByAutoTestCasesApiModel) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *CreateTestRunAndFillByAutoTestCasesApiModel) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


