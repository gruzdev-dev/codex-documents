package models

import (
	"encoding/json"
	"fmt"
)

// A search parameter that defines a named search item that can be used to search/filter on a resource.
type SearchParameter struct {
	Id                     *string                    `json:"id,omitempty" bson:"id,omitempty"`                                           // Logical id of this artifact
	Meta                   *Meta                      `json:"meta,omitempty" bson:"meta,omitempty"`                                       // Metadata about the resource
	ImplicitRules          *string                    `json:"implicitRules,omitempty" bson:"implicit_rules,omitempty"`                    // A set of rules under which this content was created
	Language               *string                    `json:"language,omitempty" bson:"language,omitempty"`                               // Language of the resource content
	Text                   *Narrative                 `json:"text,omitempty" bson:"text,omitempty"`                                       // Text summary of the resource, for human interpretation
	Contained              []json.RawMessage          `json:"contained,omitempty" bson:"contained,omitempty"`                             // Contained, inline Resources
	Url                    string                     `json:"url" bson:"url"`                                                             // Canonical identifier for this search parameter, represented as a URI (globally unique)
	Identifier             []Identifier               `json:"identifier,omitempty" bson:"identifier,omitempty"`                           // Additional identifier for the search parameter (business identifier)
	Version                *string                    `json:"version,omitempty" bson:"version,omitempty"`                                 // Business version of the search parameter
	VersionAlgorithmString *string                    `json:"versionAlgorithmString,omitempty" bson:"version_algorithm_string,omitempty"` // How to compare versions
	VersionAlgorithmCoding *Coding                    `json:"versionAlgorithmCoding,omitempty" bson:"version_algorithm_coding,omitempty"` // How to compare versions
	Name                   string                     `json:"name" bson:"name"`                                                           // Name for this search parameter (computer friendly)
	Title                  *string                    `json:"title,omitempty" bson:"title,omitempty"`                                     // Name for this search parameter (human friendly)
	DerivedFrom            *string                    `json:"derivedFrom,omitempty" bson:"derived_from,omitempty"`                        // Original definition for the search parameter
	Status                 string                     `json:"status" bson:"status"`                                                       // draft | active | retired | unknown
	Experimental           bool                       `json:"experimental,omitempty" bson:"experimental,omitempty"`                       // For testing only - never for real usage
	Date                   *string                    `json:"date,omitempty" bson:"date,omitempty"`                                       // Date last changed
	Publisher              *string                    `json:"publisher,omitempty" bson:"publisher,omitempty"`                             // Name of the publisher/steward (organization or individual)
	Contact                []ContactDetail            `json:"contact,omitempty" bson:"contact,omitempty"`                                 // Contact details for the publisher
	Description            string                     `json:"description" bson:"description"`                                             // Natural language description of the search parameter
	UseContext             []UsageContext             `json:"useContext,omitempty" bson:"use_context,omitempty"`                          // The context that the content is intended to support
	Jurisdiction           []CodeableConcept          `json:"jurisdiction,omitempty" bson:"jurisdiction,omitempty"`                       // Jurisdiction of the authority that maintains the search parameter (if applicable)
	Purpose                *string                    `json:"purpose,omitempty" bson:"purpose,omitempty"`                                 // Why this search parameter is defined
	Copyright              *string                    `json:"copyright,omitempty" bson:"copyright,omitempty"`                             // Notice about intellectual property ownership, can include restrictions on use
	CopyrightLabel         *string                    `json:"copyrightLabel,omitempty" bson:"copyright_label,omitempty"`                  // Copyright holder and year(s)
	Code                   string                     `json:"code" bson:"code"`                                                           // Recommended name for parameter in search url
	AliasCode              []string                   `json:"aliasCode,omitempty" bson:"alias_code,omitempty"`                            // Additional recommended names for parameter in search url
	Base                   []string                   `json:"base" bson:"base"`                                                           // The resource type(s) this search parameter applies to
	Type                   string                     `json:"type" bson:"type"`                                                           // number | date | string | token | reference | composite | quantity | uri | special | resource
	Expression             *string                    `json:"expression,omitempty" bson:"expression,omitempty"`                           // FHIRPath expression that extracts the values
	ProcessingMode         *string                    `json:"processingMode,omitempty" bson:"processing_mode,omitempty"`                  // normal | phonetic | other
	Constraint             *string                    `json:"constraint,omitempty" bson:"constraint,omitempty"`                           // FHIRPath expression that constraints the usage of this SearchParameter
	Target                 []string                   `json:"target,omitempty" bson:"target,omitempty"`                                   // Types of resource (if a resource reference)
	MultipleOr             bool                       `json:"multipleOr,omitempty" bson:"multiple_or,omitempty"`                          // Allow multiple values per parameter (or)
	MultipleAnd            bool                       `json:"multipleAnd,omitempty" bson:"multiple_and,omitempty"`                        // Allow multiple parameters (and)
	Comparator             []string                   `json:"comparator,omitempty" bson:"comparator,omitempty"`                           // eq | ne | gt | lt | ge | le | sa | eb | ap
	Modifier               []string                   `json:"modifier,omitempty" bson:"modifier,omitempty"`                               // missing | exact | contains | not | text | in | not-in | below | above | type | identifier | of-type | code-text | text-advanced | iterate
	Chain                  []string                   `json:"chain,omitempty" bson:"chain,omitempty"`                                     // Chained names supported
	Component              []SearchParameterComponent `json:"component,omitempty" bson:"component,omitempty"`                             // For Composite resources to define the parts
}

func (r *SearchParameter) Validate() error {
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
	if r.Description == emptyString {
		return fmt.Errorf("field 'Description' is required")
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
	if r.Code == emptyString {
		return fmt.Errorf("field 'Code' is required")
	}
	if len(r.Base) < 1 {
		return fmt.Errorf("field 'Base' must have at least 1 elements")
	}
	if r.Type == emptyString {
		return fmt.Errorf("field 'Type' is required")
	}
	for i, item := range r.Component {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Component[%d]: %w", i, err)
		}
	}
	return nil
}

type SearchParameterComponent struct {
	Id         *string `json:"id,omitempty" bson:"id,omitempty"` // Unique id for inter-element referencing
	Definition string  `json:"definition" bson:"definition"`     // Defines how the part works
	Expression string  `json:"expression" bson:"expression"`     // Sub-expression relative to main expression
}

func (r *SearchParameterComponent) Validate() error {
	var emptyString string
	if r.Definition == emptyString {
		return fmt.Errorf("field 'Definition' is required")
	}
	if r.Expression == emptyString {
		return fmt.Errorf("field 'Expression' is required")
	}
	return nil
}
