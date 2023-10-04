// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package service

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-pagerduty-go/pagerduty/v11/jsii"

	"github.com/cdktf/cdktf-provider-pagerduty-go/pagerduty/v11/service/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ServiceIncidentUrgencyRuleOutputReference interface {
	cdktf.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DuringSupportHours() ServiceIncidentUrgencyRuleDuringSupportHoursOutputReference
	DuringSupportHoursInput() *ServiceIncidentUrgencyRuleDuringSupportHours
	// Experimental.
	Fqn() *string
	InternalValue() *ServiceIncidentUrgencyRule
	SetInternalValue(val *ServiceIncidentUrgencyRule)
	OutsideSupportHours() ServiceIncidentUrgencyRuleOutsideSupportHoursOutputReference
	OutsideSupportHoursInput() *ServiceIncidentUrgencyRuleOutsideSupportHours
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Type() *string
	SetType(val *string)
	TypeInput() *string
	Urgency() *string
	SetUrgency(val *string)
	UrgencyInput() *string
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	PutDuringSupportHours(value *ServiceIncidentUrgencyRuleDuringSupportHours)
	PutOutsideSupportHours(value *ServiceIncidentUrgencyRuleOutsideSupportHours)
	ResetDuringSupportHours()
	ResetOutsideSupportHours()
	ResetUrgency()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ServiceIncidentUrgencyRuleOutputReference
type jsiiProxy_ServiceIncidentUrgencyRuleOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ServiceIncidentUrgencyRuleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceIncidentUrgencyRuleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceIncidentUrgencyRuleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceIncidentUrgencyRuleOutputReference) DuringSupportHours() ServiceIncidentUrgencyRuleDuringSupportHoursOutputReference {
	var returns ServiceIncidentUrgencyRuleDuringSupportHoursOutputReference
	_jsii_.Get(
		j,
		"duringSupportHours",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceIncidentUrgencyRuleOutputReference) DuringSupportHoursInput() *ServiceIncidentUrgencyRuleDuringSupportHours {
	var returns *ServiceIncidentUrgencyRuleDuringSupportHours
	_jsii_.Get(
		j,
		"duringSupportHoursInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceIncidentUrgencyRuleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceIncidentUrgencyRuleOutputReference) InternalValue() *ServiceIncidentUrgencyRule {
	var returns *ServiceIncidentUrgencyRule
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceIncidentUrgencyRuleOutputReference) OutsideSupportHours() ServiceIncidentUrgencyRuleOutsideSupportHoursOutputReference {
	var returns ServiceIncidentUrgencyRuleOutsideSupportHoursOutputReference
	_jsii_.Get(
		j,
		"outsideSupportHours",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceIncidentUrgencyRuleOutputReference) OutsideSupportHoursInput() *ServiceIncidentUrgencyRuleOutsideSupportHours {
	var returns *ServiceIncidentUrgencyRuleOutsideSupportHours
	_jsii_.Get(
		j,
		"outsideSupportHoursInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceIncidentUrgencyRuleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceIncidentUrgencyRuleOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceIncidentUrgencyRuleOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceIncidentUrgencyRuleOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceIncidentUrgencyRuleOutputReference) Urgency() *string {
	var returns *string
	_jsii_.Get(
		j,
		"urgency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceIncidentUrgencyRuleOutputReference) UrgencyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"urgencyInput",
		&returns,
	)
	return returns
}


func NewServiceIncidentUrgencyRuleOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ServiceIncidentUrgencyRuleOutputReference {
	_init_.Initialize()

	if err := validateNewServiceIncidentUrgencyRuleOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ServiceIncidentUrgencyRuleOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-pagerduty.service.ServiceIncidentUrgencyRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewServiceIncidentUrgencyRuleOutputReference_Override(s ServiceIncidentUrgencyRuleOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-pagerduty.service.ServiceIncidentUrgencyRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_ServiceIncidentUrgencyRuleOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ServiceIncidentUrgencyRuleOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ServiceIncidentUrgencyRuleOutputReference)SetInternalValue(val *ServiceIncidentUrgencyRule) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ServiceIncidentUrgencyRuleOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ServiceIncidentUrgencyRuleOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ServiceIncidentUrgencyRuleOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_ServiceIncidentUrgencyRuleOutputReference)SetUrgency(val *string) {
	if err := j.validateSetUrgencyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"urgency",
		val,
	)
}

func (s *jsiiProxy_ServiceIncidentUrgencyRuleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceIncidentUrgencyRuleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceIncidentUrgencyRuleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceIncidentUrgencyRuleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceIncidentUrgencyRuleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceIncidentUrgencyRuleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceIncidentUrgencyRuleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceIncidentUrgencyRuleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceIncidentUrgencyRuleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceIncidentUrgencyRuleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceIncidentUrgencyRuleOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceIncidentUrgencyRuleOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceIncidentUrgencyRuleOutputReference) PutDuringSupportHours(value *ServiceIncidentUrgencyRuleDuringSupportHours) {
	if err := s.validatePutDuringSupportHoursParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putDuringSupportHours",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_ServiceIncidentUrgencyRuleOutputReference) PutOutsideSupportHours(value *ServiceIncidentUrgencyRuleOutsideSupportHours) {
	if err := s.validatePutOutsideSupportHoursParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putOutsideSupportHours",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_ServiceIncidentUrgencyRuleOutputReference) ResetDuringSupportHours() {
	_jsii_.InvokeVoid(
		s,
		"resetDuringSupportHours",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceIncidentUrgencyRuleOutputReference) ResetOutsideSupportHours() {
	_jsii_.InvokeVoid(
		s,
		"resetOutsideSupportHours",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceIncidentUrgencyRuleOutputReference) ResetUrgency() {
	_jsii_.InvokeVoid(
		s,
		"resetUrgency",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceIncidentUrgencyRuleOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := s.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceIncidentUrgencyRuleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

