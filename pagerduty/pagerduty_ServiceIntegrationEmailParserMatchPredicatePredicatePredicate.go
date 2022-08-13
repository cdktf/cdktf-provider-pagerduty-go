// Prebuilt pagerduty Provider for Terraform CDK (cdktf)
package pagerduty


type ServiceIntegrationEmailParserMatchPredicatePredicatePredicate struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/pagerduty/r/service_integration#matcher ServiceIntegration#matcher}.
	Matcher *string `field:"required" json:"matcher" yaml:"matcher"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/pagerduty/r/service_integration#part ServiceIntegration#part}.
	Part *string `field:"required" json:"part" yaml:"part"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/pagerduty/r/service_integration#type ServiceIntegration#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}

