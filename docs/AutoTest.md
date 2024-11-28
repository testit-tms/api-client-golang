# AutoTest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalId** | **string** | External ID of the autotest | 
**Links** | Pointer to [**[]Link**](Link.md) | Collection of the autotest links | [optional] 
**ProjectId** | **string** | Unique ID of the autotest project | 
**Name** | **string** | Name of the autotest | 
**Namespace** | Pointer to **NullableString** | Name of the autotest namespace | [optional] 
**Classname** | Pointer to **NullableString** | Name of the autotest class | [optional] 
**Steps** | Pointer to [**[]AutoTestStep**](AutoTestStep.md) | Collection of the autotest steps | [optional] 
**Setup** | Pointer to [**[]AutoTestStep**](AutoTestStep.md) | Collection of the autotest setup steps | [optional] 
**Teardown** | Pointer to [**[]AutoTestStep**](AutoTestStep.md) | Collection of the autotest teardown steps | [optional] 
**Title** | Pointer to **NullableString** | Name of the autotest in autotest&#39;s card | [optional] 
**Description** | Pointer to **NullableString** | Description of the autotest in autotest&#39;s card | [optional] 
**Labels** | Pointer to [**[]Label**](Label.md) | Collection of the autotest labels | [optional] 
**IsFlaky** | Pointer to **NullableBool** | Indicates if the autotest is marked as flaky | [optional] 
**ExternalKey** | Pointer to **NullableString** | External key of the autotest | [optional] 
**GlobalId** | **int64** | Global ID of the autotest | 
**IsDeleted** | **bool** | Indicates if the autotest is deleted | 
**MustBeApproved** | **bool** | Indicates if the autotest has unapproved changes from linked work items | 
**Id** | **string** | Unique ID of the autotest | 
**CreatedDate** | **time.Time** | Creation date of the autotest | 
**ModifiedDate** | Pointer to **NullableTime** | Last modification date of the project | [optional] 
**CreatedById** | **string** | Unique ID of the project creator | 
**ModifiedById** | Pointer to **NullableString** | Unique ID of the project last editor | [optional] 
**LastTestRunId** | Pointer to **NullableString** | Unique ID of the autotest last test run | [optional] 
**LastTestRunName** | Pointer to **NullableString** | Name of the autotest last test run | [optional] 
**LastTestResultId** | Pointer to **NullableString** | Unique ID of the autotest last test result | [optional] 
**LastTestResultConfiguration** | Pointer to [**NullableConfigurationShort**](ConfigurationShort.md) | Configuration of the autotest last test result | [optional] 
**LastTestResultOutcome** | Pointer to **NullableString** | Outcome of the autotest last test result | [optional] 
**StabilityPercentage** | Pointer to **NullableInt32** | Stability percentage of the autotest | [optional] 

## Methods

### NewAutoTest

`func NewAutoTest(externalId string, projectId string, name string, globalId int64, isDeleted bool, mustBeApproved bool, id string, createdDate time.Time, createdById string, ) *AutoTest`

NewAutoTest instantiates a new AutoTest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoTestWithDefaults

`func NewAutoTestWithDefaults() *AutoTest`

NewAutoTestWithDefaults instantiates a new AutoTest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalId

`func (o *AutoTest) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *AutoTest) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *AutoTest) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.


### GetLinks

`func (o *AutoTest) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AutoTest) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AutoTest) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *AutoTest) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *AutoTest) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *AutoTest) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetProjectId

`func (o *AutoTest) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *AutoTest) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *AutoTest) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetName

`func (o *AutoTest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AutoTest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AutoTest) SetName(v string)`

SetName sets Name field to given value.


### GetNamespace

`func (o *AutoTest) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *AutoTest) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *AutoTest) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *AutoTest) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### SetNamespaceNil

`func (o *AutoTest) SetNamespaceNil(b bool)`

 SetNamespaceNil sets the value for Namespace to be an explicit nil

### UnsetNamespace
`func (o *AutoTest) UnsetNamespace()`

UnsetNamespace ensures that no value is present for Namespace, not even an explicit nil
### GetClassname

`func (o *AutoTest) GetClassname() string`

GetClassname returns the Classname field if non-nil, zero value otherwise.

### GetClassnameOk

`func (o *AutoTest) GetClassnameOk() (*string, bool)`

GetClassnameOk returns a tuple with the Classname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassname

`func (o *AutoTest) SetClassname(v string)`

SetClassname sets Classname field to given value.

### HasClassname

`func (o *AutoTest) HasClassname() bool`

HasClassname returns a boolean if a field has been set.

### SetClassnameNil

`func (o *AutoTest) SetClassnameNil(b bool)`

 SetClassnameNil sets the value for Classname to be an explicit nil

### UnsetClassname
`func (o *AutoTest) UnsetClassname()`

UnsetClassname ensures that no value is present for Classname, not even an explicit nil
### GetSteps

`func (o *AutoTest) GetSteps() []AutoTestStep`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *AutoTest) GetStepsOk() (*[]AutoTestStep, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *AutoTest) SetSteps(v []AutoTestStep)`

SetSteps sets Steps field to given value.

### HasSteps

`func (o *AutoTest) HasSteps() bool`

HasSteps returns a boolean if a field has been set.

### SetStepsNil

`func (o *AutoTest) SetStepsNil(b bool)`

 SetStepsNil sets the value for Steps to be an explicit nil

### UnsetSteps
`func (o *AutoTest) UnsetSteps()`

UnsetSteps ensures that no value is present for Steps, not even an explicit nil
### GetSetup

`func (o *AutoTest) GetSetup() []AutoTestStep`

GetSetup returns the Setup field if non-nil, zero value otherwise.

### GetSetupOk

`func (o *AutoTest) GetSetupOk() (*[]AutoTestStep, bool)`

GetSetupOk returns a tuple with the Setup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetup

`func (o *AutoTest) SetSetup(v []AutoTestStep)`

SetSetup sets Setup field to given value.

### HasSetup

`func (o *AutoTest) HasSetup() bool`

HasSetup returns a boolean if a field has been set.

### SetSetupNil

`func (o *AutoTest) SetSetupNil(b bool)`

 SetSetupNil sets the value for Setup to be an explicit nil

### UnsetSetup
`func (o *AutoTest) UnsetSetup()`

UnsetSetup ensures that no value is present for Setup, not even an explicit nil
### GetTeardown

`func (o *AutoTest) GetTeardown() []AutoTestStep`

GetTeardown returns the Teardown field if non-nil, zero value otherwise.

### GetTeardownOk

`func (o *AutoTest) GetTeardownOk() (*[]AutoTestStep, bool)`

GetTeardownOk returns a tuple with the Teardown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeardown

`func (o *AutoTest) SetTeardown(v []AutoTestStep)`

SetTeardown sets Teardown field to given value.

### HasTeardown

`func (o *AutoTest) HasTeardown() bool`

HasTeardown returns a boolean if a field has been set.

### SetTeardownNil

`func (o *AutoTest) SetTeardownNil(b bool)`

 SetTeardownNil sets the value for Teardown to be an explicit nil

### UnsetTeardown
`func (o *AutoTest) UnsetTeardown()`

UnsetTeardown ensures that no value is present for Teardown, not even an explicit nil
### GetTitle

`func (o *AutoTest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AutoTest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AutoTest) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AutoTest) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *AutoTest) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *AutoTest) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetDescription

`func (o *AutoTest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AutoTest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AutoTest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AutoTest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *AutoTest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *AutoTest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetLabels

`func (o *AutoTest) GetLabels() []Label`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *AutoTest) GetLabelsOk() (*[]Label, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *AutoTest) SetLabels(v []Label)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *AutoTest) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### SetLabelsNil

`func (o *AutoTest) SetLabelsNil(b bool)`

 SetLabelsNil sets the value for Labels to be an explicit nil

### UnsetLabels
`func (o *AutoTest) UnsetLabels()`

UnsetLabels ensures that no value is present for Labels, not even an explicit nil
### GetIsFlaky

`func (o *AutoTest) GetIsFlaky() bool`

GetIsFlaky returns the IsFlaky field if non-nil, zero value otherwise.

### GetIsFlakyOk

`func (o *AutoTest) GetIsFlakyOk() (*bool, bool)`

GetIsFlakyOk returns a tuple with the IsFlaky field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFlaky

`func (o *AutoTest) SetIsFlaky(v bool)`

SetIsFlaky sets IsFlaky field to given value.

### HasIsFlaky

`func (o *AutoTest) HasIsFlaky() bool`

HasIsFlaky returns a boolean if a field has been set.

### SetIsFlakyNil

`func (o *AutoTest) SetIsFlakyNil(b bool)`

 SetIsFlakyNil sets the value for IsFlaky to be an explicit nil

### UnsetIsFlaky
`func (o *AutoTest) UnsetIsFlaky()`

UnsetIsFlaky ensures that no value is present for IsFlaky, not even an explicit nil
### GetExternalKey

`func (o *AutoTest) GetExternalKey() string`

GetExternalKey returns the ExternalKey field if non-nil, zero value otherwise.

### GetExternalKeyOk

`func (o *AutoTest) GetExternalKeyOk() (*string, bool)`

GetExternalKeyOk returns a tuple with the ExternalKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalKey

`func (o *AutoTest) SetExternalKey(v string)`

SetExternalKey sets ExternalKey field to given value.

### HasExternalKey

`func (o *AutoTest) HasExternalKey() bool`

HasExternalKey returns a boolean if a field has been set.

### SetExternalKeyNil

`func (o *AutoTest) SetExternalKeyNil(b bool)`

 SetExternalKeyNil sets the value for ExternalKey to be an explicit nil

### UnsetExternalKey
`func (o *AutoTest) UnsetExternalKey()`

UnsetExternalKey ensures that no value is present for ExternalKey, not even an explicit nil
### GetGlobalId

`func (o *AutoTest) GetGlobalId() int64`

GetGlobalId returns the GlobalId field if non-nil, zero value otherwise.

### GetGlobalIdOk

`func (o *AutoTest) GetGlobalIdOk() (*int64, bool)`

GetGlobalIdOk returns a tuple with the GlobalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalId

`func (o *AutoTest) SetGlobalId(v int64)`

SetGlobalId sets GlobalId field to given value.


### GetIsDeleted

`func (o *AutoTest) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *AutoTest) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *AutoTest) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.


### GetMustBeApproved

`func (o *AutoTest) GetMustBeApproved() bool`

GetMustBeApproved returns the MustBeApproved field if non-nil, zero value otherwise.

### GetMustBeApprovedOk

`func (o *AutoTest) GetMustBeApprovedOk() (*bool, bool)`

GetMustBeApprovedOk returns a tuple with the MustBeApproved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMustBeApproved

`func (o *AutoTest) SetMustBeApproved(v bool)`

SetMustBeApproved sets MustBeApproved field to given value.


### GetId

`func (o *AutoTest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AutoTest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AutoTest) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedDate

`func (o *AutoTest) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *AutoTest) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *AutoTest) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.


### GetModifiedDate

`func (o *AutoTest) GetModifiedDate() time.Time`

GetModifiedDate returns the ModifiedDate field if non-nil, zero value otherwise.

### GetModifiedDateOk

`func (o *AutoTest) GetModifiedDateOk() (*time.Time, bool)`

GetModifiedDateOk returns a tuple with the ModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedDate

`func (o *AutoTest) SetModifiedDate(v time.Time)`

SetModifiedDate sets ModifiedDate field to given value.

### HasModifiedDate

`func (o *AutoTest) HasModifiedDate() bool`

HasModifiedDate returns a boolean if a field has been set.

### SetModifiedDateNil

`func (o *AutoTest) SetModifiedDateNil(b bool)`

 SetModifiedDateNil sets the value for ModifiedDate to be an explicit nil

### UnsetModifiedDate
`func (o *AutoTest) UnsetModifiedDate()`

UnsetModifiedDate ensures that no value is present for ModifiedDate, not even an explicit nil
### GetCreatedById

`func (o *AutoTest) GetCreatedById() string`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *AutoTest) GetCreatedByIdOk() (*string, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *AutoTest) SetCreatedById(v string)`

SetCreatedById sets CreatedById field to given value.


### GetModifiedById

`func (o *AutoTest) GetModifiedById() string`

GetModifiedById returns the ModifiedById field if non-nil, zero value otherwise.

### GetModifiedByIdOk

`func (o *AutoTest) GetModifiedByIdOk() (*string, bool)`

GetModifiedByIdOk returns a tuple with the ModifiedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedById

`func (o *AutoTest) SetModifiedById(v string)`

SetModifiedById sets ModifiedById field to given value.

### HasModifiedById

`func (o *AutoTest) HasModifiedById() bool`

HasModifiedById returns a boolean if a field has been set.

### SetModifiedByIdNil

`func (o *AutoTest) SetModifiedByIdNil(b bool)`

 SetModifiedByIdNil sets the value for ModifiedById to be an explicit nil

### UnsetModifiedById
`func (o *AutoTest) UnsetModifiedById()`

UnsetModifiedById ensures that no value is present for ModifiedById, not even an explicit nil
### GetLastTestRunId

`func (o *AutoTest) GetLastTestRunId() string`

GetLastTestRunId returns the LastTestRunId field if non-nil, zero value otherwise.

### GetLastTestRunIdOk

`func (o *AutoTest) GetLastTestRunIdOk() (*string, bool)`

GetLastTestRunIdOk returns a tuple with the LastTestRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTestRunId

`func (o *AutoTest) SetLastTestRunId(v string)`

SetLastTestRunId sets LastTestRunId field to given value.

### HasLastTestRunId

`func (o *AutoTest) HasLastTestRunId() bool`

HasLastTestRunId returns a boolean if a field has been set.

### SetLastTestRunIdNil

`func (o *AutoTest) SetLastTestRunIdNil(b bool)`

 SetLastTestRunIdNil sets the value for LastTestRunId to be an explicit nil

### UnsetLastTestRunId
`func (o *AutoTest) UnsetLastTestRunId()`

UnsetLastTestRunId ensures that no value is present for LastTestRunId, not even an explicit nil
### GetLastTestRunName

`func (o *AutoTest) GetLastTestRunName() string`

GetLastTestRunName returns the LastTestRunName field if non-nil, zero value otherwise.

### GetLastTestRunNameOk

`func (o *AutoTest) GetLastTestRunNameOk() (*string, bool)`

GetLastTestRunNameOk returns a tuple with the LastTestRunName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTestRunName

`func (o *AutoTest) SetLastTestRunName(v string)`

SetLastTestRunName sets LastTestRunName field to given value.

### HasLastTestRunName

`func (o *AutoTest) HasLastTestRunName() bool`

HasLastTestRunName returns a boolean if a field has been set.

### SetLastTestRunNameNil

`func (o *AutoTest) SetLastTestRunNameNil(b bool)`

 SetLastTestRunNameNil sets the value for LastTestRunName to be an explicit nil

### UnsetLastTestRunName
`func (o *AutoTest) UnsetLastTestRunName()`

UnsetLastTestRunName ensures that no value is present for LastTestRunName, not even an explicit nil
### GetLastTestResultId

`func (o *AutoTest) GetLastTestResultId() string`

GetLastTestResultId returns the LastTestResultId field if non-nil, zero value otherwise.

### GetLastTestResultIdOk

`func (o *AutoTest) GetLastTestResultIdOk() (*string, bool)`

GetLastTestResultIdOk returns a tuple with the LastTestResultId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTestResultId

`func (o *AutoTest) SetLastTestResultId(v string)`

SetLastTestResultId sets LastTestResultId field to given value.

### HasLastTestResultId

`func (o *AutoTest) HasLastTestResultId() bool`

HasLastTestResultId returns a boolean if a field has been set.

### SetLastTestResultIdNil

`func (o *AutoTest) SetLastTestResultIdNil(b bool)`

 SetLastTestResultIdNil sets the value for LastTestResultId to be an explicit nil

### UnsetLastTestResultId
`func (o *AutoTest) UnsetLastTestResultId()`

UnsetLastTestResultId ensures that no value is present for LastTestResultId, not even an explicit nil
### GetLastTestResultConfiguration

`func (o *AutoTest) GetLastTestResultConfiguration() ConfigurationShort`

GetLastTestResultConfiguration returns the LastTestResultConfiguration field if non-nil, zero value otherwise.

### GetLastTestResultConfigurationOk

`func (o *AutoTest) GetLastTestResultConfigurationOk() (*ConfigurationShort, bool)`

GetLastTestResultConfigurationOk returns a tuple with the LastTestResultConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTestResultConfiguration

`func (o *AutoTest) SetLastTestResultConfiguration(v ConfigurationShort)`

SetLastTestResultConfiguration sets LastTestResultConfiguration field to given value.

### HasLastTestResultConfiguration

`func (o *AutoTest) HasLastTestResultConfiguration() bool`

HasLastTestResultConfiguration returns a boolean if a field has been set.

### SetLastTestResultConfigurationNil

`func (o *AutoTest) SetLastTestResultConfigurationNil(b bool)`

 SetLastTestResultConfigurationNil sets the value for LastTestResultConfiguration to be an explicit nil

### UnsetLastTestResultConfiguration
`func (o *AutoTest) UnsetLastTestResultConfiguration()`

UnsetLastTestResultConfiguration ensures that no value is present for LastTestResultConfiguration, not even an explicit nil
### GetLastTestResultOutcome

`func (o *AutoTest) GetLastTestResultOutcome() string`

GetLastTestResultOutcome returns the LastTestResultOutcome field if non-nil, zero value otherwise.

### GetLastTestResultOutcomeOk

`func (o *AutoTest) GetLastTestResultOutcomeOk() (*string, bool)`

GetLastTestResultOutcomeOk returns a tuple with the LastTestResultOutcome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTestResultOutcome

`func (o *AutoTest) SetLastTestResultOutcome(v string)`

SetLastTestResultOutcome sets LastTestResultOutcome field to given value.

### HasLastTestResultOutcome

`func (o *AutoTest) HasLastTestResultOutcome() bool`

HasLastTestResultOutcome returns a boolean if a field has been set.

### SetLastTestResultOutcomeNil

`func (o *AutoTest) SetLastTestResultOutcomeNil(b bool)`

 SetLastTestResultOutcomeNil sets the value for LastTestResultOutcome to be an explicit nil

### UnsetLastTestResultOutcome
`func (o *AutoTest) UnsetLastTestResultOutcome()`

UnsetLastTestResultOutcome ensures that no value is present for LastTestResultOutcome, not even an explicit nil
### GetStabilityPercentage

`func (o *AutoTest) GetStabilityPercentage() int32`

GetStabilityPercentage returns the StabilityPercentage field if non-nil, zero value otherwise.

### GetStabilityPercentageOk

`func (o *AutoTest) GetStabilityPercentageOk() (*int32, bool)`

GetStabilityPercentageOk returns a tuple with the StabilityPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStabilityPercentage

`func (o *AutoTest) SetStabilityPercentage(v int32)`

SetStabilityPercentage sets StabilityPercentage field to given value.

### HasStabilityPercentage

`func (o *AutoTest) HasStabilityPercentage() bool`

HasStabilityPercentage returns a boolean if a field has been set.

### SetStabilityPercentageNil

`func (o *AutoTest) SetStabilityPercentageNil(b bool)`

 SetStabilityPercentageNil sets the value for StabilityPercentage to be an explicit nil

### UnsetStabilityPercentage
`func (o *AutoTest) UnsetStabilityPercentage()`

UnsetStabilityPercentage ensures that no value is present for StabilityPercentage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


