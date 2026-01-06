package models

import (
	"fmt"
)

// ExtendedContactDetail Type: Specifies contact information for a specific purpose over a period of time, might be handled/monitored by a specific named person or organization.
type ExtendedContactDetail struct {
	Id           *string          `json:"id,omitempty" bson:"id,omitempty"`                     // Unique id for inter-element referencing
	Purpose      *CodeableConcept `json:"purpose,omitempty" bson:"purpose,omitempty"`           // The type of contact
	Name         []HumanName      `json:"name,omitempty" bson:"name,omitempty"`                 // Name of an individual to contact
	Telecom      []ContactPoint   `json:"telecom,omitempty" bson:"telecom,omitempty"`           // Contact details (e.g.phone/fax/url)
	Address      *Address         `json:"address,omitempty" bson:"address,omitempty"`           // Address for the contact
	Organization *Reference       `json:"organization,omitempty" bson:"organization,omitempty"` // This contact detail is handled/monitored by a specific organization
	Period       *Period          `json:"period,omitempty" bson:"period,omitempty"`             // Period that this contact was valid for usage
}

func (r *ExtendedContactDetail) Validate() error {
	if r.Purpose != nil {
		if err := r.Purpose.Validate(); err != nil {
			return fmt.Errorf("Purpose: %w", err)
		}
	}
	for i, item := range r.Name {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Name[%d]: %w", i, err)
		}
	}
	for i, item := range r.Telecom {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Telecom[%d]: %w", i, err)
		}
	}
	if r.Address != nil {
		if err := r.Address.Validate(); err != nil {
			return fmt.Errorf("Address: %w", err)
		}
	}
	if r.Organization != nil {
		if err := r.Organization.Validate(); err != nil {
			return fmt.Errorf("Organization: %w", err)
		}
	}
	if r.Period != nil {
		if err := r.Period.Validate(); err != nil {
			return fmt.Errorf("Period: %w", err)
		}
	}
	return nil
}
