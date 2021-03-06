package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// DeliveryStreamKinesisStreamSourceConfiguration Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-kinesisstreamsourceconfiguration.html
type DeliveryStreamKinesisStreamSourceConfiguration struct {
	KinesisStreamARN interface{} `yaml:"KinesisStreamARN"`
	RoleARN          interface{} `yaml:"RoleARN"`
}

// DeliveryStreamKinesisStreamSourceConfiguration validation
func (resource DeliveryStreamKinesisStreamSourceConfiguration) Validate() []error {
	errors := []error{}

	if resource.KinesisStreamARN == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'KinesisStreamARN'"))
	}
	if resource.RoleARN == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'RoleARN'"))
	}
	return errors
}
