// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider


type PagerdutyProviderUseAppOauthScopedToken struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.3.0/docs#pd_client_id PagerdutyProvider#pd_client_id}.
	PdClientId *string `field:"required" json:"pdClientId" yaml:"pdClientId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.3.0/docs#pd_client_secret PagerdutyProvider#pd_client_secret}.
	PdClientSecret *string `field:"required" json:"pdClientSecret" yaml:"pdClientSecret"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.3.0/docs#pd_subdomain PagerdutyProvider#pd_subdomain}.
	PdSubdomain *string `field:"required" json:"pdSubdomain" yaml:"pdSubdomain"`
}

