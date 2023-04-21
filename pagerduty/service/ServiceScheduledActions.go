package service


type ServiceScheduledActions struct {
	// at block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/2.14.3/docs/resources/service#at Service#at}
	At interface{} `field:"optional" json:"at" yaml:"at"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/2.14.3/docs/resources/service#to_urgency Service#to_urgency}.
	ToUrgency *string `field:"optional" json:"toUrgency" yaml:"toUrgency"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/2.14.3/docs/resources/service#type Service#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

