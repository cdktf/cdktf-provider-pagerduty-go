//go:build no_runtime_type_checking

package serviceintegration

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_ServiceIntegrationEmailParserValueExtractorList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_ServiceIntegrationEmailParserValueExtractorList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ServiceIntegrationEmailParserValueExtractorList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ServiceIntegrationEmailParserValueExtractorList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ServiceIntegrationEmailParserValueExtractorList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ServiceIntegrationEmailParserValueExtractorList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewServiceIntegrationEmailParserValueExtractorListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

