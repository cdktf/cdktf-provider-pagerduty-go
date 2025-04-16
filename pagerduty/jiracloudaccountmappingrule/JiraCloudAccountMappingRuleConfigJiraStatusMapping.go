// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package jiracloudaccountmappingrule


type JiraCloudAccountMappingRuleConfigJiraStatusMapping struct {
	// triggered block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.24.0/docs/resources/jira_cloud_account_mapping_rule#triggered JiraCloudAccountMappingRule#triggered}
	Triggered *JiraCloudAccountMappingRuleConfigJiraStatusMappingTriggered `field:"required" json:"triggered" yaml:"triggered"`
	// acknowledged block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.24.0/docs/resources/jira_cloud_account_mapping_rule#acknowledged JiraCloudAccountMappingRule#acknowledged}
	Acknowledged *JiraCloudAccountMappingRuleConfigJiraStatusMappingAcknowledged `field:"optional" json:"acknowledged" yaml:"acknowledged"`
	// resolved block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.24.0/docs/resources/jira_cloud_account_mapping_rule#resolved JiraCloudAccountMappingRule#resolved}
	Resolved *JiraCloudAccountMappingRuleConfigJiraStatusMappingResolved `field:"optional" json:"resolved" yaml:"resolved"`
}

