# ExportProjectJsonRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SectionIds** | Pointer to **[]string** | Specifies the IDs of the sections you want to export.&lt;br /&gt;  Use this parameter if you want to export certain parts of the project.&lt;br /&gt;  In this parameter, \&quot;&lt;b&gt;string&lt;/b&gt;\&quot; values are IDs of the test library sections. | [optional] 
**WorkItemIds** | Pointer to **[]string** | Specifies the work items you want to export from a project.&lt;br /&gt;  Use this parameter if you want to export certain work items.&lt;br /&gt;  In this parameter, \&quot;&lt;b&gt;string&lt;/b&gt;\&quot; values are IDs of the work items. | [optional] 

## Methods

### NewExportProjectJsonRequest

`func NewExportProjectJsonRequest() *ExportProjectJsonRequest`

NewExportProjectJsonRequest instantiates a new ExportProjectJsonRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExportProjectJsonRequestWithDefaults

`func NewExportProjectJsonRequestWithDefaults() *ExportProjectJsonRequest`

NewExportProjectJsonRequestWithDefaults instantiates a new ExportProjectJsonRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSectionIds

`func (o *ExportProjectJsonRequest) GetSectionIds() []string`

GetSectionIds returns the SectionIds field if non-nil, zero value otherwise.

### GetSectionIdsOk

`func (o *ExportProjectJsonRequest) GetSectionIdsOk() (*[]string, bool)`

GetSectionIdsOk returns a tuple with the SectionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectionIds

`func (o *ExportProjectJsonRequest) SetSectionIds(v []string)`

SetSectionIds sets SectionIds field to given value.

### HasSectionIds

`func (o *ExportProjectJsonRequest) HasSectionIds() bool`

HasSectionIds returns a boolean if a field has been set.

### SetSectionIdsNil

`func (o *ExportProjectJsonRequest) SetSectionIdsNil(b bool)`

 SetSectionIdsNil sets the value for SectionIds to be an explicit nil

### UnsetSectionIds
`func (o *ExportProjectJsonRequest) UnsetSectionIds()`

UnsetSectionIds ensures that no value is present for SectionIds, not even an explicit nil
### GetWorkItemIds

`func (o *ExportProjectJsonRequest) GetWorkItemIds() []string`

GetWorkItemIds returns the WorkItemIds field if non-nil, zero value otherwise.

### GetWorkItemIdsOk

`func (o *ExportProjectJsonRequest) GetWorkItemIdsOk() (*[]string, bool)`

GetWorkItemIdsOk returns a tuple with the WorkItemIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkItemIds

`func (o *ExportProjectJsonRequest) SetWorkItemIds(v []string)`

SetWorkItemIds sets WorkItemIds field to given value.

### HasWorkItemIds

`func (o *ExportProjectJsonRequest) HasWorkItemIds() bool`

HasWorkItemIds returns a boolean if a field has been set.

### SetWorkItemIdsNil

`func (o *ExportProjectJsonRequest) SetWorkItemIdsNil(b bool)`

 SetWorkItemIdsNil sets the value for WorkItemIds to be an explicit nil

### UnsetWorkItemIds
`func (o *ExportProjectJsonRequest) UnsetWorkItemIds()`

UnsetWorkItemIds ensures that no value is present for WorkItemIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


