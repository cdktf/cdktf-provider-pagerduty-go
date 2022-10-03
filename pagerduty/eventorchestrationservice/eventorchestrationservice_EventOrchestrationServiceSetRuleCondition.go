package eventorchestrationservice


type EventOrchestrationServiceSetRuleCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/pagerduty/r/event_orchestration_service#expression EventOrchestrationService#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
}

