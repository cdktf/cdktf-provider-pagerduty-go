// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package eventorchestrationservice

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-pagerduty-go/pagerduty/v13/jsii"

	"github.com/cdktf/cdktf-provider-pagerduty-go/pagerduty/v13/eventorchestrationservice/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EventOrchestrationServiceSetRuleActionsAutomationActionOutputReference interface {
	cdktf.ComplexObject
	AutoSend() interface{}
	SetAutoSend(val interface{})
	AutoSendInput() interface{}
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
	Header() EventOrchestrationServiceSetRuleActionsAutomationActionHeaderList
	HeaderInput() interface{}
	InternalValue() *EventOrchestrationServiceSetRuleActionsAutomationAction
	SetInternalValue(val *EventOrchestrationServiceSetRuleActionsAutomationAction)
	Name() *string
	SetName(val *string)
	NameInput() *string
	Parameter() EventOrchestrationServiceSetRuleActionsAutomationActionParameterList
	ParameterInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Url() *string
	SetUrl(val *string)
	UrlInput() *string
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
	PutHeader(value interface{})
	PutParameter(value interface{})
	ResetAutoSend()
	ResetHeader()
	ResetParameter()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EventOrchestrationServiceSetRuleActionsAutomationActionOutputReference
type jsiiProxy_EventOrchestrationServiceSetRuleActionsAutomationActionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_EventOrchestrationServiceSetRuleActionsAutomationActionOutputReference) AutoSend() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoSend",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationServiceSetRuleActionsAutomationActionOutputReference) AutoSendInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoSendInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationServiceSetRuleActionsAutomationActionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationServiceSetRuleActionsAutomationActionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationServiceSetRuleActionsAutomationActionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationServiceSetRuleActionsAutomationActionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationServiceSetRuleActionsAutomationActionOutputReference) Header() EventOrchestrationServiceSetRuleActionsAutomationActionHeaderList {
	var returns EventOrchestrationServiceSetRuleActionsAutomationActionHeaderList
	_jsii_.Get(
		j,
		"header",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationServiceSetRuleActionsAutomationActionOutputReference) HeaderInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"headerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationServiceSetRuleActionsAutomationActionOutputReference) InternalValue() *EventOrchestrationServiceSetRuleActionsAutomationAction {
	var returns *EventOrchestrationServiceSetRuleActionsAutomationAction
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationServiceSetRuleActionsAutomationActionOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationServiceSetRuleActionsAutomationActionOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationServiceSetRuleActionsAutomationActionOutputReference) Parameter() EventOrchestrationServiceSetRuleActionsAutomationActionParameterList {
	var returns EventOrchestrationServiceSetRuleActionsAutomationActionParameterList
	_jsii_.Get(
		j,
		"parameter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationServiceSetRuleActionsAutomationActionOutputReference) ParameterInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"parameterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationServiceSetRuleActionsAutomationActionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationServiceSetRuleActionsAutomationActionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationServiceSetRuleActionsAutomationActionOutputReference) Url() *string {
	var returns *string
	_jsii_.Get(
		j,
		"url",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventOrchestrationServiceSetRuleActionsAutomationActionOutputReference) UrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"urlInput",
		&returns,
	)
	return returns
}


func NewEventOrchestrationServiceSetRuleActionsAutomationActionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) EventOrchestrationServiceSetRuleActionsAutomationActionOutputReference {
	_init_.Initialize()

	if err := validateNewEventOrchestrationServiceSetRuleActionsAutomationActionOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_EventOrchestrationServiceSetRuleActionsAutomationActionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-pagerduty.eventOrchestrationService.EventOrchestrationServiceSetRuleActionsAutomationActionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewEventOrchestrationServiceSetRuleActionsAutomationActionOutputReference_Override(e EventOrchestrationServiceSetRuleActionsAutomationActionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-pagerduty.eventOrchestrationService.EventOrchestrationServiceSetRuleActionsAutomationActionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_EventOrchestrationServiceSetRuleActionsAutomationActionOutputReference)SetAutoSend(val interface{}) {
	if err := j.validateSetAutoSendParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoSend",
		val,
	)
}

func (j *jsiiProxy_EventOrchestrationServiceSetRuleActionsAutomationActionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EventOrchestrationServiceSetRuleActionsAutomationActionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EventOrchestrationServiceSetRuleActionsAutomationActionOutputReference)SetInternalValue(val *EventOrchestrationServiceSetRuleActionsAutomationAction) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EventOrchestrationServiceSetRuleActionsAutomationActionOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_EventOrchestrationServiceSetRuleActionsAutomationActionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EventOrchestrationServiceSetRuleActionsAutomationActionOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_EventOrchestrationServiceSetRuleActionsAutomationActionOutputReference)SetUrl(val *string) {
	if err := j.validateSetUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"url",
		val,
	)
}

func (e *jsiiProxy_EventOrchestrationServiceSetRuleActionsAutomationActionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventOrchestrationServiceSetRuleActionsAutomationActionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_EventOrchestrationServiceSetRuleActionsAutomationActionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_EventOrchestrationServiceSetRuleActionsAutomationActionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_EventOrchestrationServiceSetRuleActionsAutomationActionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_EventOrchestrationServiceSetRuleActionsAutomationActionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_EventOrchestrationServiceSetRuleActionsAutomationActionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_EventOrchestrationServiceSetRuleActionsAutomationActionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_EventOrchestrationServiceSetRuleActionsAutomationActionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_EventOrchestrationServiceSetRuleActionsAutomationActionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_EventOrchestrationServiceSetRuleActionsAutomationActionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventOrchestrationServiceSetRuleActionsAutomationActionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (e *jsiiProxy_EventOrchestrationServiceSetRuleActionsAutomationActionOutputReference) PutHeader(value interface{}) {
	if err := e.validatePutHeaderParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putHeader",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventOrchestrationServiceSetRuleActionsAutomationActionOutputReference) PutParameter(value interface{}) {
	if err := e.validatePutParameterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putParameter",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventOrchestrationServiceSetRuleActionsAutomationActionOutputReference) ResetAutoSend() {
	_jsii_.InvokeVoid(
		e,
		"resetAutoSend",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventOrchestrationServiceSetRuleActionsAutomationActionOutputReference) ResetHeader() {
	_jsii_.InvokeVoid(
		e,
		"resetHeader",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventOrchestrationServiceSetRuleActionsAutomationActionOutputReference) ResetParameter() {
	_jsii_.InvokeVoid(
		e,
		"resetParameter",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventOrchestrationServiceSetRuleActionsAutomationActionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (e *jsiiProxy_EventOrchestrationServiceSetRuleActionsAutomationActionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

