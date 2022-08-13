// Prebuilt pagerduty Provider for Terraform CDK (cdktf)
package pagerduty


type ServiceEventRuleTimeFrameScheduledWeekly struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/pagerduty/r/service_event_rule#duration ServiceEventRule#duration}.
	Duration *float64 `field:"optional" json:"duration" yaml:"duration"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/pagerduty/r/service_event_rule#start_time ServiceEventRule#start_time}.
	StartTime *float64 `field:"optional" json:"startTime" yaml:"startTime"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/pagerduty/r/service_event_rule#timezone ServiceEventRule#timezone}.
	Timezone *string `field:"optional" json:"timezone" yaml:"timezone"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/pagerduty/r/service_event_rule#weekdays ServiceEventRule#weekdays}.
	Weekdays *[]*float64 `field:"optional" json:"weekdays" yaml:"weekdays"`
}

