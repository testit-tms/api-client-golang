# ApiV2WorkItemsMovePostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**NewSectionId** | **string** |  | 
**OldSectionId** | Pointer to **NullableString** |  | [optional] 
**NextWorkItemId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewApiV2WorkItemsMovePostRequest

`func NewApiV2WorkItemsMovePostRequest(id string, newSectionId string, ) *ApiV2WorkItemsMovePostRequest`

NewApiV2WorkItemsMovePostRequest instantiates a new ApiV2WorkItemsMovePostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV2WorkItemsMovePostRequestWithDefaults

`func NewApiV2WorkItemsMovePostRequestWithDefaults() *ApiV2WorkItemsMovePostRequest`

NewApiV2WorkItemsMovePostRequestWithDefaults instantiates a new ApiV2WorkItemsMovePostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApiV2WorkItemsMovePostRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiV2WorkItemsMovePostRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiV2WorkItemsMovePostRequest) SetId(v string)`

SetId sets Id field to given value.


### GetNewSectionId

`func (o *ApiV2WorkItemsMovePostRequest) GetNewSectionId() string`

GetNewSectionId returns the NewSectionId field if non-nil, zero value otherwise.

### GetNewSectionIdOk

`func (o *ApiV2WorkItemsMovePostRequest) GetNewSectionIdOk() (*string, bool)`

GetNewSectionIdOk returns a tuple with the NewSectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewSectionId

`func (o *ApiV2WorkItemsMovePostRequest) SetNewSectionId(v string)`

SetNewSectionId sets NewSectionId field to given value.


### GetOldSectionId

`func (o *ApiV2WorkItemsMovePostRequest) GetOldSectionId() string`

GetOldSectionId returns the OldSectionId field if non-nil, zero value otherwise.

### GetOldSectionIdOk

`func (o *ApiV2WorkItemsMovePostRequest) GetOldSectionIdOk() (*string, bool)`

GetOldSectionIdOk returns a tuple with the OldSectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldSectionId

`func (o *ApiV2WorkItemsMovePostRequest) SetOldSectionId(v string)`

SetOldSectionId sets OldSectionId field to given value.

### HasOldSectionId

`func (o *ApiV2WorkItemsMovePostRequest) HasOldSectionId() bool`

HasOldSectionId returns a boolean if a field has been set.

### SetOldSectionIdNil

`func (o *ApiV2WorkItemsMovePostRequest) SetOldSectionIdNil(b bool)`

 SetOldSectionIdNil sets the value for OldSectionId to be an explicit nil

### UnsetOldSectionId
`func (o *ApiV2WorkItemsMovePostRequest) UnsetOldSectionId()`

UnsetOldSectionId ensures that no value is present for OldSectionId, not even an explicit nil
### GetNextWorkItemId

`func (o *ApiV2WorkItemsMovePostRequest) GetNextWorkItemId() string`

GetNextWorkItemId returns the NextWorkItemId field if non-nil, zero value otherwise.

### GetNextWorkItemIdOk

`func (o *ApiV2WorkItemsMovePostRequest) GetNextWorkItemIdOk() (*string, bool)`

GetNextWorkItemIdOk returns a tuple with the NextWorkItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextWorkItemId

`func (o *ApiV2WorkItemsMovePostRequest) SetNextWorkItemId(v string)`

SetNextWorkItemId sets NextWorkItemId field to given value.

### HasNextWorkItemId

`func (o *ApiV2WorkItemsMovePostRequest) HasNextWorkItemId() bool`

HasNextWorkItemId returns a boolean if a field has been set.

### SetNextWorkItemIdNil

`func (o *ApiV2WorkItemsMovePostRequest) SetNextWorkItemIdNil(b bool)`

 SetNextWorkItemIdNil sets the value for NextWorkItemId to be an explicit nil

### UnsetNextWorkItemId
`func (o *ApiV2WorkItemsMovePostRequest) UnsetNextWorkItemId()`

UnsetNextWorkItemId ensures that no value is present for NextWorkItemId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


