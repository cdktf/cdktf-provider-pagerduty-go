// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package serviceeventrule


type ServiceEventRuleConditions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.18.1/docs/resources/service_event_rule#operator ServiceEventRule#operator}.
	Operator *string `field:"optional" json:"operator" yaml:"operator"`
	// subconditions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.18.1/docs/resources/service_event_rule#subconditions ServiceEventRule#subconditions}
	Subconditions interface{} `field:"optional" json:"subconditions" yaml:"subconditions"`
}

