// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package eventorchestrationrouter

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EventOrchestrationRouterConfig struct {
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
	// catch_all block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.28.1/docs/resources/event_orchestration_router#catch_all EventOrchestrationRouter#catch_all}
	CatchAll *EventOrchestrationRouterCatchAll `field:"required" json:"catchAll" yaml:"catchAll"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.28.1/docs/resources/event_orchestration_router#event_orchestration EventOrchestrationRouter#event_orchestration}.
	EventOrchestration *string `field:"required" json:"eventOrchestration" yaml:"eventOrchestration"`
	// set block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.28.1/docs/resources/event_orchestration_router#set EventOrchestrationRouter#set}
	Set *EventOrchestrationRouterSet `field:"required" json:"set" yaml:"set"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.28.1/docs/resources/event_orchestration_router#id EventOrchestrationRouter#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

