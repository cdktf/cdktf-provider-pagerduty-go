// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package addon

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_Addon) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (a *jsiiProxy_Addon) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (a *jsiiProxy_Addon) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (a *jsiiProxy_Addon) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (a *jsiiProxy_Addon) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (a *jsiiProxy_Addon) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (a *jsiiProxy_Addon) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (a *jsiiProxy_Addon) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (a *jsiiProxy_Addon) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (a *jsiiProxy_Addon) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (a *jsiiProxy_Addon) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (a *jsiiProxy_Addon) validateImportFromParameters(id *string) error {
	return nil
}

func (a *jsiiProxy_Addon) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (a *jsiiProxy_Addon) validateMoveFromIdParameters(id *string) error {
	return nil
}

func (a *jsiiProxy_Addon) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (a *jsiiProxy_Addon) validateMoveToIdParameters(id *string) error {
	return nil
}

func (a *jsiiProxy_Addon) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateAddon_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateAddon_IsConstructParameters(x interface{}) error {
	return nil
}

func validateAddon_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateAddon_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_Addon) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Addon) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Addon) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_Addon) validateSetNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Addon) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_Addon) validateSetSrcParameters(val *string) error {
	return nil
}

func validateNewAddonParameters(scope constructs.Construct, id *string, config *AddonConfig) error {
	return nil
}

