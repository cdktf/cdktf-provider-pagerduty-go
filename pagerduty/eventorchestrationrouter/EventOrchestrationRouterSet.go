package eventorchestrationrouter


type EventOrchestrationRouterSet struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/2.15.0/docs/resources/event_orchestration_router#id EventOrchestrationRouter#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
	// rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/2.15.0/docs/resources/event_orchestration_router#rule EventOrchestrationRouter#rule}
	Rule interface{} `field:"optional" json:"rule" yaml:"rule"`
}

