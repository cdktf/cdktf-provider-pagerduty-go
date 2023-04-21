package eventorchestrationrouter


type EventOrchestrationRouterCatchAll struct {
	// actions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/2.14.3/docs/resources/event_orchestration_router#actions EventOrchestrationRouter#actions}
	Actions *EventOrchestrationRouterCatchAllActions `field:"required" json:"actions" yaml:"actions"`
}

