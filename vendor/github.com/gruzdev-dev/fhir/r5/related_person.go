package models

import (
	"encoding/json"
	"fmt"
)

// Information about a person that is involved in a patient's health or the care for a patient, but who is not the primary target of healthcare.
type RelatedPerson struct {
	ResourceType  string                       `json:"resourceType" bson:"resource_type"`                       // Type of resource
	Id            *string                      `json:"id,omitempty" bson:"id,omitempty"`                        // Logical id of this artifact
	Meta          *Meta                        `json:"meta,omitempty" bson:"meta,omitempty"`                    // Metadata about the resource
	ImplicitRules *string                      `json:"implicitRules,omitempty" bson:"implicit_rules,omitempty"` // A set of rules under which this content was created
	Language      *string                      `json:"language,omitempty" bson:"language,omitempty"`            // Language of the resource content
	Text          *Narrative                   `json:"text,omitempty" bson:"text,omitempty"`                    // Text summary of the resource, for human interpretation
	Contained     []json.RawMessage            `json:"contained,omitempty" bson:"contained,omitempty"`          // Contained, inline Resources
	Identifier    []Identifier                 `json:"identifier,omitempty" bson:"identifier,omitempty"`        // A human identifier for this person
	Active        bool                         `json:"active,omitempty" bson:"active,omitempty"`                // Whether this related person's record is in active use
	Patient       *Reference                   `json:"patient" bson:"patient"`                                  // The patient this person is related to
	Relationship  []CodeableConcept            `json:"relationship,omitempty" bson:"relationship,omitempty"`    // The personal relationship of the related person to the patient
	Role          []CodeableConcept            `json:"role,omitempty" bson:"role,omitempty"`                    // The functional role of the related person to the patient
	Name          []HumanName                  `json:"name,omitempty" bson:"name,omitempty"`                    // A name associated with the person
	Telecom       []ContactPoint               `json:"telecom,omitempty" bson:"telecom,omitempty"`              // A contact detail for the person
	Gender        *string                      `json:"gender,omitempty" bson:"gender,omitempty"`                // male | female | other | unknown
	BirthDate     *string                      `json:"birthDate,omitempty" bson:"birth_date,omitempty"`         // The date on which the related person was born
	Address       []Address                    `json:"address,omitempty" bson:"address,omitempty"`              // Address where the related person can be contacted or visited
	Photo         []Attachment                 `json:"photo,omitempty" bson:"photo,omitempty"`                  // Image of the person
	Period        *Period                      `json:"period,omitempty" bson:"period,omitempty"`                // Period of time that this relationship is considered valid
	Communication []RelatedPersonCommunication `json:"communication,omitempty" bson:"communication,omitempty"`  // A language which may be used to communicate with the related person about the patient's health
}

func (r *RelatedPerson) Validate() error {
	if r.ResourceType != "RelatedPerson" {
		return fmt.Errorf("invalid resourceType: expected 'RelatedPerson', got '%s'", r.ResourceType)
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
	if r.Patient == nil {
		return fmt.Errorf("field 'Patient' is required")
	}
	if r.Patient != nil {
		if err := r.Patient.Validate(); err != nil {
			return fmt.Errorf("Patient: %w", err)
		}
	}
	for i, item := range r.Relationship {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Relationship[%d]: %w", i, err)
		}
	}
	for i, item := range r.Role {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Role[%d]: %w", i, err)
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
	for i, item := range r.Address {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Address[%d]: %w", i, err)
		}
	}
	for i, item := range r.Photo {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Photo[%d]: %w", i, err)
		}
	}
	if r.Period != nil {
		if err := r.Period.Validate(); err != nil {
			return fmt.Errorf("Period: %w", err)
		}
	}
	for i, item := range r.Communication {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Communication[%d]: %w", i, err)
		}
	}
	return nil
}

type RelatedPersonCommunication struct {
	Id        *string          `json:"id,omitempty" bson:"id,omitempty"`               // Unique id for inter-element referencing
	Language  *CodeableConcept `json:"language" bson:"language"`                       // The language which can be used to communicate with the related person about the patient's health
	Preferred bool             `json:"preferred,omitempty" bson:"preferred,omitempty"` // Language preference indicator
}

func (r *RelatedPersonCommunication) Validate() error {
	if r.Language == nil {
		return fmt.Errorf("field 'Language' is required")
	}
	if r.Language != nil {
		if err := r.Language.Validate(); err != nil {
			return fmt.Errorf("Language: %w", err)
		}
	}
	return nil
}
