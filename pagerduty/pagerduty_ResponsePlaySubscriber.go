// Prebuilt pagerduty Provider for Terraform CDK (cdktf)
package pagerduty


type ResponsePlaySubscriber struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/pagerduty/r/response_play#id ResponsePlay#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/pagerduty/r/response_play#type ResponsePlay#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}
