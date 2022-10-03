package eventorchestrationservice


type EventOrchestrationServiceCatchAll struct {
	// actions block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/pagerduty/r/event_orchestration_service#actions EventOrchestrationService#actions}
	Actions *EventOrchestrationServiceCatchAllActions `field:"required" json:"actions" yaml:"actions"`
}

