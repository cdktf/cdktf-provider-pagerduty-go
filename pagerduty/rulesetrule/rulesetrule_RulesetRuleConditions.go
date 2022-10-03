package rulesetrule


type RulesetRuleConditions struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/pagerduty/r/ruleset_rule#operator RulesetRule#operator}.
	Operator *string `field:"optional" json:"operator" yaml:"operator"`
	// subconditions block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/pagerduty/r/ruleset_rule#subconditions RulesetRule#subconditions}
	Subconditions interface{} `field:"optional" json:"subconditions" yaml:"subconditions"`
}

