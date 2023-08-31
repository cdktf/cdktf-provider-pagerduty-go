// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package eventorchestrationunrouted


type EventOrchestrationUnroutedSetRuleCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/2.16.2/docs/resources/event_orchestration_unrouted#expression EventOrchestrationUnrouted#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
}

