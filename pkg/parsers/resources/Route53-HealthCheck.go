package resources

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/pkg/parsers/properties"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
)

// Route53HealthCheck Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-healthcheck.html
type Route53HealthCheck struct {
	Type       string                       `yaml:"Type"`
	Properties Route53HealthCheckProperties `yaml:"Properties"`
	Condition  interface{}                  `yaml:"Condition,omitempty"`
	Metadata   interface{}                  `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                  `yaml:"DependsOn,omitempty"`
}

// Route53HealthCheck Properties
type Route53HealthCheckProperties struct {
	HealthCheckTags   interface{}                              `yaml:"HealthCheckTags,omitempty"`
	HealthCheckConfig *properties.HealthCheckHealthCheckConfig `yaml:"HealthCheckConfig"`
}

// NewRoute53HealthCheck constructor creates a new Route53HealthCheck
func NewRoute53HealthCheck(properties Route53HealthCheckProperties, deps ...interface{}) Route53HealthCheck {
	return Route53HealthCheck{
		Type:       "AWS::Route53::HealthCheck",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseRoute53HealthCheck parses Route53HealthCheck
func ParseRoute53HealthCheck(
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
	var resource Route53HealthCheck
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
					"Fn::Sub": "${AWS::StackName}-Route53HealthCheck-" + name,
				},
			},
		},
	}

	return
}

// ParseRoute53HealthCheck validator
func (resource Route53HealthCheck) Validate() []error {
	return resource.Properties.Validate()
}

// ParseRoute53HealthCheckProperties validator
func (resource Route53HealthCheckProperties) Validate() []error {
	errors := []error{}
	if resource.HealthCheckConfig == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'HealthCheckConfig'"))
	} else {
		errors = append(errors, resource.HealthCheckConfig.Validate()...)
	}
	return errors
}
