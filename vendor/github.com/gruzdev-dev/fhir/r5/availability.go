package models

import (
	"fmt"
)

// Availability Type: Availability data for an {item}, declaring what days/times are available, and any exceptions. The exceptions could be textual only, e.g. Public holidays, or could be time period specific and indicate a specific years dates.
type Availability struct {
	Id               *string                        `json:"id,omitempty" bson:"id,omitempty"`                               // Unique id for inter-element referencing
	Period           *Period                        `json:"period,omitempty" bson:"period,omitempty"`                       // When the availability applies
	AvailableTime    []AvailabilityAvailableTime    `json:"availableTime,omitempty" bson:"available_time,omitempty"`        // Times the {item} is available
	NotAvailableTime []AvailabilityNotAvailableTime `json:"notAvailableTime,omitempty" bson:"not_available_time,omitempty"` // Not available during this time due to provided reason
}

func (r *Availability) Validate() error {
	if r.Period != nil {
		if err := r.Period.Validate(); err != nil {
			return fmt.Errorf("Period: %w", err)
		}
	}
	for i, item := range r.AvailableTime {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("AvailableTime[%d]: %w", i, err)
		}
	}
	for i, item := range r.NotAvailableTime {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("NotAvailableTime[%d]: %w", i, err)
		}
	}
	return nil
}

type AvailabilityAvailableTime struct {
	Id                 *string  `json:"id,omitempty" bson:"id,omitempty"`                                   // Unique id for inter-element referencing
	DaysOfWeek         []string `json:"daysOfWeek,omitempty" bson:"days_of_week,omitempty"`                 // mon | tue | wed | thu | fri | sat | sun
	AllDay             *bool    `json:"allDay,omitempty" bson:"all_day,omitempty"`                          // Always available? i.e. 24 hour service
	AvailableStartTime *string  `json:"availableStartTime,omitempty" bson:"available_start_time,omitempty"` // Opening time of day (ignored if allDay = true)
	AvailableEndTime   *string  `json:"availableEndTime,omitempty" bson:"available_end_time,omitempty"`     // Closing time of day (ignored if allDay = true)
}

func (r *AvailabilityAvailableTime) Validate() error {
	return nil
}

type AvailabilityNotAvailableTime struct {
	Id          *string `json:"id,omitempty" bson:"id,omitempty"`                   // Unique id for inter-element referencing
	Description *string `json:"description,omitempty" bson:"description,omitempty"` // Reason presented to the user explaining why time not available
	During      *Period `json:"during,omitempty" bson:"during,omitempty"`           // Service not available during this period
}

func (r *AvailabilityNotAvailableTime) Validate() error {
	if r.During != nil {
		if err := r.During.Validate(); err != nil {
			return fmt.Errorf("During: %w", err)
		}
	}
	return nil
}
