package eventorchestrationservice


type EventOrchestrationServiceSetRule struct {
	// actions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/2.14.6/docs/resources/event_orchestration_service#actions EventOrchestrationService#actions}
	Actions *EventOrchestrationServiceSetRuleActions `field:"required" json:"actions" yaml:"actions"`
	// condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/2.14.6/docs/resources/event_orchestration_service#condition EventOrchestrationService#condition}
	Condition interface{} `field:"optional" json:"condition" yaml:"condition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/2.14.6/docs/resources/event_orchestration_service#disabled EventOrchestrationService#disabled}.
	Disabled interface{} `field:"optional" json:"disabled" yaml:"disabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/2.14.6/docs/resources/event_orchestration_service#label EventOrchestrationService#label}.
	Label *string `field:"optional" json:"label" yaml:"label"`
}

