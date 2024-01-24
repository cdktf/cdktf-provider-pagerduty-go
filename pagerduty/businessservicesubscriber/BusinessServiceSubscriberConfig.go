// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package businessservicesubscriber

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BusinessServiceSubscriberConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.5.0/docs/resources/business_service_subscriber#business_service_id BusinessServiceSubscriber#business_service_id}.
	BusinessServiceId *string `field:"required" json:"businessServiceId" yaml:"businessServiceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.5.0/docs/resources/business_service_subscriber#subscriber_id BusinessServiceSubscriber#subscriber_id}.
	SubscriberId *string `field:"required" json:"subscriberId" yaml:"subscriberId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.5.0/docs/resources/business_service_subscriber#subscriber_type BusinessServiceSubscriber#subscriber_type}.
	SubscriberType *string `field:"required" json:"subscriberType" yaml:"subscriberType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.5.0/docs/resources/business_service_subscriber#id BusinessServiceSubscriber#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

