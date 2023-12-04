// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package serviceintegration

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_ServiceIntegration) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (s *jsiiProxy_ServiceIntegration) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (s *jsiiProxy_ServiceIntegration) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_ServiceIntegration) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_ServiceIntegration) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_ServiceIntegration) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_ServiceIntegration) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_ServiceIntegration) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_ServiceIntegration) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_ServiceIntegration) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_ServiceIntegration) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_ServiceIntegration) validateImportFromParameters(id *string) error {
	return nil
}

func (s *jsiiProxy_ServiceIntegration) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_ServiceIntegration) validateMoveFromIdParameters(id *string) error {
	return nil
}

func (s *jsiiProxy_ServiceIntegration) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (s *jsiiProxy_ServiceIntegration) validateMoveToIdParameters(id *string) error {
	return nil
}

func (s *jsiiProxy_ServiceIntegration) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (s *jsiiProxy_ServiceIntegration) validatePutEmailFilterParameters(value interface{}) error {
	return nil
}

func (s *jsiiProxy_ServiceIntegration) validatePutEmailParserParameters(value interface{}) error {
	return nil
}

func validateServiceIntegration_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateServiceIntegration_IsConstructParameters(x interface{}) error {
	return nil
}

func validateServiceIntegration_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateServiceIntegration_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_ServiceIntegration) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ServiceIntegration) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ServiceIntegration) validateSetEmailFilterModeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ServiceIntegration) validateSetEmailIncidentCreationParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ServiceIntegration) validateSetEmailParsingFallbackParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ServiceIntegration) validateSetIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ServiceIntegration) validateSetIntegrationEmailParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ServiceIntegration) validateSetIntegrationKeyParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ServiceIntegration) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_ServiceIntegration) validateSetNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ServiceIntegration) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_ServiceIntegration) validateSetServiceParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ServiceIntegration) validateSetTypeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ServiceIntegration) validateSetVendorParameters(val *string) error {
	return nil
}

func validateNewServiceIntegrationParameters(scope constructs.Construct, id *string, config *ServiceIntegrationConfig) error {
	return nil
}

