package models

import (
	"fmt"
)

// MonetaryComponent Type: Financial line items use this datatype to commonly categorize the value, and other factors that may effect how the value should be interpreted.
type MonetaryComponent struct {
	Id     *string          `json:"id,omitempty" bson:"id,omitempty"`         // Unique id for inter-element referencing
	Type   string           `json:"type" bson:"type"`                         // base | surcharge | discount | tax | informational
	Code   *CodeableConcept `json:"code,omitempty" bson:"code,omitempty"`     // Codes may be used to differentiate between kinds of taxes, surcharges, discounts etc.
	Factor *float64         `json:"factor,omitempty" bson:"factor,omitempty"` // Factor used for calculating this component
	Amount *Money           `json:"amount,omitempty" bson:"amount,omitempty"` // Explicit value amount to be used
}

func (r *MonetaryComponent) Validate() error {
	var emptyString string
	if r.Type == emptyString {
		return fmt.Errorf("field 'Type' is required")
	}
	if r.Code != nil {
		if err := r.Code.Validate(); err != nil {
			return fmt.Errorf("Code: %w", err)
		}
	}
	if r.Amount != nil {
		if err := r.Amount.Validate(); err != nil {
			return fmt.Errorf("Amount: %w", err)
		}
	}
	return nil
}
