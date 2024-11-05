// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package jiracloudaccountmappingrule


type JiraCloudAccountMappingRuleConfigA struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.17.0/docs/resources/jira_cloud_account_mapping_rule#service JiraCloudAccountMappingRule#service}.
	Service *string `field:"required" json:"service" yaml:"service"`
	// jira block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.17.0/docs/resources/jira_cloud_account_mapping_rule#jira JiraCloudAccountMappingRule#jira}
	Jira *JiraCloudAccountMappingRuleConfigJira `field:"optional" json:"jira" yaml:"jira"`
}

