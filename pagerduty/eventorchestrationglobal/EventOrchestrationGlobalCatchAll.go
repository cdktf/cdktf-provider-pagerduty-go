// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package eventorchestrationglobal


type EventOrchestrationGlobalCatchAll struct {
	// actions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.18.2/docs/resources/event_orchestration_global#actions EventOrchestrationGlobal#actions}
	Actions *EventOrchestrationGlobalCatchAllActions `field:"required" json:"actions" yaml:"actions"`
}

