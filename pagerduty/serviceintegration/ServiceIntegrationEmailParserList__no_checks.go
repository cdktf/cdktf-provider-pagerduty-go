// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package serviceintegration

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_ServiceIntegrationEmailParserList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_ServiceIntegrationEmailParserList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_ServiceIntegrationEmailParserList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ServiceIntegrationEmailParserList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ServiceIntegrationEmailParserList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ServiceIntegrationEmailParserList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ServiceIntegrationEmailParserList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewServiceIntegrationEmailParserListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

