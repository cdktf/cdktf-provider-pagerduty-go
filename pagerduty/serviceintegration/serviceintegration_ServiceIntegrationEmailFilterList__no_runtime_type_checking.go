//go:build no_runtime_type_checking

package serviceintegration

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_ServiceIntegrationEmailFilterList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_ServiceIntegrationEmailFilterList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ServiceIntegrationEmailFilterList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ServiceIntegrationEmailFilterList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ServiceIntegrationEmailFilterList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ServiceIntegrationEmailFilterList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewServiceIntegrationEmailFilterListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

