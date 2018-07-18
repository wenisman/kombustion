package resources

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
	"log"
)

// NeptuneDBInstance Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-neptune-dbinstance.html
type NeptuneDBInstance struct {
	Type       string                      `yaml:"Type"`
	Properties NeptuneDBInstanceProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

// NeptuneDBInstance Properties
type NeptuneDBInstanceProperties struct {
	AllowMajorVersionUpgrade   interface{} `yaml:"AllowMajorVersionUpgrade,omitempty"`
	AutoMinorVersionUpgrade    interface{} `yaml:"AutoMinorVersionUpgrade,omitempty"`
	AvailabilityZone           interface{} `yaml:"AvailabilityZone,omitempty"`
	DBClusterIdentifier        interface{} `yaml:"DBClusterIdentifier,omitempty"`
	DBInstanceClass            interface{} `yaml:"DBInstanceClass"`
	DBInstanceIdentifier       interface{} `yaml:"DBInstanceIdentifier,omitempty"`
	DBParameterGroupName       interface{} `yaml:"DBParameterGroupName,omitempty"`
	DBSnapshotIdentifier       interface{} `yaml:"DBSnapshotIdentifier,omitempty"`
	DBSubnetGroupName          interface{} `yaml:"DBSubnetGroupName,omitempty"`
	PreferredMaintenanceWindow interface{} `yaml:"PreferredMaintenanceWindow,omitempty"`
	Tags                       interface{} `yaml:"Tags,omitempty"`
}

// NewNeptuneDBInstance constructor creates a new NeptuneDBInstance
func NewNeptuneDBInstance(properties NeptuneDBInstanceProperties, deps ...interface{}) NeptuneDBInstance {
	return NeptuneDBInstance{
		Type:       "AWS::Neptune::DBInstance",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseNeptuneDBInstance parses NeptuneDBInstance
func ParseNeptuneDBInstance(name string, data string) (cf types.TemplateObject, err error) {
	var resource NeptuneDBInstance
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: NeptuneDBInstance - ", err)
		}
		return
	}
	cf = types.TemplateObject{name: resource}
	return
}

// ParseNeptuneDBInstance validator
func (resource NeptuneDBInstance) Validate() []error {
	return resource.Properties.Validate()
}

// ParseNeptuneDBInstanceProperties validator
func (resource NeptuneDBInstanceProperties) Validate() []error {
	errs := []error{}
	if resource.DBInstanceClass == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'DBInstanceClass'"))
	}
	return errs
}
