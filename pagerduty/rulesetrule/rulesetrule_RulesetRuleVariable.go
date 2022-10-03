package rulesetrule


type RulesetRuleVariable struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/pagerduty/r/ruleset_rule#name RulesetRule#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// parameters block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/pagerduty/r/ruleset_rule#parameters RulesetRule#parameters}
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/pagerduty/r/ruleset_rule#type RulesetRule#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

