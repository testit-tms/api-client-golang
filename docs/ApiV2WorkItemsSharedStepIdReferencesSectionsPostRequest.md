# ApiV2WorkItemsSharedStepIdReferencesSectionsPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | Name of section | [optional] 
**CreatedByIds** | Pointer to **[]string** | Collection of identifiers of users who created work item | [optional] 
**ModifiedByIds** | Pointer to **[]string** | Collection of identifiers of users who applied last modification to work item | [optional] 
**CreatedDate** | Pointer to [**NullableSharedStepReferenceSectionsQueryFilterModelCreatedDate**](SharedStepReferenceSectionsQueryFilterModelCreatedDate.md) |  | [optional] 
**ModifiedDate** | Pointer to [**NullableSharedStepReferenceSectionsQueryFilterModelModifiedDate**](SharedStepReferenceSectionsQueryFilterModelModifiedDate.md) |  | [optional] 

## Methods

### NewApiV2WorkItemsSharedStepIdReferencesSectionsPostRequest

`func NewApiV2WorkItemsSharedStepIdReferencesSectionsPostRequest() *ApiV2WorkItemsSharedStepIdReferencesSectionsPostRequest`

NewApiV2WorkItemsSharedStepIdReferencesSectionsPostRequest instantiates a new ApiV2WorkItemsSharedStepIdReferencesSectionsPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV2WorkItemsSharedStepIdReferencesSectionsPostRequestWithDefaults

`func NewApiV2WorkItemsSharedStepIdReferencesSectionsPostRequestWithDefaults() *ApiV2WorkItemsSharedStepIdReferencesSectionsPostRequest`

NewApiV2WorkItemsSharedStepIdReferencesSectionsPostRequestWithDefaults instantiates a new ApiV2WorkItemsSharedStepIdReferencesSectionsPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ApiV2WorkItemsSharedStepIdReferencesSectionsPostRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiV2WorkItemsSharedStepIdReferencesSectionsPostRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiV2WorkItemsSharedStepIdReferencesSectionsPostRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApiV2WorkItemsSharedStepIdReferencesSectionsPostRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ApiV2WorkItemsSharedStepIdReferencesSectionsPostRequest) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ApiV2WorkItemsSharedStepIdReferencesSectionsPostRequest) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetCreatedByIds

`func (o *ApiV2WorkItemsSharedStepIdReferencesSectionsPostRequest) GetCreatedByIds() []string`

GetCreatedByIds returns the CreatedByIds field if non-nil, zero value otherwise.

### GetCreatedByIdsOk

`func (o *ApiV2WorkItemsSharedStepIdReferencesSectionsPostRequest) GetCreatedByIdsOk() (*[]string, bool)`

GetCreatedByIdsOk returns a tuple with the CreatedByIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByIds

`func (o *ApiV2WorkItemsSharedStepIdReferencesSectionsPostRequest) SetCreatedByIds(v []string)`

SetCreatedByIds sets CreatedByIds field to given value.

### HasCreatedByIds

`func (o *ApiV2WorkItemsSharedStepIdReferencesSectionsPostRequest) HasCreatedByIds() bool`

HasCreatedByIds returns a boolean if a field has been set.

### SetCreatedByIdsNil

`func (o *ApiV2WorkItemsSharedStepIdReferencesSectionsPostRequest) SetCreatedByIdsNil(b bool)`

 SetCreatedByIdsNil sets the value for CreatedByIds to be an explicit nil

### UnsetCreatedByIds
`func (o *ApiV2WorkItemsSharedStepIdReferencesSectionsPostRequest) UnsetCreatedByIds()`

UnsetCreatedByIds ensures that no value is present for CreatedByIds, not even an explicit nil
### GetModifiedByIds

`func (o *ApiV2WorkItemsSharedStepIdReferencesSectionsPostRequest) GetModifiedByIds() []string`

GetModifiedByIds returns the ModifiedByIds field if non-nil, zero value otherwise.

### GetModifiedByIdsOk

`func (o *ApiV2WorkItemsSharedStepIdReferencesSectionsPostRequest) GetModifiedByIdsOk() (*[]string, bool)`

GetModifiedByIdsOk returns a tuple with the ModifiedByIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedByIds

`func (o *ApiV2WorkItemsSharedStepIdReferencesSectionsPostRequest) SetModifiedByIds(v []string)`

SetModifiedByIds sets ModifiedByIds field to given value.

### HasModifiedByIds

`func (o *ApiV2WorkItemsSharedStepIdReferencesSectionsPostRequest) HasModifiedByIds() bool`

HasModifiedByIds returns a boolean if a field has been set.

### SetModifiedByIdsNil

`func (o *ApiV2WorkItemsSharedStepIdReferencesSectionsPostRequest) SetModifiedByIdsNil(b bool)`

 SetModifiedByIdsNil sets the value for ModifiedByIds to be an explicit nil

### UnsetModifiedByIds
`func (o *ApiV2WorkItemsSharedStepIdReferencesSectionsPostRequest) UnsetModifiedByIds()`

UnsetModifiedByIds ensures that no value is present for ModifiedByIds, not even an explicit nil
### GetCreatedDate

`func (o *ApiV2WorkItemsSharedStepIdReferencesSectionsPostRequest) GetCreatedDate() SharedStepReferenceSectionsQueryFilterModelCreatedDate`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *ApiV2WorkItemsSharedStepIdReferencesSectionsPostRequest) GetCreatedDateOk() (*SharedStepReferenceSectionsQueryFilterModelCreatedDate, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *ApiV2WorkItemsSharedStepIdReferencesSectionsPostRequest) SetCreatedDate(v SharedStepReferenceSectionsQueryFilterModelCreatedDate)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *ApiV2WorkItemsSharedStepIdReferencesSectionsPostRequest) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### SetCreatedDateNil

`func (o *ApiV2WorkItemsSharedStepIdReferencesSectionsPostRequest) SetCreatedDateNil(b bool)`

 SetCreatedDateNil sets the value for CreatedDate to be an explicit nil

### UnsetCreatedDate
`func (o *ApiV2WorkItemsSharedStepIdReferencesSectionsPostRequest) UnsetCreatedDate()`

UnsetCreatedDate ensures that no value is present for CreatedDate, not even an explicit nil
### GetModifiedDate

`func (o *ApiV2WorkItemsSharedStepIdReferencesSectionsPostRequest) GetModifiedDate() SharedStepReferenceSectionsQueryFilterModelModifiedDate`

GetModifiedDate returns the ModifiedDate field if non-nil, zero value otherwise.

### GetModifiedDateOk

`func (o *ApiV2WorkItemsSharedStepIdReferencesSectionsPostRequest) GetModifiedDateOk() (*SharedStepReferenceSectionsQueryFilterModelModifiedDate, bool)`

GetModifiedDateOk returns a tuple with the ModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedDate

`func (o *ApiV2WorkItemsSharedStepIdReferencesSectionsPostRequest) SetModifiedDate(v SharedStepReferenceSectionsQueryFilterModelModifiedDate)`

SetModifiedDate sets ModifiedDate field to given value.

### HasModifiedDate

`func (o *ApiV2WorkItemsSharedStepIdReferencesSectionsPostRequest) HasModifiedDate() bool`

HasModifiedDate returns a boolean if a field has been set.

### SetModifiedDateNil

`func (o *ApiV2WorkItemsSharedStepIdReferencesSectionsPostRequest) SetModifiedDateNil(b bool)`

 SetModifiedDateNil sets the value for ModifiedDate to be an explicit nil

### UnsetModifiedDate
`func (o *ApiV2WorkItemsSharedStepIdReferencesSectionsPostRequest) UnsetModifiedDate()`

UnsetModifiedDate ensures that no value is present for ModifiedDate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


