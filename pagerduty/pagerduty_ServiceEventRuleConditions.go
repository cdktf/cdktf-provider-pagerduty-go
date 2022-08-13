// Prebuilt pagerduty Provider for Terraform CDK (cdktf)
package pagerduty


type ServiceEventRuleConditions struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/pagerduty/r/service_event_rule#operator ServiceEventRule#operator}.
	Operator *string `field:"optional" json:"operator" yaml:"operator"`
	// subconditions block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/pagerduty/r/service_event_rule#subconditions ServiceEventRule#subconditions}
	Subconditions interface{} `field:"optional" json:"subconditions" yaml:"subconditions"`
}

