package models

import (
	"fmt"
)

// TriggerDefinition Type: A description of a triggering event. Triggering events can be named events, data events, or periodic, as determined by the type element.
type TriggerDefinition struct {
	Id                *string           `json:"id,omitempty" bson:"id,omitempty"`                                // Unique id for inter-element referencing
	Type              string            `json:"type" bson:"type"`                                                // named-event | periodic | data-changed | data-added | data-modified | data-removed | data-accessed | data-access-ended | subscription-topic
	Name              *string           `json:"name,omitempty" bson:"name,omitempty"`                            // Name or URI that identifies the event
	Code              *CodeableConcept  `json:"code,omitempty" bson:"code,omitempty"`                            // Coded definition of the event
	SubscriptionTopic *string           `json:"subscriptionTopic,omitempty" bson:"subscription_topic,omitempty"` // What event
	TimingTiming      *Timing           `json:"timingTiming,omitempty" bson:"timing_timing,omitempty"`           // Timing of the event
	TimingDate        *string           `json:"timingDate,omitempty" bson:"timing_date,omitempty"`               // Timing of the event
	TimingDateTime    *string           `json:"timingDateTime,omitempty" bson:"timing_date_time,omitempty"`      // Timing of the event
	Data              []DataRequirement `json:"data,omitempty" bson:"data,omitempty"`                            // Triggering data of the event (multiple = 'and')
	Condition         *Expression       `json:"condition,omitempty" bson:"condition,omitempty"`                  // Whether the event triggers (boolean expression)
}

func (r *TriggerDefinition) Validate() error {
	var emptyString string
	if r.Type == emptyString {
		return fmt.Errorf("field 'Type' is required")
	}
	if r.Code != nil {
		if err := r.Code.Validate(); err != nil {
			return fmt.Errorf("Code: %w", err)
		}
	}
	if r.TimingTiming != nil {
		if err := r.TimingTiming.Validate(); err != nil {
			return fmt.Errorf("TimingTiming: %w", err)
		}
	}
	for i, item := range r.Data {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Data[%d]: %w", i, err)
		}
	}
	if r.Condition != nil {
		if err := r.Condition.Validate(); err != nil {
			return fmt.Errorf("Condition: %w", err)
		}
	}
	return nil
}
