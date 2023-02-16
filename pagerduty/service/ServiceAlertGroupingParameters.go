package service


type ServiceAlertGroupingParameters struct {
	// config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/pagerduty/r/service#config Service#config}
	Config *ServiceAlertGroupingParametersConfig `field:"optional" json:"config" yaml:"config"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/pagerduty/r/service#type Service#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

