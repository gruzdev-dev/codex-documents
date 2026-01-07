package models

import (
	"encoding/json"
	"fmt"
)

// The Requirements resource is used to describe an actor - a human or an application that plays a role in data exchange, and that may have obligations associated with the role the actor plays.
type Requirements struct {
	ResourceType           string                  `json:"resourceType" bson:"resource_type"`                                          // Type of resource
	Id                     *string                 `json:"id,omitempty" bson:"id,omitempty"`                                           // Logical id of this artifact
	Meta                   *Meta                   `json:"meta,omitempty" bson:"meta,omitempty"`                                       // Metadata about the resource
	ImplicitRules          *string                 `json:"implicitRules,omitempty" bson:"implicit_rules,omitempty"`                    // A set of rules under which this content was created
	Language               *string                 `json:"language,omitempty" bson:"language,omitempty"`                               // Language of the resource content
	Text                   *Narrative              `json:"text,omitempty" bson:"text,omitempty"`                                       // Text summary of the resource, for human interpretation
	Contained              []json.RawMessage       `json:"contained,omitempty" bson:"contained,omitempty"`                             // Contained, inline Resources
	Url                    *string                 `json:"url,omitempty" bson:"url,omitempty"`                                         // Canonical identifier for this Requirements, represented as a URI (globally unique)
	Identifier             []Identifier            `json:"identifier,omitempty" bson:"identifier,omitempty"`                           // Additional identifier for the Requirements (business identifier)
	Version                *string                 `json:"version,omitempty" bson:"version,omitempty"`                                 // Business version of the Requirements
	VersionAlgorithmString *string                 `json:"versionAlgorithmString,omitempty" bson:"version_algorithm_string,omitempty"` // How to compare versions
	VersionAlgorithmCoding *Coding                 `json:"versionAlgorithmCoding,omitempty" bson:"version_algorithm_coding,omitempty"` // How to compare versions
	Name                   *string                 `json:"name,omitempty" bson:"name,omitempty"`                                       // Name for this Requirements (computer friendly)
	Title                  *string                 `json:"title,omitempty" bson:"title,omitempty"`                                     // Name for this Requirements (human friendly)
	Status                 string                  `json:"status" bson:"status"`                                                       // draft | active | retired | unknown
	Experimental           *bool                   `json:"experimental,omitempty" bson:"experimental,omitempty"`                       // For testing only - never for real usage
	Date                   *string                 `json:"date,omitempty" bson:"date,omitempty"`                                       // Date last changed
	Publisher              *string                 `json:"publisher,omitempty" bson:"publisher,omitempty"`                             // Name of the publisher/steward (organization or individual)
	Contact                []ContactDetail         `json:"contact,omitempty" bson:"contact,omitempty"`                                 // Contact details for the publisher
	Description            *string                 `json:"description,omitempty" bson:"description,omitempty"`                         // Natural language description of the requirements
	UseContext             []UsageContext          `json:"useContext,omitempty" bson:"use_context,omitempty"`                          // The context that the content is intended to support
	Jurisdiction           []CodeableConcept       `json:"jurisdiction,omitempty" bson:"jurisdiction,omitempty"`                       // Jurisdiction of the authority that maintains the Requirements (if applicable)
	Purpose                *string                 `json:"purpose,omitempty" bson:"purpose,omitempty"`                                 // Why this Requirements is defined
	Copyright              *string                 `json:"copyright,omitempty" bson:"copyright,omitempty"`                             // Notice about intellectual property ownership, can include restrictions on use
	CopyrightLabel         *string                 `json:"copyrightLabel,omitempty" bson:"copyright_label,omitempty"`                  // Copyright holder and year(s)
	DerivedFrom            []string                `json:"derivedFrom,omitempty" bson:"derived_from,omitempty"`                        // Other set of Requirements this builds on
	Imports                []RequirementsImports   `json:"imports,omitempty" bson:"imports,omitempty"`                                 // External requirements that apply here
	Reference              []string                `json:"reference,omitempty" bson:"reference,omitempty"`                             // External artifact (rule/document etc. that) created this set of requirements
	Actor                  []RequirementsActor     `json:"actor,omitempty" bson:"actor,omitempty"`                                     // Actor for these requirements
	Statement              []RequirementsStatement `json:"statement,omitempty" bson:"statement,omitempty"`                             // Conformance requirement statement
}

func (r *Requirements) Validate() error {
	if r.ResourceType != "Requirements" {
		return fmt.Errorf("invalid resourceType: expected 'Requirements', got '%s'", r.ResourceType)
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
	for i, item := range r.Imports {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Imports[%d]: %w", i, err)
		}
	}
	for i, item := range r.Actor {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Actor[%d]: %w", i, err)
		}
	}
	for i, item := range r.Statement {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Statement[%d]: %w", i, err)
		}
	}
	return nil
}

type RequirementsImports struct {
	Id        *string  `json:"id,omitempty" bson:"id,omitempty"`   // Unique id for inter-element referencing
	Reference string   `json:"reference" bson:"reference"`         // Source of imported statements
	Key       []string `json:"key,omitempty" bson:"key,omitempty"` // Statement key
}

func (r *RequirementsImports) Validate() error {
	var emptyString string
	if r.Reference == emptyString {
		return fmt.Errorf("field 'Reference' is required")
	}
	return nil
}

type RequirementsActor struct {
	Id        *string `json:"id,omitempty" bson:"id,omitempty"`   // Unique id for inter-element referencing
	Reference string  `json:"reference" bson:"reference"`         // Actor referenced
	Key       *string `json:"key,omitempty" bson:"key,omitempty"` // Unique label for actor (used in statements)
}

func (r *RequirementsActor) Validate() error {
	var emptyString string
	if r.Reference == emptyString {
		return fmt.Errorf("field 'Reference' is required")
	}
	return nil
}

type RequirementsStatement struct {
	Id             *string                           `json:"id,omitempty" bson:"id,omitempty"`                         // Unique id for inter-element referencing
	Key            string                            `json:"key" bson:"key"`                                           // Key that identifies this statement
	Label          *string                           `json:"label,omitempty" bson:"label,omitempty"`                   // Short Human label for this statement
	Conformance    []string                          `json:"conformance,omitempty" bson:"conformance,omitempty"`       // SHALL | SHOULD | MAY | SHOULD-NOT | SHALL-NOT
	Conditionality *bool                             `json:"conditionality,omitempty" bson:"conditionality,omitempty"` // Set to true if requirements statement is conditional
	Requirement    string                            `json:"requirement" bson:"requirement"`                           // The actual requirement
	Category       []Coding                          `json:"category,omitempty" bson:"category,omitempty"`             // High level categorization of a statement
	DerivedFrom    *RequirementsStatementDerivedFrom `json:"derivedFrom,omitempty" bson:"derived_from,omitempty"`      // Another statement this is refining, tightening, or establishing more context for
	PartOf         *RequirementsStatementPartOf      `json:"partOf,omitempty" bson:"part_of,omitempty"`                // Higher-level requirement or statement which this is a logical sub-requirement of
	SatisfiedBy    []string                          `json:"satisfiedBy,omitempty" bson:"satisfied_by,omitempty"`      // Design artifact that satisfies this requirement
	Reference      []string                          `json:"reference,omitempty" bson:"reference,omitempty"`           // External artifact (rule/document etc. that) created this requirement
	Source         []Reference                       `json:"source,omitempty" bson:"source,omitempty"`                 // Who asked for this statement
	Actor          []string                          `json:"actor,omitempty" bson:"actor,omitempty"`                   // Key of relevant actor
}

func (r *RequirementsStatement) Validate() error {
	var emptyString string
	if r.Key == emptyString {
		return fmt.Errorf("field 'Key' is required")
	}
	if r.Requirement == emptyString {
		return fmt.Errorf("field 'Requirement' is required")
	}
	for i, item := range r.Category {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Category[%d]: %w", i, err)
		}
	}
	if r.DerivedFrom != nil {
		if err := r.DerivedFrom.Validate(); err != nil {
			return fmt.Errorf("DerivedFrom: %w", err)
		}
	}
	if r.PartOf != nil {
		if err := r.PartOf.Validate(); err != nil {
			return fmt.Errorf("PartOf: %w", err)
		}
	}
	for i, item := range r.Source {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Source[%d]: %w", i, err)
		}
	}
	return nil
}

type RequirementsStatementDerivedFrom struct {
	Id        *string `json:"id,omitempty" bson:"id,omitempty"`               // Unique id for inter-element referencing
	Reference *string `json:"reference,omitempty" bson:"reference,omitempty"` // Pointer to Requirements instance
	Key       string  `json:"key" bson:"key"`                                 // Key of referenced statement
}

func (r *RequirementsStatementDerivedFrom) Validate() error {
	var emptyString string
	if r.Key == emptyString {
		return fmt.Errorf("field 'Key' is required")
	}
	return nil
}

type RequirementsStatementPartOf struct {
	Id        *string `json:"id,omitempty" bson:"id,omitempty"`               // Unique id for inter-element referencing
	Reference *string `json:"reference,omitempty" bson:"reference,omitempty"` // Pointer to Requirements instance
	Key       string  `json:"key" bson:"key"`                                 // Key of referenced statement
}

func (r *RequirementsStatementPartOf) Validate() error {
	var emptyString string
	if r.Key == emptyString {
		return fmt.Errorf("field 'Key' is required")
	}
	return nil
}
