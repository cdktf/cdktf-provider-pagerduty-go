// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package eventorchestration

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EventOrchestrationIntegrationParametersList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (e *jsiiProxy_EventOrchestrationIntegrationParametersList) validateGetParameters(index *float64) error {
	return nil
}

func (e *jsiiProxy_EventOrchestrationIntegrationParametersList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_EventOrchestrationIntegrationParametersList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_EventOrchestrationIntegrationParametersList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_EventOrchestrationIntegrationParametersList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewEventOrchestrationIntegrationParametersListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

