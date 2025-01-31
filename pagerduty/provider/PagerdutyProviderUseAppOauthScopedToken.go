// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider


type PagerdutyProviderUseAppOauthScopedToken struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.19.4/docs#pd_client_id PagerdutyProvider#pd_client_id}.
	PdClientId *string `field:"optional" json:"pdClientId" yaml:"pdClientId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.19.4/docs#pd_client_secret PagerdutyProvider#pd_client_secret}.
	PdClientSecret *string `field:"optional" json:"pdClientSecret" yaml:"pdClientSecret"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.19.4/docs#pd_subdomain PagerdutyProvider#pd_subdomain}.
	PdSubdomain *string `field:"optional" json:"pdSubdomain" yaml:"pdSubdomain"`
}

