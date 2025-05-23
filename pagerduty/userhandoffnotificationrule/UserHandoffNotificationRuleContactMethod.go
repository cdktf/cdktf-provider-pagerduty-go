// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package userhandoffnotificationrule


type UserHandoffNotificationRuleContactMethod struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.25.2/docs/resources/user_handoff_notification_rule#id UserHandoffNotificationRule#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
	// The type of contact method to notify for. Possible values are 'email_contact_method', 'email_contact_method_reference', 'phone_contact_method', 'phone_contact_method_reference', 'push_notification_contact_method', 'push_notification_contact_method_reference', 'sms_contact_method', 'sms_contact_method_reference'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.25.2/docs/resources/user_handoff_notification_rule#type UserHandoffNotificationRule#type}
	Type *string `field:"required" json:"type" yaml:"type"`
}

