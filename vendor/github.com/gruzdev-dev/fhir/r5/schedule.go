package models

import (
	"encoding/json"
	"fmt"
)

// A container for slots of time that may be available for booking appointments.
type Schedule struct {
	ResourceType    string              `json:"resourceType" bson:"resource_type"`                           // Type of resource
	Id              *string             `json:"id,omitempty" bson:"id,omitempty"`                            // Logical id of this artifact
	Meta            *Meta               `json:"meta,omitempty" bson:"meta,omitempty"`                        // Metadata about the resource
	ImplicitRules   *string             `json:"implicitRules,omitempty" bson:"implicit_rules,omitempty"`     // A set of rules under which this content was created
	Language        *string             `json:"language,omitempty" bson:"language,omitempty"`                // Language of the resource content
	Text            *Narrative          `json:"text,omitempty" bson:"text,omitempty"`                        // Text summary of the resource, for human interpretation
	Contained       []json.RawMessage   `json:"contained,omitempty" bson:"contained,omitempty"`              // Contained, inline Resources
	Identifier      []Identifier        `json:"identifier,omitempty" bson:"identifier,omitempty"`            // External Ids for this item
	Active          *bool               `json:"active,omitempty" bson:"active,omitempty"`                    // Whether this schedule is in active use
	ServiceCategory []CodeableConcept   `json:"serviceCategory,omitempty" bson:"service_category,omitempty"` // High-level category
	ServiceType     []CodeableReference `json:"serviceType,omitempty" bson:"service_type,omitempty"`         // Specific service
	Specialty       []CodeableConcept   `json:"specialty,omitempty" bson:"specialty,omitempty"`              // Type of specialty needed
	Name            *string             `json:"name,omitempty" bson:"name,omitempty"`                        // Human-readable label
	Actor           []Reference         `json:"actor" bson:"actor"`                                          // Resource(s) that availability information is being provided for
	PlanningHorizon *Period             `json:"planningHorizon,omitempty" bson:"planning_horizon,omitempty"` // Period of time covered by schedule
	Comment         *string             `json:"comment,omitempty" bson:"comment,omitempty"`                  // Comments on availability
}

func (r *Schedule) Validate() error {
	if r.ResourceType != "Schedule" {
		return fmt.Errorf("invalid resourceType: expected 'Schedule', got '%s'", r.ResourceType)
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
	for i, item := range r.ServiceCategory {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("ServiceCategory[%d]: %w", i, err)
		}
	}
	for i, item := range r.ServiceType {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("ServiceType[%d]: %w", i, err)
		}
	}
	for i, item := range r.Specialty {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Specialty[%d]: %w", i, err)
		}
	}
	if len(r.Actor) < 1 {
		return fmt.Errorf("field 'Actor' must have at least 1 elements")
	}
	for i, item := range r.Actor {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Actor[%d]: %w", i, err)
		}
	}
	if r.PlanningHorizon != nil {
		if err := r.PlanningHorizon.Validate(); err != nil {
			return fmt.Errorf("PlanningHorizon: %w", err)
		}
	}
	return nil
}
