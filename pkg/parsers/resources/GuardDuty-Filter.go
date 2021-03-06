package resources

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/pkg/parsers/properties"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
)

// GuardDutyFilter Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-filter.html
type GuardDutyFilter struct {
	Type       string                    `yaml:"Type"`
	Properties GuardDutyFilterProperties `yaml:"Properties"`
	Condition  interface{}               `yaml:"Condition,omitempty"`
	Metadata   interface{}               `yaml:"Metadata,omitempty"`
	DependsOn  interface{}               `yaml:"DependsOn,omitempty"`
}

// GuardDutyFilter Properties
type GuardDutyFilterProperties struct {
	Action          interface{}                       `yaml:"Action"`
	Description     interface{}                       `yaml:"Description"`
	DetectorId      interface{}                       `yaml:"DetectorId"`
	Name            interface{}                       `yaml:"Name,omitempty"`
	Rank            interface{}                       `yaml:"Rank"`
	FindingCriteria *properties.FilterFindingCriteria `yaml:"FindingCriteria"`
}

// NewGuardDutyFilter constructor creates a new GuardDutyFilter
func NewGuardDutyFilter(properties GuardDutyFilterProperties, deps ...interface{}) GuardDutyFilter {
	return GuardDutyFilter{
		Type:       "AWS::GuardDuty::Filter",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseGuardDutyFilter parses GuardDutyFilter
func ParseGuardDutyFilter(
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
	var resource GuardDutyFilter
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
					"Fn::Sub": "${AWS::StackName}-GuardDutyFilter-" + name,
				},
			},
		},
	}

	return
}

// ParseGuardDutyFilter validator
func (resource GuardDutyFilter) Validate() []error {
	return resource.Properties.Validate()
}

// ParseGuardDutyFilterProperties validator
func (resource GuardDutyFilterProperties) Validate() []error {
	errors := []error{}
	if resource.Action == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'Action'"))
	}
	if resource.Description == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'Description'"))
	}
	if resource.DetectorId == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'DetectorId'"))
	}
	if resource.Rank == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'Rank'"))
	}
	if resource.FindingCriteria == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'FindingCriteria'"))
	} else {
		errors = append(errors, resource.FindingCriteria.Validate()...)
	}
	return errors
}
