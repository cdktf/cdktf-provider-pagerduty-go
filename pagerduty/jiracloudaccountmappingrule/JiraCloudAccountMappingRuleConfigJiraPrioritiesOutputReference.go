// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package jiracloudaccountmappingrule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-pagerduty-go/pagerduty/v14/jsii"

	"github.com/cdktf/cdktf-provider-pagerduty-go/pagerduty/v14/jiracloudaccountmappingrule/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type JiraCloudAccountMappingRuleConfigJiraPrioritiesOutputReference interface {
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
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	JiraId() *string
	SetJiraId(val *string)
	JiraIdInput() *string
	PagerdutyId() *string
	SetPagerdutyId(val *string)
	PagerdutyIdInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for JiraCloudAccountMappingRuleConfigJiraPrioritiesOutputReference
type jsiiProxy_JiraCloudAccountMappingRuleConfigJiraPrioritiesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraPrioritiesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraPrioritiesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraPrioritiesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraPrioritiesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraPrioritiesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraPrioritiesOutputReference) JiraId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jiraId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraPrioritiesOutputReference) JiraIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jiraIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraPrioritiesOutputReference) PagerdutyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pagerdutyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraPrioritiesOutputReference) PagerdutyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pagerdutyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraPrioritiesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraPrioritiesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewJiraCloudAccountMappingRuleConfigJiraPrioritiesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) JiraCloudAccountMappingRuleConfigJiraPrioritiesOutputReference {
	_init_.Initialize()

	if err := validateNewJiraCloudAccountMappingRuleConfigJiraPrioritiesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_JiraCloudAccountMappingRuleConfigJiraPrioritiesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-pagerduty.jiraCloudAccountMappingRule.JiraCloudAccountMappingRuleConfigJiraPrioritiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewJiraCloudAccountMappingRuleConfigJiraPrioritiesOutputReference_Override(j JiraCloudAccountMappingRuleConfigJiraPrioritiesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-pagerduty.jiraCloudAccountMappingRule.JiraCloudAccountMappingRuleConfigJiraPrioritiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		j,
	)
}

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraPrioritiesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraPrioritiesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraPrioritiesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraPrioritiesOutputReference)SetJiraId(val *string) {
	if err := j.validateSetJiraIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jiraId",
		val,
	)
}

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraPrioritiesOutputReference)SetPagerdutyId(val *string) {
	if err := j.validateSetPagerdutyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pagerdutyId",
		val,
	)
}

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraPrioritiesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraPrioritiesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraPrioritiesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraPrioritiesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraPrioritiesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraPrioritiesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraPrioritiesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraPrioritiesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraPrioritiesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraPrioritiesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraPrioritiesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraPrioritiesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraPrioritiesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		j,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraPrioritiesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraPrioritiesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraPrioritiesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

