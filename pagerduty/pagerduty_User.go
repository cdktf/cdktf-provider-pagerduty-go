// Prebuilt pagerduty Provider for Terraform CDK (cdktf)
package pagerduty

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-pagerduty-go/pagerduty/v2/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-pagerduty-go/pagerduty/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/pagerduty/r/user pagerduty_user}.
type User interface {
	cdktf.TerraformResource
	AvatarUrl() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	Color() *string
	SetColor(val *string)
	ColorInput() *string
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() *float64
	// Experimental.
	SetCount(val *float64)
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	Email() *string
	SetEmail(val *string)
	EmailInput() *string
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
	InvitationSent() cdktf.IResolvable
	JobTitle() *string
	SetJobTitle(val *string)
	JobTitleInput() *string
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
	Role() *string
	SetRole(val *string)
	RoleInput() *string
	Teams() *[]*string
	SetTeams(val *[]*string)
	TeamsInput() *[]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	TimeZone() *string
	SetTimeZone(val *string)
	TimeZoneInput() *string
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
	ResetColor()
	ResetDescription()
	ResetId()
	ResetJobTitle()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRole()
	ResetTeams()
	ResetTimeZone()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for User
type jsiiProxy_User struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_User) AvatarUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"avatarUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) Color() *string {
	var returns *string
	_jsii_.Get(
		j,
		"color",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) ColorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"colorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) Email() *string {
	var returns *string
	_jsii_.Get(
		j,
		"email",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) EmailInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) HtmlUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"htmlUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) InvitationSent() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"invitationSent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) JobTitle() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobTitle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) JobTitleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobTitleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) Role() *string {
	var returns *string
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) RoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) Teams() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"teams",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) TeamsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"teamsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) TimeZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) TimeZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeZoneInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/pagerduty/r/user pagerduty_user} Resource.
func NewUser(scope constructs.Construct, id *string, config *UserConfig) User {
	_init_.Initialize()

	j := jsiiProxy_User{}

	_jsii_.Create(
		"@cdktf/provider-pagerduty.User",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/pagerduty/r/user pagerduty_user} Resource.
func NewUser_Override(u User, scope constructs.Construct, id *string, config *UserConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-pagerduty.User",
		[]interface{}{scope, id, config},
		u,
	)
}

func (j *jsiiProxy_User) SetColor(val *string) {
	_jsii_.Set(
		j,
		"color",
		val,
	)
}

func (j *jsiiProxy_User) SetConnection(val interface{}) {
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_User) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_User) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_User) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_User) SetEmail(val *string) {
	_jsii_.Set(
		j,
		"email",
		val,
	)
}

func (j *jsiiProxy_User) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_User) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_User) SetJobTitle(val *string) {
	_jsii_.Set(
		j,
		"jobTitle",
		val,
	)
}

func (j *jsiiProxy_User) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_User) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_User) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_User) SetProvisioners(val *[]interface{}) {
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_User) SetRole(val *string) {
	_jsii_.Set(
		j,
		"role",
		val,
	)
}

func (j *jsiiProxy_User) SetTeams(val *[]*string) {
	_jsii_.Set(
		j,
		"teams",
		val,
	)
}

func (j *jsiiProxy_User) SetTimeZone(val *string) {
	_jsii_.Set(
		j,
		"timeZone",
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
func User_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-pagerduty.User",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func User_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-pagerduty.User",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (u *jsiiProxy_User) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		u,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (u *jsiiProxy_User) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		u,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_User) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		u,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_User) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		u,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_User) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		u,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_User) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		u,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_User) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		u,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_User) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		u,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_User) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		u,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_User) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		u,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_User) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		u,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_User) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		u,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (u *jsiiProxy_User) ResetColor() {
	_jsii_.InvokeVoid(
		u,
		"resetColor",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetDescription() {
	_jsii_.InvokeVoid(
		u,
		"resetDescription",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetId() {
	_jsii_.InvokeVoid(
		u,
		"resetId",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetJobTitle() {
	_jsii_.InvokeVoid(
		u,
		"resetJobTitle",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		u,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetRole() {
	_jsii_.InvokeVoid(
		u,
		"resetRole",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetTeams() {
	_jsii_.InvokeVoid(
		u,
		"resetTeams",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetTimeZone() {
	_jsii_.InvokeVoid(
		u,
		"resetTimeZone",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		u,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (u *jsiiProxy_User) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		u,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (u *jsiiProxy_User) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		u,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (u *jsiiProxy_User) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		u,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

