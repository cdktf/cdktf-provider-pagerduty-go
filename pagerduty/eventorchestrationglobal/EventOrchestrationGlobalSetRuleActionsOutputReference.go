// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package eventorchestrationglobal

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-pagerduty-go/pagerduty/v13/jsii"

	"github.com/cdktf/cdktf-provider-pagerduty-go/pagerduty/v13/eventorchestrationglobal/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EventOrchestrationGlobalSetRuleActionsOutputReference interface {
	cdktf.ComplexObject
	Annotate() *string
	SetAnnotate(val *string)
	AnnotateInput() *string
	AutomationAction() EventOrchestrationGlobalSetRuleActionsAutomationActionOutputReference
	AutomationActionInput() *EventOrchestrationGlobalSetRuleActionsAutomationAction
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
	EscalationPolicy() *string
	SetEscalationPolicy(val *string)
	EscalationPolicyInput() *string
	EventAction() *string
	SetEventAction(val *string)
	EventActionInput() *string
	Extraction() EventOrchestrationGlobalSetRuleActionsExtractionList
	ExtractionInput() interface{}
	// Experimental.
	Fqn() *string
	IncidentCustomFieldUpdate() EventOrchestrationGlobalSetRuleActionsIncidentCustomFieldUpdateList
	IncidentCustomFieldUpdateInput() interface{}
	InternalValue() *EventOrchestrationGlobalSetRuleActions
	SetInternalValue(val *EventOrchestrationGlobalSetRuleActions)
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
	Variable() EventOrchestrationGlobalSetRuleActionsVariableList
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
	PutAutomationAction(value *EventOrchestrationGlobalSetRuleActionsAutomationAction)
	PutExtraction(value interface{})
	PutIncidentCustomFieldUpdate(value interface{})
	PutVariable(value interface{})
	ResetAnnotate()
	ResetAutomationAction()
	ResetDropEvent()
	ResetEscalationPolicy()
	ResetEventAction()
	ResetExtraction()
	ResetIncidentCustomFieldUpdate()
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

// The jsii proxy struct for EventOrchestrationGlobalSetRuleActionsOutputReference
type jsiiProxy_EventOrchestrationGlobalSetRuleActionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_EventOrchestrationGlobalSetRuleActionsOutputReference) Annotate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"annotate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationGlobalSetRuleActionsOutputReference) AnnotateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"annotateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationGlobalSetRuleActionsOutputReference) AutomationAction() EventOrchestrationGlobalSetRuleActionsAutomationActionOutputReference {
	var returns EventOrchestrationGlobalSetRuleActionsAutomationActionOutputReference
	_jsii_.Get(
		j,
		"automationAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationGlobalSetRuleActionsOutputReference) AutomationActionInput() *EventOrchestrationGlobalSetRuleActionsAutomationAction {
	var returns *EventOrchestrationGlobalSetRuleActionsAutomationAction
	_jsii_.Get(
		j,
		"automationActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationGlobalSetRuleActionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationGlobalSetRuleActionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationGlobalSetRuleActionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationGlobalSetRuleActionsOutputReference) DropEvent() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dropEvent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationGlobalSetRuleActionsOutputReference) DropEventInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dropEventInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationGlobalSetRuleActionsOutputReference) EscalationPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"escalationPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationGlobalSetRuleActionsOutputReference) EscalationPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"escalationPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationGlobalSetRuleActionsOutputReference) EventAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationGlobalSetRuleActionsOutputReference) EventActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationGlobalSetRuleActionsOutputReference) Extraction() EventOrchestrationGlobalSetRuleActionsExtractionList {
	var returns EventOrchestrationGlobalSetRuleActionsExtractionList
	_jsii_.Get(
		j,
		"extraction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationGlobalSetRuleActionsOutputReference) ExtractionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"extractionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationGlobalSetRuleActionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationGlobalSetRuleActionsOutputReference) IncidentCustomFieldUpdate() EventOrchestrationGlobalSetRuleActionsIncidentCustomFieldUpdateList {
	var returns EventOrchestrationGlobalSetRuleActionsIncidentCustomFieldUpdateList
	_jsii_.Get(
		j,
		"incidentCustomFieldUpdate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationGlobalSetRuleActionsOutputReference) IncidentCustomFieldUpdateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"incidentCustomFieldUpdateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationGlobalSetRuleActionsOutputReference) InternalValue() *EventOrchestrationGlobalSetRuleActions {
	var returns *EventOrchestrationGlobalSetRuleActions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationGlobalSetRuleActionsOutputReference) Priority() *string {
	var returns *string
	_jsii_.Get(
		j,
		"priority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationGlobalSetRuleActionsOutputReference) PriorityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"priorityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationGlobalSetRuleActionsOutputReference) RouteTo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routeTo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationGlobalSetRuleActionsOutputReference) RouteToInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routeToInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationGlobalSetRuleActionsOutputReference) Severity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"severity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationGlobalSetRuleActionsOutputReference) SeverityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"severityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationGlobalSetRuleActionsOutputReference) Suppress() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"suppress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationGlobalSetRuleActionsOutputReference) SuppressInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"suppressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationGlobalSetRuleActionsOutputReference) Suspend() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"suspend",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationGlobalSetRuleActionsOutputReference) SuspendInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"suspendInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationGlobalSetRuleActionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationGlobalSetRuleActionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationGlobalSetRuleActionsOutputReference) Variable() EventOrchestrationGlobalSetRuleActionsVariableList {
	var returns EventOrchestrationGlobalSetRuleActionsVariableList
	_jsii_.Get(
		j,
		"variable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationGlobalSetRuleActionsOutputReference) VariableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"variableInput",
		&returns,
	)
	return returns
}


func NewEventOrchestrationGlobalSetRuleActionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) EventOrchestrationGlobalSetRuleActionsOutputReference {
	_init_.Initialize()

	if err := validateNewEventOrchestrationGlobalSetRuleActionsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_EventOrchestrationGlobalSetRuleActionsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-pagerduty.eventOrchestrationGlobal.EventOrchestrationGlobalSetRuleActionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewEventOrchestrationGlobalSetRuleActionsOutputReference_Override(e EventOrchestrationGlobalSetRuleActionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-pagerduty.eventOrchestrationGlobal.EventOrchestrationGlobalSetRuleActionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_EventOrchestrationGlobalSetRuleActionsOutputReference)SetAnnotate(val *string) {
	if err := j.validateSetAnnotateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"annotate",
		val,
	)
}

func (j *jsiiProxy_EventOrchestrationGlobalSetRuleActionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EventOrchestrationGlobalSetRuleActionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EventOrchestrationGlobalSetRuleActionsOutputReference)SetDropEvent(val interface{}) {
	if err := j.validateSetDropEventParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dropEvent",
		val,
	)
}

func (j *jsiiProxy_EventOrchestrationGlobalSetRuleActionsOutputReference)SetEscalationPolicy(val *string) {
	if err := j.validateSetEscalationPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"escalationPolicy",
		val,
	)
}

func (j *jsiiProxy_EventOrchestrationGlobalSetRuleActionsOutputReference)SetEventAction(val *string) {
	if err := j.validateSetEventActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"eventAction",
		val,
	)
}

func (j *jsiiProxy_EventOrchestrationGlobalSetRuleActionsOutputReference)SetInternalValue(val *EventOrchestrationGlobalSetRuleActions) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EventOrchestrationGlobalSetRuleActionsOutputReference)SetPriority(val *string) {
	if err := j.validateSetPriorityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"priority",
		val,
	)
}

func (j *jsiiProxy_EventOrchestrationGlobalSetRuleActionsOutputReference)SetRouteTo(val *string) {
	if err := j.validateSetRouteToParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"routeTo",
		val,
	)
}

func (j *jsiiProxy_EventOrchestrationGlobalSetRuleActionsOutputReference)SetSeverity(val *string) {
	if err := j.validateSetSeverityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"severity",
		val,
	)
}

func (j *jsiiProxy_EventOrchestrationGlobalSetRuleActionsOutputReference)SetSuppress(val interface{}) {
	if err := j.validateSetSuppressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"suppress",
		val,
	)
}

func (j *jsiiProxy_EventOrchestrationGlobalSetRuleActionsOutputReference)SetSuspend(val *float64) {
	if err := j.validateSetSuspendParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"suspend",
		val,
	)
}

func (j *jsiiProxy_EventOrchestrationGlobalSetRuleActionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EventOrchestrationGlobalSetRuleActionsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_EventOrchestrationGlobalSetRuleActionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventOrchestrationGlobalSetRuleActionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_EventOrchestrationGlobalSetRuleActionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_EventOrchestrationGlobalSetRuleActionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_EventOrchestrationGlobalSetRuleActionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_EventOrchestrationGlobalSetRuleActionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_EventOrchestrationGlobalSetRuleActionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_EventOrchestrationGlobalSetRuleActionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_EventOrchestrationGlobalSetRuleActionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_EventOrchestrationGlobalSetRuleActionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_EventOrchestrationGlobalSetRuleActionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventOrchestrationGlobalSetRuleActionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (e *jsiiProxy_EventOrchestrationGlobalSetRuleActionsOutputReference) PutAutomationAction(value *EventOrchestrationGlobalSetRuleActionsAutomationAction) {
	if err := e.validatePutAutomationActionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putAutomationAction",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventOrchestrationGlobalSetRuleActionsOutputReference) PutExtraction(value interface{}) {
	if err := e.validatePutExtractionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putExtraction",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventOrchestrationGlobalSetRuleActionsOutputReference) PutIncidentCustomFieldUpdate(value interface{}) {
	if err := e.validatePutIncidentCustomFieldUpdateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putIncidentCustomFieldUpdate",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventOrchestrationGlobalSetRuleActionsOutputReference) PutVariable(value interface{}) {
	if err := e.validatePutVariableParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putVariable",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventOrchestrationGlobalSetRuleActionsOutputReference) ResetAnnotate() {
	_jsii_.InvokeVoid(
		e,
		"resetAnnotate",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventOrchestrationGlobalSetRuleActionsOutputReference) ResetAutomationAction() {
	_jsii_.InvokeVoid(
		e,
		"resetAutomationAction",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventOrchestrationGlobalSetRuleActionsOutputReference) ResetDropEvent() {
	_jsii_.InvokeVoid(
		e,
		"resetDropEvent",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventOrchestrationGlobalSetRuleActionsOutputReference) ResetEscalationPolicy() {
	_jsii_.InvokeVoid(
		e,
		"resetEscalationPolicy",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventOrchestrationGlobalSetRuleActionsOutputReference) ResetEventAction() {
	_jsii_.InvokeVoid(
		e,
		"resetEventAction",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventOrchestrationGlobalSetRuleActionsOutputReference) ResetExtraction() {
	_jsii_.InvokeVoid(
		e,
		"resetExtraction",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventOrchestrationGlobalSetRuleActionsOutputReference) ResetIncidentCustomFieldUpdate() {
	_jsii_.InvokeVoid(
		e,
		"resetIncidentCustomFieldUpdate",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventOrchestrationGlobalSetRuleActionsOutputReference) ResetPriority() {
	_jsii_.InvokeVoid(
		e,
		"resetPriority",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventOrchestrationGlobalSetRuleActionsOutputReference) ResetRouteTo() {
	_jsii_.InvokeVoid(
		e,
		"resetRouteTo",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventOrchestrationGlobalSetRuleActionsOutputReference) ResetSeverity() {
	_jsii_.InvokeVoid(
		e,
		"resetSeverity",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventOrchestrationGlobalSetRuleActionsOutputReference) ResetSuppress() {
	_jsii_.InvokeVoid(
		e,
		"resetSuppress",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventOrchestrationGlobalSetRuleActionsOutputReference) ResetSuspend() {
	_jsii_.InvokeVoid(
		e,
		"resetSuspend",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventOrchestrationGlobalSetRuleActionsOutputReference) ResetVariable() {
	_jsii_.InvokeVoid(
		e,
		"resetVariable",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventOrchestrationGlobalSetRuleActionsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (e *jsiiProxy_EventOrchestrationGlobalSetRuleActionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

