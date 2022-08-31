//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Prebuilt pagerduty Provider for Terraform CDK (cdktf)
package pagerduty

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_ServiceEventRuleActionsEventActionList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_ServiceEventRuleActionsEventActionList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ServiceEventRuleActionsEventActionList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ServiceEventRuleActionsEventActionList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ServiceEventRuleActionsEventActionList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ServiceEventRuleActionsEventActionList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewServiceEventRuleActionsEventActionListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

