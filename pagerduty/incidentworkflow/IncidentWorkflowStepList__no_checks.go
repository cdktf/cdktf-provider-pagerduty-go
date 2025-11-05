// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package incidentworkflow

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IncidentWorkflowStepList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (i *jsiiProxy_IncidentWorkflowStepList) validateGetParameters(index *float64) error {
	return nil
}

func (i *jsiiProxy_IncidentWorkflowStepList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_IncidentWorkflowStepList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_IncidentWorkflowStepList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_IncidentWorkflowStepList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_IncidentWorkflowStepList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewIncidentWorkflowStepListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

