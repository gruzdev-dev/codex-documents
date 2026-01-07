package models

import (
	"encoding/json"
	"fmt"
)

// A curated namespace that issues unique symbols within that namespace for the identification of concepts, people, devices, etc.  Represents a "System" used within the Identifier and Coding data types.
type NamingSystem struct {
	ResourceType           string                 `json:"resourceType" bson:"resource_type"`                                          // Type of resource
	Id                     *string                `json:"id,omitempty" bson:"id,omitempty"`                                           // Logical id of this artifact
	Meta                   *Meta                  `json:"meta,omitempty" bson:"meta,omitempty"`                                       // Metadata about the resource
	ImplicitRules          *string                `json:"implicitRules,omitempty" bson:"implicit_rules,omitempty"`                    // A set of rules under which this content was created
	Language               *string                `json:"language,omitempty" bson:"language,omitempty"`                               // Language of the resource content
	Text                   *Narrative             `json:"text,omitempty" bson:"text,omitempty"`                                       // Text summary of the resource, for human interpretation
	Contained              []json.RawMessage      `json:"contained,omitempty" bson:"contained,omitempty"`                             // Contained, inline Resources
	Url                    *string                `json:"url,omitempty" bson:"url,omitempty"`                                         // Canonical identifier for this naming system, represented as a URI (globally unique)
	Identifier             []Identifier           `json:"identifier,omitempty" bson:"identifier,omitempty"`                           // Additional identifier for the naming system (business identifier)
	Version                *string                `json:"version,omitempty" bson:"version,omitempty"`                                 // Business version of the naming system
	VersionAlgorithmString *string                `json:"versionAlgorithmString,omitempty" bson:"version_algorithm_string,omitempty"` // How to compare versions
	VersionAlgorithmCoding *Coding                `json:"versionAlgorithmCoding,omitempty" bson:"version_algorithm_coding,omitempty"` // How to compare versions
	Name                   string                 `json:"name" bson:"name"`                                                           // Name for this naming system (computer friendly)
	Title                  *string                `json:"title,omitempty" bson:"title,omitempty"`                                     // Title for this naming system (human friendly)
	Status                 string                 `json:"status" bson:"status"`                                                       // draft | active | retired | unknown
	Kind                   string                 `json:"kind" bson:"kind"`                                                           // codesystem | identifier | root
	Experimental           *bool                  `json:"experimental,omitempty" bson:"experimental,omitempty"`                       // For testing only - never for real usage
	Date                   string                 `json:"date" bson:"date"`                                                           // Date last changed
	Publisher              *string                `json:"publisher,omitempty" bson:"publisher,omitempty"`                             // Name of the publisher/steward (organization or individual)
	Contact                []ContactDetail        `json:"contact,omitempty" bson:"contact,omitempty"`                                 // Contact details for the publisher
	Responsible            *string                `json:"responsible,omitempty" bson:"responsible,omitempty"`                         // Who maintains system namespace?
	Type                   *CodeableConcept       `json:"type,omitempty" bson:"type,omitempty"`                                       // e.g. driver,  provider,  patient, bank etc
	Description            *string                `json:"description,omitempty" bson:"description,omitempty"`                         // Natural language description of the naming system
	UseContext             []UsageContext         `json:"useContext,omitempty" bson:"use_context,omitempty"`                          // The context that the content is intended to support
	Jurisdiction           []CodeableConcept      `json:"jurisdiction,omitempty" bson:"jurisdiction,omitempty"`                       // Jurisdiction of the authority that maintains the naming system (if applicable)
	Purpose                *string                `json:"purpose,omitempty" bson:"purpose,omitempty"`                                 // Why this naming system is defined
	Copyright              *string                `json:"copyright,omitempty" bson:"copyright,omitempty"`                             // Notice about intellectual property ownership, can include restrictions on use
	CopyrightLabel         *string                `json:"copyrightLabel,omitempty" bson:"copyright_label,omitempty"`                  // Copyright holder and year(s)
	ApprovalDate           *string                `json:"approvalDate,omitempty" bson:"approval_date,omitempty"`                      // When the NamingSystem was approved by publisher
	LastReviewDate         *string                `json:"lastReviewDate,omitempty" bson:"last_review_date,omitempty"`                 // When the NamingSystem was last reviewed by the publisher
	EffectivePeriod        *Period                `json:"effectivePeriod,omitempty" bson:"effective_period,omitempty"`                // When the NamingSystem is expected to be used
	Topic                  []CodeableConcept      `json:"topic,omitempty" bson:"topic,omitempty"`                                     // E.g. Education, Treatment, Assessment, etc
	Author                 []ContactDetail        `json:"author,omitempty" bson:"author,omitempty"`                                   // Who authored the CodeSystem
	Editor                 []ContactDetail        `json:"editor,omitempty" bson:"editor,omitempty"`                                   // Who edited the NamingSystem
	Reviewer               []ContactDetail        `json:"reviewer,omitempty" bson:"reviewer,omitempty"`                               // Who reviewed the NamingSystem
	Endorser               []ContactDetail        `json:"endorser,omitempty" bson:"endorser,omitempty"`                               // Who endorsed the NamingSystem
	RelatedArtifact        []RelatedArtifact      `json:"relatedArtifact,omitempty" bson:"related_artifact,omitempty"`                // Additional documentation, citations, etc
	Usage                  *string                `json:"usage,omitempty" bson:"usage,omitempty"`                                     // How/where is it used
	UniqueId               []NamingSystemUniqueId `json:"uniqueId" bson:"unique_id"`                                                  // Unique identifiers used for system
}

func (r *NamingSystem) Validate() error {
	if r.ResourceType != "NamingSystem" {
		return fmt.Errorf("invalid resourceType: expected 'NamingSystem', got '%s'", r.ResourceType)
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
	if r.VersionAlgorithmCoding != nil {
		if err := r.VersionAlgorithmCoding.Validate(); err != nil {
			return fmt.Errorf("VersionAlgorithmCoding: %w", err)
		}
	}
	var emptyString string
	if r.Name == emptyString {
		return fmt.Errorf("field 'Name' is required")
	}
	if r.Status == emptyString {
		return fmt.Errorf("field 'Status' is required")
	}
	if r.Kind == emptyString {
		return fmt.Errorf("field 'Kind' is required")
	}
	if r.Date == emptyString {
		return fmt.Errorf("field 'Date' is required")
	}
	for i, item := range r.Contact {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Contact[%d]: %w", i, err)
		}
	}
	if r.Type != nil {
		if err := r.Type.Validate(); err != nil {
			return fmt.Errorf("Type: %w", err)
		}
	}
	for i, item := range r.UseContext {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("UseContext[%d]: %w", i, err)
		}
	}
	for i, item := range r.Jurisdiction {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Jurisdiction[%d]: %w", i, err)
		}
	}
	if r.EffectivePeriod != nil {
		if err := r.EffectivePeriod.Validate(); err != nil {
			return fmt.Errorf("EffectivePeriod: %w", err)
		}
	}
	for i, item := range r.Topic {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Topic[%d]: %w", i, err)
		}
	}
	for i, item := range r.Author {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Author[%d]: %w", i, err)
		}
	}
	for i, item := range r.Editor {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Editor[%d]: %w", i, err)
		}
	}
	for i, item := range r.Reviewer {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Reviewer[%d]: %w", i, err)
		}
	}
	for i, item := range r.Endorser {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Endorser[%d]: %w", i, err)
		}
	}
	for i, item := range r.RelatedArtifact {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("RelatedArtifact[%d]: %w", i, err)
		}
	}
	if len(r.UniqueId) < 1 {
		return fmt.Errorf("field 'UniqueId' must have at least 1 elements")
	}
	for i, item := range r.UniqueId {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("UniqueId[%d]: %w", i, err)
		}
	}
	return nil
}

type NamingSystemUniqueId struct {
	Id            *string `json:"id,omitempty" bson:"id,omitempty"`                       // Unique id for inter-element referencing
	Type          string  `json:"type" bson:"type"`                                       // oid | uuid | uri | iri-stem | v2csmnemonic | other
	Value         string  `json:"value" bson:"value"`                                     // The unique identifier
	Preferred     *bool   `json:"preferred,omitempty" bson:"preferred,omitempty"`         // Is this the id that should be used for this type
	Comment       *string `json:"comment,omitempty" bson:"comment,omitempty"`             // Notes about identifier usage
	Period        *Period `json:"period,omitempty" bson:"period,omitempty"`               // When is identifier valid?
	Authoritative *bool   `json:"authoritative,omitempty" bson:"authoritative,omitempty"` // Whether the identifier is authoritative
}

func (r *NamingSystemUniqueId) Validate() error {
	var emptyString string
	if r.Type == emptyString {
		return fmt.Errorf("field 'Type' is required")
	}
	if r.Value == emptyString {
		return fmt.Errorf("field 'Value' is required")
	}
	if r.Period != nil {
		if err := r.Period.Validate(); err != nil {
			return fmt.Errorf("Period: %w", err)
		}
	}
	return nil
}
