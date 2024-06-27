// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package eventorchestrationservicecachevariable

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EventOrchestrationServiceCacheVariableConfig struct {
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
	// configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.14.4/docs/resources/event_orchestration_service_cache_variable#configuration EventOrchestrationServiceCacheVariable#configuration}
	Configuration *EventOrchestrationServiceCacheVariableConfiguration `field:"required" json:"configuration" yaml:"configuration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.14.4/docs/resources/event_orchestration_service_cache_variable#name EventOrchestrationServiceCacheVariable#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.14.4/docs/resources/event_orchestration_service_cache_variable#service EventOrchestrationServiceCacheVariable#service}.
	Service *string `field:"required" json:"service" yaml:"service"`
	// condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.14.4/docs/resources/event_orchestration_service_cache_variable#condition EventOrchestrationServiceCacheVariable#condition}
	Condition interface{} `field:"optional" json:"condition" yaml:"condition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.14.4/docs/resources/event_orchestration_service_cache_variable#disabled EventOrchestrationServiceCacheVariable#disabled}.
	Disabled interface{} `field:"optional" json:"disabled" yaml:"disabled"`
}

