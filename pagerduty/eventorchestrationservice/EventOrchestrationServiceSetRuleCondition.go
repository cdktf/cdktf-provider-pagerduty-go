// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package eventorchestrationservice


type EventOrchestrationServiceSetRuleCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.11.3/docs/resources/event_orchestration_service#expression EventOrchestrationService#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
}

