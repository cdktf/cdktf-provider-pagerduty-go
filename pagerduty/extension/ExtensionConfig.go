// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package extension

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ExtensionConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.15.6/docs/resources/extension#extension_objects Extension#extension_objects}.
	ExtensionObjects *[]*string `field:"required" json:"extensionObjects" yaml:"extensionObjects"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.15.6/docs/resources/extension#extension_schema Extension#extension_schema}.
	ExtensionSchema *string `field:"required" json:"extensionSchema" yaml:"extensionSchema"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.15.6/docs/resources/extension#config Extension#config}.
	Config *string `field:"optional" json:"config" yaml:"config"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.15.6/docs/resources/extension#endpoint_url Extension#endpoint_url}.
	EndpointUrl *string `field:"optional" json:"endpointUrl" yaml:"endpointUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.15.6/docs/resources/extension#name Extension#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.15.6/docs/resources/extension#type Extension#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

