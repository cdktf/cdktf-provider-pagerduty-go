package provider

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-pagerduty-go/pagerduty/v7/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-pagerduty-go/pagerduty/v7/provider/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/pagerduty/pagerduty/2.15.0/docs pagerduty}.
type PagerdutyProvider interface {
	cdktf.TerraformProvider
	Alias() *string
	SetAlias(val *string)
	AliasInput() *string
	ApiUrlOverride() *string
	SetApiUrlOverride(val *string)
	ApiUrlOverrideInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	// Experimental.
	MetaAttributes() *map[string]interface{}
	// The tree node.
	Node() constructs.Node
	// Experimental.
	RawOverrides() interface{}
	ServiceRegion() *string
	SetServiceRegion(val *string)
	ServiceRegionInput() *string
	SkipCredentialsValidation() interface{}
	SetSkipCredentialsValidation(val interface{})
	SkipCredentialsValidationInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformProviderSource() *string
	// Experimental.
	TerraformResourceType() *string
	Token() *string
	SetToken(val *string)
	TokenInput() *string
	UserToken() *string
	SetUserToken(val *string)
	UserTokenInput() *string
	// Experimental.
	AddOverride(path *string, value interface{})
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	ResetAlias()
	ResetApiUrlOverride()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetServiceRegion()
	ResetSkipCredentialsValidation()
	ResetUserToken()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for PagerdutyProvider
type jsiiProxy_PagerdutyProvider struct {
	internal.Type__cdktfTerraformProvider
}

func (j *jsiiProxy_PagerdutyProvider) Alias() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagerdutyProvider) AliasInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aliasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagerdutyProvider) ApiUrlOverride() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiUrlOverride",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagerdutyProvider) ApiUrlOverrideInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiUrlOverrideInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagerdutyProvider) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagerdutyProvider) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagerdutyProvider) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagerdutyProvider) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagerdutyProvider) MetaAttributes() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"metaAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagerdutyProvider) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagerdutyProvider) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagerdutyProvider) ServiceRegion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagerdutyProvider) ServiceRegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceRegionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagerdutyProvider) SkipCredentialsValidation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipCredentialsValidation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagerdutyProvider) SkipCredentialsValidationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipCredentialsValidationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagerdutyProvider) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagerdutyProvider) TerraformProviderSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformProviderSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagerdutyProvider) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagerdutyProvider) Token() *string {
	var returns *string
	_jsii_.Get(
		j,
		"token",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagerdutyProvider) TokenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagerdutyProvider) UserToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagerdutyProvider) UserTokenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userTokenInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/pagerduty/pagerduty/2.15.0/docs pagerduty} Resource.
func NewPagerdutyProvider(scope constructs.Construct, id *string, config *PagerdutyProviderConfig) PagerdutyProvider {
	_init_.Initialize()

	if err := validateNewPagerdutyProviderParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_PagerdutyProvider{}

	_jsii_.Create(
		"@cdktf/provider-pagerduty.provider.PagerdutyProvider",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/pagerduty/pagerduty/2.15.0/docs pagerduty} Resource.
func NewPagerdutyProvider_Override(p PagerdutyProvider, scope constructs.Construct, id *string, config *PagerdutyProviderConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-pagerduty.provider.PagerdutyProvider",
		[]interface{}{scope, id, config},
		p,
	)
}

func (j *jsiiProxy_PagerdutyProvider)SetAlias(val *string) {
	_jsii_.Set(
		j,
		"alias",
		val,
	)
}

func (j *jsiiProxy_PagerdutyProvider)SetApiUrlOverride(val *string) {
	_jsii_.Set(
		j,
		"apiUrlOverride",
		val,
	)
}

func (j *jsiiProxy_PagerdutyProvider)SetServiceRegion(val *string) {
	_jsii_.Set(
		j,
		"serviceRegion",
		val,
	)
}

func (j *jsiiProxy_PagerdutyProvider)SetSkipCredentialsValidation(val interface{}) {
	if err := j.validateSetSkipCredentialsValidationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skipCredentialsValidation",
		val,
	)
}

func (j *jsiiProxy_PagerdutyProvider)SetToken(val *string) {
	_jsii_.Set(
		j,
		"token",
		val,
	)
}

func (j *jsiiProxy_PagerdutyProvider)SetUserToken(val *string) {
	_jsii_.Set(
		j,
		"userToken",
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
func PagerdutyProvider_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePagerdutyProvider_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-pagerduty.provider.PagerdutyProvider",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func PagerdutyProvider_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePagerdutyProvider_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-pagerduty.provider.PagerdutyProvider",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func PagerdutyProvider_IsTerraformProvider(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePagerdutyProvider_IsTerraformProviderParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-pagerduty.provider.PagerdutyProvider",
		"isTerraformProvider",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func PagerdutyProvider_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-pagerduty.provider.PagerdutyProvider",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (p *jsiiProxy_PagerdutyProvider) AddOverride(path *string, value interface{}) {
	if err := p.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (p *jsiiProxy_PagerdutyProvider) OverrideLogicalId(newLogicalId *string) {
	if err := p.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (p *jsiiProxy_PagerdutyProvider) ResetAlias() {
	_jsii_.InvokeVoid(
		p,
		"resetAlias",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PagerdutyProvider) ResetApiUrlOverride() {
	_jsii_.InvokeVoid(
		p,
		"resetApiUrlOverride",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PagerdutyProvider) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		p,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PagerdutyProvider) ResetServiceRegion() {
	_jsii_.InvokeVoid(
		p,
		"resetServiceRegion",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PagerdutyProvider) ResetSkipCredentialsValidation() {
	_jsii_.InvokeVoid(
		p,
		"resetSkipCredentialsValidation",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PagerdutyProvider) ResetUserToken() {
	_jsii_.InvokeVoid(
		p,
		"resetUserToken",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PagerdutyProvider) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PagerdutyProvider) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PagerdutyProvider) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PagerdutyProvider) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

