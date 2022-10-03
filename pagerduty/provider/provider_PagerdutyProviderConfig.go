package provider


type PagerdutyProviderConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/pagerduty#token PagerdutyProvider#token}.
	Token *string `field:"required" json:"token" yaml:"token"`
	// Alias name.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/pagerduty#alias PagerdutyProvider#alias}
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/pagerduty#api_url_override PagerdutyProvider#api_url_override}.
	ApiUrlOverride *string `field:"optional" json:"apiUrlOverride" yaml:"apiUrlOverride"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/pagerduty#service_region PagerdutyProvider#service_region}.
	ServiceRegion *string `field:"optional" json:"serviceRegion" yaml:"serviceRegion"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/pagerduty#skip_credentials_validation PagerdutyProvider#skip_credentials_validation}.
	SkipCredentialsValidation interface{} `field:"optional" json:"skipCredentialsValidation" yaml:"skipCredentialsValidation"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/pagerduty#user_token PagerdutyProvider#user_token}.
	UserToken *string `field:"optional" json:"userToken" yaml:"userToken"`
}

