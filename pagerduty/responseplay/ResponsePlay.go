// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package responseplay

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-pagerduty-go/pagerduty/v12/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-pagerduty-go/pagerduty/v12/responseplay/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.0.2/docs/resources/response_play pagerduty_response_play}.
type ResponsePlay interface {
	cdktf.TerraformResource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ConferenceNumber() *string
	SetConferenceNumber(val *string)
	ConferenceNumberInput() *string
	ConferenceUrl() *string
	SetConferenceUrl(val *string)
	ConferenceUrlInput() *string
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
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	From() *string
	SetFrom(val *string)
	FromInput() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
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
	Responder() ResponsePlayResponderList
	ResponderInput() interface{}
	RespondersMessage() *string
	SetRespondersMessage(val *string)
	RespondersMessageInput() *string
	Runnability() *string
	SetRunnability(val *string)
	RunnabilityInput() *string
	Subscriber() ResponsePlaySubscriberList
	SubscriberInput() interface{}
	SubscribersMessage() *string
	SetSubscribersMessage(val *string)
	SubscribersMessageInput() *string
	Team() *string
	SetTeam(val *string)
	TeamInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Type() *string
	SetType(val *string)
	TypeInput() *string
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
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
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutResponder(value interface{})
	PutSubscriber(value interface{})
	ResetConferenceNumber()
	ResetConferenceUrl()
	ResetDescription()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetResponder()
	ResetRespondersMessage()
	ResetRunnability()
	ResetSubscriber()
	ResetSubscribersMessage()
	ResetTeam()
	ResetType()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for ResponsePlay
type jsiiProxy_ResponsePlay struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ResponsePlay) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResponsePlay) ConferenceNumber() *string {
	var returns *string
	_jsii_.Get(
		j,
		"conferenceNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResponsePlay) ConferenceNumberInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"conferenceNumberInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResponsePlay) ConferenceUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"conferenceUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResponsePlay) ConferenceUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"conferenceUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResponsePlay) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResponsePlay) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResponsePlay) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResponsePlay) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResponsePlay) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResponsePlay) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResponsePlay) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResponsePlay) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResponsePlay) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResponsePlay) From() *string {
	var returns *string
	_jsii_.Get(
		j,
		"from",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResponsePlay) FromInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fromInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResponsePlay) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResponsePlay) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResponsePlay) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResponsePlay) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResponsePlay) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResponsePlay) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResponsePlay) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResponsePlay) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResponsePlay) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResponsePlay) Responder() ResponsePlayResponderList {
	var returns ResponsePlayResponderList
	_jsii_.Get(
		j,
		"responder",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResponsePlay) ResponderInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"responderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResponsePlay) RespondersMessage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"respondersMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResponsePlay) RespondersMessageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"respondersMessageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResponsePlay) Runnability() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runnability",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResponsePlay) RunnabilityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runnabilityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResponsePlay) Subscriber() ResponsePlaySubscriberList {
	var returns ResponsePlaySubscriberList
	_jsii_.Get(
		j,
		"subscriber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResponsePlay) SubscriberInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"subscriberInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResponsePlay) SubscribersMessage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subscribersMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResponsePlay) SubscribersMessageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subscribersMessageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResponsePlay) Team() *string {
	var returns *string
	_jsii_.Get(
		j,
		"team",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResponsePlay) TeamInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"teamInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResponsePlay) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResponsePlay) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResponsePlay) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResponsePlay) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResponsePlay) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.0.2/docs/resources/response_play pagerduty_response_play} Resource.
func NewResponsePlay(scope constructs.Construct, id *string, config *ResponsePlayConfig) ResponsePlay {
	_init_.Initialize()

	if err := validateNewResponsePlayParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ResponsePlay{}

	_jsii_.Create(
		"@cdktf/provider-pagerduty.responsePlay.ResponsePlay",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.0.2/docs/resources/response_play pagerduty_response_play} Resource.
func NewResponsePlay_Override(r ResponsePlay, scope constructs.Construct, id *string, config *ResponsePlayConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-pagerduty.responsePlay.ResponsePlay",
		[]interface{}{scope, id, config},
		r,
	)
}

func (j *jsiiProxy_ResponsePlay)SetConferenceNumber(val *string) {
	if err := j.validateSetConferenceNumberParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"conferenceNumber",
		val,
	)
}

func (j *jsiiProxy_ResponsePlay)SetConferenceUrl(val *string) {
	if err := j.validateSetConferenceUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"conferenceUrl",
		val,
	)
}

func (j *jsiiProxy_ResponsePlay)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ResponsePlay)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ResponsePlay)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ResponsePlay)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_ResponsePlay)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ResponsePlay)SetFrom(val *string) {
	if err := j.validateSetFromParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"from",
		val,
	)
}

func (j *jsiiProxy_ResponsePlay)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ResponsePlay)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ResponsePlay)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ResponsePlay)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ResponsePlay)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ResponsePlay)SetRespondersMessage(val *string) {
	if err := j.validateSetRespondersMessageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"respondersMessage",
		val,
	)
}

func (j *jsiiProxy_ResponsePlay)SetRunnability(val *string) {
	if err := j.validateSetRunnabilityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runnability",
		val,
	)
}

func (j *jsiiProxy_ResponsePlay)SetSubscribersMessage(val *string) {
	if err := j.validateSetSubscribersMessageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subscribersMessage",
		val,
	)
}

func (j *jsiiProxy_ResponsePlay)SetTeam(val *string) {
	if err := j.validateSetTeamParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"team",
		val,
	)
}

func (j *jsiiProxy_ResponsePlay)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

// Generates CDKTF code for importing a ResponsePlay resource upon running "cdktf plan <stack-name>".
func ResponsePlay_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateResponsePlay_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-pagerduty.responsePlay.ResponsePlay",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
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
func ResponsePlay_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateResponsePlay_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-pagerduty.responsePlay.ResponsePlay",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ResponsePlay_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateResponsePlay_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-pagerduty.responsePlay.ResponsePlay",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ResponsePlay_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateResponsePlay_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-pagerduty.responsePlay.ResponsePlay",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ResponsePlay_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-pagerduty.responsePlay.ResponsePlay",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (r *jsiiProxy_ResponsePlay) AddMoveTarget(moveTarget *string) {
	if err := r.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (r *jsiiProxy_ResponsePlay) AddOverride(path *string, value interface{}) {
	if err := r.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (r *jsiiProxy_ResponsePlay) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (r *jsiiProxy_ResponsePlay) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (r *jsiiProxy_ResponsePlay) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (r *jsiiProxy_ResponsePlay) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (r *jsiiProxy_ResponsePlay) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (r *jsiiProxy_ResponsePlay) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (r *jsiiProxy_ResponsePlay) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (r *jsiiProxy_ResponsePlay) GetStringAttribute(terraformAttribute *string) *string {
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

func (r *jsiiProxy_ResponsePlay) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (r *jsiiProxy_ResponsePlay) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := r.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (r *jsiiProxy_ResponsePlay) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := r.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResponsePlay) MoveTo(moveTarget *string, index interface{}) {
	if err := r.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (r *jsiiProxy_ResponsePlay) OverrideLogicalId(newLogicalId *string) {
	if err := r.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (r *jsiiProxy_ResponsePlay) PutResponder(value interface{}) {
	if err := r.validatePutResponderParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putResponder",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_ResponsePlay) PutSubscriber(value interface{}) {
	if err := r.validatePutSubscriberParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putSubscriber",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_ResponsePlay) ResetConferenceNumber() {
	_jsii_.InvokeVoid(
		r,
		"resetConferenceNumber",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ResponsePlay) ResetConferenceUrl() {
	_jsii_.InvokeVoid(
		r,
		"resetConferenceUrl",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ResponsePlay) ResetDescription() {
	_jsii_.InvokeVoid(
		r,
		"resetDescription",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ResponsePlay) ResetId() {
	_jsii_.InvokeVoid(
		r,
		"resetId",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ResponsePlay) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		r,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ResponsePlay) ResetResponder() {
	_jsii_.InvokeVoid(
		r,
		"resetResponder",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ResponsePlay) ResetRespondersMessage() {
	_jsii_.InvokeVoid(
		r,
		"resetRespondersMessage",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ResponsePlay) ResetRunnability() {
	_jsii_.InvokeVoid(
		r,
		"resetRunnability",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ResponsePlay) ResetSubscriber() {
	_jsii_.InvokeVoid(
		r,
		"resetSubscriber",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ResponsePlay) ResetSubscribersMessage() {
	_jsii_.InvokeVoid(
		r,
		"resetSubscribersMessage",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ResponsePlay) ResetTeam() {
	_jsii_.InvokeVoid(
		r,
		"resetTeam",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ResponsePlay) ResetType() {
	_jsii_.InvokeVoid(
		r,
		"resetType",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ResponsePlay) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResponsePlay) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResponsePlay) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResponsePlay) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

