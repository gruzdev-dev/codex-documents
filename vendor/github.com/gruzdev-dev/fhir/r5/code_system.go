package models

import (
	"encoding/json"
	"fmt"
)

// The CodeSystem resource is used to declare the existence of and describe a code system or code system supplement and its key properties, and optionally define a part or all of its content.
type CodeSystem struct {
	Id                     *string              `json:"id,omitempty" bson:"id,omitempty"`                                           // Logical id of this artifact
	Meta                   *Meta                `json:"meta,omitempty" bson:"meta,omitempty"`                                       // Metadata about the resource
	ImplicitRules          *string              `json:"implicitRules,omitempty" bson:"implicit_rules,omitempty"`                    // A set of rules under which this content was created
	Language               *string              `json:"language,omitempty" bson:"language,omitempty"`                               // Language of the resource content
	Text                   *Narrative           `json:"text,omitempty" bson:"text,omitempty"`                                       // Text summary of the resource, for human interpretation
	Contained              []json.RawMessage    `json:"contained,omitempty" bson:"contained,omitempty"`                             // Contained, inline Resources
	Url                    *string              `json:"url,omitempty" bson:"url,omitempty"`                                         // Canonical identifier for this code system, represented as a URI (globally unique) (Coding.system)
	Identifier             []Identifier         `json:"identifier,omitempty" bson:"identifier,omitempty"`                           // Additional identifier for the code system (business identifier)
	Version                *string              `json:"version,omitempty" bson:"version,omitempty"`                                 // Business version of the code system (Coding.version)
	VersionAlgorithmString *string              `json:"versionAlgorithmString,omitempty" bson:"version_algorithm_string,omitempty"` // How to compare versions
	VersionAlgorithmCoding *Coding              `json:"versionAlgorithmCoding,omitempty" bson:"version_algorithm_coding,omitempty"` // How to compare versions
	Name                   *string              `json:"name,omitempty" bson:"name,omitempty"`                                       // Name for this code system (computer friendly)
	Title                  *string              `json:"title,omitempty" bson:"title,omitempty"`                                     // Name for this code system (human friendly)
	Status                 string               `json:"status" bson:"status"`                                                       // draft | active | retired | unknown
	Experimental           bool                 `json:"experimental,omitempty" bson:"experimental,omitempty"`                       // For testing only - never for real usage
	Date                   *string              `json:"date,omitempty" bson:"date,omitempty"`                                       // Date last changed
	Publisher              *string              `json:"publisher,omitempty" bson:"publisher,omitempty"`                             // Name of the publisher/steward (organization or individual)
	Contact                []ContactDetail      `json:"contact,omitempty" bson:"contact,omitempty"`                                 // Contact details for the publisher
	Description            *string              `json:"description,omitempty" bson:"description,omitempty"`                         // Natural language description of the code system
	UseContext             []UsageContext       `json:"useContext,omitempty" bson:"use_context,omitempty"`                          // The context that the content is intended to support
	Jurisdiction           []CodeableConcept    `json:"jurisdiction,omitempty" bson:"jurisdiction,omitempty"`                       // Jurisdiction of the authority that maintains the code system (if applicable)
	Purpose                *string              `json:"purpose,omitempty" bson:"purpose,omitempty"`                                 // Why this code system is defined
	Copyright              *string              `json:"copyright,omitempty" bson:"copyright,omitempty"`                             // Notice about intellectual property ownership, can include restrictions on use
	CopyrightLabel         *string              `json:"copyrightLabel,omitempty" bson:"copyright_label,omitempty"`                  // Copyright holder and year(s)
	ApprovalDate           *string              `json:"approvalDate,omitempty" bson:"approval_date,omitempty"`                      // When the CodeSystem was approved by publisher
	LastReviewDate         *string              `json:"lastReviewDate,omitempty" bson:"last_review_date,omitempty"`                 // When the CodeSystem was last reviewed by the publisher
	EffectivePeriod        *Period              `json:"effectivePeriod,omitempty" bson:"effective_period,omitempty"`                // When the CodeSystem is expected to be used
	Topic                  []CodeableConcept    `json:"topic,omitempty" bson:"topic,omitempty"`                                     // E.g. Education, Treatment, Assessment, etc
	Author                 []ContactDetail      `json:"author,omitempty" bson:"author,omitempty"`                                   // Who authored the CodeSystem
	Editor                 []ContactDetail      `json:"editor,omitempty" bson:"editor,omitempty"`                                   // Who edited the CodeSystem
	Reviewer               []ContactDetail      `json:"reviewer,omitempty" bson:"reviewer,omitempty"`                               // Who reviewed the CodeSystem
	Endorser               []ContactDetail      `json:"endorser,omitempty" bson:"endorser,omitempty"`                               // Who endorsed the CodeSystem
	RelatedArtifact        []RelatedArtifact    `json:"relatedArtifact,omitempty" bson:"related_artifact,omitempty"`                // Additional documentation, citations, etc
	CaseSensitive          bool                 `json:"caseSensitive,omitempty" bson:"case_sensitive,omitempty"`                    // If code comparison is case sensitive
	ValueSet               *string              `json:"valueSet,omitempty" bson:"value_set,omitempty"`                              // Canonical reference to the value set with entire code system
	HierarchyMeaning       *string              `json:"hierarchyMeaning,omitempty" bson:"hierarchy_meaning,omitempty"`              // grouped-by | is-a | part-of | classified-with
	Compositional          bool                 `json:"compositional,omitempty" bson:"compositional,omitempty"`                     // If code system defines a compositional grammar
	VersionNeeded          bool                 `json:"versionNeeded,omitempty" bson:"version_needed,omitempty"`                    // If definitions are not stable
	Content                string               `json:"content" bson:"content"`                                                     // not-present | example | fragment | complete | supplement
	Supplements            *string              `json:"supplements,omitempty" bson:"supplements,omitempty"`                         // Canonical URL of Code System this adds designations and properties to
	Count                  *int                 `json:"count,omitempty" bson:"count,omitempty"`                                     // Total concepts in the code system
	Filter                 []CodeSystemFilter   `json:"filter,omitempty" bson:"filter,omitempty"`                                   // Filter that can be used in a value set
	Property               []CodeSystemProperty `json:"property,omitempty" bson:"property,omitempty"`                               // Additional information supplied about each concept
	Concept                []CodeSystemConcept  `json:"concept,omitempty" bson:"concept,omitempty"`                                 // Concepts in the code system
}

func (r *CodeSystem) Validate() error {
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
	if r.Content == emptyString {
		return fmt.Errorf("field 'Content' is required")
	}
	for i, item := range r.Filter {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Filter[%d]: %w", i, err)
		}
	}
	for i, item := range r.Property {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Property[%d]: %w", i, err)
		}
	}
	for i, item := range r.Concept {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Concept[%d]: %w", i, err)
		}
	}
	return nil
}

type CodeSystemConceptDesignation struct {
	Id            *string  `json:"id,omitempty" bson:"id,omitempty"`                        // Unique id for inter-element referencing
	Language      *string  `json:"language,omitempty" bson:"language,omitempty"`            // Human language of the designation
	Use           *Coding  `json:"use,omitempty" bson:"use,omitempty"`                      // Details how this designation would be used
	AdditionalUse []Coding `json:"additionalUse,omitempty" bson:"additional_use,omitempty"` // Additional ways how this designation would be used
	Value         string   `json:"value" bson:"value"`                                      // The text value for this designation
}

func (r *CodeSystemConceptDesignation) Validate() error {
	if r.Use != nil {
		if err := r.Use.Validate(); err != nil {
			return fmt.Errorf("Use: %w", err)
		}
	}
	for i, item := range r.AdditionalUse {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("AdditionalUse[%d]: %w", i, err)
		}
	}
	var emptyString string
	if r.Value == emptyString {
		return fmt.Errorf("field 'Value' is required")
	}
	return nil
}

type CodeSystemConceptProperty struct {
	Id            *string  `json:"id,omitempty" bson:"id,omitempty"`     // Unique id for inter-element referencing
	Code          string   `json:"code" bson:"code"`                     // Reference to CodeSystem.property.code or a FHIR defined concept-property
	ValueCode     *string  `json:"valueCode" bson:"value_code"`          // Value of the property for this concept
	ValueCoding   *Coding  `json:"valueCoding" bson:"value_coding"`      // Value of the property for this concept
	ValueString   *string  `json:"valueString" bson:"value_string"`      // Value of the property for this concept
	ValueInteger  *int     `json:"valueInteger" bson:"value_integer"`    // Value of the property for this concept
	ValueBoolean  *bool    `json:"valueBoolean" bson:"value_boolean"`    // Value of the property for this concept
	ValueDateTime *string  `json:"valueDateTime" bson:"value_date_time"` // Value of the property for this concept
	ValueDecimal  *float64 `json:"valueDecimal" bson:"value_decimal"`    // Value of the property for this concept
}

func (r *CodeSystemConceptProperty) Validate() error {
	var emptyString string
	if r.Code == emptyString {
		return fmt.Errorf("field 'Code' is required")
	}
	if r.ValueCode == nil {
		return fmt.Errorf("field 'ValueCode' is required")
	}
	if r.ValueCoding == nil {
		return fmt.Errorf("field 'ValueCoding' is required")
	}
	if r.ValueCoding != nil {
		if err := r.ValueCoding.Validate(); err != nil {
			return fmt.Errorf("ValueCoding: %w", err)
		}
	}
	if r.ValueString == nil {
		return fmt.Errorf("field 'ValueString' is required")
	}
	if r.ValueInteger == nil {
		return fmt.Errorf("field 'ValueInteger' is required")
	}
	if r.ValueBoolean == nil {
		return fmt.Errorf("field 'ValueBoolean' is required")
	}
	if r.ValueDateTime == nil {
		return fmt.Errorf("field 'ValueDateTime' is required")
	}
	if r.ValueDecimal == nil {
		return fmt.Errorf("field 'ValueDecimal' is required")
	}
	return nil
}

type CodeSystemFilter struct {
	Id          *string  `json:"id,omitempty" bson:"id,omitempty"`                   // Unique id for inter-element referencing
	Code        string   `json:"code" bson:"code"`                                   // Code that identifies the filter
	Description *string  `json:"description,omitempty" bson:"description,omitempty"` // How or why the filter is used
	Operator    []string `json:"operator" bson:"operator"`                           // = | is-a | descendent-of | is-not-a | regex | in | not-in | generalizes | child-of | descendent-leaf | exists
	Value       string   `json:"value" bson:"value"`                                 // What to use for the value
}

func (r *CodeSystemFilter) Validate() error {
	var emptyString string
	if r.Code == emptyString {
		return fmt.Errorf("field 'Code' is required")
	}
	if len(r.Operator) < 1 {
		return fmt.Errorf("field 'Operator' must have at least 1 elements")
	}
	if r.Value == emptyString {
		return fmt.Errorf("field 'Value' is required")
	}
	return nil
}

type CodeSystemProperty struct {
	Id          *string `json:"id,omitempty" bson:"id,omitempty"`                   // Unique id for inter-element referencing
	Code        string  `json:"code" bson:"code"`                                   // Identifies the property on the concepts, and when referred to in operations
	Uri         *string `json:"uri,omitempty" bson:"uri,omitempty"`                 // Formal identifier for the property
	Description *string `json:"description,omitempty" bson:"description,omitempty"` // Why the property is defined, and/or what it conveys
	Type        string  `json:"type" bson:"type"`                                   // code | Coding | string | integer | boolean | dateTime | decimal
}

func (r *CodeSystemProperty) Validate() error {
	var emptyString string
	if r.Code == emptyString {
		return fmt.Errorf("field 'Code' is required")
	}
	if r.Type == emptyString {
		return fmt.Errorf("field 'Type' is required")
	}
	return nil
}

type CodeSystemConcept struct {
	Id          *string                        `json:"id,omitempty" bson:"id,omitempty"`                   // Unique id for inter-element referencing
	Code        string                         `json:"code" bson:"code"`                                   // Code that identifies concept
	Display     *string                        `json:"display,omitempty" bson:"display,omitempty"`         // Text to display to the user
	Definition  *string                        `json:"definition,omitempty" bson:"definition,omitempty"`   // Formal definition
	Designation []CodeSystemConceptDesignation `json:"designation,omitempty" bson:"designation,omitempty"` // Additional representations for the concept
	Property    []CodeSystemConceptProperty    `json:"property,omitempty" bson:"property,omitempty"`       // Property value for the concept
	Concept     []CodeSystemConcept            `json:"concept,omitempty" bson:"concept,omitempty"`         // Child Concepts (is-a/contains/categorizes)
}

func (r *CodeSystemConcept) Validate() error {
	var emptyString string
	if r.Code == emptyString {
		return fmt.Errorf("field 'Code' is required")
	}
	for i, item := range r.Designation {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Designation[%d]: %w", i, err)
		}
	}
	for i, item := range r.Property {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Property[%d]: %w", i, err)
		}
	}
	for i, item := range r.Concept {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Concept[%d]: %w", i, err)
		}
	}
	return nil
}
