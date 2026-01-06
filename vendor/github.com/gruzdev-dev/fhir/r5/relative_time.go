package models

import (
	"fmt"
)

// RelativeTime Type: RelativeTime expresses a time or time period as relative to the time of an event defined in data types other than dateTime.
type RelativeTime struct {
	Id                *string          `json:"id,omitempty" bson:"id,omitempty"`                                // Unique id for inter-element referencing
	ContextReference  *Reference       `json:"contextReference,omitempty" bson:"context_reference,omitempty"`   // The specific event occurrence or resource context used as a base point (reference point) in time
	ContextDefinition *string          `json:"contextDefinition,omitempty" bson:"context_definition,omitempty"` // The type of event used as a base point
	ContextPath       *string          `json:"contextPath,omitempty" bson:"context_path,omitempty"`             // Path to the element defining the basis for the relative time
	ContextCode       *CodeableConcept `json:"contextCode,omitempty" bson:"context_code,omitempty"`             // Coded representation of the event used as a base point (reference point) in time
	OffsetDuration    *Duration        `json:"offsetDuration,omitempty" bson:"offset_duration,omitempty"`       // An offset or offset range before (negative values) or after (positive values) the event
	OffsetRange       *Range           `json:"offsetRange,omitempty" bson:"offset_range,omitempty"`             // An offset or offset range before (negative values) or after (positive values) the event
	Text              *string          `json:"text,omitempty" bson:"text,omitempty"`                            // Free-text description
}

func (r *RelativeTime) Validate() error {
	if r.ContextReference != nil {
		if err := r.ContextReference.Validate(); err != nil {
			return fmt.Errorf("ContextReference: %w", err)
		}
	}
	if r.ContextCode != nil {
		if err := r.ContextCode.Validate(); err != nil {
			return fmt.Errorf("ContextCode: %w", err)
		}
	}
	if r.OffsetDuration != nil {
		if err := r.OffsetDuration.Validate(); err != nil {
			return fmt.Errorf("OffsetDuration: %w", err)
		}
	}
	if r.OffsetRange != nil {
		if err := r.OffsetRange.Validate(); err != nil {
			return fmt.Errorf("OffsetRange: %w", err)
		}
	}
	return nil
}
