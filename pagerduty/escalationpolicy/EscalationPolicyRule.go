// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package escalationpolicy


type EscalationPolicyRule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/2.16.2/docs/resources/escalation_policy#escalation_delay_in_minutes EscalationPolicy#escalation_delay_in_minutes}.
	EscalationDelayInMinutes *float64 `field:"required" json:"escalationDelayInMinutes" yaml:"escalationDelayInMinutes"`
	// target block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/2.16.2/docs/resources/escalation_policy#target EscalationPolicy#target}
	Target interface{} `field:"required" json:"target" yaml:"target"`
}

