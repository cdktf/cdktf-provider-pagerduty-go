// Prebuilt pagerduty Provider for Terraform CDK (cdktf)
package pagerduty


type ServiceIncidentUrgencyRuleDuringSupportHours struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/pagerduty/r/service#type Service#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/pagerduty/r/service#urgency Service#urgency}.
	Urgency *string `field:"optional" json:"urgency" yaml:"urgency"`
}

