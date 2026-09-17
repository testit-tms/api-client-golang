# SearchConfigurationParametersApiModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectIds** | Pointer to **[]string** |  | [optional] 
**Inquiry** | Pointer to [**NullableInquiry**](Inquiry.md) |  | [optional] 
**ValuesFilters** | Pointer to [**[]FieldFilter**](FieldFilter.md) |  | [optional] 
**ProjectsFilters** | Pointer to [**[]FieldFilter**](FieldFilter.md) |  | [optional] 

## Methods

### NewSearchConfigurationParametersApiModel

`func NewSearchConfigurationParametersApiModel() *SearchConfigurationParametersApiModel`

NewSearchConfigurationParametersApiModel instantiates a new SearchConfigurationParametersApiModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchConfigurationParametersApiModelWithDefaults

`func NewSearchConfigurationParametersApiModelWithDefaults() *SearchConfigurationParametersApiModel`

NewSearchConfigurationParametersApiModelWithDefaults instantiates a new SearchConfigurationParametersApiModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectIds

`func (o *SearchConfigurationParametersApiModel) GetProjectIds() []string`

GetProjectIds returns the ProjectIds field if non-nil, zero value otherwise.

### GetProjectIdsOk

`func (o *SearchConfigurationParametersApiModel) GetProjectIdsOk() (*[]string, bool)`

GetProjectIdsOk returns a tuple with the ProjectIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectIds

`func (o *SearchConfigurationParametersApiModel) SetProjectIds(v []string)`

SetProjectIds sets ProjectIds field to given value.

### HasProjectIds

`func (o *SearchConfigurationParametersApiModel) HasProjectIds() bool`

HasProjectIds returns a boolean if a field has been set.

### SetProjectIdsNil

`func (o *SearchConfigurationParametersApiModel) SetProjectIdsNil(b bool)`

 SetProjectIdsNil sets the value for ProjectIds to be an explicit nil

### UnsetProjectIds
`func (o *SearchConfigurationParametersApiModel) UnsetProjectIds()`

UnsetProjectIds ensures that no value is present for ProjectIds, not even an explicit nil
### GetInquiry

`func (o *SearchConfigurationParametersApiModel) GetInquiry() Inquiry`

GetInquiry returns the Inquiry field if non-nil, zero value otherwise.

### GetInquiryOk

`func (o *SearchConfigurationParametersApiModel) GetInquiryOk() (*Inquiry, bool)`

GetInquiryOk returns a tuple with the Inquiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInquiry

`func (o *SearchConfigurationParametersApiModel) SetInquiry(v Inquiry)`

SetInquiry sets Inquiry field to given value.

### HasInquiry

`func (o *SearchConfigurationParametersApiModel) HasInquiry() bool`

HasInquiry returns a boolean if a field has been set.

### SetInquiryNil

`func (o *SearchConfigurationParametersApiModel) SetInquiryNil(b bool)`

 SetInquiryNil sets the value for Inquiry to be an explicit nil

### UnsetInquiry
`func (o *SearchConfigurationParametersApiModel) UnsetInquiry()`

UnsetInquiry ensures that no value is present for Inquiry, not even an explicit nil
### GetValuesFilters

`func (o *SearchConfigurationParametersApiModel) GetValuesFilters() []FieldFilter`

GetValuesFilters returns the ValuesFilters field if non-nil, zero value otherwise.

### GetValuesFiltersOk

`func (o *SearchConfigurationParametersApiModel) GetValuesFiltersOk() (*[]FieldFilter, bool)`

GetValuesFiltersOk returns a tuple with the ValuesFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValuesFilters

`func (o *SearchConfigurationParametersApiModel) SetValuesFilters(v []FieldFilter)`

SetValuesFilters sets ValuesFilters field to given value.

### HasValuesFilters

`func (o *SearchConfigurationParametersApiModel) HasValuesFilters() bool`

HasValuesFilters returns a boolean if a field has been set.

### SetValuesFiltersNil

`func (o *SearchConfigurationParametersApiModel) SetValuesFiltersNil(b bool)`

 SetValuesFiltersNil sets the value for ValuesFilters to be an explicit nil

### UnsetValuesFilters
`func (o *SearchConfigurationParametersApiModel) UnsetValuesFilters()`

UnsetValuesFilters ensures that no value is present for ValuesFilters, not even an explicit nil
### GetProjectsFilters

`func (o *SearchConfigurationParametersApiModel) GetProjectsFilters() []FieldFilter`

GetProjectsFilters returns the ProjectsFilters field if non-nil, zero value otherwise.

### GetProjectsFiltersOk

`func (o *SearchConfigurationParametersApiModel) GetProjectsFiltersOk() (*[]FieldFilter, bool)`

GetProjectsFiltersOk returns a tuple with the ProjectsFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectsFilters

`func (o *SearchConfigurationParametersApiModel) SetProjectsFilters(v []FieldFilter)`

SetProjectsFilters sets ProjectsFilters field to given value.

### HasProjectsFilters

`func (o *SearchConfigurationParametersApiModel) HasProjectsFilters() bool`

HasProjectsFilters returns a boolean if a field has been set.

### SetProjectsFiltersNil

`func (o *SearchConfigurationParametersApiModel) SetProjectsFiltersNil(b bool)`

 SetProjectsFiltersNil sets the value for ProjectsFilters to be an explicit nil

### UnsetProjectsFilters
`func (o *SearchConfigurationParametersApiModel) UnsetProjectsFilters()`

UnsetProjectsFilters ensures that no value is present for ProjectsFilters, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


