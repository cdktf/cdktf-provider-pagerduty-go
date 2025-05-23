// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package teammembership

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type TeamMembershipConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.25.2/docs/resources/team_membership#team_id TeamMembership#team_id}.
	TeamId *string `field:"required" json:"teamId" yaml:"teamId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.25.2/docs/resources/team_membership#user_id TeamMembership#user_id}.
	UserId *string `field:"required" json:"userId" yaml:"userId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.25.2/docs/resources/team_membership#role TeamMembership#role}.
	Role *string `field:"optional" json:"role" yaml:"role"`
}

