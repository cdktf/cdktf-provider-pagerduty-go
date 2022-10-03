package eventorchestrationunrouted


type EventOrchestrationUnroutedCatchAll struct {
	// actions block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/pagerduty/r/event_orchestration_unrouted#actions EventOrchestrationUnrouted#actions}
	Actions *EventOrchestrationUnroutedCatchAllActions `field:"required" json:"actions" yaml:"actions"`
}

