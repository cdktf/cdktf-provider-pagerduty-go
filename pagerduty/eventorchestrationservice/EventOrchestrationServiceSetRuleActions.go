// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package eventorchestrationservice


type EventOrchestrationServiceSetRuleActions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.14.0/docs/resources/event_orchestration_service#annotate EventOrchestrationService#annotate}.
	Annotate *string `field:"optional" json:"annotate" yaml:"annotate"`
	// automation_action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.14.0/docs/resources/event_orchestration_service#automation_action EventOrchestrationService#automation_action}
	AutomationAction *EventOrchestrationServiceSetRuleActionsAutomationAction `field:"optional" json:"automationAction" yaml:"automationAction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.14.0/docs/resources/event_orchestration_service#event_action EventOrchestrationService#event_action}.
	EventAction *string `field:"optional" json:"eventAction" yaml:"eventAction"`
	// extraction block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.14.0/docs/resources/event_orchestration_service#extraction EventOrchestrationService#extraction}
	Extraction interface{} `field:"optional" json:"extraction" yaml:"extraction"`
	// incident_custom_field_update block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.14.0/docs/resources/event_orchestration_service#incident_custom_field_update EventOrchestrationService#incident_custom_field_update}
	IncidentCustomFieldUpdate interface{} `field:"optional" json:"incidentCustomFieldUpdate" yaml:"incidentCustomFieldUpdate"`
	// pagerduty_automation_action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.14.0/docs/resources/event_orchestration_service#pagerduty_automation_action EventOrchestrationService#pagerduty_automation_action}
	PagerdutyAutomationAction *EventOrchestrationServiceSetRuleActionsPagerdutyAutomationAction `field:"optional" json:"pagerdutyAutomationAction" yaml:"pagerdutyAutomationAction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.14.0/docs/resources/event_orchestration_service#priority EventOrchestrationService#priority}.
	Priority *string `field:"optional" json:"priority" yaml:"priority"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.14.0/docs/resources/event_orchestration_service#route_to EventOrchestrationService#route_to}.
	RouteTo *string `field:"optional" json:"routeTo" yaml:"routeTo"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.14.0/docs/resources/event_orchestration_service#severity EventOrchestrationService#severity}.
	Severity *string `field:"optional" json:"severity" yaml:"severity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.14.0/docs/resources/event_orchestration_service#suppress EventOrchestrationService#suppress}.
	Suppress interface{} `field:"optional" json:"suppress" yaml:"suppress"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.14.0/docs/resources/event_orchestration_service#suspend EventOrchestrationService#suspend}.
	Suspend *float64 `field:"optional" json:"suspend" yaml:"suspend"`
	// variable block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.14.0/docs/resources/event_orchestration_service#variable EventOrchestrationService#variable}
	Variable interface{} `field:"optional" json:"variable" yaml:"variable"`
}

