// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package service


type ServiceAlertGroupingParameters struct {
	// config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.15.1/docs/resources/service#config Service#config}
	Config *ServiceAlertGroupingParametersConfig `field:"optional" json:"config" yaml:"config"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.15.1/docs/resources/service#type Service#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

