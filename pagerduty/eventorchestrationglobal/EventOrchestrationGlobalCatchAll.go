package eventorchestrationglobal


type EventOrchestrationGlobalCatchAll struct {
	// actions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/2.15.0/docs/resources/event_orchestration_global#actions EventOrchestrationGlobal#actions}
	Actions *EventOrchestrationGlobalCatchAllActions `field:"required" json:"actions" yaml:"actions"`
}

