package properties


type Table_TableInput struct {
	
	
	
	
	
	
	
	
	
	
	Description interface{} `yaml:"Description,omitempty"`
	Name interface{} `yaml:"Name,omitempty"`
	Owner interface{} `yaml:"Owner,omitempty"`
	Parameters interface{} `yaml:"Parameters,omitempty"`
	Retention interface{} `yaml:"Retention,omitempty"`
	TableType interface{} `yaml:"TableType,omitempty"`
	ViewExpandedText interface{} `yaml:"ViewExpandedText,omitempty"`
	ViewOriginalText interface{} `yaml:"ViewOriginalText,omitempty"`
	StorageDescriptor *Table_StorageDescriptor `yaml:"StorageDescriptor,omitempty"`
	PartitionKeys interface{} `yaml:"PartitionKeys,omitempty"`
}

func (resource Table_TableInput) Validate() []error {
	errs := []error{}
	
	
	
	
	
	
	
	
	
	
	return errs
}
