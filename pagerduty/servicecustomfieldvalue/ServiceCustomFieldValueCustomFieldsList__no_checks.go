// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package servicecustomfieldvalue

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_ServiceCustomFieldValueCustomFieldsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_ServiceCustomFieldValueCustomFieldsList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_ServiceCustomFieldValueCustomFieldsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ServiceCustomFieldValueCustomFieldsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ServiceCustomFieldValueCustomFieldsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ServiceCustomFieldValueCustomFieldsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ServiceCustomFieldValueCustomFieldsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewServiceCustomFieldValueCustomFieldsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

