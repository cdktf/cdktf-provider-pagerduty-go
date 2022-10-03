package slackconnection

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SlackConnectionConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/pagerduty/r/slack_connection#channel_id SlackConnection#channel_id}.
	ChannelId *string `field:"required" json:"channelId" yaml:"channelId"`
	// config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/pagerduty/r/slack_connection#config SlackConnection#config}
	Config interface{} `field:"required" json:"config" yaml:"config"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/pagerduty/r/slack_connection#notification_type SlackConnection#notification_type}.
	NotificationType *string `field:"required" json:"notificationType" yaml:"notificationType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/pagerduty/r/slack_connection#source_id SlackConnection#source_id}.
	SourceId *string `field:"required" json:"sourceId" yaml:"sourceId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/pagerduty/r/slack_connection#source_type SlackConnection#source_type}.
	SourceType *string `field:"required" json:"sourceType" yaml:"sourceType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/pagerduty/r/slack_connection#workspace_id SlackConnection#workspace_id}.
	WorkspaceId *string `field:"required" json:"workspaceId" yaml:"workspaceId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/pagerduty/r/slack_connection#id SlackConnection#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

