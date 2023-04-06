package automationactionsaction

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-pagerduty-go/pagerduty/v5/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-pagerduty-go/pagerduty/v5/automationactionsaction/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/pagerduty/r/automation_actions_action pagerduty_automation_actions_action}.
type AutomationActionsAction interface {
	cdktf.TerraformResource
	ActionClassification() *string
	SetActionClassification(val *string)
	ActionClassificationInput() *string
	ActionDataReference() AutomationActionsActionActionDataReferenceOutputReference
	ActionDataReferenceInput() *AutomationActionsActionActionDataReference
	ActionType() *string
	SetActionType(val *string)
	ActionTypeInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() *float64
	// Experimental.
	SetCount(val *float64)
	CreationTime() *string
	SetCreationTime(val *string)
	CreationTimeInput() *string
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
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	ModifyTime() *string
	SetModifyTime(val *string)
	ModifyTimeInput() *string
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
	RunnerId() *string
	SetRunnerId(val *string)
	RunnerIdInput() *string
	RunnerType() *string
	SetRunnerType(val *string)
	RunnerTypeInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Type() *string
	SetType(val *string)
	TypeInput() *string
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
	PutActionDataReference(value *AutomationActionsActionActionDataReference)
	ResetActionClassification()
	ResetCreationTime()
	ResetDescription()
	ResetId()
	ResetModifyTime()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRunnerId()
	ResetRunnerType()
	ResetType()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for AutomationActionsAction
type jsiiProxy_AutomationActionsAction struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_AutomationActionsAction) ActionClassification() *string {
	var returns *string
	_jsii_.Get(
		j,
		"actionClassification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationActionsAction) ActionClassificationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"actionClassificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationActionsAction) ActionDataReference() AutomationActionsActionActionDataReferenceOutputReference {
	var returns AutomationActionsActionActionDataReferenceOutputReference
	_jsii_.Get(
		j,
		"actionDataReference",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationActionsAction) ActionDataReferenceInput() *AutomationActionsActionActionDataReference {
	var returns *AutomationActionsActionActionDataReference
	_jsii_.Get(
		j,
		"actionDataReferenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationActionsAction) ActionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"actionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationActionsAction) ActionTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"actionTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationActionsAction) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationActionsAction) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationActionsAction) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationActionsAction) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationActionsAction) CreationTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"creationTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationActionsAction) CreationTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"creationTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationActionsAction) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationActionsAction) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationActionsAction) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationActionsAction) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationActionsAction) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationActionsAction) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationActionsAction) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationActionsAction) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationActionsAction) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationActionsAction) ModifyTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modifyTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationActionsAction) ModifyTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modifyTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationActionsAction) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationActionsAction) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationActionsAction) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationActionsAction) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationActionsAction) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationActionsAction) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationActionsAction) RunnerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runnerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationActionsAction) RunnerIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runnerIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationActionsAction) RunnerType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runnerType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationActionsAction) RunnerTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runnerTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationActionsAction) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationActionsAction) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationActionsAction) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationActionsAction) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationActionsAction) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/pagerduty/r/automation_actions_action pagerduty_automation_actions_action} Resource.
func NewAutomationActionsAction(scope constructs.Construct, id *string, config *AutomationActionsActionConfig) AutomationActionsAction {
	_init_.Initialize()

	if err := validateNewAutomationActionsActionParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AutomationActionsAction{}

	_jsii_.Create(
		"@cdktf/provider-pagerduty.automationActionsAction.AutomationActionsAction",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/pagerduty/r/automation_actions_action pagerduty_automation_actions_action} Resource.
func NewAutomationActionsAction_Override(a AutomationActionsAction, scope constructs.Construct, id *string, config *AutomationActionsActionConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-pagerduty.automationActionsAction.AutomationActionsAction",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AutomationActionsAction)SetActionClassification(val *string) {
	if err := j.validateSetActionClassificationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"actionClassification",
		val,
	)
}

func (j *jsiiProxy_AutomationActionsAction)SetActionType(val *string) {
	if err := j.validateSetActionTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"actionType",
		val,
	)
}

func (j *jsiiProxy_AutomationActionsAction)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AutomationActionsAction)SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AutomationActionsAction)SetCreationTime(val *string) {
	if err := j.validateSetCreationTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"creationTime",
		val,
	)
}

func (j *jsiiProxy_AutomationActionsAction)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AutomationActionsAction)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_AutomationActionsAction)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AutomationActionsAction)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AutomationActionsAction)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AutomationActionsAction)SetModifyTime(val *string) {
	if err := j.validateSetModifyTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"modifyTime",
		val,
	)
}

func (j *jsiiProxy_AutomationActionsAction)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_AutomationActionsAction)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AutomationActionsAction)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AutomationActionsAction)SetRunnerId(val *string) {
	if err := j.validateSetRunnerIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runnerId",
		val,
	)
}

func (j *jsiiProxy_AutomationActionsAction)SetRunnerType(val *string) {
	if err := j.validateSetRunnerTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runnerType",
		val,
	)
}

func (j *jsiiProxy_AutomationActionsAction)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
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
func AutomationActionsAction_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAutomationActionsAction_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-pagerduty.automationActionsAction.AutomationActionsAction",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AutomationActionsAction_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAutomationActionsAction_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-pagerduty.automationActionsAction.AutomationActionsAction",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AutomationActionsAction_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAutomationActionsAction_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-pagerduty.automationActionsAction.AutomationActionsAction",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AutomationActionsAction_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-pagerduty.automationActionsAction.AutomationActionsAction",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AutomationActionsAction) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AutomationActionsAction) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationActionsAction) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationActionsAction) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationActionsAction) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationActionsAction) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationActionsAction) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationActionsAction) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationActionsAction) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationActionsAction) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationActionsAction) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationActionsAction) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AutomationActionsAction) PutActionDataReference(value *AutomationActionsActionActionDataReference) {
	if err := a.validatePutActionDataReferenceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putActionDataReference",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AutomationActionsAction) ResetActionClassification() {
	_jsii_.InvokeVoid(
		a,
		"resetActionClassification",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutomationActionsAction) ResetCreationTime() {
	_jsii_.InvokeVoid(
		a,
		"resetCreationTime",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutomationActionsAction) ResetDescription() {
	_jsii_.InvokeVoid(
		a,
		"resetDescription",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutomationActionsAction) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutomationActionsAction) ResetModifyTime() {
	_jsii_.InvokeVoid(
		a,
		"resetModifyTime",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutomationActionsAction) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutomationActionsAction) ResetRunnerId() {
	_jsii_.InvokeVoid(
		a,
		"resetRunnerId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutomationActionsAction) ResetRunnerType() {
	_jsii_.InvokeVoid(
		a,
		"resetRunnerType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutomationActionsAction) ResetType() {
	_jsii_.InvokeVoid(
		a,
		"resetType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutomationActionsAction) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationActionsAction) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationActionsAction) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationActionsAction) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}
