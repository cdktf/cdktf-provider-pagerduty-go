// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package serviceeventrule


type ServiceEventRuleTimeFrameScheduledWeekly struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.23.0/docs/resources/service_event_rule#duration ServiceEventRule#duration}.
	Duration *float64 `field:"optional" json:"duration" yaml:"duration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.23.0/docs/resources/service_event_rule#start_time ServiceEventRule#start_time}.
	StartTime *float64 `field:"optional" json:"startTime" yaml:"startTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.23.0/docs/resources/service_event_rule#timezone ServiceEventRule#timezone}.
	Timezone *string `field:"optional" json:"timezone" yaml:"timezone"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.23.0/docs/resources/service_event_rule#weekdays ServiceEventRule#weekdays}.
	Weekdays *[]*float64 `field:"optional" json:"weekdays" yaml:"weekdays"`
}

