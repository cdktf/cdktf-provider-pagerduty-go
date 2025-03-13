// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package service

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ServiceConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.21.1/docs/resources/service#escalation_policy Service#escalation_policy}.
	EscalationPolicy *string `field:"required" json:"escalationPolicy" yaml:"escalationPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.21.1/docs/resources/service#name Service#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.21.1/docs/resources/service#acknowledgement_timeout Service#acknowledgement_timeout}.
	AcknowledgementTimeout *string `field:"optional" json:"acknowledgementTimeout" yaml:"acknowledgementTimeout"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.21.1/docs/resources/service#alert_creation Service#alert_creation}.
	AlertCreation *string `field:"optional" json:"alertCreation" yaml:"alertCreation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.21.1/docs/resources/service#alert_grouping Service#alert_grouping}.
	AlertGrouping *string `field:"optional" json:"alertGrouping" yaml:"alertGrouping"`
	// alert_grouping_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.21.1/docs/resources/service#alert_grouping_parameters Service#alert_grouping_parameters}
	AlertGroupingParameters *ServiceAlertGroupingParameters `field:"optional" json:"alertGroupingParameters" yaml:"alertGroupingParameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.21.1/docs/resources/service#alert_grouping_timeout Service#alert_grouping_timeout}.
	AlertGroupingTimeout *string `field:"optional" json:"alertGroupingTimeout" yaml:"alertGroupingTimeout"`
	// auto_pause_notifications_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.21.1/docs/resources/service#auto_pause_notifications_parameters Service#auto_pause_notifications_parameters}
	AutoPauseNotificationsParameters *ServiceAutoPauseNotificationsParameters `field:"optional" json:"autoPauseNotificationsParameters" yaml:"autoPauseNotificationsParameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.21.1/docs/resources/service#auto_resolve_timeout Service#auto_resolve_timeout}.
	AutoResolveTimeout *string `field:"optional" json:"autoResolveTimeout" yaml:"autoResolveTimeout"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.21.1/docs/resources/service#description Service#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.21.1/docs/resources/service#id Service#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// incident_urgency_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.21.1/docs/resources/service#incident_urgency_rule Service#incident_urgency_rule}
	IncidentUrgencyRule *ServiceIncidentUrgencyRule `field:"optional" json:"incidentUrgencyRule" yaml:"incidentUrgencyRule"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.21.1/docs/resources/service#response_play Service#response_play}.
	ResponsePlay *string `field:"optional" json:"responsePlay" yaml:"responsePlay"`
	// scheduled_actions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.21.1/docs/resources/service#scheduled_actions Service#scheduled_actions}
	ScheduledActions interface{} `field:"optional" json:"scheduledActions" yaml:"scheduledActions"`
	// support_hours block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.21.1/docs/resources/service#support_hours Service#support_hours}
	SupportHours *ServiceSupportHours `field:"optional" json:"supportHours" yaml:"supportHours"`
}

