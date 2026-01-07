package models

import (
	"encoding/json"
	"fmt"
)

// A statement of relationships from one set of concepts to one or more other concepts - either concepts in code systems, or data element/data element concepts, or classes in class models.
type ConceptMap struct {
	Id                     *string                         `json:"id,omitempty" bson:"id,omitempty"`                                           // Logical id of this artifact
	Meta                   *Meta                           `json:"meta,omitempty" bson:"meta,omitempty"`                                       // Metadata about the resource
	ImplicitRules          *string                         `json:"implicitRules,omitempty" bson:"implicit_rules,omitempty"`                    // A set of rules under which this content was created
	Language               *string                         `json:"language,omitempty" bson:"language,omitempty"`                               // Language of the resource content
	Text                   *Narrative                      `json:"text,omitempty" bson:"text,omitempty"`                                       // Text summary of the resource, for human interpretation
	Contained              []json.RawMessage               `json:"contained,omitempty" bson:"contained,omitempty"`                             // Contained, inline Resources
	Url                    *string                         `json:"url,omitempty" bson:"url,omitempty"`                                         // Canonical identifier for this concept map, represented as a URI (globally unique)
	Identifier             []Identifier                    `json:"identifier,omitempty" bson:"identifier,omitempty"`                           // Additional identifier for the concept map
	Version                *string                         `json:"version,omitempty" bson:"version,omitempty"`                                 // Business version of the concept map
	VersionAlgorithmString *string                         `json:"versionAlgorithmString,omitempty" bson:"version_algorithm_string,omitempty"` // How to compare versions
	VersionAlgorithmCoding *Coding                         `json:"versionAlgorithmCoding,omitempty" bson:"version_algorithm_coding,omitempty"` // How to compare versions
	Name                   *string                         `json:"name,omitempty" bson:"name,omitempty"`                                       // Name for this concept map (computer friendly)
	Title                  *string                         `json:"title,omitempty" bson:"title,omitempty"`                                     // Name for this concept map (human friendly)
	Status                 string                          `json:"status" bson:"status"`                                                       // draft | active | retired | unknown
	Experimental           bool                            `json:"experimental,omitempty" bson:"experimental,omitempty"`                       // For testing only - never for real usage
	Date                   *string                         `json:"date,omitempty" bson:"date,omitempty"`                                       // Date last changed
	Publisher              *string                         `json:"publisher,omitempty" bson:"publisher,omitempty"`                             // Name of the publisher/steward (organization or individual)
	Contact                []ContactDetail                 `json:"contact,omitempty" bson:"contact,omitempty"`                                 // Contact details for the publisher
	Description            *string                         `json:"description,omitempty" bson:"description,omitempty"`                         // Natural language description of the concept map
	UseContext             []UsageContext                  `json:"useContext,omitempty" bson:"use_context,omitempty"`                          // The context that the content is intended to support
	Jurisdiction           []CodeableConcept               `json:"jurisdiction,omitempty" bson:"jurisdiction,omitempty"`                       // Jurisdiction of the authority that maintains the concept map (if applicable)
	Purpose                *string                         `json:"purpose,omitempty" bson:"purpose,omitempty"`                                 // Why this concept map is defined
	Copyright              *string                         `json:"copyright,omitempty" bson:"copyright,omitempty"`                             // Notice about intellectual property ownership, can include restrictions on use
	CopyrightLabel         *string                         `json:"copyrightLabel,omitempty" bson:"copyright_label,omitempty"`                  // Copyright holder and year(s)
	ApprovalDate           *string                         `json:"approvalDate,omitempty" bson:"approval_date,omitempty"`                      // When the ConceptMap was approved by publisher
	LastReviewDate         *string                         `json:"lastReviewDate,omitempty" bson:"last_review_date,omitempty"`                 // When the ConceptMap was last reviewed by the publisher
	EffectivePeriod        *Period                         `json:"effectivePeriod,omitempty" bson:"effective_period,omitempty"`                // When the ConceptMap is expected to be used
	Topic                  []CodeableConcept               `json:"topic,omitempty" bson:"topic,omitempty"`                                     // E.g. Education, Treatment, Assessment, etc
	Author                 []ContactDetail                 `json:"author,omitempty" bson:"author,omitempty"`                                   // Who authored the ConceptMap
	Editor                 []ContactDetail                 `json:"editor,omitempty" bson:"editor,omitempty"`                                   // Who edited the ConceptMap
	Reviewer               []ContactDetail                 `json:"reviewer,omitempty" bson:"reviewer,omitempty"`                               // Who reviewed the ConceptMap
	Endorser               []ContactDetail                 `json:"endorser,omitempty" bson:"endorser,omitempty"`                               // Who endorsed the ConceptMap
	RelatedArtifact        []RelatedArtifact               `json:"relatedArtifact,omitempty" bson:"related_artifact,omitempty"`                // Additional documentation, citations, etc
	Property               []ConceptMapProperty            `json:"property,omitempty" bson:"property,omitempty"`                               // Additional properties of the mapping
	AdditionalAttribute    []ConceptMapAdditionalAttribute `json:"additionalAttribute,omitempty" bson:"additional_attribute,omitempty"`        // Definition of an additional attribute to act as a data source or target
	SourceScopeUri         *string                         `json:"sourceScopeUri,omitempty" bson:"source_scope_uri,omitempty"`                 // The source value set that contains the concepts that are being mapped
	SourceScopeCanonical   *string                         `json:"sourceScopeCanonical,omitempty" bson:"source_scope_canonical,omitempty"`     // The source value set that contains the concepts that are being mapped
	TargetScopeUri         *string                         `json:"targetScopeUri,omitempty" bson:"target_scope_uri,omitempty"`                 // The target value set which provides context for the mappings
	TargetScopeCanonical   *string                         `json:"targetScopeCanonical,omitempty" bson:"target_scope_canonical,omitempty"`     // The target value set which provides context for the mappings
	Group                  []ConceptMapGroup               `json:"group,omitempty" bson:"group,omitempty"`                                     // Same source and target systems
}

func (r *ConceptMap) Validate() error {
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
	for i, item := range r.Property {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Property[%d]: %w", i, err)
		}
	}
	for i, item := range r.AdditionalAttribute {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("AdditionalAttribute[%d]: %w", i, err)
		}
	}
	for i, item := range r.Group {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Group[%d]: %w", i, err)
		}
	}
	return nil
}

type ConceptMapGroupElement struct {
	Id       *string                        `json:"id,omitempty" bson:"id,omitempty"`              // Unique id for inter-element referencing
	Code     *string                        `json:"code,omitempty" bson:"code,omitempty"`          // Identifies element being mapped
	Display  *string                        `json:"display,omitempty" bson:"display,omitempty"`    // Display for the code
	ValueSet *string                        `json:"valueSet,omitempty" bson:"value_set,omitempty"` // Identifies the set of concepts being mapped
	NoMap    bool                           `json:"noMap,omitempty" bson:"no_map,omitempty"`       // No mapping to a target concept for this source concept
	Comment  *string                        `json:"comment,omitempty" bson:"comment,omitempty"`    // Comments related to the mapping of the source element
	Target   []ConceptMapGroupElementTarget `json:"target,omitempty" bson:"target,omitempty"`      // Concept in target system for element
}

func (r *ConceptMapGroupElement) Validate() error {
	for i, item := range r.Target {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Target[%d]: %w", i, err)
		}
	}
	return nil
}

type ConceptMapProperty struct {
	Id          *string `json:"id,omitempty" bson:"id,omitempty"`                   // Unique id for inter-element referencing
	Code        string  `json:"code" bson:"code"`                                   // Identifies the property on the mappings, and when referred to in the $translate operation
	Uri         *string `json:"uri,omitempty" bson:"uri,omitempty"`                 // Formal identifier for the property
	Description *string `json:"description,omitempty" bson:"description,omitempty"` // Why the property is defined, and/or what it conveys
	Type        string  `json:"type" bson:"type"`                                   // Coding | string | integer | boolean | dateTime | decimal | code
	System      *string `json:"system,omitempty" bson:"system,omitempty"`           // The CodeSystem from which code values come
}

func (r *ConceptMapProperty) Validate() error {
	var emptyString string
	if r.Code == emptyString {
		return fmt.Errorf("field 'Code' is required")
	}
	if r.Type == emptyString {
		return fmt.Errorf("field 'Type' is required")
	}
	return nil
}

type ConceptMapGroup struct {
	Id       *string                  `json:"id,omitempty" bson:"id,omitempty"`             // Unique id for inter-element referencing
	Source   *string                  `json:"source,omitempty" bson:"source,omitempty"`     // Source system where concepts to be mapped are defined
	Target   *string                  `json:"target,omitempty" bson:"target,omitempty"`     // Target system that the concepts are to be mapped to
	Element  []ConceptMapGroupElement `json:"element" bson:"element"`                       // Mappings for a concept from the source set
	Unmapped *ConceptMapGroupUnmapped `json:"unmapped,omitempty" bson:"unmapped,omitempty"` // What to do when there is no mapping target for the source concept and ConceptMap.group.element.noMap is not true
}

func (r *ConceptMapGroup) Validate() error {
	if len(r.Element) < 1 {
		return fmt.Errorf("field 'Element' must have at least 1 elements")
	}
	for i, item := range r.Element {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Element[%d]: %w", i, err)
		}
	}
	if r.Unmapped != nil {
		if err := r.Unmapped.Validate(); err != nil {
			return fmt.Errorf("Unmapped: %w", err)
		}
	}
	return nil
}

type ConceptMapGroupElementTarget struct {
	Id           *string                                 `json:"id,omitempty" bson:"id,omitempty"`                // Unique id for inter-element referencing
	Code         *string                                 `json:"code,omitempty" bson:"code,omitempty"`            // Code that identifies the target element
	Display      *string                                 `json:"display,omitempty" bson:"display,omitempty"`      // Display for the code
	ValueSet     *string                                 `json:"valueSet,omitempty" bson:"value_set,omitempty"`   // Identifies the set of target concepts
	Relationship string                                  `json:"relationship" bson:"relationship"`                // related-to | equivalent | source-is-narrower-than-target | source-is-broader-than-target | not-related-to
	Comment      *string                                 `json:"comment,omitempty" bson:"comment,omitempty"`      // Comments related to the mapping to the target element
	Property     []ConceptMapGroupElementTargetProperty  `json:"property,omitempty" bson:"property,omitempty"`    // Property value for the source -> target mapping
	DependsOn    []ConceptMapGroupElementTargetDependsOn `json:"dependsOn,omitempty" bson:"depends_on,omitempty"` // Other properties required for this mapping
	Product      []ConceptMapGroupElementTargetDependsOn `json:"product,omitempty" bson:"product,omitempty"`      // Other data elements that this mapping also produces
}

func (r *ConceptMapGroupElementTarget) Validate() error {
	var emptyString string
	if r.Relationship == emptyString {
		return fmt.Errorf("field 'Relationship' is required")
	}
	for i, item := range r.Property {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Property[%d]: %w", i, err)
		}
	}
	for i, item := range r.DependsOn {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("DependsOn[%d]: %w", i, err)
		}
	}
	for i, item := range r.Product {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Product[%d]: %w", i, err)
		}
	}
	return nil
}

type ConceptMapGroupElementTargetProperty struct {
	Id            *string  `json:"id,omitempty" bson:"id,omitempty"`     // Unique id for inter-element referencing
	Code          string   `json:"code" bson:"code"`                     // Reference to ConceptMap.property.code
	ValueCoding   *Coding  `json:"valueCoding" bson:"value_coding"`      // Value of the property for this concept
	ValueString   *string  `json:"valueString" bson:"value_string"`      // Value of the property for this concept
	ValueInteger  *int     `json:"valueInteger" bson:"value_integer"`    // Value of the property for this concept
	ValueBoolean  *bool    `json:"valueBoolean" bson:"value_boolean"`    // Value of the property for this concept
	ValueDateTime *string  `json:"valueDateTime" bson:"value_date_time"` // Value of the property for this concept
	ValueDecimal  *float64 `json:"valueDecimal" bson:"value_decimal"`    // Value of the property for this concept
	ValueCode     *string  `json:"valueCode" bson:"value_code"`          // Value of the property for this concept
}

func (r *ConceptMapGroupElementTargetProperty) Validate() error {
	var emptyString string
	if r.Code == emptyString {
		return fmt.Errorf("field 'Code' is required")
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
	if r.ValueCode == nil {
		return fmt.Errorf("field 'ValueCode' is required")
	}
	return nil
}

type ConceptMapGroupElementTargetDependsOn struct {
	Id            *string   `json:"id,omitempty" bson:"id,omitempty"`                        // Unique id for inter-element referencing
	Attribute     string    `json:"attribute" bson:"attribute"`                              // A reference to a mapping attribute defined in ConceptMap.additionalAttribute
	ValueCode     *string   `json:"valueCode,omitempty" bson:"value_code,omitempty"`         // Value of the referenced data element
	ValueCoding   *Coding   `json:"valueCoding,omitempty" bson:"value_coding,omitempty"`     // Value of the referenced data element
	ValueString   *string   `json:"valueString,omitempty" bson:"value_string,omitempty"`     // Value of the referenced data element
	ValueBoolean  *bool     `json:"valueBoolean,omitempty" bson:"value_boolean,omitempty"`   // Value of the referenced data element
	ValueQuantity *Quantity `json:"valueQuantity,omitempty" bson:"value_quantity,omitempty"` // Value of the referenced data element
	ValueSet      *string   `json:"valueSet,omitempty" bson:"value_set,omitempty"`           // The mapping depends on a data element with a value from this value set
}

func (r *ConceptMapGroupElementTargetDependsOn) Validate() error {
	var emptyString string
	if r.Attribute == emptyString {
		return fmt.Errorf("field 'Attribute' is required")
	}
	if r.ValueCoding != nil {
		if err := r.ValueCoding.Validate(); err != nil {
			return fmt.Errorf("ValueCoding: %w", err)
		}
	}
	if r.ValueQuantity != nil {
		if err := r.ValueQuantity.Validate(); err != nil {
			return fmt.Errorf("ValueQuantity: %w", err)
		}
	}
	return nil
}

type ConceptMapGroupUnmapped struct {
	Id           *string `json:"id,omitempty" bson:"id,omitempty"`                     // Unique id for inter-element referencing
	Mode         string  `json:"mode" bson:"mode"`                                     // use-source-code | fixed | other-map
	Code         *string `json:"code,omitempty" bson:"code,omitempty"`                 // Fixed code when mode = fixed
	Display      *string `json:"display,omitempty" bson:"display,omitempty"`           // Display for the code
	Comment      *string `json:"comment,omitempty" bson:"comment,omitempty"`           // Comments related to the choice of how to handle unmapped elements
	ValueSet     *string `json:"valueSet,omitempty" bson:"value_set,omitempty"`        // Fixed code set when mode = fixed
	Relationship *string `json:"relationship,omitempty" bson:"relationship,omitempty"` // related-to | equivalent | source-is-narrower-than-target | source-is-broader-than-target | not-related-to
	OtherMap     *string `json:"otherMap,omitempty" bson:"other_map,omitempty"`        // canonical reference to an additional ConceptMap to use for mapping if the source concept is unmapped
}

func (r *ConceptMapGroupUnmapped) Validate() error {
	var emptyString string
	if r.Mode == emptyString {
		return fmt.Errorf("field 'Mode' is required")
	}
	return nil
}

type ConceptMapAdditionalAttribute struct {
	Id          *string `json:"id,omitempty" bson:"id,omitempty"`                   // Unique id for inter-element referencing
	Code        string  `json:"code" bson:"code"`                                   // Identifies this additional attribute through this resource
	Uri         *string `json:"uri,omitempty" bson:"uri,omitempty"`                 // Formal identifier for the data element referred to in this attribute
	Description *string `json:"description,omitempty" bson:"description,omitempty"` // Why the additional attribute is defined, and/or what the data element it refers to is
	Type        string  `json:"type" bson:"type"`                                   // code | Coding | string | boolean | Quantity
}

func (r *ConceptMapAdditionalAttribute) Validate() error {
	var emptyString string
	if r.Code == emptyString {
		return fmt.Errorf("field 'Code' is required")
	}
	if r.Type == emptyString {
		return fmt.Errorf("field 'Type' is required")
	}
	return nil
}
