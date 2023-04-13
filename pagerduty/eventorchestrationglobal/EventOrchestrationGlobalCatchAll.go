package eventorchestrationglobal


type EventOrchestrationGlobalCatchAll struct {
	// actions block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/pagerduty/r/event_orchestration_global#actions EventOrchestrationGlobal#actions}
	Actions *EventOrchestrationGlobalCatchAllActions `field:"required" json:"actions" yaml:"actions"`
}

