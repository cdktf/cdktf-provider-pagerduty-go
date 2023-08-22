package serviceeventrule


type ServiceEventRuleTimeFrame struct {
	// active_between block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/2.16.0/docs/resources/service_event_rule#active_between ServiceEventRule#active_between}
	ActiveBetween interface{} `field:"optional" json:"activeBetween" yaml:"activeBetween"`
	// scheduled_weekly block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/2.16.0/docs/resources/service_event_rule#scheduled_weekly ServiceEventRule#scheduled_weekly}
	ScheduledWeekly interface{} `field:"optional" json:"scheduledWeekly" yaml:"scheduledWeekly"`
}

