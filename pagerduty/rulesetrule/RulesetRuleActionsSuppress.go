// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package rulesetrule


type RulesetRuleActionsSuppress struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.11.3/docs/resources/ruleset_rule#threshold_time_amount RulesetRule#threshold_time_amount}.
	ThresholdTimeAmount *float64 `field:"optional" json:"thresholdTimeAmount" yaml:"thresholdTimeAmount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.11.3/docs/resources/ruleset_rule#threshold_time_unit RulesetRule#threshold_time_unit}.
	ThresholdTimeUnit *string `field:"optional" json:"thresholdTimeUnit" yaml:"thresholdTimeUnit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.11.3/docs/resources/ruleset_rule#threshold_value RulesetRule#threshold_value}.
	ThresholdValue *float64 `field:"optional" json:"thresholdValue" yaml:"thresholdValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.11.3/docs/resources/ruleset_rule#value RulesetRule#value}.
	Value interface{} `field:"optional" json:"value" yaml:"value"`
}

