package escalationpolicy


type EscalationPolicyRule struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/pagerduty/r/escalation_policy#escalation_delay_in_minutes EscalationPolicy#escalation_delay_in_minutes}.
	EscalationDelayInMinutes *float64 `field:"required" json:"escalationDelayInMinutes" yaml:"escalationDelayInMinutes"`
	// target block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/pagerduty/r/escalation_policy#target EscalationPolicy#target}
	Target interface{} `field:"required" json:"target" yaml:"target"`
}

