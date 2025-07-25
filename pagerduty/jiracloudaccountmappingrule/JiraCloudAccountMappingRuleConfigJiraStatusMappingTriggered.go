// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package jiracloudaccountmappingrule


type JiraCloudAccountMappingRuleConfigJiraStatusMappingTriggered struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.27.2/docs/resources/jira_cloud_account_mapping_rule#id JiraCloudAccountMappingRule#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.27.2/docs/resources/jira_cloud_account_mapping_rule#name JiraCloudAccountMappingRule#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
}

