package models

import (
	"encoding/json"
	"fmt"
)

// A formal computable definition of an operation (on the RESTful interface) or a named query (using the search interaction).
type OperationDefinition struct {
	ResourceType           string                         `json:"resourceType" bson:"resource_type"`                                          // Type of resource
	Id                     *string                        `json:"id,omitempty" bson:"id,omitempty"`                                           // Logical id of this artifact
	Meta                   *Meta                          `json:"meta,omitempty" bson:"meta,omitempty"`                                       // Metadata about the resource
	ImplicitRules          *string                        `json:"implicitRules,omitempty" bson:"implicit_rules,omitempty"`                    // A set of rules under which this content was created
	Language               *string                        `json:"language,omitempty" bson:"language,omitempty"`                               // Language of the resource content
	Text                   *Narrative                     `json:"text,omitempty" bson:"text,omitempty"`                                       // Text summary of the resource, for human interpretation
	Contained              []json.RawMessage              `json:"contained,omitempty" bson:"contained,omitempty"`                             // Contained, inline Resources
	Url                    *string                        `json:"url,omitempty" bson:"url,omitempty"`                                         // Canonical identifier for this operation definition, represented as an absolute URI (globally unique)
	Identifier             []Identifier                   `json:"identifier,omitempty" bson:"identifier,omitempty"`                           // Additional identifier for the implementation guide (business identifier)
	Version                *string                        `json:"version,omitempty" bson:"version,omitempty"`                                 // Business version of the operation definition
	VersionAlgorithmString *string                        `json:"versionAlgorithmString,omitempty" bson:"version_algorithm_string,omitempty"` // How to compare versions
	VersionAlgorithmCoding *Coding                        `json:"versionAlgorithmCoding,omitempty" bson:"version_algorithm_coding,omitempty"` // How to compare versions
	Name                   string                         `json:"name" bson:"name"`                                                           // Name for this operation definition (computer friendly)
	Title                  *string                        `json:"title,omitempty" bson:"title,omitempty"`                                     // Name for this operation definition (human friendly)
	Status                 string                         `json:"status" bson:"status"`                                                       // draft | active | retired | unknown
	Kind                   string                         `json:"kind" bson:"kind"`                                                           // operation | query
	Experimental           bool                           `json:"experimental,omitempty" bson:"experimental,omitempty"`                       // For testing only - never for real usage
	Date                   *string                        `json:"date,omitempty" bson:"date,omitempty"`                                       // Date last changed
	Publisher              *string                        `json:"publisher,omitempty" bson:"publisher,omitempty"`                             // Name of the publisher/steward (organization or individual)
	Contact                []ContactDetail                `json:"contact,omitempty" bson:"contact,omitempty"`                                 // Contact details for the publisher
	Description            *string                        `json:"description,omitempty" bson:"description,omitempty"`                         // Natural language description of the operation definition
	UseContext             []UsageContext                 `json:"useContext,omitempty" bson:"use_context,omitempty"`                          // The context that the content is intended to support
	Jurisdiction           []CodeableConcept              `json:"jurisdiction,omitempty" bson:"jurisdiction,omitempty"`                       // Jurisdiction of the authority that maintains the operation definition (if applicable)
	Purpose                *string                        `json:"purpose,omitempty" bson:"purpose,omitempty"`                                 // Why this operation definition is defined
	Copyright              *string                        `json:"copyright,omitempty" bson:"copyright,omitempty"`                             // Notice about intellectual property ownership, can include restrictions on use
	CopyrightLabel         *string                        `json:"copyrightLabel,omitempty" bson:"copyright_label,omitempty"`                  // Copyright holder and year(s)
	AffectsState           bool                           `json:"affectsState,omitempty" bson:"affects_state,omitempty"`                      // Whether content is changed by the operation
	Synchronicity          *string                        `json:"synchronicity,omitempty" bson:"synchronicity,omitempty"`                     // synchronous | asynchronous | either
	Code                   string                         `json:"code" bson:"code"`                                                           // Recommended name for operation in search url
	Comment                *string                        `json:"comment,omitempty" bson:"comment,omitempty"`                                 // Additional information about use
	Base                   *string                        `json:"base,omitempty" bson:"base,omitempty"`                                       // Marks this as a profile of the base
	Resource               []string                       `json:"resource,omitempty" bson:"resource,omitempty"`                               // Types this operation applies to
	System                 bool                           `json:"system" bson:"system"`                                                       // Invoke at the system level?
	Type                   bool                           `json:"type" bson:"type"`                                                           // Invoke at the type level?
	Instance               bool                           `json:"instance" bson:"instance"`                                                   // Invoke on an instance?
	InputProfile           *string                        `json:"inputProfile,omitempty" bson:"input_profile,omitempty"`                      // Validation information for in parameters
	OutputProfile          *string                        `json:"outputProfile,omitempty" bson:"output_profile,omitempty"`                    // Validation information for out parameters
	Parameter              []OperationDefinitionParameter `json:"parameter,omitempty" bson:"parameter,omitempty"`                             // Parameters for the operation/query
	Overload               []OperationDefinitionOverload  `json:"overload,omitempty" bson:"overload,omitempty"`                               // Define overloaded variants for when  generating code
}

func (r *OperationDefinition) Validate() error {
	if r.ResourceType != "OperationDefinition" {
		return fmt.Errorf("invalid resourceType: expected 'OperationDefinition', got '%s'", r.ResourceType)
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
	if r.Name == emptyString {
		return fmt.Errorf("field 'Name' is required")
	}
	if r.Status == emptyString {
		return fmt.Errorf("field 'Status' is required")
	}
	if r.Kind == emptyString {
		return fmt.Errorf("field 'Kind' is required")
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
	if r.Code == emptyString {
		return fmt.Errorf("field 'Code' is required")
	}
	for i, item := range r.Parameter {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Parameter[%d]: %w", i, err)
		}
	}
	for i, item := range r.Overload {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Overload[%d]: %w", i, err)
		}
	}
	return nil
}

type OperationDefinitionParameterReferencedFrom struct {
	Id       *string `json:"id,omitempty" bson:"id,omitempty"`              // Unique id for inter-element referencing
	Source   string  `json:"source" bson:"source"`                          // Referencing parameter
	SourceId *string `json:"sourceId,omitempty" bson:"source_id,omitempty"` // Element id of reference
}

func (r *OperationDefinitionParameterReferencedFrom) Validate() error {
	var emptyString string
	if r.Source == emptyString {
		return fmt.Errorf("field 'Source' is required")
	}
	return nil
}

type OperationDefinitionOverload struct {
	Id            *string  `json:"id,omitempty" bson:"id,omitempty"`                        // Unique id for inter-element referencing
	ParameterName []string `json:"parameterName,omitempty" bson:"parameter_name,omitempty"` // Name of parameter to include in overload
	Comment       *string  `json:"comment,omitempty" bson:"comment,omitempty"`              // Comments to go on overload
}

func (r *OperationDefinitionOverload) Validate() error {
	return nil
}

type OperationDefinitionParameter struct {
	Id             *string                                      `json:"id,omitempty" bson:"id,omitempty"`                          // Unique id for inter-element referencing
	Name           string                                       `json:"name" bson:"name"`                                          // Name in Parameters.parameter.name or in URL
	Use            string                                       `json:"use" bson:"use"`                                            // in | out
	Scope          []string                                     `json:"scope,omitempty" bson:"scope,omitempty"`                    // instance | type | system
	Min            int                                          `json:"min" bson:"min"`                                            // Minimum Cardinality
	Max            string                                       `json:"max" bson:"max"`                                            // Maximum Cardinality (a number or *)
	Documentation  *string                                      `json:"documentation,omitempty" bson:"documentation,omitempty"`    // Description of meaning/use
	Type           *string                                      `json:"type,omitempty" bson:"type,omitempty"`                      // What type this parameter has
	AllowedType    []string                                     `json:"allowedType,omitempty" bson:"allowed_type,omitempty"`       // Allowed sub-type this parameter can have (if type is abstract)
	TargetProfile  []string                                     `json:"targetProfile,omitempty" bson:"target_profile,omitempty"`   // If type is Reference | canonical, allowed targets. If type is 'Resource', then this constrains the allowed resource types
	SearchType     *string                                      `json:"searchType,omitempty" bson:"search_type,omitempty"`         // number | date | string | token | reference | composite | quantity | uri | special | resource
	Binding        *OperationDefinitionParameterBinding         `json:"binding,omitempty" bson:"binding,omitempty"`                // ValueSet details if this is coded
	ReferencedFrom []OperationDefinitionParameterReferencedFrom `json:"referencedFrom,omitempty" bson:"referenced_from,omitempty"` // References to this parameter
	Part           []OperationDefinitionParameter               `json:"part,omitempty" bson:"part,omitempty"`                      // Parts of a nested Parameter
}

func (r *OperationDefinitionParameter) Validate() error {
	var emptyString string
	if r.Name == emptyString {
		return fmt.Errorf("field 'Name' is required")
	}
	if r.Use == emptyString {
		return fmt.Errorf("field 'Use' is required")
	}
	if r.Min == 0 {
		return fmt.Errorf("field 'Min' is required")
	}
	if r.Max == emptyString {
		return fmt.Errorf("field 'Max' is required")
	}
	if r.Binding != nil {
		if err := r.Binding.Validate(); err != nil {
			return fmt.Errorf("Binding: %w", err)
		}
	}
	for i, item := range r.ReferencedFrom {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("ReferencedFrom[%d]: %w", i, err)
		}
	}
	for i, item := range r.Part {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Part[%d]: %w", i, err)
		}
	}
	return nil
}

type OperationDefinitionParameterBinding struct {
	Id       *string `json:"id,omitempty" bson:"id,omitempty"` // Unique id for inter-element referencing
	Strength string  `json:"strength" bson:"strength"`         // required | extensible | preferred | example | descriptive
	ValueSet string  `json:"valueSet" bson:"value_set"`        // Source of value set
}

func (r *OperationDefinitionParameterBinding) Validate() error {
	var emptyString string
	if r.Strength == emptyString {
		return fmt.Errorf("field 'Strength' is required")
	}
	if r.ValueSet == emptyString {
		return fmt.Errorf("field 'ValueSet' is required")
	}
	return nil
}
