// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package eventorchestrationservice


type EventOrchestrationServiceCatchAllActionsPagerdutyAutomationAction struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.24.0/docs/resources/event_orchestration_service#action_id EventOrchestrationService#action_id}.
	ActionId *string `field:"required" json:"actionId" yaml:"actionId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.24.0/docs/resources/event_orchestration_service#trigger_types EventOrchestrationService#trigger_types}.
	TriggerTypes *[]*string `field:"optional" json:"triggerTypes" yaml:"triggerTypes"`
}

