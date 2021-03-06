package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// DeploymentGroupAlarmConfiguration Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-alarmconfiguration.html
type DeploymentGroupAlarmConfiguration struct {
	Enabled                interface{} `yaml:"Enabled,omitempty"`
	IgnorePollAlarmFailure interface{} `yaml:"IgnorePollAlarmFailure,omitempty"`
	Alarms                 interface{} `yaml:"Alarms,omitempty"`
}

// DeploymentGroupAlarmConfiguration validation
func (resource DeploymentGroupAlarmConfiguration) Validate() []error {
	errors := []error{}

	return errors
}
