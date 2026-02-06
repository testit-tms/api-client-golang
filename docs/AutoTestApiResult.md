# AutoTestApiResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**ProjectId** | **string** |  | 
**ExternalId** | Pointer to **NullableString** |  | [optional] 
**Name** | **string** |  | 
**Namespace** | Pointer to **NullableString** |  | [optional] 
**Classname** | Pointer to **NullableString** |  | [optional] 
**Steps** | Pointer to [**[]AutoTestStepApiResult**](AutoTestStepApiResult.md) |  | [optional] 
**Setup** | Pointer to [**[]AutoTestStepApiResult**](AutoTestStepApiResult.md) |  | [optional] 
**Teardown** | Pointer to [**[]AutoTestStepApiResult**](AutoTestStepApiResult.md) |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**IsFlaky** | **bool** |  | 
**ExternalKey** | Pointer to **NullableString** |  | [optional] 
**GlobalId** | **int64** |  | 
**IsDeleted** | **bool** |  | 
**MustBeApproved** | **bool** |  | 
**CreatedDate** | **time.Time** |  | 
**ModifiedDate** | Pointer to **NullableTime** |  | [optional] 
**CreatedById** | **string** |  | 
**ModifiedById** | Pointer to **NullableString** |  | [optional] 
**LastTestRunId** | Pointer to **NullableString** |  | [optional] 
**LastTestRunName** | Pointer to **NullableString** |  | [optional] 
**LastTestResultId** | Pointer to **NullableString** |  | [optional] 
**LastTestResultConfiguration** | Pointer to [**NullableConfigurationShortApiResult**](ConfigurationShortApiResult.md) |  | [optional] 
**LastTestResultOutcome** | Pointer to **NullableString** |  | [optional] 
**LastTestResultStatus** | Pointer to [**NullableTestStatusApiResult**](TestStatusApiResult.md) |  | [optional] 
**StabilityPercentage** | Pointer to **NullableInt64** |  | [optional] 
**Links** | Pointer to [**[]LinkApiResult**](LinkApiResult.md) |  | [optional] 
**Labels** | Pointer to [**[]LabelApiResult**](LabelApiResult.md) |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 

## Methods

### NewAutoTestApiResult

`func NewAutoTestApiResult(id string, projectId string, name string, isFlaky bool, globalId int64, isDeleted bool, mustBeApproved bool, createdDate time.Time, createdById string, ) *AutoTestApiResult`

NewAutoTestApiResult instantiates a new AutoTestApiResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoTestApiResultWithDefaults

`func NewAutoTestApiResultWithDefaults() *AutoTestApiResult`

NewAutoTestApiResultWithDefaults instantiates a new AutoTestApiResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AutoTestApiResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AutoTestApiResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AutoTestApiResult) SetId(v string)`

SetId sets Id field to given value.


### GetProjectId

`func (o *AutoTestApiResult) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *AutoTestApiResult) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *AutoTestApiResult) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetExternalId

`func (o *AutoTestApiResult) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *AutoTestApiResult) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *AutoTestApiResult) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *AutoTestApiResult) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *AutoTestApiResult) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *AutoTestApiResult) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetName

`func (o *AutoTestApiResult) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AutoTestApiResult) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AutoTestApiResult) SetName(v string)`

SetName sets Name field to given value.


### GetNamespace

`func (o *AutoTestApiResult) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *AutoTestApiResult) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *AutoTestApiResult) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *AutoTestApiResult) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### SetNamespaceNil

`func (o *AutoTestApiResult) SetNamespaceNil(b bool)`

 SetNamespaceNil sets the value for Namespace to be an explicit nil

### UnsetNamespace
`func (o *AutoTestApiResult) UnsetNamespace()`

UnsetNamespace ensures that no value is present for Namespace, not even an explicit nil
### GetClassname

`func (o *AutoTestApiResult) GetClassname() string`

GetClassname returns the Classname field if non-nil, zero value otherwise.

### GetClassnameOk

`func (o *AutoTestApiResult) GetClassnameOk() (*string, bool)`

GetClassnameOk returns a tuple with the Classname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassname

`func (o *AutoTestApiResult) SetClassname(v string)`

SetClassname sets Classname field to given value.

### HasClassname

`func (o *AutoTestApiResult) HasClassname() bool`

HasClassname returns a boolean if a field has been set.

### SetClassnameNil

`func (o *AutoTestApiResult) SetClassnameNil(b bool)`

 SetClassnameNil sets the value for Classname to be an explicit nil

### UnsetClassname
`func (o *AutoTestApiResult) UnsetClassname()`

UnsetClassname ensures that no value is present for Classname, not even an explicit nil
### GetSteps

`func (o *AutoTestApiResult) GetSteps() []AutoTestStepApiResult`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *AutoTestApiResult) GetStepsOk() (*[]AutoTestStepApiResult, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *AutoTestApiResult) SetSteps(v []AutoTestStepApiResult)`

SetSteps sets Steps field to given value.

### HasSteps

`func (o *AutoTestApiResult) HasSteps() bool`

HasSteps returns a boolean if a field has been set.

### SetStepsNil

`func (o *AutoTestApiResult) SetStepsNil(b bool)`

 SetStepsNil sets the value for Steps to be an explicit nil

### UnsetSteps
`func (o *AutoTestApiResult) UnsetSteps()`

UnsetSteps ensures that no value is present for Steps, not even an explicit nil
### GetSetup

`func (o *AutoTestApiResult) GetSetup() []AutoTestStepApiResult`

GetSetup returns the Setup field if non-nil, zero value otherwise.

### GetSetupOk

`func (o *AutoTestApiResult) GetSetupOk() (*[]AutoTestStepApiResult, bool)`

GetSetupOk returns a tuple with the Setup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetup

`func (o *AutoTestApiResult) SetSetup(v []AutoTestStepApiResult)`

SetSetup sets Setup field to given value.

### HasSetup

`func (o *AutoTestApiResult) HasSetup() bool`

HasSetup returns a boolean if a field has been set.

### SetSetupNil

`func (o *AutoTestApiResult) SetSetupNil(b bool)`

 SetSetupNil sets the value for Setup to be an explicit nil

### UnsetSetup
`func (o *AutoTestApiResult) UnsetSetup()`

UnsetSetup ensures that no value is present for Setup, not even an explicit nil
### GetTeardown

`func (o *AutoTestApiResult) GetTeardown() []AutoTestStepApiResult`

GetTeardown returns the Teardown field if non-nil, zero value otherwise.

### GetTeardownOk

`func (o *AutoTestApiResult) GetTeardownOk() (*[]AutoTestStepApiResult, bool)`

GetTeardownOk returns a tuple with the Teardown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeardown

`func (o *AutoTestApiResult) SetTeardown(v []AutoTestStepApiResult)`

SetTeardown sets Teardown field to given value.

### HasTeardown

`func (o *AutoTestApiResult) HasTeardown() bool`

HasTeardown returns a boolean if a field has been set.

### SetTeardownNil

`func (o *AutoTestApiResult) SetTeardownNil(b bool)`

 SetTeardownNil sets the value for Teardown to be an explicit nil

### UnsetTeardown
`func (o *AutoTestApiResult) UnsetTeardown()`

UnsetTeardown ensures that no value is present for Teardown, not even an explicit nil
### GetTitle

`func (o *AutoTestApiResult) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AutoTestApiResult) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AutoTestApiResult) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AutoTestApiResult) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *AutoTestApiResult) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *AutoTestApiResult) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetDescription

`func (o *AutoTestApiResult) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AutoTestApiResult) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AutoTestApiResult) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AutoTestApiResult) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *AutoTestApiResult) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *AutoTestApiResult) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetIsFlaky

`func (o *AutoTestApiResult) GetIsFlaky() bool`

GetIsFlaky returns the IsFlaky field if non-nil, zero value otherwise.

### GetIsFlakyOk

`func (o *AutoTestApiResult) GetIsFlakyOk() (*bool, bool)`

GetIsFlakyOk returns a tuple with the IsFlaky field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFlaky

`func (o *AutoTestApiResult) SetIsFlaky(v bool)`

SetIsFlaky sets IsFlaky field to given value.


### GetExternalKey

`func (o *AutoTestApiResult) GetExternalKey() string`

GetExternalKey returns the ExternalKey field if non-nil, zero value otherwise.

### GetExternalKeyOk

`func (o *AutoTestApiResult) GetExternalKeyOk() (*string, bool)`

GetExternalKeyOk returns a tuple with the ExternalKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalKey

`func (o *AutoTestApiResult) SetExternalKey(v string)`

SetExternalKey sets ExternalKey field to given value.

### HasExternalKey

`func (o *AutoTestApiResult) HasExternalKey() bool`

HasExternalKey returns a boolean if a field has been set.

### SetExternalKeyNil

`func (o *AutoTestApiResult) SetExternalKeyNil(b bool)`

 SetExternalKeyNil sets the value for ExternalKey to be an explicit nil

### UnsetExternalKey
`func (o *AutoTestApiResult) UnsetExternalKey()`

UnsetExternalKey ensures that no value is present for ExternalKey, not even an explicit nil
### GetGlobalId

`func (o *AutoTestApiResult) GetGlobalId() int64`

GetGlobalId returns the GlobalId field if non-nil, zero value otherwise.

### GetGlobalIdOk

`func (o *AutoTestApiResult) GetGlobalIdOk() (*int64, bool)`

GetGlobalIdOk returns a tuple with the GlobalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalId

`func (o *AutoTestApiResult) SetGlobalId(v int64)`

SetGlobalId sets GlobalId field to given value.


### GetIsDeleted

`func (o *AutoTestApiResult) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *AutoTestApiResult) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *AutoTestApiResult) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.


### GetMustBeApproved

`func (o *AutoTestApiResult) GetMustBeApproved() bool`

GetMustBeApproved returns the MustBeApproved field if non-nil, zero value otherwise.

### GetMustBeApprovedOk

`func (o *AutoTestApiResult) GetMustBeApprovedOk() (*bool, bool)`

GetMustBeApprovedOk returns a tuple with the MustBeApproved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMustBeApproved

`func (o *AutoTestApiResult) SetMustBeApproved(v bool)`

SetMustBeApproved sets MustBeApproved field to given value.


### GetCreatedDate

`func (o *AutoTestApiResult) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *AutoTestApiResult) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *AutoTestApiResult) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.


### GetModifiedDate

`func (o *AutoTestApiResult) GetModifiedDate() time.Time`

GetModifiedDate returns the ModifiedDate field if non-nil, zero value otherwise.

### GetModifiedDateOk

`func (o *AutoTestApiResult) GetModifiedDateOk() (*time.Time, bool)`

GetModifiedDateOk returns a tuple with the ModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedDate

`func (o *AutoTestApiResult) SetModifiedDate(v time.Time)`

SetModifiedDate sets ModifiedDate field to given value.

### HasModifiedDate

`func (o *AutoTestApiResult) HasModifiedDate() bool`

HasModifiedDate returns a boolean if a field has been set.

### SetModifiedDateNil

`func (o *AutoTestApiResult) SetModifiedDateNil(b bool)`

 SetModifiedDateNil sets the value for ModifiedDate to be an explicit nil

### UnsetModifiedDate
`func (o *AutoTestApiResult) UnsetModifiedDate()`

UnsetModifiedDate ensures that no value is present for ModifiedDate, not even an explicit nil
### GetCreatedById

`func (o *AutoTestApiResult) GetCreatedById() string`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *AutoTestApiResult) GetCreatedByIdOk() (*string, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *AutoTestApiResult) SetCreatedById(v string)`

SetCreatedById sets CreatedById field to given value.


### GetModifiedById

`func (o *AutoTestApiResult) GetModifiedById() string`

GetModifiedById returns the ModifiedById field if non-nil, zero value otherwise.

### GetModifiedByIdOk

`func (o *AutoTestApiResult) GetModifiedByIdOk() (*string, bool)`

GetModifiedByIdOk returns a tuple with the ModifiedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedById

`func (o *AutoTestApiResult) SetModifiedById(v string)`

SetModifiedById sets ModifiedById field to given value.

### HasModifiedById

`func (o *AutoTestApiResult) HasModifiedById() bool`

HasModifiedById returns a boolean if a field has been set.

### SetModifiedByIdNil

`func (o *AutoTestApiResult) SetModifiedByIdNil(b bool)`

 SetModifiedByIdNil sets the value for ModifiedById to be an explicit nil

### UnsetModifiedById
`func (o *AutoTestApiResult) UnsetModifiedById()`

UnsetModifiedById ensures that no value is present for ModifiedById, not even an explicit nil
### GetLastTestRunId

`func (o *AutoTestApiResult) GetLastTestRunId() string`

GetLastTestRunId returns the LastTestRunId field if non-nil, zero value otherwise.

### GetLastTestRunIdOk

`func (o *AutoTestApiResult) GetLastTestRunIdOk() (*string, bool)`

GetLastTestRunIdOk returns a tuple with the LastTestRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTestRunId

`func (o *AutoTestApiResult) SetLastTestRunId(v string)`

SetLastTestRunId sets LastTestRunId field to given value.

### HasLastTestRunId

`func (o *AutoTestApiResult) HasLastTestRunId() bool`

HasLastTestRunId returns a boolean if a field has been set.

### SetLastTestRunIdNil

`func (o *AutoTestApiResult) SetLastTestRunIdNil(b bool)`

 SetLastTestRunIdNil sets the value for LastTestRunId to be an explicit nil

### UnsetLastTestRunId
`func (o *AutoTestApiResult) UnsetLastTestRunId()`

UnsetLastTestRunId ensures that no value is present for LastTestRunId, not even an explicit nil
### GetLastTestRunName

`func (o *AutoTestApiResult) GetLastTestRunName() string`

GetLastTestRunName returns the LastTestRunName field if non-nil, zero value otherwise.

### GetLastTestRunNameOk

`func (o *AutoTestApiResult) GetLastTestRunNameOk() (*string, bool)`

GetLastTestRunNameOk returns a tuple with the LastTestRunName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTestRunName

`func (o *AutoTestApiResult) SetLastTestRunName(v string)`

SetLastTestRunName sets LastTestRunName field to given value.

### HasLastTestRunName

`func (o *AutoTestApiResult) HasLastTestRunName() bool`

HasLastTestRunName returns a boolean if a field has been set.

### SetLastTestRunNameNil

`func (o *AutoTestApiResult) SetLastTestRunNameNil(b bool)`

 SetLastTestRunNameNil sets the value for LastTestRunName to be an explicit nil

### UnsetLastTestRunName
`func (o *AutoTestApiResult) UnsetLastTestRunName()`

UnsetLastTestRunName ensures that no value is present for LastTestRunName, not even an explicit nil
### GetLastTestResultId

`func (o *AutoTestApiResult) GetLastTestResultId() string`

GetLastTestResultId returns the LastTestResultId field if non-nil, zero value otherwise.

### GetLastTestResultIdOk

`func (o *AutoTestApiResult) GetLastTestResultIdOk() (*string, bool)`

GetLastTestResultIdOk returns a tuple with the LastTestResultId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTestResultId

`func (o *AutoTestApiResult) SetLastTestResultId(v string)`

SetLastTestResultId sets LastTestResultId field to given value.

### HasLastTestResultId

`func (o *AutoTestApiResult) HasLastTestResultId() bool`

HasLastTestResultId returns a boolean if a field has been set.

### SetLastTestResultIdNil

`func (o *AutoTestApiResult) SetLastTestResultIdNil(b bool)`

 SetLastTestResultIdNil sets the value for LastTestResultId to be an explicit nil

### UnsetLastTestResultId
`func (o *AutoTestApiResult) UnsetLastTestResultId()`

UnsetLastTestResultId ensures that no value is present for LastTestResultId, not even an explicit nil
### GetLastTestResultConfiguration

`func (o *AutoTestApiResult) GetLastTestResultConfiguration() ConfigurationShortApiResult`

GetLastTestResultConfiguration returns the LastTestResultConfiguration field if non-nil, zero value otherwise.

### GetLastTestResultConfigurationOk

`func (o *AutoTestApiResult) GetLastTestResultConfigurationOk() (*ConfigurationShortApiResult, bool)`

GetLastTestResultConfigurationOk returns a tuple with the LastTestResultConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTestResultConfiguration

`func (o *AutoTestApiResult) SetLastTestResultConfiguration(v ConfigurationShortApiResult)`

SetLastTestResultConfiguration sets LastTestResultConfiguration field to given value.

### HasLastTestResultConfiguration

`func (o *AutoTestApiResult) HasLastTestResultConfiguration() bool`

HasLastTestResultConfiguration returns a boolean if a field has been set.

### SetLastTestResultConfigurationNil

`func (o *AutoTestApiResult) SetLastTestResultConfigurationNil(b bool)`

 SetLastTestResultConfigurationNil sets the value for LastTestResultConfiguration to be an explicit nil

### UnsetLastTestResultConfiguration
`func (o *AutoTestApiResult) UnsetLastTestResultConfiguration()`

UnsetLastTestResultConfiguration ensures that no value is present for LastTestResultConfiguration, not even an explicit nil
### GetLastTestResultOutcome

`func (o *AutoTestApiResult) GetLastTestResultOutcome() string`

GetLastTestResultOutcome returns the LastTestResultOutcome field if non-nil, zero value otherwise.

### GetLastTestResultOutcomeOk

`func (o *AutoTestApiResult) GetLastTestResultOutcomeOk() (*string, bool)`

GetLastTestResultOutcomeOk returns a tuple with the LastTestResultOutcome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTestResultOutcome

`func (o *AutoTestApiResult) SetLastTestResultOutcome(v string)`

SetLastTestResultOutcome sets LastTestResultOutcome field to given value.

### HasLastTestResultOutcome

`func (o *AutoTestApiResult) HasLastTestResultOutcome() bool`

HasLastTestResultOutcome returns a boolean if a field has been set.

### SetLastTestResultOutcomeNil

`func (o *AutoTestApiResult) SetLastTestResultOutcomeNil(b bool)`

 SetLastTestResultOutcomeNil sets the value for LastTestResultOutcome to be an explicit nil

### UnsetLastTestResultOutcome
`func (o *AutoTestApiResult) UnsetLastTestResultOutcome()`

UnsetLastTestResultOutcome ensures that no value is present for LastTestResultOutcome, not even an explicit nil
### GetLastTestResultStatus

`func (o *AutoTestApiResult) GetLastTestResultStatus() TestStatusApiResult`

GetLastTestResultStatus returns the LastTestResultStatus field if non-nil, zero value otherwise.

### GetLastTestResultStatusOk

`func (o *AutoTestApiResult) GetLastTestResultStatusOk() (*TestStatusApiResult, bool)`

GetLastTestResultStatusOk returns a tuple with the LastTestResultStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTestResultStatus

`func (o *AutoTestApiResult) SetLastTestResultStatus(v TestStatusApiResult)`

SetLastTestResultStatus sets LastTestResultStatus field to given value.

### HasLastTestResultStatus

`func (o *AutoTestApiResult) HasLastTestResultStatus() bool`

HasLastTestResultStatus returns a boolean if a field has been set.

### SetLastTestResultStatusNil

`func (o *AutoTestApiResult) SetLastTestResultStatusNil(b bool)`

 SetLastTestResultStatusNil sets the value for LastTestResultStatus to be an explicit nil

### UnsetLastTestResultStatus
`func (o *AutoTestApiResult) UnsetLastTestResultStatus()`

UnsetLastTestResultStatus ensures that no value is present for LastTestResultStatus, not even an explicit nil
### GetStabilityPercentage

`func (o *AutoTestApiResult) GetStabilityPercentage() int64`

GetStabilityPercentage returns the StabilityPercentage field if non-nil, zero value otherwise.

### GetStabilityPercentageOk

`func (o *AutoTestApiResult) GetStabilityPercentageOk() (*int64, bool)`

GetStabilityPercentageOk returns a tuple with the StabilityPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStabilityPercentage

`func (o *AutoTestApiResult) SetStabilityPercentage(v int64)`

SetStabilityPercentage sets StabilityPercentage field to given value.

### HasStabilityPercentage

`func (o *AutoTestApiResult) HasStabilityPercentage() bool`

HasStabilityPercentage returns a boolean if a field has been set.

### SetStabilityPercentageNil

`func (o *AutoTestApiResult) SetStabilityPercentageNil(b bool)`

 SetStabilityPercentageNil sets the value for StabilityPercentage to be an explicit nil

### UnsetStabilityPercentage
`func (o *AutoTestApiResult) UnsetStabilityPercentage()`

UnsetStabilityPercentage ensures that no value is present for StabilityPercentage, not even an explicit nil
### GetLinks

`func (o *AutoTestApiResult) GetLinks() []LinkApiResult`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AutoTestApiResult) GetLinksOk() (*[]LinkApiResult, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AutoTestApiResult) SetLinks(v []LinkApiResult)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *AutoTestApiResult) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *AutoTestApiResult) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *AutoTestApiResult) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetLabels

`func (o *AutoTestApiResult) GetLabels() []LabelApiResult`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *AutoTestApiResult) GetLabelsOk() (*[]LabelApiResult, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *AutoTestApiResult) SetLabels(v []LabelApiResult)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *AutoTestApiResult) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### SetLabelsNil

`func (o *AutoTestApiResult) SetLabelsNil(b bool)`

 SetLabelsNil sets the value for Labels to be an explicit nil

### UnsetLabels
`func (o *AutoTestApiResult) UnsetLabels()`

UnsetLabels ensures that no value is present for Labels, not even an explicit nil
### GetTags

`func (o *AutoTestApiResult) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *AutoTestApiResult) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *AutoTestApiResult) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *AutoTestApiResult) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *AutoTestApiResult) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *AutoTestApiResult) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


