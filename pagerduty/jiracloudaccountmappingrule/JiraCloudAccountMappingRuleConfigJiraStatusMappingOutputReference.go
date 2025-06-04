// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package jiracloudaccountmappingrule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-pagerduty-go/pagerduty/v14/jsii"

	"github.com/cdktf/cdktf-provider-pagerduty-go/pagerduty/v14/jiracloudaccountmappingrule/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type JiraCloudAccountMappingRuleConfigJiraStatusMappingOutputReference interface {
	cdktf.ComplexObject
	Acknowledged() JiraCloudAccountMappingRuleConfigJiraStatusMappingAcknowledgedOutputReference
	AcknowledgedInput() interface{}
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
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Resolved() JiraCloudAccountMappingRuleConfigJiraStatusMappingResolvedOutputReference
	ResolvedInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Triggered() JiraCloudAccountMappingRuleConfigJiraStatusMappingTriggeredOutputReference
	TriggeredInput() interface{}
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
	PutAcknowledged(value *JiraCloudAccountMappingRuleConfigJiraStatusMappingAcknowledged)
	PutResolved(value *JiraCloudAccountMappingRuleConfigJiraStatusMappingResolved)
	PutTriggered(value *JiraCloudAccountMappingRuleConfigJiraStatusMappingTriggered)
	ResetAcknowledged()
	ResetResolved()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for JiraCloudAccountMappingRuleConfigJiraStatusMappingOutputReference
type jsiiProxy_JiraCloudAccountMappingRuleConfigJiraStatusMappingOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraStatusMappingOutputReference) Acknowledged() JiraCloudAccountMappingRuleConfigJiraStatusMappingAcknowledgedOutputReference {
	var returns JiraCloudAccountMappingRuleConfigJiraStatusMappingAcknowledgedOutputReference
	_jsii_.Get(
		j,
		"acknowledged",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraStatusMappingOutputReference) AcknowledgedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"acknowledgedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraStatusMappingOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraStatusMappingOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraStatusMappingOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraStatusMappingOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraStatusMappingOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraStatusMappingOutputReference) Resolved() JiraCloudAccountMappingRuleConfigJiraStatusMappingResolvedOutputReference {
	var returns JiraCloudAccountMappingRuleConfigJiraStatusMappingResolvedOutputReference
	_jsii_.Get(
		j,
		"resolved",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraStatusMappingOutputReference) ResolvedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resolvedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraStatusMappingOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraStatusMappingOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraStatusMappingOutputReference) Triggered() JiraCloudAccountMappingRuleConfigJiraStatusMappingTriggeredOutputReference {
	var returns JiraCloudAccountMappingRuleConfigJiraStatusMappingTriggeredOutputReference
	_jsii_.Get(
		j,
		"triggered",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraStatusMappingOutputReference) TriggeredInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"triggeredInput",
		&returns,
	)
	return returns
}


func NewJiraCloudAccountMappingRuleConfigJiraStatusMappingOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) JiraCloudAccountMappingRuleConfigJiraStatusMappingOutputReference {
	_init_.Initialize()

	if err := validateNewJiraCloudAccountMappingRuleConfigJiraStatusMappingOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_JiraCloudAccountMappingRuleConfigJiraStatusMappingOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-pagerduty.jiraCloudAccountMappingRule.JiraCloudAccountMappingRuleConfigJiraStatusMappingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewJiraCloudAccountMappingRuleConfigJiraStatusMappingOutputReference_Override(j JiraCloudAccountMappingRuleConfigJiraStatusMappingOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-pagerduty.jiraCloudAccountMappingRule.JiraCloudAccountMappingRuleConfigJiraStatusMappingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		j,
	)
}

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraStatusMappingOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraStatusMappingOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraStatusMappingOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraStatusMappingOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraStatusMappingOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraStatusMappingOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraStatusMappingOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := j.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		j,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraStatusMappingOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := j.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		j,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraStatusMappingOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := j.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		j,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraStatusMappingOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := j.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		j,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraStatusMappingOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := j.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		j,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraStatusMappingOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := j.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		j,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraStatusMappingOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := j.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		j,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraStatusMappingOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := j.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		j,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraStatusMappingOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := j.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		j,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraStatusMappingOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		j,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraStatusMappingOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := j.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		j,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraStatusMappingOutputReference) PutAcknowledged(value *JiraCloudAccountMappingRuleConfigJiraStatusMappingAcknowledged) {
	if err := j.validatePutAcknowledgedParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putAcknowledged",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraStatusMappingOutputReference) PutResolved(value *JiraCloudAccountMappingRuleConfigJiraStatusMappingResolved) {
	if err := j.validatePutResolvedParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putResolved",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraStatusMappingOutputReference) PutTriggered(value *JiraCloudAccountMappingRuleConfigJiraStatusMappingTriggered) {
	if err := j.validatePutTriggeredParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putTriggered",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraStatusMappingOutputReference) ResetAcknowledged() {
	_jsii_.InvokeVoid(
		j,
		"resetAcknowledged",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraStatusMappingOutputReference) ResetResolved() {
	_jsii_.InvokeVoid(
		j,
		"resetResolved",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraStatusMappingOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := j.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		j,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraStatusMappingOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

