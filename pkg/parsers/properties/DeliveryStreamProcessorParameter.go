package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// DeliveryStreamProcessorParameter Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-processorparameter.html
type DeliveryStreamProcessorParameter struct {
	ParameterName  interface{} `yaml:"ParameterName"`
	ParameterValue interface{} `yaml:"ParameterValue"`
}

// DeliveryStreamProcessorParameter validation
func (resource DeliveryStreamProcessorParameter) Validate() []error {
	errors := []error{}

	if resource.ParameterName == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'ParameterName'"))
	}
	if resource.ParameterValue == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'ParameterValue'"))
	}
	return errors
}
