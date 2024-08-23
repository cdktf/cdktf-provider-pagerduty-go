// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package serviceeventrule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ServiceEventRuleConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.15.5/docs/resources/service_event_rule#service ServiceEventRule#service}.
	Service *string `field:"required" json:"service" yaml:"service"`
	// actions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.15.5/docs/resources/service_event_rule#actions ServiceEventRule#actions}
	Actions *ServiceEventRuleActions `field:"optional" json:"actions" yaml:"actions"`
	// conditions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.15.5/docs/resources/service_event_rule#conditions ServiceEventRule#conditions}
	Conditions *ServiceEventRuleConditions `field:"optional" json:"conditions" yaml:"conditions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.15.5/docs/resources/service_event_rule#disabled ServiceEventRule#disabled}.
	Disabled interface{} `field:"optional" json:"disabled" yaml:"disabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.15.5/docs/resources/service_event_rule#id ServiceEventRule#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.15.5/docs/resources/service_event_rule#position ServiceEventRule#position}.
	Position *float64 `field:"optional" json:"position" yaml:"position"`
	// time_frame block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.15.5/docs/resources/service_event_rule#time_frame ServiceEventRule#time_frame}
	TimeFrame *ServiceEventRuleTimeFrame `field:"optional" json:"timeFrame" yaml:"timeFrame"`
	// variable block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.15.5/docs/resources/service_event_rule#variable ServiceEventRule#variable}
	Variable interface{} `field:"optional" json:"variable" yaml:"variable"`
}

