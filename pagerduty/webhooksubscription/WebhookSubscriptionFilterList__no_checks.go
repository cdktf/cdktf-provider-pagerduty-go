// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package webhooksubscription

// Building without runtime type checking enabled, so all the below just return nil

func (w *jsiiProxy_WebhookSubscriptionFilterList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (w *jsiiProxy_WebhookSubscriptionFilterList) validateGetParameters(index *float64) error {
	return nil
}

func (w *jsiiProxy_WebhookSubscriptionFilterList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_WebhookSubscriptionFilterList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_WebhookSubscriptionFilterList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_WebhookSubscriptionFilterList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_WebhookSubscriptionFilterList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewWebhookSubscriptionFilterListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

