// Prebuilt pagerduty Provider for Terraform CDK (cdktf)
package pagerduty


type EventOrchestrationUnroutedSetRuleCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/pagerduty/r/event_orchestration_unrouted#expression EventOrchestrationUnrouted#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
}

