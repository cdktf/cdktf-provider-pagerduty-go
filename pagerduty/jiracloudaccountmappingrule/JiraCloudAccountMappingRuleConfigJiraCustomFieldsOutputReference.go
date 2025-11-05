// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package jiracloudaccountmappingrule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-pagerduty-go/pagerduty/v14/jsii"

	"github.com/cdktf/cdktf-provider-pagerduty-go/pagerduty/v14/jiracloudaccountmappingrule/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type JiraCloudAccountMappingRuleConfigJiraCustomFieldsOutputReference interface {
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
	SourceIncidentField() *string
	SetSourceIncidentField(val *string)
	SourceIncidentFieldInput() *string
	TargetIssueField() *string
	SetTargetIssueField(val *string)
	TargetIssueFieldInput() *string
	TargetIssueFieldName() *string
	SetTargetIssueFieldName(val *string)
	TargetIssueFieldNameInput() *string
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
	Value() *string
	SetValue(val *string)
	ValueInput() *string
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	ResetSourceIncidentField()
	ResetValue()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for JiraCloudAccountMappingRuleConfigJiraCustomFieldsOutputReference
type jsiiProxy_JiraCloudAccountMappingRuleConfigJiraCustomFieldsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraCustomFieldsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraCustomFieldsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraCustomFieldsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraCustomFieldsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraCustomFieldsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraCustomFieldsOutputReference) SourceIncidentField() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceIncidentField",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraCustomFieldsOutputReference) SourceIncidentFieldInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceIncidentFieldInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraCustomFieldsOutputReference) TargetIssueField() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetIssueField",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraCustomFieldsOutputReference) TargetIssueFieldInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetIssueFieldInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraCustomFieldsOutputReference) TargetIssueFieldName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetIssueFieldName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraCustomFieldsOutputReference) TargetIssueFieldNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetIssueFieldNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraCustomFieldsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraCustomFieldsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraCustomFieldsOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraCustomFieldsOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraCustomFieldsOutputReference) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraCustomFieldsOutputReference) ValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"valueInput",
		&returns,
	)
	return returns
}


func NewJiraCloudAccountMappingRuleConfigJiraCustomFieldsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) JiraCloudAccountMappingRuleConfigJiraCustomFieldsOutputReference {
	_init_.Initialize()

	if err := validateNewJiraCloudAccountMappingRuleConfigJiraCustomFieldsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_JiraCloudAccountMappingRuleConfigJiraCustomFieldsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-pagerduty.jiraCloudAccountMappingRule.JiraCloudAccountMappingRuleConfigJiraCustomFieldsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewJiraCloudAccountMappingRuleConfigJiraCustomFieldsOutputReference_Override(j JiraCloudAccountMappingRuleConfigJiraCustomFieldsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-pagerduty.jiraCloudAccountMappingRule.JiraCloudAccountMappingRuleConfigJiraCustomFieldsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		j,
	)
}

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraCustomFieldsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraCustomFieldsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraCustomFieldsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraCustomFieldsOutputReference)SetSourceIncidentField(val *string) {
	if err := j.validateSetSourceIncidentFieldParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceIncidentField",
		val,
	)
}

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraCustomFieldsOutputReference)SetTargetIssueField(val *string) {
	if err := j.validateSetTargetIssueFieldParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetIssueField",
		val,
	)
}

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraCustomFieldsOutputReference)SetTargetIssueFieldName(val *string) {
	if err := j.validateSetTargetIssueFieldNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetIssueFieldName",
		val,
	)
}

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraCustomFieldsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraCustomFieldsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraCustomFieldsOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraCustomFieldsOutputReference)SetValue(val *string) {
	if err := j.validateSetValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"value",
		val,
	)
}

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraCustomFieldsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraCustomFieldsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraCustomFieldsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraCustomFieldsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraCustomFieldsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraCustomFieldsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraCustomFieldsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraCustomFieldsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraCustomFieldsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraCustomFieldsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraCustomFieldsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		j,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraCustomFieldsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := j.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		j,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraCustomFieldsOutputReference) ResetSourceIncidentField() {
	_jsii_.InvokeVoid(
		j,
		"resetSourceIncidentField",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraCustomFieldsOutputReference) ResetValue() {
	_jsii_.InvokeVoid(
		j,
		"resetValue",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraCustomFieldsOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := j.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		j,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraCustomFieldsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

