# ExternalIssueApiFieldSuggestion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | **string** | Value of the external issue field | 
**ExternalService** | [**ExternalIssueExternalServiceApiResult**](ExternalIssueExternalServiceApiResult.md) | Associated external service with this value | 

## Methods

### NewExternalIssueApiFieldSuggestion

`func NewExternalIssueApiFieldSuggestion(value string, externalService ExternalIssueExternalServiceApiResult, ) *ExternalIssueApiFieldSuggestion`

NewExternalIssueApiFieldSuggestion instantiates a new ExternalIssueApiFieldSuggestion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalIssueApiFieldSuggestionWithDefaults

`func NewExternalIssueApiFieldSuggestionWithDefaults() *ExternalIssueApiFieldSuggestion`

NewExternalIssueApiFieldSuggestionWithDefaults instantiates a new ExternalIssueApiFieldSuggestion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *ExternalIssueApiFieldSuggestion) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ExternalIssueApiFieldSuggestion) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ExternalIssueApiFieldSuggestion) SetValue(v string)`

SetValue sets Value field to given value.


### GetExternalService

`func (o *ExternalIssueApiFieldSuggestion) GetExternalService() ExternalIssueExternalServiceApiResult`

GetExternalService returns the ExternalService field if non-nil, zero value otherwise.

### GetExternalServiceOk

`func (o *ExternalIssueApiFieldSuggestion) GetExternalServiceOk() (*ExternalIssueExternalServiceApiResult, bool)`

GetExternalServiceOk returns a tuple with the ExternalService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalService

`func (o *ExternalIssueApiFieldSuggestion) SetExternalService(v ExternalIssueExternalServiceApiResult)`

SetExternalService sets ExternalService field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


