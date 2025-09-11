// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package eventorchestrationglobal


type EventOrchestrationGlobalSetRuleActionsAutomationAction struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.29.0/docs/resources/event_orchestration_global#name EventOrchestrationGlobal#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.29.0/docs/resources/event_orchestration_global#url EventOrchestrationGlobal#url}.
	Url *string `field:"required" json:"url" yaml:"url"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.29.0/docs/resources/event_orchestration_global#auto_send EventOrchestrationGlobal#auto_send}.
	AutoSend interface{} `field:"optional" json:"autoSend" yaml:"autoSend"`
	// header block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.29.0/docs/resources/event_orchestration_global#header EventOrchestrationGlobal#header}
	Header interface{} `field:"optional" json:"header" yaml:"header"`
	// parameter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.29.0/docs/resources/event_orchestration_global#parameter EventOrchestrationGlobal#parameter}
	Parameter interface{} `field:"optional" json:"parameter" yaml:"parameter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.29.0/docs/resources/event_orchestration_global#trigger_types EventOrchestrationGlobal#trigger_types}.
	TriggerTypes *[]*string `field:"optional" json:"triggerTypes" yaml:"triggerTypes"`
}

