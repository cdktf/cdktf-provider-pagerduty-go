// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package eventorchestrationrouter


type EventOrchestrationRouterSetRuleActions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/2.16.0/docs/resources/event_orchestration_router#route_to EventOrchestrationRouter#route_to}.
	RouteTo *string `field:"required" json:"routeTo" yaml:"routeTo"`
}

