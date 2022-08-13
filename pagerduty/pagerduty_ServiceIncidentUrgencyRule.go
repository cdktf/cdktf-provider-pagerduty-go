// Prebuilt pagerduty Provider for Terraform CDK (cdktf)
package pagerduty


type ServiceIncidentUrgencyRule struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/pagerduty/r/service#type Service#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// during_support_hours block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/pagerduty/r/service#during_support_hours Service#during_support_hours}
	DuringSupportHours *ServiceIncidentUrgencyRuleDuringSupportHours `field:"optional" json:"duringSupportHours" yaml:"duringSupportHours"`
	// outside_support_hours block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/pagerduty/r/service#outside_support_hours Service#outside_support_hours}
	OutsideSupportHours *ServiceIncidentUrgencyRuleOutsideSupportHours `field:"optional" json:"outsideSupportHours" yaml:"outsideSupportHours"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/pagerduty/r/service#urgency Service#urgency}.
	Urgency *string `field:"optional" json:"urgency" yaml:"urgency"`
}

