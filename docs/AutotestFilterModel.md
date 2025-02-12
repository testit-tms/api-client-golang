# AutoTestFilterModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectIds** | Pointer to **[]string** |  | [optional] 
**ExternalIds** | Pointer to **[]string** |  | [optional] 
**GlobalIds** | Pointer to **[]int64** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**IsFlaky** | Pointer to **NullableBool** |  | [optional] 
**MustBeApproved** | Pointer to **NullableBool** |  | [optional] 
**StabilityPercentage** | Pointer to [**NullableInt64RangeSelectorModel**](Int64RangeSelectorModel.md) |  | [optional] 
**CreatedDate** | Pointer to [**NullableDateTimeRangeSelectorModel**](DateTimeRangeSelectorModel.md) |  | [optional] 
**CreatedByIds** | Pointer to **[]string** |  | [optional] 
**ModifiedDate** | Pointer to [**NullableDateTimeRangeSelectorModel**](DateTimeRangeSelectorModel.md) |  | [optional] 
**ModifiedByIds** | Pointer to **[]string** |  | [optional] 
**IsDeleted** | Pointer to **NullableBool** |  | [optional] 
**Namespace** | Pointer to **NullableString** |  | [optional] 
**IsEmptyNamespace** | Pointer to **NullableBool** |  | [optional] 
**ClassName** | Pointer to **NullableString** |  | [optional] 
**IsEmptyClassName** | Pointer to **NullableBool** |  | [optional] 
**LastTestResultOutcome** | Pointer to [**NullableAutotestResultOutcome**](AutotestResultOutcome.md) |  | [optional] 
**LastTestResultStatusCode** | Pointer to **NullableString** |  | [optional] 
**ExternalKey** | Pointer to **NullableString** |  | [optional] 
**LastTestResultConfigurationIds** | Pointer to **[]string** |  | [optional] 

## Methods

### NewAutoTestFilterModel

`func NewAutoTestFilterModel() *AutoTestFilterModel`

NewAutoTestFilterModel instantiates a new AutoTestFilterModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoTestFilterModelWithDefaults

`func NewAutoTestFilterModelWithDefaults() *AutoTestFilterModel`

NewAutoTestFilterModelWithDefaults instantiates a new AutoTestFilterModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectIds

`func (o *AutoTestFilterModel) GetProjectIds() []string`

GetProjectIds returns the ProjectIds field if non-nil, zero value otherwise.

### GetProjectIdsOk

`func (o *AutoTestFilterModel) GetProjectIdsOk() (*[]string, bool)`

GetProjectIdsOk returns a tuple with the ProjectIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectIds

`func (o *AutoTestFilterModel) SetProjectIds(v []string)`

SetProjectIds sets ProjectIds field to given value.

### HasProjectIds

`func (o *AutoTestFilterModel) HasProjectIds() bool`

HasProjectIds returns a boolean if a field has been set.

### SetProjectIdsNil

`func (o *AutoTestFilterModel) SetProjectIdsNil(b bool)`

 SetProjectIdsNil sets the value for ProjectIds to be an explicit nil

### UnsetProjectIds
`func (o *AutoTestFilterModel) UnsetProjectIds()`

UnsetProjectIds ensures that no value is present for ProjectIds, not even an explicit nil
### GetExternalIds

`func (o *AutoTestFilterModel) GetExternalIds() []string`

GetExternalIds returns the ExternalIds field if non-nil, zero value otherwise.

### GetExternalIdsOk

`func (o *AutoTestFilterModel) GetExternalIdsOk() (*[]string, bool)`

GetExternalIdsOk returns a tuple with the ExternalIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIds

`func (o *AutoTestFilterModel) SetExternalIds(v []string)`

SetExternalIds sets ExternalIds field to given value.

### HasExternalIds

`func (o *AutoTestFilterModel) HasExternalIds() bool`

HasExternalIds returns a boolean if a field has been set.

### SetExternalIdsNil

`func (o *AutoTestFilterModel) SetExternalIdsNil(b bool)`

 SetExternalIdsNil sets the value for ExternalIds to be an explicit nil

### UnsetExternalIds
`func (o *AutoTestFilterModel) UnsetExternalIds()`

UnsetExternalIds ensures that no value is present for ExternalIds, not even an explicit nil
### GetGlobalIds

`func (o *AutoTestFilterModel) GetGlobalIds() []int64`

GetGlobalIds returns the GlobalIds field if non-nil, zero value otherwise.

### GetGlobalIdsOk

`func (o *AutoTestFilterModel) GetGlobalIdsOk() (*[]int64, bool)`

GetGlobalIdsOk returns a tuple with the GlobalIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalIds

`func (o *AutoTestFilterModel) SetGlobalIds(v []int64)`

SetGlobalIds sets GlobalIds field to given value.

### HasGlobalIds

`func (o *AutoTestFilterModel) HasGlobalIds() bool`

HasGlobalIds returns a boolean if a field has been set.

### SetGlobalIdsNil

`func (o *AutoTestFilterModel) SetGlobalIdsNil(b bool)`

 SetGlobalIdsNil sets the value for GlobalIds to be an explicit nil

### UnsetGlobalIds
`func (o *AutoTestFilterModel) UnsetGlobalIds()`

UnsetGlobalIds ensures that no value is present for GlobalIds, not even an explicit nil
### GetName

`func (o *AutoTestFilterModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AutoTestFilterModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AutoTestFilterModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AutoTestFilterModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *AutoTestFilterModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AutoTestFilterModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetIsFlaky

`func (o *AutoTestFilterModel) GetIsFlaky() bool`

GetIsFlaky returns the IsFlaky field if non-nil, zero value otherwise.

### GetIsFlakyOk

`func (o *AutoTestFilterModel) GetIsFlakyOk() (*bool, bool)`

GetIsFlakyOk returns a tuple with the IsFlaky field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFlaky

`func (o *AutoTestFilterModel) SetIsFlaky(v bool)`

SetIsFlaky sets IsFlaky field to given value.

### HasIsFlaky

`func (o *AutoTestFilterModel) HasIsFlaky() bool`

HasIsFlaky returns a boolean if a field has been set.

### SetIsFlakyNil

`func (o *AutoTestFilterModel) SetIsFlakyNil(b bool)`

 SetIsFlakyNil sets the value for IsFlaky to be an explicit nil

### UnsetIsFlaky
`func (o *AutoTestFilterModel) UnsetIsFlaky()`

UnsetIsFlaky ensures that no value is present for IsFlaky, not even an explicit nil
### GetMustBeApproved

`func (o *AutoTestFilterModel) GetMustBeApproved() bool`

GetMustBeApproved returns the MustBeApproved field if non-nil, zero value otherwise.

### GetMustBeApprovedOk

`func (o *AutoTestFilterModel) GetMustBeApprovedOk() (*bool, bool)`

GetMustBeApprovedOk returns a tuple with the MustBeApproved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMustBeApproved

`func (o *AutoTestFilterModel) SetMustBeApproved(v bool)`

SetMustBeApproved sets MustBeApproved field to given value.

### HasMustBeApproved

`func (o *AutoTestFilterModel) HasMustBeApproved() bool`

HasMustBeApproved returns a boolean if a field has been set.

### SetMustBeApprovedNil

`func (o *AutoTestFilterModel) SetMustBeApprovedNil(b bool)`

 SetMustBeApprovedNil sets the value for MustBeApproved to be an explicit nil

### UnsetMustBeApproved
`func (o *AutoTestFilterModel) UnsetMustBeApproved()`

UnsetMustBeApproved ensures that no value is present for MustBeApproved, not even an explicit nil
### GetStabilityPercentage

`func (o *AutoTestFilterModel) GetStabilityPercentage() Int64RangeSelectorModel`

GetStabilityPercentage returns the StabilityPercentage field if non-nil, zero value otherwise.

### GetStabilityPercentageOk

`func (o *AutoTestFilterModel) GetStabilityPercentageOk() (*Int64RangeSelectorModel, bool)`

GetStabilityPercentageOk returns a tuple with the StabilityPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStabilityPercentage

`func (o *AutoTestFilterModel) SetStabilityPercentage(v Int64RangeSelectorModel)`

SetStabilityPercentage sets StabilityPercentage field to given value.

### HasStabilityPercentage

`func (o *AutoTestFilterModel) HasStabilityPercentage() bool`

HasStabilityPercentage returns a boolean if a field has been set.

### SetStabilityPercentageNil

`func (o *AutoTestFilterModel) SetStabilityPercentageNil(b bool)`

 SetStabilityPercentageNil sets the value for StabilityPercentage to be an explicit nil

### UnsetStabilityPercentage
`func (o *AutoTestFilterModel) UnsetStabilityPercentage()`

UnsetStabilityPercentage ensures that no value is present for StabilityPercentage, not even an explicit nil
### GetCreatedDate

`func (o *AutoTestFilterModel) GetCreatedDate() DateTimeRangeSelectorModel`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *AutoTestFilterModel) GetCreatedDateOk() (*DateTimeRangeSelectorModel, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *AutoTestFilterModel) SetCreatedDate(v DateTimeRangeSelectorModel)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *AutoTestFilterModel) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### SetCreatedDateNil

`func (o *AutoTestFilterModel) SetCreatedDateNil(b bool)`

 SetCreatedDateNil sets the value for CreatedDate to be an explicit nil

### UnsetCreatedDate
`func (o *AutoTestFilterModel) UnsetCreatedDate()`

UnsetCreatedDate ensures that no value is present for CreatedDate, not even an explicit nil
### GetCreatedByIds

`func (o *AutoTestFilterModel) GetCreatedByIds() []string`

GetCreatedByIds returns the CreatedByIds field if non-nil, zero value otherwise.

### GetCreatedByIdsOk

`func (o *AutoTestFilterModel) GetCreatedByIdsOk() (*[]string, bool)`

GetCreatedByIdsOk returns a tuple with the CreatedByIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByIds

`func (o *AutoTestFilterModel) SetCreatedByIds(v []string)`

SetCreatedByIds sets CreatedByIds field to given value.

### HasCreatedByIds

`func (o *AutoTestFilterModel) HasCreatedByIds() bool`

HasCreatedByIds returns a boolean if a field has been set.

### SetCreatedByIdsNil

`func (o *AutoTestFilterModel) SetCreatedByIdsNil(b bool)`

 SetCreatedByIdsNil sets the value for CreatedByIds to be an explicit nil

### UnsetCreatedByIds
`func (o *AutoTestFilterModel) UnsetCreatedByIds()`

UnsetCreatedByIds ensures that no value is present for CreatedByIds, not even an explicit nil
### GetModifiedDate

`func (o *AutoTestFilterModel) GetModifiedDate() DateTimeRangeSelectorModel`

GetModifiedDate returns the ModifiedDate field if non-nil, zero value otherwise.

### GetModifiedDateOk

`func (o *AutoTestFilterModel) GetModifiedDateOk() (*DateTimeRangeSelectorModel, bool)`

GetModifiedDateOk returns a tuple with the ModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedDate

`func (o *AutoTestFilterModel) SetModifiedDate(v DateTimeRangeSelectorModel)`

SetModifiedDate sets ModifiedDate field to given value.

### HasModifiedDate

`func (o *AutoTestFilterModel) HasModifiedDate() bool`

HasModifiedDate returns a boolean if a field has been set.

### SetModifiedDateNil

`func (o *AutoTestFilterModel) SetModifiedDateNil(b bool)`

 SetModifiedDateNil sets the value for ModifiedDate to be an explicit nil

### UnsetModifiedDate
`func (o *AutoTestFilterModel) UnsetModifiedDate()`

UnsetModifiedDate ensures that no value is present for ModifiedDate, not even an explicit nil
### GetModifiedByIds

`func (o *AutoTestFilterModel) GetModifiedByIds() []string`

GetModifiedByIds returns the ModifiedByIds field if non-nil, zero value otherwise.

### GetModifiedByIdsOk

`func (o *AutoTestFilterModel) GetModifiedByIdsOk() (*[]string, bool)`

GetModifiedByIdsOk returns a tuple with the ModifiedByIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedByIds

`func (o *AutoTestFilterModel) SetModifiedByIds(v []string)`

SetModifiedByIds sets ModifiedByIds field to given value.

### HasModifiedByIds

`func (o *AutoTestFilterModel) HasModifiedByIds() bool`

HasModifiedByIds returns a boolean if a field has been set.

### SetModifiedByIdsNil

`func (o *AutoTestFilterModel) SetModifiedByIdsNil(b bool)`

 SetModifiedByIdsNil sets the value for ModifiedByIds to be an explicit nil

### UnsetModifiedByIds
`func (o *AutoTestFilterModel) UnsetModifiedByIds()`

UnsetModifiedByIds ensures that no value is present for ModifiedByIds, not even an explicit nil
### GetIsDeleted

`func (o *AutoTestFilterModel) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *AutoTestFilterModel) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *AutoTestFilterModel) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *AutoTestFilterModel) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### SetIsDeletedNil

`func (o *AutoTestFilterModel) SetIsDeletedNil(b bool)`

 SetIsDeletedNil sets the value for IsDeleted to be an explicit nil

### UnsetIsDeleted
`func (o *AutoTestFilterModel) UnsetIsDeleted()`

UnsetIsDeleted ensures that no value is present for IsDeleted, not even an explicit nil
### GetNamespace

`func (o *AutoTestFilterModel) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *AutoTestFilterModel) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *AutoTestFilterModel) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *AutoTestFilterModel) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### SetNamespaceNil

`func (o *AutoTestFilterModel) SetNamespaceNil(b bool)`

 SetNamespaceNil sets the value for Namespace to be an explicit nil

### UnsetNamespace
`func (o *AutoTestFilterModel) UnsetNamespace()`

UnsetNamespace ensures that no value is present for Namespace, not even an explicit nil
### GetIsEmptyNamespace

`func (o *AutoTestFilterModel) GetIsEmptyNamespace() bool`

GetIsEmptyNamespace returns the IsEmptyNamespace field if non-nil, zero value otherwise.

### GetIsEmptyNamespaceOk

`func (o *AutoTestFilterModel) GetIsEmptyNamespaceOk() (*bool, bool)`

GetIsEmptyNamespaceOk returns a tuple with the IsEmptyNamespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEmptyNamespace

`func (o *AutoTestFilterModel) SetIsEmptyNamespace(v bool)`

SetIsEmptyNamespace sets IsEmptyNamespace field to given value.

### HasIsEmptyNamespace

`func (o *AutoTestFilterModel) HasIsEmptyNamespace() bool`

HasIsEmptyNamespace returns a boolean if a field has been set.

### SetIsEmptyNamespaceNil

`func (o *AutoTestFilterModel) SetIsEmptyNamespaceNil(b bool)`

 SetIsEmptyNamespaceNil sets the value for IsEmptyNamespace to be an explicit nil

### UnsetIsEmptyNamespace
`func (o *AutoTestFilterModel) UnsetIsEmptyNamespace()`

UnsetIsEmptyNamespace ensures that no value is present for IsEmptyNamespace, not even an explicit nil
### GetClassName

`func (o *AutoTestFilterModel) GetClassName() string`

GetClassName returns the ClassName field if non-nil, zero value otherwise.

### GetClassNameOk

`func (o *AutoTestFilterModel) GetClassNameOk() (*string, bool)`

GetClassNameOk returns a tuple with the ClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassName

`func (o *AutoTestFilterModel) SetClassName(v string)`

SetClassName sets ClassName field to given value.

### HasClassName

`func (o *AutoTestFilterModel) HasClassName() bool`

HasClassName returns a boolean if a field has been set.

### SetClassNameNil

`func (o *AutoTestFilterModel) SetClassNameNil(b bool)`

 SetClassNameNil sets the value for ClassName to be an explicit nil

### UnsetClassName
`func (o *AutoTestFilterModel) UnsetClassName()`

UnsetClassName ensures that no value is present for ClassName, not even an explicit nil
### GetIsEmptyClassName

`func (o *AutoTestFilterModel) GetIsEmptyClassName() bool`

GetIsEmptyClassName returns the IsEmptyClassName field if non-nil, zero value otherwise.

### GetIsEmptyClassNameOk

`func (o *AutoTestFilterModel) GetIsEmptyClassNameOk() (*bool, bool)`

GetIsEmptyClassNameOk returns a tuple with the IsEmptyClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEmptyClassName

`func (o *AutoTestFilterModel) SetIsEmptyClassName(v bool)`

SetIsEmptyClassName sets IsEmptyClassName field to given value.

### HasIsEmptyClassName

`func (o *AutoTestFilterModel) HasIsEmptyClassName() bool`

HasIsEmptyClassName returns a boolean if a field has been set.

### SetIsEmptyClassNameNil

`func (o *AutoTestFilterModel) SetIsEmptyClassNameNil(b bool)`

 SetIsEmptyClassNameNil sets the value for IsEmptyClassName to be an explicit nil

### UnsetIsEmptyClassName
`func (o *AutoTestFilterModel) UnsetIsEmptyClassName()`

UnsetIsEmptyClassName ensures that no value is present for IsEmptyClassName, not even an explicit nil
### GetLastTestResultOutcome

`func (o *AutoTestFilterModel) GetLastTestResultOutcome() AutotestResultOutcome`

GetLastTestResultOutcome returns the LastTestResultOutcome field if non-nil, zero value otherwise.

### GetLastTestResultOutcomeOk

`func (o *AutoTestFilterModel) GetLastTestResultOutcomeOk() (*AutotestResultOutcome, bool)`

GetLastTestResultOutcomeOk returns a tuple with the LastTestResultOutcome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTestResultOutcome

`func (o *AutoTestFilterModel) SetLastTestResultOutcome(v AutotestResultOutcome)`

SetLastTestResultOutcome sets LastTestResultOutcome field to given value.

### HasLastTestResultOutcome

`func (o *AutoTestFilterModel) HasLastTestResultOutcome() bool`

HasLastTestResultOutcome returns a boolean if a field has been set.

### SetLastTestResultOutcomeNil

`func (o *AutoTestFilterModel) SetLastTestResultOutcomeNil(b bool)`

 SetLastTestResultOutcomeNil sets the value for LastTestResultOutcome to be an explicit nil

### UnsetLastTestResultOutcome
`func (o *AutoTestFilterModel) UnsetLastTestResultOutcome()`

UnsetLastTestResultOutcome ensures that no value is present for LastTestResultOutcome, not even an explicit nil
### GetLastTestResultStatusCode

`func (o *AutoTestFilterModel) GetLastTestResultStatusCode() string`

GetLastTestResultStatusCode returns the LastTestResultStatusCode field if non-nil, zero value otherwise.

### GetLastTestResultStatusCodeOk

`func (o *AutoTestFilterModel) GetLastTestResultStatusCodeOk() (*string, bool)`

GetLastTestResultStatusCodeOk returns a tuple with the LastTestResultStatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTestResultStatusCode

`func (o *AutoTestFilterModel) SetLastTestResultStatusCode(v string)`

SetLastTestResultStatusCode sets LastTestResultStatusCode field to given value.

### HasLastTestResultStatusCode

`func (o *AutoTestFilterModel) HasLastTestResultStatusCode() bool`

HasLastTestResultStatusCode returns a boolean if a field has been set.

### SetLastTestResultStatusCodeNil

`func (o *AutoTestFilterModel) SetLastTestResultStatusCodeNil(b bool)`

 SetLastTestResultStatusCodeNil sets the value for LastTestResultStatusCode to be an explicit nil

### UnsetLastTestResultStatusCode
`func (o *AutoTestFilterModel) UnsetLastTestResultStatusCode()`

UnsetLastTestResultStatusCode ensures that no value is present for LastTestResultStatusCode, not even an explicit nil
### GetExternalKey

`func (o *AutoTestFilterModel) GetExternalKey() string`

GetExternalKey returns the ExternalKey field if non-nil, zero value otherwise.

### GetExternalKeyOk

`func (o *AutoTestFilterModel) GetExternalKeyOk() (*string, bool)`

GetExternalKeyOk returns a tuple with the ExternalKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalKey

`func (o *AutoTestFilterModel) SetExternalKey(v string)`

SetExternalKey sets ExternalKey field to given value.

### HasExternalKey

`func (o *AutoTestFilterModel) HasExternalKey() bool`

HasExternalKey returns a boolean if a field has been set.

### SetExternalKeyNil

`func (o *AutoTestFilterModel) SetExternalKeyNil(b bool)`

 SetExternalKeyNil sets the value for ExternalKey to be an explicit nil

### UnsetExternalKey
`func (o *AutoTestFilterModel) UnsetExternalKey()`

UnsetExternalKey ensures that no value is present for ExternalKey, not even an explicit nil
### GetLastTestResultConfigurationIds

`func (o *AutoTestFilterModel) GetLastTestResultConfigurationIds() []string`

GetLastTestResultConfigurationIds returns the LastTestResultConfigurationIds field if non-nil, zero value otherwise.

### GetLastTestResultConfigurationIdsOk

`func (o *AutoTestFilterModel) GetLastTestResultConfigurationIdsOk() (*[]string, bool)`

GetLastTestResultConfigurationIdsOk returns a tuple with the LastTestResultConfigurationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTestResultConfigurationIds

`func (o *AutoTestFilterModel) SetLastTestResultConfigurationIds(v []string)`

SetLastTestResultConfigurationIds sets LastTestResultConfigurationIds field to given value.

### HasLastTestResultConfigurationIds

`func (o *AutoTestFilterModel) HasLastTestResultConfigurationIds() bool`

HasLastTestResultConfigurationIds returns a boolean if a field has been set.

### SetLastTestResultConfigurationIdsNil

`func (o *AutoTestFilterModel) SetLastTestResultConfigurationIdsNil(b bool)`

 SetLastTestResultConfigurationIdsNil sets the value for LastTestResultConfigurationIds to be an explicit nil

### UnsetLastTestResultConfigurationIds
`func (o *AutoTestFilterModel) UnsetLastTestResultConfigurationIds()`

UnsetLastTestResultConfigurationIds ensures that no value is present for LastTestResultConfigurationIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


