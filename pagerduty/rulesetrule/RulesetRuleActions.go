// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package rulesetrule


type RulesetRuleActions struct {
	// annotate block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.26.2/docs/resources/ruleset_rule#annotate RulesetRule#annotate}
	Annotate interface{} `field:"optional" json:"annotate" yaml:"annotate"`
	// event_action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.26.2/docs/resources/ruleset_rule#event_action RulesetRule#event_action}
	EventAction interface{} `field:"optional" json:"eventAction" yaml:"eventAction"`
	// extractions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.26.2/docs/resources/ruleset_rule#extractions RulesetRule#extractions}
	Extractions interface{} `field:"optional" json:"extractions" yaml:"extractions"`
	// priority block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.26.2/docs/resources/ruleset_rule#priority RulesetRule#priority}
	Priority interface{} `field:"optional" json:"priority" yaml:"priority"`
	// route block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.26.2/docs/resources/ruleset_rule#route RulesetRule#route}
	Route interface{} `field:"optional" json:"route" yaml:"route"`
	// severity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.26.2/docs/resources/ruleset_rule#severity RulesetRule#severity}
	Severity interface{} `field:"optional" json:"severity" yaml:"severity"`
	// suppress block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.26.2/docs/resources/ruleset_rule#suppress RulesetRule#suppress}
	Suppress interface{} `field:"optional" json:"suppress" yaml:"suppress"`
	// suspend block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.26.2/docs/resources/ruleset_rule#suspend RulesetRule#suspend}
	Suspend interface{} `field:"optional" json:"suspend" yaml:"suspend"`
}

