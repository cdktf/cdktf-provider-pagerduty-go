package incidentworkflow


type IncidentWorkflowStep struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/pagerduty/r/incident_workflow#action IncidentWorkflow#action}.
	Action *string `field:"required" json:"action" yaml:"action"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/pagerduty/r/incident_workflow#name IncidentWorkflow#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// input block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/pagerduty/r/incident_workflow#input IncidentWorkflow#input}
	Input interface{} `field:"optional" json:"input" yaml:"input"`
}

