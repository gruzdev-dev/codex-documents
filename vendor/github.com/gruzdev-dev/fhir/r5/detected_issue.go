package models

import (
	"encoding/json"
	"fmt"
)

// Indicates an actual or potential clinical issue with or between one or more active or proposed clinical actions for a patient; e.g. Drug-drug interaction, Ineffective treatment frequency, Procedure-condition conflict, gaps in care, etc.
type DetectedIssue struct {
	Id                 *string                   `json:"id,omitempty" bson:"id,omitempty"`                                   // Logical id of this artifact
	Meta               *Meta                     `json:"meta,omitempty" bson:"meta,omitempty"`                               // Metadata about the resource
	ImplicitRules      *string                   `json:"implicitRules,omitempty" bson:"implicit_rules,omitempty"`            // A set of rules under which this content was created
	Language           *string                   `json:"language,omitempty" bson:"language,omitempty"`                       // Language of the resource content
	Text               *Narrative                `json:"text,omitempty" bson:"text,omitempty"`                               // Text summary of the resource, for human interpretation
	Contained          []json.RawMessage         `json:"contained,omitempty" bson:"contained,omitempty"`                     // Contained, inline Resources
	Identifier         []Identifier              `json:"identifier,omitempty" bson:"identifier,omitempty"`                   // Business identifier for detected issue
	Status             string                    `json:"status" bson:"status"`                                               // preliminary | final | entered-in-error | unknown | mitigated | processing-error
	Category           []CodeableConcept         `json:"category,omitempty" bson:"category,omitempty"`                       // High level categorization of detected issue
	Code               *CodeableConcept          `json:"code,omitempty" bson:"code,omitempty"`                               // Specific type of detected issue, e.g. drug-drug, duplicate therapy, etc
	Severity           *CodeableConcept          `json:"severity,omitempty" bson:"severity,omitempty"`                       // high | moderate | low
	Subject            *Reference                `json:"subject,omitempty" bson:"subject,omitempty"`                         // Associated subject
	Encounter          *Reference                `json:"encounter,omitempty" bson:"encounter,omitempty"`                     // Encounter the detected issue is part of
	IdentifiedDateTime *string                   `json:"identifiedDateTime,omitempty" bson:"identified_date_time,omitempty"` // When detected issue occurred/is occurring
	IdentifiedPeriod   *Period                   `json:"identifiedPeriod,omitempty" bson:"identified_period,omitempty"`      // When detected issue occurred/is occurring
	IdentifiedTiming   *Timing                   `json:"identifiedTiming,omitempty" bson:"identified_timing,omitempty"`      // When detected issue occurred/is occurring
	Author             *Reference                `json:"author,omitempty" bson:"author,omitempty"`                           // The provider or device that identified the issue
	Implicated         []Reference               `json:"implicated,omitempty" bson:"implicated,omitempty"`                   // Problem resource
	Evidence           []DetectedIssueEvidence   `json:"evidence,omitempty" bson:"evidence,omitempty"`                       // Supporting evidence
	Detail             *string                   `json:"detail,omitempty" bson:"detail,omitempty"`                           // Description and context
	Reference          *string                   `json:"reference,omitempty" bson:"reference,omitempty"`                     // Authority for issue
	QualityOfEvidence  *CodeableConcept          `json:"qualityOfEvidence,omitempty" bson:"quality_of_evidence,omitempty"`   // The quality of the evidence supporting the detected issue
	ExpectedOnsetType  *CodeableConcept          `json:"expectedOnsetType,omitempty" bson:"expected_onset_type,omitempty"`   // Time frame of clinical effect
	MedicationClass    []CodeableConcept         `json:"medicationClass,omitempty" bson:"medication_class,omitempty"`        // What medication class
	ManagementCode     *CodeableConcept          `json:"managementCode,omitempty" bson:"management_code,omitempty"`          // Importance of taking action on the issue
	Mitigation         []DetectedIssueMitigation `json:"mitigation,omitempty" bson:"mitigation,omitempty"`                   // Step taken to address
}

func (r *DetectedIssue) Validate() error {
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
	for i, item := range r.Category {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Category[%d]: %w", i, err)
		}
	}
	if r.Code != nil {
		if err := r.Code.Validate(); err != nil {
			return fmt.Errorf("Code: %w", err)
		}
	}
	if r.Severity != nil {
		if err := r.Severity.Validate(); err != nil {
			return fmt.Errorf("Severity: %w", err)
		}
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
	if r.IdentifiedPeriod != nil {
		if err := r.IdentifiedPeriod.Validate(); err != nil {
			return fmt.Errorf("IdentifiedPeriod: %w", err)
		}
	}
	if r.IdentifiedTiming != nil {
		if err := r.IdentifiedTiming.Validate(); err != nil {
			return fmt.Errorf("IdentifiedTiming: %w", err)
		}
	}
	if r.Author != nil {
		if err := r.Author.Validate(); err != nil {
			return fmt.Errorf("Author: %w", err)
		}
	}
	for i, item := range r.Implicated {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Implicated[%d]: %w", i, err)
		}
	}
	for i, item := range r.Evidence {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Evidence[%d]: %w", i, err)
		}
	}
	if r.QualityOfEvidence != nil {
		if err := r.QualityOfEvidence.Validate(); err != nil {
			return fmt.Errorf("QualityOfEvidence: %w", err)
		}
	}
	if r.ExpectedOnsetType != nil {
		if err := r.ExpectedOnsetType.Validate(); err != nil {
			return fmt.Errorf("ExpectedOnsetType: %w", err)
		}
	}
	for i, item := range r.MedicationClass {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("MedicationClass[%d]: %w", i, err)
		}
	}
	if r.ManagementCode != nil {
		if err := r.ManagementCode.Validate(); err != nil {
			return fmt.Errorf("ManagementCode: %w", err)
		}
	}
	for i, item := range r.Mitigation {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Mitigation[%d]: %w", i, err)
		}
	}
	return nil
}

type DetectedIssueEvidence struct {
	Id     *string           `json:"id,omitempty" bson:"id,omitempty"`         // Unique id for inter-element referencing
	Code   []CodeableConcept `json:"code,omitempty" bson:"code,omitempty"`     // Manifestation
	Detail []Reference       `json:"detail,omitempty" bson:"detail,omitempty"` // Supporting information
}

func (r *DetectedIssueEvidence) Validate() error {
	for i, item := range r.Code {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Code[%d]: %w", i, err)
		}
	}
	for i, item := range r.Detail {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Detail[%d]: %w", i, err)
		}
	}
	return nil
}

type DetectedIssueMitigation struct {
	Id     *string          `json:"id,omitempty" bson:"id,omitempty"`         // Unique id for inter-element referencing
	Action *CodeableConcept `json:"action" bson:"action"`                     // What mitigation?
	Date   *string          `json:"date,omitempty" bson:"date,omitempty"`     // Date committed
	Author *Reference       `json:"author,omitempty" bson:"author,omitempty"` // Who is committing?
	Note   []Annotation     `json:"note,omitempty" bson:"note,omitempty"`     // Additional notes about the mitigation
}

func (r *DetectedIssueMitigation) Validate() error {
	if r.Action == nil {
		return fmt.Errorf("field 'Action' is required")
	}
	if r.Action != nil {
		if err := r.Action.Validate(); err != nil {
			return fmt.Errorf("Action: %w", err)
		}
	}
	if r.Author != nil {
		if err := r.Author.Validate(); err != nil {
			return fmt.Errorf("Author: %w", err)
		}
	}
	for i, item := range r.Note {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Note[%d]: %w", i, err)
		}
	}
	return nil
}
