# GetExternalIssueSuggestionsApiModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Field** | [**ExternalIssueApiField**](ExternalIssueApiField.md) | Field of external issue metadata to get | 
**ProjectIds** | Pointer to **[]string** | List of project identifiers where external issue is available | [optional] 
**Inquiry** | Pointer to [**NullableInquiry**](Inquiry.md) | Inquiry | [optional] 

## Methods

### NewGetExternalIssueSuggestionsApiModel

`func NewGetExternalIssueSuggestionsApiModel(field ExternalIssueApiField, ) *GetExternalIssueSuggestionsApiModel`

NewGetExternalIssueSuggestionsApiModel instantiates a new GetExternalIssueSuggestionsApiModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetExternalIssueSuggestionsApiModelWithDefaults

`func NewGetExternalIssueSuggestionsApiModelWithDefaults() *GetExternalIssueSuggestionsApiModel`

NewGetExternalIssueSuggestionsApiModelWithDefaults instantiates a new GetExternalIssueSuggestionsApiModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetField

`func (o *GetExternalIssueSuggestionsApiModel) GetField() ExternalIssueApiField`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *GetExternalIssueSuggestionsApiModel) GetFieldOk() (*ExternalIssueApiField, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *GetExternalIssueSuggestionsApiModel) SetField(v ExternalIssueApiField)`

SetField sets Field field to given value.


### GetProjectIds

`func (o *GetExternalIssueSuggestionsApiModel) GetProjectIds() []string`

GetProjectIds returns the ProjectIds field if non-nil, zero value otherwise.

### GetProjectIdsOk

`func (o *GetExternalIssueSuggestionsApiModel) GetProjectIdsOk() (*[]string, bool)`

GetProjectIdsOk returns a tuple with the ProjectIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectIds

`func (o *GetExternalIssueSuggestionsApiModel) SetProjectIds(v []string)`

SetProjectIds sets ProjectIds field to given value.

### HasProjectIds

`func (o *GetExternalIssueSuggestionsApiModel) HasProjectIds() bool`

HasProjectIds returns a boolean if a field has been set.

### SetProjectIdsNil

`func (o *GetExternalIssueSuggestionsApiModel) SetProjectIdsNil(b bool)`

 SetProjectIdsNil sets the value for ProjectIds to be an explicit nil

### UnsetProjectIds
`func (o *GetExternalIssueSuggestionsApiModel) UnsetProjectIds()`

UnsetProjectIds ensures that no value is present for ProjectIds, not even an explicit nil
### GetInquiry

`func (o *GetExternalIssueSuggestionsApiModel) GetInquiry() Inquiry`

GetInquiry returns the Inquiry field if non-nil, zero value otherwise.

### GetInquiryOk

`func (o *GetExternalIssueSuggestionsApiModel) GetInquiryOk() (*Inquiry, bool)`

GetInquiryOk returns a tuple with the Inquiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInquiry

`func (o *GetExternalIssueSuggestionsApiModel) SetInquiry(v Inquiry)`

SetInquiry sets Inquiry field to given value.

### HasInquiry

`func (o *GetExternalIssueSuggestionsApiModel) HasInquiry() bool`

HasInquiry returns a boolean if a field has been set.

### SetInquiryNil

`func (o *GetExternalIssueSuggestionsApiModel) SetInquiryNil(b bool)`

 SetInquiryNil sets the value for Inquiry to be an explicit nil

### UnsetInquiry
`func (o *GetExternalIssueSuggestionsApiModel) UnsetInquiry()`

UnsetInquiry ensures that no value is present for Inquiry, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


