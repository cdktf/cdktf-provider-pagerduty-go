// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package eventorchestrationservicecachevariable


type EventOrchestrationServiceCacheVariableConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.23.1/docs/resources/event_orchestration_service_cache_variable#type EventOrchestrationServiceCacheVariable#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.23.1/docs/resources/event_orchestration_service_cache_variable#data_type EventOrchestrationServiceCacheVariable#data_type}.
	DataType *string `field:"optional" json:"dataType" yaml:"dataType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.23.1/docs/resources/event_orchestration_service_cache_variable#regex EventOrchestrationServiceCacheVariable#regex}.
	Regex *string `field:"optional" json:"regex" yaml:"regex"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.23.1/docs/resources/event_orchestration_service_cache_variable#source EventOrchestrationServiceCacheVariable#source}.
	Source *string `field:"optional" json:"source" yaml:"source"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.23.1/docs/resources/event_orchestration_service_cache_variable#ttl_seconds EventOrchestrationServiceCacheVariable#ttl_seconds}.
	TtlSeconds *float64 `field:"optional" json:"ttlSeconds" yaml:"ttlSeconds"`
}

