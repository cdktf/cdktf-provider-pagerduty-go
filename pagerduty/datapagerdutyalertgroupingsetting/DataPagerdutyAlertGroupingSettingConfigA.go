// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datapagerdutyalertgroupingsetting


type DataPagerdutyAlertGroupingSettingConfigA struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.17.2/docs/data-sources/alert_grouping_setting#aggregate DataPagerdutyAlertGroupingSetting#aggregate}.
	Aggregate *string `field:"optional" json:"aggregate" yaml:"aggregate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.17.2/docs/data-sources/alert_grouping_setting#fields DataPagerdutyAlertGroupingSetting#fields}.
	Fields *[]*string `field:"optional" json:"fields" yaml:"fields"`
}

