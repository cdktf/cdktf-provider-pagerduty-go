// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package serviceintegration


type ServiceIntegrationEmailParserValueExtractor struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.18.2/docs/resources/service_integration#part ServiceIntegration#part}.
	Part *string `field:"required" json:"part" yaml:"part"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.18.2/docs/resources/service_integration#type ServiceIntegration#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.18.2/docs/resources/service_integration#value_name ServiceIntegration#value_name}.
	ValueName *string `field:"required" json:"valueName" yaml:"valueName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.18.2/docs/resources/service_integration#ends_before ServiceIntegration#ends_before}.
	EndsBefore *string `field:"optional" json:"endsBefore" yaml:"endsBefore"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.18.2/docs/resources/service_integration#regex ServiceIntegration#regex}.
	Regex *string `field:"optional" json:"regex" yaml:"regex"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.18.2/docs/resources/service_integration#starts_after ServiceIntegration#starts_after}.
	StartsAfter *string `field:"optional" json:"startsAfter" yaml:"startsAfter"`
}

