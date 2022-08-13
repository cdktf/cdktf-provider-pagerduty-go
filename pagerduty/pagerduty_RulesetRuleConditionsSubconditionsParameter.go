// Prebuilt pagerduty Provider for Terraform CDK (cdktf)
package pagerduty


type RulesetRuleConditionsSubconditionsParameter struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/pagerduty/r/ruleset_rule#path RulesetRule#path}.
	Path *string `field:"optional" json:"path" yaml:"path"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/pagerduty/r/ruleset_rule#value RulesetRule#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

