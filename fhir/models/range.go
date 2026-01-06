package models

// Range Type: A set of ordered Quantities defined by a low and high limit.
type Range struct {
	Id   *string   `json:"id,omitempty" bson:"id,omitempty"`     // Unique id for inter-element referencing
	Low  *Quantity `json:"low,omitempty" bson:"low,omitempty"`   // Low limit
	High *Quantity `json:"high,omitempty" bson:"high,omitempty"` // High limit
}
