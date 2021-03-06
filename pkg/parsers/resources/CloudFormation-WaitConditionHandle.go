package resources

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
)

// CloudFormationWaitConditionHandle Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waitconditionhandle.html
type CloudFormationWaitConditionHandle struct {
	Type       string                                      `yaml:"Type"`
	Properties CloudFormationWaitConditionHandleProperties `yaml:"Properties"`
	Condition  interface{}                                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                                 `yaml:"DependsOn,omitempty"`
}

// CloudFormationWaitConditionHandle Properties
type CloudFormationWaitConditionHandleProperties struct {
}

// NewCloudFormationWaitConditionHandle constructor creates a new CloudFormationWaitConditionHandle
func NewCloudFormationWaitConditionHandle(properties CloudFormationWaitConditionHandleProperties, deps ...interface{}) CloudFormationWaitConditionHandle {
	return CloudFormationWaitConditionHandle{
		Type:       "AWS::CloudFormation::WaitConditionHandle",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseCloudFormationWaitConditionHandle parses CloudFormationWaitConditionHandle
func ParseCloudFormationWaitConditionHandle(
	name string,
	data string,
) (
	source string,
	conditions types.TemplateObject,
	metadata types.TemplateObject,
	mappings types.TemplateObject,
	outputs types.TemplateObject,
	parameters types.TemplateObject,
	resources types.TemplateObject,
	transform types.TemplateObject,
	errors []error,
) {
	source = "kombustion-core-resources"

	// Resources
	var resource CloudFormationWaitConditionHandle
	err := yaml.Unmarshal([]byte(data), &resource)

	if err != nil {
		errors = append(errors, err)
		return
	}

	if validateErrs := resource.Properties.Validate(); len(errors) > 0 {
		errors = append(errors, validateErrs...)
		return
	}

	resources = types.TemplateObject{name: resource}

	// Outputs

	outputs = types.TemplateObject{
		name: types.TemplateObject{
			"Description": name + " Object",
			"Value": map[string]interface{}{
				"Ref": name,
			},
			"Export": map[string]interface{}{
				"Name": map[string]interface{}{
					"Fn::Sub": "${AWS::StackName}-CloudFormationWaitConditionHandle-" + name,
				},
			},
		},
	}

	return
}

// ParseCloudFormationWaitConditionHandle validator
func (resource CloudFormationWaitConditionHandle) Validate() []error {
	return resource.Properties.Validate()
}

// ParseCloudFormationWaitConditionHandleProperties validator
func (resource CloudFormationWaitConditionHandleProperties) Validate() []error {
	errors := []error{}
	return errors
}
