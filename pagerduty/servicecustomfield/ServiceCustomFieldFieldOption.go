// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package servicecustomfield


type ServiceCustomFieldFieldOption struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.26.0/docs/resources/service_custom_field#data_type ServiceCustomField#data_type}.
	DataType *string `field:"required" json:"dataType" yaml:"dataType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.26.0/docs/resources/service_custom_field#value ServiceCustomField#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

