package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// LoadBalancerHealthCheck Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb-health-check.html
type LoadBalancerHealthCheck struct {
	HealthyThreshold   interface{} `yaml:"HealthyThreshold"`
	Interval           interface{} `yaml:"Interval"`
	Target             interface{} `yaml:"Target"`
	Timeout            interface{} `yaml:"Timeout"`
	UnhealthyThreshold interface{} `yaml:"UnhealthyThreshold"`
}

// LoadBalancerHealthCheck validation
func (resource LoadBalancerHealthCheck) Validate() []error {
	errs := []error{}

	if resource.HealthyThreshold == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'HealthyThreshold'"))
	}
	if resource.Interval == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Interval'"))
	}
	if resource.Target == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Target'"))
	}
	if resource.Timeout == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Timeout'"))
	}
	if resource.UnhealthyThreshold == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'UnhealthyThreshold'"))
	}
	return errs
}