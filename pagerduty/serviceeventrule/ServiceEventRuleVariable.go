package serviceeventrule


type ServiceEventRuleVariable struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/2.15.3/docs/resources/service_event_rule#name ServiceEventRule#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/2.15.3/docs/resources/service_event_rule#parameters ServiceEventRule#parameters}
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/2.15.3/docs/resources/service_event_rule#type ServiceEventRule#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

