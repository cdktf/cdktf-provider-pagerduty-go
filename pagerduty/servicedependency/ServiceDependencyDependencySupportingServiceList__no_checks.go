// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package servicedependency

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_ServiceDependencyDependencySupportingServiceList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_ServiceDependencyDependencySupportingServiceList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_ServiceDependencyDependencySupportingServiceList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ServiceDependencyDependencySupportingServiceList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ServiceDependencyDependencySupportingServiceList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ServiceDependencyDependencySupportingServiceList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ServiceDependencyDependencySupportingServiceList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewServiceDependencyDependencySupportingServiceListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

