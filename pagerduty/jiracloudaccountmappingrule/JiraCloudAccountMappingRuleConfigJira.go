// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package jiracloudaccountmappingrule


type JiraCloudAccountMappingRuleConfigJira struct {
	// issue_type block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.18.2/docs/resources/jira_cloud_account_mapping_rule#issue_type JiraCloudAccountMappingRule#issue_type}
	IssueType *JiraCloudAccountMappingRuleConfigJiraIssueType `field:"required" json:"issueType" yaml:"issueType"`
	// project block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.18.2/docs/resources/jira_cloud_account_mapping_rule#project JiraCloudAccountMappingRule#project}
	Project *JiraCloudAccountMappingRuleConfigJiraProject `field:"required" json:"project" yaml:"project"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.18.2/docs/resources/jira_cloud_account_mapping_rule#autocreate_jql JiraCloudAccountMappingRule#autocreate_jql}.
	AutocreateJql *string `field:"optional" json:"autocreateJql" yaml:"autocreateJql"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.18.2/docs/resources/jira_cloud_account_mapping_rule#create_issue_on_incident_trigger JiraCloudAccountMappingRule#create_issue_on_incident_trigger}.
	CreateIssueOnIncidentTrigger interface{} `field:"optional" json:"createIssueOnIncidentTrigger" yaml:"createIssueOnIncidentTrigger"`
	// custom_fields block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.18.2/docs/resources/jira_cloud_account_mapping_rule#custom_fields JiraCloudAccountMappingRule#custom_fields}
	CustomFields interface{} `field:"optional" json:"customFields" yaml:"customFields"`
	// priorities block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.18.2/docs/resources/jira_cloud_account_mapping_rule#priorities JiraCloudAccountMappingRule#priorities}
	Priorities interface{} `field:"optional" json:"priorities" yaml:"priorities"`
	// status_mapping block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.18.2/docs/resources/jira_cloud_account_mapping_rule#status_mapping JiraCloudAccountMappingRule#status_mapping}
	StatusMapping *JiraCloudAccountMappingRuleConfigJiraStatusMapping `field:"optional" json:"statusMapping" yaml:"statusMapping"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.18.2/docs/resources/jira_cloud_account_mapping_rule#sync_notes_user JiraCloudAccountMappingRule#sync_notes_user}.
	SyncNotesUser *string `field:"optional" json:"syncNotesUser" yaml:"syncNotesUser"`
}

