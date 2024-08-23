// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package eventorchestrationunrouted


type EventOrchestrationUnroutedSetRule struct {
	// actions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.15.5/docs/resources/event_orchestration_unrouted#actions EventOrchestrationUnrouted#actions}
	Actions *EventOrchestrationUnroutedSetRuleActions `field:"required" json:"actions" yaml:"actions"`
	// condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.15.5/docs/resources/event_orchestration_unrouted#condition EventOrchestrationUnrouted#condition}
	Condition interface{} `field:"optional" json:"condition" yaml:"condition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.15.5/docs/resources/event_orchestration_unrouted#disabled EventOrchestrationUnrouted#disabled}.
	Disabled interface{} `field:"optional" json:"disabled" yaml:"disabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.15.5/docs/resources/event_orchestration_unrouted#label EventOrchestrationUnrouted#label}.
	Label *string `field:"optional" json:"label" yaml:"label"`
}

