package models

import (
	"encoding/json"
	"fmt"
)

// A slot of time on a schedule that may be available for booking appointments.
type Slot struct {
	ResourceType    string              `json:"resourceType" bson:"resource_type"`                           // Type of resource
	Id              *string             `json:"id,omitempty" bson:"id,omitempty"`                            // Logical id of this artifact
	Meta            *Meta               `json:"meta,omitempty" bson:"meta,omitempty"`                        // Metadata about the resource
	ImplicitRules   *string             `json:"implicitRules,omitempty" bson:"implicit_rules,omitempty"`     // A set of rules under which this content was created
	Language        *string             `json:"language,omitempty" bson:"language,omitempty"`                // Language of the resource content
	Text            *Narrative          `json:"text,omitempty" bson:"text,omitempty"`                        // Text summary of the resource, for human interpretation
	Contained       []json.RawMessage   `json:"contained,omitempty" bson:"contained,omitempty"`              // Contained, inline Resources
	Identifier      []Identifier        `json:"identifier,omitempty" bson:"identifier,omitempty"`            // External Ids for this item
	ServiceCategory []CodeableConcept   `json:"serviceCategory,omitempty" bson:"service_category,omitempty"` // A broad categorization of the service that is to be performed during this appointment
	ServiceType     []CodeableReference `json:"serviceType,omitempty" bson:"service_type,omitempty"`         // The type of appointments that can be booked into this slot (ideally this would be an identifiable service - which is at a location, rather than the location itself). If provided then this overrides the value provided on the Schedule resource
	Specialty       []CodeableConcept   `json:"specialty,omitempty" bson:"specialty,omitempty"`              // The specialty of a practitioner that would be required to perform the service requested in this appointment
	AppointmentType []CodeableConcept   `json:"appointmentType,omitempty" bson:"appointment_type,omitempty"` // The style of appointment or patient that may be booked in the slot (not service type)
	Schedule        *Reference          `json:"schedule" bson:"schedule"`                                    // The schedule resource that this slot defines an interval of status information
	Status          string              `json:"status" bson:"status"`                                        // busy | free | busy-unavailable | busy-tentative | entered-in-error
	Start           string              `json:"start" bson:"start"`                                          // Date/Time that the slot is to begin
	End             string              `json:"end" bson:"end"`                                              // Date/Time that the slot is to conclude
	Overbooked      bool                `json:"overbooked,omitempty" bson:"overbooked,omitempty"`            // This slot has already been overbooked, appointments are unlikely to be accepted for this time
	Comment         *string             `json:"comment,omitempty" bson:"comment,omitempty"`                  // Comments on the slot to describe any extended information. Such as custom constraints on the slot
}

func (r *Slot) Validate() error {
	if r.ResourceType != "Slot" {
		return fmt.Errorf("invalid resourceType: expected 'Slot', got '%s'", r.ResourceType)
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
	for i, item := range r.AppointmentType {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("AppointmentType[%d]: %w", i, err)
		}
	}
	if r.Schedule == nil {
		return fmt.Errorf("field 'Schedule' is required")
	}
	if r.Schedule != nil {
		if err := r.Schedule.Validate(); err != nil {
			return fmt.Errorf("Schedule: %w", err)
		}
	}
	var emptyString string
	if r.Status == emptyString {
		return fmt.Errorf("field 'Status' is required")
	}
	if r.Start == emptyString {
		return fmt.Errorf("field 'Start' is required")
	}
	if r.End == emptyString {
		return fmt.Errorf("field 'End' is required")
	}
	return nil
}
