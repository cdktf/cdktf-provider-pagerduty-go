// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package rulesetrule


type RulesetRuleActionsExtractions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.13.1/docs/resources/ruleset_rule#regex RulesetRule#regex}.
	Regex *string `field:"optional" json:"regex" yaml:"regex"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.13.1/docs/resources/ruleset_rule#source RulesetRule#source}.
	Source *string `field:"optional" json:"source" yaml:"source"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.13.1/docs/resources/ruleset_rule#target RulesetRule#target}.
	Target *string `field:"optional" json:"target" yaml:"target"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.13.1/docs/resources/ruleset_rule#template RulesetRule#template}.
	Template *string `field:"optional" json:"template" yaml:"template"`
}

