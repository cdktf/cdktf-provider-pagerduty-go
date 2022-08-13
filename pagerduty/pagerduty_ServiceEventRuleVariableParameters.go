// Prebuilt pagerduty Provider for Terraform CDK (cdktf)
package pagerduty


type ServiceEventRuleVariableParameters struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/pagerduty/r/service_event_rule#path ServiceEventRule#path}.
	Path *string `field:"optional" json:"path" yaml:"path"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/pagerduty/r/service_event_rule#value ServiceEventRule#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

