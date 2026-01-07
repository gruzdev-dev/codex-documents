package models

import (
	"encoding/json"
	"fmt"
)

// Describes the intention of how one or more practitioners intend to deliver care for a particular patient, group or community for a period of time, possibly limited to care for a specific condition or set of conditions.
type CarePlan struct {
	ResourceType   string              `json:"resourceType" bson:"resource_type"`                         // Type of resource
	Id             *string             `json:"id,omitempty" bson:"id,omitempty"`                          // Logical id of this artifact
	Meta           *Meta               `json:"meta,omitempty" bson:"meta,omitempty"`                      // Metadata about the resource
	ImplicitRules  *string             `json:"implicitRules,omitempty" bson:"implicit_rules,omitempty"`   // A set of rules under which this content was created
	Language       *string             `json:"language,omitempty" bson:"language,omitempty"`              // Language of the resource content
	Text           *Narrative          `json:"text,omitempty" bson:"text,omitempty"`                      // Text summary of the resource, for human interpretation
	Contained      []json.RawMessage   `json:"contained,omitempty" bson:"contained,omitempty"`            // Contained, inline Resources
	Identifier     []Identifier        `json:"identifier,omitempty" bson:"identifier,omitempty"`          // External Ids for this plan
	BasedOn        []Reference         `json:"basedOn,omitempty" bson:"based_on,omitempty"`               // Fulfills plan, proposal or order
	Replaces       []Reference         `json:"replaces,omitempty" bson:"replaces,omitempty"`              // CarePlan replaced by this CarePlan
	PartOf         []Reference         `json:"partOf,omitempty" bson:"part_of,omitempty"`                 // Part of referenced CarePlan
	Status         string              `json:"status" bson:"status"`                                      // draft | active | on-hold | entered-in-error | ended | completed | revoked | unknown
	Intent         string              `json:"intent" bson:"intent"`                                      // proposal | plan | order | option | directive
	Category       []CodeableConcept   `json:"category,omitempty" bson:"category,omitempty"`              // Type of plan
	Title          *string             `json:"title,omitempty" bson:"title,omitempty"`                    // Human-friendly name for the care plan
	Description    *string             `json:"description,omitempty" bson:"description,omitempty"`        // Summary of nature of plan
	Subject        *Reference          `json:"subject" bson:"subject"`                                    // Who the care plan is for
	Encounter      *Reference          `json:"encounter,omitempty" bson:"encounter,omitempty"`            // The Encounter during which this CarePlan was created
	Period         *Period             `json:"period,omitempty" bson:"period,omitempty"`                  // Time period plan covers
	Created        *string             `json:"created,omitempty" bson:"created,omitempty"`                // Date record was first recorded
	Custodian      *Reference          `json:"custodian,omitempty" bson:"custodian,omitempty"`            // Who is the designated responsible party
	Contributor    []Reference         `json:"contributor,omitempty" bson:"contributor,omitempty"`        // Who provided the content of the care plan
	CareTeam       []Reference         `json:"careTeam,omitempty" bson:"care_team,omitempty"`             // Who's involved in plan?
	Addresses      []CodeableReference `json:"addresses,omitempty" bson:"addresses,omitempty"`            // Health issues this plan addresses
	SupportingInfo []Reference         `json:"supportingInfo,omitempty" bson:"supporting_info,omitempty"` // Information considered as part of plan
	Goal           []Reference         `json:"goal,omitempty" bson:"goal,omitempty"`                      // Desired outcome of plan
	Activity       []CarePlanActivity  `json:"activity,omitempty" bson:"activity,omitempty"`              // Action to occur or has occurred as part of plan
	Note           []Annotation        `json:"note,omitempty" bson:"note,omitempty"`                      // Comments about the plan
}

func (r *CarePlan) Validate() error {
	if r.ResourceType != "CarePlan" {
		return fmt.Errorf("invalid resourceType: expected 'CarePlan', got '%s'", r.ResourceType)
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
	for i, item := range r.BasedOn {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("BasedOn[%d]: %w", i, err)
		}
	}
	for i, item := range r.Replaces {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Replaces[%d]: %w", i, err)
		}
	}
	for i, item := range r.PartOf {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("PartOf[%d]: %w", i, err)
		}
	}
	var emptyString string
	if r.Status == emptyString {
		return fmt.Errorf("field 'Status' is required")
	}
	if r.Intent == emptyString {
		return fmt.Errorf("field 'Intent' is required")
	}
	for i, item := range r.Category {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Category[%d]: %w", i, err)
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
	if r.Period != nil {
		if err := r.Period.Validate(); err != nil {
			return fmt.Errorf("Period: %w", err)
		}
	}
	if r.Custodian != nil {
		if err := r.Custodian.Validate(); err != nil {
			return fmt.Errorf("Custodian: %w", err)
		}
	}
	for i, item := range r.Contributor {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Contributor[%d]: %w", i, err)
		}
	}
	for i, item := range r.CareTeam {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("CareTeam[%d]: %w", i, err)
		}
	}
	for i, item := range r.Addresses {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Addresses[%d]: %w", i, err)
		}
	}
	for i, item := range r.SupportingInfo {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("SupportingInfo[%d]: %w", i, err)
		}
	}
	for i, item := range r.Goal {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Goal[%d]: %w", i, err)
		}
	}
	for i, item := range r.Activity {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Activity[%d]: %w", i, err)
		}
	}
	for i, item := range r.Note {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Note[%d]: %w", i, err)
		}
	}
	return nil
}

type CarePlanActivity struct {
	Id                       *string             `json:"id,omitempty" bson:"id,omitempty"`                                               // Unique id for inter-element referencing
	PerformedActivity        []CodeableReference `json:"performedActivity,omitempty" bson:"performed_activity,omitempty"`                // Activities that are completed or in progress (concept, or Appointment, Encounter, Procedure, etc.)
	Progress                 []Annotation        `json:"progress,omitempty" bson:"progress,omitempty"`                                   // Comments about the activity status/progress
	PlannedActivityReference *Reference          `json:"plannedActivityReference,omitempty" bson:"planned_activity_reference,omitempty"` // Activity that is intended to be part of the care plan
}

func (r *CarePlanActivity) Validate() error {
	for i, item := range r.PerformedActivity {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("PerformedActivity[%d]: %w", i, err)
		}
	}
	for i, item := range r.Progress {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Progress[%d]: %w", i, err)
		}
	}
	if r.PlannedActivityReference != nil {
		if err := r.PlannedActivityReference.Validate(); err != nil {
			return fmt.Errorf("PlannedActivityReference: %w", err)
		}
	}
	return nil
}
