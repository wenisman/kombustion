package properties


type Application_MappingParameters struct {
	
	
	JSONMappingParameters *Application_JSONMappingParameters `yaml:"JSONMappingParameters,omitempty"`
	CSVMappingParameters *Application_CSVMappingParameters `yaml:"CSVMappingParameters,omitempty"`
}

func (resource Application_MappingParameters) Validate() []error {
	errs := []error{}
	
	
	return errs
}
