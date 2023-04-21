package service


type ServiceAutoPauseNotificationsParameters struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/2.14.3/docs/resources/service#enabled Service#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/2.14.3/docs/resources/service#timeout Service#timeout}.
	Timeout *float64 `field:"optional" json:"timeout" yaml:"timeout"`
}

