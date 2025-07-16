// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package serviceintegration


type ServiceIntegrationEmailParserMatchPredicatePredicate struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.27.1/docs/resources/service_integration#type ServiceIntegration#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.27.1/docs/resources/service_integration#matcher ServiceIntegration#matcher}.
	Matcher *string `field:"optional" json:"matcher" yaml:"matcher"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.27.1/docs/resources/service_integration#part ServiceIntegration#part}.
	Part *string `field:"optional" json:"part" yaml:"part"`
	// predicate block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.27.1/docs/resources/service_integration#predicate ServiceIntegration#predicate}
	Predicate interface{} `field:"optional" json:"predicate" yaml:"predicate"`
}

