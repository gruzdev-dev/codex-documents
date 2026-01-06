package models

import (
	"encoding/json"
	"fmt"
)

// Risk of harmful or undesirable, physiological response which is unique to an individual and associated with exposure to a substance.
type AllergyIntolerance struct {
	Id                     *string                      `json:"id,omitempty" bson:"id,omitempty"`                                           // Logical id of this artifact
	Meta                   *Meta                        `json:"meta,omitempty" bson:"meta,omitempty"`                                       // Metadata about the resource
	ImplicitRules          *string                      `json:"implicitRules,omitempty" bson:"implicit_rules,omitempty"`                    // A set of rules under which this content was created
	Language               *string                      `json:"language,omitempty" bson:"language,omitempty"`                               // Language of the resource content
	Text                   *Narrative                   `json:"text,omitempty" bson:"text,omitempty"`                                       // Text summary of the resource, for human interpretation
	Contained              []json.RawMessage            `json:"contained,omitempty" bson:"contained,omitempty"`                             // Contained, inline Resources
	Identifier             []Identifier                 `json:"identifier,omitempty" bson:"identifier,omitempty"`                           // External ids for this item
	ClinicalStatus         *CodeableConcept             `json:"clinicalStatus,omitempty" bson:"clinical_status,omitempty"`                  // active | inactive | resolved
	VerificationStatus     *CodeableConcept             `json:"verificationStatus,omitempty" bson:"verification_status,omitempty"`          // unconfirmed | presumed | confirmed | refuted | entered-in-error
	Type                   *CodeableConcept             `json:"type,omitempty" bson:"type,omitempty"`                                       // allergy | intolerance - Underlying mechanism (if known)
	Category               []string                     `json:"category,omitempty" bson:"category,omitempty"`                               // food | medication | environment | biologic
	Criticality            *string                      `json:"criticality,omitempty" bson:"criticality,omitempty"`                         // low | high | unable-to-assess
	Code                   *CodeableConcept             `json:"code,omitempty" bson:"code,omitempty"`                                       // Code that identifies the allergy or intolerance
	Patient                *Reference                   `json:"patient" bson:"patient"`                                                     // Who the allergy or intolerance is for
	Encounter              *Reference                   `json:"encounter,omitempty" bson:"encounter,omitempty"`                             // Encounter when the allergy or intolerance was asserted
	OnsetDateTime          *string                      `json:"onsetDateTime,omitempty" bson:"onset_date_time,omitempty"`                   // When allergy or intolerance was identified
	OnsetAge               *Age                         `json:"onsetAge,omitempty" bson:"onset_age,omitempty"`                              // When allergy or intolerance was identified
	OnsetPeriod            *Period                      `json:"onsetPeriod,omitempty" bson:"onset_period,omitempty"`                        // When allergy or intolerance was identified
	OnsetRange             *Range                       `json:"onsetRange,omitempty" bson:"onset_range,omitempty"`                          // When allergy or intolerance was identified
	OnsetString            *string                      `json:"onsetString,omitempty" bson:"onset_string,omitempty"`                        // When allergy or intolerance was identified
	RecordedDate           *string                      `json:"recordedDate,omitempty" bson:"recorded_date,omitempty"`                      // Date allergy or intolerance was first recorded
	Recorder               *Reference                   `json:"recorder,omitempty" bson:"recorder,omitempty"`                               // Who recorded the sensitivity
	Asserter               *Reference                   `json:"asserter,omitempty" bson:"asserter,omitempty"`                               // Source of the information about the allergy
	LastReactionOccurrence *string                      `json:"lastReactionOccurrence,omitempty" bson:"last_reaction_occurrence,omitempty"` // Date(/time) of last known occurrence of a reaction
	Note                   []Annotation                 `json:"note,omitempty" bson:"note,omitempty"`                                       // Additional text not captured in other fields
	Reaction               []AllergyIntoleranceReaction `json:"reaction,omitempty" bson:"reaction,omitempty"`                               // Adverse Reaction Events linked to exposure to substance
}

func (r *AllergyIntolerance) Validate() error {
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
	if r.ClinicalStatus != nil {
		if err := r.ClinicalStatus.Validate(); err != nil {
			return fmt.Errorf("ClinicalStatus: %w", err)
		}
	}
	if r.VerificationStatus != nil {
		if err := r.VerificationStatus.Validate(); err != nil {
			return fmt.Errorf("VerificationStatus: %w", err)
		}
	}
	if r.Type != nil {
		if err := r.Type.Validate(); err != nil {
			return fmt.Errorf("Type: %w", err)
		}
	}
	if r.Code != nil {
		if err := r.Code.Validate(); err != nil {
			return fmt.Errorf("Code: %w", err)
		}
	}
	if r.Patient == nil {
		return fmt.Errorf("field 'Patient' is required")
	}
	if r.Patient != nil {
		if err := r.Patient.Validate(); err != nil {
			return fmt.Errorf("Patient: %w", err)
		}
	}
	if r.Encounter != nil {
		if err := r.Encounter.Validate(); err != nil {
			return fmt.Errorf("Encounter: %w", err)
		}
	}
	if r.OnsetAge != nil {
		if err := r.OnsetAge.Validate(); err != nil {
			return fmt.Errorf("OnsetAge: %w", err)
		}
	}
	if r.OnsetPeriod != nil {
		if err := r.OnsetPeriod.Validate(); err != nil {
			return fmt.Errorf("OnsetPeriod: %w", err)
		}
	}
	if r.OnsetRange != nil {
		if err := r.OnsetRange.Validate(); err != nil {
			return fmt.Errorf("OnsetRange: %w", err)
		}
	}
	if r.Recorder != nil {
		if err := r.Recorder.Validate(); err != nil {
			return fmt.Errorf("Recorder: %w", err)
		}
	}
	if r.Asserter != nil {
		if err := r.Asserter.Validate(); err != nil {
			return fmt.Errorf("Asserter: %w", err)
		}
	}
	for i, item := range r.Note {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Note[%d]: %w", i, err)
		}
	}
	for i, item := range r.Reaction {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Reaction[%d]: %w", i, err)
		}
	}
	return nil
}

type AllergyIntoleranceReaction struct {
	Id            *string             `json:"id,omitempty" bson:"id,omitempty"`                        // Unique id for inter-element referencing
	Substance     *CodeableConcept    `json:"substance,omitempty" bson:"substance,omitempty"`          // Specific substance or pharmaceutical product considered to be responsible for event
	Manifestation []CodeableReference `json:"manifestation" bson:"manifestation"`                      // Clinical symptoms/signs associated with the Event
	Description   *string             `json:"description,omitempty" bson:"description,omitempty"`      // Description of the event as a whole
	Onset         *string             `json:"onset,omitempty" bson:"onset,omitempty"`                  // Date(/time) when manifestations showed
	Severity      *string             `json:"severity,omitempty" bson:"severity,omitempty"`            // mild | moderate | severe (of event as a whole)
	ExposureRoute *CodeableConcept    `json:"exposureRoute,omitempty" bson:"exposure_route,omitempty"` // How the subject was exposed to the substance
	Note          []Annotation        `json:"note,omitempty" bson:"note,omitempty"`                    // Text about event not captured in other fields
}

func (r *AllergyIntoleranceReaction) Validate() error {
	if r.Substance != nil {
		if err := r.Substance.Validate(); err != nil {
			return fmt.Errorf("Substance: %w", err)
		}
	}
	if len(r.Manifestation) < 1 {
		return fmt.Errorf("field 'Manifestation' must have at least 1 elements")
	}
	for i, item := range r.Manifestation {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Manifestation[%d]: %w", i, err)
		}
	}
	if r.ExposureRoute != nil {
		if err := r.ExposureRoute.Validate(); err != nil {
			return fmt.Errorf("ExposureRoute: %w", err)
		}
	}
	for i, item := range r.Note {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Note[%d]: %w", i, err)
		}
	}
	return nil
}
