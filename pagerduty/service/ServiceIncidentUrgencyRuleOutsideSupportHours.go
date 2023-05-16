package service


type ServiceIncidentUrgencyRuleOutsideSupportHours struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/2.14.5/docs/resources/service#type Service#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/2.14.5/docs/resources/service#urgency Service#urgency}.
	Urgency *string `field:"optional" json:"urgency" yaml:"urgency"`
}

