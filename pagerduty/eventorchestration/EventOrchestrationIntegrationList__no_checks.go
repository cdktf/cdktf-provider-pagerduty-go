// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package eventorchestration

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EventOrchestrationIntegrationList) validateGetParameters(index *float64) error {
	return nil
}

func (e *jsiiProxy_EventOrchestrationIntegrationList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_EventOrchestrationIntegrationList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_EventOrchestrationIntegrationList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_EventOrchestrationIntegrationList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_EventOrchestrationIntegrationList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewEventOrchestrationIntegrationListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

