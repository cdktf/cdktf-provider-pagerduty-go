package serviceintegration

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-pagerduty-go/pagerduty/v6/jsii"

	"github.com/cdktf/cdktf-provider-pagerduty-go/pagerduty/v6/serviceintegration/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ServiceIntegrationEmailFilterOutputReference interface {
	cdktf.ComplexObject
	BodyMode() *string
	SetBodyMode(val *string)
	BodyModeInput() *string
	BodyRegex() *string
	SetBodyRegex(val *string)
	BodyRegexInput() *string
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
	FromEmailMode() *string
	SetFromEmailMode(val *string)
	FromEmailModeInput() *string
	FromEmailRegex() *string
	SetFromEmailRegex(val *string)
	FromEmailRegexInput() *string
	Id() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	SubjectMode() *string
	SetSubjectMode(val *string)
	SubjectModeInput() *string
	SubjectRegex() *string
	SetSubjectRegex(val *string)
	SubjectRegexInput() *string
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
	ResetBodyMode()
	ResetBodyRegex()
	ResetFromEmailMode()
	ResetFromEmailRegex()
	ResetSubjectMode()
	ResetSubjectRegex()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ServiceIntegrationEmailFilterOutputReference
type jsiiProxy_ServiceIntegrationEmailFilterOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ServiceIntegrationEmailFilterOutputReference) BodyMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bodyMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceIntegrationEmailFilterOutputReference) BodyModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bodyModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceIntegrationEmailFilterOutputReference) BodyRegex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bodyRegex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceIntegrationEmailFilterOutputReference) BodyRegexInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bodyRegexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceIntegrationEmailFilterOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceIntegrationEmailFilterOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceIntegrationEmailFilterOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceIntegrationEmailFilterOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceIntegrationEmailFilterOutputReference) FromEmailMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fromEmailMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceIntegrationEmailFilterOutputReference) FromEmailModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fromEmailModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceIntegrationEmailFilterOutputReference) FromEmailRegex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fromEmailRegex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceIntegrationEmailFilterOutputReference) FromEmailRegexInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fromEmailRegexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceIntegrationEmailFilterOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceIntegrationEmailFilterOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceIntegrationEmailFilterOutputReference) SubjectMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subjectMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceIntegrationEmailFilterOutputReference) SubjectModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subjectModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceIntegrationEmailFilterOutputReference) SubjectRegex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subjectRegex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceIntegrationEmailFilterOutputReference) SubjectRegexInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subjectRegexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceIntegrationEmailFilterOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceIntegrationEmailFilterOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewServiceIntegrationEmailFilterOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ServiceIntegrationEmailFilterOutputReference {
	_init_.Initialize()

	if err := validateNewServiceIntegrationEmailFilterOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ServiceIntegrationEmailFilterOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-pagerduty.serviceIntegration.ServiceIntegrationEmailFilterOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewServiceIntegrationEmailFilterOutputReference_Override(s ServiceIntegrationEmailFilterOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-pagerduty.serviceIntegration.ServiceIntegrationEmailFilterOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		s,
	)
}

func (j *jsiiProxy_ServiceIntegrationEmailFilterOutputReference)SetBodyMode(val *string) {
	if err := j.validateSetBodyModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bodyMode",
		val,
	)
}

func (j *jsiiProxy_ServiceIntegrationEmailFilterOutputReference)SetBodyRegex(val *string) {
	if err := j.validateSetBodyRegexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bodyRegex",
		val,
	)
}

func (j *jsiiProxy_ServiceIntegrationEmailFilterOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ServiceIntegrationEmailFilterOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ServiceIntegrationEmailFilterOutputReference)SetFromEmailMode(val *string) {
	if err := j.validateSetFromEmailModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fromEmailMode",
		val,
	)
}

func (j *jsiiProxy_ServiceIntegrationEmailFilterOutputReference)SetFromEmailRegex(val *string) {
	if err := j.validateSetFromEmailRegexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fromEmailRegex",
		val,
	)
}

func (j *jsiiProxy_ServiceIntegrationEmailFilterOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ServiceIntegrationEmailFilterOutputReference)SetSubjectMode(val *string) {
	if err := j.validateSetSubjectModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subjectMode",
		val,
	)
}

func (j *jsiiProxy_ServiceIntegrationEmailFilterOutputReference)SetSubjectRegex(val *string) {
	if err := j.validateSetSubjectRegexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subjectRegex",
		val,
	)
}

func (j *jsiiProxy_ServiceIntegrationEmailFilterOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ServiceIntegrationEmailFilterOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_ServiceIntegrationEmailFilterOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceIntegrationEmailFilterOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_ServiceIntegrationEmailFilterOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_ServiceIntegrationEmailFilterOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_ServiceIntegrationEmailFilterOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_ServiceIntegrationEmailFilterOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_ServiceIntegrationEmailFilterOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_ServiceIntegrationEmailFilterOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_ServiceIntegrationEmailFilterOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_ServiceIntegrationEmailFilterOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_ServiceIntegrationEmailFilterOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceIntegrationEmailFilterOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (s *jsiiProxy_ServiceIntegrationEmailFilterOutputReference) ResetBodyMode() {
	_jsii_.InvokeVoid(
		s,
		"resetBodyMode",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceIntegrationEmailFilterOutputReference) ResetBodyRegex() {
	_jsii_.InvokeVoid(
		s,
		"resetBodyRegex",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceIntegrationEmailFilterOutputReference) ResetFromEmailMode() {
	_jsii_.InvokeVoid(
		s,
		"resetFromEmailMode",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceIntegrationEmailFilterOutputReference) ResetFromEmailRegex() {
	_jsii_.InvokeVoid(
		s,
		"resetFromEmailRegex",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceIntegrationEmailFilterOutputReference) ResetSubjectMode() {
	_jsii_.InvokeVoid(
		s,
		"resetSubjectMode",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceIntegrationEmailFilterOutputReference) ResetSubjectRegex() {
	_jsii_.InvokeVoid(
		s,
		"resetSubjectRegex",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceIntegrationEmailFilterOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (s *jsiiProxy_ServiceIntegrationEmailFilterOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

