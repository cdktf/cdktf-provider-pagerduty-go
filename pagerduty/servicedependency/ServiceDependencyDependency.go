// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package servicedependency


type ServiceDependencyDependency struct {
	// dependent_service block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.10.1/docs/resources/service_dependency#dependent_service ServiceDependency#dependent_service}
	DependentService interface{} `field:"required" json:"dependentService" yaml:"dependentService"`
	// supporting_service block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.10.1/docs/resources/service_dependency#supporting_service ServiceDependency#supporting_service}
	SupportingService interface{} `field:"required" json:"supportingService" yaml:"supportingService"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.10.1/docs/resources/service_dependency#type ServiceDependency#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

