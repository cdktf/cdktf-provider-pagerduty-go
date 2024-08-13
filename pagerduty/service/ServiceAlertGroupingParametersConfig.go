// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package service


type ServiceAlertGroupingParametersConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.15.2/docs/resources/service#aggregate Service#aggregate}.
	Aggregate *string `field:"optional" json:"aggregate" yaml:"aggregate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.15.2/docs/resources/service#fields Service#fields}.
	Fields *[]*string `field:"optional" json:"fields" yaml:"fields"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.15.2/docs/resources/service#timeout Service#timeout}.
	Timeout *float64 `field:"optional" json:"timeout" yaml:"timeout"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.15.2/docs/resources/service#time_window Service#time_window}.
	TimeWindow *float64 `field:"optional" json:"timeWindow" yaml:"timeWindow"`
}

