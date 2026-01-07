package models

import (
	"encoding/json"
	"fmt"
)

// A ValueSet resource instance specifies a set of codes drawn from one or more code systems, intended for use in a particular context. Value sets link between [[[CodeSystem]]] definitions and their use in [coded elements](terminologies.html).
type ValueSet struct {
	ResourceType           string             `json:"resourceType" bson:"resource_type"`                                          // Type of resource
	Id                     *string            `json:"id,omitempty" bson:"id,omitempty"`                                           // Logical id of this artifact
	Meta                   *Meta              `json:"meta,omitempty" bson:"meta,omitempty"`                                       // Metadata about the resource
	ImplicitRules          *string            `json:"implicitRules,omitempty" bson:"implicit_rules,omitempty"`                    // A set of rules under which this content was created
	Language               *string            `json:"language,omitempty" bson:"language,omitempty"`                               // Language of the resource content
	Text                   *Narrative         `json:"text,omitempty" bson:"text,omitempty"`                                       // Text summary of the resource, for human interpretation
	Contained              []json.RawMessage  `json:"contained,omitempty" bson:"contained,omitempty"`                             // Contained, inline Resources
	Url                    *string            `json:"url,omitempty" bson:"url,omitempty"`                                         // Canonical identifier for this value set, represented as a URI (globally unique)
	Identifier             []Identifier       `json:"identifier,omitempty" bson:"identifier,omitempty"`                           // Additional identifier for the value set (business identifier)
	Version                *string            `json:"version,omitempty" bson:"version,omitempty"`                                 // Business version of the value set
	VersionAlgorithmString *string            `json:"versionAlgorithmString,omitempty" bson:"version_algorithm_string,omitempty"` // How to compare versions
	VersionAlgorithmCoding *Coding            `json:"versionAlgorithmCoding,omitempty" bson:"version_algorithm_coding,omitempty"` // How to compare versions
	Name                   *string            `json:"name,omitempty" bson:"name,omitempty"`                                       // Name for this value set (computer friendly)
	Title                  *string            `json:"title,omitempty" bson:"title,omitempty"`                                     // Name for this value set (human friendly)
	Status                 string             `json:"status" bson:"status"`                                                       // draft | active | retired | unknown
	Experimental           *bool              `json:"experimental,omitempty" bson:"experimental,omitempty"`                       // For testing only - never for real usage
	Date                   *string            `json:"date,omitempty" bson:"date,omitempty"`                                       // Date last changed
	Publisher              *string            `json:"publisher,omitempty" bson:"publisher,omitempty"`                             // Name of the publisher/steward (organization or individual)
	Contact                []ContactDetail    `json:"contact,omitempty" bson:"contact,omitempty"`                                 // Contact details for the publisher
	Description            *string            `json:"description,omitempty" bson:"description,omitempty"`                         // Natural language description of the value set
	UseContext             []UsageContext     `json:"useContext,omitempty" bson:"use_context,omitempty"`                          // The context that the content is intended to support
	Jurisdiction           []CodeableConcept  `json:"jurisdiction,omitempty" bson:"jurisdiction,omitempty"`                       // Jurisdiction of the authority that maintains the  value set (if applicable)
	Immutable              *bool              `json:"immutable,omitempty" bson:"immutable,omitempty"`                             // Indicates whether or not any change to the content logical definition may occur
	Purpose                *string            `json:"purpose,omitempty" bson:"purpose,omitempty"`                                 // Why this value set is defined
	Copyright              *string            `json:"copyright,omitempty" bson:"copyright,omitempty"`                             // Notice about intellectual property ownership, can include restrictions on use
	CopyrightLabel         *string            `json:"copyrightLabel,omitempty" bson:"copyright_label,omitempty"`                  // Copyright holder and year(s)
	ApprovalDate           *string            `json:"approvalDate,omitempty" bson:"approval_date,omitempty"`                      // When the ValueSet was approved by publisher
	LastReviewDate         *string            `json:"lastReviewDate,omitempty" bson:"last_review_date,omitempty"`                 // When the ValueSet was last reviewed by the publisher
	EffectivePeriod        *Period            `json:"effectivePeriod,omitempty" bson:"effective_period,omitempty"`                // When the ValueSet is expected to be used
	Topic                  []CodeableConcept  `json:"topic,omitempty" bson:"topic,omitempty"`                                     // E.g. Education, Treatment, Assessment, etc
	Author                 []ContactDetail    `json:"author,omitempty" bson:"author,omitempty"`                                   // Who authored the ValueSet
	Editor                 []ContactDetail    `json:"editor,omitempty" bson:"editor,omitempty"`                                   // Who edited the ValueSet
	Reviewer               []ContactDetail    `json:"reviewer,omitempty" bson:"reviewer,omitempty"`                               // Who reviewed the ValueSet
	Endorser               []ContactDetail    `json:"endorser,omitempty" bson:"endorser,omitempty"`                               // Who endorsed the ValueSet
	RelatedArtifact        []RelatedArtifact  `json:"relatedArtifact,omitempty" bson:"related_artifact,omitempty"`                // Additional documentation, citations, etc
	Compose                *ValueSetCompose   `json:"compose,omitempty" bson:"compose,omitempty"`                                 // Content logical definition of the value set (CLD)
	Expansion              *ValueSetExpansion `json:"expansion,omitempty" bson:"expansion,omitempty"`                             // Used when the value set is "expanded"
}

func (r *ValueSet) Validate() error {
	if r.ResourceType != "ValueSet" {
		return fmt.Errorf("invalid resourceType: expected 'ValueSet', got '%s'", r.ResourceType)
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
	if r.Compose != nil {
		if err := r.Compose.Validate(); err != nil {
			return fmt.Errorf("Compose: %w", err)
		}
	}
	if r.Expansion != nil {
		if err := r.Expansion.Validate(); err != nil {
			return fmt.Errorf("Expansion: %w", err)
		}
	}
	return nil
}

type ValueSetComposeIncludeFilter struct {
	Id       *string `json:"id,omitempty" bson:"id,omitempty"` // Unique id for inter-element referencing
	Property string  `json:"property" bson:"property"`         // A property/filter defined by the code system
	Op       string  `json:"op" bson:"op"`                     // = | is-a | descendent-of | is-not-a | regex | in | not-in | generalizes | child-of | descendent-leaf | exists
	Value    string  `json:"value" bson:"value"`               // Code from the system, or regex criteria, or boolean value for exists
}

func (r *ValueSetComposeIncludeFilter) Validate() error {
	var emptyString string
	if r.Property == emptyString {
		return fmt.Errorf("field 'Property' is required")
	}
	if r.Op == emptyString {
		return fmt.Errorf("field 'Op' is required")
	}
	if r.Value == emptyString {
		return fmt.Errorf("field 'Value' is required")
	}
	return nil
}

type ValueSetExpansion struct {
	Id         *string                      `json:"id,omitempty" bson:"id,omitempty"`                 // Unique id for inter-element referencing
	Identifier *string                      `json:"identifier,omitempty" bson:"identifier,omitempty"` // Identifies the value set expansion (business identifier)
	Next       *string                      `json:"next,omitempty" bson:"next,omitempty"`             // Opaque urls for paging through expansion results
	Timestamp  string                       `json:"timestamp" bson:"timestamp"`                       // Time ValueSet expansion happened
	Total      *int                         `json:"total,omitempty" bson:"total,omitempty"`           // Total number of codes in the expansion
	Offset     *int                         `json:"offset,omitempty" bson:"offset,omitempty"`         // Offset at which this resource starts
	Parameter  []ValueSetExpansionParameter `json:"parameter,omitempty" bson:"parameter,omitempty"`   // Parameter that controlled the expansion process
	Property   []ValueSetExpansionProperty  `json:"property,omitempty" bson:"property,omitempty"`     // Additional information supplied about each concept
	Contains   []ValueSetExpansionContains  `json:"contains,omitempty" bson:"contains,omitempty"`     // Codes in the value set
}

func (r *ValueSetExpansion) Validate() error {
	var emptyString string
	if r.Timestamp == emptyString {
		return fmt.Errorf("field 'Timestamp' is required")
	}
	for i, item := range r.Parameter {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Parameter[%d]: %w", i, err)
		}
	}
	for i, item := range r.Property {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Property[%d]: %w", i, err)
		}
	}
	for i, item := range r.Contains {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Contains[%d]: %w", i, err)
		}
	}
	return nil
}

type ValueSetExpansionParameter struct {
	Id            *string  `json:"id,omitempty" bson:"id,omitempty"`                         // Unique id for inter-element referencing
	Name          string   `json:"name" bson:"name"`                                         // Name as assigned by the client or server
	ValueString   *string  `json:"valueString,omitempty" bson:"value_string,omitempty"`      // Value of the named parameter
	ValueBoolean  *bool    `json:"valueBoolean,omitempty" bson:"value_boolean,omitempty"`    // Value of the named parameter
	ValueInteger  *int     `json:"valueInteger,omitempty" bson:"value_integer,omitempty"`    // Value of the named parameter
	ValueDecimal  *float64 `json:"valueDecimal,omitempty" bson:"value_decimal,omitempty"`    // Value of the named parameter
	ValueUri      *string  `json:"valueUri,omitempty" bson:"value_uri,omitempty"`            // Value of the named parameter
	ValueCode     *string  `json:"valueCode,omitempty" bson:"value_code,omitempty"`          // Value of the named parameter
	ValueDateTime *string  `json:"valueDateTime,omitempty" bson:"value_date_time,omitempty"` // Value of the named parameter
}

func (r *ValueSetExpansionParameter) Validate() error {
	var emptyString string
	if r.Name == emptyString {
		return fmt.Errorf("field 'Name' is required")
	}
	return nil
}

type ValueSetExpansionProperty struct {
	Id   *string `json:"id,omitempty" bson:"id,omitempty"`   // Unique id for inter-element referencing
	Code string  `json:"code" bson:"code"`                   // Identifies the property on the concepts, and when referred to in operations
	Uri  *string `json:"uri,omitempty" bson:"uri,omitempty"` // Formal identifier for the property
}

func (r *ValueSetExpansionProperty) Validate() error {
	var emptyString string
	if r.Code == emptyString {
		return fmt.Errorf("field 'Code' is required")
	}
	return nil
}

type ValueSetExpansionContains struct {
	Id          *string                                    `json:"id,omitempty" bson:"id,omitempty"`                   // Unique id for inter-element referencing
	System      *string                                    `json:"system,omitempty" bson:"system,omitempty"`           // System value for the code
	Abstract    *bool                                      `json:"abstract,omitempty" bson:"abstract,omitempty"`       // If user cannot select this entry
	Inactive    *bool                                      `json:"inactive,omitempty" bson:"inactive,omitempty"`       // If concept is inactive in the code system
	Version     *string                                    `json:"version,omitempty" bson:"version,omitempty"`         // Version in which this code/display is defined
	Code        *string                                    `json:"code,omitempty" bson:"code,omitempty"`               // Code - if blank, this is not a selectable code
	Display     *string                                    `json:"display,omitempty" bson:"display,omitempty"`         // User display for the concept
	Designation []ValueSetComposeIncludeConceptDesignation `json:"designation,omitempty" bson:"designation,omitempty"` // Additional representations for this item
	Property    []ValueSetExpansionContainsProperty        `json:"property,omitempty" bson:"property,omitempty"`       // Property value for the concept
	Contains    []ValueSetExpansionContains                `json:"contains,omitempty" bson:"contains,omitempty"`       // Codes contained under this entry
}

func (r *ValueSetExpansionContains) Validate() error {
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
	for i, item := range r.Contains {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Contains[%d]: %w", i, err)
		}
	}
	return nil
}

type ValueSetExpansionContainsProperty struct {
	Id            *string                                        `json:"id,omitempty" bson:"id,omitempty"`                    // Unique id for inter-element referencing
	Code          string                                         `json:"code" bson:"code"`                                    // Reference to ValueSet.expansion.property.code
	ValueCode     *string                                        `json:"valueCode" bson:"value_code"`                         // Value of the property for this concept
	ValueCoding   *Coding                                        `json:"valueCoding" bson:"value_coding"`                     // Value of the property for this concept
	ValueString   *string                                        `json:"valueString" bson:"value_string"`                     // Value of the property for this concept
	ValueInteger  *int                                           `json:"valueInteger" bson:"value_integer"`                   // Value of the property for this concept
	ValueBoolean  *bool                                          `json:"valueBoolean" bson:"value_boolean"`                   // Value of the property for this concept
	ValueDateTime *string                                        `json:"valueDateTime" bson:"value_date_time"`                // Value of the property for this concept
	ValueDecimal  *float64                                       `json:"valueDecimal" bson:"value_decimal"`                   // Value of the property for this concept
	SubProperty   []ValueSetExpansionContainsPropertySubProperty `json:"subProperty,omitempty" bson:"sub_property,omitempty"` // SubProperty value for the concept
}

func (r *ValueSetExpansionContainsProperty) Validate() error {
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
	for i, item := range r.SubProperty {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("SubProperty[%d]: %w", i, err)
		}
	}
	return nil
}

type ValueSetExpansionContainsPropertySubProperty struct {
	Id            *string  `json:"id,omitempty" bson:"id,omitempty"`     // Unique id for inter-element referencing
	Code          string   `json:"code" bson:"code"`                     // Reference to ValueSet.expansion.property.code
	ValueCode     *string  `json:"valueCode" bson:"value_code"`          // Value of the subproperty for this concept
	ValueCoding   *Coding  `json:"valueCoding" bson:"value_coding"`      // Value of the subproperty for this concept
	ValueString   *string  `json:"valueString" bson:"value_string"`      // Value of the subproperty for this concept
	ValueInteger  *int     `json:"valueInteger" bson:"value_integer"`    // Value of the subproperty for this concept
	ValueBoolean  *bool    `json:"valueBoolean" bson:"value_boolean"`    // Value of the subproperty for this concept
	ValueDateTime *string  `json:"valueDateTime" bson:"value_date_time"` // Value of the subproperty for this concept
	ValueDecimal  *float64 `json:"valueDecimal" bson:"value_decimal"`    // Value of the subproperty for this concept
}

func (r *ValueSetExpansionContainsPropertySubProperty) Validate() error {
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

type ValueSetCompose struct {
	Id         *string                  `json:"id,omitempty" bson:"id,omitempty"`                  // Unique id for inter-element referencing
	LockedDate *string                  `json:"lockedDate,omitempty" bson:"locked_date,omitempty"` // Fixed date for references with no specified version (transitive)
	Inactive   *bool                    `json:"inactive,omitempty" bson:"inactive,omitempty"`      // Whether inactive codes are in the value set
	Include    []ValueSetComposeInclude `json:"include" bson:"include"`                            // Include one or more codes from a code system or other value set(s)
	Exclude    []ValueSetComposeInclude `json:"exclude,omitempty" bson:"exclude,omitempty"`        // Explicitly exclude codes from a code system or other value sets
	Property   []string                 `json:"property,omitempty" bson:"property,omitempty"`      // Property to return if client doesn't override
}

func (r *ValueSetCompose) Validate() error {
	if len(r.Include) < 1 {
		return fmt.Errorf("field 'Include' must have at least 1 elements")
	}
	for i, item := range r.Include {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Include[%d]: %w", i, err)
		}
	}
	for i, item := range r.Exclude {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Exclude[%d]: %w", i, err)
		}
	}
	return nil
}

type ValueSetComposeInclude struct {
	Id        *string                         `json:"id,omitempty" bson:"id,omitempty"`               // Unique id for inter-element referencing
	System    *string                         `json:"system,omitempty" bson:"system,omitempty"`       // The system the codes come from
	Version   *string                         `json:"version,omitempty" bson:"version,omitempty"`     // Specific version of the code system referred to
	Concept   []ValueSetComposeIncludeConcept `json:"concept,omitempty" bson:"concept,omitempty"`     // A concept defined in the system
	Filter    []ValueSetComposeIncludeFilter  `json:"filter,omitempty" bson:"filter,omitempty"`       // Select codes/concepts by their properties (including relationships)
	ValueSet  []string                        `json:"valueSet,omitempty" bson:"value_set,omitempty"`  // Select the contents included in this value set
	Copyright *string                         `json:"copyright,omitempty" bson:"copyright,omitempty"` // A copyright statement for the specific code system included in the value set
}

func (r *ValueSetComposeInclude) Validate() error {
	for i, item := range r.Concept {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Concept[%d]: %w", i, err)
		}
	}
	for i, item := range r.Filter {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Filter[%d]: %w", i, err)
		}
	}
	return nil
}

type ValueSetComposeIncludeConcept struct {
	Id          *string                                    `json:"id,omitempty" bson:"id,omitempty"`                   // Unique id for inter-element referencing
	Code        string                                     `json:"code" bson:"code"`                                   // Code or expression from system
	Display     *string                                    `json:"display,omitempty" bson:"display,omitempty"`         // Text to display for this code for this value set in this valueset
	Designation []ValueSetComposeIncludeConceptDesignation `json:"designation,omitempty" bson:"designation,omitempty"` // Additional representations for this concept
}

func (r *ValueSetComposeIncludeConcept) Validate() error {
	var emptyString string
	if r.Code == emptyString {
		return fmt.Errorf("field 'Code' is required")
	}
	for i, item := range r.Designation {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Designation[%d]: %w", i, err)
		}
	}
	return nil
}

type ValueSetComposeIncludeConceptDesignation struct {
	Id            *string  `json:"id,omitempty" bson:"id,omitempty"`                        // Unique id for inter-element referencing
	Language      *string  `json:"language,omitempty" bson:"language,omitempty"`            // Human language of the designation
	Use           *Coding  `json:"use,omitempty" bson:"use,omitempty"`                      // Types of uses of designations
	AdditionalUse []Coding `json:"additionalUse,omitempty" bson:"additional_use,omitempty"` // Additional ways how this designation would be used
	Value         string   `json:"value" bson:"value"`                                      // The text value for this designation
}

func (r *ValueSetComposeIncludeConceptDesignation) Validate() error {
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
