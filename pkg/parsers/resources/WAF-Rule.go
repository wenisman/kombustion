package resources

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
)

// WAFRule Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-waf-rule.html
type WAFRule struct {
	Type       string            `yaml:"Type"`
	Properties WAFRuleProperties `yaml:"Properties"`
	Condition  interface{}       `yaml:"Condition,omitempty"`
	Metadata   interface{}       `yaml:"Metadata,omitempty"`
	DependsOn  interface{}       `yaml:"DependsOn,omitempty"`
}

// WAFRule Properties
type WAFRuleProperties struct {
	MetricName interface{} `yaml:"MetricName"`
	Name       interface{} `yaml:"Name"`
	Predicates interface{} `yaml:"Predicates,omitempty"`
}

// NewWAFRule constructor creates a new WAFRule
func NewWAFRule(properties WAFRuleProperties, deps ...interface{}) WAFRule {
	return WAFRule{
		Type:       "AWS::WAF::Rule",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseWAFRule parses WAFRule
func ParseWAFRule(
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
	var resource WAFRule
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
					"Fn::Sub": "${AWS::StackName}-WAFRule-" + name,
				},
			},
		},
	}

	return
}

// ParseWAFRule validator
func (resource WAFRule) Validate() []error {
	return resource.Properties.Validate()
}

// ParseWAFRuleProperties validator
func (resource WAFRuleProperties) Validate() []error {
	errors := []error{}
	if resource.MetricName == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'MetricName'"))
	}
	if resource.Name == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'Name'"))
	}
	return errors
}
