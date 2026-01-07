package models

// decimal Type: A rational number with implicit precision
type FHIRDecimal struct {
	Id    *string  `json:"id,omitempty" bson:"id,omitempty"`       // xml:id (or equivalent in JSON)
	Value *float64 `json:"value,omitempty" bson:"value,omitempty"` // Primitive value for decimal
}

func (r *FHIRDecimal) Validate() error {
	return nil
}
