// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package eventorchestrationrouter


type EventOrchestrationRouterSetRuleCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.29.0/docs/resources/event_orchestration_router#expression EventOrchestrationRouter#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
}

