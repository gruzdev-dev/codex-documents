package models

import (
	"fmt"
)

// Ratio Type: A relationship of two Quantity values - expressed as a numerator and a denominator.
type Ratio struct {
	Id          *string   `json:"id,omitempty" bson:"id,omitempty"`                   // Unique id for inter-element referencing
	Numerator   *Quantity `json:"numerator,omitempty" bson:"numerator,omitempty"`     // Numerator value
	Denominator *Quantity `json:"denominator,omitempty" bson:"denominator,omitempty"` // Denominator value
}

func (r *Ratio) Validate() error {
	if r.Numerator != nil {
		if err := r.Numerator.Validate(); err != nil {
			return fmt.Errorf("Numerator: %w", err)
		}
	}
	if r.Denominator != nil {
		if err := r.Denominator.Validate(); err != nil {
			return fmt.Errorf("Denominator: %w", err)
		}
	}
	return nil
}
