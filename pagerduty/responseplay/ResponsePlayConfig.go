// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package responseplay

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ResponsePlayConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.29.1/docs/resources/response_play#from ResponsePlay#from}.
	From *string `field:"required" json:"from" yaml:"from"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.29.1/docs/resources/response_play#name ResponsePlay#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.29.1/docs/resources/response_play#conference_number ResponsePlay#conference_number}.
	ConferenceNumber *string `field:"optional" json:"conferenceNumber" yaml:"conferenceNumber"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.29.1/docs/resources/response_play#conference_url ResponsePlay#conference_url}.
	ConferenceUrl *string `field:"optional" json:"conferenceUrl" yaml:"conferenceUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.29.1/docs/resources/response_play#description ResponsePlay#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.29.1/docs/resources/response_play#id ResponsePlay#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// responder block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.29.1/docs/resources/response_play#responder ResponsePlay#responder}
	Responder interface{} `field:"optional" json:"responder" yaml:"responder"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.29.1/docs/resources/response_play#responders_message ResponsePlay#responders_message}.
	RespondersMessage *string `field:"optional" json:"respondersMessage" yaml:"respondersMessage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.29.1/docs/resources/response_play#runnability ResponsePlay#runnability}.
	Runnability *string `field:"optional" json:"runnability" yaml:"runnability"`
	// subscriber block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.29.1/docs/resources/response_play#subscriber ResponsePlay#subscriber}
	Subscriber interface{} `field:"optional" json:"subscriber" yaml:"subscriber"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.29.1/docs/resources/response_play#subscribers_message ResponsePlay#subscribers_message}.
	SubscribersMessage *string `field:"optional" json:"subscribersMessage" yaml:"subscribersMessage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.29.1/docs/resources/response_play#team ResponsePlay#team}.
	Team *string `field:"optional" json:"team" yaml:"team"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.29.1/docs/resources/response_play#type ResponsePlay#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

