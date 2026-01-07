package models

import (
	"encoding/json"
	"fmt"
)

// A clinical condition, problem, diagnosis, or other event, situation, issue, or clinical concept that has risen to a level of concern.
type Condition struct {
	ResourceType       string              `json:"resourceType" bson:"resource_type"`                                 // Type of resource
	Id                 *string             `json:"id,omitempty" bson:"id,omitempty"`                                  // Logical id of this artifact
	Meta               *Meta               `json:"meta,omitempty" bson:"meta,omitempty"`                              // Metadata about the resource
	ImplicitRules      *string             `json:"implicitRules,omitempty" bson:"implicit_rules,omitempty"`           // A set of rules under which this content was created
	Language           *string             `json:"language,omitempty" bson:"language,omitempty"`                      // Language of the resource content
	Text               *Narrative          `json:"text,omitempty" bson:"text,omitempty"`                              // Text summary of the resource, for human interpretation
	Contained          []json.RawMessage   `json:"contained,omitempty" bson:"contained,omitempty"`                    // Contained, inline Resources
	Identifier         []Identifier        `json:"identifier,omitempty" bson:"identifier,omitempty"`                  // External Ids for this condition
	ClinicalStatus     *CodeableConcept    `json:"clinicalStatus" bson:"clinical_status"`                             // active | recurrence | relapse | inactive | remission | resolved | unknown
	VerificationStatus *CodeableConcept    `json:"verificationStatus,omitempty" bson:"verification_status,omitempty"` // unconfirmed | provisional | differential | confirmed | refuted | entered-in-error
	Category           []CodeableConcept   `json:"category,omitempty" bson:"category,omitempty"`                      // problem-list-item | encounter-diagnosis
	Severity           *CodeableConcept    `json:"severity,omitempty" bson:"severity,omitempty"`                      // Subjective severity of condition
	Code               *CodeableConcept    `json:"code,omitempty" bson:"code,omitempty"`                              // Identification of the condition, problem or diagnosis
	BodySite           []CodeableConcept   `json:"bodySite,omitempty" bson:"body_site,omitempty"`                     // Anatomical location, if relevant
	BodyStructure      *Reference          `json:"bodyStructure,omitempty" bson:"body_structure,omitempty"`           // Anatomical body structure
	Subject            *Reference          `json:"subject" bson:"subject"`                                            // Who has the condition?
	Encounter          *Reference          `json:"encounter,omitempty" bson:"encounter,omitempty"`                    // The Encounter during which this Condition was created
	OnsetDateTime      *string             `json:"onsetDateTime,omitempty" bson:"onset_date_time,omitempty"`          // Estimated or actual date,  date-time, or age
	OnsetAge           *Age                `json:"onsetAge,omitempty" bson:"onset_age,omitempty"`                     // Estimated or actual date,  date-time, or age
	OnsetPeriod        *Period             `json:"onsetPeriod,omitempty" bson:"onset_period,omitempty"`               // Estimated or actual date,  date-time, or age
	OnsetRange         *Range              `json:"onsetRange,omitempty" bson:"onset_range,omitempty"`                 // Estimated or actual date,  date-time, or age
	OnsetString        *string             `json:"onsetString,omitempty" bson:"onset_string,omitempty"`               // Estimated or actual date,  date-time, or age
	AbatementDateTime  *string             `json:"abatementDateTime,omitempty" bson:"abatement_date_time,omitempty"`  // When in resolution/remission
	AbatementAge       *Age                `json:"abatementAge,omitempty" bson:"abatement_age,omitempty"`             // When in resolution/remission
	AbatementPeriod    *Period             `json:"abatementPeriod,omitempty" bson:"abatement_period,omitempty"`       // When in resolution/remission
	AbatementRange     *Range              `json:"abatementRange,omitempty" bson:"abatement_range,omitempty"`         // When in resolution/remission
	AbatementString    *string             `json:"abatementString,omitempty" bson:"abatement_string,omitempty"`       // When in resolution/remission
	RecordedDate       *string             `json:"recordedDate,omitempty" bson:"recorded_date,omitempty"`             // Date condition was first recorded
	Recorder           *Reference          `json:"recorder,omitempty" bson:"recorder,omitempty"`                      // Who recorded the condition
	Asserter           *Reference          `json:"asserter,omitempty" bson:"asserter,omitempty"`                      // Person or device that asserts this condition
	Stage              []ConditionStage    `json:"stage,omitempty" bson:"stage,omitempty"`                            // Stage/grade, usually assessed formally
	Evidence           []CodeableReference `json:"evidence,omitempty" bson:"evidence,omitempty"`                      // Supporting evidence for the condition
	Note               []Annotation        `json:"note,omitempty" bson:"note,omitempty"`                              // Additional information about the Condition
}

func (r *Condition) Validate() error {
	if r.ResourceType != "Condition" {
		return fmt.Errorf("invalid resourceType: expected 'Condition', got '%s'", r.ResourceType)
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
	if r.ClinicalStatus == nil {
		return fmt.Errorf("field 'ClinicalStatus' is required")
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
	for i, item := range r.Category {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Category[%d]: %w", i, err)
		}
	}
	if r.Severity != nil {
		if err := r.Severity.Validate(); err != nil {
			return fmt.Errorf("Severity: %w", err)
		}
	}
	if r.Code != nil {
		if err := r.Code.Validate(); err != nil {
			return fmt.Errorf("Code: %w", err)
		}
	}
	for i, item := range r.BodySite {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("BodySite[%d]: %w", i, err)
		}
	}
	if r.BodyStructure != nil {
		if err := r.BodyStructure.Validate(); err != nil {
			return fmt.Errorf("BodyStructure: %w", err)
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
	if r.AbatementAge != nil {
		if err := r.AbatementAge.Validate(); err != nil {
			return fmt.Errorf("AbatementAge: %w", err)
		}
	}
	if r.AbatementPeriod != nil {
		if err := r.AbatementPeriod.Validate(); err != nil {
			return fmt.Errorf("AbatementPeriod: %w", err)
		}
	}
	if r.AbatementRange != nil {
		if err := r.AbatementRange.Validate(); err != nil {
			return fmt.Errorf("AbatementRange: %w", err)
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
	for i, item := range r.Stage {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Stage[%d]: %w", i, err)
		}
	}
	for i, item := range r.Evidence {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Evidence[%d]: %w", i, err)
		}
	}
	for i, item := range r.Note {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Note[%d]: %w", i, err)
		}
	}
	return nil
}

type ConditionStage struct {
	Id         *string          `json:"id,omitempty" bson:"id,omitempty"`                 // Unique id for inter-element referencing
	Summary    *CodeableConcept `json:"summary,omitempty" bson:"summary,omitempty"`       // Simple summary (disease specific)
	Assessment []Reference      `json:"assessment,omitempty" bson:"assessment,omitempty"` // Formal record of assessment
	Type       *CodeableConcept `json:"type,omitempty" bson:"type,omitempty"`             // Kind of staging
}

func (r *ConditionStage) Validate() error {
	if r.Summary != nil {
		if err := r.Summary.Validate(); err != nil {
			return fmt.Errorf("Summary: %w", err)
		}
	}
	for i, item := range r.Assessment {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Assessment[%d]: %w", i, err)
		}
	}
	if r.Type != nil {
		if err := r.Type.Validate(); err != nil {
			return fmt.Errorf("Type: %w", err)
		}
	}
	return nil
}
