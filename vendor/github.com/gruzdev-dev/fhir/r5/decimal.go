package models

// decimal Type: A rational number with implicit precision
type decimal struct {
}

func (r *decimal) Validate() error {
	return nil
}

type Decimal struct {
	Id    *string  `json:"id,omitempty" bson:"id,omitempty"`       // xml:id (or equivalent in JSON)
	Value *float64 `json:"value,omitempty" bson:"value,omitempty"` // Primitive value for decimal
}

func (r *Decimal) Validate() error {
	return nil
}
