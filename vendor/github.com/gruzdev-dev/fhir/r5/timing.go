package models

import (
	"fmt"
)

// Timing Type: Specifies an event that may occur multiple times. Timing schedules are used to record when things are planned, expected or requested to occur. The most common usage is in dosage instructions for medications. They are also used when planning care of various kinds, and may be used for reporting the schedule to which past regular activities were carried out.
type Timing struct {
	Id     *string          `json:"id,omitempty" bson:"id,omitempty"`         // Unique id for inter-element referencing
	Event  []string         `json:"event,omitempty" bson:"event,omitempty"`   // When the event occurs
	Repeat *TimingRepeat    `json:"repeat,omitempty" bson:"repeat,omitempty"` // When the event is to occur
	Code   *CodeableConcept `json:"code,omitempty" bson:"code,omitempty"`     // C | BID | TID | QID | AM | PM | QD | QOD | +
}

func (r *Timing) Validate() error {
	if r.Repeat != nil {
		if err := r.Repeat.Validate(); err != nil {
			return fmt.Errorf("Repeat: %w", err)
		}
	}
	if r.Code != nil {
		if err := r.Code.Validate(); err != nil {
			return fmt.Errorf("Code: %w", err)
		}
	}
	return nil
}

type TimingRepeat struct {
	Id             *string   `json:"id,omitempty" bson:"id,omitempty"`                          // Unique id for inter-element referencing
	BoundsDuration *Duration `json:"boundsDuration,omitempty" bson:"bounds_duration,omitempty"` // Length/Range of lengths, or (Start and/or end) limits
	BoundsRange    *Range    `json:"boundsRange,omitempty" bson:"bounds_range,omitempty"`       // Length/Range of lengths, or (Start and/or end) limits
	BoundsPeriod   *Period   `json:"boundsPeriod,omitempty" bson:"bounds_period,omitempty"`     // Length/Range of lengths, or (Start and/or end) limits
	Count          *int      `json:"count,omitempty" bson:"count,omitempty"`                    // Number of times to repeat
	CountMax       *int      `json:"countMax,omitempty" bson:"count_max,omitempty"`             // Maximum number of times to repeat
	Duration       *float64  `json:"duration,omitempty" bson:"duration,omitempty"`              // How long when it happens
	DurationMax    *float64  `json:"durationMax,omitempty" bson:"duration_max,omitempty"`       // How long when it happens (Max)
	DurationUnit   *string   `json:"durationUnit,omitempty" bson:"duration_unit,omitempty"`     // s | min | h | d | wk | mo | a - unit of time (UCUM)
	Frequency      *int      `json:"frequency,omitempty" bson:"frequency,omitempty"`            // Indicates the number of repetitions that should occur within a period. I.e. Event occurs frequency times per period
	FrequencyMax   *int      `json:"frequencyMax,omitempty" bson:"frequency_max,omitempty"`     // Event occurs up to frequencyMax times per period
	Period         *float64  `json:"period,omitempty" bson:"period,omitempty"`                  // The duration to which the frequency applies. I.e. Event occurs frequency times per period
	PeriodMax      *float64  `json:"periodMax,omitempty" bson:"period_max,omitempty"`           // Upper limit of period (3-4 hours)
	PeriodUnit     *string   `json:"periodUnit,omitempty" bson:"period_unit,omitempty"`         // s | min | h | d | wk | mo | a - unit of time (UCUM)
	StartOffset    *Quantity `json:"startOffset,omitempty" bson:"start_offset,omitempty"`       // Events within the repeat period do not start until startOffset has elapsed
	EndOffset      *Quantity `json:"endOffset,omitempty" bson:"end_offset,omitempty"`           // Events within the repeat period step once endOffset before the end of the period
	DayOfWeek      []string  `json:"dayOfWeek,omitempty" bson:"day_of_week,omitempty"`          // mon | tue | wed | thu | fri | sat | sun
	TimeOfDay      []string  `json:"timeOfDay,omitempty" bson:"time_of_day,omitempty"`          // Time of day for action
	When           []string  `json:"when,omitempty" bson:"when,omitempty"`                      // Code for time period of occurrence
	Offset         *int      `json:"offset,omitempty" bson:"offset,omitempty"`                  // Minutes from event (before or after)
}

func (r *TimingRepeat) Validate() error {
	if r.BoundsDuration != nil {
		if err := r.BoundsDuration.Validate(); err != nil {
			return fmt.Errorf("BoundsDuration: %w", err)
		}
	}
	if r.BoundsRange != nil {
		if err := r.BoundsRange.Validate(); err != nil {
			return fmt.Errorf("BoundsRange: %w", err)
		}
	}
	if r.BoundsPeriod != nil {
		if err := r.BoundsPeriod.Validate(); err != nil {
			return fmt.Errorf("BoundsPeriod: %w", err)
		}
	}
	if r.StartOffset != nil {
		if err := r.StartOffset.Validate(); err != nil {
			return fmt.Errorf("StartOffset: %w", err)
		}
	}
	if r.EndOffset != nil {
		if err := r.EndOffset.Validate(); err != nil {
			return fmt.Errorf("EndOffset: %w", err)
		}
	}
	return nil
}
