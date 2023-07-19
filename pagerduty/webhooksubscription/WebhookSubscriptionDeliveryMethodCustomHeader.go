package webhooksubscription


type WebhookSubscriptionDeliveryMethodCustomHeader struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/2.15.1/docs/resources/webhook_subscription#name WebhookSubscription#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/2.15.1/docs/resources/webhook_subscription#value WebhookSubscription#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

