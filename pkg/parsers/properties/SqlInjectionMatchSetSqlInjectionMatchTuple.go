package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// SqlInjectionMatchSetSqlInjectionMatchTuple Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-sqlinjectionmatchset-sqlinjectionmatchtuples.html
type SqlInjectionMatchSetSqlInjectionMatchTuple struct {
	TextTransformation interface{}                       `yaml:"TextTransformation"`
	FieldToMatch       *SqlInjectionMatchSetFieldToMatch `yaml:"FieldToMatch"`
}

// SqlInjectionMatchSetSqlInjectionMatchTuple validation
func (resource SqlInjectionMatchSetSqlInjectionMatchTuple) Validate() []error {
	errors := []error{}

	if resource.TextTransformation == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'TextTransformation'"))
	}
	if resource.FieldToMatch == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'FieldToMatch'"))
	} else {
		errors = append(errors, resource.FieldToMatch.Validate()...)
	}
	return errors
}
