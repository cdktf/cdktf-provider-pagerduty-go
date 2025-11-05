// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package servicecustomfield

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_ServiceCustomFieldFieldOptionList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_ServiceCustomFieldFieldOptionList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_ServiceCustomFieldFieldOptionList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ServiceCustomFieldFieldOptionList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ServiceCustomFieldFieldOptionList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ServiceCustomFieldFieldOptionList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ServiceCustomFieldFieldOptionList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewServiceCustomFieldFieldOptionListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

