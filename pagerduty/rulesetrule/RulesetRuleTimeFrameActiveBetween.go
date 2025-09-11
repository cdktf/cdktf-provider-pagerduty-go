// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package rulesetrule


type RulesetRuleTimeFrameActiveBetween struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.29.0/docs/resources/ruleset_rule#end_time RulesetRule#end_time}.
	EndTime *float64 `field:"optional" json:"endTime" yaml:"endTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.29.0/docs/resources/ruleset_rule#start_time RulesetRule#start_time}.
	StartTime *float64 `field:"optional" json:"startTime" yaml:"startTime"`
}

