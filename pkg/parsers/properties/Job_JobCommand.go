package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

type Job_JobCommand struct {
	Name           interface{} `yaml:"Name,omitempty"`
	ScriptLocation interface{} `yaml:"ScriptLocation,omitempty"`
}

func (resource Job_JobCommand) Validate() []error {
	errs := []error{}

	return errs
}