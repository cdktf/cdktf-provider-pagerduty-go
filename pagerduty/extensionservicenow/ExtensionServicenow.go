// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package extensionservicenow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-pagerduty-go/pagerduty/v13/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-pagerduty-go/pagerduty/v13/extensionservicenow/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.25.1/docs/resources/extension_servicenow pagerduty_extension_servicenow}.
type ExtensionServicenow interface {
	cdktf.TerraformResource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
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
	EndpointUrl() *string
	SetEndpointUrl(val *string)
	EndpointUrlInput() *string
	ExtensionObjects() *[]*string
	SetExtensionObjects(val *[]*string)
	ExtensionObjectsInput() *[]*string
	ExtensionSchema() *string
	SetExtensionSchema(val *string)
	ExtensionSchemaInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	HtmlUrl() *string
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
	Referer() *string
	SetReferer(val *string)
	RefererInput() *string
	SnowPassword() *string
	SetSnowPassword(val *string)
	SnowPasswordInput() *string
	SnowUser() *string
	SetSnowUser(val *string)
	SnowUserInput() *string
	Summary() *string
	SetSummary(val *string)
	SummaryInput() *string
	SyncOptions() *string
	SetSyncOptions(val *string)
	SyncOptionsInput() *string
	Target() *string
	SetTarget(val *string)
	TargetInput() *string
	TaskType() *string
	SetTaskType(val *string)
	TaskTypeInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Type() *string
	SetType(val *string)
	TypeInput() *string
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
	ResetEndpointUrl()
	ResetName()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetSummary()
	ResetType()
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

// The jsii proxy struct for ExtensionServicenow
type jsiiProxy_ExtensionServicenow struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ExtensionServicenow) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExtensionServicenow) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExtensionServicenow) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExtensionServicenow) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExtensionServicenow) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExtensionServicenow) EndpointUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExtensionServicenow) EndpointUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExtensionServicenow) ExtensionObjects() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"extensionObjects",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExtensionServicenow) ExtensionObjectsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"extensionObjectsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExtensionServicenow) ExtensionSchema() *string {
	var returns *string
	_jsii_.Get(
		j,
		"extensionSchema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExtensionServicenow) ExtensionSchemaInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"extensionSchemaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExtensionServicenow) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExtensionServicenow) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExtensionServicenow) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExtensionServicenow) HtmlUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"htmlUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExtensionServicenow) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExtensionServicenow) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExtensionServicenow) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExtensionServicenow) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExtensionServicenow) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExtensionServicenow) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExtensionServicenow) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExtensionServicenow) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExtensionServicenow) Referer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"referer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExtensionServicenow) RefererInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"refererInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExtensionServicenow) SnowPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snowPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExtensionServicenow) SnowPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snowPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExtensionServicenow) SnowUser() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snowUser",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExtensionServicenow) SnowUserInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snowUserInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExtensionServicenow) Summary() *string {
	var returns *string
	_jsii_.Get(
		j,
		"summary",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExtensionServicenow) SummaryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"summaryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExtensionServicenow) SyncOptions() *string {
	var returns *string
	_jsii_.Get(
		j,
		"syncOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExtensionServicenow) SyncOptionsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"syncOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExtensionServicenow) Target() *string {
	var returns *string
	_jsii_.Get(
		j,
		"target",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExtensionServicenow) TargetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExtensionServicenow) TaskType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"taskType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExtensionServicenow) TaskTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"taskTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExtensionServicenow) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExtensionServicenow) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExtensionServicenow) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExtensionServicenow) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExtensionServicenow) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.25.1/docs/resources/extension_servicenow pagerduty_extension_servicenow} Resource.
func NewExtensionServicenow(scope constructs.Construct, id *string, config *ExtensionServicenowConfig) ExtensionServicenow {
	_init_.Initialize()

	if err := validateNewExtensionServicenowParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ExtensionServicenow{}

	_jsii_.Create(
		"@cdktf/provider-pagerduty.extensionServicenow.ExtensionServicenow",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.25.1/docs/resources/extension_servicenow pagerduty_extension_servicenow} Resource.
func NewExtensionServicenow_Override(e ExtensionServicenow, scope constructs.Construct, id *string, config *ExtensionServicenowConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-pagerduty.extensionServicenow.ExtensionServicenow",
		[]interface{}{scope, id, config},
		e,
	)
}

func (j *jsiiProxy_ExtensionServicenow)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ExtensionServicenow)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ExtensionServicenow)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ExtensionServicenow)SetEndpointUrl(val *string) {
	if err := j.validateSetEndpointUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endpointUrl",
		val,
	)
}

func (j *jsiiProxy_ExtensionServicenow)SetExtensionObjects(val *[]*string) {
	if err := j.validateSetExtensionObjectsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"extensionObjects",
		val,
	)
}

func (j *jsiiProxy_ExtensionServicenow)SetExtensionSchema(val *string) {
	if err := j.validateSetExtensionSchemaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"extensionSchema",
		val,
	)
}

func (j *jsiiProxy_ExtensionServicenow)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ExtensionServicenow)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ExtensionServicenow)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ExtensionServicenow)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ExtensionServicenow)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ExtensionServicenow)SetReferer(val *string) {
	if err := j.validateSetRefererParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"referer",
		val,
	)
}

func (j *jsiiProxy_ExtensionServicenow)SetSnowPassword(val *string) {
	if err := j.validateSetSnowPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"snowPassword",
		val,
	)
}

func (j *jsiiProxy_ExtensionServicenow)SetSnowUser(val *string) {
	if err := j.validateSetSnowUserParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"snowUser",
		val,
	)
}

func (j *jsiiProxy_ExtensionServicenow)SetSummary(val *string) {
	if err := j.validateSetSummaryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"summary",
		val,
	)
}

func (j *jsiiProxy_ExtensionServicenow)SetSyncOptions(val *string) {
	if err := j.validateSetSyncOptionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"syncOptions",
		val,
	)
}

func (j *jsiiProxy_ExtensionServicenow)SetTarget(val *string) {
	if err := j.validateSetTargetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"target",
		val,
	)
}

func (j *jsiiProxy_ExtensionServicenow)SetTaskType(val *string) {
	if err := j.validateSetTaskTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"taskType",
		val,
	)
}

func (j *jsiiProxy_ExtensionServicenow)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

// Generates CDKTF code for importing a ExtensionServicenow resource upon running "cdktf plan <stack-name>".
func ExtensionServicenow_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateExtensionServicenow_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-pagerduty.extensionServicenow.ExtensionServicenow",
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
func ExtensionServicenow_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateExtensionServicenow_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-pagerduty.extensionServicenow.ExtensionServicenow",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ExtensionServicenow_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateExtensionServicenow_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-pagerduty.extensionServicenow.ExtensionServicenow",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ExtensionServicenow_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateExtensionServicenow_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-pagerduty.extensionServicenow.ExtensionServicenow",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ExtensionServicenow_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-pagerduty.extensionServicenow.ExtensionServicenow",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (e *jsiiProxy_ExtensionServicenow) AddMoveTarget(moveTarget *string) {
	if err := e.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (e *jsiiProxy_ExtensionServicenow) AddOverride(path *string, value interface{}) {
	if err := e.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (e *jsiiProxy_ExtensionServicenow) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := e.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExtensionServicenow) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := e.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExtensionServicenow) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := e.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExtensionServicenow) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := e.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExtensionServicenow) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := e.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExtensionServicenow) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := e.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExtensionServicenow) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := e.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExtensionServicenow) GetStringAttribute(terraformAttribute *string) *string {
	if err := e.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExtensionServicenow) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := e.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExtensionServicenow) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExtensionServicenow) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := e.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (e *jsiiProxy_ExtensionServicenow) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := e.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExtensionServicenow) MoveFromId(id *string) {
	if err := e.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveFromId",
		[]interface{}{id},
	)
}

func (e *jsiiProxy_ExtensionServicenow) MoveTo(moveTarget *string, index interface{}) {
	if err := e.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (e *jsiiProxy_ExtensionServicenow) MoveToId(id *string) {
	if err := e.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveToId",
		[]interface{}{id},
	)
}

func (e *jsiiProxy_ExtensionServicenow) OverrideLogicalId(newLogicalId *string) {
	if err := e.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (e *jsiiProxy_ExtensionServicenow) ResetEndpointUrl() {
	_jsii_.InvokeVoid(
		e,
		"resetEndpointUrl",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ExtensionServicenow) ResetName() {
	_jsii_.InvokeVoid(
		e,
		"resetName",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ExtensionServicenow) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		e,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ExtensionServicenow) ResetSummary() {
	_jsii_.InvokeVoid(
		e,
		"resetSummary",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ExtensionServicenow) ResetType() {
	_jsii_.InvokeVoid(
		e,
		"resetType",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ExtensionServicenow) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExtensionServicenow) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExtensionServicenow) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExtensionServicenow) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExtensionServicenow) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExtensionServicenow) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

