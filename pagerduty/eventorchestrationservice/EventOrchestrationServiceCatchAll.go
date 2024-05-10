// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package eventorchestrationservice


type EventOrchestrationServiceCatchAll struct {
	// actions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.12.0/docs/resources/event_orchestration_service#actions EventOrchestrationService#actions}
	Actions *EventOrchestrationServiceCatchAllActions `field:"required" json:"actions" yaml:"actions"`
}

