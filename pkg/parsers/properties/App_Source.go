package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

type App_Source struct {
	Password interface{} `yaml:"Password,omitempty"`
	Revision interface{} `yaml:"Revision,omitempty"`
	SshKey   interface{} `yaml:"SshKey,omitempty"`
	Type     interface{} `yaml:"Type,omitempty"`
	Url      interface{} `yaml:"Url,omitempty"`
	Username interface{} `yaml:"Username,omitempty"`
}

func (resource App_Source) Validate() []error {
	errs := []error{}

	return errs
}