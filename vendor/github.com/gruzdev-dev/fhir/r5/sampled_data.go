package models

import (
	"fmt"
)

// SampledData Type: A series of measurements taken by a device, with upper and lower limits. There may be more than one dimension in the data.
type SampledData struct {
	Id           *string   `json:"id,omitempty" bson:"id,omitempty"`                  // Unique id for inter-element referencing
	Origin       *Quantity `json:"origin" bson:"origin"`                              // Zero value and units
	Interval     *float64  `json:"interval,omitempty" bson:"interval,omitempty"`      // Number of intervalUnits between samples
	IntervalUnit string    `json:"intervalUnit" bson:"interval_unit"`                 // The measurement unit of the interval between samples
	Factor       *float64  `json:"factor,omitempty" bson:"factor,omitempty"`          // Multiply data by this before adding to origin
	LowerLimit   *float64  `json:"lowerLimit,omitempty" bson:"lower_limit,omitempty"` // Lower limit of detection
	UpperLimit   *float64  `json:"upperLimit,omitempty" bson:"upper_limit,omitempty"` // Upper limit of detection
	Dimensions   int       `json:"dimensions" bson:"dimensions"`                      // Number of sample points at each time point
	CodeMap      *string   `json:"codeMap,omitempty" bson:"code_map,omitempty"`       // Defines the codes used in the data
	Offsets      *string   `json:"offsets,omitempty" bson:"offsets,omitempty"`        // Offsets, typically in time, at which data values were taken
	Data         *string   `json:"data,omitempty" bson:"data,omitempty"`              // Decimal values with spaces, or "E" | "U" | "L", or another code
}

func (r *SampledData) Validate() error {
	if r.Origin == nil {
		return fmt.Errorf("field 'Origin' is required")
	}
	if r.Origin != nil {
		if err := r.Origin.Validate(); err != nil {
			return fmt.Errorf("Origin: %w", err)
		}
	}
	var emptyString string
	if r.IntervalUnit == emptyString {
		return fmt.Errorf("field 'IntervalUnit' is required")
	}
	if r.Dimensions == 0 {
		return fmt.Errorf("field 'Dimensions' is required")
	}
	return nil
}
