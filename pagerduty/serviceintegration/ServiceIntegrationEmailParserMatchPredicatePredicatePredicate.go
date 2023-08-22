package serviceintegration


type ServiceIntegrationEmailParserMatchPredicatePredicatePredicate struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/2.16.0/docs/resources/service_integration#matcher ServiceIntegration#matcher}.
	Matcher *string `field:"required" json:"matcher" yaml:"matcher"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/2.16.0/docs/resources/service_integration#part ServiceIntegration#part}.
	Part *string `field:"required" json:"part" yaml:"part"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/2.16.0/docs/resources/service_integration#type ServiceIntegration#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}

