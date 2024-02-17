// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package serviceintegration

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ServiceIntegrationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.8.0/docs/resources/service_integration#service ServiceIntegration#service}.
	Service *string `field:"required" json:"service" yaml:"service"`
	// email_filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.8.0/docs/resources/service_integration#email_filter ServiceIntegration#email_filter}
	EmailFilter interface{} `field:"optional" json:"emailFilter" yaml:"emailFilter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.8.0/docs/resources/service_integration#email_filter_mode ServiceIntegration#email_filter_mode}.
	EmailFilterMode *string `field:"optional" json:"emailFilterMode" yaml:"emailFilterMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.8.0/docs/resources/service_integration#email_incident_creation ServiceIntegration#email_incident_creation}.
	EmailIncidentCreation *string `field:"optional" json:"emailIncidentCreation" yaml:"emailIncidentCreation"`
	// email_parser block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.8.0/docs/resources/service_integration#email_parser ServiceIntegration#email_parser}
	EmailParser interface{} `field:"optional" json:"emailParser" yaml:"emailParser"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.8.0/docs/resources/service_integration#email_parsing_fallback ServiceIntegration#email_parsing_fallback}.
	EmailParsingFallback *string `field:"optional" json:"emailParsingFallback" yaml:"emailParsingFallback"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.8.0/docs/resources/service_integration#id ServiceIntegration#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.8.0/docs/resources/service_integration#integration_email ServiceIntegration#integration_email}.
	IntegrationEmail *string `field:"optional" json:"integrationEmail" yaml:"integrationEmail"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.8.0/docs/resources/service_integration#integration_key ServiceIntegration#integration_key}.
	IntegrationKey *string `field:"optional" json:"integrationKey" yaml:"integrationKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.8.0/docs/resources/service_integration#name ServiceIntegration#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.8.0/docs/resources/service_integration#type ServiceIntegration#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.8.0/docs/resources/service_integration#vendor ServiceIntegration#vendor}.
	Vendor *string `field:"optional" json:"vendor" yaml:"vendor"`
}

