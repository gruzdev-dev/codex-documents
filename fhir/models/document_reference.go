package models

import "encoding/json"

// A reference to a document of any kind for any purpose. While the term “document” implies a more narrow focus, for this resource this “document” encompasses *any* serialized object with a mime-type, it includes formal patient-centric documents (CDA), clinical notes, scanned paper, non-patient specific documents like policy text, as well as a photo, video, or audio recording acquired or used in healthcare.  The DocumentReference resource provides metadata about the document so that the document can be discovered and managed.  The actual content may be inline base64 encoded data or provided by direct reference.
type DocumentReference struct {
	Id              *string                      `json:"id,omitempty" bson:"id,omitempty"`                            // Logical id of this artifact
	Meta            *Meta                        `json:"meta,omitempty" bson:"meta,omitempty"`                        // Metadata about the resource
	ImplicitRules   *string                      `json:"implicitRules,omitempty" bson:"implicit_rules,omitempty"`     // A set of rules under which this content was created
	Language        *string                      `json:"language,omitempty" bson:"language,omitempty"`                // Language of the resource content
	Text            *Narrative                   `json:"text,omitempty" bson:"text,omitempty"`                        // Text summary of the resource, for human interpretation
	Contained       []json.RawMessage            `json:"contained,omitempty" bson:"contained,omitempty"`              // Contained, inline Resources
	Identifier      []Identifier                 `json:"identifier,omitempty" bson:"identifier,omitempty"`            // Business identifiers for the document
	Version         *string                      `json:"version,omitempty" bson:"version,omitempty"`                  // An explicitly assigned identifier of a variation of the content in the DocumentReference
	BasedOn         []Reference                  `json:"basedOn,omitempty" bson:"based_on,omitempty"`                 // Procedure that caused this media to be created
	Status          string                       `json:"status" bson:"status"`                                        // current | superseded | entered-in-error
	DocStatus       *string                      `json:"docStatus,omitempty" bson:"doc_status,omitempty"`             // registered | partial | preliminary | final | amended | corrected | appended | cancelled | entered-in-error | deprecated | unknown
	Modality        []CodeableConcept            `json:"modality,omitempty" bson:"modality,omitempty"`                // Imaging modality used
	Type            *CodeableConcept             `json:"type,omitempty" bson:"type,omitempty"`                        // Kind of document (LOINC if possible)
	Category        []CodeableConcept            `json:"category,omitempty" bson:"category,omitempty"`                // Categorization of document
	Subject         *Reference                   `json:"subject,omitempty" bson:"subject,omitempty"`                  // Who/what is the subject of the document
	Context         []Reference                  `json:"context,omitempty" bson:"context,omitempty"`                  // Encounter the document reference is part of
	Event           []CodeableReference          `json:"event,omitempty" bson:"event,omitempty"`                      // Main clinical acts documented
	Related         []Reference                  `json:"related,omitempty" bson:"related,omitempty"`                  // Related identifiers or resources associated with the document reference
	BodyStructure   []CodeableReference          `json:"bodyStructure,omitempty" bson:"body_structure,omitempty"`     // Body structure included
	FacilityType    *CodeableConcept             `json:"facilityType,omitempty" bson:"facility_type,omitempty"`       // Kind of facility where patient was seen
	PracticeSetting *CodeableConcept             `json:"practiceSetting,omitempty" bson:"practice_setting,omitempty"` // Additional details about where the content was created (e.g. clinical specialty)
	Period          *Period                      `json:"period,omitempty" bson:"period,omitempty"`                    // Time of service that is being documented
	Date            *string                      `json:"date,omitempty" bson:"date,omitempty"`                        // When this document reference was created
	Author          []Reference                  `json:"author,omitempty" bson:"author,omitempty"`                    // Who and/or what authored the document
	Attester        []DocumentReferenceAttester  `json:"attester,omitempty" bson:"attester,omitempty"`                // Attests to accuracy of the document
	Custodian       *Reference                   `json:"custodian,omitempty" bson:"custodian,omitempty"`              // Organization which maintains the document
	RelatesTo       []DocumentReferenceRelatesTo `json:"relatesTo,omitempty" bson:"relates_to,omitempty"`             // Relationships to other documents
	Description     *string                      `json:"description,omitempty" bson:"description,omitempty"`          // Human-readable description
	SecurityLabel   []CodeableConcept            `json:"securityLabel,omitempty" bson:"security_label,omitempty"`     // Document security-tags
	Content         []DocumentReferenceContent   `json:"content" bson:"content"`                                      // Document referenced
}

type DocumentReferenceContent struct {
	Id         *string                           `json:"id,omitempty" bson:"id,omitempty"`           // Unique id for inter-element referencing
	Attachment *Attachment                       `json:"attachment" bson:"attachment"`               // Where to access the document
	Profile    []DocumentReferenceContentProfile `json:"profile,omitempty" bson:"profile,omitempty"` // Content profile rules for the document
}

type DocumentReferenceContentProfile struct {
	Id             *string `json:"id,omitempty" bson:"id,omitempty"`      // Unique id for inter-element referencing
	ValueCoding    *Coding `json:"valueCoding" bson:"value_coding"`       // Code|uri|canonical
	ValueUri       *string `json:"valueUri" bson:"value_uri"`             // Code|uri|canonical
	ValueCanonical *string `json:"valueCanonical" bson:"value_canonical"` // Code|uri|canonical
}

type DocumentReferenceAttester struct {
	Id    *string          `json:"id,omitempty" bson:"id,omitempty"`       // Unique id for inter-element referencing
	Mode  *CodeableConcept `json:"mode" bson:"mode"`                       // personal | professional | legal | official
	Time  *string          `json:"time,omitempty" bson:"time,omitempty"`   // When the document was attested
	Party *Reference       `json:"party,omitempty" bson:"party,omitempty"` // Who attested the document
}

type DocumentReferenceRelatesTo struct {
	Id     *string          `json:"id,omitempty" bson:"id,omitempty"` // Unique id for inter-element referencing
	Code   *CodeableConcept `json:"code" bson:"code"`                 // The relationship type with another document
	Target *Reference       `json:"target" bson:"target"`             // Target of the relationship
}
