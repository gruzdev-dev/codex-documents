package models

import (
	"encoding/json"
	"fmt"
)

// The EvidenceVariable resource describes an element that knowledge (Evidence) is about.
type EvidenceVariable struct {
	ResourceType             string                               `json:"resourceType" bson:"resource_type"`                                              // Type of resource
	Id                       *string                              `json:"id,omitempty" bson:"id,omitempty"`                                               // Logical id of this artifact
	Meta                     *Meta                                `json:"meta,omitempty" bson:"meta,omitempty"`                                           // Metadata about the resource
	ImplicitRules            *string                              `json:"implicitRules,omitempty" bson:"implicit_rules,omitempty"`                        // A set of rules under which this content was created
	Language                 *string                              `json:"language,omitempty" bson:"language,omitempty"`                                   // Language of the resource content
	Text                     *Narrative                           `json:"text,omitempty" bson:"text,omitempty"`                                           // Text summary of the resource, for human interpretation
	Contained                []json.RawMessage                    `json:"contained,omitempty" bson:"contained,omitempty"`                                 // Contained, inline Resources
	Url                      *string                              `json:"url,omitempty" bson:"url,omitempty"`                                             // Canonical identifier for this evidence variable, represented as a URI (globally unique)
	Identifier               []Identifier                         `json:"identifier,omitempty" bson:"identifier,omitempty"`                               // Additional identifier for the evidence variable
	Version                  *string                              `json:"version,omitempty" bson:"version,omitempty"`                                     // Business version of the evidence variable
	VersionAlgorithmString   *string                              `json:"versionAlgorithmString,omitempty" bson:"version_algorithm_string,omitempty"`     // How to compare versions
	VersionAlgorithmCoding   *Coding                              `json:"versionAlgorithmCoding,omitempty" bson:"version_algorithm_coding,omitempty"`     // How to compare versions
	Name                     *string                              `json:"name,omitempty" bson:"name,omitempty"`                                           // Name for this evidence variable (computer friendly)
	Title                    *string                              `json:"title,omitempty" bson:"title,omitempty"`                                         // Name for this evidence variable (human friendly)
	ShortTitle               *string                              `json:"shortTitle,omitempty" bson:"short_title,omitempty"`                              // Title for use in informal contexts
	CiteAs                   *string                              `json:"citeAs,omitempty" bson:"cite_as,omitempty"`                                      // Display of how to cite this EvidenceVariable
	Status                   string                               `json:"status" bson:"status"`                                                           // draft | active | retired | unknown
	Experimental             bool                                 `json:"experimental,omitempty" bson:"experimental,omitempty"`                           // For testing only - never for real usage
	Date                     *string                              `json:"date,omitempty" bson:"date,omitempty"`                                           // Date last changed
	Author                   []ContactDetail                      `json:"author,omitempty" bson:"author,omitempty"`                                       // Who authored the content
	Publisher                *string                              `json:"publisher,omitempty" bson:"publisher,omitempty"`                                 // Name of the publisher/steward (organization or individual)
	Contact                  []ContactDetail                      `json:"contact,omitempty" bson:"contact,omitempty"`                                     // Contact details for the publisher
	Recorder                 []ContactDetail                      `json:"recorder,omitempty" bson:"recorder,omitempty"`                                   // Who entered the data for the evidence variable
	Editor                   []ContactDetail                      `json:"editor,omitempty" bson:"editor,omitempty"`                                       // Who edited the content
	Reviewer                 []ContactDetail                      `json:"reviewer,omitempty" bson:"reviewer,omitempty"`                                   // Who reviewed the content
	Endorser                 []ContactDetail                      `json:"endorser,omitempty" bson:"endorser,omitempty"`                                   // Who endorsed the content
	Description              *string                              `json:"description,omitempty" bson:"description,omitempty"`                             // Natural language description of the evidence variable
	Note                     []Annotation                         `json:"note,omitempty" bson:"note,omitempty"`                                           // Used for footnotes or explanatory notes
	UseContext               []UsageContext                       `json:"useContext,omitempty" bson:"use_context,omitempty"`                              // The context that the content is intended to support
	Purpose                  *string                              `json:"purpose,omitempty" bson:"purpose,omitempty"`                                     // Why this EvidenceVariable is defined
	Copyright                *string                              `json:"copyright,omitempty" bson:"copyright,omitempty"`                                 // Intellectual property ownership, may include restrictions on use
	CopyrightLabel           *string                              `json:"copyrightLabel,omitempty" bson:"copyright_label,omitempty"`                      // Copyright holder and year(s)
	ApprovalDate             *string                              `json:"approvalDate,omitempty" bson:"approval_date,omitempty"`                          // When the resource was approved by publisher
	LastReviewDate           *string                              `json:"lastReviewDate,omitempty" bson:"last_review_date,omitempty"`                     // When the resource was last reviewed by the publisher
	EffectivePeriod          *Period                              `json:"effectivePeriod,omitempty" bson:"effective_period,omitempty"`                    // When the resource is expected to be used
	RelatesTo                []EvidenceVariableRelatesTo          `json:"relatesTo,omitempty" bson:"relates_to,omitempty"`                                // Relationships to other Resources
	Actual                   bool                                 `json:"actual,omitempty" bson:"actual,omitempty"`                                       // Actual or conceptual
	Definition               *CodeableReference                   `json:"definition,omitempty" bson:"definition,omitempty"`                               // The meaning of the evidence variable
	DefinitionModifier       []EvidenceVariableDefinitionModifier `json:"definitionModifier,omitempty" bson:"definition_modifier,omitempty"`              // Further specification of the definition
	Handling                 *CodeableConcept                     `json:"handling,omitempty" bson:"handling,omitempty"`                                   // boolean | continuous | dichotomous | ordinal | polychotomous | time-to-event | not-specified
	Category                 []EvidenceVariableCategory           `json:"category,omitempty" bson:"category,omitempty"`                                   // A grouping for dichotomous, ordinal, or polychotomouos variables
	Conditional              *Expression                          `json:"conditional,omitempty" bson:"conditional,omitempty"`                             // Condition determining whether the data will be collected
	Classifier               []CodeableConcept                    `json:"classifier,omitempty" bson:"classifier,omitempty"`                               // Classification
	DataStorage              []EvidenceVariableDataStorage        `json:"dataStorage,omitempty" bson:"data_storage,omitempty"`                            // How the data element (value of the variable) is found
	Timing                   *RelativeTime                        `json:"timing,omitempty" bson:"timing,omitempty"`                                       // When the variable is observed
	Period                   *Period                              `json:"period,omitempty" bson:"period,omitempty"`                                       // Calendar-based timing when the variable is observed
	Constraint               []EvidenceVariableConstraint         `json:"constraint,omitempty" bson:"constraint,omitempty"`                               // Limit on acceptability of data value
	MissingDataMeaning       []CodeableConcept                    `json:"missingDataMeaning,omitempty" bson:"missing_data_meaning,omitempty"`             // How missing data can be interpreted
	UnacceptableDataHandling []CodeableConcept                    `json:"unacceptableDataHandling,omitempty" bson:"unacceptable_data_handling,omitempty"` // How erroneous data is processed
}

func (r *EvidenceVariable) Validate() error {
	if r.ResourceType != "EvidenceVariable" {
		return fmt.Errorf("invalid resourceType: expected 'EvidenceVariable', got '%s'", r.ResourceType)
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
	for i, item := range r.Author {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Author[%d]: %w", i, err)
		}
	}
	for i, item := range r.Contact {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Contact[%d]: %w", i, err)
		}
	}
	for i, item := range r.Recorder {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Recorder[%d]: %w", i, err)
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
	for i, item := range r.Note {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Note[%d]: %w", i, err)
		}
	}
	for i, item := range r.UseContext {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("UseContext[%d]: %w", i, err)
		}
	}
	if r.EffectivePeriod != nil {
		if err := r.EffectivePeriod.Validate(); err != nil {
			return fmt.Errorf("EffectivePeriod: %w", err)
		}
	}
	for i, item := range r.RelatesTo {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("RelatesTo[%d]: %w", i, err)
		}
	}
	if r.Definition != nil {
		if err := r.Definition.Validate(); err != nil {
			return fmt.Errorf("Definition: %w", err)
		}
	}
	for i, item := range r.DefinitionModifier {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("DefinitionModifier[%d]: %w", i, err)
		}
	}
	if r.Handling != nil {
		if err := r.Handling.Validate(); err != nil {
			return fmt.Errorf("Handling: %w", err)
		}
	}
	for i, item := range r.Category {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Category[%d]: %w", i, err)
		}
	}
	if r.Conditional != nil {
		if err := r.Conditional.Validate(); err != nil {
			return fmt.Errorf("Conditional: %w", err)
		}
	}
	for i, item := range r.Classifier {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Classifier[%d]: %w", i, err)
		}
	}
	for i, item := range r.DataStorage {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("DataStorage[%d]: %w", i, err)
		}
	}
	if r.Timing != nil {
		if err := r.Timing.Validate(); err != nil {
			return fmt.Errorf("Timing: %w", err)
		}
	}
	if r.Period != nil {
		if err := r.Period.Validate(); err != nil {
			return fmt.Errorf("Period: %w", err)
		}
	}
	for i, item := range r.Constraint {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Constraint[%d]: %w", i, err)
		}
	}
	for i, item := range r.MissingDataMeaning {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("MissingDataMeaning[%d]: %w", i, err)
		}
	}
	for i, item := range r.UnacceptableDataHandling {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("UnacceptableDataHandling[%d]: %w", i, err)
		}
	}
	return nil
}

type EvidenceVariableDataStorage struct {
	Id        *string                       `json:"id,omitempty" bson:"id,omitempty"`               // Unique id for inter-element referencing
	Datatype  *CodeableConcept              `json:"datatype,omitempty" bson:"datatype,omitempty"`   // Type of data used to express value of the variable
	Path      *string                       `json:"path,omitempty" bson:"path,omitempty"`           // Where to find the data element in the dataset
	Delimiter *string                       `json:"delimiter,omitempty" bson:"delimiter,omitempty"` // Character(s) separating values in a string-based list
	Component []EvidenceVariableDataStorage `json:"component,omitempty" bson:"component,omitempty"`
}

func (r *EvidenceVariableDataStorage) Validate() error {
	if r.Datatype != nil {
		if err := r.Datatype.Validate(); err != nil {
			return fmt.Errorf("Datatype: %w", err)
		}
	}
	for i, item := range r.Component {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Component[%d]: %w", i, err)
		}
	}
	return nil
}

type EvidenceVariableConstraint struct {
	Id                    *string          `json:"id,omitempty" bson:"id,omitempty"`                                          // Unique id for inter-element referencing
	Conditional           *CodeableConcept `json:"conditional,omitempty" bson:"conditional,omitempty"`                        // Condition determining whether this constraint applies
	MinimumQuantity       *Quantity        `json:"minimumQuantity,omitempty" bson:"minimum_quantity,omitempty"`               // The lowest permissible value of the variable
	MaximumQuantity       *Quantity        `json:"maximumQuantity,omitempty" bson:"maximum_quantity,omitempty"`               // The highest permissible value of the variable
	EarliestDateTime      *string          `json:"earliestDateTime,omitempty" bson:"earliest_date_time,omitempty"`            // The earliest permissible value of the variable
	LatestDateTime        *string          `json:"latestDateTime,omitempty" bson:"latest_date_time,omitempty"`                // The latest permissible value of the variable
	MinimumStringLength   *int             `json:"minimumStringLength,omitempty" bson:"minimum_string_length,omitempty"`      // The lowest number of characters allowed for a value of the variable
	MaximumStringLength   *int             `json:"maximumStringLength,omitempty" bson:"maximum_string_length,omitempty"`      // The highest number of characters allowed for a value of the variable
	Code                  *CodeableConcept `json:"code,omitempty" bson:"code,omitempty"`                                      // Rule for acceptable data values
	Expression            *Expression      `json:"expression,omitempty" bson:"expression,omitempty"`                          // Rule for acceptable data values, as an Expression
	ExpectedValueSet      *Reference       `json:"expectedValueSet,omitempty" bson:"expected_value_set,omitempty"`            // List of anticipated values used to express value of the variable
	ExpectedUnitsValueSet *Reference       `json:"expectedUnitsValueSet,omitempty" bson:"expected_units_value_set,omitempty"` // List of anticipated values used to express units for the value of the variable
	AnyValueAllowed       bool             `json:"anyValueAllowed,omitempty" bson:"any_value_allowed,omitempty"`              // Permissibility of unanticipated value used to express value of the variable
}

func (r *EvidenceVariableConstraint) Validate() error {
	if r.Conditional != nil {
		if err := r.Conditional.Validate(); err != nil {
			return fmt.Errorf("Conditional: %w", err)
		}
	}
	if r.MinimumQuantity != nil {
		if err := r.MinimumQuantity.Validate(); err != nil {
			return fmt.Errorf("MinimumQuantity: %w", err)
		}
	}
	if r.MaximumQuantity != nil {
		if err := r.MaximumQuantity.Validate(); err != nil {
			return fmt.Errorf("MaximumQuantity: %w", err)
		}
	}
	if r.Code != nil {
		if err := r.Code.Validate(); err != nil {
			return fmt.Errorf("Code: %w", err)
		}
	}
	if r.Expression != nil {
		if err := r.Expression.Validate(); err != nil {
			return fmt.Errorf("Expression: %w", err)
		}
	}
	if r.ExpectedValueSet != nil {
		if err := r.ExpectedValueSet.Validate(); err != nil {
			return fmt.Errorf("ExpectedValueSet: %w", err)
		}
	}
	if r.ExpectedUnitsValueSet != nil {
		if err := r.ExpectedUnitsValueSet.Validate(); err != nil {
			return fmt.Errorf("ExpectedUnitsValueSet: %w", err)
		}
	}
	return nil
}

type EvidenceVariableRelatesTo struct {
	Id               *string          `json:"id,omitempty" bson:"id,omitempty"`          // Unique id for inter-element referencing
	Type             *CodeableConcept `json:"type" bson:"type"`                          // documentation | justification | citation | predecessor | successor | derived-from | depends-on | composed-of | part-of | amends | amended-with | appends | appended-with | cites | cited-by | comments-on | comment-in | contains | contained-in | corrects | correction-in | replaces | replaced-with | retracts | retracted-by | signs | similar-to | supports | supported-with | transforms | transformed-into | transformed-with | documents | specification-of | created-with | cite-as | reprint | reprint-of | summarizes
	TargetUri        *string          `json:"targetUri" bson:"target_uri"`               // The artifact that is related to this EvidenceVariable
	TargetAttachment *Attachment      `json:"targetAttachment" bson:"target_attachment"` // The artifact that is related to this EvidenceVariable
	TargetCanonical  *string          `json:"targetCanonical" bson:"target_canonical"`   // The artifact that is related to this EvidenceVariable
	TargetReference  *Reference       `json:"targetReference" bson:"target_reference"`   // The artifact that is related to this EvidenceVariable
	TargetMarkdown   *string          `json:"targetMarkdown" bson:"target_markdown"`     // The artifact that is related to this EvidenceVariable
}

func (r *EvidenceVariableRelatesTo) Validate() error {
	if r.Type == nil {
		return fmt.Errorf("field 'Type' is required")
	}
	if r.Type != nil {
		if err := r.Type.Validate(); err != nil {
			return fmt.Errorf("Type: %w", err)
		}
	}
	if r.TargetUri == nil {
		return fmt.Errorf("field 'TargetUri' is required")
	}
	if r.TargetAttachment == nil {
		return fmt.Errorf("field 'TargetAttachment' is required")
	}
	if r.TargetAttachment != nil {
		if err := r.TargetAttachment.Validate(); err != nil {
			return fmt.Errorf("TargetAttachment: %w", err)
		}
	}
	if r.TargetCanonical == nil {
		return fmt.Errorf("field 'TargetCanonical' is required")
	}
	if r.TargetReference == nil {
		return fmt.Errorf("field 'TargetReference' is required")
	}
	if r.TargetReference != nil {
		if err := r.TargetReference.Validate(); err != nil {
			return fmt.Errorf("TargetReference: %w", err)
		}
	}
	if r.TargetMarkdown == nil {
		return fmt.Errorf("field 'TargetMarkdown' is required")
	}
	return nil
}

type EvidenceVariableDefinitionModifier struct {
	Id                   *string          `json:"id,omitempty" bson:"id,omitempty"`                   // Unique id for inter-element referencing
	Code                 *CodeableConcept `json:"code" bson:"code"`                                   // Attribute of the definition
	ValueCodeableConcept *CodeableConcept `json:"valueCodeableConcept" bson:"value_codeable_concept"` // Specification of the definition attribute
	ValueBoolean         *bool            `json:"valueBoolean" bson:"value_boolean"`                  // Specification of the definition attribute
	ValueQuantity        *Quantity        `json:"valueQuantity" bson:"value_quantity"`                // Specification of the definition attribute
	ValueRange           *Range           `json:"valueRange" bson:"value_range"`                      // Specification of the definition attribute
	ValuePeriod          *Period          `json:"valuePeriod" bson:"value_period"`                    // Specification of the definition attribute
	ValueRelativeTime    *RelativeTime    `json:"valueRelativeTime" bson:"value_relative_time"`       // Specification of the definition attribute
	ValueReference       *Reference       `json:"valueReference" bson:"value_reference"`              // Specification of the definition attribute
	ValueExpression      *Expression      `json:"valueExpression" bson:"value_expression"`            // Specification of the definition attribute
	ValueUri             *string          `json:"valueUri" bson:"value_uri"`                          // Specification of the definition attribute
}

func (r *EvidenceVariableDefinitionModifier) Validate() error {
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
	if r.ValuePeriod == nil {
		return fmt.Errorf("field 'ValuePeriod' is required")
	}
	if r.ValuePeriod != nil {
		if err := r.ValuePeriod.Validate(); err != nil {
			return fmt.Errorf("ValuePeriod: %w", err)
		}
	}
	if r.ValueRelativeTime == nil {
		return fmt.Errorf("field 'ValueRelativeTime' is required")
	}
	if r.ValueRelativeTime != nil {
		if err := r.ValueRelativeTime.Validate(); err != nil {
			return fmt.Errorf("ValueRelativeTime: %w", err)
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
	if r.ValueExpression == nil {
		return fmt.Errorf("field 'ValueExpression' is required")
	}
	if r.ValueExpression != nil {
		if err := r.ValueExpression.Validate(); err != nil {
			return fmt.Errorf("ValueExpression: %w", err)
		}
	}
	if r.ValueUri == nil {
		return fmt.Errorf("field 'ValueUri' is required")
	}
	return nil
}

type EvidenceVariableCategory struct {
	Id                   *string          `json:"id,omitempty" bson:"id,omitempty"`                                       // Unique id for inter-element referencing
	Name                 *string          `json:"name,omitempty" bson:"name,omitempty"`                                   // Description of the grouping
	ValueCodeableConcept *CodeableConcept `json:"valueCodeableConcept,omitempty" bson:"value_codeable_concept,omitempty"` // Definition of the grouping
	ValueQuantity        *Quantity        `json:"valueQuantity,omitempty" bson:"value_quantity,omitempty"`                // Definition of the grouping
	ValueRange           *Range           `json:"valueRange,omitempty" bson:"value_range,omitempty"`                      // Definition of the grouping
	ValueReference       *Reference       `json:"valueReference,omitempty" bson:"value_reference,omitempty"`              // Definition of the grouping
}

func (r *EvidenceVariableCategory) Validate() error {
	if r.ValueCodeableConcept != nil {
		if err := r.ValueCodeableConcept.Validate(); err != nil {
			return fmt.Errorf("ValueCodeableConcept: %w", err)
		}
	}
	if r.ValueQuantity != nil {
		if err := r.ValueQuantity.Validate(); err != nil {
			return fmt.Errorf("ValueQuantity: %w", err)
		}
	}
	if r.ValueRange != nil {
		if err := r.ValueRange.Validate(); err != nil {
			return fmt.Errorf("ValueRange: %w", err)
		}
	}
	if r.ValueReference != nil {
		if err := r.ValueReference.Validate(); err != nil {
			return fmt.Errorf("ValueReference: %w", err)
		}
	}
	return nil
}
