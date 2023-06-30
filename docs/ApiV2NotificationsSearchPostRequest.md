# ApiV2NotificationsSearchPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Types** | Pointer to [**[]NotificationTypeModel**](NotificationTypeModel.md) |  | [optional] 
**IsRead** | Pointer to **NullableBool** |  | [optional] 
**CreatedDate** | Pointer to [**NullableDateTimeRangeSelectorModel**](DateTimeRangeSelectorModel.md) |  | [optional] 

## Methods

### NewApiV2NotificationsSearchPostRequest

`func NewApiV2NotificationsSearchPostRequest() *ApiV2NotificationsSearchPostRequest`

NewApiV2NotificationsSearchPostRequest instantiates a new ApiV2NotificationsSearchPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV2NotificationsSearchPostRequestWithDefaults

`func NewApiV2NotificationsSearchPostRequestWithDefaults() *ApiV2NotificationsSearchPostRequest`

NewApiV2NotificationsSearchPostRequestWithDefaults instantiates a new ApiV2NotificationsSearchPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTypes

`func (o *ApiV2NotificationsSearchPostRequest) GetTypes() []NotificationTypeModel`

GetTypes returns the Types field if non-nil, zero value otherwise.

### GetTypesOk

`func (o *ApiV2NotificationsSearchPostRequest) GetTypesOk() (*[]NotificationTypeModel, bool)`

GetTypesOk returns a tuple with the Types field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypes

`func (o *ApiV2NotificationsSearchPostRequest) SetTypes(v []NotificationTypeModel)`

SetTypes sets Types field to given value.

### HasTypes

`func (o *ApiV2NotificationsSearchPostRequest) HasTypes() bool`

HasTypes returns a boolean if a field has been set.

### SetTypesNil

`func (o *ApiV2NotificationsSearchPostRequest) SetTypesNil(b bool)`

 SetTypesNil sets the value for Types to be an explicit nil

### UnsetTypes
`func (o *ApiV2NotificationsSearchPostRequest) UnsetTypes()`

UnsetTypes ensures that no value is present for Types, not even an explicit nil
### GetIsRead

`func (o *ApiV2NotificationsSearchPostRequest) GetIsRead() bool`

GetIsRead returns the IsRead field if non-nil, zero value otherwise.

### GetIsReadOk

`func (o *ApiV2NotificationsSearchPostRequest) GetIsReadOk() (*bool, bool)`

GetIsReadOk returns a tuple with the IsRead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRead

`func (o *ApiV2NotificationsSearchPostRequest) SetIsRead(v bool)`

SetIsRead sets IsRead field to given value.

### HasIsRead

`func (o *ApiV2NotificationsSearchPostRequest) HasIsRead() bool`

HasIsRead returns a boolean if a field has been set.

### SetIsReadNil

`func (o *ApiV2NotificationsSearchPostRequest) SetIsReadNil(b bool)`

 SetIsReadNil sets the value for IsRead to be an explicit nil

### UnsetIsRead
`func (o *ApiV2NotificationsSearchPostRequest) UnsetIsRead()`

UnsetIsRead ensures that no value is present for IsRead, not even an explicit nil
### GetCreatedDate

`func (o *ApiV2NotificationsSearchPostRequest) GetCreatedDate() DateTimeRangeSelectorModel`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *ApiV2NotificationsSearchPostRequest) GetCreatedDateOk() (*DateTimeRangeSelectorModel, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *ApiV2NotificationsSearchPostRequest) SetCreatedDate(v DateTimeRangeSelectorModel)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *ApiV2NotificationsSearchPostRequest) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### SetCreatedDateNil

`func (o *ApiV2NotificationsSearchPostRequest) SetCreatedDateNil(b bool)`

 SetCreatedDateNil sets the value for CreatedDate to be an explicit nil

### UnsetCreatedDate
`func (o *ApiV2NotificationsSearchPostRequest) UnsetCreatedDate()`

UnsetCreatedDate ensures that no value is present for CreatedDate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


