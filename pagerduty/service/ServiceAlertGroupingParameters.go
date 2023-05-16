package service


type ServiceAlertGroupingParameters struct {
	// config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/2.14.5/docs/resources/service#config Service#config}
	Config *ServiceAlertGroupingParametersConfig `field:"optional" json:"config" yaml:"config"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/2.14.5/docs/resources/service#type Service#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

