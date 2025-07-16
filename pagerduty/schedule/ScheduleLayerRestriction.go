// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schedule


type ScheduleLayerRestriction struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.27.1/docs/resources/schedule#duration_seconds Schedule#duration_seconds}.
	DurationSeconds *float64 `field:"required" json:"durationSeconds" yaml:"durationSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.27.1/docs/resources/schedule#start_time_of_day Schedule#start_time_of_day}.
	StartTimeOfDay *string `field:"required" json:"startTimeOfDay" yaml:"startTimeOfDay"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.27.1/docs/resources/schedule#type Schedule#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.27.1/docs/resources/schedule#start_day_of_week Schedule#start_day_of_week}.
	StartDayOfWeek *float64 `field:"optional" json:"startDayOfWeek" yaml:"startDayOfWeek"`
}

