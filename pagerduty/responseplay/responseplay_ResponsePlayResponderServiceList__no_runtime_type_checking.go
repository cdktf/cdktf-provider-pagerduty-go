//go:build no_runtime_type_checking

package responseplay

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_ResponsePlayResponderServiceList) validateGetParameters(index *float64) error {
	return nil
}

func (r *jsiiProxy_ResponsePlayResponderServiceList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ResponsePlayResponderServiceList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ResponsePlayResponderServiceList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ResponsePlayResponderServiceList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewResponsePlayResponderServiceListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

