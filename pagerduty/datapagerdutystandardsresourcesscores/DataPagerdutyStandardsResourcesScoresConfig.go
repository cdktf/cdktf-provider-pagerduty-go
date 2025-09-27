// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datapagerdutystandardsresourcesscores

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataPagerdutyStandardsResourcesScoresConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.29.1/docs/data-sources/standards_resources_scores#ids DataPagerdutyStandardsResourcesScores#ids}.
	Ids *[]*string `field:"required" json:"ids" yaml:"ids"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.29.1/docs/data-sources/standards_resources_scores#resource_type DataPagerdutyStandardsResourcesScores#resource_type}.
	ResourceType *string `field:"required" json:"resourceType" yaml:"resourceType"`
}

