// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package webhooksubscription


type WebhookSubscriptionDeliveryMethod struct {
	// custom_header block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.12.1/docs/resources/webhook_subscription#custom_header WebhookSubscription#custom_header}
	CustomHeader interface{} `field:"optional" json:"customHeader" yaml:"customHeader"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.12.1/docs/resources/webhook_subscription#temporarily_disabled WebhookSubscription#temporarily_disabled}.
	TemporarilyDisabled interface{} `field:"optional" json:"temporarilyDisabled" yaml:"temporarilyDisabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.12.1/docs/resources/webhook_subscription#type WebhookSubscription#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.12.1/docs/resources/webhook_subscription#url WebhookSubscription#url}.
	Url *string `field:"optional" json:"url" yaml:"url"`
}

