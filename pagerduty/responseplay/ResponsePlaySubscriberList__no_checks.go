// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package responseplay

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_ResponsePlaySubscriberList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (r *jsiiProxy_ResponsePlaySubscriberList) validateGetParameters(index *float64) error {
	return nil
}

func (r *jsiiProxy_ResponsePlaySubscriberList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ResponsePlaySubscriberList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ResponsePlaySubscriberList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ResponsePlaySubscriberList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ResponsePlaySubscriberList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewResponsePlaySubscriberListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

