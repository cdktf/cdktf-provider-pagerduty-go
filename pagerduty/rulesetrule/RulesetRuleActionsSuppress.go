package rulesetrule


type RulesetRuleActionsSuppress struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/pagerduty/r/ruleset_rule#threshold_time_amount RulesetRule#threshold_time_amount}.
	ThresholdTimeAmount *float64 `field:"optional" json:"thresholdTimeAmount" yaml:"thresholdTimeAmount"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/pagerduty/r/ruleset_rule#threshold_time_unit RulesetRule#threshold_time_unit}.
	ThresholdTimeUnit *string `field:"optional" json:"thresholdTimeUnit" yaml:"thresholdTimeUnit"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/pagerduty/r/ruleset_rule#threshold_value RulesetRule#threshold_value}.
	ThresholdValue *float64 `field:"optional" json:"thresholdValue" yaml:"thresholdValue"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/pagerduty/r/ruleset_rule#value RulesetRule#value}.
	Value interface{} `field:"optional" json:"value" yaml:"value"`
}

