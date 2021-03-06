package resources

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/pkg/parsers/properties"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
)

// CloudFrontDistribution Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-distribution.html
type CloudFrontDistribution struct {
	Type       string                           `yaml:"Type"`
	Properties CloudFrontDistributionProperties `yaml:"Properties"`
	Condition  interface{}                      `yaml:"Condition,omitempty"`
	Metadata   interface{}                      `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                      `yaml:"DependsOn,omitempty"`
}

// CloudFrontDistribution Properties
type CloudFrontDistributionProperties struct {
	Tags               interface{}                                `yaml:"Tags,omitempty"`
	DistributionConfig *properties.DistributionDistributionConfig `yaml:"DistributionConfig"`
}

// NewCloudFrontDistribution constructor creates a new CloudFrontDistribution
func NewCloudFrontDistribution(properties CloudFrontDistributionProperties, deps ...interface{}) CloudFrontDistribution {
	return CloudFrontDistribution{
		Type:       "AWS::CloudFront::Distribution",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseCloudFrontDistribution parses CloudFrontDistribution
func ParseCloudFrontDistribution(
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
	var resource CloudFrontDistribution
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
					"Fn::Sub": "${AWS::StackName}-CloudFrontDistribution-" + name,
				},
			},
		},
	}

	return
}

// ParseCloudFrontDistribution validator
func (resource CloudFrontDistribution) Validate() []error {
	return resource.Properties.Validate()
}

// ParseCloudFrontDistributionProperties validator
func (resource CloudFrontDistributionProperties) Validate() []error {
	errors := []error{}
	if resource.DistributionConfig == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'DistributionConfig'"))
	} else {
		errors = append(errors, resource.DistributionConfig.Validate()...)
	}
	return errors
}
