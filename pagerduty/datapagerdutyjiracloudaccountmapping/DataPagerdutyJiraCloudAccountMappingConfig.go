// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datapagerdutyjiracloudaccountmapping

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataPagerdutyJiraCloudAccountMappingConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.19.4/docs/data-sources/jira_cloud_account_mapping#subdomain DataPagerdutyJiraCloudAccountMapping#subdomain}.
	Subdomain *string `field:"required" json:"subdomain" yaml:"subdomain"`
}

