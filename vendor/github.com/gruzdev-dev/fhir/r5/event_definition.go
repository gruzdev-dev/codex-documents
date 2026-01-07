package models

import (
	"encoding/json"
	"fmt"
)

// The EventDefinition resource provides a reusable description of when a particular event can occur.
type EventDefinition struct {
	ResourceType           string              `json:"resourceType" bson:"resource_type"`                                          // Type of resource
	Id                     *string             `json:"id,omitempty" bson:"id,omitempty"`                                           // Logical id of this artifact
	Meta                   *Meta               `json:"meta,omitempty" bson:"meta,omitempty"`                                       // Metadata about the resource
	ImplicitRules          *string             `json:"implicitRules,omitempty" bson:"implicit_rules,omitempty"`                    // A set of rules under which this content was created
	Language               *string             `json:"language,omitempty" bson:"language,omitempty"`                               // Language of the resource content
	Text                   *Narrative          `json:"text,omitempty" bson:"text,omitempty"`                                       // Text summary of the resource, for human interpretation
	Contained              []json.RawMessage   `json:"contained,omitempty" bson:"contained,omitempty"`                             // Contained, inline Resources
	Url                    *string             `json:"url,omitempty" bson:"url,omitempty"`                                         // Canonical identifier for this event definition, represented as a URI (globally unique)
	Identifier             []Identifier        `json:"identifier,omitempty" bson:"identifier,omitempty"`                           // Additional identifier for the event definition
	Version                *string             `json:"version,omitempty" bson:"version,omitempty"`                                 // Business version of the event definition
	VersionAlgorithmString *string             `json:"versionAlgorithmString,omitempty" bson:"version_algorithm_string,omitempty"` // How to compare versions
	VersionAlgorithmCoding *Coding             `json:"versionAlgorithmCoding,omitempty" bson:"version_algorithm_coding,omitempty"` // How to compare versions
	Name                   *string             `json:"name,omitempty" bson:"name,omitempty"`                                       // Name for this event definition (computer friendly)
	Title                  *string             `json:"title,omitempty" bson:"title,omitempty"`                                     // Name for this event definition (human friendly)
	Subtitle               *string             `json:"subtitle,omitempty" bson:"subtitle,omitempty"`                               // Subordinate title of the event definition
	Status                 string              `json:"status" bson:"status"`                                                       // draft | active | retired | unknown
	Experimental           bool                `json:"experimental,omitempty" bson:"experimental,omitempty"`                       // For testing only - never for real usage
	SubjectCodeableConcept *CodeableConcept    `json:"subjectCodeableConcept,omitempty" bson:"subject_codeable_concept,omitempty"` // Type of individual the event definition is focused on
	SubjectReference       *Reference          `json:"subjectReference,omitempty" bson:"subject_reference,omitempty"`              // Type of individual the event definition is focused on
	Date                   *string             `json:"date,omitempty" bson:"date,omitempty"`                                       // Date last changed
	Publisher              *string             `json:"publisher,omitempty" bson:"publisher,omitempty"`                             // Name of the publisher/steward (organization or individual)
	Contact                []ContactDetail     `json:"contact,omitempty" bson:"contact,omitempty"`                                 // Contact details for the publisher
	Description            *string             `json:"description,omitempty" bson:"description,omitempty"`                         // Natural language description of the event definition
	UseContext             []UsageContext      `json:"useContext,omitempty" bson:"use_context,omitempty"`                          // The context that the content is intended to support
	Jurisdiction           []CodeableConcept   `json:"jurisdiction,omitempty" bson:"jurisdiction,omitempty"`                       // Jurisdiction of the authority that maintains the event definition (if applicable)
	Purpose                *string             `json:"purpose,omitempty" bson:"purpose,omitempty"`                                 // Why this event definition is defined
	Usage                  *string             `json:"usage,omitempty" bson:"usage,omitempty"`                                     // Describes the clinical usage of the event definition
	Copyright              *string             `json:"copyright,omitempty" bson:"copyright,omitempty"`                             // Notice about intellectual property ownership, can include restrictions on use
	CopyrightLabel         *string             `json:"copyrightLabel,omitempty" bson:"copyright_label,omitempty"`                  // Copyright holder and year(s)
	ApprovalDate           *string             `json:"approvalDate,omitempty" bson:"approval_date,omitempty"`                      // When the event definition was approved by publisher
	LastReviewDate         *string             `json:"lastReviewDate,omitempty" bson:"last_review_date,omitempty"`                 // When the event definition was last reviewed by the publisher
	EffectivePeriod        *Period             `json:"effectivePeriod,omitempty" bson:"effective_period,omitempty"`                // When the event definition is expected to be used
	Topic                  []CodeableConcept   `json:"topic,omitempty" bson:"topic,omitempty"`                                     // E.g. Education, Treatment, Assessment, etc
	Author                 []ContactDetail     `json:"author,omitempty" bson:"author,omitempty"`                                   // Who authored the content
	Editor                 []ContactDetail     `json:"editor,omitempty" bson:"editor,omitempty"`                                   // Who edited the content
	Reviewer               []ContactDetail     `json:"reviewer,omitempty" bson:"reviewer,omitempty"`                               // Who reviewed the content
	Endorser               []ContactDetail     `json:"endorser,omitempty" bson:"endorser,omitempty"`                               // Who endorsed the content
	RelatedArtifact        []RelatedArtifact   `json:"relatedArtifact,omitempty" bson:"related_artifact,omitempty"`                // Additional documentation, citations, etc
	Trigger                []TriggerDefinition `json:"trigger" bson:"trigger"`                                                     // "when" the event occurs (multiple = 'or')
}

func (r *EventDefinition) Validate() error {
	if r.ResourceType != "EventDefinition" {
		return fmt.Errorf("invalid resourceType: expected 'EventDefinition', got '%s'", r.ResourceType)
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
	if r.Status == emptyString {
		return fmt.Errorf("field 'Status' is required")
	}
	if r.SubjectCodeableConcept != nil {
		if err := r.SubjectCodeableConcept.Validate(); err != nil {
			return fmt.Errorf("SubjectCodeableConcept: %w", err)
		}
	}
	if r.SubjectReference != nil {
		if err := r.SubjectReference.Validate(); err != nil {
			return fmt.Errorf("SubjectReference: %w", err)
		}
	}
	for i, item := range r.Contact {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Contact[%d]: %w", i, err)
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
	if len(r.Trigger) < 1 {
		return fmt.Errorf("field 'Trigger' must have at least 1 elements")
	}
	for i, item := range r.Trigger {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Trigger[%d]: %w", i, err)
		}
	}
	return nil
}
