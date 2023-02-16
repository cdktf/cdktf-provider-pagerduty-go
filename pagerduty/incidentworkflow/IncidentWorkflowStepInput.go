package incidentworkflow


type IncidentWorkflowStepInput struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/pagerduty/r/incident_workflow#name IncidentWorkflow#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/pagerduty/r/incident_workflow#value IncidentWorkflow#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

