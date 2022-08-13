// Prebuilt pagerduty Provider for Terraform CDK (cdktf)
package pagerduty


type RulesetRuleTimeFrameActiveBetween struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/pagerduty/r/ruleset_rule#end_time RulesetRule#end_time}.
	EndTime *float64 `field:"optional" json:"endTime" yaml:"endTime"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/pagerduty/r/ruleset_rule#start_time RulesetRule#start_time}.
	StartTime *float64 `field:"optional" json:"startTime" yaml:"startTime"`
}

