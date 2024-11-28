# CreateDefectApiModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TestResultIds** | **[]string** | Linked test result IDs | 
**Form** | [**ExternalFormCreateModel**](ExternalFormCreateModel.md) | External form definition | 

## Methods

### NewCreateDefectApiModel

`func NewCreateDefectApiModel(testResultIds []string, form ExternalFormCreateModel, ) *CreateDefectApiModel`

NewCreateDefectApiModel instantiates a new CreateDefectApiModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDefectApiModelWithDefaults

`func NewCreateDefectApiModelWithDefaults() *CreateDefectApiModel`

NewCreateDefectApiModelWithDefaults instantiates a new CreateDefectApiModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTestResultIds

`func (o *CreateDefectApiModel) GetTestResultIds() []string`

GetTestResultIds returns the TestResultIds field if non-nil, zero value otherwise.

### GetTestResultIdsOk

`func (o *CreateDefectApiModel) GetTestResultIdsOk() (*[]string, bool)`

GetTestResultIdsOk returns a tuple with the TestResultIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestResultIds

`func (o *CreateDefectApiModel) SetTestResultIds(v []string)`

SetTestResultIds sets TestResultIds field to given value.


### GetForm

`func (o *CreateDefectApiModel) GetForm() ExternalFormCreateModel`

GetForm returns the Form field if non-nil, zero value otherwise.

### GetFormOk

`func (o *CreateDefectApiModel) GetFormOk() (*ExternalFormCreateModel, bool)`

GetFormOk returns a tuple with the Form field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForm

`func (o *CreateDefectApiModel) SetForm(v ExternalFormCreateModel)`

SetForm sets Form field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


