// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package enablement

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_Enablement) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (e *jsiiProxy_Enablement) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (e *jsiiProxy_Enablement) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (e *jsiiProxy_Enablement) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (e *jsiiProxy_Enablement) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (e *jsiiProxy_Enablement) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (e *jsiiProxy_Enablement) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (e *jsiiProxy_Enablement) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (e *jsiiProxy_Enablement) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (e *jsiiProxy_Enablement) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (e *jsiiProxy_Enablement) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (e *jsiiProxy_Enablement) validateImportFromParameters(id *string) error {
	return nil
}

func (e *jsiiProxy_Enablement) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (e *jsiiProxy_Enablement) validateMoveFromIdParameters(id *string) error {
	return nil
}

func (e *jsiiProxy_Enablement) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (e *jsiiProxy_Enablement) validateMoveToIdParameters(id *string) error {
	return nil
}

func (e *jsiiProxy_Enablement) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateEnablement_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateEnablement_IsConstructParameters(x interface{}) error {
	return nil
}

func validateEnablement_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateEnablement_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_Enablement) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Enablement) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Enablement) validateSetEnabledParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Enablement) validateSetEntityIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Enablement) validateSetEntityTypeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Enablement) validateSetFeatureParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Enablement) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_Enablement) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func validateNewEnablementParameters(scope constructs.Construct, id *string, config *EnablementConfig) error {
	return nil
}

