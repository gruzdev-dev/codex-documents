package models

import "encoding/json"

// A person who is directly or indirectly involved in the provisioning of healthcare or related services.
type Practitioner struct {
	Id               *string                     `json:"id,omitempty"`               // Logical id of this artifact
	Meta             *Meta                       `json:"meta,omitempty"`             // Metadata about the resource
	ImplicitRules    *string                     `json:"implicitRules,omitempty"`    // A set of rules under which this content was created
	Language         *string                     `json:"language,omitempty"`         // Language of the resource content
	Text             *Narrative                  `json:"text,omitempty"`             // Text summary of the resource, for human interpretation
	Contained        []json.RawMessage           `json:"contained,omitempty"`        // Contained, inline Resources
	Identifier       []Identifier                `json:"identifier,omitempty"`       // An identifier for the person as this agent
	Active           bool                        `json:"active,omitempty"`           // Whether this practitioner's record is in active use
	Name             []HumanName                 `json:"name,omitempty"`             // The name(s) associated with the practitioner
	Telecom          []ContactPoint              `json:"telecom,omitempty"`          // A contact detail for the practitioner (that apply to all roles)
	Gender           *string                     `json:"gender,omitempty"`           // male | female | other | unknown
	BirthDate        *string                     `json:"birthDate,omitempty"`        // The date  on which the practitioner was born
	DeceasedBoolean  *bool                       `json:"deceasedBoolean,omitempty"`  // Indicates if the practitioner is deceased or not
	DeceasedDateTime *string                     `json:"deceasedDateTime,omitempty"` // Indicates if the practitioner is deceased or not
	Address          []Address                   `json:"address,omitempty"`          // Address(es) of the practitioner that are not role specific (typically home address)
	Photo            []Attachment                `json:"photo,omitempty"`            // Image of the person
	Qualification    []PractitionerQualification `json:"qualification,omitempty"`    // Qualifications, certifications, accreditations, licenses, training, etc. pertaining to the provision of care
	Communication    []PractitionerCommunication `json:"communication,omitempty"`    // A language which may be used to communicate with the practitioner
}

type PractitionerQualification struct {
	Id         *string          `json:"id,omitempty"`         // Unique id for inter-element referencing
	Identifier []Identifier     `json:"identifier,omitempty"` // An identifier for this qualification for the practitioner
	Code       CodeableConcept  `json:"code,omitempty"`       // Coded representation of the qualification
	Status     *CodeableConcept `json:"status,omitempty"`     // Status/progress  of the qualification
	Period     *Period          `json:"period,omitempty"`     // Period during which the qualification is valid
	Issuer     *Reference       `json:"issuer,omitempty"`     // Organization that regulates and issues the qualification
}

type PractitionerCommunication struct {
	Id        *string         `json:"id,omitempty"`        // Unique id for inter-element referencing
	Language  CodeableConcept `json:"language,omitempty"`  // The language code used to communicate with the practitioner
	Preferred bool            `json:"preferred,omitempty"` // Language preference indicator
}
