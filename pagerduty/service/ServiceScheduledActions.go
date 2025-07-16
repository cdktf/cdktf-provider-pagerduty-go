// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package service


type ServiceScheduledActions struct {
	// at block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.27.1/docs/resources/service#at Service#at}
	At interface{} `field:"optional" json:"at" yaml:"at"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.27.1/docs/resources/service#to_urgency Service#to_urgency}.
	ToUrgency *string `field:"optional" json:"toUrgency" yaml:"toUrgency"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.27.1/docs/resources/service#type Service#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

