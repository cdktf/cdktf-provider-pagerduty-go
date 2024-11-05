// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package eventorchestrationunrouted


type EventOrchestrationUnroutedCatchAllActionsVariable struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.17.0/docs/resources/event_orchestration_unrouted#name EventOrchestrationUnrouted#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.17.0/docs/resources/event_orchestration_unrouted#path EventOrchestrationUnrouted#path}.
	Path *string `field:"required" json:"path" yaml:"path"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.17.0/docs/resources/event_orchestration_unrouted#type EventOrchestrationUnrouted#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.17.0/docs/resources/event_orchestration_unrouted#value EventOrchestrationUnrouted#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

