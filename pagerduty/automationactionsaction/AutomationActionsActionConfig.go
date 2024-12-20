// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package automationactionsaction

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AutomationActionsActionConfig struct {
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
	// action_data_reference block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.18.2/docs/resources/automation_actions_action#action_data_reference AutomationActionsAction#action_data_reference}
	ActionDataReference *AutomationActionsActionActionDataReference `field:"required" json:"actionDataReference" yaml:"actionDataReference"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.18.2/docs/resources/automation_actions_action#action_type AutomationActionsAction#action_type}.
	ActionType *string `field:"required" json:"actionType" yaml:"actionType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.18.2/docs/resources/automation_actions_action#name AutomationActionsAction#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.18.2/docs/resources/automation_actions_action#action_classification AutomationActionsAction#action_classification}.
	ActionClassification *string `field:"optional" json:"actionClassification" yaml:"actionClassification"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.18.2/docs/resources/automation_actions_action#creation_time AutomationActionsAction#creation_time}.
	CreationTime *string `field:"optional" json:"creationTime" yaml:"creationTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.18.2/docs/resources/automation_actions_action#description AutomationActionsAction#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.18.2/docs/resources/automation_actions_action#id AutomationActionsAction#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.18.2/docs/resources/automation_actions_action#modify_time AutomationActionsAction#modify_time}.
	ModifyTime *string `field:"optional" json:"modifyTime" yaml:"modifyTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.18.2/docs/resources/automation_actions_action#only_invocable_on_unresolved_incidents AutomationActionsAction#only_invocable_on_unresolved_incidents}.
	OnlyInvocableOnUnresolvedIncidents interface{} `field:"optional" json:"onlyInvocableOnUnresolvedIncidents" yaml:"onlyInvocableOnUnresolvedIncidents"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.18.2/docs/resources/automation_actions_action#runner_id AutomationActionsAction#runner_id}.
	RunnerId *string `field:"optional" json:"runnerId" yaml:"runnerId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.18.2/docs/resources/automation_actions_action#runner_type AutomationActionsAction#runner_type}.
	RunnerType *string `field:"optional" json:"runnerType" yaml:"runnerType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.18.2/docs/resources/automation_actions_action#type AutomationActionsAction#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

