package serviceintegration


type ServiceIntegrationEmailParser struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/2.15.1/docs/resources/service_integration#action ServiceIntegration#action}.
	Action *string `field:"required" json:"action" yaml:"action"`
	// match_predicate block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/2.15.1/docs/resources/service_integration#match_predicate ServiceIntegration#match_predicate}
	MatchPredicate *ServiceIntegrationEmailParserMatchPredicate `field:"required" json:"matchPredicate" yaml:"matchPredicate"`
	// value_extractor block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/2.15.1/docs/resources/service_integration#value_extractor ServiceIntegration#value_extractor}
	ValueExtractor interface{} `field:"optional" json:"valueExtractor" yaml:"valueExtractor"`
}

