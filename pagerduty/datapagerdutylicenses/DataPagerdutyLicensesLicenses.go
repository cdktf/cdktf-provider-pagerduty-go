// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datapagerdutylicenses


type DataPagerdutyLicensesLicenses struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.4.1/docs/data-sources/licenses#allocations_available DataPagerdutyLicenses#allocations_available}.
	AllocationsAvailable *float64 `field:"optional" json:"allocationsAvailable" yaml:"allocationsAvailable"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.4.1/docs/data-sources/licenses#current_value DataPagerdutyLicenses#current_value}.
	CurrentValue *float64 `field:"optional" json:"currentValue" yaml:"currentValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.4.1/docs/data-sources/licenses#description DataPagerdutyLicenses#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.4.1/docs/data-sources/licenses#html_url DataPagerdutyLicenses#html_url}.
	HtmlUrl *string `field:"optional" json:"htmlUrl" yaml:"htmlUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.4.1/docs/data-sources/licenses#id DataPagerdutyLicenses#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.4.1/docs/data-sources/licenses#name DataPagerdutyLicenses#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.4.1/docs/data-sources/licenses#role_group DataPagerdutyLicenses#role_group}.
	RoleGroup *string `field:"optional" json:"roleGroup" yaml:"roleGroup"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.4.1/docs/data-sources/licenses#self DataPagerdutyLicenses#self}.
	SelfAttribute *string `field:"optional" json:"selfAttribute" yaml:"selfAttribute"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.4.1/docs/data-sources/licenses#summary DataPagerdutyLicenses#summary}.
	Summary *string `field:"optional" json:"summary" yaml:"summary"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.4.1/docs/data-sources/licenses#type DataPagerdutyLicenses#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.4.1/docs/data-sources/licenses#valid_roles DataPagerdutyLicenses#valid_roles}.
	ValidRoles *[]*string `field:"optional" json:"validRoles" yaml:"validRoles"`
}

