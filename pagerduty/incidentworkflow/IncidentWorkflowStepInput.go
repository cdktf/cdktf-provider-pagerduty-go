package incidentworkflow


type IncidentWorkflowStepInput struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/2.14.6/docs/resources/incident_workflow#name IncidentWorkflow#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/2.14.6/docs/resources/incident_workflow#value IncidentWorkflow#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

