// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datapagerdutyserviceintegration

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataPagerdutyServiceIntegrationConfig struct {
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
	// examples "Amazon CloudWatch", "New Relic".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.28.0/docs/data-sources/service_integration#integration_summary DataPagerdutyServiceIntegration#integration_summary}
	IntegrationSummary *string `field:"required" json:"integrationSummary" yaml:"integrationSummary"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.28.0/docs/data-sources/service_integration#service_name DataPagerdutyServiceIntegration#service_name}.
	ServiceName *string `field:"required" json:"serviceName" yaml:"serviceName"`
}

