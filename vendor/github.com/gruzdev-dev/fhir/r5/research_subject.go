package models

import (
	"encoding/json"
	"fmt"
)

// A ResearchSubject is a participant or object which is the recipient of investigative activities in a research study.
type ResearchSubject struct {
	ResourceType     string                            `json:"resourceType" bson:"resource_type"`                             // Type of resource
	Id               *string                           `json:"id,omitempty" bson:"id,omitempty"`                              // Logical id of this artifact
	Meta             *Meta                             `json:"meta,omitempty" bson:"meta,omitempty"`                          // Metadata about the resource
	ImplicitRules    *string                           `json:"implicitRules,omitempty" bson:"implicit_rules,omitempty"`       // A set of rules under which this content was created
	Language         *string                           `json:"language,omitempty" bson:"language,omitempty"`                  // Language of the resource content
	Text             *Narrative                        `json:"text,omitempty" bson:"text,omitempty"`                          // Text summary of the resource, for human interpretation
	Contained        []json.RawMessage                 `json:"contained,omitempty" bson:"contained,omitempty"`                // Contained, inline Resources
	Identifier       []Identifier                      `json:"identifier,omitempty" bson:"identifier,omitempty"`              // Business Identifier for research subject in a study
	Status           string                            `json:"status" bson:"status"`                                          // draft | active | retired | unknown
	Period           *Period                           `json:"period,omitempty" bson:"period,omitempty"`                      // Start and end of participation
	Study            *Reference                        `json:"study" bson:"study"`                                            // Study subject is part of
	Subject          *Reference                        `json:"subject" bson:"subject"`                                        // Who or what is part of study
	SubjectState     []ResearchSubjectSubjectState     `json:"subjectState,omitempty" bson:"subject_state,omitempty"`         // A duration in the lifecycle of the ResearchSubject within a ResearchStudy
	SubjectMilestone []ResearchSubjectSubjectMilestone `json:"subjectMilestone,omitempty" bson:"subject_milestone,omitempty"` // A significant event in the progress of a ResearchSubject
	ComparisonGroup  []CodeableReference               `json:"comparisonGroup,omitempty" bson:"comparison_group,omitempty"`   // A group to which the subject is assigned
	Consent          []Reference                       `json:"consent,omitempty" bson:"consent,omitempty"`                    // Agreement to participate in study
}

func (r *ResearchSubject) Validate() error {
	if r.ResourceType != "ResearchSubject" {
		return fmt.Errorf("invalid resourceType: expected 'ResearchSubject', got '%s'", r.ResourceType)
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
	var emptyString string
	if r.Status == emptyString {
		return fmt.Errorf("field 'Status' is required")
	}
	if r.Period != nil {
		if err := r.Period.Validate(); err != nil {
			return fmt.Errorf("Period: %w", err)
		}
	}
	if r.Study == nil {
		return fmt.Errorf("field 'Study' is required")
	}
	if r.Study != nil {
		if err := r.Study.Validate(); err != nil {
			return fmt.Errorf("Study: %w", err)
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
	for i, item := range r.SubjectState {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("SubjectState[%d]: %w", i, err)
		}
	}
	for i, item := range r.SubjectMilestone {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("SubjectMilestone[%d]: %w", i, err)
		}
	}
	for i, item := range r.ComparisonGroup {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("ComparisonGroup[%d]: %w", i, err)
		}
	}
	for i, item := range r.Consent {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Consent[%d]: %w", i, err)
		}
	}
	return nil
}

type ResearchSubjectSubjectState struct {
	Id        *string          `json:"id,omitempty" bson:"id,omitempty"`            // Unique id for inter-element referencing
	Code      *CodeableConcept `json:"code" bson:"code"`                            // candidate | in-prescreening | in-screening | eligible | ineligible | on-study | on-study-intervention | in-follow-up | off-study
	StartDate string           `json:"startDate" bson:"start_date"`                 // The date a research subject entered the given state
	EndDate   *string          `json:"endDate,omitempty" bson:"end_date,omitempty"` // The date a research subject exited or left the given state
	Reason    *CodeableConcept `json:"reason,omitempty" bson:"reason,omitempty"`    // State change reason
}

func (r *ResearchSubjectSubjectState) Validate() error {
	if r.Code == nil {
		return fmt.Errorf("field 'Code' is required")
	}
	if r.Code != nil {
		if err := r.Code.Validate(); err != nil {
			return fmt.Errorf("Code: %w", err)
		}
	}
	var emptyString string
	if r.StartDate == emptyString {
		return fmt.Errorf("field 'StartDate' is required")
	}
	if r.Reason != nil {
		if err := r.Reason.Validate(); err != nil {
			return fmt.Errorf("Reason: %w", err)
		}
	}
	return nil
}

type ResearchSubjectSubjectMilestone struct {
	Id        *string           `json:"id,omitempty" bson:"id,omitempty"` // Unique id for inter-element referencing
	Milestone *CodeableConcept  `json:"milestone" bson:"milestone"`
	Date      *string           `json:"date,omitempty" bson:"date,omitempty"` // The date/time when this milestone event was completed
	Reason    []CodeableConcept `json:"reason,omitempty" bson:"reason,omitempty"`
}

func (r *ResearchSubjectSubjectMilestone) Validate() error {
	if r.Milestone == nil {
		return fmt.Errorf("field 'Milestone' is required")
	}
	if r.Milestone != nil {
		if err := r.Milestone.Validate(); err != nil {
			return fmt.Errorf("Milestone: %w", err)
		}
	}
	for i, item := range r.Reason {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Reason[%d]: %w", i, err)
		}
	}
	return nil
}
