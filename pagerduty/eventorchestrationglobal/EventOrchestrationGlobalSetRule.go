package eventorchestrationglobal


type EventOrchestrationGlobalSetRule struct {
	// actions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/2.15.0/docs/resources/event_orchestration_global#actions EventOrchestrationGlobal#actions}
	Actions *EventOrchestrationGlobalSetRuleActions `field:"required" json:"actions" yaml:"actions"`
	// condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/2.15.0/docs/resources/event_orchestration_global#condition EventOrchestrationGlobal#condition}
	Condition interface{} `field:"optional" json:"condition" yaml:"condition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/2.15.0/docs/resources/event_orchestration_global#disabled EventOrchestrationGlobal#disabled}.
	Disabled interface{} `field:"optional" json:"disabled" yaml:"disabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/2.15.0/docs/resources/event_orchestration_global#label EventOrchestrationGlobal#label}.
	Label *string `field:"optional" json:"label" yaml:"label"`
}

