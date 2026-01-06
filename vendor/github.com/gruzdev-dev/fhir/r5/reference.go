package models

import (
	"fmt"
)

// Reference Type: A reference from one resource to another.
type Reference struct {
	Id         *string     `json:"id,omitempty" bson:"id,omitempty"`                 // Unique id for inter-element referencing
	Reference  *string     `json:"reference,omitempty" bson:"reference,omitempty"`   // Literal reference, Relative, internal or absolute URL
	Type       *string     `json:"type,omitempty" bson:"type,omitempty"`             // Type the reference refers to (e.g. "Patient") - must be a resource in resources
	Identifier *Identifier `json:"identifier,omitempty" bson:"identifier,omitempty"` // Logical reference, when literal reference is not known
	Display    *string     `json:"display,omitempty" bson:"display,omitempty"`       // Text alternative for the resource
}

func (r *Reference) Validate() error {
	if r.Identifier != nil {
		if err := r.Identifier.Validate(); err != nil {
			return fmt.Errorf("Identifier: %w", err)
		}
	}
	return nil
}
