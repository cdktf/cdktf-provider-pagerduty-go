// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package rulesetrule


type RulesetRuleVariable struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.24.1/docs/resources/ruleset_rule#name RulesetRule#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.24.1/docs/resources/ruleset_rule#parameters RulesetRule#parameters}
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.24.1/docs/resources/ruleset_rule#type RulesetRule#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

