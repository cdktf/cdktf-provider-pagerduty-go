// Prebuilt pagerduty Provider for Terraform CDK (cdktf)
package pagerduty

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-pagerduty-go/pagerduty/v2/jsii"

	"github.com/hashicorp/cdktf-provider-pagerduty-go/pagerduty/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EventOrchestrationUnroutedCatchAllActionsOutputReference interface {
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
	EventAction() *string
	SetEventAction(val *string)
	EventActionInput() *string
	Extraction() EventOrchestrationUnroutedCatchAllActionsExtractionList
	ExtractionInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *EventOrchestrationUnroutedCatchAllActions
	SetInternalValue(val *EventOrchestrationUnroutedCatchAllActions)
	Severity() *string
	SetSeverity(val *string)
	SeverityInput() *string
	Suppress() cdktf.IResolvable
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Variable() EventOrchestrationUnroutedCatchAllActionsVariableList
	VariableInput() interface{}
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
	PutExtraction(value interface{})
	PutVariable(value interface{})
	ResetEventAction()
	ResetExtraction()
	ResetSeverity()
	ResetVariable()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EventOrchestrationUnroutedCatchAllActionsOutputReference
type jsiiProxy_EventOrchestrationUnroutedCatchAllActionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_EventOrchestrationUnroutedCatchAllActionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationUnroutedCatchAllActionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationUnroutedCatchAllActionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationUnroutedCatchAllActionsOutputReference) EventAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationUnroutedCatchAllActionsOutputReference) EventActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationUnroutedCatchAllActionsOutputReference) Extraction() EventOrchestrationUnroutedCatchAllActionsExtractionList {
	var returns EventOrchestrationUnroutedCatchAllActionsExtractionList
	_jsii_.Get(
		j,
		"extraction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationUnroutedCatchAllActionsOutputReference) ExtractionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"extractionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationUnroutedCatchAllActionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationUnroutedCatchAllActionsOutputReference) InternalValue() *EventOrchestrationUnroutedCatchAllActions {
	var returns *EventOrchestrationUnroutedCatchAllActions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationUnroutedCatchAllActionsOutputReference) Severity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"severity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationUnroutedCatchAllActionsOutputReference) SeverityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"severityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationUnroutedCatchAllActionsOutputReference) Suppress() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"suppress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationUnroutedCatchAllActionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationUnroutedCatchAllActionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationUnroutedCatchAllActionsOutputReference) Variable() EventOrchestrationUnroutedCatchAllActionsVariableList {
	var returns EventOrchestrationUnroutedCatchAllActionsVariableList
	_jsii_.Get(
		j,
		"variable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationUnroutedCatchAllActionsOutputReference) VariableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"variableInput",
		&returns,
	)
	return returns
}


func NewEventOrchestrationUnroutedCatchAllActionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) EventOrchestrationUnroutedCatchAllActionsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_EventOrchestrationUnroutedCatchAllActionsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-pagerduty.EventOrchestrationUnroutedCatchAllActionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewEventOrchestrationUnroutedCatchAllActionsOutputReference_Override(e EventOrchestrationUnroutedCatchAllActionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-pagerduty.EventOrchestrationUnroutedCatchAllActionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_EventOrchestrationUnroutedCatchAllActionsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EventOrchestrationUnroutedCatchAllActionsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EventOrchestrationUnroutedCatchAllActionsOutputReference) SetEventAction(val *string) {
	_jsii_.Set(
		j,
		"eventAction",
		val,
	)
}

func (j *jsiiProxy_EventOrchestrationUnroutedCatchAllActionsOutputReference) SetInternalValue(val *EventOrchestrationUnroutedCatchAllActions) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EventOrchestrationUnroutedCatchAllActionsOutputReference) SetSeverity(val *string) {
	_jsii_.Set(
		j,
		"severity",
		val,
	)
}

func (j *jsiiProxy_EventOrchestrationUnroutedCatchAllActionsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EventOrchestrationUnroutedCatchAllActionsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_EventOrchestrationUnroutedCatchAllActionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventOrchestrationUnroutedCatchAllActionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventOrchestrationUnroutedCatchAllActionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventOrchestrationUnroutedCatchAllActionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventOrchestrationUnroutedCatchAllActionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventOrchestrationUnroutedCatchAllActionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventOrchestrationUnroutedCatchAllActionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventOrchestrationUnroutedCatchAllActionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventOrchestrationUnroutedCatchAllActionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventOrchestrationUnroutedCatchAllActionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventOrchestrationUnroutedCatchAllActionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventOrchestrationUnroutedCatchAllActionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventOrchestrationUnroutedCatchAllActionsOutputReference) PutExtraction(value interface{}) {
	_jsii_.InvokeVoid(
		e,
		"putExtraction",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventOrchestrationUnroutedCatchAllActionsOutputReference) PutVariable(value interface{}) {
	_jsii_.InvokeVoid(
		e,
		"putVariable",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventOrchestrationUnroutedCatchAllActionsOutputReference) ResetEventAction() {
	_jsii_.InvokeVoid(
		e,
		"resetEventAction",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventOrchestrationUnroutedCatchAllActionsOutputReference) ResetExtraction() {
	_jsii_.InvokeVoid(
		e,
		"resetExtraction",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventOrchestrationUnroutedCatchAllActionsOutputReference) ResetSeverity() {
	_jsii_.InvokeVoid(
		e,
		"resetSeverity",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventOrchestrationUnroutedCatchAllActionsOutputReference) ResetVariable() {
	_jsii_.InvokeVoid(
		e,
		"resetVariable",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventOrchestrationUnroutedCatchAllActionsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventOrchestrationUnroutedCatchAllActionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

