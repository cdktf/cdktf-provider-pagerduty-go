// Prebuilt pagerduty Provider for Terraform CDK (cdktf)
package pagerduty


type ServiceEventRuleTimeFrame struct {
	// active_between block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/pagerduty/r/service_event_rule#active_between ServiceEventRule#active_between}
	ActiveBetween interface{} `field:"optional" json:"activeBetween" yaml:"activeBetween"`
	// scheduled_weekly block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/pagerduty/r/service_event_rule#scheduled_weekly ServiceEventRule#scheduled_weekly}
	ScheduledWeekly interface{} `field:"optional" json:"scheduledWeekly" yaml:"scheduledWeekly"`
}

