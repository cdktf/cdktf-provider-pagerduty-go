// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package automationactionsaction


type AutomationActionsActionActionDataReference struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.15.3/docs/resources/automation_actions_action#invocation_command AutomationActionsAction#invocation_command}.
	InvocationCommand *string `field:"optional" json:"invocationCommand" yaml:"invocationCommand"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.15.3/docs/resources/automation_actions_action#process_automation_job_arguments AutomationActionsAction#process_automation_job_arguments}.
	ProcessAutomationJobArguments *string `field:"optional" json:"processAutomationJobArguments" yaml:"processAutomationJobArguments"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.15.3/docs/resources/automation_actions_action#process_automation_job_id AutomationActionsAction#process_automation_job_id}.
	ProcessAutomationJobId *string `field:"optional" json:"processAutomationJobId" yaml:"processAutomationJobId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.15.3/docs/resources/automation_actions_action#process_automation_node_filter AutomationActionsAction#process_automation_node_filter}.
	ProcessAutomationNodeFilter *string `field:"optional" json:"processAutomationNodeFilter" yaml:"processAutomationNodeFilter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.15.3/docs/resources/automation_actions_action#script AutomationActionsAction#script}.
	Script *string `field:"optional" json:"script" yaml:"script"`
}

