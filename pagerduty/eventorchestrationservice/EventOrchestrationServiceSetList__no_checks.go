// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package eventorchestrationservice

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EventOrchestrationServiceSetList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (e *jsiiProxy_EventOrchestrationServiceSetList) validateGetParameters(index *float64) error {
	return nil
}

func (e *jsiiProxy_EventOrchestrationServiceSetList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_EventOrchestrationServiceSetList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_EventOrchestrationServiceSetList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_EventOrchestrationServiceSetList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_EventOrchestrationServiceSetList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewEventOrchestrationServiceSetListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

