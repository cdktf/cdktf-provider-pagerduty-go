// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package incidentworkflowtrigger


type IncidentWorkflowTriggerPermissions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.24.2/docs/resources/incident_workflow_trigger#restricted IncidentWorkflowTrigger#restricted}.
	Restricted interface{} `field:"optional" json:"restricted" yaml:"restricted"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.24.2/docs/resources/incident_workflow_trigger#team_id IncidentWorkflowTrigger#team_id}.
	TeamId *string `field:"optional" json:"teamId" yaml:"teamId"`
}

