package models

import (
	"encoding/json"
	"fmt"
)

// Common Interface declaration for conformance and knowledge artifact resources.
type CanonicalResource struct {
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
	Experimental           *bool             `json:"experimental,omitempty" bson:"experimental,omitempty"`                       // For testing only - never for real usage
	Date                   *string           `json:"date,omitempty" bson:"date,omitempty"`                                       // Date last changed
	Publisher              *string           `json:"publisher,omitempty" bson:"publisher,omitempty"`                             // Name of the publisher/steward (organization or individual)
	Contact                []ContactDetail   `json:"contact,omitempty" bson:"contact,omitempty"`                                 // Contact details for the publisher
	Description            *string           `json:"description,omitempty" bson:"description,omitempty"`                         // Natural language description of the {{title}}
	UseContext             []UsageContext    `json:"useContext,omitempty" bson:"use_context,omitempty"`                          // The context that the content is intended to support
	Jurisdiction           []CodeableConcept `json:"jurisdiction,omitempty" bson:"jurisdiction,omitempty"`                       // Jurisdiction of the authority that maintains the {{title}} (if applicable)
	Purpose                *string           `json:"purpose,omitempty" bson:"purpose,omitempty"`                                 // Why this {{title}} is defined
	Copyright              *string           `json:"copyright,omitempty" bson:"copyright,omitempty"`                             // Notice about intellectual property ownership, can include restrictions on use
	CopyrightLabel         *string           `json:"copyrightLabel,omitempty" bson:"copyright_label,omitempty"`                  // Copyright holder and year(s)
}

func (r *CanonicalResource) Validate() error {
	if r.ResourceType != "CanonicalResource" {
		return fmt.Errorf("invalid resourceType: expected 'CanonicalResource', got '%s'", r.ResourceType)
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
	return nil
}
