package properties


type Partition_StorageDescriptor struct {
	
	
	
	
	
	
	
	
	
	
	
	
	Compressed interface{} `yaml:"Compressed,omitempty"`
	InputFormat interface{} `yaml:"InputFormat,omitempty"`
	Location interface{} `yaml:"Location,omitempty"`
	NumberOfBuckets interface{} `yaml:"NumberOfBuckets,omitempty"`
	OutputFormat interface{} `yaml:"OutputFormat,omitempty"`
	Parameters interface{} `yaml:"Parameters,omitempty"`
	StoredAsSubDirectories interface{} `yaml:"StoredAsSubDirectories,omitempty"`
	SkewedInfo *Partition_SkewedInfo `yaml:"SkewedInfo,omitempty"`
	SerdeInfo *Partition_SerdeInfo `yaml:"SerdeInfo,omitempty"`
	BucketColumns interface{} `yaml:"BucketColumns,omitempty"`
	Columns interface{} `yaml:"Columns,omitempty"`
	SortColumns interface{} `yaml:"SortColumns,omitempty"`
}

func (resource Partition_StorageDescriptor) Validate() []error {
	errs := []error{}
	
	
	
	
	
	
	
	
	
	
	
	
	return errs
}
