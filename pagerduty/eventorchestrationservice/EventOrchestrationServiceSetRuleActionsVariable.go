// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package eventorchestrationservice


type EventOrchestrationServiceSetRuleActionsVariable struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.30.0/docs/resources/event_orchestration_service#name EventOrchestrationService#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.30.0/docs/resources/event_orchestration_service#path EventOrchestrationService#path}.
	Path *string `field:"required" json:"path" yaml:"path"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.30.0/docs/resources/event_orchestration_service#type EventOrchestrationService#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.30.0/docs/resources/event_orchestration_service#value EventOrchestrationService#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

