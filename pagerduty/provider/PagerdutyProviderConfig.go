// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider


type PagerdutyProviderConfig struct {
	// Alias name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.12.2/docs#alias PagerdutyProvider#alias}
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.12.2/docs#api_url_override PagerdutyProvider#api_url_override}.
	ApiUrlOverride *string `field:"optional" json:"apiUrlOverride" yaml:"apiUrlOverride"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.12.2/docs#service_region PagerdutyProvider#service_region}.
	ServiceRegion *string `field:"optional" json:"serviceRegion" yaml:"serviceRegion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.12.2/docs#skip_credentials_validation PagerdutyProvider#skip_credentials_validation}.
	SkipCredentialsValidation interface{} `field:"optional" json:"skipCredentialsValidation" yaml:"skipCredentialsValidation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.12.2/docs#token PagerdutyProvider#token}.
	Token *string `field:"optional" json:"token" yaml:"token"`
	// use_app_oauth_scoped_token block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.12.2/docs#use_app_oauth_scoped_token PagerdutyProvider#use_app_oauth_scoped_token}
	UseAppOauthScopedToken *PagerdutyProviderUseAppOauthScopedToken `field:"optional" json:"useAppOauthScopedToken" yaml:"useAppOauthScopedToken"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.12.2/docs#user_token PagerdutyProvider#user_token}.
	UserToken *string `field:"optional" json:"userToken" yaml:"userToken"`
}

