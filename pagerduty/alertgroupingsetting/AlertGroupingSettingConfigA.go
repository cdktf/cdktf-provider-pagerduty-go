// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package alertgroupingsetting


type AlertGroupingSettingConfigA struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.30.7/docs/resources/alert_grouping_setting#aggregate AlertGroupingSetting#aggregate}.
	Aggregate *string `field:"optional" json:"aggregate" yaml:"aggregate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.30.7/docs/resources/alert_grouping_setting#fields AlertGroupingSetting#fields}.
	Fields *[]*string `field:"optional" json:"fields" yaml:"fields"`
	// An array of strings which represent the iag fields with which to intelligently group against.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.30.7/docs/resources/alert_grouping_setting#iag_fields AlertGroupingSetting#iag_fields}
	IagFields *[]*string `field:"optional" json:"iagFields" yaml:"iagFields"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.30.7/docs/resources/alert_grouping_setting#timeout AlertGroupingSetting#timeout}.
	Timeout *float64 `field:"optional" json:"timeout" yaml:"timeout"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.30.7/docs/resources/alert_grouping_setting#time_window AlertGroupingSetting#time_window}.
	TimeWindow *float64 `field:"optional" json:"timeWindow" yaml:"timeWindow"`
}

