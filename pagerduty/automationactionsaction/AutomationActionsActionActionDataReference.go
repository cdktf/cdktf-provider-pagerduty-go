package automationactionsaction


type AutomationActionsActionActionDataReference struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/pagerduty/r/automation_actions_action#invocation_command AutomationActionsAction#invocation_command}.
	InvocationCommand *string `field:"optional" json:"invocationCommand" yaml:"invocationCommand"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/pagerduty/r/automation_actions_action#process_automation_job_arguments AutomationActionsAction#process_automation_job_arguments}.
	ProcessAutomationJobArguments *string `field:"optional" json:"processAutomationJobArguments" yaml:"processAutomationJobArguments"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/pagerduty/r/automation_actions_action#process_automation_job_id AutomationActionsAction#process_automation_job_id}.
	ProcessAutomationJobId *string `field:"optional" json:"processAutomationJobId" yaml:"processAutomationJobId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/pagerduty/r/automation_actions_action#process_automation_node_filter AutomationActionsAction#process_automation_node_filter}.
	ProcessAutomationNodeFilter *string `field:"optional" json:"processAutomationNodeFilter" yaml:"processAutomationNodeFilter"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/pagerduty/r/automation_actions_action#script AutomationActionsAction#script}.
	Script *string `field:"optional" json:"script" yaml:"script"`
}
