package eventorchestrationservice


type EventOrchestrationServiceSetRuleCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/2.14.4/docs/resources/event_orchestration_service#expression EventOrchestrationService#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
}

