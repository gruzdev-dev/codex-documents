package models

import (
	"encoding/json"
	"fmt"
)

// The ActorDefinition resource is used to describe an actor - a human or an application that plays a role in data exchange, and that may have obligations associated with the role the actor plays.
type ActorDefinition struct {
	ResourceType           string            `json:"resourceType" bson:"resource_type"`                                          // Type of resource
	Id                     *string           `json:"id,omitempty" bson:"id,omitempty"`                                           // Logical id of this artifact
	Meta                   *Meta             `json:"meta,omitempty" bson:"meta,omitempty"`                                       // Metadata about the resource
	ImplicitRules          *string           `json:"implicitRules,omitempty" bson:"implicit_rules,omitempty"`                    // A set of rules under which this content was created
	Language               *string           `json:"language,omitempty" bson:"language,omitempty"`                               // Language of the resource content
	Text                   *Narrative        `json:"text,omitempty" bson:"text,omitempty"`                                       // Text summary of the resource, for human interpretation
	Contained              []json.RawMessage `json:"contained,omitempty" bson:"contained,omitempty"`                             // Contained, inline Resources
	Url                    *string           `json:"url,omitempty" bson:"url,omitempty"`                                         // Canonical identifier for this actor definition, represented as a URI (globally unique)
	Identifier             []Identifier      `json:"identifier,omitempty" bson:"identifier,omitempty"`                           // Additional identifier for the actor definition (business identifier)
	Version                *string           `json:"version,omitempty" bson:"version,omitempty"`                                 // Business version of the actor definition
	VersionAlgorithmString *string           `json:"versionAlgorithmString,omitempty" bson:"version_algorithm_string,omitempty"` // How to compare versions
	VersionAlgorithmCoding *Coding           `json:"versionAlgorithmCoding,omitempty" bson:"version_algorithm_coding,omitempty"` // How to compare versions
	Name                   *string           `json:"name,omitempty" bson:"name,omitempty"`                                       // Name for this actor definition (computer friendly)
	Title                  *string           `json:"title,omitempty" bson:"title,omitempty"`                                     // Name for this actor definition (human friendly)
	Status                 string            `json:"status" bson:"status"`                                                       // draft | active | retired | unknown
	Experimental           *bool             `json:"experimental,omitempty" bson:"experimental,omitempty"`                       // For testing only - never for real usage
	Date                   *string           `json:"date,omitempty" bson:"date,omitempty"`                                       // Date last changed
	Publisher              *string           `json:"publisher,omitempty" bson:"publisher,omitempty"`                             // Name of the publisher/steward (organization or individual)
	Contact                []ContactDetail   `json:"contact,omitempty" bson:"contact,omitempty"`                                 // Contact details for the publisher
	Description            *string           `json:"description,omitempty" bson:"description,omitempty"`                         // Natural language description of the actor
	UseContext             []UsageContext    `json:"useContext,omitempty" bson:"use_context,omitempty"`                          // The context that the content is intended to support
	Jurisdiction           []CodeableConcept `json:"jurisdiction,omitempty" bson:"jurisdiction,omitempty"`                       // Jurisdiction of the authority that maintains the actor definition (if applicable)
	Purpose                *string           `json:"purpose,omitempty" bson:"purpose,omitempty"`                                 // Why this actor definition is defined
	Copyright              *string           `json:"copyright,omitempty" bson:"copyright,omitempty"`                             // Notice about intellectual property ownership, can include restrictions on use
	CopyrightLabel         *string           `json:"copyrightLabel,omitempty" bson:"copyright_label,omitempty"`                  // Copyright holder and year(s)
	Type                   string            `json:"type" bson:"type"`                                                           // person | system | collective | other
	Category               []CodeableConcept `json:"category,omitempty" bson:"category,omitempty"`                               // Further details about the type of actor
	Documentation          *string           `json:"documentation,omitempty" bson:"documentation,omitempty"`                     // Explanation and details about the actor
	Reference              []string          `json:"reference,omitempty" bson:"reference,omitempty"`                             // Reference to more information about the actor
	BaseDefinition         []string          `json:"baseDefinition,omitempty" bson:"base_definition,omitempty"`                  // Parent actor definition
}

func (r *ActorDefinition) Validate() error {
	if r.ResourceType != "ActorDefinition" {
		return fmt.Errorf("invalid resourceType: expected 'ActorDefinition', got '%s'", r.ResourceType)
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
	if r.Type == emptyString {
		return fmt.Errorf("field 'Type' is required")
	}
	for i, item := range r.Category {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Category[%d]: %w", i, err)
		}
	}
	return nil
}
