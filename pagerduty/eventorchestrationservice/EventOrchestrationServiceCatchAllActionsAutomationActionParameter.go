package eventorchestrationservice


type EventOrchestrationServiceCatchAllActionsAutomationActionParameter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/2.14.3/docs/resources/event_orchestration_service#key EventOrchestrationService#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/2.14.3/docs/resources/event_orchestration_service#value EventOrchestrationService#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

