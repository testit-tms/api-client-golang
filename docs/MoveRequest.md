# MoveRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique ID of the section | 
**OldParentId** | **string** | Unique ID of the section&#39;s current parent section | 
**ParentId** | **string** | Unique ID of the section&#39;s target parent section | 
**NextSectionId** | Pointer to **NullableString** | Unique ID of the section&#39;s following section | [optional] 

## Methods

### NewMoveRequest

`func NewMoveRequest(id string, oldParentId string, parentId string, ) *MoveRequest`

NewMoveRequest instantiates a new MoveRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMoveRequestWithDefaults

`func NewMoveRequestWithDefaults() *MoveRequest`

NewMoveRequestWithDefaults instantiates a new MoveRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MoveRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MoveRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MoveRequest) SetId(v string)`

SetId sets Id field to given value.


### GetOldParentId

`func (o *MoveRequest) GetOldParentId() string`

GetOldParentId returns the OldParentId field if non-nil, zero value otherwise.

### GetOldParentIdOk

`func (o *MoveRequest) GetOldParentIdOk() (*string, bool)`

GetOldParentIdOk returns a tuple with the OldParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldParentId

`func (o *MoveRequest) SetOldParentId(v string)`

SetOldParentId sets OldParentId field to given value.


### GetParentId

`func (o *MoveRequest) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *MoveRequest) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *MoveRequest) SetParentId(v string)`

SetParentId sets ParentId field to given value.


### GetNextSectionId

`func (o *MoveRequest) GetNextSectionId() string`

GetNextSectionId returns the NextSectionId field if non-nil, zero value otherwise.

### GetNextSectionIdOk

`func (o *MoveRequest) GetNextSectionIdOk() (*string, bool)`

GetNextSectionIdOk returns a tuple with the NextSectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextSectionId

`func (o *MoveRequest) SetNextSectionId(v string)`

SetNextSectionId sets NextSectionId field to given value.

### HasNextSectionId

`func (o *MoveRequest) HasNextSectionId() bool`

HasNextSectionId returns a boolean if a field has been set.

### SetNextSectionIdNil

`func (o *MoveRequest) SetNextSectionIdNil(b bool)`

 SetNextSectionIdNil sets the value for NextSectionId to be an explicit nil

### UnsetNextSectionId
`func (o *MoveRequest) UnsetNextSectionId()`

UnsetNextSectionId ensures that no value is present for NextSectionId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


