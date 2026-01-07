package models

import (
	"encoding/json"
	"fmt"
)

// A definition of a FHIR structure. This resource is used to describe the underlying resources, data types defined in FHIR, and also for describing extensions and constraints on resources and data types.
type StructureDefinition struct {
	ResourceType           string                           `json:"resourceType" bson:"resource_type"`                                          // Type of resource
	Id                     *string                          `json:"id,omitempty" bson:"id,omitempty"`                                           // Logical id of this artifact
	Meta                   *Meta                            `json:"meta,omitempty" bson:"meta,omitempty"`                                       // Metadata about the resource
	ImplicitRules          *string                          `json:"implicitRules,omitempty" bson:"implicit_rules,omitempty"`                    // A set of rules under which this content was created
	Language               *string                          `json:"language,omitempty" bson:"language,omitempty"`                               // Language of the resource content
	Text                   *Narrative                       `json:"text,omitempty" bson:"text,omitempty"`                                       // Text summary of the resource, for human interpretation
	Contained              []json.RawMessage                `json:"contained,omitempty" bson:"contained,omitempty"`                             // Contained, inline Resources
	Url                    string                           `json:"url" bson:"url"`                                                             // Canonical identifier for this structure definition, represented as a URI (globally unique)
	Identifier             []Identifier                     `json:"identifier,omitempty" bson:"identifier,omitempty"`                           // Additional identifier for the structure definition
	Version                *string                          `json:"version,omitempty" bson:"version,omitempty"`                                 // Business version of the structure definition
	VersionAlgorithmString *string                          `json:"versionAlgorithmString,omitempty" bson:"version_algorithm_string,omitempty"` // How to compare versions
	VersionAlgorithmCoding *Coding                          `json:"versionAlgorithmCoding,omitempty" bson:"version_algorithm_coding,omitempty"` // How to compare versions
	Name                   string                           `json:"name" bson:"name"`                                                           // Name for this structure definition (computer friendly)
	Title                  *string                          `json:"title,omitempty" bson:"title,omitempty"`                                     // Name for this structure definition (human friendly)
	Status                 string                           `json:"status" bson:"status"`                                                       // draft | active | retired | unknown
	Experimental           bool                             `json:"experimental,omitempty" bson:"experimental,omitempty"`                       // For testing only - never for real usage
	Date                   *string                          `json:"date,omitempty" bson:"date,omitempty"`                                       // Date last changed
	Publisher              *string                          `json:"publisher,omitempty" bson:"publisher,omitempty"`                             // Name of the publisher/steward (organization or individual)
	Contact                []ContactDetail                  `json:"contact,omitempty" bson:"contact,omitempty"`                                 // Contact details for the publisher
	Description            *string                          `json:"description,omitempty" bson:"description,omitempty"`                         // Natural language description of the structure definition
	UseContext             []UsageContext                   `json:"useContext,omitempty" bson:"use_context,omitempty"`                          // The context that the content is intended to support
	Jurisdiction           []CodeableConcept                `json:"jurisdiction,omitempty" bson:"jurisdiction,omitempty"`                       // Jurisdiction of the authority that maintains the structure definition (if applicable)
	Purpose                *string                          `json:"purpose,omitempty" bson:"purpose,omitempty"`                                 // Why this structure definition is defined
	Copyright              *string                          `json:"copyright,omitempty" bson:"copyright,omitempty"`                             // Notice about intellectual property ownership, can include restrictions on use
	CopyrightLabel         *string                          `json:"copyrightLabel,omitempty" bson:"copyright_label,omitempty"`                  // Copyright holder and year(s)
	Keyword                []Coding                         `json:"keyword,omitempty" bson:"keyword,omitempty"`                                 // Assist with indexing and finding
	FhirVersion            *string                          `json:"fhirVersion,omitempty" bson:"fhir_version,omitempty"`                        // FHIR Version this StructureDefinition targets
	Mapping                []StructureDefinitionMapping     `json:"mapping,omitempty" bson:"mapping,omitempty"`                                 // External specification that the content is mapped to
	Kind                   string                           `json:"kind" bson:"kind"`                                                           // primitive-type | complex-type | resource | logical
	Abstract               bool                             `json:"abstract" bson:"abstract"`                                                   // Whether the structure is abstract
	Context                []StructureDefinitionContext     `json:"context,omitempty" bson:"context,omitempty"`                                 // If an extension, where it can be used in instances
	ContextInvariant       []string                         `json:"contextInvariant,omitempty" bson:"context_invariant,omitempty"`              // FHIRPath invariants - when the extension can be used
	Type                   string                           `json:"type" bson:"type"`                                                           // Type defined or constrained by this structure
	BaseDefinition         *string                          `json:"baseDefinition,omitempty" bson:"base_definition,omitempty"`                  // Definition that this type is constrained/specialized from
	Derivation             *string                          `json:"derivation,omitempty" bson:"derivation,omitempty"`                           // specialization | constraint - How relates to base definition
	Snapshot               *StructureDefinitionSnapshot     `json:"snapshot,omitempty" bson:"snapshot,omitempty"`                               // Snapshot view of the structure
	Differential           *StructureDefinitionDifferential `json:"differential,omitempty" bson:"differential,omitempty"`                       // Differential view of the structure
}

func (r *StructureDefinition) Validate() error {
	if r.ResourceType != "StructureDefinition" {
		return fmt.Errorf("invalid resourceType: expected 'StructureDefinition', got '%s'", r.ResourceType)
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
	var emptyString string
	if r.Url == emptyString {
		return fmt.Errorf("field 'Url' is required")
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
	if r.Name == emptyString {
		return fmt.Errorf("field 'Name' is required")
	}
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
	for i, item := range r.Keyword {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Keyword[%d]: %w", i, err)
		}
	}
	for i, item := range r.Mapping {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Mapping[%d]: %w", i, err)
		}
	}
	if r.Kind == emptyString {
		return fmt.Errorf("field 'Kind' is required")
	}
	for i, item := range r.Context {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Context[%d]: %w", i, err)
		}
	}
	if r.Type == emptyString {
		return fmt.Errorf("field 'Type' is required")
	}
	if r.Snapshot != nil {
		if err := r.Snapshot.Validate(); err != nil {
			return fmt.Errorf("Snapshot: %w", err)
		}
	}
	if r.Differential != nil {
		if err := r.Differential.Validate(); err != nil {
			return fmt.Errorf("Differential: %w", err)
		}
	}
	return nil
}

type StructureDefinitionMapping struct {
	Id       *string `json:"id,omitempty" bson:"id,omitempty"`           // Unique id for inter-element referencing
	Identity string  `json:"identity" bson:"identity"`                   // Internal id when this mapping is used
	Uri      *string `json:"uri,omitempty" bson:"uri,omitempty"`         // Identifies what this mapping refers to
	Name     *string `json:"name,omitempty" bson:"name,omitempty"`       // Names what this mapping refers to
	Comment  *string `json:"comment,omitempty" bson:"comment,omitempty"` // Versions, Issues, Scope limitations etc
}

func (r *StructureDefinitionMapping) Validate() error {
	var emptyString string
	if r.Identity == emptyString {
		return fmt.Errorf("field 'Identity' is required")
	}
	return nil
}

type StructureDefinitionContext struct {
	Id         *string `json:"id,omitempty" bson:"id,omitempty"` // Unique id for inter-element referencing
	Type       string  `json:"type" bson:"type"`                 // fhirpath | element | extension
	Expression string  `json:"expression" bson:"expression"`     // Where the extension can be used in instances
}

func (r *StructureDefinitionContext) Validate() error {
	var emptyString string
	if r.Type == emptyString {
		return fmt.Errorf("field 'Type' is required")
	}
	if r.Expression == emptyString {
		return fmt.Errorf("field 'Expression' is required")
	}
	return nil
}

type StructureDefinitionSnapshot struct {
	Id      *string             `json:"id,omitempty" bson:"id,omitempty"` // Unique id for inter-element referencing
	Element []ElementDefinition `json:"element" bson:"element"`           // Definition of elements in the resource (if no StructureDefinition)
}

func (r *StructureDefinitionSnapshot) Validate() error {
	if len(r.Element) < 1 {
		return fmt.Errorf("field 'Element' must have at least 1 elements")
	}
	for i, item := range r.Element {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Element[%d]: %w", i, err)
		}
	}
	return nil
}

type StructureDefinitionDifferential struct {
	Id      *string             `json:"id,omitempty" bson:"id,omitempty"` // Unique id for inter-element referencing
	Element []ElementDefinition `json:"element" bson:"element"`           // Definition of elements in the resource (if no StructureDefinition)
}

func (r *StructureDefinitionDifferential) Validate() error {
	if len(r.Element) < 1 {
		return fmt.Errorf("field 'Element' must have at least 1 elements")
	}
	for i, item := range r.Element {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Element[%d]: %w", i, err)
		}
	}
	return nil
}
