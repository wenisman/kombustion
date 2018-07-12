package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// PipelineActionDeclaration Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-stages-actions.html
type PipelineActionDeclaration struct {
	Configuration   interface{}           `yaml:"Configuration,omitempty"`
	Name            interface{}           `yaml:"Name"`
	RoleArn         interface{}           `yaml:"RoleArn,omitempty"`
	RunOrder        interface{}           `yaml:"RunOrder,omitempty"`
	InputArtifacts  interface{}           `yaml:"InputArtifacts,omitempty"`
	OutputArtifacts interface{}           `yaml:"OutputArtifacts,omitempty"`
	ActionTypeId    *PipelineActionTypeId `yaml:"ActionTypeId"`
}

// PipelineActionDeclaration validation
func (resource PipelineActionDeclaration) Validate() []error {
	errs := []error{}

	if resource.Name == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Name'"))
	}
	if resource.ActionTypeId == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'ActionTypeId'"))
	} else {
		errs = append(errs, resource.ActionTypeId.Validate()...)
	}
	return errs
}