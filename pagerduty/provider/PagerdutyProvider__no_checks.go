// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package provider

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_PagerdutyProvider) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (p *jsiiProxy_PagerdutyProvider) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validatePagerdutyProvider_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validatePagerdutyProvider_IsConstructParameters(x interface{}) error {
	return nil
}

func validatePagerdutyProvider_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validatePagerdutyProvider_IsTerraformProviderParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_PagerdutyProvider) validateSetInsecureTlsParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_PagerdutyProvider) validateSetSkipCredentialsValidationParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_PagerdutyProvider) validateSetUseAppOauthScopedTokenParameters(val *PagerdutyProviderUseAppOauthScopedToken) error {
	return nil
}

func validateNewPagerdutyProviderParameters(scope constructs.Construct, id *string, config *PagerdutyProviderConfig) error {
	return nil
}

