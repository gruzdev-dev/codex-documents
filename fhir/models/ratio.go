package models

// Ratio Type: A relationship of two Quantity values - expressed as a numerator and a denominator.
type Ratio struct {
	Id          *string   `json:"id,omitempty" bson:"id,omitempty"`                   // Unique id for inter-element referencing
	Numerator   *Quantity `json:"numerator,omitempty" bson:"numerator,omitempty"`     // Numerator value
	Denominator *Quantity `json:"denominator,omitempty" bson:"denominator,omitempty"` // Denominator value
}
