// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package escalationpolicy


type EscalationPolicyRuleTarget struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.3.1/docs/resources/escalation_policy#id EscalationPolicy#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.3.1/docs/resources/escalation_policy#type EscalationPolicy#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

