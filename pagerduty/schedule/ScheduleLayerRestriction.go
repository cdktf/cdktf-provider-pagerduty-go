package schedule


type ScheduleLayerRestriction struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/pagerduty/r/schedule#duration_seconds Schedule#duration_seconds}.
	DurationSeconds *float64 `field:"required" json:"durationSeconds" yaml:"durationSeconds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/pagerduty/r/schedule#start_time_of_day Schedule#start_time_of_day}.
	StartTimeOfDay *string `field:"required" json:"startTimeOfDay" yaml:"startTimeOfDay"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/pagerduty/r/schedule#type Schedule#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/pagerduty/r/schedule#start_day_of_week Schedule#start_day_of_week}.
	StartDayOfWeek *float64 `field:"optional" json:"startDayOfWeek" yaml:"startDayOfWeek"`
}

