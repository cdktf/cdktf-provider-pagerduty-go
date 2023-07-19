package eventorchestrationrouter


type EventOrchestrationRouterSetRuleCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/2.15.1/docs/resources/event_orchestration_router#expression EventOrchestrationRouter#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
}

