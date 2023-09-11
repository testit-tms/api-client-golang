# AutotestsSelectModelFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectIds** | Pointer to **[]string** | Specifies an autotest projects IDs to search for | [optional] 
**ExternalIds** | Pointer to **[]string** | Specifies an autotest external IDs to search for | [optional] 
**GlobalIds** | Pointer to **[]int64** | Specifies an autotest global IDs to search for | [optional] 
**Name** | Pointer to **NullableString** | Specifies an autotest name to search for | [optional] 
**IsFlaky** | Pointer to **NullableBool** | Specifies an autotest flaky status to search for | [optional] 
**MustBeApproved** | Pointer to **NullableBool** | Specifies an autotest unapproved changes status to search for | [optional] 
**StabilityPercentage** | Pointer to [**NullableAutotestFilterModelStabilityPercentage**](AutotestFilterModelStabilityPercentage.md) |  | [optional] 
**CreatedDate** | Pointer to [**NullableAutotestFilterModelCreatedDate**](AutotestFilterModelCreatedDate.md) |  | [optional] 
**CreatedByIds** | Pointer to **[]string** | Specifies an autotest creator IDs to search for | [optional] 
**ModifiedDate** | Pointer to [**NullableAutotestFilterModelModifiedDate**](AutotestFilterModelModifiedDate.md) |  | [optional] 
**ModifiedByIds** | Pointer to **[]string** | Specifies an autotest last editor IDs to search for | [optional] 
**IsDeleted** | Pointer to **NullableBool** | Specifies an autotest deleted status to search for | [optional] 
**Namespace** | Pointer to **NullableString** | Specifies an autotest namespace to search for | [optional] 
**IsEmptyNamespace** | Pointer to **NullableBool** | Specifies an autotest namespace name presence status to search for | [optional] 
**ClassName** | Pointer to **NullableString** | Specifies an autotest class name to search for | [optional] 
**IsEmptyClassName** | Pointer to **NullableBool** | Specifies an autotest class name presence status to search for | [optional] 
**LastTestResultOutcome** | Pointer to [**NullableAutotestResultOutcome**](AutotestResultOutcome.md) |  | [optional] 
**ExternalKey** | Pointer to **NullableString** | Specifies an autotest external key to search for | [optional] 

## Methods

### NewAutotestsSelectModelFilter

`func NewAutotestsSelectModelFilter() *AutotestsSelectModelFilter`

NewAutotestsSelectModelFilter instantiates a new AutotestsSelectModelFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutotestsSelectModelFilterWithDefaults

`func NewAutotestsSelectModelFilterWithDefaults() *AutotestsSelectModelFilter`

NewAutotestsSelectModelFilterWithDefaults instantiates a new AutotestsSelectModelFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectIds

`func (o *AutotestsSelectModelFilter) GetProjectIds() []string`

GetProjectIds returns the ProjectIds field if non-nil, zero value otherwise.

### GetProjectIdsOk

`func (o *AutotestsSelectModelFilter) GetProjectIdsOk() (*[]string, bool)`

GetProjectIdsOk returns a tuple with the ProjectIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectIds

`func (o *AutotestsSelectModelFilter) SetProjectIds(v []string)`

SetProjectIds sets ProjectIds field to given value.

### HasProjectIds

`func (o *AutotestsSelectModelFilter) HasProjectIds() bool`

HasProjectIds returns a boolean if a field has been set.

### SetProjectIdsNil

`func (o *AutotestsSelectModelFilter) SetProjectIdsNil(b bool)`

 SetProjectIdsNil sets the value for ProjectIds to be an explicit nil

### UnsetProjectIds
`func (o *AutotestsSelectModelFilter) UnsetProjectIds()`

UnsetProjectIds ensures that no value is present for ProjectIds, not even an explicit nil
### GetExternalIds

`func (o *AutotestsSelectModelFilter) GetExternalIds() []string`

GetExternalIds returns the ExternalIds field if non-nil, zero value otherwise.

### GetExternalIdsOk

`func (o *AutotestsSelectModelFilter) GetExternalIdsOk() (*[]string, bool)`

GetExternalIdsOk returns a tuple with the ExternalIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIds

`func (o *AutotestsSelectModelFilter) SetExternalIds(v []string)`

SetExternalIds sets ExternalIds field to given value.

### HasExternalIds

`func (o *AutotestsSelectModelFilter) HasExternalIds() bool`

HasExternalIds returns a boolean if a field has been set.

### SetExternalIdsNil

`func (o *AutotestsSelectModelFilter) SetExternalIdsNil(b bool)`

 SetExternalIdsNil sets the value for ExternalIds to be an explicit nil

### UnsetExternalIds
`func (o *AutotestsSelectModelFilter) UnsetExternalIds()`

UnsetExternalIds ensures that no value is present for ExternalIds, not even an explicit nil
### GetGlobalIds

`func (o *AutotestsSelectModelFilter) GetGlobalIds() []int64`

GetGlobalIds returns the GlobalIds field if non-nil, zero value otherwise.

### GetGlobalIdsOk

`func (o *AutotestsSelectModelFilter) GetGlobalIdsOk() (*[]int64, bool)`

GetGlobalIdsOk returns a tuple with the GlobalIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalIds

`func (o *AutotestsSelectModelFilter) SetGlobalIds(v []int64)`

SetGlobalIds sets GlobalIds field to given value.

### HasGlobalIds

`func (o *AutotestsSelectModelFilter) HasGlobalIds() bool`

HasGlobalIds returns a boolean if a field has been set.

### SetGlobalIdsNil

`func (o *AutotestsSelectModelFilter) SetGlobalIdsNil(b bool)`

 SetGlobalIdsNil sets the value for GlobalIds to be an explicit nil

### UnsetGlobalIds
`func (o *AutotestsSelectModelFilter) UnsetGlobalIds()`

UnsetGlobalIds ensures that no value is present for GlobalIds, not even an explicit nil
### GetName

`func (o *AutotestsSelectModelFilter) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AutotestsSelectModelFilter) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AutotestsSelectModelFilter) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AutotestsSelectModelFilter) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *AutotestsSelectModelFilter) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AutotestsSelectModelFilter) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetIsFlaky

`func (o *AutotestsSelectModelFilter) GetIsFlaky() bool`

GetIsFlaky returns the IsFlaky field if non-nil, zero value otherwise.

### GetIsFlakyOk

`func (o *AutotestsSelectModelFilter) GetIsFlakyOk() (*bool, bool)`

GetIsFlakyOk returns a tuple with the IsFlaky field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFlaky

`func (o *AutotestsSelectModelFilter) SetIsFlaky(v bool)`

SetIsFlaky sets IsFlaky field to given value.

### HasIsFlaky

`func (o *AutotestsSelectModelFilter) HasIsFlaky() bool`

HasIsFlaky returns a boolean if a field has been set.

### SetIsFlakyNil

`func (o *AutotestsSelectModelFilter) SetIsFlakyNil(b bool)`

 SetIsFlakyNil sets the value for IsFlaky to be an explicit nil

### UnsetIsFlaky
`func (o *AutotestsSelectModelFilter) UnsetIsFlaky()`

UnsetIsFlaky ensures that no value is present for IsFlaky, not even an explicit nil
### GetMustBeApproved

`func (o *AutotestsSelectModelFilter) GetMustBeApproved() bool`

GetMustBeApproved returns the MustBeApproved field if non-nil, zero value otherwise.

### GetMustBeApprovedOk

`func (o *AutotestsSelectModelFilter) GetMustBeApprovedOk() (*bool, bool)`

GetMustBeApprovedOk returns a tuple with the MustBeApproved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMustBeApproved

`func (o *AutotestsSelectModelFilter) SetMustBeApproved(v bool)`

SetMustBeApproved sets MustBeApproved field to given value.

### HasMustBeApproved

`func (o *AutotestsSelectModelFilter) HasMustBeApproved() bool`

HasMustBeApproved returns a boolean if a field has been set.

### SetMustBeApprovedNil

`func (o *AutotestsSelectModelFilter) SetMustBeApprovedNil(b bool)`

 SetMustBeApprovedNil sets the value for MustBeApproved to be an explicit nil

### UnsetMustBeApproved
`func (o *AutotestsSelectModelFilter) UnsetMustBeApproved()`

UnsetMustBeApproved ensures that no value is present for MustBeApproved, not even an explicit nil
### GetStabilityPercentage

`func (o *AutotestsSelectModelFilter) GetStabilityPercentage() AutotestFilterModelStabilityPercentage`

GetStabilityPercentage returns the StabilityPercentage field if non-nil, zero value otherwise.

### GetStabilityPercentageOk

`func (o *AutotestsSelectModelFilter) GetStabilityPercentageOk() (*AutotestFilterModelStabilityPercentage, bool)`

GetStabilityPercentageOk returns a tuple with the StabilityPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStabilityPercentage

`func (o *AutotestsSelectModelFilter) SetStabilityPercentage(v AutotestFilterModelStabilityPercentage)`

SetStabilityPercentage sets StabilityPercentage field to given value.

### HasStabilityPercentage

`func (o *AutotestsSelectModelFilter) HasStabilityPercentage() bool`

HasStabilityPercentage returns a boolean if a field has been set.

### SetStabilityPercentageNil

`func (o *AutotestsSelectModelFilter) SetStabilityPercentageNil(b bool)`

 SetStabilityPercentageNil sets the value for StabilityPercentage to be an explicit nil

### UnsetStabilityPercentage
`func (o *AutotestsSelectModelFilter) UnsetStabilityPercentage()`

UnsetStabilityPercentage ensures that no value is present for StabilityPercentage, not even an explicit nil
### GetCreatedDate

`func (o *AutotestsSelectModelFilter) GetCreatedDate() AutotestFilterModelCreatedDate`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *AutotestsSelectModelFilter) GetCreatedDateOk() (*AutotestFilterModelCreatedDate, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *AutotestsSelectModelFilter) SetCreatedDate(v AutotestFilterModelCreatedDate)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *AutotestsSelectModelFilter) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### SetCreatedDateNil

`func (o *AutotestsSelectModelFilter) SetCreatedDateNil(b bool)`

 SetCreatedDateNil sets the value for CreatedDate to be an explicit nil

### UnsetCreatedDate
`func (o *AutotestsSelectModelFilter) UnsetCreatedDate()`

UnsetCreatedDate ensures that no value is present for CreatedDate, not even an explicit nil
### GetCreatedByIds

`func (o *AutotestsSelectModelFilter) GetCreatedByIds() []string`

GetCreatedByIds returns the CreatedByIds field if non-nil, zero value otherwise.

### GetCreatedByIdsOk

`func (o *AutotestsSelectModelFilter) GetCreatedByIdsOk() (*[]string, bool)`

GetCreatedByIdsOk returns a tuple with the CreatedByIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByIds

`func (o *AutotestsSelectModelFilter) SetCreatedByIds(v []string)`

SetCreatedByIds sets CreatedByIds field to given value.

### HasCreatedByIds

`func (o *AutotestsSelectModelFilter) HasCreatedByIds() bool`

HasCreatedByIds returns a boolean if a field has been set.

### SetCreatedByIdsNil

`func (o *AutotestsSelectModelFilter) SetCreatedByIdsNil(b bool)`

 SetCreatedByIdsNil sets the value for CreatedByIds to be an explicit nil

### UnsetCreatedByIds
`func (o *AutotestsSelectModelFilter) UnsetCreatedByIds()`

UnsetCreatedByIds ensures that no value is present for CreatedByIds, not even an explicit nil
### GetModifiedDate

`func (o *AutotestsSelectModelFilter) GetModifiedDate() AutotestFilterModelModifiedDate`

GetModifiedDate returns the ModifiedDate field if non-nil, zero value otherwise.

### GetModifiedDateOk

`func (o *AutotestsSelectModelFilter) GetModifiedDateOk() (*AutotestFilterModelModifiedDate, bool)`

GetModifiedDateOk returns a tuple with the ModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedDate

`func (o *AutotestsSelectModelFilter) SetModifiedDate(v AutotestFilterModelModifiedDate)`

SetModifiedDate sets ModifiedDate field to given value.

### HasModifiedDate

`func (o *AutotestsSelectModelFilter) HasModifiedDate() bool`

HasModifiedDate returns a boolean if a field has been set.

### SetModifiedDateNil

`func (o *AutotestsSelectModelFilter) SetModifiedDateNil(b bool)`

 SetModifiedDateNil sets the value for ModifiedDate to be an explicit nil

### UnsetModifiedDate
`func (o *AutotestsSelectModelFilter) UnsetModifiedDate()`

UnsetModifiedDate ensures that no value is present for ModifiedDate, not even an explicit nil
### GetModifiedByIds

`func (o *AutotestsSelectModelFilter) GetModifiedByIds() []string`

GetModifiedByIds returns the ModifiedByIds field if non-nil, zero value otherwise.

### GetModifiedByIdsOk

`func (o *AutotestsSelectModelFilter) GetModifiedByIdsOk() (*[]string, bool)`

GetModifiedByIdsOk returns a tuple with the ModifiedByIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedByIds

`func (o *AutotestsSelectModelFilter) SetModifiedByIds(v []string)`

SetModifiedByIds sets ModifiedByIds field to given value.

### HasModifiedByIds

`func (o *AutotestsSelectModelFilter) HasModifiedByIds() bool`

HasModifiedByIds returns a boolean if a field has been set.

### SetModifiedByIdsNil

`func (o *AutotestsSelectModelFilter) SetModifiedByIdsNil(b bool)`

 SetModifiedByIdsNil sets the value for ModifiedByIds to be an explicit nil

### UnsetModifiedByIds
`func (o *AutotestsSelectModelFilter) UnsetModifiedByIds()`

UnsetModifiedByIds ensures that no value is present for ModifiedByIds, not even an explicit nil
### GetIsDeleted

`func (o *AutotestsSelectModelFilter) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *AutotestsSelectModelFilter) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *AutotestsSelectModelFilter) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *AutotestsSelectModelFilter) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### SetIsDeletedNil

`func (o *AutotestsSelectModelFilter) SetIsDeletedNil(b bool)`

 SetIsDeletedNil sets the value for IsDeleted to be an explicit nil

### UnsetIsDeleted
`func (o *AutotestsSelectModelFilter) UnsetIsDeleted()`

UnsetIsDeleted ensures that no value is present for IsDeleted, not even an explicit nil
### GetNamespace

`func (o *AutotestsSelectModelFilter) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *AutotestsSelectModelFilter) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *AutotestsSelectModelFilter) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *AutotestsSelectModelFilter) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### SetNamespaceNil

`func (o *AutotestsSelectModelFilter) SetNamespaceNil(b bool)`

 SetNamespaceNil sets the value for Namespace to be an explicit nil

### UnsetNamespace
`func (o *AutotestsSelectModelFilter) UnsetNamespace()`

UnsetNamespace ensures that no value is present for Namespace, not even an explicit nil
### GetIsEmptyNamespace

`func (o *AutotestsSelectModelFilter) GetIsEmptyNamespace() bool`

GetIsEmptyNamespace returns the IsEmptyNamespace field if non-nil, zero value otherwise.

### GetIsEmptyNamespaceOk

`func (o *AutotestsSelectModelFilter) GetIsEmptyNamespaceOk() (*bool, bool)`

GetIsEmptyNamespaceOk returns a tuple with the IsEmptyNamespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEmptyNamespace

`func (o *AutotestsSelectModelFilter) SetIsEmptyNamespace(v bool)`

SetIsEmptyNamespace sets IsEmptyNamespace field to given value.

### HasIsEmptyNamespace

`func (o *AutotestsSelectModelFilter) HasIsEmptyNamespace() bool`

HasIsEmptyNamespace returns a boolean if a field has been set.

### SetIsEmptyNamespaceNil

`func (o *AutotestsSelectModelFilter) SetIsEmptyNamespaceNil(b bool)`

 SetIsEmptyNamespaceNil sets the value for IsEmptyNamespace to be an explicit nil

### UnsetIsEmptyNamespace
`func (o *AutotestsSelectModelFilter) UnsetIsEmptyNamespace()`

UnsetIsEmptyNamespace ensures that no value is present for IsEmptyNamespace, not even an explicit nil
### GetClassName

`func (o *AutotestsSelectModelFilter) GetClassName() string`

GetClassName returns the ClassName field if non-nil, zero value otherwise.

### GetClassNameOk

`func (o *AutotestsSelectModelFilter) GetClassNameOk() (*string, bool)`

GetClassNameOk returns a tuple with the ClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassName

`func (o *AutotestsSelectModelFilter) SetClassName(v string)`

SetClassName sets ClassName field to given value.

### HasClassName

`func (o *AutotestsSelectModelFilter) HasClassName() bool`

HasClassName returns a boolean if a field has been set.

### SetClassNameNil

`func (o *AutotestsSelectModelFilter) SetClassNameNil(b bool)`

 SetClassNameNil sets the value for ClassName to be an explicit nil

### UnsetClassName
`func (o *AutotestsSelectModelFilter) UnsetClassName()`

UnsetClassName ensures that no value is present for ClassName, not even an explicit nil
### GetIsEmptyClassName

`func (o *AutotestsSelectModelFilter) GetIsEmptyClassName() bool`

GetIsEmptyClassName returns the IsEmptyClassName field if non-nil, zero value otherwise.

### GetIsEmptyClassNameOk

`func (o *AutotestsSelectModelFilter) GetIsEmptyClassNameOk() (*bool, bool)`

GetIsEmptyClassNameOk returns a tuple with the IsEmptyClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEmptyClassName

`func (o *AutotestsSelectModelFilter) SetIsEmptyClassName(v bool)`

SetIsEmptyClassName sets IsEmptyClassName field to given value.

### HasIsEmptyClassName

`func (o *AutotestsSelectModelFilter) HasIsEmptyClassName() bool`

HasIsEmptyClassName returns a boolean if a field has been set.

### SetIsEmptyClassNameNil

`func (o *AutotestsSelectModelFilter) SetIsEmptyClassNameNil(b bool)`

 SetIsEmptyClassNameNil sets the value for IsEmptyClassName to be an explicit nil

### UnsetIsEmptyClassName
`func (o *AutotestsSelectModelFilter) UnsetIsEmptyClassName()`

UnsetIsEmptyClassName ensures that no value is present for IsEmptyClassName, not even an explicit nil
### GetLastTestResultOutcome

`func (o *AutotestsSelectModelFilter) GetLastTestResultOutcome() AutotestResultOutcome`

GetLastTestResultOutcome returns the LastTestResultOutcome field if non-nil, zero value otherwise.

### GetLastTestResultOutcomeOk

`func (o *AutotestsSelectModelFilter) GetLastTestResultOutcomeOk() (*AutotestResultOutcome, bool)`

GetLastTestResultOutcomeOk returns a tuple with the LastTestResultOutcome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTestResultOutcome

`func (o *AutotestsSelectModelFilter) SetLastTestResultOutcome(v AutotestResultOutcome)`

SetLastTestResultOutcome sets LastTestResultOutcome field to given value.

### HasLastTestResultOutcome

`func (o *AutotestsSelectModelFilter) HasLastTestResultOutcome() bool`

HasLastTestResultOutcome returns a boolean if a field has been set.

### SetLastTestResultOutcomeNil

`func (o *AutotestsSelectModelFilter) SetLastTestResultOutcomeNil(b bool)`

 SetLastTestResultOutcomeNil sets the value for LastTestResultOutcome to be an explicit nil

### UnsetLastTestResultOutcome
`func (o *AutotestsSelectModelFilter) UnsetLastTestResultOutcome()`

UnsetLastTestResultOutcome ensures that no value is present for LastTestResultOutcome, not even an explicit nil
### GetExternalKey

`func (o *AutotestsSelectModelFilter) GetExternalKey() string`

GetExternalKey returns the ExternalKey field if non-nil, zero value otherwise.

### GetExternalKeyOk

`func (o *AutotestsSelectModelFilter) GetExternalKeyOk() (*string, bool)`

GetExternalKeyOk returns a tuple with the ExternalKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalKey

`func (o *AutotestsSelectModelFilter) SetExternalKey(v string)`

SetExternalKey sets ExternalKey field to given value.

### HasExternalKey

`func (o *AutotestsSelectModelFilter) HasExternalKey() bool`

HasExternalKey returns a boolean if a field has been set.

### SetExternalKeyNil

`func (o *AutotestsSelectModelFilter) SetExternalKeyNil(b bool)`

 SetExternalKeyNil sets the value for ExternalKey to be an explicit nil

### UnsetExternalKey
`func (o *AutotestsSelectModelFilter) UnsetExternalKey()`

UnsetExternalKey ensures that no value is present for ExternalKey, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


