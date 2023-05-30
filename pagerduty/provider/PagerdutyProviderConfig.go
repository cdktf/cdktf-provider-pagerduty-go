package provider


type PagerdutyProviderConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/2.14.6/docs#token PagerdutyProvider#token}.
	Token *string `field:"required" json:"token" yaml:"token"`
	// Alias name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/2.14.6/docs#alias PagerdutyProvider#alias}
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/2.14.6/docs#api_url_override PagerdutyProvider#api_url_override}.
	ApiUrlOverride *string `field:"optional" json:"apiUrlOverride" yaml:"apiUrlOverride"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/2.14.6/docs#service_region PagerdutyProvider#service_region}.
	ServiceRegion *string `field:"optional" json:"serviceRegion" yaml:"serviceRegion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/2.14.6/docs#skip_credentials_validation PagerdutyProvider#skip_credentials_validation}.
	SkipCredentialsValidation interface{} `field:"optional" json:"skipCredentialsValidation" yaml:"skipCredentialsValidation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/2.14.6/docs#user_token PagerdutyProvider#user_token}.
	UserToken *string `field:"optional" json:"userToken" yaml:"userToken"`
}

