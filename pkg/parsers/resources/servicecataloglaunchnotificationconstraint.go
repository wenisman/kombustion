package resources

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
	"log"
)

// ServiceCatalogLaunchNotificationConstraint Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-launchnotificationconstraint.html
type ServiceCatalogLaunchNotificationConstraint struct {
	Type       string                                               `yaml:"Type"`
	Properties ServiceCatalogLaunchNotificationConstraintProperties `yaml:"Properties"`
	Condition  interface{}                                          `yaml:"Condition,omitempty"`
	Metadata   interface{}                                          `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                                          `yaml:"DependsOn,omitempty"`
}

// ServiceCatalogLaunchNotificationConstraint Properties
type ServiceCatalogLaunchNotificationConstraintProperties struct {
	AcceptLanguage   interface{} `yaml:"AcceptLanguage,omitempty"`
	Description      interface{} `yaml:"Description,omitempty"`
	PortfolioId      interface{} `yaml:"PortfolioId"`
	ProductId        interface{} `yaml:"ProductId"`
	NotificationArns interface{} `yaml:"NotificationArns"`
}

// NewServiceCatalogLaunchNotificationConstraint constructor creates a new ServiceCatalogLaunchNotificationConstraint
func NewServiceCatalogLaunchNotificationConstraint(properties ServiceCatalogLaunchNotificationConstraintProperties, deps ...interface{}) ServiceCatalogLaunchNotificationConstraint {
	return ServiceCatalogLaunchNotificationConstraint{
		Type:       "AWS::ServiceCatalog::LaunchNotificationConstraint",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseServiceCatalogLaunchNotificationConstraint parses ServiceCatalogLaunchNotificationConstraint
func ParseServiceCatalogLaunchNotificationConstraint(name string, data string) (cf types.TemplateObject, err error) {
	var resource ServiceCatalogLaunchNotificationConstraint
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: ServiceCatalogLaunchNotificationConstraint - ", err)
		}
		return
	}
	cf = types.TemplateObject{name: resource}
	return
}

// ParseServiceCatalogLaunchNotificationConstraint validator
func (resource ServiceCatalogLaunchNotificationConstraint) Validate() []error {
	return resource.Properties.Validate()
}

// ParseServiceCatalogLaunchNotificationConstraintProperties validator
func (resource ServiceCatalogLaunchNotificationConstraintProperties) Validate() []error {
	errs := []error{}
	if resource.PortfolioId == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'PortfolioId'"))
	}
	if resource.ProductId == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'ProductId'"))
	}
	if resource.NotificationArns == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'NotificationArns'"))
	}
	return errs
}
