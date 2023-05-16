package eventorchestrationservice


type EventOrchestrationServiceCatchAll struct {
	// actions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/2.14.5/docs/resources/event_orchestration_service#actions EventOrchestrationService#actions}
	Actions *EventOrchestrationServiceCatchAllActions `field:"required" json:"actions" yaml:"actions"`
}

