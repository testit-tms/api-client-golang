# AuditApiResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | **string** |  | 
**UserName** | Pointer to **NullableString** |  | [optional] 
**DateTime** | **time.Time** |  | 

## Methods

### NewAuditApiResult

`func NewAuditApiResult(userId string, dateTime time.Time, ) *AuditApiResult`

NewAuditApiResult instantiates a new AuditApiResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuditApiResultWithDefaults

`func NewAuditApiResultWithDefaults() *AuditApiResult`

NewAuditApiResultWithDefaults instantiates a new AuditApiResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *AuditApiResult) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *AuditApiResult) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *AuditApiResult) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetUserName

`func (o *AuditApiResult) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *AuditApiResult) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *AuditApiResult) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *AuditApiResult) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### SetUserNameNil

`func (o *AuditApiResult) SetUserNameNil(b bool)`

 SetUserNameNil sets the value for UserName to be an explicit nil

### UnsetUserName
`func (o *AuditApiResult) UnsetUserName()`

UnsetUserName ensures that no value is present for UserName, not even an explicit nil
### GetDateTime

`func (o *AuditApiResult) GetDateTime() time.Time`

GetDateTime returns the DateTime field if non-nil, zero value otherwise.

### GetDateTimeOk

`func (o *AuditApiResult) GetDateTimeOk() (*time.Time, bool)`

GetDateTimeOk returns a tuple with the DateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTime

`func (o *AuditApiResult) SetDateTime(v time.Time)`

SetDateTime sets DateTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


