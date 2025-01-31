// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package serviceintegration


type ServiceIntegrationEmailParserMatchPredicate struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.19.4/docs/resources/service_integration#type ServiceIntegration#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// predicate block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.19.4/docs/resources/service_integration#predicate ServiceIntegration#predicate}
	Predicate interface{} `field:"optional" json:"predicate" yaml:"predicate"`
}

