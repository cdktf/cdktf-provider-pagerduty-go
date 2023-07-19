package rulesetrule


type RulesetRuleTimeFrame struct {
	// active_between block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/2.15.1/docs/resources/ruleset_rule#active_between RulesetRule#active_between}
	ActiveBetween interface{} `field:"optional" json:"activeBetween" yaml:"activeBetween"`
	// scheduled_weekly block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/2.15.1/docs/resources/ruleset_rule#scheduled_weekly RulesetRule#scheduled_weekly}
	ScheduledWeekly interface{} `field:"optional" json:"scheduledWeekly" yaml:"scheduledWeekly"`
}

