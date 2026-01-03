package fhir

import "encoding/json"

// Demographics and other administrative information about an individual or animal that is the subject of potential, past, current, or future health-related care, services, or processes.
type Patient struct {
	Id                   *string                `json:"id,omitempty"`                   // Logical id of this artifact
	Meta                 *Meta                  `json:"meta,omitempty"`                 // Metadata about the resource
	ImplicitRules        *string                `json:"implicitRules,omitempty"`        // A set of rules under which this content was created
	Language             *string                `json:"language,omitempty"`             // Language of the resource content
	Text                 *Narrative             `json:"text,omitempty"`                 // Text summary of the resource, for human interpretation
	Contained            []json.RawMessage      `json:"contained,omitempty"`            // Contained, inline Resources
	Identifier           []Identifier           `json:"identifier,omitempty"`           // An identifier for this patient
	Active               bool                   `json:"active,omitempty"`               // Whether this patient's record is in active use
	Name                 []HumanName            `json:"name,omitempty"`                 // A name associated with the patient
	Telecom              []ContactPoint         `json:"telecom,omitempty"`              // A contact detail for the individual
	Gender               *string                `json:"gender,omitempty"`               // male | female | other | unknown
	BirthDate            *string                `json:"birthDate,omitempty"`            // The date of birth for the individual
	Deceased             any                    `json:"deceased[x],omitempty"`          // Indicates if/when the individual is deceased
	Address              []Address              `json:"address,omitempty"`              // An address for the individual
	MaritalStatus        *CodeableConcept       `json:"maritalStatus,omitempty"`        // Marital (civil) status of a patient
	MultipleBirth        any                    `json:"multipleBirth[x],omitempty"`     // Whether patient is part of a multiple birth
	Photo                []Attachment           `json:"photo,omitempty"`                // Image of the patient
	Contact              []PatientContact       `json:"contact,omitempty"`              // A contact party (e.g. guardian, partner, friend) for the patient
	Communication        []PatientCommunication `json:"communication,omitempty"`        // A language which may be used to communicate with the patient about his or her health
	GeneralPractitioner  []Reference            `json:"generalPractitioner,omitempty"`  // Patient's nominated primary care provider
	ManagingOrganization *Reference             `json:"managingOrganization,omitempty"` // Organization that is the custodian of the patient record
	Link                 []PatientLink          `json:"link,omitempty"`                 // Link to a Patient or RelatedPerson resource that concerns the same actual individual
}

type PatientContact struct {
	Id                *string           `json:"id,omitempty"`                // Unique id for inter-element referencing
	Relationship      []CodeableConcept `json:"relationship,omitempty"`      // The kind of personal relationship
	Role              []CodeableConcept `json:"role,omitempty"`              // The kind of functional role
	Name              *HumanName        `json:"name,omitempty"`              // A name associated with the contact person
	AdditionalName    []HumanName       `json:"additionalName,omitempty"`    // Additional names for the contact person
	Telecom           []ContactPoint    `json:"telecom,omitempty"`           // A contact detail for the person
	Address           *Address          `json:"address,omitempty"`           // Address for the contact person
	AdditionalAddress []Address         `json:"additionalAddress,omitempty"` // Additional addresses for the contact person
	Gender            *string           `json:"gender,omitempty"`            // male | female | other | unknown
	Organization      *Reference        `json:"organization,omitempty"`      // Organization that is associated with the contact
	Period            *Period           `json:"period,omitempty"`            // The period during which this contact person or organization is valid to be contacted relating to this patient
}

type PatientCommunication struct {
	Id        *string         `json:"id,omitempty"`        // Unique id for inter-element referencing
	Language  CodeableConcept `json:"language,omitempty"`  // The language which can be used to communicate with the patient about his or her health
	Preferred bool            `json:"preferred,omitempty"` // Language preference indicator
}

type PatientLink struct {
	Id    *string   `json:"id,omitempty"`    // Unique id for inter-element referencing
	Other Reference `json:"other,omitempty"` // The other patient or related person resource that the link refers to
	Type  string    `json:"type,omitempty"`  // replaced-by | replaces | refer | seealso
}
