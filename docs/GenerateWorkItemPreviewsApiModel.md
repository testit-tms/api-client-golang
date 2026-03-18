# GenerateWorkItemPreviewsApiModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalServiceId** | **string** | The ID of the external AI service to be used for generation. | 
**TaskKey** | Pointer to **NullableString** | The key of the issue in an issue tracker (e.g., JIRA-123). | [optional] 
**IssueKey** | Pointer to **NullableString** | The key of the issue in an issue tracker (e.g., JIRA-123). | [optional] 
**UserContext** | Pointer to **NullableString** | Additional user context or description of the issue if no issue key is provided. | [optional] 
**Temperature** | **float32** | Controls randomness of the AI model output. | 
**PreviewLimit** | **int32** | Number of work item previews to generate. | 

## Methods

### NewGenerateWorkItemPreviewsApiModel

`func NewGenerateWorkItemPreviewsApiModel(externalServiceId string, temperature float32, previewLimit int32, ) *GenerateWorkItemPreviewsApiModel`

NewGenerateWorkItemPreviewsApiModel instantiates a new GenerateWorkItemPreviewsApiModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenerateWorkItemPreviewsApiModelWithDefaults

`func NewGenerateWorkItemPreviewsApiModelWithDefaults() *GenerateWorkItemPreviewsApiModel`

NewGenerateWorkItemPreviewsApiModelWithDefaults instantiates a new GenerateWorkItemPreviewsApiModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalServiceId

`func (o *GenerateWorkItemPreviewsApiModel) GetExternalServiceId() string`

GetExternalServiceId returns the ExternalServiceId field if non-nil, zero value otherwise.

### GetExternalServiceIdOk

`func (o *GenerateWorkItemPreviewsApiModel) GetExternalServiceIdOk() (*string, bool)`

GetExternalServiceIdOk returns a tuple with the ExternalServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalServiceId

`func (o *GenerateWorkItemPreviewsApiModel) SetExternalServiceId(v string)`

SetExternalServiceId sets ExternalServiceId field to given value.


### GetTaskKey

`func (o *GenerateWorkItemPreviewsApiModel) GetTaskKey() string`

GetTaskKey returns the TaskKey field if non-nil, zero value otherwise.

### GetTaskKeyOk

`func (o *GenerateWorkItemPreviewsApiModel) GetTaskKeyOk() (*string, bool)`

GetTaskKeyOk returns a tuple with the TaskKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskKey

`func (o *GenerateWorkItemPreviewsApiModel) SetTaskKey(v string)`

SetTaskKey sets TaskKey field to given value.

### HasTaskKey

`func (o *GenerateWorkItemPreviewsApiModel) HasTaskKey() bool`

HasTaskKey returns a boolean if a field has been set.

### SetTaskKeyNil

`func (o *GenerateWorkItemPreviewsApiModel) SetTaskKeyNil(b bool)`

 SetTaskKeyNil sets the value for TaskKey to be an explicit nil

### UnsetTaskKey
`func (o *GenerateWorkItemPreviewsApiModel) UnsetTaskKey()`

UnsetTaskKey ensures that no value is present for TaskKey, not even an explicit nil
### GetIssueKey

`func (o *GenerateWorkItemPreviewsApiModel) GetIssueKey() string`

GetIssueKey returns the IssueKey field if non-nil, zero value otherwise.

### GetIssueKeyOk

`func (o *GenerateWorkItemPreviewsApiModel) GetIssueKeyOk() (*string, bool)`

GetIssueKeyOk returns a tuple with the IssueKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueKey

`func (o *GenerateWorkItemPreviewsApiModel) SetIssueKey(v string)`

SetIssueKey sets IssueKey field to given value.

### HasIssueKey

`func (o *GenerateWorkItemPreviewsApiModel) HasIssueKey() bool`

HasIssueKey returns a boolean if a field has been set.

### SetIssueKeyNil

`func (o *GenerateWorkItemPreviewsApiModel) SetIssueKeyNil(b bool)`

 SetIssueKeyNil sets the value for IssueKey to be an explicit nil

### UnsetIssueKey
`func (o *GenerateWorkItemPreviewsApiModel) UnsetIssueKey()`

UnsetIssueKey ensures that no value is present for IssueKey, not even an explicit nil
### GetUserContext

`func (o *GenerateWorkItemPreviewsApiModel) GetUserContext() string`

GetUserContext returns the UserContext field if non-nil, zero value otherwise.

### GetUserContextOk

`func (o *GenerateWorkItemPreviewsApiModel) GetUserContextOk() (*string, bool)`

GetUserContextOk returns a tuple with the UserContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserContext

`func (o *GenerateWorkItemPreviewsApiModel) SetUserContext(v string)`

SetUserContext sets UserContext field to given value.

### HasUserContext

`func (o *GenerateWorkItemPreviewsApiModel) HasUserContext() bool`

HasUserContext returns a boolean if a field has been set.

### SetUserContextNil

`func (o *GenerateWorkItemPreviewsApiModel) SetUserContextNil(b bool)`

 SetUserContextNil sets the value for UserContext to be an explicit nil

### UnsetUserContext
`func (o *GenerateWorkItemPreviewsApiModel) UnsetUserContext()`

UnsetUserContext ensures that no value is present for UserContext, not even an explicit nil
### GetTemperature

`func (o *GenerateWorkItemPreviewsApiModel) GetTemperature() float32`

GetTemperature returns the Temperature field if non-nil, zero value otherwise.

### GetTemperatureOk

`func (o *GenerateWorkItemPreviewsApiModel) GetTemperatureOk() (*float32, bool)`

GetTemperatureOk returns a tuple with the Temperature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemperature

`func (o *GenerateWorkItemPreviewsApiModel) SetTemperature(v float32)`

SetTemperature sets Temperature field to given value.


### GetPreviewLimit

`func (o *GenerateWorkItemPreviewsApiModel) GetPreviewLimit() int32`

GetPreviewLimit returns the PreviewLimit field if non-nil, zero value otherwise.

### GetPreviewLimitOk

`func (o *GenerateWorkItemPreviewsApiModel) GetPreviewLimitOk() (*int32, bool)`

GetPreviewLimitOk returns a tuple with the PreviewLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviewLimit

`func (o *GenerateWorkItemPreviewsApiModel) SetPreviewLimit(v int32)`

SetPreviewLimit sets PreviewLimit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


