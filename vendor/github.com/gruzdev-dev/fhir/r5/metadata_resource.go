package models

import (
	"encoding/json"
	"fmt"
)

// Common Interface declaration for conformance and knowledge artifact resources.
type MetadataResource struct {
	ResourceType           string            `json:"resourceType" bson:"resource_type"`                                          // Type of resource
	Id                     *string           `json:"id,omitempty" bson:"id,omitempty"`                                           // Logical id of this artifact
	Meta                   *Meta             `json:"meta,omitempty" bson:"meta,omitempty"`                                       // Metadata about the resource
	ImplicitRules          *string           `json:"implicitRules,omitempty" bson:"implicit_rules,omitempty"`                    // A set of rules under which this content was created
	Language               *string           `json:"language,omitempty" bson:"language,omitempty"`                               // Language of the resource content
	Text                   *Narrative        `json:"text,omitempty" bson:"text,omitempty"`                                       // Text summary of the resource, for human interpretation
	Contained              []json.RawMessage `json:"contained,omitempty" bson:"contained,omitempty"`                             // Contained, inline Resources
	Url                    *string           `json:"url,omitempty" bson:"url,omitempty"`                                         // Canonical identifier for this {{title}}, represented as an absolute URI (globally unique)
	Identifier             []Identifier      `json:"identifier,omitempty" bson:"identifier,omitempty"`                           // Additional identifier for the {{title}}
	Version                *string           `json:"version,omitempty" bson:"version,omitempty"`                                 // Canonical version of the {{title}}
	VersionAlgorithmString *string           `json:"versionAlgorithmString,omitempty" bson:"version_algorithm_string,omitempty"` // How to compare versions
	VersionAlgorithmCoding *Coding           `json:"versionAlgorithmCoding,omitempty" bson:"version_algorithm_coding,omitempty"` // How to compare versions
	Name                   *string           `json:"name,omitempty" bson:"name,omitempty"`                                       // Name for this {{title}} (computer friendly)
	Title                  *string           `json:"title,omitempty" bson:"title,omitempty"`                                     // Name for this {{title}} (human friendly)
	Status                 string            `json:"status" bson:"status"`                                                       // draft | active | retired | unknown
	Experimental           bool              `json:"experimental,omitempty" bson:"experimental,omitempty"`                       // For testing only - never for real usage
	Date                   *string           `json:"date,omitempty" bson:"date,omitempty"`                                       // Date last changed
	Publisher              *string           `json:"publisher,omitempty" bson:"publisher,omitempty"`                             // Name of the publisher/steward (organization or individual)
	Contact                []ContactDetail   `json:"contact,omitempty" bson:"contact,omitempty"`                                 // Contact details for the publisher
	Description            *string           `json:"description,omitempty" bson:"description,omitempty"`                         // Natural language description of the {{title}}
	UseContext             []UsageContext    `json:"useContext,omitempty" bson:"use_context,omitempty"`                          // The context that the content is intended to support
	Jurisdiction           []CodeableConcept `json:"jurisdiction,omitempty" bson:"jurisdiction,omitempty"`                       // Jurisdiction of the authority that maintains the {{title}} (if applicable)
	Purpose                *string           `json:"purpose,omitempty" bson:"purpose,omitempty"`                                 // Why this {{title}} is defined
	Copyright              *string           `json:"copyright,omitempty" bson:"copyright,omitempty"`                             // Notice about intellectual property ownership, can include restrictions on use
	CopyrightLabel         *string           `json:"copyrightLabel,omitempty" bson:"copyright_label,omitempty"`                  // Copyright holder and year(s)
	Usage                  *string           `json:"usage,omitempty" bson:"usage,omitempty"`                                     // Describes the clinical usage of the {{title}}
	ApprovalDate           *string           `json:"approvalDate,omitempty" bson:"approval_date,omitempty"`                      // When the {{title}} was approved by publisher
	LastReviewDate         *string           `json:"lastReviewDate,omitempty" bson:"last_review_date,omitempty"`                 // When the {{title}} was last reviewed by the publisher
	EffectivePeriod        *Period           `json:"effectivePeriod,omitempty" bson:"effective_period,omitempty"`                // When the {{title}} is expected to be used
	Topic                  []CodeableConcept `json:"topic,omitempty" bson:"topic,omitempty"`                                     // E.g. Education, Treatment, Assessment, etc
	Author                 []ContactDetail   `json:"author,omitempty" bson:"author,omitempty"`                                   // Who authored the {{title}}
	Editor                 []ContactDetail   `json:"editor,omitempty" bson:"editor,omitempty"`                                   // Who edited the {{title}}
	Reviewer               []ContactDetail   `json:"reviewer,omitempty" bson:"reviewer,omitempty"`                               // Who reviewed the {{title}}
	Endorser               []ContactDetail   `json:"endorser,omitempty" bson:"endorser,omitempty"`                               // Who endorsed the {{title}}
	RelatedArtifact        []RelatedArtifact `json:"relatedArtifact,omitempty" bson:"related_artifact,omitempty"`                // Additional documentation, citations, etc
}

func (r *MetadataResource) Validate() error {
	if r.ResourceType != "MetadataResource" {
		return fmt.Errorf("invalid resourceType: expected 'MetadataResource', got '%s'", r.ResourceType)
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
	return nil
}
