// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package jiracloudaccountmappingrule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-pagerduty-go/pagerduty/v13/jsii"

	"github.com/cdktf/cdktf-provider-pagerduty-go/pagerduty/v13/jiracloudaccountmappingrule/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type JiraCloudAccountMappingRuleConfigJiraOutputReference interface {
	cdktf.ComplexObject
	AutocreateJql() *string
	SetAutocreateJql(val *string)
	AutocreateJqlInput() *string
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
	CreateIssueOnIncidentTrigger() interface{}
	SetCreateIssueOnIncidentTrigger(val interface{})
	CreateIssueOnIncidentTriggerInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	CustomFields() JiraCloudAccountMappingRuleConfigJiraCustomFieldsList
	CustomFieldsInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	IssueType() JiraCloudAccountMappingRuleConfigJiraIssueTypeOutputReference
	IssueTypeInput() interface{}
	Priorities() JiraCloudAccountMappingRuleConfigJiraPrioritiesList
	PrioritiesInput() interface{}
	Project() JiraCloudAccountMappingRuleConfigJiraProjectOutputReference
	ProjectInput() interface{}
	StatusMapping() JiraCloudAccountMappingRuleConfigJiraStatusMappingOutputReference
	StatusMappingInput() interface{}
	SyncNotesUser() *string
	SetSyncNotesUser(val *string)
	SyncNotesUserInput() *string
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
	PutCustomFields(value interface{})
	PutIssueType(value *JiraCloudAccountMappingRuleConfigJiraIssueType)
	PutPriorities(value interface{})
	PutProject(value *JiraCloudAccountMappingRuleConfigJiraProject)
	PutStatusMapping(value *JiraCloudAccountMappingRuleConfigJiraStatusMapping)
	ResetAutocreateJql()
	ResetCreateIssueOnIncidentTrigger()
	ResetCustomFields()
	ResetPriorities()
	ResetStatusMapping()
	ResetSyncNotesUser()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for JiraCloudAccountMappingRuleConfigJiraOutputReference
type jsiiProxy_JiraCloudAccountMappingRuleConfigJiraOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraOutputReference) AutocreateJql() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autocreateJql",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraOutputReference) AutocreateJqlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autocreateJqlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraOutputReference) CreateIssueOnIncidentTrigger() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"createIssueOnIncidentTrigger",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraOutputReference) CreateIssueOnIncidentTriggerInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"createIssueOnIncidentTriggerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraOutputReference) CustomFields() JiraCloudAccountMappingRuleConfigJiraCustomFieldsList {
	var returns JiraCloudAccountMappingRuleConfigJiraCustomFieldsList
	_jsii_.Get(
		j,
		"customFields",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraOutputReference) CustomFieldsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customFieldsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraOutputReference) IssueType() JiraCloudAccountMappingRuleConfigJiraIssueTypeOutputReference {
	var returns JiraCloudAccountMappingRuleConfigJiraIssueTypeOutputReference
	_jsii_.Get(
		j,
		"issueType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraOutputReference) IssueTypeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"issueTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraOutputReference) Priorities() JiraCloudAccountMappingRuleConfigJiraPrioritiesList {
	var returns JiraCloudAccountMappingRuleConfigJiraPrioritiesList
	_jsii_.Get(
		j,
		"priorities",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraOutputReference) PrioritiesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"prioritiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraOutputReference) Project() JiraCloudAccountMappingRuleConfigJiraProjectOutputReference {
	var returns JiraCloudAccountMappingRuleConfigJiraProjectOutputReference
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraOutputReference) ProjectInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraOutputReference) StatusMapping() JiraCloudAccountMappingRuleConfigJiraStatusMappingOutputReference {
	var returns JiraCloudAccountMappingRuleConfigJiraStatusMappingOutputReference
	_jsii_.Get(
		j,
		"statusMapping",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraOutputReference) StatusMappingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"statusMappingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraOutputReference) SyncNotesUser() *string {
	var returns *string
	_jsii_.Get(
		j,
		"syncNotesUser",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraOutputReference) SyncNotesUserInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"syncNotesUserInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewJiraCloudAccountMappingRuleConfigJiraOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) JiraCloudAccountMappingRuleConfigJiraOutputReference {
	_init_.Initialize()

	if err := validateNewJiraCloudAccountMappingRuleConfigJiraOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_JiraCloudAccountMappingRuleConfigJiraOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-pagerduty.jiraCloudAccountMappingRule.JiraCloudAccountMappingRuleConfigJiraOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewJiraCloudAccountMappingRuleConfigJiraOutputReference_Override(j JiraCloudAccountMappingRuleConfigJiraOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-pagerduty.jiraCloudAccountMappingRule.JiraCloudAccountMappingRuleConfigJiraOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		j,
	)
}

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraOutputReference)SetAutocreateJql(val *string) {
	if err := j.validateSetAutocreateJqlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autocreateJql",
		val,
	)
}

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraOutputReference)SetCreateIssueOnIncidentTrigger(val interface{}) {
	if err := j.validateSetCreateIssueOnIncidentTriggerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createIssueOnIncidentTrigger",
		val,
	)
}

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraOutputReference)SetSyncNotesUser(val *string) {
	if err := j.validateSetSyncNotesUserParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"syncNotesUser",
		val,
	)
}

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		j,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraOutputReference) PutCustomFields(value interface{}) {
	if err := j.validatePutCustomFieldsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putCustomFields",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraOutputReference) PutIssueType(value *JiraCloudAccountMappingRuleConfigJiraIssueType) {
	if err := j.validatePutIssueTypeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putIssueType",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraOutputReference) PutPriorities(value interface{}) {
	if err := j.validatePutPrioritiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putPriorities",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraOutputReference) PutProject(value *JiraCloudAccountMappingRuleConfigJiraProject) {
	if err := j.validatePutProjectParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putProject",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraOutputReference) PutStatusMapping(value *JiraCloudAccountMappingRuleConfigJiraStatusMapping) {
	if err := j.validatePutStatusMappingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putStatusMapping",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraOutputReference) ResetAutocreateJql() {
	_jsii_.InvokeVoid(
		j,
		"resetAutocreateJql",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraOutputReference) ResetCreateIssueOnIncidentTrigger() {
	_jsii_.InvokeVoid(
		j,
		"resetCreateIssueOnIncidentTrigger",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraOutputReference) ResetCustomFields() {
	_jsii_.InvokeVoid(
		j,
		"resetCustomFields",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraOutputReference) ResetPriorities() {
	_jsii_.InvokeVoid(
		j,
		"resetPriorities",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraOutputReference) ResetStatusMapping() {
	_jsii_.InvokeVoid(
		j,
		"resetStatusMapping",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraOutputReference) ResetSyncNotesUser() {
	_jsii_.InvokeVoid(
		j,
		"resetSyncNotesUser",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (j *jsiiProxy_JiraCloudAccountMappingRuleConfigJiraOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

