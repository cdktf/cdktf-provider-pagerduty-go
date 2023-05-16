package rulesetrule


type RulesetRuleConditions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/2.14.5/docs/resources/ruleset_rule#operator RulesetRule#operator}.
	Operator *string `field:"optional" json:"operator" yaml:"operator"`
	// subconditions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/2.14.5/docs/resources/ruleset_rule#subconditions RulesetRule#subconditions}
	Subconditions interface{} `field:"optional" json:"subconditions" yaml:"subconditions"`
}

