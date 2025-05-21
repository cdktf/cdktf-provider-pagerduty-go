// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package incidenttypecustomfield

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IncidentTypeCustomFieldConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.25.1/docs/resources/incident_type_custom_field#data_type IncidentTypeCustomField#data_type}.
	DataType *string `field:"required" json:"dataType" yaml:"dataType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.25.1/docs/resources/incident_type_custom_field#display_name IncidentTypeCustomField#display_name}.
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.25.1/docs/resources/incident_type_custom_field#incident_type IncidentTypeCustomField#incident_type}.
	IncidentType *string `field:"required" json:"incidentType" yaml:"incidentType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.25.1/docs/resources/incident_type_custom_field#name IncidentTypeCustomField#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.25.1/docs/resources/incident_type_custom_field#default_value IncidentTypeCustomField#default_value}.
	DefaultValue *string `field:"optional" json:"defaultValue" yaml:"defaultValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.25.1/docs/resources/incident_type_custom_field#description IncidentTypeCustomField#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.25.1/docs/resources/incident_type_custom_field#enabled IncidentTypeCustomField#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.25.1/docs/resources/incident_type_custom_field#field_options IncidentTypeCustomField#field_options}.
	FieldOptions *[]*string `field:"optional" json:"fieldOptions" yaml:"fieldOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.25.1/docs/resources/incident_type_custom_field#field_type IncidentTypeCustomField#field_type}.
	FieldType *string `field:"optional" json:"fieldType" yaml:"fieldType"`
}

