package models

import (
	"encoding/json"
	"fmt"
)

// A guidance response is the formal response to a guidance request, including any output parameters returned by the evaluation, as well as the description of any proposed actions to be taken.
type GuidanceResponse struct {
	ResourceType          string              `json:"resourceType" bson:"resource_type"`                                  // Type of resource
	Id                    *string             `json:"id,omitempty" bson:"id,omitempty"`                                   // Logical id of this artifact
	Meta                  *Meta               `json:"meta,omitempty" bson:"meta,omitempty"`                               // Metadata about the resource
	ImplicitRules         *string             `json:"implicitRules,omitempty" bson:"implicit_rules,omitempty"`            // A set of rules under which this content was created
	Language              *string             `json:"language,omitempty" bson:"language,omitempty"`                       // Language of the resource content
	Text                  *Narrative          `json:"text,omitempty" bson:"text,omitempty"`                               // Text summary of the resource, for human interpretation
	Contained             []json.RawMessage   `json:"contained,omitempty" bson:"contained,omitempty"`                     // Contained, inline Resources
	RequestIdentifier     *Identifier         `json:"requestIdentifier,omitempty" bson:"request_identifier,omitempty"`    // The identifier of the request associated with this response, if any
	Identifier            []Identifier        `json:"identifier,omitempty" bson:"identifier,omitempty"`                   // Business identifier for guidance response
	ModuleUri             *string             `json:"moduleUri" bson:"module_uri"`                                        // What guidance was requested
	ModuleCanonical       *string             `json:"moduleCanonical" bson:"module_canonical"`                            // What guidance was requested
	ModuleCodeableConcept *CodeableConcept    `json:"moduleCodeableConcept" bson:"module_codeable_concept"`               // What guidance was requested
	Status                string              `json:"status" bson:"status"`                                               // success | data-requested | data-required | in-progress | failure | entered-in-error
	Subject               *Reference          `json:"subject,omitempty" bson:"subject,omitempty"`                         // Individual service was done for/to
	Encounter             *Reference          `json:"encounter,omitempty" bson:"encounter,omitempty"`                     // Encounter the guidance response is part of
	OccurrenceDateTime    *string             `json:"occurrenceDateTime,omitempty" bson:"occurrence_date_time,omitempty"` // When the guidance response was processed
	Performer             *Reference          `json:"performer,omitempty" bson:"performer,omitempty"`                     // Device returning the guidance
	Location              *Reference          `json:"location,omitempty" bson:"location,omitempty"`                       // Where guidance response occurred
	Reason                []CodeableReference `json:"reason,omitempty" bson:"reason,omitempty"`                           // Why guidance is needed
	Note                  []Annotation        `json:"note,omitempty" bson:"note,omitempty"`                               // Additional notes about the response
	EvaluationMessage     *Reference          `json:"evaluationMessage,omitempty" bson:"evaluation_message,omitempty"`    // Messages resulting from the evaluation of the artifact or artifacts
	OutputParameters      *Reference          `json:"outputParameters,omitempty" bson:"output_parameters,omitempty"`      // The output parameters of the evaluation, if any
	Result                []Reference         `json:"result,omitempty" bson:"result,omitempty"`                           // Proposed actions, if any
	DataRequirement       []DataRequirement   `json:"dataRequirement,omitempty" bson:"data_requirement,omitempty"`        // Additional required data
}

func (r *GuidanceResponse) Validate() error {
	if r.ResourceType != "GuidanceResponse" {
		return fmt.Errorf("invalid resourceType: expected 'GuidanceResponse', got '%s'", r.ResourceType)
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
	if r.RequestIdentifier != nil {
		if err := r.RequestIdentifier.Validate(); err != nil {
			return fmt.Errorf("RequestIdentifier: %w", err)
		}
	}
	for i, item := range r.Identifier {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Identifier[%d]: %w", i, err)
		}
	}
	if r.ModuleUri == nil {
		return fmt.Errorf("field 'ModuleUri' is required")
	}
	if r.ModuleCanonical == nil {
		return fmt.Errorf("field 'ModuleCanonical' is required")
	}
	if r.ModuleCodeableConcept == nil {
		return fmt.Errorf("field 'ModuleCodeableConcept' is required")
	}
	if r.ModuleCodeableConcept != nil {
		if err := r.ModuleCodeableConcept.Validate(); err != nil {
			return fmt.Errorf("ModuleCodeableConcept: %w", err)
		}
	}
	var emptyString string
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
	if r.Performer != nil {
		if err := r.Performer.Validate(); err != nil {
			return fmt.Errorf("Performer: %w", err)
		}
	}
	if r.Location != nil {
		if err := r.Location.Validate(); err != nil {
			return fmt.Errorf("Location: %w", err)
		}
	}
	for i, item := range r.Reason {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Reason[%d]: %w", i, err)
		}
	}
	for i, item := range r.Note {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Note[%d]: %w", i, err)
		}
	}
	if r.EvaluationMessage != nil {
		if err := r.EvaluationMessage.Validate(); err != nil {
			return fmt.Errorf("EvaluationMessage: %w", err)
		}
	}
	if r.OutputParameters != nil {
		if err := r.OutputParameters.Validate(); err != nil {
			return fmt.Errorf("OutputParameters: %w", err)
		}
	}
	for i, item := range r.Result {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Result[%d]: %w", i, err)
		}
	}
	for i, item := range r.DataRequirement {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("DataRequirement[%d]: %w", i, err)
		}
	}
	return nil
}
