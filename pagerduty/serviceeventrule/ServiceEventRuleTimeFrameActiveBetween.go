package serviceeventrule


type ServiceEventRuleTimeFrameActiveBetween struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/2.14.5/docs/resources/service_event_rule#end_time ServiceEventRule#end_time}.
	EndTime *float64 `field:"optional" json:"endTime" yaml:"endTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/2.14.5/docs/resources/service_event_rule#start_time ServiceEventRule#start_time}.
	StartTime *float64 `field:"optional" json:"startTime" yaml:"startTime"`
}

