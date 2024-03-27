// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package eventorchestrationrouter


type EventOrchestrationRouterCatchAllActions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.10.1/docs/resources/event_orchestration_router#route_to EventOrchestrationRouter#route_to}.
	RouteTo *string `field:"required" json:"routeTo" yaml:"routeTo"`
}

