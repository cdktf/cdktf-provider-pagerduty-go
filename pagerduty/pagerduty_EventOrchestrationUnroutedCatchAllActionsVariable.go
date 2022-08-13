// Prebuilt pagerduty Provider for Terraform CDK (cdktf)
package pagerduty


type EventOrchestrationUnroutedCatchAllActionsVariable struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/pagerduty/r/event_orchestration_unrouted#name EventOrchestrationUnrouted#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/pagerduty/r/event_orchestration_unrouted#path EventOrchestrationUnrouted#path}.
	Path *string `field:"required" json:"path" yaml:"path"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/pagerduty/r/event_orchestration_unrouted#type EventOrchestrationUnrouted#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/pagerduty/r/event_orchestration_unrouted#value EventOrchestrationUnrouted#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

