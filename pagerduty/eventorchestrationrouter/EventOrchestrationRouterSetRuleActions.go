// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package eventorchestrationrouter


type EventOrchestrationRouterSetRuleActions struct {
	// dynamic_route_to block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.28.1/docs/resources/event_orchestration_router#dynamic_route_to EventOrchestrationRouter#dynamic_route_to}
	DynamicRouteTo interface{} `field:"optional" json:"dynamicRouteTo" yaml:"dynamicRouteTo"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.28.1/docs/resources/event_orchestration_router#route_to EventOrchestrationRouter#route_to}.
	RouteTo *string `field:"optional" json:"routeTo" yaml:"routeTo"`
}

