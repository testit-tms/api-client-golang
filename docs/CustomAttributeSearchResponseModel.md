# CustomAttributeSearchResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkItemUsage** | [**[]ProjectShortestModel**](ProjectShortestModel.md) |  | 
**TestPlanUsage** | [**[]ProjectShortestModel**](ProjectShortestModel.md) |  | 
**Id** | **string** | Unique ID of the attribute | 
**Options** | [**[]CustomAttributeOptionModel**](CustomAttributeOptionModel.md) | Collection of the attribute options      Available for attributes of type &#x60;options&#x60; and &#x60;multiple options&#x60; only | 
**Type** | [**CustomAttributeTypesEnum**](CustomAttributeTypesEnum.md) | Type of the attribute | 
**IsDeleted** | **bool** | Indicates if the attribute is deleted | 
**Name** | **string** | Name of the attribute | 
**IsEnabled** | **bool** | Indicates if the attribute is enabled | 
**IsRequired** | **bool** | Indicates if the attribute value is mandatory to specify | 
**IsGlobal** | **bool** | Indicates if the attribute is available across all projects | 

## Methods

### NewCustomAttributeSearchResponseModel

`func NewCustomAttributeSearchResponseModel(workItemUsage []ProjectShortestModel, testPlanUsage []ProjectShortestModel, id string, options []CustomAttributeOptionModel, type_ CustomAttributeTypesEnum, isDeleted bool, name string, isEnabled bool, isRequired bool, isGlobal bool, ) *CustomAttributeSearchResponseModel`

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


