// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package serviceeventrule


type ServiceEventRuleActions struct {
	// annotate block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.13.1/docs/resources/service_event_rule#annotate ServiceEventRule#annotate}
	Annotate interface{} `field:"optional" json:"annotate" yaml:"annotate"`
	// event_action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.13.1/docs/resources/service_event_rule#event_action ServiceEventRule#event_action}
	EventAction interface{} `field:"optional" json:"eventAction" yaml:"eventAction"`
	// extractions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.13.1/docs/resources/service_event_rule#extractions ServiceEventRule#extractions}
	Extractions interface{} `field:"optional" json:"extractions" yaml:"extractions"`
	// priority block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.13.1/docs/resources/service_event_rule#priority ServiceEventRule#priority}
	Priority interface{} `field:"optional" json:"priority" yaml:"priority"`
	// severity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.13.1/docs/resources/service_event_rule#severity ServiceEventRule#severity}
	Severity interface{} `field:"optional" json:"severity" yaml:"severity"`
	// suppress block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.13.1/docs/resources/service_event_rule#suppress ServiceEventRule#suppress}
	Suppress interface{} `field:"optional" json:"suppress" yaml:"suppress"`
	// suspend block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.13.1/docs/resources/service_event_rule#suspend ServiceEventRule#suspend}
	Suspend interface{} `field:"optional" json:"suspend" yaml:"suspend"`
}

