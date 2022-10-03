//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package rulesetrule

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_RulesetRuleActionsEventActionList) validateGetParameters(index *float64) error {
	return nil
}

func (r *jsiiProxy_RulesetRuleActionsEventActionList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_RulesetRuleActionsEventActionList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_RulesetRuleActionsEventActionList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_RulesetRuleActionsEventActionList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_RulesetRuleActionsEventActionList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewRulesetRuleActionsEventActionListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

