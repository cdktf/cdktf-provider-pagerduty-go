// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package service


type ServiceSupportHours struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.7.1/docs/resources/service#days_of_week Service#days_of_week}.
	DaysOfWeek *[]*float64 `field:"optional" json:"daysOfWeek" yaml:"daysOfWeek"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.7.1/docs/resources/service#end_time Service#end_time}.
	EndTime *string `field:"optional" json:"endTime" yaml:"endTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.7.1/docs/resources/service#start_time Service#start_time}.
	StartTime *string `field:"optional" json:"startTime" yaml:"startTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.7.1/docs/resources/service#time_zone Service#time_zone}.
	TimeZone *string `field:"optional" json:"timeZone" yaml:"timeZone"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.7.1/docs/resources/service#type Service#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

