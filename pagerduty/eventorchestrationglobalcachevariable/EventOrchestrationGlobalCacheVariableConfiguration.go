// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package eventorchestrationglobalcachevariable


type EventOrchestrationGlobalCacheVariableConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.21.1/docs/resources/event_orchestration_global_cache_variable#type EventOrchestrationGlobalCacheVariable#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.21.1/docs/resources/event_orchestration_global_cache_variable#regex EventOrchestrationGlobalCacheVariable#regex}.
	Regex *string `field:"optional" json:"regex" yaml:"regex"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.21.1/docs/resources/event_orchestration_global_cache_variable#source EventOrchestrationGlobalCacheVariable#source}.
	Source *string `field:"optional" json:"source" yaml:"source"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.21.1/docs/resources/event_orchestration_global_cache_variable#ttl_seconds EventOrchestrationGlobalCacheVariable#ttl_seconds}.
	TtlSeconds *float64 `field:"optional" json:"ttlSeconds" yaml:"ttlSeconds"`
}

