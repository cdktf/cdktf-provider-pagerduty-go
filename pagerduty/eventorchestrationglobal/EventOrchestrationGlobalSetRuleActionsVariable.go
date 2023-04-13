package eventorchestrationglobal


type EventOrchestrationGlobalSetRuleActionsVariable struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/pagerduty/r/event_orchestration_global#name EventOrchestrationGlobal#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/pagerduty/r/event_orchestration_global#path EventOrchestrationGlobal#path}.
	Path *string `field:"required" json:"path" yaml:"path"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/pagerduty/r/event_orchestration_global#type EventOrchestrationGlobal#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/pagerduty/r/event_orchestration_global#value EventOrchestrationGlobal#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

