package serviceintegration

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-pagerduty-go/pagerduty/v5/jsii"

	"github.com/cdktf/cdktf-provider-pagerduty-go/pagerduty/v5/serviceintegration/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ServiceIntegrationEmailParserValueExtractorOutputReference interface {
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
	EndsBefore() *string
	SetEndsBefore(val *string)
	EndsBeforeInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Part() *string
	SetPart(val *string)
	PartInput() *string
	Regex() *string
	SetRegex(val *string)
	RegexInput() *string
	StartsAfter() *string
	SetStartsAfter(val *string)
	StartsAfterInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Type() *string
	SetType(val *string)
	TypeInput() *string
	ValueName() *string
	SetValueName(val *string)
	ValueNameInput() *string
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
	ResetEndsBefore()
	ResetRegex()
	ResetStartsAfter()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ServiceIntegrationEmailParserValueExtractorOutputReference
type jsiiProxy_ServiceIntegrationEmailParserValueExtractorOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ServiceIntegrationEmailParserValueExtractorOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceIntegrationEmailParserValueExtractorOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceIntegrationEmailParserValueExtractorOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceIntegrationEmailParserValueExtractorOutputReference) EndsBefore() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endsBefore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceIntegrationEmailParserValueExtractorOutputReference) EndsBeforeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endsBeforeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceIntegrationEmailParserValueExtractorOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceIntegrationEmailParserValueExtractorOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceIntegrationEmailParserValueExtractorOutputReference) Part() *string {
	var returns *string
	_jsii_.Get(
		j,
		"part",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceIntegrationEmailParserValueExtractorOutputReference) PartInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"partInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceIntegrationEmailParserValueExtractorOutputReference) Regex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceIntegrationEmailParserValueExtractorOutputReference) RegexInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceIntegrationEmailParserValueExtractorOutputReference) StartsAfter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startsAfter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceIntegrationEmailParserValueExtractorOutputReference) StartsAfterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startsAfterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceIntegrationEmailParserValueExtractorOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceIntegrationEmailParserValueExtractorOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceIntegrationEmailParserValueExtractorOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceIntegrationEmailParserValueExtractorOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceIntegrationEmailParserValueExtractorOutputReference) ValueName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"valueName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceIntegrationEmailParserValueExtractorOutputReference) ValueNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"valueNameInput",
		&returns,
	)
	return returns
}


func NewServiceIntegrationEmailParserValueExtractorOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ServiceIntegrationEmailParserValueExtractorOutputReference {
	_init_.Initialize()

	if err := validateNewServiceIntegrationEmailParserValueExtractorOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ServiceIntegrationEmailParserValueExtractorOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-pagerduty.serviceIntegration.ServiceIntegrationEmailParserValueExtractorOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewServiceIntegrationEmailParserValueExtractorOutputReference_Override(s ServiceIntegrationEmailParserValueExtractorOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-pagerduty.serviceIntegration.ServiceIntegrationEmailParserValueExtractorOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		s,
	)
}

func (j *jsiiProxy_ServiceIntegrationEmailParserValueExtractorOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ServiceIntegrationEmailParserValueExtractorOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ServiceIntegrationEmailParserValueExtractorOutputReference)SetEndsBefore(val *string) {
	if err := j.validateSetEndsBeforeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endsBefore",
		val,
	)
}

func (j *jsiiProxy_ServiceIntegrationEmailParserValueExtractorOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ServiceIntegrationEmailParserValueExtractorOutputReference)SetPart(val *string) {
	if err := j.validateSetPartParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"part",
		val,
	)
}

func (j *jsiiProxy_ServiceIntegrationEmailParserValueExtractorOutputReference)SetRegex(val *string) {
	if err := j.validateSetRegexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"regex",
		val,
	)
}

func (j *jsiiProxy_ServiceIntegrationEmailParserValueExtractorOutputReference)SetStartsAfter(val *string) {
	if err := j.validateSetStartsAfterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"startsAfter",
		val,
	)
}

func (j *jsiiProxy_ServiceIntegrationEmailParserValueExtractorOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ServiceIntegrationEmailParserValueExtractorOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ServiceIntegrationEmailParserValueExtractorOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_ServiceIntegrationEmailParserValueExtractorOutputReference)SetValueName(val *string) {
	if err := j.validateSetValueNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"valueName",
		val,
	)
}

func (s *jsiiProxy_ServiceIntegrationEmailParserValueExtractorOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceIntegrationEmailParserValueExtractorOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_ServiceIntegrationEmailParserValueExtractorOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_ServiceIntegrationEmailParserValueExtractorOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_ServiceIntegrationEmailParserValueExtractorOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_ServiceIntegrationEmailParserValueExtractorOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_ServiceIntegrationEmailParserValueExtractorOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_ServiceIntegrationEmailParserValueExtractorOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_ServiceIntegrationEmailParserValueExtractorOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_ServiceIntegrationEmailParserValueExtractorOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_ServiceIntegrationEmailParserValueExtractorOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceIntegrationEmailParserValueExtractorOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (s *jsiiProxy_ServiceIntegrationEmailParserValueExtractorOutputReference) ResetEndsBefore() {
	_jsii_.InvokeVoid(
		s,
		"resetEndsBefore",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceIntegrationEmailParserValueExtractorOutputReference) ResetRegex() {
	_jsii_.InvokeVoid(
		s,
		"resetRegex",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceIntegrationEmailParserValueExtractorOutputReference) ResetStartsAfter() {
	_jsii_.InvokeVoid(
		s,
		"resetStartsAfter",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceIntegrationEmailParserValueExtractorOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (s *jsiiProxy_ServiceIntegrationEmailParserValueExtractorOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
