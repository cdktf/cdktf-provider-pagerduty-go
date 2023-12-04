// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package eventorchestrationunrouted


type EventOrchestrationUnroutedSetRuleActions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.2.2/docs/resources/event_orchestration_unrouted#event_action EventOrchestrationUnrouted#event_action}.
	EventAction *string `field:"optional" json:"eventAction" yaml:"eventAction"`
	// extraction block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.2.2/docs/resources/event_orchestration_unrouted#extraction EventOrchestrationUnrouted#extraction}
	Extraction interface{} `field:"optional" json:"extraction" yaml:"extraction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.2.2/docs/resources/event_orchestration_unrouted#route_to EventOrchestrationUnrouted#route_to}.
	RouteTo *string `field:"optional" json:"routeTo" yaml:"routeTo"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.2.2/docs/resources/event_orchestration_unrouted#severity EventOrchestrationUnrouted#severity}.
	Severity *string `field:"optional" json:"severity" yaml:"severity"`
	// variable block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.2.2/docs/resources/event_orchestration_unrouted#variable EventOrchestrationUnrouted#variable}
	Variable interface{} `field:"optional" json:"variable" yaml:"variable"`
}

