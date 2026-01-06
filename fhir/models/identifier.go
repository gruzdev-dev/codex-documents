package models

import (
	"fmt"
)

// Identifier Type: An identifier - identifies some entity uniquely and unambiguously. Typically this is used for business identifiers.
type Identifier struct {
	Id       *string          `json:"id,omitempty" bson:"id,omitempty"`             // Unique id for inter-element referencing
	Use      *string          `json:"use,omitempty" bson:"use,omitempty"`           // usual | official | temp | secondary | old (If known)
	Type     *CodeableConcept `json:"type,omitempty" bson:"type,omitempty"`         // Description of identifier
	System   *string          `json:"system,omitempty" bson:"system,omitempty"`     // The namespace for the identifier value
	Value    *string          `json:"value,omitempty" bson:"value,omitempty"`       // The value that is unique
	Period   *Period          `json:"period,omitempty" bson:"period,omitempty"`     // Time period when id is/was valid for use
	Assigner *Reference       `json:"assigner,omitempty" bson:"assigner,omitempty"` // Organization that issued id (may be just text)
}

func (r *Identifier) Validate() error {
	if r.Type != nil {
		if err := r.Type.Validate(); err != nil {
			return fmt.Errorf("Type: %w", err)
		}
	}
	if r.Period != nil {
		if err := r.Period.Validate(); err != nil {
			return fmt.Errorf("Period: %w", err)
		}
	}
	if r.Assigner != nil {
		if err := r.Assigner.Validate(); err != nil {
			return fmt.Errorf("Assigner: %w", err)
		}
	}
	return nil
}
