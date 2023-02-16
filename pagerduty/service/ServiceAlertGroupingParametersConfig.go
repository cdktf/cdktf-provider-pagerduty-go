package service


type ServiceAlertGroupingParametersConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/pagerduty/r/service#aggregate Service#aggregate}.
	Aggregate *string `field:"optional" json:"aggregate" yaml:"aggregate"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/pagerduty/r/service#fields Service#fields}.
	Fields *[]*string `field:"optional" json:"fields" yaml:"fields"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/pagerduty/r/service#timeout Service#timeout}.
	Timeout *float64 `field:"optional" json:"timeout" yaml:"timeout"`
}

