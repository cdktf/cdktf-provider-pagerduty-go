package rulesetrule


type RulesetRuleConditionsSubconditions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/2.14.4/docs/resources/ruleset_rule#operator RulesetRule#operator}.
	Operator *string `field:"optional" json:"operator" yaml:"operator"`
	// parameter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/2.14.4/docs/resources/ruleset_rule#parameter RulesetRule#parameter}
	Parameter interface{} `field:"optional" json:"parameter" yaml:"parameter"`
}

