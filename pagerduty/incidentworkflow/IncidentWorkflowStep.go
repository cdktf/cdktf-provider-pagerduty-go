// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package incidentworkflow


type IncidentWorkflowStep struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.28.2/docs/resources/incident_workflow#action IncidentWorkflow#action}.
	Action *string `field:"required" json:"action" yaml:"action"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.28.2/docs/resources/incident_workflow#name IncidentWorkflow#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// inline_steps_input block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.28.2/docs/resources/incident_workflow#inline_steps_input IncidentWorkflow#inline_steps_input}
	InlineStepsInput interface{} `field:"optional" json:"inlineStepsInput" yaml:"inlineStepsInput"`
	// input block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.28.2/docs/resources/incident_workflow#input IncidentWorkflow#input}
	Input interface{} `field:"optional" json:"input" yaml:"input"`
}

