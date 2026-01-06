package models

import (
	"fmt"
)

// An amount of money. With regard to precision, see [Decimal Precision](datatypes.html#precision)
type MoneyQuantity struct {
	Quantity []Quantity `json:"Quantity,omitempty" bson:"quantity,omitempty"` // An amount of money. With regard to precision, see [Decimal Precision](datatypes.html#precision)
}

func (r *MoneyQuantity) Validate() error {
	for i, item := range r.Quantity {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Quantity[%d]: %w", i, err)
		}
	}
	return nil
}
