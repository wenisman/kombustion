package properties

	import "fmt"

type TopicRule_DynamoDBAction struct {
	
	
	
	
	
	
	
	
	
	HashKeyField interface{} `yaml:"HashKeyField"`
	HashKeyType interface{} `yaml:"HashKeyType,omitempty"`
	HashKeyValue interface{} `yaml:"HashKeyValue"`
	PayloadField interface{} `yaml:"PayloadField,omitempty"`
	RangeKeyField interface{} `yaml:"RangeKeyField,omitempty"`
	RangeKeyType interface{} `yaml:"RangeKeyType,omitempty"`
	RangeKeyValue interface{} `yaml:"RangeKeyValue,omitempty"`
	RoleArn interface{} `yaml:"RoleArn"`
	TableName interface{} `yaml:"TableName"`
}

func (resource TopicRule_DynamoDBAction) Validate() []error {
	errs := []error{}
	
	
	
	
	
	
	
	
	
	if resource.HashKeyField == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'HashKeyField'"))
	}
	if resource.HashKeyValue == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'HashKeyValue'"))
	}
	if resource.RoleArn == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'RoleArn'"))
	}
	if resource.TableName == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'TableName'"))
	}
	return errs
}
