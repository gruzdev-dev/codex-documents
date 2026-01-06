package models

import (
	"encoding/json"
	"fmt"
)

// A structured set of questions and their answers. The questions are ordered and grouped into coherent subsets, corresponding to the structure of the grouping of the questionnaire being responded to.
type QuestionnaireResponse struct {
	Id            *string                     `json:"id,omitempty" bson:"id,omitempty"`                        // Logical id of this artifact
	Meta          *Meta                       `json:"meta,omitempty" bson:"meta,omitempty"`                    // Metadata about the resource
	ImplicitRules *string                     `json:"implicitRules,omitempty" bson:"implicit_rules,omitempty"` // A set of rules under which this content was created
	Language      *string                     `json:"language,omitempty" bson:"language,omitempty"`            // Language of the resource content
	Text          *Narrative                  `json:"text,omitempty" bson:"text,omitempty"`                    // Text summary of the resource, for human interpretation
	Contained     []json.RawMessage           `json:"contained,omitempty" bson:"contained,omitempty"`          // Contained, inline Resources
	Identifier    []Identifier                `json:"identifier,omitempty" bson:"identifier,omitempty"`        // Business identifier for this set of answers
	BasedOn       []Reference                 `json:"basedOn,omitempty" bson:"based_on,omitempty"`             // Request fulfilled by this QuestionnaireResponse
	PartOf        []Reference                 `json:"partOf,omitempty" bson:"part_of,omitempty"`               // Part of referenced event
	Questionnaire string                      `json:"questionnaire" bson:"questionnaire"`                      // Canonical URL of Questionnaire being answered
	Status        string                      `json:"status" bson:"status"`                                    // in-progress | completed | amended | entered-in-error | stopped
	Subject       *Reference                  `json:"subject,omitempty" bson:"subject,omitempty"`              // The subject of the questions
	Encounter     *Reference                  `json:"encounter,omitempty" bson:"encounter,omitempty"`          // Encounter the questionnaire response is part of
	Authored      *string                     `json:"authored,omitempty" bson:"authored,omitempty"`            // Date the answers were gathered
	Author        *Reference                  `json:"author,omitempty" bson:"author,omitempty"`                // The individual or device that received and recorded the answers
	Source        *Reference                  `json:"source,omitempty" bson:"source,omitempty"`                // The individual or device that answered the questions
	Item          []QuestionnaireResponseItem `json:"item,omitempty" bson:"item,omitempty"`                    // Groups and questions
}

func (r *QuestionnaireResponse) Validate() error {
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
	for i, item := range r.BasedOn {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("BasedOn[%d]: %w", i, err)
		}
	}
	for i, item := range r.PartOf {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("PartOf[%d]: %w", i, err)
		}
	}
	var emptyString string
	if r.Questionnaire == emptyString {
		return fmt.Errorf("field 'Questionnaire' is required")
	}
	if r.Status == emptyString {
		return fmt.Errorf("field 'Status' is required")
	}
	if r.Subject != nil {
		if err := r.Subject.Validate(); err != nil {
			return fmt.Errorf("Subject: %w", err)
		}
	}
	if r.Encounter != nil {
		if err := r.Encounter.Validate(); err != nil {
			return fmt.Errorf("Encounter: %w", err)
		}
	}
	if r.Author != nil {
		if err := r.Author.Validate(); err != nil {
			return fmt.Errorf("Author: %w", err)
		}
	}
	if r.Source != nil {
		if err := r.Source.Validate(); err != nil {
			return fmt.Errorf("Source: %w", err)
		}
	}
	for i, item := range r.Item {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Item[%d]: %w", i, err)
		}
	}
	return nil
}

type QuestionnaireResponseItem struct {
	Id         *string                           `json:"id,omitempty" bson:"id,omitempty"`                 // Unique id for inter-element referencing
	LinkId     string                            `json:"linkId" bson:"link_id"`                            // Pointer to specific item from Questionnaire
	Definition []string                          `json:"definition,omitempty" bson:"definition,omitempty"` // ElementDefinition - details for the item
	Text       *string                           `json:"text,omitempty" bson:"text,omitempty"`             // Name for group or question text
	Answer     []QuestionnaireResponseItemAnswer `json:"answer,omitempty" bson:"answer,omitempty"`         // The response(s) to the question
	Item       []QuestionnaireResponseItem       `json:"item,omitempty" bson:"item,omitempty"`             // Child items of group item
}

func (r *QuestionnaireResponseItem) Validate() error {
	var emptyString string
	if r.LinkId == emptyString {
		return fmt.Errorf("field 'LinkId' is required")
	}
	for i, item := range r.Answer {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Answer[%d]: %w", i, err)
		}
	}
	for i, item := range r.Item {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Item[%d]: %w", i, err)
		}
	}
	return nil
}

type QuestionnaireResponseItemAnswer struct {
	Id              *string                     `json:"id,omitempty" bson:"id,omitempty"`        // Unique id for inter-element referencing
	ValueBoolean    *bool                       `json:"valueBoolean" bson:"value_boolean"`       // Single-valued answer to the question
	ValueDecimal    *float64                    `json:"valueDecimal" bson:"value_decimal"`       // Single-valued answer to the question
	ValueInteger    *int                        `json:"valueInteger" bson:"value_integer"`       // Single-valued answer to the question
	ValueDate       *string                     `json:"valueDate" bson:"value_date"`             // Single-valued answer to the question
	ValueDateTime   *string                     `json:"valueDateTime" bson:"value_date_time"`    // Single-valued answer to the question
	ValueTime       *string                     `json:"valueTime" bson:"value_time"`             // Single-valued answer to the question
	ValueString     *string                     `json:"valueString" bson:"value_string"`         // Single-valued answer to the question
	ValueUri        *string                     `json:"valueUri" bson:"value_uri"`               // Single-valued answer to the question
	ValueAttachment *Attachment                 `json:"valueAttachment" bson:"value_attachment"` // Single-valued answer to the question
	ValueCoding     *Coding                     `json:"valueCoding" bson:"value_coding"`         // Single-valued answer to the question
	ValueQuantity   *Quantity                   `json:"valueQuantity" bson:"value_quantity"`     // Single-valued answer to the question
	ValueReference  *Reference                  `json:"valueReference" bson:"value_reference"`   // Single-valued answer to the question
	Item            []QuestionnaireResponseItem `json:"item,omitempty" bson:"item,omitempty"`    // Child items of question
}

func (r *QuestionnaireResponseItemAnswer) Validate() error {
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
	for i, item := range r.Item {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Item[%d]: %w", i, err)
		}
	}
	return nil
}
