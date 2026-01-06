package models

import (
	"fmt"
)

// DosageSafety Type: Safety Details about the usage of the medication.
type DosageSafety struct {
	Id         *string                 `json:"id,omitempty" bson:"id,omitempty"`                  // Unique id for inter-element referencing
	DoseLimit  []DosageSafetyDoseLimit `json:"doseLimit,omitempty" bson:"dose_limit,omitempty"`   // A dose limit for safe use of the medication
	IfExceeded *string                 `json:"ifExceeded,omitempty" bson:"if_exceeded,omitempty"` // What to do if the instructions lead to exceeding the dose limits
}

func (r *DosageSafety) Validate() error {
	for i, item := range r.DoseLimit {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("DoseLimit[%d]: %w", i, err)
		}
	}
	return nil
}

type DosageSafetyDoseLimit struct {
	Id              *string     `json:"id,omitempty" bson:"id,omitempty"`         // Unique id for inter-element referencing
	ValueInteger    *int        `json:"valueInteger" bson:"value_integer"`        // Quantity that is safe to use
	ValueQuantity   *Quantity   `json:"valueQuantity" bson:"value_quantity"`      // Quantity that is safe to use
	ValueExpression *Expression `json:"valueExpression" bson:"value_expression"`  // Quantity that is safe to use
	Scope           string      `json:"scope" bson:"scope"`                       // dosage | period | administration | lifetime - The scope of the dose limitation
	Period          *Duration   `json:"period,omitempty" bson:"period,omitempty"` // The period over which the quantity is safe to use (if scope = period)
	Text            *string     `json:"text,omitempty" bson:"text,omitempty"`     // Additional notes about the dose limit
}

func (r *DosageSafetyDoseLimit) Validate() error {
	if r.ValueInteger == nil {
		return fmt.Errorf("field 'ValueInteger' is required")
	}
	if r.ValueQuantity == nil {
		return fmt.Errorf("field 'ValueQuantity' is required")
	}
	if r.ValueQuantity != nil {
		if err := r.ValueQuantity.Validate(); err != nil {
			return fmt.Errorf("ValueQuantity: %w", err)
		}
	}
	if r.ValueExpression == nil {
		return fmt.Errorf("field 'ValueExpression' is required")
	}
	if r.ValueExpression != nil {
		if err := r.ValueExpression.Validate(); err != nil {
			return fmt.Errorf("ValueExpression: %w", err)
		}
	}
	var emptyString string
	if r.Scope == emptyString {
		return fmt.Errorf("field 'Scope' is required")
	}
	if r.Period != nil {
		if err := r.Period.Validate(); err != nil {
			return fmt.Errorf("Period: %w", err)
		}
	}
	return nil
}
