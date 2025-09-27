// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package businessservice

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BusinessServiceConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.29.1/docs/resources/business_service#name BusinessService#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.29.1/docs/resources/business_service#description BusinessService#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.29.1/docs/resources/business_service#point_of_contact BusinessService#point_of_contact}.
	PointOfContact *string `field:"optional" json:"pointOfContact" yaml:"pointOfContact"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.29.1/docs/resources/business_service#team BusinessService#team}.
	Team *string `field:"optional" json:"team" yaml:"team"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.29.1/docs/resources/business_service#type BusinessService#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

