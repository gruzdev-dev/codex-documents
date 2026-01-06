package models

// Duration Type: A length of time.
type Duration struct {
	Id         *string  `json:"id,omitempty" bson:"id,omitempty"`                 // Unique id for inter-element referencing
	Value      *float64 `json:"value,omitempty" bson:"value,omitempty"`           // Numerical value (with implicit precision)
	Comparator *string  `json:"comparator,omitempty" bson:"comparator,omitempty"` // < | <= | >= | > | ad - how to understand the value
	Unit       *string  `json:"unit,omitempty" bson:"unit,omitempty"`             // Unit representation
	System     *string  `json:"system,omitempty" bson:"system,omitempty"`         // System that defines coded unit form
	Code       *string  `json:"code,omitempty" bson:"code,omitempty"`             // Coded form of the unit
}

func (r *Duration) Validate() error {
	return nil
}
