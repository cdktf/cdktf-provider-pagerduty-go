// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package eventorchestrationglobal


type EventOrchestrationGlobalSetRuleCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.30.5/docs/resources/event_orchestration_global#expression EventOrchestrationGlobal#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
}

