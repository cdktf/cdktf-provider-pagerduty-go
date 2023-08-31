// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package eventorchestrationservice


type EventOrchestrationServiceSetRuleActionsAutomationAction struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/2.16.2/docs/resources/event_orchestration_service#name EventOrchestrationService#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/2.16.2/docs/resources/event_orchestration_service#url EventOrchestrationService#url}.
	Url *string `field:"required" json:"url" yaml:"url"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/2.16.2/docs/resources/event_orchestration_service#auto_send EventOrchestrationService#auto_send}.
	AutoSend interface{} `field:"optional" json:"autoSend" yaml:"autoSend"`
	// header block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/2.16.2/docs/resources/event_orchestration_service#header EventOrchestrationService#header}
	Header interface{} `field:"optional" json:"header" yaml:"header"`
	// parameter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/2.16.2/docs/resources/event_orchestration_service#parameter EventOrchestrationService#parameter}
	Parameter interface{} `field:"optional" json:"parameter" yaml:"parameter"`
}

