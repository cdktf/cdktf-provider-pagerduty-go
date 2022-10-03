//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package eventorchestrationrouter

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EventOrchestrationRouterSetRuleList) validateGetParameters(index *float64) error {
	return nil
}

func (e *jsiiProxy_EventOrchestrationRouterSetRuleList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_EventOrchestrationRouterSetRuleList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_EventOrchestrationRouterSetRuleList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_EventOrchestrationRouterSetRuleList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_EventOrchestrationRouterSetRuleList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewEventOrchestrationRouterSetRuleListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

