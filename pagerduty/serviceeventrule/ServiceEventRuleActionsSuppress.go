package serviceeventrule


type ServiceEventRuleActionsSuppress struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/2.14.5/docs/resources/service_event_rule#threshold_time_amount ServiceEventRule#threshold_time_amount}.
	ThresholdTimeAmount *float64 `field:"optional" json:"thresholdTimeAmount" yaml:"thresholdTimeAmount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/2.14.5/docs/resources/service_event_rule#threshold_time_unit ServiceEventRule#threshold_time_unit}.
	ThresholdTimeUnit *string `field:"optional" json:"thresholdTimeUnit" yaml:"thresholdTimeUnit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/2.14.5/docs/resources/service_event_rule#threshold_value ServiceEventRule#threshold_value}.
	ThresholdValue *float64 `field:"optional" json:"thresholdValue" yaml:"thresholdValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/2.14.5/docs/resources/service_event_rule#value ServiceEventRule#value}.
	Value interface{} `field:"optional" json:"value" yaml:"value"`
}

