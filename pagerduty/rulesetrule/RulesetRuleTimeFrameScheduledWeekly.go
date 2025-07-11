// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package rulesetrule


type RulesetRuleTimeFrameScheduledWeekly struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.27.0/docs/resources/ruleset_rule#duration RulesetRule#duration}.
	Duration *float64 `field:"optional" json:"duration" yaml:"duration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.27.0/docs/resources/ruleset_rule#start_time RulesetRule#start_time}.
	StartTime *float64 `field:"optional" json:"startTime" yaml:"startTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.27.0/docs/resources/ruleset_rule#timezone RulesetRule#timezone}.
	Timezone *string `field:"optional" json:"timezone" yaml:"timezone"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.27.0/docs/resources/ruleset_rule#weekdays RulesetRule#weekdays}.
	Weekdays *[]*float64 `field:"optional" json:"weekdays" yaml:"weekdays"`
}

