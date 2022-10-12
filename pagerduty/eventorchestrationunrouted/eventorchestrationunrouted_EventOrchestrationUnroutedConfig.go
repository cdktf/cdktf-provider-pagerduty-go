package eventorchestrationunrouted

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EventOrchestrationUnroutedConfig struct {
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
	// catch_all block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/pagerduty/r/event_orchestration_unrouted#catch_all EventOrchestrationUnrouted#catch_all}
	CatchAll *EventOrchestrationUnroutedCatchAll `field:"required" json:"catchAll" yaml:"catchAll"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/pagerduty/r/event_orchestration_unrouted#event_orchestration EventOrchestrationUnrouted#event_orchestration}.
	EventOrchestration *string `field:"required" json:"eventOrchestration" yaml:"eventOrchestration"`
	// set block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/pagerduty/r/event_orchestration_unrouted#set EventOrchestrationUnrouted#set}
	Set interface{} `field:"required" json:"set" yaml:"set"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/pagerduty/r/event_orchestration_unrouted#id EventOrchestrationUnrouted#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}
