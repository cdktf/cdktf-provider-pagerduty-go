// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package eventorchestrationservice

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EventOrchestrationServiceSetRuleList) validateGetParameters(index *float64) error {
	return nil
}

func (e *jsiiProxy_EventOrchestrationServiceSetRuleList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_EventOrchestrationServiceSetRuleList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_EventOrchestrationServiceSetRuleList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_EventOrchestrationServiceSetRuleList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_EventOrchestrationServiceSetRuleList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewEventOrchestrationServiceSetRuleListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

