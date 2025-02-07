// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package userhandoffnotificationrule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type UserHandoffNotificationRuleConfig struct {
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
	// The number of minutes before the handoff to notify the user. Must be greater than or equal to 0.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.20.0/docs/resources/user_handoff_notification_rule#notify_advance_in_minutes UserHandoffNotificationRule#notify_advance_in_minutes}
	NotifyAdvanceInMinutes *float64 `field:"required" json:"notifyAdvanceInMinutes" yaml:"notifyAdvanceInMinutes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.20.0/docs/resources/user_handoff_notification_rule#user_id UserHandoffNotificationRule#user_id}.
	UserId *string `field:"required" json:"userId" yaml:"userId"`
	// contact_method block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.20.0/docs/resources/user_handoff_notification_rule#contact_method UserHandoffNotificationRule#contact_method}
	ContactMethod interface{} `field:"optional" json:"contactMethod" yaml:"contactMethod"`
	// The type of handoff to notify for. Possible values are 'both', 'oncall', 'offcall'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.20.0/docs/resources/user_handoff_notification_rule#handoff_type UserHandoffNotificationRule#handoff_type}
	HandoffType *string `field:"optional" json:"handoffType" yaml:"handoffType"`
}

