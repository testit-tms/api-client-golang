# ApiV2TestRunsIdTestResultsBulkPutRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Selector** | Pointer to [**NullableTestRunTestResultsPartialBulkSetModelSelector**](TestRunTestResultsPartialBulkSetModelSelector.md) |  | [optional] 
**ResultReasonIds** | Pointer to **[]string** | Unique IDs of result reasons to be assigned to test results | [optional] 
**Links** | Pointer to [**[]LinkPostModel**](LinkPostModel.md) | Collection of links to be assigned to test results | [optional] 
**Comment** | Pointer to **NullableString** | Comment to be added to test results | [optional] 
**AttachmentIds** | Pointer to **[]string** | Unique IDs of files to be attached to test results | [optional] 

## Methods

### NewApiV2TestRunsIdTestResultsBulkPutRequest

`func NewApiV2TestRunsIdTestResultsBulkPutRequest() *ApiV2TestRunsIdTestResultsBulkPutRequest`

NewApiV2TestRunsIdTestResultsBulkPutRequest instantiates a new ApiV2TestRunsIdTestResultsBulkPutRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV2TestRunsIdTestResultsBulkPutRequestWithDefaults

`func NewApiV2TestRunsIdTestResultsBulkPutRequestWithDefaults() *ApiV2TestRunsIdTestResultsBulkPutRequest`

NewApiV2TestRunsIdTestResultsBulkPutRequestWithDefaults instantiates a new ApiV2TestRunsIdTestResultsBulkPutRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSelector

`func (o *ApiV2TestRunsIdTestResultsBulkPutRequest) GetSelector() TestRunTestResultsPartialBulkSetModelSelector`

GetSelector returns the Selector field if non-nil, zero value otherwise.

### GetSelectorOk

`func (o *ApiV2TestRunsIdTestResultsBulkPutRequest) GetSelectorOk() (*TestRunTestResultsPartialBulkSetModelSelector, bool)`

GetSelectorOk returns a tuple with the Selector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelector

`func (o *ApiV2TestRunsIdTestResultsBulkPutRequest) SetSelector(v TestRunTestResultsPartialBulkSetModelSelector)`

SetSelector sets Selector field to given value.

### HasSelector

`func (o *ApiV2TestRunsIdTestResultsBulkPutRequest) HasSelector() bool`

HasSelector returns a boolean if a field has been set.

### SetSelectorNil

`func (o *ApiV2TestRunsIdTestResultsBulkPutRequest) SetSelectorNil(b bool)`

 SetSelectorNil sets the value for Selector to be an explicit nil

### UnsetSelector
`func (o *ApiV2TestRunsIdTestResultsBulkPutRequest) UnsetSelector()`

UnsetSelector ensures that no value is present for Selector, not even an explicit nil
### GetResultReasonIds

`func (o *ApiV2TestRunsIdTestResultsBulkPutRequest) GetResultReasonIds() []string`

GetResultReasonIds returns the ResultReasonIds field if non-nil, zero value otherwise.

### GetResultReasonIdsOk

`func (o *ApiV2TestRunsIdTestResultsBulkPutRequest) GetResultReasonIdsOk() (*[]string, bool)`

GetResultReasonIdsOk returns a tuple with the ResultReasonIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultReasonIds

`func (o *ApiV2TestRunsIdTestResultsBulkPutRequest) SetResultReasonIds(v []string)`

SetResultReasonIds sets ResultReasonIds field to given value.

### HasResultReasonIds

`func (o *ApiV2TestRunsIdTestResultsBulkPutRequest) HasResultReasonIds() bool`

HasResultReasonIds returns a boolean if a field has been set.

### SetResultReasonIdsNil

`func (o *ApiV2TestRunsIdTestResultsBulkPutRequest) SetResultReasonIdsNil(b bool)`

 SetResultReasonIdsNil sets the value for ResultReasonIds to be an explicit nil

### UnsetResultReasonIds
`func (o *ApiV2TestRunsIdTestResultsBulkPutRequest) UnsetResultReasonIds()`

UnsetResultReasonIds ensures that no value is present for ResultReasonIds, not even an explicit nil
### GetLinks

`func (o *ApiV2TestRunsIdTestResultsBulkPutRequest) GetLinks() []LinkPostModel`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ApiV2TestRunsIdTestResultsBulkPutRequest) GetLinksOk() (*[]LinkPostModel, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ApiV2TestRunsIdTestResultsBulkPutRequest) SetLinks(v []LinkPostModel)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ApiV2TestRunsIdTestResultsBulkPutRequest) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *ApiV2TestRunsIdTestResultsBulkPutRequest) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *ApiV2TestRunsIdTestResultsBulkPutRequest) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetComment

`func (o *ApiV2TestRunsIdTestResultsBulkPutRequest) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *ApiV2TestRunsIdTestResultsBulkPutRequest) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *ApiV2TestRunsIdTestResultsBulkPutRequest) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *ApiV2TestRunsIdTestResultsBulkPutRequest) HasComment() bool`

HasComment returns a boolean if a field has been set.

### SetCommentNil

`func (o *ApiV2TestRunsIdTestResultsBulkPutRequest) SetCommentNil(b bool)`

 SetCommentNil sets the value for Comment to be an explicit nil

### UnsetComment
`func (o *ApiV2TestRunsIdTestResultsBulkPutRequest) UnsetComment()`

UnsetComment ensures that no value is present for Comment, not even an explicit nil
### GetAttachmentIds

`func (o *ApiV2TestRunsIdTestResultsBulkPutRequest) GetAttachmentIds() []string`

GetAttachmentIds returns the AttachmentIds field if non-nil, zero value otherwise.

### GetAttachmentIdsOk

`func (o *ApiV2TestRunsIdTestResultsBulkPutRequest) GetAttachmentIdsOk() (*[]string, bool)`

GetAttachmentIdsOk returns a tuple with the AttachmentIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachmentIds

`func (o *ApiV2TestRunsIdTestResultsBulkPutRequest) SetAttachmentIds(v []string)`

SetAttachmentIds sets AttachmentIds field to given value.

### HasAttachmentIds

`func (o *ApiV2TestRunsIdTestResultsBulkPutRequest) HasAttachmentIds() bool`

HasAttachmentIds returns a boolean if a field has been set.

### SetAttachmentIdsNil

`func (o *ApiV2TestRunsIdTestResultsBulkPutRequest) SetAttachmentIdsNil(b bool)`

 SetAttachmentIdsNil sets the value for AttachmentIds to be an explicit nil

### UnsetAttachmentIds
`func (o *ApiV2TestRunsIdTestResultsBulkPutRequest) UnsetAttachmentIds()`

UnsetAttachmentIds ensures that no value is present for AttachmentIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


