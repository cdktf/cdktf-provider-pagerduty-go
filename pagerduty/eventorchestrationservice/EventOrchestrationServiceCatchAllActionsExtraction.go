// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package eventorchestrationservice


type EventOrchestrationServiceCatchAllActionsExtraction struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.23.1/docs/resources/event_orchestration_service#target EventOrchestrationService#target}.
	Target *string `field:"required" json:"target" yaml:"target"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.23.1/docs/resources/event_orchestration_service#regex EventOrchestrationService#regex}.
	Regex *string `field:"optional" json:"regex" yaml:"regex"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.23.1/docs/resources/event_orchestration_service#source EventOrchestrationService#source}.
	Source *string `field:"optional" json:"source" yaml:"source"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.23.1/docs/resources/event_orchestration_service#template EventOrchestrationService#template}.
	Template *string `field:"optional" json:"template" yaml:"template"`
}

