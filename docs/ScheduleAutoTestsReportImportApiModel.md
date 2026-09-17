# ScheduleAutoTestsReportImportApiModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttachmentId** | **string** | Identifier of attachment with JUnit XML to import | 
**ConfigurationId** | Pointer to **NullableString** | Identifier of configuration. Uses default configuration in project if omitted. | [optional] 
**LayerName** | Pointer to **NullableString** | Name of test pyramid layer to set. Uses rules if omitted. | [optional] 

## Methods

### NewScheduleAutoTestsReportImportApiModel

`func NewScheduleAutoTestsReportImportApiModel(attachmentId string, ) *ScheduleAutoTestsReportImportApiModel`

NewScheduleAutoTestsReportImportApiModel instantiates a new ScheduleAutoTestsReportImportApiModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduleAutoTestsReportImportApiModelWithDefaults

`func NewScheduleAutoTestsReportImportApiModelWithDefaults() *ScheduleAutoTestsReportImportApiModel`

NewScheduleAutoTestsReportImportApiModelWithDefaults instantiates a new ScheduleAutoTestsReportImportApiModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttachmentId

`func (o *ScheduleAutoTestsReportImportApiModel) GetAttachmentId() string`

GetAttachmentId returns the AttachmentId field if non-nil, zero value otherwise.

### GetAttachmentIdOk

`func (o *ScheduleAutoTestsReportImportApiModel) GetAttachmentIdOk() (*string, bool)`

GetAttachmentIdOk returns a tuple with the AttachmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachmentId

`func (o *ScheduleAutoTestsReportImportApiModel) SetAttachmentId(v string)`

SetAttachmentId sets AttachmentId field to given value.


### GetConfigurationId

`func (o *ScheduleAutoTestsReportImportApiModel) GetConfigurationId() string`

GetConfigurationId returns the ConfigurationId field if non-nil, zero value otherwise.

### GetConfigurationIdOk

`func (o *ScheduleAutoTestsReportImportApiModel) GetConfigurationIdOk() (*string, bool)`

GetConfigurationIdOk returns a tuple with the ConfigurationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationId

`func (o *ScheduleAutoTestsReportImportApiModel) SetConfigurationId(v string)`

SetConfigurationId sets ConfigurationId field to given value.

### HasConfigurationId

`func (o *ScheduleAutoTestsReportImportApiModel) HasConfigurationId() bool`

HasConfigurationId returns a boolean if a field has been set.

### SetConfigurationIdNil

`func (o *ScheduleAutoTestsReportImportApiModel) SetConfigurationIdNil(b bool)`

 SetConfigurationIdNil sets the value for ConfigurationId to be an explicit nil

### UnsetConfigurationId
`func (o *ScheduleAutoTestsReportImportApiModel) UnsetConfigurationId()`

UnsetConfigurationId ensures that no value is present for ConfigurationId, not even an explicit nil
### GetLayerName

`func (o *ScheduleAutoTestsReportImportApiModel) GetLayerName() string`

GetLayerName returns the LayerName field if non-nil, zero value otherwise.

### GetLayerNameOk

`func (o *ScheduleAutoTestsReportImportApiModel) GetLayerNameOk() (*string, bool)`

GetLayerNameOk returns a tuple with the LayerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayerName

`func (o *ScheduleAutoTestsReportImportApiModel) SetLayerName(v string)`

SetLayerName sets LayerName field to given value.

### HasLayerName

`func (o *ScheduleAutoTestsReportImportApiModel) HasLayerName() bool`

HasLayerName returns a boolean if a field has been set.

### SetLayerNameNil

`func (o *ScheduleAutoTestsReportImportApiModel) SetLayerNameNil(b bool)`

 SetLayerNameNil sets the value for LayerName to be an explicit nil

### UnsetLayerName
`func (o *ScheduleAutoTestsReportImportApiModel) UnsetLayerName()`

UnsetLayerName ensures that no value is present for LayerName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


