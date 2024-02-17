// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package rulesetrule


type RulesetRuleConditionsSubconditionsParameter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.8.0/docs/resources/ruleset_rule#path RulesetRule#path}.
	Path *string `field:"optional" json:"path" yaml:"path"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.8.0/docs/resources/ruleset_rule#value RulesetRule#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

