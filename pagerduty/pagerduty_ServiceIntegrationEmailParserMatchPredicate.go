// Prebuilt pagerduty Provider for Terraform CDK (cdktf)
package pagerduty


type ServiceIntegrationEmailParserMatchPredicate struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/pagerduty/r/service_integration#type ServiceIntegration#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// predicate block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/pagerduty/r/service_integration#predicate ServiceIntegration#predicate}
	Predicate interface{} `field:"optional" json:"predicate" yaml:"predicate"`
}

