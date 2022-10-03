//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package rulesetrule

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_RulesetRuleActionsExtractionsList) validateGetParameters(index *float64) error {
	return nil
}

func (r *jsiiProxy_RulesetRuleActionsExtractionsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_RulesetRuleActionsExtractionsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_RulesetRuleActionsExtractionsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_RulesetRuleActionsExtractionsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_RulesetRuleActionsExtractionsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewRulesetRuleActionsExtractionsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

