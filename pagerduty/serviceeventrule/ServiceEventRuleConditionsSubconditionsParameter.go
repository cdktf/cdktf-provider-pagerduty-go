package serviceeventrule


type ServiceEventRuleConditionsSubconditionsParameter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/2.14.6/docs/resources/service_event_rule#path ServiceEventRule#path}.
	Path *string `field:"optional" json:"path" yaml:"path"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/2.14.6/docs/resources/service_event_rule#value ServiceEventRule#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

