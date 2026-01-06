package models

import (
	"fmt"
)

// A fixed quantity (no comparator)
type SimpleQuantity struct {
	Quantity []Quantity `json:"Quantity,omitempty" bson:"quantity,omitempty"` // A fixed quantity (no comparator)
}

func (r *SimpleQuantity) Validate() error {
	for i, item := range r.Quantity {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Quantity[%d]: %w", i, err)
		}
	}
	return nil
}
