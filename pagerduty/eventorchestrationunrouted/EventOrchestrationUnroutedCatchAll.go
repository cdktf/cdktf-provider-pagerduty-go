// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package eventorchestrationunrouted


type EventOrchestrationUnroutedCatchAll struct {
	// actions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.7.1/docs/resources/event_orchestration_unrouted#actions EventOrchestrationUnrouted#actions}
	Actions *EventOrchestrationUnroutedCatchAllActions `field:"required" json:"actions" yaml:"actions"`
}

