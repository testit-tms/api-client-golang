# AutoTestFilterApiModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectIds** | Pointer to **[]string** | Specifies an autotest projects IDs to search for | [optional] 
**ExternalIds** | Pointer to **[]string** | Specifies an autotest external IDs to search for | [optional] 
**GlobalIds** | Pointer to **[]int64** | Specifies an autotest global IDs to search for | [optional] 
**Name** | Pointer to **NullableString** | Specifies an autotest name to search for | [optional] 
**IsFlaky** | Pointer to **NullableBool** | Specifies an autotest flaky status to search for | [optional] 
**MustBeApproved** | Pointer to **NullableBool** | Specifies an autotest unapproved changes status to search for | [optional] 
**StabilityPercentage** | Pointer to [**NullableInt64RangeSelectorModel**](Int64RangeSelectorModel.md) | Specifies an autotest range of stability percentage to search for | [optional] 
**CreatedDate** | Pointer to [**NullableDateTimeRangeSelectorModel**](DateTimeRangeSelectorModel.md) | Specifies an autotest range of creation date to search for | [optional] 
**CreatedByIds** | Pointer to **[]string** | Specifies an autotest creator IDs to search for | [optional] 
**ModifiedDate** | Pointer to [**NullableDateTimeRangeSelectorModel**](DateTimeRangeSelectorModel.md) | Specifies an autotest range of last modification date to search for | [optional] 
**ModifiedByIds** | Pointer to **[]string** | Specifies an autotest last editor IDs to search for | [optional] 
**IsDeleted** | Pointer to **NullableBool** | Specifies an autotest deleted status to search for | [optional] 
**Namespace** | Pointer to **NullableString** | Specifies an autotest namespace to search for | [optional] 
**IsEmptyNamespace** | Pointer to **NullableBool** | Specifies an autotest namespace name presence status to search for | [optional] 
**ClassName** | Pointer to **NullableString** | Specifies an autotest class name to search for | [optional] 
**IsEmptyClassName** | Pointer to **NullableBool** | Specifies an autotest class name presence status to search for | [optional] 
**LastTestResultOutcome** | Pointer to [**NullableAutotestResultOutcome**](AutotestResultOutcome.md) | Specifies an autotest outcome of the last test result to search for | [optional] 
**LastTestResultStatusCodes** | Pointer to **[]string** | Specifies an autotest status codes of the last test result to search for | [optional] 
**ExternalKey** | Pointer to **NullableString** | Specifies an autotest external key to search for | [optional] 
**LastTestResultConfigurationIds** | Pointer to **[]string** | Specifies an autotest configuration IDs of the last test result to search for | [optional] 

## Methods

### NewAutoTestFilterApiModel

`func NewAutoTestFilterApiModel() *AutoTestFilterApiModel`

NewAutoTestFilterApiModel instantiates a new AutoTestFilterApiModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoTestFilterApiModelWithDefaults

`func NewAutoTestFilterApiModelWithDefaults() *AutoTestFilterApiModel`

NewAutoTestFilterApiModelWithDefaults instantiates a new AutoTestFilterApiModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectIds

`func (o *AutoTestFilterApiModel) GetProjectIds() []string`

GetProjectIds returns the ProjectIds field if non-nil, zero value otherwise.

### GetProjectIdsOk

`func (o *AutoTestFilterApiModel) GetProjectIdsOk() (*[]string, bool)`

GetProjectIdsOk returns a tuple with the ProjectIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectIds

`func (o *AutoTestFilterApiModel) SetProjectIds(v []string)`

SetProjectIds sets ProjectIds field to given value.

### HasProjectIds

`func (o *AutoTestFilterApiModel) HasProjectIds() bool`

HasProjectIds returns a boolean if a field has been set.

### SetProjectIdsNil

`func (o *AutoTestFilterApiModel) SetProjectIdsNil(b bool)`

 SetProjectIdsNil sets the value for ProjectIds to be an explicit nil

### UnsetProjectIds
`func (o *AutoTestFilterApiModel) UnsetProjectIds()`

UnsetProjectIds ensures that no value is present for ProjectIds, not even an explicit nil
### GetExternalIds

`func (o *AutoTestFilterApiModel) GetExternalIds() []string`

GetExternalIds returns the ExternalIds field if non-nil, zero value otherwise.

### GetExternalIdsOk

`func (o *AutoTestFilterApiModel) GetExternalIdsOk() (*[]string, bool)`

GetExternalIdsOk returns a tuple with the ExternalIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIds

`func (o *AutoTestFilterApiModel) SetExternalIds(v []string)`

SetExternalIds sets ExternalIds field to given value.

### HasExternalIds

`func (o *AutoTestFilterApiModel) HasExternalIds() bool`

HasExternalIds returns a boolean if a field has been set.

### SetExternalIdsNil

`func (o *AutoTestFilterApiModel) SetExternalIdsNil(b bool)`

 SetExternalIdsNil sets the value for ExternalIds to be an explicit nil

### UnsetExternalIds
`func (o *AutoTestFilterApiModel) UnsetExternalIds()`

UnsetExternalIds ensures that no value is present for ExternalIds, not even an explicit nil
### GetGlobalIds

`func (o *AutoTestFilterApiModel) GetGlobalIds() []int64`

GetGlobalIds returns the GlobalIds field if non-nil, zero value otherwise.

### GetGlobalIdsOk

`func (o *AutoTestFilterApiModel) GetGlobalIdsOk() (*[]int64, bool)`

GetGlobalIdsOk returns a tuple with the GlobalIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalIds

`func (o *AutoTestFilterApiModel) SetGlobalIds(v []int64)`

SetGlobalIds sets GlobalIds field to given value.

### HasGlobalIds

`func (o *AutoTestFilterApiModel) HasGlobalIds() bool`

HasGlobalIds returns a boolean if a field has been set.

### SetGlobalIdsNil

`func (o *AutoTestFilterApiModel) SetGlobalIdsNil(b bool)`

 SetGlobalIdsNil sets the value for GlobalIds to be an explicit nil

### UnsetGlobalIds
`func (o *AutoTestFilterApiModel) UnsetGlobalIds()`

UnsetGlobalIds ensures that no value is present for GlobalIds, not even an explicit nil
### GetName

`func (o *AutoTestFilterApiModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AutoTestFilterApiModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AutoTestFilterApiModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AutoTestFilterApiModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *AutoTestFilterApiModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AutoTestFilterApiModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetIsFlaky

`func (o *AutoTestFilterApiModel) GetIsFlaky() bool`

GetIsFlaky returns the IsFlaky field if non-nil, zero value otherwise.

### GetIsFlakyOk

`func (o *AutoTestFilterApiModel) GetIsFlakyOk() (*bool, bool)`

GetIsFlakyOk returns a tuple with the IsFlaky field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFlaky

`func (o *AutoTestFilterApiModel) SetIsFlaky(v bool)`

SetIsFlaky sets IsFlaky field to given value.

### HasIsFlaky

`func (o *AutoTestFilterApiModel) HasIsFlaky() bool`

HasIsFlaky returns a boolean if a field has been set.

### SetIsFlakyNil

`func (o *AutoTestFilterApiModel) SetIsFlakyNil(b bool)`

 SetIsFlakyNil sets the value for IsFlaky to be an explicit nil

### UnsetIsFlaky
`func (o *AutoTestFilterApiModel) UnsetIsFlaky()`

UnsetIsFlaky ensures that no value is present for IsFlaky, not even an explicit nil
### GetMustBeApproved

`func (o *AutoTestFilterApiModel) GetMustBeApproved() bool`

GetMustBeApproved returns the MustBeApproved field if non-nil, zero value otherwise.

### GetMustBeApprovedOk

`func (o *AutoTestFilterApiModel) GetMustBeApprovedOk() (*bool, bool)`

GetMustBeApprovedOk returns a tuple with the MustBeApproved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMustBeApproved

`func (o *AutoTestFilterApiModel) SetMustBeApproved(v bool)`

SetMustBeApproved sets MustBeApproved field to given value.

### HasMustBeApproved

`func (o *AutoTestFilterApiModel) HasMustBeApproved() bool`

HasMustBeApproved returns a boolean if a field has been set.

### SetMustBeApprovedNil

`func (o *AutoTestFilterApiModel) SetMustBeApprovedNil(b bool)`

 SetMustBeApprovedNil sets the value for MustBeApproved to be an explicit nil

### UnsetMustBeApproved
`func (o *AutoTestFilterApiModel) UnsetMustBeApproved()`

UnsetMustBeApproved ensures that no value is present for MustBeApproved, not even an explicit nil
### GetStabilityPercentage

`func (o *AutoTestFilterApiModel) GetStabilityPercentage() Int64RangeSelectorModel`

GetStabilityPercentage returns the StabilityPercentage field if non-nil, zero value otherwise.

### GetStabilityPercentageOk

`func (o *AutoTestFilterApiModel) GetStabilityPercentageOk() (*Int64RangeSelectorModel, bool)`

GetStabilityPercentageOk returns a tuple with the StabilityPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStabilityPercentage

`func (o *AutoTestFilterApiModel) SetStabilityPercentage(v Int64RangeSelectorModel)`

SetStabilityPercentage sets StabilityPercentage field to given value.

### HasStabilityPercentage

`func (o *AutoTestFilterApiModel) HasStabilityPercentage() bool`

HasStabilityPercentage returns a boolean if a field has been set.

### SetStabilityPercentageNil

`func (o *AutoTestFilterApiModel) SetStabilityPercentageNil(b bool)`

 SetStabilityPercentageNil sets the value for StabilityPercentage to be an explicit nil

### UnsetStabilityPercentage
`func (o *AutoTestFilterApiModel) UnsetStabilityPercentage()`

UnsetStabilityPercentage ensures that no value is present for StabilityPercentage, not even an explicit nil
### GetCreatedDate

`func (o *AutoTestFilterApiModel) GetCreatedDate() DateTimeRangeSelectorModel`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *AutoTestFilterApiModel) GetCreatedDateOk() (*DateTimeRangeSelectorModel, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *AutoTestFilterApiModel) SetCreatedDate(v DateTimeRangeSelectorModel)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *AutoTestFilterApiModel) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### SetCreatedDateNil

`func (o *AutoTestFilterApiModel) SetCreatedDateNil(b bool)`

 SetCreatedDateNil sets the value for CreatedDate to be an explicit nil

### UnsetCreatedDate
`func (o *AutoTestFilterApiModel) UnsetCreatedDate()`

UnsetCreatedDate ensures that no value is present for CreatedDate, not even an explicit nil
### GetCreatedByIds

`func (o *AutoTestFilterApiModel) GetCreatedByIds() []string`

GetCreatedByIds returns the CreatedByIds field if non-nil, zero value otherwise.

### GetCreatedByIdsOk

`func (o *AutoTestFilterApiModel) GetCreatedByIdsOk() (*[]string, bool)`

GetCreatedByIdsOk returns a tuple with the CreatedByIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByIds

`func (o *AutoTestFilterApiModel) SetCreatedByIds(v []string)`

SetCreatedByIds sets CreatedByIds field to given value.

### HasCreatedByIds

`func (o *AutoTestFilterApiModel) HasCreatedByIds() bool`

HasCreatedByIds returns a boolean if a field has been set.

### SetCreatedByIdsNil

`func (o *AutoTestFilterApiModel) SetCreatedByIdsNil(b bool)`

 SetCreatedByIdsNil sets the value for CreatedByIds to be an explicit nil

### UnsetCreatedByIds
`func (o *AutoTestFilterApiModel) UnsetCreatedByIds()`

UnsetCreatedByIds ensures that no value is present for CreatedByIds, not even an explicit nil
### GetModifiedDate

`func (o *AutoTestFilterApiModel) GetModifiedDate() DateTimeRangeSelectorModel`

GetModifiedDate returns the ModifiedDate field if non-nil, zero value otherwise.

### GetModifiedDateOk

`func (o *AutoTestFilterApiModel) GetModifiedDateOk() (*DateTimeRangeSelectorModel, bool)`

GetModifiedDateOk returns a tuple with the ModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedDate

`func (o *AutoTestFilterApiModel) SetModifiedDate(v DateTimeRangeSelectorModel)`

SetModifiedDate sets ModifiedDate field to given value.

### HasModifiedDate

`func (o *AutoTestFilterApiModel) HasModifiedDate() bool`

HasModifiedDate returns a boolean if a field has been set.

### SetModifiedDateNil

`func (o *AutoTestFilterApiModel) SetModifiedDateNil(b bool)`

 SetModifiedDateNil sets the value for ModifiedDate to be an explicit nil

### UnsetModifiedDate
`func (o *AutoTestFilterApiModel) UnsetModifiedDate()`

UnsetModifiedDate ensures that no value is present for ModifiedDate, not even an explicit nil
### GetModifiedByIds

`func (o *AutoTestFilterApiModel) GetModifiedByIds() []string`

GetModifiedByIds returns the ModifiedByIds field if non-nil, zero value otherwise.

### GetModifiedByIdsOk

`func (o *AutoTestFilterApiModel) GetModifiedByIdsOk() (*[]string, bool)`

GetModifiedByIdsOk returns a tuple with the ModifiedByIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedByIds

`func (o *AutoTestFilterApiModel) SetModifiedByIds(v []string)`

SetModifiedByIds sets ModifiedByIds field to given value.

### HasModifiedByIds

`func (o *AutoTestFilterApiModel) HasModifiedByIds() bool`

HasModifiedByIds returns a boolean if a field has been set.

### SetModifiedByIdsNil

`func (o *AutoTestFilterApiModel) SetModifiedByIdsNil(b bool)`

 SetModifiedByIdsNil sets the value for ModifiedByIds to be an explicit nil

### UnsetModifiedByIds
`func (o *AutoTestFilterApiModel) UnsetModifiedByIds()`

UnsetModifiedByIds ensures that no value is present for ModifiedByIds, not even an explicit nil
### GetIsDeleted

`func (o *AutoTestFilterApiModel) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *AutoTestFilterApiModel) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *AutoTestFilterApiModel) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *AutoTestFilterApiModel) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### SetIsDeletedNil

`func (o *AutoTestFilterApiModel) SetIsDeletedNil(b bool)`

 SetIsDeletedNil sets the value for IsDeleted to be an explicit nil

### UnsetIsDeleted
`func (o *AutoTestFilterApiModel) UnsetIsDeleted()`

UnsetIsDeleted ensures that no value is present for IsDeleted, not even an explicit nil
### GetNamespace

`func (o *AutoTestFilterApiModel) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *AutoTestFilterApiModel) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *AutoTestFilterApiModel) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *AutoTestFilterApiModel) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### SetNamespaceNil

`func (o *AutoTestFilterApiModel) SetNamespaceNil(b bool)`

 SetNamespaceNil sets the value for Namespace to be an explicit nil

### UnsetNamespace
`func (o *AutoTestFilterApiModel) UnsetNamespace()`

UnsetNamespace ensures that no value is present for Namespace, not even an explicit nil
### GetIsEmptyNamespace

`func (o *AutoTestFilterApiModel) GetIsEmptyNamespace() bool`

GetIsEmptyNamespace returns the IsEmptyNamespace field if non-nil, zero value otherwise.

### GetIsEmptyNamespaceOk

`func (o *AutoTestFilterApiModel) GetIsEmptyNamespaceOk() (*bool, bool)`

GetIsEmptyNamespaceOk returns a tuple with the IsEmptyNamespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEmptyNamespace

`func (o *AutoTestFilterApiModel) SetIsEmptyNamespace(v bool)`

SetIsEmptyNamespace sets IsEmptyNamespace field to given value.

### HasIsEmptyNamespace

`func (o *AutoTestFilterApiModel) HasIsEmptyNamespace() bool`

HasIsEmptyNamespace returns a boolean if a field has been set.

### SetIsEmptyNamespaceNil

`func (o *AutoTestFilterApiModel) SetIsEmptyNamespaceNil(b bool)`

 SetIsEmptyNamespaceNil sets the value for IsEmptyNamespace to be an explicit nil

### UnsetIsEmptyNamespace
`func (o *AutoTestFilterApiModel) UnsetIsEmptyNamespace()`

UnsetIsEmptyNamespace ensures that no value is present for IsEmptyNamespace, not even an explicit nil
### GetClassName

`func (o *AutoTestFilterApiModel) GetClassName() string`

GetClassName returns the ClassName field if non-nil, zero value otherwise.

### GetClassNameOk

`func (o *AutoTestFilterApiModel) GetClassNameOk() (*string, bool)`

GetClassNameOk returns a tuple with the ClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassName

`func (o *AutoTestFilterApiModel) SetClassName(v string)`

SetClassName sets ClassName field to given value.

### HasClassName

`func (o *AutoTestFilterApiModel) HasClassName() bool`

HasClassName returns a boolean if a field has been set.

### SetClassNameNil

`func (o *AutoTestFilterApiModel) SetClassNameNil(b bool)`

 SetClassNameNil sets the value for ClassName to be an explicit nil

### UnsetClassName
`func (o *AutoTestFilterApiModel) UnsetClassName()`

UnsetClassName ensures that no value is present for ClassName, not even an explicit nil
### GetIsEmptyClassName

`func (o *AutoTestFilterApiModel) GetIsEmptyClassName() bool`

GetIsEmptyClassName returns the IsEmptyClassName field if non-nil, zero value otherwise.

### GetIsEmptyClassNameOk

`func (o *AutoTestFilterApiModel) GetIsEmptyClassNameOk() (*bool, bool)`

GetIsEmptyClassNameOk returns a tuple with the IsEmptyClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEmptyClassName

`func (o *AutoTestFilterApiModel) SetIsEmptyClassName(v bool)`

SetIsEmptyClassName sets IsEmptyClassName field to given value.

### HasIsEmptyClassName

`func (o *AutoTestFilterApiModel) HasIsEmptyClassName() bool`

HasIsEmptyClassName returns a boolean if a field has been set.

### SetIsEmptyClassNameNil

`func (o *AutoTestFilterApiModel) SetIsEmptyClassNameNil(b bool)`

 SetIsEmptyClassNameNil sets the value for IsEmptyClassName to be an explicit nil

### UnsetIsEmptyClassName
`func (o *AutoTestFilterApiModel) UnsetIsEmptyClassName()`

UnsetIsEmptyClassName ensures that no value is present for IsEmptyClassName, not even an explicit nil
### GetLastTestResultOutcome

`func (o *AutoTestFilterApiModel) GetLastTestResultOutcome() AutotestResultOutcome`

GetLastTestResultOutcome returns the LastTestResultOutcome field if non-nil, zero value otherwise.

### GetLastTestResultOutcomeOk

`func (o *AutoTestFilterApiModel) GetLastTestResultOutcomeOk() (*AutotestResultOutcome, bool)`

GetLastTestResultOutcomeOk returns a tuple with the LastTestResultOutcome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTestResultOutcome

`func (o *AutoTestFilterApiModel) SetLastTestResultOutcome(v AutotestResultOutcome)`

SetLastTestResultOutcome sets LastTestResultOutcome field to given value.

### HasLastTestResultOutcome

`func (o *AutoTestFilterApiModel) HasLastTestResultOutcome() bool`

HasLastTestResultOutcome returns a boolean if a field has been set.

### SetLastTestResultOutcomeNil

`func (o *AutoTestFilterApiModel) SetLastTestResultOutcomeNil(b bool)`

 SetLastTestResultOutcomeNil sets the value for LastTestResultOutcome to be an explicit nil

### UnsetLastTestResultOutcome
`func (o *AutoTestFilterApiModel) UnsetLastTestResultOutcome()`

UnsetLastTestResultOutcome ensures that no value is present for LastTestResultOutcome, not even an explicit nil
### GetLastTestResultStatusCodes

`func (o *AutoTestFilterApiModel) GetLastTestResultStatusCodes() []string`

GetLastTestResultStatusCodes returns the LastTestResultStatusCodes field if non-nil, zero value otherwise.

### GetLastTestResultStatusCodesOk

`func (o *AutoTestFilterApiModel) GetLastTestResultStatusCodesOk() (*[]string, bool)`

GetLastTestResultStatusCodesOk returns a tuple with the LastTestResultStatusCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTestResultStatusCodes

`func (o *AutoTestFilterApiModel) SetLastTestResultStatusCodes(v []string)`

SetLastTestResultStatusCodes sets LastTestResultStatusCodes field to given value.

### HasLastTestResultStatusCodes

`func (o *AutoTestFilterApiModel) HasLastTestResultStatusCodes() bool`

HasLastTestResultStatusCodes returns a boolean if a field has been set.

### SetLastTestResultStatusCodesNil

`func (o *AutoTestFilterApiModel) SetLastTestResultStatusCodesNil(b bool)`

 SetLastTestResultStatusCodesNil sets the value for LastTestResultStatusCodes to be an explicit nil

### UnsetLastTestResultStatusCodes
`func (o *AutoTestFilterApiModel) UnsetLastTestResultStatusCodes()`

UnsetLastTestResultStatusCodes ensures that no value is present for LastTestResultStatusCodes, not even an explicit nil
### GetExternalKey

`func (o *AutoTestFilterApiModel) GetExternalKey() string`

GetExternalKey returns the ExternalKey field if non-nil, zero value otherwise.

### GetExternalKeyOk

`func (o *AutoTestFilterApiModel) GetExternalKeyOk() (*string, bool)`

GetExternalKeyOk returns a tuple with the ExternalKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalKey

`func (o *AutoTestFilterApiModel) SetExternalKey(v string)`

SetExternalKey sets ExternalKey field to given value.

### HasExternalKey

`func (o *AutoTestFilterApiModel) HasExternalKey() bool`

HasExternalKey returns a boolean if a field has been set.

### SetExternalKeyNil

`func (o *AutoTestFilterApiModel) SetExternalKeyNil(b bool)`

 SetExternalKeyNil sets the value for ExternalKey to be an explicit nil

### UnsetExternalKey
`func (o *AutoTestFilterApiModel) UnsetExternalKey()`

UnsetExternalKey ensures that no value is present for ExternalKey, not even an explicit nil
### GetLastTestResultConfigurationIds

`func (o *AutoTestFilterApiModel) GetLastTestResultConfigurationIds() []string`

GetLastTestResultConfigurationIds returns the LastTestResultConfigurationIds field if non-nil, zero value otherwise.

### GetLastTestResultConfigurationIdsOk

`func (o *AutoTestFilterApiModel) GetLastTestResultConfigurationIdsOk() (*[]string, bool)`

GetLastTestResultConfigurationIdsOk returns a tuple with the LastTestResultConfigurationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTestResultConfigurationIds

`func (o *AutoTestFilterApiModel) SetLastTestResultConfigurationIds(v []string)`

SetLastTestResultConfigurationIds sets LastTestResultConfigurationIds field to given value.

### HasLastTestResultConfigurationIds

`func (o *AutoTestFilterApiModel) HasLastTestResultConfigurationIds() bool`

HasLastTestResultConfigurationIds returns a boolean if a field has been set.

### SetLastTestResultConfigurationIdsNil

`func (o *AutoTestFilterApiModel) SetLastTestResultConfigurationIdsNil(b bool)`

 SetLastTestResultConfigurationIdsNil sets the value for LastTestResultConfigurationIds to be an explicit nil

### UnsetLastTestResultConfigurationIds
`func (o *AutoTestFilterApiModel) UnsetLastTestResultConfigurationIds()`

UnsetLastTestResultConfigurationIds ensures that no value is present for LastTestResultConfigurationIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


