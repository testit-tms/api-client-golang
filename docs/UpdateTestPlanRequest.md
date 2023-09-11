# UpdateTestPlanRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**LockedById** | Pointer to **NullableString** |  | [optional] 
**Tags** | Pointer to [**[]TagShortModel**](TagShortModel.md) |  | [optional] 
**Name** | **string** |  | 
**StartDate** | Pointer to **NullableTime** | Used for analytics | [optional] 
**EndDate** | Pointer to **NullableTime** | Used for analytics | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Build** | Pointer to **NullableString** |  | [optional] 
**ProjectId** | **string** |  | 
**ProductName** | Pointer to **NullableString** |  | [optional] 
**HasAutomaticDurationTimer** | Pointer to **NullableBool** |  | [optional] 
**Attributes** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewUpdateTestPlanRequest

`func NewUpdateTestPlanRequest(id string, name string, projectId string, ) *UpdateTestPlanRequest`

NewUpdateTestPlanRequest instantiates a new UpdateTestPlanRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateTestPlanRequestWithDefaults

`func NewUpdateTestPlanRequestWithDefaults() *UpdateTestPlanRequest`

NewUpdateTestPlanRequestWithDefaults instantiates a new UpdateTestPlanRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdateTestPlanRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateTestPlanRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateTestPlanRequest) SetId(v string)`

SetId sets Id field to given value.


### GetLockedById

`func (o *UpdateTestPlanRequest) GetLockedById() string`

GetLockedById returns the LockedById field if non-nil, zero value otherwise.

### GetLockedByIdOk

`func (o *UpdateTestPlanRequest) GetLockedByIdOk() (*string, bool)`

GetLockedByIdOk returns a tuple with the LockedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockedById

`func (o *UpdateTestPlanRequest) SetLockedById(v string)`

SetLockedById sets LockedById field to given value.

### HasLockedById

`func (o *UpdateTestPlanRequest) HasLockedById() bool`

HasLockedById returns a boolean if a field has been set.

### SetLockedByIdNil

`func (o *UpdateTestPlanRequest) SetLockedByIdNil(b bool)`

 SetLockedByIdNil sets the value for LockedById to be an explicit nil

### UnsetLockedById
`func (o *UpdateTestPlanRequest) UnsetLockedById()`

UnsetLockedById ensures that no value is present for LockedById, not even an explicit nil
### GetTags

`func (o *UpdateTestPlanRequest) GetTags() []TagShortModel`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UpdateTestPlanRequest) GetTagsOk() (*[]TagShortModel, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UpdateTestPlanRequest) SetTags(v []TagShortModel)`

SetTags sets Tags field to given value.

### HasTags

`func (o *UpdateTestPlanRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *UpdateTestPlanRequest) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *UpdateTestPlanRequest) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetName

`func (o *UpdateTestPlanRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateTestPlanRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateTestPlanRequest) SetName(v string)`

SetName sets Name field to given value.


### GetStartDate

`func (o *UpdateTestPlanRequest) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *UpdateTestPlanRequest) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *UpdateTestPlanRequest) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *UpdateTestPlanRequest) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### SetStartDateNil

`func (o *UpdateTestPlanRequest) SetStartDateNil(b bool)`

 SetStartDateNil sets the value for StartDate to be an explicit nil

### UnsetStartDate
`func (o *UpdateTestPlanRequest) UnsetStartDate()`

UnsetStartDate ensures that no value is present for StartDate, not even an explicit nil
### GetEndDate

`func (o *UpdateTestPlanRequest) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *UpdateTestPlanRequest) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *UpdateTestPlanRequest) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *UpdateTestPlanRequest) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### SetEndDateNil

`func (o *UpdateTestPlanRequest) SetEndDateNil(b bool)`

 SetEndDateNil sets the value for EndDate to be an explicit nil

### UnsetEndDate
`func (o *UpdateTestPlanRequest) UnsetEndDate()`

UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
### GetDescription

`func (o *UpdateTestPlanRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateTestPlanRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateTestPlanRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateTestPlanRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *UpdateTestPlanRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *UpdateTestPlanRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetBuild

`func (o *UpdateTestPlanRequest) GetBuild() string`

GetBuild returns the Build field if non-nil, zero value otherwise.

### GetBuildOk

`func (o *UpdateTestPlanRequest) GetBuildOk() (*string, bool)`

GetBuildOk returns a tuple with the Build field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuild

`func (o *UpdateTestPlanRequest) SetBuild(v string)`

SetBuild sets Build field to given value.

### HasBuild

`func (o *UpdateTestPlanRequest) HasBuild() bool`

HasBuild returns a boolean if a field has been set.

### SetBuildNil

`func (o *UpdateTestPlanRequest) SetBuildNil(b bool)`

 SetBuildNil sets the value for Build to be an explicit nil

### UnsetBuild
`func (o *UpdateTestPlanRequest) UnsetBuild()`

UnsetBuild ensures that no value is present for Build, not even an explicit nil
### GetProjectId

`func (o *UpdateTestPlanRequest) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *UpdateTestPlanRequest) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *UpdateTestPlanRequest) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetProductName

`func (o *UpdateTestPlanRequest) GetProductName() string`

GetProductName returns the ProductName field if non-nil, zero value otherwise.

### GetProductNameOk

`func (o *UpdateTestPlanRequest) GetProductNameOk() (*string, bool)`

GetProductNameOk returns a tuple with the ProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductName

`func (o *UpdateTestPlanRequest) SetProductName(v string)`

SetProductName sets ProductName field to given value.

### HasProductName

`func (o *UpdateTestPlanRequest) HasProductName() bool`

HasProductName returns a boolean if a field has been set.

### SetProductNameNil

`func (o *UpdateTestPlanRequest) SetProductNameNil(b bool)`

 SetProductNameNil sets the value for ProductName to be an explicit nil

### UnsetProductName
`func (o *UpdateTestPlanRequest) UnsetProductName()`

UnsetProductName ensures that no value is present for ProductName, not even an explicit nil
### GetHasAutomaticDurationTimer

`func (o *UpdateTestPlanRequest) GetHasAutomaticDurationTimer() bool`

GetHasAutomaticDurationTimer returns the HasAutomaticDurationTimer field if non-nil, zero value otherwise.

### GetHasAutomaticDurationTimerOk

`func (o *UpdateTestPlanRequest) GetHasAutomaticDurationTimerOk() (*bool, bool)`

GetHasAutomaticDurationTimerOk returns a tuple with the HasAutomaticDurationTimer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAutomaticDurationTimer

`func (o *UpdateTestPlanRequest) SetHasAutomaticDurationTimer(v bool)`

SetHasAutomaticDurationTimer sets HasAutomaticDurationTimer field to given value.

### HasHasAutomaticDurationTimer

`func (o *UpdateTestPlanRequest) HasHasAutomaticDurationTimer() bool`

HasHasAutomaticDurationTimer returns a boolean if a field has been set.

### SetHasAutomaticDurationTimerNil

`func (o *UpdateTestPlanRequest) SetHasAutomaticDurationTimerNil(b bool)`

 SetHasAutomaticDurationTimerNil sets the value for HasAutomaticDurationTimer to be an explicit nil

### UnsetHasAutomaticDurationTimer
`func (o *UpdateTestPlanRequest) UnsetHasAutomaticDurationTimer()`

UnsetHasAutomaticDurationTimer ensures that no value is present for HasAutomaticDurationTimer, not even an explicit nil
### GetAttributes

`func (o *UpdateTestPlanRequest) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *UpdateTestPlanRequest) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *UpdateTestPlanRequest) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *UpdateTestPlanRequest) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### SetAttributesNil

`func (o *UpdateTestPlanRequest) SetAttributesNil(b bool)`

 SetAttributesNil sets the value for Attributes to be an explicit nil

### UnsetAttributes
`func (o *UpdateTestPlanRequest) UnsetAttributes()`

UnsetAttributes ensures that no value is present for Attributes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


