// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package servicecustomfield

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ServiceCustomFieldConfig struct {
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
	// The kind of data the custom field is allowed to contain.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.26.0/docs/resources/service_custom_field#data_type ServiceCustomField#data_type}
	DataType *string `field:"required" json:"dataType" yaml:"dataType"`
	// The human-readable name of the field. This must be unique across an account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.26.0/docs/resources/service_custom_field#display_name ServiceCustomField#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// The type of data this field contains. In combination with the `data_type` field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.26.0/docs/resources/service_custom_field#field_type ServiceCustomField#field_type}
	FieldType *string `field:"required" json:"fieldType" yaml:"fieldType"`
	// The name of the field.
	//
	// May include ASCII characters, specifically lowercase letters, digits, and underescores. The `name` for a Field must be unique and cannot be changed once created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.26.0/docs/resources/service_custom_field#name ServiceCustomField#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Default value for the field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.26.0/docs/resources/service_custom_field#default_value ServiceCustomField#default_value}
	DefaultValue *string `field:"optional" json:"defaultValue" yaml:"defaultValue"`
	// A description of the data this field contains.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.26.0/docs/resources/service_custom_field#description ServiceCustomField#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Whether the field is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.26.0/docs/resources/service_custom_field#enabled ServiceCustomField#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// field_option block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.26.0/docs/resources/service_custom_field#field_option ServiceCustomField#field_option}
	FieldOption interface{} `field:"optional" json:"fieldOption" yaml:"fieldOption"`
}

