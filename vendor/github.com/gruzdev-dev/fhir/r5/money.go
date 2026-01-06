package models

// Money Type: An amount of economic utility in some recognized currency.
type Money struct {
	Id       *string  `json:"id,omitempty" bson:"id,omitempty"`             // Unique id for inter-element referencing
	Value    *float64 `json:"value,omitempty" bson:"value,omitempty"`       // Numerical value (with implicit precision)
	Currency *string  `json:"currency,omitempty" bson:"currency,omitempty"` // ISO 4217 Currency Code
}

func (r *Money) Validate() error {
	return nil
}
