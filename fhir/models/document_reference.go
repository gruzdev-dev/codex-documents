package fhir

import "encoding/json"

// A reference to a document of any kind for any purpose. While the term “document” implies a more narrow focus, for this resource this “document” encompasses *any* serialized object with a mime-type, it includes formal patient-centric documents (CDA), clinical notes, scanned paper, non-patient specific documents like policy text, as well as a photo, video, or audio recording acquired or used in healthcare.  The DocumentReference resource provides metadata about the document so that the document can be discovered and managed.  The actual content may be inline base64 encoded data or provided by direct reference.
type DocumentReference struct {
	Id              *string                      `json:"id,omitempty"`              // Logical id of this artifact
	Meta            *Meta                        `json:"meta,omitempty"`            // Metadata about the resource
	ImplicitRules   *string                      `json:"implicitRules,omitempty"`   // A set of rules under which this content was created
	Language        *string                      `json:"language,omitempty"`        // Language of the resource content
	Text            *Narrative                   `json:"text,omitempty"`            // Text summary of the resource, for human interpretation
	Contained       []json.RawMessage            `json:"contained,omitempty"`       // Contained, inline Resources
	Identifier      []Identifier                 `json:"identifier,omitempty"`      // Business identifiers for the document
	Version         *string                      `json:"version,omitempty"`         // An explicitly assigned identifier of a variation of the content in the DocumentReference
	BasedOn         []Reference                  `json:"basedOn,omitempty"`         // Procedure that caused this media to be created
	Status          string                       `json:"status,omitempty"`          // current | superseded | entered-in-error
	DocStatus       *string                      `json:"docStatus,omitempty"`       // registered | partial | preliminary | final | amended | corrected | appended | cancelled | entered-in-error | deprecated | unknown
	Modality        []CodeableConcept            `json:"modality,omitempty"`        // Imaging modality used
	Type            *CodeableConcept             `json:"type,omitempty"`            // Kind of document (LOINC if possible)
	Category        []CodeableConcept            `json:"category,omitempty"`        // Categorization of document
	Subject         *Reference                   `json:"subject,omitempty"`         // Who/what is the subject of the document
	Context         []Reference                  `json:"context,omitempty"`         // Encounter the document reference is part of
	Event           []CodeableReference          `json:"event,omitempty"`           // Main clinical acts documented
	Related         []Reference                  `json:"related,omitempty"`         // Related identifiers or resources associated with the document reference
	BodyStructure   []CodeableReference          `json:"bodyStructure,omitempty"`   // Body structure included
	FacilityType    *CodeableConcept             `json:"facilityType,omitempty"`    // Kind of facility where patient was seen
	PracticeSetting *CodeableConcept             `json:"practiceSetting,omitempty"` // Additional details about where the content was created (e.g. clinical specialty)
	Period          *Period                      `json:"period,omitempty"`          // Time of service that is being documented
	Date            *string                      `json:"date,omitempty"`            // When this document reference was created
	Author          []Reference                  `json:"author,omitempty"`          // Who and/or what authored the document
	Attester        []DocumentReferenceAttester  `json:"attester,omitempty"`        // Attests to accuracy of the document
	Custodian       *Reference                   `json:"custodian,omitempty"`       // Organization which maintains the document
	RelatesTo       []DocumentReferenceRelatesTo `json:"relatesTo,omitempty"`       // Relationships to other documents
	Description     *string                      `json:"description,omitempty"`     // Human-readable description
	SecurityLabel   []CodeableConcept            `json:"securityLabel,omitempty"`   // Document security-tags
	Content         []DocumentReferenceContent   `json:"content,omitempty"`         // Document referenced
}

type DocumentReferenceAttester struct {
	Id    *string         `json:"id,omitempty"`    // Unique id for inter-element referencing
	Mode  CodeableConcept `json:"mode,omitempty"`  // personal | professional | legal | official
	Time  *string         `json:"time,omitempty"`  // When the document was attested
	Party *Reference      `json:"party,omitempty"` // Who attested the document
}

type DocumentReferenceRelatesTo struct {
	Id     *string         `json:"id,omitempty"`     // Unique id for inter-element referencing
	Code   CodeableConcept `json:"code,omitempty"`   // The relationship type with another document
	Target Reference       `json:"target,omitempty"` // Target of the relationship
}

type DocumentReferenceContent struct {
	Id         *string                           `json:"id,omitempty"`         // Unique id for inter-element referencing
	Attachment Attachment                        `json:"attachment,omitempty"` // Where to access the document
	Profile    []DocumentReferenceContentProfile `json:"profile,omitempty"`    // Content profile rules for the document
}

type DocumentReferenceContentProfile struct {
	Id    *string `json:"id,omitempty"`       // Unique id for inter-element referencing
	Value any     `json:"value[x],omitempty"` // Code|uri|canonical
}
