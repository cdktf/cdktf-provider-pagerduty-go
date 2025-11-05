// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package serviceeventrule

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_ServiceEventRuleVariableList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_ServiceEventRuleVariableList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_ServiceEventRuleVariableList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ServiceEventRuleVariableList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ServiceEventRuleVariableList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ServiceEventRuleVariableList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ServiceEventRuleVariableList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewServiceEventRuleVariableListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

