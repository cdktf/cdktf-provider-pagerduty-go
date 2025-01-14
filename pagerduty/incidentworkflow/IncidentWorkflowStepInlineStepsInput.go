// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package incidentworkflow


type IncidentWorkflowStepInlineStepsInput struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.19.1/docs/resources/incident_workflow#name IncidentWorkflow#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// step block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.19.1/docs/resources/incident_workflow#step IncidentWorkflow#step}
	Step interface{} `field:"optional" json:"step" yaml:"step"`
}

