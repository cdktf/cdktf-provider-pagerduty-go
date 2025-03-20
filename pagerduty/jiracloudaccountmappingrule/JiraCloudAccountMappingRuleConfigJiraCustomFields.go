// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package jiracloudaccountmappingrule


type JiraCloudAccountMappingRuleConfigJiraCustomFields struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.23.0/docs/resources/jira_cloud_account_mapping_rule#target_issue_field JiraCloudAccountMappingRule#target_issue_field}.
	TargetIssueField *string `field:"required" json:"targetIssueField" yaml:"targetIssueField"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.23.0/docs/resources/jira_cloud_account_mapping_rule#target_issue_field_name JiraCloudAccountMappingRule#target_issue_field_name}.
	TargetIssueFieldName *string `field:"required" json:"targetIssueFieldName" yaml:"targetIssueFieldName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.23.0/docs/resources/jira_cloud_account_mapping_rule#type JiraCloudAccountMappingRule#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.23.0/docs/resources/jira_cloud_account_mapping_rule#source_incident_field JiraCloudAccountMappingRule#source_incident_field}.
	SourceIncidentField *string `field:"optional" json:"sourceIncidentField" yaml:"sourceIncidentField"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.23.0/docs/resources/jira_cloud_account_mapping_rule#value JiraCloudAccountMappingRule#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

