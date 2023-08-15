package eventorchestrationunrouted


type EventOrchestrationUnroutedCatchAll struct {
	// actions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/2.15.3/docs/resources/event_orchestration_unrouted#actions EventOrchestrationUnrouted#actions}
	Actions *EventOrchestrationUnroutedCatchAllActions `field:"required" json:"actions" yaml:"actions"`
}

