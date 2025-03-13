// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package incidentworkflow


type IncidentWorkflowStepInlineStepsInputStepInput struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.21.1/docs/resources/incident_workflow#name IncidentWorkflow#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.21.1/docs/resources/incident_workflow#value IncidentWorkflow#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

