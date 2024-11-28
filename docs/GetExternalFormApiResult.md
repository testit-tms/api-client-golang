# GetExternalFormApiResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TestResultIds** | **[]string** | Linked test result IDs | 
**Form** | [**ExternalFormModel**](ExternalFormModel.md) | External form definition | 

## Methods

### NewGetExternalFormApiResult

`func NewGetExternalFormApiResult(testResultIds []string, form ExternalFormModel, ) *GetExternalFormApiResult`

NewGetExternalFormApiResult instantiates a new GetExternalFormApiResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetExternalFormApiResultWithDefaults

`func NewGetExternalFormApiResultWithDefaults() *GetExternalFormApiResult`

NewGetExternalFormApiResultWithDefaults instantiates a new GetExternalFormApiResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTestResultIds

`func (o *GetExternalFormApiResult) GetTestResultIds() []string`

GetTestResultIds returns the TestResultIds field if non-nil, zero value otherwise.

### GetTestResultIdsOk

`func (o *GetExternalFormApiResult) GetTestResultIdsOk() (*[]string, bool)`

GetTestResultIdsOk returns a tuple with the TestResultIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestResultIds

`func (o *GetExternalFormApiResult) SetTestResultIds(v []string)`

SetTestResultIds sets TestResultIds field to given value.


### GetForm

`func (o *GetExternalFormApiResult) GetForm() ExternalFormModel`

GetForm returns the Form field if non-nil, zero value otherwise.

### GetFormOk

`func (o *GetExternalFormApiResult) GetFormOk() (*ExternalFormModel, bool)`

GetFormOk returns a tuple with the Form field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForm

`func (o *GetExternalFormApiResult) SetForm(v ExternalFormModel)`

SetForm sets Form field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


