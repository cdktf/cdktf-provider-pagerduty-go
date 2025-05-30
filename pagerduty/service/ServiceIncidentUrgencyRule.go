// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package service


type ServiceIncidentUrgencyRule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.26.0/docs/resources/service#type Service#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// during_support_hours block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.26.0/docs/resources/service#during_support_hours Service#during_support_hours}
	DuringSupportHours *ServiceIncidentUrgencyRuleDuringSupportHours `field:"optional" json:"duringSupportHours" yaml:"duringSupportHours"`
	// outside_support_hours block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.26.0/docs/resources/service#outside_support_hours Service#outside_support_hours}
	OutsideSupportHours *ServiceIncidentUrgencyRuleOutsideSupportHours `field:"optional" json:"outsideSupportHours" yaml:"outsideSupportHours"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.26.0/docs/resources/service#urgency Service#urgency}.
	Urgency *string `field:"optional" json:"urgency" yaml:"urgency"`
}

