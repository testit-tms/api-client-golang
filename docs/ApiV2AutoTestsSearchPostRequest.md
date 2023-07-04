# ApiV2AutoTestsSearchPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filter** | Pointer to [**AutotestsSelectModelFilter**](AutotestsSelectModelFilter.md) |  | [optional] 
**Includes** | Pointer to [**AutotestsSelectModelIncludes**](AutotestsSelectModelIncludes.md) |  | [optional] 

## Methods

### NewApiV2AutoTestsSearchPostRequest

`func NewApiV2AutoTestsSearchPostRequest() *ApiV2AutoTestsSearchPostRequest`

NewApiV2AutoTestsSearchPostRequest instantiates a new ApiV2AutoTestsSearchPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV2AutoTestsSearchPostRequestWithDefaults

`func NewApiV2AutoTestsSearchPostRequestWithDefaults() *ApiV2AutoTestsSearchPostRequest`

NewApiV2AutoTestsSearchPostRequestWithDefaults instantiates a new ApiV2AutoTestsSearchPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilter

`func (o *ApiV2AutoTestsSearchPostRequest) GetFilter() AutotestsSelectModelFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *ApiV2AutoTestsSearchPostRequest) GetFilterOk() (*AutotestsSelectModelFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *ApiV2AutoTestsSearchPostRequest) SetFilter(v AutotestsSelectModelFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *ApiV2AutoTestsSearchPostRequest) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetIncludes

`func (o *ApiV2AutoTestsSearchPostRequest) GetIncludes() AutotestsSelectModelIncludes`

GetIncludes returns the Includes field if non-nil, zero value otherwise.

### GetIncludesOk

`func (o *ApiV2AutoTestsSearchPostRequest) GetIncludesOk() (*AutotestsSelectModelIncludes, bool)`

GetIncludesOk returns a tuple with the Includes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludes

`func (o *ApiV2AutoTestsSearchPostRequest) SetIncludes(v AutotestsSelectModelIncludes)`

SetIncludes sets Includes field to given value.

### HasIncludes

`func (o *ApiV2AutoTestsSearchPostRequest) HasIncludes() bool`

HasIncludes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


