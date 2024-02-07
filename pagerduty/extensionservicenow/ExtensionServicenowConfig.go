// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package extensionservicenow

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ExtensionServicenowConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.7.1/docs/resources/extension_servicenow#extension_objects ExtensionServicenow#extension_objects}.
	ExtensionObjects *[]*string `field:"required" json:"extensionObjects" yaml:"extensionObjects"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.7.1/docs/resources/extension_servicenow#extension_schema ExtensionServicenow#extension_schema}.
	ExtensionSchema *string `field:"required" json:"extensionSchema" yaml:"extensionSchema"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.7.1/docs/resources/extension_servicenow#referer ExtensionServicenow#referer}.
	Referer *string `field:"required" json:"referer" yaml:"referer"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.7.1/docs/resources/extension_servicenow#snow_password ExtensionServicenow#snow_password}.
	SnowPassword *string `field:"required" json:"snowPassword" yaml:"snowPassword"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.7.1/docs/resources/extension_servicenow#snow_user ExtensionServicenow#snow_user}.
	SnowUser *string `field:"required" json:"snowUser" yaml:"snowUser"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.7.1/docs/resources/extension_servicenow#sync_options ExtensionServicenow#sync_options}.
	SyncOptions *string `field:"required" json:"syncOptions" yaml:"syncOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.7.1/docs/resources/extension_servicenow#target ExtensionServicenow#target}.
	Target *string `field:"required" json:"target" yaml:"target"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.7.1/docs/resources/extension_servicenow#task_type ExtensionServicenow#task_type}.
	TaskType *string `field:"required" json:"taskType" yaml:"taskType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.7.1/docs/resources/extension_servicenow#endpoint_url ExtensionServicenow#endpoint_url}.
	EndpointUrl *string `field:"optional" json:"endpointUrl" yaml:"endpointUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.7.1/docs/resources/extension_servicenow#id ExtensionServicenow#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.7.1/docs/resources/extension_servicenow#name ExtensionServicenow#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.7.1/docs/resources/extension_servicenow#summary ExtensionServicenow#summary}.
	Summary *string `field:"optional" json:"summary" yaml:"summary"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.7.1/docs/resources/extension_servicenow#type ExtensionServicenow#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

