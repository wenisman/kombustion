package resources

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
)

// EC2VPCPeeringConnection Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcpeeringconnection.html
type EC2VPCPeeringConnection struct {
	Type       string                            `yaml:"Type"`
	Properties EC2VPCPeeringConnectionProperties `yaml:"Properties"`
	Condition  interface{}                       `yaml:"Condition,omitempty"`
	Metadata   interface{}                       `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                       `yaml:"DependsOn,omitempty"`
}

// EC2VPCPeeringConnection Properties
type EC2VPCPeeringConnectionProperties struct {
	PeerOwnerId interface{} `yaml:"PeerOwnerId,omitempty"`
	PeerRoleArn interface{} `yaml:"PeerRoleArn,omitempty"`
	PeerVpcId   interface{} `yaml:"PeerVpcId"`
	VpcId       interface{} `yaml:"VpcId"`
	Tags        interface{} `yaml:"Tags,omitempty"`
}

// NewEC2VPCPeeringConnection constructor creates a new EC2VPCPeeringConnection
func NewEC2VPCPeeringConnection(properties EC2VPCPeeringConnectionProperties, deps ...interface{}) EC2VPCPeeringConnection {
	return EC2VPCPeeringConnection{
		Type:       "AWS::EC2::VPCPeeringConnection",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseEC2VPCPeeringConnection parses EC2VPCPeeringConnection
func ParseEC2VPCPeeringConnection(
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
	var resource EC2VPCPeeringConnection
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
					"Fn::Sub": "${AWS::StackName}-EC2VPCPeeringConnection-" + name,
				},
			},
		},
	}

	return
}

// ParseEC2VPCPeeringConnection validator
func (resource EC2VPCPeeringConnection) Validate() []error {
	return resource.Properties.Validate()
}

// ParseEC2VPCPeeringConnectionProperties validator
func (resource EC2VPCPeeringConnectionProperties) Validate() []error {
	errors := []error{}
	if resource.PeerVpcId == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'PeerVpcId'"))
	}
	if resource.VpcId == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'VpcId'"))
	}
	return errors
}
