package models

import (
	"encoding/json"
	"fmt"
)

// Describes a physiological or technical alert condition report originated by a device.  The DeviceAlert resource is derived from the ISO/IEEE 11073-10201 Domain Information Model standard, but is more widely applicable.
type DeviceAlert struct {
	ResourceType       string                   `json:"resourceType" bson:"resource_type"`                                  // Type of resource
	Id                 *string                  `json:"id,omitempty" bson:"id,omitempty"`                                   // Logical id of this artifact
	Meta               *Meta                    `json:"meta,omitempty" bson:"meta,omitempty"`                               // Metadata about the resource
	ImplicitRules      *string                  `json:"implicitRules,omitempty" bson:"implicit_rules,omitempty"`            // A set of rules under which this content was created
	Language           *string                  `json:"language,omitempty" bson:"language,omitempty"`                       // Language of the resource content
	Text               *Narrative               `json:"text,omitempty" bson:"text,omitempty"`                               // Text summary of the resource, for human interpretation
	Contained          []json.RawMessage        `json:"contained,omitempty" bson:"contained,omitempty"`                     // Contained, inline Resources
	Identifier         []Identifier             `json:"identifier,omitempty" bson:"identifier,omitempty"`                   // Business identifier for this device alert
	Procedure          []Reference              `json:"procedure,omitempty" bson:"procedure,omitempty"`                     // Procedure during which the alert occurred
	Status             string                   `json:"status" bson:"status"`                                               // in-progress | completed | entered-in-error | unknown
	Category           []CodeableConcept        `json:"category,omitempty" bson:"category,omitempty"`                       // High level categorization of device alert
	Type               *CodeableConcept         `json:"type,omitempty" bson:"type,omitempty"`                               // physiological | technical
	Priority           *CodeableConcept         `json:"priority,omitempty" bson:"priority,omitempty"`                       // high | medium | low | info
	Code               *CodeableConcept         `json:"code" bson:"code"`                                                   // The meaning of the alert
	Subject            *Reference               `json:"subject" bson:"subject"`                                             // Who or what the alert is about
	Encounter          *Reference               `json:"encounter,omitempty" bson:"encounter,omitempty"`                     // Encounter during which the alert condition occurred
	Presence           bool                     `json:"presence" bson:"presence"`                                           // Whether the alert condition is currently active
	OccurrenceDateTime *string                  `json:"occurrenceDateTime,omitempty" bson:"occurrence_date_time,omitempty"` // When the alert condition occurred/is occurring
	OccurrencePeriod   *Period                  `json:"occurrencePeriod,omitempty" bson:"occurrence_period,omitempty"`      // When the alert condition occurred/is occurring
	Device             *Reference               `json:"device,omitempty" bson:"device,omitempty"`                           // The Device (or DeviceMetric) that detected the alert condition
	Acknowledged       bool                     `json:"acknowledged,omitempty" bson:"acknowledged,omitempty"`               // Whether the alert condition has been acknowledged
	AcknowledgedBy     *Reference               `json:"acknowledgedBy,omitempty" bson:"acknowledged_by,omitempty"`          // Who acknowledged the alert condition
	Location           *Reference               `json:"location,omitempty" bson:"location,omitempty"`                       // Location of the subject when the alert was raised
	DerivedFrom        []DeviceAlertDerivedFrom `json:"derivedFrom,omitempty" bson:"derived_from,omitempty"`                // The value causing the alert condition
	Label              *string                  `json:"label,omitempty" bson:"label,omitempty"`                             // Text to be displayed for the alert condition
	Signal             []DeviceAlertSignal      `json:"signal,omitempty" bson:"signal,omitempty"`                           // Annunciation or notification of the alert condition
}

func (r *DeviceAlert) Validate() error {
	if r.ResourceType != "DeviceAlert" {
		return fmt.Errorf("invalid resourceType: expected 'DeviceAlert', got '%s'", r.ResourceType)
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
	for i, item := range r.Procedure {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Procedure[%d]: %w", i, err)
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
	if r.Type != nil {
		if err := r.Type.Validate(); err != nil {
			return fmt.Errorf("Type: %w", err)
		}
	}
	if r.Priority != nil {
		if err := r.Priority.Validate(); err != nil {
			return fmt.Errorf("Priority: %w", err)
		}
	}
	if r.Code == nil {
		return fmt.Errorf("field 'Code' is required")
	}
	if r.Code != nil {
		if err := r.Code.Validate(); err != nil {
			return fmt.Errorf("Code: %w", err)
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
	if r.OccurrencePeriod != nil {
		if err := r.OccurrencePeriod.Validate(); err != nil {
			return fmt.Errorf("OccurrencePeriod: %w", err)
		}
	}
	if r.Device != nil {
		if err := r.Device.Validate(); err != nil {
			return fmt.Errorf("Device: %w", err)
		}
	}
	if r.AcknowledgedBy != nil {
		if err := r.AcknowledgedBy.Validate(); err != nil {
			return fmt.Errorf("AcknowledgedBy: %w", err)
		}
	}
	if r.Location != nil {
		if err := r.Location.Validate(); err != nil {
			return fmt.Errorf("Location: %w", err)
		}
	}
	for i, item := range r.DerivedFrom {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("DerivedFrom[%d]: %w", i, err)
		}
	}
	for i, item := range r.Signal {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Signal[%d]: %w", i, err)
		}
	}
	return nil
}

type DeviceAlertDerivedFrom struct {
	Id          *string    `json:"id,omitempty" bson:"id,omitempty"`               // Unique id for inter-element referencing
	Observation *Reference `json:"observation" bson:"observation"`                 // The Observation having a value causing the alert condition
	Component   *Coding    `json:"component,omitempty" bson:"component,omitempty"` // The Observation.component having a value causing the alert condition
	Limit       *Range     `json:"limit,omitempty" bson:"limit,omitempty"`         // The boundaries beyond which a value was detected to cause the alert condition
}

func (r *DeviceAlertDerivedFrom) Validate() error {
	if r.Observation == nil {
		return fmt.Errorf("field 'Observation' is required")
	}
	if r.Observation != nil {
		if err := r.Observation.Validate(); err != nil {
			return fmt.Errorf("Observation: %w", err)
		}
	}
	if r.Component != nil {
		if err := r.Component.Validate(); err != nil {
			return fmt.Errorf("Component: %w", err)
		}
	}
	if r.Limit != nil {
		if err := r.Limit.Validate(); err != nil {
			return fmt.Errorf("Limit: %w", err)
		}
	}
	return nil
}

type DeviceAlertSignal struct {
	Id              *string            `json:"id,omitempty" bson:"id,omitempty"`                       // Unique id for inter-element referencing
	ActivationState *CodeableConcept   `json:"activationState" bson:"activation_state"`                // on | off | paused
	Presence        *CodeableConcept   `json:"presence,omitempty" bson:"presence,omitempty"`           // on | latched | off | ack
	Annunciator     *CodeableReference `json:"annunciator,omitempty" bson:"annunciator,omitempty"`     // Where the signal is being annunciated
	Manifestation   *CodeableConcept   `json:"manifestation,omitempty" bson:"manifestation,omitempty"` // How the signal is being annunciated
	Type            []CodeableConcept  `json:"type,omitempty" bson:"type,omitempty"`                   // Characteristics of the signal manifestation
	Indication      *Period            `json:"indication,omitempty" bson:"indication,omitempty"`       // When the signal was being annunciated
}

func (r *DeviceAlertSignal) Validate() error {
	if r.ActivationState == nil {
		return fmt.Errorf("field 'ActivationState' is required")
	}
	if r.ActivationState != nil {
		if err := r.ActivationState.Validate(); err != nil {
			return fmt.Errorf("ActivationState: %w", err)
		}
	}
	if r.Presence != nil {
		if err := r.Presence.Validate(); err != nil {
			return fmt.Errorf("Presence: %w", err)
		}
	}
	if r.Annunciator != nil {
		if err := r.Annunciator.Validate(); err != nil {
			return fmt.Errorf("Annunciator: %w", err)
		}
	}
	if r.Manifestation != nil {
		if err := r.Manifestation.Validate(); err != nil {
			return fmt.Errorf("Manifestation: %w", err)
		}
	}
	for i, item := range r.Type {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Type[%d]: %w", i, err)
		}
	}
	if r.Indication != nil {
		if err := r.Indication.Validate(); err != nil {
			return fmt.Errorf("Indication: %w", err)
		}
	}
	return nil
}
