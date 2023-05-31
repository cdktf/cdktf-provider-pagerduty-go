package datapagerdutylicense

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataPagerdutyLicenseConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/2.15.0/docs/data-sources/license#allocations_available DataPagerdutyLicense#allocations_available}.
	AllocationsAvailable *float64 `field:"optional" json:"allocationsAvailable" yaml:"allocationsAvailable"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/2.15.0/docs/data-sources/license#current_value DataPagerdutyLicense#current_value}.
	CurrentValue *float64 `field:"optional" json:"currentValue" yaml:"currentValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/2.15.0/docs/data-sources/license#description DataPagerdutyLicense#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/2.15.0/docs/data-sources/license#html_url DataPagerdutyLicense#html_url}.
	HtmlUrl *string `field:"optional" json:"htmlUrl" yaml:"htmlUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/2.15.0/docs/data-sources/license#id DataPagerdutyLicense#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/2.15.0/docs/data-sources/license#name DataPagerdutyLicense#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/2.15.0/docs/data-sources/license#role_group DataPagerdutyLicense#role_group}.
	RoleGroup *string `field:"optional" json:"roleGroup" yaml:"roleGroup"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/2.15.0/docs/data-sources/license#self DataPagerdutyLicense#self}.
	SelfAttribute *string `field:"optional" json:"selfAttribute" yaml:"selfAttribute"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/2.15.0/docs/data-sources/license#summary DataPagerdutyLicense#summary}.
	Summary *string `field:"optional" json:"summary" yaml:"summary"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/2.15.0/docs/data-sources/license#type DataPagerdutyLicense#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/2.15.0/docs/data-sources/license#valid_roles DataPagerdutyLicense#valid_roles}.
	ValidRoles *[]*string `field:"optional" json:"validRoles" yaml:"validRoles"`
}

