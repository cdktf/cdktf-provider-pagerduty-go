// Prebuilt pagerduty Provider for Terraform CDK (cdktf)
package pagerduty


type ServiceSupportHours struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/pagerduty/r/service#days_of_week Service#days_of_week}.
	DaysOfWeek *[]*float64 `field:"optional" json:"daysOfWeek" yaml:"daysOfWeek"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/pagerduty/r/service#end_time Service#end_time}.
	EndTime *string `field:"optional" json:"endTime" yaml:"endTime"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/pagerduty/r/service#start_time Service#start_time}.
	StartTime *string `field:"optional" json:"startTime" yaml:"startTime"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/pagerduty/r/service#time_zone Service#time_zone}.
	TimeZone *string `field:"optional" json:"timeZone" yaml:"timeZone"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/pagerduty/r/service#type Service#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

