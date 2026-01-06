package models

import (
	"fmt"
)

// CodeableConcept Type: A concept that may be defined by a formal reference to a terminology or ontology or may be provided by text.
type CodeableConcept struct {
	Id     *string  `json:"id,omitempty" bson:"id,omitempty"`         // Unique id for inter-element referencing
	Coding []Coding `json:"coding,omitempty" bson:"coding,omitempty"` // Code defined by a terminology system
	Text   *string  `json:"text,omitempty" bson:"text,omitempty"`     // Plain text representation of the concept
}

func (r *CodeableConcept) Validate() error {
	for i, item := range r.Coding {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Coding[%d]: %w", i, err)
		}
	}
	return nil
}
