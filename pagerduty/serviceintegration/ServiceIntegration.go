package serviceintegration

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-pagerduty-go/pagerduty/v7/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-pagerduty-go/pagerduty/v7/serviceintegration/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/pagerduty/pagerduty/2.14.4/docs/resources/service_integration pagerduty_service_integration}.
type ServiceIntegration interface {
	cdktf.TerraformResource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	EmailFilter() ServiceIntegrationEmailFilterList
	EmailFilterInput() interface{}
	EmailFilterMode() *string
	SetEmailFilterMode(val *string)
	EmailFilterModeInput() *string
	EmailIncidentCreation() *string
	SetEmailIncidentCreation(val *string)
	EmailIncidentCreationInput() *string
	EmailParser() ServiceIntegrationEmailParserList
	EmailParserInput() interface{}
	EmailParsingFallback() *string
	SetEmailParsingFallback(val *string)
	EmailParsingFallbackInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	HtmlUrl() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	IntegrationEmail() *string
	SetIntegrationEmail(val *string)
	IntegrationEmailInput() *string
	IntegrationKey() *string
	SetIntegrationKey(val *string)
	IntegrationKeyInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	Service() *string
	SetService(val *string)
	ServiceInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Type() *string
	SetType(val *string)
	TypeInput() *string
	Vendor() *string
	SetVendor(val *string)
	VendorInput() *string
	// Experimental.
	AddOverride(path *string, value interface{})
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutEmailFilter(value interface{})
	PutEmailParser(value interface{})
	ResetEmailFilter()
	ResetEmailFilterMode()
	ResetEmailIncidentCreation()
	ResetEmailParser()
	ResetEmailParsingFallback()
	ResetId()
	ResetIntegrationEmail()
	ResetIntegrationKey()
	ResetName()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetType()
	ResetVendor()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for ServiceIntegration
type jsiiProxy_ServiceIntegration struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ServiceIntegration) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceIntegration) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceIntegration) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceIntegration) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceIntegration) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceIntegration) EmailFilter() ServiceIntegrationEmailFilterList {
	var returns ServiceIntegrationEmailFilterList
	_jsii_.Get(
		j,
		"emailFilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceIntegration) EmailFilterInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"emailFilterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceIntegration) EmailFilterMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailFilterMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceIntegration) EmailFilterModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailFilterModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceIntegration) EmailIncidentCreation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailIncidentCreation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceIntegration) EmailIncidentCreationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailIncidentCreationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceIntegration) EmailParser() ServiceIntegrationEmailParserList {
	var returns ServiceIntegrationEmailParserList
	_jsii_.Get(
		j,
		"emailParser",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceIntegration) EmailParserInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"emailParserInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceIntegration) EmailParsingFallback() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailParsingFallback",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceIntegration) EmailParsingFallbackInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailParsingFallbackInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceIntegration) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceIntegration) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceIntegration) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceIntegration) HtmlUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"htmlUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceIntegration) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceIntegration) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceIntegration) IntegrationEmail() *string {
	var returns *string
	_jsii_.Get(
		j,
		"integrationEmail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceIntegration) IntegrationEmailInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"integrationEmailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceIntegration) IntegrationKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"integrationKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceIntegration) IntegrationKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"integrationKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceIntegration) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceIntegration) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceIntegration) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceIntegration) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceIntegration) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceIntegration) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceIntegration) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceIntegration) Service() *string {
	var returns *string
	_jsii_.Get(
		j,
		"service",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceIntegration) ServiceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceIntegration) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceIntegration) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceIntegration) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceIntegration) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceIntegration) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceIntegration) Vendor() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vendor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceIntegration) VendorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vendorInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/pagerduty/pagerduty/2.14.4/docs/resources/service_integration pagerduty_service_integration} Resource.
func NewServiceIntegration(scope constructs.Construct, id *string, config *ServiceIntegrationConfig) ServiceIntegration {
	_init_.Initialize()

	if err := validateNewServiceIntegrationParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ServiceIntegration{}

	_jsii_.Create(
		"@cdktf/provider-pagerduty.serviceIntegration.ServiceIntegration",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/pagerduty/pagerduty/2.14.4/docs/resources/service_integration pagerduty_service_integration} Resource.
func NewServiceIntegration_Override(s ServiceIntegration, scope constructs.Construct, id *string, config *ServiceIntegrationConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-pagerduty.serviceIntegration.ServiceIntegration",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_ServiceIntegration)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ServiceIntegration)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ServiceIntegration)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ServiceIntegration)SetEmailFilterMode(val *string) {
	if err := j.validateSetEmailFilterModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"emailFilterMode",
		val,
	)
}

func (j *jsiiProxy_ServiceIntegration)SetEmailIncidentCreation(val *string) {
	if err := j.validateSetEmailIncidentCreationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"emailIncidentCreation",
		val,
	)
}

func (j *jsiiProxy_ServiceIntegration)SetEmailParsingFallback(val *string) {
	if err := j.validateSetEmailParsingFallbackParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"emailParsingFallback",
		val,
	)
}

func (j *jsiiProxy_ServiceIntegration)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ServiceIntegration)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ServiceIntegration)SetIntegrationEmail(val *string) {
	if err := j.validateSetIntegrationEmailParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"integrationEmail",
		val,
	)
}

func (j *jsiiProxy_ServiceIntegration)SetIntegrationKey(val *string) {
	if err := j.validateSetIntegrationKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"integrationKey",
		val,
	)
}

func (j *jsiiProxy_ServiceIntegration)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ServiceIntegration)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ServiceIntegration)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ServiceIntegration)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ServiceIntegration)SetService(val *string) {
	if err := j.validateSetServiceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"service",
		val,
	)
}

func (j *jsiiProxy_ServiceIntegration)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_ServiceIntegration)SetVendor(val *string) {
	if err := j.validateSetVendorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vendor",
		val,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func ServiceIntegration_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateServiceIntegration_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-pagerduty.serviceIntegration.ServiceIntegration",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ServiceIntegration_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateServiceIntegration_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-pagerduty.serviceIntegration.ServiceIntegration",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ServiceIntegration_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateServiceIntegration_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-pagerduty.serviceIntegration.ServiceIntegration",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ServiceIntegration_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-pagerduty.serviceIntegration.ServiceIntegration",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_ServiceIntegration) AddOverride(path *string, value interface{}) {
	if err := s.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_ServiceIntegration) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_ServiceIntegration) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_ServiceIntegration) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_ServiceIntegration) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_ServiceIntegration) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_ServiceIntegration) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_ServiceIntegration) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_ServiceIntegration) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_ServiceIntegration) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_ServiceIntegration) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceIntegration) OverrideLogicalId(newLogicalId *string) {
	if err := s.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_ServiceIntegration) PutEmailFilter(value interface{}) {
	if err := s.validatePutEmailFilterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putEmailFilter",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_ServiceIntegration) PutEmailParser(value interface{}) {
	if err := s.validatePutEmailParserParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putEmailParser",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_ServiceIntegration) ResetEmailFilter() {
	_jsii_.InvokeVoid(
		s,
		"resetEmailFilter",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceIntegration) ResetEmailFilterMode() {
	_jsii_.InvokeVoid(
		s,
		"resetEmailFilterMode",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceIntegration) ResetEmailIncidentCreation() {
	_jsii_.InvokeVoid(
		s,
		"resetEmailIncidentCreation",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceIntegration) ResetEmailParser() {
	_jsii_.InvokeVoid(
		s,
		"resetEmailParser",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceIntegration) ResetEmailParsingFallback() {
	_jsii_.InvokeVoid(
		s,
		"resetEmailParsingFallback",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceIntegration) ResetId() {
	_jsii_.InvokeVoid(
		s,
		"resetId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceIntegration) ResetIntegrationEmail() {
	_jsii_.InvokeVoid(
		s,
		"resetIntegrationEmail",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceIntegration) ResetIntegrationKey() {
	_jsii_.InvokeVoid(
		s,
		"resetIntegrationKey",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceIntegration) ResetName() {
	_jsii_.InvokeVoid(
		s,
		"resetName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceIntegration) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceIntegration) ResetType() {
	_jsii_.InvokeVoid(
		s,
		"resetType",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceIntegration) ResetVendor() {
	_jsii_.InvokeVoid(
		s,
		"resetVendor",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceIntegration) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceIntegration) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceIntegration) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceIntegration) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

