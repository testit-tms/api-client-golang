# AutoTestResultReasonGroupSearchApiModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Inquiry** | [**Inquiry**](Inquiry.md) | Inquiry details | 
**Group** | Pointer to [**NullableAutoTestResultReasonGroupApiModel**](AutoTestResultReasonGroupApiModel.md) | Group details | [optional] 

## Methods

### NewAutoTestResultReasonGroupSearchApiModel

`func NewAutoTestResultReasonGroupSearchApiModel(inquiry Inquiry, ) *AutoTestResultReasonGroupSearchApiModel`

NewAutoTestResultReasonGroupSearchApiModel instantiates a new AutoTestResultReasonGroupSearchApiModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoTestResultReasonGroupSearchApiModelWithDefaults

`func NewAutoTestResultReasonGroupSearchApiModelWithDefaults() *AutoTestResultReasonGroupSearchApiModel`

NewAutoTestResultReasonGroupSearchApiModelWithDefaults instantiates a new AutoTestResultReasonGroupSearchApiModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInquiry

`func (o *AutoTestResultReasonGroupSearchApiModel) GetInquiry() Inquiry`

GetInquiry returns the Inquiry field if non-nil, zero value otherwise.

### GetInquiryOk

`func (o *AutoTestResultReasonGroupSearchApiModel) GetInquiryOk() (*Inquiry, bool)`

GetInquiryOk returns a tuple with the Inquiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInquiry

`func (o *AutoTestResultReasonGroupSearchApiModel) SetInquiry(v Inquiry)`

SetInquiry sets Inquiry field to given value.


### GetGroup

`func (o *AutoTestResultReasonGroupSearchApiModel) GetGroup() AutoTestResultReasonGroupApiModel`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *AutoTestResultReasonGroupSearchApiModel) GetGroupOk() (*AutoTestResultReasonGroupApiModel, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *AutoTestResultReasonGroupSearchApiModel) SetGroup(v AutoTestResultReasonGroupApiModel)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *AutoTestResultReasonGroupSearchApiModel) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### SetGroupNil

`func (o *AutoTestResultReasonGroupSearchApiModel) SetGroupNil(b bool)`

 SetGroupNil sets the value for Group to be an explicit nil

### UnsetGroup
`func (o *AutoTestResultReasonGroupSearchApiModel) UnsetGroup()`

UnsetGroup ensures that no value is present for Group, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


