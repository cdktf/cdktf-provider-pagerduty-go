// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package service


type ServiceAutoPauseNotificationsParameters struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.6.0/docs/resources/service#enabled Service#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.6.0/docs/resources/service#timeout Service#timeout}.
	Timeout *float64 `field:"optional" json:"timeout" yaml:"timeout"`
}

