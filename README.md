# Api client for Test IT TMS

## Getting Started

### Compatibility

| Test IT | API Client      |
|---------|-----------------|
| 4.2     | 1.0.0           |
| 4.3     | 1.1.0           |
| 4.4     | 1.2.0           |
| 4.5     | 1.3.0           |
| 4.6     | 1.4.0           |
| 5.0     | 2.0.3           |
| 5.2     | 2.1.0           |
| 5.2.2   | 2.2.0-tms-5.2.2 |
| 5.3     | 3.0.1-tms-5.3   |
| 5.4     | 3.2.0-tms-5.4   |
| 5.4.1   | 3.2.1-tms-5.4.1 |
| Cloud   | 3.2.1           |

### Installation

```sh
go get github.com/testit-tms/api-client-golang/v3
```

## Examples

Please follow the [installation](#installation) instruction and execute the following Golang code:

```go
import (
	"context"
	"fmt"

	tmsclient "github.com/testit-tms/api-client-golang/v2"
)

const (
	host   = "Your TMS address"
	key    = "Your private token"
	scheme = "https or http"
)

func main() {
	ctx := context.WithValue(context.Background(), tmsclient.ContextAPIKeys, map[string]tmsclient.APIKey{
		"Bearer or PrivateToken": {
			Key:    key,
			Prefix: "PrivateToken",
		},
	})

	configuration := tmsclient.NewConfiguration()
	configuration.Host = host
	configuration.Scheme = scheme
	apiClient := tmsclient.NewAPIClient(configuration)

	id := "123"

	_, err := apiClient.AttachmentsApi.ApiV2AttachmentsIdDelete(ctx, id).Execute()
	if err != nil {
		fmt.Println(err)
	}
}
```

## Documentation for API Endpoints

All URIs are relative to *http://localhost*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AttachmentsAPI* | [**ApiV2AttachmentsIdDelete**](docs/AttachmentsAPI.md#apiv2attachmentsiddelete) | **Delete** /api/v2/attachments/{id} | Delete attachment file
*AttachmentsAPI* | [**ApiV2AttachmentsIdGet**](docs/AttachmentsAPI.md#apiv2attachmentsidget) | **Get** /api/v2/attachments/{id} | Download attachment file
*AttachmentsAPI* | [**ApiV2AttachmentsIdMetadataGet**](docs/AttachmentsAPI.md#apiv2attachmentsidmetadataget) | **Get** /api/v2/attachments/{id}/metadata | Get attachment metadata
*AttachmentsAPI* | [**ApiV2AttachmentsOccupiedFileStorageSizeGet**](docs/AttachmentsAPI.md#apiv2attachmentsoccupiedfilestoragesizeget) | **Get** /api/v2/attachments/occupiedFileStorageSize | Get size of attachments storage in bytes
*AttachmentsAPI* | [**ApiV2AttachmentsPost**](docs/AttachmentsAPI.md#apiv2attachmentspost) | **Post** /api/v2/attachments | Upload new attachment file
*AutoTestsAPI* | [**ApiV2AutoTestsDelete**](docs/AutoTestsAPI.md#apiv2autotestsdelete) | **Delete** /api/v2/autoTests | Delete autotests
*AutoTestsAPI* | [**ApiV2AutoTestsFlakyBulkPost**](docs/AutoTestsAPI.md#apiv2autotestsflakybulkpost) | **Post** /api/v2/autoTests/flaky/bulk | Set \&quot;Flaky\&quot; status for multiple autotests
*AutoTestsAPI* | [**ApiV2AutoTestsIdPatch**](docs/AutoTestsAPI.md#apiv2autotestsidpatch) | **Patch** /api/v2/autoTests/{id} | Patch auto test
*AutoTestsAPI* | [**ApiV2AutoTestsIdTestResultsSearchPost**](docs/AutoTestsAPI.md#apiv2autotestsidtestresultssearchpost) | **Post** /api/v2/autoTests/{id}/testResults/search | Get test results history for autotest
*AutoTestsAPI* | [**ApiV2AutoTestsIdWorkItemsChangedIdGet**](docs/AutoTestsAPI.md#apiv2autotestsidworkitemschangedidget) | **Get** /api/v2/autoTests/{id}/workItems/changed/id | Get identifiers of changed linked work items
*AutoTestsAPI* | [**ApiV2AutoTestsIdWorkItemsChangedWorkItemIdApprovePost**](docs/AutoTestsAPI.md#apiv2autotestsidworkitemschangedworkitemidapprovepost) | **Post** /api/v2/autoTests/{id}/workItems/changed/{workItemId}/approve | Approve changes to work items linked to autotest
*AutoTestsAPI* | [**ApiV2AutoTestsSearchPost**](docs/AutoTestsAPI.md#apiv2autotestssearchpost) | **Post** /api/v2/autoTests/search | Search for autotests
*AutoTestsAPI* | [**CreateAutoTest**](docs/AutoTestsAPI.md#createautotest) | **Post** /api/v2/autoTests | Create autotest
*AutoTestsAPI* | [**CreateMultiple**](docs/AutoTestsAPI.md#createmultiple) | **Post** /api/v2/autoTests/bulk | Create multiple autotests
*AutoTestsAPI* | [**DeleteAutoTest**](docs/AutoTestsAPI.md#deleteautotest) | **Delete** /api/v2/autoTests/{id} | Delete autotest
*AutoTestsAPI* | [**DeleteAutoTestLinkFromWorkItem**](docs/AutoTestsAPI.md#deleteautotestlinkfromworkitem) | **Delete** /api/v2/autoTests/{id}/workItems | Unlink autotest from work item
*AutoTestsAPI* | [**GetAllAutoTests**](docs/AutoTestsAPI.md#getallautotests) | **Get** /api/v2/autoTests | 
*AutoTestsAPI* | [**GetAutoTestAverageDuration**](docs/AutoTestsAPI.md#getautotestaverageduration) | **Get** /api/v2/autoTests/{id}/averageDuration | Get average autotest duration
*AutoTestsAPI* | [**GetAutoTestById**](docs/AutoTestsAPI.md#getautotestbyid) | **Get** /api/v2/autoTests/{id} | Get autotest by internal or global ID
*AutoTestsAPI* | [**GetAutoTestChronology**](docs/AutoTestsAPI.md#getautotestchronology) | **Get** /api/v2/autoTests/{id}/chronology | Get autotest chronology
*AutoTestsAPI* | [**GetTestRuns**](docs/AutoTestsAPI.md#gettestruns) | **Get** /api/v2/autoTests/{id}/testRuns | Get completed tests runs for autotests
*AutoTestsAPI* | [**GetWorkItemsLinkedToAutoTest**](docs/AutoTestsAPI.md#getworkitemslinkedtoautotest) | **Get** /api/v2/autoTests/{id}/workItems | Get work items linked to autotest
*AutoTestsAPI* | [**LinkAutoTestToWorkItem**](docs/AutoTestsAPI.md#linkautotesttoworkitem) | **Post** /api/v2/autoTests/{id}/workItems | Link autotest with work items
*AutoTestsAPI* | [**UpdateAutoTest**](docs/AutoTestsAPI.md#updateautotest) | **Put** /api/v2/autoTests | Update autotest
*AutoTestsAPI* | [**UpdateMultiple**](docs/AutoTestsAPI.md#updatemultiple) | **Put** /api/v2/autoTests/bulk | Update multiple autotests
*BackgroundJobsAPI* | [**ApiV2BackgroundJobsCompletedDelete**](docs/BackgroundJobsAPI.md#apiv2backgroundjobscompleteddelete) | **Delete** /api/v2/backgroundJobs/completed | Delete all completed background jobs
*BackgroundJobsAPI* | [**ApiV2BackgroundJobsGet**](docs/BackgroundJobsAPI.md#apiv2backgroundjobsget) | **Get** /api/v2/backgroundJobs | 
*BackgroundJobsAPI* | [**ApiV2BackgroundJobsIdCancelPost**](docs/BackgroundJobsAPI.md#apiv2backgroundjobsidcancelpost) | **Post** /api/v2/backgroundJobs/{id}/cancel | Cancel current user background job
*BackgroundJobsAPI* | [**ApiV2BackgroundJobsIdGet**](docs/BackgroundJobsAPI.md#apiv2backgroundjobsidget) | **Get** /api/v2/backgroundJobs/{id} | Get background job by ID
*BackgroundJobsAPI* | [**ApiV2BackgroundJobsIdStatusGet**](docs/BackgroundJobsAPI.md#apiv2backgroundjobsidstatusget) | **Get** /api/v2/backgroundJobs/{id}/status | Get background job status by job ID
*BackgroundJobsAPI* | [**ApiV2BackgroundJobsSearchPost**](docs/BackgroundJobsAPI.md#apiv2backgroundjobssearchpost) | **Post** /api/v2/backgroundJobs/search | Search for user background jobs
*ConfigurationsAPI* | [**ApiV2ConfigurationsCreateByParametersPost**](docs/ConfigurationsAPI.md#apiv2configurationscreatebyparameterspost) | **Post** /api/v2/configurations/createByParameters | Create configurations by parameters
*ConfigurationsAPI* | [**ApiV2ConfigurationsDeleteBulkPost**](docs/ConfigurationsAPI.md#apiv2configurationsdeletebulkpost) | **Post** /api/v2/configurations/delete/bulk | Delete multiple configurations
*ConfigurationsAPI* | [**ApiV2ConfigurationsIdDelete**](docs/ConfigurationsAPI.md#apiv2configurationsiddelete) | **Delete** /api/v2/configurations/{id} | Delete configuration
*ConfigurationsAPI* | [**ApiV2ConfigurationsIdPatch**](docs/ConfigurationsAPI.md#apiv2configurationsidpatch) | **Patch** /api/v2/configurations/{id} | Patch configuration
*ConfigurationsAPI* | [**ApiV2ConfigurationsIdPurgePost**](docs/ConfigurationsAPI.md#apiv2configurationsidpurgepost) | **Post** /api/v2/configurations/{id}/purge | Permanently delete configuration from archive
*ConfigurationsAPI* | [**ApiV2ConfigurationsIdRestorePost**](docs/ConfigurationsAPI.md#apiv2configurationsidrestorepost) | **Post** /api/v2/configurations/{id}/restore | Restore configuration from the archive
*ConfigurationsAPI* | [**ApiV2ConfigurationsPurgeBulkPost**](docs/ConfigurationsAPI.md#apiv2configurationspurgebulkpost) | **Post** /api/v2/configurations/purge/bulk | Permanently delete multiple archived configurations
*ConfigurationsAPI* | [**ApiV2ConfigurationsPut**](docs/ConfigurationsAPI.md#apiv2configurationsput) | **Put** /api/v2/configurations | Edit configuration
*ConfigurationsAPI* | [**ApiV2ConfigurationsRestoreBulkPost**](docs/ConfigurationsAPI.md#apiv2configurationsrestorebulkpost) | **Post** /api/v2/configurations/restore/bulk | Restore multiple configurations from the archive
*ConfigurationsAPI* | [**ApiV2ConfigurationsSearchPost**](docs/ConfigurationsAPI.md#apiv2configurationssearchpost) | **Post** /api/v2/configurations/search | Search for configurations
*ConfigurationsAPI* | [**CreateConfiguration**](docs/ConfigurationsAPI.md#createconfiguration) | **Post** /api/v2/configurations | Create Configuration
*ConfigurationsAPI* | [**GetConfigurationById**](docs/ConfigurationsAPI.md#getconfigurationbyid) | **Get** /api/v2/configurations/{id} | Get configuration by internal or global ID
*CustomAttributeTemplatesAPI* | [**ApiV2CustomAttributesTemplatesExistsGet**](docs/CustomAttributeTemplatesAPI.md#apiv2customattributestemplatesexistsget) | **Get** /api/v2/customAttributes/templates/exists | 
*CustomAttributeTemplatesAPI* | [**ApiV2CustomAttributesTemplatesIdCustomAttributesExcludePost**](docs/CustomAttributeTemplatesAPI.md#apiv2customattributestemplatesidcustomattributesexcludepost) | **Post** /api/v2/customAttributes/templates/{id}/customAttributes/exclude | Exclude CustomAttributes from CustomAttributeTemplate
*CustomAttributeTemplatesAPI* | [**ApiV2CustomAttributesTemplatesIdCustomAttributesIncludePost**](docs/CustomAttributeTemplatesAPI.md#apiv2customattributestemplatesidcustomattributesincludepost) | **Post** /api/v2/customAttributes/templates/{id}/customAttributes/include | Include CustomAttributes to CustomAttributeTemplate
*CustomAttributeTemplatesAPI* | [**ApiV2CustomAttributesTemplatesIdDelete**](docs/CustomAttributeTemplatesAPI.md#apiv2customattributestemplatesiddelete) | **Delete** /api/v2/customAttributes/templates/{id} | Delete CustomAttributeTemplate
*CustomAttributeTemplatesAPI* | [**ApiV2CustomAttributesTemplatesIdGet**](docs/CustomAttributeTemplatesAPI.md#apiv2customattributestemplatesidget) | **Get** /api/v2/customAttributes/templates/{id} | Get CustomAttributeTemplate by ID
*CustomAttributeTemplatesAPI* | [**ApiV2CustomAttributesTemplatesNameGet**](docs/CustomAttributeTemplatesAPI.md#apiv2customattributestemplatesnameget) | **Get** /api/v2/customAttributes/templates/{name} | Get CustomAttributeTemplate by name
*CustomAttributeTemplatesAPI* | [**ApiV2CustomAttributesTemplatesPost**](docs/CustomAttributeTemplatesAPI.md#apiv2customattributestemplatespost) | **Post** /api/v2/customAttributes/templates | Create CustomAttributeTemplate
*CustomAttributeTemplatesAPI* | [**ApiV2CustomAttributesTemplatesPut**](docs/CustomAttributeTemplatesAPI.md#apiv2customattributestemplatesput) | **Put** /api/v2/customAttributes/templates | Update custom attributes template
*CustomAttributeTemplatesAPI* | [**ApiV2CustomAttributesTemplatesSearchPost**](docs/CustomAttributeTemplatesAPI.md#apiv2customattributestemplatessearchpost) | **Post** /api/v2/customAttributes/templates/search | Search CustomAttributeTemplates
*CustomAttributesAPI* | [**ApiV2CustomAttributesExistsGet**](docs/CustomAttributesAPI.md#apiv2customattributesexistsget) | **Get** /api/v2/customAttributes/exists | 
*CustomAttributesAPI* | [**ApiV2CustomAttributesGlobalIdDelete**](docs/CustomAttributesAPI.md#apiv2customattributesglobaliddelete) | **Delete** /api/v2/customAttributes/global/{id} | Delete global attribute
*CustomAttributesAPI* | [**ApiV2CustomAttributesGlobalIdPut**](docs/CustomAttributesAPI.md#apiv2customattributesglobalidput) | **Put** /api/v2/customAttributes/global/{id} | Edit global attribute
*CustomAttributesAPI* | [**ApiV2CustomAttributesGlobalPost**](docs/CustomAttributesAPI.md#apiv2customattributesglobalpost) | **Post** /api/v2/customAttributes/global | Create global attribute
*CustomAttributesAPI* | [**ApiV2CustomAttributesIdGet**](docs/CustomAttributesAPI.md#apiv2customattributesidget) | **Get** /api/v2/customAttributes/{id} | Get attribute
*CustomAttributesAPI* | [**ApiV2CustomAttributesSearchPost**](docs/CustomAttributesAPI.md#apiv2customattributessearchpost) | **Post** /api/v2/customAttributes/search | Search for attributes
*NotificationsAPI* | [**ApiV2NotificationsCountGet**](docs/NotificationsAPI.md#apiv2notificationscountget) | **Get** /api/v2/notifications/count | Get unread Notifications total in last 7 days
*NotificationsAPI* | [**ApiV2NotificationsGet**](docs/NotificationsAPI.md#apiv2notificationsget) | **Get** /api/v2/notifications | Get all Notifications for current User
*NotificationsAPI* | [**ApiV2NotificationsIdReadPost**](docs/NotificationsAPI.md#apiv2notificationsidreadpost) | **Post** /api/v2/notifications/{id}/read | Set Notification as read
*NotificationsAPI* | [**ApiV2NotificationsReadPost**](docs/NotificationsAPI.md#apiv2notificationsreadpost) | **Post** /api/v2/notifications/read | Set all Notifications as read
*NotificationsAPI* | [**ApiV2NotificationsSearchPost**](docs/NotificationsAPI.md#apiv2notificationssearchpost) | **Post** /api/v2/notifications/search | Search Notifications for current User
*ParametersAPI* | [**ApiV2ParametersBulkPost**](docs/ParametersAPI.md#apiv2parametersbulkpost) | **Post** /api/v2/parameters/bulk | Create multiple parameters
*ParametersAPI* | [**ApiV2ParametersBulkPut**](docs/ParametersAPI.md#apiv2parametersbulkput) | **Put** /api/v2/parameters/bulk | Update multiple parameters
*ParametersAPI* | [**ApiV2ParametersGroupsGet**](docs/ParametersAPI.md#apiv2parametersgroupsget) | **Get** /api/v2/parameters/groups | Get parameters as group
*ParametersAPI* | [**ApiV2ParametersKeyNameNameExistsGet**](docs/ParametersAPI.md#apiv2parameterskeynamenameexistsget) | **Get** /api/v2/parameters/key/name/{name}/exists | Check existence parameter key in system
*ParametersAPI* | [**ApiV2ParametersKeyValuesGet**](docs/ParametersAPI.md#apiv2parameterskeyvaluesget) | **Get** /api/v2/parameters/{key}/values | Get all parameter key values
*ParametersAPI* | [**ApiV2ParametersKeysGet**](docs/ParametersAPI.md#apiv2parameterskeysget) | **Get** /api/v2/parameters/keys | Get all parameter keys
*ParametersAPI* | [**ApiV2ParametersSearchGroupsPost**](docs/ParametersAPI.md#apiv2parameterssearchgroupspost) | **Post** /api/v2/parameters/search/groups | Search for parameters as group
*ParametersAPI* | [**ApiV2ParametersSearchPost**](docs/ParametersAPI.md#apiv2parameterssearchpost) | **Post** /api/v2/parameters/search | Search for parameters
*ParametersAPI* | [**CreateParameter**](docs/ParametersAPI.md#createparameter) | **Post** /api/v2/parameters | Create parameter
*ParametersAPI* | [**DeleteByName**](docs/ParametersAPI.md#deletebyname) | **Delete** /api/v2/parameters/name/{name} | Delete parameter by name
*ParametersAPI* | [**DeleteByParameterKeyId**](docs/ParametersAPI.md#deletebyparameterkeyid) | **Delete** /api/v2/parameters/keyId/{keyId} | Delete parameters by parameter key identifier
*ParametersAPI* | [**DeleteParameter**](docs/ParametersAPI.md#deleteparameter) | **Delete** /api/v2/parameters/{id} | Delete parameter
*ParametersAPI* | [**GetAllParameters**](docs/ParametersAPI.md#getallparameters) | **Get** /api/v2/parameters | Get all parameters
*ParametersAPI* | [**GetParameterById**](docs/ParametersAPI.md#getparameterbyid) | **Get** /api/v2/parameters/{id} | Get parameter by ID
*ParametersAPI* | [**UpdateParameter**](docs/ParametersAPI.md#updateparameter) | **Put** /api/v2/parameters | Update parameter
*ProjectAttributeTemplatesAPI* | [**ApiV2ProjectsProjectIdAttributesTemplatesSearchPost**](docs/ProjectAttributeTemplatesAPI.md#apiv2projectsprojectidattributestemplatessearchpost) | **Post** /api/v2/projects/{projectId}/attributes/templates/search | Search for custom attributes templates
*ProjectAttributeTemplatesAPI* | [**ApiV2ProjectsProjectIdAttributesTemplatesTemplateIdDelete**](docs/ProjectAttributeTemplatesAPI.md#apiv2projectsprojectidattributestemplatestemplateiddelete) | **Delete** /api/v2/projects/{projectId}/attributes/templates/{templateId} | Delete CustomAttributeTemplate from Project
*ProjectAttributeTemplatesAPI* | [**ApiV2ProjectsProjectIdAttributesTemplatesTemplateIdPost**](docs/ProjectAttributeTemplatesAPI.md#apiv2projectsprojectidattributestemplatestemplateidpost) | **Post** /api/v2/projects/{projectId}/attributes/templates/{templateId} | Add CustomAttributeTemplate to Project
*ProjectAttributesAPI* | [**CreateProjectsAttribute**](docs/ProjectAttributesAPI.md#createprojectsattribute) | **Post** /api/v2/projects/{projectId}/attributes | Create project attribute
*ProjectAttributesAPI* | [**DeleteProjectsAttribute**](docs/ProjectAttributesAPI.md#deleteprojectsattribute) | **Delete** /api/v2/projects/{projectId}/attributes/{attributeId} | Delete project attribute
*ProjectAttributesAPI* | [**GetAttributeByProjectId**](docs/ProjectAttributesAPI.md#getattributebyprojectid) | **Get** /api/v2/projects/{projectId}/attributes/{attributeId} | Get project attribute
*ProjectAttributesAPI* | [**GetAttributesByProjectId**](docs/ProjectAttributesAPI.md#getattributesbyprojectid) | **Get** /api/v2/projects/{projectId}/attributes | Get project attributes
*ProjectAttributesAPI* | [**SearchAttributesInProject**](docs/ProjectAttributesAPI.md#searchattributesinproject) | **Post** /api/v2/projects/{projectId}/attributes/search | Search for attributes used in the project
*ProjectAttributesAPI* | [**UpdateProjectsAttribute**](docs/ProjectAttributesAPI.md#updateprojectsattribute) | **Put** /api/v2/projects/{projectId}/attributes | Edit attribute of the project
*ProjectConfigurationsAPI* | [**GetConfigurationsByProjectId**](docs/ProjectConfigurationsAPI.md#getconfigurationsbyprojectid) | **Get** /api/v2/projects/{projectId}/configurations | Get project configurations
*ProjectSectionsAPI* | [**GetSectionsByProjectId**](docs/ProjectSectionsAPI.md#getsectionsbyprojectid) | **Get** /api/v2/projects/{projectId}/sections | Get project sections
*ProjectSettingsAPI* | [**ApiV2ProjectsProjectIdSettingsAutotestsPost**](docs/ProjectSettingsAPI.md#apiv2projectsprojectidsettingsautotestspost) | **Post** /api/v2/projects/{projectId}/settings/autotests | Set autotest project settings.
*ProjectSettingsAPI* | [**GetAutotestProjectSettings**](docs/ProjectSettingsAPI.md#getautotestprojectsettings) | **Get** /api/v2/projects/{projectId}/settings/autotests | Get autotest project settings.
*ProjectTestPlanAttributesAPI* | [**CreateCustomAttributeTestPlanProjectRelations**](docs/ProjectTestPlanAttributesAPI.md#createcustomattributetestplanprojectrelations) | **Post** /api/v2/projects/{projectId}/testPlans/attributes | Add attributes to project&#39;s test plans
*ProjectTestPlanAttributesAPI* | [**DeleteCustomAttributeTestPlanProjectRelations**](docs/ProjectTestPlanAttributesAPI.md#deletecustomattributetestplanprojectrelations) | **Delete** /api/v2/projects/{projectId}/testPlans/attributes/{attributeId} | Delete attribute from project&#39;s test plans
*ProjectTestPlanAttributesAPI* | [**GetCustomAttributeTestPlanProjectRelations**](docs/ProjectTestPlanAttributesAPI.md#getcustomattributetestplanprojectrelations) | **Get** /api/v2/projects/{projectId}/testPlans/attributes | Get project&#39;s test plan attributes
*ProjectTestPlanAttributesAPI* | [**SearchTestPlanAttributesInProject**](docs/ProjectTestPlanAttributesAPI.md#searchtestplanattributesinproject) | **Post** /api/v2/projects/{projectId}/testPlans/attributes/search | Search for attributes used in the project test plans
*ProjectTestPlanAttributesAPI* | [**UpdateCustomAttributeTestPlanProjectRelations**](docs/ProjectTestPlanAttributesAPI.md#updatecustomattributetestplanprojectrelations) | **Put** /api/v2/projects/{projectId}/testPlans/attributes | Update attribute of project&#39;s test plans
*ProjectTestPlansAPI* | [**ApiV2ProjectsProjectIdTestPlansAnalyticsGet**](docs/ProjectTestPlansAPI.md#apiv2projectsprojectidtestplansanalyticsget) | **Get** /api/v2/projects/{projectId}/testPlans/analytics | Get TestPlans analytics
*ProjectTestPlansAPI* | [**ApiV2ProjectsProjectIdTestPlansDeleteBulkPost**](docs/ProjectTestPlansAPI.md#apiv2projectsprojectidtestplansdeletebulkpost) | **Post** /api/v2/projects/{projectId}/testPlans/delete/bulk | Delete multiple test plans
*ProjectTestPlansAPI* | [**ApiV2ProjectsProjectIdTestPlansNameExistsGet**](docs/ProjectTestPlansAPI.md#apiv2projectsprojectidtestplansnameexistsget) | **Get** /api/v2/projects/{projectId}/testPlans/{name}/exists | Checks if TestPlan exists with the specified name exists for the project
*ProjectTestPlansAPI* | [**ApiV2ProjectsProjectIdTestPlansPurgeBulkPost**](docs/ProjectTestPlansAPI.md#apiv2projectsprojectidtestplanspurgebulkpost) | **Post** /api/v2/projects/{projectId}/testPlans/purge/bulk | Permanently delete multiple archived test plans
*ProjectTestPlansAPI* | [**ApiV2ProjectsProjectIdTestPlansRestoreBulkPost**](docs/ProjectTestPlansAPI.md#apiv2projectsprojectidtestplansrestorebulkpost) | **Post** /api/v2/projects/{projectId}/testPlans/restore/bulk | Restore multiple test plans
*ProjectTestPlansAPI* | [**ApiV2ProjectsProjectIdTestPlansSearchPost**](docs/ProjectTestPlansAPI.md#apiv2projectsprojectidtestplanssearchpost) | **Post** /api/v2/projects/{projectId}/testPlans/search | Get Project TestPlans with analytics
*ProjectWorkItemsAPI* | [**ApiV2ProjectsProjectIdWorkItemsSearchGroupedPost**](docs/ProjectWorkItemsAPI.md#apiv2projectsprojectidworkitemssearchgroupedpost) | **Post** /api/v2/projects/{projectId}/workItems/search/grouped | Search for work items and group results by attribute
*ProjectWorkItemsAPI* | [**ApiV2ProjectsProjectIdWorkItemsSearchIdPost**](docs/ProjectWorkItemsAPI.md#apiv2projectsprojectidworkitemssearchidpost) | **Post** /api/v2/projects/{projectId}/workItems/search/id | Search for work items and extract IDs only
*ProjectWorkItemsAPI* | [**ApiV2ProjectsProjectIdWorkItemsSearchPost**](docs/ProjectWorkItemsAPI.md#apiv2projectsprojectidworkitemssearchpost) | **Post** /api/v2/projects/{projectId}/workItems/search | Search for work items
*ProjectWorkItemsAPI* | [**ApiV2ProjectsProjectIdWorkItemsTagsGet**](docs/ProjectWorkItemsAPI.md#apiv2projectsprojectidworkitemstagsget) | **Get** /api/v2/projects/{projectId}/workItems/tags | Get WorkItems Tags
*ProjectWorkItemsAPI* | [**GetWorkItemsByProjectId**](docs/ProjectWorkItemsAPI.md#getworkitemsbyprojectid) | **Get** /api/v2/projects/{projectId}/workItems | Get project work items
*ProjectsAPI* | [**AddGlobaAttributesToProject**](docs/ProjectsAPI.md#addglobaattributestoproject) | **Post** /api/v2/projects/{id}/globalAttributes | Add global attributes to project
*ProjectsAPI* | [**ApiV2ProjectsDemoPost**](docs/ProjectsAPI.md#apiv2projectsdemopost) | **Post** /api/v2/projects/demo | 
*ProjectsAPI* | [**ApiV2ProjectsIdDelete**](docs/ProjectsAPI.md#apiv2projectsiddelete) | **Delete** /api/v2/projects/{id} | Archive project
*ProjectsAPI* | [**ApiV2ProjectsIdFailureClassesGet**](docs/ProjectsAPI.md#apiv2projectsidfailureclassesget) | **Get** /api/v2/projects/{id}/failureClasses | Get failure classes
*ProjectsAPI* | [**ApiV2ProjectsIdFavoritePut**](docs/ProjectsAPI.md#apiv2projectsidfavoriteput) | **Put** /api/v2/projects/{id}/favorite | Mark Project as favorite
*ProjectsAPI* | [**ApiV2ProjectsIdFiltersGet**](docs/ProjectsAPI.md#apiv2projectsidfiltersget) | **Get** /api/v2/projects/{id}/filters | Get Project filters
*ProjectsAPI* | [**ApiV2ProjectsIdPatch**](docs/ProjectsAPI.md#apiv2projectsidpatch) | **Patch** /api/v2/projects/{id} | Patch project
*ProjectsAPI* | [**ApiV2ProjectsIdPurgePost**](docs/ProjectsAPI.md#apiv2projectsidpurgepost) | **Post** /api/v2/projects/{id}/purge | Purge the project
*ProjectsAPI* | [**ApiV2ProjectsIdRestorePost**](docs/ProjectsAPI.md#apiv2projectsidrestorepost) | **Post** /api/v2/projects/{id}/restore | Restore archived project
*ProjectsAPI* | [**ApiV2ProjectsIdTestPlansAttributeAttributeIdDelete**](docs/ProjectsAPI.md#apiv2projectsidtestplansattributeattributeiddelete) | **Delete** /api/v2/projects/{id}/testPlans/attribute/{attributeId} | Delete attribute from project&#39;s test plans
*ProjectsAPI* | [**ApiV2ProjectsIdTestPlansAttributePut**](docs/ProjectsAPI.md#apiv2projectsidtestplansattributeput) | **Put** /api/v2/projects/{id}/testPlans/attribute | Update attribute of project&#39;s test plans
*ProjectsAPI* | [**ApiV2ProjectsIdTestRunsActiveGet**](docs/ProjectsAPI.md#apiv2projectsidtestrunsactiveget) | **Get** /api/v2/projects/{id}/testRuns/active | Get active Project TestRuns
*ProjectsAPI* | [**ApiV2ProjectsIdTestRunsFullGet**](docs/ProjectsAPI.md#apiv2projectsidtestrunsfullget) | **Get** /api/v2/projects/{id}/testRuns/full | Get Project TestRuns full models
*ProjectsAPI* | [**ApiV2ProjectsNameNameExistsGet**](docs/ProjectsAPI.md#apiv2projectsnamenameexistsget) | **Get** /api/v2/projects/name/{name}/exists | 
*ProjectsAPI* | [**ApiV2ProjectsPurgeBulkPost**](docs/ProjectsAPI.md#apiv2projectspurgebulkpost) | **Post** /api/v2/projects/purge/bulk | Purge multiple projects
*ProjectsAPI* | [**ApiV2ProjectsRestoreBulkPost**](docs/ProjectsAPI.md#apiv2projectsrestorebulkpost) | **Post** /api/v2/projects/restore/bulk | Restore multiple projects
*ProjectsAPI* | [**ApiV2ProjectsSearchPost**](docs/ProjectsAPI.md#apiv2projectssearchpost) | **Post** /api/v2/projects/search | Search for projects
*ProjectsAPI* | [**ApiV2ProjectsShortsPost**](docs/ProjectsAPI.md#apiv2projectsshortspost) | **Post** /api/v2/projects/shorts | Get projects short models
*ProjectsAPI* | [**CreateProject**](docs/ProjectsAPI.md#createproject) | **Post** /api/v2/projects | Create project
*ProjectsAPI* | [**DeleteProjectAutoTests**](docs/ProjectsAPI.md#deleteprojectautotests) | **Delete** /api/v2/projects/{id}/autoTests | Delete all autotests from project
*ProjectsAPI* | [**GetAllProjects**](docs/ProjectsAPI.md#getallprojects) | **Get** /api/v2/projects | Get all projects
*ProjectsAPI* | [**GetAutoTestsNamespaces**](docs/ProjectsAPI.md#getautotestsnamespaces) | **Get** /api/v2/projects/{id}/autoTestsNamespaces | Get namespaces of autotests in project
*ProjectsAPI* | [**GetProjectById**](docs/ProjectsAPI.md#getprojectbyid) | **Get** /api/v2/projects/{id} | Get project by ID
*ProjectsAPI* | [**GetTestPlansByProjectId**](docs/ProjectsAPI.md#gettestplansbyprojectid) | **Get** /api/v2/projects/{id}/testPlans | Get project test plans
*ProjectsAPI* | [**GetTestRunsByProjectId**](docs/ProjectsAPI.md#gettestrunsbyprojectid) | **Get** /api/v2/projects/{id}/testRuns | Get project test runs
*ProjectsAPI* | [**UpdateProject**](docs/ProjectsAPI.md#updateproject) | **Put** /api/v2/projects | Update project
*SearchAPI* | [**ApiV2SearchGlobalSearchPost**](docs/SearchAPI.md#apiv2searchglobalsearchpost) | **Post** /api/v2/search/globalSearch | 
*SectionsAPI* | [**ApiV2SectionsIdPatch**](docs/SectionsAPI.md#apiv2sectionsidpatch) | **Patch** /api/v2/sections/{id} | Patch section
*SectionsAPI* | [**CreateSection**](docs/SectionsAPI.md#createsection) | **Post** /api/v2/sections | Create section
*SectionsAPI* | [**DeleteSection**](docs/SectionsAPI.md#deletesection) | **Delete** /api/v2/sections/{id} | Delete section
*SectionsAPI* | [**GetSectionById**](docs/SectionsAPI.md#getsectionbyid) | **Get** /api/v2/sections/{id} | Get section
*SectionsAPI* | [**GetWorkItemsBySectionId**](docs/SectionsAPI.md#getworkitemsbysectionid) | **Get** /api/v2/sections/{id}/workItems | Get section work items
*SectionsAPI* | [**Move**](docs/SectionsAPI.md#move) | **Post** /api/v2/sections/move | Move section with all work items into another section
*SectionsAPI* | [**Rename**](docs/SectionsAPI.md#rename) | **Post** /api/v2/sections/rename | Rename section
*SectionsAPI* | [**UpdateSection**](docs/SectionsAPI.md#updatesection) | **Put** /api/v2/sections | Update section
*TagsAPI* | [**ApiV2TagsDelete**](docs/TagsAPI.md#apiv2tagsdelete) | **Delete** /api/v2/tags | Delete tags
*TagsAPI* | [**ApiV2TagsIdDelete**](docs/TagsAPI.md#apiv2tagsiddelete) | **Delete** /api/v2/tags/{id} | Delete tag
*TagsAPI* | [**ApiV2TagsPost**](docs/TagsAPI.md#apiv2tagspost) | **Post** /api/v2/tags | Create tag
*TagsAPI* | [**ApiV2TagsPut**](docs/TagsAPI.md#apiv2tagsput) | **Put** /api/v2/tags | Update tag
*TagsAPI* | [**ApiV2TagsSearchGet**](docs/TagsAPI.md#apiv2tagssearchget) | **Get** /api/v2/tags/search | Search tags
*TagsAPI* | [**ApiV2TagsTestPlansTagsGet**](docs/TagsAPI.md#apiv2tagstestplanstagsget) | **Get** /api/v2/tags/testPlansTags | Get all Tags that are used in TestPlans
*TestPlansAPI* | [**AddTestPointsWithSections**](docs/TestPlansAPI.md#addtestpointswithsections) | **Post** /api/v2/testPlans/{id}/test-points/withSections | Add test-points to TestPlan with sections
*TestPlansAPI* | [**AddWorkItemsWithSections**](docs/TestPlansAPI.md#addworkitemswithsections) | **Post** /api/v2/testPlans/{id}/workItems/withSections | Add WorkItems to TestPlan with Sections as TestSuites
*TestPlansAPI* | [**ApiV2TestPlansIdAnalyticsGet**](docs/TestPlansAPI.md#apiv2testplansidanalyticsget) | **Get** /api/v2/testPlans/{id}/analytics | Get analytics by TestPlan
*TestPlansAPI* | [**ApiV2TestPlansIdAutobalancePost**](docs/TestPlansAPI.md#apiv2testplansidautobalancepost) | **Post** /api/v2/testPlans/{id}/autobalance | Distribute test points between the users
*TestPlansAPI* | [**ApiV2TestPlansIdConfigurationsGet**](docs/TestPlansAPI.md#apiv2testplansidconfigurationsget) | **Get** /api/v2/testPlans/{id}/configurations | Get TestPlan configurations
*TestPlansAPI* | [**ApiV2TestPlansIdExportTestPointsXlsxPost**](docs/TestPlansAPI.md#apiv2testplansidexporttestpointsxlsxpost) | **Post** /api/v2/testPlans/{id}/export/testPoints/xlsx | Export TestPoints from TestPlan in xls format
*TestPlansAPI* | [**ApiV2TestPlansIdExportTestResultHistoryXlsxPost**](docs/TestPlansAPI.md#apiv2testplansidexporttestresulthistoryxlsxpost) | **Post** /api/v2/testPlans/{id}/export/testResultHistory/xlsx | Export TestResults history from TestPlan in xls format
*TestPlansAPI* | [**ApiV2TestPlansIdHistoryGet**](docs/TestPlansAPI.md#apiv2testplansidhistoryget) | **Get** /api/v2/testPlans/{id}/history | Get TestPlan history
*TestPlansAPI* | [**ApiV2TestPlansIdLinksGet**](docs/TestPlansAPI.md#apiv2testplansidlinksget) | **Get** /api/v2/testPlans/{id}/links | Get Links of TestPlan
*TestPlansAPI* | [**ApiV2TestPlansIdPatch**](docs/TestPlansAPI.md#apiv2testplansidpatch) | **Patch** /api/v2/testPlans/{id} | Patch test plan
*TestPlansAPI* | [**ApiV2TestPlansIdSummariesGet**](docs/TestPlansAPI.md#apiv2testplansidsummariesget) | **Get** /api/v2/testPlans/{id}/summaries | Get summary by TestPlan
*TestPlansAPI* | [**ApiV2TestPlansIdTestPointsLastResultsGet**](docs/TestPlansAPI.md#apiv2testplansidtestpointslastresultsget) | **Get** /api/v2/testPlans/{id}/testPoints/lastResults | Get TestPoints with last result from TestPlan
*TestPlansAPI* | [**ApiV2TestPlansIdTestPointsResetPost**](docs/TestPlansAPI.md#apiv2testplansidtestpointsresetpost) | **Post** /api/v2/testPlans/{id}/testPoints/reset | Reset TestPoints status of TestPlan
*TestPlansAPI* | [**ApiV2TestPlansIdTestPointsTesterDelete**](docs/TestPlansAPI.md#apiv2testplansidtestpointstesterdelete) | **Delete** /api/v2/testPlans/{id}/testPoints/tester | Unassign users from multiple test points
*TestPlansAPI* | [**ApiV2TestPlansIdTestPointsTesterUserIdPost**](docs/TestPlansAPI.md#apiv2testplansidtestpointstesteruseridpost) | **Post** /api/v2/testPlans/{id}/testPoints/tester/{userId} | Assign user as a tester to multiple test points
*TestPlansAPI* | [**ApiV2TestPlansIdTestRunsGet**](docs/TestPlansAPI.md#apiv2testplansidtestrunsget) | **Get** /api/v2/testPlans/{id}/testRuns | Get TestRuns of TestPlan
*TestPlansAPI* | [**ApiV2TestPlansIdTestRunsSearchPost**](docs/TestPlansAPI.md#apiv2testplansidtestrunssearchpost) | **Post** /api/v2/testPlans/{id}/testRuns/search | Search TestRuns of TestPlan
*TestPlansAPI* | [**ApiV2TestPlansIdTestRunsTestResultsLastModifiedModifiedDateGet**](docs/TestPlansAPI.md#apiv2testplansidtestrunstestresultslastmodifiedmodifieddateget) | **Get** /api/v2/testPlans/{id}/testRuns/testResults/lastModified/modifiedDate | Get last modification date of test plan&#39;s test results
*TestPlansAPI* | [**ApiV2TestPlansIdUnlockRequestPost**](docs/TestPlansAPI.md#apiv2testplansidunlockrequestpost) | **Post** /api/v2/testPlans/{id}/unlock/request | Send unlock TestPlan notification
*TestPlansAPI* | [**ApiV2TestPlansShortsPost**](docs/TestPlansAPI.md#apiv2testplansshortspost) | **Post** /api/v2/testPlans/shorts | Get TestPlans short models by Project identifiers
*TestPlansAPI* | [**Clone**](docs/TestPlansAPI.md#clone) | **Post** /api/v2/testPlans/{id}/clone | Clone TestPlan
*TestPlansAPI* | [**Complete**](docs/TestPlansAPI.md#complete) | **Post** /api/v2/testPlans/{id}/complete | Complete TestPlan
*TestPlansAPI* | [**CreateTestPlan**](docs/TestPlansAPI.md#createtestplan) | **Post** /api/v2/testPlans | Create TestPlan
*TestPlansAPI* | [**DeleteTestPlan**](docs/TestPlansAPI.md#deletetestplan) | **Delete** /api/v2/testPlans/{id} | Delete TestPlan
*TestPlansAPI* | [**GetTestPlanById**](docs/TestPlansAPI.md#gettestplanbyid) | **Get** /api/v2/testPlans/{id} | Get TestPlan by Id
*TestPlansAPI* | [**GetTestSuitesById**](docs/TestPlansAPI.md#gettestsuitesbyid) | **Get** /api/v2/testPlans/{id}/testSuites | Get TestSuites Tree By Id
*TestPlansAPI* | [**Pause**](docs/TestPlansAPI.md#pause) | **Post** /api/v2/testPlans/{id}/pause | Pause TestPlan
*TestPlansAPI* | [**PurgeTestPlan**](docs/TestPlansAPI.md#purgetestplan) | **Post** /api/v2/testPlans/{id}/purge | Permanently delete test plan from archive
*TestPlansAPI* | [**RestoreTestPlan**](docs/TestPlansAPI.md#restoretestplan) | **Post** /api/v2/testPlans/{id}/restore | Restore TestPlan
*TestPlansAPI* | [**Start**](docs/TestPlansAPI.md#start) | **Post** /api/v2/testPlans/{id}/start | Start TestPlan
*TestPlansAPI* | [**UpdateTestPlan**](docs/TestPlansAPI.md#updatetestplan) | **Put** /api/v2/testPlans | Update TestPlan
*TestPointsAPI* | [**ApiV2TestPointsIdTestRunsGet**](docs/TestPointsAPI.md#apiv2testpointsidtestrunsget) | **Get** /api/v2/testPoints/{id}/testRuns | Get all test runs which use test point
*TestPointsAPI* | [**ApiV2TestPointsIdWorkItemGet**](docs/TestPointsAPI.md#apiv2testpointsidworkitemget) | **Get** /api/v2/testPoints/{id}/workItem | Get work item represented by test point
*TestPointsAPI* | [**ApiV2TestPointsSearchIdPost**](docs/TestPointsAPI.md#apiv2testpointssearchidpost) | **Post** /api/v2/testPoints/search/id | Search for test points and extract IDs only
*TestPointsAPI* | [**ApiV2TestPointsSearchPost**](docs/TestPointsAPI.md#apiv2testpointssearchpost) | **Post** /api/v2/testPoints/search | Search for test points
*TestResultsAPI* | [**ApiV2TestResultsExternalProjectsExternalProjectIdDefectsExternalFormsPost**](docs/TestResultsAPI.md#apiv2testresultsexternalprojectsexternalprojectiddefectsexternalformspost) | **Post** /api/v2/testResults/external-projects/{externalProjectId}/defects/external-forms | 
*TestResultsAPI* | [**ApiV2TestResultsExternalProjectsExternalProjectIdDefectsPost**](docs/TestResultsAPI.md#apiv2testresultsexternalprojectsexternalprojectiddefectspost) | **Post** /api/v2/testResults/external-projects/{externalProjectId}/defects | 
*TestResultsAPI* | [**ApiV2TestResultsIdAggregatedGet**](docs/TestResultsAPI.md#apiv2testresultsidaggregatedget) | **Get** /api/v2/testResults/{id}/aggregated | Get test result by ID aggregated with previous results
*TestResultsAPI* | [**ApiV2TestResultsIdAttachmentsAttachmentIdPut**](docs/TestResultsAPI.md#apiv2testresultsidattachmentsattachmentidput) | **Put** /api/v2/testResults/{id}/attachments/{attachmentId} | Attach file to the test result
*TestResultsAPI* | [**ApiV2TestResultsIdAttachmentsInfoGet**](docs/TestResultsAPI.md#apiv2testresultsidattachmentsinfoget) | **Get** /api/v2/testResults/{id}/attachments/info | Get test result attachments meta-information
*TestResultsAPI* | [**ApiV2TestResultsIdGet**](docs/TestResultsAPI.md#apiv2testresultsidget) | **Get** /api/v2/testResults/{id} | Get test result by ID
*TestResultsAPI* | [**ApiV2TestResultsIdPut**](docs/TestResultsAPI.md#apiv2testresultsidput) | **Put** /api/v2/testResults/{id} | Edit test result by ID
*TestResultsAPI* | [**ApiV2TestResultsIdRerunsGet**](docs/TestResultsAPI.md#apiv2testresultsidrerunsget) | **Get** /api/v2/testResults/{id}/reruns | Get reruns
*TestResultsAPI* | [**ApiV2TestResultsSearchPost**](docs/TestResultsAPI.md#apiv2testresultssearchpost) | **Post** /api/v2/testResults/search | Search for test results
*TestResultsAPI* | [**ApiV2TestResultsStatisticsFilterPost**](docs/TestResultsAPI.md#apiv2testresultsstatisticsfilterpost) | **Post** /api/v2/testResults/statistics/filter | Search for test results and extract statistics
*TestResultsAPI* | [**CreateAttachment**](docs/TestResultsAPI.md#createattachment) | **Post** /api/v2/testResults/{id}/attachments | Upload and link attachment to TestResult
*TestResultsAPI* | [**DeleteAttachment**](docs/TestResultsAPI.md#deleteattachment) | **Delete** /api/v2/testResults/{id}/attachments/{attachmentId} | Remove attachment and unlink from TestResult
*TestResultsAPI* | [**DownloadAttachment**](docs/TestResultsAPI.md#downloadattachment) | **Get** /api/v2/testResults/{id}/attachments/{attachmentId} | Get attachment of TestResult
*TestResultsAPI* | [**GetAttachment**](docs/TestResultsAPI.md#getattachment) | **Get** /api/v2/testResults/{id}/attachments/{attachmentId}/info | Get Metadata of TestResult&#39;s attachment
*TestResultsAPI* | [**GetAttachments**](docs/TestResultsAPI.md#getattachments) | **Get** /api/v2/testResults/{id}/attachments | Get all attachments of TestResult
*TestRunsAPI* | [**ApiV2TestRunsDelete**](docs/TestRunsAPI.md#apiv2testrunsdelete) | **Delete** /api/v2/testRuns | Delete multiple test runs
*TestRunsAPI* | [**ApiV2TestRunsIdAutoTestsNamespacesGet**](docs/TestRunsAPI.md#apiv2testrunsidautotestsnamespacesget) | **Get** /api/v2/testRuns/{id}/autoTestsNamespaces | Get autotest classes and namespaces in test run
*TestRunsAPI* | [**ApiV2TestRunsIdDelete**](docs/TestRunsAPI.md#apiv2testrunsiddelete) | **Delete** /api/v2/testRuns/{id} | Delete test run
*TestRunsAPI* | [**ApiV2TestRunsIdPurgePost**](docs/TestRunsAPI.md#apiv2testrunsidpurgepost) | **Post** /api/v2/testRuns/{id}/purge | Permanently delete test run from archive
*TestRunsAPI* | [**ApiV2TestRunsIdRerunsPost**](docs/TestRunsAPI.md#apiv2testrunsidrerunspost) | **Post** /api/v2/testRuns/{id}/reruns | Manual autotests rerun in test run
*TestRunsAPI* | [**ApiV2TestRunsIdRestorePost**](docs/TestRunsAPI.md#apiv2testrunsidrestorepost) | **Post** /api/v2/testRuns/{id}/restore | Restore test run from the archive
*TestRunsAPI* | [**ApiV2TestRunsIdStatisticsFilterPost**](docs/TestRunsAPI.md#apiv2testrunsidstatisticsfilterpost) | **Post** /api/v2/testRuns/{id}/statistics/filter | Search for the test run test results and build statistics
*TestRunsAPI* | [**ApiV2TestRunsIdTestPointsResultsGet**](docs/TestRunsAPI.md#apiv2testrunsidtestpointsresultsget) | **Get** /api/v2/testRuns/{id}/testPoints/results | Get test results from the test run grouped by test points
*TestRunsAPI* | [**ApiV2TestRunsIdTestResultsBulkPut**](docs/TestRunsAPI.md#apiv2testrunsidtestresultsbulkput) | **Put** /api/v2/testRuns/{id}/testResults/bulk | Partial edit of multiple test results in the test run
*TestRunsAPI* | [**ApiV2TestRunsIdTestResultsLastModifiedModificationDateGet**](docs/TestRunsAPI.md#apiv2testrunsidtestresultslastmodifiedmodificationdateget) | **Get** /api/v2/testRuns/{id}/testResults/lastModified/modificationDate | Get modification date of last test result of the test run
*TestRunsAPI* | [**ApiV2TestRunsPurgeBulkPost**](docs/TestRunsAPI.md#apiv2testrunspurgebulkpost) | **Post** /api/v2/testRuns/purge/bulk | Permanently delete multiple test runs from archive
*TestRunsAPI* | [**ApiV2TestRunsRestoreBulkPost**](docs/TestRunsAPI.md#apiv2testrunsrestorebulkpost) | **Post** /api/v2/testRuns/restore/bulk | Restore multiple test runs from the archive
*TestRunsAPI* | [**ApiV2TestRunsSearchPost**](docs/TestRunsAPI.md#apiv2testrunssearchpost) | **Post** /api/v2/testRuns/search | Search for test runs
*TestRunsAPI* | [**ApiV2TestRunsUpdateMultiplePost**](docs/TestRunsAPI.md#apiv2testrunsupdatemultiplepost) | **Post** /api/v2/testRuns/updateMultiple | Update multiple test runs
*TestRunsAPI* | [**CompleteTestRun**](docs/TestRunsAPI.md#completetestrun) | **Post** /api/v2/testRuns/{id}/complete | Complete TestRun
*TestRunsAPI* | [**CreateAndFillByAutoTests**](docs/TestRunsAPI.md#createandfillbyautotests) | **Post** /api/v2/testRuns/byAutoTests | Create test runs based on autotests and configurations
*TestRunsAPI* | [**CreateAndFillByConfigurations**](docs/TestRunsAPI.md#createandfillbyconfigurations) | **Post** /api/v2/testRuns/byConfigurations | Create test runs picking the needed test points
*TestRunsAPI* | [**CreateAndFillByWorkItems**](docs/TestRunsAPI.md#createandfillbyworkitems) | **Post** /api/v2/testRuns/byWorkItems | Create test run based on configurations and work items
*TestRunsAPI* | [**CreateEmpty**](docs/TestRunsAPI.md#createempty) | **Post** /api/v2/testRuns | Create empty TestRun
*TestRunsAPI* | [**GetTestRunById**](docs/TestRunsAPI.md#gettestrunbyid) | **Get** /api/v2/testRuns/{id} | Get TestRun by Id
*TestRunsAPI* | [**SetAutoTestResultsForTestRun**](docs/TestRunsAPI.md#setautotestresultsfortestrun) | **Post** /api/v2/testRuns/{id}/testResults | Send test results to the test runs in the system
*TestRunsAPI* | [**StartTestRun**](docs/TestRunsAPI.md#starttestrun) | **Post** /api/v2/testRuns/{id}/start | Start TestRun
*TestRunsAPI* | [**StopTestRun**](docs/TestRunsAPI.md#stoptestrun) | **Post** /api/v2/testRuns/{id}/stop | Stop TestRun
*TestRunsAPI* | [**UpdateEmpty**](docs/TestRunsAPI.md#updateempty) | **Put** /api/v2/testRuns | Update empty TestRun
*TestSuitesAPI* | [**AddTestPointsToTestSuite**](docs/TestSuitesAPI.md#addtestpointstotestsuite) | **Post** /api/v2/testSuites/{id}/test-points | Add test-points to test suite
*TestSuitesAPI* | [**ApiV2TestSuitesIdPatch**](docs/TestSuitesAPI.md#apiv2testsuitesidpatch) | **Patch** /api/v2/testSuites/{id} | Patch test suite
*TestSuitesAPI* | [**ApiV2TestSuitesIdRefreshPost**](docs/TestSuitesAPI.md#apiv2testsuitesidrefreshpost) | **Post** /api/v2/testSuites/{id}/refresh | Refresh test suite. Only dynamic test suites are supported by this method
*TestSuitesAPI* | [**ApiV2TestSuitesIdWorkItemsPost**](docs/TestSuitesAPI.md#apiv2testsuitesidworkitemspost) | **Post** /api/v2/testSuites/{id}/workItems | Set work items for test suite
*TestSuitesAPI* | [**ApiV2TestSuitesPost**](docs/TestSuitesAPI.md#apiv2testsuitespost) | **Post** /api/v2/testSuites | Create test suite
*TestSuitesAPI* | [**ApiV2TestSuitesPut**](docs/TestSuitesAPI.md#apiv2testsuitesput) | **Put** /api/v2/testSuites | Edit test suite
*TestSuitesAPI* | [**DeleteTestSuite**](docs/TestSuitesAPI.md#deletetestsuite) | **Delete** /api/v2/testSuites/{id} | Delete TestSuite
*TestSuitesAPI* | [**GetConfigurationsByTestSuiteId**](docs/TestSuitesAPI.md#getconfigurationsbytestsuiteid) | **Get** /api/v2/testSuites/{id}/configurations | Get Configurations By Id
*TestSuitesAPI* | [**GetTestPointsById**](docs/TestSuitesAPI.md#gettestpointsbyid) | **Get** /api/v2/testSuites/{id}/testPoints | Get TestPoints By Id
*TestSuitesAPI* | [**GetTestResultsById**](docs/TestSuitesAPI.md#gettestresultsbyid) | **Get** /api/v2/testSuites/{id}/testResults | Get TestResults By Id
*TestSuitesAPI* | [**GetTestSuiteById**](docs/TestSuitesAPI.md#gettestsuitebyid) | **Get** /api/v2/testSuites/{id} | Get TestSuite by Id
*TestSuitesAPI* | [**SearchWorkItems**](docs/TestSuitesAPI.md#searchworkitems) | **Post** /api/v2/testSuites/{id}/workItems/search | Search WorkItems
*TestSuitesAPI* | [**SetConfigurationsByTestSuiteId**](docs/TestSuitesAPI.md#setconfigurationsbytestsuiteid) | **Post** /api/v2/testSuites/{id}/configurations | Set Configurations By TestSuite Id
*UsersAPI* | [**ApiV2UsersExistsGet**](docs/UsersAPI.md#apiv2usersexistsget) | **Get** /api/v2/users/exists | 
*WebhooksAPI* | [**ApiV2WebhooksDelete**](docs/WebhooksAPI.md#apiv2webhooksdelete) | **Delete** /api/v2/webhooks | 
*WebhooksAPI* | [**ApiV2WebhooksGet**](docs/WebhooksAPI.md#apiv2webhooksget) | **Get** /api/v2/webhooks | Get all webhooks
*WebhooksAPI* | [**ApiV2WebhooksIdDelete**](docs/WebhooksAPI.md#apiv2webhooksiddelete) | **Delete** /api/v2/webhooks/{id} | Delete webhook by ID
*WebhooksAPI* | [**ApiV2WebhooksIdGet**](docs/WebhooksAPI.md#apiv2webhooksidget) | **Get** /api/v2/webhooks/{id} | Get webhook by ID
*WebhooksAPI* | [**ApiV2WebhooksIdPut**](docs/WebhooksAPI.md#apiv2webhooksidput) | **Put** /api/v2/webhooks/{id} | Edit webhook by ID
*WebhooksAPI* | [**ApiV2WebhooksPost**](docs/WebhooksAPI.md#apiv2webhookspost) | **Post** /api/v2/webhooks | Create webhook
*WebhooksAPI* | [**ApiV2WebhooksPut**](docs/WebhooksAPI.md#apiv2webhooksput) | **Put** /api/v2/webhooks | 
*WebhooksAPI* | [**ApiV2WebhooksSearchPost**](docs/WebhooksAPI.md#apiv2webhookssearchpost) | **Post** /api/v2/webhooks/search | Search for webhooks
*WebhooksAPI* | [**ApiV2WebhooksSpecialVariablesGet**](docs/WebhooksAPI.md#apiv2webhooksspecialvariablesget) | **Get** /api/v2/webhooks/specialVariables | Get special variables for webhook event type
*WebhooksAPI* | [**ApiV2WebhooksTestPost**](docs/WebhooksAPI.md#apiv2webhookstestpost) | **Post** /api/v2/webhooks/test | Test webhook&#39;s url
*WebhooksLogsAPI* | [**ApiV2WebhooksLogsGet**](docs/WebhooksLogsAPI.md#apiv2webhookslogsget) | **Get** /api/v2/webhooks/logs | Get all webhook logs
*WebhooksLogsAPI* | [**ApiV2WebhooksLogsIdDelete**](docs/WebhooksLogsAPI.md#apiv2webhookslogsiddelete) | **Delete** /api/v2/webhooks/logs/{id} | Delete webhook log by ID
*WebhooksLogsAPI* | [**ApiV2WebhooksLogsIdGet**](docs/WebhooksLogsAPI.md#apiv2webhookslogsidget) | **Get** /api/v2/webhooks/logs/{id} | Get webhook log by ID
*WorkItemsAPI* | [**ApiV2WorkItemsIdAttachmentsPost**](docs/WorkItemsAPI.md#apiv2workitemsidattachmentspost) | **Post** /api/v2/workItems/{id}/attachments | Upload and link attachment to WorkItem
*WorkItemsAPI* | [**ApiV2WorkItemsIdCheckListTransformToTestCasePost**](docs/WorkItemsAPI.md#apiv2workitemsidchecklisttransformtotestcasepost) | **Post** /api/v2/workItems/{id}/checkList/transformTo/testCase | Transform CheckList to TestCase
*WorkItemsAPI* | [**ApiV2WorkItemsIdHistoryGet**](docs/WorkItemsAPI.md#apiv2workitemsidhistoryget) | **Get** /api/v2/workItems/{id}/history | Get change history of WorkItem
*WorkItemsAPI* | [**ApiV2WorkItemsIdLikeDelete**](docs/WorkItemsAPI.md#apiv2workitemsidlikedelete) | **Delete** /api/v2/workItems/{id}/like | Delete like from WorkItem
*WorkItemsAPI* | [**ApiV2WorkItemsIdLikePost**](docs/WorkItemsAPI.md#apiv2workitemsidlikepost) | **Post** /api/v2/workItems/{id}/like | Set like to WorkItem
*WorkItemsAPI* | [**ApiV2WorkItemsIdLikesCountGet**](docs/WorkItemsAPI.md#apiv2workitemsidlikescountget) | **Get** /api/v2/workItems/{id}/likes/count | Get likes count of WorkItem
*WorkItemsAPI* | [**ApiV2WorkItemsIdLikesGet**](docs/WorkItemsAPI.md#apiv2workitemsidlikesget) | **Get** /api/v2/workItems/{id}/likes | Get likes of WorkItem
*WorkItemsAPI* | [**ApiV2WorkItemsIdTestResultsHistoryGet**](docs/WorkItemsAPI.md#apiv2workitemsidtestresultshistoryget) | **Get** /api/v2/workItems/{id}/testResults/history | Get test results history of WorkItem
*WorkItemsAPI* | [**ApiV2WorkItemsIdVersionVersionIdActualPost**](docs/WorkItemsAPI.md#apiv2workitemsidversionversionidactualpost) | **Post** /api/v2/workItems/{id}/version/{versionId}/actual | Set WorkItem as actual
*WorkItemsAPI* | [**ApiV2WorkItemsLinksUrlsSearchPost**](docs/WorkItemsAPI.md#apiv2workitemslinksurlssearchpost) | **Post** /api/v2/workItems/links/urls/search | 
*WorkItemsAPI* | [**ApiV2WorkItemsMovePost**](docs/WorkItemsAPI.md#apiv2workitemsmovepost) | **Post** /api/v2/workItems/move | Move WorkItem to another section
*WorkItemsAPI* | [**ApiV2WorkItemsSearchPost**](docs/WorkItemsAPI.md#apiv2workitemssearchpost) | **Post** /api/v2/workItems/search | Search for work items
*WorkItemsAPI* | [**ApiV2WorkItemsSharedStepIdReferencesSectionsPost**](docs/WorkItemsAPI.md#apiv2workitemssharedstepidreferencessectionspost) | **Post** /api/v2/workItems/{sharedStepId}/references/sections | Get SharedStep references in sections
*WorkItemsAPI* | [**ApiV2WorkItemsSharedStepIdReferencesWorkItemsPost**](docs/WorkItemsAPI.md#apiv2workitemssharedstepidreferencesworkitemspost) | **Post** /api/v2/workItems/{sharedStepId}/references/workItems | Get SharedStep references in work items
*WorkItemsAPI* | [**ApiV2WorkItemsSharedStepsSharedStepIdReferencesGet**](docs/WorkItemsAPI.md#apiv2workitemssharedstepssharedstepidreferencesget) | **Get** /api/v2/workItems/sharedSteps/{sharedStepId}/references | Get SharedStep references
*WorkItemsAPI* | [**CreateWorkItem**](docs/WorkItemsAPI.md#createworkitem) | **Post** /api/v2/workItems | Create Test Case, Checklist or Shared Step
*WorkItemsAPI* | [**DeleteAllWorkItemsFromAutoTest**](docs/WorkItemsAPI.md#deleteallworkitemsfromautotest) | **Delete** /api/v2/workItems/{id}/autoTests | Delete all links AutoTests from WorkItem by Id or GlobalId
*WorkItemsAPI* | [**DeleteWorkItem**](docs/WorkItemsAPI.md#deleteworkitem) | **Delete** /api/v2/workItems/{id} | Delete Test Case, Checklist or Shared Step by Id or GlobalId
*WorkItemsAPI* | [**GetAutoTestsForWorkItem**](docs/WorkItemsAPI.md#getautotestsforworkitem) | **Get** /api/v2/workItems/{id}/autoTests | Get all AutoTests linked to WorkItem by Id or GlobalId
*WorkItemsAPI* | [**GetIterations**](docs/WorkItemsAPI.md#getiterations) | **Get** /api/v2/workItems/{id}/iterations | Get iterations by work item Id or GlobalId
*WorkItemsAPI* | [**GetWorkItemById**](docs/WorkItemsAPI.md#getworkitembyid) | **Get** /api/v2/workItems/{id} | Get Test Case, Checklist or Shared Step by Id or GlobalId
*WorkItemsAPI* | [**GetWorkItemChronology**](docs/WorkItemsAPI.md#getworkitemchronology) | **Get** /api/v2/workItems/{id}/chronology | Get WorkItem chronology by Id or GlobalId
*WorkItemsAPI* | [**GetWorkItemVersions**](docs/WorkItemsAPI.md#getworkitemversions) | **Get** /api/v2/workItems/{id}/versions | Get WorkItem versions
*WorkItemsAPI* | [**PurgeWorkItem**](docs/WorkItemsAPI.md#purgeworkitem) | **Post** /api/v2/workItems/{id}/purge | Permanently delete test case, checklist or shared steps from archive
*WorkItemsAPI* | [**RestoreWorkItem**](docs/WorkItemsAPI.md#restoreworkitem) | **Post** /api/v2/workItems/{id}/restore | Restore test case, checklist or shared steps from archive
*WorkItemsAPI* | [**UpdateWorkItem**](docs/WorkItemsAPI.md#updateworkitem) | **Put** /api/v2/workItems | Update Test Case, Checklist or Shared Step
*WorkItemsCommentsAPI* | [**ApiV2WorkItemsCommentsCommentIdDelete**](docs/WorkItemsCommentsAPI.md#apiv2workitemscommentscommentiddelete) | **Delete** /api/v2/workItems/comments/{commentId} | Delete WorkItem comment
*WorkItemsCommentsAPI* | [**ApiV2WorkItemsCommentsPost**](docs/WorkItemsCommentsAPI.md#apiv2workitemscommentspost) | **Post** /api/v2/workItems/comments | Create WorkItem comment
*WorkItemsCommentsAPI* | [**ApiV2WorkItemsCommentsPut**](docs/WorkItemsCommentsAPI.md#apiv2workitemscommentsput) | **Put** /api/v2/workItems/comments | Update work item comment
*WorkItemsCommentsAPI* | [**ApiV2WorkItemsIdCommentsCountGet**](docs/WorkItemsCommentsAPI.md#apiv2workitemsidcommentscountget) | **Get** /api/v2/workItems/{id}/comments/count | Get work item comments count
*WorkItemsCommentsAPI* | [**ApiV2WorkItemsIdCommentsGet**](docs/WorkItemsCommentsAPI.md#apiv2workitemsidcommentsget) | **Get** /api/v2/workItems/{id}/comments | Get work item comments


## Documentation for Models

You can see the documentation [here](https://github.com/testit-tms/api-client-golang/blob/main/docs/Readme.md)

# Contributing

You can help to develop the project. Any contributions are **greatly appreciated**.

* If you have suggestions for adding or removing projects, feel free to [open an issue](https://github.com/testit-tms/api-client-golang/issues/new) to discuss it, or directly create a pull request after you edit the *README.md* file with necessary changes.
* Please make sure you check your spelling and grammar.
* Create individual PR for each suggestion.
* Please also read through the [Code Of Conduct](https://github.com/testit-tms/api-client-golang/blob/main/CODE_OF_CONDUCT.md) before posting your first idea as well.

# License

Distributed under the Apache-2.0 License. See [LICENSE](https://github.com/testit-tms/api-client-golang/blob/main/LICENSE.md) for more information.
