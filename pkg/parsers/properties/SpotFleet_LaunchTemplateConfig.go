package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

type SpotFleet_LaunchTemplateConfig struct {
	Overrides                   interface{}                                 `yaml:"Overrides,omitempty"`
	LaunchTemplateSpecification *SpotFleet_FleetLaunchTemplateSpecification `yaml:"LaunchTemplateSpecification,omitempty"`
}

func (resource SpotFleet_LaunchTemplateConfig) Validate() []error {
	errs := []error{}

	return errs
}