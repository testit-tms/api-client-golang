# GetXlsxTestPointsByTestPlanApiModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IncludeName** | **bool** |  | 
**IncludeSection** | **bool** |  | 
**IncludePriority** | **bool** |  | 
**IncludeSourceType** | **bool** |  | 
**IncludeAutomated** | **bool** |  | 
**IncludeStatus** | **bool** |  | 
**IncludeDuration** | **bool** |  | 
**IncludeCreationDate** | **bool** |  | 
**IncludeAuthor** | **bool** |  | 
**IncludeModificationDate** | **bool** |  | 
**IncludeModifiedBy** | **bool** |  | 
**IncludeTags** | **bool** |  | 
**IncludeIterations** | **bool** |  | 
**CustomAttributesIds** | Pointer to **[]string** |  | [optional] 
**Filter** | Pointer to [**NullableTestPlanTestPointsSearchApiModel**](TestPlanTestPointsSearchApiModel.md) |  | [optional] 
**Order** | [**[]Order**](Order.md) |  | 
**ConfigurationIds** | Pointer to **[]string** |  | [optional] 

## Methods

### NewGetXlsxTestPointsByTestPlanApiModel

`func NewGetXlsxTestPointsByTestPlanApiModel(includeName bool, includeSection bool, includePriority bool, includeSourceType bool, includeAutomated bool, includeStatus bool, includeDuration bool, includeCreationDate bool, includeAuthor bool, includeModificationDate bool, includeModifiedBy bool, includeTags bool, includeIterations bool, order []Order, ) *GetXlsxTestPointsByTestPlanApiModel`

NewGetXlsxTestPointsByTestPlanApiModel instantiates a new GetXlsxTestPointsByTestPlanApiModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetXlsxTestPointsByTestPlanApiModelWithDefaults

`func NewGetXlsxTestPointsByTestPlanApiModelWithDefaults() *GetXlsxTestPointsByTestPlanApiModel`

NewGetXlsxTestPointsByTestPlanApiModelWithDefaults instantiates a new GetXlsxTestPointsByTestPlanApiModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncludeName

`func (o *GetXlsxTestPointsByTestPlanApiModel) GetIncludeName() bool`

GetIncludeName returns the IncludeName field if non-nil, zero value otherwise.

### GetIncludeNameOk

`func (o *GetXlsxTestPointsByTestPlanApiModel) GetIncludeNameOk() (*bool, bool)`

GetIncludeNameOk returns a tuple with the IncludeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeName

`func (o *GetXlsxTestPointsByTestPlanApiModel) SetIncludeName(v bool)`

SetIncludeName sets IncludeName field to given value.


### GetIncludeSection

`func (o *GetXlsxTestPointsByTestPlanApiModel) GetIncludeSection() bool`

GetIncludeSection returns the IncludeSection field if non-nil, zero value otherwise.

### GetIncludeSectionOk

`func (o *GetXlsxTestPointsByTestPlanApiModel) GetIncludeSectionOk() (*bool, bool)`

GetIncludeSectionOk returns a tuple with the IncludeSection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeSection

`func (o *GetXlsxTestPointsByTestPlanApiModel) SetIncludeSection(v bool)`

SetIncludeSection sets IncludeSection field to given value.


### GetIncludePriority

`func (o *GetXlsxTestPointsByTestPlanApiModel) GetIncludePriority() bool`

GetIncludePriority returns the IncludePriority field if non-nil, zero value otherwise.

### GetIncludePriorityOk

`func (o *GetXlsxTestPointsByTestPlanApiModel) GetIncludePriorityOk() (*bool, bool)`

GetIncludePriorityOk returns a tuple with the IncludePriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludePriority

`func (o *GetXlsxTestPointsByTestPlanApiModel) SetIncludePriority(v bool)`

SetIncludePriority sets IncludePriority field to given value.


### GetIncludeSourceType

`func (o *GetXlsxTestPointsByTestPlanApiModel) GetIncludeSourceType() bool`

GetIncludeSourceType returns the IncludeSourceType field if non-nil, zero value otherwise.

### GetIncludeSourceTypeOk

`func (o *GetXlsxTestPointsByTestPlanApiModel) GetIncludeSourceTypeOk() (*bool, bool)`

GetIncludeSourceTypeOk returns a tuple with the IncludeSourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeSourceType

`func (o *GetXlsxTestPointsByTestPlanApiModel) SetIncludeSourceType(v bool)`

SetIncludeSourceType sets IncludeSourceType field to given value.


### GetIncludeAutomated

`func (o *GetXlsxTestPointsByTestPlanApiModel) GetIncludeAutomated() bool`

GetIncludeAutomated returns the IncludeAutomated field if non-nil, zero value otherwise.

### GetIncludeAutomatedOk

`func (o *GetXlsxTestPointsByTestPlanApiModel) GetIncludeAutomatedOk() (*bool, bool)`

GetIncludeAutomatedOk returns a tuple with the IncludeAutomated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAutomated

`func (o *GetXlsxTestPointsByTestPlanApiModel) SetIncludeAutomated(v bool)`

SetIncludeAutomated sets IncludeAutomated field to given value.


### GetIncludeStatus

`func (o *GetXlsxTestPointsByTestPlanApiModel) GetIncludeStatus() bool`

GetIncludeStatus returns the IncludeStatus field if non-nil, zero value otherwise.

### GetIncludeStatusOk

`func (o *GetXlsxTestPointsByTestPlanApiModel) GetIncludeStatusOk() (*bool, bool)`

GetIncludeStatusOk returns a tuple with the IncludeStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeStatus

`func (o *GetXlsxTestPointsByTestPlanApiModel) SetIncludeStatus(v bool)`

SetIncludeStatus sets IncludeStatus field to given value.


### GetIncludeDuration

`func (o *GetXlsxTestPointsByTestPlanApiModel) GetIncludeDuration() bool`

GetIncludeDuration returns the IncludeDuration field if non-nil, zero value otherwise.

### GetIncludeDurationOk

`func (o *GetXlsxTestPointsByTestPlanApiModel) GetIncludeDurationOk() (*bool, bool)`

GetIncludeDurationOk returns a tuple with the IncludeDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeDuration

`func (o *GetXlsxTestPointsByTestPlanApiModel) SetIncludeDuration(v bool)`

SetIncludeDuration sets IncludeDuration field to given value.


### GetIncludeCreationDate

`func (o *GetXlsxTestPointsByTestPlanApiModel) GetIncludeCreationDate() bool`

GetIncludeCreationDate returns the IncludeCreationDate field if non-nil, zero value otherwise.

### GetIncludeCreationDateOk

`func (o *GetXlsxTestPointsByTestPlanApiModel) GetIncludeCreationDateOk() (*bool, bool)`

GetIncludeCreationDateOk returns a tuple with the IncludeCreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeCreationDate

`func (o *GetXlsxTestPointsByTestPlanApiModel) SetIncludeCreationDate(v bool)`

SetIncludeCreationDate sets IncludeCreationDate field to given value.


### GetIncludeAuthor

`func (o *GetXlsxTestPointsByTestPlanApiModel) GetIncludeAuthor() bool`

GetIncludeAuthor returns the IncludeAuthor field if non-nil, zero value otherwise.

### GetIncludeAuthorOk

`func (o *GetXlsxTestPointsByTestPlanApiModel) GetIncludeAuthorOk() (*bool, bool)`

GetIncludeAuthorOk returns a tuple with the IncludeAuthor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAuthor

`func (o *GetXlsxTestPointsByTestPlanApiModel) SetIncludeAuthor(v bool)`

SetIncludeAuthor sets IncludeAuthor field to given value.


### GetIncludeModificationDate

`func (o *GetXlsxTestPointsByTestPlanApiModel) GetIncludeModificationDate() bool`

GetIncludeModificationDate returns the IncludeModificationDate field if non-nil, zero value otherwise.

### GetIncludeModificationDateOk

`func (o *GetXlsxTestPointsByTestPlanApiModel) GetIncludeModificationDateOk() (*bool, bool)`

GetIncludeModificationDateOk returns a tuple with the IncludeModificationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeModificationDate

`func (o *GetXlsxTestPointsByTestPlanApiModel) SetIncludeModificationDate(v bool)`

SetIncludeModificationDate sets IncludeModificationDate field to given value.


### GetIncludeModifiedBy

`func (o *GetXlsxTestPointsByTestPlanApiModel) GetIncludeModifiedBy() bool`

GetIncludeModifiedBy returns the IncludeModifiedBy field if non-nil, zero value otherwise.

### GetIncludeModifiedByOk

`func (o *GetXlsxTestPointsByTestPlanApiModel) GetIncludeModifiedByOk() (*bool, bool)`

GetIncludeModifiedByOk returns a tuple with the IncludeModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeModifiedBy

`func (o *GetXlsxTestPointsByTestPlanApiModel) SetIncludeModifiedBy(v bool)`

SetIncludeModifiedBy sets IncludeModifiedBy field to given value.


### GetIncludeTags

`func (o *GetXlsxTestPointsByTestPlanApiModel) GetIncludeTags() bool`

GetIncludeTags returns the IncludeTags field if non-nil, zero value otherwise.

### GetIncludeTagsOk

`func (o *GetXlsxTestPointsByTestPlanApiModel) GetIncludeTagsOk() (*bool, bool)`

GetIncludeTagsOk returns a tuple with the IncludeTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeTags

`func (o *GetXlsxTestPointsByTestPlanApiModel) SetIncludeTags(v bool)`

SetIncludeTags sets IncludeTags field to given value.


### GetIncludeIterations

`func (o *GetXlsxTestPointsByTestPlanApiModel) GetIncludeIterations() bool`

GetIncludeIterations returns the IncludeIterations field if non-nil, zero value otherwise.

### GetIncludeIterationsOk

`func (o *GetXlsxTestPointsByTestPlanApiModel) GetIncludeIterationsOk() (*bool, bool)`

GetIncludeIterationsOk returns a tuple with the IncludeIterations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeIterations

`func (o *GetXlsxTestPointsByTestPlanApiModel) SetIncludeIterations(v bool)`

SetIncludeIterations sets IncludeIterations field to given value.


### GetCustomAttributesIds

`func (o *GetXlsxTestPointsByTestPlanApiModel) GetCustomAttributesIds() []string`

GetCustomAttributesIds returns the CustomAttributesIds field if non-nil, zero value otherwise.

### GetCustomAttributesIdsOk

`func (o *GetXlsxTestPointsByTestPlanApiModel) GetCustomAttributesIdsOk() (*[]string, bool)`

GetCustomAttributesIdsOk returns a tuple with the CustomAttributesIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomAttributesIds

`func (o *GetXlsxTestPointsByTestPlanApiModel) SetCustomAttributesIds(v []string)`

SetCustomAttributesIds sets CustomAttributesIds field to given value.

### HasCustomAttributesIds

`func (o *GetXlsxTestPointsByTestPlanApiModel) HasCustomAttributesIds() bool`

HasCustomAttributesIds returns a boolean if a field has been set.

### SetCustomAttributesIdsNil

`func (o *GetXlsxTestPointsByTestPlanApiModel) SetCustomAttributesIdsNil(b bool)`

 SetCustomAttributesIdsNil sets the value for CustomAttributesIds to be an explicit nil

### UnsetCustomAttributesIds
`func (o *GetXlsxTestPointsByTestPlanApiModel) UnsetCustomAttributesIds()`

UnsetCustomAttributesIds ensures that no value is present for CustomAttributesIds, not even an explicit nil
### GetFilter

`func (o *GetXlsxTestPointsByTestPlanApiModel) GetFilter() TestPlanTestPointsSearchApiModel`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *GetXlsxTestPointsByTestPlanApiModel) GetFilterOk() (*TestPlanTestPointsSearchApiModel, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *GetXlsxTestPointsByTestPlanApiModel) SetFilter(v TestPlanTestPointsSearchApiModel)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *GetXlsxTestPointsByTestPlanApiModel) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### SetFilterNil

`func (o *GetXlsxTestPointsByTestPlanApiModel) SetFilterNil(b bool)`

 SetFilterNil sets the value for Filter to be an explicit nil

### UnsetFilter
`func (o *GetXlsxTestPointsByTestPlanApiModel) UnsetFilter()`

UnsetFilter ensures that no value is present for Filter, not even an explicit nil
### GetOrder

`func (o *GetXlsxTestPointsByTestPlanApiModel) GetOrder() []Order`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *GetXlsxTestPointsByTestPlanApiModel) GetOrderOk() (*[]Order, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *GetXlsxTestPointsByTestPlanApiModel) SetOrder(v []Order)`

SetOrder sets Order field to given value.


### GetConfigurationIds

`func (o *GetXlsxTestPointsByTestPlanApiModel) GetConfigurationIds() []string`

GetConfigurationIds returns the ConfigurationIds field if non-nil, zero value otherwise.

### GetConfigurationIdsOk

`func (o *GetXlsxTestPointsByTestPlanApiModel) GetConfigurationIdsOk() (*[]string, bool)`

GetConfigurationIdsOk returns a tuple with the ConfigurationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationIds

`func (o *GetXlsxTestPointsByTestPlanApiModel) SetConfigurationIds(v []string)`

SetConfigurationIds sets ConfigurationIds field to given value.

### HasConfigurationIds

`func (o *GetXlsxTestPointsByTestPlanApiModel) HasConfigurationIds() bool`

HasConfigurationIds returns a boolean if a field has been set.

### SetConfigurationIdsNil

`func (o *GetXlsxTestPointsByTestPlanApiModel) SetConfigurationIdsNil(b bool)`

 SetConfigurationIdsNil sets the value for ConfigurationIds to be an explicit nil

### UnsetConfigurationIds
`func (o *GetXlsxTestPointsByTestPlanApiModel) UnsetConfigurationIds()`

UnsetConfigurationIds ensures that no value is present for ConfigurationIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


