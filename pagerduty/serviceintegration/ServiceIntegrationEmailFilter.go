// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package serviceintegration


type ServiceIntegrationEmailFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.2.2/docs/resources/service_integration#body_mode ServiceIntegration#body_mode}.
	BodyMode *string `field:"optional" json:"bodyMode" yaml:"bodyMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.2.2/docs/resources/service_integration#body_regex ServiceIntegration#body_regex}.
	BodyRegex *string `field:"optional" json:"bodyRegex" yaml:"bodyRegex"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.2.2/docs/resources/service_integration#from_email_mode ServiceIntegration#from_email_mode}.
	FromEmailMode *string `field:"optional" json:"fromEmailMode" yaml:"fromEmailMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.2.2/docs/resources/service_integration#from_email_regex ServiceIntegration#from_email_regex}.
	FromEmailRegex *string `field:"optional" json:"fromEmailRegex" yaml:"fromEmailRegex"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.2.2/docs/resources/service_integration#subject_mode ServiceIntegration#subject_mode}.
	SubjectMode *string `field:"optional" json:"subjectMode" yaml:"subjectMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.2.2/docs/resources/service_integration#subject_regex ServiceIntegration#subject_regex}.
	SubjectRegex *string `field:"optional" json:"subjectRegex" yaml:"subjectRegex"`
}

