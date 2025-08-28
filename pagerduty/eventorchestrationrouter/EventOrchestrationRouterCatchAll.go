// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package eventorchestrationrouter


type EventOrchestrationRouterCatchAll struct {
	// actions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.28.2/docs/resources/event_orchestration_router#actions EventOrchestrationRouter#actions}
	Actions *EventOrchestrationRouterCatchAllActions `field:"required" json:"actions" yaml:"actions"`
}

