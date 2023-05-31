package incidentworkflow


type IncidentWorkflowStep struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/2.15.0/docs/resources/incident_workflow#action IncidentWorkflow#action}.
	Action *string `field:"required" json:"action" yaml:"action"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/2.15.0/docs/resources/incident_workflow#name IncidentWorkflow#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// input block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/2.15.0/docs/resources/incident_workflow#input IncidentWorkflow#input}
	Input interface{} `field:"optional" json:"input" yaml:"input"`
}

