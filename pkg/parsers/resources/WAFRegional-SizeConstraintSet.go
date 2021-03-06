package resources

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
)

// WAFRegionalSizeConstraintSet Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-sizeconstraintset.html
type WAFRegionalSizeConstraintSet struct {
	Type       string                                 `yaml:"Type"`
	Properties WAFRegionalSizeConstraintSetProperties `yaml:"Properties"`
	Condition  interface{}                            `yaml:"Condition,omitempty"`
	Metadata   interface{}                            `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                            `yaml:"DependsOn,omitempty"`
}

// WAFRegionalSizeConstraintSet Properties
type WAFRegionalSizeConstraintSetProperties struct {
	Name            interface{} `yaml:"Name"`
	SizeConstraints interface{} `yaml:"SizeConstraints,omitempty"`
}

// NewWAFRegionalSizeConstraintSet constructor creates a new WAFRegionalSizeConstraintSet
func NewWAFRegionalSizeConstraintSet(properties WAFRegionalSizeConstraintSetProperties, deps ...interface{}) WAFRegionalSizeConstraintSet {
	return WAFRegionalSizeConstraintSet{
		Type:       "AWS::WAFRegional::SizeConstraintSet",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseWAFRegionalSizeConstraintSet parses WAFRegionalSizeConstraintSet
func ParseWAFRegionalSizeConstraintSet(
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
	var resource WAFRegionalSizeConstraintSet
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
					"Fn::Sub": "${AWS::StackName}-WAFRegionalSizeConstraintSet-" + name,
				},
			},
		},
	}

	return
}

// ParseWAFRegionalSizeConstraintSet validator
func (resource WAFRegionalSizeConstraintSet) Validate() []error {
	return resource.Properties.Validate()
}

// ParseWAFRegionalSizeConstraintSetProperties validator
func (resource WAFRegionalSizeConstraintSetProperties) Validate() []error {
	errors := []error{}
	if resource.Name == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'Name'"))
	}
	return errors
}
