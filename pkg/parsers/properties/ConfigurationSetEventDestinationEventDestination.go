package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// ConfigurationSetEventDestinationEventDestination Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-configurationseteventdestination-eventdestination.html
type ConfigurationSetEventDestinationEventDestination struct {
	Enabled                    interface{}                                                 `yaml:"Enabled,omitempty"`
	Name                       interface{}                                                 `yaml:"Name,omitempty"`
	MatchingEventTypes         interface{}                                                 `yaml:"MatchingEventTypes"`
	KinesisFirehoseDestination *ConfigurationSetEventDestinationKinesisFirehoseDestination `yaml:"KinesisFirehoseDestination,omitempty"`
	CloudWatchDestination      *ConfigurationSetEventDestinationCloudWatchDestination      `yaml:"CloudWatchDestination,omitempty"`
}

// ConfigurationSetEventDestinationEventDestination validation
func (resource ConfigurationSetEventDestinationEventDestination) Validate() []error {
	errors := []error{}

	if resource.MatchingEventTypes == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'MatchingEventTypes'"))
	}
	return errors
}
