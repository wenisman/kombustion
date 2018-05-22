package properties

	import "fmt"

type Project_Artifacts struct {
	
	
	
	
	
	
	Location interface{} `yaml:"Location,omitempty"`
	Name interface{} `yaml:"Name,omitempty"`
	NamespaceType interface{} `yaml:"NamespaceType,omitempty"`
	Packaging interface{} `yaml:"Packaging,omitempty"`
	Path interface{} `yaml:"Path,omitempty"`
	Type interface{} `yaml:"Type"`
}

func (resource Project_Artifacts) Validate() []error {
	errs := []error{}
	
	
	
	
	
	
	if resource.Type == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Type'"))
	}
	return errs
}
