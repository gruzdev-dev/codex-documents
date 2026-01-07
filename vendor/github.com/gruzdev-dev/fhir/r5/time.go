package models

// time Type: A time during the day, with no date specified
type time struct {
	Id    *string `json:"id,omitempty" bson:"id,omitempty"`       // xml:id (or equivalent in JSON)
	Value *string `json:"value,omitempty" bson:"value,omitempty"` // Primitive value for time
}

func (r *time) Validate() error {
	return nil
}
