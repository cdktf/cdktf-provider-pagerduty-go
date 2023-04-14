package datapagerdutylicenses

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-pagerduty-go/pagerduty/v6/jsii"

	"github.com/cdktf/cdktf-provider-pagerduty-go/pagerduty/v6/datapagerdutylicenses/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataPagerdutyLicensesLicensesOutputReference interface {
	cdktf.ComplexObject
	AllocationsAvailable() *float64
	SetAllocationsAvailable(val *float64)
	AllocationsAvailableInput() *float64
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
	CurrentValue() *float64
	SetCurrentValue(val *float64)
	CurrentValueInput() *float64
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	// Experimental.
	Fqn() *string
	HtmlUrl() *string
	SetHtmlUrl(val *string)
	HtmlUrlInput() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Name() *string
	SetName(val *string)
	NameInput() *string
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
	ValidRoles() *[]*string
	SetValidRoles(val *[]*string)
	ValidRolesInput() *[]*string
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
	ResetAllocationsAvailable()
	ResetCurrentValue()
	ResetDescription()
	ResetHtmlUrl()
	ResetId()
	ResetName()
	ResetRoleGroup()
	ResetSelfAttribute()
	ResetSummary()
	ResetType()
	ResetValidRoles()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataPagerdutyLicensesLicensesOutputReference
type jsiiProxy_DataPagerdutyLicensesLicensesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataPagerdutyLicensesLicensesOutputReference) AllocationsAvailable() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"allocationsAvailable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPagerdutyLicensesLicensesOutputReference) AllocationsAvailableInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"allocationsAvailableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPagerdutyLicensesLicensesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPagerdutyLicensesLicensesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPagerdutyLicensesLicensesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPagerdutyLicensesLicensesOutputReference) CurrentValue() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"currentValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPagerdutyLicensesLicensesOutputReference) CurrentValueInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"currentValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPagerdutyLicensesLicensesOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPagerdutyLicensesLicensesOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPagerdutyLicensesLicensesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPagerdutyLicensesLicensesOutputReference) HtmlUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"htmlUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPagerdutyLicensesLicensesOutputReference) HtmlUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"htmlUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPagerdutyLicensesLicensesOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPagerdutyLicensesLicensesOutputReference) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPagerdutyLicensesLicensesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPagerdutyLicensesLicensesOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPagerdutyLicensesLicensesOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPagerdutyLicensesLicensesOutputReference) RoleGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPagerdutyLicensesLicensesOutputReference) RoleGroupInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPagerdutyLicensesLicensesOutputReference) SelfAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"selfAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPagerdutyLicensesLicensesOutputReference) SelfAttributeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"selfAttributeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPagerdutyLicensesLicensesOutputReference) Summary() *string {
	var returns *string
	_jsii_.Get(
		j,
		"summary",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPagerdutyLicensesLicensesOutputReference) SummaryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"summaryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPagerdutyLicensesLicensesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPagerdutyLicensesLicensesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPagerdutyLicensesLicensesOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPagerdutyLicensesLicensesOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPagerdutyLicensesLicensesOutputReference) ValidRoles() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"validRoles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPagerdutyLicensesLicensesOutputReference) ValidRolesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"validRolesInput",
		&returns,
	)
	return returns
}


func NewDataPagerdutyLicensesLicensesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataPagerdutyLicensesLicensesOutputReference {
	_init_.Initialize()

	if err := validateNewDataPagerdutyLicensesLicensesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataPagerdutyLicensesLicensesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-pagerduty.dataPagerdutyLicenses.DataPagerdutyLicensesLicensesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataPagerdutyLicensesLicensesOutputReference_Override(d DataPagerdutyLicensesLicensesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-pagerduty.dataPagerdutyLicenses.DataPagerdutyLicensesLicensesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataPagerdutyLicensesLicensesOutputReference)SetAllocationsAvailable(val *float64) {
	if err := j.validateSetAllocationsAvailableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allocationsAvailable",
		val,
	)
}

func (j *jsiiProxy_DataPagerdutyLicensesLicensesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataPagerdutyLicensesLicensesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataPagerdutyLicensesLicensesOutputReference)SetCurrentValue(val *float64) {
	if err := j.validateSetCurrentValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"currentValue",
		val,
	)
}

func (j *jsiiProxy_DataPagerdutyLicensesLicensesOutputReference)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_DataPagerdutyLicensesLicensesOutputReference)SetHtmlUrl(val *string) {
	if err := j.validateSetHtmlUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"htmlUrl",
		val,
	)
}

func (j *jsiiProxy_DataPagerdutyLicensesLicensesOutputReference)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataPagerdutyLicensesLicensesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataPagerdutyLicensesLicensesOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DataPagerdutyLicensesLicensesOutputReference)SetRoleGroup(val *string) {
	if err := j.validateSetRoleGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleGroup",
		val,
	)
}

func (j *jsiiProxy_DataPagerdutyLicensesLicensesOutputReference)SetSelfAttribute(val *string) {
	if err := j.validateSetSelfAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"selfAttribute",
		val,
	)
}

func (j *jsiiProxy_DataPagerdutyLicensesLicensesOutputReference)SetSummary(val *string) {
	if err := j.validateSetSummaryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"summary",
		val,
	)
}

func (j *jsiiProxy_DataPagerdutyLicensesLicensesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataPagerdutyLicensesLicensesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataPagerdutyLicensesLicensesOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_DataPagerdutyLicensesLicensesOutputReference)SetValidRoles(val *[]*string) {
	if err := j.validateSetValidRolesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"validRoles",
		val,
	)
}

func (d *jsiiProxy_DataPagerdutyLicensesLicensesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataPagerdutyLicensesLicensesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataPagerdutyLicensesLicensesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataPagerdutyLicensesLicensesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataPagerdutyLicensesLicensesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataPagerdutyLicensesLicensesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataPagerdutyLicensesLicensesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataPagerdutyLicensesLicensesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataPagerdutyLicensesLicensesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataPagerdutyLicensesLicensesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataPagerdutyLicensesLicensesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataPagerdutyLicensesLicensesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataPagerdutyLicensesLicensesOutputReference) ResetAllocationsAvailable() {
	_jsii_.InvokeVoid(
		d,
		"resetAllocationsAvailable",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataPagerdutyLicensesLicensesOutputReference) ResetCurrentValue() {
	_jsii_.InvokeVoid(
		d,
		"resetCurrentValue",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataPagerdutyLicensesLicensesOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		d,
		"resetDescription",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataPagerdutyLicensesLicensesOutputReference) ResetHtmlUrl() {
	_jsii_.InvokeVoid(
		d,
		"resetHtmlUrl",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataPagerdutyLicensesLicensesOutputReference) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataPagerdutyLicensesLicensesOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		d,
		"resetName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataPagerdutyLicensesLicensesOutputReference) ResetRoleGroup() {
	_jsii_.InvokeVoid(
		d,
		"resetRoleGroup",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataPagerdutyLicensesLicensesOutputReference) ResetSelfAttribute() {
	_jsii_.InvokeVoid(
		d,
		"resetSelfAttribute",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataPagerdutyLicensesLicensesOutputReference) ResetSummary() {
	_jsii_.InvokeVoid(
		d,
		"resetSummary",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataPagerdutyLicensesLicensesOutputReference) ResetType() {
	_jsii_.InvokeVoid(
		d,
		"resetType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataPagerdutyLicensesLicensesOutputReference) ResetValidRoles() {
	_jsii_.InvokeVoid(
		d,
		"resetValidRoles",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataPagerdutyLicensesLicensesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataPagerdutyLicensesLicensesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

