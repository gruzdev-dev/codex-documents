package models

import (
	"encoding/json"
	"fmt"
)

// Demographics and other administrative information about an individual or animal that is the subject of potential, past, current, or future health-related care, services, or processes.
type Patient struct {
	Id                   *string                `json:"id,omitempty" bson:"id,omitempty"`                                       // Logical id of this artifact
	Meta                 *Meta                  `json:"meta,omitempty" bson:"meta,omitempty"`                                   // Metadata about the resource
	ImplicitRules        *string                `json:"implicitRules,omitempty" bson:"implicit_rules,omitempty"`                // A set of rules under which this content was created
	Language             *string                `json:"language,omitempty" bson:"language,omitempty"`                           // Language of the resource content
	Text                 *Narrative             `json:"text,omitempty" bson:"text,omitempty"`                                   // Text summary of the resource, for human interpretation
	Contained            []json.RawMessage      `json:"contained,omitempty" bson:"contained,omitempty"`                         // Contained, inline Resources
	Identifier           []Identifier           `json:"identifier,omitempty" bson:"identifier,omitempty"`                       // An identifier for this patient
	Active               bool                   `json:"active,omitempty" bson:"active,omitempty"`                               // Whether this patient's record is in active use
	Name                 []HumanName            `json:"name,omitempty" bson:"name,omitempty"`                                   // A name associated with the patient
	Telecom              []ContactPoint         `json:"telecom,omitempty" bson:"telecom,omitempty"`                             // A contact detail for the individual
	Gender               *string                `json:"gender,omitempty" bson:"gender,omitempty"`                               // male | female | other | unknown
	BirthDate            *string                `json:"birthDate,omitempty" bson:"birth_date,omitempty"`                        // The date of birth for the individual
	DeceasedBoolean      *bool                  `json:"deceasedBoolean,omitempty" bson:"deceased_boolean,omitempty"`            // Indicates if/when the individual is deceased
	DeceasedDateTime     *string                `json:"deceasedDateTime,omitempty" bson:"deceased_date_time,omitempty"`         // Indicates if/when the individual is deceased
	Address              []Address              `json:"address,omitempty" bson:"address,omitempty"`                             // An address for the individual
	MaritalStatus        *CodeableConcept       `json:"maritalStatus,omitempty" bson:"marital_status,omitempty"`                // Marital (civil) status of a patient
	MultipleBirthBoolean *bool                  `json:"multipleBirthBoolean,omitempty" bson:"multiple_birth_boolean,omitempty"` // Whether patient is part of a multiple birth
	MultipleBirthInteger *int                   `json:"multipleBirthInteger,omitempty" bson:"multiple_birth_integer,omitempty"` // Whether patient is part of a multiple birth
	Photo                []Attachment           `json:"photo,omitempty" bson:"photo,omitempty"`                                 // Image of the patient
	Contact              []PatientContact       `json:"contact,omitempty" bson:"contact,omitempty"`                             // A contact party (e.g. guardian, partner, friend) for the patient
	Communication        []PatientCommunication `json:"communication,omitempty" bson:"communication,omitempty"`                 // A language which may be used to communicate with the patient about his or her health
	GeneralPractitioner  []Reference            `json:"generalPractitioner,omitempty" bson:"general_practitioner,omitempty"`    // Patient's nominated primary care provider
	ManagingOrganization *Reference             `json:"managingOrganization,omitempty" bson:"managing_organization,omitempty"`  // Organization that is the custodian of the patient record
	Link                 []PatientLink          `json:"link,omitempty" bson:"link,omitempty"`                                   // Link to a Patient or RelatedPerson resource that concerns the same actual individual
}

func (r *Patient) Validate() error {
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
	if r.MaritalStatus != nil {
		if err := r.MaritalStatus.Validate(); err != nil {
			return fmt.Errorf("MaritalStatus: %w", err)
		}
	}
	for i, item := range r.Photo {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Photo[%d]: %w", i, err)
		}
	}
	for i, item := range r.Contact {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Contact[%d]: %w", i, err)
		}
	}
	for i, item := range r.Communication {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Communication[%d]: %w", i, err)
		}
	}
	for i, item := range r.GeneralPractitioner {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("GeneralPractitioner[%d]: %w", i, err)
		}
	}
	if r.ManagingOrganization != nil {
		if err := r.ManagingOrganization.Validate(); err != nil {
			return fmt.Errorf("ManagingOrganization: %w", err)
		}
	}
	for i, item := range r.Link {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Link[%d]: %w", i, err)
		}
	}
	return nil
}

type PatientContact struct {
	Id                *string           `json:"id,omitempty" bson:"id,omitempty"`                                // Unique id for inter-element referencing
	Relationship      []CodeableConcept `json:"relationship,omitempty" bson:"relationship,omitempty"`            // The kind of personal relationship
	Role              []CodeableConcept `json:"role,omitempty" bson:"role,omitempty"`                            // The kind of functional role
	Name              *HumanName        `json:"name,omitempty" bson:"name,omitempty"`                            // A name associated with the contact person
	AdditionalName    []HumanName       `json:"additionalName,omitempty" bson:"additional_name,omitempty"`       // Additional names for the contact person
	Telecom           []ContactPoint    `json:"telecom,omitempty" bson:"telecom,omitempty"`                      // A contact detail for the person
	Address           *Address          `json:"address,omitempty" bson:"address,omitempty"`                      // Address for the contact person
	AdditionalAddress []Address         `json:"additionalAddress,omitempty" bson:"additional_address,omitempty"` // Additional addresses for the contact person
	Gender            *string           `json:"gender,omitempty" bson:"gender,omitempty"`                        // male | female | other | unknown
	Organization      *Reference        `json:"organization,omitempty" bson:"organization,omitempty"`            // Organization that is associated with the contact
	Period            *Period           `json:"period,omitempty" bson:"period,omitempty"`                        // The period during which this contact person or organization is valid to be contacted relating to this patient
}

func (r *PatientContact) Validate() error {
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
	if r.Name != nil {
		if err := r.Name.Validate(); err != nil {
			return fmt.Errorf("Name: %w", err)
		}
	}
	for i, item := range r.AdditionalName {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("AdditionalName[%d]: %w", i, err)
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
	for i, item := range r.AdditionalAddress {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("AdditionalAddress[%d]: %w", i, err)
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

type PatientCommunication struct {
	Id        *string          `json:"id,omitempty" bson:"id,omitempty"`               // Unique id for inter-element referencing
	Language  *CodeableConcept `json:"language" bson:"language"`                       // The language which can be used to communicate with the patient about his or her health
	Preferred bool             `json:"preferred,omitempty" bson:"preferred,omitempty"` // Language preference indicator
}

func (r *PatientCommunication) Validate() error {
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

type PatientLink struct {
	Id    *string    `json:"id,omitempty" bson:"id,omitempty"` // Unique id for inter-element referencing
	Other *Reference `json:"other" bson:"other"`               // The other patient or related person resource that the link refers to
	Type  string     `json:"type" bson:"type"`                 // replaced-by | replaces | refer | seealso
}

func (r *PatientLink) Validate() error {
	if r.Other == nil {
		return fmt.Errorf("field 'Other' is required")
	}
	if r.Other != nil {
		if err := r.Other.Validate(); err != nil {
			return fmt.Errorf("Other: %w", err)
		}
	}
	var emptyString string
	if r.Type == emptyString {
		return fmt.Errorf("field 'Type' is required")
	}
	return nil
}
