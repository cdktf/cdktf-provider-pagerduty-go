// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package eventorchestrationservice

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-pagerduty-go/pagerduty/v13/jsii"

	"github.com/cdktf/cdktf-provider-pagerduty-go/pagerduty/v13/eventorchestrationservice/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EventOrchestrationServiceSetRuleActionsOutputReference interface {
	cdktf.ComplexObject
	Annotate() *string
	SetAnnotate(val *string)
	AnnotateInput() *string
	AutomationAction() EventOrchestrationServiceSetRuleActionsAutomationActionOutputReference
	AutomationActionInput() *EventOrchestrationServiceSetRuleActionsAutomationAction
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
	EscalationPolicy() *string
	SetEscalationPolicy(val *string)
	EscalationPolicyInput() *string
	EventAction() *string
	SetEventAction(val *string)
	EventActionInput() *string
	Extraction() EventOrchestrationServiceSetRuleActionsExtractionList
	ExtractionInput() interface{}
	// Experimental.
	Fqn() *string
	IncidentCustomFieldUpdate() EventOrchestrationServiceSetRuleActionsIncidentCustomFieldUpdateList
	IncidentCustomFieldUpdateInput() interface{}
	InternalValue() *EventOrchestrationServiceSetRuleActions
	SetInternalValue(val *EventOrchestrationServiceSetRuleActions)
	PagerdutyAutomationAction() EventOrchestrationServiceSetRuleActionsPagerdutyAutomationActionOutputReference
	PagerdutyAutomationActionInput() *EventOrchestrationServiceSetRuleActionsPagerdutyAutomationAction
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
	Variable() EventOrchestrationServiceSetRuleActionsVariableList
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
	PutAutomationAction(value *EventOrchestrationServiceSetRuleActionsAutomationAction)
	PutExtraction(value interface{})
	PutIncidentCustomFieldUpdate(value interface{})
	PutPagerdutyAutomationAction(value *EventOrchestrationServiceSetRuleActionsPagerdutyAutomationAction)
	PutVariable(value interface{})
	ResetAnnotate()
	ResetAutomationAction()
	ResetEscalationPolicy()
	ResetEventAction()
	ResetExtraction()
	ResetIncidentCustomFieldUpdate()
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

// The jsii proxy struct for EventOrchestrationServiceSetRuleActionsOutputReference
type jsiiProxy_EventOrchestrationServiceSetRuleActionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_EventOrchestrationServiceSetRuleActionsOutputReference) Annotate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"annotate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationServiceSetRuleActionsOutputReference) AnnotateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"annotateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationServiceSetRuleActionsOutputReference) AutomationAction() EventOrchestrationServiceSetRuleActionsAutomationActionOutputReference {
	var returns EventOrchestrationServiceSetRuleActionsAutomationActionOutputReference
	_jsii_.Get(
		j,
		"automationAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationServiceSetRuleActionsOutputReference) AutomationActionInput() *EventOrchestrationServiceSetRuleActionsAutomationAction {
	var returns *EventOrchestrationServiceSetRuleActionsAutomationAction
	_jsii_.Get(
		j,
		"automationActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationServiceSetRuleActionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationServiceSetRuleActionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationServiceSetRuleActionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationServiceSetRuleActionsOutputReference) EscalationPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"escalationPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationServiceSetRuleActionsOutputReference) EscalationPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"escalationPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationServiceSetRuleActionsOutputReference) EventAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationServiceSetRuleActionsOutputReference) EventActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationServiceSetRuleActionsOutputReference) Extraction() EventOrchestrationServiceSetRuleActionsExtractionList {
	var returns EventOrchestrationServiceSetRuleActionsExtractionList
	_jsii_.Get(
		j,
		"extraction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationServiceSetRuleActionsOutputReference) ExtractionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"extractionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationServiceSetRuleActionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationServiceSetRuleActionsOutputReference) IncidentCustomFieldUpdate() EventOrchestrationServiceSetRuleActionsIncidentCustomFieldUpdateList {
	var returns EventOrchestrationServiceSetRuleActionsIncidentCustomFieldUpdateList
	_jsii_.Get(
		j,
		"incidentCustomFieldUpdate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationServiceSetRuleActionsOutputReference) IncidentCustomFieldUpdateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"incidentCustomFieldUpdateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationServiceSetRuleActionsOutputReference) InternalValue() *EventOrchestrationServiceSetRuleActions {
	var returns *EventOrchestrationServiceSetRuleActions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationServiceSetRuleActionsOutputReference) PagerdutyAutomationAction() EventOrchestrationServiceSetRuleActionsPagerdutyAutomationActionOutputReference {
	var returns EventOrchestrationServiceSetRuleActionsPagerdutyAutomationActionOutputReference
	_jsii_.Get(
		j,
		"pagerdutyAutomationAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationServiceSetRuleActionsOutputReference) PagerdutyAutomationActionInput() *EventOrchestrationServiceSetRuleActionsPagerdutyAutomationAction {
	var returns *EventOrchestrationServiceSetRuleActionsPagerdutyAutomationAction
	_jsii_.Get(
		j,
		"pagerdutyAutomationActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationServiceSetRuleActionsOutputReference) Priority() *string {
	var returns *string
	_jsii_.Get(
		j,
		"priority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationServiceSetRuleActionsOutputReference) PriorityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"priorityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationServiceSetRuleActionsOutputReference) RouteTo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routeTo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationServiceSetRuleActionsOutputReference) RouteToInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routeToInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationServiceSetRuleActionsOutputReference) Severity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"severity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationServiceSetRuleActionsOutputReference) SeverityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"severityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationServiceSetRuleActionsOutputReference) Suppress() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"suppress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationServiceSetRuleActionsOutputReference) SuppressInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"suppressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationServiceSetRuleActionsOutputReference) Suspend() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"suspend",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationServiceSetRuleActionsOutputReference) SuspendInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"suspendInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationServiceSetRuleActionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationServiceSetRuleActionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationServiceSetRuleActionsOutputReference) Variable() EventOrchestrationServiceSetRuleActionsVariableList {
	var returns EventOrchestrationServiceSetRuleActionsVariableList
	_jsii_.Get(
		j,
		"variable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationServiceSetRuleActionsOutputReference) VariableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"variableInput",
		&returns,
	)
	return returns
}


func NewEventOrchestrationServiceSetRuleActionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) EventOrchestrationServiceSetRuleActionsOutputReference {
	_init_.Initialize()

	if err := validateNewEventOrchestrationServiceSetRuleActionsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_EventOrchestrationServiceSetRuleActionsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-pagerduty.eventOrchestrationService.EventOrchestrationServiceSetRuleActionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewEventOrchestrationServiceSetRuleActionsOutputReference_Override(e EventOrchestrationServiceSetRuleActionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-pagerduty.eventOrchestrationService.EventOrchestrationServiceSetRuleActionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_EventOrchestrationServiceSetRuleActionsOutputReference)SetAnnotate(val *string) {
	if err := j.validateSetAnnotateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"annotate",
		val,
	)
}

func (j *jsiiProxy_EventOrchestrationServiceSetRuleActionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EventOrchestrationServiceSetRuleActionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EventOrchestrationServiceSetRuleActionsOutputReference)SetEscalationPolicy(val *string) {
	if err := j.validateSetEscalationPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"escalationPolicy",
		val,
	)
}

func (j *jsiiProxy_EventOrchestrationServiceSetRuleActionsOutputReference)SetEventAction(val *string) {
	if err := j.validateSetEventActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"eventAction",
		val,
	)
}

func (j *jsiiProxy_EventOrchestrationServiceSetRuleActionsOutputReference)SetInternalValue(val *EventOrchestrationServiceSetRuleActions) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EventOrchestrationServiceSetRuleActionsOutputReference)SetPriority(val *string) {
	if err := j.validateSetPriorityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"priority",
		val,
	)
}

func (j *jsiiProxy_EventOrchestrationServiceSetRuleActionsOutputReference)SetRouteTo(val *string) {
	if err := j.validateSetRouteToParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"routeTo",
		val,
	)
}

func (j *jsiiProxy_EventOrchestrationServiceSetRuleActionsOutputReference)SetSeverity(val *string) {
	if err := j.validateSetSeverityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"severity",
		val,
	)
}

func (j *jsiiProxy_EventOrchestrationServiceSetRuleActionsOutputReference)SetSuppress(val interface{}) {
	if err := j.validateSetSuppressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"suppress",
		val,
	)
}

func (j *jsiiProxy_EventOrchestrationServiceSetRuleActionsOutputReference)SetSuspend(val *float64) {
	if err := j.validateSetSuspendParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"suspend",
		val,
	)
}

func (j *jsiiProxy_EventOrchestrationServiceSetRuleActionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EventOrchestrationServiceSetRuleActionsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_EventOrchestrationServiceSetRuleActionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventOrchestrationServiceSetRuleActionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_EventOrchestrationServiceSetRuleActionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_EventOrchestrationServiceSetRuleActionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_EventOrchestrationServiceSetRuleActionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_EventOrchestrationServiceSetRuleActionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_EventOrchestrationServiceSetRuleActionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_EventOrchestrationServiceSetRuleActionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_EventOrchestrationServiceSetRuleActionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_EventOrchestrationServiceSetRuleActionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_EventOrchestrationServiceSetRuleActionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventOrchestrationServiceSetRuleActionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (e *jsiiProxy_EventOrchestrationServiceSetRuleActionsOutputReference) PutAutomationAction(value *EventOrchestrationServiceSetRuleActionsAutomationAction) {
	if err := e.validatePutAutomationActionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putAutomationAction",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventOrchestrationServiceSetRuleActionsOutputReference) PutExtraction(value interface{}) {
	if err := e.validatePutExtractionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putExtraction",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventOrchestrationServiceSetRuleActionsOutputReference) PutIncidentCustomFieldUpdate(value interface{}) {
	if err := e.validatePutIncidentCustomFieldUpdateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putIncidentCustomFieldUpdate",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventOrchestrationServiceSetRuleActionsOutputReference) PutPagerdutyAutomationAction(value *EventOrchestrationServiceSetRuleActionsPagerdutyAutomationAction) {
	if err := e.validatePutPagerdutyAutomationActionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putPagerdutyAutomationAction",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventOrchestrationServiceSetRuleActionsOutputReference) PutVariable(value interface{}) {
	if err := e.validatePutVariableParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putVariable",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventOrchestrationServiceSetRuleActionsOutputReference) ResetAnnotate() {
	_jsii_.InvokeVoid(
		e,
		"resetAnnotate",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventOrchestrationServiceSetRuleActionsOutputReference) ResetAutomationAction() {
	_jsii_.InvokeVoid(
		e,
		"resetAutomationAction",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventOrchestrationServiceSetRuleActionsOutputReference) ResetEscalationPolicy() {
	_jsii_.InvokeVoid(
		e,
		"resetEscalationPolicy",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventOrchestrationServiceSetRuleActionsOutputReference) ResetEventAction() {
	_jsii_.InvokeVoid(
		e,
		"resetEventAction",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventOrchestrationServiceSetRuleActionsOutputReference) ResetExtraction() {
	_jsii_.InvokeVoid(
		e,
		"resetExtraction",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventOrchestrationServiceSetRuleActionsOutputReference) ResetIncidentCustomFieldUpdate() {
	_jsii_.InvokeVoid(
		e,
		"resetIncidentCustomFieldUpdate",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventOrchestrationServiceSetRuleActionsOutputReference) ResetPagerdutyAutomationAction() {
	_jsii_.InvokeVoid(
		e,
		"resetPagerdutyAutomationAction",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventOrchestrationServiceSetRuleActionsOutputReference) ResetPriority() {
	_jsii_.InvokeVoid(
		e,
		"resetPriority",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventOrchestrationServiceSetRuleActionsOutputReference) ResetRouteTo() {
	_jsii_.InvokeVoid(
		e,
		"resetRouteTo",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventOrchestrationServiceSetRuleActionsOutputReference) ResetSeverity() {
	_jsii_.InvokeVoid(
		e,
		"resetSeverity",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventOrchestrationServiceSetRuleActionsOutputReference) ResetSuppress() {
	_jsii_.InvokeVoid(
		e,
		"resetSuppress",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventOrchestrationServiceSetRuleActionsOutputReference) ResetSuspend() {
	_jsii_.InvokeVoid(
		e,
		"resetSuspend",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventOrchestrationServiceSetRuleActionsOutputReference) ResetVariable() {
	_jsii_.InvokeVoid(
		e,
		"resetVariable",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventOrchestrationServiceSetRuleActionsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (e *jsiiProxy_EventOrchestrationServiceSetRuleActionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

