# ExternalIssueApiMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Identifier of external issue in external service | 
**Title** | **string** | Title of external issue in external service | 
**Code** | Pointer to **NullableString** | Code of external issue in external service | [optional] 
**Description** | Pointer to **NullableString** | Description of external issue in external service | [optional] 
**Status** | Pointer to **NullableString** | Status of external issue in external service | [optional] 
**Assignee** | Pointer to **NullableString** | Assignee of external issue in external service | [optional] 
**Type** | Pointer to [**NullableExternalIssueApiType**](ExternalIssueApiType.md) | Type of external issue in external service | [optional] 
**Priority** | Pointer to [**NullableExternalIssueApiPriority**](ExternalIssueApiPriority.md) | Priority of external issue in external service | [optional] 

## Methods

### NewExternalIssueApiMetadata

`func NewExternalIssueApiMetadata(id string, title string, ) *ExternalIssueApiMetadata`

NewExternalIssueApiMetadata instantiates a new ExternalIssueApiMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalIssueApiMetadataWithDefaults

`func NewExternalIssueApiMetadataWithDefaults() *ExternalIssueApiMetadata`

NewExternalIssueApiMetadataWithDefaults instantiates a new ExternalIssueApiMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ExternalIssueApiMetadata) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExternalIssueApiMetadata) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExternalIssueApiMetadata) SetId(v string)`

SetId sets Id field to given value.


### GetTitle

`func (o *ExternalIssueApiMetadata) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ExternalIssueApiMetadata) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ExternalIssueApiMetadata) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetCode

`func (o *ExternalIssueApiMetadata) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ExternalIssueApiMetadata) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ExternalIssueApiMetadata) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ExternalIssueApiMetadata) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *ExternalIssueApiMetadata) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *ExternalIssueApiMetadata) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetDescription

`func (o *ExternalIssueApiMetadata) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ExternalIssueApiMetadata) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ExternalIssueApiMetadata) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ExternalIssueApiMetadata) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ExternalIssueApiMetadata) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ExternalIssueApiMetadata) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetStatus

`func (o *ExternalIssueApiMetadata) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ExternalIssueApiMetadata) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ExternalIssueApiMetadata) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ExternalIssueApiMetadata) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *ExternalIssueApiMetadata) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *ExternalIssueApiMetadata) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetAssignee

`func (o *ExternalIssueApiMetadata) GetAssignee() string`

GetAssignee returns the Assignee field if non-nil, zero value otherwise.

### GetAssigneeOk

`func (o *ExternalIssueApiMetadata) GetAssigneeOk() (*string, bool)`

GetAssigneeOk returns a tuple with the Assignee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignee

`func (o *ExternalIssueApiMetadata) SetAssignee(v string)`

SetAssignee sets Assignee field to given value.

### HasAssignee

`func (o *ExternalIssueApiMetadata) HasAssignee() bool`

HasAssignee returns a boolean if a field has been set.

### SetAssigneeNil

`func (o *ExternalIssueApiMetadata) SetAssigneeNil(b bool)`

 SetAssigneeNil sets the value for Assignee to be an explicit nil

### UnsetAssignee
`func (o *ExternalIssueApiMetadata) UnsetAssignee()`

UnsetAssignee ensures that no value is present for Assignee, not even an explicit nil
### GetType

`func (o *ExternalIssueApiMetadata) GetType() ExternalIssueApiType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ExternalIssueApiMetadata) GetTypeOk() (*ExternalIssueApiType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ExternalIssueApiMetadata) SetType(v ExternalIssueApiType)`

SetType sets Type field to given value.

### HasType

`func (o *ExternalIssueApiMetadata) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *ExternalIssueApiMetadata) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *ExternalIssueApiMetadata) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetPriority

`func (o *ExternalIssueApiMetadata) GetPriority() ExternalIssueApiPriority`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *ExternalIssueApiMetadata) GetPriorityOk() (*ExternalIssueApiPriority, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *ExternalIssueApiMetadata) SetPriority(v ExternalIssueApiPriority)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *ExternalIssueApiMetadata) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### SetPriorityNil

`func (o *ExternalIssueApiMetadata) SetPriorityNil(b bool)`

 SetPriorityNil sets the value for Priority to be an explicit nil

### UnsetPriority
`func (o *ExternalIssueApiMetadata) UnsetPriority()`

UnsetPriority ensures that no value is present for Priority, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


