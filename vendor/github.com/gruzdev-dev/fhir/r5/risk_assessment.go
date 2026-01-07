package models

import (
	"encoding/json"
	"fmt"
)

// An assessment of the likely outcome(s) for a patient or other subject as well as the likelihood of each outcome.
type RiskAssessment struct {
	ResourceType       string                     `json:"resourceType" bson:"resource_type"`                                  // Type of resource
	Id                 *string                    `json:"id,omitempty" bson:"id,omitempty"`                                   // Logical id of this artifact
	Meta               *Meta                      `json:"meta,omitempty" bson:"meta,omitempty"`                               // Metadata about the resource
	ImplicitRules      *string                    `json:"implicitRules,omitempty" bson:"implicit_rules,omitempty"`            // A set of rules under which this content was created
	Language           *string                    `json:"language,omitempty" bson:"language,omitempty"`                       // Language of the resource content
	Text               *Narrative                 `json:"text,omitempty" bson:"text,omitempty"`                               // Text summary of the resource, for human interpretation
	Contained          []json.RawMessage          `json:"contained,omitempty" bson:"contained,omitempty"`                     // Contained, inline Resources
	Identifier         []Identifier               `json:"identifier,omitempty" bson:"identifier,omitempty"`                   // Unique identifier for the assessment
	BasedOn            *Reference                 `json:"basedOn,omitempty" bson:"based_on,omitempty"`                        // Request fulfilled by this assessment
	Parent             *Reference                 `json:"parent,omitempty" bson:"parent,omitempty"`                           // Part of this occurrence
	Status             string                     `json:"status" bson:"status"`                                               // registered | specimen-in-process | preliminary | final | amended | corrected | appended | cancelled | entered-in-error | unknown | cannot-be-obtained
	Method             *CodeableConcept           `json:"method,omitempty" bson:"method,omitempty"`                           // Evaluation mechanism
	Code               *CodeableConcept           `json:"code,omitempty" bson:"code,omitempty"`                               // Type of assessment
	Subject            *Reference                 `json:"subject" bson:"subject"`                                             // Who/what does assessment apply to?
	Encounter          *Reference                 `json:"encounter,omitempty" bson:"encounter,omitempty"`                     // Where was assessment performed?
	OccurrenceDateTime *string                    `json:"occurrenceDateTime,omitempty" bson:"occurrence_date_time,omitempty"` // When was assessment made?
	OccurrencePeriod   *Period                    `json:"occurrencePeriod,omitempty" bson:"occurrence_period,omitempty"`      // When was assessment made?
	Condition          *Reference                 `json:"condition,omitempty" bson:"condition,omitempty"`                     // Condition assessed
	Performer          *Reference                 `json:"performer,omitempty" bson:"performer,omitempty"`                     // Who did assessment?
	Reason             []CodeableReference        `json:"reason,omitempty" bson:"reason,omitempty"`                           // Why was the assessment necessary?
	Basis              []Reference                `json:"basis,omitempty" bson:"basis,omitempty"`                             // Information used in assessment
	Prediction         []RiskAssessmentPrediction `json:"prediction,omitempty" bson:"prediction,omitempty"`                   // Outcome predicted
	Mitigation         *string                    `json:"mitigation,omitempty" bson:"mitigation,omitempty"`                   // How to reduce risk
	Note               []Annotation               `json:"note,omitempty" bson:"note,omitempty"`                               // Comments on the risk assessment
}

func (r *RiskAssessment) Validate() error {
	if r.ResourceType != "RiskAssessment" {
		return fmt.Errorf("invalid resourceType: expected 'RiskAssessment', got '%s'", r.ResourceType)
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
	if r.BasedOn != nil {
		if err := r.BasedOn.Validate(); err != nil {
			return fmt.Errorf("BasedOn: %w", err)
		}
	}
	if r.Parent != nil {
		if err := r.Parent.Validate(); err != nil {
			return fmt.Errorf("Parent: %w", err)
		}
	}
	var emptyString string
	if r.Status == emptyString {
		return fmt.Errorf("field 'Status' is required")
	}
	if r.Method != nil {
		if err := r.Method.Validate(); err != nil {
			return fmt.Errorf("Method: %w", err)
		}
	}
	if r.Code != nil {
		if err := r.Code.Validate(); err != nil {
			return fmt.Errorf("Code: %w", err)
		}
	}
	if r.Subject == nil {
		return fmt.Errorf("field 'Subject' is required")
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
	if r.OccurrencePeriod != nil {
		if err := r.OccurrencePeriod.Validate(); err != nil {
			return fmt.Errorf("OccurrencePeriod: %w", err)
		}
	}
	if r.Condition != nil {
		if err := r.Condition.Validate(); err != nil {
			return fmt.Errorf("Condition: %w", err)
		}
	}
	if r.Performer != nil {
		if err := r.Performer.Validate(); err != nil {
			return fmt.Errorf("Performer: %w", err)
		}
	}
	for i, item := range r.Reason {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Reason[%d]: %w", i, err)
		}
	}
	for i, item := range r.Basis {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Basis[%d]: %w", i, err)
		}
	}
	for i, item := range r.Prediction {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Prediction[%d]: %w", i, err)
		}
	}
	for i, item := range r.Note {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Note[%d]: %w", i, err)
		}
	}
	return nil
}

type RiskAssessmentPrediction struct {
	Id                  *string          `json:"id,omitempty" bson:"id,omitempty"`                                    // Unique id for inter-element referencing
	Outcome             *CodeableConcept `json:"outcome,omitempty" bson:"outcome,omitempty"`                          // Possible outcome for the subject
	ProbabilityDecimal  *float64         `json:"probabilityDecimal,omitempty" bson:"probability_decimal,omitempty"`   // Likelihood of specified outcome
	ProbabilityQuantity *Quantity        `json:"probabilityQuantity,omitempty" bson:"probability_quantity,omitempty"` // Likelihood of specified outcome
	ProbabilityRange    *Range           `json:"probabilityRange,omitempty" bson:"probability_range,omitempty"`       // Likelihood of specified outcome
	QualitativeRisk     *CodeableConcept `json:"qualitativeRisk,omitempty" bson:"qualitative_risk,omitempty"`         // Likelihood of specified outcome as a qualitative value
	RelativeRisk        *float64         `json:"relativeRisk,omitempty" bson:"relative_risk,omitempty"`               // Relative likelihood
	WhenPeriod          *Period          `json:"whenPeriod,omitempty" bson:"when_period,omitempty"`                   // Timeframe or age range
	WhenRange           *Range           `json:"whenRange,omitempty" bson:"when_range,omitempty"`                     // Timeframe or age range
	Rationale           *string          `json:"rationale,omitempty" bson:"rationale,omitempty"`                      // Explanation of prediction
}

func (r *RiskAssessmentPrediction) Validate() error {
	if r.Outcome != nil {
		if err := r.Outcome.Validate(); err != nil {
			return fmt.Errorf("Outcome: %w", err)
		}
	}
	if r.ProbabilityQuantity != nil {
		if err := r.ProbabilityQuantity.Validate(); err != nil {
			return fmt.Errorf("ProbabilityQuantity: %w", err)
		}
	}
	if r.ProbabilityRange != nil {
		if err := r.ProbabilityRange.Validate(); err != nil {
			return fmt.Errorf("ProbabilityRange: %w", err)
		}
	}
	if r.QualitativeRisk != nil {
		if err := r.QualitativeRisk.Validate(); err != nil {
			return fmt.Errorf("QualitativeRisk: %w", err)
		}
	}
	if r.WhenPeriod != nil {
		if err := r.WhenPeriod.Validate(); err != nil {
			return fmt.Errorf("WhenPeriod: %w", err)
		}
	}
	if r.WhenRange != nil {
		if err := r.WhenRange.Validate(); err != nil {
			return fmt.Errorf("WhenRange: %w", err)
		}
	}
	return nil
}
