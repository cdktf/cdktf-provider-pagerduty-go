// Prebuilt pagerduty Provider for Terraform CDK (cdktf)
package pagerduty


type SlackConnectionConfigA struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/pagerduty/r/slack_connection#events SlackConnection#events}.
	Events *[]*string `field:"required" json:"events" yaml:"events"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/pagerduty/r/slack_connection#priorities SlackConnection#priorities}.
	Priorities *[]*string `field:"optional" json:"priorities" yaml:"priorities"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/pagerduty/r/slack_connection#urgency SlackConnection#urgency}.
	Urgency *string `field:"optional" json:"urgency" yaml:"urgency"`
}

