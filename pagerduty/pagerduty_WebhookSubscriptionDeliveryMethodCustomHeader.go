// Prebuilt pagerduty Provider for Terraform CDK (cdktf)
package pagerduty


type WebhookSubscriptionDeliveryMethodCustomHeader struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/pagerduty/r/webhook_subscription#name WebhookSubscription#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/pagerduty/r/webhook_subscription#value WebhookSubscription#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

