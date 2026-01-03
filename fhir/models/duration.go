package fhir

// Duration Type: A length of time.
type Duration struct {
	Id         *string  `json:"id,omitempty"`         // Unique id for inter-element referencing
	Value      *float64 `json:"value,omitempty"`      // Numerical value (with implicit precision)
	Comparator *string  `json:"comparator,omitempty"` // < | <= | >= | > | ad - how to understand the value
	Unit       *string  `json:"unit,omitempty"`       // Unit representation
	System     *string  `json:"system,omitempty"`     // System that defines coded unit form
	Code       *string  `json:"code,omitempty"`       // Coded form of the unit
}
