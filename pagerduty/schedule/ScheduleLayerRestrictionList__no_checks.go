// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package schedule

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_ScheduleLayerRestrictionList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_ScheduleLayerRestrictionList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_ScheduleLayerRestrictionList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ScheduleLayerRestrictionList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ScheduleLayerRestrictionList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ScheduleLayerRestrictionList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ScheduleLayerRestrictionList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewScheduleLayerRestrictionListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

