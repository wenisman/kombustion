package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// ScalingPolicyTargetTrackingScalingPolicyConfiguration Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-targettrackingscalingpolicyconfiguration.html
type ScalingPolicyTargetTrackingScalingPolicyConfiguration struct {
	DisableScaleIn                interface{}                                 `yaml:"DisableScaleIn,omitempty"`
	ScaleInCooldown               interface{}                                 `yaml:"ScaleInCooldown,omitempty"`
	ScaleOutCooldown              interface{}                                 `yaml:"ScaleOutCooldown,omitempty"`
	TargetValue                   interface{}                                 `yaml:"TargetValue"`
	PredefinedMetricSpecification *ScalingPolicyPredefinedMetricSpecification `yaml:"PredefinedMetricSpecification,omitempty"`
	CustomizedMetricSpecification *ScalingPolicyCustomizedMetricSpecification `yaml:"CustomizedMetricSpecification,omitempty"`
}

// ScalingPolicyTargetTrackingScalingPolicyConfiguration validation
func (resource ScalingPolicyTargetTrackingScalingPolicyConfiguration) Validate() []error {
	errors := []error{}

	if resource.TargetValue == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'TargetValue'"))
	}
	return errors
}
