package eventorchestrationglobal


type EventOrchestrationGlobalSetRuleCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/2.14.2/docs/resources/event_orchestration_global#expression EventOrchestrationGlobal#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
}
