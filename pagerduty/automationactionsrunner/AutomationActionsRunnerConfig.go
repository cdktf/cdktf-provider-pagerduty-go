// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package automationactionsrunner

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AutomationActionsRunnerConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.27.1/docs/resources/automation_actions_runner#name AutomationActionsRunner#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.27.1/docs/resources/automation_actions_runner#runner_type AutomationActionsRunner#runner_type}.
	RunnerType *string `field:"required" json:"runnerType" yaml:"runnerType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.27.1/docs/resources/automation_actions_runner#description AutomationActionsRunner#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.27.1/docs/resources/automation_actions_runner#id AutomationActionsRunner#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.27.1/docs/resources/automation_actions_runner#last_seen AutomationActionsRunner#last_seen}.
	LastSeen *string `field:"optional" json:"lastSeen" yaml:"lastSeen"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.27.1/docs/resources/automation_actions_runner#runbook_api_key AutomationActionsRunner#runbook_api_key}.
	RunbookApiKey *string `field:"optional" json:"runbookApiKey" yaml:"runbookApiKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.27.1/docs/resources/automation_actions_runner#runbook_base_uri AutomationActionsRunner#runbook_base_uri}.
	RunbookBaseUri *string `field:"optional" json:"runbookBaseUri" yaml:"runbookBaseUri"`
}

