package resources

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
)

// CertificateManagerCertificate Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-certificatemanager-certificate.html
type CertificateManagerCertificate struct {
	Type       string                                  `yaml:"Type"`
	Properties CertificateManagerCertificateProperties `yaml:"Properties"`
	Condition  interface{}                             `yaml:"Condition,omitempty"`
	Metadata   interface{}                             `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                             `yaml:"DependsOn,omitempty"`
}

// CertificateManagerCertificate Properties
type CertificateManagerCertificateProperties struct {
	DomainName              interface{} `yaml:"DomainName"`
	ValidationMethod        interface{} `yaml:"ValidationMethod,omitempty"`
	DomainValidationOptions interface{} `yaml:"DomainValidationOptions,omitempty"`
	SubjectAlternativeNames interface{} `yaml:"SubjectAlternativeNames,omitempty"`
	Tags                    interface{} `yaml:"Tags,omitempty"`
}

// NewCertificateManagerCertificate constructor creates a new CertificateManagerCertificate
func NewCertificateManagerCertificate(properties CertificateManagerCertificateProperties, deps ...interface{}) CertificateManagerCertificate {
	return CertificateManagerCertificate{
		Type:       "AWS::CertificateManager::Certificate",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseCertificateManagerCertificate parses CertificateManagerCertificate
func ParseCertificateManagerCertificate(
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
	var resource CertificateManagerCertificate
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

	return
}

// ParseCertificateManagerCertificate validator
func (resource CertificateManagerCertificate) Validate() []error {
	return resource.Properties.Validate()
}

// ParseCertificateManagerCertificateProperties validator
func (resource CertificateManagerCertificateProperties) Validate() []error {
	errors := []error{}
	if resource.DomainName == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'DomainName'"))
	}
	return errors
}