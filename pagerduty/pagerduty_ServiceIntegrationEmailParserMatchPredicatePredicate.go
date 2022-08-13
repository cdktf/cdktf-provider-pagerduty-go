// Prebuilt pagerduty Provider for Terraform CDK (cdktf)
package pagerduty


type ServiceIntegrationEmailParserMatchPredicatePredicate struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/pagerduty/r/service_integration#type ServiceIntegration#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/pagerduty/r/service_integration#matcher ServiceIntegration#matcher}.
	Matcher *string `field:"optional" json:"matcher" yaml:"matcher"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/pagerduty/r/service_integration#part ServiceIntegration#part}.
	Part *string `field:"optional" json:"part" yaml:"part"`
	// predicate block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/pagerduty/r/service_integration#predicate ServiceIntegration#predicate}
	Predicate interface{} `field:"optional" json:"predicate" yaml:"predicate"`
}

