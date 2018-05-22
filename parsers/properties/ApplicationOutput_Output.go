package properties

	import "fmt"

type ApplicationOutput_Output struct {
	
	
	
	
	
	Name interface{} `yaml:"Name,omitempty"`
	LambdaOutput *ApplicationOutput_LambdaOutput `yaml:"LambdaOutput,omitempty"`
	KinesisStreamsOutput *ApplicationOutput_KinesisStreamsOutput `yaml:"KinesisStreamsOutput,omitempty"`
	KinesisFirehoseOutput *ApplicationOutput_KinesisFirehoseOutput `yaml:"KinesisFirehoseOutput,omitempty"`
	DestinationSchema *ApplicationOutput_DestinationSchema `yaml:"DestinationSchema"`
}

func (resource ApplicationOutput_Output) Validate() []error {
	errs := []error{}
	
	
	
	
	
	if resource.DestinationSchema == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'DestinationSchema'"))
	} else {
		errs = append(errs, resource.DestinationSchema.Validate()...)
	}
	return errs
}
