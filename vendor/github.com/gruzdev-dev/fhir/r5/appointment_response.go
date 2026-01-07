package models

import (
	"encoding/json"
	"fmt"
)

// A reply to an appointment request for a patient and/or practitioner(s), such as a confirmation or rejection.
type AppointmentResponse struct {
	ResourceType      string            `json:"resourceType" bson:"resource_type"`                            // Type of resource
	Id                *string           `json:"id,omitempty" bson:"id,omitempty"`                             // Logical id of this artifact
	Meta              *Meta             `json:"meta,omitempty" bson:"meta,omitempty"`                         // Metadata about the resource
	ImplicitRules     *string           `json:"implicitRules,omitempty" bson:"implicit_rules,omitempty"`      // A set of rules under which this content was created
	Language          *string           `json:"language,omitempty" bson:"language,omitempty"`                 // Language of the resource content
	Text              *Narrative        `json:"text,omitempty" bson:"text,omitempty"`                         // Text summary of the resource, for human interpretation
	Contained         []json.RawMessage `json:"contained,omitempty" bson:"contained,omitempty"`               // Contained, inline Resources
	Identifier        []Identifier      `json:"identifier,omitempty" bson:"identifier,omitempty"`             // External Ids for this item
	Appointment       *Reference        `json:"appointment" bson:"appointment"`                               // Appointment this response relates to
	ProposedNewTime   *bool             `json:"proposedNewTime,omitempty" bson:"proposed_new_time,omitempty"` // Indicator for a counter proposal
	Start             *string           `json:"start,omitempty" bson:"start,omitempty"`                       // Time from appointment, or requested new start time
	End               *string           `json:"end,omitempty" bson:"end,omitempty"`                           // Time from appointment, or requested new end time
	ParticipantType   []CodeableConcept `json:"participantType,omitempty" bson:"participant_type,omitempty"`  // Role of participant in the appointment
	Actor             *Reference        `json:"actor,omitempty" bson:"actor,omitempty"`                       // Person(s), Location, HealthcareService, or Device
	ParticipantStatus string            `json:"participantStatus" bson:"participant_status"`                  // accepted | declined | tentative | needs-action | entered-in-error
	Comment           *string           `json:"comment,omitempty" bson:"comment,omitempty"`                   // Additional comments
	Recurring         *bool             `json:"recurring,omitempty" bson:"recurring,omitempty"`               // This response is for all occurrences in a recurring request
	OccurrenceDate    *string           `json:"occurrenceDate,omitempty" bson:"occurrence_date,omitempty"`    // Original date within a recurring request
	RecurrenceId      *int              `json:"recurrenceId,omitempty" bson:"recurrence_id,omitempty"`        // The recurrence ID of the specific recurring request
}

func (r *AppointmentResponse) Validate() error {
	if r.ResourceType != "AppointmentResponse" {
		return fmt.Errorf("invalid resourceType: expected 'AppointmentResponse', got '%s'", r.ResourceType)
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
	if r.Appointment == nil {
		return fmt.Errorf("field 'Appointment' is required")
	}
	if r.Appointment != nil {
		if err := r.Appointment.Validate(); err != nil {
			return fmt.Errorf("Appointment: %w", err)
		}
	}
	for i, item := range r.ParticipantType {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("ParticipantType[%d]: %w", i, err)
		}
	}
	if r.Actor != nil {
		if err := r.Actor.Validate(); err != nil {
			return fmt.Errorf("Actor: %w", err)
		}
	}
	var emptyString string
	if r.ParticipantStatus == emptyString {
		return fmt.Errorf("field 'ParticipantStatus' is required")
	}
	return nil
}
