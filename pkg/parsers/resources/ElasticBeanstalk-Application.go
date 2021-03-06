package resources

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"github.com/KablamoOSS/kombustion/pkg/parsers/properties"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
)

// ElasticBeanstalkApplication Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk.html
type ElasticBeanstalkApplication struct {
	Type       string                                `yaml:"Type"`
	Properties ElasticBeanstalkApplicationProperties `yaml:"Properties"`
	Condition  interface{}                           `yaml:"Condition,omitempty"`
	Metadata   interface{}                           `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                           `yaml:"DependsOn,omitempty"`
}

// ElasticBeanstalkApplication Properties
type ElasticBeanstalkApplicationProperties struct {
	ApplicationName         interface{}                                               `yaml:"ApplicationName,omitempty"`
	Description             interface{}                                               `yaml:"Description,omitempty"`
	ResourceLifecycleConfig *properties.ApplicationApplicationResourceLifecycleConfig `yaml:"ResourceLifecycleConfig,omitempty"`
}

// NewElasticBeanstalkApplication constructor creates a new ElasticBeanstalkApplication
func NewElasticBeanstalkApplication(properties ElasticBeanstalkApplicationProperties, deps ...interface{}) ElasticBeanstalkApplication {
	return ElasticBeanstalkApplication{
		Type:       "AWS::ElasticBeanstalk::Application",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseElasticBeanstalkApplication parses ElasticBeanstalkApplication
func ParseElasticBeanstalkApplication(
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
	var resource ElasticBeanstalkApplication
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
					"Fn::Sub": "${AWS::StackName}-ElasticBeanstalkApplication-" + name,
				},
			},
		},
	}

	return
}

// ParseElasticBeanstalkApplication validator
func (resource ElasticBeanstalkApplication) Validate() []error {
	return resource.Properties.Validate()
}

// ParseElasticBeanstalkApplicationProperties validator
func (resource ElasticBeanstalkApplicationProperties) Validate() []error {
	errors := []error{}
	return errors
}
