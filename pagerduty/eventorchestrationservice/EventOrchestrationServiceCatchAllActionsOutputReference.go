// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package eventorchestrationservice

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-pagerduty-go/pagerduty/v11/jsii"

	"github.com/cdktf/cdktf-provider-pagerduty-go/pagerduty/v11/eventorchestrationservice/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EventOrchestrationServiceCatchAllActionsOutputReference interface {
	cdktf.ComplexObject
	Annotate() *string
	SetAnnotate(val *string)
	AnnotateInput() *string
	AutomationAction() EventOrchestrationServiceCatchAllActionsAutomationActionOutputReference
	AutomationActionInput() *EventOrchestrationServiceCatchAllActionsAutomationAction
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
	Extraction() EventOrchestrationServiceCatchAllActionsExtractionList
	ExtractionInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *EventOrchestrationServiceCatchAllActions
	SetInternalValue(val *EventOrchestrationServiceCatchAllActions)
	PagerdutyAutomationAction() EventOrchestrationServiceCatchAllActionsPagerdutyAutomationActionOutputReference
	PagerdutyAutomationActionInput() *EventOrchestrationServiceCatchAllActionsPagerdutyAutomationAction
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
	Variable() EventOrchestrationServiceCatchAllActionsVariableList
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
	PutAutomationAction(value *EventOrchestrationServiceCatchAllActionsAutomationAction)
	PutExtraction(value interface{})
	PutPagerdutyAutomationAction(value *EventOrchestrationServiceCatchAllActionsPagerdutyAutomationAction)
	PutVariable(value interface{})
	ResetAnnotate()
	ResetAutomationAction()
	ResetEventAction()
	ResetExtraction()
	ResetPagerdutyAutomationAction()
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

// The jsii proxy struct for EventOrchestrationServiceCatchAllActionsOutputReference
type jsiiProxy_EventOrchestrationServiceCatchAllActionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_EventOrchestrationServiceCatchAllActionsOutputReference) Annotate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"annotate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationServiceCatchAllActionsOutputReference) AnnotateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"annotateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationServiceCatchAllActionsOutputReference) AutomationAction() EventOrchestrationServiceCatchAllActionsAutomationActionOutputReference {
	var returns EventOrchestrationServiceCatchAllActionsAutomationActionOutputReference
	_jsii_.Get(
		j,
		"automationAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationServiceCatchAllActionsOutputReference) AutomationActionInput() *EventOrchestrationServiceCatchAllActionsAutomationAction {
	var returns *EventOrchestrationServiceCatchAllActionsAutomationAction
	_jsii_.Get(
		j,
		"automationActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationServiceCatchAllActionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationServiceCatchAllActionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationServiceCatchAllActionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationServiceCatchAllActionsOutputReference) EventAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationServiceCatchAllActionsOutputReference) EventActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationServiceCatchAllActionsOutputReference) Extraction() EventOrchestrationServiceCatchAllActionsExtractionList {
	var returns EventOrchestrationServiceCatchAllActionsExtractionList
	_jsii_.Get(
		j,
		"extraction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationServiceCatchAllActionsOutputReference) ExtractionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"extractionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationServiceCatchAllActionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationServiceCatchAllActionsOutputReference) InternalValue() *EventOrchestrationServiceCatchAllActions {
	var returns *EventOrchestrationServiceCatchAllActions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationServiceCatchAllActionsOutputReference) PagerdutyAutomationAction() EventOrchestrationServiceCatchAllActionsPagerdutyAutomationActionOutputReference {
	var returns EventOrchestrationServiceCatchAllActionsPagerdutyAutomationActionOutputReference
	_jsii_.Get(
		j,
		"pagerdutyAutomationAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationServiceCatchAllActionsOutputReference) PagerdutyAutomationActionInput() *EventOrchestrationServiceCatchAllActionsPagerdutyAutomationAction {
	var returns *EventOrchestrationServiceCatchAllActionsPagerdutyAutomationAction
	_jsii_.Get(
		j,
		"pagerdutyAutomationActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationServiceCatchAllActionsOutputReference) Priority() *string {
	var returns *string
	_jsii_.Get(
		j,
		"priority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationServiceCatchAllActionsOutputReference) PriorityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"priorityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationServiceCatchAllActionsOutputReference) RouteTo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routeTo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationServiceCatchAllActionsOutputReference) RouteToInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routeToInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationServiceCatchAllActionsOutputReference) Severity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"severity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationServiceCatchAllActionsOutputReference) SeverityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"severityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationServiceCatchAllActionsOutputReference) Suppress() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"suppress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationServiceCatchAllActionsOutputReference) SuppressInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"suppressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationServiceCatchAllActionsOutputReference) Suspend() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"suspend",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationServiceCatchAllActionsOutputReference) SuspendInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"suspendInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationServiceCatchAllActionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationServiceCatchAllActionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationServiceCatchAllActionsOutputReference) Variable() EventOrchestrationServiceCatchAllActionsVariableList {
	var returns EventOrchestrationServiceCatchAllActionsVariableList
	_jsii_.Get(
		j,
		"variable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationServiceCatchAllActionsOutputReference) VariableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"variableInput",
		&returns,
	)
	return returns
}


func NewEventOrchestrationServiceCatchAllActionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) EventOrchestrationServiceCatchAllActionsOutputReference {
	_init_.Initialize()

	if err := validateNewEventOrchestrationServiceCatchAllActionsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_EventOrchestrationServiceCatchAllActionsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-pagerduty.eventOrchestrationService.EventOrchestrationServiceCatchAllActionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewEventOrchestrationServiceCatchAllActionsOutputReference_Override(e EventOrchestrationServiceCatchAllActionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-pagerduty.eventOrchestrationService.EventOrchestrationServiceCatchAllActionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_EventOrchestrationServiceCatchAllActionsOutputReference)SetAnnotate(val *string) {
	if err := j.validateSetAnnotateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"annotate",
		val,
	)
}

func (j *jsiiProxy_EventOrchestrationServiceCatchAllActionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EventOrchestrationServiceCatchAllActionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EventOrchestrationServiceCatchAllActionsOutputReference)SetEventAction(val *string) {
	if err := j.validateSetEventActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"eventAction",
		val,
	)
}

func (j *jsiiProxy_EventOrchestrationServiceCatchAllActionsOutputReference)SetInternalValue(val *EventOrchestrationServiceCatchAllActions) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EventOrchestrationServiceCatchAllActionsOutputReference)SetPriority(val *string) {
	if err := j.validateSetPriorityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"priority",
		val,
	)
}

func (j *jsiiProxy_EventOrchestrationServiceCatchAllActionsOutputReference)SetRouteTo(val *string) {
	if err := j.validateSetRouteToParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"routeTo",
		val,
	)
}

func (j *jsiiProxy_EventOrchestrationServiceCatchAllActionsOutputReference)SetSeverity(val *string) {
	if err := j.validateSetSeverityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"severity",
		val,
	)
}

func (j *jsiiProxy_EventOrchestrationServiceCatchAllActionsOutputReference)SetSuppress(val interface{}) {
	if err := j.validateSetSuppressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"suppress",
		val,
	)
}

func (j *jsiiProxy_EventOrchestrationServiceCatchAllActionsOutputReference)SetSuspend(val *float64) {
	if err := j.validateSetSuspendParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"suspend",
		val,
	)
}

func (j *jsiiProxy_EventOrchestrationServiceCatchAllActionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EventOrchestrationServiceCatchAllActionsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_EventOrchestrationServiceCatchAllActionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventOrchestrationServiceCatchAllActionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_EventOrchestrationServiceCatchAllActionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_EventOrchestrationServiceCatchAllActionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_EventOrchestrationServiceCatchAllActionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_EventOrchestrationServiceCatchAllActionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_EventOrchestrationServiceCatchAllActionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_EventOrchestrationServiceCatchAllActionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_EventOrchestrationServiceCatchAllActionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_EventOrchestrationServiceCatchAllActionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_EventOrchestrationServiceCatchAllActionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventOrchestrationServiceCatchAllActionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (e *jsiiProxy_EventOrchestrationServiceCatchAllActionsOutputReference) PutAutomationAction(value *EventOrchestrationServiceCatchAllActionsAutomationAction) {
	if err := e.validatePutAutomationActionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putAutomationAction",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventOrchestrationServiceCatchAllActionsOutputReference) PutExtraction(value interface{}) {
	if err := e.validatePutExtractionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putExtraction",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventOrchestrationServiceCatchAllActionsOutputReference) PutPagerdutyAutomationAction(value *EventOrchestrationServiceCatchAllActionsPagerdutyAutomationAction) {
	if err := e.validatePutPagerdutyAutomationActionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putPagerdutyAutomationAction",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventOrchestrationServiceCatchAllActionsOutputReference) PutVariable(value interface{}) {
	if err := e.validatePutVariableParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putVariable",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventOrchestrationServiceCatchAllActionsOutputReference) ResetAnnotate() {
	_jsii_.InvokeVoid(
		e,
		"resetAnnotate",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventOrchestrationServiceCatchAllActionsOutputReference) ResetAutomationAction() {
	_jsii_.InvokeVoid(
		e,
		"resetAutomationAction",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventOrchestrationServiceCatchAllActionsOutputReference) ResetEventAction() {
	_jsii_.InvokeVoid(
		e,
		"resetEventAction",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventOrchestrationServiceCatchAllActionsOutputReference) ResetExtraction() {
	_jsii_.InvokeVoid(
		e,
		"resetExtraction",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventOrchestrationServiceCatchAllActionsOutputReference) ResetPagerdutyAutomationAction() {
	_jsii_.InvokeVoid(
		e,
		"resetPagerdutyAutomationAction",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventOrchestrationServiceCatchAllActionsOutputReference) ResetPriority() {
	_jsii_.InvokeVoid(
		e,
		"resetPriority",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventOrchestrationServiceCatchAllActionsOutputReference) ResetRouteTo() {
	_jsii_.InvokeVoid(
		e,
		"resetRouteTo",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventOrchestrationServiceCatchAllActionsOutputReference) ResetSeverity() {
	_jsii_.InvokeVoid(
		e,
		"resetSeverity",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventOrchestrationServiceCatchAllActionsOutputReference) ResetSuppress() {
	_jsii_.InvokeVoid(
		e,
		"resetSuppress",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventOrchestrationServiceCatchAllActionsOutputReference) ResetSuspend() {
	_jsii_.InvokeVoid(
		e,
		"resetSuspend",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventOrchestrationServiceCatchAllActionsOutputReference) ResetVariable() {
	_jsii_.InvokeVoid(
		e,
		"resetVariable",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventOrchestrationServiceCatchAllActionsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (e *jsiiProxy_EventOrchestrationServiceCatchAllActionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

