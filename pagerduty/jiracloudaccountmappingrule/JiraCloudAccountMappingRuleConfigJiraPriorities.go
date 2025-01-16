// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package jiracloudaccountmappingrule


type JiraCloudAccountMappingRuleConfigJiraPriorities struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.19.2/docs/resources/jira_cloud_account_mapping_rule#jira_id JiraCloudAccountMappingRule#jira_id}.
	JiraId *string `field:"required" json:"jiraId" yaml:"jiraId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.19.2/docs/resources/jira_cloud_account_mapping_rule#pagerduty_id JiraCloudAccountMappingRule#pagerduty_id}.
	PagerdutyId *string `field:"required" json:"pagerdutyId" yaml:"pagerdutyId"`
}

