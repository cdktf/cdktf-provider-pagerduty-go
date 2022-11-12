//go:build no_runtime_type_checking

package rulesetrule

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_RulesetRuleVariableParametersList) validateGetParameters(index *float64) error {
	return nil
}

func (r *jsiiProxy_RulesetRuleVariableParametersList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_RulesetRuleVariableParametersList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_RulesetRuleVariableParametersList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_RulesetRuleVariableParametersList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_RulesetRuleVariableParametersList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewRulesetRuleVariableParametersListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

