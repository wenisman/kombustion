package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

type UserPool_SmsConfiguration struct {
	ExternalId   interface{} `yaml:"ExternalId,omitempty"`
	SnsCallerArn interface{} `yaml:"SnsCallerArn,omitempty"`
}

func (resource UserPool_SmsConfiguration) Validate() []error {
	errs := []error{}

	return errs
}