package models

import (
	"encoding/json"
	"fmt"
)

// A formally or informally recognized grouping of people or organizations formed for the purpose of achieving some form of collective action.
type Organization struct {
	ResourceType  string                      `json:"resourceType" bson:"resource_type"`                       // Type of resource
	Id            *string                     `json:"id,omitempty" bson:"id,omitempty"`                        // Logical id of this artifact
	Meta          *Meta                       `json:"meta,omitempty" bson:"meta,omitempty"`                    // Metadata about the resource
	ImplicitRules *string                     `json:"implicitRules,omitempty" bson:"implicit_rules,omitempty"` // A set of rules under which this content was created
	Language      *string                     `json:"language,omitempty" bson:"language,omitempty"`            // Language of the resource content
	Text          *Narrative                  `json:"text,omitempty" bson:"text,omitempty"`                    // Text summary of the resource, for human interpretation
	Contained     []json.RawMessage           `json:"contained,omitempty" bson:"contained,omitempty"`          // Contained, inline Resources
	Identifier    []Identifier                `json:"identifier,omitempty" bson:"identifier,omitempty"`        // Identifies this organization  across multiple systems
	Active        bool                        `json:"active,omitempty" bson:"active,omitempty"`                // Whether the organization's record is still in active use
	Type          []CodeableConcept           `json:"type,omitempty" bson:"type,omitempty"`                    // Kind of organization
	Name          *string                     `json:"name,omitempty" bson:"name,omitempty"`                    // Name used for the organization
	Alias         []string                    `json:"alias,omitempty" bson:"alias,omitempty"`                  // A list of alternate names that the organization is known as, or was known as in the past
	Description   *string                     `json:"description,omitempty" bson:"description,omitempty"`      // Additional details about the Organization that could be displayed as further information to identify the Organization beyond its name
	Contact       []ExtendedContactDetail     `json:"contact,omitempty" bson:"contact,omitempty"`              // Official contact details for the Organization
	PartOf        *Reference                  `json:"partOf,omitempty" bson:"part_of,omitempty"`               // The organization of which this organization forms a part
	Endpoint      []Reference                 `json:"endpoint,omitempty" bson:"endpoint,omitempty"`            // Technical endpoints providing access to services operated for the organization
	Qualification []OrganizationQualification `json:"qualification,omitempty" bson:"qualification,omitempty"`  // Qualifications, certifications, accreditations, licenses, training, etc. pertaining to the provision of care
}

func (r *Organization) Validate() error {
	if r.ResourceType != "Organization" {
		return fmt.Errorf("invalid resourceType: expected 'Organization', got '%s'", r.ResourceType)
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
	for i, item := range r.Type {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Type[%d]: %w", i, err)
		}
	}
	for i, item := range r.Contact {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Contact[%d]: %w", i, err)
		}
	}
	if r.PartOf != nil {
		if err := r.PartOf.Validate(); err != nil {
			return fmt.Errorf("PartOf: %w", err)
		}
	}
	for i, item := range r.Endpoint {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Endpoint[%d]: %w", i, err)
		}
	}
	for i, item := range r.Qualification {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Qualification[%d]: %w", i, err)
		}
	}
	return nil
}

type OrganizationQualification struct {
	Id         *string          `json:"id,omitempty" bson:"id,omitempty"`                 // Unique id for inter-element referencing
	Identifier []Identifier     `json:"identifier,omitempty" bson:"identifier,omitempty"` // An identifier for this qualification for the organization
	Code       *CodeableConcept `json:"code" bson:"code"`                                 // Coded representation of the qualification
	Status     *CodeableConcept `json:"status,omitempty" bson:"status,omitempty"`         // Status/progress of the qualification
	Period     *Period          `json:"period,omitempty" bson:"period,omitempty"`         // Period during which the qualification is valid
	Issuer     *Reference       `json:"issuer,omitempty" bson:"issuer,omitempty"`         // Organization that regulates and issues the qualification
}

func (r *OrganizationQualification) Validate() error {
	for i, item := range r.Identifier {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Identifier[%d]: %w", i, err)
		}
	}
	if r.Code == nil {
		return fmt.Errorf("field 'Code' is required")
	}
	if r.Code != nil {
		if err := r.Code.Validate(); err != nil {
			return fmt.Errorf("Code: %w", err)
		}
	}
	if r.Status != nil {
		if err := r.Status.Validate(); err != nil {
			return fmt.Errorf("Status: %w", err)
		}
	}
	if r.Period != nil {
		if err := r.Period.Validate(); err != nil {
			return fmt.Errorf("Period: %w", err)
		}
	}
	if r.Issuer != nil {
		if err := r.Issuer.Validate(); err != nil {
			return fmt.Errorf("Issuer: %w", err)
		}
	}
	return nil
}
