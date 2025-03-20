// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package eventorchestrationrouter


type EventOrchestrationRouterSetRuleActionsDynamicRouteTo struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.23.0/docs/resources/event_orchestration_router#lookup_by EventOrchestrationRouter#lookup_by}.
	LookupBy *string `field:"required" json:"lookupBy" yaml:"lookupBy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.23.0/docs/resources/event_orchestration_router#regex EventOrchestrationRouter#regex}.
	Regex *string `field:"required" json:"regex" yaml:"regex"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.23.0/docs/resources/event_orchestration_router#source EventOrchestrationRouter#source}.
	Source *string `field:"required" json:"source" yaml:"source"`
}

