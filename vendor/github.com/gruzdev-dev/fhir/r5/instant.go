package models

// instant Type: An instant in time - known at least to the second
type instant struct {
	Id    *string `json:"id,omitempty" bson:"id,omitempty"`       // xml:id (or equivalent in JSON)
	Value *string `json:"value,omitempty" bson:"value,omitempty"` // Primitive value for instant
}

func (r *instant) Validate() error {
	return nil
}
