package eventorchestrationglobal

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-pagerduty-go/pagerduty/v7/jsii"

	"github.com/cdktf/cdktf-provider-pagerduty-go/pagerduty/v7/eventorchestrationglobal/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EventOrchestrationGlobalCatchAllActionsOutputReference interface {
	cdktf.ComplexObject
	Annotate() *string
	SetAnnotate(val *string)
	AnnotateInput() *string
	AutomationAction() EventOrchestrationGlobalCatchAllActionsAutomationActionOutputReference
	AutomationActionInput() *EventOrchestrationGlobalCatchAllActionsAutomationAction
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
	DropEvent() interface{}
	SetDropEvent(val interface{})
	DropEventInput() interface{}
	EventAction() *string
	SetEventAction(val *string)
	EventActionInput() *string
	Extraction() EventOrchestrationGlobalCatchAllActionsExtractionList
	ExtractionInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *EventOrchestrationGlobalCatchAllActions
	SetInternalValue(val *EventOrchestrationGlobalCatchAllActions)
	Priority() *string
	SetPriority(val *string)
	PriorityInput() *string
	RouteTo() *string
	SetRouteTo(val *string)
	RouteToInput() *string
	Severity() *string
	SetSeverity(val *string)
	SeverityInput() *string
	Suppress() interface{}
	SetSuppress(val interface{})
	SuppressInput() interface{}
	Suspend() *float64
	SetSuspend(val *float64)
	SuspendInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Variable() EventOrchestrationGlobalCatchAllActionsVariableList
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
	PutAutomationAction(value *EventOrchestrationGlobalCatchAllActionsAutomationAction)
	PutExtraction(value interface{})
	PutVariable(value interface{})
	ResetAnnotate()
	ResetAutomationAction()
	ResetDropEvent()
	ResetEventAction()
	ResetExtraction()
	ResetPriority()
	ResetRouteTo()
	ResetSeverity()
	ResetSuppress()
	ResetSuspend()
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

// The jsii proxy struct for EventOrchestrationGlobalCatchAllActionsOutputReference
type jsiiProxy_EventOrchestrationGlobalCatchAllActionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_EventOrchestrationGlobalCatchAllActionsOutputReference) Annotate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"annotate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationGlobalCatchAllActionsOutputReference) AnnotateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"annotateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationGlobalCatchAllActionsOutputReference) AutomationAction() EventOrchestrationGlobalCatchAllActionsAutomationActionOutputReference {
	var returns EventOrchestrationGlobalCatchAllActionsAutomationActionOutputReference
	_jsii_.Get(
		j,
		"automationAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationGlobalCatchAllActionsOutputReference) AutomationActionInput() *EventOrchestrationGlobalCatchAllActionsAutomationAction {
	var returns *EventOrchestrationGlobalCatchAllActionsAutomationAction
	_jsii_.Get(
		j,
		"automationActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationGlobalCatchAllActionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationGlobalCatchAllActionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationGlobalCatchAllActionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationGlobalCatchAllActionsOutputReference) DropEvent() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dropEvent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationGlobalCatchAllActionsOutputReference) DropEventInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dropEventInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationGlobalCatchAllActionsOutputReference) EventAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationGlobalCatchAllActionsOutputReference) EventActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationGlobalCatchAllActionsOutputReference) Extraction() EventOrchestrationGlobalCatchAllActionsExtractionList {
	var returns EventOrchestrationGlobalCatchAllActionsExtractionList
	_jsii_.Get(
		j,
		"extraction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationGlobalCatchAllActionsOutputReference) ExtractionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"extractionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationGlobalCatchAllActionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationGlobalCatchAllActionsOutputReference) InternalValue() *EventOrchestrationGlobalCatchAllActions {
	var returns *EventOrchestrationGlobalCatchAllActions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationGlobalCatchAllActionsOutputReference) Priority() *string {
	var returns *string
	_jsii_.Get(
		j,
		"priority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationGlobalCatchAllActionsOutputReference) PriorityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"priorityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationGlobalCatchAllActionsOutputReference) RouteTo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routeTo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationGlobalCatchAllActionsOutputReference) RouteToInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routeToInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationGlobalCatchAllActionsOutputReference) Severity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"severity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationGlobalCatchAllActionsOutputReference) SeverityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"severityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationGlobalCatchAllActionsOutputReference) Suppress() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"suppress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationGlobalCatchAllActionsOutputReference) SuppressInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"suppressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationGlobalCatchAllActionsOutputReference) Suspend() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"suspend",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationGlobalCatchAllActionsOutputReference) SuspendInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"suspendInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationGlobalCatchAllActionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationGlobalCatchAllActionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationGlobalCatchAllActionsOutputReference) Variable() EventOrchestrationGlobalCatchAllActionsVariableList {
	var returns EventOrchestrationGlobalCatchAllActionsVariableList
	_jsii_.Get(
		j,
		"variable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationGlobalCatchAllActionsOutputReference) VariableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"variableInput",
		&returns,
	)
	return returns
}


func NewEventOrchestrationGlobalCatchAllActionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) EventOrchestrationGlobalCatchAllActionsOutputReference {
	_init_.Initialize()

	if err := validateNewEventOrchestrationGlobalCatchAllActionsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_EventOrchestrationGlobalCatchAllActionsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-pagerduty.eventOrchestrationGlobal.EventOrchestrationGlobalCatchAllActionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewEventOrchestrationGlobalCatchAllActionsOutputReference_Override(e EventOrchestrationGlobalCatchAllActionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-pagerduty.eventOrchestrationGlobal.EventOrchestrationGlobalCatchAllActionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_EventOrchestrationGlobalCatchAllActionsOutputReference)SetAnnotate(val *string) {
	if err := j.validateSetAnnotateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"annotate",
		val,
	)
}

func (j *jsiiProxy_EventOrchestrationGlobalCatchAllActionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EventOrchestrationGlobalCatchAllActionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EventOrchestrationGlobalCatchAllActionsOutputReference)SetDropEvent(val interface{}) {
	if err := j.validateSetDropEventParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dropEvent",
		val,
	)
}

func (j *jsiiProxy_EventOrchestrationGlobalCatchAllActionsOutputReference)SetEventAction(val *string) {
	if err := j.validateSetEventActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"eventAction",
		val,
	)
}

func (j *jsiiProxy_EventOrchestrationGlobalCatchAllActionsOutputReference)SetInternalValue(val *EventOrchestrationGlobalCatchAllActions) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EventOrchestrationGlobalCatchAllActionsOutputReference)SetPriority(val *string) {
	if err := j.validateSetPriorityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"priority",
		val,
	)
}

func (j *jsiiProxy_EventOrchestrationGlobalCatchAllActionsOutputReference)SetRouteTo(val *string) {
	if err := j.validateSetRouteToParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"routeTo",
		val,
	)
}

func (j *jsiiProxy_EventOrchestrationGlobalCatchAllActionsOutputReference)SetSeverity(val *string) {
	if err := j.validateSetSeverityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"severity",
		val,
	)
}

func (j *jsiiProxy_EventOrchestrationGlobalCatchAllActionsOutputReference)SetSuppress(val interface{}) {
	if err := j.validateSetSuppressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"suppress",
		val,
	)
}

func (j *jsiiProxy_EventOrchestrationGlobalCatchAllActionsOutputReference)SetSuspend(val *float64) {
	if err := j.validateSetSuspendParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"suspend",
		val,
	)
}

func (j *jsiiProxy_EventOrchestrationGlobalCatchAllActionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EventOrchestrationGlobalCatchAllActionsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_EventOrchestrationGlobalCatchAllActionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventOrchestrationGlobalCatchAllActionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_EventOrchestrationGlobalCatchAllActionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_EventOrchestrationGlobalCatchAllActionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_EventOrchestrationGlobalCatchAllActionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_EventOrchestrationGlobalCatchAllActionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_EventOrchestrationGlobalCatchAllActionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_EventOrchestrationGlobalCatchAllActionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_EventOrchestrationGlobalCatchAllActionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_EventOrchestrationGlobalCatchAllActionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_EventOrchestrationGlobalCatchAllActionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventOrchestrationGlobalCatchAllActionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := e.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventOrchestrationGlobalCatchAllActionsOutputReference) PutAutomationAction(value *EventOrchestrationGlobalCatchAllActionsAutomationAction) {
	if err := e.validatePutAutomationActionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putAutomationAction",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventOrchestrationGlobalCatchAllActionsOutputReference) PutExtraction(value interface{}) {
	if err := e.validatePutExtractionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putExtraction",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventOrchestrationGlobalCatchAllActionsOutputReference) PutVariable(value interface{}) {
	if err := e.validatePutVariableParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putVariable",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventOrchestrationGlobalCatchAllActionsOutputReference) ResetAnnotate() {
	_jsii_.InvokeVoid(
		e,
		"resetAnnotate",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventOrchestrationGlobalCatchAllActionsOutputReference) ResetAutomationAction() {
	_jsii_.InvokeVoid(
		e,
		"resetAutomationAction",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventOrchestrationGlobalCatchAllActionsOutputReference) ResetDropEvent() {
	_jsii_.InvokeVoid(
		e,
		"resetDropEvent",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventOrchestrationGlobalCatchAllActionsOutputReference) ResetEventAction() {
	_jsii_.InvokeVoid(
		e,
		"resetEventAction",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventOrchestrationGlobalCatchAllActionsOutputReference) ResetExtraction() {
	_jsii_.InvokeVoid(
		e,
		"resetExtraction",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventOrchestrationGlobalCatchAllActionsOutputReference) ResetPriority() {
	_jsii_.InvokeVoid(
		e,
		"resetPriority",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventOrchestrationGlobalCatchAllActionsOutputReference) ResetRouteTo() {
	_jsii_.InvokeVoid(
		e,
		"resetRouteTo",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventOrchestrationGlobalCatchAllActionsOutputReference) ResetSeverity() {
	_jsii_.InvokeVoid(
		e,
		"resetSeverity",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventOrchestrationGlobalCatchAllActionsOutputReference) ResetSuppress() {
	_jsii_.InvokeVoid(
		e,
		"resetSuppress",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventOrchestrationGlobalCatchAllActionsOutputReference) ResetSuspend() {
	_jsii_.InvokeVoid(
		e,
		"resetSuspend",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventOrchestrationGlobalCatchAllActionsOutputReference) ResetVariable() {
	_jsii_.InvokeVoid(
		e,
		"resetVariable",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventOrchestrationGlobalCatchAllActionsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := e.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventOrchestrationGlobalCatchAllActionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

