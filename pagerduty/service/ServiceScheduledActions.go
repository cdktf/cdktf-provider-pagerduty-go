package service


type ServiceScheduledActions struct {
	// at block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/pagerduty/r/service#at Service#at}
	At interface{} `field:"optional" json:"at" yaml:"at"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/pagerduty/r/service#to_urgency Service#to_urgency}.
	ToUrgency *string `field:"optional" json:"toUrgency" yaml:"toUrgency"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/pagerduty/r/service#type Service#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

