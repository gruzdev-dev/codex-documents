package models

import (
	"encoding/json"
	"fmt"
)

// A Map of relationships between 2 structures that can be used to transform data.
type StructureMap struct {
	Id                     *string                 `json:"id,omitempty" bson:"id,omitempty"`                                           // Logical id of this artifact
	Meta                   *Meta                   `json:"meta,omitempty" bson:"meta,omitempty"`                                       // Metadata about the resource
	ImplicitRules          *string                 `json:"implicitRules,omitempty" bson:"implicit_rules,omitempty"`                    // A set of rules under which this content was created
	Language               *string                 `json:"language,omitempty" bson:"language,omitempty"`                               // Language of the resource content
	Text                   *Narrative              `json:"text,omitempty" bson:"text,omitempty"`                                       // Text summary of the resource, for human interpretation
	Contained              []json.RawMessage       `json:"contained,omitempty" bson:"contained,omitempty"`                             // Contained, inline Resources
	Url                    string                  `json:"url" bson:"url"`                                                             // Canonical identifier for this structure map, represented as a URI (globally unique)
	Identifier             []Identifier            `json:"identifier,omitempty" bson:"identifier,omitempty"`                           // Additional identifier for the structure map
	Version                *string                 `json:"version,omitempty" bson:"version,omitempty"`                                 // Business version of the structure map
	VersionAlgorithmString *string                 `json:"versionAlgorithmString,omitempty" bson:"version_algorithm_string,omitempty"` // How to compare versions
	VersionAlgorithmCoding *Coding                 `json:"versionAlgorithmCoding,omitempty" bson:"version_algorithm_coding,omitempty"` // How to compare versions
	Name                   string                  `json:"name" bson:"name"`                                                           // Name for this structure map (computer friendly)
	Title                  *string                 `json:"title,omitempty" bson:"title,omitempty"`                                     // Name for this structure map (human friendly)
	Status                 string                  `json:"status" bson:"status"`                                                       // draft | active | retired | unknown
	Experimental           bool                    `json:"experimental,omitempty" bson:"experimental,omitempty"`                       // For testing only - never for real usage
	Date                   *string                 `json:"date,omitempty" bson:"date,omitempty"`                                       // Date last changed
	Publisher              *string                 `json:"publisher,omitempty" bson:"publisher,omitempty"`                             // Name of the publisher/steward (organization or individual)
	Contact                []ContactDetail         `json:"contact,omitempty" bson:"contact,omitempty"`                                 // Contact details for the publisher
	Description            *string                 `json:"description,omitempty" bson:"description,omitempty"`                         // Natural language description of the structure map
	UseContext             []UsageContext          `json:"useContext,omitempty" bson:"use_context,omitempty"`                          // The context that the content is intended to support
	Jurisdiction           []CodeableConcept       `json:"jurisdiction,omitempty" bson:"jurisdiction,omitempty"`                       // Jurisdiction of the authority that maintains the structure map (if applicable)
	Purpose                *string                 `json:"purpose,omitempty" bson:"purpose,omitempty"`                                 // Why this structure map is defined
	Copyright              *string                 `json:"copyright,omitempty" bson:"copyright,omitempty"`                             // Notice about intellectual property ownership, can include restrictions on use
	CopyrightLabel         *string                 `json:"copyrightLabel,omitempty" bson:"copyright_label,omitempty"`                  // Copyright holder and year(s)
	Structure              []StructureMapStructure `json:"structure,omitempty" bson:"structure,omitempty"`                             // Structure Definition used by this map
	Import                 []string                `json:"import,omitempty" bson:"import,omitempty"`                                   // Other maps used by this map (canonical URLs)
	Const                  []StructureMapConst     `json:"const,omitempty" bson:"const,omitempty"`                                     // Definition of the constant value used in the map rules
	Group                  []StructureMapGroup     `json:"group" bson:"group"`                                                         // Named sections for reader convenience
}

func (r *StructureMap) Validate() error {
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
	for i, item := range r.Structure {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Structure[%d]: %w", i, err)
		}
	}
	for i, item := range r.Const {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Const[%d]: %w", i, err)
		}
	}
	if len(r.Group) < 1 {
		return fmt.Errorf("field 'Group' must have at least 1 elements")
	}
	for i, item := range r.Group {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Group[%d]: %w", i, err)
		}
	}
	return nil
}

type StructureMapStructure struct {
	Id            *string `json:"id,omitempty" bson:"id,omitempty"`                       // Unique id for inter-element referencing
	Url           string  `json:"url" bson:"url"`                                         // Canonical reference to structure definition
	Mode          string  `json:"mode" bson:"mode"`                                       // source | queried | target | produced
	Alias         *string `json:"alias,omitempty" bson:"alias,omitempty"`                 // Name for type in this map
	Documentation *string `json:"documentation,omitempty" bson:"documentation,omitempty"` // Documentation on use of structure
}

func (r *StructureMapStructure) Validate() error {
	var emptyString string
	if r.Url == emptyString {
		return fmt.Errorf("field 'Url' is required")
	}
	if r.Mode == emptyString {
		return fmt.Errorf("field 'Mode' is required")
	}
	return nil
}

type StructureMapConst struct {
	Id    *string `json:"id,omitempty" bson:"id,omitempty"`       // Unique id for inter-element referencing
	Name  *string `json:"name,omitempty" bson:"name,omitempty"`   // Constant name
	Value *string `json:"value,omitempty" bson:"value,omitempty"` // FHIRPath exression - value of the constant
}

func (r *StructureMapConst) Validate() error {
	return nil
}

type StructureMapGroup struct {
	Id            *string                  `json:"id,omitempty" bson:"id,omitempty"`                       // Unique id for inter-element referencing
	Name          string                   `json:"name" bson:"name"`                                       // Human-readable label
	Extends       *string                  `json:"extends,omitempty" bson:"extends,omitempty"`             // Another group that this group adds rules to
	TypeMode      *string                  `json:"typeMode,omitempty" bson:"type_mode,omitempty"`          // types | type-and-types
	Documentation *string                  `json:"documentation,omitempty" bson:"documentation,omitempty"` // Additional description/explanation for group
	Input         []StructureMapGroupInput `json:"input" bson:"input"`                                     // Named instance provided when invoking the map
	Rule          []StructureMapGroupRule  `json:"rule,omitempty" bson:"rule,omitempty"`                   // Transform Rule from source to target
}

func (r *StructureMapGroup) Validate() error {
	var emptyString string
	if r.Name == emptyString {
		return fmt.Errorf("field 'Name' is required")
	}
	if len(r.Input) < 1 {
		return fmt.Errorf("field 'Input' must have at least 1 elements")
	}
	for i, item := range r.Input {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Input[%d]: %w", i, err)
		}
	}
	for i, item := range r.Rule {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Rule[%d]: %w", i, err)
		}
	}
	return nil
}

type StructureMapGroupInput struct {
	Id            *string `json:"id,omitempty" bson:"id,omitempty"`                       // Unique id for inter-element referencing
	Name          string  `json:"name" bson:"name"`                                       // Name for this instance of data
	Type          *string `json:"type,omitempty" bson:"type,omitempty"`                   // Type for this instance of data
	Mode          string  `json:"mode" bson:"mode"`                                       // source | target
	Documentation *string `json:"documentation,omitempty" bson:"documentation,omitempty"` // Documentation for this instance of data
}

func (r *StructureMapGroupInput) Validate() error {
	var emptyString string
	if r.Name == emptyString {
		return fmt.Errorf("field 'Name' is required")
	}
	if r.Mode == emptyString {
		return fmt.Errorf("field 'Mode' is required")
	}
	return nil
}

type StructureMapGroupRule struct {
	Id            *string                          `json:"id,omitempty" bson:"id,omitempty"`                       // Unique id for inter-element referencing
	Name          *string                          `json:"name,omitempty" bson:"name,omitempty"`                   // Name of the rule for internal references
	Source        []StructureMapGroupRuleSource    `json:"source" bson:"source"`                                   // Source inputs to the mapping
	Target        []StructureMapGroupRuleTarget    `json:"target,omitempty" bson:"target,omitempty"`               // Content to create because of this mapping rule
	Rule          []StructureMapGroupRule          `json:"rule,omitempty" bson:"rule,omitempty"`                   // Rules contained in this rule
	Dependent     []StructureMapGroupRuleDependent `json:"dependent,omitempty" bson:"dependent,omitempty"`         // Which other rules to apply in the context of this rule
	Documentation *string                          `json:"documentation,omitempty" bson:"documentation,omitempty"` // Documentation for this instance of data
}

func (r *StructureMapGroupRule) Validate() error {
	if len(r.Source) < 1 {
		return fmt.Errorf("field 'Source' must have at least 1 elements")
	}
	for i, item := range r.Source {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Source[%d]: %w", i, err)
		}
	}
	for i, item := range r.Target {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Target[%d]: %w", i, err)
		}
	}
	for i, item := range r.Rule {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Rule[%d]: %w", i, err)
		}
	}
	for i, item := range r.Dependent {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Dependent[%d]: %w", i, err)
		}
	}
	return nil
}

type StructureMapGroupRuleSource struct {
	Id           *string `json:"id,omitempty" bson:"id,omitempty"`                      // Unique id for inter-element referencing
	Context      string  `json:"context" bson:"context"`                                // Type or variable this rule applies to
	Min          *int    `json:"min,omitempty" bson:"min,omitempty"`                    // Specified minimum cardinality
	Max          *string `json:"max,omitempty" bson:"max,omitempty"`                    // Specified maximum cardinality (number or *)
	Type         *string `json:"type,omitempty" bson:"type,omitempty"`                  // Rule only applies if source has this type
	DefaultValue *string `json:"defaultValue,omitempty" bson:"default_value,omitempty"` // Default value if no value exists
	Element      *string `json:"element,omitempty" bson:"element,omitempty"`            // Optional field for this source
	ListMode     *string `json:"listMode,omitempty" bson:"list_mode,omitempty"`         // first | not_first | last | not_last | only_one
	Variable     *string `json:"variable,omitempty" bson:"variable,omitempty"`          // Named context for field, if a field is specified
	Condition    *string `json:"condition,omitempty" bson:"condition,omitempty"`        // FHIRPath expression  - must be true or the rule does not apply
	Check        *string `json:"check,omitempty" bson:"check,omitempty"`                // FHIRPath expression  - must be true or the mapping engine throws an error instead of completing
	LogMessage   *string `json:"logMessage,omitempty" bson:"log_message,omitempty"`     // Message to put in log if source exists (FHIRPath)
}

func (r *StructureMapGroupRuleSource) Validate() error {
	var emptyString string
	if r.Context == emptyString {
		return fmt.Errorf("field 'Context' is required")
	}
	return nil
}

type StructureMapGroupRuleTarget struct {
	Id         *string                                `json:"id,omitempty" bson:"id,omitempty"`                   // Unique id for inter-element referencing
	Context    *string                                `json:"context,omitempty" bson:"context,omitempty"`         // Variable this rule applies to
	Element    *string                                `json:"element,omitempty" bson:"element,omitempty"`         // Field to create in the context
	Variable   *string                                `json:"variable,omitempty" bson:"variable,omitempty"`       // Named context for field, if desired, and a field is specified
	ListMode   []string                               `json:"listMode,omitempty" bson:"list_mode,omitempty"`      // first | share | last | single
	ListRuleId *string                                `json:"listRuleId,omitempty" bson:"list_rule_id,omitempty"` // Internal rule reference for shared list items
	Transform  *string                                `json:"transform,omitempty" bson:"transform,omitempty"`     // create | copy +
	Parameter  []StructureMapGroupRuleTargetParameter `json:"parameter,omitempty" bson:"parameter,omitempty"`     // Parameters to the transform
}

func (r *StructureMapGroupRuleTarget) Validate() error {
	for i, item := range r.Parameter {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Parameter[%d]: %w", i, err)
		}
	}
	return nil
}

type StructureMapGroupRuleTargetParameter struct {
	Id            *string  `json:"id,omitempty" bson:"id,omitempty"`     // Unique id for inter-element referencing
	ValueId       *string  `json:"valueId" bson:"value_id"`              // Parameter value - variable or literal
	ValueString   *string  `json:"valueString" bson:"value_string"`      // Parameter value - variable or literal
	ValueBoolean  *bool    `json:"valueBoolean" bson:"value_boolean"`    // Parameter value - variable or literal
	ValueInteger  *int     `json:"valueInteger" bson:"value_integer"`    // Parameter value - variable or literal
	ValueDecimal  *float64 `json:"valueDecimal" bson:"value_decimal"`    // Parameter value - variable or literal
	ValueDate     *string  `json:"valueDate" bson:"value_date"`          // Parameter value - variable or literal
	ValueTime     *string  `json:"valueTime" bson:"value_time"`          // Parameter value - variable or literal
	ValueDateTime *string  `json:"valueDateTime" bson:"value_date_time"` // Parameter value - variable or literal
}

func (r *StructureMapGroupRuleTargetParameter) Validate() error {
	if r.ValueId == nil {
		return fmt.Errorf("field 'ValueId' is required")
	}
	if r.ValueString == nil {
		return fmt.Errorf("field 'ValueString' is required")
	}
	if r.ValueBoolean == nil {
		return fmt.Errorf("field 'ValueBoolean' is required")
	}
	if r.ValueInteger == nil {
		return fmt.Errorf("field 'ValueInteger' is required")
	}
	if r.ValueDecimal == nil {
		return fmt.Errorf("field 'ValueDecimal' is required")
	}
	if r.ValueDate == nil {
		return fmt.Errorf("field 'ValueDate' is required")
	}
	if r.ValueTime == nil {
		return fmt.Errorf("field 'ValueTime' is required")
	}
	if r.ValueDateTime == nil {
		return fmt.Errorf("field 'ValueDateTime' is required")
	}
	return nil
}

type StructureMapGroupRuleDependent struct {
	Id        *string                                `json:"id,omitempty" bson:"id,omitempty"` // Unique id for inter-element referencing
	Name      string                                 `json:"name" bson:"name"`                 // Name of a rule or group to apply
	Parameter []StructureMapGroupRuleTargetParameter `json:"parameter" bson:"parameter"`       // Parameter to pass to the rule or group
}

func (r *StructureMapGroupRuleDependent) Validate() error {
	var emptyString string
	if r.Name == emptyString {
		return fmt.Errorf("field 'Name' is required")
	}
	if len(r.Parameter) < 1 {
		return fmt.Errorf("field 'Parameter' must have at least 1 elements")
	}
	for i, item := range r.Parameter {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Parameter[%d]: %w", i, err)
		}
	}
	return nil
}
