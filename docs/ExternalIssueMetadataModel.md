# ExternalIssueMetadataModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Title** | **string** |  | 
**Code** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to **NullableString** |  | [optional] 
**Assignee** | Pointer to **NullableString** |  | [optional] 
**Type** | Pointer to [**NullableExternalIssueTypeModel**](ExternalIssueTypeModel.md) |  | [optional] 
**Priority** | Pointer to [**NullableExternalIssuePriorityModel**](ExternalIssuePriorityModel.md) |  | [optional] 

## Methods

### NewExternalIssueMetadataModel

`func NewExternalIssueMetadataModel(id string, title string, ) *ExternalIssueMetadataModel`

NewExternalIssueMetadataModel instantiates a new ExternalIssueMetadataModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalIssueMetadataModelWithDefaults

`func NewExternalIssueMetadataModelWithDefaults() *ExternalIssueMetadataModel`

NewExternalIssueMetadataModelWithDefaults instantiates a new ExternalIssueMetadataModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ExternalIssueMetadataModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExternalIssueMetadataModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExternalIssueMetadataModel) SetId(v string)`

SetId sets Id field to given value.


### GetTitle

`func (o *ExternalIssueMetadataModel) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ExternalIssueMetadataModel) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ExternalIssueMetadataModel) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetCode

`func (o *ExternalIssueMetadataModel) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ExternalIssueMetadataModel) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ExternalIssueMetadataModel) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ExternalIssueMetadataModel) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *ExternalIssueMetadataModel) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *ExternalIssueMetadataModel) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetDescription

`func (o *ExternalIssueMetadataModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ExternalIssueMetadataModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ExternalIssueMetadataModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ExternalIssueMetadataModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ExternalIssueMetadataModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ExternalIssueMetadataModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetStatus

`func (o *ExternalIssueMetadataModel) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ExternalIssueMetadataModel) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ExternalIssueMetadataModel) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ExternalIssueMetadataModel) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *ExternalIssueMetadataModel) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *ExternalIssueMetadataModel) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetAssignee

`func (o *ExternalIssueMetadataModel) GetAssignee() string`

GetAssignee returns the Assignee field if non-nil, zero value otherwise.

### GetAssigneeOk

`func (o *ExternalIssueMetadataModel) GetAssigneeOk() (*string, bool)`

GetAssigneeOk returns a tuple with the Assignee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignee

`func (o *ExternalIssueMetadataModel) SetAssignee(v string)`

SetAssignee sets Assignee field to given value.

### HasAssignee

`func (o *ExternalIssueMetadataModel) HasAssignee() bool`

HasAssignee returns a boolean if a field has been set.

### SetAssigneeNil

`func (o *ExternalIssueMetadataModel) SetAssigneeNil(b bool)`

 SetAssigneeNil sets the value for Assignee to be an explicit nil

### UnsetAssignee
`func (o *ExternalIssueMetadataModel) UnsetAssignee()`

UnsetAssignee ensures that no value is present for Assignee, not even an explicit nil
### GetType

`func (o *ExternalIssueMetadataModel) GetType() ExternalIssueTypeModel`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ExternalIssueMetadataModel) GetTypeOk() (*ExternalIssueTypeModel, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ExternalIssueMetadataModel) SetType(v ExternalIssueTypeModel)`

SetType sets Type field to given value.

### HasType

`func (o *ExternalIssueMetadataModel) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *ExternalIssueMetadataModel) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *ExternalIssueMetadataModel) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetPriority

`func (o *ExternalIssueMetadataModel) GetPriority() ExternalIssuePriorityModel`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *ExternalIssueMetadataModel) GetPriorityOk() (*ExternalIssuePriorityModel, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *ExternalIssueMetadataModel) SetPriority(v ExternalIssuePriorityModel)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *ExternalIssueMetadataModel) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### SetPriorityNil

`func (o *ExternalIssueMetadataModel) SetPriorityNil(b bool)`

 SetPriorityNil sets the value for Priority to be an explicit nil

### UnsetPriority
`func (o *ExternalIssueMetadataModel) UnsetPriority()`

UnsetPriority ensures that no value is present for Priority, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


