// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package escalationpolicy

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EscalationPolicyConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.17.2/docs/resources/escalation_policy#name EscalationPolicy#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.17.2/docs/resources/escalation_policy#rule EscalationPolicy#rule}
	Rule interface{} `field:"required" json:"rule" yaml:"rule"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.17.2/docs/resources/escalation_policy#description EscalationPolicy#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.17.2/docs/resources/escalation_policy#id EscalationPolicy#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.17.2/docs/resources/escalation_policy#num_loops EscalationPolicy#num_loops}.
	NumLoops *float64 `field:"optional" json:"numLoops" yaml:"numLoops"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.17.2/docs/resources/escalation_policy#teams EscalationPolicy#teams}.
	Teams *[]*string `field:"optional" json:"teams" yaml:"teams"`
}

