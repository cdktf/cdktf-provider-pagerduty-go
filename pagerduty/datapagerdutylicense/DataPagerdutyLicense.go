// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datapagerdutylicense

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-pagerduty-go/pagerduty/v12/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-pagerduty-go/pagerduty/v12/datapagerdutylicense/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.3.1/docs/data-sources/license pagerduty_license}.
type DataPagerdutyLicense interface {
	cdktf.TerraformDataSource
	AllocationsAvailable() *float64
	SetAllocationsAvailable(val *float64)
	AllocationsAvailableInput() *float64
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CurrentValue() *float64
	SetCurrentValue(val *float64)
	CurrentValueInput() *float64
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	HtmlUrl() *string
	SetHtmlUrl(val *string)
	HtmlUrlInput() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
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
	RawOverrides() interface{}
	RoleGroup() *string
	SetRoleGroup(val *string)
	RoleGroupInput() *string
	SelfAttribute() *string
	SetSelfAttribute(val *string)
	SelfAttributeInput() *string
	Summary() *string
	SetSummary(val *string)
	SummaryInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Type() *string
	SetType(val *string)
	TypeInput() *string
	ValidRoles() *[]*string
	SetValidRoles(val *[]*string)
	ValidRolesInput() *[]*string
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	ResetAllocationsAvailable()
	ResetCurrentValue()
	ResetDescription()
	ResetHtmlUrl()
	ResetId()
	ResetName()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRoleGroup()
	ResetSelfAttribute()
	ResetSummary()
	ResetType()
	ResetValidRoles()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for DataPagerdutyLicense
type jsiiProxy_DataPagerdutyLicense struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataPagerdutyLicense) AllocationsAvailable() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"allocationsAvailable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPagerdutyLicense) AllocationsAvailableInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"allocationsAvailableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPagerdutyLicense) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPagerdutyLicense) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPagerdutyLicense) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPagerdutyLicense) CurrentValue() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"currentValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPagerdutyLicense) CurrentValueInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"currentValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPagerdutyLicense) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPagerdutyLicense) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPagerdutyLicense) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPagerdutyLicense) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPagerdutyLicense) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPagerdutyLicense) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPagerdutyLicense) HtmlUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"htmlUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPagerdutyLicense) HtmlUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"htmlUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPagerdutyLicense) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPagerdutyLicense) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPagerdutyLicense) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPagerdutyLicense) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPagerdutyLicense) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPagerdutyLicense) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPagerdutyLicense) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPagerdutyLicense) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPagerdutyLicense) RoleGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPagerdutyLicense) RoleGroupInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPagerdutyLicense) SelfAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"selfAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPagerdutyLicense) SelfAttributeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"selfAttributeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPagerdutyLicense) Summary() *string {
	var returns *string
	_jsii_.Get(
		j,
		"summary",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPagerdutyLicense) SummaryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"summaryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPagerdutyLicense) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPagerdutyLicense) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPagerdutyLicense) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPagerdutyLicense) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPagerdutyLicense) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPagerdutyLicense) ValidRoles() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"validRoles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPagerdutyLicense) ValidRolesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"validRolesInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.3.1/docs/data-sources/license pagerduty_license} Data Source.
func NewDataPagerdutyLicense(scope constructs.Construct, id *string, config *DataPagerdutyLicenseConfig) DataPagerdutyLicense {
	_init_.Initialize()

	if err := validateNewDataPagerdutyLicenseParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataPagerdutyLicense{}

	_jsii_.Create(
		"@cdktf/provider-pagerduty.dataPagerdutyLicense.DataPagerdutyLicense",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.3.1/docs/data-sources/license pagerduty_license} Data Source.
func NewDataPagerdutyLicense_Override(d DataPagerdutyLicense, scope constructs.Construct, id *string, config *DataPagerdutyLicenseConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-pagerduty.dataPagerdutyLicense.DataPagerdutyLicense",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataPagerdutyLicense)SetAllocationsAvailable(val *float64) {
	if err := j.validateSetAllocationsAvailableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allocationsAvailable",
		val,
	)
}

func (j *jsiiProxy_DataPagerdutyLicense)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataPagerdutyLicense)SetCurrentValue(val *float64) {
	if err := j.validateSetCurrentValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"currentValue",
		val,
	)
}

func (j *jsiiProxy_DataPagerdutyLicense)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataPagerdutyLicense)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_DataPagerdutyLicense)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataPagerdutyLicense)SetHtmlUrl(val *string) {
	if err := j.validateSetHtmlUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"htmlUrl",
		val,
	)
}

func (j *jsiiProxy_DataPagerdutyLicense)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataPagerdutyLicense)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataPagerdutyLicense)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DataPagerdutyLicense)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataPagerdutyLicense)SetRoleGroup(val *string) {
	if err := j.validateSetRoleGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleGroup",
		val,
	)
}

func (j *jsiiProxy_DataPagerdutyLicense)SetSelfAttribute(val *string) {
	if err := j.validateSetSelfAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"selfAttribute",
		val,
	)
}

func (j *jsiiProxy_DataPagerdutyLicense)SetSummary(val *string) {
	if err := j.validateSetSummaryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"summary",
		val,
	)
}

func (j *jsiiProxy_DataPagerdutyLicense)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_DataPagerdutyLicense)SetValidRoles(val *[]*string) {
	if err := j.validateSetValidRolesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"validRoles",
		val,
	)
}

// Generates CDKTF code for importing a DataPagerdutyLicense resource upon running "cdktf plan <stack-name>".
func DataPagerdutyLicense_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDataPagerdutyLicense_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-pagerduty.dataPagerdutyLicense.DataPagerdutyLicense",
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
func DataPagerdutyLicense_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataPagerdutyLicense_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-pagerduty.dataPagerdutyLicense.DataPagerdutyLicense",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataPagerdutyLicense_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataPagerdutyLicense_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-pagerduty.dataPagerdutyLicense.DataPagerdutyLicense",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataPagerdutyLicense_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataPagerdutyLicense_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-pagerduty.dataPagerdutyLicense.DataPagerdutyLicense",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataPagerdutyLicense_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-pagerduty.dataPagerdutyLicense.DataPagerdutyLicense",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataPagerdutyLicense) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataPagerdutyLicense) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataPagerdutyLicense) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataPagerdutyLicense) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataPagerdutyLicense) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataPagerdutyLicense) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataPagerdutyLicense) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataPagerdutyLicense) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataPagerdutyLicense) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataPagerdutyLicense) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataPagerdutyLicense) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataPagerdutyLicense) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataPagerdutyLicense) ResetAllocationsAvailable() {
	_jsii_.InvokeVoid(
		d,
		"resetAllocationsAvailable",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataPagerdutyLicense) ResetCurrentValue() {
	_jsii_.InvokeVoid(
		d,
		"resetCurrentValue",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataPagerdutyLicense) ResetDescription() {
	_jsii_.InvokeVoid(
		d,
		"resetDescription",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataPagerdutyLicense) ResetHtmlUrl() {
	_jsii_.InvokeVoid(
		d,
		"resetHtmlUrl",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataPagerdutyLicense) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataPagerdutyLicense) ResetName() {
	_jsii_.InvokeVoid(
		d,
		"resetName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataPagerdutyLicense) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataPagerdutyLicense) ResetRoleGroup() {
	_jsii_.InvokeVoid(
		d,
		"resetRoleGroup",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataPagerdutyLicense) ResetSelfAttribute() {
	_jsii_.InvokeVoid(
		d,
		"resetSelfAttribute",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataPagerdutyLicense) ResetSummary() {
	_jsii_.InvokeVoid(
		d,
		"resetSummary",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataPagerdutyLicense) ResetType() {
	_jsii_.InvokeVoid(
		d,
		"resetType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataPagerdutyLicense) ResetValidRoles() {
	_jsii_.InvokeVoid(
		d,
		"resetValidRoles",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataPagerdutyLicense) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataPagerdutyLicense) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataPagerdutyLicense) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataPagerdutyLicense) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

