// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package alertgroupingsetting


type AlertGroupingSettingConfigA struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.17.1/docs/resources/alert_grouping_setting#aggregate AlertGroupingSetting#aggregate}.
	Aggregate *string `field:"optional" json:"aggregate" yaml:"aggregate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.17.1/docs/resources/alert_grouping_setting#fields AlertGroupingSetting#fields}.
	Fields *[]*string `field:"optional" json:"fields" yaml:"fields"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.17.1/docs/resources/alert_grouping_setting#timeout AlertGroupingSetting#timeout}.
	Timeout *float64 `field:"optional" json:"timeout" yaml:"timeout"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.17.1/docs/resources/alert_grouping_setting#time_window AlertGroupingSetting#time_window}.
	TimeWindow *float64 `field:"optional" json:"timeWindow" yaml:"timeWindow"`
}

