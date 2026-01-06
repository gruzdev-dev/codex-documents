package models

import (
	"fmt"
)

// CodeableReference Type: A reference to a resource (by instance), or instead, a reference to a concept defined in a terminology or ontology (by class).
type CodeableReference struct {
	Id        *string          `json:"id,omitempty" bson:"id,omitempty"`               // Unique id for inter-element referencing
	Concept   *CodeableConcept `json:"concept,omitempty" bson:"concept,omitempty"`     // Reference to a concept (by class)
	Reference *Reference       `json:"reference,omitempty" bson:"reference,omitempty"` // Reference to a resource (by instance)
}

func (r *CodeableReference) Validate() error {
	if r.Concept != nil {
		if err := r.Concept.Validate(); err != nil {
			return fmt.Errorf("Concept: %w", err)
		}
	}
	if r.Reference != nil {
		if err := r.Reference.Validate(); err != nil {
			return fmt.Errorf("Reference: %w", err)
		}
	}
	return nil
}
