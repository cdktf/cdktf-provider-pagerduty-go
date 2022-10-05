package serviceeventrule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-pagerduty-go/pagerduty/v3/jsii"

	"github.com/cdktf/cdktf-provider-pagerduty-go/pagerduty/v3/serviceeventrule/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ServiceEventRuleActionsOutputReference interface {
	cdktf.ComplexObject
	Annotate() ServiceEventRuleActionsAnnotateList
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
	EventAction() ServiceEventRuleActionsEventActionList
	EventActionInput() interface{}
	Extractions() ServiceEventRuleActionsExtractionsList
	ExtractionsInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *ServiceEventRuleActions
	SetInternalValue(val *ServiceEventRuleActions)
	Priority() ServiceEventRuleActionsPriorityList
	PriorityInput() interface{}
	Severity() ServiceEventRuleActionsSeverityList
	SeverityInput() interface{}
	Suppress() ServiceEventRuleActionsSuppressList
	SuppressInput() interface{}
	Suspend() ServiceEventRuleActionsSuspendList
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
	PutSeverity(value interface{})
	PutSuppress(value interface{})
	PutSuspend(value interface{})
	ResetAnnotate()
	ResetEventAction()
	ResetExtractions()
	ResetPriority()
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

// The jsii proxy struct for ServiceEventRuleActionsOutputReference
type jsiiProxy_ServiceEventRuleActionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ServiceEventRuleActionsOutputReference) Annotate() ServiceEventRuleActionsAnnotateList {
	var returns ServiceEventRuleActionsAnnotateList
	_jsii_.Get(
		j,
		"annotate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceEventRuleActionsOutputReference) AnnotateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"annotateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceEventRuleActionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceEventRuleActionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceEventRuleActionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceEventRuleActionsOutputReference) EventAction() ServiceEventRuleActionsEventActionList {
	var returns ServiceEventRuleActionsEventActionList
	_jsii_.Get(
		j,
		"eventAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceEventRuleActionsOutputReference) EventActionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"eventActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceEventRuleActionsOutputReference) Extractions() ServiceEventRuleActionsExtractionsList {
	var returns ServiceEventRuleActionsExtractionsList
	_jsii_.Get(
		j,
		"extractions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceEventRuleActionsOutputReference) ExtractionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"extractionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceEventRuleActionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceEventRuleActionsOutputReference) InternalValue() *ServiceEventRuleActions {
	var returns *ServiceEventRuleActions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceEventRuleActionsOutputReference) Priority() ServiceEventRuleActionsPriorityList {
	var returns ServiceEventRuleActionsPriorityList
	_jsii_.Get(
		j,
		"priority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceEventRuleActionsOutputReference) PriorityInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"priorityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceEventRuleActionsOutputReference) Severity() ServiceEventRuleActionsSeverityList {
	var returns ServiceEventRuleActionsSeverityList
	_jsii_.Get(
		j,
		"severity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceEventRuleActionsOutputReference) SeverityInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"severityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceEventRuleActionsOutputReference) Suppress() ServiceEventRuleActionsSuppressList {
	var returns ServiceEventRuleActionsSuppressList
	_jsii_.Get(
		j,
		"suppress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceEventRuleActionsOutputReference) SuppressInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"suppressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceEventRuleActionsOutputReference) Suspend() ServiceEventRuleActionsSuspendList {
	var returns ServiceEventRuleActionsSuspendList
	_jsii_.Get(
		j,
		"suspend",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceEventRuleActionsOutputReference) SuspendInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"suspendInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceEventRuleActionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceEventRuleActionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewServiceEventRuleActionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ServiceEventRuleActionsOutputReference {
	_init_.Initialize()

	if err := validateNewServiceEventRuleActionsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ServiceEventRuleActionsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-pagerduty.serviceEventRule.ServiceEventRuleActionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewServiceEventRuleActionsOutputReference_Override(s ServiceEventRuleActionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-pagerduty.serviceEventRule.ServiceEventRuleActionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_ServiceEventRuleActionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ServiceEventRuleActionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ServiceEventRuleActionsOutputReference)SetInternalValue(val *ServiceEventRuleActions) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ServiceEventRuleActionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ServiceEventRuleActionsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_ServiceEventRuleActionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceEventRuleActionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceEventRuleActionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceEventRuleActionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceEventRuleActionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceEventRuleActionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceEventRuleActionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceEventRuleActionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceEventRuleActionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceEventRuleActionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceEventRuleActionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceEventRuleActionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceEventRuleActionsOutputReference) PutAnnotate(value interface{}) {
	if err := s.validatePutAnnotateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putAnnotate",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_ServiceEventRuleActionsOutputReference) PutEventAction(value interface{}) {
	if err := s.validatePutEventActionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putEventAction",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_ServiceEventRuleActionsOutputReference) PutExtractions(value interface{}) {
	if err := s.validatePutExtractionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putExtractions",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_ServiceEventRuleActionsOutputReference) PutPriority(value interface{}) {
	if err := s.validatePutPriorityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putPriority",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_ServiceEventRuleActionsOutputReference) PutSeverity(value interface{}) {
	if err := s.validatePutSeverityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putSeverity",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_ServiceEventRuleActionsOutputReference) PutSuppress(value interface{}) {
	if err := s.validatePutSuppressParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putSuppress",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_ServiceEventRuleActionsOutputReference) PutSuspend(value interface{}) {
	if err := s.validatePutSuspendParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putSuspend",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_ServiceEventRuleActionsOutputReference) ResetAnnotate() {
	_jsii_.InvokeVoid(
		s,
		"resetAnnotate",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceEventRuleActionsOutputReference) ResetEventAction() {
	_jsii_.InvokeVoid(
		s,
		"resetEventAction",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceEventRuleActionsOutputReference) ResetExtractions() {
	_jsii_.InvokeVoid(
		s,
		"resetExtractions",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceEventRuleActionsOutputReference) ResetPriority() {
	_jsii_.InvokeVoid(
		s,
		"resetPriority",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceEventRuleActionsOutputReference) ResetSeverity() {
	_jsii_.InvokeVoid(
		s,
		"resetSeverity",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceEventRuleActionsOutputReference) ResetSuppress() {
	_jsii_.InvokeVoid(
		s,
		"resetSuppress",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceEventRuleActionsOutputReference) ResetSuspend() {
	_jsii_.InvokeVoid(
		s,
		"resetSuspend",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceEventRuleActionsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := s.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceEventRuleActionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

