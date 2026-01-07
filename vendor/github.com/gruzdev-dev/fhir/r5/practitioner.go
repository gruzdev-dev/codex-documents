package models

import (
	"encoding/json"
	"fmt"
)

// A person who is directly or indirectly involved in the provisioning of healthcare or related services.
type Practitioner struct {
	ResourceType     string                      `json:"resourceType" bson:"resource_type"`                              // Type of resource
	Id               *string                     `json:"id,omitempty" bson:"id,omitempty"`                               // Logical id of this artifact
	Meta             *Meta                       `json:"meta,omitempty" bson:"meta,omitempty"`                           // Metadata about the resource
	ImplicitRules    *string                     `json:"implicitRules,omitempty" bson:"implicit_rules,omitempty"`        // A set of rules under which this content was created
	Language         *string                     `json:"language,omitempty" bson:"language,omitempty"`                   // Language of the resource content
	Text             *Narrative                  `json:"text,omitempty" bson:"text,omitempty"`                           // Text summary of the resource, for human interpretation
	Contained        []json.RawMessage           `json:"contained,omitempty" bson:"contained,omitempty"`                 // Contained, inline Resources
	Identifier       []Identifier                `json:"identifier,omitempty" bson:"identifier,omitempty"`               // An identifier for the person as this agent
	Active           *bool                       `json:"active,omitempty" bson:"active,omitempty"`                       // Whether this practitioner's record is in active use
	Name             []HumanName                 `json:"name,omitempty" bson:"name,omitempty"`                           // The name(s) associated with the practitioner
	Telecom          []ContactPoint              `json:"telecom,omitempty" bson:"telecom,omitempty"`                     // A contact detail for the practitioner (that apply to all roles)
	Gender           *string                     `json:"gender,omitempty" bson:"gender,omitempty"`                       // male | female | other | unknown
	BirthDate        *string                     `json:"birthDate,omitempty" bson:"birth_date,omitempty"`                // The date  on which the practitioner was born
	DeceasedBoolean  *bool                       `json:"deceasedBoolean,omitempty" bson:"deceased_boolean,omitempty"`    // Indicates if the practitioner is deceased or not
	DeceasedDateTime *string                     `json:"deceasedDateTime,omitempty" bson:"deceased_date_time,omitempty"` // Indicates if the practitioner is deceased or not
	Address          []Address                   `json:"address,omitempty" bson:"address,omitempty"`                     // Address(es) of the practitioner that are not role specific (typically home address)
	Photo            []Attachment                `json:"photo,omitempty" bson:"photo,omitempty"`                         // Image of the person
	Qualification    []PractitionerQualification `json:"qualification,omitempty" bson:"qualification,omitempty"`         // Qualifications, certifications, accreditations, licenses, training, etc. pertaining to the provision of care
	Communication    []PractitionerCommunication `json:"communication,omitempty" bson:"communication,omitempty"`         // A language which may be used to communicate with the practitioner
}

func (r *Practitioner) Validate() error {
	if r.ResourceType != "Practitioner" {
		return fmt.Errorf("invalid resourceType: expected 'Practitioner', got '%s'", r.ResourceType)
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
	for i, item := range r.Qualification {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Qualification[%d]: %w", i, err)
		}
	}
	for i, item := range r.Communication {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Communication[%d]: %w", i, err)
		}
	}
	return nil
}

type PractitionerQualification struct {
	Id         *string          `json:"id,omitempty" bson:"id,omitempty"`                 // Unique id for inter-element referencing
	Identifier []Identifier     `json:"identifier,omitempty" bson:"identifier,omitempty"` // An identifier for this qualification for the practitioner
	Code       *CodeableConcept `json:"code" bson:"code"`                                 // Coded representation of the qualification
	Status     *CodeableConcept `json:"status,omitempty" bson:"status,omitempty"`         // Status/progress  of the qualification
	Period     *Period          `json:"period,omitempty" bson:"period,omitempty"`         // Period during which the qualification is valid
	Issuer     *Reference       `json:"issuer,omitempty" bson:"issuer,omitempty"`         // Organization that regulates and issues the qualification
}

func (r *PractitionerQualification) Validate() error {
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

type PractitionerCommunication struct {
	Id        *string          `json:"id,omitempty" bson:"id,omitempty"`               // Unique id for inter-element referencing
	Language  *CodeableConcept `json:"language" bson:"language"`                       // The language code used to communicate with the practitioner
	Preferred *bool            `json:"preferred,omitempty" bson:"preferred,omitempty"` // Language preference indicator
}

func (r *PractitionerCommunication) Validate() error {
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
