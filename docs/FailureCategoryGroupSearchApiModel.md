# FailureCategoryGroupSearchApiModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Inquiry** | [**Inquiry**](Inquiry.md) | Inquiry details | 
**Group** | Pointer to [**NullableFailureCategoryGroupApiModel**](FailureCategoryGroupApiModel.md) | Group details | [optional] 

## Methods

### NewFailureCategoryGroupSearchApiModel

`func NewFailureCategoryGroupSearchApiModel(inquiry Inquiry, ) *FailureCategoryGroupSearchApiModel`

NewFailureCategoryGroupSearchApiModel instantiates a new FailureCategoryGroupSearchApiModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFailureCategoryGroupSearchApiModelWithDefaults

`func NewFailureCategoryGroupSearchApiModelWithDefaults() *FailureCategoryGroupSearchApiModel`

NewFailureCategoryGroupSearchApiModelWithDefaults instantiates a new FailureCategoryGroupSearchApiModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInquiry

`func (o *FailureCategoryGroupSearchApiModel) GetInquiry() Inquiry`

GetInquiry returns the Inquiry field if non-nil, zero value otherwise.

### GetInquiryOk

`func (o *FailureCategoryGroupSearchApiModel) GetInquiryOk() (*Inquiry, bool)`

GetInquiryOk returns a tuple with the Inquiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInquiry

`func (o *FailureCategoryGroupSearchApiModel) SetInquiry(v Inquiry)`

SetInquiry sets Inquiry field to given value.


### GetGroup

`func (o *FailureCategoryGroupSearchApiModel) GetGroup() FailureCategoryGroupApiModel`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *FailureCategoryGroupSearchApiModel) GetGroupOk() (*FailureCategoryGroupApiModel, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *FailureCategoryGroupSearchApiModel) SetGroup(v FailureCategoryGroupApiModel)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *FailureCategoryGroupSearchApiModel) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### SetGroupNil

`func (o *FailureCategoryGroupSearchApiModel) SetGroupNil(b bool)`

 SetGroupNil sets the value for Group to be an explicit nil

### UnsetGroup
`func (o *FailureCategoryGroupSearchApiModel) UnsetGroup()`

UnsetGroup ensures that no value is present for Group, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


