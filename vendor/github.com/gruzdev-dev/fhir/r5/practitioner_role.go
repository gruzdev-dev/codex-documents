package models

import (
	"encoding/json"
	"fmt"
)

// A specific set of Roles/Locations/specialties/services that a practitioner may perform, or has performed at an organization during a period of time.
type PractitionerRole struct {
	ResourceType      string                  `json:"resourceType" bson:"resource_type"`                               // Type of resource
	Id                *string                 `json:"id,omitempty" bson:"id,omitempty"`                                // Logical id of this artifact
	Meta              *Meta                   `json:"meta,omitempty" bson:"meta,omitempty"`                            // Metadata about the resource
	ImplicitRules     *string                 `json:"implicitRules,omitempty" bson:"implicit_rules,omitempty"`         // A set of rules under which this content was created
	Language          *string                 `json:"language,omitempty" bson:"language,omitempty"`                    // Language of the resource content
	Text              *Narrative              `json:"text,omitempty" bson:"text,omitempty"`                            // Text summary of the resource, for human interpretation
	Contained         []json.RawMessage       `json:"contained,omitempty" bson:"contained,omitempty"`                  // Contained, inline Resources
	Identifier        []Identifier            `json:"identifier,omitempty" bson:"identifier,omitempty"`                // Identifiers for a role/location
	Active            bool                    `json:"active,omitempty" bson:"active,omitempty"`                        // Whether this practitioner role record is in active use
	Period            *Period                 `json:"period,omitempty" bson:"period,omitempty"`                        // The period during which the practitioner is authorized to perform in these role(s)
	Practitioner      *Reference              `json:"practitioner,omitempty" bson:"practitioner,omitempty"`            // Practitioner that provides services for the organization
	Organization      *Reference              `json:"organization,omitempty" bson:"organization,omitempty"`            // Organization where the role is available
	Network           []Reference             `json:"network,omitempty" bson:"network,omitempty"`                      // The network in which the PractitionerRole provides the role's services (if defined) at the indicated locations (if defined)
	Code              []CodeableConcept       `json:"code,omitempty" bson:"code,omitempty"`                            // Roles which this practitioner may perform
	Display           *string                 `json:"display,omitempty" bson:"display,omitempty"`                      // Denormalized practitioner name, role, organization and location
	Specialty         []CodeableConcept       `json:"specialty,omitempty" bson:"specialty,omitempty"`                  // Specific specialty of the practitioner
	Location          []Reference             `json:"location,omitempty" bson:"location,omitempty"`                    // Location(s) where the practitioner provides care
	HealthcareService []Reference             `json:"healthcareService,omitempty" bson:"healthcare_service,omitempty"` // Healthcare services provided for this role's Organization/Location(s)
	Contact           []ExtendedContactDetail `json:"contact,omitempty" bson:"contact,omitempty"`                      // Official contact details relating to this PractitionerRole
	Characteristic    []CodeableConcept       `json:"characteristic,omitempty" bson:"characteristic,omitempty"`        // Collection of characteristics (attributes)
	Communication     []CodeableConcept       `json:"communication,omitempty" bson:"communication,omitempty"`          // A language the practitioner (in this role) can use in patient communication
	Availability      *Availability           `json:"availability,omitempty" bson:"availability,omitempty"`            // Times the Practitioner is available at this location and/or healthcare service (including exceptions)
	Endpoint          []Reference             `json:"endpoint,omitempty" bson:"endpoint,omitempty"`                    // Endpoints for interacting with the practitioner in this role
}

func (r *PractitionerRole) Validate() error {
	if r.ResourceType != "PractitionerRole" {
		return fmt.Errorf("invalid resourceType: expected 'PractitionerRole', got '%s'", r.ResourceType)
	}
	if r.Meta != nil {
		if err := r.Meta.Validate(); err != nil {
			return fmt.Errorf("Meta: %w", err)
		}
	}
	if r.Text != nil {
		if err := r.Text.Validate(); err != nil {
			return fmt.Errorf("Text: %w", err)
		}
	}
	for i, item := range r.Identifier {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Identifier[%d]: %w", i, err)
		}
	}
	if r.Period != nil {
		if err := r.Period.Validate(); err != nil {
			return fmt.Errorf("Period: %w", err)
		}
	}
	if r.Practitioner != nil {
		if err := r.Practitioner.Validate(); err != nil {
			return fmt.Errorf("Practitioner: %w", err)
		}
	}
	if r.Organization != nil {
		if err := r.Organization.Validate(); err != nil {
			return fmt.Errorf("Organization: %w", err)
		}
	}
	for i, item := range r.Network {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Network[%d]: %w", i, err)
		}
	}
	for i, item := range r.Code {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Code[%d]: %w", i, err)
		}
	}
	for i, item := range r.Specialty {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Specialty[%d]: %w", i, err)
		}
	}
	for i, item := range r.Location {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Location[%d]: %w", i, err)
		}
	}
	for i, item := range r.HealthcareService {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("HealthcareService[%d]: %w", i, err)
		}
	}
	for i, item := range r.Contact {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Contact[%d]: %w", i, err)
		}
	}
	for i, item := range r.Characteristic {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Characteristic[%d]: %w", i, err)
		}
	}
	for i, item := range r.Communication {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Communication[%d]: %w", i, err)
		}
	}
	if r.Availability != nil {
		if err := r.Availability.Validate(); err != nil {
			return fmt.Errorf("Availability: %w", err)
		}
	}
	for i, item := range r.Endpoint {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Endpoint[%d]: %w", i, err)
		}
	}
	return nil
}
