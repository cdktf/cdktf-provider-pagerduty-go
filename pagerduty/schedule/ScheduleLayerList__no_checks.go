// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package schedule

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_ScheduleLayerList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_ScheduleLayerList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_ScheduleLayerList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ScheduleLayerList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ScheduleLayerList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ScheduleLayerList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ScheduleLayerList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewScheduleLayerListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

