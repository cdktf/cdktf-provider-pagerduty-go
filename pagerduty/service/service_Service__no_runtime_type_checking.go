//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package service

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_Service) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (s *jsiiProxy_Service) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Service) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Service) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Service) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Service) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Service) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Service) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Service) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Service) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Service) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Service) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (s *jsiiProxy_Service) validatePutAlertGroupingParametersParameters(value *ServiceAlertGroupingParameters) error {
	return nil
}

func (s *jsiiProxy_Service) validatePutAutoPauseNotificationsParametersParameters(value *ServiceAutoPauseNotificationsParameters) error {
	return nil
}

func (s *jsiiProxy_Service) validatePutIncidentUrgencyRuleParameters(value *ServiceIncidentUrgencyRule) error {
	return nil
}

func (s *jsiiProxy_Service) validatePutScheduledActionsParameters(value interface{}) error {
	return nil
}

func (s *jsiiProxy_Service) validatePutSupportHoursParameters(value *ServiceSupportHours) error {
	return nil
}

func validateService_IsConstructParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_Service) validateSetAcknowledgementTimeoutParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Service) validateSetAlertCreationParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Service) validateSetAlertGroupingParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Service) validateSetAlertGroupingTimeoutParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Service) validateSetAutoResolveTimeoutParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Service) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Service) validateSetDescriptionParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Service) validateSetEscalationPolicyParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Service) validateSetIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Service) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_Service) validateSetNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Service) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_Service) validateSetResponsePlayParameters(val *string) error {
	return nil
}

func validateNewServiceParameters(scope constructs.Construct, id *string, config *ServiceConfig) error {
	return nil
}

