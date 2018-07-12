package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// ClusterScriptBootstrapActionConfig Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-scriptbootstrapactionconfig.html
type ClusterScriptBootstrapActionConfig struct {
	Path interface{} `yaml:"Path"`
	Args interface{} `yaml:"Args,omitempty"`
}

// ClusterScriptBootstrapActionConfig validation
func (resource ClusterScriptBootstrapActionConfig) Validate() []error {
	errs := []error{}

	if resource.Path == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Path'"))
	}
	return errs
}