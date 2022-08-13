// Prebuilt pagerduty Provider for Terraform CDK (cdktf)
package pagerduty


type EventOrchestrationUnroutedCatchAllActionsExtraction struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/pagerduty/r/event_orchestration_unrouted#target EventOrchestrationUnrouted#target}.
	Target *string `field:"required" json:"target" yaml:"target"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/pagerduty/r/event_orchestration_unrouted#regex EventOrchestrationUnrouted#regex}.
	Regex *string `field:"optional" json:"regex" yaml:"regex"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/pagerduty/r/event_orchestration_unrouted#source EventOrchestrationUnrouted#source}.
	Source *string `field:"optional" json:"source" yaml:"source"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/pagerduty/r/event_orchestration_unrouted#template EventOrchestrationUnrouted#template}.
	Template *string `field:"optional" json:"template" yaml:"template"`
}

