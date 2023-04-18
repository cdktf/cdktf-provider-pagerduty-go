package automationactionsaction

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-pagerduty-go/pagerduty/v7/jsii"

	"github.com/cdktf/cdktf-provider-pagerduty-go/pagerduty/v7/automationactionsaction/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AutomationActionsActionActionDataReferenceOutputReference interface {
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
	InternalValue() *AutomationActionsActionActionDataReference
	SetInternalValue(val *AutomationActionsActionActionDataReference)
	InvocationCommand() *string
	SetInvocationCommand(val *string)
	InvocationCommandInput() *string
	ProcessAutomationJobArguments() *string
	SetProcessAutomationJobArguments(val *string)
	ProcessAutomationJobArgumentsInput() *string
	ProcessAutomationJobId() *string
	SetProcessAutomationJobId(val *string)
	ProcessAutomationJobIdInput() *string
	ProcessAutomationNodeFilter() *string
	SetProcessAutomationNodeFilter(val *string)
	ProcessAutomationNodeFilterInput() *string
	Script() *string
	SetScript(val *string)
	ScriptInput() *string
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
	ResetInvocationCommand()
	ResetProcessAutomationJobArguments()
	ResetProcessAutomationJobId()
	ResetProcessAutomationNodeFilter()
	ResetScript()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AutomationActionsActionActionDataReferenceOutputReference
type jsiiProxy_AutomationActionsActionActionDataReferenceOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AutomationActionsActionActionDataReferenceOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationActionsActionActionDataReferenceOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationActionsActionActionDataReferenceOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationActionsActionActionDataReferenceOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationActionsActionActionDataReferenceOutputReference) InternalValue() *AutomationActionsActionActionDataReference {
	var returns *AutomationActionsActionActionDataReference
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationActionsActionActionDataReferenceOutputReference) InvocationCommand() *string {
	var returns *string
	_jsii_.Get(
		j,
		"invocationCommand",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationActionsActionActionDataReferenceOutputReference) InvocationCommandInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"invocationCommandInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationActionsActionActionDataReferenceOutputReference) ProcessAutomationJobArguments() *string {
	var returns *string
	_jsii_.Get(
		j,
		"processAutomationJobArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationActionsActionActionDataReferenceOutputReference) ProcessAutomationJobArgumentsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"processAutomationJobArgumentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationActionsActionActionDataReferenceOutputReference) ProcessAutomationJobId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"processAutomationJobId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationActionsActionActionDataReferenceOutputReference) ProcessAutomationJobIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"processAutomationJobIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationActionsActionActionDataReferenceOutputReference) ProcessAutomationNodeFilter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"processAutomationNodeFilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationActionsActionActionDataReferenceOutputReference) ProcessAutomationNodeFilterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"processAutomationNodeFilterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationActionsActionActionDataReferenceOutputReference) Script() *string {
	var returns *string
	_jsii_.Get(
		j,
		"script",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationActionsActionActionDataReferenceOutputReference) ScriptInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scriptInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationActionsActionActionDataReferenceOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationActionsActionActionDataReferenceOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAutomationActionsActionActionDataReferenceOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AutomationActionsActionActionDataReferenceOutputReference {
	_init_.Initialize()

	if err := validateNewAutomationActionsActionActionDataReferenceOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AutomationActionsActionActionDataReferenceOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-pagerduty.automationActionsAction.AutomationActionsActionActionDataReferenceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAutomationActionsActionActionDataReferenceOutputReference_Override(a AutomationActionsActionActionDataReferenceOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-pagerduty.automationActionsAction.AutomationActionsActionActionDataReferenceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AutomationActionsActionActionDataReferenceOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AutomationActionsActionActionDataReferenceOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AutomationActionsActionActionDataReferenceOutputReference)SetInternalValue(val *AutomationActionsActionActionDataReference) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AutomationActionsActionActionDataReferenceOutputReference)SetInvocationCommand(val *string) {
	if err := j.validateSetInvocationCommandParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"invocationCommand",
		val,
	)
}

func (j *jsiiProxy_AutomationActionsActionActionDataReferenceOutputReference)SetProcessAutomationJobArguments(val *string) {
	if err := j.validateSetProcessAutomationJobArgumentsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"processAutomationJobArguments",
		val,
	)
}

func (j *jsiiProxy_AutomationActionsActionActionDataReferenceOutputReference)SetProcessAutomationJobId(val *string) {
	if err := j.validateSetProcessAutomationJobIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"processAutomationJobId",
		val,
	)
}

func (j *jsiiProxy_AutomationActionsActionActionDataReferenceOutputReference)SetProcessAutomationNodeFilter(val *string) {
	if err := j.validateSetProcessAutomationNodeFilterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"processAutomationNodeFilter",
		val,
	)
}

func (j *jsiiProxy_AutomationActionsActionActionDataReferenceOutputReference)SetScript(val *string) {
	if err := j.validateSetScriptParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"script",
		val,
	)
}

func (j *jsiiProxy_AutomationActionsActionActionDataReferenceOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AutomationActionsActionActionDataReferenceOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AutomationActionsActionActionDataReferenceOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationActionsActionActionDataReferenceOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AutomationActionsActionActionDataReferenceOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AutomationActionsActionActionDataReferenceOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AutomationActionsActionActionDataReferenceOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AutomationActionsActionActionDataReferenceOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AutomationActionsActionActionDataReferenceOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AutomationActionsActionActionDataReferenceOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AutomationActionsActionActionDataReferenceOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AutomationActionsActionActionDataReferenceOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AutomationActionsActionActionDataReferenceOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationActionsActionActionDataReferenceOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationActionsActionActionDataReferenceOutputReference) ResetInvocationCommand() {
	_jsii_.InvokeVoid(
		a,
		"resetInvocationCommand",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutomationActionsActionActionDataReferenceOutputReference) ResetProcessAutomationJobArguments() {
	_jsii_.InvokeVoid(
		a,
		"resetProcessAutomationJobArguments",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutomationActionsActionActionDataReferenceOutputReference) ResetProcessAutomationJobId() {
	_jsii_.InvokeVoid(
		a,
		"resetProcessAutomationJobId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutomationActionsActionActionDataReferenceOutputReference) ResetProcessAutomationNodeFilter() {
	_jsii_.InvokeVoid(
		a,
		"resetProcessAutomationNodeFilter",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutomationActionsActionActionDataReferenceOutputReference) ResetScript() {
	_jsii_.InvokeVoid(
		a,
		"resetScript",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutomationActionsActionActionDataReferenceOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := a.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationActionsActionActionDataReferenceOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

