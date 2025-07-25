// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package usernotificationrule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type UserNotificationRuleConfig struct {
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
	// contact_method block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.27.2/docs/resources/user_notification_rule#contact_method UserNotificationRule#contact_method}
	ContactMethod *UserNotificationRuleContactMethod `field:"required" json:"contactMethod" yaml:"contactMethod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.27.2/docs/resources/user_notification_rule#start_delay_in_minutes UserNotificationRule#start_delay_in_minutes}.
	StartDelayInMinutes *float64 `field:"required" json:"startDelayInMinutes" yaml:"startDelayInMinutes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.27.2/docs/resources/user_notification_rule#urgency UserNotificationRule#urgency}.
	Urgency *string `field:"required" json:"urgency" yaml:"urgency"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.27.2/docs/resources/user_notification_rule#user_id UserNotificationRule#user_id}.
	UserId *string `field:"required" json:"userId" yaml:"userId"`
}

