package schedule


type ScheduleLayer struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/pagerduty/r/schedule#rotation_turn_length_seconds Schedule#rotation_turn_length_seconds}.
	RotationTurnLengthSeconds *float64 `field:"required" json:"rotationTurnLengthSeconds" yaml:"rotationTurnLengthSeconds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/pagerduty/r/schedule#rotation_virtual_start Schedule#rotation_virtual_start}.
	RotationVirtualStart *string `field:"required" json:"rotationVirtualStart" yaml:"rotationVirtualStart"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/pagerduty/r/schedule#start Schedule#start}.
	Start *string `field:"required" json:"start" yaml:"start"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/pagerduty/r/schedule#users Schedule#users}.
	Users *[]*string `field:"required" json:"users" yaml:"users"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/pagerduty/r/schedule#end Schedule#end}.
	End *string `field:"optional" json:"end" yaml:"end"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/pagerduty/r/schedule#name Schedule#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// restriction block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/pagerduty/r/schedule#restriction Schedule#restriction}
	Restriction interface{} `field:"optional" json:"restriction" yaml:"restriction"`
}

