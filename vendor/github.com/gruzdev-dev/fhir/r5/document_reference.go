package models

import (
	"encoding/json"
	"fmt"
)

// A reference to a document of any kind for any purpose. While the term “document” implies a more narrow focus, for this resource this “document” encompasses *any* serialized object with a mime-type, it includes formal patient-centric documents (CDA), clinical notes, scanned paper, non-patient specific documents like policy text, as well as a photo, video, or audio recording acquired or used in healthcare.  The DocumentReference resource provides metadata about the document so that the document can be discovered and managed.  The actual content may be inline base64 encoded data or provided by direct reference.
type DocumentReference struct {
	ResourceType    string                       `json:"resourceType" bson:"resource_type"`                           // Type of resource
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

func (r *DocumentReference) Validate() error {
	if r.ResourceType != "DocumentReference" {
		return fmt.Errorf("invalid resourceType: expected 'DocumentReference', got '%s'", r.ResourceType)
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
	for i, item := range r.BasedOn {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("BasedOn[%d]: %w", i, err)
		}
	}
	var emptyString string
	if r.Status == emptyString {
		return fmt.Errorf("field 'Status' is required")
	}
	for i, item := range r.Modality {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Modality[%d]: %w", i, err)
		}
	}
	if r.Type != nil {
		if err := r.Type.Validate(); err != nil {
			return fmt.Errorf("Type: %w", err)
		}
	}
	for i, item := range r.Category {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Category[%d]: %w", i, err)
		}
	}
	if r.Subject != nil {
		if err := r.Subject.Validate(); err != nil {
			return fmt.Errorf("Subject: %w", err)
		}
	}
	for i, item := range r.Context {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Context[%d]: %w", i, err)
		}
	}
	for i, item := range r.Event {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Event[%d]: %w", i, err)
		}
	}
	for i, item := range r.Related {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Related[%d]: %w", i, err)
		}
	}
	for i, item := range r.BodyStructure {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("BodyStructure[%d]: %w", i, err)
		}
	}
	if r.FacilityType != nil {
		if err := r.FacilityType.Validate(); err != nil {
			return fmt.Errorf("FacilityType: %w", err)
		}
	}
	if r.PracticeSetting != nil {
		if err := r.PracticeSetting.Validate(); err != nil {
			return fmt.Errorf("PracticeSetting: %w", err)
		}
	}
	if r.Period != nil {
		if err := r.Period.Validate(); err != nil {
			return fmt.Errorf("Period: %w", err)
		}
	}
	for i, item := range r.Author {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Author[%d]: %w", i, err)
		}
	}
	for i, item := range r.Attester {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Attester[%d]: %w", i, err)
		}
	}
	if r.Custodian != nil {
		if err := r.Custodian.Validate(); err != nil {
			return fmt.Errorf("Custodian: %w", err)
		}
	}
	for i, item := range r.RelatesTo {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("RelatesTo[%d]: %w", i, err)
		}
	}
	for i, item := range r.SecurityLabel {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("SecurityLabel[%d]: %w", i, err)
		}
	}
	if len(r.Content) < 1 {
		return fmt.Errorf("field 'Content' must have at least 1 elements")
	}
	for i, item := range r.Content {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Content[%d]: %w", i, err)
		}
	}
	return nil
}

type DocumentReferenceAttester struct {
	Id    *string          `json:"id,omitempty" bson:"id,omitempty"`       // Unique id for inter-element referencing
	Mode  *CodeableConcept `json:"mode" bson:"mode"`                       // personal | professional | legal | official
	Time  *string          `json:"time,omitempty" bson:"time,omitempty"`   // When the document was attested
	Party *Reference       `json:"party,omitempty" bson:"party,omitempty"` // Who attested the document
}

func (r *DocumentReferenceAttester) Validate() error {
	if r.Mode == nil {
		return fmt.Errorf("field 'Mode' is required")
	}
	if r.Mode != nil {
		if err := r.Mode.Validate(); err != nil {
			return fmt.Errorf("Mode: %w", err)
		}
	}
	if r.Party != nil {
		if err := r.Party.Validate(); err != nil {
			return fmt.Errorf("Party: %w", err)
		}
	}
	return nil
}

type DocumentReferenceRelatesTo struct {
	Id     *string          `json:"id,omitempty" bson:"id,omitempty"` // Unique id for inter-element referencing
	Code   *CodeableConcept `json:"code" bson:"code"`                 // The relationship type with another document
	Target *Reference       `json:"target" bson:"target"`             // Target of the relationship
}

func (r *DocumentReferenceRelatesTo) Validate() error {
	if r.Code == nil {
		return fmt.Errorf("field 'Code' is required")
	}
	if r.Code != nil {
		if err := r.Code.Validate(); err != nil {
			return fmt.Errorf("Code: %w", err)
		}
	}
	if r.Target == nil {
		return fmt.Errorf("field 'Target' is required")
	}
	if r.Target != nil {
		if err := r.Target.Validate(); err != nil {
			return fmt.Errorf("Target: %w", err)
		}
	}
	return nil
}

type DocumentReferenceContent struct {
	Id         *string                           `json:"id,omitempty" bson:"id,omitempty"`           // Unique id for inter-element referencing
	Attachment *Attachment                       `json:"attachment" bson:"attachment"`               // Where to access the document
	Profile    []DocumentReferenceContentProfile `json:"profile,omitempty" bson:"profile,omitempty"` // Content profile rules for the document
}

func (r *DocumentReferenceContent) Validate() error {
	if r.Attachment == nil {
		return fmt.Errorf("field 'Attachment' is required")
	}
	if r.Attachment != nil {
		if err := r.Attachment.Validate(); err != nil {
			return fmt.Errorf("Attachment: %w", err)
		}
	}
	for i, item := range r.Profile {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Profile[%d]: %w", i, err)
		}
	}
	return nil
}

type DocumentReferenceContentProfile struct {
	Id             *string `json:"id,omitempty" bson:"id,omitempty"`      // Unique id for inter-element referencing
	ValueCoding    *Coding `json:"valueCoding" bson:"value_coding"`       // Code|uri|canonical
	ValueUri       *string `json:"valueUri" bson:"value_uri"`             // Code|uri|canonical
	ValueCanonical *string `json:"valueCanonical" bson:"value_canonical"` // Code|uri|canonical
}

func (r *DocumentReferenceContentProfile) Validate() error {
	if r.ValueCoding == nil {
		return fmt.Errorf("field 'ValueCoding' is required")
	}
	if r.ValueCoding != nil {
		if err := r.ValueCoding.Validate(); err != nil {
			return fmt.Errorf("ValueCoding: %w", err)
		}
	}
	if r.ValueUri == nil {
		return fmt.Errorf("field 'ValueUri' is required")
	}
	if r.ValueCanonical == nil {
		return fmt.Errorf("field 'ValueCanonical' is required")
	}
	return nil
}
