package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// LaunchTemplateInstanceMarketOptions Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-launchtemplatedata-instancemarketoptions.html
type LaunchTemplateInstanceMarketOptions struct {
	MarketType  interface{}                `yaml:"MarketType,omitempty"`
	SpotOptions *LaunchTemplateSpotOptions `yaml:"SpotOptions,omitempty"`
}

// LaunchTemplateInstanceMarketOptions validation
func (resource LaunchTemplateInstanceMarketOptions) Validate() []error {
	errs := []error{}

	return errs
}