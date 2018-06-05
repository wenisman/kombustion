package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

type TopicRule_S3Action struct {
	BucketName interface{} `yaml:"BucketName"`
	Key        interface{} `yaml:"Key"`
	RoleArn    interface{} `yaml:"RoleArn"`
}

func (resource TopicRule_S3Action) Validate() []error {
	errs := []error{}

	if resource.BucketName == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'BucketName'"))
	}
	if resource.Key == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Key'"))
	}
	if resource.RoleArn == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'RoleArn'"))
	}
	return errs
}