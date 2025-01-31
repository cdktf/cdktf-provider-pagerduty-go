// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package escalationpolicy


type EscalationPolicyRule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.19.4/docs/resources/escalation_policy#escalation_delay_in_minutes EscalationPolicy#escalation_delay_in_minutes}.
	EscalationDelayInMinutes *float64 `field:"required" json:"escalationDelayInMinutes" yaml:"escalationDelayInMinutes"`
	// target block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.19.4/docs/resources/escalation_policy#target EscalationPolicy#target}
	Target interface{} `field:"required" json:"target" yaml:"target"`
	// escalation_rule_assignment_strategy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.19.4/docs/resources/escalation_policy#escalation_rule_assignment_strategy EscalationPolicy#escalation_rule_assignment_strategy}
	EscalationRuleAssignmentStrategy *EscalationPolicyRuleEscalationRuleAssignmentStrategy `field:"optional" json:"escalationRuleAssignmentStrategy" yaml:"escalationRuleAssignmentStrategy"`
}

