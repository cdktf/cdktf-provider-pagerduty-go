//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package rulesetrule

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_RulesetRuleActionsRouteList) validateGetParameters(index *float64) error {
	return nil
}

func (r *jsiiProxy_RulesetRuleActionsRouteList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_RulesetRuleActionsRouteList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_RulesetRuleActionsRouteList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_RulesetRuleActionsRouteList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_RulesetRuleActionsRouteList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewRulesetRuleActionsRouteListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

