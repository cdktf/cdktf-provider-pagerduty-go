// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package eventorchestrationglobal


type EventOrchestrationGlobalSetRuleActions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.30.6/docs/resources/event_orchestration_global#annotate EventOrchestrationGlobal#annotate}.
	Annotate *string `field:"optional" json:"annotate" yaml:"annotate"`
	// automation_action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.30.6/docs/resources/event_orchestration_global#automation_action EventOrchestrationGlobal#automation_action}
	AutomationAction *EventOrchestrationGlobalSetRuleActionsAutomationAction `field:"optional" json:"automationAction" yaml:"automationAction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.30.6/docs/resources/event_orchestration_global#drop_event EventOrchestrationGlobal#drop_event}.
	DropEvent interface{} `field:"optional" json:"dropEvent" yaml:"dropEvent"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.30.6/docs/resources/event_orchestration_global#escalation_policy EventOrchestrationGlobal#escalation_policy}.
	EscalationPolicy *string `field:"optional" json:"escalationPolicy" yaml:"escalationPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.30.6/docs/resources/event_orchestration_global#event_action EventOrchestrationGlobal#event_action}.
	EventAction *string `field:"optional" json:"eventAction" yaml:"eventAction"`
	// extraction block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.30.6/docs/resources/event_orchestration_global#extraction EventOrchestrationGlobal#extraction}
	Extraction interface{} `field:"optional" json:"extraction" yaml:"extraction"`
	// incident_custom_field_update block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.30.6/docs/resources/event_orchestration_global#incident_custom_field_update EventOrchestrationGlobal#incident_custom_field_update}
	IncidentCustomFieldUpdate interface{} `field:"optional" json:"incidentCustomFieldUpdate" yaml:"incidentCustomFieldUpdate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.30.6/docs/resources/event_orchestration_global#priority EventOrchestrationGlobal#priority}.
	Priority *string `field:"optional" json:"priority" yaml:"priority"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.30.6/docs/resources/event_orchestration_global#route_to EventOrchestrationGlobal#route_to}.
	RouteTo *string `field:"optional" json:"routeTo" yaml:"routeTo"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.30.6/docs/resources/event_orchestration_global#severity EventOrchestrationGlobal#severity}.
	Severity *string `field:"optional" json:"severity" yaml:"severity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.30.6/docs/resources/event_orchestration_global#suppress EventOrchestrationGlobal#suppress}.
	Suppress interface{} `field:"optional" json:"suppress" yaml:"suppress"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.30.6/docs/resources/event_orchestration_global#suspend EventOrchestrationGlobal#suspend}.
	Suspend *float64 `field:"optional" json:"suspend" yaml:"suspend"`
	// variable block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.30.6/docs/resources/event_orchestration_global#variable EventOrchestrationGlobal#variable}
	Variable interface{} `field:"optional" json:"variable" yaml:"variable"`
}

