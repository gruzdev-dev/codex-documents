package models

import (
	"fmt"
)

// ParameterDefinition Type: The parameters to the module. This collection specifies both the input and output parameters. Input parameters are provided by the caller as part of the $evaluate operation. Output parameters are included in the GuidanceResponse.
type ParameterDefinition struct {
	Id            *string `json:"id,omitempty" bson:"id,omitempty"`                       // Unique id for inter-element referencing
	Name          *string `json:"name,omitempty" bson:"name,omitempty"`                   // Name used to access the parameter value
	Use           string  `json:"use" bson:"use"`                                         // in | out
	Min           *int    `json:"min,omitempty" bson:"min,omitempty"`                     // Minimum cardinality
	Max           *string `json:"max,omitempty" bson:"max,omitempty"`                     // Maximum cardinality (a number of *)
	Documentation *string `json:"documentation,omitempty" bson:"documentation,omitempty"` // A brief description of the parameter
	Type          string  `json:"type" bson:"type"`                                       // What type of value
	Profile       *string `json:"profile,omitempty" bson:"profile,omitempty"`             // What profile the value is expected to be
}

func (r *ParameterDefinition) Validate() error {
	var emptyString string
	if r.Use == emptyString {
		return fmt.Errorf("field 'Use' is required")
	}
	if r.Type == emptyString {
		return fmt.Errorf("field 'Type' is required")
	}
	return nil
}
