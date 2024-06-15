// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package webhooksubscription


type WebhookSubscriptionDeliveryMethodCustomHeader struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.14.0/docs/resources/webhook_subscription#name WebhookSubscription#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.14.0/docs/resources/webhook_subscription#value WebhookSubscription#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

