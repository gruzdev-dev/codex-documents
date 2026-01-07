package models

import (
	"encoding/json"
	"fmt"
)

// A structured set of questions intended to guide the collection of answers from end-users. Questionnaires provide detailed control over order, presentation, phraseology and grouping to allow coherent, consistent data collection.
type Questionnaire struct {
	ResourceType           string              `json:"resourceType" bson:"resource_type"`                                          // Type of resource
	Id                     *string             `json:"id,omitempty" bson:"id,omitempty"`                                           // Logical id of this artifact
	Meta                   *Meta               `json:"meta,omitempty" bson:"meta,omitempty"`                                       // Metadata about the resource
	ImplicitRules          *string             `json:"implicitRules,omitempty" bson:"implicit_rules,omitempty"`                    // A set of rules under which this content was created
	Language               *string             `json:"language,omitempty" bson:"language,omitempty"`                               // Language of the resource content
	Text                   *Narrative          `json:"text,omitempty" bson:"text,omitempty"`                                       // Text summary of the resource, for human interpretation
	Contained              []json.RawMessage   `json:"contained,omitempty" bson:"contained,omitempty"`                             // Contained, inline Resources
	Url                    *string             `json:"url,omitempty" bson:"url,omitempty"`                                         // Canonical identifier for this questionnaire, represented as an absolute URI (globally unique)
	Identifier             []Identifier        `json:"identifier,omitempty" bson:"identifier,omitempty"`                           // Business identifier for questionnaire
	Version                *string             `json:"version,omitempty" bson:"version,omitempty"`                                 // Business version of the questionnaire
	VersionAlgorithmString *string             `json:"versionAlgorithmString,omitempty" bson:"version_algorithm_string,omitempty"` // How to compare versions
	VersionAlgorithmCoding *Coding             `json:"versionAlgorithmCoding,omitempty" bson:"version_algorithm_coding,omitempty"` // How to compare versions
	Name                   *string             `json:"name,omitempty" bson:"name,omitempty"`                                       // Name for this questionnaire (computer friendly)
	Title                  *string             `json:"title,omitempty" bson:"title,omitempty"`                                     // Name for this questionnaire (human friendly)
	DerivedFrom            []string            `json:"derivedFrom,omitempty" bson:"derived_from,omitempty"`                        // Based on Questionnaire
	Status                 string              `json:"status" bson:"status"`                                                       // draft | active | retired | unknown
	Experimental           *bool               `json:"experimental,omitempty" bson:"experimental,omitempty"`                       // For testing only - never for real usage
	SubjectType            []string            `json:"subjectType,omitempty" bson:"subject_type,omitempty"`                        // Resource that can be subject of QuestionnaireResponse
	Date                   *string             `json:"date,omitempty" bson:"date,omitempty"`                                       // Date last changed
	Publisher              *string             `json:"publisher,omitempty" bson:"publisher,omitempty"`                             // Name of the publisher/steward (organization or individual)
	Contact                []ContactDetail     `json:"contact,omitempty" bson:"contact,omitempty"`                                 // Contact details for the publisher
	Description            *string             `json:"description,omitempty" bson:"description,omitempty"`                         // Natural language description of the questionnaire
	UseContext             []UsageContext      `json:"useContext,omitempty" bson:"use_context,omitempty"`                          // The context that the content is intended to support
	Jurisdiction           []CodeableConcept   `json:"jurisdiction,omitempty" bson:"jurisdiction,omitempty"`                       // Jurisdiction of the authority that maintains the questionnaire (if applicable)
	Purpose                *string             `json:"purpose,omitempty" bson:"purpose,omitempty"`                                 // Why this questionnaire is defined
	Copyright              *string             `json:"copyright,omitempty" bson:"copyright,omitempty"`                             // Notice about intellectual property ownership, can include restrictions on use
	CopyrightLabel         *string             `json:"copyrightLabel,omitempty" bson:"copyright_label,omitempty"`                  // Copyright holder and year(s)
	ApprovalDate           *string             `json:"approvalDate,omitempty" bson:"approval_date,omitempty"`                      // When the questionnaire was approved by publisher
	LastReviewDate         *string             `json:"lastReviewDate,omitempty" bson:"last_review_date,omitempty"`                 // When the questionnaire was last reviewed by the publisher
	EffectivePeriod        *Period             `json:"effectivePeriod,omitempty" bson:"effective_period,omitempty"`                // When the questionnaire is expected to be used
	Code                   []Coding            `json:"code,omitempty" bson:"code,omitempty"`                                       // Concept that represents the overall questionnaire
	Item                   []QuestionnaireItem `json:"item,omitempty" bson:"item,omitempty"`                                       // Questions and sections within the Questionnaire
}

func (r *Questionnaire) Validate() error {
	if r.ResourceType != "Questionnaire" {
		return fmt.Errorf("invalid resourceType: expected 'Questionnaire', got '%s'", r.ResourceType)
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
	for i, item := range r.Code {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Code[%d]: %w", i, err)
		}
	}
	for i, item := range r.Item {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Item[%d]: %w", i, err)
		}
	}
	return nil
}

type QuestionnaireItem struct {
	Id               *string                         `json:"id,omitempty" bson:"id,omitempty"`                              // Unique id for inter-element referencing
	LinkId           string                          `json:"linkId" bson:"link_id"`                                         // Unique id for item in questionnaire
	Definition       []string                        `json:"definition,omitempty" bson:"definition,omitempty"`              // ElementDefinition - details for the item
	Code             []Coding                        `json:"code,omitempty" bson:"code,omitempty"`                          // Corresponding concept for this item in a terminology
	Prefix           *string                         `json:"prefix,omitempty" bson:"prefix,omitempty"`                      // E.g. "1(a)", "2.5.3"
	Text             *string                         `json:"text,omitempty" bson:"text,omitempty"`                          // Primary text for the item
	Type             string                          `json:"type" bson:"type"`                                              // group | display | boolean | decimal | integer | date | dateTime +
	EnableWhen       []QuestionnaireItemEnableWhen   `json:"enableWhen,omitempty" bson:"enable_when,omitempty"`             // Only allow data when
	EnableBehavior   *string                         `json:"enableBehavior,omitempty" bson:"enable_behavior,omitempty"`     // all | any
	DisabledDisplay  *string                         `json:"disabledDisplay,omitempty" bson:"disabled_display,omitempty"`   // hidden | protected
	Required         *bool                           `json:"required,omitempty" bson:"required,omitempty"`                  // Whether the item must be included in data results
	Repeats          *bool                           `json:"repeats,omitempty" bson:"repeats,omitempty"`                    // Whether the item may repeat
	ReadOnly         *bool                           `json:"readOnly,omitempty" bson:"read_only,omitempty"`                 // Don't allow human editing
	MaxLength        *int                            `json:"maxLength,omitempty" bson:"max_length,omitempty"`               // No more than these many characters
	AnswerConstraint *string                         `json:"answerConstraint,omitempty" bson:"answer_constraint,omitempty"` // optionsOnly | optionsOrType | optionsOrString
	AnswerValueSet   *string                         `json:"answerValueSet,omitempty" bson:"answer_value_set,omitempty"`    // ValueSet containing permitted answers
	AnswerOption     []QuestionnaireItemAnswerOption `json:"answerOption,omitempty" bson:"answer_option,omitempty"`         // Permitted answer
	Initial          []QuestionnaireItemInitial      `json:"initial,omitempty" bson:"initial,omitempty"`                    // Initial value(s) when item is first rendered
	Item             []QuestionnaireItem             `json:"item,omitempty" bson:"item,omitempty"`                          // Nested questionnaire items
}

func (r *QuestionnaireItem) Validate() error {
	var emptyString string
	if r.LinkId == emptyString {
		return fmt.Errorf("field 'LinkId' is required")
	}
	for i, item := range r.Code {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Code[%d]: %w", i, err)
		}
	}
	if r.Type == emptyString {
		return fmt.Errorf("field 'Type' is required")
	}
	for i, item := range r.EnableWhen {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("EnableWhen[%d]: %w", i, err)
		}
	}
	for i, item := range r.AnswerOption {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("AnswerOption[%d]: %w", i, err)
		}
	}
	for i, item := range r.Initial {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Initial[%d]: %w", i, err)
		}
	}
	for i, item := range r.Item {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Item[%d]: %w", i, err)
		}
	}
	return nil
}

type QuestionnaireItemEnableWhen struct {
	Id               *string     `json:"id,omitempty" bson:"id,omitempty"`          // Unique id for inter-element referencing
	Question         string      `json:"question" bson:"question"`                  // The linkId of question that determines whether item is enabled/disabled
	Operator         string      `json:"operator" bson:"operator"`                  // exists | = | != | > | < | >= | <=
	AnswerBoolean    *bool       `json:"answerBoolean" bson:"answer_boolean"`       // Value for question comparison based on operator
	AnswerDecimal    *float64    `json:"answerDecimal" bson:"answer_decimal"`       // Value for question comparison based on operator
	AnswerInteger    *int        `json:"answerInteger" bson:"answer_integer"`       // Value for question comparison based on operator
	AnswerDate       *string     `json:"answerDate" bson:"answer_date"`             // Value for question comparison based on operator
	AnswerDateTime   *string     `json:"answerDateTime" bson:"answer_date_time"`    // Value for question comparison based on operator
	AnswerTime       *string     `json:"answerTime" bson:"answer_time"`             // Value for question comparison based on operator
	AnswerString     *string     `json:"answerString" bson:"answer_string"`         // Value for question comparison based on operator
	AnswerCoding     *Coding     `json:"answerCoding" bson:"answer_coding"`         // Value for question comparison based on operator
	AnswerQuantity   *Quantity   `json:"answerQuantity" bson:"answer_quantity"`     // Value for question comparison based on operator
	AnswerReference  *Reference  `json:"answerReference" bson:"answer_reference"`   // Value for question comparison based on operator
	AnswerUri        *string     `json:"answerUri" bson:"answer_uri"`               // Value for question comparison based on operator
	AnswerAttachment *Attachment `json:"answerAttachment" bson:"answer_attachment"` // Value for question comparison based on operator
}

func (r *QuestionnaireItemEnableWhen) Validate() error {
	var emptyString string
	if r.Question == emptyString {
		return fmt.Errorf("field 'Question' is required")
	}
	if r.Operator == emptyString {
		return fmt.Errorf("field 'Operator' is required")
	}
	if r.AnswerBoolean == nil {
		return fmt.Errorf("field 'AnswerBoolean' is required")
	}
	if r.AnswerDecimal == nil {
		return fmt.Errorf("field 'AnswerDecimal' is required")
	}
	if r.AnswerInteger == nil {
		return fmt.Errorf("field 'AnswerInteger' is required")
	}
	if r.AnswerDate == nil {
		return fmt.Errorf("field 'AnswerDate' is required")
	}
	if r.AnswerDateTime == nil {
		return fmt.Errorf("field 'AnswerDateTime' is required")
	}
	if r.AnswerTime == nil {
		return fmt.Errorf("field 'AnswerTime' is required")
	}
	if r.AnswerString == nil {
		return fmt.Errorf("field 'AnswerString' is required")
	}
	if r.AnswerCoding == nil {
		return fmt.Errorf("field 'AnswerCoding' is required")
	}
	if r.AnswerCoding != nil {
		if err := r.AnswerCoding.Validate(); err != nil {
			return fmt.Errorf("AnswerCoding: %w", err)
		}
	}
	if r.AnswerQuantity == nil {
		return fmt.Errorf("field 'AnswerQuantity' is required")
	}
	if r.AnswerQuantity != nil {
		if err := r.AnswerQuantity.Validate(); err != nil {
			return fmt.Errorf("AnswerQuantity: %w", err)
		}
	}
	if r.AnswerReference == nil {
		return fmt.Errorf("field 'AnswerReference' is required")
	}
	if r.AnswerReference != nil {
		if err := r.AnswerReference.Validate(); err != nil {
			return fmt.Errorf("AnswerReference: %w", err)
		}
	}
	if r.AnswerUri == nil {
		return fmt.Errorf("field 'AnswerUri' is required")
	}
	if r.AnswerAttachment == nil {
		return fmt.Errorf("field 'AnswerAttachment' is required")
	}
	if r.AnswerAttachment != nil {
		if err := r.AnswerAttachment.Validate(); err != nil {
			return fmt.Errorf("AnswerAttachment: %w", err)
		}
	}
	return nil
}

type QuestionnaireItemAnswerOption struct {
	Id              *string    `json:"id,omitempty" bson:"id,omitempty"`                            // Unique id for inter-element referencing
	ValueDecimal    *float64   `json:"valueDecimal" bson:"value_decimal"`                           // Answer value
	ValueInteger    *int       `json:"valueInteger" bson:"value_integer"`                           // Answer value
	ValueDate       *string    `json:"valueDate" bson:"value_date"`                                 // Answer value
	ValueDateTime   *string    `json:"valueDateTime" bson:"value_date_time"`                        // Answer value
	ValueTime       *string    `json:"valueTime" bson:"value_time"`                                 // Answer value
	ValueString     *string    `json:"valueString" bson:"value_string"`                             // Answer value
	ValueUri        *string    `json:"valueUri" bson:"value_uri"`                                   // Answer value
	ValueCoding     *Coding    `json:"valueCoding" bson:"value_coding"`                             // Answer value
	ValueQuantity   *Quantity  `json:"valueQuantity" bson:"value_quantity"`                         // Answer value
	ValueReference  *Reference `json:"valueReference" bson:"value_reference"`                       // Answer value
	InitialSelected *bool      `json:"initialSelected,omitempty" bson:"initial_selected,omitempty"` // Whether option is selected by default
}

func (r *QuestionnaireItemAnswerOption) Validate() error {
	if r.ValueDecimal == nil {
		return fmt.Errorf("field 'ValueDecimal' is required")
	}
	if r.ValueInteger == nil {
		return fmt.Errorf("field 'ValueInteger' is required")
	}
	if r.ValueDate == nil {
		return fmt.Errorf("field 'ValueDate' is required")
	}
	if r.ValueDateTime == nil {
		return fmt.Errorf("field 'ValueDateTime' is required")
	}
	if r.ValueTime == nil {
		return fmt.Errorf("field 'ValueTime' is required")
	}
	if r.ValueString == nil {
		return fmt.Errorf("field 'ValueString' is required")
	}
	if r.ValueUri == nil {
		return fmt.Errorf("field 'ValueUri' is required")
	}
	if r.ValueCoding == nil {
		return fmt.Errorf("field 'ValueCoding' is required")
	}
	if r.ValueCoding != nil {
		if err := r.ValueCoding.Validate(); err != nil {
			return fmt.Errorf("ValueCoding: %w", err)
		}
	}
	if r.ValueQuantity == nil {
		return fmt.Errorf("field 'ValueQuantity' is required")
	}
	if r.ValueQuantity != nil {
		if err := r.ValueQuantity.Validate(); err != nil {
			return fmt.Errorf("ValueQuantity: %w", err)
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
	return nil
}

type QuestionnaireItemInitial struct {
	Id              *string     `json:"id,omitempty" bson:"id,omitempty"`        // Unique id for inter-element referencing
	ValueBoolean    *bool       `json:"valueBoolean" bson:"value_boolean"`       // Actual value for initializing the question
	ValueDecimal    *float64    `json:"valueDecimal" bson:"value_decimal"`       // Actual value for initializing the question
	ValueInteger    *int        `json:"valueInteger" bson:"value_integer"`       // Actual value for initializing the question
	ValueDate       *string     `json:"valueDate" bson:"value_date"`             // Actual value for initializing the question
	ValueDateTime   *string     `json:"valueDateTime" bson:"value_date_time"`    // Actual value for initializing the question
	ValueTime       *string     `json:"valueTime" bson:"value_time"`             // Actual value for initializing the question
	ValueString     *string     `json:"valueString" bson:"value_string"`         // Actual value for initializing the question
	ValueUri        *string     `json:"valueUri" bson:"value_uri"`               // Actual value for initializing the question
	ValueAttachment *Attachment `json:"valueAttachment" bson:"value_attachment"` // Actual value for initializing the question
	ValueCoding     *Coding     `json:"valueCoding" bson:"value_coding"`         // Actual value for initializing the question
	ValueQuantity   *Quantity   `json:"valueQuantity" bson:"value_quantity"`     // Actual value for initializing the question
	ValueReference  *Reference  `json:"valueReference" bson:"value_reference"`   // Actual value for initializing the question
}

func (r *QuestionnaireItemInitial) Validate() error {
	if r.ValueBoolean == nil {
		return fmt.Errorf("field 'ValueBoolean' is required")
	}
	if r.ValueDecimal == nil {
		return fmt.Errorf("field 'ValueDecimal' is required")
	}
	if r.ValueInteger == nil {
		return fmt.Errorf("field 'ValueInteger' is required")
	}
	if r.ValueDate == nil {
		return fmt.Errorf("field 'ValueDate' is required")
	}
	if r.ValueDateTime == nil {
		return fmt.Errorf("field 'ValueDateTime' is required")
	}
	if r.ValueTime == nil {
		return fmt.Errorf("field 'ValueTime' is required")
	}
	if r.ValueString == nil {
		return fmt.Errorf("field 'ValueString' is required")
	}
	if r.ValueUri == nil {
		return fmt.Errorf("field 'ValueUri' is required")
	}
	if r.ValueAttachment == nil {
		return fmt.Errorf("field 'ValueAttachment' is required")
	}
	if r.ValueAttachment != nil {
		if err := r.ValueAttachment.Validate(); err != nil {
			return fmt.Errorf("ValueAttachment: %w", err)
		}
	}
	if r.ValueCoding == nil {
		return fmt.Errorf("field 'ValueCoding' is required")
	}
	if r.ValueCoding != nil {
		if err := r.ValueCoding.Validate(); err != nil {
			return fmt.Errorf("ValueCoding: %w", err)
		}
	}
	if r.ValueQuantity == nil {
		return fmt.Errorf("field 'ValueQuantity' is required")
	}
	if r.ValueQuantity != nil {
		if err := r.ValueQuantity.Validate(); err != nil {
			return fmt.Errorf("ValueQuantity: %w", err)
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
	return nil
}
