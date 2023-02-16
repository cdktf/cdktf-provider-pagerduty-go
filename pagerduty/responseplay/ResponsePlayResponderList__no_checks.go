//go:build no_runtime_type_checking

package responseplay

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_ResponsePlayResponderList) validateGetParameters(index *float64) error {
	return nil
}

func (r *jsiiProxy_ResponsePlayResponderList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ResponsePlayResponderList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ResponsePlayResponderList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ResponsePlayResponderList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ResponsePlayResponderList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewResponsePlayResponderListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

