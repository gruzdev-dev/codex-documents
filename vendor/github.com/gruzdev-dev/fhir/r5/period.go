package models

// Period Type: A time period defined by a start and end date and optionally time.
type Period struct {
	Id    *string `json:"id,omitempty" bson:"id,omitempty"`       // Unique id for inter-element referencing
	Start *string `json:"start,omitempty" bson:"start,omitempty"` // Starting time with inclusive boundary
	End   *string `json:"end,omitempty" bson:"end,omitempty"`     // End time with inclusive boundary, if not ongoing
}

func (r *Period) Validate() error {
	return nil
}
