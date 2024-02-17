// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package serviceeventrule


type ServiceEventRuleVariableParameters struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.8.0/docs/resources/service_event_rule#path ServiceEventRule#path}.
	Path *string `field:"optional" json:"path" yaml:"path"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.8.0/docs/resources/service_event_rule#value ServiceEventRule#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

