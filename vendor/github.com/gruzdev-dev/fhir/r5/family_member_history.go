package models

import (
	"encoding/json"
	"fmt"
)

// Significant health conditions for a person related to the patient relevant in the context of care for the patient.
type FamilyMemberHistory struct {
	Id               *string                        `json:"id,omitempty" bson:"id,omitempty"`                               // Logical id of this artifact
	Meta             *Meta                          `json:"meta,omitempty" bson:"meta,omitempty"`                           // Metadata about the resource
	ImplicitRules    *string                        `json:"implicitRules,omitempty" bson:"implicit_rules,omitempty"`        // A set of rules under which this content was created
	Language         *string                        `json:"language,omitempty" bson:"language,omitempty"`                   // Language of the resource content
	Text             *Narrative                     `json:"text,omitempty" bson:"text,omitempty"`                           // Text summary of the resource, for human interpretation
	Contained        []json.RawMessage              `json:"contained,omitempty" bson:"contained,omitempty"`                 // Contained, inline Resources
	Identifier       []Identifier                   `json:"identifier,omitempty" bson:"identifier,omitempty"`               // External Id(s) for this record
	Status           string                         `json:"status" bson:"status"`                                           // partial | completed | entered-in-error | health-unknown
	DataAbsentReason *CodeableConcept               `json:"dataAbsentReason,omitempty" bson:"data_absent_reason,omitempty"` // subject-unknown | withheld | unable-to-obtain | deferred
	Patient          *Reference                     `json:"patient" bson:"patient"`                                         // Patient history is about
	Date             *string                        `json:"date,omitempty" bson:"date,omitempty"`                           // When history was recorded or last updated
	Recorder         *Reference                     `json:"recorder,omitempty" bson:"recorder,omitempty"`                   // Who recorded the family member history
	Asserter         *Reference                     `json:"asserter,omitempty" bson:"asserter,omitempty"`                   // Person or device that asserts this family member history
	Name             *string                        `json:"name,omitempty" bson:"name,omitempty"`                           // The family member described
	Relationship     *CodeableConcept               `json:"relationship" bson:"relationship"`                               // Relationship to the subject
	Sex              *CodeableConcept               `json:"sex,omitempty" bson:"sex,omitempty"`                             // male | female | other | unknown
	BornPeriod       *Period                        `json:"bornPeriod,omitempty" bson:"born_period,omitempty"`              // (approximate) date of birth
	BornDate         *string                        `json:"bornDate,omitempty" bson:"born_date,omitempty"`                  // (approximate) date of birth
	BornString       *string                        `json:"bornString,omitempty" bson:"born_string,omitempty"`              // (approximate) date of birth
	AgeAge           *Age                           `json:"ageAge,omitempty" bson:"age_age,omitempty"`                      // (approximate) age
	AgeRange         *Range                         `json:"ageRange,omitempty" bson:"age_range,omitempty"`                  // (approximate) age
	AgeString        *string                        `json:"ageString,omitempty" bson:"age_string,omitempty"`                // (approximate) age
	EstimatedAge     bool                           `json:"estimatedAge,omitempty" bson:"estimated_age,omitempty"`          // Age is estimated?
	DeceasedBoolean  *bool                          `json:"deceasedBoolean,omitempty" bson:"deceased_boolean,omitempty"`    // Dead? How old/when?
	DeceasedAge      *Age                           `json:"deceasedAge,omitempty" bson:"deceased_age,omitempty"`            // Dead? How old/when?
	DeceasedRange    *Range                         `json:"deceasedRange,omitempty" bson:"deceased_range,omitempty"`        // Dead? How old/when?
	DeceasedDate     *string                        `json:"deceasedDate,omitempty" bson:"deceased_date,omitempty"`          // Dead? How old/when?
	DeceasedString   *string                        `json:"deceasedString,omitempty" bson:"deceased_string,omitempty"`      // Dead? How old/when?
	Reason           []CodeableReference            `json:"reason,omitempty" bson:"reason,omitempty"`                       // Why was family member history performed?
	Note             []Annotation                   `json:"note,omitempty" bson:"note,omitempty"`                           // General note about related person
	Condition        []FamilyMemberHistoryCondition `json:"condition,omitempty" bson:"condition,omitempty"`                 // Condition that the related person had
	Procedure        []FamilyMemberHistoryProcedure `json:"procedure,omitempty" bson:"procedure,omitempty"`                 // Procedures that the related person had
}

func (r *FamilyMemberHistory) Validate() error {
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
	var emptyString string
	if r.Status == emptyString {
		return fmt.Errorf("field 'Status' is required")
	}
	if r.DataAbsentReason != nil {
		if err := r.DataAbsentReason.Validate(); err != nil {
			return fmt.Errorf("DataAbsentReason: %w", err)
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
	if r.Relationship == nil {
		return fmt.Errorf("field 'Relationship' is required")
	}
	if r.Relationship != nil {
		if err := r.Relationship.Validate(); err != nil {
			return fmt.Errorf("Relationship: %w", err)
		}
	}
	if r.Sex != nil {
		if err := r.Sex.Validate(); err != nil {
			return fmt.Errorf("Sex: %w", err)
		}
	}
	if r.BornPeriod != nil {
		if err := r.BornPeriod.Validate(); err != nil {
			return fmt.Errorf("BornPeriod: %w", err)
		}
	}
	if r.AgeAge != nil {
		if err := r.AgeAge.Validate(); err != nil {
			return fmt.Errorf("AgeAge: %w", err)
		}
	}
	if r.AgeRange != nil {
		if err := r.AgeRange.Validate(); err != nil {
			return fmt.Errorf("AgeRange: %w", err)
		}
	}
	if r.DeceasedAge != nil {
		if err := r.DeceasedAge.Validate(); err != nil {
			return fmt.Errorf("DeceasedAge: %w", err)
		}
	}
	if r.DeceasedRange != nil {
		if err := r.DeceasedRange.Validate(); err != nil {
			return fmt.Errorf("DeceasedRange: %w", err)
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
	for i, item := range r.Condition {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Condition[%d]: %w", i, err)
		}
	}
	for i, item := range r.Procedure {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Procedure[%d]: %w", i, err)
		}
	}
	return nil
}

type FamilyMemberHistoryCondition struct {
	Id                 *string          `json:"id,omitempty" bson:"id,omitempty"`                                   // Unique id for inter-element referencing
	Code               *CodeableConcept `json:"code" bson:"code"`                                                   // Condition, allergy, or intolerance suffered by relation
	Outcome            *CodeableConcept `json:"outcome,omitempty" bson:"outcome,omitempty"`                         // deceased | permanent disability | etc
	ContributedToDeath bool             `json:"contributedToDeath,omitempty" bson:"contributed_to_death,omitempty"` // Whether the condition contributed to the cause of death
	OnsetAge           *Age             `json:"onsetAge,omitempty" bson:"onset_age,omitempty"`                      // When condition first manifested
	OnsetRange         *Range           `json:"onsetRange,omitempty" bson:"onset_range,omitempty"`                  // When condition first manifested
	OnsetPeriod        *Period          `json:"onsetPeriod,omitempty" bson:"onset_period,omitempty"`                // When condition first manifested
	OnsetString        *string          `json:"onsetString,omitempty" bson:"onset_string,omitempty"`                // When condition first manifested
	Note               []Annotation     `json:"note,omitempty" bson:"note,omitempty"`                               // Extra information about condition
}

func (r *FamilyMemberHistoryCondition) Validate() error {
	if r.Code == nil {
		return fmt.Errorf("field 'Code' is required")
	}
	if r.Code != nil {
		if err := r.Code.Validate(); err != nil {
			return fmt.Errorf("Code: %w", err)
		}
	}
	if r.Outcome != nil {
		if err := r.Outcome.Validate(); err != nil {
			return fmt.Errorf("Outcome: %w", err)
		}
	}
	if r.OnsetAge != nil {
		if err := r.OnsetAge.Validate(); err != nil {
			return fmt.Errorf("OnsetAge: %w", err)
		}
	}
	if r.OnsetRange != nil {
		if err := r.OnsetRange.Validate(); err != nil {
			return fmt.Errorf("OnsetRange: %w", err)
		}
	}
	if r.OnsetPeriod != nil {
		if err := r.OnsetPeriod.Validate(); err != nil {
			return fmt.Errorf("OnsetPeriod: %w", err)
		}
	}
	for i, item := range r.Note {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Note[%d]: %w", i, err)
		}
	}
	return nil
}

type FamilyMemberHistoryProcedure struct {
	Id                 *string          `json:"id,omitempty" bson:"id,omitempty"`                                   // Unique id for inter-element referencing
	Code               *CodeableConcept `json:"code" bson:"code"`                                                   // Procedures performed on the related person
	Outcome            *CodeableConcept `json:"outcome,omitempty" bson:"outcome,omitempty"`                         // What happened following the procedure
	ContributedToDeath bool             `json:"contributedToDeath,omitempty" bson:"contributed_to_death,omitempty"` // Whether the procedure contributed to the cause of death
	PerformedAge       *Age             `json:"performedAge,omitempty" bson:"performed_age,omitempty"`              // When the procedure was performed
	PerformedRange     *Range           `json:"performedRange,omitempty" bson:"performed_range,omitempty"`          // When the procedure was performed
	PerformedPeriod    *Period          `json:"performedPeriod,omitempty" bson:"performed_period,omitempty"`        // When the procedure was performed
	PerformedString    *string          `json:"performedString,omitempty" bson:"performed_string,omitempty"`        // When the procedure was performed
	PerformedDateTime  *string          `json:"performedDateTime,omitempty" bson:"performed_date_time,omitempty"`   // When the procedure was performed
	Note               []Annotation     `json:"note,omitempty" bson:"note,omitempty"`                               // Extra information about the procedure
}

func (r *FamilyMemberHistoryProcedure) Validate() error {
	if r.Code == nil {
		return fmt.Errorf("field 'Code' is required")
	}
	if r.Code != nil {
		if err := r.Code.Validate(); err != nil {
			return fmt.Errorf("Code: %w", err)
		}
	}
	if r.Outcome != nil {
		if err := r.Outcome.Validate(); err != nil {
			return fmt.Errorf("Outcome: %w", err)
		}
	}
	if r.PerformedAge != nil {
		if err := r.PerformedAge.Validate(); err != nil {
			return fmt.Errorf("PerformedAge: %w", err)
		}
	}
	if r.PerformedRange != nil {
		if err := r.PerformedRange.Validate(); err != nil {
			return fmt.Errorf("PerformedRange: %w", err)
		}
	}
	if r.PerformedPeriod != nil {
		if err := r.PerformedPeriod.Validate(); err != nil {
			return fmt.Errorf("PerformedPeriod: %w", err)
		}
	}
	for i, item := range r.Note {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Note[%d]: %w", i, err)
		}
	}
	return nil
}
