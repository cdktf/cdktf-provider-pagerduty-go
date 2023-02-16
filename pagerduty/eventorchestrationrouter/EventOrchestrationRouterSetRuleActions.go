package eventorchestrationrouter


type EventOrchestrationRouterSetRuleActions struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/pagerduty/r/event_orchestration_router#route_to EventOrchestrationRouter#route_to}.
	RouteTo *string `field:"required" json:"routeTo" yaml:"routeTo"`
}

