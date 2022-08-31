//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Prebuilt pagerduty Provider for Terraform CDK (cdktf)
package pagerduty

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EscalationPolicyRuleList) validateGetParameters(index *float64) error {
	return nil
}

func (e *jsiiProxy_EscalationPolicyRuleList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_EscalationPolicyRuleList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_EscalationPolicyRuleList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_EscalationPolicyRuleList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_EscalationPolicyRuleList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewEscalationPolicyRuleListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

