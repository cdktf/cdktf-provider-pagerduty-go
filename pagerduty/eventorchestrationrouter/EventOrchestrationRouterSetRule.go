package eventorchestrationrouter


type EventOrchestrationRouterSetRule struct {
	// actions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/2.15.1/docs/resources/event_orchestration_router#actions EventOrchestrationRouter#actions}
	Actions *EventOrchestrationRouterSetRuleActions `field:"required" json:"actions" yaml:"actions"`
	// condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/2.15.1/docs/resources/event_orchestration_router#condition EventOrchestrationRouter#condition}
	Condition interface{} `field:"optional" json:"condition" yaml:"condition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/2.15.1/docs/resources/event_orchestration_router#disabled EventOrchestrationRouter#disabled}.
	Disabled interface{} `field:"optional" json:"disabled" yaml:"disabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/2.15.1/docs/resources/event_orchestration_router#label EventOrchestrationRouter#label}.
	Label *string `field:"optional" json:"label" yaml:"label"`
}

