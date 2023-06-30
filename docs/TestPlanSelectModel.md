# TestPlanSelectModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filter** | Pointer to [**ApiV2ProjectsIdTestPlansSearchPostRequest**](ApiV2ProjectsIdTestPlansSearchPostRequest.md) |  | [optional] 
**ExtractionModel** | Pointer to [**NullableTestPlanExtractionModel**](TestPlanExtractionModel.md) |  | [optional] 

## Methods

### NewTestPlanSelectModel

`func NewTestPlanSelectModel() *TestPlanSelectModel`

NewTestPlanSelectModel instantiates a new TestPlanSelectModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestPlanSelectModelWithDefaults

`func NewTestPlanSelectModelWithDefaults() *TestPlanSelectModel`

NewTestPlanSelectModelWithDefaults instantiates a new TestPlanSelectModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilter

`func (o *TestPlanSelectModel) GetFilter() ApiV2ProjectsIdTestPlansSearchPostRequest`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *TestPlanSelectModel) GetFilterOk() (*ApiV2ProjectsIdTestPlansSearchPostRequest, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *TestPlanSelectModel) SetFilter(v ApiV2ProjectsIdTestPlansSearchPostRequest)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *TestPlanSelectModel) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetExtractionModel

`func (o *TestPlanSelectModel) GetExtractionModel() TestPlanExtractionModel`

GetExtractionModel returns the ExtractionModel field if non-nil, zero value otherwise.

### GetExtractionModelOk

`func (o *TestPlanSelectModel) GetExtractionModelOk() (*TestPlanExtractionModel, bool)`

GetExtractionModelOk returns a tuple with the ExtractionModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtractionModel

`func (o *TestPlanSelectModel) SetExtractionModel(v TestPlanExtractionModel)`

SetExtractionModel sets ExtractionModel field to given value.

### HasExtractionModel

`func (o *TestPlanSelectModel) HasExtractionModel() bool`

HasExtractionModel returns a boolean if a field has been set.

### SetExtractionModelNil

`func (o *TestPlanSelectModel) SetExtractionModelNil(b bool)`

 SetExtractionModelNil sets the value for ExtractionModel to be an explicit nil

### UnsetExtractionModel
`func (o *TestPlanSelectModel) UnsetExtractionModel()`

UnsetExtractionModel ensures that no value is present for ExtractionModel, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


