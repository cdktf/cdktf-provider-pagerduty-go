// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package jiracloudaccountmappingrule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-pagerduty-go/pagerduty/v13/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-pagerduty-go/pagerduty/v13/jiracloudaccountmappingrule/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.17.2/docs/resources/jira_cloud_account_mapping_rule pagerduty_jira_cloud_account_mapping_rule}.
type JiraCloudAccountMappingRule interface {
	cdktf.TerraformResource
	AccountMapping() *string
	SetAccountMapping(val *string)
	AccountMappingInput() *string
	AutocreateJqlDisabledReason() *string
	AutocreateJqlDisabledUntil() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	Config() JiraCloudAccountMappingRuleConfigAOutputReference
	ConfigInput() interface{}
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
	// Experimental.
	AddOverride(path *string, value interface{})
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
	HasResourceMove() interface{}
	// Experimental.
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using it's instance function.
	// Experimental.
	MoveFromId(id *string)
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutConfig(value *JiraCloudAccountMappingRuleConfigA)
	ResetConfig()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for JiraCloudAccountMappingRule
type jsiiProxy_JiraCloudAccountMappingRule struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_JiraCloudAccountMappingRule) AccountMapping() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountMapping",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JiraCloudAccountMappingRule) AccountMappingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountMappingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JiraCloudAccountMappingRule) AutocreateJqlDisabledReason() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autocreateJqlDisabledReason",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JiraCloudAccountMappingRule) AutocreateJqlDisabledUntil() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autocreateJqlDisabledUntil",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JiraCloudAccountMappingRule) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JiraCloudAccountMappingRule) Config() JiraCloudAccountMappingRuleConfigAOutputReference {
	var returns JiraCloudAccountMappingRuleConfigAOutputReference
	_jsii_.Get(
		j,
		"config",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JiraCloudAccountMappingRule) ConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"configInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JiraCloudAccountMappingRule) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JiraCloudAccountMappingRule) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JiraCloudAccountMappingRule) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JiraCloudAccountMappingRule) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JiraCloudAccountMappingRule) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JiraCloudAccountMappingRule) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JiraCloudAccountMappingRule) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JiraCloudAccountMappingRule) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JiraCloudAccountMappingRule) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JiraCloudAccountMappingRule) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JiraCloudAccountMappingRule) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JiraCloudAccountMappingRule) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JiraCloudAccountMappingRule) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JiraCloudAccountMappingRule) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JiraCloudAccountMappingRule) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JiraCloudAccountMappingRule) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JiraCloudAccountMappingRule) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JiraCloudAccountMappingRule) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.17.2/docs/resources/jira_cloud_account_mapping_rule pagerduty_jira_cloud_account_mapping_rule} Resource.
func NewJiraCloudAccountMappingRule(scope constructs.Construct, id *string, config *JiraCloudAccountMappingRuleConfig) JiraCloudAccountMappingRule {
	_init_.Initialize()

	if err := validateNewJiraCloudAccountMappingRuleParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_JiraCloudAccountMappingRule{}

	_jsii_.Create(
		"@cdktf/provider-pagerduty.jiraCloudAccountMappingRule.JiraCloudAccountMappingRule",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.17.2/docs/resources/jira_cloud_account_mapping_rule pagerduty_jira_cloud_account_mapping_rule} Resource.
func NewJiraCloudAccountMappingRule_Override(j JiraCloudAccountMappingRule, scope constructs.Construct, id *string, config *JiraCloudAccountMappingRuleConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-pagerduty.jiraCloudAccountMappingRule.JiraCloudAccountMappingRule",
		[]interface{}{scope, id, config},
		j,
	)
}

func (j *jsiiProxy_JiraCloudAccountMappingRule)SetAccountMapping(val *string) {
	if err := j.validateSetAccountMappingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accountMapping",
		val,
	)
}

func (j *jsiiProxy_JiraCloudAccountMappingRule)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_JiraCloudAccountMappingRule)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_JiraCloudAccountMappingRule)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_JiraCloudAccountMappingRule)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_JiraCloudAccountMappingRule)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_JiraCloudAccountMappingRule)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_JiraCloudAccountMappingRule)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_JiraCloudAccountMappingRule)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

// Generates CDKTF code for importing a JiraCloudAccountMappingRule resource upon running "cdktf plan <stack-name>".
func JiraCloudAccountMappingRule_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateJiraCloudAccountMappingRule_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-pagerduty.jiraCloudAccountMappingRule.JiraCloudAccountMappingRule",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func JiraCloudAccountMappingRule_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateJiraCloudAccountMappingRule_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-pagerduty.jiraCloudAccountMappingRule.JiraCloudAccountMappingRule",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func JiraCloudAccountMappingRule_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateJiraCloudAccountMappingRule_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-pagerduty.jiraCloudAccountMappingRule.JiraCloudAccountMappingRule",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func JiraCloudAccountMappingRule_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateJiraCloudAccountMappingRule_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-pagerduty.jiraCloudAccountMappingRule.JiraCloudAccountMappingRule",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func JiraCloudAccountMappingRule_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-pagerduty.jiraCloudAccountMappingRule.JiraCloudAccountMappingRule",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JiraCloudAccountMappingRule) AddMoveTarget(moveTarget *string) {
	if err := j.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (j *jsiiProxy_JiraCloudAccountMappingRule) AddOverride(path *string, value interface{}) {
	if err := j.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (j *jsiiProxy_JiraCloudAccountMappingRule) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (j *jsiiProxy_JiraCloudAccountMappingRule) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (j *jsiiProxy_JiraCloudAccountMappingRule) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (j *jsiiProxy_JiraCloudAccountMappingRule) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (j *jsiiProxy_JiraCloudAccountMappingRule) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (j *jsiiProxy_JiraCloudAccountMappingRule) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (j *jsiiProxy_JiraCloudAccountMappingRule) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (j *jsiiProxy_JiraCloudAccountMappingRule) GetStringAttribute(terraformAttribute *string) *string {
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

func (j *jsiiProxy_JiraCloudAccountMappingRule) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (j *jsiiProxy_JiraCloudAccountMappingRule) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		j,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JiraCloudAccountMappingRule) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := j.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (j *jsiiProxy_JiraCloudAccountMappingRule) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (j *jsiiProxy_JiraCloudAccountMappingRule) MoveFromId(id *string) {
	if err := j.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"moveFromId",
		[]interface{}{id},
	)
}

func (j *jsiiProxy_JiraCloudAccountMappingRule) MoveTo(moveTarget *string, index interface{}) {
	if err := j.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (j *jsiiProxy_JiraCloudAccountMappingRule) MoveToId(id *string) {
	if err := j.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"moveToId",
		[]interface{}{id},
	)
}

func (j *jsiiProxy_JiraCloudAccountMappingRule) OverrideLogicalId(newLogicalId *string) {
	if err := j.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (j *jsiiProxy_JiraCloudAccountMappingRule) PutConfig(value *JiraCloudAccountMappingRuleConfigA) {
	if err := j.validatePutConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putConfig",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JiraCloudAccountMappingRule) ResetConfig() {
	_jsii_.InvokeVoid(
		j,
		"resetConfig",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JiraCloudAccountMappingRule) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		j,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JiraCloudAccountMappingRule) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		j,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JiraCloudAccountMappingRule) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		j,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JiraCloudAccountMappingRule) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		j,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JiraCloudAccountMappingRule) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		j,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JiraCloudAccountMappingRule) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JiraCloudAccountMappingRule) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		j,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

