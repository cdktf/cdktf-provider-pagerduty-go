package eventorchestrationrouter


type EventOrchestrationRouterCatchAll struct {
	// actions block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/pagerduty/r/event_orchestration_router#actions EventOrchestrationRouter#actions}
	Actions *EventOrchestrationRouterCatchAllActions `field:"required" json:"actions" yaml:"actions"`
}

