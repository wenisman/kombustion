package properties


type Table_StorageDescriptor struct {
	
	
	
	
	
	
	
	
	
	
	
	
	Compressed interface{} `yaml:"Compressed,omitempty"`
	InputFormat interface{} `yaml:"InputFormat,omitempty"`
	Location interface{} `yaml:"Location,omitempty"`
	NumberOfBuckets interface{} `yaml:"NumberOfBuckets,omitempty"`
	OutputFormat interface{} `yaml:"OutputFormat,omitempty"`
	Parameters interface{} `yaml:"Parameters,omitempty"`
	StoredAsSubDirectories interface{} `yaml:"StoredAsSubDirectories,omitempty"`
	SkewedInfo *Table_SkewedInfo `yaml:"SkewedInfo,omitempty"`
	SerdeInfo *Table_SerdeInfo `yaml:"SerdeInfo,omitempty"`
	BucketColumns interface{} `yaml:"BucketColumns,omitempty"`
	Columns interface{} `yaml:"Columns,omitempty"`
	SortColumns interface{} `yaml:"SortColumns,omitempty"`
}

func (resource Table_StorageDescriptor) Validate() []error {
	errs := []error{}
	
	
	
	
	
	
	
	
	
	
	
	
	return errs
}
