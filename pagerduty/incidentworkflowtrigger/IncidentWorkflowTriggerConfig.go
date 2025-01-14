// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package incidentworkflowtrigger

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IncidentWorkflowTriggerConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.19.1/docs/resources/incident_workflow_trigger#subscribed_to_all_services IncidentWorkflowTrigger#subscribed_to_all_services}.
	SubscribedToAllServices interface{} `field:"required" json:"subscribedToAllServices" yaml:"subscribedToAllServices"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.19.1/docs/resources/incident_workflow_trigger#type IncidentWorkflowTrigger#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.19.1/docs/resources/incident_workflow_trigger#workflow IncidentWorkflowTrigger#workflow}.
	Workflow *string `field:"required" json:"workflow" yaml:"workflow"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.19.1/docs/resources/incident_workflow_trigger#condition IncidentWorkflowTrigger#condition}.
	Condition *string `field:"optional" json:"condition" yaml:"condition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.19.1/docs/resources/incident_workflow_trigger#id IncidentWorkflowTrigger#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// permissions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.19.1/docs/resources/incident_workflow_trigger#permissions IncidentWorkflowTrigger#permissions}
	Permissions *IncidentWorkflowTriggerPermissions `field:"optional" json:"permissions" yaml:"permissions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.19.1/docs/resources/incident_workflow_trigger#services IncidentWorkflowTrigger#services}.
	Services *[]*string `field:"optional" json:"services" yaml:"services"`
}

