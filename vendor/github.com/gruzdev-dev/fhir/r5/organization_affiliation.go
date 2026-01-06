package models

import (
	"encoding/json"
	"fmt"
)

// Defines an affiliation/assotiation/relationship between 2 distinct organizations, that is not a part-of relationship/sub-division relationship.
type OrganizationAffiliation struct {
	Id                        *string                 `json:"id,omitempty" bson:"id,omitempty"`                                                // Logical id of this artifact
	Meta                      *Meta                   `json:"meta,omitempty" bson:"meta,omitempty"`                                            // Metadata about the resource
	ImplicitRules             *string                 `json:"implicitRules,omitempty" bson:"implicit_rules,omitempty"`                         // A set of rules under which this content was created
	Language                  *string                 `json:"language,omitempty" bson:"language,omitempty"`                                    // Language of the resource content
	Text                      *Narrative              `json:"text,omitempty" bson:"text,omitempty"`                                            // Text summary of the resource, for human interpretation
	Contained                 []json.RawMessage       `json:"contained,omitempty" bson:"contained,omitempty"`                                  // Contained, inline Resources
	Identifier                []Identifier            `json:"identifier,omitempty" bson:"identifier,omitempty"`                                // Business identifiers that are specific to this role
	Active                    bool                    `json:"active,omitempty" bson:"active,omitempty"`                                        // Whether this organization affiliation record is in active use
	Period                    *Period                 `json:"period,omitempty" bson:"period,omitempty"`                                        // The period during which the participatingOrganization is affiliated with the primary organization
	Organization              *Reference              `json:"organization,omitempty" bson:"organization,omitempty"`                            // Organization where the role is available
	ParticipatingOrganization *Reference              `json:"participatingOrganization,omitempty" bson:"participating_organization,omitempty"` // Organization that provides/performs the role (e.g. providing services or is a member of)
	Network                   []Reference             `json:"network,omitempty" bson:"network,omitempty"`                                      // The network in which the participatingOrganization provides the role's services (if defined) at the indicated locations (if defined)
	Code                      []CodeableConcept       `json:"code,omitempty" bson:"code,omitempty"`                                            // Definition of the role the participatingOrganization plays
	Specialty                 []CodeableConcept       `json:"specialty,omitempty" bson:"specialty,omitempty"`                                  // Specific specialty of the participatingOrganization in the context of the role
	Location                  []Reference             `json:"location,omitempty" bson:"location,omitempty"`                                    // The location(s) at which the role occurs
	HealthcareService         []Reference             `json:"healthcareService,omitempty" bson:"healthcare_service,omitempty"`                 // Healthcare services provided through the role
	Contact                   []ExtendedContactDetail `json:"contact,omitempty" bson:"contact,omitempty"`                                      // Official contact details at the participatingOrganization relevant to this Affiliation
	Endpoint                  []Reference             `json:"endpoint,omitempty" bson:"endpoint,omitempty"`                                    // Technical endpoints providing access to services operated for this role
}

func (r *OrganizationAffiliation) Validate() error {
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
	if r.Organization != nil {
		if err := r.Organization.Validate(); err != nil {
			return fmt.Errorf("Organization: %w", err)
		}
	}
	if r.ParticipatingOrganization != nil {
		if err := r.ParticipatingOrganization.Validate(); err != nil {
			return fmt.Errorf("ParticipatingOrganization: %w", err)
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
	for i, item := range r.Endpoint {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Endpoint[%d]: %w", i, err)
		}
	}
	return nil
}
