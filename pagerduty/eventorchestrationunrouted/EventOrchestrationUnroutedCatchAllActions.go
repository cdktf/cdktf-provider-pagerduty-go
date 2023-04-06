package eventorchestrationunrouted


type EventOrchestrationUnroutedCatchAllActions struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/pagerduty/r/event_orchestration_unrouted#event_action EventOrchestrationUnrouted#event_action}.
	EventAction *string `field:"optional" json:"eventAction" yaml:"eventAction"`
	// extraction block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/pagerduty/r/event_orchestration_unrouted#extraction EventOrchestrationUnrouted#extraction}
	Extraction interface{} `field:"optional" json:"extraction" yaml:"extraction"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/pagerduty/r/event_orchestration_unrouted#severity EventOrchestrationUnrouted#severity}.
	Severity *string `field:"optional" json:"severity" yaml:"severity"`
	// variable block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/pagerduty/r/event_orchestration_unrouted#variable EventOrchestrationUnrouted#variable}
	Variable interface{} `field:"optional" json:"variable" yaml:"variable"`
}
