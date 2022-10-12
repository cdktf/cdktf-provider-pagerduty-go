package eventorchestrationunrouted


type EventOrchestrationUnroutedSetRule struct {
	// actions block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/pagerduty/r/event_orchestration_unrouted#actions EventOrchestrationUnrouted#actions}
	Actions *EventOrchestrationUnroutedSetRuleActions `field:"required" json:"actions" yaml:"actions"`
	// condition block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/pagerduty/r/event_orchestration_unrouted#condition EventOrchestrationUnrouted#condition}
	Condition interface{} `field:"optional" json:"condition" yaml:"condition"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/pagerduty/r/event_orchestration_unrouted#disabled EventOrchestrationUnrouted#disabled}.
	Disabled interface{} `field:"optional" json:"disabled" yaml:"disabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/pagerduty/r/event_orchestration_unrouted#label EventOrchestrationUnrouted#label}.
	Label *string `field:"optional" json:"label" yaml:"label"`
}
