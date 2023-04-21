package service


type ServiceAlertGroupingParametersConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/2.14.3/docs/resources/service#aggregate Service#aggregate}.
	Aggregate *string `field:"optional" json:"aggregate" yaml:"aggregate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/2.14.3/docs/resources/service#fields Service#fields}.
	Fields *[]*string `field:"optional" json:"fields" yaml:"fields"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/2.14.3/docs/resources/service#timeout Service#timeout}.
	Timeout *float64 `field:"optional" json:"timeout" yaml:"timeout"`
}

