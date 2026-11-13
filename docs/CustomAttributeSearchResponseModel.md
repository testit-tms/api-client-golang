# CustomAttributeSearchResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkItemUsage** | [**[]ProjectShortestModel**](ProjectShortestModel.md) |  | 
**TestPlanUsage** | [**[]ProjectShortestModel**](ProjectShortestModel.md) |  | 
**Id** | **string** | Unique ID of the attribute. | 
**Code** | Pointer to **NullableString** | Optional code identifier for the attribute. | [optional] 
**Type** | [**CustomAttributeTypesEnum**](CustomAttributeTypesEnum.md) | Type of the attribute. | 
**Options** | [**[]CustomAttributeOptionModel**](CustomAttributeOptionModel.md) | Collection of the attribute options. | 
**Targets** | **[]string** | Collection of the attribute targets.   Defines where the attribute can be used (e.g., TestCases, AutoTestCases, TestPlans). | 
**IsReadOnly** | **bool** | Indicates if the attribute is read-only. | 
**IsDeleted** | **bool** | Indicates if the attribute is deleted. | 
**IsSystem** | **bool** | Indicates if the attribute is system. | 
**Name** | **string** | Name of the attribute | 
**IsEnabled** | **bool** | Indicates if the attribute is enabled | 
**IsRequired** | **bool** | Indicates if the attribute value is mandatory to specify | 
**IsGlobal** | **bool** | Indicates if the attribute is available across all projects | 

## Methods

### NewCustomAttributeSearchResponseModel

`func NewCustomAttributeSearchResponseModel(workItemUsage []ProjectShortestModel, testPlanUsage []ProjectShortestModel, id string, type_ CustomAttributeTypesEnum, options []CustomAttributeOptionModel, targets []string, isReadOnly bool, isDeleted bool, isSystem bool, name string, isEnabled bool, isRequired bool, isGlobal bool, ) *CustomAttributeSearchResponseModel`

NewCustomAttributeSearchResponseModel instantiates a new CustomAttributeSearchResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomAttributeSearchResponseModelWithDefaults

`func NewCustomAttributeSearchResponseModelWithDefaults() *CustomAttributeSearchResponseModel`

NewCustomAttributeSearchResponseModelWithDefaults instantiates a new CustomAttributeSearchResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkItemUsage

`func (o *CustomAttributeSearchResponseModel) GetWorkItemUsage() []ProjectShortestModel`

GetWorkItemUsage returns the WorkItemUsage field if non-nil, zero value otherwise.

### GetWorkItemUsageOk

`func (o *CustomAttributeSearchResponseModel) GetWorkItemUsageOk() (*[]ProjectShortestModel, bool)`

GetWorkItemUsageOk returns a tuple with the WorkItemUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkItemUsage

`func (o *CustomAttributeSearchResponseModel) SetWorkItemUsage(v []ProjectShortestModel)`

SetWorkItemUsage sets WorkItemUsage field to given value.


### GetTestPlanUsage

`func (o *CustomAttributeSearchResponseModel) GetTestPlanUsage() []ProjectShortestModel`

GetTestPlanUsage returns the TestPlanUsage field if non-nil, zero value otherwise.

### GetTestPlanUsageOk

`func (o *CustomAttributeSearchResponseModel) GetTestPlanUsageOk() (*[]ProjectShortestModel, bool)`

GetTestPlanUsageOk returns a tuple with the TestPlanUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestPlanUsage

`func (o *CustomAttributeSearchResponseModel) SetTestPlanUsage(v []ProjectShortestModel)`

SetTestPlanUsage sets TestPlanUsage field to given value.


### GetId

`func (o *CustomAttributeSearchResponseModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomAttributeSearchResponseModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomAttributeSearchResponseModel) SetId(v string)`

SetId sets Id field to given value.


### GetCode

`func (o *CustomAttributeSearchResponseModel) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *CustomAttributeSearchResponseModel) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *CustomAttributeSearchResponseModel) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *CustomAttributeSearchResponseModel) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *CustomAttributeSearchResponseModel) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *CustomAttributeSearchResponseModel) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetType

`func (o *CustomAttributeSearchResponseModel) GetType() CustomAttributeTypesEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CustomAttributeSearchResponseModel) GetTypeOk() (*CustomAttributeTypesEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CustomAttributeSearchResponseModel) SetType(v CustomAttributeTypesEnum)`

SetType sets Type field to given value.


### GetOptions

`func (o *CustomAttributeSearchResponseModel) GetOptions() []CustomAttributeOptionModel`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *CustomAttributeSearchResponseModel) GetOptionsOk() (*[]CustomAttributeOptionModel, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *CustomAttributeSearchResponseModel) SetOptions(v []CustomAttributeOptionModel)`

SetOptions sets Options field to given value.


### GetTargets

`func (o *CustomAttributeSearchResponseModel) GetTargets() []string`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *CustomAttributeSearchResponseModel) GetTargetsOk() (*[]string, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *CustomAttributeSearchResponseModel) SetTargets(v []string)`

SetTargets sets Targets field to given value.


### GetIsReadOnly

`func (o *CustomAttributeSearchResponseModel) GetIsReadOnly() bool`

GetIsReadOnly returns the IsReadOnly field if non-nil, zero value otherwise.

### GetIsReadOnlyOk

`func (o *CustomAttributeSearchResponseModel) GetIsReadOnlyOk() (*bool, bool)`

GetIsReadOnlyOk returns a tuple with the IsReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReadOnly

`func (o *CustomAttributeSearchResponseModel) SetIsReadOnly(v bool)`

SetIsReadOnly sets IsReadOnly field to given value.


### GetIsDeleted

`func (o *CustomAttributeSearchResponseModel) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *CustomAttributeSearchResponseModel) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *CustomAttributeSearchResponseModel) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.


### GetIsSystem

`func (o *CustomAttributeSearchResponseModel) GetIsSystem() bool`

GetIsSystem returns the IsSystem field if non-nil, zero value otherwise.

### GetIsSystemOk

`func (o *CustomAttributeSearchResponseModel) GetIsSystemOk() (*bool, bool)`

GetIsSystemOk returns a tuple with the IsSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSystem

`func (o *CustomAttributeSearchResponseModel) SetIsSystem(v bool)`

SetIsSystem sets IsSystem field to given value.


### GetName

`func (o *CustomAttributeSearchResponseModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomAttributeSearchResponseModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustomAttributeSearchResponseModel) SetName(v string)`

SetName sets Name field to given value.


### GetIsEnabled

`func (o *CustomAttributeSearchResponseModel) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *CustomAttributeSearchResponseModel) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *CustomAttributeSearchResponseModel) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.


### GetIsRequired

`func (o *CustomAttributeSearchResponseModel) GetIsRequired() bool`

GetIsRequired returns the IsRequired field if non-nil, zero value otherwise.

### GetIsRequiredOk

`func (o *CustomAttributeSearchResponseModel) GetIsRequiredOk() (*bool, bool)`

GetIsRequiredOk returns a tuple with the IsRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRequired

`func (o *CustomAttributeSearchResponseModel) SetIsRequired(v bool)`

SetIsRequired sets IsRequired field to given value.


### GetIsGlobal

`func (o *CustomAttributeSearchResponseModel) GetIsGlobal() bool`

GetIsGlobal returns the IsGlobal field if non-nil, zero value otherwise.

### GetIsGlobalOk

`func (o *CustomAttributeSearchResponseModel) GetIsGlobalOk() (*bool, bool)`

GetIsGlobalOk returns a tuple with the IsGlobal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGlobal

`func (o *CustomAttributeSearchResponseModel) SetIsGlobal(v bool)`

SetIsGlobal sets IsGlobal field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


