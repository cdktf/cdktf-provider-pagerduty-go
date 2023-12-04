// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package rulesetrule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type RulesetRuleConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.2.2/docs/resources/ruleset_rule#ruleset RulesetRule#ruleset}.
	Ruleset *string `field:"required" json:"ruleset" yaml:"ruleset"`
	// actions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.2.2/docs/resources/ruleset_rule#actions RulesetRule#actions}
	Actions *RulesetRuleActions `field:"optional" json:"actions" yaml:"actions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.2.2/docs/resources/ruleset_rule#catch_all RulesetRule#catch_all}.
	CatchAll interface{} `field:"optional" json:"catchAll" yaml:"catchAll"`
	// conditions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.2.2/docs/resources/ruleset_rule#conditions RulesetRule#conditions}
	Conditions *RulesetRuleConditions `field:"optional" json:"conditions" yaml:"conditions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.2.2/docs/resources/ruleset_rule#disabled RulesetRule#disabled}.
	Disabled interface{} `field:"optional" json:"disabled" yaml:"disabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.2.2/docs/resources/ruleset_rule#id RulesetRule#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.2.2/docs/resources/ruleset_rule#position RulesetRule#position}.
	Position *float64 `field:"optional" json:"position" yaml:"position"`
	// time_frame block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.2.2/docs/resources/ruleset_rule#time_frame RulesetRule#time_frame}
	TimeFrame *RulesetRuleTimeFrame `field:"optional" json:"timeFrame" yaml:"timeFrame"`
	// variable block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.2.2/docs/resources/ruleset_rule#variable RulesetRule#variable}
	Variable interface{} `field:"optional" json:"variable" yaml:"variable"`
}

