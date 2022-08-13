// Prebuilt pagerduty Provider for Terraform CDK (cdktf)
package pagerduty

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-pagerduty-go/pagerduty/v2/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-pagerduty-go/pagerduty/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/pagerduty/r/service pagerduty_service}.
type Service interface {
	cdktf.TerraformResource
	AcknowledgementTimeout() *string
	SetAcknowledgementTimeout(val *string)
	AcknowledgementTimeoutInput() *string
	AlertCreation() *string
	SetAlertCreation(val *string)
	AlertCreationInput() *string
	AlertGrouping() *string
	SetAlertGrouping(val *string)
	AlertGroupingInput() *string
	AlertGroupingParameters() ServiceAlertGroupingParametersOutputReference
	AlertGroupingParametersInput() *ServiceAlertGroupingParameters
	AlertGroupingTimeout() *string
	SetAlertGroupingTimeout(val *string)
	AlertGroupingTimeoutInput() *string
	AutoResolveTimeout() *string
	SetAutoResolveTimeout(val *string)
	AutoResolveTimeoutInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
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
	CreatedAt() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	EscalationPolicy() *string
	SetEscalationPolicy(val *string)
	EscalationPolicyInput() *string
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
	IncidentUrgencyRule() ServiceIncidentUrgencyRuleOutputReference
	IncidentUrgencyRuleInput() *ServiceIncidentUrgencyRule
	LastIncidentTimestamp() *string
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
	ScheduledActions() ServiceScheduledActionsList
	ScheduledActionsInput() interface{}
	Status() *string
	SupportHours() ServiceSupportHoursOutputReference
	SupportHoursInput() *ServiceSupportHours
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Type() *string
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
	PutAlertGroupingParameters(value *ServiceAlertGroupingParameters)
	PutIncidentUrgencyRule(value *ServiceIncidentUrgencyRule)
	PutScheduledActions(value interface{})
	PutSupportHours(value *ServiceSupportHours)
	ResetAcknowledgementTimeout()
	ResetAlertCreation()
	ResetAlertGrouping()
	ResetAlertGroupingParameters()
	ResetAlertGroupingTimeout()
	ResetAutoResolveTimeout()
	ResetDescription()
	ResetId()
	ResetIncidentUrgencyRule()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetScheduledActions()
	ResetSupportHours()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for Service
type jsiiProxy_Service struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_Service) AcknowledgementTimeout() *string {
	var returns *string
	_jsii_.Get(
		j,
		"acknowledgementTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) AcknowledgementTimeoutInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"acknowledgementTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) AlertCreation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alertCreation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) AlertCreationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alertCreationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) AlertGrouping() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alertGrouping",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) AlertGroupingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alertGroupingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) AlertGroupingParameters() ServiceAlertGroupingParametersOutputReference {
	var returns ServiceAlertGroupingParametersOutputReference
	_jsii_.Get(
		j,
		"alertGroupingParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) AlertGroupingParametersInput() *ServiceAlertGroupingParameters {
	var returns *ServiceAlertGroupingParameters
	_jsii_.Get(
		j,
		"alertGroupingParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) AlertGroupingTimeout() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alertGroupingTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) AlertGroupingTimeoutInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alertGroupingTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) AutoResolveTimeout() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoResolveTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) AutoResolveTimeoutInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoResolveTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) EscalationPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"escalationPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) EscalationPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"escalationPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) HtmlUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"htmlUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) IncidentUrgencyRule() ServiceIncidentUrgencyRuleOutputReference {
	var returns ServiceIncidentUrgencyRuleOutputReference
	_jsii_.Get(
		j,
		"incidentUrgencyRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) IncidentUrgencyRuleInput() *ServiceIncidentUrgencyRule {
	var returns *ServiceIncidentUrgencyRule
	_jsii_.Get(
		j,
		"incidentUrgencyRuleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) LastIncidentTimestamp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastIncidentTimestamp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) ScheduledActions() ServiceScheduledActionsList {
	var returns ServiceScheduledActionsList
	_jsii_.Get(
		j,
		"scheduledActions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) ScheduledActionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scheduledActionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) SupportHours() ServiceSupportHoursOutputReference {
	var returns ServiceSupportHoursOutputReference
	_jsii_.Get(
		j,
		"supportHours",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) SupportHoursInput() *ServiceSupportHours {
	var returns *ServiceSupportHours
	_jsii_.Get(
		j,
		"supportHoursInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Service) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/pagerduty/r/service pagerduty_service} Resource.
func NewService(scope constructs.Construct, id *string, config *ServiceConfig) Service {
	_init_.Initialize()

	j := jsiiProxy_Service{}

	_jsii_.Create(
		"@cdktf/provider-pagerduty.Service",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/pagerduty/r/service pagerduty_service} Resource.
func NewService_Override(s Service, scope constructs.Construct, id *string, config *ServiceConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-pagerduty.Service",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_Service) SetAcknowledgementTimeout(val *string) {
	_jsii_.Set(
		j,
		"acknowledgementTimeout",
		val,
	)
}

func (j *jsiiProxy_Service) SetAlertCreation(val *string) {
	_jsii_.Set(
		j,
		"alertCreation",
		val,
	)
}

func (j *jsiiProxy_Service) SetAlertGrouping(val *string) {
	_jsii_.Set(
		j,
		"alertGrouping",
		val,
	)
}

func (j *jsiiProxy_Service) SetAlertGroupingTimeout(val *string) {
	_jsii_.Set(
		j,
		"alertGroupingTimeout",
		val,
	)
}

func (j *jsiiProxy_Service) SetAutoResolveTimeout(val *string) {
	_jsii_.Set(
		j,
		"autoResolveTimeout",
		val,
	)
}

func (j *jsiiProxy_Service) SetConnection(val interface{}) {
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_Service) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_Service) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Service) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_Service) SetEscalationPolicy(val *string) {
	_jsii_.Set(
		j,
		"escalationPolicy",
		val,
	)
}

func (j *jsiiProxy_Service) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_Service) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_Service) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_Service) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_Service) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_Service) SetProvisioners(val *[]interface{}) {
	_jsii_.Set(
		j,
		"provisioners",
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
func Service_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-pagerduty.Service",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func Service_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-pagerduty.Service",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_Service) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_Service) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Service) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Service) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Service) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Service) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Service) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Service) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Service) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Service) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Service) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Service) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_Service) PutAlertGroupingParameters(value *ServiceAlertGroupingParameters) {
	_jsii_.InvokeVoid(
		s,
		"putAlertGroupingParameters",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_Service) PutIncidentUrgencyRule(value *ServiceIncidentUrgencyRule) {
	_jsii_.InvokeVoid(
		s,
		"putIncidentUrgencyRule",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_Service) PutScheduledActions(value interface{}) {
	_jsii_.InvokeVoid(
		s,
		"putScheduledActions",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_Service) PutSupportHours(value *ServiceSupportHours) {
	_jsii_.InvokeVoid(
		s,
		"putSupportHours",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_Service) ResetAcknowledgementTimeout() {
	_jsii_.InvokeVoid(
		s,
		"resetAcknowledgementTimeout",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Service) ResetAlertCreation() {
	_jsii_.InvokeVoid(
		s,
		"resetAlertCreation",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Service) ResetAlertGrouping() {
	_jsii_.InvokeVoid(
		s,
		"resetAlertGrouping",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Service) ResetAlertGroupingParameters() {
	_jsii_.InvokeVoid(
		s,
		"resetAlertGroupingParameters",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Service) ResetAlertGroupingTimeout() {
	_jsii_.InvokeVoid(
		s,
		"resetAlertGroupingTimeout",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Service) ResetAutoResolveTimeout() {
	_jsii_.InvokeVoid(
		s,
		"resetAutoResolveTimeout",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Service) ResetDescription() {
	_jsii_.InvokeVoid(
		s,
		"resetDescription",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Service) ResetId() {
	_jsii_.InvokeVoid(
		s,
		"resetId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Service) ResetIncidentUrgencyRule() {
	_jsii_.InvokeVoid(
		s,
		"resetIncidentUrgencyRule",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Service) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Service) ResetScheduledActions() {
	_jsii_.InvokeVoid(
		s,
		"resetScheduledActions",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Service) ResetSupportHours() {
	_jsii_.InvokeVoid(
		s,
		"resetSupportHours",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Service) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Service) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Service) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Service) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

