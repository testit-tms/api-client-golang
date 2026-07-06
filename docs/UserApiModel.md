# UserApiModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**UserName** | **string** |  | 
**FirstName** | **string** |  | 
**LastName** | **string** |  | 
**DisplayName** | **string** |  | 
**Email** | **string** |  | 
**ProviderId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewUserApiModel

`func NewUserApiModel(id string, userName string, firstName string, lastName string, displayName string, email string, ) *UserApiModel`

NewUserApiModel instantiates a new UserApiModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserApiModelWithDefaults

`func NewUserApiModelWithDefaults() *UserApiModel`

NewUserApiModelWithDefaults instantiates a new UserApiModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserApiModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserApiModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserApiModel) SetId(v string)`

SetId sets Id field to given value.


### GetUserName

`func (o *UserApiModel) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *UserApiModel) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *UserApiModel) SetUserName(v string)`

SetUserName sets UserName field to given value.


### GetFirstName

`func (o *UserApiModel) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *UserApiModel) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *UserApiModel) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetLastName

`func (o *UserApiModel) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *UserApiModel) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *UserApiModel) SetLastName(v string)`

SetLastName sets LastName field to given value.


### GetDisplayName

`func (o *UserApiModel) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *UserApiModel) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *UserApiModel) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetEmail

`func (o *UserApiModel) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserApiModel) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserApiModel) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetProviderId

`func (o *UserApiModel) GetProviderId() string`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *UserApiModel) GetProviderIdOk() (*string, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *UserApiModel) SetProviderId(v string)`

SetProviderId sets ProviderId field to given value.

### HasProviderId

`func (o *UserApiModel) HasProviderId() bool`

HasProviderId returns a boolean if a field has been set.

### SetProviderIdNil

`func (o *UserApiModel) SetProviderIdNil(b bool)`

 SetProviderIdNil sets the value for ProviderId to be an explicit nil

### UnsetProviderId
`func (o *UserApiModel) UnsetProviderId()`

UnsetProviderId ensures that no value is present for ProviderId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


