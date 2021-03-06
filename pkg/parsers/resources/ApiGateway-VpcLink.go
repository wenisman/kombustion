package resources

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
)

// ApiGatewayVpcLink Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-vpclink.html
type ApiGatewayVpcLink struct {
	Type       string                      `yaml:"Type"`
	Properties ApiGatewayVpcLinkProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

// ApiGatewayVpcLink Properties
type ApiGatewayVpcLinkProperties struct {
	Description interface{} `yaml:"Description,omitempty"`
	Name        interface{} `yaml:"Name"`
	TargetArns  interface{} `yaml:"TargetArns"`
}

// NewApiGatewayVpcLink constructor creates a new ApiGatewayVpcLink
func NewApiGatewayVpcLink(properties ApiGatewayVpcLinkProperties, deps ...interface{}) ApiGatewayVpcLink {
	return ApiGatewayVpcLink{
		Type:       "AWS::ApiGateway::VpcLink",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseApiGatewayVpcLink parses ApiGatewayVpcLink
func ParseApiGatewayVpcLink(
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
	var resource ApiGatewayVpcLink
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
					"Fn::Sub": "${AWS::StackName}-ApiGatewayVpcLink-" + name,
				},
			},
		},
	}

	return
}

// ParseApiGatewayVpcLink validator
func (resource ApiGatewayVpcLink) Validate() []error {
	return resource.Properties.Validate()
}

// ParseApiGatewayVpcLinkProperties validator
func (resource ApiGatewayVpcLinkProperties) Validate() []error {
	errors := []error{}
	if resource.Name == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'Name'"))
	}
	if resource.TargetArns == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'TargetArns'"))
	}
	return errors
}
