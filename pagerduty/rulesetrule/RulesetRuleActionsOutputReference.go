// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package rulesetrule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-pagerduty-go/pagerduty/v13/jsii"

	"github.com/cdktf/cdktf-provider-pagerduty-go/pagerduty/v13/rulesetrule/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type RulesetRuleActionsOutputReference interface {
	cdktf.ComplexObject
	Annotate() RulesetRuleActionsAnnotateList
	AnnotateInput() interface{}
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
	EventAction() RulesetRuleActionsEventActionList
	EventActionInput() interface{}
	Extractions() RulesetRuleActionsExtractionsList
	ExtractionsInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *RulesetRuleActions
	SetInternalValue(val *RulesetRuleActions)
	Priority() RulesetRuleActionsPriorityList
	PriorityInput() interface{}
	Route() RulesetRuleActionsRouteList
	RouteInput() interface{}
	Severity() RulesetRuleActionsSeverityList
	SeverityInput() interface{}
	Suppress() RulesetRuleActionsSuppressList
	SuppressInput() interface{}
	Suspend() RulesetRuleActionsSuspendList
	SuspendInput() interface{}
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
	PutAnnotate(value interface{})
	PutEventAction(value interface{})
	PutExtractions(value interface{})
	PutPriority(value interface{})
	PutRoute(value interface{})
	PutSeverity(value interface{})
	PutSuppress(value interface{})
	PutSuspend(value interface{})
	ResetAnnotate()
	ResetEventAction()
	ResetExtractions()
	ResetPriority()
	ResetRoute()
	ResetSeverity()
	ResetSuppress()
	ResetSuspend()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for RulesetRuleActionsOutputReference
type jsiiProxy_RulesetRuleActionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_RulesetRuleActionsOutputReference) Annotate() RulesetRuleActionsAnnotateList {
	var returns RulesetRuleActionsAnnotateList
	_jsii_.Get(
		j,
		"annotate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRuleActionsOutputReference) AnnotateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"annotateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRuleActionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRuleActionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRuleActionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRuleActionsOutputReference) EventAction() RulesetRuleActionsEventActionList {
	var returns RulesetRuleActionsEventActionList
	_jsii_.Get(
		j,
		"eventAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRuleActionsOutputReference) EventActionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"eventActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRuleActionsOutputReference) Extractions() RulesetRuleActionsExtractionsList {
	var returns RulesetRuleActionsExtractionsList
	_jsii_.Get(
		j,
		"extractions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRuleActionsOutputReference) ExtractionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"extractionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRuleActionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRuleActionsOutputReference) InternalValue() *RulesetRuleActions {
	var returns *RulesetRuleActions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRuleActionsOutputReference) Priority() RulesetRuleActionsPriorityList {
	var returns RulesetRuleActionsPriorityList
	_jsii_.Get(
		j,
		"priority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRuleActionsOutputReference) PriorityInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"priorityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRuleActionsOutputReference) Route() RulesetRuleActionsRouteList {
	var returns RulesetRuleActionsRouteList
	_jsii_.Get(
		j,
		"route",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRuleActionsOutputReference) RouteInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"routeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRuleActionsOutputReference) Severity() RulesetRuleActionsSeverityList {
	var returns RulesetRuleActionsSeverityList
	_jsii_.Get(
		j,
		"severity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRuleActionsOutputReference) SeverityInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"severityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRuleActionsOutputReference) Suppress() RulesetRuleActionsSuppressList {
	var returns RulesetRuleActionsSuppressList
	_jsii_.Get(
		j,
		"suppress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRuleActionsOutputReference) SuppressInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"suppressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRuleActionsOutputReference) Suspend() RulesetRuleActionsSuspendList {
	var returns RulesetRuleActionsSuspendList
	_jsii_.Get(
		j,
		"suspend",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRuleActionsOutputReference) SuspendInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"suspendInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRuleActionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RulesetRuleActionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewRulesetRuleActionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) RulesetRuleActionsOutputReference {
	_init_.Initialize()

	if err := validateNewRulesetRuleActionsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_RulesetRuleActionsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-pagerduty.rulesetRule.RulesetRuleActionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewRulesetRuleActionsOutputReference_Override(r RulesetRuleActionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-pagerduty.rulesetRule.RulesetRuleActionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		r,
	)
}

func (j *jsiiProxy_RulesetRuleActionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_RulesetRuleActionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_RulesetRuleActionsOutputReference)SetInternalValue(val *RulesetRuleActions) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_RulesetRuleActionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_RulesetRuleActionsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (r *jsiiProxy_RulesetRuleActionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RulesetRuleActionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := r.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RulesetRuleActionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := r.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RulesetRuleActionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := r.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		r,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RulesetRuleActionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := r.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		r,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RulesetRuleActionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := r.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		r,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RulesetRuleActionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := r.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		r,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RulesetRuleActionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := r.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		r,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RulesetRuleActionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := r.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		r,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RulesetRuleActionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := r.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		r,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RulesetRuleActionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RulesetRuleActionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := r.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RulesetRuleActionsOutputReference) PutAnnotate(value interface{}) {
	if err := r.validatePutAnnotateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putAnnotate",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RulesetRuleActionsOutputReference) PutEventAction(value interface{}) {
	if err := r.validatePutEventActionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putEventAction",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RulesetRuleActionsOutputReference) PutExtractions(value interface{}) {
	if err := r.validatePutExtractionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putExtractions",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RulesetRuleActionsOutputReference) PutPriority(value interface{}) {
	if err := r.validatePutPriorityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putPriority",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RulesetRuleActionsOutputReference) PutRoute(value interface{}) {
	if err := r.validatePutRouteParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putRoute",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RulesetRuleActionsOutputReference) PutSeverity(value interface{}) {
	if err := r.validatePutSeverityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putSeverity",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RulesetRuleActionsOutputReference) PutSuppress(value interface{}) {
	if err := r.validatePutSuppressParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putSuppress",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RulesetRuleActionsOutputReference) PutSuspend(value interface{}) {
	if err := r.validatePutSuspendParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putSuspend",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RulesetRuleActionsOutputReference) ResetAnnotate() {
	_jsii_.InvokeVoid(
		r,
		"resetAnnotate",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RulesetRuleActionsOutputReference) ResetEventAction() {
	_jsii_.InvokeVoid(
		r,
		"resetEventAction",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RulesetRuleActionsOutputReference) ResetExtractions() {
	_jsii_.InvokeVoid(
		r,
		"resetExtractions",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RulesetRuleActionsOutputReference) ResetPriority() {
	_jsii_.InvokeVoid(
		r,
		"resetPriority",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RulesetRuleActionsOutputReference) ResetRoute() {
	_jsii_.InvokeVoid(
		r,
		"resetRoute",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RulesetRuleActionsOutputReference) ResetSeverity() {
	_jsii_.InvokeVoid(
		r,
		"resetSeverity",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RulesetRuleActionsOutputReference) ResetSuppress() {
	_jsii_.InvokeVoid(
		r,
		"resetSuppress",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RulesetRuleActionsOutputReference) ResetSuspend() {
	_jsii_.InvokeVoid(
		r,
		"resetSuspend",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RulesetRuleActionsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := r.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		r,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RulesetRuleActionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

