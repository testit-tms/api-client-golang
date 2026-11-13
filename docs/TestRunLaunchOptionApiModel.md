# TestRunLaunchOptionApiModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomParameters** | Pointer to **map[string]string** | Customers test run parameters. | [optional] 
**WebhookIds** | **[]string** | Enabled webhooks. | 

## Methods

### NewTestRunLaunchOptionApiModel

`func NewTestRunLaunchOptionApiModel(webhookIds []string, ) *TestRunLaunchOptionApiModel`

NewTestRunLaunchOptionApiModel instantiates a new TestRunLaunchOptionApiModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestRunLaunchOptionApiModelWithDefaults

`func NewTestRunLaunchOptionApiModelWithDefaults() *TestRunLaunchOptionApiModel`

NewTestRunLaunchOptionApiModelWithDefaults instantiates a new TestRunLaunchOptionApiModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomParameters

`func (o *TestRunLaunchOptionApiModel) GetCustomParameters() map[string]string`

GetCustomParameters returns the CustomParameters field if non-nil, zero value otherwise.

### GetCustomParametersOk

`func (o *TestRunLaunchOptionApiModel) GetCustomParametersOk() (*map[string]string, bool)`

GetCustomParametersOk returns a tuple with the CustomParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomParameters

`func (o *TestRunLaunchOptionApiModel) SetCustomParameters(v map[string]string)`

SetCustomParameters sets CustomParameters field to given value.

### HasCustomParameters

`func (o *TestRunLaunchOptionApiModel) HasCustomParameters() bool`

HasCustomParameters returns a boolean if a field has been set.

### SetCustomParametersNil

`func (o *TestRunLaunchOptionApiModel) SetCustomParametersNil(b bool)`

 SetCustomParametersNil sets the value for CustomParameters to be an explicit nil

### UnsetCustomParameters
`func (o *TestRunLaunchOptionApiModel) UnsetCustomParameters()`

UnsetCustomParameters ensures that no value is present for CustomParameters, not even an explicit nil
### GetWebhookIds

`func (o *TestRunLaunchOptionApiModel) GetWebhookIds() []string`

GetWebhookIds returns the WebhookIds field if non-nil, zero value otherwise.

### GetWebhookIdsOk

`func (o *TestRunLaunchOptionApiModel) GetWebhookIdsOk() (*[]string, bool)`

GetWebhookIdsOk returns a tuple with the WebhookIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookIds

`func (o *TestRunLaunchOptionApiModel) SetWebhookIds(v []string)`

SetWebhookIds sets WebhookIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


