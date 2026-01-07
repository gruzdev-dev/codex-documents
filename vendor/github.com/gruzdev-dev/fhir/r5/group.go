package models

import (
	"encoding/json"
	"fmt"
)

// Represents a defined collection of entities that may be discussed or acted upon collectively but which are not typically expected to act collectively*. These collections are also not typically formally or legally recognized.\r\n\r\n*NOTE: Group may be used to define families or households, which in some circumstances may act collectively or have a degree of legal or formal recognition. This should be considered an exception. When Group is used for types of entities other than Patient or RelatedPerson, the expectation remains that the Group will not act collectively or have formal recognition - use Organization if these behaviors are needed. See more discussion [below](group.html#group-usage)
type Group struct {
	ResourceType           string                `json:"resourceType" bson:"resource_type"`                                          // Type of resource
	Id                     *string               `json:"id,omitempty" bson:"id,omitempty"`                                           // Logical id of this artifact
	Meta                   *Meta                 `json:"meta,omitempty" bson:"meta,omitempty"`                                       // Metadata about the resource
	ImplicitRules          *string               `json:"implicitRules,omitempty" bson:"implicit_rules,omitempty"`                    // A set of rules under which this content was created
	Language               *string               `json:"language,omitempty" bson:"language,omitempty"`                               // Language of the resource content
	Text                   *Narrative            `json:"text,omitempty" bson:"text,omitempty"`                                       // Text summary of the resource, for human interpretation
	Contained              []json.RawMessage     `json:"contained,omitempty" bson:"contained,omitempty"`                             // Contained, inline Resources
	Url                    *string               `json:"url,omitempty" bson:"url,omitempty"`                                         // Canonical identifier for this Group, represented as an absolute URI (globally unique)
	Identifier             []Identifier          `json:"identifier,omitempty" bson:"identifier,omitempty"`                           // Business Identifier for this Group
	Version                *string               `json:"version,omitempty" bson:"version,omitempty"`                                 // Business version of the Group
	VersionAlgorithmString *string               `json:"versionAlgorithmString,omitempty" bson:"version_algorithm_string,omitempty"` // How to compare versions
	VersionAlgorithmCoding *Coding               `json:"versionAlgorithmCoding,omitempty" bson:"version_algorithm_coding,omitempty"` // How to compare versions
	Name                   *string               `json:"name,omitempty" bson:"name,omitempty"`                                       // Label for Group
	Title                  *string               `json:"title,omitempty" bson:"title,omitempty"`                                     // Name for this Group (human friendly)
	Status                 *string               `json:"status,omitempty" bson:"status,omitempty"`                                   // draft | active | retired | unknown
	Experimental           bool                  `json:"experimental,omitempty" bson:"experimental,omitempty"`                       // For testing only - never for real usage
	Date                   *string               `json:"date,omitempty" bson:"date,omitempty"`                                       // Date last changed
	Publisher              *string               `json:"publisher,omitempty" bson:"publisher,omitempty"`                             // Name of the publisher/steward (organization or individual)
	Contact                []ContactDetail       `json:"contact,omitempty" bson:"contact,omitempty"`                                 // Contact details for the publisher
	Description            *string               `json:"description,omitempty" bson:"description,omitempty"`                         // Natural language description of the group
	UseContext             []UsageContext        `json:"useContext,omitempty" bson:"use_context,omitempty"`                          // The context that the content is intended to support
	Purpose                *string               `json:"purpose,omitempty" bson:"purpose,omitempty"`                                 // Why this Group is defined
	Copyright              *string               `json:"copyright,omitempty" bson:"copyright,omitempty"`                             // Notice about intellectual property ownership, can include restrictions on use
	CopyrightLabel         *string               `json:"copyrightLabel,omitempty" bson:"copyright_label,omitempty"`                  // Copyright holder and year(s)
	Type                   *string               `json:"type,omitempty" bson:"type,omitempty"`                                       // person | animal | practitioner | device | careteam | healthcareservice | location | organization | relatedperson | specimen | medication | medicinalproductdefinition | substance | substancedefinition | biologicallyDerivedProduct | nutritionProduct
	Membership             string                `json:"membership" bson:"membership"`                                               // definitional | conceptual | enumerated
	Code                   *CodeableConcept      `json:"code,omitempty" bson:"code,omitempty"`                                       // Use of the Group (and by implication, kind of members)
	Quantity               *int                  `json:"quantity,omitempty" bson:"quantity,omitempty"`                               // Number of members
	ManagingEntity         *Reference            `json:"managingEntity,omitempty" bson:"managing_entity,omitempty"`                  // Entity that is the custodian of the Group's definition
	CombinationMethod      *string               `json:"combinationMethod,omitempty" bson:"combination_method,omitempty"`            // all-of | any-of | at-least | at-most | except-subset
	CombinationThreshold   *int                  `json:"combinationThreshold,omitempty" bson:"combination_threshold,omitempty"`      // Provides the value of "n" when "at-least" or "at-most" codes are used
	Characteristic         []GroupCharacteristic `json:"characteristic,omitempty" bson:"characteristic,omitempty"`                   // Include / Exclude group members by Trait
	Member                 []GroupMember         `json:"member,omitempty" bson:"member,omitempty"`                                   // Who or what is in group
}

func (r *Group) Validate() error {
	if r.ResourceType != "Group" {
		return fmt.Errorf("invalid resourceType: expected 'Group', got '%s'", r.ResourceType)
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
	var emptyString string
	if r.Membership == emptyString {
		return fmt.Errorf("field 'Membership' is required")
	}
	if r.Code != nil {
		if err := r.Code.Validate(); err != nil {
			return fmt.Errorf("Code: %w", err)
		}
	}
	if r.ManagingEntity != nil {
		if err := r.ManagingEntity.Validate(); err != nil {
			return fmt.Errorf("ManagingEntity: %w", err)
		}
	}
	for i, item := range r.Characteristic {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Characteristic[%d]: %w", i, err)
		}
	}
	for i, item := range r.Member {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Member[%d]: %w", i, err)
		}
	}
	return nil
}

type GroupCharacteristic struct {
	Id                   *string           `json:"id,omitempty" bson:"id,omitempty"`                                       // Unique id for inter-element referencing
	Code                 *CodeableConcept  `json:"code" bson:"code"`                                                       // Kind of characteristic
	ValueCodeableConcept *CodeableConcept  `json:"valueCodeableConcept" bson:"value_codeable_concept"`                     // Value held by characteristic
	ValueBoolean         *bool             `json:"valueBoolean" bson:"value_boolean"`                                      // Value held by characteristic
	ValueQuantity        *Quantity         `json:"valueQuantity" bson:"value_quantity"`                                    // Value held by characteristic
	ValueRange           *Range            `json:"valueRange" bson:"value_range"`                                          // Value held by characteristic
	ValueReference       *Reference        `json:"valueReference" bson:"value_reference"`                                  // Value held by characteristic
	ValueUri             *string           `json:"valueUri" bson:"value_uri"`                                              // Value held by characteristic
	ValueExpression      *Expression       `json:"valueExpression" bson:"value_expression"`                                // Value held by characteristic
	Exclude              bool              `json:"exclude" bson:"exclude"`                                                 // Group includes or excludes
	Description          *string           `json:"description,omitempty" bson:"description,omitempty"`                     // Natural language description of the characteristic
	Method               []CodeableConcept `json:"method,omitempty" bson:"method,omitempty"`                               // Method for how the characteristic value was determined
	Formula              *Expression       `json:"formula,omitempty" bson:"formula,omitempty"`                             // Formal algorithm to derive the value
	Determiner           *Reference        `json:"determiner,omitempty" bson:"determiner,omitempty"`                       // Who determines the value
	Offset               *CodeableConcept  `json:"offset,omitempty" bson:"offset,omitempty"`                               // Reference point for comparison
	InstancesUnsignedInt *int              `json:"instancesUnsignedInt,omitempty" bson:"instances_unsigned_int,omitempty"` // Number of occurrences meeting the characteristic
	InstancesRange       *Range            `json:"instancesRange,omitempty" bson:"instances_range,omitempty"`              // Number of occurrences meeting the characteristic
	DurationDuration     *Duration         `json:"durationDuration,omitempty" bson:"duration_duration,omitempty"`          // Length of time in which the characteristic is met
	DurationRange        *Range            `json:"durationRange,omitempty" bson:"duration_range,omitempty"`                // Length of time in which the characteristic is met
	Period               *Period           `json:"period,omitempty" bson:"period,omitempty"`                               // Period over which characteristic is tested
	Timing               []RelativeTime    `json:"timing,omitempty" bson:"timing,omitempty"`                               // Timing in which the characteristic is determined
}

func (r *GroupCharacteristic) Validate() error {
	if r.Code == nil {
		return fmt.Errorf("field 'Code' is required")
	}
	if r.Code != nil {
		if err := r.Code.Validate(); err != nil {
			return fmt.Errorf("Code: %w", err)
		}
	}
	if r.ValueCodeableConcept == nil {
		return fmt.Errorf("field 'ValueCodeableConcept' is required")
	}
	if r.ValueCodeableConcept != nil {
		if err := r.ValueCodeableConcept.Validate(); err != nil {
			return fmt.Errorf("ValueCodeableConcept: %w", err)
		}
	}
	if r.ValueBoolean == nil {
		return fmt.Errorf("field 'ValueBoolean' is required")
	}
	if r.ValueQuantity == nil {
		return fmt.Errorf("field 'ValueQuantity' is required")
	}
	if r.ValueQuantity != nil {
		if err := r.ValueQuantity.Validate(); err != nil {
			return fmt.Errorf("ValueQuantity: %w", err)
		}
	}
	if r.ValueRange == nil {
		return fmt.Errorf("field 'ValueRange' is required")
	}
	if r.ValueRange != nil {
		if err := r.ValueRange.Validate(); err != nil {
			return fmt.Errorf("ValueRange: %w", err)
		}
	}
	if r.ValueReference == nil {
		return fmt.Errorf("field 'ValueReference' is required")
	}
	if r.ValueReference != nil {
		if err := r.ValueReference.Validate(); err != nil {
			return fmt.Errorf("ValueReference: %w", err)
		}
	}
	if r.ValueUri == nil {
		return fmt.Errorf("field 'ValueUri' is required")
	}
	if r.ValueExpression == nil {
		return fmt.Errorf("field 'ValueExpression' is required")
	}
	if r.ValueExpression != nil {
		if err := r.ValueExpression.Validate(); err != nil {
			return fmt.Errorf("ValueExpression: %w", err)
		}
	}
	for i, item := range r.Method {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Method[%d]: %w", i, err)
		}
	}
	if r.Formula != nil {
		if err := r.Formula.Validate(); err != nil {
			return fmt.Errorf("Formula: %w", err)
		}
	}
	if r.Determiner != nil {
		if err := r.Determiner.Validate(); err != nil {
			return fmt.Errorf("Determiner: %w", err)
		}
	}
	if r.Offset != nil {
		if err := r.Offset.Validate(); err != nil {
			return fmt.Errorf("Offset: %w", err)
		}
	}
	if r.InstancesRange != nil {
		if err := r.InstancesRange.Validate(); err != nil {
			return fmt.Errorf("InstancesRange: %w", err)
		}
	}
	if r.DurationDuration != nil {
		if err := r.DurationDuration.Validate(); err != nil {
			return fmt.Errorf("DurationDuration: %w", err)
		}
	}
	if r.DurationRange != nil {
		if err := r.DurationRange.Validate(); err != nil {
			return fmt.Errorf("DurationRange: %w", err)
		}
	}
	if r.Period != nil {
		if err := r.Period.Validate(); err != nil {
			return fmt.Errorf("Period: %w", err)
		}
	}
	for i, item := range r.Timing {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Timing[%d]: %w", i, err)
		}
	}
	return nil
}

type GroupMember struct {
	Id          *string           `json:"id,omitempty" bson:"id,omitempty"`                   // Unique id for inter-element referencing
	Entity      *Reference        `json:"entity" bson:"entity"`                               // Reference to the group member
	Involvement []CodeableConcept `json:"involvement,omitempty" bson:"involvement,omitempty"` // Code that describes how user is part of the group
	Period      *Period           `json:"period,omitempty" bson:"period,omitempty"`           // Period member belonged to the group
	Inactive    bool              `json:"inactive,omitempty" bson:"inactive,omitempty"`       // If member is no longer in group
}

func (r *GroupMember) Validate() error {
	if r.Entity == nil {
		return fmt.Errorf("field 'Entity' is required")
	}
	if r.Entity != nil {
		if err := r.Entity.Validate(); err != nil {
			return fmt.Errorf("Entity: %w", err)
		}
	}
	for i, item := range r.Involvement {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Involvement[%d]: %w", i, err)
		}
	}
	if r.Period != nil {
		if err := r.Period.Validate(); err != nil {
			return fmt.Errorf("Period: %w", err)
		}
	}
	return nil
}
