package serviceeventrule


type ServiceEventRuleConditionsSubconditions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/2.14.4/docs/resources/service_event_rule#operator ServiceEventRule#operator}.
	Operator *string `field:"optional" json:"operator" yaml:"operator"`
	// parameter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/2.14.4/docs/resources/service_event_rule#parameter ServiceEventRule#parameter}
	Parameter interface{} `field:"optional" json:"parameter" yaml:"parameter"`
}

