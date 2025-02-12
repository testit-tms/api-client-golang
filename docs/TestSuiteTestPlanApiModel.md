# TestSuiteTestPlanApiModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Test suite nane | 
**ConfigurationIds** | Pointer to **[]string** | Configuration identifiers. Empty configurations means using default configurations | [optional] 
**Type** | Pointer to [**NullableTestSuiteType**](TestSuiteType.md) | Type of the test suite | [optional] 
**SaveStructure** | Pointer to **NullableBool** | Indicates if the test suite retains section tree structure | [optional] 
**WorkItemsSelector** | [**WorkItemSelectModel**](WorkItemSelectModel.md) | Model containing options to filter work items | 

## Methods

### NewTestSuiteTestPlanApiModel

`func NewTestSuiteTestPlanApiModel(name string, workItemsSelector WorkItemSelectModel, ) *TestSuiteTestPlanApiModel`

NewTestSuiteTestPlanApiModel instantiates a new TestSuiteTestPlanApiModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestSuiteTestPlanApiModelWithDefaults

`func NewTestSuiteTestPlanApiModelWithDefaults() *TestSuiteTestPlanApiModel`

NewTestSuiteTestPlanApiModelWithDefaults instantiates a new TestSuiteTestPlanApiModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TestSuiteTestPlanApiModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TestSuiteTestPlanApiModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TestSuiteTestPlanApiModel) SetName(v string)`

SetName sets Name field to given value.


### GetConfigurationIds

`func (o *TestSuiteTestPlanApiModel) GetConfigurationIds() []string`

GetConfigurationIds returns the ConfigurationIds field if non-nil, zero value otherwise.

### GetConfigurationIdsOk

`func (o *TestSuiteTestPlanApiModel) GetConfigurationIdsOk() (*[]string, bool)`

GetConfigurationIdsOk returns a tuple with the ConfigurationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationIds

`func (o *TestSuiteTestPlanApiModel) SetConfigurationIds(v []string)`

SetConfigurationIds sets ConfigurationIds field to given value.

### HasConfigurationIds

`func (o *TestSuiteTestPlanApiModel) HasConfigurationIds() bool`

HasConfigurationIds returns a boolean if a field has been set.

### SetConfigurationIdsNil

`func (o *TestSuiteTestPlanApiModel) SetConfigurationIdsNil(b bool)`

 SetConfigurationIdsNil sets the value for ConfigurationIds to be an explicit nil

### UnsetConfigurationIds
`func (o *TestSuiteTestPlanApiModel) UnsetConfigurationIds()`

UnsetConfigurationIds ensures that no value is present for ConfigurationIds, not even an explicit nil
### GetType

`func (o *TestSuiteTestPlanApiModel) GetType() TestSuiteType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TestSuiteTestPlanApiModel) GetTypeOk() (*TestSuiteType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TestSuiteTestPlanApiModel) SetType(v TestSuiteType)`

SetType sets Type field to given value.

### HasType

`func (o *TestSuiteTestPlanApiModel) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *TestSuiteTestPlanApiModel) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *TestSuiteTestPlanApiModel) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetSaveStructure

`func (o *TestSuiteTestPlanApiModel) GetSaveStructure() bool`

GetSaveStructure returns the SaveStructure field if non-nil, zero value otherwise.

### GetSaveStructureOk

`func (o *TestSuiteTestPlanApiModel) GetSaveStructureOk() (*bool, bool)`

GetSaveStructureOk returns a tuple with the SaveStructure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaveStructure

`func (o *TestSuiteTestPlanApiModel) SetSaveStructure(v bool)`

SetSaveStructure sets SaveStructure field to given value.

### HasSaveStructure

`func (o *TestSuiteTestPlanApiModel) HasSaveStructure() bool`

HasSaveStructure returns a boolean if a field has been set.

### SetSaveStructureNil

`func (o *TestSuiteTestPlanApiModel) SetSaveStructureNil(b bool)`

 SetSaveStructureNil sets the value for SaveStructure to be an explicit nil

### UnsetSaveStructure
`func (o *TestSuiteTestPlanApiModel) UnsetSaveStructure()`

UnsetSaveStructure ensures that no value is present for SaveStructure, not even an explicit nil
### GetWorkItemsSelector

`func (o *TestSuiteTestPlanApiModel) GetWorkItemsSelector() WorkItemSelectModel`

GetWorkItemsSelector returns the WorkItemsSelector field if non-nil, zero value otherwise.

### GetWorkItemsSelectorOk

`func (o *TestSuiteTestPlanApiModel) GetWorkItemsSelectorOk() (*WorkItemSelectModel, bool)`

GetWorkItemsSelectorOk returns a tuple with the WorkItemsSelector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkItemsSelector

`func (o *TestSuiteTestPlanApiModel) SetWorkItemsSelector(v WorkItemSelectModel)`

SetWorkItemsSelector sets WorkItemsSelector field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


