package models

import (
	"fmt"
)

// ContactDetail Type: Specifies contact information for a person or organization.
type ContactDetail struct {
	Id      *string        `json:"id,omitempty" bson:"id,omitempty"`           // Unique id for inter-element referencing
	Name    *string        `json:"name,omitempty" bson:"name,omitempty"`       // Name of an individual to contact
	Telecom []ContactPoint `json:"telecom,omitempty" bson:"telecom,omitempty"` // Contact details for individual or organization
}

func (r *ContactDetail) Validate() error {
	for i, item := range r.Telecom {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Telecom[%d]: %w", i, err)
		}
	}
	return nil
}
