// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package escalationpolicy

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EscalationPolicyRuleTargetList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (e *jsiiProxy_EscalationPolicyRuleTargetList) validateGetParameters(index *float64) error {
	return nil
}

func (e *jsiiProxy_EscalationPolicyRuleTargetList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_EscalationPolicyRuleTargetList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_EscalationPolicyRuleTargetList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_EscalationPolicyRuleTargetList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_EscalationPolicyRuleTargetList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewEscalationPolicyRuleTargetListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

