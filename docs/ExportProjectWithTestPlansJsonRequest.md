# ExportProjectWithTestPlansJsonRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TestPlansIds** | Pointer to **[]string** | Specifies the IDs of test plans to be exported with the project.&lt;br /&gt;  In this parameter, \&quot;&lt;b&gt;string&lt;/b&gt;\&quot; values are IDs of the test plans.&lt;br /&gt;  To get the test plan IDs, use the &#x60;GET /api/v2/projects/{projectId}/testPlans&#x60; method. | [optional] 

## Methods

### NewExportProjectWithTestPlansJsonRequest

`func NewExportProjectWithTestPlansJsonRequest() *ExportProjectWithTestPlansJsonRequest`

NewExportProjectWithTestPlansJsonRequest instantiates a new ExportProjectWithTestPlansJsonRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExportProjectWithTestPlansJsonRequestWithDefaults

`func NewExportProjectWithTestPlansJsonRequestWithDefaults() *ExportProjectWithTestPlansJsonRequest`

NewExportProjectWithTestPlansJsonRequestWithDefaults instantiates a new ExportProjectWithTestPlansJsonRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTestPlansIds

`func (o *ExportProjectWithTestPlansJsonRequest) GetTestPlansIds() []string`

GetTestPlansIds returns the TestPlansIds field if non-nil, zero value otherwise.

### GetTestPlansIdsOk

`func (o *ExportProjectWithTestPlansJsonRequest) GetTestPlansIdsOk() (*[]string, bool)`

GetTestPlansIdsOk returns a tuple with the TestPlansIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestPlansIds

`func (o *ExportProjectWithTestPlansJsonRequest) SetTestPlansIds(v []string)`

SetTestPlansIds sets TestPlansIds field to given value.

### HasTestPlansIds

`func (o *ExportProjectWithTestPlansJsonRequest) HasTestPlansIds() bool`

HasTestPlansIds returns a boolean if a field has been set.

### SetTestPlansIdsNil

`func (o *ExportProjectWithTestPlansJsonRequest) SetTestPlansIdsNil(b bool)`

 SetTestPlansIdsNil sets the value for TestPlansIds to be an explicit nil

### UnsetTestPlansIds
`func (o *ExportProjectWithTestPlansJsonRequest) UnsetTestPlansIds()`

UnsetTestPlansIds ensures that no value is present for TestPlansIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


