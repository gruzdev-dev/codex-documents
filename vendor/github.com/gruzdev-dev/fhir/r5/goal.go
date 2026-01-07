package models

import (
	"encoding/json"
	"fmt"
)

// Describes the intended objective(s) for a patient, group, or organizational care.  Examples include a patient's weight loss, restoration of an activity of daily living for a patient, obtaining herd immunity via immunization for a group, meeting a process improvement objective for an organization, etc.
type Goal struct {
	ResourceType         string            `json:"resourceType" bson:"resource_type"`                                      // Type of resource
	Id                   *string           `json:"id,omitempty" bson:"id,omitempty"`                                       // Logical id of this artifact
	Meta                 *Meta             `json:"meta,omitempty" bson:"meta,omitempty"`                                   // Metadata about the resource
	ImplicitRules        *string           `json:"implicitRules,omitempty" bson:"implicit_rules,omitempty"`                // A set of rules under which this content was created
	Language             *string           `json:"language,omitempty" bson:"language,omitempty"`                           // Language of the resource content
	Text                 *Narrative        `json:"text,omitempty" bson:"text,omitempty"`                                   // Text summary of the resource, for human interpretation
	Contained            []json.RawMessage `json:"contained,omitempty" bson:"contained,omitempty"`                         // Contained, inline Resources
	Identifier           []Identifier      `json:"identifier,omitempty" bson:"identifier,omitempty"`                       // External Ids for this goal
	LifecycleStatus      string            `json:"lifecycleStatus" bson:"lifecycle_status"`                                // proposed | planned | accepted | active | on-hold | completed | cancelled | entered-in-error | rejected
	AchievementStatus    *CodeableConcept  `json:"achievementStatus,omitempty" bson:"achievement_status,omitempty"`        // in-progress | improving | worsening | no-change | achieved | sustaining | not-achieved | no-progress | not-attainable
	Category             []CodeableConcept `json:"category,omitempty" bson:"category,omitempty"`                           // E.g. Treatment, dietary, behavioral, etc
	Continuous           *bool             `json:"continuous,omitempty" bson:"continuous,omitempty"`                       // After meeting the goal, ongoing activity is needed to sustain the goal objective
	Priority             *CodeableConcept  `json:"priority,omitempty" bson:"priority,omitempty"`                           // high-priority | medium-priority | low-priority
	Description          *CodeableConcept  `json:"description" bson:"description"`                                         // Code or text describing goal
	Subject              *Reference        `json:"subject" bson:"subject"`                                                 // Who this goal is intended for
	StartDate            *string           `json:"startDate,omitempty" bson:"start_date,omitempty"`                        // When goal pursuit begins
	StartCodeableConcept *CodeableConcept  `json:"startCodeableConcept,omitempty" bson:"start_codeable_concept,omitempty"` // When goal pursuit begins
	Acceptance           []GoalAcceptance  `json:"acceptance,omitempty" bson:"acceptance,omitempty"`                       // Individual acceptance of goal
	Target               []GoalTarget      `json:"target,omitempty" bson:"target,omitempty"`                               // Target outcome for the goal
	StatusDate           *string           `json:"statusDate,omitempty" bson:"status_date,omitempty"`                      // When goal achievment status took effect
	StatusReason         []CodeableConcept `json:"statusReason,omitempty" bson:"status_reason,omitempty"`                  // Reason for current lifecycle status
	Recorder             *Reference        `json:"recorder,omitempty" bson:"recorder,omitempty"`                           // Who recorded the goal
	Source               *Reference        `json:"source,omitempty" bson:"source,omitempty"`                               // Who's responsible for creating Goal?
	Addresses            []Reference       `json:"addresses,omitempty" bson:"addresses,omitempty"`                         // Issues addressed by this goal
	Note                 []Annotation      `json:"note,omitempty" bson:"note,omitempty"`                                   // Comments about the goal
}

func (r *Goal) Validate() error {
	if r.ResourceType != "Goal" {
		return fmt.Errorf("invalid resourceType: expected 'Goal', got '%s'", r.ResourceType)
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
	if r.LifecycleStatus == emptyString {
		return fmt.Errorf("field 'LifecycleStatus' is required")
	}
	if r.AchievementStatus != nil {
		if err := r.AchievementStatus.Validate(); err != nil {
			return fmt.Errorf("AchievementStatus: %w", err)
		}
	}
	for i, item := range r.Category {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Category[%d]: %w", i, err)
		}
	}
	if r.Priority != nil {
		if err := r.Priority.Validate(); err != nil {
			return fmt.Errorf("Priority: %w", err)
		}
	}
	if r.Description == nil {
		return fmt.Errorf("field 'Description' is required")
	}
	if r.Description != nil {
		if err := r.Description.Validate(); err != nil {
			return fmt.Errorf("Description: %w", err)
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
	if r.StartCodeableConcept != nil {
		if err := r.StartCodeableConcept.Validate(); err != nil {
			return fmt.Errorf("StartCodeableConcept: %w", err)
		}
	}
	for i, item := range r.Acceptance {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Acceptance[%d]: %w", i, err)
		}
	}
	for i, item := range r.Target {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Target[%d]: %w", i, err)
		}
	}
	for i, item := range r.StatusReason {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("StatusReason[%d]: %w", i, err)
		}
	}
	if r.Recorder != nil {
		if err := r.Recorder.Validate(); err != nil {
			return fmt.Errorf("Recorder: %w", err)
		}
	}
	if r.Source != nil {
		if err := r.Source.Validate(); err != nil {
			return fmt.Errorf("Source: %w", err)
		}
	}
	for i, item := range r.Addresses {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Addresses[%d]: %w", i, err)
		}
	}
	for i, item := range r.Note {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Note[%d]: %w", i, err)
		}
	}
	return nil
}

type GoalTarget struct {
	Id                    *string          `json:"id,omitempty" bson:"id,omitempty"`                                         // Unique id for inter-element referencing
	Measure               *CodeableConcept `json:"measure,omitempty" bson:"measure,omitempty"`                               // The parameter whose value is being tracked
	DetailQuantity        *Quantity        `json:"detailQuantity,omitempty" bson:"detail_quantity,omitempty"`                // The target value to be achieved
	DetailRange           *Range           `json:"detailRange,omitempty" bson:"detail_range,omitempty"`                      // The target value to be achieved
	DetailCodeableConcept *CodeableConcept `json:"detailCodeableConcept,omitempty" bson:"detail_codeable_concept,omitempty"` // The target value to be achieved
	DetailString          *string          `json:"detailString,omitempty" bson:"detail_string,omitempty"`                    // The target value to be achieved
	DetailBoolean         *bool            `json:"detailBoolean,omitempty" bson:"detail_boolean,omitempty"`                  // The target value to be achieved
	DetailInteger         *int             `json:"detailInteger,omitempty" bson:"detail_integer,omitempty"`                  // The target value to be achieved
	DetailRatio           *Ratio           `json:"detailRatio,omitempty" bson:"detail_ratio,omitempty"`                      // The target value to be achieved
	DueDate               *string          `json:"dueDate,omitempty" bson:"due_date,omitempty"`                              // Reach goal on or before
	DueDuration           *Duration        `json:"dueDuration,omitempty" bson:"due_duration,omitempty"`                      // Reach goal on or before
}

func (r *GoalTarget) Validate() error {
	if r.Measure != nil {
		if err := r.Measure.Validate(); err != nil {
			return fmt.Errorf("Measure: %w", err)
		}
	}
	if r.DetailQuantity != nil {
		if err := r.DetailQuantity.Validate(); err != nil {
			return fmt.Errorf("DetailQuantity: %w", err)
		}
	}
	if r.DetailRange != nil {
		if err := r.DetailRange.Validate(); err != nil {
			return fmt.Errorf("DetailRange: %w", err)
		}
	}
	if r.DetailCodeableConcept != nil {
		if err := r.DetailCodeableConcept.Validate(); err != nil {
			return fmt.Errorf("DetailCodeableConcept: %w", err)
		}
	}
	if r.DetailRatio != nil {
		if err := r.DetailRatio.Validate(); err != nil {
			return fmt.Errorf("DetailRatio: %w", err)
		}
	}
	if r.DueDuration != nil {
		if err := r.DueDuration.Validate(); err != nil {
			return fmt.Errorf("DueDuration: %w", err)
		}
	}
	return nil
}

type GoalAcceptance struct {
	Id          *string          `json:"id,omitempty" bson:"id,omitempty"`             // Unique id for inter-element referencing
	Participant *Reference       `json:"participant" bson:"participant"`               // Individual or organization whose acceptance is reflected
	Status      *string          `json:"status,omitempty" bson:"status,omitempty"`     // agree | disagree | pending
	Priority    *CodeableConcept `json:"priority,omitempty" bson:"priority,omitempty"` // Priority of goal for individual
}

func (r *GoalAcceptance) Validate() error {
	if r.Participant == nil {
		return fmt.Errorf("field 'Participant' is required")
	}
	if r.Participant != nil {
		if err := r.Participant.Validate(); err != nil {
			return fmt.Errorf("Participant: %w", err)
		}
	}
	if r.Priority != nil {
		if err := r.Priority.Validate(); err != nil {
			return fmt.Errorf("Priority: %w", err)
		}
	}
	return nil
}
