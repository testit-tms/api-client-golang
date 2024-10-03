# AutotestSelectModelFilter

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

### NewAutotestSelectModelFilter

`func NewAutotestSelectModelFilter() *AutotestSelectModelFilter`

NewAutotestSelectModelFilter instantiates a new AutotestSelectModelFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutotestSelectModelFilterWithDefaults

`func NewAutotestSelectModelFilterWithDefaults() *AutotestSelectModelFilter`

NewAutotestSelectModelFilterWithDefaults instantiates a new AutotestSelectModelFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectIds

`func (o *AutotestSelectModelFilter) GetProjectIds() []string`

GetProjectIds returns the ProjectIds field if non-nil, zero value otherwise.

### GetProjectIdsOk

`func (o *AutotestSelectModelFilter) GetProjectIdsOk() (*[]string, bool)`

GetProjectIdsOk returns a tuple with the ProjectIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectIds

`func (o *AutotestSelectModelFilter) SetProjectIds(v []string)`

SetProjectIds sets ProjectIds field to given value.

### HasProjectIds

`func (o *AutotestSelectModelFilter) HasProjectIds() bool`

HasProjectIds returns a boolean if a field has been set.

### SetProjectIdsNil

`func (o *AutotestSelectModelFilter) SetProjectIdsNil(b bool)`

 SetProjectIdsNil sets the value for ProjectIds to be an explicit nil

### UnsetProjectIds
`func (o *AutotestSelectModelFilter) UnsetProjectIds()`

UnsetProjectIds ensures that no value is present for ProjectIds, not even an explicit nil
### GetExternalIds

`func (o *AutotestSelectModelFilter) GetExternalIds() []string`

GetExternalIds returns the ExternalIds field if non-nil, zero value otherwise.

### GetExternalIdsOk

`func (o *AutotestSelectModelFilter) GetExternalIdsOk() (*[]string, bool)`

GetExternalIdsOk returns a tuple with the ExternalIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIds

`func (o *AutotestSelectModelFilter) SetExternalIds(v []string)`

SetExternalIds sets ExternalIds field to given value.

### HasExternalIds

`func (o *AutotestSelectModelFilter) HasExternalIds() bool`

HasExternalIds returns a boolean if a field has been set.

### SetExternalIdsNil

`func (o *AutotestSelectModelFilter) SetExternalIdsNil(b bool)`

 SetExternalIdsNil sets the value for ExternalIds to be an explicit nil

### UnsetExternalIds
`func (o *AutotestSelectModelFilter) UnsetExternalIds()`

UnsetExternalIds ensures that no value is present for ExternalIds, not even an explicit nil
### GetGlobalIds

`func (o *AutotestSelectModelFilter) GetGlobalIds() []int64`

GetGlobalIds returns the GlobalIds field if non-nil, zero value otherwise.

### GetGlobalIdsOk

`func (o *AutotestSelectModelFilter) GetGlobalIdsOk() (*[]int64, bool)`

GetGlobalIdsOk returns a tuple with the GlobalIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalIds

`func (o *AutotestSelectModelFilter) SetGlobalIds(v []int64)`

SetGlobalIds sets GlobalIds field to given value.

### HasGlobalIds

`func (o *AutotestSelectModelFilter) HasGlobalIds() bool`

HasGlobalIds returns a boolean if a field has been set.

### SetGlobalIdsNil

`func (o *AutotestSelectModelFilter) SetGlobalIdsNil(b bool)`

 SetGlobalIdsNil sets the value for GlobalIds to be an explicit nil

### UnsetGlobalIds
`func (o *AutotestSelectModelFilter) UnsetGlobalIds()`

UnsetGlobalIds ensures that no value is present for GlobalIds, not even an explicit nil
### GetName

`func (o *AutotestSelectModelFilter) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AutotestSelectModelFilter) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AutotestSelectModelFilter) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AutotestSelectModelFilter) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *AutotestSelectModelFilter) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AutotestSelectModelFilter) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetIsFlaky

`func (o *AutotestSelectModelFilter) GetIsFlaky() bool`

GetIsFlaky returns the IsFlaky field if non-nil, zero value otherwise.

### GetIsFlakyOk

`func (o *AutotestSelectModelFilter) GetIsFlakyOk() (*bool, bool)`

GetIsFlakyOk returns a tuple with the IsFlaky field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFlaky

`func (o *AutotestSelectModelFilter) SetIsFlaky(v bool)`

SetIsFlaky sets IsFlaky field to given value.

### HasIsFlaky

`func (o *AutotestSelectModelFilter) HasIsFlaky() bool`

HasIsFlaky returns a boolean if a field has been set.

### SetIsFlakyNil

`func (o *AutotestSelectModelFilter) SetIsFlakyNil(b bool)`

 SetIsFlakyNil sets the value for IsFlaky to be an explicit nil

### UnsetIsFlaky
`func (o *AutotestSelectModelFilter) UnsetIsFlaky()`

UnsetIsFlaky ensures that no value is present for IsFlaky, not even an explicit nil
### GetMustBeApproved

`func (o *AutotestSelectModelFilter) GetMustBeApproved() bool`

GetMustBeApproved returns the MustBeApproved field if non-nil, zero value otherwise.

### GetMustBeApprovedOk

`func (o *AutotestSelectModelFilter) GetMustBeApprovedOk() (*bool, bool)`

GetMustBeApprovedOk returns a tuple with the MustBeApproved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMustBeApproved

`func (o *AutotestSelectModelFilter) SetMustBeApproved(v bool)`

SetMustBeApproved sets MustBeApproved field to given value.

### HasMustBeApproved

`func (o *AutotestSelectModelFilter) HasMustBeApproved() bool`

HasMustBeApproved returns a boolean if a field has been set.

### SetMustBeApprovedNil

`func (o *AutotestSelectModelFilter) SetMustBeApprovedNil(b bool)`

 SetMustBeApprovedNil sets the value for MustBeApproved to be an explicit nil

### UnsetMustBeApproved
`func (o *AutotestSelectModelFilter) UnsetMustBeApproved()`

UnsetMustBeApproved ensures that no value is present for MustBeApproved, not even an explicit nil
### GetStabilityPercentage

`func (o *AutotestSelectModelFilter) GetStabilityPercentage() AutotestFilterModelStabilityPercentage`

GetStabilityPercentage returns the StabilityPercentage field if non-nil, zero value otherwise.

### GetStabilityPercentageOk

`func (o *AutotestSelectModelFilter) GetStabilityPercentageOk() (*AutotestFilterModelStabilityPercentage, bool)`

GetStabilityPercentageOk returns a tuple with the StabilityPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStabilityPercentage

`func (o *AutotestSelectModelFilter) SetStabilityPercentage(v AutotestFilterModelStabilityPercentage)`

SetStabilityPercentage sets StabilityPercentage field to given value.

### HasStabilityPercentage

`func (o *AutotestSelectModelFilter) HasStabilityPercentage() bool`

HasStabilityPercentage returns a boolean if a field has been set.

### SetStabilityPercentageNil

`func (o *AutotestSelectModelFilter) SetStabilityPercentageNil(b bool)`

 SetStabilityPercentageNil sets the value for StabilityPercentage to be an explicit nil

### UnsetStabilityPercentage
`func (o *AutotestSelectModelFilter) UnsetStabilityPercentage()`

UnsetStabilityPercentage ensures that no value is present for StabilityPercentage, not even an explicit nil
### GetCreatedDate

`func (o *AutotestSelectModelFilter) GetCreatedDate() AutotestFilterModelCreatedDate`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *AutotestSelectModelFilter) GetCreatedDateOk() (*AutotestFilterModelCreatedDate, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *AutotestSelectModelFilter) SetCreatedDate(v AutotestFilterModelCreatedDate)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *AutotestSelectModelFilter) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### SetCreatedDateNil

`func (o *AutotestSelectModelFilter) SetCreatedDateNil(b bool)`

 SetCreatedDateNil sets the value for CreatedDate to be an explicit nil

### UnsetCreatedDate
`func (o *AutotestSelectModelFilter) UnsetCreatedDate()`

UnsetCreatedDate ensures that no value is present for CreatedDate, not even an explicit nil
### GetCreatedByIds

`func (o *AutotestSelectModelFilter) GetCreatedByIds() []string`

GetCreatedByIds returns the CreatedByIds field if non-nil, zero value otherwise.

### GetCreatedByIdsOk

`func (o *AutotestSelectModelFilter) GetCreatedByIdsOk() (*[]string, bool)`

GetCreatedByIdsOk returns a tuple with the CreatedByIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByIds

`func (o *AutotestSelectModelFilter) SetCreatedByIds(v []string)`

SetCreatedByIds sets CreatedByIds field to given value.

### HasCreatedByIds

`func (o *AutotestSelectModelFilter) HasCreatedByIds() bool`

HasCreatedByIds returns a boolean if a field has been set.

### SetCreatedByIdsNil

`func (o *AutotestSelectModelFilter) SetCreatedByIdsNil(b bool)`

 SetCreatedByIdsNil sets the value for CreatedByIds to be an explicit nil

### UnsetCreatedByIds
`func (o *AutotestSelectModelFilter) UnsetCreatedByIds()`

UnsetCreatedByIds ensures that no value is present for CreatedByIds, not even an explicit nil
### GetModifiedDate

`func (o *AutotestSelectModelFilter) GetModifiedDate() AutotestFilterModelModifiedDate`

GetModifiedDate returns the ModifiedDate field if non-nil, zero value otherwise.

### GetModifiedDateOk

`func (o *AutotestSelectModelFilter) GetModifiedDateOk() (*AutotestFilterModelModifiedDate, bool)`

GetModifiedDateOk returns a tuple with the ModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedDate

`func (o *AutotestSelectModelFilter) SetModifiedDate(v AutotestFilterModelModifiedDate)`

SetModifiedDate sets ModifiedDate field to given value.

### HasModifiedDate

`func (o *AutotestSelectModelFilter) HasModifiedDate() bool`

HasModifiedDate returns a boolean if a field has been set.

### SetModifiedDateNil

`func (o *AutotestSelectModelFilter) SetModifiedDateNil(b bool)`

 SetModifiedDateNil sets the value for ModifiedDate to be an explicit nil

### UnsetModifiedDate
`func (o *AutotestSelectModelFilter) UnsetModifiedDate()`

UnsetModifiedDate ensures that no value is present for ModifiedDate, not even an explicit nil
### GetModifiedByIds

`func (o *AutotestSelectModelFilter) GetModifiedByIds() []string`

GetModifiedByIds returns the ModifiedByIds field if non-nil, zero value otherwise.

### GetModifiedByIdsOk

`func (o *AutotestSelectModelFilter) GetModifiedByIdsOk() (*[]string, bool)`

GetModifiedByIdsOk returns a tuple with the ModifiedByIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedByIds

`func (o *AutotestSelectModelFilter) SetModifiedByIds(v []string)`

SetModifiedByIds sets ModifiedByIds field to given value.

### HasModifiedByIds

`func (o *AutotestSelectModelFilter) HasModifiedByIds() bool`

HasModifiedByIds returns a boolean if a field has been set.

### SetModifiedByIdsNil

`func (o *AutotestSelectModelFilter) SetModifiedByIdsNil(b bool)`

 SetModifiedByIdsNil sets the value for ModifiedByIds to be an explicit nil

### UnsetModifiedByIds
`func (o *AutotestSelectModelFilter) UnsetModifiedByIds()`

UnsetModifiedByIds ensures that no value is present for ModifiedByIds, not even an explicit nil
### GetIsDeleted

`func (o *AutotestSelectModelFilter) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *AutotestSelectModelFilter) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *AutotestSelectModelFilter) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *AutotestSelectModelFilter) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### SetIsDeletedNil

`func (o *AutotestSelectModelFilter) SetIsDeletedNil(b bool)`

 SetIsDeletedNil sets the value for IsDeleted to be an explicit nil

### UnsetIsDeleted
`func (o *AutotestSelectModelFilter) UnsetIsDeleted()`

UnsetIsDeleted ensures that no value is present for IsDeleted, not even an explicit nil
### GetNamespace

`func (o *AutotestSelectModelFilter) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *AutotestSelectModelFilter) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *AutotestSelectModelFilter) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *AutotestSelectModelFilter) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### SetNamespaceNil

`func (o *AutotestSelectModelFilter) SetNamespaceNil(b bool)`

 SetNamespaceNil sets the value for Namespace to be an explicit nil

### UnsetNamespace
`func (o *AutotestSelectModelFilter) UnsetNamespace()`

UnsetNamespace ensures that no value is present for Namespace, not even an explicit nil
### GetIsEmptyNamespace

`func (o *AutotestSelectModelFilter) GetIsEmptyNamespace() bool`

GetIsEmptyNamespace returns the IsEmptyNamespace field if non-nil, zero value otherwise.

### GetIsEmptyNamespaceOk

`func (o *AutotestSelectModelFilter) GetIsEmptyNamespaceOk() (*bool, bool)`

GetIsEmptyNamespaceOk returns a tuple with the IsEmptyNamespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEmptyNamespace

`func (o *AutotestSelectModelFilter) SetIsEmptyNamespace(v bool)`

SetIsEmptyNamespace sets IsEmptyNamespace field to given value.

### HasIsEmptyNamespace

`func (o *AutotestSelectModelFilter) HasIsEmptyNamespace() bool`

HasIsEmptyNamespace returns a boolean if a field has been set.

### SetIsEmptyNamespaceNil

`func (o *AutotestSelectModelFilter) SetIsEmptyNamespaceNil(b bool)`

 SetIsEmptyNamespaceNil sets the value for IsEmptyNamespace to be an explicit nil

### UnsetIsEmptyNamespace
`func (o *AutotestSelectModelFilter) UnsetIsEmptyNamespace()`

UnsetIsEmptyNamespace ensures that no value is present for IsEmptyNamespace, not even an explicit nil
### GetClassName

`func (o *AutotestSelectModelFilter) GetClassName() string`

GetClassName returns the ClassName field if non-nil, zero value otherwise.

### GetClassNameOk

`func (o *AutotestSelectModelFilter) GetClassNameOk() (*string, bool)`

GetClassNameOk returns a tuple with the ClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassName

`func (o *AutotestSelectModelFilter) SetClassName(v string)`

SetClassName sets ClassName field to given value.

### HasClassName

`func (o *AutotestSelectModelFilter) HasClassName() bool`

HasClassName returns a boolean if a field has been set.

### SetClassNameNil

`func (o *AutotestSelectModelFilter) SetClassNameNil(b bool)`

 SetClassNameNil sets the value for ClassName to be an explicit nil

### UnsetClassName
`func (o *AutotestSelectModelFilter) UnsetClassName()`

UnsetClassName ensures that no value is present for ClassName, not even an explicit nil
### GetIsEmptyClassName

`func (o *AutotestSelectModelFilter) GetIsEmptyClassName() bool`

GetIsEmptyClassName returns the IsEmptyClassName field if non-nil, zero value otherwise.

### GetIsEmptyClassNameOk

`func (o *AutotestSelectModelFilter) GetIsEmptyClassNameOk() (*bool, bool)`

GetIsEmptyClassNameOk returns a tuple with the IsEmptyClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEmptyClassName

`func (o *AutotestSelectModelFilter) SetIsEmptyClassName(v bool)`

SetIsEmptyClassName sets IsEmptyClassName field to given value.

### HasIsEmptyClassName

`func (o *AutotestSelectModelFilter) HasIsEmptyClassName() bool`

HasIsEmptyClassName returns a boolean if a field has been set.

### SetIsEmptyClassNameNil

`func (o *AutotestSelectModelFilter) SetIsEmptyClassNameNil(b bool)`

 SetIsEmptyClassNameNil sets the value for IsEmptyClassName to be an explicit nil

### UnsetIsEmptyClassName
`func (o *AutotestSelectModelFilter) UnsetIsEmptyClassName()`

UnsetIsEmptyClassName ensures that no value is present for IsEmptyClassName, not even an explicit nil
### GetLastTestResultOutcome

`func (o *AutotestSelectModelFilter) GetLastTestResultOutcome() AutotestResultOutcome`

GetLastTestResultOutcome returns the LastTestResultOutcome field if non-nil, zero value otherwise.

### GetLastTestResultOutcomeOk

`func (o *AutotestSelectModelFilter) GetLastTestResultOutcomeOk() (*AutotestResultOutcome, bool)`

GetLastTestResultOutcomeOk returns a tuple with the LastTestResultOutcome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTestResultOutcome

`func (o *AutotestSelectModelFilter) SetLastTestResultOutcome(v AutotestResultOutcome)`

SetLastTestResultOutcome sets LastTestResultOutcome field to given value.

### HasLastTestResultOutcome

`func (o *AutotestSelectModelFilter) HasLastTestResultOutcome() bool`

HasLastTestResultOutcome returns a boolean if a field has been set.

### SetLastTestResultOutcomeNil

`func (o *AutotestSelectModelFilter) SetLastTestResultOutcomeNil(b bool)`

 SetLastTestResultOutcomeNil sets the value for LastTestResultOutcome to be an explicit nil

### UnsetLastTestResultOutcome
`func (o *AutotestSelectModelFilter) UnsetLastTestResultOutcome()`

UnsetLastTestResultOutcome ensures that no value is present for LastTestResultOutcome, not even an explicit nil
### GetExternalKey

`func (o *AutotestSelectModelFilter) GetExternalKey() string`

GetExternalKey returns the ExternalKey field if non-nil, zero value otherwise.

### GetExternalKeyOk

`func (o *AutotestSelectModelFilter) GetExternalKeyOk() (*string, bool)`

GetExternalKeyOk returns a tuple with the ExternalKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalKey

`func (o *AutotestSelectModelFilter) SetExternalKey(v string)`

SetExternalKey sets ExternalKey field to given value.

### HasExternalKey

`func (o *AutotestSelectModelFilter) HasExternalKey() bool`

HasExternalKey returns a boolean if a field has been set.

### SetExternalKeyNil

`func (o *AutotestSelectModelFilter) SetExternalKeyNil(b bool)`

 SetExternalKeyNil sets the value for ExternalKey to be an explicit nil

### UnsetExternalKey
`func (o *AutotestSelectModelFilter) UnsetExternalKey()`

UnsetExternalKey ensures that no value is present for ExternalKey, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


