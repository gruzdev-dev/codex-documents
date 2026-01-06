package models

import (
	"fmt"
)

// RatioRange Type: A range of ratios expressed as a low and high numerator and a denominator.
type RatioRange struct {
	Id            *string   `json:"id,omitempty" bson:"id,omitempty"`                        // Unique id for inter-element referencing
	LowNumerator  *Quantity `json:"lowNumerator,omitempty" bson:"low_numerator,omitempty"`   // Low Numerator limit
	HighNumerator *Quantity `json:"highNumerator,omitempty" bson:"high_numerator,omitempty"` // High Numerator limit
	Denominator   *Quantity `json:"denominator,omitempty" bson:"denominator,omitempty"`      // Denominator value
}

func (r *RatioRange) Validate() error {
	if r.LowNumerator != nil {
		if err := r.LowNumerator.Validate(); err != nil {
			return fmt.Errorf("LowNumerator: %w", err)
		}
	}
	if r.HighNumerator != nil {
		if err := r.HighNumerator.Validate(); err != nil {
			return fmt.Errorf("HighNumerator: %w", err)
		}
	}
	if r.Denominator != nil {
		if err := r.Denominator.Validate(); err != nil {
			return fmt.Errorf("Denominator: %w", err)
		}
	}
	return nil
}
