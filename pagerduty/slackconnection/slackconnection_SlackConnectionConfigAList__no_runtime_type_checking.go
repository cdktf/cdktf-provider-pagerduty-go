//go:build no_runtime_type_checking

package slackconnection

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SlackConnectionConfigAList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_SlackConnectionConfigAList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_SlackConnectionConfigAList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SlackConnectionConfigAList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SlackConnectionConfigAList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_SlackConnectionConfigAList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewSlackConnectionConfigAListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

