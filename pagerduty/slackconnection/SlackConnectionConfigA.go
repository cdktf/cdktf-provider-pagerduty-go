package slackconnection


type SlackConnectionConfigA struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/2.14.4/docs/resources/slack_connection#events SlackConnection#events}.
	Events *[]*string `field:"required" json:"events" yaml:"events"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/2.14.4/docs/resources/slack_connection#priorities SlackConnection#priorities}.
	Priorities *[]*string `field:"optional" json:"priorities" yaml:"priorities"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/2.14.4/docs/resources/slack_connection#urgency SlackConnection#urgency}.
	Urgency *string `field:"optional" json:"urgency" yaml:"urgency"`
}

