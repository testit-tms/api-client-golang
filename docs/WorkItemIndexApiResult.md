# WorkItemIndexApiResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentIndex** | **int64** | Current index (position) of the element in a collection | 
**TotalCount** | **int64** | Total count of elements in a collection | 

## Methods

### NewWorkItemIndexApiResult

`func NewWorkItemIndexApiResult(currentIndex int64, totalCount int64, ) *WorkItemIndexApiResult`

NewWorkItemIndexApiResult instantiates a new WorkItemIndexApiResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkItemIndexApiResultWithDefaults

`func NewWorkItemIndexApiResultWithDefaults() *WorkItemIndexApiResult`

NewWorkItemIndexApiResultWithDefaults instantiates a new WorkItemIndexApiResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentIndex

`func (o *WorkItemIndexApiResult) GetCurrentIndex() int64`

GetCurrentIndex returns the CurrentIndex field if non-nil, zero value otherwise.

### GetCurrentIndexOk

`func (o *WorkItemIndexApiResult) GetCurrentIndexOk() (*int64, bool)`

GetCurrentIndexOk returns a tuple with the CurrentIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentIndex

`func (o *WorkItemIndexApiResult) SetCurrentIndex(v int64)`

SetCurrentIndex sets CurrentIndex field to given value.


### GetTotalCount

`func (o *WorkItemIndexApiResult) GetTotalCount() int64`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *WorkItemIndexApiResult) GetTotalCountOk() (*int64, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *WorkItemIndexApiResult) SetTotalCount(v int64)`

SetTotalCount sets TotalCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


