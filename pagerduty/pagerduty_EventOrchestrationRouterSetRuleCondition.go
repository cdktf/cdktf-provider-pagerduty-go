// Prebuilt pagerduty Provider for Terraform CDK (cdktf)
package pagerduty


type EventOrchestrationRouterSetRuleCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/pagerduty/r/event_orchestration_router#expression EventOrchestrationRouter#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
}

