// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package eventorchestrationglobal


type EventOrchestrationGlobalSetRuleActionsExtraction struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.11.0/docs/resources/event_orchestration_global#target EventOrchestrationGlobal#target}.
	Target *string `field:"required" json:"target" yaml:"target"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.11.0/docs/resources/event_orchestration_global#regex EventOrchestrationGlobal#regex}.
	Regex *string `field:"optional" json:"regex" yaml:"regex"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.11.0/docs/resources/event_orchestration_global#source EventOrchestrationGlobal#source}.
	Source *string `field:"optional" json:"source" yaml:"source"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.11.0/docs/resources/event_orchestration_global#template EventOrchestrationGlobal#template}.
	Template *string `field:"optional" json:"template" yaml:"template"`
}

