// Prebuilt pagerduty Provider for Terraform CDK (cdktf)
package pagerduty

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BusinessServiceConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/pagerduty/r/business_service#name BusinessService#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/pagerduty/r/business_service#description BusinessService#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/pagerduty/r/business_service#id BusinessService#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/pagerduty/r/business_service#point_of_contact BusinessService#point_of_contact}.
	PointOfContact *string `field:"optional" json:"pointOfContact" yaml:"pointOfContact"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/pagerduty/r/business_service#team BusinessService#team}.
	Team *string `field:"optional" json:"team" yaml:"team"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/pagerduty/r/business_service#type BusinessService#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

