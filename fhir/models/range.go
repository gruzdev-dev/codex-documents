package models

import (
	"fmt"
)

// Range Type: A set of ordered Quantities defined by a low and high limit.
type Range struct {
	Id   *string   `json:"id,omitempty" bson:"id,omitempty"`     // Unique id for inter-element referencing
	Low  *Quantity `json:"low,omitempty" bson:"low,omitempty"`   // Low limit
	High *Quantity `json:"high,omitempty" bson:"high,omitempty"` // High limit
}

func (r *Range) Validate() error {
	if r.Low != nil {
		if err := r.Low.Validate(); err != nil {
			return fmt.Errorf("Low: %w", err)
		}
	}
	if r.High != nil {
		if err := r.High.Validate(); err != nil {
			return fmt.Errorf("High: %w", err)
		}
	}
	return nil
}
