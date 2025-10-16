// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package tagassignment

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type TagAssignmentConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.30.3/docs/resources/tag_assignment#entity_id TagAssignment#entity_id}.
	EntityId *string `field:"required" json:"entityId" yaml:"entityId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.30.3/docs/resources/tag_assignment#entity_type TagAssignment#entity_type}.
	EntityType *string `field:"required" json:"entityType" yaml:"entityType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.30.3/docs/resources/tag_assignment#tag_id TagAssignment#tag_id}.
	TagId *string `field:"required" json:"tagId" yaml:"tagId"`
}

